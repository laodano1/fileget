# GRPC

Contains examples for using [go-micro/service/grpc](https://github.com/micro/go-micro/service/grpc)

- [greeter](greeter) - A greeter example
- [gateway](gateway) - A grpc gateway example

## New service

Check out the [greeter](greeter) example using go-micro/service/grpc

### Import go-micro/service/grpc

```
import "github.com/micro/go-micro/service/grpc"
```

### Create micro.Service

```
service := grpc.NewService()
```

## Pre-existing Service

What if you want to add grpc to a pre-existing service? Use the build pattern for plugins but swap out the client/server.

### Create a plugin file

```
package main

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	cli "github.com/micro/go-plugins/client/grpc"
	srv "github.com/micro/go-plugins/server/grpc"
)

func init() {
	// set the default client
	client.DefaultClient = cli.NewClient()
	// set the default server
	server.DefaultServer = srv.NewServer()
}
```

### Build the binary

```
// For local use
go build -i -o service ./main.go ./plugins.go
```

### Run

Because the default client/server have been replaced we can just run as usual

```
./service
```
