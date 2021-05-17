package mock

import "thelight/models"

var onecomment models.Comment = models.Comment{
	ID:   "1b",
	Name: "Anonymous",
	Text: "Thanks for sharing!",
}

var Comments []models.Comment = []models.Comment{
	onecomment,
	onecomment,
	onecomment,
	onecomment,
	onecomment,
	onecomment,
}
