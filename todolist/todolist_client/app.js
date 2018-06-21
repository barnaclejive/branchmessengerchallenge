const messages = require('./proto/todolist_pb');
const services = require('./proto/todolist_grpc_pb');
const grpc = require('grpc');
const grpcClient = new services.TodoListServiceClient('localhost:50051', grpc.credentials.createInsecure());

// DEMO
// Instead of building a true UI that exercises the gRPC services, I synchonized service calls to simulate usage in an app.
async function runDemo() {

  // create two lists
  let list1 = await createTodoList("Birthday Party");
  let list2 = await createTodoList("Cleanup Office");

  // add todos to to each list
  let todo1_1 = await createTodo(list1.getId(), "Order cake");
  let todo1_2 = await createTodo(list1.getId(), "Hire clown");
  let todo1_3 = await createTodo(list1.getId(), "Buy beer");
  let todo1_4 = await createTodo(list1.getId(), "Order balloons");

  let todo2_1 = await createTodo(list2.getId(), "Dust monitors");
  let todo2_2 = await createTodo(list2.getId(), "Vaccuum chairs");
  let todo2_3 = await createTodo(list2.getId(), "Sell camcorder");

  // delete a todo
  await deleteTodo(todo1_2.getId());

  // set note and due date on a todo, mark complete
  await updateTodo(todo1_1.getId(), "Order cake", 1, "Tracking #ABCD1234", 0);
  await updateTodo(todo2_1.getId(), "Dust monitors", 1, "Do not use the harsh cleanser!", 1529519641);

  // view todo lists
  let todoLists = await getTodoLists();

  // view all todos for a given list
  let todos_1 = await getTodos(todoLists[0][0]);

  // search all todos - filter by complete status: All, Complete, Incomplete
  // this could be used to build an autocomplete search field where the searchTodos function is called on each keydown event of a search box text field
  let s1_all = await searchTodos("Order", messages.SearchTodosRequest.CompleteStatus.ALL)
  let s1_complete = await searchTodos("Order", messages.SearchTodosRequest.CompleteStatus.COMPLETE)
  let s1_incomplete = await searchTodos("Order", messages.SearchTodosRequest.CompleteStatus.INCOMPLETE)
}

runDemo();


// Service client calls wrapped in js functions that return promises
function createTodoList(name) {

  var request = new messages.CreateTodoListRequest();
  request.setName(name);

  return new Promise((resolve, reject) => {

    grpcClient.createTodoList(request, function(err, response) {

      if (err != null) {
        reject("createTodoList failure")
      }

      console.log('\n* Created todo list', response.getId(), response.getName());
      resolve(response)
    });
  })
}

function getTodoLists() {

  return new Promise((resolve, reject) => {

    grpcClient.getTodoLists(new messages.Empty(), function(err, response) {

      if (err != null) {
        reject("getTodoLists failure")
      }

      todoLists = response.array[0]

      console.log('\n* Got TodoLists:');
      for (var i = 0; i < todoLists.length; i++) {
        console.log('-', 'ID:', todoLists[i][0], 'Name:', todoLists[i][1]);
      }
      resolve(response.array[0])
    });
  })
}

function createTodo(listId, name) {

  var request = new messages.CreateTodoRequest();
  request.setListId(listId);
  request.setName(name);

  return new Promise((resolve, reject) => {

    grpcClient.createTodo(request, function(err, response) {

      if (err != null) {
        reject("createTodo failure")
      }

      console.log('\n* Created todo', response.getId(), response.getName(), 'on todolist', response.getListId());
      resolve(response)
    });
  });
}

function updateTodo(todoId, name, complete, notes, dueDate) {

  var request = new messages.UpdateTodoRequest();
  request.setId(todoId);
  request.setName(name);
  request.setComplete(complete);
  request.setNotes(notes);
  request.setDueDate(dueDate);

  return new Promise((resolve, reject) => {
    grpcClient.updateTodo(request, function(err, response) {

      if (err != null) {
        reject("updateTodo failure")
      }

      console.log('\n* Updated todo', todoId);
      resolve(todoId)
    });
  });
}

function getTodos(listId) {

  var request = new messages.GetTodosForListRequest();
  request.setListId(listId);

  return new Promise((resolve, reject) => {

    grpcClient.getTodosForList(request, function(err, response) {

      if (err != null) {
        reject("getTodos failure")
      }

      todos = response.array[0]

      console.log('\n* Got ', todos.length, 'todos for list', listId);
      for (var i = 0; i < todos.length; i++) {
        console.log('-', 'ID:', todos[i][0], 'List ID:', todos[i][1], 'Name:', todos[i][2]);
      }
      resolve(response.array[0])
    });
  });
}

async function searchTodos(term, status) {

  var request = new messages.SearchTodosRequest();
  request.setTerm(term);
  request.setStatus(status);

  return new Promise((resolve, reject) => {

    grpcClient.searchTodos(request, function(err, response) {

      if (err != null) {
        reject("searchTodos failure")
      }

      todos = response.array[0]

      console.log('\n* Found ', todos.length, 'todos for search term ', term, 'for SearchTodosRequest.CompleteStatus:', status);
      for (var i = 0; i < todos.length; i++) {
        console.log('-', 'ID:', todos[i][0], 'List ID:', todos[i][1], 'Name:', todos[i][2]);
      }
      resolve(response.array[0])
    });
  });
}

function deleteTodo(todoId) {

  var request = new messages.DeleteTodoRequest();
  request.setId(todoId);

  return new Promise((resolve, reject) => {

    grpcClient.deleteTodo(request, function(err, response) {

      if (err != null) {
        reject("deleteTodo failure")
      }

      console.log('\n* Deleted todo', response.getId());
      resolve(true);
    });
  });
}

// I started to go down the path of exposing a real UI to exercise the gRPC client.
// For the purpose of this example though I instead chose to serialize requests to the client to demonostrate the functionality.
// I've left this configurtion of the UI handling here to show where I was going with it.
//
// const express = require('express')
// const path = require('path')
// var exphbs  = require('express-handlebars');
// const app = express()
//
// app.engine('handlebars', exphbs({defaultLayout: 'main'}));
// app.set('view engine', 'handlebars');
//
// app.get('/', function (req, res) {
//   getTodoLists(function(todolists) {
//     res.render('todolists', {title: "Todo Lists", todolists: todolists.array[0]});
//   });
// });
//
// app.get('/todos/:listId', function (req, res) {
//   getTodos(req.params['listId'], function(todos) {
//     res.render('todos', {title: "Todos", todos: todos.array[0]});
//   });
// });
//
// app.get('/todos/edit/:id', function (req, res) {
//   getTodo(req.params['id'], function(todos) {
//     res.render('todos', {title: "Todos", todos: todos.array[0]});
//   });
// });
//
// app.listen(8191);
