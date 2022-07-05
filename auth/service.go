package auth

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Service interface{
	GenerateToken(userId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct{

}

var SECRET_KEY = []byte("INI rahasia")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService)GenerateToken(userId int)(string, error){
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil{
		return signedToken, err
	}


	return signedToken, nil
}

func (s *jwtService)ValidateToken(encodeToken string) (*jwt.Token, error){
	token , err := jwt.Parse(encodeToken, func(token *jwt.Token)(interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid Token")
		}
		return []byte(SECRET_KEY), nil

	})

	if err != nil{
		fmt.Println("error :", err)
		return token, err
	}

	return token, err
}

