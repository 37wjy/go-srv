# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: msg.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tmsg.proto\x12\x02pb\"\x8e\x01\n\tBroadCast\x12\x0e\n\x06source\x18\x02 \x01(\t\x12\x0e\n\x06Target\x18\x03 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x04 \x01(\x0c\"S\n\x08OperType\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x12\n\x0e\x42ROADCAST_GAME\x10\x01\x12\x12\n\x0e\x42ROADCAST_ROOM\x10\x02\x12\x12\n\x0e\x42ROADCAST_RANK\x10\x03\x42\x03Z\x01.b\x06proto3')



_BROADCAST = DESCRIPTOR.message_types_by_name['BroadCast']
_BROADCAST_OPERTYPE = _BROADCAST.enum_types_by_name['OperType']
BroadCast = _reflection.GeneratedProtocolMessageType('BroadCast', (_message.Message,), {
  'DESCRIPTOR' : _BROADCAST,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.BroadCast)
  })
_sym_db.RegisterMessage(BroadCast)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\001.'
  _BROADCAST._serialized_start=18
  _BROADCAST._serialized_end=160
  _BROADCAST_OPERTYPE._serialized_start=77
  _BROADCAST_OPERTYPE._serialized_end=160
# @@protoc_insertion_point(module_scope)
