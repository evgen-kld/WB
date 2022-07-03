package db

import (
	"Task0/JSON"
)

func Insert(s JSON.StructJSON) {

	Connect()

	inOrder := `
	INSERT INTO orders(id, track_number, entry, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`
	db.Exec(inOrder, s.OrderUID, s.TrackNumber, s.Entry, s.Locale, s.InternalSignature, s.CustomerID, s.DeliveryService, s.Shardkey, s.SmID, s.DateCreated, s.OofShard)

	inDelivery := `
	INSERT INTO delivery(id, name, phone, zip, city, address, region, email) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	db.Exec(inDelivery, s.OrderUID, s.Delivery.Name, s.Delivery.Phone, s.Delivery.Zip, s.Delivery.City, s.Delivery.Address, s.Delivery.Region, s.Delivery.Email)
	inPayment := `
	INSERT INTO payment(id, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

	db.Exec(inPayment, s.OrderUID, s.Payment.Transaction, s.Payment.RequestID, s.Payment.Currency, s.Payment.Provider, s.Payment.Amount, s.Payment.PaymentDt, s.Payment.Bank, s.Payment.DeliveryCost, s.Payment.GoodsTotal, s.Payment.CustomFee)
	for _, elem := range s.Items {
		inItem := `
			INSERT INTO item(id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
		db.Exec(inItem, s.OrderUID, elem.ChrtID, elem.TrackNumber, elem.Price, elem.Rid, elem.Name, elem.Sale, elem.Size, elem.TotalPrice, elem.NmID, elem.Brand, elem.Status)
	}

	db.Close()
}
