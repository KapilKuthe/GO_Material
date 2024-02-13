package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12"
)

var secretKey = []byte("kYa_phuk=Ke-aya?Be")

func GenerateToken(email string, userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 12).Unix(),
	})
	return token.SignedString(secretKey)
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing Method!")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("unable to parse token")
	}

	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claim")
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}

func AuthMiddleware(ctx iris.Context) {
	token := ctx.GetHeader("Authorization")

	if token == "" {
		ctx.StopWithJSON(iris.StatusUnauthorized, iris.Map{"message": "Not Authorized!"})
		return
	}

	uid, err := VerifyToken(token)
	if err != nil {
		ctx.StopWithJSON(iris.StatusUnauthorized, iris.Map{"message": "Not Authorized!"})
		return
	}

	ctx.SetID(uid)
	ctx.Next()
}
