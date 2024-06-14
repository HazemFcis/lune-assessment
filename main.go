package main

import (
	"assignment/database"
	"assignment/handlers"
	"assignment/import"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDatabase()

	_import.ImportMovies("movies.csv")
	_import.ImportReviews("reviews.csv")

	app.Get("/movies", handlers.GetMovies)
	app.Get("/search", handlers.SearchMoviesByActor)
	app.Get("/movies/sorted", handlers.GetMoviesSortedByRating)

	app.Listen(":3000")
}
