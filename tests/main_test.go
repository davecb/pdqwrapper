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
func Test_pdq(t *testing.T) {
	var progName = "pdq"
	tests := []struct {
		name        string
		think       float64
		serviceTime float64
		from        float64
		to          float64
		by          float64
	}{
		{
			name:        "fred",
			think:       1.0,
			serviceTime: 1.0,
			from:        1.0,
			to:          1.0,
			by:          1.0,
		},
	}
	var verbose, debug bool
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pdq(progName, tt.think, tt.serviceTime, tt.from, tt.to, tt.by, verbose, debug)
		})
	}

	var p PositiveFloat64
	p = p.New() // Instantiate a positive-number iterator

	var q float64
	var ok bool
	q, ok = p.Next()
	fmt.Printf("after initialization, q %g == 1.0, ok %t -= true\n", q, ok)

}

// PositiveFloat64 is a number > 0. It's subject to change or reconsideration entirely.
// Barely started (:-)), should be an iterator, so I can range over it.
type PositiveFloat64 struct {
	Value float64
	Legal bool
}

// PositiveFloat64 is a container for a positive number
// New returns an initialized PositiveFloat64
func (p PositiveFloat64) New() PositiveFloat64 {
	p.Value = 0 // the initial Next() will return 1, FIXME
	p.Legal = true
	return p
}

// Next returns the next integer value of p
func (p PositiveFloat64) Next() (float64, bool) {
	p.Value++
	return p.Value, p.Legal
}

// iteration should return (-MAXFLOAT, false), ... (0, false), (1, true) ...
