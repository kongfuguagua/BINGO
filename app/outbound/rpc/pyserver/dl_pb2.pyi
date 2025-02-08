from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class DLapp(_message.Message):
    __slots__ = ("metadata", "spec")
    METADATA_FIELD_NUMBER: _ClassVar[int]
    SPEC_FIELD_NUMBER: _ClassVar[int]
    metadata: DLMetadata
    spec: DLSpec
    def __init__(self, metadata: _Optional[_Union[DLMetadata, _Mapping]] = ..., spec: _Optional[_Union[DLSpec, _Mapping]] = ...) -> None: ...

class DLMetadata(_message.Message):
    __slots__ = ("id", "Namespace", "DLName")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    DLNAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    Namespace: str
    DLName: str
    def __init__(self, id: _Optional[str] = ..., Namespace: _Optional[str] = ..., DLName: _Optional[str] = ...) -> None: ...

class DLSpec(_message.Message):
    __slots__ = ("model", "dataObj")
    MODEL_FIELD_NUMBER: _ClassVar[int]
    DATAOBJ_FIELD_NUMBER: _ClassVar[int]
    model: DLModel
    dataObj: DLDataOBJ
    def __init__(self, model: _Optional[_Union[DLModel, _Mapping]] = ..., dataObj: _Optional[_Union[DLDataOBJ, _Mapping]] = ...) -> None: ...

class DLCreateSpec(_message.Message):
    __slots__ = ("Namespace", "DLName", "modelName")
    NAMESPACE_FIELD_NUMBER: _ClassVar[int]
    DLNAME_FIELD_NUMBER: _ClassVar[int]
    MODELNAME_FIELD_NUMBER: _ClassVar[int]
    Namespace: str
    DLName: str
    modelName: str
    def __init__(self, Namespace: _Optional[str] = ..., DLName: _Optional[str] = ..., modelName: _Optional[str] = ...) -> None: ...

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

class DLGetRequestById(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DLGetResponseById(_message.Message):
    __slots__ = ("dlApp",)
    DLAPP_FIELD_NUMBER: _ClassVar[int]
    dlApp: DLapp
    def __init__(self, dlApp: _Optional[_Union[DLapp, _Mapping]] = ...) -> None: ...

class DLCreateRequest(_message.Message):
    __slots__ = ("spec",)
    SPEC_FIELD_NUMBER: _ClassVar[int]
    spec: DLCreateSpec
    def __init__(self, spec: _Optional[_Union[DLCreateSpec, _Mapping]] = ...) -> None: ...

class DLCreateResponse(_message.Message):
    __slots__ = ("dlApp",)
    DLAPP_FIELD_NUMBER: _ClassVar[int]
    dlApp: DLapp
    def __init__(self, dlApp: _Optional[_Union[DLapp, _Mapping]] = ...) -> None: ...

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
