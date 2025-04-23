package middleware

import (
	"gametify/repositories"
	"gametify/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userId, err := utils.ValidateToken(tokenString)
		if err != nil || userId == 0 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		user, err := userRepo.FindByID(userId)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("user_id", userId)
		c.Set("user_role", user.Role)
		c.Next()
	}
}
