export default () => {
  return {
    email: '',
    password: '',
    submit() {
      console.log('submit called')
      this.callApi()
    },
    callApi() {
      const apiUrl =
        window.location.pathname === '/user/login'
          ? '/user/login'
          : '/admin/login'
      window.htmx
        .ajax('POST', apiUrl, {
          values: {
            email: this.email,
            password: this.password,
          },
          swap: 'none',
        })
        .finally(() => {
          this.email = ''
          this.password = ''
        })
    },
  }
}
