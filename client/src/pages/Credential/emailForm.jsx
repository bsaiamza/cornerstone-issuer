import { useState } from 'react'
import { Divider } from '@mui/material'
import { toast } from 'react-toastify'
import axios from 'axios'
import validate from 'za-id-validator'
// components
import BoxComponent from '../../components/Box'
import ButtonComponent from '../../components/Button'
import FormikComponent from '../../components/Formik'
import FormComponent from '../../components/Form'
import GridComponent from '../../components/Grid'
import TextFieldComponent from '../../components/TextField'
import TypographyComponent from '../../components/Typography'
// images
import bg from '../../assets/images/bg1.png'

const idValidation = data => {
  const errors = {}

  if (validate(data.id_number)) {
  } else {
    errors.id_number = 'Invalid ZA ID number!'
  }

  return errors
}

const GetCredentialEmailForm = () => {
  const [submitting, setSubmitting] = useState(false)

  const sendOffer = async data => {
    setSubmitting(true)
    let apiURL = '/api/v1/cornerstone/issuer/email-credential'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/email-credential'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      await toast.promise(
        axios
          .post(process.env.REACT_APP_API_URL + 'email-credential', data)
          .then(response => {
            toast.success('Sent credential request via email!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Sending email...',
        }
      )
    } else {
      await toast.promise(
        axios
          .post(apiURL, data)
          .then(response => {
            toast.success('Sent credential request via email!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Sending email...',
        }
      )
    }
    setSubmitting(false)
  }

  return (
    <>
      <TypographyComponent variant="h5">
        Get my Cornerstone Credential
      </TypographyComponent>
      <Divider />
      <div style={{ marginTop: '1rem' }}>
        <FormikComponent
          initialValues={{
            email: '',
            id_number: '',
            surname: '',
            forenames: '',
            gender: '',
            date_of_birth: '',
            country_of_birth: '',
          }}
          validate={idValidation}
          onSubmit={(values, { resetForm }) => {
            sendOffer(values)
            // resetForm()
          }}
        >
          {({ values, handleChange, touched, errors }) => (
            <FormComponent>
              <div>
                <TextFieldComponent
                  error={touched.id_number && Boolean(errors.id_number)}
                  helperText={touched.id_number && errors.id_number}
                  id="id_number"
                  name="id_number"
                  value={values.id_number}
                  onChange={handleChange}
                  label="ID Number"
                  sx={{ m: '1rem' }}
                  required
                />

                <TextFieldComponent
                  id="forenames"
                  name="forenames"
                  value={values.forenames}
                  onChange={handleChange}
                  label="Forenames"
                  sx={{ m: '1rem' }}
                  required
                />
              </div>

              <div>
                <TextFieldComponent
                  id="surname"
                  name="surname"
                  value={values.surname}
                  onChange={handleChange}
                  label="Surname"
                  sx={{ m: '1rem' }}
                  required
                />
                <TextFieldComponent
                  id="gender"
                  name="gender"
                  // value={values.gender}
                  value={
                    (values.gender =
                      values.id_number.substring(6, 7) > 4 ? 'Male' : 'Female')
                  }
                  onChange={handleChange}
                  label="Gender"
                  sx={{ m: '1rem' }}
                  required
                  disabled
                />
              </div>

              <div>
                <TextFieldComponent
                  id="date_of_birth"
                  name="date_of_birth"
                  // type="date"
                  // value={values.date_of_birth}
                  value={
                    (values.date_of_birth =
                      values.id_number.substring(0, 1) > 2
                        ? '19' +
                          values.id_number.substring(0, 2) +
                          '-' +
                          values.id_number.substring(2, 4) +
                          '-' +
                          values.id_number.substring(4, 6)
                        : '20' +
                          values.id_number.substring(0, 2) +
                          '-' +
                          values.id_number.substring(2, 4) +
                          '-' +
                          values.id_number.substring(4, 6))
                  }
                  onChange={handleChange}
                  label="D.O.B"
                  // sx={{ m: '1rem', width: '15rem' }}
                  sx={{ m: '1rem' }}
                  required
                  // InputLabelProps={{
                  //   shrink: true,
                  // }}
                  disabled
                />

                <TextFieldComponent
                  id="country_of_birth"
                  name="country_of_birth"
                  value={
                    (values.country_of_birth =
                      values.id_number.substring(10, 11) === '0'
                        ? 'South Africa'
                        : '')
                  }
                  onChange={handleChange}
                  label="Country of Birth"
                  sx={{ m: '1rem' }}
                  required
                  disabled={values.id_number.substring(10, 11) === '0'}
                />
              </div>

              <div>
                <TextFieldComponent
                  id="email"
                  name="email"
                  value={values.email}
                  onChange={handleChange}
                  label="Email"
                  sx={{ m: '1rem' }}
                  required
                />
              </div>

              <div>
                <ButtonComponent
                  variant="contained"
                  type="submit"
                  sx={{ color: '#fff', m: '1rem' }}
                  disabled={submitting}
                >
                  Email
                </ButtonComponent>
              </div>
            </FormComponent>
          )}
        </FormikComponent>
      </div>
    </>
  )
}

export default GetCredentialEmailForm
