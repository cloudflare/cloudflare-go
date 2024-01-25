// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccountDlpProfileCustomService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDlpProfileCustomService] method instead.
type AccountDlpProfileCustomService struct {
	Options []option.RequestOption
}

// NewAccountDlpProfileCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpProfileCustomService(opts ...option.RequestOption) (r *AccountDlpProfileCustomService) {
	r = &AccountDlpProfileCustomService{}
	r.Options = opts
	return
}

// Fetches a custom DLP profile.
func (r *AccountDlpProfileCustomService) Get(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfileCustomGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a DLP custom profile.
func (r *AccountDlpProfileCustomService) Update(ctx context.Context, accountIdentifier string, profileID string, body AccountDlpProfileCustomUpdateParams, opts ...option.RequestOption) (res *AccountDlpProfileCustomUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a DLP custom profile.
func (r *AccountDlpProfileCustomService) Delete(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfileCustomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a set of DLP custom profiles.
func (r *AccountDlpProfileCustomService) DlpProfilesNewCustomProfiles(ctx context.Context, accountIdentifier string, body AccountDlpProfileCustomDlpProfilesNewCustomProfilesParams, opts ...option.RequestOption) (res *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDlpProfileCustomGetResponse struct {
	Errors   []AccountDlpProfileCustomGetResponseError   `json:"errors"`
	Messages []AccountDlpProfileCustomGetResponseMessage `json:"messages"`
	Result   AccountDlpProfileCustomGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDlpProfileCustomGetResponseSuccess `json:"success"`
	JSON    accountDlpProfileCustomGetResponseJSON    `json:"-"`
}

// accountDlpProfileCustomGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileCustomGetResponse]
type accountDlpProfileCustomGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDlpProfileCustomGetResponseErrorJSON `json:"-"`
}

// accountDlpProfileCustomGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomGetResponseError]
type accountDlpProfileCustomGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountDlpProfileCustomGetResponseMessageJSON `json:"-"`
}

// accountDlpProfileCustomGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomGetResponseMessage]
type accountDlpProfileCustomGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomGetResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []AccountDlpProfileCustomGetResponseResultEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      AccountDlpProfileCustomGetResponseResultType `json:"type"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileCustomGetResponseResultJSON `json:"-"`
}

// accountDlpProfileCustomGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomGetResponseResult]
type accountDlpProfileCustomGetResponseResultJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDlpProfileCustomGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A custom entry that matches a profile
type AccountDlpProfileCustomGetResponseResultEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern AccountDlpProfileCustomGetResponseResultEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                       `json:"profile_id"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileCustomGetResponseResultEntryJSON `json:"-"`
}

// accountDlpProfileCustomGetResponseResultEntryJSON contains the JSON metadata for
// the struct [AccountDlpProfileCustomGetResponseResultEntry]
type accountDlpProfileCustomGetResponseResultEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomGetResponseResultEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type AccountDlpProfileCustomGetResponseResultEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation AccountDlpProfileCustomGetResponseResultEntriesPatternValidation `json:"validation"`
	JSON       accountDlpProfileCustomGetResponseResultEntriesPatternJSON       `json:"-"`
}

// accountDlpProfileCustomGetResponseResultEntriesPatternJSON contains the JSON
// metadata for the struct [AccountDlpProfileCustomGetResponseResultEntriesPattern]
type accountDlpProfileCustomGetResponseResultEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomGetResponseResultEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type AccountDlpProfileCustomGetResponseResultEntriesPatternValidation string

const (
	AccountDlpProfileCustomGetResponseResultEntriesPatternValidationLuhn AccountDlpProfileCustomGetResponseResultEntriesPatternValidation = "luhn"
)

// The type of the profile.
type AccountDlpProfileCustomGetResponseResultType string

const (
	AccountDlpProfileCustomGetResponseResultTypeCustom AccountDlpProfileCustomGetResponseResultType = "custom"
)

// Whether the API call was successful
type AccountDlpProfileCustomGetResponseSuccess bool

const (
	AccountDlpProfileCustomGetResponseSuccessTrue AccountDlpProfileCustomGetResponseSuccess = true
)

type AccountDlpProfileCustomUpdateResponse struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []AccountDlpProfileCustomUpdateResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      AccountDlpProfileCustomUpdateResponseType `json:"type"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileCustomUpdateResponseJSON `json:"-"`
}

// accountDlpProfileCustomUpdateResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomUpdateResponse]
type accountDlpProfileCustomUpdateResponseJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDlpProfileCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A custom entry that matches a profile
type AccountDlpProfileCustomUpdateResponseEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern AccountDlpProfileCustomUpdateResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                    `json:"profile_id"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileCustomUpdateResponseEntryJSON `json:"-"`
}

// accountDlpProfileCustomUpdateResponseEntryJSON contains the JSON metadata for
// the struct [AccountDlpProfileCustomUpdateResponseEntry]
type accountDlpProfileCustomUpdateResponseEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomUpdateResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type AccountDlpProfileCustomUpdateResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation AccountDlpProfileCustomUpdateResponseEntriesPatternValidation `json:"validation"`
	JSON       accountDlpProfileCustomUpdateResponseEntriesPatternJSON       `json:"-"`
}

// accountDlpProfileCustomUpdateResponseEntriesPatternJSON contains the JSON
// metadata for the struct [AccountDlpProfileCustomUpdateResponseEntriesPattern]
type accountDlpProfileCustomUpdateResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomUpdateResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type AccountDlpProfileCustomUpdateResponseEntriesPatternValidation string

const (
	AccountDlpProfileCustomUpdateResponseEntriesPatternValidationLuhn AccountDlpProfileCustomUpdateResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type AccountDlpProfileCustomUpdateResponseType string

const (
	AccountDlpProfileCustomUpdateResponseTypeCustom AccountDlpProfileCustomUpdateResponseType = "custom"
)

type AccountDlpProfileCustomDeleteResponse struct {
	Errors   []AccountDlpProfileCustomDeleteResponseError   `json:"errors"`
	Messages []AccountDlpProfileCustomDeleteResponseMessage `json:"messages"`
	Result   AccountDlpProfileCustomDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDlpProfileCustomDeleteResponseSuccess `json:"success"`
	JSON    accountDlpProfileCustomDeleteResponseJSON    `json:"-"`
}

// accountDlpProfileCustomDeleteResponseJSON contains the JSON metadata for the
// struct [AccountDlpProfileCustomDeleteResponse]
type accountDlpProfileCustomDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountDlpProfileCustomDeleteResponseErrorJSON `json:"-"`
}

// accountDlpProfileCustomDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountDlpProfileCustomDeleteResponseError]
type accountDlpProfileCustomDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountDlpProfileCustomDeleteResponseMessageJSON `json:"-"`
}

// accountDlpProfileCustomDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountDlpProfileCustomDeleteResponseMessage]
type accountDlpProfileCustomDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountDlpProfileCustomDeleteResponseResultUnknown] or
// [shared.UnionString].
type AccountDlpProfileCustomDeleteResponseResult interface {
	ImplementsAccountDlpProfileCustomDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDlpProfileCustomDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type AccountDlpProfileCustomDeleteResponseSuccess bool

const (
	AccountDlpProfileCustomDeleteResponseSuccessTrue AccountDlpProfileCustomDeleteResponseSuccess = true
)

type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponse struct {
	Errors     []AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseError    `json:"errors"`
	Messages   []AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseMessage  `json:"messages"`
	Result     []AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResult   `json:"result"`
	ResultInfo AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseSuccess `json:"success"`
	JSON    accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseJSON    `json:"-"`
}

// accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseJSON contains the
// JSON metadata for the struct
// [AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponse]
type accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseError struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseErrorJSON `json:"-"`
}

// accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseError]
type accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseMessage struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseMessageJSON `json:"-"`
}

// accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseMessage]
type accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultType `json:"type"`
	UpdatedAt time.Time                                                             `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultJSON `json:"-"`
}

// accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultJSON contains
// the JSON metadata for the struct
// [AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResult]
type accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A custom entry that matches a profile
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                                                `json:"profile_id"`
	UpdatedAt time.Time                                                                  `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntryJSON `json:"-"`
}

// accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntryJSON
// contains the JSON metadata for the struct
// [AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntry]
type accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPatternValidation `json:"validation"`
	JSON       accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPatternJSON       `json:"-"`
}

// accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPatternJSON
// contains the JSON metadata for the struct
// [AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPattern]
type accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPatternValidation string

const (
	AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPatternValidationLuhn AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultEntriesPatternValidation = "luhn"
)

// The type of the profile.
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultType string

const (
	AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultTypeCustom AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultType = "custom"
)

type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                   `json:"total_count"`
	JSON       accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultInfoJSON `json:"-"`
}

// accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultInfo]
type accountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseSuccess bool

const (
	AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseSuccessTrue AccountDlpProfileCustomDlpProfilesNewCustomProfilesResponseSuccess = true
)

type AccountDlpProfileCustomUpdateParams struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The custom entries for this profile. Array elements with IDs are modifying the
	// existing entry with that ID. Elements without ID will create new entries. Any
	// entry not in the list will be deleted.
	Entries param.Field[[]AccountDlpProfileCustomUpdateParamsEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]AccountDlpProfileCustomUpdateParamsSharedEntry] `json:"shared_entries"`
}

func (r AccountDlpProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom entry that matches a profile
type AccountDlpProfileCustomUpdateParamsEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
	// The name of the entry.
	Name param.Field[string] `json:"name"`
	// A pattern that matches an entry
	Pattern param.Field[AccountDlpProfileCustomUpdateParamsEntriesPattern] `json:"pattern"`
	// ID of the parent profile
	ProfileID param.Field[interface{}] `json:"profile_id"`
}

func (r AccountDlpProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A pattern that matches an entry
type AccountDlpProfileCustomUpdateParamsEntriesPattern struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation param.Field[AccountDlpProfileCustomUpdateParamsEntriesPatternValidation] `json:"validation"`
}

func (r AccountDlpProfileCustomUpdateParamsEntriesPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type AccountDlpProfileCustomUpdateParamsEntriesPatternValidation string

const (
	AccountDlpProfileCustomUpdateParamsEntriesPatternValidationLuhn AccountDlpProfileCustomUpdateParamsEntriesPatternValidation = "luhn"
)

// Properties of a predefined entry in a custom profile
//
// Satisfied by
// [AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined],
// [AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdateIntegration].
type AccountDlpProfileCustomUpdateParamsSharedEntry interface {
	implementsAccountDlpProfileCustomUpdateParamsSharedEntry()
}

// Properties of a predefined entry in a custom profile
type AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdatePredefined) implementsAccountDlpProfileCustomUpdateParamsSharedEntry() {
}

// Properties of an integration entry in a custom profile
type AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdateIntegration struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdateIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountDlpProfileCustomUpdateParamsSharedEntriesDlpSharedEntryUpdateIntegration) implementsAccountDlpProfileCustomUpdateParamsSharedEntry() {
}

type AccountDlpProfileCustomDlpProfilesNewCustomProfilesParams struct {
	Profiles param.Field[[]AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfile] `json:"profiles,required"`
}

func (r AccountDlpProfileCustomDlpProfilesNewCustomProfilesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfile struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The entries for this profile.
	Entries param.Field[[]AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
}

func (r AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom entry create payload
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The name of the entry.
	Name param.Field[string] `json:"name,required"`
	// A pattern that matches an entry
	Pattern param.Field[AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern] `json:"pattern,required"`
}

func (r AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A pattern that matches an entry
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation param.Field[AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidation] `json:"validation"`
}

func (r AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidation string

const (
	AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn AccountDlpProfileCustomDlpProfilesNewCustomProfilesParamsProfilesEntriesPatternValidation = "luhn"
)
