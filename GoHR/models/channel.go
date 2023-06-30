package models

import (
	"time"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Channel struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null;unique;size:50" json:"name"`
	Slug        string    `gorm:"size:255" json:"slug"`
	Description string    `gorm:"not null;size:500" json:"description"`
	PostCount   int       `json:"postCount"`
	Photo       string    `gorm:"default:no-photo.jpg" json:"photo"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	Posts       []Post    `gorm:"many2many:channel_posts;" json:"posts"`
}

// Refrence: https://gorm.io/docs/hooks.html
// BeforeSave hook
func (c *Channel) BeforeSave(tx *gorm.DB) (err error) {
	c.Slug = slug.Make(c.Name)

	return nil
}

// BeforeDelete hook
func (c *Channel) BeforeDelete(tx *gorm.DB) (err error) {
	return nil
}
