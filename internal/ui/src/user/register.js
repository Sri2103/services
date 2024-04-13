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
      if (window.location.pathname === '/user/register') {
        this.callApi('user')
      } else {
        this.callApi('admin')
      }
    },
    callApi(role) {
      window.htmx
        .ajax('POST', `/${role}/register`, {
          values: {
            name: this.name,
            username: this.username,
            email: this.email,
            password: this.password,
            role: role,
          },
          swap: 'none',
        })
        .finally(() => {
          this.name = ''
          this.username = ''
          this.email = ''
          this.password = ''
          this.role = role
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
