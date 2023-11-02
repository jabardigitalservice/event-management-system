package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/object/transport/handler/http/request"
	"github.com/lib/pq"
)

func (r *Repository) CreateObject(ctx context.Context, obj request.Object) (request.Object, error) {
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return request.Object{}, err
	}

	query := `
        INSERT INTO "objects" ("name", "address", "description", "banner", "logo", "social_media", "organizer", "status", "created_at", "updated_at")
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

func (r *Repository) GetObjects(ctx context.Context, params request.QueryParam) ([]entity.Object, error) {
	binds := make([]interface{}, 0)

	var query = fmt.Sprintf(`SELECT
        id,
        name,
        address,
        description,
        banner,
        logo,
        social_media,
        organizer,
        status,
        created_at,
        updated_at
    FROM objects
    WHERE 1 = 1 %s `,
		r.filterObjectQuery(params, &binds))

	result, err := r.getObjects(ctx, query, binds...)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (r *Repository) getObjects(ctx context.Context, query string, args ...interface{}) ([]entity.Object, error) {
	rows, err := r.db.Slave.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, _errors.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	result := make([]entity.Object, 0)

	for rows.Next() {
		object := new(entity.Object)

		var banner pq.StringArray
		var socialMediaJSON []byte

		if err := rows.Scan(
			&object.ID,
			&object.Name,
			&object.Address,
			&object.Description,
			&banner,
			&object.Logo,
			&socialMediaJSON,
			&object.Organizer,
			&object.Status,
			&object.CreatedAt,
			&object.UpdatedAt,
		); err != nil {
			return nil, err
		}

		object.Banner = []string(banner)

		if err := json.Unmarshal(socialMediaJSON, &object.SocialMedia); err != nil {
			return nil, err
		}

		result = append(result, *object)
	}

	return result, nil
}

func (r *Repository) filterObjectQuery(params request.QueryParam, binds *[]interface{}) string {
	var query string
	counter := 1

	if params.Status != "" {
		*binds = append(*binds, params.Status)
		query = fmt.Sprintf(`%s AND status = $%d`, query, counter)
		counter++
	}
	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query = fmt.Sprintf(`%s AND (name ILIKE $%d OR description ILIKE $%d)`, query, counter, counter)
		counter++
	}

	if params.StartDate != "" && params.EndDate != "" {
		*binds = append(*binds, params.StartDate, params.EndDate)
		query = fmt.Sprintf("%s AND (DATE(created_at) BETWEEN $%d AND $%d)", query, counter, counter+1)
		counter += 2
	}

	sortBy := ` ORDER BY created_at DESC`
	if params.SortBy != "" {
		sortBy = fmt.Sprintf(` ORDER BY %s %s`, params.SortBy, params.SortOrder)
	}

	query += sortBy

	if params.PageSize > 0 {
		*binds = append(*binds, params.PageSize)
		query = fmt.Sprintf(`%s LIMIT $%d`, query, counter)
		counter++
	}

	if params.Offset > 0 {
		*binds = append(*binds, params.Offset)
		query = fmt.Sprintf(`%s OFFSET $%d`, query, counter)
	}

	return query
}

func (r *Repository) CountFilteredObjects(ctx context.Context, params request.QueryParam) (*int, error) {
	binds := make([]interface{}, 0)

	query := fmt.Sprintf(`SELECT COUNT(1) FROM objects WHERE 1 = 1 %s`, r.filterObjectCountQuery(params, &binds))

	count, err := r.countfilteredObjects(ctx, query, binds...)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func (r *Repository) countfilteredObjects(ctx context.Context, query string, args ...interface{}) (int, error) {
	var count int
	err := r.db.Slave.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) filterObjectCountQuery(params request.QueryParam, binds *[]interface{}) string {
	var query string
	counter := 1

	if params.Status != "" {
		*binds = append(*binds, params.Status)
		query = fmt.Sprintf(`%s AND status = $%d`, query, counter)
		counter++
	}
	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query = fmt.Sprintf(`%s AND (name ILIKE $%d OR description ILIKE $%d)`, query, counter, counter)
		counter++
	}

	if params.StartDate != "" && params.EndDate != "" {
		*binds = append(*binds, params.StartDate, params.EndDate)
		query = fmt.Sprintf("%s AND (DATE(created_at) BETWEEN $%d AND $%d)", query, counter, counter+1)
		counter += 2
	}

	query += ";"

	return query
}

func (r *Repository) GetObjectByID(ctx context.Context, id *uuid.UUID) (*entity.Object, error) {
	query := `
        SELECT
            id,
            name,
            address,
            description,
            banner,
            logo,
            social_media,
            organizer,
            status,
            created_at,
            updated_at
        FROM objects
        WHERE id = $1`

	var result entity.Object
	var banner pq.StringArray
	var socialMediaJSON []byte

	err := r.db.Slave.QueryRowContext(ctx, query, id).Scan(
		&result.ID,
		&result.Name,
		&result.Address,
		&result.Description,
		&banner,
		&result.Logo,
		&socialMediaJSON,
		&result.Organizer,
		&result.Status,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, _errors.ErrNotFound
		}
		return nil, err
	}
	if err := json.Unmarshal(socialMediaJSON, &result.SocialMedia); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) UpdateObject(ctx context.Context, obj *request.Object) (*request.Object, error) {
	socialMediaJSON, err := json.Marshal(obj.SocialMedia)
	if err != nil {
		return nil, err
	}
	bannerArray := pq.StringArray(obj.Banner)

	query := `
        UPDATE "objects" SET "name" = $2, "address" = $3, "description" = $4, "banner" = $5, "logo" = $6, "social_media" = $7, "organizer" = $8, "status" = $9, "updated_at" = $10
        WHERE "id" = $1
    `

	_, err = r.db.Master.ExecContext(ctx,
		query,
		obj.ID, obj.Name, obj.Address, obj.Description, bannerArray, obj.Logo, socialMediaJSON, obj.Organizer, obj.Status, time.Now(),
	)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *Repository) DeleteObject(ctx context.Context, id *uuid.UUID) error {
	query := "DELETE FROM objects WHERE id = $1"

	_, err := r.db.Master.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
