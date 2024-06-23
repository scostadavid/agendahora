/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
               "./cmd/web/**/*.html", "./cmd/web/**/*.templ",
    ],
    theme: {
        extend: {
            animation: {
            fadeIn: 'fadeIn 0.5s ease-out forwards',
            fadeOut: 'fadeOut 0.5s ease-in forwards',
            },
            keyframes: {
            fadeIn: {
                '0%': { opacity: 0, transform: 'translateY(-10%)' },
                '100%': { opacity: 1, transform: 'translateY(0)' },
            },
            fadeOut: {
                '0%': { opacity: 1, transform: 'translateY(0)' },
                '100%': { opacity: 0, transform: 'translateY(-10%)' },
            },
            },
        },
    },
    plugins: [],
}

