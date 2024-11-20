package pdqWrapper

// this simulates cc -g -o ./pdq pdq.c ../pdq5/lib/*.o -lm
// see also http://www.perfdynamics.com/Tools/PDQman.html

// #include <stdio.h>
// #include <math.h>
// #include "pdq5/lib/PDQ_Lib.h"
import "C"

const (
	TERM  = 0
	CEN   = 1
	FCFS  = 2
	EXACT = 3
)

func Random() int {
	return int(C.random())
}

// Init is the startup function for the pdq library
func Init(s string) {
	C.PDQ_Init(C.string(s))
	// eg, PDQ_Init("closed uniserver")
}

// CreatClosed creates a  closed queue model, with a name, load and think-time
func CreateClosed(modelName string, load, think float64) {
	C.CreateClosed(modelName, C.double(load), C.double(think))
	// Caution, the library used to return a count, which is unused
	// the modelName is used several places. For a simple model, use "work"
	// eg, 	modelName = "work"; streams = PDQ_CreateClosed(modelName, TERM, load, think)
}

// CreateNodes creates a named server node
func CreateNode(serverName string) {
	C.CreateNode(C.string(serverName), C.int(CEN), C.int(FCFS))
	// Caution, the library used to return a count, which is unused
	//the nodeName is used several places. For a simple model, use "server0"
	//eg nodeName = "server0"; nodes = PDQ_CreateNode(nodeName, CEN, FCFS)
}

// SetDemand set the service-time, AKA service-demand, of a specified server and model
func SetDemand(nodeName, modelName string, serviceTime float64) {
	C.PDQ_SetDemand(C.string(nodeName), C.string(modelName), C.double(serviceTime))
	// eg, PDQ_SetDemand("server0", "work", serviceDemand)
}

// Solve runs the pdq modeller
func Solve() {
	C.PDQ_Solve(C.int(EXACT))
}

// Report is the data from a single run at a given load
type Report struct {
	Load          float64
	Throughput    float64
	Utilization   float64
	QueueLength   float64
	ResidenceTime float64
	ResponseTime  float64
}

// Results returns the dat for a single specified losd
func Results(load float64) Report {
	return Report{
		load,
		C.PDQ_GetThruput(C.int(TERM), C.string("work")),
		C.PDQ_GetUtilization("server0", "work", TERM),
		C.PDQ_GetQueueLength("server0", "work", TERM),
		C.PDQ_GetResidenceTime("server0", "work", TERM),
		C.PDQ_GetResponse(TERM, "work"),
	}
}

/*
	in C:

func doOneStep(load, think, serviceDemand float64, verbose bool) {
 void doOneStep(double load, double think, double serviceTime, int verbose) {
	extern int	nodes, streams;
	static char server_name[80] = "";

	PDQ_Init("closed uniserver")
	streams = PDQ_CreateClosed("work", TERM, load, think)

	nodeName = "server0"
	nodes = PDQ_CreateNode(nodeName, CEN, FCFS)
	PDQ_SetDemand(nodeName, "work", serviceDemand)

	PDQ_Solve(EXACT)

	if verbose {
		fmt.Printf("%d\t%f\t%f\t%f\t%f\t%f\n",
			int(load),
			PDQ_GetThruput(TERM, "work"),
			PDQ_GetUtilization("server0", "work", TERM),
			PDQ_GetQueueLength("server0", "work", TERM),
			PDQ_GetResidenceTime("server0", "work", TERM),
			PDQ_GetResponse(TERM, "work"))
	} else {
		fmt.Printf("%d,\t%f\n", int(load), PDQ_GetResponse(TERM, "work"))
	}
}
in Go:
func doOneStep(load, thinkTime, serviceTime float64) {

	s := "closed uniserver"
	Init(s)

	modelName := "work"
	CreateClosed(modelName, load, thinkTime) {

	nodeName := "server0"
	CreateNode(nodeName)

	SetDemand(nodeName, modelName, serviceTime)

	Solve() // FIXME exact?
	r := Result(load)
	fmt.Printf("%d\t%f\t%f\t%f\t%f\t%f\n",
		int(load),
		PDQ_GetThruput(TERM, "work"),
		PDQ_GetUtilization("server0", "work", TERM),
		PDQ_GetQueueLength("server0", "work", TERM),
		PDQ_GetResidenceTime("server0", "work", TERM),
		PDQ_GetResponse(TERM, "work"))
}
*/
