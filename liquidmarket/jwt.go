package liquidmarket

import (
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
)

func insecureGetUserFromJWT(r *http.Request) (*User, error) {

	tokenString, err := extractToken(r)

	if err != nil {
		return nil, err
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return token, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)

	user := User{
		GoogleID: claims["sub"].(string),
		FirstName: claims["given_name"].(string),
		LastName: claims["family_name"].(string),
		Email: claims["email"].(string),
	}
	
	return &user, nil
}

func secureGetUserFromJWT(r *http.Request) (*User, error) {
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		hmacSampleSecret := []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	user := User{ 
		GoogleID: "111307193244795741281",
		FirstName: "Thomas",
		LastName: "Horrobin",
		Email: "thomasroberthorrobin@gmail.com",
	}

	return &user, nil
}

func extractToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", fmt.Errorf("No header provided.")
	}
	split := strings.Split(header, " ");
	if len(split) < 2 {
		return "", fmt.Errorf("Authorization incorectly formatted")
	}
	return split[1], nil
}