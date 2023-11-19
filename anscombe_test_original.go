package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

// Unit test for RegressCoef function ensuring that the slopes and intercepts are 0.5 and 3 for each dataset.
func TestRegressCoef(t *testing.T) {
	slope := 0.5
	intercept := 3.0

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
	coef1 := RegressCoef(data1)
	coef2 := RegressCoef(data2)
	coef3 := RegressCoef(data3)
	coef4 := RegressCoef(data4)

	if coef1[0] != slope {
		t.Error("The slope in the 1st dataset is incorrect: expected 0.5, got", coef1[0])
	}
	if coef2[0] != slope {
		t.Error("The slope in the 2nd dataset is incorrect: expected 0.5, got", coef2[0])
	}
	if coef3[0] != slope {
		t.Error("The slope in the 3rd dataset is incorrect: expected 0.5, got", coef3[0])
	}
	if coef4[0] != slope {
		t.Error("The slope in the 4th dataset is incorrect: expected 0.5, got", coef4[0])
	}
	if coef1[1] != intercept {
		t.Error("The intercept in the 1st dataset is incorrect: expected 3.0, got", coef1[1])
	}
	if coef2[1] != intercept {
		t.Error("The intercept in the 2nd dataset is incorrect: expected 3.0, got", coef2[1])
	}
	if coef3[1] != intercept {
		t.Error("The intercept in the 3rd dataset is incorrect: expected 3.0, got", coef3[1])
	}
	if coef4[1] != intercept {
		t.Error("The intercept in the 4th dataset is incorrect: expected 3.0, got", coef4[1])
	}
}

// Write result to blackholes for benchmark test
var blackhole1 []float64
var blackhole2 []float64
var blackhole3 []float64
var blackhole4 []float64

// Benchmark test for RegressCoef function
func BenchmarkRegressCoef(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
		coef1 := RegressCoef(data1)
		coef2 := RegressCoef(data2)
		coef3 := RegressCoef(data3)
		coef4 := RegressCoef(data4)

		blackhole1 = coef1
		blackhole2 = coef2
		blackhole3 = coef3
		blackhole4 = coef4

	}
}
