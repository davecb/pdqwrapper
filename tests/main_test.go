package main

import (
	"fmt"
	"math"
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
	//var progName = "pdq"
	//tests := []struct {
	//	name        string
	//	think       float64
	//	serviceTime float64
	//	from        float64
	//	to          float64
	//	by          float64
	//}{
	//	{
	//		name:        "fred",
	//		think:       1.0,
	//		serviceTime: 1.0,
	//		from:        1.0,
	//		to:          1.0,
	//		by:          1.0,
	//	},
	//}
	////var verbose, debug bool
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		pdq(progName, tt.think, tt.serviceTime, tt.from, tt.to, tt.by, verbose, debug)
	//	})
	//}

	var p PositiveFloat64
	var limit int
	p, limit = p.New() // Instantiate a positive-number iterator
	fmt.Printf("after initialization, s = %g, ok == %t, limit == %d\n", p.Value, p.Legal, limit)

	var d1 float64
	var ok bool
	for i := 0; i < 15; d1, ok = p.Next() {
		fmt.Printf("%d: d1 = %g, ok == %t\n", i, d1, ok)
		i++ //foo(d1, ok)
	}
}

//func foo(sample float64, ok bool) {
//	fmt.Printf("after initialization, s = %g, ok == %t\n", sample, ok)
//}

// PositiveFloat64 is a number > 0. It's subject to change or reconsideration entirely.
// Barely started (:-)), should be an iterator, so I can range over it.
type PositiveFloat64 struct {
	Value          float64
	Legal          bool
	index          int
	sampleFloat64s []float64
}

// PositiveFloat64 is a container for a positive number
// New returns an initialized PositiveFloat64
func (p PositiveFloat64) New() (PositiveFloat64, int) {
	p.Value = 0.0
	p.Legal = true
	p.index = 0
	p.sampleFloat64s = []float64{
		-math.MaxFloat64,
		-math.MaxFloat64 + 1,
		-math.MaxFloat64 / 2,
		-math.MaxFloat64/2 + 1,
		-3, -2, -1, 0, 1, 2, 3,
		math.MaxFloat64/2 - 1,
		math.MaxFloat64 / 2,
		math.MaxFloat64 - 1,
		math.MaxFloat64,
	}
	return p, len(p.sampleFloat64s)
}

// Next returns a counter, the next value from the list of samples, and true if
// the value is positive.
func (p *PositiveFloat64) Next() (float64, bool) {

	p.Value = p.sampleFloat64s[p.index]
	p.index++

	if p.Value > 0.0 {
		p.Legal = true
	} else {
		p.Legal = false
	}
	return p.Value, p.Legal
}
