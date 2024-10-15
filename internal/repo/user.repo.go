package repo

import (
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"github.com/nhh57/go-ecommerce-backend-api/internal/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
}
type userRepository struct {
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
func (u userRepository) GetUserByEmail(email string) bool {
	// SELECT * FROM user WHERE email = '??' ORDER BY email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("user_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	return row != NumberNull
}
