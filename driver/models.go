package driver

import "gorm.io/gorm"

//Draft is how article draft is modeled for database
type Draft struct {
	gorm.Model
	Title    string
	Body     string
	Tag      string
	ImageURL string
}

//Article is how article is modeled for database
type Article struct {
	gorm.Model
	UserID   uint
	User     User
	Comment  []Comment `gorm:"foreignKey:ArticleID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title    string
	Body     string
	Tag      string
	ImageURL string
}

//User is how user is modeled for database, a user in this case is admin and a writer.
type User struct {
	gorm.Model
	Articles  []Article `gorm:"foreignKey:UserID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Media     Media     `gorm:"foreignKey:UserID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string
	Pass      string
	Bio       string
	AvatarURL string
}

//Comment is basically how comment is modeled for database
type Comment struct {
	gorm.Model
	ArticleID uint
	Name      string
	Text      string
}

//Media is basically how media is modeled for database
type Media struct {
	gorm.Model
	UserID   uint
	ImageURL string
}
