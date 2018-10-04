package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// AccessApplication represents an Access application.
type AccessApplication struct {
	ID              string     `json:"id,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	AUD             string     `json:"aud,omitempty"`
	Name            string     `json:"name"`
	Domain          string     `json:"domain"`
	SessionDuration string     `json:"session_duration,omitempty"`
}

// AccessApplicationListResponse represents the response from the list
// access applications endpoint.
type AccessApplicationListResponse struct {
	Result []AccessApplication `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccessApplicationDetailResponse is the API response, containing a single
// access application.
type AccessApplicationDetailResponse struct {
	Success  bool              `json:"success"`
	Errors   []string          `json:"errors"`
	Messages []string          `json:"messages"`
	Result   AccessApplication `json:"result"`
}

// AccessApplications returns all applications within a zone.
//
// API reference: https://api.cloudflare.com/#access-applications-list-access-applications
func (api *API) AccessApplications(zoneID string, pageOpts PaginationOptions) ([]AccessApplication, ResultInfo, error) {
	v := url.Values{}
	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}
	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	uri := "/zones/" + zoneID + "/access/apps"
	if len(v) > 0 {
		uri = uri + "?" + v.Encode()
	}

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []AccessApplication{}, ResultInfo{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessApplicationListResponse AccessApplicationListResponse
	err = json.Unmarshal(res, &accessApplicationListResponse)
	if err != nil {
		return []AccessApplication{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessApplicationListResponse.Result, accessApplicationListResponse.ResultInfo, nil
}

// AccessApplication returns a single application based on the
// application ID.
//
// API reference: https://api.cloudflare.com/#access-applications-access-applications-details
func (api *API) AccessApplication(zoneID, applicationID string) (AccessApplication, error) {
	uri := fmt.Sprintf(
		"/zones/%s/access/apps/%s",
		zoneID,
		applicationID,
	)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return AccessApplication{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessApplicationDetailResponse AccessApplicationDetailResponse
	err = json.Unmarshal(res, &accessApplicationDetailResponse)
	if err != nil {
		return AccessApplication{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessApplicationDetailResponse.Result, nil
}

