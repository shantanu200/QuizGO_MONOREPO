package middlewares

import (
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"quizGo/constants"
)

func CreateAuthToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"_id": userId,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SECRET := constants.EnvConstant("JWTSECRET")
	t, err := token.SignedString([]byte(SECRET))

	if err != nil {
		return "", err
	}

	return t, nil
}

func SetupJWT() fiber.Handler {
	SECRET := constants.EnvConstant("JWTSECRET")
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(SECRET)},
	})
}

func GetUserId(c *fiber.Ctx) (string, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := claims["_id"]

	return id.(string), nil
}
