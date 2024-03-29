export default () => {
  // create data for the form field

  return {
    modalOpen: false,
    name: '',
    price: '',
    category: '',
    description: '',
    show() {
      this.modalOpen = true
    },
    hide() {
      this.name = ''
      this.price = ''
      this.category = ''
      this.description = ''
      this.modalOpen = false
    },
    /**
     * Submits the form data to the server using AJAX.
     */
    submit() {
      console.log('$el', this.$el)
      window.htmx
        .ajax('POST', '/products/add', {
          values: {
            name: this.name,
            price: this.price,
            category: this.category,
            description: this.description,
          },
          swap: 'none',
        })
        .then(() => {
          window.Swal.fire({
            title: 'Success',
            text: 'Product added successfully',
            icon: 'success',
            confirmButtonText: 'OK',
          })
        })
        .catch((error) => {
          window.Swal.fire({
            title: 'Error',
            text: error,
            icon: 'error',
            confirmButtonText: 'OK',
          })
        })
        .finally(() => this.hide())
    },
  }
}
