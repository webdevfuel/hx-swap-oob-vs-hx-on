/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./template/**/*.templ"],
  theme: {
    extend: {
      colors: {
        background: "#151515",
        border: "#525252",
        primary: "#1d49f7",
        muted: "#222222",
      },
    },
  },
  plugins: [require("@tailwindcss/forms")],
};
