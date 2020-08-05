package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

func GenToken(data AuthModel) string {
	authClaims := jwt.MapClaims{
		"id":    data.ID,
		"phone": data.Phone,
		"email": data.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	auth := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)

	token, _ := auth.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	return token
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
