// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *DLPProfileCustomService) New(ctx context.Context, params DLPProfileCustomNewParams, opts ...option.RequestOption) (res *[]DLPProfileCustomNewResponse, err error) {
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
func (r *DLPProfileCustomService) Update(ctx context.Context, profileID string, params DLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *DLPProfileCustomUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Deletes a DLP custom profile.
func (r *DLPProfileCustomService) Delete(ctx context.Context, profileID string, body DLPProfileCustomDeleteParams, opts ...option.RequestOption) (res *DLPProfileCustomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile.
func (r *DLPProfileCustomService) Get(ctx context.Context, profileID string, query DLPProfileCustomGetParams, opts ...option.RequestOption) (res *DLPProfileCustomGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileCustomGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPProfileCustomNewResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfileCustomNewResponseContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                   `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileCustomNewResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileCustomNewResponseType `json:"type"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomNewResponseJSON `json:"-"`
}

// dlpProfileCustomNewResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomNewResponse]
type dlpProfileCustomNewResponseJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	CreatedAt         apijson.Field
	Description       apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileCustomNewResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfileCustomNewResponseContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfileCustomNewResponseContextAwarenessJSON `json:"-"`
}

// dlpProfileCustomNewResponseContextAwarenessJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseContextAwareness]
type dlpProfileCustomNewResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileCustomNewResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                `json:"files,required"`
	JSON  dlpProfileCustomNewResponseContextAwarenessSkipJSON `json:"-"`
}

// dlpProfileCustomNewResponseContextAwarenessSkipJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseContextAwarenessSkip]
type dlpProfileCustomNewResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type DLPProfileCustomNewResponseEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern DLPProfileCustomNewResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                          `json:"profile_id"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomNewResponseEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseEntryJSON contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseEntry]
type dlpProfileCustomNewResponseEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type DLPProfileCustomNewResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation DLPProfileCustomNewResponseEntriesPatternValidation `json:"validation"`
	JSON       dlpProfileCustomNewResponseEntriesPatternJSON       `json:"-"`
}

// dlpProfileCustomNewResponseEntriesPatternJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEntriesPattern]
type dlpProfileCustomNewResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileCustomNewResponseEntriesPatternValidation string

const (
	DLPProfileCustomNewResponseEntriesPatternValidationLuhn DLPProfileCustomNewResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type DLPProfileCustomNewResponseType string

const (
	DLPProfileCustomNewResponseTypeCustom DLPProfileCustomNewResponseType = "custom"
)

type DLPProfileCustomUpdateResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfileCustomUpdateResponseContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                      `json:"created_at" format:"date-time"`
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
	ContextAwareness  apijson.Field
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

func (r dlpProfileCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileCustomUpdateResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfileCustomUpdateResponseContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfileCustomUpdateResponseContextAwarenessJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseContextAwarenessJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseContextAwareness]
type dlpProfileCustomUpdateResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileCustomUpdateResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                   `json:"files,required"`
	JSON  dlpProfileCustomUpdateResponseContextAwarenessSkipJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseContextAwarenessSkipJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseContextAwarenessSkip]
type dlpProfileCustomUpdateResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomUpdateResponseEntryJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomUpdateResponseEntriesPatternJSON) RawJSON() string {
	return r.raw
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

type DLPProfileCustomGetResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfileCustomGetResponseContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                   `json:"created_at" format:"date-time"`
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
	ContextAwareness  apijson.Field
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

func (r dlpProfileCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileCustomGetResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfileCustomGetResponseContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfileCustomGetResponseContextAwarenessJSON `json:"-"`
}

// dlpProfileCustomGetResponseContextAwarenessJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseContextAwareness]
type dlpProfileCustomGetResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileCustomGetResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                `json:"files,required"`
	JSON  dlpProfileCustomGetResponseContextAwarenessSkipJSON `json:"-"`
}

// dlpProfileCustomGetResponseContextAwarenessSkipJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseContextAwarenessSkip]
type dlpProfileCustomGetResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomGetResponseEntryJSON) RawJSON() string {
	return r.raw
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

func (r dlpProfileCustomGetResponseEntriesPatternJSON) RawJSON() string {
	return r.raw
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

type DLPProfileCustomNewResponseEnvelope struct {
	Errors   []DLPProfileCustomNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []DLPProfileCustomNewResponse                 `json:"result,required,nullable"`
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
	AccountID param.Field[string] `path:"account_id,required"`
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

type DLPProfileCustomGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
