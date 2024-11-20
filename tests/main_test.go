package main

import (
	"github.com/davecb/pdqwrapper/tests/testIterator"
	"testing"
	"time"
)

// The System Under Test (SUT) is func pdq(progName string, think, serviceDemand, from, to, by float64, verbose, debug bool)
// t is a positive number, upper bound unknown
// z sleep is synonymous with t
// s is service time, positive number, upper bound unknown
// from is a small positive number, defaulting to 1
// to is a small positive number, invalid if less than from, probably limited to 1,000
// by is a small positive number, smaller than to-from
// v is a verbose flag
// d is a debug flag
// h is a usage flag, causing an exit

// Test_pdq test the top-level non-main function
func Test_pdq(t *testing.T) {

	// initialize vars, in part for use in developing the loops
	var z, s, from, to, by int
	var zStruct = testIterator.FloatSample{1.0, true}
	var sStruct = testIterator.FloatSample{1.0, true}
	var fromStruct = testIterator.IntSample{1, true}
	var toStruct = testIterator.IntSample{10, true}
	var byStruct = testIterator.IntSample{1, true}
	const verbose = false
	const debug = false

	var start = time.Now()
	var count int

	for s, sStruct = range testIterator.PositiveFloat {
		// s is service time, a positive number, upper bound unknown
		for z, zStruct = range testIterator.PositiveFloat {
			// think time "t" or "z" is a positive number, upper bound unknown
			for from, fromStruct = range testIterator.SmallPositiveCounter {
				// from is the initial load
				for to, toStruct = range testIterator.SmallPositiveCounter {
					// to is the final load,
					for by, byStruct = range testIterator.SmallPositiveCounter {
						//by is the step size

						// inner test
						count++
						legal := testIterator.AllTrue(sStruct.Legal, zStruct.Legal, fromStruct.Legal, toStruct.Legal, byStruct.Legal)
						t.Logf("debug, with z(%d) == %v %t, s(%d) == %v %t, from(%d) == %d %t, to(%d) == %d %t, by(%d) == %d %t, legal == %t\n",
							z, zStruct.Value, zStruct.Legal, s, sStruct.Value, sStruct.Legal, from, fromStruct.Value, fromStruct.Legal, to, toStruct.Value, toStruct.Legal, by, byStruct.Value, byStruct.Legal, legal)
						err := pdq("unit test", zStruct.Value, sStruct.Value, float64(fromStruct.Value), float64(toStruct.Value), float64(byStruct.Value), verbose, debug)
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
	t.Logf("%d tests run in %v\n", count, time.Since(start))
}
