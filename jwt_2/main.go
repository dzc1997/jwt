package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func mySigningKey() []byte {
	return []byte("AllYourBase")
}

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Unix() + 5,
		"iss":      "nn",
		"nbg":      time.Now().Unix() - 5,
		"username": "dzc",
	})
	ss, err := token.SignedString(mySigningKey())
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(ss)

	token, err = jwt.ParseWithClaims(ss, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey(), nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println(token.Claims.(*jwt.MapClaims))
	fmt.Println((*(token.Claims.(*jwt.MapClaims)))["username"])
}
