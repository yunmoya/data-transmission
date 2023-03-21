import asyncio
import configparser
import logging
import os
import random
import threading
import time
from concurrent import futures
from copy import copy

import grpc
from blspy import PrivateKey, PublicKeyMPL

import gocosi_pb2_grpc
from bls import BLS
from gocosi_pb2 import Gossip
from gocosi_rpc_server import GoCosiRPCServicer
import gocosi_rpc_client as rpc_client
from utils import get_host_ip, write_to_toml, read_from_toml


def setup_logger(name, log_file, _formatter, _level=logging.INFO):
    """Function setup as many loggers as you want"""

    handler = logging.FileHandler(log_file)
    handler.setFormatter(_formatter)

    _logger = logging.getLogger(name)
    _logger.setLevel(_level)
    _logger.addHandler(handler)
    return _logger


def load_nodes_info(file_names, sum_nodes):
    node_list = [0] * sum_nodes
    public_ley_dict = dict()
    info = read_from_toml(file_names)
    nodes = info['public']
    for node in nodes:
        address = node['address']
        index = int(node['index'])
        pub_key = node['public_key']
        public_ley_dict[address] = pub_key
        node_list[index] = address
    print(f"node_list: {node_list} \n public_key_dict: {public_ley_dict}")
    return node_list, public_ley_dict


def load_self_info(filename):
    if not os.path.exists(filename):
        print(f"failed to find {filename}")
        return None, None, None
    else:
        info = read_from_toml(filename)
        server = info['server']
        index = int(server['index'])
        key = server['key']
        pri_key_hex = key['private_key']
        pub_key_hex = key['public_key']
        return index, BLS.deserialize(pri_key_hex, PrivateKey), BLS.deserialize(pub_key_hex, PublicKeyMPL)


def read_config(grpc_server, port, sum_nodes):
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

    key_file_path = config.get('server', 'key_file_path')
    key_file = f"{key_file_path}/local{port}.toml"
    index, pri_key, pub_key = load_self_info(key_file)
    grpc_server.index = index
    grpc_server.private_key = pri_key
    grpc_server.public_key = pub_key

    node_list, public_key_dict = load_nodes_info(key_file, sum_nodes)
    grpc_server.public_key_dict = public_key_dict
    grpc_server.node_list = node_list
    grpc_server.view_list = copy(node_list)
    grpc_server.sub_list = copy(node_list)
    while len(grpc_server.view_list) > grpc_server.max_views:
        target = random.choice(grpc_server.view_list)
        grpc_server.view_list.remove(target)
        if target not in grpc_server.sub_list:
            grpc_server.sub_list.append(target)
    while len(grpc_server.sub_list) > grpc_server.max_subs:
        grpc_server.sub_list.remove(random.choice(grpc_server.sub_list))

    _max_neighbours = int(config.get('client', 'max_neighbours'))
    print(f"my port: {port}, my known server: {node_list}")
    return _max_neighbours, node_list


def init_server(_args):
    port = _args.port
    index = _args.index
    parent_path = "certs"
    if not os.path.exists(parent_path):
        os.mkdir(parent_path)
    file_path = f"{parent_path}/local{port}.toml"
    ip = get_host_ip()
    pri_key, pub_key = BLS.createKey()
    pri_key_hex = bytes(pri_key).hex()
    pub_key_hex = bytes(pub_key).hex()
    key = {
        'private_key': pri_key_hex,
        'public_key': pub_key_hex
    }

    server = {
        'address': f"{ip}:{port}",
        'index': index,
        'key': key
    }

    public = {
        'address': server['address'],
        'index': index,
        'public_key': pub_key_hex
    }

    info = {
        'server': server,
        'public': [public]
    }

    write_to_toml(info, file_path)


def start_server(_args):
    port = _args.port
    sum_nodes = _args.sum
    host = get_host_ip()
    cert_path = f"certs/local{port}.toml"

    # Check validation
    if not os.path.exists(cert_path):
        print(f"cert file is not exists, please initiate a gocosi server first.")
        return

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

    max_neighbours, known_servers = read_config(grpc_server, port, sum_nodes)
    thread = threading.Thread(target=gossip_sender, args=(grpc_server, max_neighbours, logger))
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
    serve(grpc_server, logger)


def serve(_server, logger):
    listen_port = _server.port
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=50))
    gocosi_pb2_grpc.add_GocosiRPCServicer_to_server(_server, server)
    server.add_insecure_port('[::]:' + str(listen_port))
    server.start()
    logger.info("GRPC server started, listening on " + str(listen_port))
    server.wait_for_termination()


def gossip_sender(_server: GoCosiRPCServicer, _max_neighbours, _logger):
    _logger.info("gossip_sender starts working")
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

            _logger.debug("gossip_sender's multicast info: " + str(request))
            _logger.debug("gossip_sender's multicast neighbours: " + str(neighbours))
            threading.Thread(target=gossip_sender_thread, args=(request, neighbours, _server)).start()
            _server.event_dict.clear()


def gossip_sender_thread(request: Gossip, neighbours, _server):
    for node in neighbours:
        asyncio.run(rpc_client.gossip_async(_server, node, request))


if __name__ == '__main__':
    from argparse import ArgumentParser

    parser = ArgumentParser(description="Initiate or start a gocosi server")
    subparsers = parser.add_subparsers(title='sub-command', help='sub-command help')
    parser_init = subparsers.add_parser('init', help="initiate a gocosi server")
    parser_init.add_argument('-p', '--port', default=20000, type=int, help='server\'s port')
    parser_init.add_argument('-i', '--index', default=0, type=int, help='server\'s index in gocosi')
    parser_start = subparsers.add_parser('start', help="start a gocosi server")
    parser_start.add_argument('-p', '--port', default=20000, type=int, help='server\'s port')
    parser_start.add_argument('-s', '--sum', default=10, type=int, help='numbers of gocosi members')
    parser_start.set_defaults(func=start_server)
    parser_init.set_defaults(func=init_server)
    args = parser.parse_args()
    args.func(args)
