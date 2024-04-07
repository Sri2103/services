export default () => {
  return {
    items: [
      {
        name: 'Product A',
        price: 99.99,
        quantity: 1,
        image: 'https://via.placeholder.com/50',
      },
      {
        name: 'Product B',
        price: 49.99,
        quantity: 1,
        image: 'https://via.placeholder.com/50',
      },
      {
        name: 'Product C',
        price: 29.99,
        quantity: 1,
        image: 'https://via.placeholder.com/50',
      },
    ],
    taxRate: 0.1,
    coupon: null,
    applyCoupon() {
      if (this.coupon === 'DISCOUNT10') {
        // Apply 10% discount
        return 0.1 // 10% discount
      } else {
        // No discount
        return 0
      }
    },
    subtotal() {
      return this.items.reduce(
        (acc, item) => acc + item.price * item.quantity,
        0
      )
    },
    tax() {
      return this.subtotal() * this.taxRate
    },
    discount() {
      // Calculate discount based on coupon
      return this.applyCoupon() * this.subtotal()
    },
    total() {
      return this.subtotal() + this.tax() - this.discount()
    },
    removeItem(index) {
      this.items.splice(index, 1)
    },
  }
}
