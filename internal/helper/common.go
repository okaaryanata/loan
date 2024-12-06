package helper

import (
	"time"

	"github.com/jackc/pgx/v5"
)

func StructScan(row pgx.Row, dest interface{}) error {
	return row.Scan(dest)
}

func StructScanAll(rows pgx.Rows, dest interface{}) error {
	return rows.Scan(dest)
}

func JakartaTime() (*time.Time, error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	jakartaTime := time.Now().In(location)
	return &jakartaTime, nil
}
