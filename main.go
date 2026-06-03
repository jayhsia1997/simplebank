package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jayhsia1997/simplebank/api"
	db "github.com/jayhsia1997/simplebank/db/sqlc"
	"github.com/jayhsia1997/simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)
	err = server.Start(config.Host + ":" + config.Port)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
