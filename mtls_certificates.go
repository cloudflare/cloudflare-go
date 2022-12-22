package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// MTLSAssociationDetails represents the metadata for an existing association between a user-uploaded mTLS certificate and a Cloudflare service.
type MTLSAssociationDetails struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

// MTLSAssociationResponse represents the response from the retrieval endpoint for mTLS certificate associations.
type MTLSAssociationResponse struct {
	Response
	Result []MTLSAssociationDetails `json:"result"`
}

// MTLSCertificateDetails represents the metadata for a user-uploaded mTLS certificate.
type MTLSCertificateDetails struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Issuer       string    `json:"issuer"`
	Signature    string    `json:"signature"`
	SerialNumber string    `json:"serial_number"`
	Certificates string    `json:"certificates"`
	CA           bool      `json:"ca"`
	UploadedOn   time.Time `json:"uploaded_on"`
	UpdatedAt    time.Time `json:"updated_at"`
	ExpiresOn    time.Time `json:"expires_on"`
}

// MTLSCertificateResponse represents the response from endpoints relating to retrieving, creating, and deleting an mTLS certificate.
type MTLSCertificateResponse struct {
	Response
	Result MTLSCertificateDetails `json:"result"`
}

// MTLSCertificatesResponse represents the response from the mTLS certificate list endpoint.
type MTLSCertificatesResponse struct {
	Response
	Result     []MTLSCertificateDetails `json:"result"`
	ResultInfo `json:"result_info"`
}

// MTLSCertificateParams represents the data related to the mTLS certificate being uploaded. Name is an optional field.
type CreateMTLSCertificateParams struct {
	Name         string `json:"name"`
	Certificates string `json:"certificates"`
	PrivateKey   string `json:"private_key"`
	CA           bool   `json:"ca"`
}

type ListMTLSCertificatesParams struct {
	PaginationOptions
	Limit  int    `url:"limit,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Name   string `url:"name,omitempty"`
	CA     bool   `url:"ca,omitempty"`
}

var (
	ErrMissingCertificateID = errors.New("missing required certificate ID")
)

// ListMTLSCertificates returns a list of all user-uploaded mTLS certificates.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-list-mtls-certificates
func (api *API) ListMTLSCertificates(ctx context.Context, rc *ResourceContainer, params ListMTLSCertificatesParams) ([]MTLSCertificateDetails, ResultInfo, error) {
	switch {
	case rc.Level != AccountRouteLevel:
		return []MTLSCertificateDetails{}, ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	case rc.Identifier == "":
		return []MTLSCertificateDetails{}, ResultInfo{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return []MTLSCertificateDetails{}, ResultInfo{}, err
	}
	var r MTLSCertificatesResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []MTLSCertificateDetails{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, r.ResultInfo, err
}

// GetMTLSCertificateDetails returns the metadata associated with a user-uploaded mTLS certificate.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-get-mtls-certificate
func (api *API) GetMTLSCertificateDetails(ctx context.Context, rc *ResourceContainer, certificateID string) (MTLSCertificateDetails, error) {
	switch {
	case rc.Level != AccountRouteLevel:
		return MTLSCertificateDetails{}, ErrRequiredAccountLevelResourceContainer
	case rc.Identifier == "":
		return MTLSCertificateDetails{}, ErrMissingAccountID
	case certificateID == "":
		return MTLSCertificateDetails{}, ErrMissingCertificateID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s", rc.Identifier, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return MTLSCertificateDetails{}, err
	}
	var r MTLSCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return MTLSCertificateDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListMTLSCertificateAssociations returns a list of all existing associations between the mTLS certificate and Cloudflare services.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-list-mtls-certificate-associations
func (api *API) ListMTLSCertificateAssociations(ctx context.Context, rc *ResourceContainer, certificateID string) ([]MTLSAssociationDetails, error) {
	switch {
	case rc.Level != AccountRouteLevel:
		return []MTLSAssociationDetails{}, ErrRequiredAccountLevelResourceContainer
	case rc.Identifier == "":
		return []MTLSAssociationDetails{}, ErrMissingAccountID
	case certificateID == "":
		return []MTLSAssociationDetails{}, ErrMissingCertificateID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s/associations", rc.Identifier, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []MTLSAssociationDetails{}, err
	}
	var r MTLSAssociationResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []MTLSAssociationDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateMTLSCertificate will create the provided certificate for use with mTLS enabled Cloudflare services.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-upload-mtls-certificate
func (api *API) CreateMTLSCertificate(ctx context.Context, rc *ResourceContainer, params CreateMTLSCertificateParams) (MTLSCertificateDetails, error) {
	switch {
	case rc.Level != AccountRouteLevel:
		return MTLSCertificateDetails{}, ErrRequiredAccountLevelResourceContainer
	case rc.Identifier == "":
		return MTLSCertificateDetails{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return MTLSCertificateDetails{}, err
	}
	var r MTLSCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return MTLSCertificateDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteMTLSCertificate will delete the specified mTLS certificate.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-delete-mtls-certificate
func (api *API) DeleteMTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string) (MTLSCertificateDetails, error) {
	switch {
	case rc.Level != AccountRouteLevel:
		return MTLSCertificateDetails{}, ErrRequiredAccountLevelResourceContainer
	case rc.Identifier == "":
		return MTLSCertificateDetails{}, ErrMissingAccountID
	case certificateID == "":
		return MTLSCertificateDetails{}, ErrMissingCertificateID
	}

	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s", rc.Identifier, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return MTLSCertificateDetails{}, err
	}
	var r MTLSCertificateResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return MTLSCertificateDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
