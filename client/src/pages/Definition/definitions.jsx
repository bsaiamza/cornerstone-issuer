import { useEffect, useState } from 'react'
import axios from 'axios'
import { toast } from 'react-toastify'
// components
import ButtonComponent from '../../components/Button'
import CardComponent from '../../components/Card'
import GridComponent from '../../components/Grid'
import ListItemComponent from '../../components/ListItem'
import ListItemTextComponent from '../../components/ListItemText'

const Definitions = () => {
  const [didData, setDidData] = useState([])
  const [definitionData, setDefinitionData] = useState([])
  const [data, setData] = useState([])

  useEffect(() => {
    let didApiURL = '/api/v1/cornerstone/issuer/did'
    let definitionApiURL = '/api/v1/cornerstone/issuer/definitions'

    if (process.env.API_BASE_URL) {
      didApiURL = process.env.API_BASE_URL + '/cornerstone/issuer/did'

      definitionApiURL =
        process.env.API_BASE_URL + '/cornerstone/issuer/definitions'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      if (process.env.DID) {
        axios
          .get(
            process.env.REACT_APP_API_URL +
              'definitions?issuer_did=' +
              process.env.DID
          )
          .then(response => {
            setDefinitionData(response.data.credential_definition_ids)
          })
          .catch(error => console.log(error))
      } else {
        axios
          .get(process.env.REACT_APP_API_URL + 'did')
          .then(response => {
            setDidData(response.data)
          })
          .catch(error => console.log(error))

        if (didData.did !== 'undefined') {
          axios
            .get(
              process.env.REACT_APP_API_URL +
                'definitions?issuer_did=' +
                didData.did
            )
            .then(response => {
              setDefinitionData(response.data.credential_definition_ids)
            })
            .catch(error => console.log(error))
        }
      }
    } else {
      if (process.env.DID) {
        axios
          .get(definitionApiURL + '?issuer_did=' + process.env.DID)
          .then(response => {
            setDefinitionData(response.data.credential_definition_ids)
          })
          .catch(error => console.log(error))
      } else {
        axios
          .get(didApiURL)
          .then(response => {
            setDidData(response.data)
          })
          .catch(error => console.log(error))

        if (didData.did !== 'undefined') {
          axios
            .get(definitionApiURL + '?issuer_did=' + didData.did)
            .then(response => {
              setDefinitionData(response.data.credential_definition_ids)
            })
            .catch(error => console.log(error))
        }
      }
    }
  }, [didData.did])

  const getDefinition = creDefId => {
    let apiURL = '/api/v1/cornerstone/issuer/definition'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/definition'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      toast.promise(
        axios
          .get(
            process.env.REACT_APP_API_URL + 'definition?cred_def_id=' + creDefId
          )
          .then(response => {
            setData(response.data)
            toast.success('Fetched credential definition!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Fetching...',
        }
      )
    } else {
      toast.promise(
        axios
          .get(apiURL + '?cred_def_id=' + creDefId)
          .then(response => {
            setData(response.data)
            toast.success('Fetched credential definition!')
          })
          .catch(error => {
            toast.error(error.response.data.msg)
          }),
        {
          pending: 'Fetching...',
        }
      )
    }
  }

  return (
    <GridComponent container spacing={2}>
      <GridComponent item xs={12} md={4}>
        {definitionData ? (
          definitionData.length === 0 ? (
            <CardComponent
              sx={{
                m: '1rem',
              }}
            >
              No credential definitions available!
            </CardComponent>
          ) : (
            definitionData.map((definition, index) => (
              <CardComponent
                key={index}
                sx={{
                  m: '1rem',
                  overflowWrap: 'break-word',
                }}
              >
                <ListItemComponent>
                  Credential Definition ID:
                  <ButtonComponent
                    onClick={() => getDefinition(definition)}
                    sx={{ color: '#0645AD !important' }}
                  >
                    <ListItemTextComponent primary={definition} />
                  </ButtonComponent>
                </ListItemComponent>
              </CardComponent>
            ))
          )
        ) : (
          <CardComponent
            sx={{
              m: '1rem',
            }}
          >
            No credential definitions available!
          </CardComponent>
        )}
      </GridComponent>
      <GridComponent item xs={12} md={8}>
        <CardComponent
          sx={{
            m: '1rem',
          }}
        >
          <ListItemComponent>
            Credential Definition ID:{' '}
            <ListItemTextComponent primary={data.id} />
          </ListItemComponent>
          <ListItemComponent>
            Schema ID: <ListItemTextComponent primary={data.schemaId} />
          </ListItemComponent>
          <ListItemComponent>
            Type: <ListItemTextComponent primary={data.type} />
          </ListItemComponent>
          <ListItemComponent>
            Tag:
            <ListItemTextComponent primary={data.tag} />
          </ListItemComponent>
        </CardComponent>
      </GridComponent>
    </GridComponent>
  )
}

export default Definitions
