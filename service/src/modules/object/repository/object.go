package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/lib/pq"
)

func (r *Repository) CreateObject(ctx context.Context, obj request.Object) (request.Object, error) {
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return request.Object{}, err
	}

	query := `
        INSERT INTO "object" ("name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        RETURNING "id", "created_at", "updated_at"
    `
	err = r.db.Slave.QueryRowContext(ctx,
		query,
		obj.Name, obj.Address, obj.Description, pq.Array(obj.Banner), obj.Logo, socialMediaJSON, obj.Organizer, obj.Status, time.Now(), time.Now(),
	).Scan(&obj.ID, &obj.CreatedAt, &obj.UpdatedAt)

	if err != nil {
		return request.Object{}, err
	}

	return obj, nil
}
