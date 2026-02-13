package controllers

import (
	"context"
	"net/http"

	"property-backend/models"
	"property-backend/services"
	"property-backend/utils"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication endpoints
type AuthController struct {
	svc services.AuthService
}

// NewAuthController creates a new AuthController
func NewAuthController(svc services.AuthService) *AuthController {
	return &AuthController{svc: svc}
}

// SignUp godoc
// @Summary Sign up a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/signup [post]
func (a *AuthController) SignUp(c *gin.Context) {
	var req struct {
		FirstName   string `json:"first_name" binding:"required"`
		LastName    string `json:"last_name" binding:"required"`
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,min=8"`
		PhoneNumber string `json:"phone_number"`
		RoleID      uint   `json:"role_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password"})
		return
	}
	user := models.User{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		HashedPassword: hashed,
		PhoneNumber:    req.PhoneNumber,
		RoleID:         req.RoleID,
	}
	authResp, err := a.svc.SignUp(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, authResp)
}

// SignIn godoc
// @Summary Sign in a user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/v1/auth/signin [post]
func (a *AuthController) SignIn(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authResp, err := a.svc.SignIn(context.Background(), req.Email, req.Password)
	if err != nil {
		// treat invalid credentials distinctly
		if err == services.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authResp)
}

// ForgotPassword godoc
// @Summary Request a password reset
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/forgot-password [post]
func (a *AuthController) ForgotPassword(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.svc.ForgotPassword(context.Background(), req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to your email"})
}

// ResetPassword godoc
// @Summary Reset password with token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/v1/auth/reset-password [post]
func (a *AuthController) ResetPassword(c *gin.Context) {
	var req struct {
		Token       string `json:"token" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.svc.ResetPassword(context.Background(), req.Token, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
