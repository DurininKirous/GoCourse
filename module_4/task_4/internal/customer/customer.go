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

func CalcPrice(discounter Discounter, price int) (int, error) {
	discount, error := discounter.CalcDiscount()
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

func (cust *Customer) WrOffDebt() error {
	cust.Debt = 0

	return nil
}

type Discounter interface {
	CalcDiscount() (int, error)
}

type Debtor interface {
	WrOffDebt() error
}
