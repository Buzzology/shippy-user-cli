package main

import (
	"context"
	"fmt"
	"log"
	proto "github.com/Buzzology/shippy-service-user/proto/user"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

func createUser(ctx context.Context, service micro.Service, user *proto.User) error {
	client := proto.NewUserService("shippy.service.user", service.Client())
	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	// print the response
	fmt.Println("Response: ", rsp.User)

	return nil
}

func main() {

	// Create and initialise a new service
	service := micro.NewService(
		micro.Name("go.micro.srv.user-cli")
		micro.Version("latest")
	)

	// Init will parse command line flags
	srv.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", micro.DefaultClient)

	name := "Chris"
	email := "sales@accor.com.au"
	password := "test123"
	company := "ACC"

	r, err := client.Create(context.TODO(), &pb.User {
		Name: name,
		Email: email,
		Password: password,
		Company: company,
	})

	if err != nil {
		log.Fataf("Could not create: %v", err)
	}

	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}

	for _, user := range getAll.Users {
		log.Println(user)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User {
		Email: email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Print("Your access token is: %s \n", authResponse.Token)

	// Let's just exit because
	os.Exit(0)
}