package mock

import "thelight/models"

var onearticle models.Article = models.Article{
	ID:       "1",
	Date:     "",
	Title:    "",
	Preview:  "",
	Body:     "",
	Tag:      []string{"", "", ""},
	ImageURL: "",
	WriterInfo: models.WriterInfo{
		ID:        "",
		Name:      "",
		AvatarURL: "",
		Bio:       "",
	},
}

var Articles []models.Article = []models.Article{
	onearticle,
	onearticle,
	onearticle,
	onearticle,
	onearticle,
	onearticle,
}
