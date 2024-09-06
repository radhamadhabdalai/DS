* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */

/*
 The Client is used as a tool for connecting to and interacting with servers.
 */


#ifndef Client_h
#define Client_h

#include <sys/socket.h>
#include <netinet/in.h>



struct Client
{
   
    // The network socket for handling connections.
    int socket;
    // Variables dealing with the specifics of a connection.
    int domain;
    int service;
    int protocol;
    int port;
    u_long interface;
  
    // The request method allows a client to make a request of a specified server.
    char * (*request)(struct Client *client, char *server_ip, void *request, unsigned long size);
};




struct Client client_constructor(int domain, int service, int protocol, int port, u_long interface);


#endif
