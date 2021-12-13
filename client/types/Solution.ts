export type Solution = {
    width: number;
    height: number;
    hexCodes: string[]
}

export type SolverResponse = {
    width: number;
    height: number;
    calc_time: number;
    solutions: string[][];
}