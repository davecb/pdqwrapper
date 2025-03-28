/*
 * pdq -- report on a closed queuing centre, based on
 *	closed_center.c from Gunther. This is what I use
 *	from the command-line.
 *	 
 */
#include <stdio.h>
#include <stdlib.h>	/* For exit(). */
#include <libgen.h>	/* For basename(). */
#include <math.h>
/* #include "/usr/local/pdq/lib/PDQ_Lib.h"  -- old location*/
/* #include "/home/davecb/projects/PDQ2/pdq/pdq5/lib/PDQ_Lib.h" */
#include "../pdq5/lib/PDQ_Lib.h"

#define PREFORK
#define STRESS	0
#define HOMEPG	1

static char *ProgName = NULL;
void doOneStep(double load, double think, double serviceTime, int verbose);

 void
usage() {
	fprintf(stderr, "Usage: %s [-z sleepTime][-s serviceTime][-vd] "
		"from to by\n", ProgName);
}

 int
main(int argc, char **argv) {
	double	from = 1.0,
		to = 0.0,
		by = 0.0,
		sleep = 0.0,
		serviceTime = 0.0, 
		load;
	int	verbose = 0,
		debug = 0;
	int	i;


	ProgName = basename(argv[0]);
	if (argc == 1) {
		usage();
		exit(1);
	}
	for (i=1; i < argc; i++) {
		/* printf("argv[%d] = %s\n", i, argv[i]); */
		if (argv[i][0] == '-') {
			switch (argv[i][1]) {
			case 'z':
			case 't':
			    sleep = atof(argv[++i]);
				break;
			case 'h':
				usage(ProgName);
				exit(0);
			case 'd': debug = 1;
				break;
			case 's': serviceTime = atof(argv[++i]);
				break;
			case 'v': verbose = 1;
				break;
			default:
				(void) fprintf(stderr,
				       "%s: unknown option -%c, ignored.\n",
				       ProgName, argv[i][1]);
			}
		}
		else {
			break;
		}
	}
	/* Check options */
	if (serviceTime <= 0) {
	    (void) fprintf(stderr, "%s: -s is <= 0.0 which is not supported. Halting.\n",
            ProgName);
        	exit(1);
	}
	if (sleep < 0.0) {
	    (void) fprintf(stderr, "%s: -z is < 0.0 which is not supported. Halting.\n",
                ProgName);
            	exit(1);
	}


	/* collect from to and by parameters */
	if (i < argc) {
		from = atof(argv[i++]);
	}
	if (i < argc) {
		to = atof(argv[i++]);
	}
	if (i < argc) {
		by = atof(argv[i]);
	}

	/* Check parameters. */
	if (from < 0.0) {
		(void) fprintf(stderr, "%s: from is negative, which is not defined. Halting.",
		    ProgName);
		    exit(1);
	}
    if (from == 0.0) {
        from = 1.0;
    }
	if (to <= 0.0) { 	/* FIXME check that we have less than 1000 users */
		to = from;
	}
	if (by <= 0.0) {
		by = 1.0;
	}

	if (debug == 1) {
    	    (void) printf(
                "service time = %g "
                "sleep time = %g "
                "from %g "
                "to = %g "
                "by = %g\n",
    	        serviceTime, sleep, from, to, by);
    }
    /* Print headers. */
    printf("General closed solution from PDQ where "
    	"serviceTime = %g sleep time = %g\n",
           serviceTime, sleep);
    if (verbose) {
    	printf("Load\tThroughput\tUtilization\tQueueLen\t"
    		"Residence\tResponse\n");
    }
    else {
    	printf("Load\tResponse\n");
    }
	for (load=from; load <= to; load += by) {
		doOneStep(load, sleep, serviceTime, verbose);
	}
	/* if (debug == 1) {
	 *	PDQ_Report(); optional
	 * }
	 */
	exit(0);
}


/*
 * doOneStep -- do one solution step
 */
 void
doOneStep(double load, double sleep, double serviceTime, int verbose) {
	extern int	nodes, streams;
	static char server_name[80] = "";
	int	i;

	/* optionally, one can set PDQ_SetDebug(1); */
	PDQ_Init("closed uniserver"); /* Name model. */

	/* Define workload and queuing circuit type. */
	streams = PDQ_CreateClosed("work", TERM, load, sleep); // * TERM, etc is defined in .h file */

	/* Create a single node, with a demand of serviceTime */
	(void) sprintf(server_name, "server0");
	nodes = PDQ_CreateNode(server_name, CEN, FCFS);
	PDQ_SetDemand(server_name, "work", serviceTime);
	PDQ_Solve(EXACT); /* EXACT allows a load of <= 1000 users */

	if (verbose) {
		printf("%d\t%f\t%f\t%f\t%f\t%f\n",
			(int) load,
			PDQ_GetThruput(TERM, "work"),
	 		PDQ_GetUtilization("server0", "work", TERM),
			PDQ_GetQueueLength("server0", "work", TERM),
			PDQ_GetResidenceTime("server0", "work", TERM),
			PDQ_GetResponse(TERM, "work"));
	}
	else {
		printf("%d\t%f\n",(int) load, PDQ_GetResponse(TERM, "work"));
	}

} /* doOneStep */

