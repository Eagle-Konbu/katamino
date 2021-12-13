// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import axios, { AxiosError, AxiosResponse } from 'axios';
import type { NextApiRequest, NextApiResponse } from 'next'
import type { Solution, SolverResponse } from '../../types/Solution'

type Data = {
  name: string
}

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<SolverResponse>
) {
  const width = req.query.width as string || "20";
  const height = req.query.height as string || "3"
  axios.get(`/api/solve/${height}/${width}`).then((r: AxiosResponse<SolverResponse>) => {
    const { data, status } = r;
    console.log(data);

    res.status(200).json(data);
  }).catch((err: AxiosError<{ error: string }>) => {
    console.error(err);
    res.status(500);
  });
}
