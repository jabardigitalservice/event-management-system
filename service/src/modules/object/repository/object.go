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

func (r *Repository) GetObjects(ctx context.Context, page int, perPage int) ([]request.Object, error) {
	offset := (page - 1) * perPage

	var objects []request.Object
	query := `SELECT * FROM "object" ORDER BY "id" LIMIT $1 OFFSET $2`
	rows, err := r.db.Slave.QueryContext(ctx, query, perPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var obj request.Object
		var banners pq.StringArray
		var socialMediaData []byte

		if err := rows.Scan(&obj.ID, &obj.Name, &obj.Address, &obj.Description, &banners, &obj.Logo, &socialMediaData, &obj.Organizer, &obj.Status, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			return nil, err
		}

		obj.Banner = []string(banners)

		if err := json.Unmarshal(socialMediaData, &obj.SocialMedia); err != nil {
			return nil, err
		}

		objects = append(objects, obj)
	}

	return objects, nil
}
