
/*
 * pdq -- report on a single queuing centre, based on
 *	closed_center.c from Gunther. This is what I use
 *	from the command-line 
 *	 
 */
#include <stdio.h>
#include <stdlib.h>	/* For exit(). */
#include <libgen.h>	/* For basename(). */
#include <math.h>
/* #include "/usr/local/pdq/lib/PDQ_Lib.h" */
#include "/home/davecb/projects/PDQ2/pdq/pdq5/lib/PDQ_Lib.h"

#define PREFORK
#define STRESS	0
#define HOMEPG	1

static char *ProgName = NULL;
void doOneStep(double load, double think, double serviceDemand, double dmax, 
	int verbose);

 int
usage() {
	fprintf(stderr, "Usage: %s [-z think][-s service][-d dmax][-vx] "
		"-c centers from to by\n", ProgName);
}

 int
main(int argc, char **argv) {
	double	from = 1.0,
		to = 0.0,
		by = 0.0,
		think = 0.0,   
		serviceDemand = 0.0, 
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
			case 'z': think = atof(argv[++i]);
				break;
			case 'x': debug = 1;
				break;
			case 'd': dmax = atof(argv[++i]);
				break;
			case 's': serviceDemand = atof(argv[++i]);
				break;
			case 'c': centers = atof(argv[++i]);
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
	if (i < argc) {
		from = atof(argv[i++]);
	}
	if (i < argc) {
		to = atof(argv[i++]);
	}
	if (i < argc) {
		by = atof(argv[i]);
	}

	/* Check target. */
	if (to == 0.0) {
		to = from;
	}

	/* Print headers. */
	printf("\"# General closed solution from PDQ where "
		"serviceDemand = %g, centers = %g,\"\n"
	       "\"# think time = %g dmax = %g\"\n",
	       serviceDemand, centers, think, dmax);
	if (verbose) {
		printf("Load\tThroughput\tUtilization\tQueueLen\t"
			"Residence\tResponse\n");
	}
	else {
		printf("\"# Load,\" Response\n");
	}

	/* Adjust Dmax if we have more than one center. */
	if (dmax == 0.0 && centers != 1) {
		printf("Dmax must be non-zero for this model.\n");
		exit(3);
	}
	else {
		dmax = dmax / centers;
	}
	for (load=from; load < (to + 1.0); load += by) {
		doOneStep(load, think, serviceDemand, dmax, verbose);
	}
	if (debug == 1) {
		PDQ_Report();
	}
	exit(0);
}


/*
 * doOneStep -- do one solution step
 */
 void
doOneStep(double load, double think, double serviceDemand, double dmax, 
	int verbose) {
	extern int	nodes, streams;
	static char server_name[80] = "";
	int	i;

	extern double	PDQ_GetResponse();
	extern double	PDQ_GetThruput();
	extern double	PDQ_GetUtilization();

	/* PDQ_SetDebug(1); */
	PDQ_Init("closed uniserver"); /* Name model. */

	/* Define workload and queuing circuit type. */
	streams = PDQ_CreateClosed("work", TERM, load, think);

	if (dmax == 0.0) {
		/* Create a single node, of serviceDemand */
		(void) sprintf(server_name, "server0"); 
		nodes = PDQ_CreateNode(server_name, CEN, FCFS);
		PDQ_SetDemand(server_name, "work", serviceDemand);
	}
	else {
		/* Construct a dmax node and then a list of nodes, 
 		 * of size << dmax, all totalling serviceDemand.
		 */
		(void) sprintf(server_name, "server0");
		nodes = PDQ_CreateNode(server_name, CEN, FCFS);
		PDQ_SetDemand(server_name, "work", dmax);
		serviceDemand -= dmax;

		for (i=1; serviceDemand > 0.0; i++) {
			(void) sprintf(server_name, "server%d", i);
			nodes = PDQ_CreateNode(server_name, CEN, FCFS);
			if (serviceDemand > (dmax/2)) {
				/* Do half a dmax. */
				PDQ_SetDemand(server_name, "work", dmax/2);
				serviceDemand -= (dmax/2);
			}
			else {
				/* Last one, do the remainder. */
				PDQ_SetDemand(server_name, "work", 
					serviceDemand);
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
		printf("%d,\t%f\n",(int) load, PDQ_GetResponse(TERM, "work"));
	}

} /* doOneStep */
