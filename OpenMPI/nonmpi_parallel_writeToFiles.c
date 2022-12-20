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