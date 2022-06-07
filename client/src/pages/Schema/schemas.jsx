import { useEffect, useState } from 'react'
import axios from 'axios'
import { toast } from 'react-toastify'
// components
import ButtonComponent from '../../components/Button'
import CardComponent from '../../components/Card'
import GridComponent from '../../components/Grid'
import ListItemComponent from '../../components/ListItem'
import ListItemTextComponent from '../../components/ListItemText'

const Schemas = () => {
  const [didData, setDidData] = useState([])
  const [schemaData, setSchemaData] = useState([])
  const [data, setData] = useState([])

  useEffect(() => {
    let didApiURL = '/api/v1/cornerstone/issuer/did'
    let schemaApiURL = '/api/v1/cornerstone/issuer/schemas'

    if (process.env.API_BASE_URL) {
      didApiURL = process.env.API_BASE_URL + '/cornerstone/issuer/did'

      schemaApiURL = process.env.API_BASE_URL + '/cornerstone/issuer/schemas'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      if (process.env.DID) {
        axios
          .get(
            process.env.REACT_APP_API_URL +
              'schemas?schema_issuer_did=' +
              process.env.DID
          )
          .then(response => {
            setSchemaData(response.data.schema_ids)
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
                'schemas?schema_issuer_did=' +
                didData.did
            )
            .then(response => {
              setSchemaData(response.data.schema_ids)
            })
            .catch(error => console.log(error))
        }
      }
    } else {
      if (process.env.DID) {
        axios
          .get(schemaApiURL + '?schema_issuer_did=' + process.env.DID)
          .then(response => {
            setSchemaData(response.data.schema_ids)
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
            .get(schemaApiURL + '?schema_issuer_did=' + didData.did)
            .then(response => {
              setSchemaData(response.data.schema_ids)
            })
            .catch(error => console.log(error))
        }
      }
    }
  }, [didData.did])

  const getSchema = schemaId => {
    let apiURL = '/api/v1/cornerstone/issuer/schema'

    if (process.env.API_BASE_URL) {
      apiURL = process.env.API_BASE_URL + '/cornerstone/issuer/schema'
    }

    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      toast.promise(
        axios
          .get(process.env.REACT_APP_API_URL + 'schema?schema_id=' + schemaId)
          .then(response => {
            setData(response.data.schema)
            toast.success('Fetched schema!')
          })
          .catch(error => {
            toast.error(error.msg)
          }),
        {
          pending: 'Fetching...',
        }
      )
    } else {
      toast.promise(
        axios
          .get(apiURL + '?schema_id=' + schemaId)
          .then(response => {
            setData(response.data.schema)
            toast.success('Fetched schema!')
          })
          .catch(error => {
            toast.error(error.msg)
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
        {schemaData ? (
          schemaData.length === 0 ? (
            <CardComponent
              sx={{
                m: '1rem',
              }}
            >
              No schemas available!
            </CardComponent>
          ) : (
            schemaData.map((schema, index) => (
              <CardComponent
                key={index}
                sx={{
                  m: '1rem',
                  overflowWrap: 'break-word',
                }}
              >
                <ListItemComponent>
                  Schema ID:
                  <ButtonComponent
                    onClick={() => getSchema(schema)}
                    sx={{ color: '#0645AD !important' }}
                  >
                    <ListItemTextComponent primary={schema} />
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
            No schemas available!
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
            Schema ID: <ListItemTextComponent primary={data.id} />
          </ListItemComponent>
          <ListItemComponent>
            Schema Name: <ListItemTextComponent primary={data.name} />
          </ListItemComponent>
          <ListItemComponent>
            Schema Version: <ListItemTextComponent primary={data.version} />
          </ListItemComponent>
          <ListItemComponent>
            Schema Attributes:
            <ListItemTextComponent primary={data.attrNames} />
          </ListItemComponent>
        </CardComponent>
      </GridComponent>
    </GridComponent>
  )
}

export default Schemas
