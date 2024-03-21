package main

import (
	"context"
	"dependency_injection/repository"
)

type Database interface {
	GetUsers(ctx context.Context) ([]repository.UserTable, error)
}
