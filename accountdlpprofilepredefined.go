// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDlpProfilePredefinedService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDlpProfilePredefinedService] method instead.
type AccountDlpProfilePredefinedService struct {
	Options []option.RequestOption
}

// NewAccountDlpProfilePredefinedService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDlpProfilePredefinedService(opts ...option.RequestOption) (r *AccountDlpProfilePredefinedService) {
	r = &AccountDlpProfilePredefinedService{}
	r.Options = opts
	return
}

// Fetches a predefined DLP profile.
func (r *AccountDlpProfilePredefinedService) Get(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *PredefinedProfileResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *AccountDlpProfilePredefinedService) Update(ctx context.Context, accountIdentifier string, profileID string, body AccountDlpProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *PredefinedProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type PredefinedProfile struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// The entries for this profile.
	Entries []PredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type PredefinedProfileType `json:"type"`
	JSON predefinedProfileJSON `json:"-"`
}

// predefinedProfileJSON contains the JSON metadata for the struct
// [PredefinedProfile]
type predefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r PredefinedProfile) implementsEitherProfileResponseSNgAtLbhResult() {}

func (r PredefinedProfile) implementsProfilesResponseCollectionResult() {}

// A predefined entry that matches a profile
type PredefinedProfileEntry struct {
	// UUID
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                `json:"profile_id"`
	JSON      predefinedProfileEntryJSON `json:"-"`
}

// predefinedProfileEntryJSON contains the JSON metadata for the struct
// [PredefinedProfileEntry]
type predefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type PredefinedProfileType string

const (
	PredefinedProfileTypePredefined PredefinedProfileType = "predefined"
)

type PredefinedProfileResponse struct {
	Errors   []PredefinedProfileResponseError   `json:"errors"`
	Messages []PredefinedProfileResponseMessage `json:"messages"`
	Result   PredefinedProfile                  `json:"result"`
	// Whether the API call was successful
	Success PredefinedProfileResponseSuccess `json:"success"`
	JSON    predefinedProfileResponseJSON    `json:"-"`
}

// predefinedProfileResponseJSON contains the JSON metadata for the struct
// [PredefinedProfileResponse]
type predefinedProfileResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PredefinedProfileResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PredefinedProfileResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    predefinedProfileResponseErrorJSON `json:"-"`
}

// predefinedProfileResponseErrorJSON contains the JSON metadata for the struct
// [PredefinedProfileResponseError]
type predefinedProfileResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PredefinedProfileResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PredefinedProfileResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    predefinedProfileResponseMessageJSON `json:"-"`
}

// predefinedProfileResponseMessageJSON contains the JSON metadata for the struct
// [PredefinedProfileResponseMessage]
type predefinedProfileResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PredefinedProfileResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PredefinedProfileResponseSuccess bool

const (
	PredefinedProfileResponseSuccessTrue PredefinedProfileResponseSuccess = true
)

type AccountDlpProfilePredefinedUpdateParams struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// The entries for this profile.
	Entries param.Field[[]AccountDlpProfilePredefinedUpdateParamsEntry] `json:"entries"`
}

func (r AccountDlpProfilePredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDlpProfilePredefinedUpdateParamsEntry struct {
	// Wheter the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountDlpProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
