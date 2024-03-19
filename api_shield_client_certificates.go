package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// HostnameAssociation represents the metadata for an existing association
// between a user-uploaded mTLS certificate and hostnames within a given zone.
type HostnameAssociation struct {
	CertificateID string   `json:"mtls_certificate_id,omitempty"`
	Hostnames     []string `json:"hostnames"`
}

type HostnameAssociationResponse struct {
	Response
	Result HostnameAssociation `json:"result"`
}

type HostnameAssociationParams struct {
	CertificateID string `json:"mtls_certificate_id,omitempty"`
}

// ListHostnameAssociations returns a list of all domain associations for a given certificate_id.
//
// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-list-hostname-associations
func (api *API) ListHostnameAssociations(ctx context.Context, rc *ResourceContainer, params HostnameAssociationParams) (HostnameAssociation, error) {
	if rc.Level != ZoneRouteLevel {
		return HostnameAssociation{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return HostnameAssociation{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/certificate_authorities/hostname_associations", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return HostnameAssociation{}, err
	}
	var r HostnameAssociationResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, err
}
