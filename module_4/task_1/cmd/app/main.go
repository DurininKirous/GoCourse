package main

import (
	"errors"
	"fmt"
	"module_4/task_1/internal/customer"
)

const DEFAULT_DISCOUNT = 500

func main() {
	cust := customer.NewCustomer("Ivan", 24, 1000, 300, true)
	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount is not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	discount, _ := cust.CalcDiscount()
	price, _ := customer.CalcPrice(*cust, 500)
	fmt.Println(cust, "\nDiscount:", discount, "\nTotal price:", price)
}
