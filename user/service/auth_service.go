package service

import "github.com/proyectos01-a/user/dto/req"

type AuthService interface {
	UserLogin(req *req.LoginRequest) (string, error)
}
