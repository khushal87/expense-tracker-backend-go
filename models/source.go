package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Source struct {
	bun.BaseModel `bun:"table:sources"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	Name          string    `bun:"name,notnull" json:"name"`
	Type          string    `bun:"type,notnull" json:"type"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"UpdatedAt"`
}
