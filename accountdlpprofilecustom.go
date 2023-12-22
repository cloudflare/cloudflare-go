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
func (r *AccountDlpProfileCustomService) Get(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *CustomProfileResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a DLP custom profile.
func (r *AccountDlpProfileCustomService) Update(ctx context.Context, accountIdentifier string, profileID string, body AccountDlpProfileCustomUpdateParams, opts ...option.RequestOption) (res *CustomProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a DLP custom profile.
func (r *AccountDlpProfileCustomService) Delete(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a set of DLP custom profiles.
func (r *AccountDlpProfileCustomService) DlpProfilesNewCustomProfiles(ctx context.Context, accountIdentifier string, body AccountDlpProfileCustomDlpProfilesNewCustomProfilesParams, opts ...option.RequestOption) (res *CreateCustomProfileResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseSingle struct {
	Errors   []APIResponseSingleError   `json:"errors"`
	Messages []APIResponseSingleMessage `json:"messages"`
	Result   APIResponseSingleResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success APIResponseSingleSuccess `json:"success"`
	JSON    apiResponseSingleJSON    `json:"-"`
}

// apiResponseSingleJSON contains the JSON metadata for the struct
// [APIResponseSingle]
type apiResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r APIResponseSingle) implementsPackageResponseSingle() {}

type APIResponseSingleError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    apiResponseSingleErrorJSON `json:"-"`
}

// apiResponseSingleErrorJSON contains the JSON metadata for the struct
// [APIResponseSingleError]
type apiResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type APIResponseSingleMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    apiResponseSingleMessageJSON `json:"-"`
}

// apiResponseSingleMessageJSON contains the JSON metadata for the struct
// [APIResponseSingleMessage]
type apiResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [APIResponseSingleResultObject] or [shared.UnionString].
type APIResponseSingleResult interface {
	ImplementsAPIResponseSingleResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseSingleResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type APIResponseSingleSuccess bool

const (
	APIResponseSingleSuccessTrue APIResponseSingleSuccess = true
)

type CreateCustomProfileResponse struct {
	Errors     []CreateCustomProfileResponseError    `json:"errors"`
	Messages   []CreateCustomProfileResponseMessage  `json:"messages"`
	Result     []CustomProfile                       `json:"result"`
	ResultInfo CreateCustomProfileResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CreateCustomProfileResponseSuccess `json:"success"`
	JSON    createCustomProfileResponseJSON    `json:"-"`
}

// createCustomProfileResponseJSON contains the JSON metadata for the struct
// [CreateCustomProfileResponse]
type createCustomProfileResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateCustomProfileResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreateCustomProfileResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    createCustomProfileResponseErrorJSON `json:"-"`
}

// createCustomProfileResponseErrorJSON contains the JSON metadata for the struct
// [CreateCustomProfileResponseError]
type createCustomProfileResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateCustomProfileResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreateCustomProfileResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    createCustomProfileResponseMessageJSON `json:"-"`
}

// createCustomProfileResponseMessageJSON contains the JSON metadata for the struct
// [CreateCustomProfileResponseMessage]
type createCustomProfileResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateCustomProfileResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CreateCustomProfileResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       createCustomProfileResponseResultInfoJSON `json:"-"`
}

// createCustomProfileResponseResultInfoJSON contains the JSON metadata for the
// struct [CreateCustomProfileResponseResultInfo]
type createCustomProfileResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateCustomProfileResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CreateCustomProfileResponseSuccess bool

const (
	CreateCustomProfileResponseSuccessTrue CreateCustomProfileResponseSuccess = true
)

type CustomProfile struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []CustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      CustomProfileType `json:"type"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      customProfileJSON `json:"-"`
}

// customProfileJSON contains the JSON metadata for the struct [CustomProfile]
type customProfileJSON struct {
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

func (r *CustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r CustomProfile) implementsEitherProfileResponseSNgAtLbhResult() {}

func (r CustomProfile) implementsProfilesResponseCollectionResult() {}

// A custom entry that matches a profile
type CustomProfileEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern CustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}            `json:"profile_id"`
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	JSON      customProfileEntryJSON `json:"-"`
}

// customProfileEntryJSON contains the JSON metadata for the struct
// [CustomProfileEntry]
type customProfileEntryJSON struct {
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

func (r *CustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type CustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation CustomProfileEntriesPatternValidation `json:"validation"`
	JSON       customProfileEntriesPatternJSON       `json:"-"`
}

// customProfileEntriesPatternJSON contains the JSON metadata for the struct
// [CustomProfileEntriesPattern]
type customProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type CustomProfileEntriesPatternValidation string

const (
	CustomProfileEntriesPatternValidationLuhn CustomProfileEntriesPatternValidation = "luhn"
)

// The type of the profile.
type CustomProfileType string

const (
	CustomProfileTypeCustom CustomProfileType = "custom"
)

type CustomProfileResponse struct {
	Errors   []CustomProfileResponseError   `json:"errors"`
	Messages []CustomProfileResponseMessage `json:"messages"`
	Result   CustomProfile                  `json:"result"`
	// Whether the API call was successful
	Success CustomProfileResponseSuccess `json:"success"`
	JSON    customProfileResponseJSON    `json:"-"`
}

// customProfileResponseJSON contains the JSON metadata for the struct
// [CustomProfileResponse]
type customProfileResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProfileResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomProfileResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    customProfileResponseErrorJSON `json:"-"`
}

// customProfileResponseErrorJSON contains the JSON metadata for the struct
// [CustomProfileResponseError]
type customProfileResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProfileResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomProfileResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    customProfileResponseMessageJSON `json:"-"`
}

// customProfileResponseMessageJSON contains the JSON metadata for the struct
// [CustomProfileResponseMessage]
type customProfileResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomProfileResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomProfileResponseSuccess bool

const (
	CustomProfileResponseSuccessTrue CustomProfileResponseSuccess = true
)

type AccountDlpProfileCustomUpdateParams struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The entries for this profile. Array elements with IDs are modifying the existing
	// entry with that ID. Elements without ID will create new entries. Any entry not
	// in the list will be deleted.
	Entries param.Field[[]AccountDlpProfileCustomUpdateParamsEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
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
