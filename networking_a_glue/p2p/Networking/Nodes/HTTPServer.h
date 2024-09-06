/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */



/*
 The HTTPServer prog. is the basis for servers to read and recieve HTTP protocols.
 */

#ifndef HTTPServer_h
#define HTTPServer_h

#include "Server.h"
#include "../Protocols/HTTPRequest.h"



struct HTTPServer
{
     // A generic server object to connect to the network with the appropriate protocols.
    struct Server server;
    // A dictionary of routes registerred on the server with URL's as keys.
    struct Dictionary routes;
    
     // This method is used to register URL's as routes to the server.
    void (*register_routes)(struct HTTPServer *server, char * (*route_function)(struct HTTPServer *server, struct HTTPRequest *request), char *uri, int num_methods, ...);
    // The launch sequence begins an infinite loop where the server listens for and handles incoming connections.
    void (*launch)(struct HTTPServer *server);
};

// The HTTPMethods enum lists the various HTTP methods for easy referral.
enum HTTPMethods
{
    CONNECT,
    DELETE,
    GET,
    HEAD,
    OPTIONS,
    PATCH,
    POST,
    PUT,
    TRACE
};


struct HTTPServer http_server_constructor(void);




// This function combines the contents of files into a single string.
// This is used to create web pages out of multiple templates.
char *render_template(int num_templates, ...);

#endif 

