package repository

import (
	"database/sql"
	"sanbercode-golang-final-project-nando/pkg/structs"
)

func InsertDtrans(db *sql.DB, dtrans structs.Dtrans) error {
	sql := "INSERT INTO dtrans (order_id, menu_id, dtrans_quantity, dtrans_total) VALUES($1, $2, $3, $4)"

	errs := db.QueryRow(sql,
		dtrans.OrderId,
		dtrans.MenuId,
		dtrans.DtransQuantity,
		dtrans.DtransTotal,
	)

	return errs.Err()
}

func GetDtransFromGivenId(db *sql.DB, order_id string) (results []structs.Dtrans, err error) {
	sql := "SELECT * FROM dtrans WHERE order_id = $1"

	rows, err := db.Query(sql, order_id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var dtrans = structs.Dtrans{}

		err = rows.Scan(&dtrans.OrderId, &dtrans.MenuId, &dtrans.DtransQuantity, &dtrans.DtransTotal)
		if err != nil {
			panic(err)
		}

		results = append(results, dtrans)
	}

	return
}
