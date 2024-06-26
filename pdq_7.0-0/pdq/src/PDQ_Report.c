/*******************************************************************************/
/*  Copyright (C) 1994--2021, Performance Dynamics Company                    */
/*                                                                             */
/*  This software is licensed as described in the file COPYING, which          */
/*  you should have received as part of this distribution. The terms           */
/*  are also available at http://www.perfdynamics.com/Tools/copyright.html.    */
/*                                                                             */
/*  You may opt to use, copy, modify, merge, publish, distribute and/or sell   */
/*  copies of the Software, and permit persons to whom the Software is         */
/*  furnished to do so, under the terms of the COPYING file.                   */
/*                                                                             */
/*  This software is distributed on an "AS IS" basis, WITHOUT WARRANTY OF ANY  */
/*  KIND, either express or implied.                                           */
/*******************************************************************************/

/*
 * PDQ_Report.c
 *
 * Revised by NJG on Fri Aug  2 10:29:48  2002
 * Revised by NJG on Thu Oct  7 20:02:27 PDT 2004
 * Updated by NJG on Mon, Apr 2, 2007
 * Updated by NJG on Wed, Apr 4, 2007: Added Waiting line and time
 * Updated by NJG on Friday, July 10, 2009: fixed dev utilization reporting
 * Updated by PJP on Sat, Nov 3, 2012: Added R support
 * Updated by NJG on Friday, January 11, 2013: 
 *    o Widened top header for new Z.x.y version format
 *    o Centered Report title w/o stars and less glitter
 *    o Modified Model INPUTS to show number of servers numerically for 
 *      both *single* and multi node comparison
 * Updated by NJG on Saturday, January 12, 2013: 
 *    o Fixed wUnit to be tUnit in WORKLOAD Parameters section
 *    o Queue was sometimes wrong for MSQ (too many divides by m)
 * NJG on Monday, February 25, 2013 removed blank line b/w Workload parameters
 * NJG on Sunday, May 15, 2016 moved incomplete circuit warning (line 154 ff.) to Solve()
 * NJG on Tuesday, May 24, 2016. 
      o Cleaned up compiler wornings about unused variables
      o PDQ_VERSION is now a #defined comstant in PDQ_Lib.h
 * NJG on Thursday, December 07, 2017
      o Changed Demand field in WORKLOAD Parameters section to display small 
        service times, e.g., micro-seconds
 * Updated by NJG on Sat Dec 29, 2018  - Format changes for new MSO, MSC multi-server devtypes
 * Updated by NJG on Tue Nov 23, 2020  - Formatting mods for PDQ release 7.0
 *
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include "PDQ_Lib.h"
#include "PDQ_Global.h"


//-------------------------------------------------------------------------

int            syshdr;
int            jobhdr;
int            nodhdr;
int            devhdr;


//----- Prototypes of internal print layout routins -----------------------

void print_node_head(void);
void print_nodes(void);
void print_job(int c, int should_be_class);
void print_sys_head(void);
void print_job_head(int should_be_class);
void print_dev_head(void);
void print_system_stats(int c, int should_be_class);
void print_node_stats(int c, int should_be_class);
void banner_stars(void);
void banner_dash(void);		// Added by NJG on Mon, Apr 2, 2007
void banner_chars(char *s);

//-------------------------------------------------------------------------

void PDQ_Report_null(void)
{
	PRINTF("foo!\n");  //From PDQ_Lib.h #define PRINTF Rprintf
	
}

//-------------------------------------------------------------------------

void PDQ_Report(void)
{
	extern char     model[];
	extern char     s1[], s2[], s3[], s4[];
	extern int      streams, nodes, demands, PDQ_DEBUG;
	extern          JOB_TYPE *job;

	int             c;
	time_t          clock;
	char           *tstamp;
	size_t          fillbase = 26; // was 25
	size_t          fill;
	char           *pad = "                        "; // 24 was 23
	double          allusers = 0.0;
	char           *p = "PDQ_Report()";

	if (PDQ_DEBUG == 1)
	{
		/*debug(p, "Entering");*/
		PRINTF("Entering PDQ_Report()\n");
	}

	resets(s1);
	resets(s2);
	resets(s3);
	resets(s4);

	syshdr = FALSE;
	jobhdr = FALSE;
	nodhdr = FALSE;
	devhdr = FALSE;

	if ((clock = time(0)) == -1) {
		errmsg(p, "Failed to get date");
	}

    tstamp = (char *) ctime(&clock);
    // e.g., "Thu Jan 10 21:19:40 2013" is 24 chars + \n\0 
    // see http://www.thinkage.ca/english/gcos/expl/c/lib/ctime.html
	strncpy(s4, tstamp, 24); // toss embedded \n char
	fill = fillbase - strlen(s4);
	strcpy(s1, s4);
	strncat(s1, pad, fill);
	
	fill = fillbase - strlen(model);
	strcpy(s2, model);
	strncat(s2, pad, fill);

    // VERSION is now a #defined comstant in PDQ_Lib.h
	fill = fillbase - strlen(PDQ_VERSION);
	strcpy(s3, PDQ_VERSION);
	strncat(s3, pad, fill);

	PRINTF("\n");
	// NJG on Friday, January 11, 2013
	// Center the Report title w/o stars
	PRINTF("%15s%9s%24s%9s\n", " ", " ","PRETTY DAMN QUICK REPORT"," ");
	banner_dash();
	PRINTF("               ***  on   %s   ***\n", s1);
	PRINTF("               ***  for  %s   ***\n", s2);
	PRINTF("               ***  PDQ  %s   ***\n", s3);
	banner_dash();

	resets(s1);
	resets(s2);
	resets(s3);
	resets(s4);

	PRINTF("\n");
	
    // The following logic was moved to Solve() by NJG on Sunday, May 15, 2016
    // NJG on Wednesday, August 19, 2015 Added incomplete PDQ circuit detection.
	//    if (!streams) PRINTF("PDQ_Report warning: No PDQ workload defined.\n");
	//    if (!nodes) PRINTF("PDQ_Report warning: No PDQ nodes defined.\n");
	//    if (!demands) PRINTF("PDQ_Report warning: No PDQ service demands defined.\n");

    // Append any user comments
    if (strlen(Comment)) {
		PRINTF("COMMENT: ");
		PRINTF("%s\n\n", Comment);  // In PDQ_Global.c
	}
		

	/* Show INPUT Parameters */
	banner_dash();
	banner_chars("      PDQ Model INPUTS");
	banner_dash();
	PRINTF("\n");
	print_nodes();

	/* OUTPUT Statistics */

	for (c = 0; c < streams; c++) {
		switch (job[c].should_be_class) {
			case TERM:
				allusers += job[c].term->pop;
				break;			
			case BATCH:
				allusers += job[c].batch->pop;
				break;
			case TRANS:
				allusers = 0;
				break;
			default:
				resets(s2);
				sprintf(s2, "Unknown job should_be_class: %d", job[c].should_be_class);
				errmsg(p, s2);
				break;
		}
	}  /* loop over c */

	PRINTF("\n");
	//PRINTF("Queueing Network Parameters\n"); -- edited out by NJG on Tue Nov 17, 2020
	
	switch (job[0].should_be_class) { // can ony be one type of network
			case TERM:
			case BATCH:
				typetostr(s1, CLOSED);
				break;
			case TRANS:
				typetostr(s1, OPEN);
				break;
			case VOID:	
			default:
				typetostr(s1, VOID);
				break;
		}
	
	PRINTF("Network type: %8s\n", s1);
	PRINTF("Workload streams: %4d\n", streams);
	PRINTF("Queueing nodes:   %4d\n\n", nodes);

	//PRINTF("WORKLOAD Parameters:\n");

	for (c = 0; c < streams; c++) {
		switch (job[c].should_be_class) {
			case TERM:
				print_job(c, TERM);
				break;
			case BATCH:
				print_job(c, BATCH);
				break;
			case TRANS:
				print_job(c, TRANS);
				break;
			default:
				typetostr(s1, job[c].should_be_class);
				sprintf(s2, "Unknown job should_be_class: %s", s1);
				errmsg(p, s2);
				break;
		}
	}  /* loop over c */


	for (c = 0; c < streams; c++) {
		switch (job[c].should_be_class) {
			case TERM:
				print_system_stats(c, TERM);
				break;
			case BATCH:
				print_system_stats(c, BATCH);
				break;
			case TRANS:
				print_system_stats(c, TRANS);
				break;
			default:
				typetostr(s1, job[c].should_be_class);
				sprintf(s2, "Unknown job should_be_class: %s", s1);
				errmsg(p, s2);
				break;
		}
	}  /* loop over c */

	PRINTF("\n");

	for (c = 0; c < streams; c++) {
		switch (job[c].should_be_class) {
			case TERM:
				print_node_stats(c, TERM);
				break;
			case BATCH:
				print_node_stats(c, BATCH);
				break;
			case TRANS:
				print_node_stats(c, TRANS);
				break;
			default:
				typetostr(s1, job[c].should_be_class);
				sprintf(s2, "Unknown job should_be_class: %s", s1);
				errmsg(p, s2);
				break;
		}
	}  /* over c */

	PRINTF("\n");

	if (PDQ_DEBUG)
		debug(p, "Exiting");
		
}  /* end of PDQ_Report() */





//=======================================
//   Internal print layout routines 
//=======================================

void print_node_head(void)
{
	extern int      demand_ext, PDQ_DEBUG;
	extern char     model[];
	extern char     s1[];
	extern JOB_TYPE *job;

	char           *dmdfmt = "%-4s %-5s %-10s %-10s %-5s    %12s\n";
	char           *visfmt = "%-4s %-5s %-10s %-10s %-5s %10s %10s %10s\n";

	if (PDQ_DEBUG) {
		typetostr(s1, job[0].network);
		PRINTF("%s Network: \"%s\"\n", s1, model);
		resets(s1);
	}
	
	// was PRINTF("Workload Parameters\n"); - edited by NJG on Tue Nov 17, 2020
	PRINTF("Queueing Network Parameters\n\n");

    //Edited by NJG on Saturday, December 29, 2018
    // for new constant defs in PDQ_Lib.h
    // Column head "Sched" now called Node "Type"
	switch (demand_ext) {
	case DEMAND:
		PRINTF(dmdfmt,
		  "Node", "Sched", "Resource", "Workload", "Class", "Service time");
		PRINTF(dmdfmt,
		  "----", "-----", "--------", "--------", "-----", "------------");
		break;
	case VISITS:
		PRINTF(visfmt,
		  "Node", "Sched", "Resource", "Workload", "Class", "Visits", "Service", "Demand");
		PRINTF(visfmt,
		  "----", "-----", "--------", "--------", "-----", "------", "-------", "------");
		break;
	default:
		errmsg("print_node_head()", "Unknown file type");
		break;
	}

	nodhdr = TRUE;
}  // print_node_head 

//-------------------------------------------------------------------------

void print_nodes(void)
{
	extern char       s1[], s2[], s3[], s4[];
	extern int        demand_ext, PDQ_DEBUG;
	extern int        streams, nodes;
	extern NODE_TYPE *node;
	extern JOB_TYPE  *job;

	int               c, k;
	char             *p = "print_nodes()";

	if (PDQ_DEBUG)
		debug(p, "Entering");

	if (!nodhdr)
		print_node_head();

	for (c = 0; c < streams; c++) {
		for (k = 0; k < nodes; k++) {
			resets(s1);
			resets(s2);
			resets(s3);
			resets(s4);

            //Edited by NJG on Saturday, December 29, 2018
            // for new constant defs in PDQ_Lib.h
			typetostr(s1, node[k].devtype);
			typetostr(s3, node[k].sched);
			
			/* 
			* The following MSQ hackery was disabled on Saturday, December 29, 2018
			*
            * NJG: Friday, January 11, 2013
			* if (node[k].devtype == MSQ) {
			* Function CreateMultiNode(), number of MSO servers is in node.servers
			*	sprintf(s1, "%3d", node[k].devtype); 
			*} else {
			* In CreateNode() function, node.devtype == CEN
			* To be consistent with MSQ reporting that shows number of servers under "Node"
            * column in the WORKLOAD Parameters section of Report(), show single server from
            * CreateNode() as a numeric 1 in "Node" column.
			* node.sched still displays as FCFS
			*	typetostr(s1, node[k].devtype);
		    *  sprintf(s1, "%3d", 1);
			*}
			*/

			getjob_name(s2, c);

			switch (job[c].should_be_class) {
				case TERM:
					typetostr(s4, TERM);
					break;
				case BATCH:
					typetostr(s4, BATCH);
					break;
				case TRANS:
					typetostr(s4, TRANS);
					break;
				default:
					typetostr(s4, VOID);
					break;
			}

			switch (demand_ext) {
				case DEMAND:
					PRINTF("%-4s %-5s %-10s %-10s %-5s %15.6lf\n",
					  s1,
					  s3,
					  node[k].devname,
					  s2,
					  s4,
					  node[k].demand[c]
					);
					break;
				case VISITS:
					PRINTF("%-4s %-4s %-10s %-10s %-5s %10.4f %10.4lf %10.4lf\n",
					  s1,
					  s3,
					  node[k].devname,
					  s2,
					  s4,
					  node[k].visits[c],
					  node[k].service[c],
					  node[k].demand[c]
					);
					break;
				default:
					errmsg("print_nodes()", "Unknown file type");
					break;
			}  /* switch */
		}  /* over k */

		//PRINTF("\n");
	}  /* over c */


	if (PDQ_DEBUG)
		debug(p, "Exiting");

	nodhdr = FALSE;
}  /* print_nodes */

//-------------------------------------------------------------------------

void print_job(int c, int should_be_class)
{
	extern int      PDQ_DEBUG;
	extern JOB_TYPE *job;
	char           *p = "print_job()";

	if (PDQ_DEBUG)
		debug(p, "Entering");

	switch (should_be_class) {
		case TERM:
			print_job_head(TERM);
			PRINTF("%-10s   %6.2f    %10.4lf         %6.2f\n",
		  	job[c].term->name,
		  	job[c].term->pop,
		  	job[c].term->sys->minRT,
		  	job[c].term->think
				);
			break;
		case BATCH:
			print_job_head(BATCH);
			PRINTF("%-10s   %6.2f    %10.4lf\n",
		  	job[c].batch->name,
		  	job[c].batch->pop,
		  	job[c].batch->sys->minRT
				);
			break;
		case TRANS:
			print_job_head(TRANS);
			PRINTF("%-10s     %10.4f    %10.4lf\n",
		  	job[c].trans->name,          // "Arrivals" 
		  	job[c].trans->arrival_rate,  // "Rate"
		  	job[c].trans->sys->minRT     // "Min R" 
				);
			break;
		default:
			errmsg("print_job()", "Unknown job type");
			break;
	}

	if (PDQ_DEBUG)
		debug(p, "Exiting");
}  /* print_job */

//-------------------------------------------------------------------------
//
// The following stats appear in the section labeled
//
//               ******   PDQ Model OUTPUTS   *******

void print_sys_head(void)
{
	extern double   tolerance;
	extern char     s1[];
	extern int      method;
	extern int      iterations;
	extern int      PDQ_DEBUG;
	char           *p = "print_sys_head()";

	if (PDQ_DEBUG)
		debug(p, "Entering");

	PRINTF("\n\n");
	banner_dash();
	banner_chars("     PDQ Model OUTPUTS");
	banner_dash();
	PRINTF("\n");
	typetostr(s1, method);
	PRINTF("Solution method: %s", s1);

	if (method == APPROX)
		PRINTF("        (Iterations: %d; Accuracy: %3.4lf%%)",
		  iterations,
		  tolerance * 100.0
		);

	PRINTF("\n\n");
	banner_chars("   SYSTEM Performance");
	PRINTF("\n");

	PRINTF("Metric                   Value      Unit\n");
	PRINTF("------                  -------     ----\n");

	syshdr = TRUE;

	if (PDQ_DEBUG)
		debug(p, "Exiting");
}  /* print_sys_head */

//-------------------------------------------------------------------------

int             trmhdr = FALSE;
int             bathdr = FALSE;
int             trxhdr = FALSE;

//-------------------------------------------------------------------------

void print_job_head(int should_be_class)
{
	extern char      tUnit[];

	switch (should_be_class) {
		case TERM:
			if (!trmhdr) {
				PRINTF("\n");
				PRINTF("Workload      Users        R minimum      Thinktime\n");
				PRINTF("--------      -----        ---------      ---------\n");
				trmhdr = TRUE;
				bathdr = trxhdr = FALSE;
			}
			break;
		case BATCH:
			if (!bathdr) {
				PRINTF("\n");
				PRINTF("Workload       Jobs        R minimum\n");
				PRINTF("--------       ----        ----------\n");
				bathdr = TRUE;
				trmhdr = trxhdr = FALSE;
			}
			break;
		case TRANS:
			if (!trxhdr) {
				PRINTF("Arrivals          Rate          R minimum\n");
				PRINTF("--------       ----------       ---------\n");
				trxhdr = TRUE;
				trmhdr = bathdr = FALSE;
			}
			break;
		default:
			errmsg("print_job_head()", "Unknown workload type");
			break;
	}
}  /* print_job_head */

//-------------------------------------------------------------------------

void print_dev_head(void)
{
	banner_chars("   RESOURCE Performance");
	PRINTF("\n");
	PRINTF("Metric          Resource     Work               Value    Unit\n");
	PRINTF("------          --------     ----              -------   ----\n");

	devhdr = TRUE;
}  /* print_dev_head */

//-------------------------------------------------------------------------
//
// The following stats appear in the section labeled
//
//               ******   SYSTEM Performance   *******

void print_system_stats(int c, int should_be_class)
{
	extern char      tUnit[];
	extern char      wUnit[];
	extern int       PDQ_DEBUG;
	extern char      s1[], s2[];
	extern JOB_TYPE *job;
	char            *ps = "SYSTEM section: print_system_stats()";
	char            *pw = "Workload section: print_system_stats()";
	char            *pb = "Bounds section: print_system_stats()";

	if (PDQ_DEBUG)
		debug(ps, "Entering");

	if (!syshdr)
		print_sys_head();

    // This is the Workload section of SYSTEM Performance

	switch (should_be_class) {
		case TERM:
			if (job[c].term->sys->thruput == 0) {
				getjob_name(s2, c);
				sprintf(s1, "\nX = %6.4f TERM workname = %s", job[c].term->sys->thruput, s2);
				errmsg(pw, s1);
			}
			PRINTF("Workload: \"%s\"\n", job[c].term->name);
			PRINTF("Mean concurrency      %10.4lf    %s\n",
		  		job[c].term->sys->residency, wUnit);
			PRINTF("Mean throughput       %10.4lf    %s/%s\n",
		  		job[c].term->sys->thruput, wUnit, tUnit);
			PRINTF("Response time         %10.4lf    %s\n",
		  		job[c].term->sys->response, tUnit);
			PRINTF("Round trip time       %10.4lf    %s\n",
		  		job[c].term->sys->response + job[c].term->think, tUnit);
		  	PRINTF("Stretch factor        %10.4lf\n",
		  		job[c].term->sys->response / job[c].term->sys->minRT);
			break;
		case BATCH:
			if (job[c].batch->sys->thruput == 0) {
				getjob_name(s2, c);
				sprintf(s1, "\nX = %6.4f for BATCH workname = %s", job[c].batch->sys->thruput, s2);
				errmsg(pw, s1);
			}
			PRINTF("Workload: \"%s\"\n", job[c].batch->name);
			PRINTF("Mean concurrency      %10.4lf    %s\n",
		  		job[c].batch->sys->residency, wUnit);
			PRINTF("Mean throughput       %10.4lf    %s/%s\n",
		  		job[c].batch->sys->thruput, wUnit, tUnit);
			PRINTF("Response time         %10.4lf    %s\n",
		  		job[c].batch->sys->response, tUnit);
			PRINTF("Stretch factor        %10.4lf\n",
		  		job[c].batch->sys->response / job[c].batch->sys->minRT);
			break;
		case TRANS:
			if (job[c].trans->sys->thruput == 0) {
				getjob_name(s2, c);
				sprintf(s1, "\nX = %6.4f for workname = %s", job[c].trans->sys->thruput, s2);
				errmsg(pw, s1);
			}
			PRINTF("Workload: \"%s\"\n", job[c].trans->name);
			PRINTF("Number in system      %10.4lf    %s\n",
		  		job[c].term->sys->residency, wUnit);
			PRINTF("Mean throughput       %10.4lf    %s/%s\n",
		  		job[c].trans->sys->thruput, wUnit, tUnit);
			PRINTF("Response time         %10.4lf    %s\n",
		  		job[c].trans->sys->response, tUnit);
			PRINTF("Stretch factor        %10.4lf\n",
		  		job[c].term->sys->response / job[c].term->sys->minRT);
		  		break;
		default:
			break;
	}

	PRINTF("\nBounds Analysis:\n");

	switch (should_be_class) {
		case TERM:
			if (job[c].term->sys->thruput == 0) {
				getjob_name(s2, c);
				sprintf(s1, "\nX = %6.4f for workname = %s", job[c].term->sys->thruput, s2);
				errmsg(pb, s1);
			}
			// wUnit and tUnit defined in PDQ_Global.c
			PRINTF("Max throughput        %10.4lf    %s/%s\n",
		  		job[c].term->sys->maxTP, wUnit, tUnit);
			PRINTF("Min response          %10.4lf    %s\n",
		  		job[c].term->sys->minRT, tUnit);
			PRINTF("Max demand            %10.4lf    %s\n",
		  		1 / job[c].term->sys->maxTP,  tUnit);
			PRINTF("Total demand          %10.4lf    %s\n",
		  		job[c].term->sys->minRT, tUnit);
			PRINTF("Think time            %10.4lf    %s\n",
		  		job[c].term->think, tUnit);
			// PRINTF("Optimal load          %10.4lf    %s\n",
// 		  		(job[c].term->think + job[c].term->sys->minRT) * 
// 		  		job[c].term->sys->maxTP, wUnit); 
			PRINTF("Optimal load          %10.4lf    %s\n",
		  		job[c].term->sys->Nopt, wUnit); 
			break;
		case BATCH:
			if (job[c].batch->sys->thruput == 0) {
				getjob_name(s2, c);
				sprintf(s1, "X = %6.4f for workname = %s", job[c].batch->sys->thruput, s2);
				errmsg(pb, s1);
			}
			PRINTF("Max throughput        %10.4lf    %s/%s\n",
		  		job[c].batch->sys->maxTP, wUnit, tUnit);
			PRINTF("Min response          %10.4lf    %s\n",
		  		job[c].batch->sys->minRT, tUnit);
			PRINTF("Max demand            %10.4lf    %s\n",
		  		1 / job[c].batch->sys->maxTP,  tUnit);
			PRINTF("Total demand          %10.4lf    %s\n",
		  		job[c].batch->sys->minRT, tUnit);
			PRINTF("Optimal jobs          %10.4f    %s\n",
				job[c].batch->sys->Nopt, "Jobs");
			break;
		case TRANS:
			PRINTF("Max throughput        %10.4lf    %s/%s\n",
		  		job[c].trans->saturation_rate, wUnit, tUnit);
			PRINTF("Min response          %10.4lf    %s\n",
		  		job[c].trans->sys->minRT, tUnit);
		  	break;
		default:
			break;
	}

	PRINTF("\n");

	if (PDQ_DEBUG)
		debug(ps, "Exiting");
}  /* print_system_stats */

//-------------------------------------------------------------------------
//
// The following stats appear in the section labeled
//
//               ******   RESOURCE Performance   *******


void print_node_stats(int c, int should_be_class)
{

	extern char       s1[];
	extern char       tUnit[];
	extern char       wUnit[];
	extern int        PDQ_DEBUG, demand_ext, nodes;
	extern JOB_TYPE  *job;
	extern NODE_TYPE *node;
	extern char       s3[], s4[];

	double            	X;
	double            	devR;
	double            	devD;
	double          	devU;
	double          	devQ;
	double          	devW;
	double          	devL;
	int               	k;
	int               	mservers;
	char             	*p = "print_node_stats()";

	if (PDQ_DEBUG)
		debug(p, "Entering");

	if (!devhdr)
		print_dev_head();

	getjob_name(s1, c);

	switch (should_be_class) {
		case TERM:
			X = job[c].term->sys->thruput;
			break;
		case BATCH:
			X = job[c].batch->sys->thruput;
			break;
		case TRANS:
			X = job[c].trans->arrival_rate;
			break;
		default:
			break;
	}

	for (k = 0; k < nodes; k++) {
		if (node[k].demand[c] == 0)
			continue;

		if (demand_ext == VISITS) {
			resets(s4);
			strcpy(s4, "Visits");
			strcat(s4, "/");
			strcat(s4, tUnit);
		} else {
			resets(s4);
			strcpy(s4, wUnit);
			strcat(s4, "/");
			strcat(s4, tUnit);
		}

// Updated by NJG on Saturday, December 29, 2018
// Removed hack of using sched type carrying server number (of Friday, January 11, 2013) 
// New metrics: MSO and MSC for server capacity; the 'm' in M/M/m
		resets(s3);
		typetostr(s3, node[k].devtype);
		if (node[k].devtype == MSO || node[k].devtype == MSC) {
			mservers = node[k].servers; 
		} else {
			mservers = 1;
		}
		// Now, display mservers metric	
		PRINTF("%-14s  %-10s   %-10s   %12d   %s\n",
		  "Capacity",
		  node[k].devname,
		  s1,
		  mservers,
		  "Servers"
		);


		PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
		  "Throughput",
		  node[k].devname,
		  s1,
		  (demand_ext == VISITS) ? node[k].visits[c] * X : X,
		  s4
		);


		/******* Calculate other stats *******/
		// Was switch (node[k].sched) { ... like FCFS, LCFS,etc.
		// Now specified via node device type in 7.0 (NJG on Tue Nov 17, 2020)
		switch (node[k].devtype) {
			case MSO: // NJG added this devtype on Dec 29, 2018
			 	// devU is per-server from U < 1 test in MVA_Canon.c
			 	devU = node[k].utiliz[c];
				//mservers = node[k].servers;
				// X is aggregate arrival rate into queueing network
				devQ = X * node[k].resit[c]; // Little's law
				devW = node[k].resit[c] - node[k].demand[c];
				devL = X * devW;
				break;
			case CEN:
			case MSC: // NJG added this devtype on Dec 29, 2018
				// Updated by NJG on Tue Nov 17, 2020
				// These metrics are now computed and assigned to the appropriate 
				// PDQ data structure by MServerFESC() in PDQ_MServer.c 
				devU = node[k].utiliz[c]; // per-server utilization
				devQ = node[k].qsize[c];
				devW = node[k].resit[c] - node[k].demand[c];
				devL = devQ - mservers * devU;
                break;
			case DLY:
			default:
				devU = 100.0;
				devQ = 0.0;
				devW = node[k].demand[c];
				devL = 0.0;
				break;
		}

// NJG: Friday, January 11, 2013 
		PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
		  "In service",
		  node[k].devname,
		  s1,
		  devU * mservers, // total utilization 
		  wUnit
		);
			
		PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
		  "Utilization",
		  node[k].devname,
		  s1,
		  devU * 100, // percent
		  "Percent"
		);
	
		PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
		  "Queue length",
		  node[k].devname,
		  s1,
		  devQ,
		  wUnit
		);
		
		PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
			"Waiting line",
			node[k].devname,
			s1,
			devL,
			wUnit
		);
		
		PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
			"Waiting time",
			node[k].devname,
			s1,
			devW,
			tUnit
		);
		
		PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
			"Residence time",
			node[k].devname,
			s1,
			(node[k].sched == ISRV) ? node[k].demand[c] : node[k].resit[c],
			tUnit
		);

		// Only if visits are used ...
		if (demand_ext == VISITS) {
			/* don't do this if service-time is unspecified */
			devD = node[k].demand[c];
			devR = node[k].resit[c];

			PRINTF("%-14s  %-10s   %-10s   %12.4lf   %s\n",
				"Waiting time",
				node[k].devname,
				s1,
		        (node[k].sched == ISRV) ? devD : devR - devD,
				tUnit
			);
		}
		PRINTF("\n");
	}

	if (PDQ_DEBUG)
		debug(p, "Exiting");
		
}  // print_node_stats




//-------------------------------------------------------------------------

void banner_stars(void)
{
	PRINTF("               ******************************************\n");

}  /* banner_stars */

//-------------------------------------------------------------------------

void banner_dash(void)
{
	PRINTF("               ==========================================\n");

}  // banner_dash




//-------------------------------------------------------------------------
void banner_chars(char *s)
{

	PRINTF("               ********%-26s********\n", s);

}  /* banner_chars */

//-------------------------------------------------------------------------

