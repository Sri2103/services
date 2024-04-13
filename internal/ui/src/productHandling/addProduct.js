import { uploadImages } from '../firebase'

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
     * @param {type} event - description of event parameter
     */
    async submit() {
      const currentElement = this.$el
      /**
       * Handle response error event.
       *
       * @param {type} event - description of event parameter
       */
      function handleResponseError(event) {
        if (event.target == currentElement && event.detail.xhr.status >= 399) {
          window.Swal.fire({
            title: 'Error',
            text: event.detail.xhr.responseText['message'],
            icon: 'error',
            confirmButtonText: 'OK',
          })
        }
      }
      /**
       * Handles the event after making a request.
       *
       * @param {Event} event - The event object
       */
      function handleAfterRequest(event) {
        if (event.target == currentElement && !event.detail.failed) {
          window.Swal.fire({
            title: 'Success',
            text: 'Product added successfully',
            icon: 'success',
            confirmButtonText: 'OK',
          })
        }
        if (event.target == currentElement) {
          document.removeEventListener(
            'htmx:responseError',
            handleResponseError
          )
        }
      }
      try {
        document.addEventListener('htmx:responseError', handleResponseError)
        document.addEventListener('htmx:afterRequest', handleAfterRequest, {
          once: true,
        })
        const images = document.getElementById('multiple_files')
        const files = images && images.files
        // Checks if there are any selected files in the input field

        console.log(files, 'files to upload')
        const str = files.length > 0 && (await uploadImages(files[0]))
        console.log(str, 'download string')
        if (str.length === 0) {
          window.Swal.fire({
            title: 'Error',
            text: 'Please select a file',
            icon: 'error',
            confirmButtonText: 'OK',
          })
          return
        }

        const colorElement = document.getElementById('color')
        const colorValue = colorElement.value

        const imageLinks = document.getElementById('image-links').value
        let newImageLinks = ''
        if (!!imageLinks && imageLinks.length > 0) {
          newImageLinks = imageLinks + ',' + str
        } else {
          newImageLinks = str
        }
        console.log(newImageLinks, 'new image links')

        await window.htmx.ajax('POST', '/products/add', {
          values: {
            name: this.name,
            price: this.price,
            category: this.category,
            description: this.description,
            image: [str],
            color: colorValue,
          },
          source: this.$el,
          swap: 'none',
        })
      } catch (error) {
        console.log(error, 'error from the ajax add product')
        window.Swal.fire({
          title: 'Error',
          text: error,
          icon: 'error',
          confirmButtonText: 'OK',
        })
      }
      this.hide()
    },
  }
}
