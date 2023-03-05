from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class CommonResp(_message.Message):
    __slots__ = ["message"]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class Empty(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetPubkeyResp(_message.Message):
    __slots__ = ["publickeys"]
    class PublickeysEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    PUBLICKEYS_FIELD_NUMBER: _ClassVar[int]
    publickeys: _containers.ScalarMap[str, str]
    def __init__(self, publickeys: _Optional[_Mapping[str, str]] = ...) -> None: ...

class Gossip(_message.Message):
    __slots__ = ["event_list", "subs", "unsubs"]
    class EventListEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    EVENT_LIST_FIELD_NUMBER: _ClassVar[int]
    SUBS_FIELD_NUMBER: _ClassVar[int]
    UNSUBS_FIELD_NUMBER: _ClassVar[int]
    event_list: _containers.ScalarMap[str, str]
    subs: _containers.RepeatedScalarFieldContainer[str]
    unsubs: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, subs: _Optional[_Iterable[str]] = ..., unsubs: _Optional[_Iterable[str]] = ..., event_list: _Optional[_Mapping[str, str]] = ...) -> None: ...

class InfoResp(_message.Message):
    __slots__ = ["latency", "neighbours", "pubkey", "round_time", "subs", "unsubs"]
    LATENCY_FIELD_NUMBER: _ClassVar[int]
    NEIGHBOURS_FIELD_NUMBER: _ClassVar[int]
    PUBKEY_FIELD_NUMBER: _ClassVar[int]
    ROUND_TIME_FIELD_NUMBER: _ClassVar[int]
    SUBS_FIELD_NUMBER: _ClassVar[int]
    UNSUBS_FIELD_NUMBER: _ClassVar[int]
    latency: _containers.RepeatedScalarFieldContainer[float]
    neighbours: int
    pubkey: _containers.RepeatedScalarFieldContainer[str]
    round_time: str
    subs: _containers.RepeatedScalarFieldContainer[str]
    unsubs: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, subs: _Optional[_Iterable[str]] = ..., unsubs: _Optional[_Iterable[str]] = ..., pubkey: _Optional[_Iterable[str]] = ..., latency: _Optional[_Iterable[float]] = ..., round_time: _Optional[str] = ..., neighbours: _Optional[int] = ...) -> None: ...

class NewMsgReq(_message.Message):
    __slots__ = ["msg"]
    MSG_FIELD_NUMBER: _ClassVar[int]
    msg: str
    def __init__(self, msg: _Optional[str] = ...) -> None: ...

class NewMsgResp(_message.Message):
    __slots__ = ["signature", "signers"]
    class SignersEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: int
        def __init__(self, key: _Optional[str] = ..., value: _Optional[int] = ...) -> None: ...
    SIGNATURE_FIELD_NUMBER: _ClassVar[int]
    SIGNERS_FIELD_NUMBER: _ClassVar[int]
    signature: str
    signers: _containers.ScalarMap[str, int]
    def __init__(self, signature: _Optional[str] = ..., signers: _Optional[_Mapping[str, int]] = ...) -> None: ...

class RegisterNodeReq(_message.Message):
    __slots__ = ["nodes"]
    NODES_FIELD_NUMBER: _ClassVar[int]
    nodes: str
    def __init__(self, nodes: _Optional[str] = ...) -> None: ...
