package main

import (
	"fmt"
	"os"
)

func main() {
	to, id := os.Getenv("SMS_TO"), os.Getenv("SMS_TEMPLATE_ID")
	if to == "" || id == "" {
		fmt.Println("set SMS_TO and SMS_TEMPLATE_ID")
		return
	}
	e := PaymentEvent{1999, "USD", to, id, map[string]string{"amount": "19.99"}}
	if err := NotifyPayment(NewInfraiClient(), e); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("payment notification accepted")
}
