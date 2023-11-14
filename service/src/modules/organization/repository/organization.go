package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_errors "github.com/jabardigitalservice/super-app-services/event/src/error"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
)

func (r *Repository) CreateOrganization(ctx context.Context, obj entity.Organization) (*entity.Organization, error) {
	insertQuery := `
        INSERT INTO organizations (name, email, address, pic_phone, description, logo, created_at, updated_at,
        province, city, district, village, google_map, pic_name, pic_position, province_id, city_id, district_id, village_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
    `

	_, err := r.db.Slave.ExecContext(ctx, insertQuery,
		obj.Name, obj.Email, obj.Address, obj.PicPhone, obj.Description, obj.Logo, time.Now(), time.Now(),
		obj.Province, obj.City, obj.District, obj.Village, obj.GoogleMap, obj.PicName, obj.PicPosition, obj.ProvinceID, obj.CityID, obj.DistrictID, obj.VillageID)

	if err != nil {
		return nil, err
	}

	selectQuery := `
        SELECT id, created_at, updated_at
        FROM organizations
        WHERE name = $1 AND email = $2
    `

	err = r.db.Slave.QueryRowContext(ctx, selectQuery, obj.Name, obj.Email).Scan(&obj.Id, &obj.CreatedAt, &obj.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *Repository) GetOrganizations(ctx context.Context, params request.QueryParam) ([]entity.Organization, error) {
	binds := make([]interface{}, 0)

	storageURL := r.app.GetStorageBaseUrl()
	storageURL = storageURL + "/"

	var query = fmt.Sprintf(`SELECT
        id,
        name,
        email,
        address,
		pic_phone,
        description,
        logo,
        created_at,
        updated_at,
        province,
        city,
        district,
        village,
        google_map,
        pic_name,
        pic_position,
		province_id, 
		city_id, 
		district_id, 
		village_id
    FROM organizations
    WHERE 1 = 1 %s `,
		r.filterOrganizationQuery(params, &binds))

	result, err := r.getOrganizations(ctx, query, binds...)
	if err != nil {
		return nil, err
	}

	for i := range result {
		result[i].Logo = storageURL + result[i].Logo
	}

	return result, nil
}

func (r *Repository) getOrganizations(ctx context.Context, query string, args ...interface{}) ([]entity.Organization, error) {
	rows, err := r.db.Slave.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, _errors.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	result := make([]entity.Organization, 0)

	for rows.Next() {
		organization := new(entity.Organization)

		if err := rows.Scan(
			&organization.Id,
			&organization.Name,
			&organization.Email,
			&organization.Address,
			&organization.PicPhone,
			&organization.Description,
			&organization.Logo,
			&organization.CreatedAt,
			&organization.UpdatedAt,
			&organization.Province,
			&organization.City,
			&organization.District,
			&organization.Village,
			&organization.GoogleMap,
			&organization.PicName,
			&organization.PicPosition,
			&organization.ProvinceID,
			&organization.CityID,
			&organization.DistrictID,
			&organization.VillageID,
		); err != nil {
			return nil, err
		}

		result = append(result, *organization)
	}

	return result, nil
}

func (r *Repository) filterOrganizationQuery(params request.QueryParam, binds *[]interface{}) string {
	var query string
	counter := 1

	if params.Status != "" {
		*binds = append(*binds, params.Status)
		query = fmt.Sprintf(`%s AND status = $%d`, query, counter)
		counter++
	}
	if params.Keyword != "" {
		*binds = append(*binds, "%"+params.Keyword+"%")
		query = fmt.Sprintf(`%s AND (name ILIKE $%d OR description ILIKE $%d)`, query, counter, counter)
		counter++
	}

	sortBy := " ORDER BY created_at DESC"
	if params.SortBy != "" {
		sortBy = fmt.Sprintf(" ORDER BY %s %s", params.SortBy, params.SortOrder)
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

func (r *Repository) CountFilteredOrganizations(ctx context.Context, params request.QueryParam) (int, error) {
	binds := make([]interface{}, 0)

	query := fmt.Sprintf(`SELECT COUNT(1) FROM organizations WHERE 1 = 1 %s`, r.filterOrganizationCountQuery(params, &binds))

	count, err := r.countFilteredOrganizations(ctx, query, binds...)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) countFilteredOrganizations(ctx context.Context, query string, args ...interface{}) (int, error) {
	var count int
	err := r.db.Slave.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) filterOrganizationCountQuery(params request.QueryParam, binds *[]interface{}) string {
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

	return query
}

func (r *Repository) GetOrganizationByID(ctx context.Context, id *uuid.UUID) (*entity.Organization, error) {
	query := `
        SELECT
        id,
        name,
        email,
        address,
		pic_phone,
        description,
        logo,
        created_at,
        updated_at,
        province,
        city,
        district,
        village,
        google_map,
        pic_name,
        pic_position,
		province_id, 
		city_id, 
		district_id, 
		village_id
    FROM organizations
        WHERE id = $1`

	storageURL := r.app.GetStorageBaseUrl()
	storageURL = storageURL + "/"

	var result entity.Organization

	err := r.db.Slave.QueryRowContext(ctx, query, id).Scan(
		&result.Id,
		&result.Name,
		&result.Email,
		&result.Address,
		&result.PicPhone,
		&result.Description,
		&result.Logo,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Province,
		&result.City,
		&result.District,
		&result.Village,
		&result.GoogleMap,
		&result.PicName,
		&result.PicPosition,
		&result.ProvinceID,
		&result.CityID,
		&result.DistrictID,
		&result.VillageID,
	)

	result.Logo = storageURL + result.Logo

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, _errors.ErrNotFound
		}
		return nil, err
	}

	return &result, nil
}

func (r *Repository) UpdateOrganization(ctx context.Context, obj *request.Organization) (*request.Organization, error) {
	query := `
        UPDATE organizations
        SET name = $2, email = $3, address = $4, pic_phone = $5, description = $6, logo = $7,
        province = $8, city = $9, district = $10, village = $11, google_map = $12, pic_name = $13, pic_position = $14, province_id = $15, 
		city_id = $16, 	district_id = $17, 	village_id = $18, updated_at = $19
        WHERE id = $1
    `

	_, err := r.db.Master.ExecContext(ctx, query, obj.Id, obj.Name, obj.Email, obj.Address, obj.PicPhone, obj.Description, obj.Logo,
		obj.Province, obj.City, obj.District, obj.Village, obj.GoogleMap, obj.PicName, obj.PicPosition, obj.ProvinceID, obj.CityID, obj.DistrictID, obj.VillageID, time.Now())

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (r *Repository) DeleteOrganization(ctx context.Context, id *uuid.UUID) error {
	query := "DELETE FROM organizations WHERE id = $1"

	_, err := r.db.Master.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
