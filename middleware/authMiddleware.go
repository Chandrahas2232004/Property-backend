package middleware

import (
	"net/http"
	"strings"

	"property-backend/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token from Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Store claims in context for use in handlers
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role_id", claims.RoleID)
		c.Set("user_claims", claims)

		c.Next()
	}
}

// GetUserIDFromContext retrieves the user ID from the request context
func GetUserIDFromContext(c *gin.Context) uint {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(uint)
}

// GetUserClaimsFromContext retrieves the JWT claims from the request context
func GetUserClaimsFromContext(c *gin.Context) *utils.JWTClaims {
	claims, exists := c.Get("user_claims")
	if !exists {
		return nil
	}
	return claims.(*utils.JWTClaims)
}
