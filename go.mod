module github.com/Buzzology/shippy-user-cli

go 1.16

replace github.com/Buzzology/shippy-service-user => ../shippy-service-user

require (
	github.com/Buzzology/shippy-service-user v0.0.0-00010101000000-000000000000
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
)
