package cart_service

type mockService struct {
}

// AddProductToCart implements Service.
func (m *mockService) AddProductToCart(userId string, productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

// AddProductToNonUserCart implements Service.
func (m *mockService) AddProductToNonUserCart(productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

// EmptyCart implements Service.
func (m *mockService) EmptyCart(userId string) (Cart, error) {
	panic("unimplemented")
}

// EmptyNonUserCart implements Service.
func (m *mockService) EmptyNonUserCart() (Cart, error) {
	panic("unimplemented")
}

// GetNonUserCart implements Service.
func (m *mockService) GetNonUserCart() (Cart, error) {
	panic("unimplemented")
}

// GetUserCart implements Service.
func (m *mockService) GetUserCart(userId string) (Cart, error) {
	panic("unimplemented")
}

// RemoveProductFromCart implements Service.
func (m *mockService) RemoveProductFromCart(userId string, productId string) (Cart, error) {
	panic("unimplemented")
}

// RemoveProductFromNonUserCart implements Service.
func (m *mockService) RemoveProductFromNonUserCart(productId string) (Cart, error) {
	panic("unimplemented")
}

// UpdateNonUserQuantity implements Service.
func (m *mockService) UpdateNonUserQuantity(productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

// UpdateQuantity implements Service.
func (m *mockService) UpdateQuantity(userId string, productId string, quantity int) (Cart, error) {
	panic("unimplemented")
}

func NewMockService() Service {

	return &mockService{}
}
