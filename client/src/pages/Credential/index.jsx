import { useState } from 'react'
// components
import BoxComponent from '../../components/Box'
import TabsComponent from '../../components/Tabs'
import TabComponent from '../../components/Tab'
import { a11yProps, TabPanel } from '../../components/TabPanel'
// credential
import CredentialRecords from './records'

const DefinitionPage = () => {
  const [value, setValue] = useState(0)

  const handleChange = (event, newValue) => {
    setValue(newValue)
  }

  return (
    <div style={{ margin: '1rem' }}>
      <BoxComponent>
        <TabsComponent
          value={value}
          onChange={handleChange}
          ariaLabel="Credentials Tabs"
        >
          <TabComponent label="Requested" {...a11yProps(0)} />
        </TabsComponent>
      </BoxComponent>
      <TabPanel value={value} index={0}>
        <CredentialRecords />
      </TabPanel>
    </div>
  )
}

export default DefinitionPage
