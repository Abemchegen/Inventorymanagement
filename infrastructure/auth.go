package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware that checks for the presence of a valid JWT token
// in the Authorization header of incoming requests.
type AuthMiddleware struct {
	jwtService *JWTService
}

// NewAuthMiddleware creates a new instance of AuthMiddleware with the provided JWTService.
func NewAuthMiddleware(jwtService *JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

// Middleware function to check if the request is authenticated by validating the JWT token.
func (a *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the "Authorization" header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort() // Stop further processing of the request
			return
		}

		// The token is usually passed with "Bearer <token>"
		if len(tokenString) > 7 && strings.ToLower(tokenString[:7]) == "bearer " {
			tokenString = tokenString[7:] // Remove the "Bearer " part
		}

		// Validate the token
		claims, err := a.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort() // Stop further processing of the request
			return
		}

		userID, userIDExists := claims["userid"]
		email, emailExists := claims["email"]
		role, roleExists := claims["role"]
		if !userIDExists || !emailExists || !roleExists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token does not contain necessary claims"})
			c.Abort()
			return
		}

		c.Set("userid", userID)
		c.Set("email", email)
		c.Set("role", role)
		c.Next()
	}
}
