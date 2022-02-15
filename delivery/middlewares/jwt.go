package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("rahasia"),
	})
}

func CreateToken(userid int, userName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userid
	claims["username"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("rahasia"))
}

func GetUserName(e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userName := claims["username"].(string)
		if userName == "" {
			return userName, fmt.Errorf("empty username")
		}
		return userName, nil
	}
	return "", fmt.Errorf("invalid user")
}

func ExtractToken(jwtSecret string, e echo.Context) (int,string,error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userid := int(claims["id"].(float64))
		userName := claims["username"].(string)
		if userid == 0 {
			return userid,userName, fmt.Errorf("invalid id")
		}
		return userid,userName, nil
	}
	return 0,"", fmt.Errorf("invalid user")
}