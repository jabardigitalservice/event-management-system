package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
	"github.com/lib/pq"
)

func (r *Repository) CreateObject(ctx context.Context, obj entity.Object) (uint64, time.Time, time.Time, error) {
	var id uint64
	var createdAt, updatedAt time.Time

	// Convert the SocialMedia slice to a format compatible with PostgreSQL JSONB array.
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return 0, time.Time{}, time.Time{}, err
	}
	fmt.Println("socialMediaJSON:", string(socialMediaJSON))

	err = r.db.Slave.QueryRowContext(ctx, `INSERT INTO "object" ("name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING "id", "created_at", "updated_at"`,
		obj.Name, obj.Address, obj.Description, pq.Array(obj.Banner), obj.Logo, socialMediaJSON, obj.Organizer, obj.Status, time.Now(), time.Now()).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return 0, time.Time{}, time.Time{}, err
	}

	return id, createdAt, updatedAt, nil
}

func (r *Repository) GetObjects(ctx context.Context, page int, perPage int) ([]entity.Object, error) {
	// Calculate the offset based on the page and perPage
	offset := (page - 1) * perPage

	// Query the database to retrieve a page of objects
	var objects []entity.Object
	query := `SELECT * FROM "object" ORDER BY "id" LIMIT $1 OFFSET $2`
	rows, err := r.db.Slave.QueryContext(ctx, query, perPage, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var obj entity.Object

		// Initialize the pq.StringArray for banners
		var banners pq.StringArray

		// Define a variable to scan the social_media data from the database
		var socialMediaData []byte

		if err := rows.Scan(&obj.ID, &obj.Name, &obj.Address, &obj.Description, &banners, &obj.Logo, &socialMediaData, &obj.Organizer, &obj.Status, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			return nil, err
		}

		// Unmarshal the JSONB data into the entity.SocialMedia slice
		var socialMedia []entity.SocialMedia
		if err := json.Unmarshal(socialMediaData, &socialMedia); err != nil {
			return nil, err
		}

		// Assign the unmarshaled socialMedia data to the object
		obj.SocialMedia = socialMedia

		objects = append(objects, obj)
	}

	return objects, nil
}
