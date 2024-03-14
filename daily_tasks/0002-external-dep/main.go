package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/alexellis/hmac"
)

func main() {
	var msgInput string
	var secretInput string

	flag.StringVar(&msgInput, "msg", "", "plain text message")
	flag.StringVar(&secretInput, "secret", "", "private key used for hashing")
	flag.Parse()

	msgInput = strings.TrimSpace(msgInput)
	secretInput = strings.TrimSpace(secretInput)

	if msgInput == "" || secretInput == "" {
		log.Fatal("msg or secret can not be empty")
	}

	msg := []byte(msgInput)
	secret := []byte(secretInput)

	hash := hmac.Sign(msg, secret)

	fmt.Printf("Given text: %s\n", msgInput)
	fmt.Printf("Hash generated %x\n", hash)

	// checkInput := []byte(`Hello World!`)
	// validateErr := hmac.Validate(checkInput, fmt.Sprintf("sha1=%x", hash), string(secret))

	// if validateErr != nil {
	// 	log.Fatal(validateErr)
	// }

	// fmt.Println("Checking passed!!!")
}
