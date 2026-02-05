package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"property-backend/models"
	"property-backend/services"
	"property-backend/utils"
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
	id, err := a.svc.SignUp(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
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
	user, err := a.svc.SignIn(context.Background(), req.Email, req.Password)
	if err != nil {
		// treat invalid credentials distinctly
		if err == services.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
