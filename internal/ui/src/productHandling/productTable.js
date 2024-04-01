export default (products) => {
  return {
    search: '',
    productsInSearch: [],
    editName: '',
    editPrice: '',
    editColor: '',
    editCategory: '',
    currentEditID: '',

    init() {
      this.productsInSearch = products
    },
    printProductsToConsole() {
      console.log(products)
    },
    findProductById(productId) {
      return products.find((product) => product.productId === productId)
    },
    get getProducts() {
      const filterItems =
        this.productsInSearch &&
        this.productsInSearch.filter((item) => {
          return item.productName
            .toLowerCase()
            .startsWith(this.search.toLowerCase())
        })

      if (
        filterItems.length < this.productsInSearch.length &&
        filterItems.length > 0
      ) {
        this.productsInSearch = filterItems
      } else {
        if (this.search) {
          this.productsInSearch = []
        } else {
          this.productsInSearch = products
        }
      }

      return this.productsInSearch
    },
    clearSearch() {
      this.search = ''
      this.productsInSearch = products
    },
    editProduct(productId) {
      this.currentEditID = productId
      console.log(productId, 'editProduct')

      window.htmx.ajax('GET', `/products/edit/${productId}`, {
        target: `#${productId}`,
        swap: 'outerHTML',
      })
    },
    saveProduct(productId) {
      const pr = this.findProductById(productId)
      this.editCategory = pr.productCategory
      this.editColor = pr.productColor
      this.editName = pr.productName
      this.editPrice = pr.productPrice
      window.htmx
        .ajax('POST', `/products/update/${productId}`, {
          values: {
            productCategory: this.editCategory,
            productColor: this.editColor,
            productName: this.editName,
            productPrice: this.editPrice,
            productId: productId,
          },
          target: `#${productId}`,
          swap: 'outerHTML',
        })
        .finally(() => {
          this.editCategory = ''
          this.editColor = ''
          this.editName = ''
          this.editPrice = ''
        })
    },
  }
}
