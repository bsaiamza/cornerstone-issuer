import QRCode from 'react-qr-code'
import { useTheme } from '@mui/material/styles'

import { LIGHT_MODE_THEME } from '../../utils/constants'

const ScanQR = () => {
	const theme = useTheme()
	const credential = sessionStorage.getItem('credential')

	return (
		<div
			style={{
				backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '' : '#F5F5F5',
				padding: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
				borderRadius: theme.palette.mode === LIGHT_MODE_THEME ? '' : 5,
			}}>
			{/*@ts-ignore*/}
			{credential ? <QRCode value={credential} /> : 'No QR code to scan!'}
		</div>
	)
}

export default ScanQR
