package jwt

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/thomas-bousquet/user-service/errors"
	. "github.com/thomas-bousquet/user-service/errors"
	. "github.com/thomas-bousquet/user-service/models"
	uuid "github.com/thomas-bousquet/user-service/utils/id-generator"
	"time"
)

type JWT struct {
	SecretKey     []byte
	SigningMethod jwtGo.SigningMethod
}

func New(secretKey []byte) JWT {
	return JWT{
		SecretKey:     secretKey,
		SigningMethod: jwtGo.SigningMethodHS512,
	}
}

func (jwt JWT) CreateToken(user User) (*string, error) {
	now := time.Now()
	claims := jwtGo.StandardClaims{
		Id:        uuid.New().String(),
		IssuedAt:  now.Unix(),
		ExpiresAt: 0,
		Subject:   user.Id,
	}

	token := jwtGo.NewWithClaims(jwt.SigningMethod, claims)
	signedToken, err := token.SignedString(jwt.SecretKey)

	if err != nil {
		return nil, NewUnexpectedError(nil, nil)
	}

	return &signedToken, nil
}

func (jwt JWT) ParseToken(token string) (*jwtGo.Token, error) {
	claims := &jwtGo.StandardClaims{}

	t, err := jwtGo.ParseWithClaims(token, claims, func(token *jwtGo.Token) (interface{}, error) {
		return []byte(jwt.SecretKey), nil
	})

	if err != nil {
		return nil, errors.NewAuthorizationError(nil)
	}

	return t, nil
}
