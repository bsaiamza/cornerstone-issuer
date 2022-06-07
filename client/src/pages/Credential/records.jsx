import axios from 'axios'
import { useEffect, useState } from 'react'
// import { toast } from 'react-toastify'
// components
// import ButtonComponent from '../../components/Button'
import CardComponent from '../../components/Card'
// import FormikComponent from '../../components/Formik'
// import FormComponent from '../../components/Form'
import GridComponent from '../../components/Grid'
import ListItemComponent from '../../components/ListItem'
import ListItemTextComponent from '../../components/ListItemText'
// import TextFieldComponent from '../../components/TextField'

const CredentialRecords = () => {
  const [data, setData] = useState([])
  // const [submitting, setSubmitting] = useState(false)
  // let idNumber = useState('')

  useEffect(() => {
    let apiURL = '/api/v1/cornerstone/issuer/credentials'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/credentials'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      axios
        .get(process.env.REACT_APP_API_URL + 'credentials')
        .then(response => {
          console.log(response.data)
          setData(response.data)
        })
        .catch(error => console.log(error))
    } else {
      axios
        .get(apiURL)
        .then(response => {
          console.log(response.data)
          setData(response.data)
        })
        .catch(error => console.log(error))
    }
  }, [])

  // const sendOffer = async (credExID, id_no, schema_id, cred_def_id) => {
  //   await toast.promise(
  //     axios
  //       .post(
  //         '/api/v1/cornerstone/issuer/credential/offer?cred_ex_id=' + credExID,
  //         {
  //           id_no,
  //           schema_id,
  //           cred_def_id,
  //         }
  //       )
  //       // .post(
  //       //   process.env.REACT_APP_API_URL +
  //       //     'credential/offer?cred_ex_id=' +
  //       //     credExID,
  //       //   { id_no, schema_id, cred_def_id }
  //       // )
  //       .then(response => {
  //         setSubmitting(true)
  //         toast.success('Approved!')
  //       })
  //       .catch(error => {
  //         toast.error(error.response.data.msg)
  //       }),
  //     {
  //       pending: 'Approving...',
  //     }
  //   )
  // }

  // const issueCredential = async (credExID, comment) => {
  //   await toast.promise(
  //     axios
  //       .post(
  //         '/api/v1/cornerstone/issuer/credential/issue?cred_ex_id=' + credExID,
  //         { comment }
  //       )
  //       // .post(
  //       //   process.env.REACT_APP_API_URL +
  //       //     'credential/issue?cred_ex_id=' +
  //       //     credExID,
  //       //   { comment }
  //       // )
  //       .then(response => {
  //         setSubmitting(true)
  //         toast.success('Issued!')
  //       })
  //       .catch(error => {
  //         toast.error(error.response.data.msg)
  //       }),
  //     {
  //       pending: 'Issuing...',
  //     }
  //   )
  // }

  return (
    <GridComponent container justify="center" spacing={2}>
      {data ? (
        data.length === 0 ? (
          <CardComponent
            sx={{
              m: '1rem',
            }}
          >
            No credential requests available!
          </CardComponent>
        ) : (
          data.map((request, index) => (
            <GridComponent item xs={12} md={4}>
              <CardComponent
                key={index}
                sx={{
                  m: '1rem',
                  wordWrap: 'break-word',
                }}
              >
                <ListItemComponent>
                  <ListItemTextComponent
                    primary={'Created: ' + request.created_at}
                  />
                </ListItemComponent>
                <ListItemComponent>
                  <ListItemTextComponent
                    primary={'Connection ID: ' + request.connection_id}
                  />
                </ListItemComponent>
                <ListItemComponent>
                  <ListItemTextComponent
                    primary={
                      'Credential Exchange ID: ' +
                      request.credential_exchange_id
                    }
                  />
                </ListItemComponent>

                {request.error_msg.length !== 0 ? (
                  <ListItemComponent>
                    <ListItemTextComponent primary={request.error_msg} />
                  </ListItemComponent>
                ) : (
                  <ListItemComponent>
                    <ListItemTextComponent
                      primary={'State: ' + request.state}
                    />
                  </ListItemComponent>
                )}
                {/* <ListItemComponent>
                  <ListItemTextComponent primary={'Credential Preview: '} />
                </ListItemComponent> */}
                {/* {request.cred_ex_record.cred_preview.attributes.map(
                  (attribute, index) => {
                    // if (attribute.name === 'id_number') {
                    //   idNumber = attribute.value
                    // }

                    return (
                      <ListItemComponent key={index}>
                        <ListItemTextComponent
                          primary={
                            'Attribute: ' +
                            attribute.name +
                            ', Value: ' +
                            attribute.value
                          }
                        />
                      </ListItemComponent>
                    )
                  }
                )} */}
                {/* {request.cred_ex_record.state === 'proposal-received' ? (
                  <FormikComponent
                    initialValues={{
                      id_num: idNumber,
                      schema_id: '',
                      cred_def_id: '',
                    }}
                    onSubmit={(values, { resetForm }) => {
                      sendOffer(
                        request.cred_ex_record.cred_ex_id,
                        values.id_num,
                        values.schema_id,
                        values.cred_def_id
                      )
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
                          <TextFieldComponent
                            id="cred_def_id"
                            name="cred_def_id"
                            value={values.cred_def_id}
                            onChange={handleChange}
                            label="Credential Definition ID"
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
                            Approve
                          </ButtonComponent>
                        </div>
                      </FormComponent>
                    )}
                  </FormikComponent>
                ) : (
                  ''
                )}
                {request.cred_ex_record.state === 'request-received' ? (
                  <FormikComponent
                    initialValues={{
                      comment: '',
                    }}
                    onSubmit={(values, { resetForm }) => {
                      issueCredential(
                        request.cred_ex_record.cred_ex_id,
                        values.comment
                      )
                      resetForm()
                    }}
                  >
                    {({ values, handleChange }) => (
                      <FormComponent>
                        <div>
                          <TextFieldComponent
                            id="comment"
                            name="comment"
                            value={values.comment}
                            onChange={handleChange}
                            label="Comment"
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
                            Issue
                          </ButtonComponent>
                        </div>
                      </FormComponent>
                    )}
                  </FormikComponent>
                ) : (
                  ''
                )} */}
                <ListItemComponent sx={{ backgroundColor: '#eee' }}>
                  <ListItemTextComponent
                    primary={'Last updated at: ' + request.updated_at}
                  />
                </ListItemComponent>
              </CardComponent>
            </GridComponent>
          ))
        )
      ) : (
        <CardComponent
          sx={{
            m: '1rem',
          }}
        >
          No credential requests available!
        </CardComponent>
      )}
    </GridComponent>
  )
}

export default CredentialRecords
