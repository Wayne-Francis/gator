package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Wayne_Francis/gator/internal/config"
	"github.com/Wayne_Francis/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	db, err := sql.Open("postgres", cfg.Dburl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	s := &state{cfg: &cfg, db: dbQueries}
	cmds := commands{
		handlers: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerFeeds)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerDeleteFeedFollow))
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("please type commands")
	}
	c := command{
		name: args[1],
		Args: args[2:],
	}
	err = cmds.run(s, c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
