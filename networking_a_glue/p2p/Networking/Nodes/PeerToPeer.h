/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */


/*
 The PeerToPeer program is to vreate an object for a node with client and server capacities.

 */

#ifndef PeerToPeer_h
#define PeerToPeer_h

#include "Server.h"
#include "../../DataStructures/Lists/LinkedList.h"

struct PeerToPeer
{
    struct Server server;
    struct LinkedList known_hosts;
    
    int domain;
    int service;
    int protocol;
    int port;
    u_long interface;
    
    // The user_portal function is activated to launch the server and client applications.
    void (*user_portal)(struct PeerToPeer *peer_to_peer);
    void * (*server_function)(void *arg);
    void * (*client_function)(void *arg);
};

struct PeerToPeer peer_to_peer_constructor(int domain, int service, int protocol, int port, u_long interface, void * (*server_function)(void *arg), void * (*client_function)(void *arg));

#endif 
