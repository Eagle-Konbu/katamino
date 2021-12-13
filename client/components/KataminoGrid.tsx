import { Box, Grid } from "grommet";
import type { Solution } from '../types/Solution'

function KataminoGrid(props: { solution: Solution }) {
  const grids = () => {
    let elements = [];
    for (let i = 0; i < props.solution.height; i++) {
      for (let j = 0; j < props.solution.width; j++) {
        elements.push((
          <Box background={props.solution.hexCodes[i * props.solution.width + j]} />
        ))
      }
    }
    return elements;
  }
  return (
    <Grid
      columns={Array(props.solution.width).fill('xxsmall')}
      rows={Array(props.solution.height).fill('xxsmall')}
      gap='xxsmall'
    >
      {grids()}
    </Grid>
  )
}

export default KataminoGrid;