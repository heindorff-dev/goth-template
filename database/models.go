// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"
)

type ApiErrorLog struct {
	ErrorID      int64
	Timestamp    time.Time
	ErrorMessage string
}
