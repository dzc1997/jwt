package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func mySigningKey() []byte {
	return []byte("AllYourBase")
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	c := MyClaims{
		Username: "nn",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "nn",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, err := token.SignedString(mySigningKey())
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(ss)

	token, err = jwt.ParseWithClaims(ss, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey(), nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println(token.Claims.(*MyClaims).Username)
}
