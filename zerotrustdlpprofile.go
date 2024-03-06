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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDLPProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDLPProfileService]
// method instead.
type ZeroTrustDLPProfileService struct {
	Options     []option.RequestOption
	Customs     *ZeroTrustDLPProfileCustomService
	Predefineds *ZeroTrustDLPProfilePredefinedService
}

// NewZeroTrustDLPProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDLPProfileService(opts ...option.RequestOption) (r *ZeroTrustDLPProfileService) {
	r = &ZeroTrustDLPProfileService{}
	r.Options = opts
	r.Customs = NewZeroTrustDLPProfileCustomService(opts...)
	r.Predefineds = NewZeroTrustDLPProfilePredefinedService(opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *ZeroTrustDLPProfileService) List(ctx context.Context, query ZeroTrustDLPProfileListParams, opts ...option.RequestOption) (res *[]ZeroTrustDLPProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfileListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a DLP profile by ID. Supports both predefined and custom profiles
func (r *ZeroTrustDLPProfileService) Get(ctx context.Context, profileID string, query ZeroTrustDLPProfileGetParams, opts ...option.RequestOption) (res *ZeroTrustDLPProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfileGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ZeroTrustDLPProfileListResponseDLPPredefinedProfile],
// [ZeroTrustDLPProfileListResponseDLPCustomProfile] or
// [ZeroTrustDLPProfileListResponseDLPIntegrationProfile].
type ZeroTrustDLPProfileListResponse interface {
	implementsZeroTrustDLPProfileListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustDLPProfileListResponse)(nil)).Elem(), "")
}

type ZeroTrustDLPProfileListResponseDLPPredefinedProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwareness `json:"context_awareness"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileListResponseDLPPredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type ZeroTrustDLPProfileListResponseDLPPredefinedProfileType `json:"type"`
	JSON zeroTrustDLPProfileListResponseDLPPredefinedProfileJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPPredefinedProfileJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileListResponseDLPPredefinedProfile]
type zeroTrustDLPProfileListResponseDLPPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPPredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustDLPProfileListResponseDLPPredefinedProfile) implementsZeroTrustDLPProfileListResponse() {
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip `json:"skip,required"`
	JSON zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessJSON contains
// the JSON metadata for the struct
// [ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwareness]
type zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                                        `json:"files,required"`
	JSON  zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON
// contains the JSON metadata for the struct
// [ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip]
type zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A predefined entry that matches a profile
type ZeroTrustDLPProfileListResponseDLPPredefinedProfileEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                  `json:"profile_id"`
	JSON      zeroTrustDLPProfileListResponseDLPPredefinedProfileEntryJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPPredefinedProfileEntryJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileListResponseDLPPredefinedProfileEntry]
type zeroTrustDLPProfileListResponseDLPPredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPPredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPPredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type ZeroTrustDLPProfileListResponseDLPPredefinedProfileType string

const (
	ZeroTrustDLPProfileListResponseDLPPredefinedProfileTypePredefined ZeroTrustDLPProfileListResponseDLPPredefinedProfileType = "predefined"
)

type ZeroTrustDLPProfileListResponseDLPCustomProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                                       `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileListResponseDLPCustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileListResponseDLPCustomProfileType `json:"type"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileListResponseDLPCustomProfileJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPCustomProfileJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileListResponseDLPCustomProfile]
type zeroTrustDLPProfileListResponseDLPCustomProfileJSON struct {
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

func (r *ZeroTrustDLPProfileListResponseDLPCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustDLPProfileListResponseDLPCustomProfile) implementsZeroTrustDLPProfileListResponse() {
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkip `json:"skip,required"`
	JSON zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessJSON contains the
// JSON metadata for the struct
// [ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwareness]
type zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                                    `json:"files,required"`
	JSON  zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkipJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkipJSON contains
// the JSON metadata for the struct
// [ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkip]
type zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPCustomProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type ZeroTrustDLPProfileListResponseDLPCustomProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                              `json:"profile_id"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileListResponseDLPCustomProfileEntryJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPCustomProfileEntryJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileListResponseDLPCustomProfileEntry]
type zeroTrustDLPProfileListResponseDLPCustomProfileEntryJSON struct {
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

func (r *ZeroTrustDLPProfileListResponseDLPCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternValidation `json:"validation"`
	JSON       zeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternJSON       `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternJSON contains the
// JSON metadata for the struct
// [ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPattern]
type zeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternValidation string

const (
	ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternValidationLuhn ZeroTrustDLPProfileListResponseDLPCustomProfileEntriesPatternValidation = "luhn"
)

// The type of the profile.
type ZeroTrustDLPProfileListResponseDLPCustomProfileType string

const (
	ZeroTrustDLPProfileListResponseDLPCustomProfileTypeCustom ZeroTrustDLPProfileListResponseDLPCustomProfileType = "custom"
)

type ZeroTrustDLPProfileListResponseDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileListResponseDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileListResponseDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileListResponseDLPIntegrationProfileJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPIntegrationProfileJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileListResponseDLPIntegrationProfile]
type zeroTrustDLPProfileListResponseDLPIntegrationProfileJSON struct {
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

func (r *ZeroTrustDLPProfileListResponseDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustDLPProfileListResponseDLPIntegrationProfile) implementsZeroTrustDLPProfileListResponse() {
}

// An entry derived from an integration
type ZeroTrustDLPProfileListResponseDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                   `json:"profile_id"`
	UpdatedAt time.Time                                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileListResponseDLPIntegrationProfileEntryJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseDLPIntegrationProfileEntryJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileListResponseDLPIntegrationProfileEntry]
type zeroTrustDLPProfileListResponseDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type ZeroTrustDLPProfileListResponseDLPIntegrationProfileType string

const (
	ZeroTrustDLPProfileListResponseDLPIntegrationProfileTypeIntegration ZeroTrustDLPProfileListResponseDLPIntegrationProfileType = "integration"
)

// Union satisfied by [ZeroTrustDLPProfileGetResponseDLPPredefinedProfile],
// [ZeroTrustDLPProfileGetResponseDLPCustomProfile] or
// [ZeroTrustDLPProfileGetResponseDLPIntegrationProfile].
type ZeroTrustDLPProfileGetResponse interface {
	implementsZeroTrustDLPProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustDLPProfileGetResponse)(nil)).Elem(), "")
}

type ZeroTrustDLPProfileGetResponseDLPPredefinedProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwareness `json:"context_awareness"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileGetResponseDLPPredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type ZeroTrustDLPProfileGetResponseDLPPredefinedProfileType `json:"type"`
	JSON zeroTrustDLPProfileGetResponseDLPPredefinedProfileJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPPredefinedProfileJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileGetResponseDLPPredefinedProfile]
type zeroTrustDLPProfileGetResponseDLPPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPPredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustDLPProfileGetResponseDLPPredefinedProfile) implementsZeroTrustDLPProfileGetResponse() {
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip `json:"skip,required"`
	JSON zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessJSON contains
// the JSON metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwareness]
type zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                                       `json:"files,required"`
	JSON  zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON
// contains the JSON metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip]
type zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A predefined entry that matches a profile
type ZeroTrustDLPProfileGetResponseDLPPredefinedProfileEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                 `json:"profile_id"`
	JSON      zeroTrustDLPProfileGetResponseDLPPredefinedProfileEntryJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPPredefinedProfileEntryJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPPredefinedProfileEntry]
type zeroTrustDLPProfileGetResponseDLPPredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPPredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPPredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type ZeroTrustDLPProfileGetResponseDLPPredefinedProfileType string

const (
	ZeroTrustDLPProfileGetResponseDLPPredefinedProfileTypePredefined ZeroTrustDLPProfileGetResponseDLPPredefinedProfileType = "predefined"
)

type ZeroTrustDLPProfileGetResponseDLPCustomProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                                      `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileGetResponseDLPCustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileGetResponseDLPCustomProfileType `json:"type"`
	UpdatedAt time.Time                                          `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileGetResponseDLPCustomProfileJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPCustomProfileJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileGetResponseDLPCustomProfile]
type zeroTrustDLPProfileGetResponseDLPCustomProfileJSON struct {
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

func (r *ZeroTrustDLPProfileGetResponseDLPCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustDLPProfileGetResponseDLPCustomProfile) implementsZeroTrustDLPProfileGetResponse() {}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkip `json:"skip,required"`
	JSON zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessJSON contains the
// JSON metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwareness]
type zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                                   `json:"files,required"`
	JSON  zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON contains
// the JSON metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkip]
type zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type ZeroTrustDLPProfileGetResponseDLPCustomProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                             `json:"profile_id"`
	UpdatedAt time.Time                                               `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileGetResponseDLPCustomProfileEntryJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPCustomProfileEntryJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileGetResponseDLPCustomProfileEntry]
type zeroTrustDLPProfileGetResponseDLPCustomProfileEntryJSON struct {
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

func (r *ZeroTrustDLPProfileGetResponseDLPCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternValidation `json:"validation"`
	JSON       zeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternJSON       `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternJSON contains the
// JSON metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPattern]
type zeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternValidation string

const (
	ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternValidationLuhn ZeroTrustDLPProfileGetResponseDLPCustomProfileEntriesPatternValidation = "luhn"
)

// The type of the profile.
type ZeroTrustDLPProfileGetResponseDLPCustomProfileType string

const (
	ZeroTrustDLPProfileGetResponseDLPCustomProfileTypeCustom ZeroTrustDLPProfileGetResponseDLPCustomProfileType = "custom"
)

type ZeroTrustDLPProfileGetResponseDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileGetResponseDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                               `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileGetResponseDLPIntegrationProfileJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPIntegrationProfileJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileGetResponseDLPIntegrationProfile]
type zeroTrustDLPProfileGetResponseDLPIntegrationProfileJSON struct {
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

func (r *ZeroTrustDLPProfileGetResponseDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustDLPProfileGetResponseDLPIntegrationProfile) implementsZeroTrustDLPProfileGetResponse() {
}

// An entry derived from an integration
type ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                  `json:"profile_id"`
	UpdatedAt time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileGetResponseDLPIntegrationProfileEntryJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPIntegrationProfileEntryJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry]
type zeroTrustDLPProfileGetResponseDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type ZeroTrustDLPProfileGetResponseDLPIntegrationProfileType string

const (
	ZeroTrustDLPProfileGetResponseDLPIntegrationProfileTypeIntegration ZeroTrustDLPProfileGetResponseDLPIntegrationProfileType = "integration"
)

type ZeroTrustDLPProfileListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPProfileListResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfileListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfileListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDLPProfileListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustDLPProfileListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDLPProfileListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPProfileListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileListResponseEnvelope]
type zeroTrustDLPProfileListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustDLPProfileListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDLPProfileListResponseEnvelopeErrors]
type zeroTrustDLPProfileListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDLPProfileListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileListResponseEnvelopeMessages]
type zeroTrustDLPProfileListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDLPProfileListResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfileListResponseEnvelopeSuccessTrue ZeroTrustDLPProfileListResponseEnvelopeSuccess = true
)

type ZeroTrustDLPProfileListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       zeroTrustDLPProfileListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileListResponseEnvelopeResultInfo]
type zeroTrustDLPProfileListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPProfileGetResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfileGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDLPProfileGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPProfileGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPProfileGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileGetResponseEnvelope]
type zeroTrustDLPProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustDLPProfileGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDLPProfileGetResponseEnvelopeErrors]
type zeroTrustDLPProfileGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDLPProfileGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustDLPProfileGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileGetResponseEnvelopeMessages]
type zeroTrustDLPProfileGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDLPProfileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustDLPProfileGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfileGetResponseEnvelopeSuccessTrue ZeroTrustDLPProfileGetResponseEnvelopeSuccess = true
)
