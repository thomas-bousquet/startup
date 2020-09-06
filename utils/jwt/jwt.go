package jwt

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/models"
	uuid "github.com/thomas-bousquet/startup/utils/id-generator"
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
	expirationTime := now.Add(1 * time.Hour)
	claims := jwtGo.StandardClaims{
		Id:        uuid.New().String(),
		IssuedAt:  now.Unix(),
		ExpiresAt: expirationTime.Unix(),
		Subject:   user.Id,
	}

	token := jwtGo.NewWithClaims(jwt.SigningMethod, claims)
	signedToken, err := token.SignedString(jwt.SecretKey)

	if err != nil {
		return nil, NewUnexpectedError()
	}

	return &signedToken, nil
}

func (jwt JWT) ParseToken(token string) (*jwtGo.Token, error) {
	claims := &jwtGo.StandardClaims{}

	t, err := jwtGo.ParseWithClaims(token, claims, func(token *jwtGo.Token) (interface{}, error) {
		return []byte(jwt.SecretKey), nil
	})

	if err != nil {
		return nil, errors.NewAuthorizationError(err.Error())
	}

	return t, nil
}
