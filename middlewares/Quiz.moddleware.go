package middlewares

import (
	"errors"
	"quizGo/constants"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateQuizToken(userId string, quizId string, duration int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"quizId": quizId,
		"exp":    time.Now().Add(time.Duration(10) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SECRET := constants.EnvConstant("QUIZSECRET")
	t, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", nil
	}

	return t, nil
}

func ValidateToken(token string) (bool, error) {
	SECRET := constants.EnvConstant("QUIZSECRET")

	result, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return false, errors.New("unable to attempt quiz. your quiz attempt is expired/over")
	}

	if _, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		return true, errors.New("your quiz is active. Please complete it before expiration")
	}
	return true, nil
}
