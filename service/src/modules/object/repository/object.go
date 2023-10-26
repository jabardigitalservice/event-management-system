package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/lib/pq"
)

func (r *Repository) CreateObject(ctx context.Context, obj entity.Object) (entity.Object, error) {
	var id uint64
	var createdAt, updatedAt time.Time

	// Convert the SocialMedia slice to a format compatible with PostgreSQL JSONB array.
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return entity.Object{}, err
	}

	err = r.db.Slave.QueryRowContext(ctx, `INSERT INTO "object" ("name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING "id", "created_at", "updated_at"`,
		obj.Name, obj.Address, obj.Description, pq.Array(obj.Banner), obj.Logo, socialMediaJSON, obj.Organizer, obj.Status, time.Now(), time.Now()).
		Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return entity.Object{}, err
	}

	// Assign the values to the original obj and return it
	obj.ID = id
	obj.CreatedAt = createdAt
	obj.UpdatedAt = updatedAt

	return obj, nil
}
