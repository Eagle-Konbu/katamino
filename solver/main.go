package main

import (
	"fmt"
	"net/http"
	"solver/solvepkg"

	"github.com/labstack/echo"
)

type Response struct {
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	CalcTime  float64    `json:"calc_time"`
	Solutions [][]string `json:"solutions"`
}

func solveHandler(c echo.Context) error {
	solutions, calcTime := solvepkg.Solve()
	var res Response
	res.Width = 5
	res.Height = 5
	res.CalcTime = calcTime
	res.Solutions = solutions

	fmt.Println(res)
	return c.JSON(http.StatusOK, res)
}

func main() {
	echo := echo.New()

	echo.GET("/", solveHandler)
	echo.Logger.Fatal(echo.Start("0.0.0.0:8080"))
}
