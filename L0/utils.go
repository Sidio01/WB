package main

import (
	"fmt"
)

func validate(o Order) error {
	err := o.validateOrder()
	if err != nil {
		return fmt.Errorf("order error - %v", err)
	}
	err = o.Delivery.validateDelivery()
	if err != nil {
		return fmt.Errorf("delivery error - %v", err)
	}
	err = o.Payment.validatePayment()
	if err != nil {
		return fmt.Errorf("payment error - %v", err)
	}
	for _, i := range o.Items {
		err = i.validateItem()
		if err != nil {
			return fmt.Errorf("item error - %v", err)
		}
	}
	return nil
}
