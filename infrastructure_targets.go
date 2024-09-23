package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var ErrMissingTargetId = errors.New("required target id missing")

// Target represents an Infrastructure Target.
type Target struct {
	Hostname   string    `json:"hostname"`
	ID         string    `json:"id"`
	IP         IPInfo    `json:"ip"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type IPDetails struct {
	IPAddr           string `json:"ip_addr"`
	VirtualNetworkId string `json:"virtual_network_id"`
}

type IPInfo struct {
	IPV4 *IPDetails `json:"ipv4,omitempty"`
	IPV6 *IPDetails `json:"ipv6,omitempty"`
}

type InfrastructureTargetParams struct {
	Hostname string `json:"hostname"`
	IP       IPInfo `json:"ip"`
}

type CreateInfrastructureTargetParams struct {
	InfrastructureTargetParams
}

type UpdateInfrastructureTargetParams struct {
	ID           string                     `json:"hostname"`
	ModifyParams InfrastructureTargetParams `json:"modify_params"`
}

// TargetDetailResponse is the API response, containing a single target.
type TargetDetailResponse struct {
	Result Target `json:"result"`
	Response
}

type TargetListDetailResponse struct {
	Result []Target `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type TargetListParams struct {
	CreatedAfter     string `url:"created_after,omitempty"`
	Hostname         string `url:"hostname,omitempty"`
	HostnameContains string `url:"hostname_contains,omitempty"`
	IPV4             string `url:"ip_v4,omitempty"`
	IPV6             string `url:"ip_v6,omitempty"`
	ModifedAfter     string `url:"modified_after,omitempty"`
	VirtualNetworkId string `url:"virtual_network_id,omitempty"`

	ResultInfo
}

// ListInfrastructureTargets returns all applications within an account or zone.
//
// API reference: https://developers.cloudflare.com/api/operations/infra-targets-list
func (api *API) ListInfrastructureTargets(ctx context.Context, rc *ResourceContainer, params TargetListParams) ([]Target, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []Target{}, &ResultInfo{}, ErrMissingAccountID
	}

	baseURL := fmt.Sprintf("/%s/%s/infrastructure/targets", rc.Level, rc.Identifier)

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

	var applications []Target
	var r TargetListDetailResponse

	for {
		uri := buildURI(baseURL, params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []Target{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []Target{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		applications = append(applications, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return applications, &r.ResultInfo, nil
}

// CreateInfrastructureTarget creates a new infrastructure target.
//
// API reference: https://developers.cloudflare.com/api/operations/infra-targets-post
func (api *API) CreateInfrastructureTarget(ctx context.Context, rc *ResourceContainer, params CreateInfrastructureTargetParams) (Target, error) {
	if rc.Identifier == "" {
		return Target{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/%s/%s/infrastructure/targets", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return Target{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var targetDetailResponse TargetDetailResponse
	err = json.Unmarshal(res, &targetDetailResponse)
	if err != nil {
		return Target{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return targetDetailResponse.Result, nil
}

// UpdateInfrastructureTarget updates an existing infrastructure target.
//
// API reference: https://developers.cloudflare.com/api/operations/infra-targets-put
func (api *API) UpdateInfrastructureTarget(ctx context.Context, rc *ResourceContainer, params UpdateInfrastructureTargetParams) (Target, error) {
	if rc.Identifier == "" {
		return Target{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return Target{}, ErrMissingTargetId
	}

	uri := fmt.Sprintf(
		"/%s/%s/infrastructure/targets/%s",
		rc.Level,
		rc.Identifier,
		params.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.ModifyParams)
	if err != nil {
		return Target{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var targetDetailResponse TargetDetailResponse
	err = json.Unmarshal(res, &targetDetailResponse)
	if err != nil {
		return Target{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return targetDetailResponse.Result, nil
}

// GetInfrastructureTarget returns a single infrastructure target based on target ID
// ID for either account or zone.
//
// API reference: https://developers.cloudflare.com/api/operations/infra-targets-get
func (api *API) GetInfrastructureTarget(ctx context.Context, rc *ResourceContainer, targetID string) (Target, error) {
	if rc.Identifier == "" {
		return Target{}, ErrMissingAccountID
	}

	if targetID == "" {
		return Target{}, ErrMissingTargetId
	}

	uri := fmt.Sprintf(
		"/%s/%s/infrastructure/targets/%s",
		rc.Level,
		rc.Identifier,
		targetID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Target{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var targetDetailResponse TargetDetailResponse
	err = json.Unmarshal(res, &targetDetailResponse)
	if err != nil {
		return Target{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return targetDetailResponse.Result, nil
}

// DeleteInfrastructureTarget deletes an infrastructure target.
//
// API reference: https://developers.cloudflare.com/api/operations/infra-targets-delete
func (api *API) DeleteInfrastructureTarget(ctx context.Context, rc *ResourceContainer, targetID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if targetID == "" {
		return ErrMissingTargetId
	}

	uri := fmt.Sprintf(
		"/%s/%s/infrastructure/targets/%s",
		rc.Level,
		rc.Identifier,
		targetID,
	)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	return nil
}
