package main
import (
    "context"
    "fmt"
    "os"
    "time"
    "github.com/Wayne_Francis/gator/internal/config"
    "github.com/Wayne_Francis/gator/internal/database"
    "github.com/google/uuid"
)

type state struct { 
 db  *database.Queries
 cfg *config.Config
}

type command struct {
	name string
	Args [] string 
}

type commands struct {
	handlers map[string]func(*state,command) error
}


func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
        return fmt.Errorf("please enter a username")
    }
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println(err)
    os.Exit(1)
	}
    err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}
	fmt.Printf("username %v has been set", cmd.Args[0])
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	if len(cmd.name) < 1 {
        return fmt.Errorf("please enter a command to run")
    	}
	handler, ok := c.handlers[cmd.name] 
	if !ok {
	return fmt.Errorf("handler does not exist")
    	}
	return handler(s,cmd)
	}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] =  f
}

func handlerRegister(s *state, cmd command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("please enter a command to run")
    	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
    Name: name,
	ID: uuid.New(),
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	})
	if err != nil {
    fmt.Println(err)
    os.Exit(1)
	}
	err = s.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("new user:name: %v,ID: %v,created: %v,updated: %v", user.Name,user.ID,user.CreatedAt,user.UpdatedAt )
	return nil
	}

	func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
        return fmt.Errorf("reset takes no arguments")
    	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
    	return err
	}
	fmt.Printf("Users have been deleted")
	return nil
}
