package service

import (
	"fmt"
	"github.com/nhh57/go-ecommerce-backend-api/internal/repo"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/random"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/nhh57/go-ecommerce-backend-api/pkg/response"
	"strconv"
	"time"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthenRepository
}

func NewUserService(
	userRepo repo.IUserRepository,
	userAuthRepo repo.IUserAuthenRepository,
) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("hashEmail::%s", hashEmail)

	// 5. check otp is available

	// 6. user spam

	// 1. check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	// 2. new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("OTP is ::%d\n", otp)
	// 3. save OTP in redis with exp time
	err := us.userAuthRepo.AddOTP(email, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. send Email OTP
	err = sendto.SentTextEmailOtp([]string{email}, "nguyenhoanghai0507@gmail.com", strconv.Itoa(otp))
	if err != nil {
		return response.ErrSendEmailOTP
	}

	return response.ErrCodeSuccess
}
