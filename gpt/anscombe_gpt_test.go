package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

// TestRegression runs tests for the regression calculations
func TestRegression(t *testing.T) {
	t.Run("Set I", func(t *testing.T) {
		testRegression(t, "Set I", []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			[]float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68})
	})

	t.Run("Set II", func(t *testing.T) {
		testRegression(t, "Set II", []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
			[]float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74})
	})

	// Add tests for other sets as needed
}

func testRegression(t *testing.T, setName string, x, y []float64) {
	// Run the regression function
	setDesignMatrix := createSeries(x, y)
	setModel, _ := stats.LinearRegression(setDesignMatrix)

	// Test the calculated intercept and slope
	if err := checkRegression(setName, setModel); err != nil {
		t.Error(err)
	}
}

func checkRegression(setName string, model stats.Series) error {
	// Find the first pair of non-equal coordinates
	var firstCoord, lastCoord stats.Coordinate
	for i := 0; i < len(model)-1; i++ {
		if model[i] != model[i+1] {
			firstCoord = model[i]
			lastCoord = model[i+1]
			break
		}
	}

	// Checking if all coordinates are equal
	if firstCoord == lastCoord {
		return fmt.Errorf("all coordinates are equal for regression on %s. Unable to calculate slope and intercept", setName)
	}

	// Calculating the slope and intercept from the coordinates
	slope := (lastCoord.Y - firstCoord.Y) / (lastCoord.X - firstCoord.X)
	intercept := firstCoord.Y - slope*firstCoord.X

	// Check if the calculated intercept and slope are almost equal to the expected values
	interceptTolerance := 0.0001
	slopeTolerance := 0.0001
	if !almostEqual(intercept, expectedIntercept[setName], interceptTolerance) || !almostEqual(slope, expectedSlope[setName], slopeTolerance) {
		return fmt.Errorf("failed for %s. Expected Intercept: %.4f, Slope: %.4f. Got Intercept: %.4f, Slope: %.4f",
			setName, expectedIntercept[setName], expectedSlope[setName], intercept, slope)
	}

	return nil
}

// Utility function to check if two float64 values are almost equal
func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
}

// Add expected intercept and slope values for each set
var expectedIntercept = map[string]float64{
	"Set I":  3.0001,
	"Set II": 3.0009,
	// Add expected values for other sets
}

var expectedSlope = map[string]float64{
	"Set I":  0.5001,
	"Set II": 0.5000,
	// Add expected values for other sets
}

// Anscombe datasets
var anscombeX = [][]float64{
	{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
	{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
}

var anscombeY = [][]float64{
	{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
	{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
	{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
	{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
}

func BenchmarkLinearRegression(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(anscombeX); j++ {
			_, _ = stats.LinearRegression(createSeries(anscombeX[j], anscombeY[j]))
		}
	}
}
