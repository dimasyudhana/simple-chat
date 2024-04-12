package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestResponse struct {
	Code       int         `json:"code,omitempty"`
	Message    string      `json:"message,omitempty"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func ResponseFormat(code int, message string, data interface{}, pagination interface{}) RequestResponse {
	result := RequestResponse{
		Code:       code,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}

	return result
}

func BadRequestError(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ResponseFormat(http.StatusBadRequest, message, nil, nil))
}

func NotFoundError(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ResponseFormat(http.StatusNotFound, message, nil, nil))
}

func UnauthorizedError(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, ResponseFormat(http.StatusUnauthorized, message, nil, nil))
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, ResponseFormat(http.StatusInternalServerError, message, nil, nil))
}
