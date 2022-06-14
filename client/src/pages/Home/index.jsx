import { Divider } from '@mui/material'
// components
import BoxComponent from '../../components/Box'
import ButtonComponent from '../../components/Button'
import GridComponent from '../../components/Grid'
import TypographyComponent from '../../components/Typography'
// images
import bg from '../../assets/images/bg1.png'
import { Link } from 'react-router-dom'

const HomePage = () => {
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
          </BoxComponent>
          <ButtonComponent variant="contained" sx={{ m: '2rem' }}>
            <Link
              to="issue-credential"
              style={{
                color: 'white',
                textDecoration: 'none',
              }}
            >
              Get my Cornerstone Credential
            </Link>
          </ButtonComponent>
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
