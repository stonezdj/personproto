# Protocol Buff Example

Install protocol buffers:

```bash
brew install protobuf
```

Create directory for generated code:

```bash
mkdir -p protofiles
```

Create a file person.proto in this directory

Run command to generate code in go

```bash
protoc --go_out=$GOPATH/src person.proto
```