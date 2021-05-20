package mock

import "thelight/models"

var oneMedia models.Media = models.Media{
	ID:       "1m",
	ImageURL: "https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg",
}

var Medias []models.Media = []models.Media{
	oneMedia,
	oneMedia,
	oneMedia,
	oneMedia,
	oneMedia,
	oneMedia,
}
