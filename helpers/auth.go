package helpers

import (
	"fmt"
	"os"

	"santara.co.id/set/market-common-service-grpc/models"
	"santara.co.id/set/market-common-service-grpc/pkg/proto/common"

	"github.com/dgrijalva/jwt-go"
)

// Auth ...
func Auth(header string) (data *common.DataGetUser, err error) {
	runes := []rune(header)
	bearer := string(runes[0:6])
	if bearer != "Bearer" {
		return data, err
	}
	tokenString := string(runes[7:])
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	if !token.Valid {
		fmt.Println(err)
		return data, err
	}

	id := int64(int(claims["uid"].(float64)))
	user, err := models.FindByID(id)

	data = user
	if user.Id == 0 {
		return data, err
	}
	return data, err
}
