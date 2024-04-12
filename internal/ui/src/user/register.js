export default () => {
  return {
    name: '',
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
    role: '',
    submit() {
      //
      console.log(this.name, this.email, this.password)
      console.log(window.location, 'register')
      //   this.callApi()
    },
    callApi() {
      window.htmx
        .ajax('POST', '/user/register', {
          values: {
            name: this.name,
            username: this.username,
            email: this.email,
            password: this.password,
            role: '',
          },
          swap: 'none',
        })
        .finally(() => {
          this.name = ''
          this.username = ''
          this.email = ''
          this.password = ''
        })
        .then((response) => {
          window.Swal.fire({
            title: 'Success',
            text: response.message,
            icon: 'success',
            confirmButtonText: 'OK',
          })
        })
    },
  }
}
