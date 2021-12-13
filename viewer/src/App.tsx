import { Box, Button, Card, CardBody, Form, Grid, Header, Heading, RadioButtonGroup } from 'grommet'
import React from 'react';

import KataminoGrid from './components/KataminoGrid'
import type { Solution } from './types/Solution'

function App() {
  const [size, setSize] = React.useState('6x10');

  const solution: Solution = {
    width: 10,
    height: 6,
    hexCodes: Array(60).fill("#CC7700")
  };

  return (
    <div>
      <Header background="brand">
        <Heading size="xsmall" margin="medium">Katamino Solver</Heading>
      </Header>
      <Box direction="column" pad="medium" align="center" background="light-2">
        <Box pad="medium" align="center">
          <RadioButtonGroup
            name="size_radio"
            options={['6x10', '5x12', '4x15', '3x20']}
            value={size}
            onChange={(e) => { setSize(e.target.value) }}
          />
          <Button primary label="Solve" style={{ marginTop: '20px' }} />
        </Box>
        <Card background="light-1">
          <CardBody pad="medium">
            <KataminoGrid solution={solution} />
          </CardBody>
        </Card>
      </Box>
    </div>
  );
}

export default App;
