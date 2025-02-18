package middleware

import (
	"log"
	"net/http"
	"strings"

	tokenservice "auth/token_service"
	user1 "user/domain"

	"github.com/gin-gonic/gin"
)

func LoggedIn(TS tokenservice.TokenService_imp) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for the Authorization header
		auth := c.GetHeader("Authorization")
		if auth == "" {
			log.Println("Authorization header missing")
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authorization header missing"})
			c.Abort()
			return
		}
		log.Println("Authorization header:", auth)

		// Extract the Bearer token
		authSplit := strings.Split(auth, " ")
		if len(authSplit) != 2 || strings.ToLower(authSplit[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authorization format is incorrect"})
			c.Abort()
			return
		}

		accessToken := authSplit[1]
		user, err := TS.ValidateAccessToken(accessToken)
		if err != nil {
			// Check for token expiration or invalid token
			if err.Error() == "token has expired" {
				// Try to get the refresh token from cookies
				refreshToken, err := c.Cookie("refresh_token")
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"msg": "Refresh token missing or expired"})
					c.Abort()
					return
				}

				// Validate the refresh token
				user, err = TS.ValidateRefreshToken(refreshToken)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid refresh token"})
					c.Abort()
					return
				}

				// Generate a new access token
				newAccessToken, err := TS.GenerateAccessToken(*user)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"msg": "Failed to generate new access token"})
					c.Abort()
					return
				}

				// Send the new access token in response headers (optional)
				c.Header("Authorization", "Bearer "+newAccessToken)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid access token"})
				c.Abort()
				return
			}
		}

		// Token is valid, store the user in the context

		usr := user1.CreateResponseUser(*user)
		c.Set("user", usr)

		// Proceed to the next handler
		c.Next()
	}
}
