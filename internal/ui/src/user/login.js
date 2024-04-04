export default () => {
  return {
    email: '',
    password: '',
    submit() {
      console.log('submit called')
      this.callApi()
    },
    callApi() {
      window.htmx
        .ajax('POST', '/user/login', {
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
