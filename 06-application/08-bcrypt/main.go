package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash, err := bcrypt.GenerateFromPassword([]byte(`pwd1234`), bcrypt.MinCost)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println("Hashed password:", string(hash))

	if err := bcrypt.CompareHashAndPassword(hash, []byte(`pwd1234`)); err != nil {
		fmt.Println("password doesnt match")
	} else {
		fmt.Println("password is a match")
	}
}
