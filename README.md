



Be sure you have the protobufs on your GOPATH

```
go generate google.golang.org/grpc/examples/helloworld/...

```

# Build the wasm sauce

be sure you have some go1.11

```
go get golang.org/x/build/version/go1.11beta1
go1.11beta1 download
``` 

```
GOARCH=wasm GOOS=js go1.11beta1 build -o static/client.wasm ./wasmclient
``


Run the server

```
go1.11beta1 install
wasmgrpc

open localhost:5001 
```



play with the go client

```
go run goclient/main.go
```

# References

https://medium.zenika.com/go-1-11-webassembly-for-the-gophers-ae4bb8b1ee03

