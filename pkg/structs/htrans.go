package structs

import "time"

type Htrans struct {
	OrderId        string    `json:"order_id"`
	HtransSubtotal string    `json:"htrans_subtotal"`
	HtransStatus   int       `json:"htrans_status"`
	Created_at     *time.Time `json:"created_at"`
	Updated_at     *time.Time `json:"updated_at"`
}
