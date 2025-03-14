package main

import (
	"github.com/devanfer02/ratemyubprof/internal/infra/server"
)

func main() {
	server := server.NewHttpServer()
	
	server.Start()

	server.GracefullyShutdown()
}