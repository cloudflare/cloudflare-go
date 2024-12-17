package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type TeamsCertificate struct {
	InUse         *bool      `json:"in_use"`
	ID            string     `json:"id"`
	BindingStatus string     `json:"binding_status"`
	QsPackId      string     `json:"qs_pack_id"`
	Type          string     `json:"type"`
	UpdatedAt     *time.Time `json:"updated_at"`
	UploadedOn    *time.Time `json:"uploaded_on"`
	CreatedAt     *time.Time `json:"created_at"`
	ExpiresOn     *time.Time `json:"expires_on"`
}

type TeamsCertificateCreateRequest struct {
	ValidityPeriodDays int `json:"validity_period_days,omitempty"`
}

const DEFAULT_VALIDITY_PERIOD_DAYS = 1826

// TeamsCertificateResponse is the API response, containing a single certificate.
type TeamsCertificateResponse struct {
	Response
	Result TeamsCertificate `json:"result"`
}

// TeamsCertificatesResponse is the API response, containing an array of certificates.
type TeamsCertificatesResponse struct {
	Response
	Result []TeamsCertificate `json:"result"`
}

// TeamsCertificates returns all certificates in an account
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/gateway/subresources/certificates/methods/list/
func (api *API) TeamsCertificates(ctx context.Context, accountID string) ([]TeamsCertificate, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/certificates", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsCertificate{}, err
	}

	var teamsCertificatesResponse TeamsCertificatesResponse
	err = json.Unmarshal(res, &teamsCertificatesResponse)
	if err != nil {
		return []TeamsCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsCertificatesResponse.Result, nil
}

// TeamsCertificate returns teams account certificate.
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/gateway/subresources/certificates/methods/get/
func (api *API) TeamsCertificate(ctx context.Context, accountID string, certificateId string) (TeamsCertificate, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/certificates/%s", accountID, certificateId)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsCertificate{}, err
	}

	var teamsCertificateResponse TeamsCertificateResponse
	err = json.Unmarshal(res, &teamsCertificateResponse)
	if err != nil {
		return TeamsCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsCertificateResponse.Result, nil
}

// TeamsGenerateCertificate generates a new gateway managed certificate
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/gateway/subresources/certificates/methods/create/
func (api *API) TeamsGenerateCertificate(ctx context.Context, accountID string, certificateRequest TeamsCertificateCreateRequest) (TeamsCertificate, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/certificates", accountID)

	if certificateRequest.ValidityPeriodDays == 0 {
		certificateRequest.ValidityPeriodDays = DEFAULT_VALIDITY_PERIOD_DAYS
	}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, certificateRequest)
	if err != nil {
		return TeamsCertificate{}, err
	}

	var teamsCertResponse TeamsCertificateResponse
	err = json.Unmarshal(res, &teamsCertResponse)
	if err != nil {
		return TeamsCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsCertResponse.Result, nil
}

// TeamsActivateCertificate activates a certificate
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/gateway/subresources/certificates/methods/activate/
func (api *API) TeamsActivateCertificate(ctx context.Context, accountID string, certificateId string) (TeamsCertificate, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/certificates/%s/activate", accountID, certificateId)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return TeamsCertificate{}, err
	}

	var teamsCertResponse TeamsCertificateResponse
	err = json.Unmarshal(res, &teamsCertResponse)
	if err != nil {
		return TeamsCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsCertResponse.Result, nil
}

// TeamsDectivateCertificate deactivates a certificate
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/gateway/subresources/certificates/methods/deactivate/
func (api *API) TeamsDeactivateCertificate(ctx context.Context, accountID string, certificateId string) (TeamsCertificate, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/certificates/%s/deactivate", accountID, certificateId)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return TeamsCertificate{}, err
	}

	var teamsCertResponse TeamsCertificateResponse
	err = json.Unmarshal(res, &teamsCertResponse)
	if err != nil {
		return TeamsCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsCertResponse.Result, nil
}

// TeamsDeleteCertificate deletes a certificate.
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/gateway/subresources/certificates/methods/delete/
func (api *API) TeamsDeleteCertificate(ctx context.Context, accountID string, certificateId string) error {
	uri := fmt.Sprintf("/accounts/%s/gateway/certificates/%s", accountID, certificateId)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
