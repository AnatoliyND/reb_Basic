package internal

import "errors"

type Customer struct {
	*Overduer
	Name        string
	Age         int
	discount    bool
	resDiscount int

	//Balance     int
	//Debt        int
	//CalcDiscount func() (int, error)
}

type Overduer struct {
	Balans int
	Debt   int
}

type Debtor interface {
	WrOffDebt() error
}

const DEFAULT_DISCOUNT = 500

func NewCustomer(name string, age int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		discount: discount,
		//Balance:  balance,
		//Debt:     debt,
	}
}

func (c *Customer) WrOffDebt() error {

	if !c.discount {
		return errors.New("Discount not available")
	}
	c.resDiscount = DEFAULT_DISCOUNT - c.Overduer.Debt
	if c.resDiscount < 0 {
		c.Overduer.Debt = c.Overduer.Debt - DEFAULT_DISCOUNT
		c.resDiscount = 0
	}
	return nil
}

/*
func (c *Customer) ResDiscount() error {

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
*/
/*04_02
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
