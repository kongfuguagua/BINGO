# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: dl.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'dl.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x08\x64l.proto\x12\x05mnist\"\\\n\x07\x44LModel\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04path\x18\x02 \x01(\t\x12\x0e\n\x06status\x18\x03 \x01(\x08\x12\x11\n\tinputType\x18\x04 \x01(\t\x12\x12\n\noutputType\x18\x05 \x01(\t\"X\n\tDLDataOBJ\x12\x10\n\x08\x44\x61tabase\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06status\x18\x03 \x01(\x08\x12\r\n\x05total\x18\x04 \x01(\x03\x12\x0c\n\x04type\x18\x05 \x01(\t\"5\n\x06\x44LData\x12\r\n\x05index\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\x08\x12\x0c\n\x04\x64\x65\x61l\x18\x03 \x01(\t\"!\n\x11setDLModelRequest\x12\x0c\n\x04path\x18\x01 \x01(\t\"$\n\x10setDLDataRequest\x12\x10\n\x08\x44\x61tabase\x18\x01 \x01(\t2~\n\nDLfunction\x12\x37\n\x0binitDLModel\x12\x18.mnist.setDLModelRequest\x1a\x0e.mnist.DLModel\x12\x37\n\ninitDLData\x12\x17.mnist.setDLDataRequest\x1a\x10.mnist.DLDataOBJb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'dl_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_DLMODEL']._serialized_start=19
  _globals['_DLMODEL']._serialized_end=111
  _globals['_DLDATAOBJ']._serialized_start=113
  _globals['_DLDATAOBJ']._serialized_end=201
  _globals['_DLDATA']._serialized_start=203
  _globals['_DLDATA']._serialized_end=256
  _globals['_SETDLMODELREQUEST']._serialized_start=258
  _globals['_SETDLMODELREQUEST']._serialized_end=291
  _globals['_SETDLDATAREQUEST']._serialized_start=293
  _globals['_SETDLDATAREQUEST']._serialized_end=329
  _globals['_DLFUNCTION']._serialized_start=331
  _globals['_DLFUNCTION']._serialized_end=457
# @@protoc_insertion_point(module_scope)
