# gator

This is an RSS feed agregator backed by PostgreSQL.

## Setup

The prerequisites are [PostgreSQL](https://www.postgresql.org/download/) and [Go](https://go.dev/dl/). For development, [goose](https://github.com/pressly/goose/) is used for migrations in `sql/schema` and [SQLC](https://github.com/sqlc-dev/sqlc/) for compiling queries in `sql/queries`.

After installing Postgres and starting the server (ex. `systemctl start postgresql`), run `psql` as postgres (ex. `sudo -u postgres psql`) and create a database (ex. `CREATE DATABASE gator;`); you may set the password using `ALTER USER postgres PASSWORD 'password';`.

By default, the config file is found at `$HOME/.config/gatorconfig.json`. Here you need to specify the database URL:

```
{
    "db_url": "postgres://postgres:<PASSWORD>@localhost:5432/<DATABASE>?sslmode=disable",
    "currentuser_name": ""
}

```

## Usage

`go install` to install. `gator` followed by one of the following parameters does as follows:

- `register <username>`: creates a user of username <username>.
- `login <username>`: logs into the user.
- `users`: produces a list of existing users.
- `addfeed <feed name> <feed url>`: adds a feed to the database, naming it <name>, and the current user follows it.
- `agg <time duration>`: fetches all of the feeds the user currently follows every <time duration> (minimum time duration: 1 second).
- `follow <feed url>`: adds an existing feed to the current user's following list.
- `unfollow <feed url>`: removes a feed from the current user's following list.
- `feeds`: lists feeds followed by all users (the name, URL, and who the feed was created by).
- `browse <limit>`: prints the title and description of the <limit> most recent posts (default: 2).
- `following`: prints a list of feeds the current users follows.
- `reset`: deletes everything in the database.

## Functionality

The Postgres database consists of four tables: `users`, `feeds`, `feed_follows`, and `posts`. 

A user is created (`register`), the user creates a feed (`addfeed`) and the user/feed pair is added to feed_follows. No posts are as of yet stored in the database, so posts are aggregated periodically, fetching the least recently fetched feed first (`agg`), and stored as a post. Then, they are shown to the user (`browse`).
