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

type ListHostnameAssociationParams struct {
	CertificateID string `url:"mtls_certificate_id,omitempty"`
}

type ReplaceHostnameAssociationParams struct {
	CertificateID string   `json:"mtls_certificate_id,omitempty"`
	Hostnames     []string `json:"hostnames"`
}
type ClientCertificateAuthority struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientCertificate struct {
	Certificate          string                     `json:"certificate"`
	CertificateAuthority ClientCertificateAuthority `json:"certificate_authority"`
	CommonName           string                     `json:"common_name"`
	Country              string                     `json:"country"`
	CSR                  string                     `json:"csr"`
	ExpiresOn            string                     `json:"expires_on"`
	FingerprintSHA256    string                     `json:"fingerprint_sha256"`
	ID                   string                     `json:"id"`
	IssuedOn             string                     `json:"issued_on"`
	Location             string                     `json:"location"`
	Organization         string                     `json:"organization"`
	OrganizationalUnit   string                     `json:"organizational_unit"`
	SerialNumber         string                     `json:"serial_number"`
	Signature            string                     `json:"signature"`
	SKI                  string                     `json:"ski"`
	State                string                     `json:"state"`
	Status               string                     `json:"status"`
	ValidityDays         int                        `json:"validity_days"`
}

type ListClientCertificatesResponse struct {
	Response
	Result []ClientCertificate `json:"result"`
	ResultInfo
}

type ListClientCertificatesParams struct {
	ResultInfo
}

type CreateClientCertificateParams struct {
	CSR          string `json:"csr"`
	ValidityDays int    `json:"validity_days"`
}

// ListHostnameAssociations returns a list of all domain associations for a given certificate_id.
//
// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-list-hostname-associations
func (api *API) ListHostnameAssociations(ctx context.Context, rc *ResourceContainer, params ListHostnameAssociationParams) (HostnameAssociation, error) {
	if rc.Level != ZoneRouteLevel {
		return HostnameAssociation{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return HostnameAssociation{}, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/certificate_authorities/hostname_associations", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return HostnameAssociation{}, err
	}
	var r HostnameAssociationResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, err
}

// ReplaceHostnameAssociations associates a list of hostnames with a given certificate_id.
//
// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-put-hostname-associations
func (api *API) ReplaceHostnameAssociations(ctx context.Context, rc *ResourceContainer, params ReplaceHostnameAssociationParams) (HostnameAssociation, error) {
	if rc.Level != ZoneRouteLevel {
		return HostnameAssociation{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return HostnameAssociation{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/certificate_authorities/hostname_associations", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return HostnameAssociation{}, err
	}
	var r HostnameAssociationResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameAssociation{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, err
}

// ListClientCertificates lists all of your Zone's API Shield mTLS Client Certificates by Status and/or using Pagination
//
// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-list-client-certificates
func (api *API) ListClientCertificates(ctx context.Context, rc *ResourceContainer, params ListClientCertificatesParams) ([]ClientCertificate, *ResultInfo, error) {
	if rc.Level != ZoneRouteLevel {
		return []ClientCertificate{}, &ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []ClientCertificate{}, &ResultInfo{}, ErrMissingZoneID
	}

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

	var clientCertificates []ClientCertificate
	var r ListClientCertificatesResponse

	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/client_certificates", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []ClientCertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		if err := json.Unmarshal(res, &r); err != nil {
			return []ClientCertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}

		clientCertificates = append(clientCertificates, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return clientCertificates, &r.ResultInfo, nil
}

// CreateClientCertificate creates a new API Shield mTLS Client Certificate
//
// API reference: https://developers.cloudflare.com/api/operations/client-certificate-for-a-zone-create-client-certificate
func (api *API) CreateClientCertificate(ctx context.Context, rc *ResourceContainer, params CreateClientCertificateParams) (ClientCertificate, error) {
	if rc.Level != ZoneRouteLevel {
		return ClientCertificate{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return ClientCertificate{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/client_certificates", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return ClientCertificate{}, err
	}
	var r ClientCertificate
	if err := json.Unmarshal(res, &r); err != nil {
		return ClientCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, err
}
