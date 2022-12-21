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
#include <mpi.h>

#define BUFF_SIZE 100 

int main (int argc, char *argv[])
{
   FILE *fp;
   int myrank;
   int buff[BUFF_SIZE];
   char filename[50];
   int i;


   MPI_Init(NULL, NULL);
   MPI_Comm_rank(MPI_COMM_WORLD, &myrank);

   for(i = 0 ; i< BUFF_SIZE ; i++)
     buff[i] = myrank * i + BUFF_SIZE;

    sprintf(filename, "textfile%d.txt", myrank);  
   fp = fopen( filename , "w");
   fwrite(buff, sizeof(int), BUFF_SIZE,fp);
    
   fclose(fp);
  
   MPI_Finalize();  

  return 0;  
}
