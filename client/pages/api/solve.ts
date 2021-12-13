// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import axios, { AxiosError, AxiosResponse } from 'axios';
import type { NextApiRequest, NextApiResponse } from 'next'
import type { Solution, SolverResponse } from '../../types/Solution'

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<SolverResponse>
) {
  const width = req.query.width as string || "20";
  const height = req.query.height as string || "3"
  axios.get(`http://solver:8080/solve/${height}/${width}`).then((r: AxiosResponse<SolverResponse>) => {
    const { data, status } = r;

    res.status(200).json(data);
  }).catch((err: AxiosError<{ error: string }>) => {
    console.error(err);
    res.status(500);
  });
}
