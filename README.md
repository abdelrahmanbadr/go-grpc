# go-grpc

run command  protoc --go_out=plugins=grpc:. protos/helloworld.proto

problem : protoc-gen-go: program not found or is not executable

solution: 

1- run  which protoc --> to get the path (/usr/local/bin/protoc)

2- run export PATH=$PATH:/usr/local/bin/protoc

3- source ~/.bash_profile
