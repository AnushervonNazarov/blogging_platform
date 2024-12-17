package service

import (
	"blogging_platform/errs"
	"blogging_platform/internal/models"
	repository "blogging_platform/internal/repositories"
	"blogging_platform/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		// Role:     role,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func SignIn(username, password string) (accessToken string, err error) {
	password = utils.GenerateHash(password)
	user, err := repository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return "", errs.ErrIncorrectUsernameOrPassword
		}
		return "", err
	}

	accessToken, err = GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
