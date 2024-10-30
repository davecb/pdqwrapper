#!/usr/bin/perl
###############################################################################
#  Copyright (C) 1994 - 2009, Performance Dynamics Company                    #
#                                                                             #
#  This software is licensed as described in the file COPYING, which          #
#  you should have received as part of this distribution. The terms           #
#  are also available at http://www.perfdynamics.com/Tools/copyright.html.    #
#                                                                             #
#  You may opt to use, copy, modify, merge, publish, distribute and/or sell   #
#  copies of the Software, and permit persons to whom the Software is         #
#  furnished to do so, under the terms of the COPYING file.                   #
#                                                                             #
#  This software is distributed on an "AS IS" basis, WITHOUT WARRANTY OF ANY  #
#  KIND, either express or implied.                                           #
###############################################################################

#
#  $Id: closed1.pl,v 4.3 2009/03/26 02:55:32 pfeller Exp $
#
#------------------------------------------------------------------------------

use pdq;

#------------------------------------------------------------------------------
#
# Based on time_share.c
#
# Illustrates PDQ solver for closed uni-server queue.  Compare with repair.c
#
#------------------------------------------------------------------------------

#---- Model specific variables ------------------------------------------------

#$pop      = 200.0;
$pop      = 100.0;
$think    = 300.0;
$servt    = 0.63;

#---- Initialize the model giving it a name -----------------------------------

pdq::Init("Time Share Computer");
pdq::SetComment("This is just a simple M/M/1 queue.");

#---- Define the workload and circuit type ------------------------------------

$streams = pdq::CreateClosed("compile", $pdq::TERM, $pop, $think);

#  So the value of $streams should correspond to the value of $pdq::streams

#---- Define the queueing center ----------------------------------------------

$nodes  = pdq::CreateNode("CPU", $pdq::CEN, $pdq::FCFS);

#  So the value of $nodes should correspond to the value of $pdq::nodes

#---- Define service demand ---------------------------------------------------

pdq::SetDemand("CPU", "compile", $servt);

#---- Solve the model ---------------------------------------------------------

pdq::Solve($pdq::EXACT);

pdq::Report();

