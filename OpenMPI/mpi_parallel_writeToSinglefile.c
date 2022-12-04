/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */

/* This program is about -  using OpenMPI Parallel write to a single file */

#include<stdio.h>
#include<mpi.h>
#define BUFF_SIZE 100

int main(int argc, char *argv[])
{
 /* initialize variables*/
  MPI_File mpifile;
  int i, myrank, buff[BUFF_SIZE];
  

  /* start MPI declarations*/
  MPI_Init(argc, &argv);

  MPI_Comm_rank(MPI_COMM_WORLD, &myrank); 
  for(i = 0 ; i < BUFF_SIZE ; i++)

     buff[i] = i* BUFF_SIZE + myrank;


   MPI_File_open(MPI_COMM_WORLD, "testfile", MPI_MODE_CREATE | MPI_MODE_WRONLY , MPI_INFO_NULL, &mpifile);

   MPI_File_set_view(mpifile, myrank * BUFF_SIZE * sizeof(int), MPI_INT, MPI_INT, "native", MPI_INFO_NULL);

   MPI_File_Write(mpifile, buff, BUFF_SIZE, MPI_INT, MPI_STATUS_IGNORE);

   MPI_File_close(&mpifile);



  MPI_Finalize();

return 0;

}

