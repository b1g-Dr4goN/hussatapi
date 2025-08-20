package main

import (
	"database/sql"
	"log"

	"github.com/b1g-Dr4goN/hussatapi/cmd/api"
	"github.com/b1g-Dr4goN/hussatapi/configs/db"
	"github.com/b1g-Dr4goN/hussatapi/configs/env"
)

func main() {
	database, err := db.NewPostgreSQLStorage(db.PostgresConfig{
		Host:     env.Envs.DBHost,
		Port:     env.Envs.DBPort,
		User:     env.Envs.DBUser,
		Password: env.Envs.DBPassword,
		DBName:   env.Envs.DBName,
		SSLMode:  env.Envs.DBSSLMode,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(database)

	server := api.NewAPIServer(env.Envs.Port, database)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database: Successfully connected!")
}
