// boot.dev blog aggregator guided project
package main

import (
	"database/sql"
	"log"
	"os"
	"github.com/PotatoesInHD/GatorDontPlay/internal/config"
	"github.com/PotatoesInHD/GatorDontPlay/internal/database"
	_ "github.com/lib/pq"
)


type state struct {
	configPtr	*config.Config
	dbPtr		*database.Queries
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

	//opening database
    db, err := sql.Open("postgres", cfg.DbUrl)
    if err != nil {
    	log.Fatalf("Error opening database: %v", err)
    }
    gatorState.dbPtr = database.New(db)

	//register commands
	cmds := commands {
		registeredCmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
