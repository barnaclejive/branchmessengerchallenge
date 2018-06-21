package main

import (
  "log"
  "net"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/barnaclejive/branchmessengerchallenge/todolist/todolist_server/todolist"
  "google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func main()  {

  log.Printf("Configuring listener on port %s...", port)
  lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

  log.Printf("Creating gRPC server...")
	s := grpc.NewServer()

  log.Printf("Registering RegisterTodoListServiceServer...")
	pb.RegisterTodoListServiceServer(s, &server{})
	reflection.Register(s)

  log.Printf("Serving...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// GRPC SERVICE FUNCTIONS
func (s *server) CreateTodoList(ctx context.Context, in *pb.CreateTodoListRequest) (*pb.TodoList, error) {

  todoList := TodoListDao.createTodoList(*in)
	return &todoList, nil
}

func (s *server) GetTodoLists(ctx context.Context, in *pb.Empty) (*pb.GetTodoListsResponse, error) {

  todoLists := TodoListDao.getTodoLists()
  getTodoListsResponse := pb.GetTodoListsResponse{TodoLists: todoLists}
	return &getTodoListsResponse, nil
}

func (s *server) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.Todo, error) {

  todo := TodoListDao.createTodo(*in)
	return &todo, nil
}

func (s *server) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (*pb.Empty, error) {

  TodoListDao.updateTodo(*in)
	return &pb.Empty{}, nil
}

func (s *server) GetTodosForList(ctx context.Context, in *pb.GetTodosForListRequest) (*pb.TodosResponse, error) {

  todos := TodoListDao.getTodosForList(*in)
  todosResponse := pb.TodosResponse{Todos: todos}
	return &todosResponse, nil
}

func (s *server) SearchTodos(ctx context.Context, in *pb.SearchTodosRequest) (*pb.TodosResponse, error) {

  todos := TodoListDao.searchTodos(*in)
  todosResponse := pb.TodosResponse{Todos: todos}
	return &todosResponse, nil
}

func (s *server) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {

  deleteTodoResponse := TodoListDao.deleteTodo(*in)
	return &deleteTodoResponse, nil
}
