package product_service

type Product struct {
	Id          string  `json:id,omitempty`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
