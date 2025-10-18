// model - user

package model

import (
	"time"
)

type User struct {
	ID string `gorm:"primaryKey"`
	Name string `gorm:"size:50;not null"`
	Username string `gorm:"size:30;unique;not null"`
	Bio string `gorm:"size:255"`
	IconURL string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts []Post
	Articles []Article
	Followers []Follow `gorm:"foreignKey:FollowingID"`
	Following []Follow `gorm:"foreignKey:FollowerID"`
}
