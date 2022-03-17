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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tmsg.proto\x12\x02pb\"\xda\x01\n\tBroadCast\x12-\n\x08opertype\x18\x01 \x01(\x0e\x32\x16.pb.BroadCast.OperTypeH\x00\x88\x01\x01\x12\x13\n\x06source\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\x0e\n\x06Target\x18\x03 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x04 \x01(\x0c\"S\n\x08OperType\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x12\n\x0e\x42ROADCAST_GAME\x10\x01\x12\x12\n\x0e\x42ROADCAST_ROOM\x10\x02\x12\x12\n\x0e\x42ROADCAST_RANK\x10\x03\x42\x0b\n\t_opertypeB\t\n\x07_source\"Q\n\tHandShake\x12\x0e\n\x06s_host\x18\x01 \x01(\t\x12\x0e\n\x06s_name\x18\x02 \x01(\t\x12\x13\n\x0b\x63urr_branch\x18\x03 \x01(\t\x12\x0f\n\x07s_group\x18\x04 \x01(\x05\"J\n\x04\x45\x63ho\x12\x0e\n\x06s_host\x18\x01 \x01(\t\x12\x0e\n\x06s_name\x18\x02 \x01(\t\x12\x15\n\x08s_online\x18\x03 \x01(\x05H\x00\x88\x01\x01\x42\x0b\n\t_s_online\"\xa8\x01\n\x06Server\x12\x0e\n\x06s_host\x18\x01 \x01(\t\x12\x13\n\x06s_name\x18\x02 \x01(\tH\x00\x88\x01\x01\x12\x18\n\x0b\x63urr_branch\x18\x03 \x01(\tH\x01\x88\x01\x01\x12\x15\n\x08s_online\x18\x04 \x01(\x05H\x02\x88\x01\x01\x12\x14\n\x07s_group\x18\x05 \x01(\x05H\x03\x88\x01\x01\x42\t\n\x07_s_nameB\x0e\n\x0c_curr_branchB\x0b\n\t_s_onlineB\n\n\x08_s_group\"\xa5\x01\n\x0eGameServerList\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12@\n\x10game_server_list\x18\x02 \x03(\x0b\x32&.pb.GameServerList.GameServerListEntry\x1a\x41\n\x13GameServerListEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x19\n\x05value\x18\x02 \x01(\x0b\x32\n.pb.Server:\x02\x38\x01\"\xa3\x02\n\x11SpecialServerList\x12\x43\n\x10room_server_list\x18\x01 \x03(\x0b\x32).pb.SpecialServerList.RoomServerListEntry\x12\x43\n\x10rank_server_list\x18\x02 \x03(\x0b\x32).pb.SpecialServerList.RankServerListEntry\x1a\x41\n\x13RoomServerListEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x19\n\x05value\x18\x02 \x01(\x0b\x32\n.pb.Server:\x02\x38\x01\x1a\x41\n\x13RankServerListEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x19\n\x05value\x18\x02 \x01(\x0b\x32\n.pb.Server:\x02\x38\x01\"\x16\n\x05\x45rror\x12\r\n\x05\x65rror\x18\x01 \x01(\tB\x06Z\x04./pbb\x06proto3')



_BROADCAST = DESCRIPTOR.message_types_by_name['BroadCast']
_HANDSHAKE = DESCRIPTOR.message_types_by_name['HandShake']
_ECHO = DESCRIPTOR.message_types_by_name['Echo']
_SERVER = DESCRIPTOR.message_types_by_name['Server']
_GAMESERVERLIST = DESCRIPTOR.message_types_by_name['GameServerList']
_GAMESERVERLIST_GAMESERVERLISTENTRY = _GAMESERVERLIST.nested_types_by_name['GameServerListEntry']
_SPECIALSERVERLIST = DESCRIPTOR.message_types_by_name['SpecialServerList']
_SPECIALSERVERLIST_ROOMSERVERLISTENTRY = _SPECIALSERVERLIST.nested_types_by_name['RoomServerListEntry']
_SPECIALSERVERLIST_RANKSERVERLISTENTRY = _SPECIALSERVERLIST.nested_types_by_name['RankServerListEntry']
_ERROR = DESCRIPTOR.message_types_by_name['Error']
_BROADCAST_OPERTYPE = _BROADCAST.enum_types_by_name['OperType']
BroadCast = _reflection.GeneratedProtocolMessageType('BroadCast', (_message.Message,), {
  'DESCRIPTOR' : _BROADCAST,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.BroadCast)
  })
_sym_db.RegisterMessage(BroadCast)

HandShake = _reflection.GeneratedProtocolMessageType('HandShake', (_message.Message,), {
  'DESCRIPTOR' : _HANDSHAKE,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.HandShake)
  })
_sym_db.RegisterMessage(HandShake)

Echo = _reflection.GeneratedProtocolMessageType('Echo', (_message.Message,), {
  'DESCRIPTOR' : _ECHO,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.Echo)
  })
_sym_db.RegisterMessage(Echo)

Server = _reflection.GeneratedProtocolMessageType('Server', (_message.Message,), {
  'DESCRIPTOR' : _SERVER,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.Server)
  })
_sym_db.RegisterMessage(Server)

GameServerList = _reflection.GeneratedProtocolMessageType('GameServerList', (_message.Message,), {

  'GameServerListEntry' : _reflection.GeneratedProtocolMessageType('GameServerListEntry', (_message.Message,), {
    'DESCRIPTOR' : _GAMESERVERLIST_GAMESERVERLISTENTRY,
    '__module__' : 'msg_pb2'
    # @@protoc_insertion_point(class_scope:pb.GameServerList.GameServerListEntry)
    })
  ,
  'DESCRIPTOR' : _GAMESERVERLIST,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.GameServerList)
  })
_sym_db.RegisterMessage(GameServerList)
_sym_db.RegisterMessage(GameServerList.GameServerListEntry)

SpecialServerList = _reflection.GeneratedProtocolMessageType('SpecialServerList', (_message.Message,), {

  'RoomServerListEntry' : _reflection.GeneratedProtocolMessageType('RoomServerListEntry', (_message.Message,), {
    'DESCRIPTOR' : _SPECIALSERVERLIST_ROOMSERVERLISTENTRY,
    '__module__' : 'msg_pb2'
    # @@protoc_insertion_point(class_scope:pb.SpecialServerList.RoomServerListEntry)
    })
  ,

  'RankServerListEntry' : _reflection.GeneratedProtocolMessageType('RankServerListEntry', (_message.Message,), {
    'DESCRIPTOR' : _SPECIALSERVERLIST_RANKSERVERLISTENTRY,
    '__module__' : 'msg_pb2'
    # @@protoc_insertion_point(class_scope:pb.SpecialServerList.RankServerListEntry)
    })
  ,
  'DESCRIPTOR' : _SPECIALSERVERLIST,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.SpecialServerList)
  })
_sym_db.RegisterMessage(SpecialServerList)
_sym_db.RegisterMessage(SpecialServerList.RoomServerListEntry)
_sym_db.RegisterMessage(SpecialServerList.RankServerListEntry)

Error = _reflection.GeneratedProtocolMessageType('Error', (_message.Message,), {
  'DESCRIPTOR' : _ERROR,
  '__module__' : 'msg_pb2'
  # @@protoc_insertion_point(class_scope:pb.Error)
  })
_sym_db.RegisterMessage(Error)

if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\004./pb'
  _GAMESERVERLIST_GAMESERVERLISTENTRY._options = None
  _GAMESERVERLIST_GAMESERVERLISTENTRY._serialized_options = b'8\001'
  _SPECIALSERVERLIST_ROOMSERVERLISTENTRY._options = None
  _SPECIALSERVERLIST_ROOMSERVERLISTENTRY._serialized_options = b'8\001'
  _SPECIALSERVERLIST_RANKSERVERLISTENTRY._options = None
  _SPECIALSERVERLIST_RANKSERVERLISTENTRY._serialized_options = b'8\001'
  _BROADCAST._serialized_start=18
  _BROADCAST._serialized_end=236
  _BROADCAST_OPERTYPE._serialized_start=129
  _BROADCAST_OPERTYPE._serialized_end=212
  _HANDSHAKE._serialized_start=238
  _HANDSHAKE._serialized_end=319
  _ECHO._serialized_start=321
  _ECHO._serialized_end=395
  _SERVER._serialized_start=398
  _SERVER._serialized_end=566
  _GAMESERVERLIST._serialized_start=569
  _GAMESERVERLIST._serialized_end=734
  _GAMESERVERLIST_GAMESERVERLISTENTRY._serialized_start=669
  _GAMESERVERLIST_GAMESERVERLISTENTRY._serialized_end=734
  _SPECIALSERVERLIST._serialized_start=737
  _SPECIALSERVERLIST._serialized_end=1028
  _SPECIALSERVERLIST_ROOMSERVERLISTENTRY._serialized_start=896
  _SPECIALSERVERLIST_ROOMSERVERLISTENTRY._serialized_end=961
  _SPECIALSERVERLIST_RANKSERVERLISTENTRY._serialized_start=963
  _SPECIALSERVERLIST_RANKSERVERLISTENTRY._serialized_end=1028
  _ERROR._serialized_start=1030
  _ERROR._serialized_end=1052
# @@protoc_insertion_point(module_scope)
