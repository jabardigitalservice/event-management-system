package endpoint

import (
	"context"

	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/request"
	"github.com/jabardigitalservice/super-app-services/event/src/modules/organization/transport/handler/http/response"
	"github.com/jinzhu/copier"
)

func (e *Endpoint) GetOrganizations(ctx context.Context, params request.QueryParam) ([]response.Organization, int, error) {
	organizations, count, err := e.usecase.GetOrganizations(ctx, params)
	if err != nil {
		return nil, 0, err
	}

	responseOrganizations := make([]response.Organization, len(organizations))

	for i, org := range organizations {
		responseOrg := &response.Organization{}
		if err := copier.Copy(responseOrg, org); err != nil {
			return nil, 0, err
		}
		responseOrganizations[i] = *responseOrg
	}

	return responseOrganizations, count, nil
}
