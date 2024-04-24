package helpers

import (
	"github.com/gin-gonic/gin"
)

// APIResponse is a helper function to format API response
func APIResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
