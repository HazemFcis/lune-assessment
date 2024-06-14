# Lune Assesment App

This is a simple Go application built with the Fiber framework and PostgreSQL for managing movie information and reviews. It provides endpoints to import movie data from CSV files, display movie details, search for movies by actor, and sort movies by average ratings.

## Features

- Import movie and review data from CSV files.
- Display a list of all movies.
- Search for movies by actor.
- Sort movies by average rating.

## Requirements

- Go 1.16 or higher
- PostgreSQL
             to run it locally use : docker run --name postgres_container -e POSTGRES_PASSWORD=mysecretpassword -d -p 5432:5432 -v postgres_data:/var/lib/postgresql/data postgres

## Installation

1. Clone the repository:

```sh
git clone https://github.com/HazemFcis/lune-assessment.git

