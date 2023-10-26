package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
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

		// Convert pq.StringArray to []string for banners
		obj.Banner = []string(banners)

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

func (r *Repository) GetObjectByID(ctx context.Context, objectID uint64) (entity.Object, error) {
	// Query the database to retrieve the object by ID
	var obj entity.Object
	query := `SELECT * FROM "object" WHERE "id" = $1`
	row := r.db.Slave.QueryRowContext(ctx, query, objectID)

	// Initialize the pq.StringArray for banners
	var banners pq.StringArray

	// Define a variable to scan the social_media data from the database
	var socialMediaData []byte

	err := row.Scan(&obj.ID, &obj.Name, &obj.Address, &obj.Description, &banners, &obj.Logo, &socialMediaData, &obj.Organizer, &obj.Status, &obj.CreatedAt, &obj.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Object{}, _errors.ErrNotFound
		}
		return entity.Object{}, err
	}

	// Convert pq.StringArray to []string for banners
	obj.Banner = []string(banners)

	// Unmarshal the JSONB data into the entity.SocialMedia slice
	var socialMedia []entity.SocialMedia
	if err := json.Unmarshal(socialMediaData, &socialMedia); err != nil {
		return entity.Object{}, err
	}

	// Assign the unmarshaled socialMedia data to the object
	obj.SocialMedia = socialMedia

	return obj, nil
}

func (r *Repository) UpdateObject(ctx context.Context, obj entity.Object) (entity.Object, error) {
	var updatedObject entity.Object

	// Convert the SocialMedia slice to a format compatible with PostgreSQL JSONB array.
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return entity.Object{}, err
	}

	// Create a slice of strings for the "banner" field.
	bannerArray := pq.StringArray(obj.Banner)

	// Perform the update operation in the database and return the updated data
	err = r.db.Master.QueryRowContext(ctx,
		`UPDATE "object" SET "name" = $2, "address" = $3, "description" = $4, "banner" = $5, "logo" = $6, "social_media" = $7, "organizer" = $8, "status" = $9, "updated_at" = $10
         WHERE "id" = $1 RETURNING "id", "name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at"`,
		obj.ID, obj.Name, obj.Address, obj.Description, bannerArray, obj.Logo, socialMediaJSON, obj.Organizer, obj.Status, time.Now(),
	).Scan(&updatedObject.ID, &updatedObject.Name, &updatedObject.Address, &updatedObject.Description, &bannerArray, &updatedObject.Logo, &socialMediaJSON, &updatedObject.Organizer, &updatedObject.Status, &updatedObject.CreatedAt, &updatedObject.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			// Object not found, return a custom "not found" error
			return entity.Object{}, _errors.ErrNotFound
		}
		return entity.Object{}, err
	}

	// Retrieve the "banner" values from the pq.StringArray.
	updatedObject.Banner = []string(bannerArray)

	// Unmarshal the "social_media" JSONB data into the entity.SocialMedia slice.
	if err := json.Unmarshal([]byte(socialMediaJSON), &updatedObject.SocialMedia); err != nil {
		return entity.Object{}, err
	}

	return updatedObject, nil
}

func (r *Repository) DeleteObject(ctx context.Context, objectID uint64) error {
	_, err := r.db.Master.ExecContext(ctx, `DELETE FROM "object" WHERE "id" = $1`, objectID)
	return err
}
