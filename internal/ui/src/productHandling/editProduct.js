export default (product) => {
  return {
    modalOpen: false,
    editName: product.productName,
    editPrice: product.productPrice,
    editColor: product.productColor,
    editCategory: product.productCategory,
    currentEditID: product.productId,
    editDescription: product.productDescription,
    editImages: product.productImages || [],
    show() {
      this.modalOpen = true
    },
    hide() {
      this.modalOpen = false
    },
    addImageField() {
      this.editImages.push('')
    },
    removeImageField(index) {
      this.editImages.splice(index, 1)
    },
    submit() {
      window.htmx
        .ajax('POST', '/products/update/' + this.currentEditID, {
          values: {
            productCategory: this.editCategory,
            productColor: this.editColor,
            productName: this.editName,
            productPrice: this.editPrice,
            productId: this.currentEditID,
            productDescription: this.editDescription,
            productImages: this.editImages,
          },
          target: `#${this.currentEditID}`,
          swap: 'outerHTML',
        })
        .finally(() => {
          this.modalOpen = false
        })
    },
  }
}
