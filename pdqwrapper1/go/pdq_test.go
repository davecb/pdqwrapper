package main

//
import (
	"github.com/davecb/pdqwrapper/tests/testIterator"
	"testing"
	"time"
)

// The System Under Test (SUT) is func pdq(progName string, think, serviceDemand, from, to, by float64, verbose, debug bool)
// t, think, is a positive number, upper bound unknown
// z sleep is synonymous with t
// s is service time, positive number, upper bound unknown
// from is a small positive number, defaulting to 1
// to is a small positive number, invalid if less than from, probably limited to 1,000
// by is a small positive number, smaller than to-from
// v is a verbose flag
// d is a debug flag
// h is a usage flag, causing an exit

// Test_pdq test the top-level non-main function. This tests everything
// except startup and command-line parsing done in main()
func TestPdq(t *testing.T) {

	var z, s, from, to, by int
	var zStruct = testIterator.FloatSample{}
	var sStruct = testIterator.FloatSample{}
	var fromStruct = testIterator.IntSample{}
	var toStruct = testIterator.IntSample{}
	var byStruct = testIterator.IntSample{}

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
						// any composite tests go here.
						steps := (to - from)
						switch {
						case steps <= 0:

							legal = false
						case by <= 0:
							legal = false
						case steps/by > 999:
							legal = false
						}

						err := pdq("unit test", zStruct.Value, sStruct.Value, float64(fromStruct.Value), float64(toStruct.Value), float64(byStruct.Value), count, legal)
						if err != nil {
							// failure case
							if legal {
								// provide detailed information about the failure, then stop.
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
