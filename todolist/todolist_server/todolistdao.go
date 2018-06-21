package main

import (
  "log"
  "os"

  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  pb "github.com/barnaclejive/branchmessengerchallenge/todolist/todolist_server/todolist"
)

// The database will be exposed as a singleton of the TodoListDaoManager type.
// Database operations are exposed as methods on the TodoListDao object.
type TodoListDaoManager struct {
  db *sql.DB
}
var TodoListDao TodoListDaoManager

// Initalize the database.
// To keep things simple for this exercise I chose to use a on disk sqlite database.
// You can re-create it on startup my manually setting 'initialDatabase' to true.
const initialDatabase = true

func init() {

  log.Println("Initializing database...")

  if initialDatabase {
    os.Remove("./todolist.db")
  }

  db, err := sql.Open("sqlite3", "./todolist.db")
  if err != nil {
    log.Fatal(err)
  }

  if initialDatabase {

    schemaSql := `
    CREATE TABLE todolist (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT
    );
    CREATE TABLE todo (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT,
      complete TINYINT,
      notes TEXT,
      due_date INTEGER,
      list_id,
      FOREIGN KEY(list_id) REFERENCES todolist(id)
    );
    `
    _, err = db.Exec(schemaSql)
    if err != nil {
      log.Fatal(err)
    }
  }

  TodoListDao = TodoListDaoManager{db}
}

// CREATE TODOLIST
func (m TodoListDaoManager) createTodoList(createTodoListRequest pb.CreateTodoListRequest) pb.TodoList {

  tx, err := m.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`INSERT INTO todolist (name) VALUES (?);`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

  res, err := stmt.Exec(createTodoListRequest.Name)
  if err != nil && res != nil {
    log.Fatal(err)
  }
  tx.Commit()
  id, err := res.LastInsertId()

  todoList := pb.TodoList{Id: int32(id), Name: createTodoListRequest.Name}
  return todoList
}

// GET TODOLISTS
func (m TodoListDaoManager) getTodoLists() []*pb.TodoList {

  selectSql := `SELECT id, name FROM todolist ORDER BY id;`
  stmt, err := m.db.Prepare(selectSql)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  rows, err := stmt.Query()
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  results := []*pb.TodoList{}
  for rows.Next() {

    var id int32
    var name string

    err = rows.Scan(&id, &name)
    if err != nil {
      log.Fatal(err)
    }

    todoList := pb.TodoList{Id: id, Name: name}
    results = append(results, &todoList)
  }

  return results
}

// CREATE TODO
func (m TodoListDaoManager) createTodo(createTodoRequest pb.CreateTodoRequest) pb.Todo {

  tx, err := m.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO todo (name, complete, notes, due_date, list_id) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

  complete := 0
  if createTodoRequest.Complete {
    complete = 1
  }

  res, err := stmt.Exec(createTodoRequest.Name, complete, createTodoRequest.Notes, createTodoRequest.DueDate, createTodoRequest.ListId)
  if err != nil && res != nil {
    log.Fatal(err)
  }
  tx.Commit()
  id, err := res.LastInsertId()

  todo := pb.Todo{Id: int32(id), Name: createTodoRequest.Name, ListId: createTodoRequest.ListId}
  return todo
}

// UPDATE TODO
func (m TodoListDaoManager) updateTodo(updateTodoRequest pb.UpdateTodoRequest) {

  tx, err := m.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("UPDATE todo SET name = ?, complete = ?, notes = ?, due_date = ? WHERE id = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

  complete := 0
  if updateTodoRequest.Complete {
    complete = 1
  }

  _, err = stmt.Exec(updateTodoRequest.Name, complete, updateTodoRequest.Notes, updateTodoRequest.DueDate, updateTodoRequest.Id)
  if err != nil {
    log.Fatal(err)
  }
  tx.Commit()
}

// GET TODOS FOR LIST
func (m TodoListDaoManager) getTodosForList(getTodosForListRequest pb.GetTodosForListRequest) []*pb.Todo {

  selectSql := `SELECT id, name, complete, notes, due_date, list_id FROM todo WHERE list_id = ? ORDER BY id;`
  stmt, err := m.db.Prepare(selectSql)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  rows, err := stmt.Query(getTodosForListRequest.ListId)
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  results := []*pb.Todo{}
  for rows.Next() {

    var id int32
    var name string
    var complete bool
    var notes string
    var dueDate int32
    var listId int32

    err = rows.Scan(&id, &name, &complete, &notes, &dueDate, &listId)
    if err != nil {
      log.Fatal(err)
    }

    todo := pb.Todo{Id: id, ListId: listId, Name: name, Notes: notes, DueDate: dueDate, Complete: complete}
    results = append(results, &todo)
  }

  return results
}

// SEARCH TODOS
func (m TodoListDaoManager) searchTodos(searchTodosRequest pb.SearchTodosRequest) []*pb.Todo {

  selectSql := `
  SELECT id, name, complete, notes, due_date, list_id FROM todo WHERE name LIKE ?
  `

  var complete int = -1
  if searchTodosRequest.GetStatus() != pb.SearchTodosRequest_ALL {
    selectSql += `
    AND complete = ?
    `
    if searchTodosRequest.GetStatus() == pb.SearchTodosRequest_COMPLETE {
      complete = 1
    } else if searchTodosRequest.GetStatus() == pb.SearchTodosRequest_INCOMPLETE {
      complete = 0
    }
  }

  selectSql += `
  ORDER BY id;
  `

  stmt, err := m.db.Prepare(selectSql)
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  likeArg := `%` + searchTodosRequest.GetTerm() + `%`
  var rows *sql.Rows
  if complete != -1 {
    rows, err = stmt.Query(likeArg, complete)
  } else {
    rows, err = stmt.Query(likeArg)
  }

  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  results := []*pb.Todo{}
  for rows.Next() {

    var id int32
    var name string
    var complete bool
    var notes string
    var dueDate int32
    var listId int32

    err = rows.Scan(&id, &name, &complete, &notes, &dueDate, &listId)
    if err != nil {
      log.Fatal(err)
    }

    todo := pb.Todo{Id: id, ListId: listId, Name: name, Notes: notes, DueDate: dueDate, Complete: complete}
    results = append(results, &todo)
  }

  return results
}

// DELETE TODO
func (m TodoListDaoManager) deleteTodo(deleteTodoRequest pb.DeleteTodoRequest) pb.DeleteTodoResponse {

  tx, err := m.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("DELETE FROM todo WHERE id = ?;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

  _, err = stmt.Exec(deleteTodoRequest.Id)
  if err != nil {
    log.Fatal(err)
  }
  tx.Commit()

  deleteTodoResponse := pb.DeleteTodoResponse{Id: deleteTodoRequest.Id}
  return deleteTodoResponse
}
