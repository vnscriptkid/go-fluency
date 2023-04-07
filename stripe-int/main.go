package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type TheStripe struct {
	API *client.API
}

func NewTheStripe(key string) *TheStripe {
	s := &TheStripe{
		API: &client.API{},
	}

	s.Init(key)

	return s
}

func (s *TheStripe) Init(key string) {
	s.API.Init(key, nil)
}

func (s *TheStripe) CreatePI(params *stripe.PaymentIntentParams) (pi *stripe.PaymentIntent, err error) {
	pi, err = s.API.PaymentIntents.New(params)
	return
}

func (s *TheStripe) ConfirmPI(id string) (pi *stripe.PaymentIntent, err error) {
	pi, err = s.API.PaymentIntents.Confirm(id, nil)
	return
}

func (s *TheStripe) CapturePI(id string) (pi *stripe.PaymentIntent, err error) {
	pi, err = s.API.PaymentIntents.Capture(id, nil)
	return
}

func (s *TheStripe) RefundPI(id string, amount int64) (refund *stripe.Refund, err error) {
	params := stripe.RefundParams{
		Charge: &id,
	}

	if amount > 0 {
		params.Amount = &amount
	}

	refund, err = s.API.Refunds.New(&params)
	return
}

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	stripeApiKey := os.Getenv("STRIPE_API_KEY")

	if stripeApiKey == "" || !strings.HasPrefix(stripeApiKey, "sk_") {
		log.Fatal("invalid stripe pKey")
	}

	theStripe := NewTheStripe(stripeApiKey)

	fmt.Println(theStripe, stripeApiKey)

	var amount int64 = 10000
	customerID := "cus_NfYNUsEnyFTMck"
	currency := "usd"
	paymentMethod := "card_1MuDKZFZszgLqIoI1rbcDItn"

	pi, err := theStripe.CreatePI(&stripe.PaymentIntentParams{
		Amount:        &amount,
		Customer:      &customerID,
		Currency:      &currency,
		PaymentMethod: &paymentMethod,
	})

	if err != nil {
		fmt.Println("err creating pi")
		return
	}

	fmt.Printf("pi created, client secret: %v, pi: %v", pi.ClientSecret, pi)

	//// CONFIRM
	pi, err = theStripe.ConfirmPI(pi.ID)

	if err != nil {
		fmt.Println("err confirm pi")
		return
	}
	fmt.Printf("pi confirmed: %v | pi charge: %v", pi, pi.Charges.Data[0].ID)

	////// CAPTURE
	// _, err = theStripe.CapturePI(pi.ID)

	// if err != nil {
	// 	fmt.Println("err captured pi")
	// 	return
	// }
	// fmt.Println("pi captured")

	////// REFUND
	_, err = theStripe.RefundPI(pi.Charges.Data[0].ID, 0)

	if err != nil {
		fmt.Println("err refund pi")
		return
	}
	fmt.Println("pi refunded")

}
