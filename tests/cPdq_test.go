package main

import (
	"fmt"
	"github.com/davecb/pdqwrapper/tests/testIterator"
	"log"
	"os"
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

// Test_cPdq writes a file to use to test the top-level non-main function
// Run with  go test cPdq_test.go to produce cPdq_test.csv
func Test_cPdq(t *testing.T) {

	const verbose = false
	const debug = false
	// initialize vars, in part for use in developing the loops
	var zStruct = testIterator.FloatSample{1.0, true}
	var sStruct = testIterator.FloatSample{1.0, true}
	var fromStruct = testIterator.IntSample{1, true}
	var toStruct = testIterator.IntSample{10, true}
	var byStruct = testIterator.IntSample{1, true}

	fileName := "cPdq_test.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed to create file, error = %q\n", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("failed to close file, error = %q\n", err)
		}
	}()

	var start = time.Now()
	var count int

	// write .csv header
	fmt.Fprintf(file, "#z, s, from, to, by, legal\n")

	for _, sStruct = range testIterator.PositiveFloat {
		// s is service time, a positive number, upper bound unknown
		for _, zStruct = range testIterator.PositiveFloat {
			// think time "t" or "z" is a positive number, upper bound unknown
			for _, fromStruct = range testIterator.SmallPositiveCounter {
				// from is the initial load
				for _, toStruct = range testIterator.SmallPositiveCounter {
					// to is the final load,
					for _, byStruct = range testIterator.SmallPositiveCounter {
						//by is the step size

						// inner test generation step
						count++
						legal := testIterator.AllTrue(sStruct.Legal, zStruct.Legal, fromStruct.Legal, toStruct.Legal, byStruct.Legal)
						fmt.Fprintf(file, "%v, %v, %v, %v, %v, %t\n",
							zStruct.Value, sStruct.Value, float64(fromStruct.Value), float64(toStruct.Value), float64(byStruct.Value), legal)
					}
				}
			}
		}
	}
	t.Logf("file %q, containing %d tests, generated in %v\n",
		fileName, count, time.Since(start))
}
