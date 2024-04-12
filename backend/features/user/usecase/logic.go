package usecase

import (
	"errors"
	"strings"

	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/dimasyudhana/simple-chat/features/user"
)

var log = middlewares.Log()

type Service struct {
	query user.Repository
}

func New(ud user.Repository) user.UseCase {
	return &Service{
		query: ud,
	}
}

// Register implements user.UseCase.
func (us *Service) Register(request user.UserCore) (user.UserCore, error) {
	if request.Username == "" || request.Email == "" || request.Password == "" {
		log.Error("request cannot be empty")
		return user.UserCore{}, errors.New("request cannot be empty")
	}

	result, err := us.query.Register(request)
	if err != nil {
		message := ""
		switch {
		case strings.Contains(err.Error(), "error while creating id for user"):
			log.Error("error while creating id for user")
			message = "error while creating id for user"
		case strings.Contains(err.Error(), "error while hashing password"):
			log.Error("error while hashing password")
			message = "error while hashing password"
		case strings.Contains(err.Error(), "error insert data, duplicate input"):
			log.Error("error insert data, duplicate input")
			message = "error insert data, duplicate input"
		case strings.Contains(err.Error(), "no row affected"):
			log.Error("no row affected")
			message = "no row affected"
		default:
			log.Error("internal server error")
			message = "internal server error"
		}
		return user.UserCore{}, errors.New(message)
	}
	return result, nil
}

// Login implements user.UseCase.
func (us *Service) Login(request user.UserCore) (user.UserCore, string, error) {
	if request.Username == "" || request.Password == "" {
		log.Error("username and password cannot be empty")
		return user.UserCore{}, "", errors.New("username and password cannot be empty")
	}

	result, token, err := us.query.Login(request)
	if err != nil {
		message := ""
		if strings.Contains(err.Error(), "invalid username and password") {
			log.Error("invalid username and password")
			message = "invalid username and password"
		} else if strings.Contains(err.Error(), "password does not match") {
			log.Error("password does not match")
			message = "password does not match"
		} else if strings.Contains(err.Error(), "error while creating jwt token") {
			log.Error("error while creating jwt token")
			message = "error while creating jwt token"
		} else {
			log.Error("internal server error")
			message = "internal server error"
		}
		return user.UserCore{}, "", errors.New(message)
	}

	return result, token, nil
}
