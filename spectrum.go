package cloudflare

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// SpectrumApplication defines a single Spectrum Application.
type SpectrumApplication struct {
	ID            string                 `json:"id"`
	Protocol      string                 `json:"protocol"`
	IPv4          bool                   `json:"ipv4"`
	DNS           SpectrumApplicationDNS `json:"dns"`
	OriginDirect  []string               `json:"origin_direct"`
	OriginPort    int                    `json:"origin_port"`
	IPFirewall    bool                   `json:"ip_firewall"`
	ProxyProtocol bool                   `json:"proxy_protocol"`
	TLS           string                 `json:"tls"`
	CreatedOn     *time.Time             `json:"created_on"`
	ModifiedOn    *time.Time             `json:"modified_on"`
}

// SpectrumApplicationDNS holds the origin DNS configuration for a Spectrum
// Application.
type SpectrumApplicationDNS struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// SpectrumApplicationDetailResponse is the structure of the detailed response
// from the API.
type SpectrumApplicationDetailResponse struct {
	Success  bool                `json:"success"`
	Errors   []string            `json:"errors"`
	Messages []string            `json:"messages"`
	Result   SpectrumApplication `json:"result"`
}

// SpectrumApplicationsDetailResponse is the structure of the detailed response
// from the API.
type SpectrumApplicationsDetailResponse struct {
	Success  bool                  `json:"success"`
	Errors   []string              `json:"errors"`
	Messages []string              `json:"messages"`
	Result   []SpectrumApplication `json:"result"`
}

// DeletedSpectrumApplicationResponse defines the API response when deleting a
// Spectrum application.
type DeletedSpectrumApplicationResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   DeletedSpectrumApplicationResult `json:"result"`
}

// DeletedSpectrumApplicationResult defines the struct for a deleted application
// result.
type DeletedSpectrumApplicationResult struct {
	ID string `json:"id"`
}

// SpectrumApplications fetches all of the Spectrum applications for a zone.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#list-spectrum-applications
func (api *API) SpectrumApplications(zoneID string) ([]SpectrumApplication, error) {
	uri := "/zones/" + zoneID + "/spectrum/apps"

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []SpectrumApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApplications SpectrumApplicationsDetailResponse
	err = json.Unmarshal(res, &spectrumApplications)
	if err != nil {
		return []SpectrumApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApplications.Result, nil
}

// SpectrumApplication fetches a single Spectrum application based on the ID.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#list-spectrum-applications
func (api *API) SpectrumApplication(zoneID string, applicationID string) (SpectrumApplication, error) {
	uri := fmt.Sprintf(
		"/zones/%s/spectrum/apps/%s",
		zoneID,
		applicationID,
	)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApplication SpectrumApplicationDetailResponse
	err = json.Unmarshal(res, &spectrumApplication)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApplication.Result, nil
}

// UpdateSpectrumApplication fetches a single Spectrum application based on the ID.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#update-a-spectrum-application
func (api *API) UpdateSpectrumApplication(zoneID, appID string, appDetails SpectrumApplication) (SpectrumApplication, error) {
	uri := fmt.Sprintf(
		"/zones/%s/spectrum/apps/%s",
		zoneID,
		appID,
	)

	res, err := api.makeRequest("PUT", uri, appDetails)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var spectrumApplication SpectrumApplicationDetailResponse
	err = json.Unmarshal(res, &spectrumApplication)
	if err != nil {
		return SpectrumApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return spectrumApplication.Result, nil
}

// DeleteSpectrumApplication removes a Spectrum application based on the ID.
//
// API reference: https://developers.cloudflare.com/spectrum/api-reference/#delete-a-spectrum-application
func (api *API) DeleteSpectrumApplication(zoneID string, applicationID string) (DeletedSpectrumApplicationResult, error) {
	uri := fmt.Sprintf(
		"/zones/%s/spectrum/apps/%s",
		zoneID,
		applicationID,
	)

	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return DeletedSpectrumApplicationResult{}, errors.Wrap(err, errMakeRequestError)
	}

	var deletedSpectrumApplication DeletedSpectrumApplicationResponse
	err = json.Unmarshal(res, &deletedSpectrumApplication)
	if err != nil {
		return DeletedSpectrumApplicationResult{}, errors.Wrap(err, errUnmarshalError)
	}

	return deletedSpectrumApplication.Result, nil
}
