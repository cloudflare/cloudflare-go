package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AccessMutualTLSCertificate is the structure of a single Access Mutual TLS
// certificate.
type AccessMutualTLSCertificate struct {
	ID                  string    `json:"id,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	ExpiresOn           time.Time `json:"expires_on,omitempty"`
	Name                string    `json:"name,omitempty"`
	Fingerprint         string    `json:"fingerprint,omitempty"`
	Certificate         string    `json:"certificate,omitempty"`
	AssociatedHostnames []string  `json:"associated_hostnames,omitempty"`
}

// AccessMutualTLSCertificateListResponse is the API response for all Access
// Mutual TLS certificates.
type AccessMutualTLSCertificateListResponse struct {
	Response
	Result     []AccessMutualTLSCertificate `json:"result"`
	ResultInfo `json:"result_info"`
}

// AccessMutualTLSCertificateDetailResponse is the API response for a single
// Access Mutual TLS certificate.
type AccessMutualTLSCertificateDetailResponse struct {
	Response
	Result AccessMutualTLSCertificate `json:"result"`
}

// ListAccessMutualTLSCertificates returns all Access TLS certificates
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-list-mtls-certificates
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-list-mtls-certificates
func (api *API) ListAccessMutualTLSCertificates(ctx context.Context, rc *ResourceContainer, pageOpts PaginationOptions) ([]AccessMutualTLSCertificate, *ResultInfo, error) {
	baseURL := fmt.Sprintf(
		"/%s/%s/access/certificates",
		rc.Level,
		rc.Identifier,
	)

	autoPaginate := true
	if pageOpts.PerPage >= 1 || pageOpts.Page >= 1 {
		autoPaginate = false
	}

	if pageOpts.PerPage < 1 {
		pageOpts.PerPage = 25
	}

	if pageOpts.Page < 1 {
		pageOpts.Page = 1
	}

	resultInfo := ResultInfo{
		Page:    pageOpts.Page,
		PerPage: pageOpts.PerPage,
	}

	var accessCertificates []AccessMutualTLSCertificate
	var r AccessMutualTLSCertificateListResponse

	for {
		uri := buildURI(baseURL, resultInfo)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []AccessMutualTLSCertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []AccessMutualTLSCertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		accessCertificates = append(accessCertificates, r.Result...)
		resultInfo = r.ResultInfo.Next()
		if resultInfo.Done() || autoPaginate {
			break
		}
	}

	return accessCertificates, &r.ResultInfo, nil
}

// GetAccessMutualTLSCertificate returns a single Access Mutual TLS
// certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-get-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-get-an-mtls-certificate
func (api *API) GetAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		rc.Level,
		rc.Identifier,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// CreateAccessMutualTLSCertificate creates an Access TLS Mutual
// certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-add-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-add-an-mtls-certificate
func (api *API) CreateAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, certificate)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// UpdateAccessMutualTLSCertificate updates an account level Access TLS Mutual
// certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-update-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-update-an-mtls-certificate
func (api *API) UpdateAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		rc.Level,
		rc.Identifier,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, certificate)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// DeleteAccessMutualTLSCertificate destroys an Access Mutual
// TLS certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-delete-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-delete-an-mtls-certificate
func (api *API) DeleteAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		rc.Level,
		rc.Identifier,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}
