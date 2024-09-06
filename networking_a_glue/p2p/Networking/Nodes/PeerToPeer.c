/* Copyright (C) 2022 Radhamadhab Dalai - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the license, which unfortunately won't be
 * written for another century.
 *
 * You should have received a copy of the license with
 * this file. If not, please write to: radhamadhabdalai@soa.ac.in.
 */


#include "PeerToPeer.h"
#include "Client.h"
#include "Server.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <pthread.h>


void user_portal(struct PeerToPeer *peer_to_peer);
char * known_hosts(void *arg);


struct PeerToPeer peer_to_peer_constructor(int domain, int service, int protocol, int port, u_long interface, void * (*server_function)(void *arg), void * (*client_function)(void *arg))
{
    struct PeerToPeer peer_to_peer;
    peer_to_peer.domain = domain;
    peer_to_peer.service = service;
    peer_to_peer.protocol = protocol;
    peer_to_peer.port = port;
    peer_to_peer.interface = interface;
    
    peer_to_peer.known_hosts = linked_list_constructor();
    peer_to_peer.known_hosts.insert(&peer_to_peer.known_hosts, 0, "127.0.0.1", 10);
    
    peer_to_peer.server = server_constructor(domain, service, protocol, interface, port, 20);
    peer_to_peer.server.register_routes(&peer_to_peer.server, known_hosts, "/known_hosts\n");

    
    peer_to_peer.user_portal = user_portal;
    peer_to_peer.server_function = server_function;
    peer_to_peer.client_function = client_function;
    
    return peer_to_peer;
}


void user_portal(struct PeerToPeer *peer_to_peer)
{
    // Launch the server as a thread.
    pthread_t server_thread;
    pthread_create(&server_thread, NULL, peer_to_peer->server_function, peer_to_peer);
    // Launch the client.
    peer_to_peer->client_function(peer_to_peer);
}

char * known_hosts(void *arg)
{
    // Returns a list of Nodes that have previously been communicated with.
    printf("Known hosts requested\n");
    struct PeerToPeer *peer_to_peer = arg;
    char *hosts = malloc(peer_to_peer->known_hosts.length * 17);
    memset(hosts, 0, peer_to_peer->known_hosts.length * 17);
    for (int i = 0; i < peer_to_peer->known_hosts.length; i++)
    {
        strcat(hosts, peer_to_peer->known_hosts.retrieve(&peer_to_peer->known_hosts, i));
        strcat(hosts, ",");
    }
    return hosts;
}
