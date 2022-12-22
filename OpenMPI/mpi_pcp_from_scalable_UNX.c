/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */

/* This program is about -  using OpenMPI Parallel write to a single file */



#include "mpi.h"
#include <stdio.h>
#include<sys/types.h>
#include<sys/stat.h>
#include<fcntl.h>

#define BUFSIZE 256*1024
#define CMDSIZE 100
#define MAXPATHLEN 20

/* A parallel copy program*/

int main(int argc, char *argv[])
{

int num_hosts, mystatus, allstatus, done, numread;
int infd, outfd;
char buf[BUFSIZE];
char outfilename[MAXPATHLEN], controlmsg[CMDSIZE];
char soft_limit[20];
MPI_Info hostinfo;
MPI_Comm pcpslaves, all_processes;

MPI_Init(&argc, &argv);
makehostlist(argv[1], "targets", &num_hosts);
MPI_Info_create(&hostinfo);
sprintf(soft_limit, "0:%d", num_hosts);
MPI_Info_set(hostinfo, "soft", soft_limit);

MPI_Comm_spawn("pcp_slave", MPI_ARGV_NULL, num_hosts, hostinfo,0, MPI_COMM_SELF, &pcpslaves, MPI_ERRCODES_IGNORE);

MPI_Info_free(&hostinfo);
MPI_Intercomm_merge(pcpslaves,  0, &all_processes);

strcpy(outfilename, argv[3]);

if((infd == open(argv[2], O_RDONLY))==1 )
{
fprintf(stderr, "input %s doesnot exist ", argv[2]);
sprintf(controlmsg, "exit");
MPI_Bcast( controlmsg, CMDSIZE , MPI_CHAR , 0 , all_processes);
MPI_Finalize();
return -1;

}

else 
{
sprintf(controlmsg, "ready");
MPI_Bcast( controlmsg, CMDSIZE , MPI_CHAR , 0 , all_processes);
}

MPI_Bcast(outfilename, MAXPATHLEN, MPI_CHAR, 0 , all_processes);


if(outfd = open(outfilename , O_CREAT|O_TRUNC|O_WRONLY|S_IRWXU)==-1)
mystatus =-1;
else 
mystatus = 0;

MPI_Allreduce(&mystatus, &allstatus, 1, MPI_INT, MPI_MIN, all_processes);

if(allstatus ==-1) 
{

MPI_Finalize();
fprintf(stderr, "error", outfilename);
return -1;
}

}
