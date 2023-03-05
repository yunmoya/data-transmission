import time
from json import dumps

import grpc

import gocosi_pb2_grpc
import numpy
from gocosi_pb2 import Empty, NewMsgReq

service_config_json = dumps({
    "methodConfig": [{
        # To apply retry to all methods, put [{}] in the "name" field
        "name": [{}],
        "retryPolicy": {
            "maxAttempts": 5,
            "initialBackoff": "0.1s",
            "maxBackoff": "1s",
            "backoffMultiplier": 2,
            "retryableStatusCodes": ["UNAVAILABLE"],
        },
    }]
})
options = [("grpc.enable_retries", 1), ("grpc.service_config", service_config_json)]
# NOTE: the retry feature will be enabled by default >=v1.40.0


timeout = 600  # seconds
logger = None


async def gossip_async(server, address, request):
    try:
        async with grpc.aio.insecure_channel(address, options=options) as channel:
            stub = gocosi_pb2_grpc.GocosiRPCStub(channel)
            resp = await stub.GossipReq(request, timeout=timeout)
            # logger.info(resp)
    except grpc.RpcError as rpc_error:
        handle_exception(rpc_error, address, server)


def get_public_key(server, address):
    try:
        with grpc.insecure_channel(address, options=options) as channel:
            stub = gocosi_pb2_grpc.GocosiRPCStub(channel)
            request = Empty()
            resp = stub.GetPubkey(request, timeout=timeout)
            logger.info(f"get {address}'s public keys")
            return resp.publickeys
    except grpc.RpcError as rpc_error:
        handle_exception(rpc_error, address, server)
        return None


def handle_exception(rpc_error, address, server):
    if rpc_error.code() == grpc.StatusCode.UNAVAILABLE:
        logger.error('lose connection from ' + address)
        # if address not in server.unsub_list:
        #     server.unsub_list.append(address)
        # if address in server.sub_list:
        #     server.sub_list.remove(address)
    else:
        logger.error(f"Received unknown RPC error: code={rpc_error.code()} message={rpc_error.details()}")


def get_info(address):
    try:
        with grpc.insecure_channel(address, options=options) as channel:
            stub = gocosi_pb2_grpc.GocosiRPCStub(channel)
            request = Empty()
            resp = stub.Info(request, timeout=timeout)
            return resp
    except grpc.RpcError as rpc_error:
        print(f"Received unknown RPC error: code={rpc_error.code()} message={rpc_error.details()}")
        return None


def sig_msg(address, msg):
    try:
        with grpc.insecure_channel(address, options=options) as channel:
            stub = gocosi_pb2_grpc.GocosiRPCStub(channel)
            request = NewMsgReq(msg=msg)
            resp = stub.NewMsg(request, timeout=timeout)
            return resp
    except grpc.RpcError as rpc_error:
        print(f"Received unknown RPC error: code={rpc_error.code()} message={rpc_error.details()}")
        return None


if __name__ == '__main__':
    time.sleep(60)
    while True:
        latency = []
        for i in range(30000, 30050):
            l = get_info(f"localhost:{i}").latency
            print(l)
        time.sleep(5)