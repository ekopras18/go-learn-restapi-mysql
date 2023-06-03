package models

import (
	"html/template"
	"time"
)

type Blog struct {
	Id        int           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Title     string        `gorm:"type:varchar(255)" json:"title"`
	Author    string        `gorm:"type:varchar(255)" json:"author"`
	Tags      string        `gorm:"type:varchar(255)" json:"tags"` // []string is alias for array
	Content   template.HTML `gorm:"type:longtext" json:"content"`  // []byte is alias for longtext
	CreatedAt time.Time     `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time     `gorm:"type:datetime" json:"updated_at"`
}

func (b *Blog) TableName() string {
	return "blog"
}
