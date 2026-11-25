//boot.dev blog aggregator guided project
package main

import (
	"github.com/PotatoesInHD/GatorDontPlay/internal/config"
	"log"
	"os"
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
	gatorState := state {
		configPtr: &cfg,
	}

	//register commands
	cmds := commands {
		registeredCmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	//get args from cli
	args := os.Args
    if len(args) < 2 {
    	log.Fatal("Error: No Args Given!")
    }

    //setup command struct
    cmd := command {
    	name: args[1],
     	args: args[2:],
    }

    //run command
    if err = cmds.run(&gatorState, cmd); err != nil {
    	log.Fatalf("Error: %v", err)
    }
}
