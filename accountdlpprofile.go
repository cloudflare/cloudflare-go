// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDlpProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDlpProfileService] method
// instead.
type AccountDlpProfileService struct {
	Options     []option.RequestOption
	Customs     *AccountDlpProfileCustomService
	Predefineds *AccountDlpProfilePredefinedService
}

// NewAccountDlpProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpProfileService(opts ...option.RequestOption) (r *AccountDlpProfileService) {
	r = &AccountDlpProfileService{}
	r.Options = opts
	r.Customs = NewAccountDlpProfileCustomService(opts...)
	r.Predefineds = NewAccountDlpProfilePredefinedService(opts...)
	return
}

// Fetches a DLP profile by ID. Supports both predefined and custom profiles
func (r *AccountDlpProfileService) Get(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *AccountDlpProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *AccountDlpProfileService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountDlpProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDlpProfileGetResponse struct {
	Errors   []AccountDlpProfileGetResponseError   `json:"errors"`
	Messages []AccountDlpProfileGetResponseMessage `json:"messages"`
	Result   AccountDlpProfileGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDlpProfileGetResponseSuccess `json:"success"`
	JSON    accountDlpProfileGetResponseJSON    `json:"-"`
}

// accountDlpProfileGetResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileGetResponse]
type accountDlpProfileGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountDlpProfileGetResponseErrorJSON `json:"-"`
}

// accountDlpProfileGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountDlpProfileGetResponseError]
type accountDlpProfileGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountDlpProfileGetResponseMessageJSON `json:"-"`
}

// accountDlpProfileGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDlpProfileGetResponseMessage]
type accountDlpProfileGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountDlpProfileGetResponseResultDlpPredefinedProfile],
// [AccountDlpProfileGetResponseResultDlpCustomProfile] or
// [AccountDlpProfileGetResponseResultDlpIntegrationProfile].
type AccountDlpProfileGetResponseResult interface {
	implementsAccountDlpProfileGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountDlpProfileGetResponseResult)(nil)).Elem(), "")
}

type AccountDlpProfileGetResponseResultDlpPredefinedProfile struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// The entries for this profile.
	Entries []AccountDlpProfileGetResponseResultDlpPredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type AccountDlpProfileGetResponseResultDlpPredefinedProfileType `json:"type"`
	JSON accountDlpProfileGetResponseResultDlpPredefinedProfileJSON `json:"-"`
}

// accountDlpProfileGetResponseResultDlpPredefinedProfileJSON contains the JSON
// metadata for the struct [AccountDlpProfileGetResponseResultDlpPredefinedProfile]
type accountDlpProfileGetResponseResultDlpPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponseResultDlpPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDlpProfileGetResponseResultDlpPredefinedProfile) implementsAccountDlpProfileGetResponseResult() {
}

// A predefined entry that matches a profile
type AccountDlpProfileGetResponseResultDlpPredefinedProfileEntry struct {
	// UUID
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                     `json:"profile_id"`
	JSON      accountDlpProfileGetResponseResultDlpPredefinedProfileEntryJSON `json:"-"`
}

// accountDlpProfileGetResponseResultDlpPredefinedProfileEntryJSON contains the
// JSON metadata for the struct
// [AccountDlpProfileGetResponseResultDlpPredefinedProfileEntry]
type accountDlpProfileGetResponseResultDlpPredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponseResultDlpPredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type AccountDlpProfileGetResponseResultDlpPredefinedProfileType string

const (
	AccountDlpProfileGetResponseResultDlpPredefinedProfileTypePredefined AccountDlpProfileGetResponseResultDlpPredefinedProfileType = "predefined"
)

type AccountDlpProfileGetResponseResultDlpCustomProfile struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []AccountDlpProfileGetResponseResultDlpCustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      AccountDlpProfileGetResponseResultDlpCustomProfileType `json:"type"`
	UpdatedAt time.Time                                              `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileGetResponseResultDlpCustomProfileJSON `json:"-"`
}

// accountDlpProfileGetResponseResultDlpCustomProfileJSON contains the JSON
// metadata for the struct [AccountDlpProfileGetResponseResultDlpCustomProfile]
type accountDlpProfileGetResponseResultDlpCustomProfileJSON struct {
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

func (r *AccountDlpProfileGetResponseResultDlpCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDlpProfileGetResponseResultDlpCustomProfile) implementsAccountDlpProfileGetResponseResult() {
}

// A custom entry that matches a profile
type AccountDlpProfileGetResponseResultDlpCustomProfileEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                                 `json:"profile_id"`
	UpdatedAt time.Time                                                   `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileGetResponseResultDlpCustomProfileEntryJSON `json:"-"`
}

// accountDlpProfileGetResponseResultDlpCustomProfileEntryJSON contains the JSON
// metadata for the struct
// [AccountDlpProfileGetResponseResultDlpCustomProfileEntry]
type accountDlpProfileGetResponseResultDlpCustomProfileEntryJSON struct {
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

func (r *AccountDlpProfileGetResponseResultDlpCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPatternValidation `json:"validation"`
	JSON       accountDlpProfileGetResponseResultDlpCustomProfileEntriesPatternJSON       `json:"-"`
}

// accountDlpProfileGetResponseResultDlpCustomProfileEntriesPatternJSON contains
// the JSON metadata for the struct
// [AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPattern]
type accountDlpProfileGetResponseResultDlpCustomProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPatternValidation string

const (
	AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPatternValidationLuhn AccountDlpProfileGetResponseResultDlpCustomProfileEntriesPatternValidation = "luhn"
)

// The type of the profile.
type AccountDlpProfileGetResponseResultDlpCustomProfileType string

const (
	AccountDlpProfileGetResponseResultDlpCustomProfileTypeCustom AccountDlpProfileGetResponseResultDlpCustomProfileType = "custom"
)

type AccountDlpProfileGetResponseResultDlpIntegrationProfile struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []AccountDlpProfileGetResponseResultDlpIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      AccountDlpProfileGetResponseResultDlpIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                                   `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileGetResponseResultDlpIntegrationProfileJSON `json:"-"`
}

// accountDlpProfileGetResponseResultDlpIntegrationProfileJSON contains the JSON
// metadata for the struct
// [AccountDlpProfileGetResponseResultDlpIntegrationProfile]
type accountDlpProfileGetResponseResultDlpIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponseResultDlpIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDlpProfileGetResponseResultDlpIntegrationProfile) implementsAccountDlpProfileGetResponseResult() {
}

// An entry derived from an integration
type AccountDlpProfileGetResponseResultDlpIntegrationProfileEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                      `json:"profile_id"`
	UpdatedAt time.Time                                                        `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileGetResponseResultDlpIntegrationProfileEntryJSON `json:"-"`
}

// accountDlpProfileGetResponseResultDlpIntegrationProfileEntryJSON contains the
// JSON metadata for the struct
// [AccountDlpProfileGetResponseResultDlpIntegrationProfileEntry]
type accountDlpProfileGetResponseResultDlpIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileGetResponseResultDlpIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type AccountDlpProfileGetResponseResultDlpIntegrationProfileType string

const (
	AccountDlpProfileGetResponseResultDlpIntegrationProfileTypeIntegration AccountDlpProfileGetResponseResultDlpIntegrationProfileType = "integration"
)

// Whether the API call was successful
type AccountDlpProfileGetResponseSuccess bool

const (
	AccountDlpProfileGetResponseSuccessTrue AccountDlpProfileGetResponseSuccess = true
)

type AccountDlpProfileListResponse struct {
	Errors     []AccountDlpProfileListResponseError    `json:"errors"`
	Messages   []AccountDlpProfileListResponseMessage  `json:"messages"`
	Result     []AccountDlpProfileListResponseResult   `json:"result"`
	ResultInfo AccountDlpProfileListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountDlpProfileListResponseSuccess `json:"success"`
	JSON    accountDlpProfileListResponseJSON    `json:"-"`
}

// accountDlpProfileListResponseJSON contains the JSON metadata for the struct
// [AccountDlpProfileListResponse]
type accountDlpProfileListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountDlpProfileListResponseErrorJSON `json:"-"`
}

// accountDlpProfileListResponseErrorJSON contains the JSON metadata for the struct
// [AccountDlpProfileListResponseError]
type accountDlpProfileListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDlpProfileListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountDlpProfileListResponseMessageJSON `json:"-"`
}

// accountDlpProfileListResponseMessageJSON contains the JSON metadata for the
// struct [AccountDlpProfileListResponseMessage]
type accountDlpProfileListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccountDlpProfileListResponseResultDlpPredefinedProfile],
// [AccountDlpProfileListResponseResultDlpCustomProfile] or
// [AccountDlpProfileListResponseResultDlpIntegrationProfile].
type AccountDlpProfileListResponseResult interface {
	implementsAccountDlpProfileListResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountDlpProfileListResponseResult)(nil)).Elem(), "")
}

type AccountDlpProfileListResponseResultDlpPredefinedProfile struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// The entries for this profile.
	Entries []AccountDlpProfileListResponseResultDlpPredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type AccountDlpProfileListResponseResultDlpPredefinedProfileType `json:"type"`
	JSON accountDlpProfileListResponseResultDlpPredefinedProfileJSON `json:"-"`
}

// accountDlpProfileListResponseResultDlpPredefinedProfileJSON contains the JSON
// metadata for the struct
// [AccountDlpProfileListResponseResultDlpPredefinedProfile]
type accountDlpProfileListResponseResultDlpPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseResultDlpPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDlpProfileListResponseResultDlpPredefinedProfile) implementsAccountDlpProfileListResponseResult() {
}

// A predefined entry that matches a profile
type AccountDlpProfileListResponseResultDlpPredefinedProfileEntry struct {
	// UUID
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                      `json:"profile_id"`
	JSON      accountDlpProfileListResponseResultDlpPredefinedProfileEntryJSON `json:"-"`
}

// accountDlpProfileListResponseResultDlpPredefinedProfileEntryJSON contains the
// JSON metadata for the struct
// [AccountDlpProfileListResponseResultDlpPredefinedProfileEntry]
type accountDlpProfileListResponseResultDlpPredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseResultDlpPredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type AccountDlpProfileListResponseResultDlpPredefinedProfileType string

const (
	AccountDlpProfileListResponseResultDlpPredefinedProfileTypePredefined AccountDlpProfileListResponseResultDlpPredefinedProfileType = "predefined"
)

type AccountDlpProfileListResponseResultDlpCustomProfile struct {
	// UUID
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []AccountDlpProfileListResponseResultDlpCustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      AccountDlpProfileListResponseResultDlpCustomProfileType `json:"type"`
	UpdatedAt time.Time                                               `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileListResponseResultDlpCustomProfileJSON `json:"-"`
}

// accountDlpProfileListResponseResultDlpCustomProfileJSON contains the JSON
// metadata for the struct [AccountDlpProfileListResponseResultDlpCustomProfile]
type accountDlpProfileListResponseResultDlpCustomProfileJSON struct {
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

func (r *AccountDlpProfileListResponseResultDlpCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDlpProfileListResponseResultDlpCustomProfile) implementsAccountDlpProfileListResponseResult() {
}

// A custom entry that matches a profile
type AccountDlpProfileListResponseResultDlpCustomProfileEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern AccountDlpProfileListResponseResultDlpCustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                                  `json:"profile_id"`
	UpdatedAt time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileListResponseResultDlpCustomProfileEntryJSON `json:"-"`
}

// accountDlpProfileListResponseResultDlpCustomProfileEntryJSON contains the JSON
// metadata for the struct
// [AccountDlpProfileListResponseResultDlpCustomProfileEntry]
type accountDlpProfileListResponseResultDlpCustomProfileEntryJSON struct {
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

func (r *AccountDlpProfileListResponseResultDlpCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type AccountDlpProfileListResponseResultDlpCustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation AccountDlpProfileListResponseResultDlpCustomProfileEntriesPatternValidation `json:"validation"`
	JSON       accountDlpProfileListResponseResultDlpCustomProfileEntriesPatternJSON       `json:"-"`
}

// accountDlpProfileListResponseResultDlpCustomProfileEntriesPatternJSON contains
// the JSON metadata for the struct
// [AccountDlpProfileListResponseResultDlpCustomProfileEntriesPattern]
type accountDlpProfileListResponseResultDlpCustomProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseResultDlpCustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type AccountDlpProfileListResponseResultDlpCustomProfileEntriesPatternValidation string

const (
	AccountDlpProfileListResponseResultDlpCustomProfileEntriesPatternValidationLuhn AccountDlpProfileListResponseResultDlpCustomProfileEntriesPatternValidation = "luhn"
)

// The type of the profile.
type AccountDlpProfileListResponseResultDlpCustomProfileType string

const (
	AccountDlpProfileListResponseResultDlpCustomProfileTypeCustom AccountDlpProfileListResponseResultDlpCustomProfileType = "custom"
)

type AccountDlpProfileListResponseResultDlpIntegrationProfile struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []AccountDlpProfileListResponseResultDlpIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      AccountDlpProfileListResponseResultDlpIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileListResponseResultDlpIntegrationProfileJSON `json:"-"`
}

// accountDlpProfileListResponseResultDlpIntegrationProfileJSON contains the JSON
// metadata for the struct
// [AccountDlpProfileListResponseResultDlpIntegrationProfile]
type accountDlpProfileListResponseResultDlpIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseResultDlpIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountDlpProfileListResponseResultDlpIntegrationProfile) implementsAccountDlpProfileListResponseResult() {
}

// An entry derived from an integration
type AccountDlpProfileListResponseResultDlpIntegrationProfileEntry struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                       `json:"profile_id"`
	UpdatedAt time.Time                                                         `json:"updated_at" format:"date-time"`
	JSON      accountDlpProfileListResponseResultDlpIntegrationProfileEntryJSON `json:"-"`
}

// accountDlpProfileListResponseResultDlpIntegrationProfileEntryJSON contains the
// JSON metadata for the struct
// [AccountDlpProfileListResponseResultDlpIntegrationProfileEntry]
type accountDlpProfileListResponseResultDlpIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseResultDlpIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type AccountDlpProfileListResponseResultDlpIntegrationProfileType string

const (
	AccountDlpProfileListResponseResultDlpIntegrationProfileTypeIntegration AccountDlpProfileListResponseResultDlpIntegrationProfileType = "integration"
)

type AccountDlpProfileListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       accountDlpProfileListResponseResultInfoJSON `json:"-"`
}

// accountDlpProfileListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountDlpProfileListResponseResultInfo]
type accountDlpProfileListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDlpProfileListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDlpProfileListResponseSuccess bool

const (
	AccountDlpProfileListResponseSuccessTrue AccountDlpProfileListResponseSuccess = true
)
