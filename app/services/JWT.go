package services

import (
	"app/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kataras/iris/v12"
)

func GenerateJWT(user *models.User) (string, error) {
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":        user.Username,
		"id":          user.ID.Hex(),
		"date_issued": time.Now().UTC().String(),
		"expiry":      time.Now().Add(time.Hour * 24 * 7).UTC().String(), // expires after 7 days
	})

	return token.SignedString(hmacSampleSecret)
}

func VerifyJWT(tokenString string) (*string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		layout := "2006-01-02 15:04:05 -0700 MST"
		ExpiryString, es_type_ok := claims["expiry"].(string)
		Username, u_type_ok := claims["name"].(string)
		Expiry, exp_type_err := time.Parse(layout, ExpiryString)

		if !es_type_ok || !u_type_ok || exp_type_err != nil {
			return nil, fmt.Errorf("unexpected format")
		}

		if time.Now().After(Expiry) {
			return nil, fmt.Errorf("token expired")
		}

		return &Username, nil
	} else {
		return nil, fmt.Errorf("token not valid")
	}
}

func GetAndVerifyCookie(ctx iris.Context) (*string, error) {
	cookie := ctx.GetCookie("dylank-io-auth")
	if cookie == "" {
		return nil, fmt.Errorf("no cookie found")
	}

	username, verify_err := VerifyJWT(cookie)

	if verify_err != nil {
		return nil, verify_err
	} else {
		return username, nil
	}
}
