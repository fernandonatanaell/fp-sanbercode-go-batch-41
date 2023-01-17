package repository

import (
	"database/sql"
	"sanbercode-golang-final-project-nando/pkg/structs"
	"time"
)

func GetOrder(db *sql.DB, orderID string) (results structs.Htrans, err error) {
	sql := "SELECT * FROM htrans WHERE order_id = $1 LIMIT 1"

	err = db.QueryRow(sql, orderID).Scan(
		&results.OrderId,
		&results.HtransSubtotal,
		&results.HtransStatus,
		&results.Created_at,
		&results.Updated_at,
	)

	return
}

func GetAllOrders(db *sql.DB) (results []structs.Htrans, err error) {
	sql := "SELECT * FROM htrans"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var htrans = structs.Htrans{}

		err = rows.Scan(
			&htrans.OrderId,
			&htrans.HtransSubtotal,
			&htrans.HtransStatus,
			&htrans.Created_at,
			&htrans.Updated_at,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, htrans)
	}

	return
}

func CreateNewOrder(db *sql.DB, order_id string) error {
	sql := "INSERT INTO htrans (order_id, htrans_subtotal, htrans_status, created_at) VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql,
		order_id,
		"0",
		0,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	return errs.Err()
}

func UpdateOrder(db *sql.DB, htrans structs.Htrans) error {
	sql := "UPDATE htrans SET htrans_subtotal = $1, htrans_status = $2, updated_at = $3 WHERE order_id = $4"

	errs := db.QueryRow(sql,
		htrans.HtransSubtotal,
		1,
		time.Now().Format("2006-01-02 15:04:05"),
		htrans.OrderId,
	)

	return errs.Err()
}
