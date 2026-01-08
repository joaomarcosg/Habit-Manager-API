package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosg/Habit-Manager-API/internal/api"
	"github.com/joaomarcosg/Habit-Manager-API/internal/services"
	"github.com/joaomarcosg/Habit-Manager-API/internal/store/pgstore"
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
		log.Fatalf("database did not respond: %v", err)
	}

	log.Println("connected to PostgreSQL!")

	userStore := pgstore.NewPGUserStore(pool)

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: *services.NewUserService(&userStore),
	}

	api.BindRoutes()

	fmt.Println("Starting server on port :3080")
	if err := http.ListenAndServe(":3080", api.Router); err != nil {
		panic(err)
	}
}
