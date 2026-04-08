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


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

    return func(s *state, cmd command) error {

        user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

            if err != nil {

            return err

                        }

                 return handler(s, cmd, user)

              }

}


func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
        return fmt.Errorf("please enter a username\n")
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
	fmt.Printf("username %v has been set\n", cmd.Args[0])
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	if len(cmd.name) < 1 {
        return fmt.Errorf("please enter a command to run\n")
    	}
	handler, ok := c.handlers[cmd.name] 
	if !ok {
	return fmt.Errorf("handler does not exist\n")
    	}
	return handler(s,cmd)
	}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] =  f
}

func handlerRegister(s *state, cmd command) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("please enter a command to run\n")
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
	fmt.Printf("new user:name: %v,ID: %v,created: %v,updated: %v\n", user.Name,user.ID,user.CreatedAt,user.UpdatedAt )
	return nil
	}

	func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
        return fmt.Errorf("reset takes no arguments\n")
    	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
    	return err
	}
	fmt.Printf("Users have been deleted\n")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
        return fmt.Errorf("users takes no arguments\n")
    	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
     if s.cfg.CurrentUserName == user.Name {
		fmt.Printf("* %v (current)\n", user.Name)
	 } else {
	 fmt.Printf("* %v\n", user.Name)
	 }
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
        return fmt.Errorf("agg takes no arguments\n")
    	}
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	 fmt.Printf("%+v\n", feed)
	return nil
	}


	func handlerAddFeed(s *state, cmd command, user database.User) error {
	 if len(cmd.Args) != 2 {
        return fmt.Errorf("addfeed requires 2 arguments")
    }
    name := cmd.Args[0]
    url := cmd.Args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
    Name: name,
	ID: uuid.New(),
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Url: url,
	UserID: user.ID,
	})
	if err != nil {
    return err
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
	ID: uuid.New(),
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	UserID: user.ID,
	FeedID: feed.ID,
	})
	if err != nil {
    return err
	}
	fmt.Printf("new feed:\nname: %v,\nID: %v,\ncreated: %v,\nupdated: %v\nUrl: %v", feed.Name,feed.ID,feed.CreatedAt,feed.UpdatedAt,feed.Url)
	return nil
	}

	func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
        return fmt.Errorf("feeds takes no arguments\n")
    	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
    user, err := s.db.GetUserById(context.Background(), feed.UserID)
    if err != nil {
        return err
    }
    fmt.Printf("Name: %v\nUrl: %v\nCreated by: %v\n", feed.Name, feed.Url, user.Name)
}
	return nil
}