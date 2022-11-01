import { useCallback, useState } from 'react'
import Webcam from 'react-webcam'
import { CameraOptions, useFaceDetection } from 'react-use-face-detection'
import FaceDetection from '@mediapipe/face_detection'
import { Camera } from '@mediapipe/camera_utils'
import { Alert, AlertTitle, Button, Collapse, Grid, IconButton, Tooltip } from '@mui/material'
import { Close, Lightbulb, Person, PhotoCamera, Visibility } from '@mui/icons-material'

const CaptureImage = () => {
	const [open, setOpen] = useState<boolean | undefined>(true)
	const [imgSrc, setImgSrc] = useState<string | null>(null)

	const { webcamRef, boundingBox, isLoading, detected, facesDetected } = useFaceDetection({
		faceDetectionOptions: {
			model: 'short',
		},
		faceDetection: new FaceDetection.FaceDetection({
			locateFile: (file) => `https://cdn.jsdelivr.net/npm/@mediapipe/face_detection/${file}`,
		}),
		camera: ({ mediaSrc, onFrame, width, height }: CameraOptions) =>
			new Camera(mediaSrc, {
				onFrame,
				width,
				height,
			}),
	})

	const capture = useCallback(() => {
		// @ts-ignore
		const imageSrc = webcamRef?.current.getUserMedia()
		setImgSrc(imageSrc)

		sessionStorage.setItem('img', JSON.stringify(imageSrc))
	}, [webcamRef, setImgSrc])

	return (
		<>
			<Collapse in={open}>
				<Alert
					severity='info'
					action={
						<Tooltip title='Close'>
							<IconButton
								aria-label='close'
								color='inherit'
								size='small'
								onClick={() => {
									setOpen(false)
								}}>
								<Close fontSize='inherit' />
							</IconButton>
						</Tooltip>
					}
					sx={{ mb: 2, textAlign: 'left' }}>
					<AlertTitle>INFO</AlertTitle>
					MAKE SURE YOU: <br />
					<Button
						size='small'
						color='info'
						startIcon={<Lightbulb />}
						sx={{ cursor: 'help', '&:hover': { backgroundColor: 'transparent' } }}
						disableRipple>
						{' '}
						Have good lighting.
					</Button>
					<br />
					<Button
						size='small'
						color='info'
						startIcon={<Visibility />}
						sx={{ cursor: 'help', '&:hover': { backgroundColor: 'transparent' } }}
						disableRipple>
						{' '}
						Look directly into the camera.
					</Button>
					<br />
					<Button
						size='small'
						color='info'
						startIcon={<Person />}
						sx={{ cursor: 'help', '&:hover': { backgroundColor: 'transparent' } }}
						disableRipple>
						{' '}
						Capture a head to shoulders close up.
					</Button>
				</Alert>
			</Collapse>
			<Grid container spacing={1}>
				<Grid item xs={12} md={6}>
					<Webcam
						ref={webcamRef}
						// forceScreenshotSourceSize
						// screenshotFormat='image/jpeg'
						// screenshotQuality={1}
						// height={300}
						mirrored
						forceScreenshotSourceSize
						screenshotFormat='image/jpeg'
						// videoConstraints={{
						// 	height: 720,
						// 	width: 1280,
						// }}
						height='180'
						width='320'
						style={{ border: detected ? '5px solid green' : '5px solid red' }}
					/>
				</Grid>
				<Grid item xs={12} md={6}>
					{imgSrc && <img src={imgSrc} height={150} width={220} alt='You' />}
				</Grid>
			</Grid>
			{/* <div style={{ width: 132, height: 170, borderTopRightRadius: 100, textAlign: 'center' }}> */}
			{/* {boundingBox.map((box, index) => 
					// <div
					// 	key={`${index + 1}`}
					// 	style={{
					// 		border: '4px solid green',
					// 		borderRadius: '80px',
					// 		position: 'absolute',
					// 		top: `${box.yCenter * 95}%`,
					// 		left: `${box.xCenter * 145}%`,
					// 		width: `${box.width * 30}%`,
					// 		height: `${box.height * 60}%`,
					// 		zIndex: 1,
					// 	}}
					// />
				// )}*/}

			{/* </div> */}

			{detected && facesDetected === 1 ? (
				<>
					<Tooltip title='Capture'>
						<IconButton
							color='primary'
							aria-label='capture picture'
							component='label'
							onClick={capture}
							sx={{
								border: '0.05rem solid',
								'&:hover': { backgroundColor: '#FAA61A', border: '0.05rem solid #fff', color: '#fff' },
								mt: 5,
							}}>
							<PhotoCamera />
						</IconButton>
					</Tooltip>
				</>
			) : (
				<>
					<IconButton
						color='primary'
						disabled
						sx={{
							border: '0.05rem solid',
							mt: 5,
						}}>
						<PhotoCamera />
					</IconButton>
				</>
			)}
		</>
	)
}

export default CaptureImage
