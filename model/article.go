// model - article => articles

package model

import (
	"time"
)

type Article struct {
	ID uint `gorm:"primaryKey"`
	UserID string `gorm:"index;not null"`
	Title string `gorm:"size:100;not null"`
	Content string `gorm:"type:text;not null"`
	Thumbnail string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User
}
