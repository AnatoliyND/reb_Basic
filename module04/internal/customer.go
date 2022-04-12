package internal

import "errors"

type Customer struct {
	Name        string
	Age         int
	Balance     int
	Debt        int
	discount    bool
	resDiscount int
	//CalcDiscount func() (int, error)
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

func (c *Customer) CalcDiscount() (int, error) {

	if !c.discount {
		return 0, errors.New("Discount not available")
	}
	c.resDiscount = DEFAULT_DISCOUNT - c.Debt
	if c.resDiscount < 0 {
		c.Debt = c.Debt - DEFAULT_DISCOUNT
		c.resDiscount = 0
	}
	return c.resDiscount, nil
}

/*
func (c *Customer) ResDiscount() error {

	if !c.discount {
		return errors.New("Discount not available")
	}
	c.resDiscount = DEFAULT_DISCOUNT - c.Debt
	if c.resDiscount < 0 {
		c.Debt = c.Debt - DEFAULT_DISCOUNT
		c.resDiscount = 0
	}
	return nil
}
*/
/*
	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	} */
