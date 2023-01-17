package structs

type Dtrans struct {
	OrderId        string `json:"order_id"`
	MenuId         int    `json:"menu_id"`
	DtransQuantity int    `json:"dtrans_quantity"`
	DtransTotal    string `json:"dtrans_total"`
}
