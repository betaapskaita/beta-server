
proto-all: proto-server proto-frontend

proto-server:
	protoc -I proto proto/*.proto --proto_path=$(GOPATH)/src:. --go_out=plugins=grpc:.

proto-frontend:
	protoc -I proto proto/*.proto --js_out=import_style=commonjs:../beta-client/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../beta-client/proto
