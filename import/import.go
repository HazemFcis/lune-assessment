package _import

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"

	"assignment/database"
	"assignment/models"
)

func ImportMovies(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	recordCh := make(chan []string, len(records)-1)

	for _, record := range records[1:] {
		recordCh <- record
	}
	close(recordCh)

	for i := 0; i < 10; i++ { // 10 goroutines for parallel processing
		wg.Add(1)
		go func() {
			defer wg.Done()
			for record := range recordCh {
				year, _ := strconv.Atoi(record[2])
				movie := models.Movie{
					Title:           record[0],
					Description:     record[1],
					Year:            year,
					Director:        record[3],
					Actor:           record[4],
					FilmingLocation: record[5],
					Country:         record[6],
				}
				database.DB.Create(&movie)
			}
		}()
	}

	wg.Wait()
}

func ImportReviews(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	recordCh := make(chan []string, len(records)-1)

	for _, record := range records[1:] {
		recordCh <- record
	}
	close(recordCh)

	for i := 0; i < 10; i++ { // 10 goroutines for parallel processing
		wg.Add(1)
		go func() {
			defer wg.Done()
			for record := range recordCh {
				stars, _ := strconv.Atoi(record[2])
				var movie models.Movie
				database.DB.Where("title = ?", record[0]).First(&movie)

				review := models.Review{
					MovieID: movie.ID,
					User:    record[1],
					Stars:   stars,
					Review:  record[3],
				}
				database.DB.Create(&review)
			}
		}()
	}

	wg.Wait()
}
