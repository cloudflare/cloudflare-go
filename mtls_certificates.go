package cloudflare

import (
	"context"
	"encoding/json"
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
	Result []MTLSCertificateDetails `json:"result"`
}

// MTLSCertificateParams represents the data related to the mTLS certificate being uploaded. Name is an optional field.
type MTLSCertificateParams struct {
	Name         string `json:"name"`
	Certificates string `json:"certificates"`
	PrivateKey   string `json:"private_key"`
	CA           bool   `json:"ca"`
}

// ListMTLSCertificates returns a list of all user-uploaded mTLS certificates.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-list-mtls-certificates
func (api *API) ListMTLSCertificates(ctx context.Context, accountID string) ([]MTLSCertificateDetails, error) {
	uri := fmt.Sprintf("/accounts/%s/mtls_certificates", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []MTLSCertificateDetails{}, err
	}
	var r MTLSCertificatesResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []MTLSCertificateDetails{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetMTLSCertificateDetails returns the metadata associated with a user-uploaded mTLS certificate.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-get-mtls-certificate
func (api *API) GetMTLSCertificateDetails(ctx context.Context, accountID, certificateID string) (MTLSCertificateDetails, error) {
	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s", accountID, certificateID)
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
func (api *API) ListMTLSCertificateAssociations(ctx context.Context, accountID, certificateID string) ([]MTLSAssociationDetails, error) {
	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s/associations", accountID, certificateID)
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

// UploadMTLSCertificate will upload the provided certificate for use with mTLS enabled Cloudflare services.
//
// API reference: https://api.cloudflare.com/#mtls-certificate-management-upload-mtls-certificate
func (api *API) UploadMTLSCertificate(ctx context.Context, accountID string, params MTLSCertificateParams) (MTLSCertificateDetails, error) {
	uri := fmt.Sprintf("/accounts/%s/mtls_certificates", accountID)
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
func (api *API) DeleteMTLSCertificate(ctx context.Context, accountID, certificateID string) (MTLSCertificateDetails, error) {
	uri := fmt.Sprintf("/accounts/%s/mtls_certificates/%s", accountID, certificateID)
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
