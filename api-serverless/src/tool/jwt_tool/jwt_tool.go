package jwt_tool

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSignMethod = jwt.SigningMethodES256
)

func GenerateToken[T any](data T, secretKey string, expired time.Duration) string {
	claim := jwt.MapClaims{
		"userData":  data,
		"expiredAt": time.Now().Add(expired).Unix(),
	}

	t := jwt.NewWithClaims(jwtSignMethod, claim)

	token, _ := t.SignedString(secretKey)
	return token
}

func GetData[T any](token string, secretKey string) (T, error) {
	var data T
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return data, err
	}

	if claims, ok :=
		parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		expiredAt, _ := claims["expiredAt"].(int64)
		if expiredAt < time.Now().Unix() {
			return data, errors.New("token expired")
		}
		data, _ = claims["userData"].(T)
	} else {
		return data, errors.New("Invalid token")
	}
	return data, nil
}
