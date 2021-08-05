const colors = require('tailwindcss/colors')

module.exports = {
  purge: ['./pages/**/*.{js,ts,jsx,tsx}', './components/**/*.{js,ts,jsx,tsx}'],
  darkMode: 'media',
  theme: {
    colors: {
      transparent: 'transparent',
      current: 'currentColor',
      primary: colors.red,
      secondary: colors.teal,
      grey: colors.warmGray,
      error: colors.red,
      black: colors.black,
      white: colors.white
    },
  },
  variants: {
    extend: {
      placeholderColor: ['hover'],
      ringWidth: ['hover']
    },
  },
  plugins: [],
}
