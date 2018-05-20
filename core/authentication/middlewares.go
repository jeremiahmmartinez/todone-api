package authentication

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"todone-api/data"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := InitJWTAuthenticationBackend()
	sharedData := data.SharedData()

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid {
		sharedData.SetToken(token)

		claims := token.Claims.(jwt.MapClaims)
		userId := uint64(claims["sub"].(float64))

		sharedData.SetUserId(&userId)

		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
