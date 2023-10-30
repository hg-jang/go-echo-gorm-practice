package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID			 			uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Email		 			string
	NickName				string
	FirstName				string
	LastName				string
	Password				string
	RefreshToken			sql.NullString
	RefreshTokenExpiration  sql.NullTime
	CreatedAt				time.Time
	UpdatedAt				time.Time
	DeletedAt				sql.NullTime
}

func (user *User) TableName() string {
	return "users"
}