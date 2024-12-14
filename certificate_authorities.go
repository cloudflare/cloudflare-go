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

type HostnameAssociationsResponse struct {
	Response
	Result []HostnameAssociation `json:"result"`
}

type HostnameAssociation = string

// List Hostname Associations
//
// API Reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-list-hostname-associations
func (api *API) ListCertificateAuthoritiesHostnameAssociations(ctx context.Context, rc *ResourceContainer, params ListCertificateAuthoritiesHostnameAssociationsParams) ([]HostnameAssociation, error) {

	uri := fmt.Sprintf(
		"/%s/%s/certificate_authorities/hostname_associations",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var hostnameAssociationsResponse HostnameAssociationsResponse
	err = json.Unmarshal(res, &hostnameAssociationsResponse)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return hostnameAssociationsResponse.Result, nil
}

// Replace Hostname Associations
//
// API Reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-put-hostname-associations
func (api *API) UpdateCertificateAuthoritiesHostnameAssociations(ctx context.Context, rc *ResourceContainer, params UpdateCertificateAuthoritiesHostnameAssociationsParams) ([]HostnameAssociation, error) {
	uri := fmt.Sprintf(
		"/%s/%s/certificate_authorities/hostname_associations",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var hostnameAssociationsResponse HostnameAssociationsResponse
	err = json.Unmarshal(res, &hostnameAssociationsResponse)
	if err != nil {
		return []HostnameAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return hostnameAssociationsResponse.Result, nil
}
