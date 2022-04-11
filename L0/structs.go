package main

import "fmt"

type Order struct {
	CustomerID        string   `json:"customer_id"`
	DateCreated       string   `json:"date_created"`
	Delivery          Delivery `json:"delivery"`
	DeliveryService   string   `json:"delivery_service"`
	Entry             string   `json:"entry"`
	InternalSignature string   `json:"internal_signature"`
	Items             []Item   `json:"items"`
	Locale            string   `json:"locale"`
	OofShard          string   `json:"oof_shard"`
	OrderUID          string   `json:"order_uid"`
	Payment           Payment  `json:"payment"`
	Shardkey          string   `json:"shardkey"`
	SmID              int64    `json:"sm_id"`
	TrackNumber       string   `json:"track_number"`
}

func (o *Order) validateOrder() error {
	switch {
	case o.CustomerID == "":
		return fmt.Errorf("the customer_id field is required")
	case o.DateCreated == "":
		return fmt.Errorf("the date_created field is required")
	case o.DeliveryService == "":
		return fmt.Errorf("the delivery_service field is required")
	case o.Entry == "":
		return fmt.Errorf("the entry field is required")
	case o.InternalSignature == "":
		return fmt.Errorf("the internal_signature field is required")
	case o.Locale == "":
		return fmt.Errorf("the locale field is required")
	case o.OofShard == "":
		return fmt.Errorf("the oof_shard field is required")
	case o.OrderUID == "":
		return fmt.Errorf("the order_uid field is required")
	case o.Shardkey == "":
		return fmt.Errorf("the shardkey field is required")
	case o.TrackNumber == "":
		return fmt.Errorf("the track_number field is required")
	case o.SmID <= 0:
		return fmt.Errorf("the sm_id field must be positive integer")
	}
	return nil
}

type Delivery struct {
	Address string `json:"address"`
	City    string `json:"city"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Region  string `json:"region"`
	Zip     string `json:"zip"`
}

func (d *Delivery) validateDelivery() error {
	switch {
	case d.Address == "":
		return fmt.Errorf("the address field is required")
	case d.City == "":
		return fmt.Errorf("the city field is required")
	case d.Email == "":
		return fmt.Errorf("the email field is required")
	case d.Name == "":
		return fmt.Errorf("the name field is required")
	case d.Phone == "":
		return fmt.Errorf("the phone field is required")
	case d.Region == "":
		return fmt.Errorf("the region field is required")
	case d.Zip == "":
		return fmt.Errorf("the zip field is required")
	}
	return nil
}

type Item struct {
	Brand       string `json:"brand"`
	ChrtID      int64  `json:"chrt_id"`
	Name        string `json:"name"`
	NmID        int64  `json:"nm_id"`
	Price       int64  `json:"price"`
	Rid         string `json:"rid"`
	Sale        int64  `json:"sale"`
	Size        string `json:"size"`
	Status      int64  `json:"status"`
	TotalPrice  int64  `json:"total_price"`
	TrackNumber string `json:"track_number"`
}

func (i *Item) validateItem() error {
	switch {
	case i.Brand == "":
		return fmt.Errorf("the brand field is required")
	case i.ChrtID <= 0:
		return fmt.Errorf("the chrt_id field must be positive integer")
	case i.Name == "":
		return fmt.Errorf("the name field is required")
	case i.NmID <= 0:
		return fmt.Errorf("the nm_id field must be positive integer")
	case i.Price <= 0:
		return fmt.Errorf("the price field must be positive integer")
	case i.Rid == "":
		return fmt.Errorf("the rid field is required")
	case i.Sale < 0 && i.Sale > 100:
		return fmt.Errorf("the sale field must be integer from 0 to 100")
	case i.Size == "":
		return fmt.Errorf("the size field is required")
	case i.Status <= 0:
		return fmt.Errorf("the status field must be positive integer")
	case i.TotalPrice <= 0:
		return fmt.Errorf("the total_price must be positive integer")
	case i.TrackNumber == "":
		return fmt.Errorf("the track_number field is required")
	}
	return nil
}

type Payment struct {
	Amount       int64  `json:"amount"`
	Bank         string `json:"bank"`
	Currency     string `json:"currency"`
	CustomFee    int64  `json:"custom_fee"`
	DeliveryCost int64  `json:"delivery_cost"`
	GoodsTotal   int64  `json:"goods_total"`
	PaymentDt    int64  `json:"payment_dt"`
	Provider     string `json:"provider"`
	RequestID    string `json:"request_id"`
	Transaction  string `json:"transaction"`
}

func (p *Payment) validatePayment() error {
	switch {
	case p.Amount <= 0:
		return fmt.Errorf("the amount field must be positive integer")
	case p.Bank == "":
		return fmt.Errorf("the bank field is required")
	case p.Currency == "":
		return fmt.Errorf("the currency field is required")
	case p.CustomFee < 0:
		return fmt.Errorf("the custom_fee field must be greater than or equal to zero")
	case p.DeliveryCost <= 0:
		return fmt.Errorf("the delivery_cost field must be positive integer")
	case p.GoodsTotal <= 0:
		return fmt.Errorf("the goods_total field must be positive integer")
	case p.PaymentDt <= 0:
		return fmt.Errorf("the payment_dt field must be positive integer")
	case p.Provider == "":
		return fmt.Errorf("the provider field is required")
	case p.RequestID == "":
		return fmt.Errorf("the request_id field is required")
	case p.Transaction == "":
		return fmt.Errorf("the transaction field is required")
	}
	return nil
}
