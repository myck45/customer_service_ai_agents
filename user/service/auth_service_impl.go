package service

import (
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/user/auth"
	"github.com/proyectos01-a/user/dto/req"
	"github.com/sirupsen/logrus"
)

type AuthServiceImpl struct {
	auth       auth.Auth
	bcryptUtil auth.Bcrypt
	userRepo   data.UserRepository
}

// UserLogin implements AuthService.
func (a *AuthServiceImpl) UserLogin(req *req.LoginRequest) (string, error) {

	user, err := a.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		logrus.WithError(err).Error("*** [AuthServiceImpl] Error getting user by email")
		return "", err
	}

	err = a.bcryptUtil.ComparePassword(user.Password, req.Password)
	if err != nil {
		logrus.WithError(err).Error("*** [AuthServiceImpl] Error comparing password")
		return "", err
	}

	token, err := a.auth.GenerateToken(user.ID, user.UserEmail, user.Role)
	if err != nil {
		logrus.WithError(err).Error("*** [AuthServiceImpl] Error generating token")
		return "", err
	}

	return token, nil
}

func NewAuthServiceImpl(auth auth.Auth, bcryptUtil auth.Bcrypt, userRepo data.UserRepository) AuthService {
	return &AuthServiceImpl{
		auth:       auth,
		bcryptUtil: bcryptUtil,
		userRepo:   userRepo,
	}
}
