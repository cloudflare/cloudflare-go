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

// DLPProfileCustomService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDLPProfileCustomService] method
// instead.
type DLPProfileCustomService struct {
	Options []option.RequestOption
}

// NewDLPProfileCustomService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfileCustomService(opts ...option.RequestOption) (r *DLPProfileCustomService) {
	r = &DLPProfileCustomService{}
	r.Options = opts
	return
}

// Updates a DLP custom profile.
func (r *DLPProfileCustomService) Update(ctx context.Context, accountID string, profileID string, body DLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *DLPProfileCustomUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a DLP custom profile.
func (r *DLPProfileCustomService) Delete(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *DLPProfileCustomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a set of DLP custom profiles.
func (r *DLPProfileCustomService) DLPProfilesNewCustomProfiles(ctx context.Context, accountID string, body DLPProfileCustomDLPProfilesNewCustomProfilesParams, opts ...option.RequestOption) (res *[]DLPProfileCustomDLPProfilesNewCustomProfilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile.
func (r *DLPProfileCustomService) Get(ctx context.Context, accountID string, profileID string, opts ...option.RequestOption) (res *DLPProfileCustomGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", accountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPProfileCustomUpdateResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileCustomUpdateResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileCustomUpdateResponseType `json:"type"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponse]
type dlpProfileCustomUpdateResponseJSON struct {
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

func (r *DLPProfileCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A custom entry that matches a profile
type DLPProfileCustomUpdateResponseEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern DLPProfileCustomUpdateResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                             `json:"profile_id"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEntryJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseEntry]
type dlpProfileCustomUpdateResponseEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type DLPProfileCustomUpdateResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation DLPProfileCustomUpdateResponseEntriesPatternValidation `json:"validation"`
	JSON       dlpProfileCustomUpdateResponseEntriesPatternJSON       `json:"-"`
}

// dlpProfileCustomUpdateResponseEntriesPatternJSON contains the JSON metadata for
// the struct [DLPProfileCustomUpdateResponseEntriesPattern]
type dlpProfileCustomUpdateResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileCustomUpdateResponseEntriesPatternValidation string

const (
	DLPProfileCustomUpdateResponseEntriesPatternValidationLuhn DLPProfileCustomUpdateResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type DLPProfileCustomUpdateResponseType string

const (
	DLPProfileCustomUpdateResponseTypeCustom DLPProfileCustomUpdateResponseType = "custom"
)

// Union satisfied by [DLPProfileCustomDeleteResponseUnknown] or
// [shared.UnionString].
type DLPProfileCustomDeleteResponse interface {
	ImplementsDLPProfileCustomDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DLPProfileCustomDLPProfilesNewCustomProfilesResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileCustomDLPProfilesNewCustomProfilesResponseType `json:"type"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomDLPProfilesNewCustomProfilesResponseJSON `json:"-"`
}

// dlpProfileCustomDLPProfilesNewCustomProfilesResponseJSON contains the JSON
// metadata for the struct [DLPProfileCustomDLPProfilesNewCustomProfilesResponse]
type dlpProfileCustomDLPProfilesNewCustomProfilesResponseJSON struct {
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

func (r *DLPProfileCustomDLPProfilesNewCustomProfilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A custom entry that matches a profile
type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                                   `json:"profile_id"`
	UpdatedAt time.Time                                                     `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomDLPProfilesNewCustomProfilesResponseEntryJSON `json:"-"`
}

// dlpProfileCustomDLPProfilesNewCustomProfilesResponseEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntry]
type dlpProfileCustomDLPProfilesNewCustomProfilesResponseEntryJSON struct {
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

func (r *DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPatternValidation `json:"validation"`
	JSON       dlpProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPatternJSON       `json:"-"`
}

// dlpProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPatternJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPattern]
type dlpProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPatternValidation string

const (
	DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPatternValidationLuhn DLPProfileCustomDLPProfilesNewCustomProfilesResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type DLPProfileCustomDLPProfilesNewCustomProfilesResponseType string

const (
	DLPProfileCustomDLPProfilesNewCustomProfilesResponseTypeCustom DLPProfileCustomDLPProfilesNewCustomProfilesResponseType = "custom"
)

type DLPProfileCustomGetResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64   `json:"allowed_match_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileCustomGetResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileCustomGetResponseType `json:"type"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomGetResponseJSON `json:"-"`
}

// dlpProfileCustomGetResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomGetResponse]
type dlpProfileCustomGetResponseJSON struct {
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

func (r *DLPProfileCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A custom entry that matches a profile
type DLPProfileCustomGetResponseEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern DLPProfileCustomGetResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                          `json:"profile_id"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomGetResponseEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseEntryJSON contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseEntry]
type dlpProfileCustomGetResponseEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A pattern that matches an entry
type DLPProfileCustomGetResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation DLPProfileCustomGetResponseEntriesPatternValidation `json:"validation"`
	JSON       dlpProfileCustomGetResponseEntriesPatternJSON       `json:"-"`
}

// dlpProfileCustomGetResponseEntriesPatternJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEntriesPattern]
type dlpProfileCustomGetResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileCustomGetResponseEntriesPatternValidation string

const (
	DLPProfileCustomGetResponseEntriesPatternValidationLuhn DLPProfileCustomGetResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type DLPProfileCustomGetResponseType string

const (
	DLPProfileCustomGetResponseTypeCustom DLPProfileCustomGetResponseType = "custom"
)

type DLPProfileCustomUpdateParams struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The custom entries for this profile. Array elements with IDs are modifying the
	// existing entry with that ID. Elements without ID will create new entries. Any
	// entry not in the list will be deleted.
	Entries param.Field[[]DLPProfileCustomUpdateParamsEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomUpdateParamsSharedEntry] `json:"shared_entries"`
}

func (r DLPProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom entry that matches a profile
type DLPProfileCustomUpdateParamsEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
	// The name of the entry.
	Name param.Field[string] `json:"name"`
	// A pattern that matches an entry
	Pattern param.Field[DLPProfileCustomUpdateParamsEntriesPattern] `json:"pattern"`
	// ID of the parent profile
	ProfileID param.Field[interface{}] `json:"profile_id"`
}

func (r DLPProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A pattern that matches an entry
type DLPProfileCustomUpdateParamsEntriesPattern struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation param.Field[DLPProfileCustomUpdateParamsEntriesPatternValidation] `json:"validation"`
}

func (r DLPProfileCustomUpdateParamsEntriesPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileCustomUpdateParamsEntriesPatternValidation string

const (
	DLPProfileCustomUpdateParamsEntriesPatternValidationLuhn DLPProfileCustomUpdateParamsEntriesPatternValidation = "luhn"
)

// Properties of a predefined entry in a custom profile
//
// Satisfied by
// [DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined],
// [DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration].
type DLPProfileCustomUpdateParamsSharedEntry interface {
	implementsDLPProfileCustomUpdateParamsSharedEntry()
}

// Properties of a predefined entry in a custom profile
type DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined) implementsDLPProfileCustomUpdateParamsSharedEntry() {
}

// Properties of an integration entry in a custom profile
type DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration) implementsDLPProfileCustomUpdateParamsSharedEntry() {
}

type DLPProfileCustomDeleteResponseEnvelope struct {
	Errors   []DLPProfileCustomDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPProfileCustomDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPProfileCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomDeleteResponseEnvelope]
type dlpProfileCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfileCustomDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfileCustomDeleteResponseEnvelopeErrors]
type dlpProfileCustomDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfileCustomDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfileCustomDeleteResponseEnvelopeMessages]
type dlpProfileCustomDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPProfileCustomDeleteResponseEnvelopeSuccess bool

const (
	DLPProfileCustomDeleteResponseEnvelopeSuccessTrue DLPProfileCustomDeleteResponseEnvelopeSuccess = true
)

type DLPProfileCustomDLPProfilesNewCustomProfilesParams struct {
	Profiles param.Field[[]DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfile] `json:"profiles,required"`
}

func (r DLPProfileCustomDLPProfilesNewCustomProfilesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfile struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The entries for this profile.
	Entries param.Field[[]DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
}

func (r DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom entry create payload
type DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The name of the entry.
	Name param.Field[string] `json:"name,required"`
	// A pattern that matches an entry
	Pattern param.Field[DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntriesPattern] `json:"pattern,required"`
}

func (r DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A pattern that matches an entry
type DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntriesPattern struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation param.Field[DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntriesPatternValidation] `json:"validation"`
}

func (r DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntriesPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntriesPatternValidation string

const (
	DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntriesPatternValidationLuhn DLPProfileCustomDLPProfilesNewCustomProfilesParamsProfilesEntriesPatternValidation = "luhn"
)

type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelope struct {
	Errors   []DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeMessages `json:"messages,required"`
	Result   []DLPProfileCustomDLPProfilesNewCustomProfilesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelope]
type dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeErrors]
type dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeMessages]
type dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeSuccess bool

const (
	DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeSuccessTrue DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeSuccess = true
)

type DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                    `json:"total_count"`
	JSON       dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeResultInfo]
type dlpProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDLPProfilesNewCustomProfilesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfileCustomGetResponseEnvelope struct {
	Errors   []DLPProfileCustomGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPProfileCustomGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPProfileCustomGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpProfileCustomGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelope]
type dlpProfileCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfileCustomGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    dlpProfileCustomGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelopeErrors]
type dlpProfileCustomGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DLPProfileCustomGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dlpProfileCustomGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseEnvelopeMessages]
type dlpProfileCustomGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DLPProfileCustomGetResponseEnvelopeSuccess bool

const (
	DLPProfileCustomGetResponseEnvelopeSuccessTrue DLPProfileCustomGetResponseEnvelopeSuccess = true
)
