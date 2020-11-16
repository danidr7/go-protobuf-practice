# go-protobuf-practice

A project for test some things around protobuf

### how to play

The objects are defined in `person.proto` file, 
this file is the input for 'protobuf' compiler generate all the required code for play around.

For compile it:
```
make compile
```
Then it can be ran:
```
make run
```
For a simple benchmark between protobuf and json:
```
make benchmark
```

Benchmark results:
```
BenchmarkProtobufJson/protobuf_encoding-8                3516054               302 ns/op              32 B/op          1 allocs/op
BenchmarkProtobufJson/protobuf_decoding-8                2424274               511 ns/op             184 B/op          5 allocs/op
BenchmarkProtobufJson/json_encoding-8                    2018937               558 ns/op              80 B/op          1 allocs/op
BenchmarkProtobufJson/json_decoding-8                     520821              2411 ns/op             464 B/op         12 allocs/op
```

### references
- https://developers.google.com/protocol-buffers
- https://developers.google.com/protocol-buffers/docs/gotutorial
- https://tutorialedge.net/golang/go-protocol-buffer-tutorial/