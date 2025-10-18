package model

import "time"

type Follow struct {
    ID           uint      `gorm:"primaryKey"`
    FollowerID   string    `gorm:"index;not null"`
    FollowingID  string    `gorm:"index;not null"`
    CreatedAt    time.Time
}