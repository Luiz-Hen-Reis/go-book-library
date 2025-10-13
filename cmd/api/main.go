package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Luiz-Hen-Reis/go-book-library/internal/api"
	"github.com/Luiz-Hen-Reis/go-book-library/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("GOBOOK_DATABASE_USER"),
		os.Getenv("GOBOOK_DATABASE_PASSWORD"),
		os.Getenv("GOBOOK_DATABASE_HOST"),
		os.Getenv("GOBOOK_DATABASE_PORT"),
		os.Getenv("GOBOOK_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.Api{
		Router:        chi.NewMux(),
		AuthorService: services.NewAuthorService(pool),
	}

	api.BindRoutes()

	fmt.Println("Starting Server on port :3080")
	if err := http.ListenAndServe("localhost:3080", api.Router); err != nil {
		panic(err)
	}
}