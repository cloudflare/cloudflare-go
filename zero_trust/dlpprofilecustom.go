// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Creates a set of DLP custom profiles.
func (r *DLPProfileCustomService) New(ctx context.Context, params DLPProfileCustomNewParams, opts ...option.RequestOption) (res *[]DLPCustomProfile, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP custom profile.
func (r *DLPProfileCustomService) Update(ctx context.Context, profileID string, params DLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *DLPCustomProfile, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Deletes a DLP custom profile.
func (r *DLPProfileCustomService) Delete(ctx context.Context, profileID string, params DLPProfileCustomDeleteParams, opts ...option.RequestOption) (res *DLPProfileCustomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile.
func (r *DLPProfileCustomService) Get(ctx context.Context, profileID string, query DLPProfileCustomGetParams, opts ...option.RequestOption) (res *DLPCustomProfile, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPCustomProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPCustomProfileContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                        `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPCustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled bool `json:"ocr_enabled"`
	// The type of the profile.
	Type      DLPCustomProfileType `json:"type"`
	UpdatedAt time.Time            `json:"updated_at" format:"date-time"`
	JSON      dlpCustomProfileJSON `json:"-"`
}

// dlpCustomProfileJSON contains the JSON metadata for the struct
// [DLPCustomProfile]
type dlpCustomProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	OCREnabled        apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPCustomProfile) implementsZeroTrustDLPProfiles() {}

func (r DLPCustomProfile) implementsZeroTrustDLPProfileGetResponse() {}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPCustomProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPCustomProfileContextAwarenessSkip `json:"skip,required"`
	JSON dlpCustomProfileContextAwarenessJSON `json:"-"`
}

// dlpCustomProfileContextAwarenessJSON contains the JSON metadata for the struct
// [DLPCustomProfileContextAwareness]
type dlpCustomProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPCustomProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                     `json:"files,required"`
	JSON  dlpCustomProfileContextAwarenessSkipJSON `json:"-"`
}

// dlpCustomProfileContextAwarenessSkipJSON contains the JSON metadata for the
// struct [DLPCustomProfileContextAwarenessSkip]
type dlpCustomProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type DLPCustomProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern DLPCustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}               `json:"profile_id"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      dlpCustomProfileEntryJSON `json:"-"`
}

// dlpCustomProfileEntryJSON contains the JSON metadata for the struct
// [DLPCustomProfileEntry]
type dlpCustomProfileEntryJSON struct {
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

func (r *DLPCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type DLPCustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation DLPCustomProfileEntriesPatternValidation `json:"validation"`
	JSON       dlpCustomProfileEntriesPatternJSON       `json:"-"`
}

// dlpCustomProfileEntriesPatternJSON contains the JSON metadata for the struct
// [DLPCustomProfileEntriesPattern]
type dlpCustomProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomProfileEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPCustomProfileEntriesPatternValidation string

const (
	DLPCustomProfileEntriesPatternValidationLuhn DLPCustomProfileEntriesPatternValidation = "luhn"
)

func (r DLPCustomProfileEntriesPatternValidation) IsKnown() bool {
	switch r {
	case DLPCustomProfileEntriesPatternValidationLuhn:
		return true
	}
	return false
}

// The type of the profile.
type DLPCustomProfileType string

const (
	DLPCustomProfileTypeCustom DLPCustomProfileType = "custom"
)

func (r DLPCustomProfileType) IsKnown() bool {
	switch r {
	case DLPCustomProfileTypeCustom:
		return true
	}
	return false
}

// Union satisfied by [zero_trust.DLPProfileCustomDeleteResponseUnknown] or
// [shared.UnionString].
type DLPProfileCustomDeleteResponse interface {
	ImplementsZeroTrustDLPProfileCustomDeleteResponse()
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

type DLPProfileCustomNewParams struct {
	// Identifier
	AccountID param.Field[string]                             `path:"account_id,required"`
	Profiles  param.Field[[]DLPProfileCustomNewParamsProfile] `json:"profiles,required"`
}

func (r DLPProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewParamsProfile struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[DLPProfileCustomNewParamsProfilesContextAwareness] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The entries for this profile.
	Entries param.Field[[]DLPProfileCustomNewParamsProfilesEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled param.Field[bool] `json:"ocr_enabled"`
}

func (r DLPProfileCustomNewParamsProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileCustomNewParamsProfilesContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[DLPProfileCustomNewParamsProfilesContextAwarenessSkip] `json:"skip,required"`
}

func (r DLPProfileCustomNewParamsProfilesContextAwareness) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileCustomNewParamsProfilesContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r DLPProfileCustomNewParamsProfilesContextAwarenessSkip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom entry create payload
type DLPProfileCustomNewParamsProfilesEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The name of the entry.
	Name param.Field[string] `json:"name,required"`
	// A pattern that matches an entry
	Pattern param.Field[DLPProfileCustomNewParamsProfilesEntriesPattern] `json:"pattern,required"`
}

func (r DLPProfileCustomNewParamsProfilesEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A pattern that matches an entry
type DLPProfileCustomNewParamsProfilesEntriesPattern struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation param.Field[DLPProfileCustomNewParamsProfilesEntriesPatternValidation] `json:"validation"`
}

func (r DLPProfileCustomNewParamsProfilesEntriesPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileCustomNewParamsProfilesEntriesPatternValidation string

const (
	DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn DLPProfileCustomNewParamsProfilesEntriesPatternValidation = "luhn"
)

func (r DLPProfileCustomNewParamsProfilesEntriesPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseEnvelope struct {
	Errors   []DLPProfileCustomNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []DLPCustomProfile                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DLPProfileCustomNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DLPProfileCustomNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileCustomNewResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelope]
type dlpProfileCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    dlpProfileCustomNewResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelopeErrors]
type dlpProfileCustomNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dlpProfileCustomNewResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseEnvelopeMessages]
type dlpProfileCustomNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileCustomNewResponseEnvelopeSuccess bool

const (
	DLPProfileCustomNewResponseEnvelopeSuccessTrue DLPProfileCustomNewResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       dlpProfileCustomNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseEnvelopeResultInfo]
type dlpProfileCustomNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[DLPProfileCustomUpdateParamsContextAwareness] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The custom entries for this profile. Array elements with IDs are modifying the
	// existing entry with that ID. Elements without ID will create new entries. Any
	// entry not in the list will be deleted.
	Entries param.Field[[]DLPProfileCustomUpdateParamsEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
	// If true, scan images via OCR to determine if any text present matches filters.
	OCREnabled param.Field[bool] `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomUpdateParamsSharedEntry] `json:"shared_entries"`
}

func (r DLPProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileCustomUpdateParamsContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[DLPProfileCustomUpdateParamsContextAwarenessSkip] `json:"skip,required"`
}

func (r DLPProfileCustomUpdateParamsContextAwareness) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileCustomUpdateParamsContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r DLPProfileCustomUpdateParamsContextAwarenessSkip) MarshalJSON() (data []byte, err error) {
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

func (r DLPProfileCustomUpdateParamsEntriesPatternValidation) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateParamsEntriesPatternValidationLuhn:
		return true
	}
	return false
}

// Properties of a predefined entry in a custom profile
//
// Satisfied by
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined],
// [zero_trust.DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration].
type DLPProfileCustomUpdateParamsSharedEntry interface {
	implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntry()
}

// Properties of a predefined entry in a custom profile
type DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntry() {
}

// Properties of an integration entry in a custom profile
type DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntry() {
}

type DLPProfileCustomDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r DLPProfileCustomDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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

func (r dlpProfileCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileCustomDeleteResponseEnvelopeSuccess bool

const (
	DLPProfileCustomDeleteResponseEnvelopeSuccessTrue DLPProfileCustomDeleteResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomGetResponseEnvelope struct {
	Errors   []DLPProfileCustomGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPCustomProfile                              `json:"result,required"`
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

func (r dlpProfileCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileCustomGetResponseEnvelopeSuccess bool

const (
	DLPProfileCustomGetResponseEnvelopeSuccessTrue DLPProfileCustomGetResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
