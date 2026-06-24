package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/dto"
	"github.com/kyawzinkoko-dev/online-learning-platform/internal/modules/auth/service"
	"github.com/kyawzinkoko-dev/online-learning-platform/pkg/response"
)

type AuthHandler struct {
	authService *service.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Malformed JSON request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	if err := h.authService.Register(req); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Registration processed successfully", nil)

}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Malformed request body")
		return
	}

	if err := h.validate.Struct(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	token, err := h.authService.Login(req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	response.Success(c, http.StatusOK, "Authentication successful", gin.H{
		"token":      token,
		"token_type": "Bearer",
	})
}
