package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"github.com/nhh57/go-ecommerce-backend-api/internal/repo"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/random"
	"github.com/nhh57/go-ecommerce-backend-api/pkg/response"
	"github.com/segmentio/kafka-go"
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
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}
	// 4. send Email OTP
	//err = sendto.SendTemplateEmailOtp([]string{email}, "datnvpk02264@fpt.edu.vn", "otp-auth.html",
	//	map[string]interface{}{"otp": strconv.Itoa(otp)})
	//if err != nil {
	//	return response.ErrSendEmailOTP
	//}
	// send otp by java
	//err = sendto.SendEmailToJavaByAPI(strconv.Itoa(otp), email, "otp-auth.html")
	//fmt.Println("err sendto:Java::%d\n", err)
	//if err != nil {
	//	return response.ErrSendEmailOTP
	//}

	// send otp via Kafka Java
	body := make(map[string]interface{})
	body["otp"] = otp
	body["email"] = email
	bodyRequest, _ := json.Marshal(body)
	message := kafka.Message{
		Key:   []byte("otp-auth"),
		Value: []byte(bodyRequest),
		Time:  time.Now(),
	}
	err = global.KafkaProducer.WriteMessages(context.Background(), message)
	if err != nil {
		fmt.Println("err send to Kafka::%d\n", err)
		return response.ErrSendEmailOTP
	}
	return response.ErrCodeSuccess
}
