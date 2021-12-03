compile:
	protoc -I ./api --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false ./api/*.proto

#for mac
#	protoc --go_out=plugins=grpc:api ./api/club.proto