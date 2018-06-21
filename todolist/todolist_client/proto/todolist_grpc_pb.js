// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var proto_todolist_pb = require('../proto/todolist_pb.js');

function serialize_todolist_CreateTodoListRequest(arg) {
  if (!(arg instanceof proto_todolist_pb.CreateTodoListRequest)) {
    throw new Error('Expected argument of type todolist.CreateTodoListRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_CreateTodoListRequest(buffer_arg) {
  return proto_todolist_pb.CreateTodoListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_CreateTodoRequest(arg) {
  if (!(arg instanceof proto_todolist_pb.CreateTodoRequest)) {
    throw new Error('Expected argument of type todolist.CreateTodoRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_CreateTodoRequest(buffer_arg) {
  return proto_todolist_pb.CreateTodoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_DeleteTodoRequest(arg) {
  if (!(arg instanceof proto_todolist_pb.DeleteTodoRequest)) {
    throw new Error('Expected argument of type todolist.DeleteTodoRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_DeleteTodoRequest(buffer_arg) {
  return proto_todolist_pb.DeleteTodoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_DeleteTodoResponse(arg) {
  if (!(arg instanceof proto_todolist_pb.DeleteTodoResponse)) {
    throw new Error('Expected argument of type todolist.DeleteTodoResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_DeleteTodoResponse(buffer_arg) {
  return proto_todolist_pb.DeleteTodoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_Empty(arg) {
  if (!(arg instanceof proto_todolist_pb.Empty)) {
    throw new Error('Expected argument of type todolist.Empty');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_Empty(buffer_arg) {
  return proto_todolist_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_GetTodoListsResponse(arg) {
  if (!(arg instanceof proto_todolist_pb.GetTodoListsResponse)) {
    throw new Error('Expected argument of type todolist.GetTodoListsResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_GetTodoListsResponse(buffer_arg) {
  return proto_todolist_pb.GetTodoListsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_GetTodosForListRequest(arg) {
  if (!(arg instanceof proto_todolist_pb.GetTodosForListRequest)) {
    throw new Error('Expected argument of type todolist.GetTodosForListRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_GetTodosForListRequest(buffer_arg) {
  return proto_todolist_pb.GetTodosForListRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_SearchTodosRequest(arg) {
  if (!(arg instanceof proto_todolist_pb.SearchTodosRequest)) {
    throw new Error('Expected argument of type todolist.SearchTodosRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_SearchTodosRequest(buffer_arg) {
  return proto_todolist_pb.SearchTodosRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_Todo(arg) {
  if (!(arg instanceof proto_todolist_pb.Todo)) {
    throw new Error('Expected argument of type todolist.Todo');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_Todo(buffer_arg) {
  return proto_todolist_pb.Todo.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_TodoList(arg) {
  if (!(arg instanceof proto_todolist_pb.TodoList)) {
    throw new Error('Expected argument of type todolist.TodoList');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_TodoList(buffer_arg) {
  return proto_todolist_pb.TodoList.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_TodosResponse(arg) {
  if (!(arg instanceof proto_todolist_pb.TodosResponse)) {
    throw new Error('Expected argument of type todolist.TodosResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_TodosResponse(buffer_arg) {
  return proto_todolist_pb.TodosResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_todolist_UpdateTodoRequest(arg) {
  if (!(arg instanceof proto_todolist_pb.UpdateTodoRequest)) {
    throw new Error('Expected argument of type todolist.UpdateTodoRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_todolist_UpdateTodoRequest(buffer_arg) {
  return proto_todolist_pb.UpdateTodoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var TodoListServiceService = exports.TodoListServiceService = {
  createTodoList: {
    path: '/todolist.TodoListService/CreateTodoList',
    requestStream: false,
    responseStream: false,
    requestType: proto_todolist_pb.CreateTodoListRequest,
    responseType: proto_todolist_pb.TodoList,
    requestSerialize: serialize_todolist_CreateTodoListRequest,
    requestDeserialize: deserialize_todolist_CreateTodoListRequest,
    responseSerialize: serialize_todolist_TodoList,
    responseDeserialize: deserialize_todolist_TodoList,
  },
  getTodoLists: {
    path: '/todolist.TodoListService/GetTodoLists',
    requestStream: false,
    responseStream: false,
    requestType: proto_todolist_pb.Empty,
    responseType: proto_todolist_pb.GetTodoListsResponse,
    requestSerialize: serialize_todolist_Empty,
    requestDeserialize: deserialize_todolist_Empty,
    responseSerialize: serialize_todolist_GetTodoListsResponse,
    responseDeserialize: deserialize_todolist_GetTodoListsResponse,
  },
  createTodo: {
    path: '/todolist.TodoListService/CreateTodo',
    requestStream: false,
    responseStream: false,
    requestType: proto_todolist_pb.CreateTodoRequest,
    responseType: proto_todolist_pb.Todo,
    requestSerialize: serialize_todolist_CreateTodoRequest,
    requestDeserialize: deserialize_todolist_CreateTodoRequest,
    responseSerialize: serialize_todolist_Todo,
    responseDeserialize: deserialize_todolist_Todo,
  },
  updateTodo: {
    path: '/todolist.TodoListService/UpdateTodo',
    requestStream: false,
    responseStream: false,
    requestType: proto_todolist_pb.UpdateTodoRequest,
    responseType: proto_todolist_pb.Empty,
    requestSerialize: serialize_todolist_UpdateTodoRequest,
    requestDeserialize: deserialize_todolist_UpdateTodoRequest,
    responseSerialize: serialize_todolist_Empty,
    responseDeserialize: deserialize_todolist_Empty,
  },
  getTodosForList: {
    path: '/todolist.TodoListService/GetTodosForList',
    requestStream: false,
    responseStream: false,
    requestType: proto_todolist_pb.GetTodosForListRequest,
    responseType: proto_todolist_pb.TodosResponse,
    requestSerialize: serialize_todolist_GetTodosForListRequest,
    requestDeserialize: deserialize_todolist_GetTodosForListRequest,
    responseSerialize: serialize_todolist_TodosResponse,
    responseDeserialize: deserialize_todolist_TodosResponse,
  },
  searchTodos: {
    path: '/todolist.TodoListService/SearchTodos',
    requestStream: false,
    responseStream: false,
    requestType: proto_todolist_pb.SearchTodosRequest,
    responseType: proto_todolist_pb.TodosResponse,
    requestSerialize: serialize_todolist_SearchTodosRequest,
    requestDeserialize: deserialize_todolist_SearchTodosRequest,
    responseSerialize: serialize_todolist_TodosResponse,
    responseDeserialize: deserialize_todolist_TodosResponse,
  },
  deleteTodo: {
    path: '/todolist.TodoListService/DeleteTodo',
    requestStream: false,
    responseStream: false,
    requestType: proto_todolist_pb.DeleteTodoRequest,
    responseType: proto_todolist_pb.DeleteTodoResponse,
    requestSerialize: serialize_todolist_DeleteTodoRequest,
    requestDeserialize: deserialize_todolist_DeleteTodoRequest,
    responseSerialize: serialize_todolist_DeleteTodoResponse,
    responseDeserialize: deserialize_todolist_DeleteTodoResponse,
  },
};

exports.TodoListServiceClient = grpc.makeGenericClientConstructor(TodoListServiceService);
