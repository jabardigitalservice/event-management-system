package sql

import (
	"database/sql"
	"time"
)

func ParseNullString(d sql.NullString) *string {
	if d.Valid {
		return &d.String
	} else {
		return nil
	}
}

func ParseNullTime(d sql.NullTime) *time.Time {
	if d.Valid {
		return &d.Time
	} else {
		return nil
	}
}

func ParseNullInt32(d sql.NullInt32) *int32 {
	if d.Valid {
		return &d.Int32
	} else {
		return nil
	}
}
