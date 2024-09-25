package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type InfraProtocol string

const (
	InfraSSH InfraProtocol = "SSH"
	RDP      InfraProtocol = "RDP"
)

// WARNING : The following table has a constraint. Each organization can specify at maximum one protocol for each of their ports.
// There can be at maximum a single (OrgID, Port, Protocol) combination, which can be shared across multiple records.
type InfraTargetContext struct {
	TargetAttributes map[string][]string `json:"target_attributes"`
	Port             int                 `json:"port"`
	Protocol         InfraProtocol       `json:"protocol"`
}

type InfrastructureApplication struct {
	// Items common to both SAML and OIDC
	ID             string                       `json:"id,omitempty"`
	Type           AccessApplicationType        `json:"type"`
	Name           string                       `json:"name"`
	Aud            string                       `json:"aud,omitempty"`
	CreatedAt      string                       `json:"created_at,omitempty"`
	UpdatedAt      string                       `json:"updated_at,omitempty"`
	ScimConfig     *AccessApplicationSCIMConfig `json:"scim_config,omitempty"`
	TargetContexts []InfraTargetContext         `json:"target_criteria"`
	Policies       []AccessPolicy               `json:"policies"`
}

type InfrastructureApplicationParams struct {
	Name           string                `json:"name"`
	Type           AccessApplicationType `json:"type"`
	TargetContexts []InfraTargetContext  `json:"target_criteria"`
	Policies       []AccessPolicy        `json:"policies"`
}

type CreateInfrastructureApplicationParams struct {
	InfrastructureApplicationParams
}

type UpdateInfrastructureApplicationParams struct {
	ID string `json:"-"`
	InfrastructureApplicationParams
}

// ListAccessInfrastructureApplications returns all applications within an account or zone.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-applications-list-access-applications
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-applications-list-access-applications
func (api *API) ListAccessInfrastructureApplications(ctx context.Context, rc *ResourceContainer, params ListAccessApplicationsParams) ([]InfrastructureApplication, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []InfrastructureApplication{}, &ResultInfo{}, ErrMissingAccountID
	}
	baseURL := fmt.Sprintf("/%s/%s/access/apps?app_type=infrastructure", rc.Level, rc.Identifier)

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = 25
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var infraApps []InfrastructureApplication
	var r AccessApplicationListResponse

	for {
		uri := buildURI(baseURL, params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []InfrastructureApplication{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []InfrastructureApplication{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		applications := r.Result
		for _, app := range applications {
			infraApps = append(infraApps, InfrastructureApplication{
				ID:             app.ID,
				Type:           app.Type,
				Name:           app.Name,
				Aud:            app.AUD,
				CreatedAt:      app.CreatedAt.String(),
				UpdatedAt:      app.UpdatedAt.String(),
				ScimConfig:     app.SCIMConfig,
				TargetContexts: app.TargetContexts,
				Policies:       app.Policies,
			})
		}
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return infraApps, &r.ResultInfo, nil
}

// CreateAccessInfrastructureApplication creates a new access application.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-applications-add-an-application
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-applications-add-a-bookmark-application
func (api *API) CreateAccessInfrastructureApplication(ctx context.Context, rc *ResourceContainer, params CreateInfrastructureApplicationParams) (InfrastructureApplication, error) {
	if rc.Identifier == "" {
		return InfrastructureApplication{}, ErrMissingAccountID
	}
	uri := fmt.Sprintf("/%s/%s/access/apps", rc.Level, rc.Identifier)
	return api.sendInfrastructureApplicationRequest(ctx, http.MethodPost, uri, params.InfrastructureApplicationParams)
}

// UpdateAccessInfrastructureApplication creates a new access application.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-applications-add-an-application
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-applications-add-a-bookmark-application
func (api *API) UpdateAccessInfrastructureApplication(ctx context.Context, rc *ResourceContainer, params UpdateInfrastructureApplicationParams) (InfrastructureApplication, error) {
	if rc.Identifier == "" {
		return InfrastructureApplication{}, ErrMissingAccountID
	}
	if params.ID == "" {
		return InfrastructureApplication{}, ErrMissingTargetId
	}
	uri := fmt.Sprintf("/%s/%s/access/apps/%s", rc.Level, rc.Identifier, params.ID)
	return api.sendInfrastructureApplicationRequest(ctx, http.MethodPut, uri, params.InfrastructureApplicationParams)
}

func (api *API) sendInfrastructureApplicationRequest(ctx context.Context, httpMethod string, uri string, params InfrastructureApplicationParams) (InfrastructureApplication, error) {
	res, err := api.makeRequestContext(ctx, httpMethod, uri, params)
	if err != nil {
		return InfrastructureApplication{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessApplicationDetailResponse AccessApplicationDetailResponse
	err = json.Unmarshal(res, &accessApplicationDetailResponse)
	if err != nil {
		return InfrastructureApplication{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	accessApp := accessApplicationDetailResponse.Result
	return InfrastructureApplication{
		ID:             accessApp.ID,
		Type:           accessApp.Type,
		Name:           accessApp.Name,
		Aud:            accessApp.AUD,
		CreatedAt:      accessApp.CreatedAt.String(),
		UpdatedAt:      accessApp.UpdatedAt.String(),
		ScimConfig:     accessApp.SCIMConfig,
		TargetContexts: accessApp.TargetContexts,
		Policies:       accessApp.Policies,
	}, nil
}

// GetAccessInfrastructureApplication returns a single infrastructure application based on the application ID for either account or zone.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-applications-get-an-access-application
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-applications-get-an-access-application
func (api *API) GetAccessInfrastructureApplication(ctx context.Context, rc *ResourceContainer, applicationID string) (InfrastructureApplication, error) {
	if rc.Identifier == "" {
		return InfrastructureApplication{}, ErrMissingAccountID
	}
	if applicationID == "" {
		return InfrastructureApplication{}, ErrMissingTargetId
	}
	infraAccessApp, err := api.GetAccessApplication(ctx, rc, applicationID)
	if err != nil {
		return InfrastructureApplication{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return InfrastructureApplication{
		ID:             infraAccessApp.ID,
		Type:           infraAccessApp.Type,
		Name:           infraAccessApp.Name,
		Aud:            infraAccessApp.AUD,
		CreatedAt:      infraAccessApp.CreatedAt.String(),
		UpdatedAt:      infraAccessApp.UpdatedAt.String(),
		ScimConfig:     infraAccessApp.SCIMConfig,
		TargetContexts: infraAccessApp.TargetContexts,
		Policies:       infraAccessApp.Policies,
	}, nil
}

// DeleteAccessInfrastructureApplication deletes a single infrastructure application.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-applications-delete-an-access-application
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-applications-delete-an-access-application
func (api *API) DeleteAccessInfrastructureApplication(ctx context.Context, rc *ResourceContainer, applicationID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}
	if applicationID == "" {
		return ErrMissingTargetId
	}
	return api.DeleteAccessApplication(ctx, rc, applicationID)
}
