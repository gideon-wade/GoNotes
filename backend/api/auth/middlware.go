package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	apiError "github.com/gonotes/api/errors"
)

func AuthMiddleware(jwtSecret []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			errResp := apiError.NewUnauthorizedError(
				"Missing authorization header.",
				"An Authorization header with a valid access token is required to access this resource.",
				"401-unauthorized",
			)
			ctx.AbortWithStatusJSON(errResp.Status, errResp)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			errResp := apiError.NewUnauthorizedError(
				"Invalid authorization header format.",
				"Ensure the Authorization header is in the format: 'Bearer <access_token>'.",
				"401-unauthorized",
			)
			ctx.AbortWithStatusJSON(errResp.Status, errResp)
			return
		}

		tokenString := parts[1]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			errResp := apiError.NewUnauthorizedError(
				"Invalid or expired access token.",
				"The provided access token is invalid or has expired. Log in again to obtain a new token.",
				"401-unauthorized")
			ctx.AbortWithStatusJSON(errResp.Status, errResp)
			return
		}

		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
