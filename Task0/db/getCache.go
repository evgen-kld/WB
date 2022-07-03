package db

import (
	"Task0/JSON"
)

var items_elem struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func GetCashe(cache map[string]JSON.StructJSON) map[string]JSON.StructJSON {
	Connect()

	// Получить все id з БД
	mas := make([]string, 0)
	rows, _ := db.Query("SELECT id FROM Orders")
	for rows.Next() {
		var a string
		rows.Scan(&a)
		mas = append(mas, a)
	}

	// пройти по всем id, на основе каждого свормировать JSON.StructJSON добавить в map
	for _, elem := range mas {
		var str JSON.StructJSON
		query := `SELECT * FROM orders WHERE orders.id=$1`
		db.QueryRow(query, elem).Scan(&str.OrderUID, &str.TrackNumber, &str.Entry, &str.Locale,
			&str.InternalSignature, &str.CustomerID, &str.DeliveryService,
			&str.Shardkey, &str.SmID, &str.DateCreated, &str.OofShard)

		query = `SELECT * FROM delivery WHERE id=$1`
		db.QueryRow(query, elem).Scan(&str.OrderUID, &str.Delivery.Name, &str.Delivery.Phone,
			&str.Delivery.Zip, &str.Delivery.City, &str.Delivery.Address, &str.Delivery.Region, &str.Delivery.Email)

		query = `SELECT * FROM payment WHERE id=$1`
		db.QueryRow(query, elem).Scan(&str.OrderUID, &str.Payment.Transaction, &str.Payment.RequestID,
			&str.Payment.Currency, &str.Payment.Provider, &str.Payment.Amount, &str.Payment.PaymentDt,
			&str.Payment.Bank, &str.Payment.DeliveryCost, &str.Payment.GoodsTotal, &str.Payment.CustomFee)

		query = `SELECT * FROM item WHERE id=$1`
		rows, _ := db.Query(query, elem)
		for rows.Next() {
			a := items_elem
			rows.Scan(&str.OrderUID, &a.ChrtID, &a.TrackNumber, &a.Price, &a.Rid, &a.Name,
				&a.Sale, &a.Size, &a.TotalPrice, &a.NmID, &a.Brand, &a.Status)
			str.Items = append(str.Items, a)
		}
		cache[str.OrderUID] = str
	}
	db.Close()

	return cache
}
