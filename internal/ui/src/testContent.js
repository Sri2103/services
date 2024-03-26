import Alpine from 'alpinejs'
import swal from 'sweetalert'
Alpine.data('testContent', () => {
  return {
    email: '',
    name: '',
    message: '',
    open: false,
    alertMe() {
      swal('Hello world!')
    },
  }
})
