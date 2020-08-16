package jwt

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	. "github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/models"
	uuid "github.com/thomas-bousquet/startup/utils/id-generator"
	"time"
)

type JWT struct {
	SecretKey []byte
}

func New(secretKey []byte) JWT {
	return JWT{
		SecretKey: secretKey,
	}
}

func (jwt JWT) CreateToken(user User) (*string, error) {
	now := time.Now()
	expirationTime := now.Add(5 * time.Minute)
	claims := jwtGo.StandardClaims{
		Id:        uuid.New().String(),
		IssuedAt:  now.Unix(),
		ExpiresAt: expirationTime.Unix(),
		Subject:   user.Id,
	}

	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS512, claims)
	signedToken, err := token.SignedString(jwt.SecretKey)

	if err != nil {
		println(err)
		return nil, NewUnexpectedError()
	}

	return &signedToken, nil
}
