package models

import (
	"errors"
	"github.com/MicBun/go-job-list-api/utils/jwtAuth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
}

func (u *User) GetJWTToken(ctx *gin.Context) (string, error) {
	db := ctx.MustGet("db").(*gorm.DB)
	var user User
	err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	if err != nil {
		return "", errors.New("user not found")
	}
	token, _ := jwtAuth.GenerateToken()
	return token, nil
}
