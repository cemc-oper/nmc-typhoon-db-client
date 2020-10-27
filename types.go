package nmc_typhoon_db_client

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type DateTime struct {
	time.Time
}

func (date DateTime) MarshalCSV() ([]byte, error) {
	return []byte(date.Format("200601021504")), nil
}

type NullDateTime struct {
	mysql.NullTime
}

func (d NullDateTime) MarshalCSV() ([]byte, error) {
	if d.Valid {
		return []byte(d.Time.Format("200601021504")), nil
	} else {
		return nil, nil
	}
}

type DataString struct {
	sql.NullString
}

func (d DataString) MarshalCSV() ([]byte, error) {
	return []byte(d.String), nil
}

type DataInt32 struct {
	sql.NullInt32
}

func (d DataInt32) MarshalCSV() ([]byte, error) {
	if d.Valid {
		return []byte(fmt.Sprintf("%d", d.Int32)), nil
	} else {
		return nil, nil
	}
}

type DataFloat64 struct {
	sql.NullFloat64
}

func (d *DataFloat64) MarshalCSV() ([]byte, error) {
	if d.Valid {
		return []byte(fmt.Sprintf("%f", d.Float64)), nil
	} else {
		return nil, nil
	}
}
