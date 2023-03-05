import asyncio
import configparser
import logging
import os
import random
import socket
import threading
import time
from concurrent import futures
from copy import copy

import grpc

import gocosi_pb2_grpc
from gocosi_pb2 import Gossip
from gocosi_rpc_server import GoCosiRPCServicer
import gocosi_rpc_client as rpc_client


def setup_logger(name, log_file, _formatter, _level=logging.INFO):
    """Function setup as many loggers as you want"""

    handler = logging.FileHandler(log_file)
    handler.setFormatter(_formatter)

    _logger = logging.getLogger(name)
    _logger.setLevel(_level)
    _logger.addHandler(handler)
    return _logger


def read_config():
    config = configparser.RawConfigParser()
    config.read('configuration.properties')
    grpc_server.max_subs = int(config.get('grpc', 'max_subs'))
    grpc_server.max_unsubs = int(config.get('grpc', 'max_unsubs'))
    grpc_server.max_views = int(config.get('grpc', 'max_views'))
    grpc_server.max_events = int(config.get('grpc', 'max_events'))
    grpc_server.max_event_ids = int(config.get('grpc', 'max_event_ids'))
    grpc_server.max_msgs = int(config.get('grpc', 'max_msgs'))
    grpc_server.thread_max_worker = int(config.get('grpc', 'thread_max_worker'))

    grpc_server.gossip_round_time = float(config.get('time', 'gossip_round_time'))
    grpc_server.msg_time_out = float(config.get('time', 'msg_time_out'))

    _max_neighbours = int(config.get('client', 'max_neighbours'))
    # print(config.get('client', 'known_servers'))
    _known_servers_list = config.get('client', 'known_servers').split(',')
    index = port % 10000
    _known_servers = []
    if index == 0:
        _known_servers.append(_known_servers_list[0])
    elif index >= 10:
        _known_servers.append(_known_servers_list[index % 10])
    else:
        _known_servers.append(_known_servers_list[index - 1])
    print(f"my port: {port}, my known server: {_known_servers}")
    return _max_neighbours, _known_servers


def serve(_server):
    listen_port = _server.port
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=50))
    # server = grpc.server(futures.ThreadPoolExecutor())
    gocosi_pb2_grpc.add_GocosiRPCServicer_to_server(_server, server)
    server.add_insecure_port('[::]:' + str(listen_port))
    server.start()
    logger.info("GRPC server started, listening on " + str(listen_port))
    _server.gen_register_node(known_servers)
    # server.stop(10)
    server.wait_for_termination()


def gossip_sender(_server: GoCosiRPCServicer, _max_neighbours):
    logger.info("gossip_sender starts working")
    while len(_server.view_list) == 0:
        time.sleep(0.1)  # 100ms
    while True:
        time.sleep(_server.gossip_round_time)
        if len(_server.event_dict) > 0:
            subs = copy(_server.sub_list)
            unsubs = copy(_server.unsub_list)
            views = copy(_server.view_list)
            events = copy(_server.event_dict)
            if _server.address not in subs:
                subs.append(_server.address)
            request = Gossip(
                subs=subs,
                unsubs=unsubs,
                event_list=events
            )

            neighbours = []
            if len(views) <= _max_neighbours:
                neighbours = views
            else:
                while len(neighbours) < _max_neighbours:
                    index = random.choice(range(len(views)))
                    neighbours.append(views.pop(index))

            # logger.debug("gossip_sender's multicast info: " + str(request))
            # logger.debug("gossip_sender's multicast neighbours: " + str(neighbours))
            threading.Thread(target=gossip_sender_thread, args=(request, neighbours, _server)).start()
            _server.event_dict.clear()


def gossip_sender_thread(request: Gossip, neighbours, _server):
    for node in neighbours:
        asyncio.run(rpc_client.gossip_async(_server, node, request))


def get_host_ip():
    s = None
    _host = ''
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.connect(("8.8.8.8", 80))
        _host = s.getsockname()[0]
    finally:
        s.close()
    return _host


if __name__ == '__main__':
    from argparse import ArgumentParser

    parser = ArgumentParser()
    parser.add_argument('-p', '--port', default=20000, type=int, help='port to listen on')
    parser.add_argument('-s', '--sum', default=10, type=int, help='port to listen on')
    args = parser.parse_args()
    port = args.port
    sum_nodes = args.sum
    host = get_host_ip()

    # Instantiate loggers
    LOG_FILENAME = "logFiles/" + str(port) + "_1" + time.strftime("-%m%d-%H%M") + ".log"
    GRPC_LOG_FILENAME = "logFiles/" + str(port) + time.strftime("-%m%d-%H%M") + ".log"

    if not os.path.isdir("logFiles"):
        os.makedirs("logFiles")

    formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')

    LOG_FILENAME = os.path.join(os.path.dirname(os.path.realpath(__file__)), LOG_FILENAME)
    GRPC_LOG_FILENAME = os.path.join(os.path.dirname(os.path.realpath(__file__)), GRPC_LOG_FILENAME)

    # first file logger
    logger = setup_logger('gocosi_logger', LOG_FILENAME, formatter, _level=logging.DEBUG)
    rpc_client.logger = logger
    # second file logger
    grpc_logger = setup_logger('grpc_logger', GRPC_LOG_FILENAME, formatter, _level=logging.DEBUG)
    logger.debug("Author   : Cuican Li")
    logger.debug("Application : GoCosi")
    logger.debug("")

    grpc_logger.debug("Author   : Cuican Li")
    grpc_logger.debug("Application : GoCosi GRPC Server")
    grpc_logger.debug("")

    # config grpc
    grpc_server = GoCosiRPCServicer(host, port, sum_nodes, grpc_logger)

    max_neighbours, known_servers = read_config()
    thread = threading.Thread(target=gossip_sender, args=(grpc_server, max_neighbours))
    thread.setDaemon(True)
    thread.start()

    # keep gossiping msg thread
    thread1 = threading.Thread(target=grpc_server.gossip_sig_thread)
    thread1.setDaemon(True)
    thread1.start()

    # keep handle new sig event thread
    thread2 = threading.Thread(target=grpc_server.handle_new_sig_thread)
    thread2.setDaemon(True)
    thread2.start()

    # grpc server start
    serve(grpc_server)
