package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"main/component/api"
	"main/component/models"
	"main/component/services"
	"main/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (mid *ApiMiddleware) AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if tokenString == auth {
			c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
			c.Abort()
			return
		}

		token, errValidateJWT := services.ValidateJWT(tokenString, services.GetSecretKey())
		if errValidateJWT != nil {
			log.Errorf("Error validate JWT, error %v", errValidateJWT)
			api.Error(c, http.StatusUnauthorized)
			c.Abort()
			return
		}

		if !token.Valid {
			api.ErrorWithMessage(c, http.StatusUnauthorized, "Could not find bearer token in Authorization header")
			c.Abort()
			return
		}

		// Parse token to struct user and store into context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			api.ErrorWithMessage(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		utils.ShowInfoLogs(fmt.Sprintf("User login: %v", utils.LogFull(claims)))
		var currentUser models.AuthClaim
		errDecode := mapstructure.Decode(claims, &currentUser)
		if errDecode != nil {
			log.Errorf("Cannot decode claims %v to current user, error %v", utils.LogFull(claims), errDecode)
		}
		c.Set(utils.CurrentUser, currentUser)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://js-test-axios-msvvkj.stackblitz.io")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
