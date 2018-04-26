package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	theta0 float64
	theta1 float64
	input  string
)

func getEstimatePrice(mileage float64) float64 {
	return theta0 + (theta1 * mileage)
}

func getInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a mileage: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
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
	getInput()
	mileage, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	price := getEstimatePrice(mileage)
	fmt.Printf("Price: %.0f€\n", price)
}
