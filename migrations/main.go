package migrations

import (
	"context"
	"expense-tracker/models"
	"github.com/uptrace/bun"
	"log"
)

func RunMigrations(db *bun.DB) {
	ctx := context.Background()

	if _, err := db.NewCreateTable().Model((*models.Source)(nil)).IfNotExists().Exec(ctx); err != nil {
		log.Fatalf("Failed to run migrations for sources, %v", err)
	}

	if _, err := db.NewCreateTable().Model((*models.Transaction)(nil)).
		ForeignKey(`("source_id") REFERENCES "sources"("id") ON DELETE CASCADE`).
		IfNotExists().Exec(ctx); err != nil {
		log.Fatalf("Failed to run migrations for transaction, %v", err)
	}

	log.Println("Migrations completed successfully!")
}
