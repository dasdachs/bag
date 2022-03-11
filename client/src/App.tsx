import { Suspense } from "react"
import { BrowserRouter, Route, Routes } from "react-router-dom"
import { CircularProgress, ThemeProvider } from "@mui/material"
import CssBaseline from "@mui/material/CssBaseline"
import { theme } from "./theme"

function App() {
	return (
		<ThemeProvider theme={theme}>
			<Suspense fallback={<CircularProgress />}>
				<CssBaseline />
				<BrowserRouter>
					<Routes>
						<Route path="/" element={<div>Hello</div>} />
					</Routes>
				</BrowserRouter>
			</Suspense>
		</ThemeProvider>
	)
}

export default App
