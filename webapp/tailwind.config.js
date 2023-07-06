/** @type {import('tailwindcss').Config} */

const daisyui = require('daisyui')

module.exports = {
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],

  theme: {
    extend: {}
  },
  plugins: [require('daisyui')],
  daisyui: {
    themes: ['forest']
  }
}
