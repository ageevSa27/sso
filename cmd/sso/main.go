package main

import (
	"fmt"

	"sso/internal/config"
)

func main() {
	// init config
	cfg := config.MustLoad()

	fmt.Printf("%+v", cfg)
	// init logger

	//init app

	// run gRPC server

}
