package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type JSONResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ValidationErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func Success(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, JSONResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string) {
	c.JSON(status, JSONResponse{
		Success: false,
		Message: message,
	})
}

func ValidationError(c *gin.Context, err error) {
	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusBadRequest, JSONResponse{
			Success: false,
			Message: "Invalid request payload format",
		})
		return
	}

	errMap := make(map[string]string)

	for _, f := range ve {
		field := strings.ToLower(f.Field())

		switch f.Tag() {
		case "required":
			errMap[field] = field + " is required"
		case "email":
			errMap[field] = "Must be a valid email address"
		case "min":
			errMap[field] = "Must be at least " + f.Param() + " characters long"
		default:
			errMap[field] = "Validation failed on rule: " + f.Tag()
		}
	}

	c.JSON(http.StatusUnprocessableEntity, ValidationErrorResponse{
		Success: false,
		Message: "Validation fail",
		Errors:  errMap,
	})

}
