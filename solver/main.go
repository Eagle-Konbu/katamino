package main

import (
	"fmt"
	"net/http"
	"solver/solvepkg"
	"strconv"

	"github.com/labstack/echo"
)

type Response struct {
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	CalcTime  float64    `json:"calc_time"`
	Solutions [][]string `json:"solutions"`
}

func solveHandler(c echo.Context) error {
	width, _ := strconv.Atoi(c.Param("width"))
	height, _ := strconv.Atoi(c.Param("height"))

	solutions, calcTime := solvepkg.Solve(width, height)
	res := Response{Width: width, Height: height, CalcTime: calcTime, Solutions: solutions}

	fmt.Println(res)
	return c.JSON(http.StatusOK, res)
}

func main() {
	echo := echo.New()

	echo.GET("/solve/:height/:width", solveHandler)
	echo.Logger.Fatal(echo.Start("0.0.0.0:8080"))
}
