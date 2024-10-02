package main

import (
	"fmt"
	"log"
	"math"
	"os"

	linears "linears/handleData"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide data file.")
	}

	path := os.Args[1]
	data, err := linears.ReadData(path)
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	n := float64(len(data))
	if n == 0 {
		log.Fatal("provide data, file is empty")
	}

	xSum := 0.0
	ySum := 0.0
	xSquaredSum := 0.0
	xySum := 0.0

	for i, y := range data {
		x := float64(i)
		xSum += x
		ySum += y
		xSquaredSum += x * x
		xySum += x * y
	}

	slope := (n*xySum - xSum*ySum) / (n*xSquaredSum - xSum*xSum)
	intercept := (ySum - slope*xSum) / n

	var numerator float64
	var xMean float64
	var yMean float64
	for _, y := range data {
		x := float64(linears.IndexData(data, y))
		xMean += x
		yMean += y
	}
	xMean /= n
	yMean /= n

	for _, y := range data {
		x := float64(linears.IndexData(data, y))
		numerator += (x - xMean) * (y - yMean)
	}

	var xDenominator float64
	var yDenominator float64
	for _, y := range data {
		x := float64(linears.IndexData(data, y))
		xDenominator += (x - xMean) * (x - xMean)
		yDenominator += (y - yMean) * (y - yMean)
	}

	r := numerator / math.Sqrt(xDenominator*yDenominator)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
