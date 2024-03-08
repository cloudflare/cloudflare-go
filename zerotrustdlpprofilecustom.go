// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// ZeroTrustDLPProfileCustomService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustDLPProfileCustomService] method instead.
type ZeroTrustDLPProfileCustomService struct {
	Options []option.RequestOption
}

// NewZeroTrustDLPProfileCustomService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDLPProfileCustomService(opts ...option.RequestOption) (r *ZeroTrustDLPProfileCustomService) {
	r = &ZeroTrustDLPProfileCustomService{}
	r.Options = opts
	return
}

// Creates a set of DLP custom profiles.
func (r *ZeroTrustDLPProfileCustomService) New(ctx context.Context, params ZeroTrustDLPProfileCustomNewParams, opts ...option.RequestOption) (res *[]ZeroTrustDLPProfileCustomNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfileCustomNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP custom profile.
func (r *ZeroTrustDLPProfileCustomService) Update(ctx context.Context, profileID string, params ZeroTrustDLPProfileCustomUpdateParams, opts ...option.RequestOption) (res *ZeroTrustDLPProfileCustomUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Deletes a DLP custom profile.
func (r *ZeroTrustDLPProfileCustomService) Delete(ctx context.Context, profileID string, body ZeroTrustDLPProfileCustomDeleteParams, opts ...option.RequestOption) (res *ZeroTrustDLPProfileCustomDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfileCustomDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile.
func (r *ZeroTrustDLPProfileCustomService) Get(ctx context.Context, profileID string, query ZeroTrustDLPProfileCustomGetParams, opts ...option.RequestOption) (res *ZeroTrustDLPProfileCustomGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfileCustomGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDLPProfileCustomNewResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ZeroTrustDLPProfileCustomNewResponseContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                            `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileCustomNewResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileCustomNewResponseType `json:"type"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileCustomNewResponseJSON `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileCustomNewResponse]
type zeroTrustDLPProfileCustomNewResponseJSON struct {
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

func (r *ZeroTrustDLPProfileCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileCustomNewResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ZeroTrustDLPProfileCustomNewResponseContextAwarenessSkip `json:"skip,required"`
	JSON zeroTrustDLPProfileCustomNewResponseContextAwarenessJSON `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseContextAwarenessJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomNewResponseContextAwareness]
type zeroTrustDLPProfileCustomNewResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomNewResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileCustomNewResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                         `json:"files,required"`
	JSON  zeroTrustDLPProfileCustomNewResponseContextAwarenessSkipJSON `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseContextAwarenessSkipJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileCustomNewResponseContextAwarenessSkip]
type zeroTrustDLPProfileCustomNewResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomNewResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type ZeroTrustDLPProfileCustomNewResponseEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern ZeroTrustDLPProfileCustomNewResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                   `json:"profile_id"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileCustomNewResponseEntryJSON `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseEntryJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileCustomNewResponseEntry]
type zeroTrustDLPProfileCustomNewResponseEntryJSON struct {
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

func (r *ZeroTrustDLPProfileCustomNewResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type ZeroTrustDLPProfileCustomNewResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation ZeroTrustDLPProfileCustomNewResponseEntriesPatternValidation `json:"validation"`
	JSON       zeroTrustDLPProfileCustomNewResponseEntriesPatternJSON       `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseEntriesPatternJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomNewResponseEntriesPattern]
type zeroTrustDLPProfileCustomNewResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomNewResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type ZeroTrustDLPProfileCustomNewResponseEntriesPatternValidation string

const (
	ZeroTrustDLPProfileCustomNewResponseEntriesPatternValidationLuhn ZeroTrustDLPProfileCustomNewResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type ZeroTrustDLPProfileCustomNewResponseType string

const (
	ZeroTrustDLPProfileCustomNewResponseTypeCustom ZeroTrustDLPProfileCustomNewResponseType = "custom"
)

type ZeroTrustDLPProfileCustomUpdateResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ZeroTrustDLPProfileCustomUpdateResponseContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                               `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileCustomUpdateResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileCustomUpdateResponseType `json:"type"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileCustomUpdateResponseJSON `json:"-"`
}

// zeroTrustDLPProfileCustomUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileCustomUpdateResponse]
type zeroTrustDLPProfileCustomUpdateResponseJSON struct {
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

func (r *ZeroTrustDLPProfileCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileCustomUpdateResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ZeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkip `json:"skip,required"`
	JSON zeroTrustDLPProfileCustomUpdateResponseContextAwarenessJSON `json:"-"`
}

// zeroTrustDLPProfileCustomUpdateResponseContextAwarenessJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileCustomUpdateResponseContextAwareness]
type zeroTrustDLPProfileCustomUpdateResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomUpdateResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomUpdateResponseContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                            `json:"files,required"`
	JSON  zeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkipJSON `json:"-"`
}

// zeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkipJSON contains the
// JSON metadata for the struct
// [ZeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkip]
type zeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomUpdateResponseContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type ZeroTrustDLPProfileCustomUpdateResponseEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern ZeroTrustDLPProfileCustomUpdateResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                      `json:"profile_id"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileCustomUpdateResponseEntryJSON `json:"-"`
}

// zeroTrustDLPProfileCustomUpdateResponseEntryJSON contains the JSON metadata for
// the struct [ZeroTrustDLPProfileCustomUpdateResponseEntry]
type zeroTrustDLPProfileCustomUpdateResponseEntryJSON struct {
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

func (r *ZeroTrustDLPProfileCustomUpdateResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomUpdateResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type ZeroTrustDLPProfileCustomUpdateResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation ZeroTrustDLPProfileCustomUpdateResponseEntriesPatternValidation `json:"validation"`
	JSON       zeroTrustDLPProfileCustomUpdateResponseEntriesPatternJSON       `json:"-"`
}

// zeroTrustDLPProfileCustomUpdateResponseEntriesPatternJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomUpdateResponseEntriesPattern]
type zeroTrustDLPProfileCustomUpdateResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomUpdateResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomUpdateResponseEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type ZeroTrustDLPProfileCustomUpdateResponseEntriesPatternValidation string

const (
	ZeroTrustDLPProfileCustomUpdateResponseEntriesPatternValidationLuhn ZeroTrustDLPProfileCustomUpdateResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type ZeroTrustDLPProfileCustomUpdateResponseType string

const (
	ZeroTrustDLPProfileCustomUpdateResponseTypeCustom ZeroTrustDLPProfileCustomUpdateResponseType = "custom"
)

// Union satisfied by [ZeroTrustDLPProfileCustomDeleteResponseUnknown] or
// [shared.UnionString].
type ZeroTrustDLPProfileCustomDeleteResponse interface {
	ImplementsZeroTrustDLPProfileCustomDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustDLPProfileCustomDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustDLPProfileCustomGetResponse struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ZeroTrustDLPProfileCustomGetResponseContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                            `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileCustomGetResponseEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileCustomGetResponseType `json:"type"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileCustomGetResponseJSON `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileCustomGetResponse]
type zeroTrustDLPProfileCustomGetResponseJSON struct {
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

func (r *ZeroTrustDLPProfileCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileCustomGetResponseContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ZeroTrustDLPProfileCustomGetResponseContextAwarenessSkip `json:"skip,required"`
	JSON zeroTrustDLPProfileCustomGetResponseContextAwarenessJSON `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseContextAwarenessJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomGetResponseContextAwareness]
type zeroTrustDLPProfileCustomGetResponseContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomGetResponseContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileCustomGetResponseContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                         `json:"files,required"`
	JSON  zeroTrustDLPProfileCustomGetResponseContextAwarenessSkipJSON `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseContextAwarenessSkipJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileCustomGetResponseContextAwarenessSkip]
type zeroTrustDLPProfileCustomGetResponseContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomGetResponseContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type ZeroTrustDLPProfileCustomGetResponseEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern ZeroTrustDLPProfileCustomGetResponseEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                   `json:"profile_id"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileCustomGetResponseEntryJSON `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseEntryJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileCustomGetResponseEntry]
type zeroTrustDLPProfileCustomGetResponseEntryJSON struct {
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

func (r *ZeroTrustDLPProfileCustomGetResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type ZeroTrustDLPProfileCustomGetResponseEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation ZeroTrustDLPProfileCustomGetResponseEntriesPatternValidation `json:"validation"`
	JSON       zeroTrustDLPProfileCustomGetResponseEntriesPatternJSON       `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseEntriesPatternJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomGetResponseEntriesPattern]
type zeroTrustDLPProfileCustomGetResponseEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomGetResponseEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type ZeroTrustDLPProfileCustomGetResponseEntriesPatternValidation string

const (
	ZeroTrustDLPProfileCustomGetResponseEntriesPatternValidationLuhn ZeroTrustDLPProfileCustomGetResponseEntriesPatternValidation = "luhn"
)

// The type of the profile.
type ZeroTrustDLPProfileCustomGetResponseType string

const (
	ZeroTrustDLPProfileCustomGetResponseTypeCustom ZeroTrustDLPProfileCustomGetResponseType = "custom"
)

type ZeroTrustDLPProfileCustomNewParams struct {
	// Identifier
	AccountID param.Field[string]                                      `path:"account_id,required"`
	Profiles  param.Field[[]ZeroTrustDLPProfileCustomNewParamsProfile] `json:"profiles,required"`
}

func (r ZeroTrustDLPProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustDLPProfileCustomNewParamsProfile struct {
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ZeroTrustDLPProfileCustomNewParamsProfilesContextAwareness] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The entries for this profile.
	Entries param.Field[[]ZeroTrustDLPProfileCustomNewParamsProfilesEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
}

func (r ZeroTrustDLPProfileCustomNewParamsProfile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileCustomNewParamsProfilesContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[ZeroTrustDLPProfileCustomNewParamsProfilesContextAwarenessSkip] `json:"skip,required"`
}

func (r ZeroTrustDLPProfileCustomNewParamsProfilesContextAwareness) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileCustomNewParamsProfilesContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r ZeroTrustDLPProfileCustomNewParamsProfilesContextAwarenessSkip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom entry create payload
type ZeroTrustDLPProfileCustomNewParamsProfilesEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The name of the entry.
	Name param.Field[string] `json:"name,required"`
	// A pattern that matches an entry
	Pattern param.Field[ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern] `json:"pattern,required"`
}

func (r ZeroTrustDLPProfileCustomNewParamsProfilesEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A pattern that matches an entry
type ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation param.Field[ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidation] `json:"validation"`
}

func (r ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidation string

const (
	ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidationLuhn ZeroTrustDLPProfileCustomNewParamsProfilesEntriesPatternValidation = "luhn"
)

type ZeroTrustDLPProfileCustomNewResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfileCustomNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfileCustomNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDLPProfileCustomNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustDLPProfileCustomNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDLPProfileCustomNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPProfileCustomNewResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDLPProfileCustomNewResponseEnvelope]
type zeroTrustDLPProfileCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileCustomNewResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDLPProfileCustomNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomNewResponseEnvelopeErrors]
type zeroTrustDLPProfileCustomNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileCustomNewResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDLPProfileCustomNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomNewResponseEnvelopeMessages]
type zeroTrustDLPProfileCustomNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDLPProfileCustomNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfileCustomNewResponseEnvelopeSuccessTrue ZeroTrustDLPProfileCustomNewResponseEnvelopeSuccess = true
)

type ZeroTrustDLPProfileCustomNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       zeroTrustDLPProfileCustomNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPProfileCustomNewResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomNewResponseEnvelopeResultInfo]
type zeroTrustDLPProfileCustomNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileCustomUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount param.Field[float64] `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ZeroTrustDLPProfileCustomUpdateParamsContextAwareness] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// The custom entries for this profile. Array elements with IDs are modifying the
	// existing entry with that ID. Elements without ID will create new entries. Any
	// entry not in the list will be deleted.
	Entries param.Field[[]ZeroTrustDLPProfileCustomUpdateParamsEntry] `json:"entries"`
	// The name of the profile.
	Name param.Field[string] `json:"name"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]ZeroTrustDLPProfileCustomUpdateParamsSharedEntry] `json:"shared_entries"`
}

func (r ZeroTrustDLPProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileCustomUpdateParamsContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[ZeroTrustDLPProfileCustomUpdateParamsContextAwarenessSkip] `json:"skip,required"`
}

func (r ZeroTrustDLPProfileCustomUpdateParamsContextAwareness) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileCustomUpdateParamsContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r ZeroTrustDLPProfileCustomUpdateParamsContextAwarenessSkip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom entry that matches a profile
type ZeroTrustDLPProfileCustomUpdateParamsEntry struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
	// The name of the entry.
	Name param.Field[string] `json:"name"`
	// A pattern that matches an entry
	Pattern param.Field[ZeroTrustDLPProfileCustomUpdateParamsEntriesPattern] `json:"pattern"`
	// ID of the parent profile
	ProfileID param.Field[interface{}] `json:"profile_id"`
}

func (r ZeroTrustDLPProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A pattern that matches an entry
type ZeroTrustDLPProfileCustomUpdateParamsEntriesPattern struct {
	// The regex pattern.
	Regex param.Field[string] `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation param.Field[ZeroTrustDLPProfileCustomUpdateParamsEntriesPatternValidation] `json:"validation"`
}

func (r ZeroTrustDLPProfileCustomUpdateParamsEntriesPattern) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type ZeroTrustDLPProfileCustomUpdateParamsEntriesPatternValidation string

const (
	ZeroTrustDLPProfileCustomUpdateParamsEntriesPatternValidationLuhn ZeroTrustDLPProfileCustomUpdateParamsEntriesPatternValidation = "luhn"
)

// Properties of a predefined entry in a custom profile
//
// Satisfied by
// [ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined],
// [ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration].
type ZeroTrustDLPProfileCustomUpdateParamsSharedEntry interface {
	implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntry()
}

// Properties of a predefined entry in a custom profile
type ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdatePredefined) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntry() {
}

// Properties of an integration entry in a custom profile
type ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration struct {
	// Whether the entry is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustDLPProfileCustomUpdateParamsSharedEntriesDLPSharedEntryUpdateIntegration) implementsZeroTrustDLPProfileCustomUpdateParamsSharedEntry() {
}

type ZeroTrustDLPProfileCustomDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPProfileCustomDeleteResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfileCustomDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfileCustomDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDLPProfileCustomDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPProfileCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPProfileCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPProfileCustomDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileCustomDeleteResponseEnvelope]
type zeroTrustDLPProfileCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileCustomDeleteResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustDLPProfileCustomDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfileCustomDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomDeleteResponseEnvelopeErrors]
type zeroTrustDLPProfileCustomDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileCustomDeleteResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustDLPProfileCustomDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfileCustomDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileCustomDeleteResponseEnvelopeMessages]
type zeroTrustDLPProfileCustomDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDLPProfileCustomDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfileCustomDeleteResponseEnvelopeSuccessTrue ZeroTrustDLPProfileCustomDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustDLPProfileCustomGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPProfileCustomGetResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfileCustomGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfileCustomGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDLPProfileCustomGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPProfileCustomGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPProfileCustomGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDLPProfileCustomGetResponseEnvelope]
type zeroTrustDLPProfileCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileCustomGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDLPProfileCustomGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomGetResponseEnvelopeErrors]
type zeroTrustDLPProfileCustomGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileCustomGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDLPProfileCustomGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfileCustomGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileCustomGetResponseEnvelopeMessages]
type zeroTrustDLPProfileCustomGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileCustomGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileCustomGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDLPProfileCustomGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfileCustomGetResponseEnvelopeSuccessTrue ZeroTrustDLPProfileCustomGetResponseEnvelopeSuccess = true
)
