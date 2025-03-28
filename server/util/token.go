package util

import (
	"errors"
	"time"
	"webback/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var jwtkey = []byte(config.Conf.JWT.Key)

func keyFunc(t *jwt.Token) (interface{}, error) {
	return jwtkey, nil
}

const TokenExpireDuration = time.Hour * 24

type MyClaims struct {
	Name string
	jwt.StandardClaims
}

func GenerateToken(name string) (string, error) {
	c := MyClaims{
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "blockche",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(jwtkey)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, keyFunc)
	if err != nil {
		logrus.Errorf("ParseWithClaims error: %v", err)
		return nil, err
	}

	if c, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return c, nil
	} else {
		return nil, errors.New("Claims无效")
	}
}
