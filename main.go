//boot.dev blog aggregator guided project
package main

import (
	"github.com/PotatoesInHD/GatorDontPlay/internal/config"
	"log"
	"fmt"
)


type state struct {
	configPtr	*config.Config
}


func main() {
	//read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	//set config file user
	if err = cfg.SetUser("bobs"); err != nil {
		log.Fatalf("Error: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("%+v\n",cfg)
}
