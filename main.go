package main

import (
	"github.com/PotatoesInHD/GatorDontPlay/internal/config"
	"log"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err = cfg.SetUser("bobs"); err != nil {
		log.Fatalf("Error: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(cfg)
}
