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
