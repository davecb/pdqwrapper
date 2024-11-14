package main

import "os"

func main() {
	os.Exit(0)
}

//
//import (
//	"fmt"
//	"os"
//	"path/filepath"
//	"strconv"
//)
//
//// PDQ constants (you'll need to implement or import these)
//const (
//	TERM  = 0
//	CEN   = 1
//	FCFS  = 2
//	EXACT = 3
//)
//
//var (
//	nodes   int
//	streams int
//)
//
//func usage(progName string) {
//	fmt.Fprintf(os.Stderr, "Usage: %s [-z think][-s service][-vd] from to by\n", progName)
//}
//
//func main() {
//	var (
//		from          = 1.0
//		to            = 0.0
//		by            = 0.0
//		think         = 0.0
//		serviceDemand = 0.0
//		verbose       = false
//		debug         = false
//	)
//
//	progName := filepath.Base(os.Args[0])
//
//	if len(os.Args) == 1 {
//		usage(progName)
//		os.Exit(1)
//	}
//
//	// Parse command line arguments
//	i := 1
//	for i < len(os.Args) && os.Args[i][0] == '-' {
//		switch os.Args[i][1] {
//		case 'h':
//			usage(progName)
//			os.Exit(0)
//		case 'z', 't':
//			i++
//			think, _ = strconv.ParseFloat(os.Args[i], 64)
//		case 'd', 'x':
//			debug = true
//		case 's':
//			i++
//			serviceDemand, _ = strconv.ParseFloat(os.Args[i], 64)
//		case 'v':
//			verbose = true
//		default:
//			fmt.Fprintf(os.Stderr, "%s: unknown option -%c, ignored.\n", progName, os.Args[i][1])
//		}
//		i++
//	}
//
//	// Parse remaining arguments
//	if i < len(os.Args) {
//		from, _ = strconv.ParseFloat(os.Args[i], 64)
//		i++
//	}
//	if i < len(os.Args) {
//		to, _ = strconv.ParseFloat(os.Args[i], 64)
//		i++
//	}
//	if i < len(os.Args) {
//		by, _ = strconv.ParseFloat(os.Args[i], 64)
//	}
//	pdq(progName, think, serviceDemand, from, to, by, verbose, debug)
//}
//
//func pdq(progName string, think, serviceDemand, from, to, by float64, verbose, debug bool) {
//	// Check parameters
//	if from < 0.0 {
//		fmt.Fprintf(os.Stderr, "%s: from is negative, which is not defined. Halting.", progName)
//		os.Exit(1)
//	}
//	if from == 0.0 {
//		from = 1.0
//	}
//	if to <= 0.0 {
//		to = from
//	}
//	if by <= 0.0 {
//		by = 1.0
//	}
//
//	//// Print headers
//	//fmt.Printf("General closed solution from PDQ where serviceTime = %g think time = %g\n",
//	//	serviceDemand, think)
//	//
//	//if verbose {
//	//	fmt.Printf("Load\tThroughput\tUtilization\tQueueLen\tResidence\tResponse\n")
//	//} else {
//	//	fmt.Printf("\"# Load,\" Response\n")
//	//}
//
//	for load := from; load <= to; load += by {
//		doOneStep(load, think, serviceDemand, verbose)
//	}
//
//}
//
//func doOneStep(load, think, serviceDemand float64, verbose bool) {
//	fmt.Fprintf(os.Stderr, "load = %g think = %g "+
//		"serviceTime = %g verbose = %t\n", load, think, serviceDemand, verbose)
//}
