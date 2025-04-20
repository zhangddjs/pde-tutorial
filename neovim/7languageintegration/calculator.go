package main

import (
	"fmt"
	"math"
	"sort"
)

type Calculator struct {
	data []float64
}

// Add adds a number to the dataset.
func (c *Calculator) Add(x float64) {
	c.data = append(c.data, x)
}

// Sum returns the sum of all numbers in the dataset.
func (c Calculator) Sum() float64 {
	sum := 0.0
	for _, v := range c.data {
		sum += v
	}
	return sum
}

// Mean returns the average of the dataset.
func (c Calculator) Mean() float64 {
	if len(c.data) == 0 {
		return 0
	}
	return c.Sum() / float64(len(c.data))
}

// Max returns the maximum value in the dataset.
func (c Calculator) Max() float64 {
	if len(c.data) == 0 {
		return 0
	}
	max := c.data[0]
	for _, v := range c.data {
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimum value in the dataset.
func (c Calculator) Min() float64 {
	if len(c.data) == 0 {
		return 0
	}
	min := c.data[0]
	for _, v := range c.data {
		if v < min {
			min = v
		}
	}
	return min
}

// StdDev returns the standard deviation of the dataset.
func (c Calculator) StdDev() float64 {
	if len(c.data) == 0 {
		return 0
	}
	mean := c.Mean()
	var sumSquares float64
	for _, v := range c.data {
		diff := v - mean
		sumSquares += diff * diff
	}
	return math.Sqrt(sumSquares / float64(len(c.data)))
}

// Median returns the median value of the dataset.
func (c Calculator) Median() float64 {
	n := len(c.data)
	if n == 0 {
		return 0
	}
	sorted := make([]float64, n)
	copy(sorted, c.data)
	sort.Float64s(sorted)
	mid := n / 2
	if n%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

// Clear resets the dataset.
func (c *Calculator) Clear() {
	c.data = []float64{}
}

// Count returns the number of elements in the dataset.
func (c Calculator) Count() int {
	return len(c.data)
}

// Print displays the current dataset.
func (c Calculator) Print() {
	fmt.Println("Data:", c.data)
}

func main() {
	c := Calculator{}
	c.Add(10)
	c.Add(20)
	c.Add(30)
	c.Print()
	fmt.Println("Sum:", c.Sum())
	fmt.Println("Mean:", c.Mean())
	fmt.Println("Max:", c.Max())
	fmt.Println("Min:", c.Min())
	fmt.Println("StdDev:", c.StdDev())
	fmt.Println("Median:", c.Median())
	fmt.Println("Count:", c.Count())
	c.Clear()
	c.Print()
}
