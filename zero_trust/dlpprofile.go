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
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// DLPProfileService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDLPProfileService] method instead.
type DLPProfileService struct {
	Options     []option.RequestOption
	Customs     *DLPProfileCustomService
	Predefineds *DLPProfilePredefinedService
}

// NewDLPProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPProfileService(opts ...option.RequestOption) (r *DLPProfileService) {
	r = &DLPProfileService{}
	r.Options = opts
	r.Customs = NewDLPProfileCustomService(opts...)
	r.Predefineds = NewDLPProfilePredefinedService(opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) List(ctx context.Context, query DLPProfileListParams, opts ...option.RequestOption) (res *[]DLPProfileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a DLP profile by ID. Supports both predefined and custom profiles
func (r *DLPProfileService) Get(ctx context.Context, profileID string, query DLPProfileGetParams, opts ...option.RequestOption) (res *DLPProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPProfileGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zero_trust.DLPProfileListResponseDLPPredefinedProfile],
// [zero_trust.DLPProfileListResponseDLPCustomProfile] or
// [zero_trust.DLPProfileListResponseDLPIntegrationProfile].
type DLPProfileListResponse interface {
	implementsZeroTrustDLPProfileListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileListResponseDLPPredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileListResponseDLPCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileListResponseDLPIntegrationProfile{}),
		},
	)
}

type DLPProfileListResponseDLPPredefinedProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfileListResponseDLPPredefinedProfileContextAwareness `json:"context_awareness"`
	// The entries for this profile.
	Entries []DLPProfileListResponseDLPPredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type DLPProfileListResponseDLPPredefinedProfileType `json:"type"`
	JSON dlpProfileListResponseDLPPredefinedProfileJSON `json:"-"`
}

// dlpProfileListResponseDLPPredefinedProfileJSON contains the JSON metadata for
// the struct [DLPProfileListResponseDLPPredefinedProfile]
type dlpProfileListResponseDLPPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPPredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseDLPPredefinedProfile) implementsZeroTrustDLPProfileListResponse() {}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileListResponseDLPPredefinedProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfileListResponseDLPPredefinedProfileContextAwarenessJSON `json:"-"`
}

// dlpProfileListResponseDLPPredefinedProfileContextAwarenessJSON contains the JSON
// metadata for the struct
// [DLPProfileListResponseDLPPredefinedProfileContextAwareness]
type dlpProfileListResponseDLPPredefinedProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPPredefinedProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPPredefinedProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                               `json:"files,required"`
	JSON  dlpProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON `json:"-"`
}

// dlpProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON contains the
// JSON metadata for the struct
// [DLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip]
type dlpProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPPredefinedProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPPredefinedProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A predefined entry that matches a profile
type DLPProfileListResponseDLPPredefinedProfileEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                         `json:"profile_id"`
	JSON      dlpProfileListResponseDLPPredefinedProfileEntryJSON `json:"-"`
}

// dlpProfileListResponseDLPPredefinedProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileListResponseDLPPredefinedProfileEntry]
type dlpProfileListResponseDLPPredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPPredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPPredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type DLPProfileListResponseDLPPredefinedProfileType string

const (
	DLPProfileListResponseDLPPredefinedProfileTypePredefined DLPProfileListResponseDLPPredefinedProfileType = "predefined"
)

type DLPProfileListResponseDLPCustomProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfileListResponseDLPCustomProfileContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                              `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileListResponseDLPCustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileListResponseDLPCustomProfileType `json:"type"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      dlpProfileListResponseDLPCustomProfileJSON `json:"-"`
}

// dlpProfileListResponseDLPCustomProfileJSON contains the JSON metadata for the
// struct [DLPProfileListResponseDLPCustomProfile]
type dlpProfileListResponseDLPCustomProfileJSON struct {
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

func (r *DLPProfileListResponseDLPCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseDLPCustomProfile) implementsZeroTrustDLPProfileListResponse() {}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileListResponseDLPCustomProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfileListResponseDLPCustomProfileContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfileListResponseDLPCustomProfileContextAwarenessJSON `json:"-"`
}

// dlpProfileListResponseDLPCustomProfileContextAwarenessJSON contains the JSON
// metadata for the struct [DLPProfileListResponseDLPCustomProfileContextAwareness]
type dlpProfileListResponseDLPCustomProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPCustomProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPCustomProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileListResponseDLPCustomProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                           `json:"files,required"`
	JSON  dlpProfileListResponseDLPCustomProfileContextAwarenessSkipJSON `json:"-"`
}

// dlpProfileListResponseDLPCustomProfileContextAwarenessSkipJSON contains the JSON
// metadata for the struct
// [DLPProfileListResponseDLPCustomProfileContextAwarenessSkip]
type dlpProfileListResponseDLPCustomProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPCustomProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPCustomProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type DLPProfileListResponseDLPCustomProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern DLPProfileListResponseDLPCustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                     `json:"profile_id"`
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	JSON      dlpProfileListResponseDLPCustomProfileEntryJSON `json:"-"`
}

// dlpProfileListResponseDLPCustomProfileEntryJSON contains the JSON metadata for
// the struct [DLPProfileListResponseDLPCustomProfileEntry]
type dlpProfileListResponseDLPCustomProfileEntryJSON struct {
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

func (r *DLPProfileListResponseDLPCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type DLPProfileListResponseDLPCustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation DLPProfileListResponseDLPCustomProfileEntriesPatternValidation `json:"validation"`
	JSON       dlpProfileListResponseDLPCustomProfileEntriesPatternJSON       `json:"-"`
}

// dlpProfileListResponseDLPCustomProfileEntriesPatternJSON contains the JSON
// metadata for the struct [DLPProfileListResponseDLPCustomProfileEntriesPattern]
type dlpProfileListResponseDLPCustomProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPCustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPCustomProfileEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileListResponseDLPCustomProfileEntriesPatternValidation string

const (
	DLPProfileListResponseDLPCustomProfileEntriesPatternValidationLuhn DLPProfileListResponseDLPCustomProfileEntriesPatternValidation = "luhn"
)

// The type of the profile.
type DLPProfileListResponseDLPCustomProfileType string

const (
	DLPProfileListResponseDLPCustomProfileTypeCustom DLPProfileListResponseDLPCustomProfileType = "custom"
)

type DLPProfileListResponseDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileListResponseDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileListResponseDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	JSON      dlpProfileListResponseDLPIntegrationProfileJSON `json:"-"`
}

// dlpProfileListResponseDLPIntegrationProfileJSON contains the JSON metadata for
// the struct [DLPProfileListResponseDLPIntegrationProfile]
type dlpProfileListResponseDLPIntegrationProfileJSON struct {
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

func (r *DLPProfileListResponseDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileListResponseDLPIntegrationProfile) implementsZeroTrustDLPProfileListResponse() {}

// An entry derived from an integration
type DLPProfileListResponseDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                          `json:"profile_id"`
	UpdatedAt time.Time                                            `json:"updated_at" format:"date-time"`
	JSON      dlpProfileListResponseDLPIntegrationProfileEntryJSON `json:"-"`
}

// dlpProfileListResponseDLPIntegrationProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileListResponseDLPIntegrationProfileEntry]
type dlpProfileListResponseDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type DLPProfileListResponseDLPIntegrationProfileType string

const (
	DLPProfileListResponseDLPIntegrationProfileTypeIntegration DLPProfileListResponseDLPIntegrationProfileType = "integration"
)

// Union satisfied by [zero_trust.DLPProfileGetResponseDLPPredefinedProfile],
// [zero_trust.DLPProfileGetResponseDLPCustomProfile] or
// [zero_trust.DLPProfileGetResponseDLPIntegrationProfile].
type DLPProfileGetResponse interface {
	implementsZeroTrustDLPProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseDLPPredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseDLPCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseDLPIntegrationProfile{}),
		},
	)
}

type DLPProfileGetResponseDLPPredefinedProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfileGetResponseDLPPredefinedProfileContextAwareness `json:"context_awareness"`
	// The entries for this profile.
	Entries []DLPProfileGetResponseDLPPredefinedProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type DLPProfileGetResponseDLPPredefinedProfileType `json:"type"`
	JSON dlpProfileGetResponseDLPPredefinedProfileJSON `json:"-"`
}

// dlpProfileGetResponseDLPPredefinedProfileJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseDLPPredefinedProfile]
type dlpProfileGetResponseDLPPredefinedProfileJSON struct {
	ID                apijson.Field
	AllowedMatchCount apijson.Field
	ContextAwareness  apijson.Field
	Entries           apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPPredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPPredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseDLPPredefinedProfile) implementsZeroTrustDLPProfileGetResponse() {}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileGetResponseDLPPredefinedProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfileGetResponseDLPPredefinedProfileContextAwarenessJSON `json:"-"`
}

// dlpProfileGetResponseDLPPredefinedProfileContextAwarenessJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseDLPPredefinedProfileContextAwareness]
type dlpProfileGetResponseDLPPredefinedProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPPredefinedProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPPredefinedProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                              `json:"files,required"`
	JSON  dlpProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON `json:"-"`
}

// dlpProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip]
type dlpProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPPredefinedProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPPredefinedProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A predefined entry that matches a profile
type DLPProfileGetResponseDLPPredefinedProfileEntry struct {
	// The ID for this entry
	ID string `json:"id"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                        `json:"profile_id"`
	JSON      dlpProfileGetResponseDLPPredefinedProfileEntryJSON `json:"-"`
}

// dlpProfileGetResponseDLPPredefinedProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseDLPPredefinedProfileEntry]
type dlpProfileGetResponseDLPPredefinedProfileEntryJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPPredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPPredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type DLPProfileGetResponseDLPPredefinedProfileType string

const (
	DLPProfileGetResponseDLPPredefinedProfileTypePredefined DLPProfileGetResponseDLPPredefinedProfileType = "predefined"
)

type DLPProfileGetResponseDLPCustomProfile struct {
	// The ID for this profile
	ID string `json:"id"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount float64 `json:"allowed_match_count"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness DLPProfileGetResponseDLPCustomProfileContextAwareness `json:"context_awareness"`
	CreatedAt        time.Time                                             `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileGetResponseDLPCustomProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileGetResponseDLPCustomProfileType `json:"type"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPCustomProfileJSON `json:"-"`
}

// dlpProfileGetResponseDLPCustomProfileJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseDLPCustomProfile]
type dlpProfileGetResponseDLPCustomProfileJSON struct {
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

func (r *DLPProfileGetResponseDLPCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseDLPCustomProfile) implementsZeroTrustDLPProfileGetResponse() {}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
type DLPProfileGetResponseDLPCustomProfileContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip DLPProfileGetResponseDLPCustomProfileContextAwarenessSkip `json:"skip,required"`
	JSON dlpProfileGetResponseDLPCustomProfileContextAwarenessJSON `json:"-"`
}

// dlpProfileGetResponseDLPCustomProfileContextAwarenessJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseDLPCustomProfileContextAwareness]
type dlpProfileGetResponseDLPCustomProfileContextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPCustomProfileContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPCustomProfileContextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type DLPProfileGetResponseDLPCustomProfileContextAwarenessSkip struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                                                          `json:"files,required"`
	JSON  dlpProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON `json:"-"`
}

// dlpProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseDLPCustomProfileContextAwarenessSkip]
type dlpProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPCustomProfileContextAwarenessSkip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPCustomProfileContextAwarenessSkipJSON) RawJSON() string {
	return r.raw
}

// A custom entry that matches a profile
type DLPProfileGetResponseDLPCustomProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// A pattern that matches an entry
	Pattern DLPProfileGetResponseDLPCustomProfileEntriesPattern `json:"pattern"`
	// ID of the parent profile
	ProfileID interface{}                                    `json:"profile_id"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPCustomProfileEntryJSON `json:"-"`
}

// dlpProfileGetResponseDLPCustomProfileEntryJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseDLPCustomProfileEntry]
type dlpProfileGetResponseDLPCustomProfileEntryJSON struct {
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

func (r *DLPProfileGetResponseDLPCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

// A pattern that matches an entry
type DLPProfileGetResponseDLPCustomProfileEntriesPattern struct {
	// The regex pattern.
	Regex string `json:"regex,required"`
	// Validation algorithm for the pattern. This algorithm will get run on potential
	// matches, and if it returns false, the entry will not be matched.
	Validation DLPProfileGetResponseDLPCustomProfileEntriesPatternValidation `json:"validation"`
	JSON       dlpProfileGetResponseDLPCustomProfileEntriesPatternJSON       `json:"-"`
}

// dlpProfileGetResponseDLPCustomProfileEntriesPatternJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseDLPCustomProfileEntriesPattern]
type dlpProfileGetResponseDLPCustomProfileEntriesPatternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPCustomProfileEntriesPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPCustomProfileEntriesPatternJSON) RawJSON() string {
	return r.raw
}

// Validation algorithm for the pattern. This algorithm will get run on potential
// matches, and if it returns false, the entry will not be matched.
type DLPProfileGetResponseDLPCustomProfileEntriesPatternValidation string

const (
	DLPProfileGetResponseDLPCustomProfileEntriesPatternValidationLuhn DLPProfileGetResponseDLPCustomProfileEntriesPatternValidation = "luhn"
)

// The type of the profile.
type DLPProfileGetResponseDLPCustomProfileType string

const (
	DLPProfileGetResponseDLPCustomProfileTypeCustom DLPProfileGetResponseDLPCustomProfileType = "custom"
)

type DLPProfileGetResponseDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfileGetResponseDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfileGetResponseDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPIntegrationProfileJSON `json:"-"`
}

// dlpProfileGetResponseDLPIntegrationProfileJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseDLPIntegrationProfile]
type dlpProfileGetResponseDLPIntegrationProfileJSON struct {
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

func (r *DLPProfileGetResponseDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseDLPIntegrationProfile) implementsZeroTrustDLPProfileGetResponse() {}

// An entry derived from an integration
type DLPProfileGetResponseDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                         `json:"profile_id"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseDLPIntegrationProfileEntryJSON `json:"-"`
}

// dlpProfileGetResponseDLPIntegrationProfileEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseDLPIntegrationProfileEntry]
type dlpProfileGetResponseDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseDLPIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

// The type of the profile.
type DLPProfileGetResponseDLPIntegrationProfileType string

const (
	DLPProfileGetResponseDLPIntegrationProfileTypeIntegration DLPProfileGetResponseDLPIntegrationProfileType = "integration"
)

type DLPProfileListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileListResponseEnvelope struct {
	Errors   []DLPProfileListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DLPProfileListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DLPProfileListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DLPProfileListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dlpProfileListResponseEnvelopeJSON       `json:"-"`
}

// dlpProfileListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPProfileListResponseEnvelope]
type dlpProfileListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dlpProfileListResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfileListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileListResponseEnvelopeErrors]
type dlpProfileListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dlpProfileListResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfileListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPProfileListResponseEnvelopeMessages]
type dlpProfileListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileListResponseEnvelopeSuccess bool

const (
	DLPProfileListResponseEnvelopeSuccessTrue DLPProfileListResponseEnvelopeSuccess = true
)

type DLPProfileListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       dlpProfileListResponseEnvelopeResultInfoJSON `json:"-"`
}

// dlpProfileListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DLPProfileListResponseEnvelopeResultInfo]
type dlpProfileListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetResponseEnvelope struct {
	Errors   []DLPProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DLPProfileGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpProfileGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseEnvelope]
type dlpProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dlpProfileGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeErrors]
type dlpProfileGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dlpProfileGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeMessages]
type dlpProfileGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPProfileGetResponseEnvelopeSuccess bool

const (
	DLPProfileGetResponseEnvelopeSuccessTrue DLPProfileGetResponseEnvelopeSuccess = true
)
