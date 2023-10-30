package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type TennisMatch struct {
	ID			 			uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	MatchType				tennisMatchType `gorm:"type:tennis_match_type"`
	SexType					tennisMatchSexType `gorm:"type:tennis_match_sex_type"`
	MatchStatus				MatchStatus `gorm:"type:match_status"`
	CourtId					uuid.UUID `gorm:"type:uuid;"`
	Court					TennisCourt
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	sql.NullTime
}