package middleware

import "github.com/golang-jwt/jwt"

func GenerateJWTToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID

	tokenString, err := token.SignedString([]byte("SuperSecretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
