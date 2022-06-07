import { useState } from 'react'
// components
import BoxComponent from '../../components/Box'
import TabsComponent from '../../components/Tabs'
import TabComponent from '../../components/Tab'
import { a11yProps, TabPanel } from '../../components/TabPanel'
// schema
import Schemas from './schemas'
import NewSchemaForm from './form'

const SchemaPage = () => {
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
          ariaLabel="Schema Tabs"
        >
          <TabComponent label="All" {...a11yProps(0)} />
          <TabComponent label="New" {...a11yProps(1)} />
        </TabsComponent>
      </BoxComponent>
      <TabPanel value={value} index={0}>
        <Schemas />
      </TabPanel>
      <TabPanel value={value} index={1}>
        <NewSchemaForm />
      </TabPanel>
    </div>
  )
}

export default SchemaPage
