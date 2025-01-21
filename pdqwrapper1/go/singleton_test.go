package main

import (
	"flag"
	"testing"
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

var z, s, to, from, by float64
var legal bool

func init() {
	flag.Float64Var(&z, "z", 0.0, "sleepTime value")
	flag.Float64Var(&s, "s", 0.0, "serviceTime value")
	flag.Float64Var(&from, "from", 0, "from")
	flag.Float64Var(&to, "to", 0.0, "to")
	flag.Float64Var(&by, "by", 0.0, "by")
	flag.BoolVar(&legal, "legal", false, "legal flag value")
}

func TestSingleton(t *testing.T) {
	t.Logf("z = %v\n", z)
	t.Logf("s = %v\n", s)
	t.Logf("from = %v\n", from)
	t.Logf("to = %v\n", to)
	t.Logf("nb = %v\n", by)
	t.Logf("legal = %v\n", legal)

	err := pdq("unit test", z, s, float64(from), float64(to), float64(by), -1, legal)
	if err != nil {
		// failure case
		if legal {
			// provide detailed information about the failure, then stop.
			t.Errorf("missing success, with err == '%v', z == %v, s == %v, from == %v, to == %v, by == %v, legal == %t\n",
				err, z, s, from, to, by, legal)
		}
	} else {
		// success case
		if !legal {
			t.Errorf("missing failure, with err == '%v', z == %v, s == %v, from == %v, to == %v, by == %v, legal == %t\n",
				err, z, s, from, to, by, legal)
		}
	}
}
