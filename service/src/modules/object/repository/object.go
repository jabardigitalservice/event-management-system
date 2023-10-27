package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/lib/pq"
)

func (r *Repository) CreateObject(ctx context.Context, obj entity.Object) (entity.Object, error) {
	// Convert the SocialMedia slice to a format compatible with PostgreSQL JSONB array.
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return entity.Object{}, err
	}

	// Define a query that inserts the data and returns the id, created_at, and updated_at.
	query := `
		INSERT INTO "object" ("name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING "id", "created_at", "updated_at"
	`

	// Prepare the statement.
	stmt, err := r.db.Slave.PrepareContext(ctx, query)
	if err != nil {
		return entity.Object{}, err
	}
	defer stmt.Close()

	// Execute the prepared statement and scan the result into the obj.
	err = stmt.QueryRowContext(ctx,
		obj.Name, obj.Address, obj.Description, pq.Array(obj.Banner), obj.Logo, socialMediaJSON, obj.Organizer, obj.Status, time.Now(), time.Now()).
		Scan(&obj.ID, &obj.CreatedAt, &obj.UpdatedAt)

	if err != nil {
		return entity.Object{}, err
	}
	// Return the obj with the retrieved values.
	return obj, nil
}
