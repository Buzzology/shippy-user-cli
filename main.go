package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/<YourUsername>/shippy/shippy-service-user/proto/user"
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
	service := micro.NewService()
	service.Init(

		// TODO: Look into micro actions
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			ctx := context.Background()
			user := &proto.User{
				Name:     name,
				Email:    email,
				Company:  company,
				Password: password,
			}

			if err := createUser(ctx, service, user); err != nil {
				log.Println("Error creating user: ", err.Error())
				return err
			}

			return nil
		}),
	)
}