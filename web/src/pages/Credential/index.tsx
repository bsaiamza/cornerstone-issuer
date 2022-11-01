// import { useState } from 'react'
// import {
// 	Alert,
// 	Box,
// 	Button,
// 	Grid,
// 	Paper,
// 	Step,
// 	Stepper,
// 	StepContent,
// 	StepLabel,
// 	styled,
// 	Typography,
// } from '@mui/material'
// import { NavLink } from 'react-router-dom'

// import { SEO } from '../../components/SEO'
// import CaptureImage from './CaptureImage'
// import EnterDetails from './EnterDetails'
// import ScanQR from './ScanQR'

// const steps = [
// 	{
// 		label: 'Capture Image',
// 		description: `In order to begin the process of receiving your cornerstone credential,
//                   we need to capture an image of you to perform a liveliness test and store
//                   the image in your credential.`,
// 	},
// 	{
// 		label: 'Enter Details',
// 		description: 'Please capture your details to store in your credential.',
// 	},
// 	{
// 		label: 'Scan QR Code',
// 		description: 'Please scan the QR code with your mobile wallet app to receive your Cornerstone Credential.',
// 	},
// ]

// const Credential = () => {
// 	const [activeStep, setActiveStep] = useState(0)

// 	const handleNext = () => {
// 		setActiveStep((prevActiveStep) => prevActiveStep + 1)
// 	}

// 	const handleBack = () => {
// 		setActiveStep((prevActiveStep) => prevActiveStep - 1)
// 	}

// 	const handleReset = () => {
// 		setActiveStep(0)
// 	}

// 	return (
// 		<>
// 			<SEO title='Credential' />

// 			<Grid container spacing={1}>
// 				<Grid item xs={12} md={4} sx={{ borderRight: { xs: 0, md: 1 }, borderBottom: { xs: 1, md: 0 } }}>
// 					<Box sx={{ textAlign: 'left', p: 1 }}>
// 						<Stepper activeStep={activeStep} orientation='vertical'>
// 							{steps.map((step, index) => (
// 								<Step key={step.label}>
// 									{/* <StepLabel optional={index === 2 ? <Typography variant='caption'>Last step</Typography> : null}> */}
// 									<StepLabel>{step.label}</StepLabel>
// 									<StepContent>
// 										<Typography>{step.description}</Typography>
// 										<Box sx={{ mb: 2 }}>
// 											<div>
// 												{index === steps.length - 3 ? (
// 													<Button variant='contained' size='small' onClick={handleNext} sx={{ mt: 1, mr: 1 }}>
// 														Continue
// 													</Button>
// 												) : (
// 													''
// 												)}

// 												{index === steps.length - 2 || index === steps.length - 1 ? (
// 													<Button variant='outlined' size='small' onClick={handleBack} sx={{ mt: 1, mr: 1 }}>
// 														Back
// 													</Button>
// 												) : (
// 													''
// 												)}
// 											</div>
// 										</Box>
// 									</StepContent>
// 								</Step>
// 							))}
// 						</Stepper>
// 					</Box>
// 					{activeStep === steps.length - 1 && (
// 						<>
// 							<Alert severity='info' sx={{ m: 2, textAlign: 'left' }}>
// 								<Button variant='outlined' size='small' onClick={handleReset} color='info' sx={{ mt: 1, mr: 1 }}>
// 									Reset
// 								</Button>
// 								-- OR --{' '}
// 								<Button variant='outlined' size='small' color='info' sx={{ mt: 1, mr: 1 }}>
// 									<StyledNavLink to='/'>Return Home</StyledNavLink>
// 								</Button>
// 							</Alert>
// 						</>
// 					)}
// 				</Grid>
// 				<Grid
// 					item
// 					xs={12}
// 					md={8}
// 					sx={{
// 						alignItems: 'center',
// 						justifyContent: 'center',
// 						textAlign: 'center',
// 					}}
// 					justifyContent='center'
// 					justifyItems='center'>
// 					{activeStep === steps.length - 3 && <CaptureImage />}
// 					{activeStep === steps.length - 2 && <EnterDetails handleNext={handleNext} />}
// 					{activeStep === steps.length - 1 && <ScanQR />}
// 				</Grid>
// 			</Grid>
// 		</>
// 	)
// }

// const StyledNavLink = styled(NavLink)`
// 	text-decoration: none;
// 	color: inherit;
// `

// export default Credential

import { useState } from 'react'
import {
	Alert,
	Box,
	Button,
	Grid,
	Paper,
	Step,
	Stepper,
	StepContent,
	StepLabel,
	styled,
	Typography,
} from '@mui/material'
import { NavLink } from 'react-router-dom'

import { SEO } from '../../components/SEO'
import EnterDetails from './EnterDetails'
import ScanQR from './ScanQR'

const steps = [
	{
		label: 'Enter Details',
		description: 'Please capture your details that will be stored in your credential.',
	},
	{
		label: 'Scan QR Code',
		description: 'Please scan the QR code with your mobile wallet app to receive your Cornerstone Credential.',
	},
]

const Credential = () => {
	const [activeStep, setActiveStep] = useState(0)

	const handleNext = () => {
		setActiveStep((prevActiveStep) => prevActiveStep + 1)
	}

	const handleBack = () => {
		setActiveStep((prevActiveStep) => prevActiveStep - 1)
	}

	const handleReset = () => {
		setActiveStep(0)
	}

	return (
		<>
			<SEO title='Credential' />

			<Grid container spacing={1}>
				<Grid item xs={12} md={4} sx={{ borderRight: { xs: 0, md: 1 }, borderBottom: { xs: 1, md: 0 } }}>
					<Box sx={{ textAlign: 'left', p: 1 }}>
						<Stepper activeStep={activeStep} orientation='vertical'>
							{steps.map((step, index) => (
								<Step key={step.label}>
									{/* <StepLabel optional={index === 2 ? <Typography variant='caption'>Last step</Typography> : null}> */}
									<StepLabel>{step.label}</StepLabel>
									<StepContent>
										<Typography>{step.description}</Typography>
										<Box sx={{ mb: 2 }}>
											<div>
												{index === steps.length - 3 ? (
													<Button variant='contained' size='small' onClick={handleNext} sx={{ mt: 1, mr: 1 }}>
														Continue
													</Button>
												) : (
													''
												)}

												{index === steps.length - 1 ? (
													<Button variant='outlined' size='small' onClick={handleBack} sx={{ mt: 1, mr: 1 }}>
														Back
													</Button>
												) : (
													''
												)}
											</div>
										</Box>
									</StepContent>
								</Step>
							))}
						</Stepper>
					</Box>
					{activeStep === steps.length - 1 && (
						<>
							<Alert severity='info' sx={{ m: 2, textAlign: 'left' }}>
								<Button variant='outlined' size='small' color='info' sx={{ mt: 1, mr: 1 }}>
									<StyledNavLink to='/'>Return Home</StyledNavLink>
								</Button>
							</Alert>
						</>
					)}
				</Grid>
				<Grid
					item
					xs={12}
					md={8}
					sx={{
						alignItems: 'center',
						justifyContent: 'center',
						textAlign: 'center',
					}}
					justifyContent='center'
					justifyItems='center'>
					{activeStep === steps.length - 2 && <EnterDetails handleNext={handleNext} />}
					{activeStep === steps.length - 1 && <ScanQR />}
				</Grid>
			</Grid>
		</>
	)
}

const StyledNavLink = styled(NavLink)`
	text-decoration: none;
	color: inherit;
`

export default Credential
