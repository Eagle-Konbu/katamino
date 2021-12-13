import axios, { AxiosError, AxiosResponse } from 'axios';
import { Box, Button, Card, CardBody, Header, Heading, RadioButtonGroup, Spinner, Text } from 'grommet'
import React from 'react';

import KataminoGrid from './components/KataminoGrid'
import type { Solution, SolverResponse } from './types/Solution'

function App() {
  const [size, setSize] = React.useState('6x10');
  const [isLoading, setIsLoading] = React.useState(false);

  const [solutions, setSolutions] = React.useState<Solution[]>([]);
  const [calcTime, setCalcTime] = React.useState(0);

  const handleClick = () => {
    setIsLoading(true);
    const width = Number(size.substring(2));
    const height = Number(size.substring(0, 1));
    axios.get(`/solve/${height}/${width}`).then((res: AxiosResponse<SolverResponse>) => {
      const { data, status } = res;
      console.log(data);
      
      setSolutions(data.solutions.map(s => {
        return {
          width: data.width,
          height: data.height,
          hexCodes: s
        }
      }));
      setCalcTime(data.calc_time);
    }).catch((err: AxiosError<{ error: string }>) => {
      console.error(err);
    }).finally(() => {
      setIsLoading(false);
    })
  }

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
          <Button primary label="Solve" onClick={handleClick} style={{ marginTop: '20px' }} />
        </Box>
        <div style={{ display: isLoading ? 'block' : 'none' }}><Spinner /></div>
        <Text style={{ display: solutions.length != 0 ? 'block' : 'none' }}><span>計算時間:{calcTime}s</span><span style={{ marginLeft: '20px' }}>解の個数: {solutions.length}個</span></Text>
        <Box direction="column" pad="medium">
          {
            solutions.map(s => {
              return (
                <Card background="light-1" style={{ marginBottom: '30px' }}>
                  <CardBody pad="medium">
                    <KataminoGrid solution={s} />
                  </CardBody>
                </Card>
              )
            })
          }
        </Box>
      </Box>
    </div>
  );
}

export default App;
