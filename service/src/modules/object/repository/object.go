package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/lib/pq"
)

func (r *Repository) CreateObject(ctx context.Context, obj request.Object) (request.Object, error) {
	// Convert the SocialMedia slice to a format compatible with PostgreSQL JSONB array.
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return request.Object{}, err
	}

	// Define a query that inserts the data and returns the id, created_at, and updated_at.
	query := `
        INSERT INTO "object" ("name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING "id", "created_at", "updated_at"
    `

	// Execute the query using ExecContext.
	_, err = r.db.Slave.ExecContext(ctx,
		query, obj.Name, obj.Address, obj.Description, pq.Array(obj.Banner), obj.Logo, socialMediaJSON, obj.Organizer, obj.Status, time.Now(), time.Now())

	if err != nil {
		return request.Object{}, err
	}

	// Retrieve the generated values by scanning the result into the obj.
	err = r.db.Slave.QueryRowContext(ctx, "SELECT * FROM object WHERE id = (SELECT last_value FROM object_id_seq);").Scan(
		&obj.ID, &obj.Name, &obj.Address, &obj.Description, &obj.Banner, &obj.Logo, &socialMediaJSON, &obj.Organizer, &obj.Status, &obj.CreatedAt, &obj.UpdatedAt)

	if err != nil {
		return request.Object{}, err
	}

	return obj, nil
}
