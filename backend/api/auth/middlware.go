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
			errResp := apiError.NewUnauthorizedError("Missing authorization header.")
			ctx.AbortWithStatusJSON(errResp.Status, errResp)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			errResp := apiError.NewUnauthorizedError("Invalid authorization header format. Expected 'Bearer <token>'.")
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
			errResp := apiError.NewUnauthorizedError("Invalid or expired access token.")
			ctx.AbortWithStatusJSON(errResp.Status, errResp)
			return
		}

		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}
