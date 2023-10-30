package model

// tennisMatchType
type tennisMatchType string

const (
	SINGLES tennisMatchType = "SINGLES"
	DOUBLES tennisMatchType = "DOUBLES"
)

func (mt *tennisMatchType) Scan(value interface{}) error {
	*mt = tennisMatchType(value.([]byte))
	return nil
}

func (mt tennisMatchType) Value() (interface{}, error) {
	return string(mt), nil
}

// tennisMatchSexType
type tennisMatchSexType string

const (
	MENS	tennisMatchSexType = "MENS"
	WOMENS	tennisMatchSexType = "WOMENS"
	MIXED	tennisMatchSexType = "MIXED"
)

func (mt *tennisMatchSexType) Scan(value interface{}) error {
	*mt = tennisMatchSexType(value.([]byte))
	return nil
}

func (mt tennisMatchSexType) Value() (interface{}, error) {
	return string(mt), nil
}

// MatchStatus
type MatchStatus string

const (
	REQUESTED		MatchStatus = "REQUESTED"
	ACCEPTED		MatchStatus = "ACCEPTED"
	REFUSED			MatchStatus = "REFUSED"
	IN_PROGRESS		MatchStatus = "IN_PROGRESS"
	COMPLETED		MatchStatus = "COMPLETED"
	CANCELLED		MatchStatus = "CANCELLED"
)

func (ms *MatchStatus) Scan(value interface{}) error {
	*ms = MatchStatus(value.([]byte))
	return nil
}

func (ms MatchStatus) Value() (interface{}, error) {
	return string(ms), nil
}