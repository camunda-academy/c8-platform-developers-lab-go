package services

import (
	"fmt"
)

func ChargeAmount(cardNumber string, cvc string, expiryDate string, amount float64) {
	fmt.Printf("charging card %v that expires on %v and has cvc %v with amount of %v\n", cardNumber, expiryDate, cvc, amount)
	fmt.Println("Payment completed")
}