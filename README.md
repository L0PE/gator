# gator

A command-line RSS feed aggregator and user manager written in Go.

## Features
- Register and manage users
- Add, follow, and unfollow RSS feeds
- Aggregate and fetch posts from feeds
- List users, feeds, and followings
- PostgreSQL-backed storage

## Prerequisites
- Go 1.23+
- PostgreSQL database

## Setup
1. **Clone the repository:**
   ```sh
   git clone https://github.com/L0PE/gator.git
   cd gator
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Configure the application:**
   Create a JSON config file at `~/.gatorconfig.json` with the following structure:
   ```json
   {
     "db_url": "postgres://user:password@localhost:5432/dbname?sslmode=disable",
     "current_user_name": ""
   }
   ```
   Replace the `db_url` with your PostgreSQL connection string.

4. **Run database migrations:**
   (Ensure your database schema matches the files in `sql/schema/`.)

5. **Build the CLI:**
   ```sh
   go build -o gator
   ```

## Usage
Run the CLI with a command and its arguments:
```sh
./gator <command> [arguments]
```

### Available Commands
- `register <username>`: Register a new user
- `login <username>`: Set the current user
- `reset`: Delete all users
- `users`: List all users (current user is marked)
- `agg <duration>`: Start feed aggregation loop (e.g., `agg 10m`)
- `addfeed <name> <url>`: Add a new RSS feed (must be logged in)
- `feeds`: List all feeds and their owners
- `follow <feed_url>`: Follow a feed (must be logged in)
- `following`: List feeds the current user is following
- `unfollow <feed_url>`: Unfollow a feed (must be logged in)

### Example
```sh
./gator register alice
./gator login alice
./gator addfeed "Go Blog" "https://blog.golang.org/feed.atom"
./gator feeds
./gator follow "https://blog.golang.org/feed.atom"
./gator following
```

## Configuration
- The config file is located at `~/.gatorconfig.json`.
- It stores the database URL and the current user.

## Database Setup with Goose

This project uses [goose](https://github.com/pressly/goose) for managing PostgreSQL database migrations.

### 1. Install Goose
You can install goose via Go:
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```
Make sure `$GOPATH/bin` (or Go's bin directory) is in your `$PATH`.

### 2. Create Your Database
Create a PostgreSQL database if you haven't already:
```sh
createdb <your_db_name>
```

### 3. Run Migrations
From the `gator` directory, run:
```sh
goose postgres "postgres://user:password@localhost:5432/<your_db_name>?sslmode=disable" up
```
Replace the connection string with your actual database credentials. The migration files are located in `sql/schema/`.

### 4. Migration Directory
If you are not in the `gator` directory, specify the migration directory:
```sh
goose -dir ./sql/schema postgres "postgres://user:password@localhost:5432/<your_db_name>?sslmode=disable" up
```

For more commands and options, see the [goose documentation](https://github.com/pressly/goose). 

## License
MIT