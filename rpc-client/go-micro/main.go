package main

import (
	"context"
	"fmt"

	micro "github.com/asim/go-micro/v3"
	"github.com/micro/cli"
	proto "github.com/yveshield/go-microservices/go-micro-greeter/pb"
)

func main() {
	service := micro.NewService(
		micro.Name("go-micro-srv-greeter.client"),
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Name to greet",
			}),
	)

	var name string

	service.Init(
		micro.Action(func(c *cli.Context) {
			name = "gomicro RPC call"
			if len(c.String("name")) > 0 {
				name = c.String("name")
			}
		}),
	)

	client := proto.NewGreeterClient("go-micro-srv-greeter", service.Client())

	rsp, err := client.Greeting(context.Background(), &proto.GreetingRequest{Name: name})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Greeting)
}
