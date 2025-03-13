package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"os"
)

// jwt parser file
var MySigninKey = []byte(os.Getenv("SECRET_KEY"))

func handleRequest() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super secret information")
}

// take an endpoint make sure current person with token is valid and  return the endpoint
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf(("invalid signing method"))
				}

				audience := "billing.jwtgo.io"
				issuer := "jwtgo.io"

				checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
				if !checkAudience {
					return nil, fmt.Errorf(("invalid audience"))
				}

				checkIssuer := token.Claims.(jwt.MapClaims).VerifyIssuer(issuer, false)
				if !checkIssuer {
					return nil, fmt.Errorf("invalid issuer")
				}

				return MySigninKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {
				endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "No authorization token provided")
		}
	})

}

func main() {
	fmt.Print("Server")
	handleRequest()
}
