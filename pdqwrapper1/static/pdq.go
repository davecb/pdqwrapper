package main

import "C"
import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unsafe"
)

// FIXME, ${SRCDIR} is not recognized by cgo in the following

/*
#cgo  LDFLAGS: -static -L/home/davecb/go/src/github.com/davecb/pdqwrapper/pdqwrapper1/pdq5/lib -lpdq
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
#include "/home/davecb/go/src/github.com/davecb/pdqwrapper/pdqwrapper1/pdq5/lib/PDQ_Lib.h"
double floor(double x) {
	double intpart;

   if (x >= 0.0) {
        return (double)((long long)x);
    }
    intpart = (double)((long long)x);
    return (intpart == x) ? x : intpart - 1;
}
double ceil(double x) {
 	 double intpart;
    if (x <= 0.0) {
        return (double)((long long)x);
    }
    intpart = (double)((long long)x);
    return (intpart == x) ? x : intpart + 1;
}
*/
import "C"

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
	fmt.Printf("General closed solution from PDQ where serviceTime = %g thinkTime time = %g\n",
		serviceTime, thinkTime)
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
	const (
		TERM  = 11
		CEN   = 4
		FCFS  = 8
		EXACT = 14
	)

	var s *C.char = C.CString("closed uniserver")
	defer C.free(unsafe.Pointer(s))
	C.PDQ_Init(s)

	var modelName *C.char = C.CString("work")
	defer C.free(unsafe.Pointer(modelName))
	C.PDQ_CreateClosed(modelName, TERM, C.double(load), C.double(thinkTime)) // could not determine kind of name for C.CreateClosed

	var nodeName *C.char = C.CString("server0")
	defer C.free(unsafe.Pointer(nodeName))
	C.PDQ_CreateNode(nodeName, CEN, FCFS)

	C.PDQ_SetDemand(nodeName, modelName, C.double(serviceTime))

	C.PDQ_Solve(EXACT)

	var work *C.char = C.CString("work")
	defer C.free(unsafe.Pointer(work))

	fmt.Printf("%d\t%f\t%f\t%f\t%f\t%f\n",
		int(load),
		C.PDQ_GetThruput(TERM, work),
		C.PDQ_GetUtilization(nodeName, work, TERM),
		C.PDQ_GetQueueLength(nodeName, work, TERM),
		C.PDQ_GetResidenceTime(nodeName, work, TERM),
		C.PDQ_GetResponse(TERM, work))

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

/*
intyrinsics missing
double exact(double d) {
   return 0.0;
}

*/
