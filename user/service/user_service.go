package service

import (
	"github.com/proyectos01-a/user/dto/req"
	"github.com/proyectos01-a/user/dto/res"
)

type UserService interface {
	CreateUser(user *req.CreateUserReq) (*res.UserResponse, error)
	GetUserByID(id uint) (*res.UserResponse, error)
	GetUserByEmail(email string) (*res.UserResponse, error)
	GetAllUsers() ([]res.UserResponse, error)
	UpdateUser(id uint, user *req.UpdateUserReq) error
	DeleteUser(id uint) error
}
