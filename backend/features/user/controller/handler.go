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
		err := c.ShouldBindJSON(&request)
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
				message = "Bad request, duplicate input"
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
		c.JSON(http.StatusCreated, response.ResponseFormat(http.StatusCreated, "Successfully operation", CoreToResponse(result), nil))
	}
}

// Login implements user.Controller.
func (*Controller) Login() gin.HandlerFunc {
	panic("unimplemented")
}
