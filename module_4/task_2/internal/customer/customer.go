package customer

import "errors"

type Customer struct {
	Name     string
	Age      int
	Balance  int
	Debt     int
	discount bool
}

const DEFAULT_DISCOUNT = 500

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		discount: discount,
	}
}

func CalcPrice(cust Customer, price int) (int, error) {
	discount, error := cust.CalcDiscount()
	if error != nil {
		return 0, errors.New("Discount is not available")
	}
	return price - discount, nil
}

func (cust *Customer) CalcDiscount() (int, error) {
	if !cust.discount {
		return 0, errors.New("Discount is not available")
	}
	result := DEFAULT_DISCOUNT - cust.Debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}
