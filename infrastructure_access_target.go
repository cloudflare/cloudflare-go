package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

var ErrMissingTargetId = errors.New("required target id missing")

type InfrastructureAccessTarget struct {
	Hostname   string                           `json:"hostname"`
	ID         string                           `json:"id"`
	IP         InfrastructureAccessTargetIPInfo `json:"ip"`
	CreatedAt  string                           `json:"created_at"`
	ModifiedAt string                           `json:"modified_at"`
}

type InfrastructureAccessTargetIPInfo struct {
	IPV4 *InfrastructureAccessTargetIPDetails `json:"ipv4,omitempty"`
	IPV6 *InfrastructureAccessTargetIPDetails `json:"ipv6,omitempty"`
}

type InfrastructureAccessTargetIPDetails struct {
	IPAddr           string `json:"ip_addr"`
	VirtualNetworkId string `json:"virtual_network_id"`
}

type InfrastructureAccessTargetParams struct {
	Hostname string                           `json:"hostname"`
	IP       InfrastructureAccessTargetIPInfo `json:"ip"`
}

type CreateInfrastructureAccessTargetParams struct {
	InfrastructureAccessTargetParams
}

type UpdateInfrastructureAccessTargetParams struct {
	ID           string                           `json:"-"`
	ModifyParams InfrastructureAccessTargetParams `json:"modify_params"`
}

// InfrastructureAccessTargetDetailResponse is the API response, containing a single target.
type InfrastructureAccessTargetDetailResponse struct {
	Result InfrastructureAccessTarget `json:"result"`
	Response
}

type InfrastructureAccessTargetListDetailResponse struct {
	Result []InfrastructureAccessTarget `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type InfrastructureAccessTargetListParams struct {
	Hostname         string `url:"hostname,omitempty"`
	HostnameContains string `url:"hostname_contains,omitempty"`
	IPV4             string `url:"ip_v4,omitempty"`
	IPV6             string `url:"ip_v6,omitempty"`
	CreatedAfter     string `url:"created_after,omitempty"`
	ModifedAfter     string `url:"modified_after,omitempty"`
	VirtualNetworkId string `url:"virtual_network_id,omitempty"`

	ResultInfo
}

// ListInfrastructureAccessTargets returns all infrastructure access targets within an account.
//
// Account API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/access/subresources/infrastructure/subresources/targets/methods/list/
func (api *API) ListInfrastructureAccessTargets(ctx context.Context, rc *ResourceContainer, params InfrastructureAccessTargetListParams) ([]InfrastructureAccessTarget, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []InfrastructureAccessTarget{}, &ResultInfo{}, ErrMissingAccountID
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

	var targets []InfrastructureAccessTarget
	var r InfrastructureAccessTargetListDetailResponse

	for {
		uri := buildURI(baseURL, params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []InfrastructureAccessTarget{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []InfrastructureAccessTarget{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		targets = append(targets, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return targets, &r.ResultInfo, nil
}

// CreateInfrastructureAccessTarget creates a new infrastructure access target.
//
// Account API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/access/subresources/infrastructure/subresources/targets/methods/create/
func (api *API) CreateInfrastructureAccessTarget(ctx context.Context, rc *ResourceContainer, params CreateInfrastructureAccessTargetParams) (InfrastructureAccessTarget, error) {
	if rc.Identifier == "" {
		return InfrastructureAccessTarget{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/%s/%s/infrastructure/targets", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return InfrastructureAccessTarget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var targetDetailResponse InfrastructureAccessTargetDetailResponse
	err = json.Unmarshal(res, &targetDetailResponse)
	if err != nil {
		return InfrastructureAccessTarget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return targetDetailResponse.Result, nil
}

// UpdateInfrastructureAccessTarget updates an existing infrastructure access target.
//
// Account API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/access/subresources/infrastructure/subresources/targets/methods/update/
func (api *API) UpdateInfrastructureAccessTarget(ctx context.Context, rc *ResourceContainer, params UpdateInfrastructureAccessTargetParams) (InfrastructureAccessTarget, error) {
	if rc.Identifier == "" {
		return InfrastructureAccessTarget{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return InfrastructureAccessTarget{}, ErrMissingTargetId
	}

	uri := fmt.Sprintf(
		"/%s/%s/infrastructure/targets/%s",
		rc.Level,
		rc.Identifier,
		params.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.ModifyParams)
	if err != nil {
		return InfrastructureAccessTarget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var targetDetailResponse InfrastructureAccessTargetDetailResponse
	err = json.Unmarshal(res, &targetDetailResponse)
	if err != nil {
		return InfrastructureAccessTarget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return targetDetailResponse.Result, nil
}

// GetInfrastructureAccessTarget returns a single infrastructure access target based on target ID
// ID for either account or zone.
//
// Account API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/access/subresources/infrastructure/subresources/targets/methods/get/
func (api *API) GetInfrastructureAccessTarget(ctx context.Context, rc *ResourceContainer, targetID string) (InfrastructureAccessTarget, error) {
	if rc.Identifier == "" {
		return InfrastructureAccessTarget{}, ErrMissingAccountID
	}

	if targetID == "" {
		return InfrastructureAccessTarget{}, ErrMissingTargetId
	}

	uri := fmt.Sprintf(
		"/%s/%s/infrastructure/targets/%s",
		rc.Level,
		rc.Identifier,
		targetID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return InfrastructureAccessTarget{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var targetDetailResponse InfrastructureAccessTargetDetailResponse
	err = json.Unmarshal(res, &targetDetailResponse)
	if err != nil {
		return InfrastructureAccessTarget{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return targetDetailResponse.Result, nil
}

// DeleteInfrastructureAccessTarget deletes an infrastructure access target.
//
// Account API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/access/subresources/infrastructure/subresources/targets/methods/delete/
func (api *API) DeleteInfrastructureAccessTarget(ctx context.Context, rc *ResourceContainer, targetID string) error {
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
