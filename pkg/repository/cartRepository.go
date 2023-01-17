package repository

import (
	"database/sql"
	"sanbercode-golang-final-project-nando/pkg/structs"
)

func GetCart(db *sql.DB, cart structs.Cart) (results structs.Cart, err error) {
	sql := "SELECT * FROM carts WHERE order_id = $1 AND menu_id = $2 LIMIT 1"

	err = db.QueryRow(sql, cart.OrderId, cart.MenuId).Scan(
		&results.OrderId,
		&results.MenuId,
		&results.CartQuantity,
	)

	return
}

func GetAllCarts(db *sql.DB, order_id string) (results []structs.Cart, err error) {
	sql := "SELECT * FROM carts WHERE order_id = $1"

	rows, err := db.Query(sql, order_id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var cart = structs.Cart{}

		err = rows.Scan(&cart.OrderId, &cart.MenuId, &cart.CartQuantity)
		if err != nil {
			panic(err)
		}

		results = append(results, cart)
	}

	return
}

func AddToCart(db *sql.DB, cart structs.Cart) error {
	sql := "INSERT INTO carts (order_id, menu_id, cart_quantity) VALUES($1, $2, $3)"

	errs := db.QueryRow(sql,
		cart.OrderId,
		cart.MenuId,
		cart.CartQuantity,
	)

	return errs.Err()
}

func UpdateCart(db *sql.DB, cart structs.Cart) error {
	sql := "UPDATE carts SET cart_quantity = $1 WHERE order_id = $2 AND menu_id = $3"

	errs := db.QueryRow(sql,
		cart.CartQuantity,
		cart.OrderId,
		cart.MenuId,
	)

	return errs.Err()
}

func DeleteFromCart(db *sql.DB, cart structs.Cart) error {
	sql := "DELETE FROM carts WHERE order_id = $1 AND menu_id = $2"

	errs := db.QueryRow(sql,
		cart.OrderId,
		cart.MenuId,
	)

	return errs.Err()
}

func ClearCart(db *sql.DB, order_id string) error {
	sql := "DELETE FROM carts WHERE order_id = $1"

	errs := db.QueryRow(sql, order_id)

	return errs.Err()
}
