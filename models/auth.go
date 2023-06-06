package models

import "time"

type Users struct {
	Id        int       `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Username  string    `validate:"required" gorm:"type:varchar(255)" json:"username"`
	Name      string    `validate:"required" gorm:"type:varchar(255)" json:"name"`
	Email     string    `validate:"required,email" gorm:"type:varchar(255)" json:"email"`
	Password  string    `validate:"required" gorm:"type:varchar(255)" json:"password"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
}

func (b *Users) TableName() string {
	return "users"
}
