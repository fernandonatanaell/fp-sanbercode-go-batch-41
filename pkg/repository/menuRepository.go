package repository

import (
	"database/sql"
	"sanbercode-golang-final-project-nando/pkg/structs"
	"time"
)

func SearchMenu(db *sql.DB, menu_name string) (results []structs.Menu, err error) {
	sql := "SELECT * FROM menu WHERE menu_name ILIKE $1 AND deleted_at IS NULL"

	rows, err := db.Query(sql, "%"+menu_name+"%")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var menu = structs.Menu{}

		err = rows.Scan(
			&menu.MenuId,
			&menu.MenuName,
			&menu.MenuDescription,
			&menu.MenuPrice,
			&menu.Created_at,
			&menu.Updated_at,
			&menu.Deleted_at,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, menu)
	}

	return
}

func GetMenu(db *sql.DB, menu_id int) (results structs.Menu, err error) {
	sql := "SELECT * FROM menu WHERE menu_id = $1 AND deleted_at IS NULL LIMIT 1"

	err = db.QueryRow(sql, menu_id).Scan(
		&results.MenuId,
		&results.MenuName,
		&results.MenuDescription,
		&results.MenuPrice,
		&results.Created_at,
		&results.Updated_at,
		&results.Deleted_at,
	)

	return
}

func GetAllMenu(db *sql.DB) (results []structs.Menu, err error) {
	sql := "SELECT * FROM menu WHERE deleted_at IS NULL"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var menu = structs.Menu{}

		err = rows.Scan(
			&menu.MenuId,
			&menu.MenuName,
			&menu.MenuDescription,
			&menu.MenuPrice,
			&menu.Created_at,
			&menu.Updated_at,
			&menu.Deleted_at,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, menu)
	}

	return
}

func GetMenuWithTrashed(db *sql.DB, menu_id int) (results structs.Menu, err error) {
	sql := "SELECT * FROM menu WHERE menu_id = $1 LIMIT 1"

	err = db.QueryRow(sql, menu_id).Scan(
		&results.MenuId,
		&results.MenuName,
		&results.MenuDescription,
		&results.MenuPrice,
		&results.Created_at,
		&results.Updated_at,
		&results.Deleted_at,
	)

	return
}

func GetAllMenuWithTrash(db *sql.DB) (results []structs.Menu, err error) {
	sql := "SELECT * FROM menu"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var menu = structs.Menu{}

		err = rows.Scan(
			&menu.MenuId,
			&menu.MenuName,
			&menu.MenuDescription,
			&menu.MenuPrice,
			&menu.Created_at,
			&menu.Updated_at,
			&menu.Deleted_at,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, menu)
	}

	return
}

func CreateNewMenu(db *sql.DB, menu structs.Menu) error {
	sql := "INSERT INTO menu (menu_name, menu_description, menu_price, created_at) VALUES($1, $2, $3, $4)"

	errs := db.QueryRow(sql,
		menu.MenuName,
		menu.MenuDescription,
		menu.MenuPrice,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	return errs.Err()
}

func UpdateMenu(db *sql.DB, menu structs.Menu) error {
	sql := "UPDATE menu SET menu_name = $1, menu_description = $2, menu_price = $3, updated_at = $4 WHERE menu_id = $5"

	errs := db.QueryRow(sql,
		menu.MenuName,
		menu.MenuDescription,
		menu.MenuPrice,
		time.Now().Format("2006-01-02 15:04:05"),
		menu.MenuId,
	)

	return errs.Err()
}

func DeleteMenu(db *sql.DB, menu structs.Menu) error {
	sql := "UPDATE menu SET deleted_at = $1 WHERE menu_id = $2"
	var deletedAt interface{}

	if menu.Deleted_at == nil {
		deletedAt = time.Now().Format("2006-01-02 15:04:05")
	} else {
		deletedAt = nil
	}

	errs := db.QueryRow(sql, deletedAt, menu.MenuId)
	return errs.Err()
}
