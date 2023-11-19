package main

//go:generate gotests -w -all anscombe.go

//Importing packages including stats
import (
	"fmt"
	"math"

	"github.com/montanaflynn/stats"
)

// Function to round floats taken from https://gosamples.dev/round-float/
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// Calculate the coefficients (slope and intercept) of the linear regression line for the data points
func RegressCoef(data []stats.Coordinate) []float64 {

	// Run the least squares regression using Go's stats package
	r, _ := stats.LinearRegression(data)

	// stats.LinearRegression returns the expected values for each provided point
	// rather than the coefficients like Python and R, so need to calculate coefficients
	// Need to use 7th and 8th position to avoid zero value in 4th set
	slope := (r[7].Y - r[6].Y) / (r[7].X - r[6].X)
	intercept := (r[7].Y - (r[7].X * slope))

	// coefficients slice
	coef := []float64{roundFloat(slope, 1), roundFloat(intercept, 1)}
	return coef
}

func main() {
	// data variables defines the data for the different plots.
	// I wanted to define the data in a map to make it similar to the format for Python and R, however,
	// could not come up with a easy transformation to the required coordinate format that
	// the stats package required at this time. I will investigate further.
	data1 := []stats.Coordinate{
		{X: 10, Y: 8.04},
		{X: 8, Y: 6.95},
		{X: 13, Y: 7.58},
		{X: 9, Y: 8.81},
		{X: 11, Y: 8.33},
		{X: 14, Y: 9.96},
		{X: 6, Y: 7.24},
		{X: 4, Y: 4.26},
		{X: 12, Y: 10.84},
		{X: 7, Y: 4.82},
		{X: 5, Y: 5.68},
	}
	data2 := []stats.Coordinate{
		{X: 10, Y: 9.14},
		{X: 8, Y: 8.14},
		{X: 13, Y: 8.74},
		{X: 9, Y: 8.77},
		{X: 11, Y: 9.26},
		{X: 14, Y: 8.1},
		{X: 6, Y: 6.13},
		{X: 4, Y: 3.1},
		{X: 12, Y: 9.13},
		{X: 7, Y: 7.26},
		{X: 5, Y: 4.74},
	}
	data3 := []stats.Coordinate{
		{X: 10, Y: 7.46},
		{X: 8, Y: 6.77},
		{X: 13, Y: 12.74},
		{X: 9, Y: 7.11},
		{X: 11, Y: 7.81},
		{X: 14, Y: 8.84},
		{X: 6, Y: 6.08},
		{X: 4, Y: 5.39},
		{X: 12, Y: 8.15},
		{X: 7, Y: 6.42},
		{X: 5, Y: 5.73},
	}
	data4 := []stats.Coordinate{
		{X: 8, Y: 6.58},
		{X: 8, Y: 5.76},
		{X: 8, Y: 7.71},
		{X: 8, Y: 8.84},
		{X: 8, Y: 8.47},
		{X: 8, Y: 7.04},
		{X: 8, Y: 5.25},
		{X: 19, Y: 12.5},
		{X: 8, Y: 5.56},
		{X: 8, Y: 7.91},
		{X: 8, Y: 6.89},
	}

	// Run function to calculate regression coefficients
	coef1 := RegressCoef(data1)
	coef2 := RegressCoef(data2)
	coef3 := RegressCoef(data3)
	coef4 := RegressCoef(data4)

	fmt.Println(coef1)
	fmt.Println(coef2)
	fmt.Println(coef3)
	fmt.Println(coef4)
}
