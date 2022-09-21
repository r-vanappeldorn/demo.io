/** @type {import('tailwindcss').Config} */
const colors = require("tailwindcss/colors")

module.exports = {
  mode: "jit",
  content: [
    "./src/pages/*{ts,tsx,js,jsx}",
    "./node_modules/@demo.io/utils/dist/components/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    colors: {
      dark: {
        100: "#A9A9A9",
        600: "#3C3C3C",
        700: "#2D2D2D",
        800: "#1E1E1E",
        900: "#121212",
      },
      purpel: {
        400: "#BB86FC",
        500: "#6200EE",
        600: "#5a00d9",
        900: "#3A2D49",
      },
      ...colors,
    },
    extend: {},
  },
  plugins: [require("@tailwindcss/forms"), require("tw-elements/dist/plugin")],
}
