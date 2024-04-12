package controller

import (
	"context"
	"net/http"
	"strings"

	"github.com/dimasyudhana/simple-chat/app/middlewares"
	"github.com/dimasyudhana/simple-chat/features/user"
	"github.com/dimasyudhana/simple-chat/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/mold/v4/modifiers"
)

var log = middlewares.Log()

type Controller struct {
	service user.UseCase
}

func New(us user.UseCase) user.Controller {
	return &Controller{
		service: us,
	}
}

// Register implements user.Controller.
func (uc *Controller) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := RegisterRequest{}
		conform := modifiers.New()
		err := c.ShouldBind(&request)
		if err != nil {
			log.Error("error on bind input")
			response.BadRequestError(c, "Bad request")
			return
		}

		err = conform.Struct(context.Background(), &request)
		result, err := uc.service.Register(RequestToCore(request))
		if err != nil {
			message := ""
			switch {
			case strings.Contains(err.Error(), "request cannot be empty"):
				log.Error("request cannot be empty")
				message = "Bad request, request cannot be empty"
				response.BadRequestError(c, message)
			case strings.Contains(err.Error(), "error insert data, duplicate input"):
				log.Error("error insert data, duplicate input")
				message = "Bad request, duplicate input"
				response.BadRequestError(c, message)
			case strings.Contains(err.Error(), "no row affected"):
				log.Error("no row affected")
				message = "Bad request, duplicate entry"
				response.BadRequestError(c, message)
			case strings.Contains(err.Error(), "error while creating id for user"):
				log.Error("error while creating id for user")
				message = "Internal server error"
				response.InternalServerError(c, message)
			case strings.Contains(err.Error(), "error while hashing password"):
				log.Error("error while hashing password")
				message = "Internal server error"
				response.InternalServerError(c, message)
			default:
				log.Error("internal server error")
				message = "Internal server error"
				response.InternalServerError(c, message)
			}
			return
		}
		c.JSON(http.StatusCreated, response.ResponseFormat(http.StatusCreated, "Successfully operation", Register(result), nil))
	}
}

// Login implements user.Controller.
func (uc *Controller) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := LoginRequest{}
		conform := modifiers.New()
		err := c.ShouldBind(&request)
		if err != nil {
			log.Error("error on bind input")
			response.BadRequestError(c, "Bad request")
			return
		}

		err = conform.Struct(context.Background(), &request)
		resp, token, err := uc.service.Login(RequestToCore(request))
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "invalid email format"):
				log.Error("bad request, invalid email format")
				response.BadRequestError(c, "Bad request, invalid email format")
			case strings.Contains(err.Error(), "password cannot be empty"):
				log.Error("bad request, password cannot be empty")
				response.BadRequestError(c, "Bad request, password cannot be empty")
			case strings.Contains(err.Error(), "invalid email and password"):
				log.Error("bad request, invalid email and password")
				response.BadRequestError(c, "Bad request, invalid email and password")
			case strings.Contains(err.Error(), "password does not match"):
				log.Error("bad request, password does not match")
				response.BadRequestError(c, "Bad request, password does not match")
			case strings.Contains(err.Error(), "no row affected"):
				log.Error("no row affected")
				response.NotFoundError(c, "The requested resource was not found")
			case strings.Contains(err.Error(), "error while creating jwt token"):
				log.Error("internal server error, error while creating jwt token")
				response.InternalServerError(c, "Internal server error")
			default:
				log.Error("internal server error")
				response.InternalServerError(c, "Internal server error")

			}
			return
		}
		c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, response.ResponseFormat(http.StatusOK, "Successful login", Login(resp), nil))
	}
}

// Logout implements user.Controller.
func (*Controller) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("jwt", "", -1, "", "", false, true)
		c.JSON(http.StatusOK, response.ResponseFormat(http.StatusOK, "Successful logout", nil, nil))
	}
}
