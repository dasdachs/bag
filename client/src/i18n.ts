import { initReactI18next } from "react-i18next"
import * as i18n from "i18next"
import LanguageDetector from "i18next-browser-languagedetector"
import Backend from "i18next-http-backend"

i18n.use(Backend)
	.use(LanguageDetector)
	.use(initReactI18next)
	.init({
		fallbackLng: "en",
		debug: true,
		interpolation: {
			escapeValue: false,
		},
		saveMissing: true,
	})
	.catch((e) => console.error(e))

export default i18n
