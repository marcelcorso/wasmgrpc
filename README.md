# wasmgrpc

I wish I could grpc from the browser. With wasm I can "run" go in the browser. Can I then call a grpc server from there? 

No. 

This repo tries to do it. And what you get is 

`2018/07/09 21:05:58 rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = 'transport: Error while dialing dial tcp: Protocol not available'`

on the js console. 


# Want to play?

Be sure you have the helloworld grpc protobufs on your GOPATH

```
go generate google.golang.org/grpc/examples/helloworld/...

```

# Build the wasm sauce

be sure you have some go1.11

```
go get golang.org/x/build/version/go1.11beta1
go1.11beta1 download
``` 

Do generate

```
GOARCH=wasm GOOS=js go1.11beta1 build -o static/client.wasm ./wasmclient
```


# Run the server

The server does 2 things: 

 * static file server that serves the html, js and the wasm 
 * a grpc server that implements the helloworld server
  

build it and run it

```
go1.11beta1 install
wasmgrpc

open localhost:5001 
```

# References

https://medium.zenika.com/go-1-11-webassembly-for-the-gophers-ae4bb8b1ee03

