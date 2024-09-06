
/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */


#ifndef p2p.h
#define p2p.h
#include "Server.h"
#include "../../DataStructures/Lists/LinkedList.h"

struct p3p 
{

struct Server server;
struct LinkedList *known_hosts;

int domain;
int service;
int protocol;
int port;
u_long interface;

}

struct p2p peer_to_peer_constructor(int domain., int service, int protocol, int port, u_long interface);

#endif
