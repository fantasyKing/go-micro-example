package main

import (
	"fmt"

	greeter "github.com/fantasyKing/go-micro-example/proto"
	srv "github.com/fantasyKing/go-micro-example/service"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	
	greeter.RegisterGreeterHandler(service.Server(), new(srv.Service))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
