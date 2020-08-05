package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"khanhnguyen234/api-gateway/common"
	"net/http"
)

type AuthResponse struct {
	ID    string
	Phone string
	Email string
	Token string
}

func SignUpController(validator SignUpValidator) SignUpRequest {
	auth := validator.Request

	validator.auth.Email = validator.Request.Email
	validator.auth.setPassword(validator.Request.Password)

	InsertAuth(validator.auth)
	return auth
}

func SignInController(validator SignUpValidator) (AuthResponse, common.ErrorFields) {
	authModel, err := FindOneUser(validator.Request.Email)

	if err != nil {
		return AuthResponse{}, common.NewErrorField("email", errors.New("Not Registered email"))
	}

	if authModel.checkPassword(validator.Request.Password) != nil {
		return AuthResponse{}, common.NewErrorField("password", errors.New("Invalid password"))
	}

	token := GenToken(authModel)

	return AuthResponse{
		ID:    authModel.ID.Hex(),
		Phone: authModel.Phone,
		Email: authModel.Email,
		Token: token,
	}, common.ErrorFields{}
}

func AuthContextController(r *http.Request) (AuthResponse, error) {
	token, err := VerifyToken(r)

	if err != nil {
		return AuthResponse{}, errors.New("token invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return AuthResponse{}, errors.New("token claims error")
	}

	return AuthResponse{
		ID:    claims["id"].(string),
		Phone: claims["phone"].(string),
		Email: claims["email"].(string),
	}, nil
}
