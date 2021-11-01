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
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ozonmp/ise_car_api/v1/ise_car_api.proto',
  package='ozonmp.omp_template_api.v1',
  syntax='proto3',
  serialized_options=_b('Z>github.com/ozonmp/ise-car-api/pkg/ise-car-api;omp_template_api'),
  serialized_pb=_b('\n\'ozonmp/ise_car_api/v1/ise_car_api.proto\x12\x1aozonmp.omp_template_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"]\n\x03\x43\x61r\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x10\n\x03\x66oo\x18\x02 \x01(\x04R\x03\x66oo\x12\x34\n\x07\x63reated\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x63reated\"E\n\x19\x44\x65scribeTemplateV1Request\x12(\n\x0btemplate_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\ntemplateId\"S\n\x1a\x44\x65scribeTemplateV1Response\x12\x35\n\x05value\x18\x01 \x01(\x0b\x32\x1f.ozonmp.omp_template_api.v1.CarR\x05value2\xc2\x01\n\x15OmpTemplateApiService\x12\xa8\x01\n\x12\x44\x65scribeTemplateV1\x12\x35.ozonmp.omp_template_api.v1.DescribeTemplateV1Request\x1a\x36.ozonmp.omp_template_api.v1.DescribeTemplateV1Response\"#\x82\xd3\xe4\x93\x02\x1d\x12\x1b/v1/templates/{template_id}B@Z>github.com/ozonmp/ise-car-api/pkg/ise-car-api;omp_template_apib\x06proto3')
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_CAR = _descriptor.Descriptor(
  name='Car',
  full_name='ozonmp.omp_template_api.v1.Car',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ozonmp.omp_template_api.v1.Car.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='id', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='foo', full_name='ozonmp.omp_template_api.v1.Car.foo', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='foo', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created', full_name='ozonmp.omp_template_api.v1.Car.created', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='created', file=DESCRIPTOR),
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
  serialized_start=159,
  serialized_end=252,
)


_DESCRIBETEMPLATEV1REQUEST = _descriptor.Descriptor(
  name='DescribeTemplateV1Request',
  full_name='ozonmp.omp_template_api.v1.DescribeTemplateV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='template_id', full_name='ozonmp.omp_template_api.v1.DescribeTemplateV1Request.template_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\372B\0042\002 \000'), json_name='templateId', file=DESCRIPTOR),
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
  serialized_start=254,
  serialized_end=323,
)


_DESCRIBETEMPLATEV1RESPONSE = _descriptor.Descriptor(
  name='DescribeTemplateV1Response',
  full_name='ozonmp.omp_template_api.v1.DescribeTemplateV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='ozonmp.omp_template_api.v1.DescribeTemplateV1Response.value', index=0,
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
  serialized_start=325,
  serialized_end=408,
)

_CAR.fields_by_name['created'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_DESCRIBETEMPLATEV1RESPONSE.fields_by_name['value'].message_type = _CAR
DESCRIPTOR.message_types_by_name['Car'] = _CAR
DESCRIPTOR.message_types_by_name['DescribeTemplateV1Request'] = _DESCRIBETEMPLATEV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeTemplateV1Response'] = _DESCRIBETEMPLATEV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Car = _reflection.GeneratedProtocolMessageType('Car', (_message.Message,), dict(
  DESCRIPTOR = _CAR,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.omp_template_api.v1.Car)
  ))
_sym_db.RegisterMessage(Car)

DescribeTemplateV1Request = _reflection.GeneratedProtocolMessageType('DescribeTemplateV1Request', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBETEMPLATEV1REQUEST,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.omp_template_api.v1.DescribeTemplateV1Request)
  ))
_sym_db.RegisterMessage(DescribeTemplateV1Request)

DescribeTemplateV1Response = _reflection.GeneratedProtocolMessageType('DescribeTemplateV1Response', (_message.Message,), dict(
  DESCRIPTOR = _DESCRIBETEMPLATEV1RESPONSE,
  __module__ = 'ozonmp.ise_car_api.v1.ise_car_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.omp_template_api.v1.DescribeTemplateV1Response)
  ))
_sym_db.RegisterMessage(DescribeTemplateV1Response)


DESCRIPTOR._options = None
_DESCRIBETEMPLATEV1REQUEST.fields_by_name['template_id']._options = None

_OMPTEMPLATEAPISERVICE = _descriptor.ServiceDescriptor(
  name='OmpTemplateApiService',
  full_name='ozonmp.omp_template_api.v1.OmpTemplateApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=411,
  serialized_end=605,
  methods=[
  _descriptor.MethodDescriptor(
    name='DescribeTemplateV1',
    full_name='ozonmp.omp_template_api.v1.OmpTemplateApiService.DescribeTemplateV1',
    index=0,
    containing_service=None,
    input_type=_DESCRIBETEMPLATEV1REQUEST,
    output_type=_DESCRIBETEMPLATEV1RESPONSE,
    serialized_options=_b('\202\323\344\223\002\035\022\033/v1/templates/{template_id}'),
  ),
])
_sym_db.RegisterServiceDescriptor(_OMPTEMPLATEAPISERVICE)

DESCRIPTOR.services_by_name['OmpTemplateApiService'] = _OMPTEMPLATEAPISERVICE

# @@protoc_insertion_point(module_scope)
