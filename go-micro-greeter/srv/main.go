package main

import (
	"log"

	"github.com/asim/go-micro/v3"
	pb "github.com/yveshield/go-microservices/go-micro-greeter/pb"
	"golang.org/x/net/context"
)

// Greeter implements greeter service.
type Greeter struct{}

// Greeting method implementation.
func (g *Greeter) Greeting(ctx context.Context, in *pb.GreetingRequest, out *pb.GreetingResponse) error {
	out.Greeting = "GO-MICRO Hello " + in.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go-micro-srv-greeter"),
		micro.Version("latest"),
	)

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
