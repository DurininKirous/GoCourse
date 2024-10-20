package main

import (
	"errors"
	"fmt"
	"module_4/task_1/internal/customer"
)

func main() {
	cust := customer.NewCustomer("Ivan", 24, 1000, 300, true)
	fmt.Println(cust)
	error := startTransactionDynamic(cust)
	fmt.Println(error)
}

func startTransactionDynamic(debtor interface{}) error {
	cust, ok := (debtor).(customer.Debtor)
	if !ok {
		return errors.New("Incorrect type")
	} else {
		cust.WrOffDebt()
	}
	return nil
}
