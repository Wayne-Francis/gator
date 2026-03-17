package main
import (
	"github.com/Wayne_Francis/gator/internal/config"
	"fmt"
)

type state struct { 
 cfg *config.Config
}

type command struct {
	name string
	args [] string 
}

type commands struct {
	handlers map[string]func(*state,command) error
}


func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
        return fmt.Errorf("please enter a username")
    }
    err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("username %v has been set", cmd.args[0])
	return nil
}