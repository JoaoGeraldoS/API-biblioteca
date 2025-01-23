package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("SECRETKEY"))

type Claims struct {
	User string `json:"user"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GerarToken(user string, role string) (string, error) {

	expriracao := time.Now().Add(2 * time.Hour)

	claims := &Claims{
		User: user,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expriracao),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   user,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			ctx.Abort()
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			ctx.Abort()
			return
		}

		tokenString = parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(*Claims)
		ctx.Set("User", claims.User)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		if role != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Acesso permitido apenas para administradores", "debug_role": role})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
