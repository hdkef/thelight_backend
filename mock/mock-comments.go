package mock

import "thelight/models"

var onecomment models.Comment = models.Comment{
	ID:   "",
	Name: "",
	Text: "",
}

var Comments []models.Comment = []models.Comment{
	onecomment,
	onecomment,
	onecomment,
	onecomment,
	onecomment,
	onecomment,
}
