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

### Example:

```
Server: Starting new server
Server: Sending new max: 2
Server: Sending new max: 4
Server: Sending new max: 7
Server: Exit
```

```
Client: 0 sent
Client: 0 sent
Client: 2 sent
Client: New max 2 received
Client: 2 sent
Client: 2 sent
Client: 4 sent
Client: New max 4 received
Client: 4 sent
Client: 2 sent
Client: 2 sent
Client: 7 sent
Client: New max 7 received
Client: Finished with max of 7!
```