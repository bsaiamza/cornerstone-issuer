import { useEffect, useState } from 'react'
import { Divider } from '@mui/material'
import QRCode from 'react-qr-code'
import axios from 'axios'
// components
import BoxComponent from '../../components/Box'
// import CardComponent from '../../components/Card'
import GridComponent from '../../components/Grid'
import TypographyComponent from '../../components/Typography'
// images
import bg from '../../assets/images/bg1.png'
import scan from '../../assets/images/scan.png'

const HomePage = () => {
  const [url, setURL] = useState('')
  // const [object, setObject] = useState({})

  useEffect(() => {
    let apiURL = '/api/v1/cornerstone/issuer/connection/multi-invitation'

    if (process.env.API_BASE_URL) {
      apiURL =
        process.env.API_BASE_URL +
        '/cornerstone/issuer/connection/multi-invitation'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      axios
        .get(process.env.REACT_APP_API_URL + 'connection/multi-invitation')
        .then(response => {
          console.log(response.data)
          setURL(response.data.invitationURL)
          // setObject(response.data.invitation)
        })
        .catch(error => console.log(error))
    } else {
      axios
        .get(apiURL)
        .then(response => {
          console.log(response.data)
          setURL(response.data.invitationURL)
          // setObject(response.data.invitation)
        })
        .catch(error => console.log(error))
    }
  }, [])

  // const invitation = JSON.stringify(object)

  return (
    <BoxComponent sx={{ m: '2rem' }}>
      <GridComponent container spacing={2}>
        <GridComponent item xs={12} md={6}>
          <BoxComponent sx={{ textAlign: 'center' }}>
            <TypographyComponent variant="h2" color="Grey">
              Iamza Cornerstone Issuer
            </TypographyComponent>
            <TypographyComponent
              variant="h5"
              color="Grey"
              sx={{ marginBottom: '1rem' }}
            >
              Identity credentials made reliable and easy.
            </TypographyComponent>
            <Divider />
            <TypographyComponent
              variant="h6"
              color="Grey"
              sx={{ marginTop: '1rem', marginBottom: '1rem' }}
            >
              Connect with us to get your digital credentials now, scan the QR
              Code below:
            </TypographyComponent>
            <BoxComponent
              alt="DI"
              component="img"
              src={scan}
              sx={{
                height: '10rem',
                width: '10rem',
                opacity: 0.5,
              }}
            />
            <div>
              <QRCode value={url} />
            </div>
            {/* <CardComponent
              sx={{
                width: 'auto',
                p: '2rem',
                m: '1rem',
                overflowWrap: 'break-word',
              }}
            >
              <div className="json-text">{invitation}</div>
            </CardComponent> */}
          </BoxComponent>
        </GridComponent>
        <GridComponent item xs={12} md={6}>
          <BoxComponent
            alt="DI"
            component="img"
            src={bg}
            sx={{
              height: 'auto',
              width: '100%',
              opacity: 0.5,
            }}
          />
        </GridComponent>
      </GridComponent>
    </BoxComponent>
  )
}

export default HomePage
