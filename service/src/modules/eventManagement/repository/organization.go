package repository

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (r *Repository) CreateOrganization(ctx context.Context, org entity.Organization) (uint64, error) {
	var id uint64
	err := r.db.Slave.QueryRowContext(ctx, "INSERT INTO organization (name, email, address, phone_number, logo) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		org.Name, org.Email, org.Address, org.PhoneNumber, org.Logo).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
