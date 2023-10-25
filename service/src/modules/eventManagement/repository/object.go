package repository

import (
	"context"
	"strings"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
	"github.com/lib/pq"
)

func (r *Repository) CreateObject(ctx context.Context, obj entity.Object) (uint64, time.Time, time.Time, error) {
	var id uint64
	var createdAt, updatedAt time.Time

	// Create a slice of []interface{} to build the parameters for the query
	// Ensure that the order of parameters matches the order in the SQL query
	params := []interface{}{obj.Name, obj.Address, obj.Description, pq.Array(obj.Banner), obj.Logo}
	var socialMediaArray pq.StringArray
	for _, media := range obj.SocialMedia {
		// Append each platform and link as a sub-array
		socialMediaArray = append(socialMediaArray, media[0]+","+media[1])
	}
	params = append(params, pq.Array(socialMediaArray), obj.Organizer, obj.Status, time.Now(), time.Now())

	// Insert the object into the database.
	err := r.db.Slave.QueryRowContext(ctx, `INSERT INTO "object" ("name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING "id", "created_at", "updated_at"`,
		params...).Scan(&id, &createdAt, &updatedAt)

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
		var socialMediaArray pq.StringArray

		if err := rows.Scan(&obj.ID, &obj.Name, &obj.Address, &obj.Description, &banners, &obj.Logo, &socialMediaArray, &obj.Organizer, &obj.Status, &obj.CreatedAt, &obj.UpdatedAt); err != nil {
			return nil, err
		}

		// Convert the pq.StringArray to a regular []string for banners
		obj.Banner = []string(banners)

		// Convert the pq.StringArray of arrays to the socialMedia field
		for _, sm := range socialMediaArray {
			// Split each sub-array on the comma to separate the platform and link
			parts := strings.Split(sm, ",")
			if len(parts) == 2 {
				obj.SocialMedia = append(obj.SocialMedia, []string{parts[0], parts[1]})
			}
		}

		objects = append(objects, obj)
	}

	return objects, nil
}
