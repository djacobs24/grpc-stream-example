# GRPC Stream Example

## Installation
### MacOS
```
brew install go
brew install dep
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Dep
This repo uses dep to manage its dependencies.

`dep init` initializes dep

`dep ensure` syncs the projects dependencies

`dep ensure -update` updates all dependencies

## Generate the Model
`make proto` to generate the protobuf file for the model.

## Run
1. Open two terminal sessions
2. In the first run `make server`
3. In the second run `make client`