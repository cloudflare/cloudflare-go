package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
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
	Result []AccessMutualTLSCertificate `json:"result"`
}

// AccessMutualTLSCertificateDetailResponse is the API response for a single
// Access Mutual TLS certificate.
type AccessMutualTLSCertificateDetailResponse struct {
	Response
	Result AccessMutualTLSCertificate `json:"result"`
}

// AccessMutualTLSCertificates returns all Access TLS certificates for the account
// level.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-properties
func (api *API) AccessMutualTLSCertificates(accountID string) ([]AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificates(accountID, AccountRouteRoot)
}

// ZoneAccessMutualTLSCertificates returns all Access TLS certificates for the
// zone level.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-properties
func (api *API) ZoneAccessMutualTLSCertificates(zoneID string) ([]AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificates(zoneID, ZoneRouteRoot)
}

func (api *API) accessMutualTLSCertificates(id string, routeRoot RouteRoot) ([]AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates",
		routeRoot,
		id,
	)

	res, err := api.makeRequestContext(context.Background(), "GET", uri, nil)
	if err != nil {
		return []AccessMutualTLSCertificate{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessMutualTLSCertificateListResponse AccessMutualTLSCertificateListResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateListResponse)
	if err != nil {
		return []AccessMutualTLSCertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessMutualTLSCertificateListResponse.Result, nil
}

// AccessMutualTLSCertificate returns a single account level Access Mutual TLS
// certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-access-certificate-details
func (api *API) AccessMutualTLSCertificate(accountID, certificateID string) (AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificate(accountID, certificateID, AccountRouteRoot)
}

// ZoneAccessMutualTLSCertificate returns a single zone level Access Mutual TLS
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-access-certificate-details
func (api *API) ZoneAccessMutualTLSCertificate(zoneID, certificateID string) (AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificate(zoneID, certificateID, ZoneRouteRoot)
}

func (api *API) accessMutualTLSCertificate(id, certificateID string, routeRoot RouteRoot) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		routeRoot,
		id,
		certificateID,
	)

	res, err := api.makeRequestContext(context.Background(), "GET", uri, nil)
	if err != nil {
		return AccessMutualTLSCertificate{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// CreateAccessMutualTLSCertificate creates an account level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-create-access-certificate
func (api *API) CreateAccessMutualTLSCertificate(accountID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.createAccessMutualTLSCertificate(accountID, certificate, AccountRouteRoot)
}

// CreateZoneAccessMutualTLSCertificate creates a zone level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-create-access-certificate
func (api *API) CreateZoneAccessMutualTLSCertificate(zoneID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.createAccessMutualTLSCertificate(zoneID, certificate, ZoneRouteRoot)
}

func (api *API) createAccessMutualTLSCertificate(id string, certificate AccessMutualTLSCertificate, routeRoot RouteRoot) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates",
		routeRoot,
		id,
	)

	res, err := api.makeRequestContext(context.Background(), "POST", uri, certificate)
	if err != nil {
		return AccessMutualTLSCertificate{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// UpdateAccessMutualTLSCertificate updates an account level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-update-access-certificate
func (api *API) UpdateAccessMutualTLSCertificate(accountID, certificateID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.updateAccessMutualTLSCertificate(accountID, certificateID, certificate, AccountRouteRoot)
}

// UpdateZoneAccessMutualTLSCertificate updates a zone level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-update-access-certificate
func (api *API) UpdateZoneAccessMutualTLSCertificate(zoneID, certificateID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.updateAccessMutualTLSCertificate(zoneID, certificateID, certificate, ZoneRouteRoot)
}

func (api *API) updateAccessMutualTLSCertificate(id string, certificateID string, certificate AccessMutualTLSCertificate, routeRoot RouteRoot) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		routeRoot,
		id,
		certificateID,
	)

	res, err := api.makeRequestContext(context.Background(), "PUT", uri, certificate)
	if err != nil {
		return AccessMutualTLSCertificate{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// DeleteAccessMutualTLSCertificate destroys an account level Access Mutual
// TLS certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-update-access-certificate
func (api *API) DeleteAccessMutualTLSCertificate(accountID, certificateID string) error {
	return api.deleteAccessMutualTLSCertificate(accountID, certificateID, AccountRouteRoot)
}

// DeleteZoneAccessMutualTLSCertificate destroys a zone level Access Mutual TLS
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-update-access-certificate
func (api *API) DeleteZoneAccessMutualTLSCertificate(zoneID, certificateID string) error {
	return api.deleteAccessMutualTLSCertificate(zoneID, certificateID, ZoneRouteRoot)
}

func (api *API) deleteAccessMutualTLSCertificate(id, certificateID string, routeRoot RouteRoot) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		routeRoot,
		id,
		certificateID,
	)

	res, err := api.makeRequestContext(context.Background(), "DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}

	return nil
}
