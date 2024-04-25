package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRequired(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the session token from the headers
		authHeader := c.GetHeader("Authorization")

		// If no token is provided, return an error
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Remove the "Bearer " prefix
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token with the database or JWT here
		var expiresAt time.Time
		err := db.QueryRow("SELECT expires_at FROM sessions WHERE token=$1", token).Scan(&expiresAt)
		if err != nil {
			fmt.Printf("Error querying for token %s: %v\n", token, err)
			if err == sql.ErrNoRows {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate session token"})
			}
			return
		}

		// Check if the session has expired
		if time.Now().After(expiresAt) {
			// Delete the session from the database
			db.Exec("DELETE FROM sessions WHERE token = $1", token)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session token has expired"})
			return
		}

		// If the token is valid, continue to the next handler
		c.Next()
	}
}
