package authorization

import "github.com/golang-jwt/jwt/v5"

const JWT_SIGNING_KEY = "mysecretkey"

func GenerateJwtToken(username string) (string, error) {

	claims := jwt.MapClaims{"sub": username}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	return token.SignedString([]byte(JWT_SIGNING_KEY))

}
