import { useState } from 'react'
// components
import BoxComponent from '../../components/Box'
import TabsComponent from '../../components/Tabs'
import TabComponent from '../../components/Tab'
import { a11yProps, TabPanel } from '../../components/TabPanel'
// definition
import Definitions from './definitions'
import NewDefinitionForm from './form'

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
          ariaLabel="Definition Tabs"
        >
          <TabComponent label="All" {...a11yProps(0)} />
          <TabComponent label="New" {...a11yProps(1)} />
        </TabsComponent>
      </BoxComponent>
      <TabPanel value={value} index={0}>
        <Definitions />
      </TabPanel>
      <TabPanel value={value} index={1}>
        <NewDefinitionForm />
      </TabPanel>
    </div>
  )
}

export default DefinitionPage
