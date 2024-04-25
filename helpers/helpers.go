package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// APIResponse is a helper function to format API response
func APIResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateJWT(userID uint64) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(10 * time.Minute).Unix(),
	})

	tokenString, err = token.SignedString([]byte("your-secret-key"))
	return
}
