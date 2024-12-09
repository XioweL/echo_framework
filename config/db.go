package config

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
	"log"
	"time"
)

var Pool *pgxpool.Pool

func InitDB() {
	connStr := "postgresql://postgres:123456@localhost:5432/echo_framework"

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal("Error parsing connection string: ", err)
	}

	config.ConnConfig.ConnectTimeout = 5 * time.Second

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err)
	}

	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println("Connected to the database")
}

func CloseDB() {
	Pool.Close()
}
