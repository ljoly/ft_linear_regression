package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func floatToString() (string, string) {
	return strconv.FormatFloat(theta0, 'f', -1, 64), strconv.FormatFloat(theta1, 'f', -1, 64)
}

func saveThetas() {
	str0, str1 := floatToString()
	str := str0 + "," + str1
	file, err := os.Create("../../assets/theta.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	fmt.Fprintf(file, str)
	file.Close()
}
