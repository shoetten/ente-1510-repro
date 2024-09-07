package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/zalando/go-keyring"
)

const (
	secretService = "ente"
	secretUser    = "ente-1510-repro"
)

func main() {
	// get password
	secret, err := keyring.Get(secretService, secretUser)
	fmt.Println([]byte(secret))
	fmt.Println(len([]byte(secret)))

	if err != nil {
		key := make([]byte, 32)
		_, err = rand.Read(key)
		if err != nil {
			log.Fatal(fmt.Errorf("error generating key: %w", err))
		}
		secret = string(key)
		fmt.Println("setting cli secret to:")
		fmt.Println(key)
		fmt.Println(len(key))

		keySetErr := keyring.Set(secretService, secretUser, secret)

		if keySetErr != nil {
			log.Fatal(fmt.Errorf("error setting password in keyring: %w", keySetErr))
		}
	}
	fmt.Println("Returning..")
	keyringSecret, _ := keyring.Get(secretService, secretUser)
	fmt.Println([]byte(keyringSecret))
}
