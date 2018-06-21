Notes:
Uses protobug and gRPC.
The server is built with Go, the client is built with node.js.
The reason the client is not also in Go is that I thought I mightend up build a full UI, or at least a JavaScript library.
I first went down the path of trying to export the proto using gRPC-Web so I could built a fairly static web UI whose JS called the server directly.
I didn't find great documentation or support for gRPC-Web (still beta, just recently made public https://github.com/grpc/grpc-web)
So, I endup building the client as a set of functions running on a node.js server. I started to go down the path of building a UI using the node.js server as it's backend (backend for frontend).
The UI code ended up being pretty time consuming and unrelated to the service backend api, so I scraped it in favor of just have inthe client server run some demo interactions with the Go backend.

Running:

# From the repo root run the following command to generate the server (Go) and client (js) messages and service stubs.
protoc --proto_path=proto/ --go_out=plugins=grpc:./todolist_server/todolist/ proto/todolist.proto
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./todolist_client/ --grpc_out=./todolist_client/ --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` proto/todolist.proto
go install ./todolist_server/todolist

# run server on port 50051
cd todolist_server/
go run *.go

# run client
cd todolist_client/
npm install
node app.js
