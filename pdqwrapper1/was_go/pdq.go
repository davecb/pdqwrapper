package main

import (
	"fmt"
	"github.com/davecb/pdqwrapper/tests/pdqWrapperDirs/pdqWrapper"
	"os"
	"path/filepath"
	"strconv"
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
		thinkTime   = 0.0
		serviceTime = 0.0
		from        = 1.0
		to          = 0.0
		by          = 0.0
		verbose     = false // unused
	)

	progName := filepath.Base(os.Args[0])

	if len(os.Args) == 1 {
		usage(progName)
		os.Exit(1)
	}

	// Parse command line arguments  FIXME, use flags package
	i := 1
	for i < len(os.Args) && os.Args[i][0] == '-' {
		switch os.Args[i][1] {
		case 'h':
			usage(progName)
			os.Exit(0)
		case 'z', 't':
			i++
			thinkTime, _ = strconv.ParseFloat(os.Args[i], 64)
		case 's':
			i++
			serviceTime, _ = strconv.ParseFloat(os.Args[i], 64)
		case 'v':
			verbose = true
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option -%c, ignored.\n", progName, os.Args[i][1])
		}
		i++
	}

	// Parse remaining arguments
	if i < len(os.Args) {
		from, _ = strconv.ParseFloat(os.Args[i], 64)
		i++
	}
	if i < len(os.Args) {
		to, _ = strconv.ParseFloat(os.Args[i], 64)
		i++
	}
	if i < len(os.Args) {
		by, _ = strconv.ParseFloat(os.Args[i], 64)
	}

	// Check parameters
	if from < 0.0 {
		fmt.Fprintf(os.Stderr, "%s: from is negative, which is not defined. Halting.", progName)
		os.Exit(1)
	}
	if from == 0.0 {
		from = 1.0
	}
	if to <= 0.0 {
		to = from
	}
	if by <= 0.0 {
		by = 1.0
	}

	// Print headers for call to shared library
	fmt.Printf("General closed solution from PDQ where serviceTime = %g thinkTime time = %g from = %g, for = %g by=%g\n",
		serviceTime, thinkTime, from, to, by)
	if verbose {
		fmt.Printf("Load\tThroughput\tUtilization\tQueueLen\tResidence\tResponse\n")
	} else {
		fmt.Printf("\"# Load,\" Response\n")
	}

	for load := from; load <= to; load += by {
		doOneStep(load, thinkTime, serviceTime)
	}
}

func doOneStep(load, thinkTime, serviceTime float64) {

	s := "closed uniserver"
	pdqWrapper.Init(s)

	modelName := "work"
	pdqWrapper.CreateClosed(modelName, load, thinkTime) {

		nodeName := "server0"
		pdqWrapper.CreateNode(nodeName)

		pdqWrapper.SetDemand(nodeName, modelName, serviceTime)

		pdqWrapper.Solve() // FIXME exact?
		r := pdqWrapper.Results(load)
		fmt.Printf("%f\t%f\t%f\t%f\t%f\t%f\n",
			load,
			r.Throughput,
			r.Utilization,
			r.QueueLength,
			r.ResidenceTime,
			r.ResponseTime)
	}
}

/*****
Note that this translation assumes the existence of PDQ-related functions that would need to be implemented or imported from a Go PDQ library. The functions that need to be implemented include:

PDQ_Init
PDQ_CreateClosed
PDQ_CreateNode
PDQ_SetDemand
PDQ_Solve
PDQ_GetThruput
PDQ_GetUtilization
PDQ_GetQueueLength
PDQ_GetResidenceTime
PDQ_GetResponse
PDQ_Report

see https://karthikkaranth.me/blog/calling-c-code-from-go/
and
static:
    gcc -c gb.c
    ar -rcs libgb.a gb.o
    go build -ldflags "-linkmode external -extldflags -static" bridge.go

dynamic:
    gcc -shared -o libgb.so gb.c
    go build bridge.go
*/
