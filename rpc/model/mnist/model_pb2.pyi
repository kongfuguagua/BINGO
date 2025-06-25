from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class DLModel(_message.Message):
    __slots__ = ("name", "path", "status", "inputType", "outputType")
    NAME_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    INPUTTYPE_FIELD_NUMBER: _ClassVar[int]
    OUTPUTTYPE_FIELD_NUMBER: _ClassVar[int]
    name: str
    path: str
    status: bool
    inputType: str
    outputType: str
    def __init__(self, name: _Optional[str] = ..., path: _Optional[str] = ..., status: bool = ..., inputType: _Optional[str] = ..., outputType: _Optional[str] = ...) -> None: ...

class DLDataOBJ(_message.Message):
    __slots__ = ("Database", "name", "status", "total", "type")
    DATABASE_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    Database: str
    name: str
    status: bool
    total: int
    type: str
    def __init__(self, Database: _Optional[str] = ..., name: _Optional[str] = ..., status: bool = ..., total: _Optional[int] = ..., type: _Optional[str] = ...) -> None: ...

class DLData(_message.Message):
    __slots__ = ("index", "status", "deal")
    INDEX_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    DEAL_FIELD_NUMBER: _ClassVar[int]
    index: str
    status: bool
    deal: str
    def __init__(self, index: _Optional[str] = ..., status: bool = ..., deal: _Optional[str] = ...) -> None: ...

class setDLModelRequest(_message.Message):
    __slots__ = ("path",)
    PATH_FIELD_NUMBER: _ClassVar[int]
    path: str
    def __init__(self, path: _Optional[str] = ...) -> None: ...

class setDLDataRequest(_message.Message):
    __slots__ = ("Database",)
    DATABASE_FIELD_NUMBER: _ClassVar[int]
    Database: str
    def __init__(self, Database: _Optional[str] = ...) -> None: ...
