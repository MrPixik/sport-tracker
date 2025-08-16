package main

import (
	"fmt"
	"github.com/mrpixik/sport-tracker/internal/config"
)

func main() {
	// $env:auth_config_path="configs/auth/local.yaml"
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//TODO: init logger

	//TODO: init storage

	//TODO: init router

	//TODO: init server
}
