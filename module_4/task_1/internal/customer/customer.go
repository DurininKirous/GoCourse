package customer

import "errors"

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}

func CalcPrice(cust Customer, price int) (int, error) {
	discount, error := cust.CalcDiscount()
	if error != nil {
		return 0, errors.New("Discount is not available")
	}
	return price - discount, nil
}
