# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: gocosi.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0cgocosi.proto\x12\x06gocosi\" \n\x0fRegisterNodeReq\x12\r\n\x05nodes\x18\x01 \x01(\t\"\x18\n\tNewMsgReq\x12\x0b\n\x03msg\x18\x01 \x01(\t\"\x8b\x01\n\x06Gossip\x12\x0c\n\x04subs\x18\x01 \x03(\t\x12\x0e\n\x06unsubs\x18\x02 \x03(\t\x12\x31\n\nevent_list\x18\x03 \x03(\x0b\x32\x1d.gocosi.Gossip.EventListEntry\x1a\x30\n\x0e\x45ventListEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"q\n\x08InfoResp\x12\x0c\n\x04subs\x18\x01 \x03(\t\x12\x0e\n\x06unsubs\x18\x02 \x03(\t\x12\x0e\n\x06pubkey\x18\x03 \x03(\t\x12\x0f\n\x07latency\x18\x04 \x03(\x01\x12\x12\n\nround_time\x18\x05 \x01(\t\x12\x12\n\nneighbours\x18\x06 \x01(\x05\"}\n\rGetPubkeyResp\x12\x39\n\npublickeys\x18\x01 \x03(\x0b\x32%.gocosi.GetPubkeyResp.PublickeysEntry\x1a\x31\n\x0fPublickeysEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"0\n\nNewMsgResp\x12\x11\n\tsignature\x18\x01 \x01(\t\x12\x0f\n\x07signers\x18\x02 \x03(\x05\"\x1d\n\nCommonResp\x12\x0f\n\x07message\x18\x01 \x01(\t\"\x07\n\x05\x45mpty2\x90\x02\n\tGocosiRPC\x12=\n\x0cRegisterNode\x12\x17.gocosi.RegisterNodeReq\x1a\x12.gocosi.CommonResp\"\x00\x12\x31\n\x06NewMsg\x12\x11.gocosi.NewMsgReq\x1a\x12.gocosi.NewMsgResp\"\x00\x12\x31\n\tGossipReq\x12\x0e.gocosi.Gossip\x1a\x12.gocosi.CommonResp\"\x00\x12)\n\x04Info\x12\r.gocosi.Empty\x1a\x10.gocosi.InfoResp\"\x00\x12\x33\n\tGetPubkey\x12\r.gocosi.Empty\x1a\x15.gocosi.GetPubkeyResp\"\x00\x42\x13Z\x11\x61pp/gocosi/gocosib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'gocosi_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\021app/gocosi/gocosi'
  _GOSSIP_EVENTLISTENTRY._options = None
  _GOSSIP_EVENTLISTENTRY._serialized_options = b'8\001'
  _GETPUBKEYRESP_PUBLICKEYSENTRY._options = None
  _GETPUBKEYRESP_PUBLICKEYSENTRY._serialized_options = b'8\001'
  _REGISTERNODEREQ._serialized_start=24
  _REGISTERNODEREQ._serialized_end=56
  _NEWMSGREQ._serialized_start=58
  _NEWMSGREQ._serialized_end=82
  _GOSSIP._serialized_start=85
  _GOSSIP._serialized_end=224
  _GOSSIP_EVENTLISTENTRY._serialized_start=176
  _GOSSIP_EVENTLISTENTRY._serialized_end=224
  _INFORESP._serialized_start=226
  _INFORESP._serialized_end=339
  _GETPUBKEYRESP._serialized_start=341
  _GETPUBKEYRESP._serialized_end=466
  _GETPUBKEYRESP_PUBLICKEYSENTRY._serialized_start=417
  _GETPUBKEYRESP_PUBLICKEYSENTRY._serialized_end=466
  _NEWMSGRESP._serialized_start=468
  _NEWMSGRESP._serialized_end=516
  _COMMONRESP._serialized_start=518
  _COMMONRESP._serialized_end=547
  _EMPTY._serialized_start=549
  _EMPTY._serialized_end=556
  _GOCOSIRPC._serialized_start=559
  _GOCOSIRPC._serialized_end=831
# @@protoc_insertion_point(module_scope)
