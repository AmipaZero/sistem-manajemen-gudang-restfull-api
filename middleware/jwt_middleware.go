package middleware

import (
	"net/http"
	"sistem-manajemen-gudang/config"
	"sistem-manajemen-gudang/model/domain"
	"sistem-manajemen-gudang/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		 if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
            return
        }
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token tidak sesuai"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := util.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		mapClaims := claims.(jwt.MapClaims)
		userID := uint(mapClaims["user_id"].(float64))

		var user domain.User
		if err := config.DB.First(&user, userID).Error; err != nil || user.Token == nil || *user.Token != tokenString {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		c.Set("userID", userID)
		c.Set("role", user.Role)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {

		
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "akses khusus admin"})
			return
		}
		//cast domain
		userRole, ok := role.(domain.Role)
		if !ok || string(userRole) != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "akses khusus admin"})
			return
		}
		c.Next()
	}
}
func StaffOrAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		 if c.Request.Method == "OPTIONS" {
            c.Next()
            return
        }
		role, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "akses khusus staff dan admin"})
			return
		}
				//cast domain

		userRole, ok := role.(domain.Role)
		if !ok || string(userRole) != "staff" && string(userRole) != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "akses khusus staff dan admin"})
			return
		}
		c.Next()
	}
}