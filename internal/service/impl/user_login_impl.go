package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nhh57/go-ecommerce-backend-api/global"
	consts "github.com/nhh57/go-ecommerce-backend-api/internal/const"
	"github.com/nhh57/go-ecommerce-backend-api/internal/database"
	"github.com/nhh57/go-ecommerce-backend-api/internal/model"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/random"
	"github.com/nhh57/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/nhh57/go-ecommerce-backend-api/pkg/response"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"strings"
	"time"
)

type sUserLogin struct {
	// Implement the IUserLogin interface here.
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	// Implement login logic here.
	return nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// Implement login logic here.
	// 1 .hash email
	fmt.Printf("VerifyKey %s\n", in.VerifyKey)
	fmt.Printf("VeirifyType %d\n", in.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey:: %s\n", hashKey)
	// 2. check user exists in user base
	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}
	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}
	// 3. Create OTP
	userKey := utils.GetUserKey(hashKey)

	otpFound, err := global.Rdb.Get(ctx, userKey).Result()
	// util...
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed ::", err)
		return response.ErrInvalidOTP, err
	case otpFound != "":
		return response.ErrCodeUserHasExists, fmt.Errorf("")
	}

	// 4. Generate OTP

	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("OTP is ::%d\n", otpNew)
	fmt.Printf("Register:: userKey::%s\n", userKey)
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalidOTP, err
	}
	//return response.ErrCodeSuccess, nil
	// 6. Send OTP
	fmt.Printf("in.VerifyType::%s", in.VerifyType)
	switch in.VerifyType {

	case consts.EMAIL:
		err = sendto.SentTextEmailOtp([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		// 7. Save OTP to MYSQL
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
		})

		if err != nil {
			return response.ErrSendEmailOTP, err
		}

		// 8. getlastID
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		fmt.Println("lastIdVerifyUser::%s", lastIdVerifyUser)
		log.Println("lastIdVerifyUser", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil
	case consts.MOBILE:
		return response.ErrCodeSuccess, nil
	}
	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	// login

	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	fmt.Printf("hashKey::%s", hashKey)
	// get OTP
	otpFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}
	fmt.Printf("otpFound::%s", otpFound)
	if in.VerifyCode != otpFound {
		// Neu nhu sai 3 lan trong 1 phut thi xoa
		return out, fmt.Errorf("OTP not match")
	}
	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, err
	}
	// update status verified
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)

	if err != nil {
		return out, err
	}
	// output
	// tot nhat nen tao key_secret
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"

	return out, err
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	// Implement login logic here.
	return nil
}
