package main

import (
	"time"
)

type Movie struct {
	Plot            string    `bson:"plot"`
	Genres          []string  `bson:"genres"`
	Runtime         int       `bson:"runtime"`
	Rated           string    `bson:"rated"`
	Cast            []string  `bson:"cast"`
	NumMflixComment int       `bson:"num_mflix_comment"`
	Poster          string    `bson:"poster"`
	Title           string    `bson:"title"`
	Fullplot        string    `bson:"fullplot"`
	Languages       []string  `bson:"languages"`
	Release         time.Time `bson:"release"`
	Directors       []string  `bson:"directors"`
	Writter         []string  `bson:"writter"`
}
