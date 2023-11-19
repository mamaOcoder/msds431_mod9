package main

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

func runRegression() {
	// Define the Anscombe data
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	x2 := x1
	x3 := x1
	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	// Fit linear regression models by ordinary least squares
	setIDesignMatrix := createSeries(x1, y1)
	setIModel, _ := stats.LinearRegression(setIDesignMatrix)
	printRegressionSummary("Set I", setIModel)

	setIIDesignMatrix := createSeries(x2, y2)
	setIIModel, _ := stats.LinearRegression(setIIDesignMatrix)
	printRegressionSummary("Set II", setIIModel)

	setIIIDesignMatrix := createSeries(x3, y3)
	setIIIModel, _ := stats.LinearRegression(setIIIDesignMatrix)
	printRegressionSummary("Set III", setIIIModel)

	setIVDesignMatrix := createSeries(x4, y4)
	setIVModel, _ := stats.LinearRegression(setIVDesignMatrix)
	printRegressionSummary("Set IV", setIVModel)
}

func createSeries(x, y []float64) stats.Series {
	series := make(stats.Series, len(x))
	for i := range x {
		series[i] = stats.Coordinate{X: x[i], Y: y[i]}
	}
	return series
}

func printRegressionSummary(setName string, model stats.Series) {
	intercept := model[0].Y
	slope := model[1].Y
	fmt.Printf("%s Regression:\n", setName)
	fmt.Printf("Intercept: %.4f\n", intercept)
	fmt.Printf("Slope: %.4f\n", slope)
	fmt.Println()
}

func main() {
	// Run the regression function
	runRegression()
}
