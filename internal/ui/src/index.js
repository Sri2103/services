/* eslint-disable no-unused-vars */
import 'flowbite'
import Alpine from 'alpinejs'
import focus from '@alpinejs/focus'
import addProduct from './productHandling/addProduct'
import productTable from './productHandling/productTable'
import Swal from 'sweetalert2'
// import 'htmx.org'
window.Alpine = Alpine
window.Swal = Swal
Alpine.data('productTable', productTable)
Alpine.data('addProduct', addProduct)

window.htmx = require('htmx.org')

Alpine.plugin(focus)
Alpine.start()
