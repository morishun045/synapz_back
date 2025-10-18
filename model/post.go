// model - post => []posts

package model

import (
	"time"
)

type Post struct {
	ID uint `gorm:"primaryKey"`
	UserID string `gorm:"index;not null"`
	Content string `gorm:"type:text;not null"`
	MediaURL string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
	User User
	Replies []Reply
	Likes []Like
	Retweets []Retweet
	Bookmarks []Bookmark
}
