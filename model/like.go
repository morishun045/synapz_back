// model - like => likes

package model

import (
	"time"
)

type Like struct {
	ID uint `gorm:"primaryKey"`
	PostID uint `gorm:"index;not null"`
	UserID string `gorm:"index;not null"`
	CreatedAt time.Time

	Post Post
	User User
}
