package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Player struct {
	UserId		uuid.UUID `gorm:"type:uuid;primary_key;"`
	User 	  	User
	Rating		uint
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	sql.NullTime
}