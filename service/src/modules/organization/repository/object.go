package repository

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
)

func (r *Repository) CreateOrganization(ctx context.Context, obj entity.Organization) (*entity.Organization, error) {
	insertQuery := `
        INSERT INTO organizations (name, email, address, phone_number, description, logo, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_, err := r.db.Slave.ExecContext(ctx, insertQuery,
		obj.Name, obj.Email, obj.Address, obj.Phone_number, obj.Description, obj.Logo, time.Now(), time.Now())

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
