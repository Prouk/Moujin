/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/**/*.{js,css}", "./src/*.templ"],
  theme: {
    extend: {
      transitionProperty: {
        'max-height': 'max-height'
      },
      maxHeight: {
        '128': '32rem',
        '256': '64rem',
      },
      maxWidth: {
        '128': '32rem',
        '256': '64rem',
      }
    }
  },
  plugins: [],
}

