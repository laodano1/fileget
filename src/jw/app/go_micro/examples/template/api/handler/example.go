package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"

	"github.com/micro/examples/template/api/client"
	example "github.com/micro/examples/template/srv/proto/example"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
)

type Example struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Example.Call is called by the API as /template/example/call with post body {"name": "foo"}
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received Example.Call request")

	// extract the client from the context
	exampleClient, ok := client.ExampleFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.template.example.call", "example client not found")
	}

	// make request
	response, err := exampleClient.Call(ctx, &example.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.template.example.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
