package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	Amount        int64     `bun:"amount,notnull" json:"amount"`
	Received      bool      `bun:"received,notnull,default:false" json:"received"`
	SourceID      int64     `bun:"source_id,notnull" json:"sourceId"`              // Foreign key to source
	Source        *Source   `bun:"rel:belongs-to,join:source_id=id" json:"source"` // Relation to source
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"UpdatedAt"`
}
