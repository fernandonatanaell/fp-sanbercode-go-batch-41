package structs

import "time"

type Menu struct {
	MenuId          int       `json:"menu_id"`
	MenuName        string    `json:"menu_name"`
	MenuDescription string    `json:"menu_description"`
	MenuPrice       int       `json:"menu_price"`
	Created_at      *time.Time `json:"created_at"`
	Updated_at      *time.Time `json:"updated_at"`
	Deleted_at      *time.Time `json:"deleted_at"`
}
