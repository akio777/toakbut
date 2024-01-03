package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Breaks struct {
	bun.BaseModel `bun:"breaks,alias:breaks"`
	ID            int        `bun:"id,autoincrement,pk" json:"id"`
	UserID        string     `bun:"user_id" json:"user_id"`
	BreakIn       *time.Time `bun:"break_in" json:"break_in,omitempty"`
	BreakOut      *time.Time `bun:"break_out" json:"break_out,omitempty"`
	CreatedAt     *time.Time `bun:"created_at" json:"created_at,omitempty"`
	UpdatedAt     *time.Time `bun:"updated_at" json:"updated_at,omitempty"`
}
