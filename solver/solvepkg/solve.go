package solvepkg

func Solve(width int, height int) ([][]string, float64) {
	var solTmp []string
	// var solutions [][]string
	solutions := make([][]string, 4)

	for i := 0; i < 30; i++ {
		solTmp = append(solTmp, "#AA0000")
	}
	for i := 0; i < 30; i++ {
		solTmp = append(solTmp, "#00AA00")
	}

	for i := 0; i < 4; i++ {
		tmp := make([]string, len(solTmp))
		copy(tmp, solTmp)
		solutions[i] = tmp
	}

	return solutions, 2.0
}
