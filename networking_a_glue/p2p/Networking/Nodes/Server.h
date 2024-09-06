/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */

/*
 The Server struct is used for nodes that need to operate as servers.

 */

#ifndef Server_h
#define Server_h

#include "../../DataStructures/Dictionary/Dictionary.h"

#include <sys/socket.h>
#include <netinet/in.h>


struct Server
{
    /* PUBLIC MEMBER VARIABLES */
    int domain;
    int service;
    int protocol;
    u_long interface;
    int port;
    int backlog;
    struct sockaddr_in address;
    int socket;
    
    struct Dictionary routes;
    
    void (*register_routes)(struct Server *server, char *(*route_function)(void *arg), char *path);
};

struct ServerRoute
{
    char * (*route_function)(void *arg);
};


struct Server server_constructor(int domain, int service, int protocol, u_long interface, int port, int backlog);

#endif 
