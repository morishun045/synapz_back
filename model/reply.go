// model - reply => replys

package model

import (
	"time"
)

type Reply struct {
	ID uint `gorm:"primaryKey"`
	PostID uint `gorm:"index;not null"`
	UserID string `gorm:"index;not null"`
	Content string `gorm:"type:text;not null"`
	CreatedAt time.Time

	Post Post
	User User
}
