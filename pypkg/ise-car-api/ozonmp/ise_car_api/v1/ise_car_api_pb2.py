# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/ise_car_api/v1/ise_car_api.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ozonmp/ise_car_api/v1/ise_car_api.proto',
  package='ozonmp.ise_car_api.v1',
  syntax='proto3',
  serialized_options=_b('Z9github.com/ozonmp/ise-car-api/pkg/ise-car-api;ise_car_api'),
  serialized_pb=_b('\n\'ozonmp/ise_car_api/v1/ise_car_api.proto\x12\x15ozonmp.ise_car_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\"+\n\x03\x43\x61r\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x14\n\x05title\x18\x02 \x01(\tR\x05title\"6\n\x14\x44\x65scribeCarV1Request\x12\x1e\n\x06\x63\x61r_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x05\x63\x61rId\"I\n\x15\x44\x65scribeCarV1Response\x12\x30\n\x05value\x18\x01 \x01(\x0b\x32\x1a.ozonmp.ise_car_api.v1.CarR\x05value\"5\n\x12\x43reateCarV1Request\x12\x1f\n\x05title\x18\x01 \x01(\tB\t\xfa\x42\x06r\x04\x10\x01\x18\x64R\x05title\",\n\x13\x43reateCarV1Response\x12\x15\n\x06\x63\x61r_id\x18\x01 \x01(\x04R\x05\x63\x61rId\"\x13\n\x11ListCarsV1Request\"F\n\x12ListCarsV1Response\x12\x30\n\x05items\x18\x01 \x03(\x0b\x32\x1a.ozonmp.ise_car_api.v1.CarR\x05items\"4\n\x12RemoveCarV1Request\x12\x1e\n\x06\x63\x61r_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x05\x63\x61rId\"+\n\x13RemoveCarV1Response\x12\x14\n\x05\x66ound\x18\x01 \x01(\x08R\x05\x66ound2\x90\x04\n\x10IseCarApiService\x12\x85\x01\n\rDescribeCarV1\x12+.ozonmp.ise_car_api.v1.DescribeCarV1Request\x1a,.ozonmp.ise_car_api.v1.DescribeCarV1Response\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/cars/{car_id}\x12s\n\nListCarsV1\x12(.ozonmp.ise_car_api.v1.ListCarsV1Request\x1a).ozonmp.ise_car_api.v1.ListCarsV1Response\"\x10\x82\xd3\xe4\x93\x02\n\x12\x08/v1/cars\x12~\n\x0b\x43reateCarV1\x12).ozonmp.ise_car_api.v1.CreateCarV1Request\x1a*.ozonmp.ise_car_api.v1.CreateCarV1Response\"\x18\x82\xd3\xe4\x93\x02\x12\"\x10/v1/cars/{title}\x12\x7f\n\x0bRemoveCarV1\x12).ozonmp.ise_car_api.v1.RemoveCarV1Request\x1a*.ozonmp.ise_car_api.v1.RemoveCarV1Response\"\x19\x82\xd3\xe4\x93\x02\x13*\x11/v1/cars/{car_id}B;Z9github.com/ozonmp/ise-car-api/pkg/ise-car-api;ise_car_apib\x06proto3')
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,])




_CAR = _descriptor.Descriptor(
  name='Car',
  full_name='ozonmp.ise_car_api.v1.Car',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ozonmp.ise_car_api.v1.Car.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='id', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='ozonmp.ise_car_api.v1.Car.title', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='title', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=121,
  serialized_end=164,
)


_DESCRIBECARV1REQUEST = _descriptor.Descriptor(
  name='DescribeCarV1Request',
  full_name='ozonmp.ise_car_api.v1.DescribeCarV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='car_id', full_name='ozonmp.ise_car_api.v1.DescribeCarV1Request.car_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='carId', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=166,
  serialized_end=220,
)


_DESCRIBECARV1RESPONSE = _descriptor.Descriptor(
  name='DescribeCarV1Response',
  full_name='ozonmp.ise_car_api.v1.DescribeCarV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='ozonmp.ise_car_api.v1.DescribeCarV1Response.value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=222,
  serialized_end=295,
)


_CREATECARV1REQUEST = _descriptor.Descriptor(
  name='CreateCarV1Request',
  full_name='ozonmp.ise_car_api.v1.CreateCarV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='title', full_name='ozonmp.ise_car_api.v1.CreateCarV1Request.title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\006r\004\020\001\030d'), json_name='title', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=297,
  serialized_end=350,
)


_CREATECARV1RESPONSE = _descriptor.Descriptor(
  name='CreateCarV1Response',
  full_name='ozonmp.ise_car_api.v1.CreateCarV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='car_id', full_name='ozonmp.ise_car_api.v1.CreateCarV1Response.car_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='carId', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=352,
  serialized_end=396,
)


_LISTCARSV1REQUEST = _descriptor.Descriptor(
  name='ListCarsV1Request',
  full_name='ozonmp.ise_car_api.v1.ListCarsV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=398,
  serialized_end=417,
)


_LISTCARSV1RESPONSE = _descriptor.Descriptor(
  name='ListCarsV1Response',
  full_name='ozonmp.ise_car_api.v1.ListCarsV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='items', full_name='ozonmp.ise_car_api.v1.ListCarsV1Response.items', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='items', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=419,
  serialized_end=489,
)


_REMOVECARV1REQUEST = _descriptor.Descriptor(
  name='RemoveCarV1Request',
  full_name='ozonmp.ise_car_api.v1.RemoveCarV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='car_id', full_name='ozonmp.ise_car_api.v1.RemoveCarV1Request.car_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='carId', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=491,
  serialized_end=543,
)


_REMOVECARV1RESPONSE = _descriptor.Descriptor(
  name='RemoveCarV1Response',
  full_name='ozonmp.ise_car_api.v1.RemoveCarV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='found', full_name='ozonmp.ise_car_api.v1.RemoveCarV1Response.found', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='found', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=545,
  serialized_end=588,
)

_DESCRIBECARV1RESPONSE.fields_by_name['value'].message_type = _CAR
_LISTCARSV1RESPONSE.fields_by_name['items'].message_type = _CAR
DESCRIPTOR.message_types_by_name['Car'] = _CAR
DESCRIPTOR.message_types_by_name['DescribeCarV1Request'] = _DESCRIBECARV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeCarV1Response'] = _DESCRIBECARV1RESPONSE
DESCRIPTOR.message_types_by_name['CreateCarV1Request'] = _CREATECARV1REQUEST
DESCRIPTOR.message_types_by_name['CreateCarV1Response'] = _CREATECARV1RESPONSE
DESCRIPTOR.message_types_by_name['ListCarsV1Request'] = _LISTCARSV1REQUEST
DESCRIPTOR.message_types_by_name['ListCarsV1Response'] = _LISTCARSV1RESPONSE
DESCRIPTOR.message_types_by_name['RemoveCarV1Request'] = _REMOVECARV1REQUEST
DESCRIPTOR.message_types_by_name['RemoveCarV1Response'] = _REMOVECARV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Car = _reflection.GeneratedProtocolMessageType('Car', (_message.Message,), dict(
  DESCRIPTOR = _CAR,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.Car)
  ))
_sym_db.RegisterMessage(Car)

DescribeCarV1Request = _reflection.GeneratedProtocolMessageType('DescribeCarV1Request', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBECARV1REQUEST,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.DescribeCarV1Request)
  ))
_sym_db.RegisterMessage(DescribeCarV1Request)

DescribeCarV1Response = _reflection.GeneratedProtocolMessageType('DescribeCarV1Response', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBECARV1RESPONSE,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.DescribeCarV1Response)
  ))
_sym_db.RegisterMessage(DescribeCarV1Response)

CreateCarV1Request = _reflection.GeneratedProtocolMessageType('CreateCarV1Request', (_message.Message,), dict(
  DESCRIPTOR = _CREATECARV1REQUEST,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.CreateCarV1Request)
  ))
_sym_db.RegisterMessage(CreateCarV1Request)

CreateCarV1Response = _reflection.GeneratedProtocolMessageType('CreateCarV1Response', (_message.Message,), dict(
  DESCRIPTOR = _CREATECARV1RESPONSE,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.CreateCarV1Response)
  ))
_sym_db.RegisterMessage(CreateCarV1Response)

ListCarsV1Request = _reflection.GeneratedProtocolMessageType('ListCarsV1Request', (_message.Message,), dict(
  DESCRIPTOR = _LISTCARSV1REQUEST,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.ListCarsV1Request)
  ))
_sym_db.RegisterMessage(ListCarsV1Request)

ListCarsV1Response = _reflection.GeneratedProtocolMessageType('ListCarsV1Response', (_message.Message,), dict(
  DESCRIPTOR = _LISTCARSV1RESPONSE,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.ListCarsV1Response)
  ))
_sym_db.RegisterMessage(ListCarsV1Response)

RemoveCarV1Request = _reflection.GeneratedProtocolMessageType('RemoveCarV1Request', (_message.Message,), dict(
  DESCRIPTOR = _REMOVECARV1REQUEST,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.RemoveCarV1Request)
  ))
_sym_db.RegisterMessage(RemoveCarV1Request)

RemoveCarV1Response = _reflection.GeneratedProtocolMessageType('RemoveCarV1Response', (_message.Message,), dict(
  DESCRIPTOR = _REMOVECARV1RESPONSE,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.ise_car_api.v1.RemoveCarV1Response)
  ))
_sym_db.RegisterMessage(RemoveCarV1Response)


DESCRIPTOR._options = None
_DESCRIBECARV1REQUEST.fields_by_name['car_id']._options = None
_CREATECARV1REQUEST.fields_by_name['title']._options = None
_REMOVECARV1REQUEST.fields_by_name['car_id']._options = None

_ISECARAPISERVICE = _descriptor.ServiceDescriptor(
  name='IseCarApiService',
  full_name='ozonmp.ise_car_api.v1.IseCarApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=591,
  serialized_end=1119,
  methods=[
  _descriptor.MethodDescriptor(
    name='DescribeCarV1',
    full_name='ozonmp.ise_car_api.v1.IseCarApiService.DescribeCarV1',
    index=0,
    containing_service=None,
    input_type=_DESCRIBECARV1REQUEST,
    output_type=_DESCRIBECARV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\023\022\021/v1/cars/{car_id}'),
  ),
  _descriptor.MethodDescriptor(
    name='ListCarsV1',
    full_name='ozonmp.ise_car_api.v1.IseCarApiService.ListCarsV1',
    index=1,
    containing_service=None,
    input_type=_LISTCARSV1REQUEST,
    output_type=_LISTCARSV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\n\022\010/v1/cars'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateCarV1',
    full_name='ozonmp.ise_car_api.v1.IseCarApiService.CreateCarV1',
    index=2,
    containing_service=None,
    input_type=_CREATECARV1REQUEST,
    output_type=_CREATECARV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\022\"\020/v1/cars/{title}'),
  ),
  _descriptor.MethodDescriptor(
    name='RemoveCarV1',
    full_name='ozonmp.ise_car_api.v1.IseCarApiService.RemoveCarV1',
    index=3,
    containing_service=None,
    input_type=_REMOVECARV1REQUEST,
    output_type=_REMOVECARV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\023*\021/v1/cars/{car_id}'),
  ),
])
_sym_db.RegisterServiceDescriptor(_ISECARAPISERVICE)

DESCRIPTOR.services_by_name['IseCarApiService'] = _ISECARAPISERVICE

# @@protoc_insertion_point(module_scope)
