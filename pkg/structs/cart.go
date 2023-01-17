package structs

type Cart struct {
	OrderId      string `json:"order_id"`
	MenuId       int `json:"menu_id"`
	CartQuantity int    `json:"cart_quantity"`
}
