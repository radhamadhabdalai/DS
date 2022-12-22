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

int main(int argc, char ** argv) 
{

int * a;
MPI_Win win;

MPI_Init(&argc, &argv);

MPI_Win_allocate(1000*sizeof(int), sizeof(int), MPI_INFO_NULL, MPI_COMM_WORLD, &a, &win);

MPI_Win_free(&win);

MPI_Finalize();

return 0;

}
