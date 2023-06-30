package models

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	ChannelID int       `gorm:"not null" json:"channelId"`
	Channel   Channel   `gorm:"foreignKey:ChannelID"`
	UserID    uint      `json:"userId"`
}
