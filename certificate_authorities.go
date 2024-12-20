package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type ListCertificateAuthoritiesHostnameAssociationsParams struct {
	MTLSCertificateID string `url:"mtls_certificate_id,omitempty"`
}

type UpdateCertificateAuthoritiesHostnameAssociationsParams struct {
	Hostnames         []HostnameAssociation `json:"hostnames,omitempty"`
	MTLSCertificateID string                `json:"mtls_certificate_id,omitempty"`
}

type HostnameAssociationsUpdateRequest struct {
	Hostnames         []HostnameAssociation `json:"hostnames,omitempty"`
	MTLSCertificateID string                `json:"mtls_certificate_id,omitempty"`
}

type HostnameAssociationsResponse struct {
	Response
	Result HostnameAssociations `json:"result"`
}

type HostnameAssociations struct {
	Hostnames []HostnameAssociation `json:"hostnames"`
}

type HostnameAssociation = string

// List Hostname Associations
//
// API Reference: https://developers.cloudflare.com/api/resources/certificate_authorities/subresources/hostname_associations/methods/get/
func (api *API) ListCertificateAuthoritiesHostnameAssociations(ctx context.Context, rc *ResourceContainer, params ListCertificateAuthoritiesHostnameAssociationsParams) ([]HostnameAssociation, error) {
	if rc.Level != ZoneRouteLevel {
		return []HostnameAssociation{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/certificate_authorities/hostname_associations", rc.Identifier), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var hostnameAssociationsResponse HostnameAssociationsResponse
	err = json.Unmarshal(res, &hostnameAssociationsResponse)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return hostnameAssociationsResponse.Result.Hostnames, nil
}

// Replace Hostname Associations
//
// API Reference: https://developers.cloudflare.com/api/resources/certificate_authorities/subresources/hostname_associations/methods/update/
func (api *API) UpdateCertificateAuthoritiesHostnameAssociations(ctx context.Context, rc *ResourceContainer, params UpdateCertificateAuthoritiesHostnameAssociationsParams) ([]HostnameAssociation, error) {
	if rc.Level != ZoneRouteLevel {
		return []HostnameAssociation{}, ErrRequiredZoneLevelResourceContainer
	}
	
	uri := fmt.Sprintf("/zones/%s/certificate_authorities/hostname_associations", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var hostnameAssociationsResponse HostnameAssociationsResponse
	err = json.Unmarshal(res, &hostnameAssociationsResponse)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return hostnameAssociationsResponse.Result.Hostnames, nil
}
