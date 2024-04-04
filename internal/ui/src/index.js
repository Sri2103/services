/* eslint-disable no-unused-vars */
import 'flowbite'
import Alpine from 'alpinejs'
import focus from '@alpinejs/focus'
import addProduct from './productHandling/addProduct'
import productTable from './productHandling/productTable'
import Swal from 'sweetalert2'
import editProduct from './productHandling/editProduct'
import carousel from './views/carousel'
import login from './user/login'
import register from './user/register'

// import 'htmx.org'
window.Alpine = Alpine
window.Swal = Swal
Alpine.data('productTable', productTable)
Alpine.data('addProduct', addProduct)
Alpine.data('editProduct', editProduct)
Alpine.data('carousel', carousel)
Alpine.data('login', login)
Alpine.data('register', register)

window.htmx = require('htmx.org')

Alpine.plugin(focus)
Alpine.start()
