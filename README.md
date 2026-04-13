# Gator

Gator is a multi-user command line RSS feed aggregator, written as part of the Boot.dev back-end developer course. It can be used to follow updates to websites and also follow other people's feeds as well.

## Prerequisites

- [Go](https://golang.org/dl/) (latest version) - Required to build and install the `gator` CLI.
- [Postgres](https://www.postgresql.org/download/) - Used as the database to store users, feeds, and posts.

## Installation

To install gator on your computer use the following CLI command:

```bash
go install github.com/Wayne-Francis/gator@latest
```
go install will fetch the most up to date version of the program from the repo and install it onto your computer ready for use, if you have met the prerequisites.

## Configuration

To configure gator you need to do the following 3 steps:

1 - Create a .gatorconfig.json in your home directory
2 - The config file is a JSON block , example shown below:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```
3 - Replace the db_url value with your own Postgres connection details.

## Usage 

Gator is used to monitor and gather RSS feeds via the command line. Below is a list of available commands.

To use a command type `gator <command>` and any required arguments:

- `login <name>` - Log in as an existing user
- `register <name>` - Register a new user
- `reset` - Resets the database, this is unrecoverable so only use when necessary
- `users` - Lists all current users in the database
- `agg <interval>` - Continuously collects RSS feeds at a given time interval (e.g. `gator agg 30s`)
- `addfeed <name> <url>` - Add a new feed to follow
- `feeds` - Lists all feeds in the database
- `follow <url>` - Follow an existing feed
- `unfollow <url>` - Unfollow a feed
- `browse [limit]` - Browse your posts, optionally provide a number to limit results
