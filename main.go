package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"https://github.com/Sylph-Ops/Project_2/api"
	db "https://github.com/Sylph-Ops/Project_2/db/sqlc"
)

var (
	dbDriver      = "postgres"
	dbSource      = os.Getenv("DATABASE_URL")
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
