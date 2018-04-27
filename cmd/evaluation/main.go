package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Point struct {
	x  float64
	y  float64
	y2 float64
}

var (
	points   []Point
	meanYs   float64
	theta0   float64
	theta1   float64
	rSquared float64
)

func getEstimateOrdinate(x float64) float64 {
	return theta0 + (theta1 * x)
}

func getYMean() {
	for _, p := range points {
		meanYs += p.y
	}
	meanYs /= float64(len(points))
}

func getrSquared() {
	var (
		modelSquaredError  float64
		meanYsSquaredError float64
	)
	for _, p := range points {
		modelSquaredError += (p.y - p.y2) * (p.y - p.y2)
		meanYsSquaredError += (p.y - meanYs) * (p.y - meanYs)
	}
	rSquared = 1 - (modelSquaredError / meanYsSquaredError)
}

func getData() {
	csvFile, _ := os.Open("../../assets/data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		var (
			a float64
			b float64
		)
		fmt.Sscanf(line[0], "%f", &a)
		fmt.Sscanf(line[1], "%f", &b)
		if !(a == 0 && b == 0) {
			points = append(points, Point{
				x:  a,
				y:  b,
				y2: getEstimateOrdinate(a),
			},
			)
		}
	}
}

func getThetas() {
	csvFile, _ := os.Open("../../assets/theta.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sscanf(line[0], "%f", &theta0)
	fmt.Sscanf(line[1], "%f", &theta1)
}

func main() {
	getThetas()
	getData()
	getYMean()
	getrSquared()
	fmt.Printf("Model accuracy: %0.1f%%\n", rSquared*100)
}
