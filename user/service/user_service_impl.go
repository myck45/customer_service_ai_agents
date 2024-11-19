package service

import (
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/models"
	"github.com/proyectos01-a/shared/utils"
	"github.com/proyectos01-a/user/auth"
	"github.com/proyectos01-a/user/dto/req"
	"github.com/proyectos01-a/user/dto/res"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	userRepo   data.UserRepository
	bcryptUtil auth.Bcrypt
	utils      utils.Utils
}

// CreateUser implements UserService.
func (u *UserServiceImpl) CreateUser(user *req.CreateUserReq) (*res.UserResponse, error) {

	birthDate, err := u.utils.ParseStringToDateTime(user.BirthDate)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error parsing birth date")
		return nil, err
	}

	hashedPassword, err := u.bcryptUtil.HashPassword(user.Password)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error hashing password")
		return nil, err
	}

	userToSave := &models.User{
		Name:      user.Name,
		LastName:  user.LastName,
		BirthDate: *birthDate,
		UserEmail: user.UserEmail,
		Password:  hashedPassword,
		PhoneNum:  user.PhoneNum,
		Role:      user.Role,
	}

	userToSave, err = u.userRepo.SaveUser(userToSave)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error saving user")
		return nil, err
	}

	return &res.UserResponse{
		ID:   userToSave.ID,
		Name: userToSave.Name,
	}, nil
}

// DeleteUser implements UserService.
func (u *UserServiceImpl) DeleteUser(id uint) error {

	err := u.userRepo.DeleteUser(id)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error deleting user")
		return err
	}

	return nil
}

// GetAllUsers implements UserService.
func (u *UserServiceImpl) GetAllUsers() ([]res.UserResponse, error) {

	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error fetching users")
		return nil, err
	}

	userList := make([]res.UserResponse, 0, len(users))
	for _, user := range users {
		userList = append(userList, res.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			LastName:  user.LastName,
			BirthDate: user.BirthDate.Format("2006-01-02"),
			UserEmail: user.UserEmail,
			PhoneNum:  user.PhoneNum,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return userList, nil
}

// GetUserByEmail implements UserService.
func (u *UserServiceImpl) GetUserByEmail(email string) (*res.UserResponse, error) {

	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error fetching user")
		return nil, err
	}

	return &res.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		BirthDate: user.BirthDate.Format("2006-01-02"),
		UserEmail: user.UserEmail,
		PhoneNum:  user.PhoneNum,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// GetUserByID implements UserService.
func (u *UserServiceImpl) GetUserByID(id uint) (*res.UserResponse, error) {

	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error fetching user")
		return nil, err
	}

	return &res.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		BirthDate: user.BirthDate.Format("2006-01-02"),
		UserEmail: user.UserEmail,
		PhoneNum:  user.PhoneNum,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

// UpdateUser implements UserService.
func (u *UserServiceImpl) UpdateUser(id uint, user *req.UpdateUserReq) error {

	birthDate, err := u.utils.ParseStringToDateTime(user.BirthDate)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error parsing birth date")
		return err
	}

	hashedPassword, err := u.bcryptUtil.HashPassword(user.Password)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error hashing password")
		return err
	}

	userToUpdate := &models.User{
		Name:      user.Name,
		LastName:  user.LastName,
		BirthDate: *birthDate,
		UserEmail: user.UserEmail,
		Password:  hashedPassword,
		PhoneNum:  user.PhoneNum,
	}

	err = u.userRepo.UpdateUser(userToUpdate)
	if err != nil {
		logrus.WithError(err).Error("*** [UserServiceImpl] Error updating user")
		return err
	}

	return nil
}

func NewUserServiceImpl(userRepo data.UserRepository, bcryptUtil auth.Bcrypt, utils utils.Utils) UserService {
	return &UserServiceImpl{
		userRepo:   userRepo,
		bcryptUtil: bcryptUtil,
		utils:      utils,
	}
}
