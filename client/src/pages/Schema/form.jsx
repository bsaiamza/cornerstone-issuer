import { useState } from 'react'
import { toast } from 'react-toastify'
import { FieldArray } from 'formik'
import { Remove } from '@mui/icons-material'
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
import TooltipComponent from '../../components/Tooltip'

const NewSchemaForm = () => {
  const [data, setData] = useState([])
  const [submitting, setSubmitting] = useState(false)

  const handleSubmit = async schema => {
    setSubmitting(true)
    let apiURL = '/api/v1/cornerstone/issuer/schema/create'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/schema/create'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      await toast.promise(
        axios
          .post(process.env.REACT_APP_API_URL + 'schema/create', schema)
          .then(response => {
            setData(response.data.schema)
            toast.success('Created Schema!')
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
          .post(apiURL, schema)
          .then(response => {
            setData(response.data.schema)
            toast.success('Created Schema!')
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
          <h2>Add New Schema</h2>
          <FormikComponent
            initialValues={{
              schema_name: '',
              schema_version: '',
              attributes: [],
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
                    id="schema_name"
                    name="schema_name"
                    value={values.schema_name}
                    onChange={handleChange}
                    label="Schema Name"
                    sx={{ m: '1rem' }}
                    required
                  />
                </div>
                <div>
                  <TextFieldComponent
                    id="schema_version"
                    name="schema_version"
                    value={values.schema_version}
                    onChange={handleChange}
                    label="Schema Version"
                    sx={{ m: '1rem' }}
                    required
                  />
                </div>
                <FieldArray
                  name="attributes"
                  render={arrayHelpers => (
                    <div>
                      {values.attributes.map((attribute, index) => (
                        <div key={index}>
                          <TextFieldComponent
                            id={`attributes.${index}`}
                            name={`attributes.${index}`}
                            value={attribute}
                            onChange={handleChange}
                            label="Schema Attribute Name"
                            sx={{ m: '1rem' }}
                            required
                          />
                          <TooltipComponent title="Remove attribute">
                            <ButtonComponent
                              onClick={() => arrayHelpers.remove(index)} // remove an attribute from the list
                              sx={{ m: '1rem' }}
                            >
                              <Remove />
                            </ButtonComponent>
                          </TooltipComponent>
                        </div>
                      ))}
                      <ButtonComponent
                        variant="outlined"
                        sx={{ m: '1rem' }}
                        onClick={() => arrayHelpers.push('')}
                      >
                        {/* show this when user has removed all attributes from the list */}
                        Add an Attribute
                      </ButtonComponent>
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
                    </div>
                  )}
                />
              </FormComponent>
            )}
          </FormikComponent>
        </GridComponent>
        <GridComponent item xs={12} md={6}>
          {data ? (
            <div style={{ padding: '2rem' }}>
              <ListItemComponent>
                <ListItemTextComponent primary={'Schema ID: ' + data.id} />
              </ListItemComponent>
              <ListItemComponent>
                <ListItemTextComponent primary={'Schema Name: ' + data.name} />
              </ListItemComponent>
              <ListItemComponent>
                <ListItemTextComponent
                  primary={'Schema Version: ' + data.version}
                />
              </ListItemComponent>
              <ListItemComponent>
                <ListItemTextComponent
                  primary={'Schema Attributes: ' + data.attrNames}
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

export default NewSchemaForm
