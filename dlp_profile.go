package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DLPPattern represents a DLP Pattern that matches an entry.
type DLPPattern struct {
	Regex      string `json:"regex,omitempty"`
	Validation string `json:"validation,omitempty"`
}

// DLPEntry represents a DLP Entry, which can be matched in HTTP bodies or files.
type DLPEntry struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ProfileID string `json:"profile_id,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Type      string `json:"type,omitempty"`

	// The following fields are only present for custom entries.

	Pattern   *DLPPattern `json:"pattern,omitempty"`
	CreatedAt *time.Time  `json:"created_at,omitempty"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty"`
}

// DLPProfile represents a DLP Profile, which contains a set
// of entries.
type DLPProfile struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`

	// The following fields are omitted for predefined DLP
	// profiles
	Entries   []DLPEntry `json:"entries,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// DLPProfilesCreateRequest represents a request to create a
// set of profiles.
type DLPProfilesCreateRequest struct {
	Profiles []DLPProfile `json:"profiles"`
}

// DLPProfileListResponse represents the response from the list
// dlp profiles endpoint.
type DLPProfileListResponse struct {
	Result []DLPProfile `json:"result"`
	Response
}

// DLPProfileResponse is the API response, containing a single
// access application.
type DLPProfileResponse struct {
	Success  bool       `json:"success"`
	Errors   []string   `json:"errors"`
	Messages []string   `json:"messages"`
	Result   DLPProfile `json:"result"`
}

// DLPProfiles returns all DLP profiles within an account.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-list-all-profiles
func (api *API) DLPProfiles(ctx context.Context, accountID string) ([]DLPProfile, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles", AccountRouteRoot, accountID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DLPProfile{}, err
	}

	var dlpProfilesListResponse DLPProfileListResponse
	err = json.Unmarshal(res, &dlpProfilesListResponse)
	if err != nil {
		return []DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpProfilesListResponse.Result, nil
}

// DLPProfile returns a single DLP profile (custom or predefined) based on the
// profile ID.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-get-dlp-profile
func (api *API) DLPProfile(ctx context.Context, accountID, profileID string) (DLPProfile, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/%s", AccountRouteRoot, accountID, profileID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DLPProfile{}, err
	}

	var dlpProfileResponse DLPProfileResponse
	err = json.Unmarshal(res, &dlpProfileResponse)
	if err != nil {
		return DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpProfileResponse.Result, nil
}

// CreateDLPCustomProfiles creates a set of custom DLP Profile.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-create-custom-profiles
func (api *API) CreateDLPCustomProfiles(ctx context.Context, accountID string, profiles []DLPProfile) ([]DLPProfile, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/custom", AccountRouteRoot, accountID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, DLPProfilesCreateRequest{
		Profiles: profiles,
	})
	if err != nil {
		return []DLPProfile{}, err
	}

	var dLPCustomProfilesResponse DLPProfileListResponse
	err = json.Unmarshal(res, &dLPCustomProfilesResponse)
	if err != nil {
		return []DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dLPCustomProfilesResponse.Result, nil
}

// CreateDLPCustomProfile is a wrapper for CreateDLPCustomProfiles that creates a single custom DLP Profile.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-create-custom-profiles
func (api *API) CreateDLPCustomProfile(ctx context.Context, accountID string, profile DLPProfile) (DLPProfile, error) {
	profiles, err := api.CreateDLPCustomProfiles(ctx, accountID, []DLPProfile{profile})
	if err != nil {
		return DLPProfile{}, err
	}
	if len(profiles) == 0 {
		return DLPProfile{}, fmt.Errorf("%s: no profile found in the response", errUnmarshalError)
	}
	return profiles[0], nil
}

// DeleteDLPCustomProfile deletes a DLP custom profile.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-delete-custom-profile
func (api *API) DeleteDLPCustomProfile(ctx context.Context, accountID, profileID string) error {
	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/custom/%s", AccountRouteRoot, accountID, profileID), nil)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}

// UpdateDLPCustomProfile updates a DLP custom profile.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-update-custom-profile
func (api *API) UpdateDLPCustomProfile(ctx context.Context, accountID, profileID string, profile DLPProfile) (DLPProfile, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/custom/%s", AccountRouteRoot, accountID, profileID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, profile)
	if err != nil {
		return DLPProfile{}, err
	}

	var dlpProfileResponse DLPProfileResponse
	err = json.Unmarshal(res, &dlpProfileResponse)
	if err != nil {
		return DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpProfileResponse.Result, nil
}

// UpdateDLPPredefinedProfile updates a DLP predefined profile.
//
// API reference: https://api.cloudflare.com/#dlp-profiles-update-predefined-profile
func (api *API) UpdateDLPPredefinedProfile(ctx context.Context, accountID, profileID string, profile DLPProfile) (DLPProfile, error) {
	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/profiles/predefined/%s", AccountRouteRoot, accountID, profileID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, profile)
	if err != nil {
		return DLPProfile{}, err
	}

	var dlpProfileResponse DLPProfileResponse
	err = json.Unmarshal(res, &dlpProfileResponse)
	if err != nil {
		return DLPProfile{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpProfileResponse.Result, nil
}
