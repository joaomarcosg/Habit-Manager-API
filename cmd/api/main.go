package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("HABIT_MANAGER_DATABASE_USER"),
		os.Getenv("HABIT_MANAGER_DATABASE_PASSWORD"),
		os.Getenv("HABIT_MANAGER_DATABASE_HOST"),
		os.Getenv("HABIT_MANAGER_DATABASE_PORT"),
		os.Getenv("HABIT_MANAGER_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("database did not responsd: %v", err)
	}

	log.Println("connected to PostgreSQL!")
}
