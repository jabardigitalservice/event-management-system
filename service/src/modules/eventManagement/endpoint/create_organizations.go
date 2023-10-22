package endpoint

import (
	"context"
	"time"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/eventManagement/entity"
)

func (e *Endpoint) CreateOrganizations(ctx context.Context, orgData entity.Organization) (uint64, entity.Organization, error) {
	id, org, err := e.usecase.CreateOrganizations(ctx, orgData)
	if err != nil {
		return 0, entity.Organization{}, err
	}

	createdAt := time.Now()
	// Create an instance of the Organization struct with the data
	organization := entity.Organization{
		Id:          id,
		Name:        org.Name,
		Email:       org.Email,
		Address:     org.Address,
		PhoneNumber: org.PhoneNumber,
		Logo:        org.Logo,
		CreatedAt:   createdAt,
	}

	return id, organization, nil
}
