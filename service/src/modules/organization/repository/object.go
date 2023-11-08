package repository

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/entity"
)

func (r *Repository) CreateOrganization(ctx context.Context, obj entity.Organization) (entity.Organization, error) {
	query := `
        INSERT INTO organizations (name, email, address, phone_number, description, logo, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_, err := r.db.Slave.ExecContext(ctx, query,
		obj.Name, obj.Email, obj.Address, obj.Phone_number, obj.Description, obj.Logo, time.Now(), time.Now())

	return obj, err
}
