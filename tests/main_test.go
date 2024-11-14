package main

import (
	"math"
	"testing"
)

// func Pdq(progName string, think, serviceDemand, from, to, by float64, verbose, debug bool) {
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

	type Sample struct {
		Value int64
		Legal bool
	}

	positiveInt := []Sample{
		{-math.MaxInt64, false},
		{-math.MaxInt64 + 1, false},
		{-math.MaxInt64 / 2, false},
		{-(math.MaxInt64 / 2) + 1, false},
		{-3, false},
		{-2, false},
		{-1, false},
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{(math.MaxInt64 / 2) - 1, true},
		{(math.MaxInt64 / 2), true},
		{math.MaxInt64 - 1, true},
		{math.MaxInt64, true},
	}

	for z, zStruct := range positiveInt {
		// think time "t" is a positive number, upper bound unknown
		// sleep time "z" is synonymous with t
		//fmt.Printf("z%d: Value = %d, Legal = %t\n",
		//	z, zStruct.Value, zStruct.Legal)
		for s, sStruct := range positiveInt {
			// s is service time, a positive number, upper bound unknown
			//fmt.Printf("s%d: Value = %d, Legal = %t\n",
			//	s, sStruct.Value, sStruct.Legal)
			// inner test
			err := Pdq("unit test", float64(zStruct.Value), float64(sStruct.Value), 1, 2, 1, false, false)
			if err != nil {
				t.Errorf("fail with z(%d) = %g, s(%d), %g\n",
					z, float64(zStruct.Value), s, float64(sStruct.Value))
			}

		}
	}
}
