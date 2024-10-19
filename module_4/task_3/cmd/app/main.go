package main

import (
	"fmt"
	"module_4/task_1/internal/customer"
)

func main() {
	cust := customer.NewCustomer("Ivan", 24, 1000, 300, true)
	discount, _ := cust.CalcDiscount()
	price, _ := customer.CalcPrice(cust, 500)
	fmt.Println(cust, "\nDiscount:", discount, "\nTotal price:", price)
}
