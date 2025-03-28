 /*
  * this contains the dmax and centers code, for an experiment
 in building a calvin model with a small dmax and a larger set of delay nodes
 which MIGHT work
 */
/*
 * pdq -- report on a closed queuing centre, based on
 *	closed_center.c from Gunther. This is what I use
 *	from the command-line 
 *	 
 */
#include <stdio.h>
#include <stdlib.h>	/* For exit(). */
#include <libgen.h>	/* For basename(). */
#include <math.h>
/* #include "/usr/local/pdq/lib/PDQ_Lib.h" */
/* #include "/home/davecb/projects/PDQ2/pdq/pdq5/lib/PDQ_Lib.h" */
#include "../pdq5/lib/PDQ_Lib.h"

#define PREFORK
#define STRESS	0
#define HOMEPG	1

static char *ProgName = NULL;
void doOneStep(double load, double think, double serviceTime, double dmax, 
	int verbose, int debug);

 void
usage() {
	fprintf(stderr, "Usage: %s -t think -s service[-d dmax][-c centers][-vx] from to by\n",
        ProgName);
}
 double max(double a, double b) {
    if (a >= b) {
        return a;
    }
    else {
        return b;
    }
 }

 int
main(int argc, char **argv) {
	double	from = 1.0,
		to = 0.0,
		by = 0.0,
		think = 0.0,   
		serviceTime = 0.0, 
		dmax = 0.0,
		centers = 1.0,
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
			    think = atof(argv[++i]);
				break;
			case 'x': debug = 1;
				break;
			case 'd': dmax = atof(argv[++i]);
				break;
			case 's': serviceTime = atof(argv[++i]);
				break;
			case 'c': centers = atof(argv[++i]); /* FIXME misleading and redundant */
				break;
			case 'v': verbose = 1;
				break;
			case 'h': 
				usage();
				exit(0);
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
	if (think < 0.0) {
	    (void) fprintf(stderr, "%s: -t is < 0.0 which is not supported. Halting.\n",
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
	if (to <= 0.0) {
		to = from;
	}
	if (by <= 0.0) {
		by = 1.0;
	}

	/* Adjust Dmax if we have more than one center. */
	if (dmax == 0.0 && centers != 1) {
		printf("Dmax must be non-zero for multi-center models.\n");
		exit(3);
	}
	else {
		dmax = dmax / centers;
	}
	if (debug == 1) {
    	    (void) printf(
                "serviceTime = %g "
                "think time = %g "
                "dmax = %g "
                "centers = %g "
                "from = %g "
                "to = %g "
                "by = %g\n",
    	        serviceTime, think, dmax, centers, from, to, by);
    }
    /* Print headers. */
    printf("General closed solution from PDQ where "
    	"serviceTime = %g centers = %g "
           "think time = %g dmax = %g\n",
           serviceTime, centers, think, dmax);
    if (verbose) {
    	printf("Load\tThroughput\tUtilization\tQueueLen\t"
    		"Residence\tResponse\n");
    }
    else {
    	printf("Load\tResponse\n");
    }
	for (load=from; load <= to; load += by) {
		doOneStep(load, think, serviceTime, dmax, verbose, debug);
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
doOneStep(double load, double think, double serviceTime, double dmax, 
	int verbose, int debug) {
	extern int	nodes, streams;
	static char server_name[80] = "";
	int	i;
	fprintf(stderr, "doOneStep(double load = %g, double think = %g, "
	    "double serviceTime = %g, double dmax = %g, "
	    "int verbose = %d, int debug= %d)\n",
	    load, think, serviceTime, dmax, verbose, debug);

	/* optionally, one can set PDQ_SetDebug(1); */
	PDQ_Init("closed uniserver"); /* Name model. */

	/* Define workload and queuing circuit type. */
	streams = PDQ_CreateClosed("work", TERM, load, think);

	if (dmax == 0.0) {
		/* Create a single node, with a demand of serviceTime */
		(void) sprintf(server_name, "server0"); 
		nodes = PDQ_CreateNode(server_name, CEN, FCFS);
		PDQ_SetDemand(server_name, "work", serviceTime);
	}
	else {
	    /* wants service > 0 && center = true, FIXME */
		/* Construct a dmax node and then a list of nodes, 
 		 * of size << dmax, all totalling serviceTime.
		 */
		(void) sprintf(server_name, "server0");
		nodes = PDQ_CreateNode(server_name, CEN, FCFS);
		PDQ_SetDemand(server_name, "work", dmax);
		serviceTime -= dmax;
		fprintf(stderr, "%s dmax = %g, remaining = %g, i = %d\n",
		    server_name, dmax, serviceTime, 0);

		for (i=1; serviceTime > 0.0; i++) {
			(void) sprintf(server_name, "server%d", i);
			nodes = PDQ_CreateNode(server_name, CEN, FCFS);
			if (serviceTime > (dmax/2)) {
				/* Do half of a dmax. */
				PDQ_SetDemand(server_name, "work", dmax/2);
				serviceTime -= (dmax/2);

				fprintf(stderr, "%s demand used = %g, remaining = %g, i = %d\n",
                	server_name, dmax/2, serviceTime, i);
			}
			else {
				/* Last one, do the remainder. */
				PDQ_SetDemand(server_name, "work", serviceTime);
				fprintf(stderr, "%s demand used = %g, remaining = %g, i = %d\n",
                    server_name, serviceTime, max(0.0, serviceTime - dmax), i);
				break;
			}
		}
    }
	PDQ_Solve(EXACT);

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

