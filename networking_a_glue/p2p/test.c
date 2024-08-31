/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */

/* This program is about -  Concurrent programming using thread */

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <ctype.h>

void *server_function(void *arg)
{
//IFNET6 for ipv6 but for ipv4 AF_IFNET
struct Server server = server_construction(AF_INET, SOCK_STREAM, 0, INADDR_ANY, 1248, 20);
struct sockaddr *address = (struct sockaddr*) &server.address;
socklen_t = address_length = (socklen_t) sizeof(server.address);

while(1) {

    int client = accept(server.socket, address, &address_length);
    char request[255];
    memset(request, 0, 255);
    write(client, request, 255);
    printf("%s\n", request);
    close(client);

}
return NULL;
}


void client_function(char *request){
 struct Client client = client_constructor(AF_INET, SOCK_STREAM, 0,1248, );

}

//server thread program
int main()
{
pthread_t server_thread;
pthread_create(server_thread, NULL, server_function, NULL);


int i = 0;
int err = 0;
void *ret;
 //pthread_mutex_t rw;
 //pthread_mutex_lock(&rw); 
while(i < 3)
{

    if (pthread_create(&tid[i], NULL, printToConsole, (void *)i) != 0) {
    perror("pthread_create() error");
    exit(1);
  }
//  err = pthread_create(&(tid[i]), NULL,printToConsole(i), NULL);
//         if (err != 0) {
//             printf("\nCan't Create thread :[%s]", strerror(err));
//         }
        i++;

}




// pthread_join(tid[0], NULL);
// pthread_join(tid[1], NULL);
// pthread_join(tid[2], NULL);
  if (pthread_join(tid[0], &ret) != 0) {
    perror("pthread_create() error");
    exit(3);
  }
    if (pthread_join(tid[1], &ret) != 0) {
    perror("pthread_create() error");
    exit(3);
  }
    if (pthread_join(tid[2], &ret) != 0) {
    perror("pthread_create() error");
    exit(3);
  }

  return 0;  
}
