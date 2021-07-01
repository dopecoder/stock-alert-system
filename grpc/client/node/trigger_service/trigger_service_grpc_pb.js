// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var trigger_service_pb = require('./trigger_service_pb.js');

function serialize_trigger_service_CheckServiceHealthReq(arg) {
  if (!(arg instanceof trigger_service_pb.CheckServiceHealthReq)) {
    throw new Error('Expected argument of type trigger_service.CheckServiceHealthReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_CheckServiceHealthReq(buffer_arg) {
  return trigger_service_pb.CheckServiceHealthReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_CheckServiceHealthRes(arg) {
  if (!(arg instanceof trigger_service_pb.CheckServiceHealthRes)) {
    throw new Error('Expected argument of type trigger_service.CheckServiceHealthRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_CheckServiceHealthRes(buffer_arg) {
  return trigger_service_pb.CheckServiceHealthRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_CreateTriggerReq(arg) {
  if (!(arg instanceof trigger_service_pb.CreateTriggerReq)) {
    throw new Error('Expected argument of type trigger_service.CreateTriggerReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_CreateTriggerReq(buffer_arg) {
  return trigger_service_pb.CreateTriggerReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_CreateTriggerRes(arg) {
  if (!(arg instanceof trigger_service_pb.CreateTriggerRes)) {
    throw new Error('Expected argument of type trigger_service.CreateTriggerRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_CreateTriggerRes(buffer_arg) {
  return trigger_service_pb.CreateTriggerRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_DeleteTriggerReq(arg) {
  if (!(arg instanceof trigger_service_pb.DeleteTriggerReq)) {
    throw new Error('Expected argument of type trigger_service.DeleteTriggerReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_DeleteTriggerReq(buffer_arg) {
  return trigger_service_pb.DeleteTriggerReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_DeleteTriggerRes(arg) {
  if (!(arg instanceof trigger_service_pb.DeleteTriggerRes)) {
    throw new Error('Expected argument of type trigger_service.DeleteTriggerRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_DeleteTriggerRes(buffer_arg) {
  return trigger_service_pb.DeleteTriggerRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_GetTriggerReq(arg) {
  if (!(arg instanceof trigger_service_pb.GetTriggerReq)) {
    throw new Error('Expected argument of type trigger_service.GetTriggerReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_GetTriggerReq(buffer_arg) {
  return trigger_service_pb.GetTriggerReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_GetTriggerRes(arg) {
  if (!(arg instanceof trigger_service_pb.GetTriggerRes)) {
    throw new Error('Expected argument of type trigger_service.GetTriggerRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_GetTriggerRes(buffer_arg) {
  return trigger_service_pb.GetTriggerRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_StartServiceReq(arg) {
  if (!(arg instanceof trigger_service_pb.StartServiceReq)) {
    throw new Error('Expected argument of type trigger_service.StartServiceReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_StartServiceReq(buffer_arg) {
  return trigger_service_pb.StartServiceReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_StartServiceRes(arg) {
  if (!(arg instanceof trigger_service_pb.StartServiceRes)) {
    throw new Error('Expected argument of type trigger_service.StartServiceRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_StartServiceRes(buffer_arg) {
  return trigger_service_pb.StartServiceRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_StopServiceReq(arg) {
  if (!(arg instanceof trigger_service_pb.StopServiceReq)) {
    throw new Error('Expected argument of type trigger_service.StopServiceReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_StopServiceReq(buffer_arg) {
  return trigger_service_pb.StopServiceReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_StopServiceRes(arg) {
  if (!(arg instanceof trigger_service_pb.StopServiceRes)) {
    throw new Error('Expected argument of type trigger_service.StopServiceRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_StopServiceRes(buffer_arg) {
  return trigger_service_pb.StopServiceRes.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_TriggerStatusReq(arg) {
  if (!(arg instanceof trigger_service_pb.TriggerStatusReq)) {
    throw new Error('Expected argument of type trigger_service.TriggerStatusReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_TriggerStatusReq(buffer_arg) {
  return trigger_service_pb.TriggerStatusReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_trigger_service_TriggerStatusRes(arg) {
  if (!(arg instanceof trigger_service_pb.TriggerStatusRes)) {
    throw new Error('Expected argument of type trigger_service.TriggerStatusRes');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_trigger_service_TriggerStatusRes(buffer_arg) {
  return trigger_service_pb.TriggerStatusRes.deserializeBinary(new Uint8Array(buffer_arg));
}


var TriggerServiceService = exports.TriggerServiceService = {
  createTrigger: {
    path: '/trigger_service.TriggerService/createTrigger',
    requestStream: false,
    responseStream: false,
    requestType: trigger_service_pb.CreateTriggerReq,
    responseType: trigger_service_pb.CreateTriggerRes,
    requestSerialize: serialize_trigger_service_CreateTriggerReq,
    requestDeserialize: deserialize_trigger_service_CreateTriggerReq,
    responseSerialize: serialize_trigger_service_CreateTriggerRes,
    responseDeserialize: deserialize_trigger_service_CreateTriggerRes,
  },
  getTrigger: {
    path: '/trigger_service.TriggerService/getTrigger',
    requestStream: false,
    responseStream: false,
    requestType: trigger_service_pb.GetTriggerReq,
    responseType: trigger_service_pb.GetTriggerRes,
    requestSerialize: serialize_trigger_service_GetTriggerReq,
    requestDeserialize: deserialize_trigger_service_GetTriggerReq,
    responseSerialize: serialize_trigger_service_GetTriggerRes,
    responseDeserialize: deserialize_trigger_service_GetTriggerRes,
  },
  getTriggerStatus: {
    path: '/trigger_service.TriggerService/getTriggerStatus',
    requestStream: false,
    responseStream: false,
    requestType: trigger_service_pb.TriggerStatusReq,
    responseType: trigger_service_pb.TriggerStatusRes,
    requestSerialize: serialize_trigger_service_TriggerStatusReq,
    requestDeserialize: deserialize_trigger_service_TriggerStatusReq,
    responseSerialize: serialize_trigger_service_TriggerStatusRes,
    responseDeserialize: deserialize_trigger_service_TriggerStatusRes,
  },
  startService: {
    path: '/trigger_service.TriggerService/startService',
    requestStream: false,
    responseStream: false,
    requestType: trigger_service_pb.StartServiceReq,
    responseType: trigger_service_pb.StartServiceRes,
    requestSerialize: serialize_trigger_service_StartServiceReq,
    requestDeserialize: deserialize_trigger_service_StartServiceReq,
    responseSerialize: serialize_trigger_service_StartServiceRes,
    responseDeserialize: deserialize_trigger_service_StartServiceRes,
  },
  stopService: {
    path: '/trigger_service.TriggerService/stopService',
    requestStream: false,
    responseStream: false,
    requestType: trigger_service_pb.StopServiceReq,
    responseType: trigger_service_pb.StopServiceRes,
    requestSerialize: serialize_trigger_service_StopServiceReq,
    requestDeserialize: deserialize_trigger_service_StopServiceReq,
    responseSerialize: serialize_trigger_service_StopServiceRes,
    responseDeserialize: deserialize_trigger_service_StopServiceRes,
  },
  deteTrigger: {
    path: '/trigger_service.TriggerService/deteTrigger',
    requestStream: false,
    responseStream: false,
    requestType: trigger_service_pb.DeleteTriggerReq,
    responseType: trigger_service_pb.DeleteTriggerRes,
    requestSerialize: serialize_trigger_service_DeleteTriggerReq,
    requestDeserialize: deserialize_trigger_service_DeleteTriggerReq,
    responseSerialize: serialize_trigger_service_DeleteTriggerRes,
    responseDeserialize: deserialize_trigger_service_DeleteTriggerRes,
  },
  checkServiceHealth: {
    path: '/trigger_service.TriggerService/checkServiceHealth',
    requestStream: false,
    responseStream: false,
    requestType: trigger_service_pb.CheckServiceHealthReq,
    responseType: trigger_service_pb.CheckServiceHealthRes,
    requestSerialize: serialize_trigger_service_CheckServiceHealthReq,
    requestDeserialize: deserialize_trigger_service_CheckServiceHealthReq,
    responseSerialize: serialize_trigger_service_CheckServiceHealthRes,
    responseDeserialize: deserialize_trigger_service_CheckServiceHealthRes,
  },
};

exports.TriggerServiceClient = grpc.makeGenericClientConstructor(TriggerServiceService);
