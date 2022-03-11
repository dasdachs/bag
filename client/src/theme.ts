import { createTheme } from "@mui/material"

export const theme = createTheme({
	typography: {
		fontFamily: [
			'"Press Start 2P"',
			'"Segoe UI"',
			"Roboto",
			'"Helvetica Neue"',
			"sans-serif",
			"cursive",
			'"Segoe UI Emoji"',
			'"Segoe UI Symbol"',
		].join(","),
	},
	palette: {
		mode: "dark",
		text: {
			primary: "#ffeb3b",
			secondary: "#ffeb3b",
		},
		background: {
			default: "#4a148c",
			paper: "#311b92",
		},
	},
})
