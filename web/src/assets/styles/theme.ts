import { createTheme, responsiveFontSizes } from '@mui/material'

// constants
import { DARK_MODE_THEME, LIGHT_MODE_THEME } from '../../utils/constants'

export const getAppTheme = (mode: typeof LIGHT_MODE_THEME | typeof DARK_MODE_THEME) => {
	let theme = createTheme({
		palette: {
			mode,
			...(mode === 'light'
				? {
						// palette values for light mode
						primary: {
							light: '#F5F5F5',
							main: '#FAA61A',
							dark: '#4D4D4D',
							contrastText: '#ffffff',
						},
						divider: '#333333',
						text: {
							primary: '#333333',
							secondary: '#333333',
						},
						background: {
							default: '#fff',
							paper: '#F5F5F5',
						},
				  }
				: {
						// palette values for dark mode
						primary: {
							light: '#F5F5F5',
							main: '#FAA61A',
							dark: '#4D4D4D',
							contrastText: '#ffffff',
						},
						divider: '#F5F5F5',
						background: {
							default: '#08131B',
							paper: '#08131B',
						},
						text: {
							primary: '#ffffff',
							secondary: '#cccccc',
						},
				  }),
		},
	})
	theme = responsiveFontSizes(theme)
	return theme
}
