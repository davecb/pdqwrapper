package main

import (
	"math"
	"testing"
)

// func Wrapper(progName string, think, serviceDemand, from, to, by float64, verbose, debug bool) {
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
	//type intSample struct {
	//	Value int64
	//	Legal bool
	//}
	type floatSample struct {
		Value float64
		Legal bool
	}
	//var positiveInt []intSample
	var positiveFloat []floatSample

	//positiveInt = []intSample{
	//	{-math.MaxInt64, false},
	//	{-math.MaxInt64 + 1, false},
	//	{-math.MaxInt64 / 2, false},
	//	{-(math.MaxInt64 / 2) + 1, false},
	//	{-3, false},
	//	{-2, false},
	//	{-1, false},
	//	{0, false},
	//	{1, true},
	//	{2, true},
	//	{3, true},
	//	{4, true},
	//	{(math.MaxInt64 / 2) - 1, true},
	//	{(math.MaxInt64 / 2), true},
	//	{math.MaxInt64 - 1, true},
	//	{math.MaxInt64, true},
	//}
	positiveFloat = []floatSample{
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

	for z, zStruct := range positiveFloat {
		// think time "t" is a positive number, upper bound unknown
		// sleep time "z" is synonymous with t
		for s, sStruct := range positiveFloat {
			// s is service time, a positive number, upper bound unknown
			for from, fromStruct := range positiveFloat {
				// from is the initial load
				for to, toStruct := range positiveFloat {
					// to is the final load
					// by is the step size

					// inner test
					legal := allTrue(zStruct.Legal, sStruct.Legal, fromStruct.Legal, toStruct.Legal)
					err := Wrapper("unit test", zStruct.Value, sStruct.Value, fromStruct.Value, toStruct.Value, 10, false, false)
					t.Logf("err == %s, with z(%d) == %g, s(%d) == %g, from(%d) == %g\n",
						err, z, zStruct.Value, s, sStruct.Value, from, fromStruct.Value)
					if err != nil {
						// failure case
						if legal {
							t.Fatalf("err == %v, z(%d) == %g, s(%d) == %g, from(%d) == %g, to(%d) == %g\n",
								err, z, zStruct.Value, s, sStruct.Value, from, fromStruct.Value, to, toStruct.Value)
						}
					} else {
						// success case
						if !legal {
							t.Fatalf("missing failure, with z(%d) == %g, s(%d) == %g, from(%d) == %g\n",
								z, zStruct.Value, s, sStruct.Value, from, fromStruct.Value)
						}
					}
				}
			}
		}
	}
}

// allTrue looks for any false values in a vector of booleans
func allTrue(b ...bool) bool {
	for _, v := range b {
		if !v {
			return false
		}
	}
	return true
}
