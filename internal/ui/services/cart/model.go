package cart_service

type Cart struct {
	CartId    string     `json:"cart_id,omitempty"`
	UserId    string     `json:"user_id,omitempty"`
	CartItems []CartItem `json:"cart_items,omitempty"`
}

type CartItem struct {
	ProductId string `json:"product_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}
