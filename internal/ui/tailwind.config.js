/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './views/components/*.templ',
    './views/layouts/*.templ',
    './views/pages/*.templ',
    './views/**/*.templ',
    './node_modules/flowbite/**/*.js',
  ],
  theme: {
    extend: {},
  },
  plugins: [require('flowbite/plugin')],
}
