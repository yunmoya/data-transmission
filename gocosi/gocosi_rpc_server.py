import configparser
import time
import traceback
from _sha256 import sha256
from concurrent.futures.thread import ThreadPoolExecutor
from copy import copy
from json import dumps, loads
import random
from urllib.parse import urlparse

import grpc
from blspy import SignatureMPL, BasicSchemeMPL, PublicKeyMPL
from grpc import StatusCode

import gocosi_pb2
import gocosi_pb2_grpc
from bls import BLS

import gocosi_rpc_client as client


def get_msg_hash(msg):
    if type(msg) == str:
        return sha256(msg.encode()).hexdigest()
    return sha256(dumps(msg).encode()).hexdigest()


class GoCosiRPCServicer(gocosi_pb2_grpc.GocosiRPCServicer):
    def __init__(self, _host, _port, _sum_nodes, _logger):
        self.host = _host
        self.port = _port
        self.sum_nodes = _sum_nodes
        self.logger = _logger
        self.address = _host + ":" + str(_port)
        self.index = 0

        self.node_list = []
        self.sub_list = []
        self.unsub_list = []
        self.view_list = []
        self.event_id_list = []
        self.sig_event_id_list = []
        self.msg_latency = []
        self.msg_list = []
        self.event_dict = dict()
        self.sig_event_dict = dict()
        self.msg_sig_dict = dict()
        self.public_key_dict = dict()

        self.max_subs = 100
        self.max_unsubs = 100
        self.max_views = 50
        self.max_events = 50
        self.max_event_ids = 50
        self.max_msgs = 50
        self.thread_max_worker = 10
        self.gossip_round_time = 0.01  # seconds
        self.msg_time_out = 600  # s

        self.thread_pool = ThreadPoolExecutor(self.thread_max_worker)

        self.private_key = None
        self.public_key = None

    def RegisterNode(self, request: gocosi_pb2.RegisterNodeReq, context):
        try:
            message = self.gen_register_node(request.nodes)
            return gocosi_pb2.CommonResp(message=message)
        except Exception as e:
            traceback.print_exc()
            return gocosi_pb2.CommonResp(message=str(e))

    def NewMsg(self, request: gocosi_pb2.NewMsgReq, context):
        try:
            if len(self.sub_list) <= 0:
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details("I'm not ready.")
            else:
                signers, co_sig = self.new_msg(request.msg)
                return gocosi_pb2.NewMsgResp(signature=co_sig, signers=signers)
        except Exception as e:
            traceback.print_exc()
            return gocosi_pb2.CommonResp(message=str(e))

    def GossipReq(self, request: gocosi_pb2.Gossip, context):
        try:
            # print(f"receive {request}")
            unsubs = list(request.unsubs)
            subs = list(request.subs)
            # phase 1
            for unsub in unsubs:
                if unsub in self.view_list:
                    self.view_list.remove(unsub)
                if unsub not in self.unsub_list:
                    self.unsub_list.append(unsub)
            while len(self.unsub_list) > self.max_unsubs:
                self.unsub_list.remove(random.choice(self.unsub_list))
            # phase 2
            for sub in subs:
                if sub == self.address:
                    pass
                elif sub not in self.view_list:
                    self.view_list.append(sub)
                    if sub not in self.sub_list:
                        self.sub_list.append(sub)
            while len(self.view_list) > self.max_views:
                target = random.choice(self.view_list)
                self.view_list.remove(target)
                if target not in self.sub_list:
                    self.sub_list.append(target)
            while len(self.sub_list) > self.max_subs:
                self.sub_list.remove(random.choice(self.sub_list))
            # phase 3
            # print(str(request.event_list.keys()))
            for e_id in list(request.event_list.keys()):
                if e_id not in self.event_id_list:
                    self.event_dict[e_id] = request.event_list[e_id]
                    event = loads(request.event_list[e_id])
                    event_type = event['type']
                    if event_type == 0:  # register node
                        msg = self.process_register(event['values'])
                    elif event_type == 1:  # signature
                        if e_id not in self.sig_event_dict:
                            self.sig_event_id_list.append(e_id)
                            self.sig_event_dict[e_id] = event
                            while len(self.event_id_list) > self.max_events:
                                self.event_id_list.pop(0)
                                self.event_dict.pop(random.choice(list(self.event_dict)))
                    else:
                        context.set_code(StatusCode.INVALID_ARGUMENT)
                        context.set_details("There has invalid type in events")
                    self.event_id_list.append(e_id)
            while len(self.event_id_list) > self.max_event_ids:
                self.event_id_list.pop(0)
            while len(self.event_dict) > self.max_events:
                self.event_dict.popitem()
            return gocosi_pb2.CommonResp(message='receive successfully')
        except Exception as e:
            traceback.print_exc()
            return gocosi_pb2.CommonResp(message=str(e))

    def GetPubkey(self, request, context):
        return gocosi_pb2.GetPubkeyResp(publickeys=self.public_key_dict)

    def Info(self, request, context):
        print("recieve request")
        config = configparser.RawConfigParser()
        config.read('configuration.properties')
        max_neighbours = int(config.get('client', 'max_neighbours'))
        return gocosi_pb2.InfoResp(
            subs=self.sub_list,
            unsubs=self.unsub_list,
            pubkey=self.public_key_dict,
            latency=self.msg_latency,
            round_time=str(self.gossip_round_time),
            neighbours=max_neighbours
        )

    def gen_register_node(self, nodes):
        if nodes is None:
            return "please supply a valid list of nodes"
        if type(nodes) == str:
            nodes = nodes.split(",")
        for node in nodes:
            parsed_url = urlparse(node)
            if parsed_url.netloc or parsed_url.path:
                if node == self.address:
                    pass
                elif node not in self.sub_list:
                    # print(self.address)
                    # self.logger.info("gen_register_node | " + node + " is not in my sub list")
                    self.sub_list.append(node)
                    if node not in self.view_list:
                        self.view_list.append(node)
            else:
                return "unvalid node"
        values = {
            'address': self.address,
            'pub_key': self.public_key_dict[self.address]
        }
        event = {
            'values': values,
            'type': 0
        }
        event = dumps(event)
        self.event_dict[get_msg_hash(event)] = event
        return "add node successfully"

    def process_register(self, values):
        address = values['address']
        pub_key = values['pub_key']
        self.public_key_dict[address] = pub_key
        if address == self.address:
            pass
        elif address not in self.sub_list:
            # self.logger.info("process_register | " + address + " is not in my sub list")
            self.sub_list.append(address)
            if address not in self.view_list:
                self.view_list.append(address)
        return "add node successfully"

    def sig_for_msg(self, values):
        message = values.get('msg')
        sigs = values.get('sigs')
        co_sig = sigs.get('co_sig')
        signers = sigs.get('signers')
        self.logger.debug(f"sig_for_msg | recieve signers {signers}")
        if message is None or type(message) is not str or signers is None:
            return "tempered Data"
        m_h = get_msg_hash(message)
        if m_h in self.msg_sig_dict:
            signers_length = sum(1 for i in self.msg_sig_dict[m_h].get('signers') if i > 0)
            if signers_length / self.sum_nodes > 0.66:
                self.logger.debug(f"I have enough signatures, length:{signers_length}")
                return "I have enough signatures"
        co_sig = BLS.deserialize(co_sig, SignatureMPL)
        self.logger.debug(f"sig_for_msg | {message} is valid, signers:{signers}")
        if message not in self.msg_list:
            if self.is_valid_msg(message, co_sig, signers):
                # TODO check msg from blockchain
                temp_array = []
                for c in message:
                    temp_array.append(ord(c))
                msg = bytes(temp_array)
                sig = BasicSchemeMPL.sign(self.private_key, msg)
                co_sig = BasicSchemeMPL.aggregate([co_sig, sig])
                co_sig_str = bytes(co_sig).hex()
                signers[self.index] = 1
                sigs = {
                    'co_sig': co_sig_str,
                    'signers': signers
                }
                self.add_msg(message, sigs)
                return 'add signature successfully'
            else:
                self.logger.error(f'sig_for_msg | invalid msg:{message}')
                return 'invalid msg'

        else:
            my_sigs = self.msg_sig_dict[m_h]
            self.logger.debug(f"sig_for_msg | agg signers {sigs.get('signers')} to {my_sigs.get('signers')}")
            my_co_sig = BLS.deserialize(my_sigs.get('co_sig'), SignatureMPL)
            my_signers = copy(my_sigs.get('signers'))

            update = False
            new_signers = [0] * self.sum_nodes
            for i in range(self.sum_nodes):
                if signers[i] > 0 and my_signers[i] == 0:
                    update = True
                    new_signers[i] = signers[i]
            # unkown signers exist
            if update:
                for i in range(self.sum_nodes):
                    if new_signers[i] > 0:
                        my_signers[i] = new_signers[i]
                    else:
                        my_signers[i] += signers[i]

                new_co_sig = BasicSchemeMPL.aggregate([my_co_sig, co_sig])
                new_co_sig_str = bytes(new_co_sig).hex()
                sigs = {
                    'co_sig': new_co_sig_str,
                    'signers': my_signers
                }
                self.msg_sig_dict[m_h] = sigs
                signers_length = sum(1 for i in my_signers if i > 0)
                self.logger.debug(
                    f"sig_for_msg | agg successfully, new sigs signers length:{signers_length}")
            else:
                self.logger.debug("sig_for_msg | no new signers")
            return 'add signature successfully'

    def new_msg(self, message):
        if message is None or type(message) is not str:
            self.logger.error("new_msg | tempered Data")
            return dict(), None
        m_h = get_msg_hash(message)
        if m_h in self.msg_sig_dict:
            signers_length = sum(1 for i in self.msg_sig_dict[m_h].get('signers') if i > 0)
            if signers_length / self.sum_nodes > 0.66:
                signers = self.msg_sig_dict[m_h].get('signers')
                co_sig = self.msg_sig_dict[m_h].get('co_sig')
                return signers, co_sig
        # TODO check msg from blockchain
        temp_array = []
        for c in message:
            temp_array.append(ord(c))
        msg = bytes(temp_array)
        sig = BasicSchemeMPL.sign(self.private_key, msg)
        sig_str = bytes(sig).hex()
        signers = [0] * self.sum_nodes
        signers[self.index] = 1
        sigs = {
            'co_sig': sig_str,
            'signers': signers
        }
        self.add_msg(message, sigs)
        future = self.thread_pool.submit(self.get_msg_cosig, message)
        signers, co_sig = future.result()
        # print(signers)
        # print(co_sig)
        return signers, co_sig

    def is_valid_msg(self, msg, co_sig, signers):
        if type(co_sig) is not SignatureMPL:
            co_sig = BLS.deserialize(co_sig, SignatureMPL)
        temp_array = []
        for c in msg:
            temp_array.append(ord(c))
        msg = bytes(temp_array)
        agg_pk = None
        for i in range(len(signers)):
            if signers[i] > 0:
                node = self.node_list[i]
                if node in self.public_key_dict:
                    pk = BLS.deserialize(self.public_key_dict[node], PublicKeyMPL)
                    temp_pk = None
                    for j in range(signers[i]):
                        if temp_pk is None:
                            temp_pk = pk
                        else:
                            temp_pk += pk
                    if agg_pk is None:
                        agg_pk = temp_pk
                    else:
                        agg_pk += temp_pk
                else:
                    public_keys = client.get_public_key(self, node)
                    if public_keys:
                        for new_node in list(public_keys.keys()):
                            if new_node not in self.public_key_dict:
                                self.public_key_dict[new_node] = public_keys[new_node]
                        self.logger.debug(f"my pub key dict keys: {self.public_key_dict.keys()}")
                        pk = BLS.deserialize(self.public_key_dict[node], PublicKeyMPL)
                        temp_pk = None
                        for j in range(signers[i]):
                            if temp_pk is None:
                                temp_pk = pk
                            else:
                                temp_pk += pk
                        if agg_pk is None:
                            agg_pk = temp_pk
                        else:
                            agg_pk += temp_pk
                    else:
                        self.logger.error("is_valid_msg | can not get " + node + "'s public key")
                        return False
        if agg_pk is None:
            self.logger.error("is_valid_msg | agg_pk is None")
            return False
        print(f'agg_pk: {agg_pk}')
        print(f'signature: {co_sig}')
        print(f'verify msg: {msg}')
        return BasicSchemeMPL.verify(agg_pk, msg, co_sig)

    def get_msg_cosig(self, msg):
        start = time.time()
        while time.time() - start < self.msg_time_out:
            m_h = get_msg_hash(msg)
            sigs = self.msg_sig_dict[m_h]
            signers = sigs.get('signers')
            # due to self address is not in sub_list, so we compute len(self.sub_list) + 1 to represent node numbers
            signers_length = sum(1 for i in signers if i > 0)
            if signers_length / self.sum_nodes > 0.66:
                spent_time = time.time() - start
                self.logger.info(f"get_msg_cosig | {msg} spent time: {spent_time}")
                self.msg_latency.append(spent_time)
                co_sig = sigs.get('co_sig')
                return signers, co_sig
        return None, None

    def add_msg(self, message, sigs):
        self.msg_list.append(message)
        self.msg_sig_dict[get_msg_hash(message)] = sigs
        self.logger.debug(
            f"add_msg | add new msg: {message} signers: {self.msg_sig_dict[get_msg_hash(message)].get('signers')}")
        if len(self.msg_list) > self.max_msgs:
            oldest_msg = self.msg_list.pop(0)
            o_msg_hash = get_msg_hash(oldest_msg)
            self.msg_sig_dict.pop(o_msg_hash, None)

    def gossip_sig_thread(self):
        print("gossip_sig_thread start")
        last_gossip = dict()
        while True:
            if len(self.msg_list) == 0:
                time.sleep(0.2)
            else:
                time.sleep(0.1)
                for msg in self.msg_list[::-1]:
                    m_h = get_msg_hash(msg)
                    values = {
                        'msg': msg,
                        'sigs': self.msg_sig_dict[m_h]
                    }
                    val_hash = get_msg_hash(values)
                    if msg in last_gossip and val_hash is last_gossip[m_h]:
                        pass
                    else:
                        last_gossip[m_h] = val_hash
                        event = {
                            'values': values,
                            'type': 1
                        }
                        event = dumps(event)
                        self.event_dict[get_msg_hash(event)] = event
                self.logger.info(f"gossip_sig_thread | {self.event_dict}")

    def handle_new_sig_thread(self):
        while True:
            if self.sig_event_id_list:
                e_id = self.sig_event_id_list.pop()
                event = self.sig_event_dict.pop(e_id)
                msg = self.sig_for_msg(event['values'])
                self.logger.info(f"handle_new_sig_thread | response e_id: {e_id} message: {msg}")
            else:
                time.sleep(0.1)
