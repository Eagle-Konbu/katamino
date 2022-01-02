package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"solver/solvepkg"
)

type Response struct {
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	CalcTime  float64    `json:"calc_time"`
	Solutions [][]string `json:"solutions"`
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	solutions, calcTime := solvepkg.Solve()
	var res Response
	res.Width = 5
	res.Height = 5
	res.CalcTime = calcTime
	res.Solutions = solutions

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	w.Header().Set("Content-Type", "application/json")

	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", solveHandler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
