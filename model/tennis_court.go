package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TennisCourt struct {
	ID			 			uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	name					string
	address					string
	lat						float64
	lng						float64
	CreatedAt   	 		time.Time
	UpdatedAt    			time.Time
	DeletedAt    			sql.NullTime
}