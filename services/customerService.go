package services

import (
	"fmt"
	"strconv"
)

func getCustomerCredit(customerId string) float64 {
	a := len(customerId) - 2
	b := len(customerId)

	i, _ := strconv.ParseFloat(customerId[a:b], 64)
	return i
}

func DeductCredit(customerId string, amount float64) float64 {

	credit := getCustomerCredit(customerId);
	openAmount := 0.
	deductedCredit := 0.
	if (credit > amount) {
		deductedCredit = amount;
		openAmount = 0.0;
	} else {
		openAmount = amount - credit;
		deductedCredit = credit;
	}
	fmt.Printf("charged %v from the credit, open amount is %v \n", deductedCredit, openAmount)
	return openAmount;
}
