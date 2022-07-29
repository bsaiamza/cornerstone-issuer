import { useState } from 'react'
import { Divider } from '@mui/material'
import { toast } from 'react-toastify'
import axios from 'axios'
import QRCode from 'react-qr-code'
import validate from 'za-id-validator'
// components
import ButtonComponent from '../../components/Button'
import FormikComponent from '../../components/Formik'
import FormComponent from '../../components/Form'
import TextFieldComponent from '../../components/TextField'
import TypographyComponent from '../../components/Typography'
// utils
import { CORNERSTONE_ISSUER_URL } from '../../utils'

const apiURL = CORNERSTONE_ISSUER_URL + '/credential'

const idValidation = data => {
  const errors = {}

  if (validate(data.id_number)) {
  } else {
    errors.id_number = 'Invalid ZA ID number!'
  }

  return errors
}

const GetCredentialForm = () => {
  const [submitting, setSubmitting] = useState(false)
  const [success, setSuccess] = useState(false)
  const [data, setData] = useState([])

  const sendOffer = async data => {
    setSubmitting(true)

    await toast.promise(
      axios
        .post(apiURL, data)
        .then(response => {
          setData(response.data)
          setSuccess(true)
          toast.success('Credential request generated!')
        })
        .catch(error => {
          toast.error(error.response.data.msg)
        }),
      {
        pending: 'Generating request...',
      }
    )

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
            id_number: '',
            first_names: '',
            surname: '',
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
                  id="first_names"
                  name="first_names"
                  value={values.first_names}
                  onChange={handleChange}
                  label="First Names"
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
                  sx={{ m: '1rem' }}
                  required
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
                <ButtonComponent
                  variant="contained"
                  type="submit"
                  sx={{ color: '#fff', m: '1rem' }}
                  disabled={submitting}
                >
                  Submit
                </ButtonComponent>
              </div>
            </FormComponent>
          )}
        </FormikComponent>
      </div>

      {success && data.credential ? <QRCode value={data.credential} /> : ''}
    </>
  )
}

export default GetCredentialForm
