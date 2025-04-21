package middleware

import (
	"net/http"
	"strings"
	"github.com/breezjirasak/triptales/internal/auth"
	"github.com/gin-gonic/gin"
)

// JWTMiddleware validates JWT tokens from the Authorization header
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}
		
		// Check the format of the header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}
		
		// Extract the token
		tokenString := parts[1]
		
		// Validate the token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}
		
		// Set user ID and username in context for handlers to use
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		
		c.Next()
	}
}

// OptionalJWTMiddleware attempts to validate a JWT token but doesn't require one
func OptionalJWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		
		// If no header is present, continue without setting user context
		if authHeader == "" {
			c.Next()
			return
		}
		
		// Check the format of the header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			// Invalid format but we don't abort since this is optional
			c.Next()
			return
		}
		
		// Extract the token
		tokenString := parts[1]
		
		// Validate the token
		claims, err := auth.ValidateToken(tokenString)
		if err == nil {
			// Set user ID and username in context for handlers to use
			c.Set("userID", claims.UserID)
			c.Set("username", claims.Username)
		}
		
		c.Next()
	}
}

// AdminMiddleware checks if the user has admin privileges
// This is an example - you'll need to implement admin checking based on your requirements
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// First validate JWT token
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		
		// This is a placeholder - implement your own admin checking logic
		// For example, you might check a user's role in the database
		isAdmin := checkIfUserIsAdmin(userID.(string))
		
		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin privileges required"})
			c.Abort()
			return
		}
		
		c.Next()
	}
}

// Placeholder function - implement this based on your requirements
func checkIfUserIsAdmin(userID string) bool {
	// Query your database to check if the user has admin role
	// This is just a placeholder
	return false
}