package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// PDQ constants (you'll need to implement or import these)
const (
	TERM  = 0
	CEN   = 1
	FCFS  = 2
	EXACT = 3
)

var (
	nodes   int
	streams int
)

func usage(progName string) {
	fmt.Fprintf(os.Stderr, "Usage: %s [-z think][-s service][-vd] from to by\n", progName)
}

func main() {
	var (
		from          = 1
		to            = 0
		by            = 0
		think         = 0.0
		serviceDemand = 0.0
		verbose       = false
		debug         = false
	)

	progName := filepath.Base(os.Args[0])

	if len(os.Args) == 1 {
		usage(progName)
		os.Exit(1)
	}

	// Parse command line arguments
	i := 1
	for i < len(os.Args) && os.Args[i][0] == '-' {
		switch os.Args[i][1] {
		case 'h':
			usage(progName)
			os.Exit(0)
		case 'z', 't':
			i++
			think, _ = strconv.ParseFloat(os.Args[i], 64)
		case 'd', 'x':
			debug = true
		case 's':
			i++
			serviceDemand, _ = strconv.ParseFloat(os.Args[i], 64)
		case 'v':
			verbose = true
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option -%c, ignored.\n", progName, os.Args[i][1])
		}
		i++
	}

	// Parse remaining arguments FIXME, these are floats
	if i < len(os.Args) {
		from, _ = strconv.Atoi(os.Args[i])
		i++
	}
	if i < len(os.Args) {
		to, _ = strconv.Atoi(os.Args[i])
		i++
	}
	if i < len(os.Args) {
		by, _ = strconv.Atoi(os.Args[i])
	}
	Wrapper(progName, think, serviceDemand, from, to, by, verbose, debug)
}

// Wrapper is the code that calls the pdq library
func Wrapper(progName string, thinkTime, serviceTime float64, from, to, by int, verbose, debug bool) error {
	// Check parameters
	// FIXME can these be floats??? The library accepts them, so yes. Future extension
	if from < 0 {
		return fmt.Errorf("%s: \"from\" is negative, which is not defined. Halting.", progName)
	}
	if from == 0 {
		// from is only well-defined for positives, but choosing is a common, harmless user error
		from = 1
	}
	if to <= 0 {
		// this is bad code!   FIXME in the refactor, this was dumb
		// FIXME also check that we have less than 1000 users
		to = from
	}
	if by <= 0 {
		// FIXME ditto
		by = 1
	}

	// Print headers
	fmt.Printf("General closed solution from PDQ where serviceTime = %g thinkTime time = %g\n",
		serviceTime, thinkTime)

	if verbose {
		fmt.Printf("Load\tThroughput\tUtilization\tQueueLen\tResidence\tResponse\n")
	} else {
		fmt.Printf("\"# Load,\" Response\n")
	}

	for load := from; load <= to; load += by {
		doOneStep(load, thinkTime, serviceTime, verbose)
	}
	return nil
}

func doOneStep(load int, thinkTime, serviceTime float64, verbose bool) {
	fmt.Fprintf(os.Stderr, "load = %d thinkTime = %g "+
		"serviceTime = %g verbose = %t\n", load, thinkTime, serviceTime, verbose)
}
