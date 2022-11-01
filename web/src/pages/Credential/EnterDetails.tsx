// import { useState } from 'react'
// import {
// 	Alert,
// 	AlertTitle,
// 	Button,
// 	Collapse,
// 	Grid,
// 	IconButton,
// 	Paper,
// 	TextField,
// 	Tooltip,
// 	useTheme,
// } from '@mui/material'
// import { Close } from '@mui/icons-material'
// import { Formik, Form } from 'formik'
// import { toast } from 'react-toastify'
// import axios from 'axios'
// import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
// import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider'
// import { DatePicker } from '@mui/x-date-pickers/DatePicker'
// import Moment from 'moment'

// import { CORNERSTONE_ISSUER_URL, LIGHT_MODE_THEME } from '../../utils/constants'

// interface EnterDetailsProps {
// 	handleNext: () => void
// }

// const apiURL = CORNERSTONE_ISSUER_URL + '/credential'

// const validSouthAfricanId = require('valid-south-african-id')

// const idValidation = (data: any) => {
// 	const errors = {}

// 	// eslint-disable-next-line no-empty
// 	if (validSouthAfricanId(data.identity_number)) {
// 	} else {
// 		// @ts-ignore
// 		errors.identity_number = 'Invalid ZA ID number!'
// 	}

// 	return errors
// }

// const EnterDetails = ({ handleNext }: EnterDetailsProps) => {
// 	const theme = useTheme()

// 	const [submitting, setSubmitting] = useState(false)
// 	const [open, setOpen] = useState(true)

// 	// @ts-ignore
// 	const idPhoto = JSON.parse(sessionStorage.getItem('img'))

// 	const formatDate = Moment().format('YYYY-MM-DD')

// 	const sendOffer = async (values: any) => {
// 		setSubmitting(true)

// 		await toast.promise(
// 			axios
// 				.post(apiURL, values)
// 				.then((response: any) => {
// 					// @ts-ignore
// 					sessionStorage.setItem('credential', response.data.credential)
// 					toast.success('Credential generated!')
// 					setTimeout(() => {
// 						handleNext()
// 					}, 3000)
// 				})
// 				.catch((error: any) => {
// 					toast.error(error.response.data.msg)
// 				}),
// 			{
// 				pending: 'Generating credential...',
// 			}
// 		)
// 		setSubmitting(false)
// 	}

// 	return (
// 		<Grid container>
// 			<Grid item xs={0} md={2} />
// 			<Grid item xs={12} md={8}>
// 				{!idPhoto ? (
// 					<Collapse in={open}>
// 						<Alert
// 							severity='error'
// 							action={
// 								<Tooltip title='Close'>
// 									<IconButton
// 										aria-label='close'
// 										color='inherit'
// 										size='small'
// 										onClick={() => {
// 											setOpen(false)
// 										}}>
// 										<Close fontSize='inherit' />
// 									</IconButton>
// 								</Tooltip>
// 							}
// 							sx={{ mb: 2, textAlign: 'left' }}>
// 							<AlertTitle>Please go back and take a picture to continue!</AlertTitle>
// 						</Alert>
// 					</Collapse>
// 				) : (
// 					''
// 				)}

// 				<Paper
// 					square
// 					elevation={2}
// 					sx={{ p: 3, width: { md: '100%' }, backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '#fff' : '' }}>
// 					<Formik
// 						initialValues={{
// 							identity_number: '',
// 							names: '',
// 							surname: '',
// 							gender: '',
// 							date_of_birth: '',
// 							country_of_birth: '',
// 							nationality: '',
// 							status: '',
// 							date_of_issue: formatDate,
// 							identity_photo: idPhoto,
// 						}}
// 						validate={idValidation}
// 						onSubmit={(values, { resetForm }) => {
// 							sendOffer(values)
// 							// resetForm()
// 						}}>
// 						{({ values, handleChange, touched, errors, setFieldValue }) => (
// 							<Form>
// 								<div>
// 									<TextField
// 										error={touched.identity_number && Boolean(errors.identity_number)}
// 										helperText={touched.identity_number && errors.identity_number}
// 										id='identity_number'
// 										name='identity_number'
// 										value={values.identity_number}
// 										onChange={handleChange}
// 										label='ID Number'
// 										sx={{ m: '1rem' }}
// 										required
// 										disabled={!idPhoto}
// 									/>

// 									<TextField
// 										id='names'
// 										name='names'
// 										value={values.names}
// 										onChange={handleChange}
// 										label='Names'
// 										sx={{ m: '1rem' }}
// 										required
// 										disabled={!idPhoto}
// 									/>
// 								</div>

// 								<div>
// 									<TextField
// 										id='surname'
// 										name='surname'
// 										value={values.surname}
// 										onChange={handleChange}
// 										label='Surname'
// 										sx={{ m: '1rem' }}
// 										required
// 										disabled={!idPhoto}
// 									/>

// 									<TextField
// 										id='gender'
// 										name='gender'
// 										// @ts-ignore
// 										value={
// 											(values.gender = values.identity_number
// 												? // @ts-ignore
// 												  values.identity_number.substring(6, 7) > 4
// 													? 'Male'
// 													: 'Female'
// 												: '')
// 										}
// 										onChange={handleChange}
// 										label='Gender'
// 										sx={{ m: '1rem' }}
// 										required
// 										disabled={!idPhoto}
// 									/>
// 								</div>

// 								<div>
// 									<TextField
// 										id='date_of_birth'
// 										name='date_of_birth'
// 										value={
// 											// @ts-ignore
// 											(values.date_of_birth =
// 												// @ts-ignore
// 												values.identity_number
// 													? // @ts-ignore
// 													  values.identity_number.substring(0, 1) > 2
// 														? '19' +
// 														  values.identity_number.substring(0, 2) +
// 														  '-' +
// 														  values.identity_number.substring(2, 4) +
// 														  '-' +
// 														  values.identity_number.substring(4, 6)
// 														: '20' +
// 														  values.identity_number.substring(0, 2) +
// 														  '-' +
// 														  values.identity_number.substring(2, 4) +
// 														  '-' +
// 														  values.identity_number.substring(4, 6)
// 													: '')
// 										}
// 										onChange={handleChange}
// 										label='D.O.B'
// 										sx={{ m: '1rem' }}
// 										required
// 										disabled={!idPhoto}
// 									/>

// 									<TextField
// 										id='country_of_birth'
// 										name='country_of_birth'
// 										value={
// 											(values.country_of_birth = values.identity_number.substring(10, 11) === '0' ? 'South Africa' : '')
// 										}
// 										onChange={handleChange}
// 										label='Country of Birth'
// 										sx={{ m: '1rem' }}
// 										required
// 										// disabled={values.identity_number.substring(10, 11) === '0' || !idPhoto}
// 										disabled={!idPhoto}
// 									/>
// 								</div>

// 								<div>
// 									<TextField
// 										id='nationality'
// 										name='nationality'
// 										value={
// 											(values.nationality = values.identity_number.substring(10, 11) === '0' ? 'South African' : '')
// 										}
// 										onChange={handleChange}
// 										label='Nationality'
// 										sx={{ m: '1rem' }}
// 										required
// 										// disabled={values.identity_number.substring(10, 11) === '0' || !idPhoto}
// 										disabled={!idPhoto}
// 									/>

// 									<TextField
// 										id='status'
// 										name='status'
// 										value={(values.status = values.identity_number.substring(10, 11) === '0' ? 'Citizen' : '')}
// 										onChange={handleChange}
// 										label='Status'
// 										sx={{ m: '1rem' }}
// 										required
// 										// disabled={values.identity_number.substring(10, 11) === '0' || !idPhoto}
// 										disabled={!idPhoto}
// 									/>
// 								</div>

// 								<div>
// 									<LocalizationProvider dateAdapter={AdapterDayjs}>
// 										<DatePicker
// 											label='Date of Issue'
// 											value={values.date_of_issue}
// 											onChange={(value) => setFieldValue('date_of_issue', value, true)}
// 											renderInput={(params) => (
// 												<TextField
// 													id='date_of_issue'
// 													name='date_of_issue'
// 													required
// 													{...params}
// 													sx={{ m: '1rem' }}
// 													disabled={!idPhoto}
// 												/>
// 											)}
// 										/>
// 									</LocalizationProvider>
// 								</div>

// 								<div>
// 									<Button
// 										variant='contained'
// 										size='small'
// 										type='submit'
// 										sx={{ color: '#fff', m: '1rem' }}
// 										disabled={submitting}>
// 										Submit
// 									</Button>
// 								</div>
// 							</Form>
// 						)}
// 					</Formik>
// 				</Paper>
// 			</Grid>
// 			<Grid item xs={0} md={2} />
// 		</Grid>
// 	)
// }

// export default EnterDetails

import { useState } from 'react'
import { Alert, Button, Grid, Paper, TextField, useTheme } from '@mui/material'
import { Formik, Form } from 'formik'
import { toast } from 'react-toastify'
import axios from 'axios'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider'
import { DatePicker } from '@mui/x-date-pickers/DatePicker'

import { CORNERSTONE_ISSUER_URL, LIGHT_MODE_THEME } from '../../utils/constants'

interface EnterDetailsProps {
	handleNext: () => void
}

const apiURL = CORNERSTONE_ISSUER_URL + '/credential'
const apiEmailURL = CORNERSTONE_ISSUER_URL + '/email-credential'

const validSouthAfricanId = require('valid-south-african-id')

const idValidation = (data: any) => {
	const errors = {}

	// eslint-disable-next-line no-empty
	if (validSouthAfricanId(data.identity_number)) {
	} else {
		// @ts-ignore
		errors.identity_number = 'Invalid ZA ID number!'
	}

	return errors
}

const EnterDetails = ({ handleNext }: EnterDetailsProps) => {
	const theme = useTheme()

	const [submitting, setSubmitting] = useState<boolean | undefined>(false)

	const sendOffer = async (values: any) => {
		setSubmitting(true)

		await toast.promise(
			axios
				.post(apiURL, values)
				.then((response: any) => {
					// @ts-ignore
					sessionStorage.setItem('credential', response.data.credential)
					toast.success('Credential generated!')
					setTimeout(() => {
						handleNext()
					}, 1000)
				})
				.catch((error: any) => {
					toast.error(error.response.data.msg)
				}),
			{
				pending: 'Generating credential...',
			}
		)
		setSubmitting(false)
	}

	const emailOffer = async (values: any) => {
		setSubmitting(true)

		await toast.promise(
			axios
				.post(apiEmailURL, values)
				.then((response: any) => {
					// @ts-ignore
					sessionStorage.setItem('credential', response.data.credential)
					toast.success('Emailed credential!')
				})
				.catch((error: any) => {
					toast.error(error.response.data.msg)
				}),
			{
				pending: 'Emailing credential...',
			}
		)
		setSubmitting(false)
	}

	return (
		<Grid container>
			<Grid item xs={0} md={2} />
			<Grid item xs={12} md={8}>
				<Paper
					square
					elevation={2}
					sx={{
						p: 3,
						width: { md: '100%' },
						backgroundColor: theme.palette.mode === LIGHT_MODE_THEME ? '#fff' : '',
						borderRadius: 5,
					}}>
					<Formik
						initialValues={{
							identity_number: '',
							names: '',
							surname: '',
							gender: '',
							date_of_birth: '',
							country_of_birth: '',
							nationality: '',
							citizen_status: '',
							email: '',
						}}
						validate={idValidation}
						onSubmit={(values, { resetForm }) => {
							values.email === '' ? sendOffer(values) : emailOffer(values)
							// resetForm()
						}}>
						{({ values, handleChange, touched, errors, setFieldValue }) => (
							<Form>
								<div>
									<TextField
										error={touched.identity_number && Boolean(errors.identity_number)}
										helperText={touched.identity_number && errors.identity_number}
										id='identity_number'
										name='identity_number'
										value={values.identity_number}
										onChange={handleChange}
										label='ID Number'
										sx={{ m: '1rem' }}
										required
									/>

									<TextField
										id='names'
										name='names'
										value={values.names}
										onChange={handleChange}
										label='Names'
										sx={{ m: '1rem' }}
										required
									/>
								</div>

								<div>
									<TextField
										id='surname'
										name='surname'
										value={values.surname}
										onChange={handleChange}
										label='Surname'
										sx={{ m: '1rem' }}
										required
									/>

									<TextField
										id='gender'
										name='gender'
										// @ts-ignore
										value={
											(values.gender = values.identity_number
												? // @ts-ignore
												  values.identity_number.substring(6, 7) > 4
													? 'Male'
													: 'Female'
												: '')
										}
										onChange={handleChange}
										label='Gender'
										sx={{ m: '1rem' }}
										required
										disabled
									/>
								</div>

								<div>
									<LocalizationProvider dateAdapter={AdapterDayjs}>
										<DatePicker
											label='Date of Birth'
											value={
												// @ts-ignore
												(values.date_of_birth =
													// @ts-ignore
													values.identity_number
														? // @ts-ignore
														  values.identity_number.substring(0, 1) > 2
															? '19' +
															  values.identity_number.substring(0, 2) +
															  '-' +
															  values.identity_number.substring(2, 4) +
															  '-' +
															  values.identity_number.substring(4, 6)
															: '20' +
															  values.identity_number.substring(0, 2) +
															  '-' +
															  values.identity_number.substring(2, 4) +
															  '-' +
															  values.identity_number.substring(4, 6)
														: '')
											}
											onChange={(value) => setFieldValue('date_of_birth', value, true)}
											renderInput={(params) => (
												<TextField
													id='date_of_birth'
													name='date_of_birth'
													{...params}
													sx={{ m: '1rem', width: '14.5rem' }}
													required
												/>
											)}
											disabled
										/>
									</LocalizationProvider>

									<TextField
										id='country_of_birth'
										name='country_of_birth'
										value={
											(values.country_of_birth = values.identity_number.substring(10, 11) === '0' ? 'RSA' : 'Other')
										}
										onChange={handleChange}
										label='Country of Birth'
										sx={{ m: '1rem' }}
										required
										disabled
									/>
								</div>

								<div>
									<TextField
										id='nationality'
										name='nationality'
										value={(values.nationality = values.identity_number.substring(10, 11) === '0' ? 'RSA' : 'Other')}
										onChange={handleChange}
										label='Nationality'
										sx={{ m: '1rem' }}
										required
										disabled
									/>

									<TextField
										id='citizen_status'
										name='citizen_status'
										value={
											(values.citizen_status =
												values.identity_number.substring(10, 11) === '0' ? 'Citizen' : 'Non-Citizen')
										}
										onChange={handleChange}
										label='Citizen Status'
										sx={{ m: '1rem' }}
										required
										disabled
									/>
								</div>

								<div>
									<TextField
										id='email'
										name='email'
										type='email'
										value={values.email}
										onChange={handleChange}
										label='Email'
										sx={{ m: '1rem' }}
										// helperText='Optional: If you want to receive your credential via email.'
										helperText={
											<Alert severity='info' sx={{ backgroundColor: 'transparent' }}>
												Optional: If you want to receive your credential via email.
											</Alert>
										}
									/>
								</div>

								<div>
									<Button
										variant='contained'
										size='small'
										type='submit'
										sx={{ color: '#fff', m: '1rem' }}
										disabled={submitting}>
										Submit
									</Button>
								</div>
							</Form>
						)}
					</Formik>
				</Paper>
			</Grid>
			<Grid item xs={0} md={2} />
		</Grid>
	)
}

export default EnterDetails
