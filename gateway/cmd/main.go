package main

import (
	"flag"
	"log"
	"github.com/prathusingh/mint-services/gateway/config"
	"github.com/prathusingh/mint-services/gateway/internal/server"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}


	server.Server()
}