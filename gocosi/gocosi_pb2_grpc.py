# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import gocosi_pb2 as gocosi__pb2


class GocosiRPCStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RegisterNode = channel.unary_unary(
                '/gocosi.GocosiRPC/RegisterNode',
                request_serializer=gocosi__pb2.RegisterNodeReq.SerializeToString,
                response_deserializer=gocosi__pb2.CommonResp.FromString,
                )
        self.NewMsg = channel.unary_unary(
                '/gocosi.GocosiRPC/NewMsg',
                request_serializer=gocosi__pb2.NewMsgReq.SerializeToString,
                response_deserializer=gocosi__pb2.NewMsgResp.FromString,
                )
        self.GossipReq = channel.unary_unary(
                '/gocosi.GocosiRPC/GossipReq',
                request_serializer=gocosi__pb2.Gossip.SerializeToString,
                response_deserializer=gocosi__pb2.CommonResp.FromString,
                )
        self.Info = channel.unary_unary(
                '/gocosi.GocosiRPC/Info',
                request_serializer=gocosi__pb2.Empty.SerializeToString,
                response_deserializer=gocosi__pb2.InfoResp.FromString,
                )
        self.GetPubkey = channel.unary_unary(
                '/gocosi.GocosiRPC/GetPubkey',
                request_serializer=gocosi__pb2.Empty.SerializeToString,
                response_deserializer=gocosi__pb2.GetPubkeyResp.FromString,
                )


class GocosiRPCServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RegisterNode(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NewMsg(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GossipReq(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Info(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetPubkey(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GocosiRPCServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RegisterNode': grpc.unary_unary_rpc_method_handler(
                    servicer.RegisterNode,
                    request_deserializer=gocosi__pb2.RegisterNodeReq.FromString,
                    response_serializer=gocosi__pb2.CommonResp.SerializeToString,
            ),
            'NewMsg': grpc.unary_unary_rpc_method_handler(
                    servicer.NewMsg,
                    request_deserializer=gocosi__pb2.NewMsgReq.FromString,
                    response_serializer=gocosi__pb2.NewMsgResp.SerializeToString,
            ),
            'GossipReq': grpc.unary_unary_rpc_method_handler(
                    servicer.GossipReq,
                    request_deserializer=gocosi__pb2.Gossip.FromString,
                    response_serializer=gocosi__pb2.CommonResp.SerializeToString,
            ),
            'Info': grpc.unary_unary_rpc_method_handler(
                    servicer.Info,
                    request_deserializer=gocosi__pb2.Empty.FromString,
                    response_serializer=gocosi__pb2.InfoResp.SerializeToString,
            ),
            'GetPubkey': grpc.unary_unary_rpc_method_handler(
                    servicer.GetPubkey,
                    request_deserializer=gocosi__pb2.Empty.FromString,
                    response_serializer=gocosi__pb2.GetPubkeyResp.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'gocosi.GocosiRPC', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class GocosiRPC(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RegisterNode(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/gocosi.GocosiRPC/RegisterNode',
            gocosi__pb2.RegisterNodeReq.SerializeToString,
            gocosi__pb2.CommonResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NewMsg(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/gocosi.GocosiRPC/NewMsg',
            gocosi__pb2.NewMsgReq.SerializeToString,
            gocosi__pb2.NewMsgResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GossipReq(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/gocosi.GocosiRPC/GossipReq',
            gocosi__pb2.Gossip.SerializeToString,
            gocosi__pb2.CommonResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Info(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/gocosi.GocosiRPC/Info',
            gocosi__pb2.Empty.SerializeToString,
            gocosi__pb2.InfoResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetPubkey(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/gocosi.GocosiRPC/GetPubkey',
            gocosi__pb2.Empty.SerializeToString,
            gocosi__pb2.GetPubkeyResp.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
