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
func Test_pdf(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name          string
		think         float64
		serviceDemand float64
		from          float64
		to            float64
		by            float64
	}{
		{
			name:          "fred",
			think:         1.0,
			serviceDemand: 1.0,
			from:          1.0,
			to:            1.0,
			by:            1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdf(tt.think, tt.serviceDemand, tt.from, tt.to, tt.by)
		})
	}

	var p Positive

	p = p.New() // Instantiate a positive
	var q float64
	q = p.Next()
	fmt.Printf("after initialization, q %g == 1.0\n", q)

}

// Positive is a number > 0. It is initially implemented as a float64 to get
// more range than an int64, but that's mostly an artifact of how it was
// developed. It's subject to change or reconsideration entirely.
type Positive struct {
	Value float64
}

// Positive is a container for a positive number

// New returns an initialized Positive, p
func (p Positive) New() Positive {
	p.Value = 0 // the initial Next() will return 1
	return p
}

// Next returns the next integer value of p
func (p Positive) Next() float64 {
	q := p.Value
	p.Value++
	return q
}
