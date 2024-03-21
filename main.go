package main

import (
	"context"
	"database/sql"
	"dependency_injection/repository"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type User struct {
	name string
	db   Database
}

func main() {
	connStr := "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf(": %s", err)
	}
	queries := repository.New(db)

	user := User{"Sarthak", queries}
	fmt.Println(user.db.GetUsers(context.Background()))
}
