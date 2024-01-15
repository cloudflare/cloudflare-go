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
func (r *AccountDlpProfilePredefinedService) Get(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfilePredefinedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a DLP predefined profile. Only supports enabling/disabling entries.
func (r *AccountDlpProfilePredefinedService) Update(ctx context.Context, accountIdentifier string, profileID string, body AccountDlpProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *AccountDlpProfilePredefinedUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountDlpProfilePredefinedGetResponse struct {
	Errors   []AccountDlpProfilePredefinedGetResponseError   `json:"errors"`
	Messages []AccountDlpProfilePredefinedGetResponseMessage `json:"messages"`
	Result   AccountDlpProfilePredefinedGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDlpProfilePredefinedGetResponseSuccess `json:"success"`
	JSON    accountDlpProfilePredefinedGetResponseJSON    `json:"-"`
}

// accountDlpProfilePredefinedGetResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfilePredefinedGetResponse]
type accountDlpProfilePredefinedGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfilePredefinedGetResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountDlpProfilePredefinedGetResponseErrorJSON `json:"-"`
}

// accountDlpProfilePredefinedGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountDlpProfilePredefinedGetResponseError]
type accountDlpProfilePredefinedGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfilePredefinedGetResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountDlpProfilePredefinedGetResponseMessageJSON `json:"-"`
}

// accountDlpProfilePredefinedGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountDlpProfilePredefinedGetResponseMessage]
type accountDlpProfilePredefinedGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfilePredefinedGetResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// The entries for this profile.
	Entries []AccountDlpProfilePredefinedGetResponseResultEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type AccountDlpProfilePredefinedGetResponseResultType `json:"type"`
	JSON accountDlpProfilePredefinedGetResponseResultJSON `json:"-"`
}

// accountDlpProfilePredefinedGetResponseResultJSON contains the JSON metadata for
// the struct [AccountDlpProfilePredefinedGetResponseResult]
type accountDlpProfilePredefinedGetResponseResultJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A predefined entry that matches a profile
type AccountDlpProfilePredefinedGetResponseResultEntry struct {
	// UUID
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                           `json:"profile_id"`
	JSON      accountDlpProfilePredefinedGetResponseResultEntryJSON `json:"-"`
}

// accountDlpProfilePredefinedGetResponseResultEntryJSON contains the JSON metadata
// for the struct [AccountDlpProfilePredefinedGetResponseResultEntry]
type accountDlpProfilePredefinedGetResponseResultEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedGetResponseResultEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type AccountDlpProfilePredefinedGetResponseResultType string

const (
	AccountDlpProfilePredefinedGetResponseResultTypePredefined AccountDlpProfilePredefinedGetResponseResultType = "predefined"
)

// Whether the API call was successful
type AccountDlpProfilePredefinedGetResponseSuccess bool

const (
	AccountDlpProfilePredefinedGetResponseSuccessTrue AccountDlpProfilePredefinedGetResponseSuccess = true
)

type AccountDlpProfilePredefinedUpdateResponse struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// The entries for this profile.
	Entries []AccountDlpProfilePredefinedUpdateResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type AccountDlpProfilePredefinedUpdateResponseType `json:"type"`
	JSON accountDlpProfilePredefinedUpdateResponseJSON `json:"-"`
}

// accountDlpProfilePredefinedUpdateResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfilePredefinedUpdateResponse]
type accountDlpProfilePredefinedUpdateResponseJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A predefined entry that matches a profile
type AccountDlpProfilePredefinedUpdateResponseEntry struct {
	// UUID
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                        `json:"profile_id"`
	JSON      accountDlpProfilePredefinedUpdateResponseEntryJSON `json:"-"`
}

// accountDlpProfilePredefinedUpdateResponseEntryJSON contains the JSON metadata
// for the struct [AccountDlpProfilePredefinedUpdateResponseEntry]
type accountDlpProfilePredefinedUpdateResponseEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfilePredefinedUpdateResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type AccountDlpProfilePredefinedUpdateResponseType string

const (
	AccountDlpProfilePredefinedUpdateResponseTypePredefined AccountDlpProfilePredefinedUpdateResponseType = "predefined"
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
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountDlpProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
