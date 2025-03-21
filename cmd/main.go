package main

import (
	"context"

	"github.com/matzefriedrich/parsley/pkg/bootstrap"

	"go.db.restapi/server"
)

func main() {
	context := context.Background()

	// Runs a Fiber instance as a Parsley-enabled app
	bootstrap.RunParsleyApplication(context, server.NewApp)
}
