package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"main/component/models"
	"main/utils"
	"time"
)

func GetSecretKey() []byte {
	return []byte(utils.AppConfig.Auth.JwtSecretKey)
}

func ValidateJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return token, nil
}

func GenerateToken(claim models.AuthClaim) (string, string, error) {
	claims := jwt.MapClaims{
		"authorized":  claim.Authorized,
		"id":          claim.Id,
		"roles":       claim.Roles,
		"fullName":    claim.FullName,
		"permissions": claim.Permissions,
		"email":       claim.Email,
	}

	// Gen access token
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenRes, errAccessToken := accessToken.SignedString(GetSecretKey())
	if errAccessToken != nil {
		log.Errorf("Error create access token %v", errAccessToken)
		return "", "", errAccessToken
	}

	// Gen refresh token
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenRes, errRefreshToken := refreshToken.SignedString(GetSecretKey())
	if errRefreshToken != nil {
		log.Errorf("Error create refresh token %v", errRefreshToken)
		return "", "", errRefreshToken
	}
	return accessTokenRes, refreshTokenRes, nil
}
