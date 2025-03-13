package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
	"time"
)

var MySigninKey = []byte(os.Getenv("SECRET_KEY"))

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJwt()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Fprintf(w, string(validToken))
}

func GetJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "aditya"
	claims["audience"] = "billing.jwtgo.io"
	claims["issuer"] = "jwtgo.io"
	claims["expiration"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(MySigninKey)

	if err != nil {
		log.Printf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleRequests()
}
