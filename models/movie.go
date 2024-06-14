package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title           string
	Description     string
	Year            int
	Director        string
	Actor           string
	FilmingLocation string
	Country         string
	Reviews         []Review `gorm:"foreignKey:MovieID"`
}

type Review struct {
	gorm.Model
	MovieID uint
	User    string
	Stars   int
	Review  string
}
