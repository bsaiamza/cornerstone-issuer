import { Box, Button, Divider, Grid, styled, Typography } from '@mui/material'
import { NavLink } from 'react-router-dom'

import bg from '../../assets/images/bg1.png'
import { SEO } from '../../components/SEO'

const Home = () => {
	return (
		<>
			<SEO title='Welcome' />

			<Grid container spacing={2}>
				<Grid item xs={12} md={6}>
					<Box sx={{ textAlign: 'center' }}>
						<Typography variant='h2' color='Grey'>
							IAMZA <br />
							Cornerstone Issuer
						</Typography>
						<Typography variant='h5' color='Grey' sx={{ marginBottom: '1rem' }}>
							Identity credentials made reliable and easy.
						</Typography>
						<Divider />
						<Button variant='contained' sx={{ m: '2rem' }}>
							<StyledNavLink to='get-my-credential'>Get started</StyledNavLink>
						</Button>
					</Box>
				</Grid>
				<Grid
					item
					xs={12}
					md={6}
					sx={{
						alignItems: 'center',
						justifyContent: 'center',
						textAlign: 'center',
					}}>
					<Box
						alt='DI'
						component='img'
						src={bg}
						sx={{
							height: 'auto',
							width: '100%',
							opacity: 0.5,
						}}
					/>
				</Grid>
			</Grid>
		</>
	)
}

const StyledNavLink = styled(NavLink)`
	text-decoration: none;
	color: inherit;
`

export default Home
