package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/dto"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": formatValidationErrors(ve)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload format"})
		return
	}

	if err := h.authService.Register(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}

func formatValidationErrors(ve validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range ve {
		errs[f.Field()] = f.Tag()
	}
	return errs
}
