package main

import (
	"fmt"
	"log"
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
		err         error
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
			thinkTime, err = strconv.ParseFloat(os.Args[i], 64)
			if err != nil {
				log.Fatalf("parseFloat failed, can't happen. err = %v\n", err)
			}
		case 's':
			i++
			serviceTime, err = strconv.ParseFloat(os.Args[i], 64)
			if err != nil {
				log.Fatalf("parseFloat failed, can't happen. err = %v\n", err)
			}
		default:
			fmt.Fprintf(os.Stderr, "%s: unknown option -%c, ignored.\n", progName, os.Args[i][1])
		}
		i++
	}

	// Parse remaining arguments
	if i < len(os.Args) {
		from, err = strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			log.Fatalf("parseFloat failed, can't happen. err = %v\n", err)
		}
		i++
	}
	if i < len(os.Args) {
		to, err = strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			log.Fatalf("parseFloat failed, can't happen. err = %v\n", err)
		}
		i++
	}
	if i < len(os.Args) {
		by, err = strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			log.Fatalf("parseFloat failed, can't happen. err = %v\n", err)
		}
	}
	err = pdq(progName, thinkTime, serviceTime, from, to, by, -1, true)
	if err != nil {
		log.Fatalf("pdq error = %v, halting", err)
	}
}

// pdq is the code that calls the pdq library. it is distinct from main
func pdq(progName string, thinkTime, serviceTime, from, to, by float64, line int, legal bool) error {
	args := fmt.Sprintf("-z=%v -s=%v -from=%v -to=%v -by=%v legal=%v", thinkTime, serviceTime, from, to, by, legal)
	if line == -1 {
		// don't report line
		fmt.Printf("General closed solution from PDQ where %s\n", args)
	} else {
		fmt.Printf("General closed solution from PDQ where line=%d %s\n", line, args)
	}
	fmt.Printf("Load\tThroughput\tUtilization\tQueueLen\tResidence\tResponse\n")

	// Check parameters
	if thinkTime <= 0.0 {
		return fmt.Errorf("%s: thinkTime == %g, which is non-positive and not valid: %s", progName, thinkTime, args)
	}
	if serviceTime <= 0.0 {
		return fmt.Errorf("%s: serviceTime == %g, which is non-positive and not valid: %s", progName, serviceTime, args)
	}
	if from <= 0.0 {
		return fmt.Errorf("%s: from == %g, which is non-positive and not valid: %s", progName, from, args)
	}
	if to <= 0.0 {
		return fmt.Errorf("%s: to == %g, which is non-positive and not valid: %s", progName, to, args)
	}
	if by == 0 {
		by = 1 // Silly heuristic, but for a common usage error
	}
	if by <= 0.0 {
		return fmt.Errorf("%s: by == %g, which is non-positive and not valid: %s", progName, to, args)
	}
	// if there are interrelationship limits, test them here

	// FIXME, that was just plain wrong! Tried for EXACT
	//steps := (to - from)
	//if steps < 0.0 || steps/by > 999 {
	//	return fmt.Errorf("%s: (to-from)/by  == %g, which is 1000 or more, and not valid: %s", progName, steps/by, args)
	//}
	//}

	for load := from; load <= to; load += by {
		doOneStep(load, thinkTime, serviceTime)
	}
	return nil
}

// doOneStep does a model for one value of load and exits with an error on failure.
// the latter is a feature of the library
func doOneStep(load, thinkTime, serviceTime float64) {
	const (
		TERM   = 11
		CEN    = 4
		FCFS   = 8
		EXACT  = 14
		APPROX = 15
	)
	var (
		s         *C.char = C.CString("closed uniserver")
		modelName *C.char = C.CString("work")
		nodeName  *C.char = C.CString("server0")
		work      *C.char = C.CString("work")
	)
	defer C.free(unsafe.Pointer(s))
	defer C.free(unsafe.Pointer(modelName))
	defer C.free(unsafe.Pointer(nodeName))
	defer C.free(unsafe.Pointer(work))

	C.PDQ_Init(s)
	C.PDQ_CreateClosed(modelName, TERM, C.double(load), C.double(thinkTime))
	C.PDQ_CreateNode(nodeName, CEN, FCFS)
	C.PDQ_SetDemand(nodeName, modelName, C.double(serviceTime))

	//if approximate FIXME, make this conditional
	C.PDQ_Solve(APPROX)
	//} else {
	//	C.PDQ_Solve(EXACT)
	//}
	fmt.Printf("%d\t%f\t%f\t%f\t%f\t%f\n",
		int(load),
		C.PDQ_GetThruput(TERM, work),
		C.PDQ_GetUtilization(nodeName, work, TERM),
		C.PDQ_GetQueueLength(nodeName, work, TERM),
		C.PDQ_GetResidenceTime(nodeName, work, TERM),
		C.PDQ_GetResponse(TERM, work))
}
