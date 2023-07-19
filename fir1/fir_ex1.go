package main

import (
	"fmt"

	tsp "github.com/erdetn/go-tsp"
)

func main() {

	c := []float64{float64(0.1), 0.2, 0.3, 0.2, 0.1, -0.1, -0.2, 0.2}
	signal := []float64{float64(1), 2.0, 3.0, 4.0, 5.0, 6.0, 7.0,
		8.0, 9.0, 8.0, 7.0, 6.0, 5.0, 4.0,
		0.3, 2.0, 1.0, 0.5, 2.0, 8.0, 2.0}

	filter := tsp.NewFirFilter(c)

	i := 0
	y := float64(0)
	for _, sample := range signal {
		y = filter.Filter(sample)
		fmt.Printf("<[%d] fir(%.4f) -> y = %.4f>\n", i, sample, y)
		i++
	}
}
