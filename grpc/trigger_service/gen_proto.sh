
protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative trigger_service.proto 

npm install -g grpc-tools
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:../client/node/trigger_service/ --grpc_out=grpc_js:../client/node/trigger_service/ trigger_service.proto