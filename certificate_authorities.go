package cloudflare

import "context"

type ListCertificateAuthoritiesHostnameAssociationsParams struct {
	MTLSCertificateID *string `url:"mtls_certificate_id,omitempty"`
}

type UpdateCertificateAuthoritiesHostnameAssociationsParams struct {
	Hostnames         []HostnameAssociation `json:"hostnames,omitempty"`
	MTLSCertificateID *string               `json:"mtls_certificate_id,omitempty"`
}

type HostnameAssociation = string

// List Hostname Associations
//
// API Reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-list-hostname-associations
func (api *API) ListCertificateAuthoritiesHostnameAssociations(ctx context.Context, rc *ResourceContainer, params ListCertificateAuthoritiesHostnameAssociationsParams) ([]string, error) {
	return []string{}, nil
}

// Replace Hostname Associations
//
// API Reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-put-hostname-associations
func (api *API) UpdateCertificateAuthoritiesHostnameAssociations(ctx context.Context, rc *ResourceContainer, params UpdateCertificateAuthoritiesHostnameAssociationsParams) ([]string, error) {
	return []string{}, nil
}
