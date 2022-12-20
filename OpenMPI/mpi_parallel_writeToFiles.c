#include <stdio.h>
#include <mpi.h>
#define BUFF_SIZE 100


int main(int argc, char* argv[])
{
   /* initialize variables*/
   MPI_File mpifp;
   int myrank;
   int i = 0;
   int buff[BUFF_SIZE];
   char filename[50];
   

   /* initialize MPI */
   MPI_Init(NULL, NULL);
   MPI_Comm_rank( MPI_COMM_WORLD , &myrank); 

    for(i = 0 ; i < BUFF_SIZE ; i++)
     
     buff[i] = i * BUFF_SIZE + myrank;

    sprintf(filename, "filename%d.txt", myrank);

    MPI_File_open( MPI_COMM_SELF , filename , MPI_MODE_WRONLY | MPI_MODE_CREATE, MPI_INFO_NULL , &mpifp);

    MPI_File_write(mpifp, buff, BUFF_SIZE, MPI_INT, MPI_STATUS_IGNORE);

    MPI_File_close(&mpifp);
    MPI_Finalize();

return 0;
}

