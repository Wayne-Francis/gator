package main
import _ "github.com/lib/pq"
import (
	"log"
	 "github.com/Wayne_Francis/gator/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	s := &state{cfg: &cfg}
	cmds := commands{
    	handlers: map[string]func(*state, command) error{},
	}
    	cmds.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("please type commands")
}
	c := command{
	name : args[1],
	args : args[2:],
	}
	err = cmds.run(s, c) 
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}