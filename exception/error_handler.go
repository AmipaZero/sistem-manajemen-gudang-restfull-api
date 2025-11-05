package exception

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("internal server error: %v", r)})
                c.Abort()
            }
        }()
        c.Next()
    }
}