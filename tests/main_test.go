package main

import (
	"fmt"
	"testing"
)

// t is positive number, upper bound unknown
// z sleep is synonymous with t
// s is service time, positibe number, upper bound unknown
// v is a verbose flag
// d is a debug flag
// h is a usage flag, causing an exit
// from is a positive number, defaulting to 1
// to is a positive number, invalid if less than from
// by is a positive number, smaller than to-from
func Test_main(t *testing.T) {

	type Sample struct {
		Value float64
		Legal bool
	}

	positiveInt := []Sample{
		Sample{
			Value: 1.0,
			Legal: true,
		},
	}

	for i, sample := range positiveInt {
		fmt.Printf("%d: Value = %g, Legal = %t\n",
			i, sample.Value, sample.Legal)
	}
}
