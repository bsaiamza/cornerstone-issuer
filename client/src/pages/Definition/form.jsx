import { useState } from 'react'
import { toast } from 'react-toastify'
import { FormControl, InputLabel, MenuItem, Select } from '@mui/material'
import axios from 'axios'
// components
import ButtonComponent from '../../components/Button'
import CardComponent from '../../components/Card'
import GridComponent from '../../components/Grid'
import FormikComponent from '../../components/Formik'
import FormComponent from '../../components/Form'
import ListItemComponent from '../../components/ListItem'
import ListItemTextComponent from '../../components/ListItemText'
import TextFieldComponent from '../../components/TextField'

const NewDefinitionForm = () => {
  const [data, setData] = useState([])
  const [submitting, setSubmitting] = useState(false)

  const handleSubmit = async definition => {
    setSubmitting(true)
    let apiURL = '/api/v1/cornerstone/issuer/definition/create'

    if (process.env.API_BASE_URL) {
      apiURL =
        process.env.API_BASE_URL + '/cornerstone/issuer/definition/create'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      await toast.promise(
        axios
          .post(process.env.REACT_APP_API_URL + 'definition/create', definition)
          .then(response => {
            console.log(response.data)
            setData(response.data)
            toast.success('Created Credential Definition!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Creating...',
        }
      )
    } else {
      await toast.promise(
        axios
          .post(apiURL, definition)
          .then(response => {
            console.log(response.data)
            setData(response.data)
            toast.success('Created Credential Definition!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Creating...',
        }
      )
    }
    setSubmitting(false)
  }

  return (
    <CardComponent>
      <GridComponent container spacing={2}>
        <GridComponent item xs={12} md={6}>
          <h2>Add New Credential Definition</h2>
          <FormikComponent
            initialValues={{
              schema_id: '',
              support_revocation: false,
              revocation_registry_size: 0,
              tag: '',
            }}
            onSubmit={(values, { resetForm }) => {
              handleSubmit(values)
              resetForm()
            }}
          >
            {({ values, handleChange }) => (
              <FormComponent>
                <div>
                  <TextFieldComponent
                    id="schema_id"
                    name="schema_id"
                    value={values.schema_id}
                    onChange={handleChange}
                    label="Schema ID"
                    sx={{ m: '1rem' }}
                    required
                  />
                </div>
                <div>
                  <FormControl sx={{ width: '16.5rem' }}>
                    <InputLabel
                      id="support_revocation_label"
                      sx={{ margin: '1rem 0 0 1rem' }}
                    >
                      Support Revocation
                    </InputLabel>
                    <Select
                      labelId="support_revocation_label"
                      id="support_revocation"
                      name="support_revocation"
                      value={values.support_revocation}
                      label="Support Revocation"
                      onChange={handleChange}
                      sx={{ m: '1rem' }}
                    >
                      <MenuItem value={false}>False</MenuItem>
                      <MenuItem value={true}>True</MenuItem>
                    </Select>
                  </FormControl>
                </div>
                <div>
                  <TextFieldComponent
                    id="revocation_registry_size"
                    name="revocation_registry_size"
                    type="number"
                    value={values.revocation_registry_size}
                    onChange={handleChange}
                    label="Revocation Registry Size"
                    sx={{ m: '1rem' }}
                  />
                </div>
                <div>
                  <TextFieldComponent
                    id="tag"
                    name="tag"
                    value={values.tag}
                    onChange={handleChange}
                    label="Tag"
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
                    Create
                  </ButtonComponent>
                </div>
              </FormComponent>
            )}
          </FormikComponent>
        </GridComponent>
        <GridComponent item xs={12} md={6}>
          {data ? (
            <div style={{ padding: '2rem' }}>
              <ListItemComponent>
                <ListItemTextComponent
                  primary={
                    'Credential Definition ID: ' + data.credential_definition_id
                  }
                />
              </ListItemComponent>
            </div>
          ) : (
            ''
          )}
        </GridComponent>
      </GridComponent>
    </CardComponent>
  )
}

export default NewDefinitionForm
