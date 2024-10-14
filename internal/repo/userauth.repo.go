package repo

import (
	"fmt"
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"time"
)

type IUserAuthenRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthenRepository struct{}

func (u userAuthenRepository) AddOTP(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr%d:otp", email) //usr:email:otp
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthenRepository() IUserAuthenRepository {
	return &userAuthenRepository{}
}
