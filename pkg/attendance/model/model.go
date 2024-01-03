package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Attendance struct {
	bun.BaseModel `bun:"attendance,alias:attendance"`
	ID            int        `bun:"id,autoincrement,pk" json:"id"`
	UserID        string     `bun:"user_id" json:"user_id"`
	CheckIn       *time.Time `bun:"check_in" json:"check_in,omitempty"`
	CheckOut      *time.Time `bun:"check_out" json:"check_out,omitempty"`
	WorkType      string     `bun:"work_type" json:"work_type"`
	CreatedAt     *time.Time `bun:"created_at" json:"created_at,omitempty"`
	UpdatedAt     *time.Time `bun:"updated_at" json:"updated_at,omitempty"`
}
