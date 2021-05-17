package models

type Article struct {
	ID         string
	Date       string
	Title      string
	Preview    string
	Body       string
	Tag        []string
	ImageURL   string
	WriterInfo WriterInfo
}
