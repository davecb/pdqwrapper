package main

import (
	"log"
	"math"
	"testing"
)

// The SUT is func Wrapper(progName string, think, serviceDemand, from, to, by float64, verbose, debug bool)
// t is a positive number, upper bound unknown
// z sleep is synonymous with t
// s is service time, positive number, upper bound unknown
// v is a verbose flag
// d is a debug flag
// h is a usage flag, causing an exit
// from is a positive number, defaulting to 1
// to is a positive number, invalid if less than from
// by is a positive number, smaller than to-from

// Test_Wrapper test the top-level non-main function
func Test_Wrapper(t *testing.T) {

	// inlined value-generators. FIXME make methodical (:-))
	type floatSample struct {
		Value float64
		Legal bool
	}
	type intSample struct {
		Value int
		Legal bool
	}

	// a sample set for positive. Notice e inject some decimal
	// values, not just integer values
	var PositiveFloat = []floatSample{
		{-math.MaxFloat64 + 0, false},
		{-math.MaxFloat64 + 1.1, false},
		{-math.MaxFloat64/2 + .2, false},
		{-(math.MaxFloat64 / 2) + 1.3, false},
		{-3.4, false},
		{-2.5, false},
		{-1.6, false},
		{0, false},
		{1.7, true},
		{2.8, true},
		{3.9, true},
		{4, true},
		{(math.MaxFloat64 / 2) - 1, true},
		{math.MaxFloat64 / 2, true},
		{math.MaxFloat64 - 1, true},
		{math.MaxFloat64, true},
	}

	// a sample-set for small positive counters
	var SmallPositiveCounter = []intSample{
		{-10, false},
		{-5, false},
		{-3, false},
		{-2, false},
		{-1, false},
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{5, true},
		{10, true},
	}

	// initialize vars, for use developing the loops
	var z, s, from, to, by int
	var zStruct = floatSample{1.0, true}
	var sStruct = floatSample{1.0, true}
	var fromStruct = intSample{1, true}
	var toStruct = intSample{10, true}
	var byStruct = intSample{1, true}

	for s, sStruct = range PositiveFloat {
		// s is service time, a positive number, upper bound unknown
		for z, zStruct = range PositiveFloat {
			// think time "t" or "z" is a positive number, upper bound unknown
			for from, fromStruct = range SmallPositiveCounter {
				// from is the initial load
				for to, toStruct = range SmallPositiveCounter {
					// to is the final load,
					for by, byStruct = range SmallPositiveCounter {
						//by is the step size

						// inner test
						legal := allTrue(zStruct.Legal, sStruct.Legal, fromStruct.Legal, toStruct.Legal, byStruct.Legal)
						t.Logf("debug, with z(%d) == %v %t, s(%d) == %v %t, from(%d) == %d %t, to(%d) == %d %t, by(%d) == %d %t, legal == %t\n",
							z, zStruct.Value, zStruct.Legal, s, sStruct.Value, sStruct.Legal, from, fromStruct.Value, fromStruct.Legal, to, toStruct.Value, toStruct.Legal, by, byStruct.Value, byStruct.Legal, legal)
						err := Wrapper("unit test", zStruct.Value, sStruct.Value, float64(fromStruct.Value), float64(toStruct.Value), float64(byStruct.Value), false, false)
						if err != nil {
							// failure case
							if legal {
								t.Fatalf("missing success, with err == '%v', z(%d) == %v, s(%d) == %v, from(%d) == %v, to(%d) == %v, by(%d) == %v, legal == %t\n",
									err, z, zStruct.Value, s, sStruct.Value, from, fromStruct.Value, to, toStruct.Value, by, byStruct.Value, legal)
							}
						} else {
							// success case
							if !legal {
								t.Fatalf("missing failure, with err == '%v', z(%d) == %v, s(%d) == %v, from(%d) == %v, legal == %t\n",
									err, z, zStruct.Value, s, sStruct.Value, from, fromStruct.Value, legal)
							}
						}
					}
				}
			}
		}
	}
}

// allTrue looks for any false values in a vector of booleans
func allTrue(b ...bool) bool {
	log.Printf("allTrue, b == %v\n", b)
	for _, v := range b {
		if !v {
			return false
		}
	}
	return true
}
