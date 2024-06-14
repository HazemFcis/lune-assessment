package handlers

import (
	"assignment/database"
	"assignment/models"

	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	var movies []models.Movie
	database.DB.Preload("Reviews").Find(&movies)
	return c.JSON(movies)
}

func SearchMoviesByActor(c *fiber.Ctx) error {
	actor := c.Query("actor")
	var movies []models.Movie
	database.DB.Where("actor LIKE ?", "%"+actor+"%").Preload("Reviews").Find(&movies)
	return c.JSON(movies)
}

func GetMoviesSortedByRating(c *fiber.Ctx) error {
	type MovieWithRating struct {
		models.Movie
		AverageRating float64 `gorm:"column:average_rating"`
	}
	var movies []MovieWithRating
	database.DB.Raw(`
		SELECT movies.*, AVG(reviews.stars) as average_rating
		FROM movies
		RIGHT JOIN reviews ON reviews.movie_id = movies.id
		GROUP BY movies.id
		ORDER BY average_rating DESC
	`).Scan(&movies)
	return c.JSON(movies)
}
