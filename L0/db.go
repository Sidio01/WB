package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func dbConnection() (*sql.DB, error) {
	envs, err := godotenv.Read(".env")
	if err != nil {
		return new(sql.DB), err
	}
	host := envs["DB_HOST"]
	port, _ := strconv.Atoi(envs["DB_PORT"])
	db_user := envs["DB_USER"]
	password := envs["DB_PASS"]
	dbname := envs["DB_NAME"]

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, db_user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return new(sql.DB), err
	}
	return db, nil
}

func dbInsertItem(i Item) (int, error) {
	db, err := dbConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id int
	query := `
	INSERT INTO "item" ("brand", "chrt_id", "name", "nm_id", "price", "rid", "sale", "size", "status", "total_price", "track_number") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING item_id`
	err = db.QueryRow(query, i.Brand, i.ChrtID, i.Name, i.NmID, i.Price, i.Rid, i.Sale, i.Size, i.Status, i.TotalPrice, i.TrackNumber).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func dbInsertPayment(p Payment) (int, error) {
	db, err := dbConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id int
	query := `
	INSERT INTO "payment" ("amount", "bank", "currency", "custom_fee", "delivery_cost", "goods_total", "payment_dt", "provider", "request_id", "transaction") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING payment_id`
	err = db.QueryRow(query, p.Amount, p.Bank, p.Currency, p.CustomFee, p.DeliveryCost, p.GoodsTotal, p.PaymentDt, p.Provider, p.RequestID, p.Transaction).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func dbInsertDelivery(d Delivery) (int, error) {
	db, err := dbConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id int
	query := `
	INSERT INTO "delivery" ("address", "city", "email", "name", "phone", "region", "zip") 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING delivery_id`
	err = db.QueryRow(query, d.Address, d.City, d.Email, d.Name, d.Phone, d.Region, d.Zip).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func dbInsertOrder(o Order, deliveryId, paymentId int) (int, error) {
	db, err := dbConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var id int
	query := `
	INSERT INTO "order" ("customer_id", "date_created", "delivery_service", "entry", "internal_signature", "locale", "oof_shard", "order_uid", "shardkey", "sm_id", "track_number", "payment_id", "delivery_id") 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING order_id`
	err = db.QueryRow(query, o.CustomerID, o.DateCreated, o.DeliveryService, o.Entry, o.InternalSignature, o.Locale, o.OofShard, o.OrderUID, o.Shardkey, o.SmID, o.TrackNumber, paymentId, deliveryId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func dbInsertOrderItems(orderId, itemId int) error {
	db, err := dbConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO "order-item" ("order_id", "item_id") VALUES ($1, $2)`
	_, err = db.Exec(query, orderId, itemId)
	if err != nil {
		return err
	}
	return nil
}

func updateDb(o Order) (int, error) {
	paymentId, err := dbInsertPayment(o.Payment)
	if err != nil {
		return 0, err
	}
	deliveryId, err := dbInsertDelivery(o.Delivery)
	if err != nil {
		return 0, err
	}
	orderId, err := dbInsertOrder(o, deliveryId, paymentId)
	if err != nil {
		return 0, err
	}
	for _, i := range o.Items {
		itemId, err := dbInsertItem(i)
		if err != nil {
			return 0, err
		}
		err = dbInsertOrderItems(orderId, itemId)
		if err != nil {
			return 0, err
		}
	}
	return orderId, nil
}

func dbGet() (map[string]interface{}, error) {
	cache := make(map[string]interface{})

	db, err := dbConnection()
	if err != nil {
		return cache, err
	}
	defer db.Close()

	main_query := `
	SELECT "order_id", "customer_id", "date_created", "delivery_service", "entry", "internal_signature", "locale", "oof_shard", "order_uid", "shardkey", "sm_id", "track_number",
	"amount", "bank", "currency", "custom_fee", "delivery_cost", "goods_total", "payment_dt", "provider", "request_id", "transaction", "address", "city", "email", "name",
	"phone", "region", "zip"
	FROM "order"
	JOIN payment ON "order".payment_id = payment.payment_id
	JOIN delivery ON "order".delivery_id = delivery.delivery_id`
	main_rows, err := db.Query(main_query)
	if err != nil {
		return cache, err
	}
	for main_rows.Next() {
		var (
			itemsList         []Item
			orderId           int
			customerId        string
			dateCreated       string
			deliveryService   string
			entry             string
			internalSignature string
			locale            string
			oofShard          string
			orderUid          string
			shardkey          string
			smId              int64
			orderTrackNumber  string
			amount            int64
			bank              string
			currency          string
			customFee         int64
			deliveryCost      int64
			goodsTotal        int64
			paymentDt         int64
			provider          string
			requestId         string
			transaction       string
			address           string
			city              string
			email             string
			name              string
			phone             string
			region            string
			zip               string
		)
		err := main_rows.Scan(&orderId,
			&customerId,
			&dateCreated,
			&deliveryService,
			&entry,
			&internalSignature,
			&locale,
			&oofShard,
			&orderUid,
			&shardkey,
			&smId,
			&orderTrackNumber,
			&amount,
			&bank,
			&currency,
			&customFee,
			&deliveryCost,
			&goodsTotal,
			&paymentDt,
			&provider,
			&requestId,
			&transaction,
			&address,
			&city,
			&email,
			&name,
			&phone,
			&region,
			&zip)
		if err != nil {
			return cache, err
		}

		item_query := `
		SELECT "brand", "chrt_id", "name", "nm_id", "price", "rid", "sale", "size", "status", "total_price", "item"."track_number"
		FROM "order"
		JOIN "order-item" ON "order".order_id = "order-item".order_id
		JOIN item ON item.item_id = "order-item".item_id
		WHERE "order".order_id = $1`
		item_rows, err := db.Query(item_query, orderId)
		if err != nil {
			return cache, err
		}
		for item_rows.Next() {
			var (
				brand           string
				chrtId          int64
				name            string
				nmId            int64
				price           int64
				rid             string
				sale            int64
				size            string
				status          int64
				totalPrice      int64
				itemTrackNumber string
			)
			err := item_rows.Scan(&brand,
				&chrtId,
				&name,
				&nmId,
				&price,
				&rid,
				&sale,
				&size,
				&status,
				&totalPrice,
				&itemTrackNumber)
			if err != nil {
				return cache, err
			}
			itemsList = append(itemsList, Item{Brand: brand,
				ChrtID:      chrtId,
				Name:        name,
				NmID:        nmId,
				Price:       price,
				Rid:         rid,
				Sale:        sale,
				Size:        size,
				Status:      status,
				TotalPrice:  totalPrice,
				TrackNumber: itemTrackNumber,
			})
		}

		cache[strconv.Itoa(orderId)] = Order{
			CustomerID:  customerId,
			DateCreated: dateCreated,
			Delivery: Delivery{
				Address: address,
				City:    city,
				Email:   email,
				Name:    name,
				Phone:   phone,
				Region:  region,
				Zip:     zip,
			},
			DeliveryService:   deliveryService,
			Entry:             entry,
			InternalSignature: internalSignature,
			Items:             itemsList,
			Locale:            locale,
			OofShard:          oofShard,
			OrderUID:          orderUid,
			Payment: Payment{
				Amount:       amount,
				Bank:         bank,
				Currency:     currency,
				CustomFee:    customFee,
				DeliveryCost: deliveryCost,
				GoodsTotal:   goodsTotal,
				PaymentDt:    paymentDt,
				Provider:     provider,
				RequestID:    requestId,
				Transaction:  transaction,
			},
			Shardkey:    shardkey,
			SmID:        smId,
			TrackNumber: orderTrackNumber,
		}
	}
	return cache, nil
}
