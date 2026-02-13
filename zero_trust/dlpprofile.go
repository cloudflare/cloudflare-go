// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// DLPProfileService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfileService] method instead.
type DLPProfileService struct {
	Options    []option.RequestOption
	Custom     *DLPProfileCustomService
	Predefined *DLPProfilePredefinedService
}

// NewDLPProfileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPProfileService(opts ...option.RequestOption) (r *DLPProfileService) {
	r = &DLPProfileService{}
	r.Options = opts
	r.Custom = NewDLPProfileCustomService(opts...)
	r.Predefined = NewDLPProfilePredefinedService(opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) List(ctx context.Context, params DLPProfileListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Profile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all DLP profiles in an account.
func (r *DLPProfileService) ListAutoPaging(ctx context.Context, params DLPProfileListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Profile] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Fetches a DLP profile by ID.
func (r *DLPProfileService) Get(ctx context.Context, profileID string, query DLPProfileGetParams, opts ...option.RequestOption) (res *DLPProfileGetResponse, err error) {
	var env DLPProfileGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
//
// Deprecated: deprecated
type ContextAwareness struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled bool `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip SkipConfiguration    `json:"skip,required"`
	JSON contextAwarenessJSON `json:"-"`
}

// contextAwarenessJSON contains the JSON metadata for the struct
// [ContextAwareness]
type contextAwarenessJSON struct {
	Enabled     apijson.Field
	Skip        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContextAwareness) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextAwarenessJSON) RawJSON() string {
	return r.raw
}

// Scan the context of predefined entries to only return matches surrounded by
// keywords.
//
// Deprecated: deprecated
type ContextAwarenessParam struct {
	// If true, scan the context of predefined entries to only return matches
	// surrounded by keywords.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Content types to exclude from context analysis and return all matches.
	Skip param.Field[SkipConfigurationParam] `json:"skip,required"`
}

func (r ContextAwarenessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Profile struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile.
	Name             string      `json:"name,required"`
	Type             ProfileType `json:"type,required"`
	AIContextEnabled bool        `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64                      `json:"allowed_match_count"`
	ConfidenceThreshold ProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// This field can have the runtime type of [[]ProfileCustomProfileEntry],
	// [[]ProfilePredefinedProfileEntry], [[]ProfileIntegrationProfileEntry].
	Entries    interface{} `json:"entries"`
	OCREnabled bool        `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool `json:"open_access"`
	// This field can have the runtime type of [[]ProfileCustomProfileSharedEntry],
	// [[]ProfileIntegrationProfileSharedEntry].
	SharedEntries interface{} `json:"shared_entries"`
	// When the profile was lasted updated.
	UpdatedAt time.Time   `json:"updated_at" format:"date-time"`
	JSON      profileJSON `json:"-"`
	union     ProfileUnion
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	ID                  apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	AllowedMatchCount   apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Description         apijson.Field
	Entries             apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	SharedEntries       apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r profileJSON) RawJSON() string {
	return r.raw
}

func (r *Profile) UnmarshalJSON(data []byte) (err error) {
	*r = Profile{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ProfileCustomProfile],
// [ProfilePredefinedProfile], [ProfileIntegrationProfile].
func (r Profile) AsUnion() ProfileUnion {
	return r.union
}

// Union satisfied by [ProfileCustomProfile], [ProfilePredefinedProfile] or
// [ProfileIntegrationProfile].
type ProfileUnion interface {
	implementsProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfile{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfile{}),
		},
	)
}

type ProfileCustomProfile struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the profile.
	Name       string                   `json:"name,required"`
	OCREnabled bool                     `json:"ocr_enabled,required"`
	Type       ProfileCustomProfileType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time                               `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool                                    `json:"ai_context_enabled"`
	ConfidenceThreshold ProfileCustomProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// Deprecated: deprecated
	Entries       []ProfileCustomProfileEntry       `json:"entries"`
	SharedEntries []ProfileCustomProfileSharedEntry `json:"shared_entries"`
	JSON          profileCustomProfileJSON          `json:"-"`
}

// profileCustomProfileJSON contains the JSON metadata for the struct
// [ProfileCustomProfile]
type profileCustomProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	CreatedAt           apijson.Field
	Name                apijson.Field
	OCREnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	Description         apijson.Field
	Entries             apijson.Field
	SharedEntries       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileCustomProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfile) implementsProfile() {}

type ProfileCustomProfileType string

const (
	ProfileCustomProfileTypeCustom ProfileCustomProfileType = "custom"
)

func (r ProfileCustomProfileType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileTypeCustom:
		return true
	}
	return false
}

type ProfileCustomProfileConfidenceThreshold string

const (
	ProfileCustomProfileConfidenceThresholdLow      ProfileCustomProfileConfidenceThreshold = "low"
	ProfileCustomProfileConfidenceThresholdMedium   ProfileCustomProfileConfidenceThreshold = "medium"
	ProfileCustomProfileConfidenceThresholdHigh     ProfileCustomProfileConfidenceThreshold = "high"
	ProfileCustomProfileConfidenceThresholdVeryHigh ProfileCustomProfileConfidenceThreshold = "very_high"
)

func (r ProfileCustomProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfileCustomProfileConfidenceThresholdLow, ProfileCustomProfileConfidenceThresholdMedium, ProfileCustomProfileConfidenceThresholdHigh, ProfileCustomProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type ProfileCustomProfileEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                            `json:"enabled,required"`
	Name    string                          `json:"name,required"`
	Type    ProfileCustomProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfileCustomProfileEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [ProfileCustomProfileEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                   `json:"word_list"`
	JSON     profileCustomProfileEntryJSON `json:"-"`
	union    ProfileCustomProfileEntriesUnion
}

// profileCustomProfileEntryJSON contains the JSON metadata for the struct
// [ProfileCustomProfileEntry]
type profileCustomProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profileCustomProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileCustomProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileCustomProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileCustomProfileEntriesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfileCustomProfileEntriesCustomEntry],
// [ProfileCustomProfileEntriesPredefinedEntry],
// [ProfileCustomProfileEntriesIntegrationEntry],
// [ProfileCustomProfileEntriesExactDataEntry],
// [ProfileCustomProfileEntriesDocumentFingerprintEntry],
// [ProfileCustomProfileEntriesWordListEntry].
func (r ProfileCustomProfileEntry) AsUnion() ProfileCustomProfileEntriesUnion {
	return r.union
}

// Union satisfied by [ProfileCustomProfileEntriesCustomEntry],
// [ProfileCustomProfileEntriesPredefinedEntry],
// [ProfileCustomProfileEntriesIntegrationEntry],
// [ProfileCustomProfileEntriesExactDataEntry],
// [ProfileCustomProfileEntriesDocumentFingerprintEntry] or
// [ProfileCustomProfileEntriesWordListEntry].
type ProfileCustomProfileEntriesUnion interface {
	implementsProfileCustomProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileCustomProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileEntriesWordListEntry{}),
		},
	)
}

type ProfileCustomProfileEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                       `json:"enabled,required"`
	Name        string                                     `json:"name,required"`
	Pattern     Pattern                                    `json:"pattern,required"`
	Type        ProfileCustomProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                  `json:"updated_at,required" format:"date-time"`
	Description string                                     `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileEntriesCustomEntryJSON `json:"-"`
}

// profileCustomProfileEntriesCustomEntryJSON contains the JSON metadata for the
// struct [ProfileCustomProfileEntriesCustomEntry]
type profileCustomProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesCustomEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesCustomEntryType string

const (
	ProfileCustomProfileEntriesCustomEntryTypeCustom ProfileCustomProfileEntriesCustomEntryType = "custom"
)

func (r ProfileCustomProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesPredefinedEntry struct {
	ID         string                                               `json:"id,required" format:"uuid"`
	Confidence ProfileCustomProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                 `json:"enabled,required"`
	Name       string                                               `json:"name,required"`
	Type       ProfileCustomProfileEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	Variant   ProfileCustomProfileEntriesPredefinedEntryVariant `json:"variant"`
	JSON      profileCustomProfileEntriesPredefinedEntryJSON    `json:"-"`
}

// profileCustomProfileEntriesPredefinedEntryJSON contains the JSON metadata for
// the struct [ProfileCustomProfileEntriesPredefinedEntry]
type profileCustomProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesPredefinedEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                     `json:"available,required"`
	JSON      profileCustomProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profileCustomProfileEntriesPredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct [ProfileCustomProfileEntriesPredefinedEntryConfidence]
type profileCustomProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileCustomProfileEntriesPredefinedEntryType string

const (
	ProfileCustomProfileEntriesPredefinedEntryTypePredefined ProfileCustomProfileEntriesPredefinedEntryType = "predefined"
)

func (r ProfileCustomProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesPredefinedEntryVariant struct {
	TopicType   ProfileCustomProfileEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        ProfileCustomProfileEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                     `json:"description,nullable"`
	JSON        profileCustomProfileEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// profileCustomProfileEntriesPredefinedEntryVariantJSON contains the JSON metadata
// for the struct [ProfileCustomProfileEntriesPredefinedEntryVariant]
type profileCustomProfileEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type ProfileCustomProfileEntriesPredefinedEntryVariantTopicType string

const (
	ProfileCustomProfileEntriesPredefinedEntryVariantTopicTypeIntent  ProfileCustomProfileEntriesPredefinedEntryVariantTopicType = "Intent"
	ProfileCustomProfileEntriesPredefinedEntryVariantTopicTypeContent ProfileCustomProfileEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r ProfileCustomProfileEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesPredefinedEntryVariantTopicTypeIntent, ProfileCustomProfileEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesPredefinedEntryVariantType string

const (
	ProfileCustomProfileEntriesPredefinedEntryVariantTypePromptTopic ProfileCustomProfileEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r ProfileCustomProfileEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesIntegrationEntry struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Type      ProfileCustomProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileEntriesIntegrationEntryJSON `json:"-"`
}

// profileCustomProfileEntriesIntegrationEntryJSON contains the JSON metadata for
// the struct [ProfileCustomProfileEntriesIntegrationEntry]
type profileCustomProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesIntegrationEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesIntegrationEntryType string

const (
	ProfileCustomProfileEntriesIntegrationEntryTypeIntegration ProfileCustomProfileEntriesIntegrationEntryType = "integration"
)

func (r ProfileCustomProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                          `json:"case_sensitive,required"`
	CreatedAt     time.Time                                     `json:"created_at,required" format:"date-time"`
	Enabled       bool                                          `json:"enabled,required"`
	Name          string                                        `json:"name,required"`
	Secret        bool                                          `json:"secret,required"`
	Type          ProfileCustomProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                     `json:"updated_at,required" format:"date-time"`
	JSON          profileCustomProfileEntriesExactDataEntryJSON `json:"-"`
}

// profileCustomProfileEntriesExactDataEntryJSON contains the JSON metadata for the
// struct [ProfileCustomProfileEntriesExactDataEntry]
type profileCustomProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesExactDataEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesExactDataEntryType string

const (
	ProfileCustomProfileEntriesExactDataEntryTypeExactData ProfileCustomProfileEntriesExactDataEntryType = "exact_data"
)

func (r ProfileCustomProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesDocumentFingerprintEntry struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      ProfileCustomProfileEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	JSON      profileCustomProfileEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profileCustomProfileEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct [ProfileCustomProfileEntriesDocumentFingerprintEntry]
type profileCustomProfileEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesDocumentFingerprintEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesDocumentFingerprintEntryType string

const (
	ProfileCustomProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfileCustomProfileEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfileCustomProfileEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesWordListEntry struct {
	ID        string                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                         `json:"enabled,required"`
	Name      string                                       `json:"name,required"`
	Type      ProfileCustomProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                    `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                  `json:"word_list,required"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileEntriesWordListEntryJSON `json:"-"`
}

// profileCustomProfileEntriesWordListEntryJSON contains the JSON metadata for the
// struct [ProfileCustomProfileEntriesWordListEntry]
type profileCustomProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileEntriesWordListEntry) implementsProfileCustomProfileEntry() {}

type ProfileCustomProfileEntriesWordListEntryType string

const (
	ProfileCustomProfileEntriesWordListEntryTypeWordList ProfileCustomProfileEntriesWordListEntryType = "word_list"
)

func (r ProfileCustomProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfileCustomProfileEntriesType string

const (
	ProfileCustomProfileEntriesTypeCustom              ProfileCustomProfileEntriesType = "custom"
	ProfileCustomProfileEntriesTypePredefined          ProfileCustomProfileEntriesType = "predefined"
	ProfileCustomProfileEntriesTypeIntegration         ProfileCustomProfileEntriesType = "integration"
	ProfileCustomProfileEntriesTypeExactData           ProfileCustomProfileEntriesType = "exact_data"
	ProfileCustomProfileEntriesTypeDocumentFingerprint ProfileCustomProfileEntriesType = "document_fingerprint"
	ProfileCustomProfileEntriesTypeWordList            ProfileCustomProfileEntriesType = "word_list"
)

func (r ProfileCustomProfileEntriesType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileEntriesTypeCustom, ProfileCustomProfileEntriesTypePredefined, ProfileCustomProfileEntriesTypeIntegration, ProfileCustomProfileEntriesTypeExactData, ProfileCustomProfileEntriesTypeDocumentFingerprint, ProfileCustomProfileEntriesTypeWordList:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                  `json:"enabled,required"`
	Name    string                                `json:"name,required"`
	Type    ProfileCustomProfileSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfileCustomProfileSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [ProfileCustomProfileSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                         `json:"word_list"`
	JSON     profileCustomProfileSharedEntryJSON `json:"-"`
	union    ProfileCustomProfileSharedEntriesUnion
}

// profileCustomProfileSharedEntryJSON contains the JSON metadata for the struct
// [ProfileCustomProfileSharedEntry]
type profileCustomProfileSharedEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profileCustomProfileSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileCustomProfileSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileCustomProfileSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileCustomProfileSharedEntriesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfileCustomProfileSharedEntriesCustomEntry],
// [ProfileCustomProfileSharedEntriesPredefinedEntry],
// [ProfileCustomProfileSharedEntriesIntegrationEntry],
// [ProfileCustomProfileSharedEntriesExactDataEntry],
// [ProfileCustomProfileSharedEntriesDocumentFingerprintEntry],
// [ProfileCustomProfileSharedEntriesWordListEntry].
func (r ProfileCustomProfileSharedEntry) AsUnion() ProfileCustomProfileSharedEntriesUnion {
	return r.union
}

// Union satisfied by [ProfileCustomProfileSharedEntriesCustomEntry],
// [ProfileCustomProfileSharedEntriesPredefinedEntry],
// [ProfileCustomProfileSharedEntriesIntegrationEntry],
// [ProfileCustomProfileSharedEntriesExactDataEntry],
// [ProfileCustomProfileSharedEntriesDocumentFingerprintEntry] or
// [ProfileCustomProfileSharedEntriesWordListEntry].
type ProfileCustomProfileSharedEntriesUnion interface {
	implementsProfileCustomProfileSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileCustomProfileSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileCustomProfileSharedEntriesWordListEntry{}),
		},
	)
}

type ProfileCustomProfileSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                             `json:"enabled,required"`
	Name        string                                           `json:"name,required"`
	Pattern     Pattern                                          `json:"pattern,required"`
	Type        ProfileCustomProfileSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                        `json:"updated_at,required" format:"date-time"`
	Description string                                           `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileSharedEntriesCustomEntryJSON `json:"-"`
}

// profileCustomProfileSharedEntriesCustomEntryJSON contains the JSON metadata for
// the struct [ProfileCustomProfileSharedEntriesCustomEntry]
type profileCustomProfileSharedEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileSharedEntriesCustomEntry) implementsProfileCustomProfileSharedEntry() {}

type ProfileCustomProfileSharedEntriesCustomEntryType string

const (
	ProfileCustomProfileSharedEntriesCustomEntryTypeCustom ProfileCustomProfileSharedEntriesCustomEntryType = "custom"
)

func (r ProfileCustomProfileSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesPredefinedEntry struct {
	ID         string                                                     `json:"id,required" format:"uuid"`
	Confidence ProfileCustomProfileSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                       `json:"enabled,required"`
	Name       string                                                     `json:"name,required"`
	Type       ProfileCustomProfileSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	Variant   ProfileCustomProfileSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      profileCustomProfileSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// profileCustomProfileSharedEntriesPredefinedEntryJSON contains the JSON metadata
// for the struct [ProfileCustomProfileSharedEntriesPredefinedEntry]
type profileCustomProfileSharedEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileSharedEntriesPredefinedEntry) implementsProfileCustomProfileSharedEntry() {
}

type ProfileCustomProfileSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                           `json:"available,required"`
	JSON      profileCustomProfileSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profileCustomProfileSharedEntriesPredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct
// [ProfileCustomProfileSharedEntriesPredefinedEntryConfidence]
type profileCustomProfileSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileCustomProfileSharedEntriesPredefinedEntryType string

const (
	ProfileCustomProfileSharedEntriesPredefinedEntryTypePredefined ProfileCustomProfileSharedEntriesPredefinedEntryType = "predefined"
)

func (r ProfileCustomProfileSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesPredefinedEntryVariant struct {
	TopicType   ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        ProfileCustomProfileSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                           `json:"description,nullable"`
	JSON        profileCustomProfileSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// profileCustomProfileSharedEntriesPredefinedEntryVariantJSON contains the JSON
// metadata for the struct
// [ProfileCustomProfileSharedEntriesPredefinedEntryVariant]
type profileCustomProfileSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicType string

const (
	ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicTypeIntent  ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicTypeContent ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicTypeIntent, ProfileCustomProfileSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesPredefinedEntryVariantType string

const (
	ProfileCustomProfileSharedEntriesPredefinedEntryVariantTypePromptTopic ProfileCustomProfileSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r ProfileCustomProfileSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesIntegrationEntry struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Type      ProfileCustomProfileSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileSharedEntriesIntegrationEntryJSON `json:"-"`
}

// profileCustomProfileSharedEntriesIntegrationEntryJSON contains the JSON metadata
// for the struct [ProfileCustomProfileSharedEntriesIntegrationEntry]
type profileCustomProfileSharedEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileSharedEntriesIntegrationEntry) implementsProfileCustomProfileSharedEntry() {
}

type ProfileCustomProfileSharedEntriesIntegrationEntryType string

const (
	ProfileCustomProfileSharedEntriesIntegrationEntryTypeIntegration ProfileCustomProfileSharedEntriesIntegrationEntryType = "integration"
)

func (r ProfileCustomProfileSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                `json:"case_sensitive,required"`
	CreatedAt     time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                `json:"enabled,required"`
	Name          string                                              `json:"name,required"`
	Secret        bool                                                `json:"secret,required"`
	Type          ProfileCustomProfileSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                           `json:"updated_at,required" format:"date-time"`
	JSON          profileCustomProfileSharedEntriesExactDataEntryJSON `json:"-"`
}

// profileCustomProfileSharedEntriesExactDataEntryJSON contains the JSON metadata
// for the struct [ProfileCustomProfileSharedEntriesExactDataEntry]
type profileCustomProfileSharedEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileSharedEntriesExactDataEntry) implementsProfileCustomProfileSharedEntry() {
}

type ProfileCustomProfileSharedEntriesExactDataEntryType string

const (
	ProfileCustomProfileSharedEntriesExactDataEntryTypeExactData ProfileCustomProfileSharedEntriesExactDataEntryType = "exact_data"
)

func (r ProfileCustomProfileSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Type      ProfileCustomProfileSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	JSON      profileCustomProfileSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profileCustomProfileSharedEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct
// [ProfileCustomProfileSharedEntriesDocumentFingerprintEntry]
type profileCustomProfileSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileSharedEntriesDocumentFingerprintEntry) implementsProfileCustomProfileSharedEntry() {
}

type ProfileCustomProfileSharedEntriesDocumentFingerprintEntryType string

const (
	ProfileCustomProfileSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfileCustomProfileSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfileCustomProfileSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesWordListEntry struct {
	ID        string                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                               `json:"enabled,required"`
	Name      string                                             `json:"name,required"`
	Type      ProfileCustomProfileSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                          `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                        `json:"word_list,required"`
	ProfileID string                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomProfileSharedEntriesWordListEntryJSON `json:"-"`
}

// profileCustomProfileSharedEntriesWordListEntryJSON contains the JSON metadata
// for the struct [ProfileCustomProfileSharedEntriesWordListEntry]
type profileCustomProfileSharedEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomProfileSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomProfileSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomProfileSharedEntriesWordListEntry) implementsProfileCustomProfileSharedEntry() {}

type ProfileCustomProfileSharedEntriesWordListEntryType string

const (
	ProfileCustomProfileSharedEntriesWordListEntryTypeWordList ProfileCustomProfileSharedEntriesWordListEntryType = "word_list"
)

func (r ProfileCustomProfileSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfileCustomProfileSharedEntriesType string

const (
	ProfileCustomProfileSharedEntriesTypeCustom              ProfileCustomProfileSharedEntriesType = "custom"
	ProfileCustomProfileSharedEntriesTypePredefined          ProfileCustomProfileSharedEntriesType = "predefined"
	ProfileCustomProfileSharedEntriesTypeIntegration         ProfileCustomProfileSharedEntriesType = "integration"
	ProfileCustomProfileSharedEntriesTypeExactData           ProfileCustomProfileSharedEntriesType = "exact_data"
	ProfileCustomProfileSharedEntriesTypeDocumentFingerprint ProfileCustomProfileSharedEntriesType = "document_fingerprint"
	ProfileCustomProfileSharedEntriesTypeWordList            ProfileCustomProfileSharedEntriesType = "word_list"
)

func (r ProfileCustomProfileSharedEntriesType) IsKnown() bool {
	switch r {
	case ProfileCustomProfileSharedEntriesTypeCustom, ProfileCustomProfileSharedEntriesTypePredefined, ProfileCustomProfileSharedEntriesTypeIntegration, ProfileCustomProfileSharedEntriesTypeExactData, ProfileCustomProfileSharedEntriesTypeDocumentFingerprint, ProfileCustomProfileSharedEntriesTypeWordList:
		return true
	}
	return false
}

type ProfilePredefinedProfile struct {
	// The id of the predefined profile (uuid).
	ID                string `json:"id,required" format:"uuid"`
	AllowedMatchCount int64  `json:"allowed_match_count,required"`
	// Deprecated: deprecated
	Entries []ProfilePredefinedProfileEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                                      `json:"name,required"`
	Type                ProfilePredefinedProfileType                `json:"type,required"`
	AIContextEnabled    bool                                        `json:"ai_context_enabled"`
	ConfidenceThreshold ProfilePredefinedProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                         `json:"open_access"`
	JSON       profilePredefinedProfileJSON `json:"-"`
}

// profilePredefinedProfileJSON contains the JSON metadata for the struct
// [ProfilePredefinedProfile]
type profilePredefinedProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfilePredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfile) implementsProfile() {}

type ProfilePredefinedProfileEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                `json:"enabled,required"`
	Name    string                              `json:"name,required"`
	Type    ProfilePredefinedProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfilePredefinedProfileEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [ProfilePredefinedProfileEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                       `json:"word_list"`
	JSON     profilePredefinedProfileEntryJSON `json:"-"`
	union    ProfilePredefinedProfileEntriesUnion
}

// profilePredefinedProfileEntryJSON contains the JSON metadata for the struct
// [ProfilePredefinedProfileEntry]
type profilePredefinedProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profilePredefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfilePredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfilePredefinedProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfilePredefinedProfileEntriesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfilePredefinedProfileEntriesCustomEntry],
// [ProfilePredefinedProfileEntriesPredefinedEntry],
// [ProfilePredefinedProfileEntriesIntegrationEntry],
// [ProfilePredefinedProfileEntriesExactDataEntry],
// [ProfilePredefinedProfileEntriesDocumentFingerprintEntry],
// [ProfilePredefinedProfileEntriesWordListEntry].
func (r ProfilePredefinedProfileEntry) AsUnion() ProfilePredefinedProfileEntriesUnion {
	return r.union
}

// Union satisfied by [ProfilePredefinedProfileEntriesCustomEntry],
// [ProfilePredefinedProfileEntriesPredefinedEntry],
// [ProfilePredefinedProfileEntriesIntegrationEntry],
// [ProfilePredefinedProfileEntriesExactDataEntry],
// [ProfilePredefinedProfileEntriesDocumentFingerprintEntry] or
// [ProfilePredefinedProfileEntriesWordListEntry].
type ProfilePredefinedProfileEntriesUnion interface {
	implementsProfilePredefinedProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfilePredefinedProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfilePredefinedProfileEntriesWordListEntry{}),
		},
	)
}

type ProfilePredefinedProfileEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                           `json:"enabled,required"`
	Name        string                                         `json:"name,required"`
	Pattern     Pattern                                        `json:"pattern,required"`
	Type        ProfilePredefinedProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                      `json:"updated_at,required" format:"date-time"`
	Description string                                         `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedProfileEntriesCustomEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesCustomEntryJSON contains the JSON metadata for
// the struct [ProfilePredefinedProfileEntriesCustomEntry]
type profilePredefinedProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesCustomEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesCustomEntryType string

const (
	ProfilePredefinedProfileEntriesCustomEntryTypeCustom ProfilePredefinedProfileEntriesCustomEntryType = "custom"
)

func (r ProfilePredefinedProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesPredefinedEntry struct {
	ID         string                                                   `json:"id,required" format:"uuid"`
	Confidence ProfilePredefinedProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                     `json:"enabled,required"`
	Name       string                                                   `json:"name,required"`
	Type       ProfilePredefinedProfileEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	Variant   ProfilePredefinedProfileEntriesPredefinedEntryVariant `json:"variant"`
	JSON      profilePredefinedProfileEntriesPredefinedEntryJSON    `json:"-"`
}

// profilePredefinedProfileEntriesPredefinedEntryJSON contains the JSON metadata
// for the struct [ProfilePredefinedProfileEntriesPredefinedEntry]
type profilePredefinedProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesPredefinedEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                         `json:"available,required"`
	JSON      profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct
// [ProfilePredefinedProfileEntriesPredefinedEntryConfidence]
type profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfilePredefinedProfileEntriesPredefinedEntryType string

const (
	ProfilePredefinedProfileEntriesPredefinedEntryTypePredefined ProfilePredefinedProfileEntriesPredefinedEntryType = "predefined"
)

func (r ProfilePredefinedProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesPredefinedEntryVariant struct {
	TopicType   ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        ProfilePredefinedProfileEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                         `json:"description,nullable"`
	JSON        profilePredefinedProfileEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// profilePredefinedProfileEntriesPredefinedEntryVariantJSON contains the JSON
// metadata for the struct [ProfilePredefinedProfileEntriesPredefinedEntryVariant]
type profilePredefinedProfileEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicType string

const (
	ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicTypeIntent  ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicType = "Intent"
	ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicTypeContent ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicTypeIntent, ProfilePredefinedProfileEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesPredefinedEntryVariantType string

const (
	ProfilePredefinedProfileEntriesPredefinedEntryVariantTypePromptTopic ProfilePredefinedProfileEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r ProfilePredefinedProfileEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesIntegrationEntry struct {
	ID        string                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                `json:"enabled,required"`
	Name      string                                              `json:"name,required"`
	Type      ProfilePredefinedProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedProfileEntriesIntegrationEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesIntegrationEntryJSON contains the JSON metadata
// for the struct [ProfilePredefinedProfileEntriesIntegrationEntry]
type profilePredefinedProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesIntegrationEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesIntegrationEntryType string

const (
	ProfilePredefinedProfileEntriesIntegrationEntryTypeIntegration ProfilePredefinedProfileEntriesIntegrationEntryType = "integration"
)

func (r ProfilePredefinedProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                              `json:"case_sensitive,required"`
	CreatedAt     time.Time                                         `json:"created_at,required" format:"date-time"`
	Enabled       bool                                              `json:"enabled,required"`
	Name          string                                            `json:"name,required"`
	Secret        bool                                              `json:"secret,required"`
	Type          ProfilePredefinedProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                         `json:"updated_at,required" format:"date-time"`
	JSON          profilePredefinedProfileEntriesExactDataEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesExactDataEntryJSON contains the JSON metadata for
// the struct [ProfilePredefinedProfileEntriesExactDataEntry]
type profilePredefinedProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesExactDataEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesExactDataEntryType string

const (
	ProfilePredefinedProfileEntriesExactDataEntryTypeExactData ProfilePredefinedProfileEntriesExactDataEntryType = "exact_data"
)

func (r ProfilePredefinedProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesDocumentFingerprintEntry struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      ProfilePredefinedProfileEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	JSON      profilePredefinedProfileEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct
// [ProfilePredefinedProfileEntriesDocumentFingerprintEntry]
type profilePredefinedProfileEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesDocumentFingerprintEntry) implementsProfilePredefinedProfileEntry() {
}

type ProfilePredefinedProfileEntriesDocumentFingerprintEntryType string

const (
	ProfilePredefinedProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfilePredefinedProfileEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfilePredefinedProfileEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesWordListEntry struct {
	ID        string                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                             `json:"enabled,required"`
	Name      string                                           `json:"name,required"`
	Type      ProfilePredefinedProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                      `json:"word_list,required"`
	ProfileID string                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedProfileEntriesWordListEntryJSON `json:"-"`
}

// profilePredefinedProfileEntriesWordListEntryJSON contains the JSON metadata for
// the struct [ProfilePredefinedProfileEntriesWordListEntry]
type profilePredefinedProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedProfileEntriesWordListEntry) implementsProfilePredefinedProfileEntry() {}

type ProfilePredefinedProfileEntriesWordListEntryType string

const (
	ProfilePredefinedProfileEntriesWordListEntryTypeWordList ProfilePredefinedProfileEntriesWordListEntryType = "word_list"
)

func (r ProfilePredefinedProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfilePredefinedProfileEntriesType string

const (
	ProfilePredefinedProfileEntriesTypeCustom              ProfilePredefinedProfileEntriesType = "custom"
	ProfilePredefinedProfileEntriesTypePredefined          ProfilePredefinedProfileEntriesType = "predefined"
	ProfilePredefinedProfileEntriesTypeIntegration         ProfilePredefinedProfileEntriesType = "integration"
	ProfilePredefinedProfileEntriesTypeExactData           ProfilePredefinedProfileEntriesType = "exact_data"
	ProfilePredefinedProfileEntriesTypeDocumentFingerprint ProfilePredefinedProfileEntriesType = "document_fingerprint"
	ProfilePredefinedProfileEntriesTypeWordList            ProfilePredefinedProfileEntriesType = "word_list"
)

func (r ProfilePredefinedProfileEntriesType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileEntriesTypeCustom, ProfilePredefinedProfileEntriesTypePredefined, ProfilePredefinedProfileEntriesTypeIntegration, ProfilePredefinedProfileEntriesTypeExactData, ProfilePredefinedProfileEntriesTypeDocumentFingerprint, ProfilePredefinedProfileEntriesTypeWordList:
		return true
	}
	return false
}

type ProfilePredefinedProfileType string

const (
	ProfilePredefinedProfileTypePredefined ProfilePredefinedProfileType = "predefined"
)

func (r ProfilePredefinedProfileType) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileTypePredefined:
		return true
	}
	return false
}

type ProfilePredefinedProfileConfidenceThreshold string

const (
	ProfilePredefinedProfileConfidenceThresholdLow      ProfilePredefinedProfileConfidenceThreshold = "low"
	ProfilePredefinedProfileConfidenceThresholdMedium   ProfilePredefinedProfileConfidenceThreshold = "medium"
	ProfilePredefinedProfileConfidenceThresholdHigh     ProfilePredefinedProfileConfidenceThreshold = "high"
	ProfilePredefinedProfileConfidenceThresholdVeryHigh ProfilePredefinedProfileConfidenceThreshold = "very_high"
)

func (r ProfilePredefinedProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfilePredefinedProfileConfidenceThresholdLow, ProfilePredefinedProfileConfidenceThresholdMedium, ProfilePredefinedProfileConfidenceThresholdHigh, ProfilePredefinedProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type ProfileIntegrationProfile struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Entries       []ProfileIntegrationProfileEntry       `json:"entries,required"`
	Name          string                                 `json:"name,required"`
	SharedEntries []ProfileIntegrationProfileSharedEntry `json:"shared_entries,required"`
	Type          ProfileIntegrationProfileType          `json:"type,required"`
	UpdatedAt     time.Time                              `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                        `json:"description,nullable"`
	JSON        profileIntegrationProfileJSON `json:"-"`
}

// profileIntegrationProfileJSON contains the JSON metadata for the struct
// [ProfileIntegrationProfile]
type profileIntegrationProfileJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Entries       apijson.Field
	Name          apijson.Field
	SharedEntries apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	Description   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfile) implementsProfile() {}

type ProfileIntegrationProfileEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                 `json:"enabled,required"`
	Name    string                               `json:"name,required"`
	Type    ProfileIntegrationProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfileIntegrationProfileEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [ProfileIntegrationProfileEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                        `json:"word_list"`
	JSON     profileIntegrationProfileEntryJSON `json:"-"`
	union    ProfileIntegrationProfileEntriesUnion
}

// profileIntegrationProfileEntryJSON contains the JSON metadata for the struct
// [ProfileIntegrationProfileEntry]
type profileIntegrationProfileEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profileIntegrationProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileIntegrationProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileIntegrationProfileEntriesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfileIntegrationProfileEntriesCustomEntry],
// [ProfileIntegrationProfileEntriesPredefinedEntry],
// [ProfileIntegrationProfileEntriesIntegrationEntry],
// [ProfileIntegrationProfileEntriesExactDataEntry],
// [ProfileIntegrationProfileEntriesDocumentFingerprintEntry],
// [ProfileIntegrationProfileEntriesWordListEntry].
func (r ProfileIntegrationProfileEntry) AsUnion() ProfileIntegrationProfileEntriesUnion {
	return r.union
}

// Union satisfied by [ProfileIntegrationProfileEntriesCustomEntry],
// [ProfileIntegrationProfileEntriesPredefinedEntry],
// [ProfileIntegrationProfileEntriesIntegrationEntry],
// [ProfileIntegrationProfileEntriesExactDataEntry],
// [ProfileIntegrationProfileEntriesDocumentFingerprintEntry] or
// [ProfileIntegrationProfileEntriesWordListEntry].
type ProfileIntegrationProfileEntriesUnion interface {
	implementsProfileIntegrationProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileIntegrationProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileEntriesWordListEntry{}),
		},
	)
}

type ProfileIntegrationProfileEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                            `json:"enabled,required"`
	Name        string                                          `json:"name,required"`
	Pattern     Pattern                                         `json:"pattern,required"`
	Type        ProfileIntegrationProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                       `json:"updated_at,required" format:"date-time"`
	Description string                                          `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileEntriesCustomEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesCustomEntryJSON contains the JSON metadata for
// the struct [ProfileIntegrationProfileEntriesCustomEntry]
type profileIntegrationProfileEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesCustomEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesCustomEntryType string

const (
	ProfileIntegrationProfileEntriesCustomEntryTypeCustom ProfileIntegrationProfileEntriesCustomEntryType = "custom"
)

func (r ProfileIntegrationProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesPredefinedEntry struct {
	ID         string                                                    `json:"id,required" format:"uuid"`
	Confidence ProfileIntegrationProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                      `json:"enabled,required"`
	Name       string                                                    `json:"name,required"`
	Type       ProfileIntegrationProfileEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	Variant   ProfileIntegrationProfileEntriesPredefinedEntryVariant `json:"variant"`
	JSON      profileIntegrationProfileEntriesPredefinedEntryJSON    `json:"-"`
}

// profileIntegrationProfileEntriesPredefinedEntryJSON contains the JSON metadata
// for the struct [ProfileIntegrationProfileEntriesPredefinedEntry]
type profileIntegrationProfileEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesPredefinedEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                          `json:"available,required"`
	JSON      profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct
// [ProfileIntegrationProfileEntriesPredefinedEntryConfidence]
type profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileIntegrationProfileEntriesPredefinedEntryType string

const (
	ProfileIntegrationProfileEntriesPredefinedEntryTypePredefined ProfileIntegrationProfileEntriesPredefinedEntryType = "predefined"
)

func (r ProfileIntegrationProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesPredefinedEntryVariant struct {
	TopicType   ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        ProfileIntegrationProfileEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                          `json:"description,nullable"`
	JSON        profileIntegrationProfileEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// profileIntegrationProfileEntriesPredefinedEntryVariantJSON contains the JSON
// metadata for the struct [ProfileIntegrationProfileEntriesPredefinedEntryVariant]
type profileIntegrationProfileEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicType string

const (
	ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicTypeIntent  ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicType = "Intent"
	ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicTypeContent ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicTypeIntent, ProfileIntegrationProfileEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesPredefinedEntryVariantType string

const (
	ProfileIntegrationProfileEntriesPredefinedEntryVariantTypePromptTopic ProfileIntegrationProfileEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r ProfileIntegrationProfileEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesIntegrationEntry struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      ProfileIntegrationProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	ProfileID string                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileEntriesIntegrationEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesIntegrationEntryJSON contains the JSON metadata
// for the struct [ProfileIntegrationProfileEntriesIntegrationEntry]
type profileIntegrationProfileEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesIntegrationEntry) implementsProfileIntegrationProfileEntry() {
}

type ProfileIntegrationProfileEntriesIntegrationEntryType string

const (
	ProfileIntegrationProfileEntriesIntegrationEntryTypeIntegration ProfileIntegrationProfileEntriesIntegrationEntryType = "integration"
)

func (r ProfileIntegrationProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                               `json:"case_sensitive,required"`
	CreatedAt     time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled       bool                                               `json:"enabled,required"`
	Name          string                                             `json:"name,required"`
	Secret        bool                                               `json:"secret,required"`
	Type          ProfileIntegrationProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                          `json:"updated_at,required" format:"date-time"`
	JSON          profileIntegrationProfileEntriesExactDataEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesExactDataEntryJSON contains the JSON metadata
// for the struct [ProfileIntegrationProfileEntriesExactDataEntry]
type profileIntegrationProfileEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesExactDataEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesExactDataEntryType string

const (
	ProfileIntegrationProfileEntriesExactDataEntryTypeExactData ProfileIntegrationProfileEntriesExactDataEntryType = "exact_data"
)

func (r ProfileIntegrationProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesDocumentFingerprintEntry struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      ProfileIntegrationProfileEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	JSON      profileIntegrationProfileEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct
// [ProfileIntegrationProfileEntriesDocumentFingerprintEntry]
type profileIntegrationProfileEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesDocumentFingerprintEntry) implementsProfileIntegrationProfileEntry() {
}

type ProfileIntegrationProfileEntriesDocumentFingerprintEntryType string

const (
	ProfileIntegrationProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfileIntegrationProfileEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfileIntegrationProfileEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesWordListEntry struct {
	ID        string                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                              `json:"enabled,required"`
	Name      string                                            `json:"name,required"`
	Type      ProfileIntegrationProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                         `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                       `json:"word_list,required"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileEntriesWordListEntryJSON `json:"-"`
}

// profileIntegrationProfileEntriesWordListEntryJSON contains the JSON metadata for
// the struct [ProfileIntegrationProfileEntriesWordListEntry]
type profileIntegrationProfileEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileEntriesWordListEntry) implementsProfileIntegrationProfileEntry() {}

type ProfileIntegrationProfileEntriesWordListEntryType string

const (
	ProfileIntegrationProfileEntriesWordListEntryTypeWordList ProfileIntegrationProfileEntriesWordListEntryType = "word_list"
)

func (r ProfileIntegrationProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationProfileEntriesType string

const (
	ProfileIntegrationProfileEntriesTypeCustom              ProfileIntegrationProfileEntriesType = "custom"
	ProfileIntegrationProfileEntriesTypePredefined          ProfileIntegrationProfileEntriesType = "predefined"
	ProfileIntegrationProfileEntriesTypeIntegration         ProfileIntegrationProfileEntriesType = "integration"
	ProfileIntegrationProfileEntriesTypeExactData           ProfileIntegrationProfileEntriesType = "exact_data"
	ProfileIntegrationProfileEntriesTypeDocumentFingerprint ProfileIntegrationProfileEntriesType = "document_fingerprint"
	ProfileIntegrationProfileEntriesTypeWordList            ProfileIntegrationProfileEntriesType = "word_list"
)

func (r ProfileIntegrationProfileEntriesType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileEntriesTypeCustom, ProfileIntegrationProfileEntriesTypePredefined, ProfileIntegrationProfileEntriesTypeIntegration, ProfileIntegrationProfileEntriesTypeExactData, ProfileIntegrationProfileEntriesTypeDocumentFingerprint, ProfileIntegrationProfileEntriesTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                       `json:"enabled,required"`
	Name    string                                     `json:"name,required"`
	Type    ProfileIntegrationProfileSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [ProfileIntegrationProfileSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [ProfileIntegrationProfileSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                              `json:"word_list"`
	JSON     profileIntegrationProfileSharedEntryJSON `json:"-"`
	union    ProfileIntegrationProfileSharedEntriesUnion
}

// profileIntegrationProfileSharedEntryJSON contains the JSON metadata for the
// struct [ProfileIntegrationProfileSharedEntry]
type profileIntegrationProfileSharedEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r profileIntegrationProfileSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileIntegrationProfileSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileIntegrationProfileSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileIntegrationProfileSharedEntriesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProfileIntegrationProfileSharedEntriesCustomEntry],
// [ProfileIntegrationProfileSharedEntriesPredefinedEntry],
// [ProfileIntegrationProfileSharedEntriesIntegrationEntry],
// [ProfileIntegrationProfileSharedEntriesExactDataEntry],
// [ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntry],
// [ProfileIntegrationProfileSharedEntriesWordListEntry].
func (r ProfileIntegrationProfileSharedEntry) AsUnion() ProfileIntegrationProfileSharedEntriesUnion {
	return r.union
}

// Union satisfied by [ProfileIntegrationProfileSharedEntriesCustomEntry],
// [ProfileIntegrationProfileSharedEntriesPredefinedEntry],
// [ProfileIntegrationProfileSharedEntriesIntegrationEntry],
// [ProfileIntegrationProfileSharedEntriesExactDataEntry],
// [ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntry] or
// [ProfileIntegrationProfileSharedEntriesWordListEntry].
type ProfileIntegrationProfileSharedEntriesUnion interface {
	implementsProfileIntegrationProfileSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileIntegrationProfileSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProfileIntegrationProfileSharedEntriesWordListEntry{}),
		},
	)
}

type ProfileIntegrationProfileSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                  `json:"enabled,required"`
	Name        string                                                `json:"name,required"`
	Pattern     Pattern                                               `json:"pattern,required"`
	Type        ProfileIntegrationProfileSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                             `json:"updated_at,required" format:"date-time"`
	Description string                                                `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileSharedEntriesCustomEntryJSON `json:"-"`
}

// profileIntegrationProfileSharedEntriesCustomEntryJSON contains the JSON metadata
// for the struct [ProfileIntegrationProfileSharedEntriesCustomEntry]
type profileIntegrationProfileSharedEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileSharedEntriesCustomEntry) implementsProfileIntegrationProfileSharedEntry() {
}

type ProfileIntegrationProfileSharedEntriesCustomEntryType string

const (
	ProfileIntegrationProfileSharedEntriesCustomEntryTypeCustom ProfileIntegrationProfileSharedEntriesCustomEntryType = "custom"
)

func (r ProfileIntegrationProfileSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesPredefinedEntry struct {
	ID         string                                                          `json:"id,required" format:"uuid"`
	Confidence ProfileIntegrationProfileSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                            `json:"enabled,required"`
	Name       string                                                          `json:"name,required"`
	Type       ProfileIntegrationProfileSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	Variant   ProfileIntegrationProfileSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      profileIntegrationProfileSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// profileIntegrationProfileSharedEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct [ProfileIntegrationProfileSharedEntriesPredefinedEntry]
type profileIntegrationProfileSharedEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileSharedEntriesPredefinedEntry) implementsProfileIntegrationProfileSharedEntry() {
}

type ProfileIntegrationProfileSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                `json:"available,required"`
	JSON      profileIntegrationProfileSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// profileIntegrationProfileSharedEntriesPredefinedEntryConfidenceJSON contains the
// JSON metadata for the struct
// [ProfileIntegrationProfileSharedEntriesPredefinedEntryConfidence]
type profileIntegrationProfileSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileIntegrationProfileSharedEntriesPredefinedEntryType string

const (
	ProfileIntegrationProfileSharedEntriesPredefinedEntryTypePredefined ProfileIntegrationProfileSharedEntriesPredefinedEntryType = "predefined"
)

func (r ProfileIntegrationProfileSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesPredefinedEntryVariant struct {
	TopicType   ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                `json:"description,nullable"`
	JSON        profileIntegrationProfileSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// profileIntegrationProfileSharedEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [ProfileIntegrationProfileSharedEntriesPredefinedEntryVariant]
type profileIntegrationProfileSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicType string

const (
	ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicTypeIntent  ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicTypeContent ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicTypeIntent, ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantType string

const (
	ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTypePromptTopic ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesIntegrationEntry struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      ProfileIntegrationProfileSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileSharedEntriesIntegrationEntryJSON `json:"-"`
}

// profileIntegrationProfileSharedEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct [ProfileIntegrationProfileSharedEntriesIntegrationEntry]
type profileIntegrationProfileSharedEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileSharedEntriesIntegrationEntry) implementsProfileIntegrationProfileSharedEntry() {
}

type ProfileIntegrationProfileSharedEntriesIntegrationEntryType string

const (
	ProfileIntegrationProfileSharedEntriesIntegrationEntryTypeIntegration ProfileIntegrationProfileSharedEntriesIntegrationEntryType = "integration"
)

func (r ProfileIntegrationProfileSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                     `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                     `json:"enabled,required"`
	Name          string                                                   `json:"name,required"`
	Secret        bool                                                     `json:"secret,required"`
	Type          ProfileIntegrationProfileSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                `json:"updated_at,required" format:"date-time"`
	JSON          profileIntegrationProfileSharedEntriesExactDataEntryJSON `json:"-"`
}

// profileIntegrationProfileSharedEntriesExactDataEntryJSON contains the JSON
// metadata for the struct [ProfileIntegrationProfileSharedEntriesExactDataEntry]
type profileIntegrationProfileSharedEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileSharedEntriesExactDataEntry) implementsProfileIntegrationProfileSharedEntry() {
}

type ProfileIntegrationProfileSharedEntriesExactDataEntryType string

const (
	ProfileIntegrationProfileSharedEntriesExactDataEntryTypeExactData ProfileIntegrationProfileSharedEntriesExactDataEntryType = "exact_data"
)

func (r ProfileIntegrationProfileSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	JSON      profileIntegrationProfileSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// profileIntegrationProfileSharedEntriesDocumentFingerprintEntryJSON contains the
// JSON metadata for the struct
// [ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntry]
type profileIntegrationProfileSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntry) implementsProfileIntegrationProfileSharedEntry() {
}

type ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntryType string

const (
	ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesWordListEntry struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      ProfileIntegrationProfileSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                             `json:"word_list,required"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationProfileSharedEntriesWordListEntryJSON `json:"-"`
}

// profileIntegrationProfileSharedEntriesWordListEntryJSON contains the JSON
// metadata for the struct [ProfileIntegrationProfileSharedEntriesWordListEntry]
type profileIntegrationProfileSharedEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationProfileSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationProfileSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationProfileSharedEntriesWordListEntry) implementsProfileIntegrationProfileSharedEntry() {
}

type ProfileIntegrationProfileSharedEntriesWordListEntryType string

const (
	ProfileIntegrationProfileSharedEntriesWordListEntryTypeWordList ProfileIntegrationProfileSharedEntriesWordListEntryType = "word_list"
)

func (r ProfileIntegrationProfileSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationProfileSharedEntriesType string

const (
	ProfileIntegrationProfileSharedEntriesTypeCustom              ProfileIntegrationProfileSharedEntriesType = "custom"
	ProfileIntegrationProfileSharedEntriesTypePredefined          ProfileIntegrationProfileSharedEntriesType = "predefined"
	ProfileIntegrationProfileSharedEntriesTypeIntegration         ProfileIntegrationProfileSharedEntriesType = "integration"
	ProfileIntegrationProfileSharedEntriesTypeExactData           ProfileIntegrationProfileSharedEntriesType = "exact_data"
	ProfileIntegrationProfileSharedEntriesTypeDocumentFingerprint ProfileIntegrationProfileSharedEntriesType = "document_fingerprint"
	ProfileIntegrationProfileSharedEntriesTypeWordList            ProfileIntegrationProfileSharedEntriesType = "word_list"
)

func (r ProfileIntegrationProfileSharedEntriesType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileSharedEntriesTypeCustom, ProfileIntegrationProfileSharedEntriesTypePredefined, ProfileIntegrationProfileSharedEntriesTypeIntegration, ProfileIntegrationProfileSharedEntriesTypeExactData, ProfileIntegrationProfileSharedEntriesTypeDocumentFingerprint, ProfileIntegrationProfileSharedEntriesTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationProfileType string

const (
	ProfileIntegrationProfileTypeIntegration ProfileIntegrationProfileType = "integration"
)

func (r ProfileIntegrationProfileType) IsKnown() bool {
	switch r {
	case ProfileIntegrationProfileTypeIntegration:
		return true
	}
	return false
}

type ProfileType string

const (
	ProfileTypeCustom      ProfileType = "custom"
	ProfileTypePredefined  ProfileType = "predefined"
	ProfileTypeIntegration ProfileType = "integration"
)

func (r ProfileType) IsKnown() bool {
	switch r {
	case ProfileTypeCustom, ProfileTypePredefined, ProfileTypeIntegration:
		return true
	}
	return false
}

type ProfileConfidenceThreshold string

const (
	ProfileConfidenceThresholdLow      ProfileConfidenceThreshold = "low"
	ProfileConfidenceThresholdMedium   ProfileConfidenceThreshold = "medium"
	ProfileConfidenceThresholdHigh     ProfileConfidenceThreshold = "high"
	ProfileConfidenceThresholdVeryHigh ProfileConfidenceThreshold = "very_high"
)

func (r ProfileConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfileConfidenceThresholdLow, ProfileConfidenceThresholdMedium, ProfileConfidenceThresholdHigh, ProfileConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

// Content types to exclude from context analysis and return all matches.
type SkipConfiguration struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files bool                  `json:"files,required"`
	JSON  skipConfigurationJSON `json:"-"`
}

// skipConfigurationJSON contains the JSON metadata for the struct
// [SkipConfiguration]
type skipConfigurationJSON struct {
	Files       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkipConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipConfigurationJSON) RawJSON() string {
	return r.raw
}

// Content types to exclude from context analysis and return all matches.
type SkipConfigurationParam struct {
	// If the content type is a file, skip context analysis and return all matches.
	Files param.Field[bool] `json:"files,required"`
}

func (r SkipConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileGetResponse struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile.
	Name             string                    `json:"name,required"`
	Type             DLPProfileGetResponseType `json:"type,required"`
	AIContextEnabled bool                      `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64                                    `json:"allowed_match_count"`
	ConfidenceThreshold DLPProfileGetResponseConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// This field can have the runtime type of [[]DLPProfileGetResponseCustomEntry],
	// [[]DLPProfileGetResponsePredefinedEntry],
	// [[]DLPProfileGetResponseIntegrationEntry].
	Entries    interface{} `json:"entries"`
	OCREnabled bool        `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool `json:"open_access"`
	// This field can have the runtime type of
	// [[]DLPProfileGetResponseCustomSharedEntry],
	// [[]DLPProfileGetResponseIntegrationSharedEntry].
	SharedEntries interface{} `json:"shared_entries"`
	// When the profile was lasted updated.
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      dlpProfileGetResponseJSON `json:"-"`
	union     DLPProfileGetResponseUnion
}

// dlpProfileGetResponseJSON contains the JSON metadata for the struct
// [DLPProfileGetResponse]
type dlpProfileGetResponseJSON struct {
	ID                  apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	AllowedMatchCount   apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Description         apijson.Field
	Entries             apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	SharedEntries       apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r dlpProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [DLPProfileGetResponseCustom],
// [DLPProfileGetResponsePredefined], [DLPProfileGetResponseIntegration].
func (r DLPProfileGetResponse) AsUnion() DLPProfileGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPProfileGetResponseCustom],
// [DLPProfileGetResponsePredefined] or [DLPProfileGetResponseIntegration].
type DLPProfileGetResponseUnion interface {
	implementsDLPProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileGetResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileGetResponseCustom struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the profile.
	Name       string                          `json:"name,required"`
	OCREnabled bool                            `json:"ocr_enabled,required"`
	Type       DLPProfileGetResponseCustomType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time                                      `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool                                           `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileGetResponseCustomConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// Deprecated: deprecated
	Entries       []DLPProfileGetResponseCustomEntry       `json:"entries"`
	SharedEntries []DLPProfileGetResponseCustomSharedEntry `json:"shared_entries"`
	JSON          dlpProfileGetResponseCustomJSON          `json:"-"`
}

// dlpProfileGetResponseCustomJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseCustom]
type dlpProfileGetResponseCustomJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	CreatedAt           apijson.Field
	Name                apijson.Field
	OCREnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	Description         apijson.Field
	Entries             apijson.Field
	SharedEntries       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustom) implementsDLPProfileGetResponse() {}

type DLPProfileGetResponseCustomType string

const (
	DLPProfileGetResponseCustomTypeCustom DLPProfileGetResponseCustomType = "custom"
)

func (r DLPProfileGetResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomConfidenceThreshold string

const (
	DLPProfileGetResponseCustomConfidenceThresholdLow      DLPProfileGetResponseCustomConfidenceThreshold = "low"
	DLPProfileGetResponseCustomConfidenceThresholdMedium   DLPProfileGetResponseCustomConfidenceThreshold = "medium"
	DLPProfileGetResponseCustomConfidenceThresholdHigh     DLPProfileGetResponseCustomConfidenceThreshold = "high"
	DLPProfileGetResponseCustomConfidenceThresholdVeryHigh DLPProfileGetResponseCustomConfidenceThreshold = "very_high"
)

func (r DLPProfileGetResponseCustomConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomConfidenceThresholdLow, DLPProfileGetResponseCustomConfidenceThresholdMedium, DLPProfileGetResponseCustomConfidenceThresholdHigh, DLPProfileGetResponseCustomConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                   `json:"enabled,required"`
	Name    string                                 `json:"name,required"`
	Type    DLPProfileGetResponseCustomEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseCustomEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseCustomEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                          `json:"word_list"`
	JSON     dlpProfileGetResponseCustomEntryJSON `json:"-"`
	union    DLPProfileGetResponseCustomEntriesUnion
}

// dlpProfileGetResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseCustomEntry]
type dlpProfileGetResponseCustomEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseCustomEntriesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileGetResponseCustomEntriesCustomEntry],
// [DLPProfileGetResponseCustomEntriesPredefinedEntry],
// [DLPProfileGetResponseCustomEntriesIntegrationEntry],
// [DLPProfileGetResponseCustomEntriesExactDataEntry],
// [DLPProfileGetResponseCustomEntriesDocumentFingerprintEntry],
// [DLPProfileGetResponseCustomEntriesWordListEntry].
func (r DLPProfileGetResponseCustomEntry) AsUnion() DLPProfileGetResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileGetResponseCustomEntriesCustomEntry],
// [DLPProfileGetResponseCustomEntriesPredefinedEntry],
// [DLPProfileGetResponseCustomEntriesIntegrationEntry],
// [DLPProfileGetResponseCustomEntriesExactDataEntry],
// [DLPProfileGetResponseCustomEntriesDocumentFingerprintEntry] or
// [DLPProfileGetResponseCustomEntriesWordListEntry].
type DLPProfileGetResponseCustomEntriesUnion interface {
	implementsDLPProfileGetResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseCustomEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomEntriesWordListEntry{}),
		},
	)
}

type DLPProfileGetResponseCustomEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                              `json:"enabled,required"`
	Name        string                                            `json:"name,required"`
	Pattern     Pattern                                           `json:"pattern,required"`
	Type        DLPProfileGetResponseCustomEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                         `json:"updated_at,required" format:"date-time"`
	Description string                                            `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesCustomEntryJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseCustomEntriesCustomEntry]
type dlpProfileGetResponseCustomEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesCustomEntry) implementsDLPProfileGetResponseCustomEntry() {}

type DLPProfileGetResponseCustomEntriesCustomEntryType string

const (
	DLPProfileGetResponseCustomEntriesCustomEntryTypeCustom DLPProfileGetResponseCustomEntriesCustomEntryType = "custom"
)

func (r DLPProfileGetResponseCustomEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesPredefinedEntry struct {
	ID         string                                                      `json:"id,required" format:"uuid"`
	Confidence DLPProfileGetResponseCustomEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                        `json:"enabled,required"`
	Name       string                                                      `json:"name,required"`
	Type       DLPProfileGetResponseCustomEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileGetResponseCustomEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileGetResponseCustomEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileGetResponseCustomEntriesPredefinedEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseCustomEntriesPredefinedEntry]
type dlpProfileGetResponseCustomEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesPredefinedEntry) implementsDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                            `json:"available,required"`
	JSON      dlpProfileGetResponseCustomEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesPredefinedEntryConfidenceJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseCustomEntriesPredefinedEntryConfidence]
type dlpProfileGetResponseCustomEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseCustomEntriesPredefinedEntryType string

const (
	DLPProfileGetResponseCustomEntriesPredefinedEntryTypePredefined DLPProfileGetResponseCustomEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileGetResponseCustomEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileGetResponseCustomEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                            `json:"description,nullable"`
	JSON        dlpProfileGetResponseCustomEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileGetResponseCustomEntriesPredefinedEntryVariantJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseCustomEntriesPredefinedEntryVariant]
type dlpProfileGetResponseCustomEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicTypeContent DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesPredefinedEntryVariantType string

const (
	DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTypePromptTopic DLPProfileGetResponseCustomEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileGetResponseCustomEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesIntegrationEntry struct {
	ID        string                                                 `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                              `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                   `json:"enabled,required"`
	Name      string                                                 `json:"name,required"`
	Type      DLPProfileGetResponseCustomEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                              `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseCustomEntriesIntegrationEntry]
type dlpProfileGetResponseCustomEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesIntegrationEntry) implementsDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesIntegrationEntryType string

const (
	DLPProfileGetResponseCustomEntriesIntegrationEntryTypeIntegration DLPProfileGetResponseCustomEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileGetResponseCustomEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                 `json:"case_sensitive,required"`
	CreatedAt     time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                 `json:"enabled,required"`
	Name          string                                               `json:"name,required"`
	Secret        bool                                                 `json:"secret,required"`
	Type          DLPProfileGetResponseCustomEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                            `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileGetResponseCustomEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesExactDataEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseCustomEntriesExactDataEntry]
type dlpProfileGetResponseCustomEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesExactDataEntry) implementsDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesExactDataEntryType string

const (
	DLPProfileGetResponseCustomEntriesExactDataEntryTypeExactData DLPProfileGetResponseCustomEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileGetResponseCustomEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesDocumentFingerprintEntry struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfileGetResponseCustomEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponseCustomEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseCustomEntriesDocumentFingerprintEntry]
type dlpProfileGetResponseCustomEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesDocumentFingerprintEntry) implementsDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesDocumentFingerprintEntryType string

const (
	DLPProfileGetResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileGetResponseCustomEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileGetResponseCustomEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesWordListEntry struct {
	ID        string                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                `json:"enabled,required"`
	Name      string                                              `json:"name,required"`
	Type      DLPProfileGetResponseCustomEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                           `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                         `json:"word_list,required"`
	ProfileID string                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomEntriesWordListEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponseCustomEntriesWordListEntry]
type dlpProfileGetResponseCustomEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomEntriesWordListEntry) implementsDLPProfileGetResponseCustomEntry() {
}

type DLPProfileGetResponseCustomEntriesWordListEntryType string

const (
	DLPProfileGetResponseCustomEntriesWordListEntryTypeWordList DLPProfileGetResponseCustomEntriesWordListEntryType = "word_list"
)

func (r DLPProfileGetResponseCustomEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomEntriesType string

const (
	DLPProfileGetResponseCustomEntriesTypeCustom              DLPProfileGetResponseCustomEntriesType = "custom"
	DLPProfileGetResponseCustomEntriesTypePredefined          DLPProfileGetResponseCustomEntriesType = "predefined"
	DLPProfileGetResponseCustomEntriesTypeIntegration         DLPProfileGetResponseCustomEntriesType = "integration"
	DLPProfileGetResponseCustomEntriesTypeExactData           DLPProfileGetResponseCustomEntriesType = "exact_data"
	DLPProfileGetResponseCustomEntriesTypeDocumentFingerprint DLPProfileGetResponseCustomEntriesType = "document_fingerprint"
	DLPProfileGetResponseCustomEntriesTypeWordList            DLPProfileGetResponseCustomEntriesType = "word_list"
)

func (r DLPProfileGetResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomEntriesTypeCustom, DLPProfileGetResponseCustomEntriesTypePredefined, DLPProfileGetResponseCustomEntriesTypeIntegration, DLPProfileGetResponseCustomEntriesTypeExactData, DLPProfileGetResponseCustomEntriesTypeDocumentFingerprint, DLPProfileGetResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                         `json:"enabled,required"`
	Name    string                                       `json:"name,required"`
	Type    DLPProfileGetResponseCustomSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseCustomSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                `json:"word_list"`
	JSON     dlpProfileGetResponseCustomSharedEntryJSON `json:"-"`
	union    DLPProfileGetResponseCustomSharedEntriesUnion
}

// dlpProfileGetResponseCustomSharedEntryJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseCustomSharedEntry]
type dlpProfileGetResponseCustomSharedEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileGetResponseCustomSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponseCustomSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponseCustomSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseCustomSharedEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileGetResponseCustomSharedEntriesCustomEntry],
// [DLPProfileGetResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileGetResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileGetResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntry],
// [DLPProfileGetResponseCustomSharedEntriesWordListEntry].
func (r DLPProfileGetResponseCustomSharedEntry) AsUnion() DLPProfileGetResponseCustomSharedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileGetResponseCustomSharedEntriesCustomEntry],
// [DLPProfileGetResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileGetResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileGetResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntry] or
// [DLPProfileGetResponseCustomSharedEntriesWordListEntry].
type DLPProfileGetResponseCustomSharedEntriesUnion interface {
	implementsDLPProfileGetResponseCustomSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseCustomSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseCustomSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileGetResponseCustomSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                    `json:"enabled,required"`
	Name        string                                                  `json:"name,required"`
	Pattern     Pattern                                                 `json:"pattern,required"`
	Type        DLPProfileGetResponseCustomSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                               `json:"updated_at,required" format:"date-time"`
	Description string                                                  `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesCustomEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseCustomSharedEntriesCustomEntry]
type dlpProfileGetResponseCustomSharedEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomSharedEntriesCustomEntry) implementsDLPProfileGetResponseCustomSharedEntry() {
}

type DLPProfileGetResponseCustomSharedEntriesCustomEntryType string

const (
	DLPProfileGetResponseCustomSharedEntriesCustomEntryTypeCustom DLPProfileGetResponseCustomSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileGetResponseCustomSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesPredefinedEntry struct {
	ID         string                                                            `json:"id,required" format:"uuid"`
	Confidence DLPProfileGetResponseCustomSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                              `json:"enabled,required"`
	Name       string                                                            `json:"name,required"`
	Type       DLPProfileGetResponseCustomSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileGetResponseCustomSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseCustomSharedEntriesPredefinedEntry]
type dlpProfileGetResponseCustomSharedEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomSharedEntriesPredefinedEntry) implementsDLPProfileGetResponseCustomSharedEntry() {
}

type DLPProfileGetResponseCustomSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                  `json:"available,required"`
	JSON      dlpProfileGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON contains
// the JSON metadata for the struct
// [DLPProfileGetResponseCustomSharedEntriesPredefinedEntryConfidence]
type dlpProfileGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseCustomSharedEntriesPredefinedEntryType string

const (
	DLPProfileGetResponseCustomSharedEntriesPredefinedEntryTypePredefined DLPProfileGetResponseCustomSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileGetResponseCustomSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                  `json:"description,nullable"`
	JSON        dlpProfileGetResponseCustomSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariant]
type dlpProfileGetResponseCustomSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesIntegrationEntry struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfileGetResponseCustomSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseCustomSharedEntriesIntegrationEntry]
type dlpProfileGetResponseCustomSharedEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomSharedEntriesIntegrationEntry) implementsDLPProfileGetResponseCustomSharedEntry() {
}

type DLPProfileGetResponseCustomSharedEntriesIntegrationEntryType string

const (
	DLPProfileGetResponseCustomSharedEntriesIntegrationEntryTypeIntegration DLPProfileGetResponseCustomSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileGetResponseCustomSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                       `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                       `json:"enabled,required"`
	Name          string                                                     `json:"name,required"`
	Secret        bool                                                       `json:"secret,required"`
	Type          DLPProfileGetResponseCustomSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                  `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileGetResponseCustomSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesExactDataEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseCustomSharedEntriesExactDataEntry]
type dlpProfileGetResponseCustomSharedEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomSharedEntriesExactDataEntry) implementsDLPProfileGetResponseCustomSharedEntry() {
}

type DLPProfileGetResponseCustomSharedEntriesExactDataEntryType string

const (
	DLPProfileGetResponseCustomSharedEntriesExactDataEntryTypeExactData DLPProfileGetResponseCustomSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileGetResponseCustomSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                 `json:"enabled,required"`
	Name      string                                                               `json:"name,required"`
	Type      DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                            `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntry]
type dlpProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntry) implementsDLPProfileGetResponseCustomSharedEntry() {
}

type DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesWordListEntry struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Type      DLPProfileGetResponseCustomSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                               `json:"word_list,required"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseCustomSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileGetResponseCustomSharedEntriesWordListEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseCustomSharedEntriesWordListEntry]
type dlpProfileGetResponseCustomSharedEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseCustomSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseCustomSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseCustomSharedEntriesWordListEntry) implementsDLPProfileGetResponseCustomSharedEntry() {
}

type DLPProfileGetResponseCustomSharedEntriesWordListEntryType string

const (
	DLPProfileGetResponseCustomSharedEntriesWordListEntryTypeWordList DLPProfileGetResponseCustomSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileGetResponseCustomSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseCustomSharedEntriesType string

const (
	DLPProfileGetResponseCustomSharedEntriesTypeCustom              DLPProfileGetResponseCustomSharedEntriesType = "custom"
	DLPProfileGetResponseCustomSharedEntriesTypePredefined          DLPProfileGetResponseCustomSharedEntriesType = "predefined"
	DLPProfileGetResponseCustomSharedEntriesTypeIntegration         DLPProfileGetResponseCustomSharedEntriesType = "integration"
	DLPProfileGetResponseCustomSharedEntriesTypeExactData           DLPProfileGetResponseCustomSharedEntriesType = "exact_data"
	DLPProfileGetResponseCustomSharedEntriesTypeDocumentFingerprint DLPProfileGetResponseCustomSharedEntriesType = "document_fingerprint"
	DLPProfileGetResponseCustomSharedEntriesTypeWordList            DLPProfileGetResponseCustomSharedEntriesType = "word_list"
)

func (r DLPProfileGetResponseCustomSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseCustomSharedEntriesTypeCustom, DLPProfileGetResponseCustomSharedEntriesTypePredefined, DLPProfileGetResponseCustomSharedEntriesTypeIntegration, DLPProfileGetResponseCustomSharedEntriesTypeExactData, DLPProfileGetResponseCustomSharedEntriesTypeDocumentFingerprint, DLPProfileGetResponseCustomSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefined struct {
	// The id of the predefined profile (uuid).
	ID                string `json:"id,required" format:"uuid"`
	AllowedMatchCount int64  `json:"allowed_match_count,required"`
	// Deprecated: deprecated
	Entries []DLPProfileGetResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                                             `json:"name,required"`
	Type                DLPProfileGetResponsePredefinedType                `json:"type,required"`
	AIContextEnabled    bool                                               `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileGetResponsePredefinedConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                                `json:"open_access"`
	JSON       dlpProfileGetResponsePredefinedJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedJSON contains the JSON metadata for the struct
// [DLPProfileGetResponsePredefined]
type dlpProfileGetResponsePredefinedJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AIContextEnabled    apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefined) implementsDLPProfileGetResponse() {}

type DLPProfileGetResponsePredefinedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                       `json:"enabled,required"`
	Name    string                                     `json:"name,required"`
	Type    DLPProfileGetResponsePredefinedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileGetResponsePredefinedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                              `json:"word_list"`
	JSON     dlpProfileGetResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileGetResponsePredefinedEntriesUnion
}

// dlpProfileGetResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPProfileGetResponsePredefinedEntry]
type dlpProfileGetResponsePredefinedEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponsePredefinedEntriesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileGetResponsePredefinedEntriesCustomEntry],
// [DLPProfileGetResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileGetResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileGetResponsePredefinedEntriesExactDataEntry],
// [DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntry],
// [DLPProfileGetResponsePredefinedEntriesWordListEntry].
func (r DLPProfileGetResponsePredefinedEntry) AsUnion() DLPProfileGetResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileGetResponsePredefinedEntriesCustomEntry],
// [DLPProfileGetResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileGetResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileGetResponsePredefinedEntriesExactDataEntry],
// [DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntry] or
// [DLPProfileGetResponsePredefinedEntriesWordListEntry].
type DLPProfileGetResponsePredefinedEntriesUnion interface {
	implementsDLPProfileGetResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponsePredefinedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponsePredefinedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileGetResponsePredefinedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                  `json:"enabled,required"`
	Name        string                                                `json:"name,required"`
	Pattern     Pattern                                               `json:"pattern,required"`
	Type        DLPProfileGetResponsePredefinedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                             `json:"updated_at,required" format:"date-time"`
	Description string                                                `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponsePredefinedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesCustomEntryJSON contains the JSON metadata
// for the struct [DLPProfileGetResponsePredefinedEntriesCustomEntry]
type dlpProfileGetResponsePredefinedEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesCustomEntry) implementsDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesCustomEntryType string

const (
	DLPProfileGetResponsePredefinedEntriesCustomEntryTypeCustom DLPProfileGetResponsePredefinedEntriesCustomEntryType = "custom"
)

func (r DLPProfileGetResponsePredefinedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesPredefinedEntry struct {
	ID         string                                                          `json:"id,required" format:"uuid"`
	Confidence DLPProfileGetResponsePredefinedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                            `json:"enabled,required"`
	Name       string                                                          `json:"name,required"`
	Type       DLPProfileGetResponsePredefinedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileGetResponsePredefinedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponsePredefinedEntriesPredefinedEntry]
type dlpProfileGetResponsePredefinedEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesPredefinedEntry) implementsDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                `json:"available,required"`
	JSON      dlpProfileGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponsePredefinedEntriesPredefinedEntryConfidence]
type dlpProfileGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponsePredefinedEntriesPredefinedEntryType string

const (
	DLPProfileGetResponsePredefinedEntriesPredefinedEntryTypePredefined DLPProfileGetResponsePredefinedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileGetResponsePredefinedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                `json:"description,nullable"`
	JSON        dlpProfileGetResponsePredefinedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariant]
type dlpProfileGetResponsePredefinedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantType string

const (
	DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesIntegrationEntry struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPProfileGetResponsePredefinedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponsePredefinedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponsePredefinedEntriesIntegrationEntry]
type dlpProfileGetResponsePredefinedEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesIntegrationEntry) implementsDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesIntegrationEntryType string

const (
	DLPProfileGetResponsePredefinedEntriesIntegrationEntryTypeIntegration DLPProfileGetResponsePredefinedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileGetResponsePredefinedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                     `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                     `json:"enabled,required"`
	Name          string                                                   `json:"name,required"`
	Secret        bool                                                     `json:"secret,required"`
	Type          DLPProfileGetResponsePredefinedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileGetResponsePredefinedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesExactDataEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponsePredefinedEntriesExactDataEntry]
type dlpProfileGetResponsePredefinedEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesExactDataEntry) implementsDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesExactDataEntryType string

const (
	DLPProfileGetResponsePredefinedEntriesExactDataEntryTypeExactData DLPProfileGetResponsePredefinedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileGetResponsePredefinedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntry struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponsePredefinedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesDocumentFingerprintEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntry]
type dlpProfileGetResponsePredefinedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntry) implementsDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesWordListEntry struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      DLPProfileGetResponsePredefinedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                             `json:"word_list,required"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponsePredefinedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileGetResponsePredefinedEntriesWordListEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponsePredefinedEntriesWordListEntry]
type dlpProfileGetResponsePredefinedEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponsePredefinedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponsePredefinedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponsePredefinedEntriesWordListEntry) implementsDLPProfileGetResponsePredefinedEntry() {
}

type DLPProfileGetResponsePredefinedEntriesWordListEntryType string

const (
	DLPProfileGetResponsePredefinedEntriesWordListEntryTypeWordList DLPProfileGetResponsePredefinedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileGetResponsePredefinedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedEntriesType string

const (
	DLPProfileGetResponsePredefinedEntriesTypeCustom              DLPProfileGetResponsePredefinedEntriesType = "custom"
	DLPProfileGetResponsePredefinedEntriesTypePredefined          DLPProfileGetResponsePredefinedEntriesType = "predefined"
	DLPProfileGetResponsePredefinedEntriesTypeIntegration         DLPProfileGetResponsePredefinedEntriesType = "integration"
	DLPProfileGetResponsePredefinedEntriesTypeExactData           DLPProfileGetResponsePredefinedEntriesType = "exact_data"
	DLPProfileGetResponsePredefinedEntriesTypeDocumentFingerprint DLPProfileGetResponsePredefinedEntriesType = "document_fingerprint"
	DLPProfileGetResponsePredefinedEntriesTypeWordList            DLPProfileGetResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileGetResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedEntriesTypeCustom, DLPProfileGetResponsePredefinedEntriesTypePredefined, DLPProfileGetResponsePredefinedEntriesTypeIntegration, DLPProfileGetResponsePredefinedEntriesTypeExactData, DLPProfileGetResponsePredefinedEntriesTypeDocumentFingerprint, DLPProfileGetResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedType string

const (
	DLPProfileGetResponsePredefinedTypePredefined DLPProfileGetResponsePredefinedType = "predefined"
)

func (r DLPProfileGetResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponsePredefinedConfidenceThreshold string

const (
	DLPProfileGetResponsePredefinedConfidenceThresholdLow      DLPProfileGetResponsePredefinedConfidenceThreshold = "low"
	DLPProfileGetResponsePredefinedConfidenceThresholdMedium   DLPProfileGetResponsePredefinedConfidenceThreshold = "medium"
	DLPProfileGetResponsePredefinedConfidenceThresholdHigh     DLPProfileGetResponsePredefinedConfidenceThreshold = "high"
	DLPProfileGetResponsePredefinedConfidenceThresholdVeryHigh DLPProfileGetResponsePredefinedConfidenceThreshold = "very_high"
)

func (r DLPProfileGetResponsePredefinedConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileGetResponsePredefinedConfidenceThresholdLow, DLPProfileGetResponsePredefinedConfidenceThresholdMedium, DLPProfileGetResponsePredefinedConfidenceThresholdHigh, DLPProfileGetResponsePredefinedConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegration struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Entries       []DLPProfileGetResponseIntegrationEntry       `json:"entries,required"`
	Name          string                                        `json:"name,required"`
	SharedEntries []DLPProfileGetResponseIntegrationSharedEntry `json:"shared_entries,required"`
	Type          DLPProfileGetResponseIntegrationType          `json:"type,required"`
	UpdatedAt     time.Time                                     `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                               `json:"description,nullable"`
	JSON        dlpProfileGetResponseIntegrationJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseIntegration]
type dlpProfileGetResponseIntegrationJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Entries       apijson.Field
	Name          apijson.Field
	SharedEntries apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	Description   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegration) implementsDLPProfileGetResponse() {}

type DLPProfileGetResponseIntegrationEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                        `json:"enabled,required"`
	Name    string                                      `json:"name,required"`
	Type    DLPProfileGetResponseIntegrationEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseIntegrationEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                               `json:"word_list"`
	JSON     dlpProfileGetResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileGetResponseIntegrationEntriesUnion
}

// dlpProfileGetResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseIntegrationEntry]
type dlpProfileGetResponseIntegrationEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseIntegrationEntriesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileGetResponseIntegrationEntriesCustomEntry],
// [DLPProfileGetResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileGetResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileGetResponseIntegrationEntriesExactDataEntry],
// [DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntry],
// [DLPProfileGetResponseIntegrationEntriesWordListEntry].
func (r DLPProfileGetResponseIntegrationEntry) AsUnion() DLPProfileGetResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileGetResponseIntegrationEntriesCustomEntry],
// [DLPProfileGetResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileGetResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileGetResponseIntegrationEntriesExactDataEntry],
// [DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntry] or
// [DLPProfileGetResponseIntegrationEntriesWordListEntry].
type DLPProfileGetResponseIntegrationEntriesUnion interface {
	implementsDLPProfileGetResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseIntegrationEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationEntriesWordListEntry{}),
		},
	)
}

type DLPProfileGetResponseIntegrationEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                   `json:"enabled,required"`
	Name        string                                                 `json:"name,required"`
	Pattern     Pattern                                                `json:"pattern,required"`
	Type        DLPProfileGetResponseIntegrationEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                              `json:"updated_at,required" format:"date-time"`
	Description string                                                 `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesCustomEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseIntegrationEntriesCustomEntry]
type dlpProfileGetResponseIntegrationEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesCustomEntry) implementsDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesCustomEntryType string

const (
	DLPProfileGetResponseIntegrationEntriesCustomEntryTypeCustom DLPProfileGetResponseIntegrationEntriesCustomEntryType = "custom"
)

func (r DLPProfileGetResponseIntegrationEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesPredefinedEntry struct {
	ID         string                                                           `json:"id,required" format:"uuid"`
	Confidence DLPProfileGetResponseIntegrationEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                             `json:"enabled,required"`
	Name       string                                                           `json:"name,required"`
	Type       DLPProfileGetResponseIntegrationEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileGetResponseIntegrationEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseIntegrationEntriesPredefinedEntry]
type dlpProfileGetResponseIntegrationEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesPredefinedEntry) implementsDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                 `json:"available,required"`
	JSON      dlpProfileGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON contains
// the JSON metadata for the struct
// [DLPProfileGetResponseIntegrationEntriesPredefinedEntryConfidence]
type dlpProfileGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseIntegrationEntriesPredefinedEntryType string

const (
	DLPProfileGetResponseIntegrationEntriesPredefinedEntryTypePredefined DLPProfileGetResponseIntegrationEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileGetResponseIntegrationEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                 `json:"description,nullable"`
	JSON        dlpProfileGetResponseIntegrationEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariant]
type dlpProfileGetResponseIntegrationEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantType string

const (
	DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesIntegrationEntry struct {
	ID        string                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                        `json:"enabled,required"`
	Name      string                                                      `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                   `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseIntegrationEntriesIntegrationEntry]
type dlpProfileGetResponseIntegrationEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesIntegrationEntry) implementsDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesIntegrationEntryType string

const (
	DLPProfileGetResponseIntegrationEntriesIntegrationEntryTypeIntegration DLPProfileGetResponseIntegrationEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileGetResponseIntegrationEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                      `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                      `json:"enabled,required"`
	Name          string                                                    `json:"name,required"`
	Secret        bool                                                      `json:"secret,required"`
	Type          DLPProfileGetResponseIntegrationEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                 `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileGetResponseIntegrationEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesExactDataEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseIntegrationEntriesExactDataEntry]
type dlpProfileGetResponseIntegrationEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesExactDataEntry) implementsDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesExactDataEntryType string

const (
	DLPProfileGetResponseIntegrationEntriesExactDataEntryTypeExactData DLPProfileGetResponseIntegrationEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileGetResponseIntegrationEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntry struct {
	ID        string                                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                `json:"enabled,required"`
	Name      string                                                              `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                           `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponseIntegrationEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesDocumentFingerprintEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntry]
type dlpProfileGetResponseIntegrationEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntry) implementsDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntryType string

const (
	DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesWordListEntry struct {
	ID        string                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                     `json:"enabled,required"`
	Name      string                                                   `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                              `json:"word_list,required"`
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationEntriesWordListEntryJSON contains the JSON
// metadata for the struct [DLPProfileGetResponseIntegrationEntriesWordListEntry]
type dlpProfileGetResponseIntegrationEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationEntriesWordListEntry) implementsDLPProfileGetResponseIntegrationEntry() {
}

type DLPProfileGetResponseIntegrationEntriesWordListEntryType string

const (
	DLPProfileGetResponseIntegrationEntriesWordListEntryTypeWordList DLPProfileGetResponseIntegrationEntriesWordListEntryType = "word_list"
)

func (r DLPProfileGetResponseIntegrationEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationEntriesType string

const (
	DLPProfileGetResponseIntegrationEntriesTypeCustom              DLPProfileGetResponseIntegrationEntriesType = "custom"
	DLPProfileGetResponseIntegrationEntriesTypePredefined          DLPProfileGetResponseIntegrationEntriesType = "predefined"
	DLPProfileGetResponseIntegrationEntriesTypeIntegration         DLPProfileGetResponseIntegrationEntriesType = "integration"
	DLPProfileGetResponseIntegrationEntriesTypeExactData           DLPProfileGetResponseIntegrationEntriesType = "exact_data"
	DLPProfileGetResponseIntegrationEntriesTypeDocumentFingerprint DLPProfileGetResponseIntegrationEntriesType = "document_fingerprint"
	DLPProfileGetResponseIntegrationEntriesTypeWordList            DLPProfileGetResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileGetResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationEntriesTypeCustom, DLPProfileGetResponseIntegrationEntriesTypePredefined, DLPProfileGetResponseIntegrationEntriesTypeIntegration, DLPProfileGetResponseIntegrationEntriesTypeExactData, DLPProfileGetResponseIntegrationEntriesTypeDocumentFingerprint, DLPProfileGetResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                              `json:"enabled,required"`
	Name    string                                            `json:"name,required"`
	Type    DLPProfileGetResponseIntegrationSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                     `json:"word_list"`
	JSON     dlpProfileGetResponseIntegrationSharedEntryJSON `json:"-"`
	union    DLPProfileGetResponseIntegrationSharedEntriesUnion
}

// dlpProfileGetResponseIntegrationSharedEntryJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseIntegrationSharedEntry]
type dlpProfileGetResponseIntegrationSharedEntryJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Description   apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpProfileGetResponseIntegrationSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileGetResponseIntegrationSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileGetResponseIntegrationSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileGetResponseIntegrationSharedEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileGetResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesWordListEntry].
func (r DLPProfileGetResponseIntegrationSharedEntry) AsUnion() DLPProfileGetResponseIntegrationSharedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileGetResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntry] or
// [DLPProfileGetResponseIntegrationSharedEntriesWordListEntry].
type DLPProfileGetResponseIntegrationSharedEntriesUnion interface {
	implementsDLPProfileGetResponseIntegrationSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileGetResponseIntegrationSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileGetResponseIntegrationSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileGetResponseIntegrationSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                         `json:"enabled,required"`
	Name        string                                                       `json:"name,required"`
	Pattern     Pattern                                                      `json:"pattern,required"`
	Type        DLPProfileGetResponseIntegrationSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                    `json:"updated_at,required" format:"date-time"`
	Description string                                                       `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesCustomEntry]
type dlpProfileGetResponseIntegrationSharedEntriesCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationSharedEntriesCustomEntry) implementsDLPProfileGetResponseIntegrationSharedEntry() {
}

type DLPProfileGetResponseIntegrationSharedEntriesCustomEntryType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesCustomEntryTypeCustom DLPProfileGetResponseIntegrationSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntry struct {
	ID         string                                                                 `json:"id,required" format:"uuid"`
	Confidence DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                   `json:"enabled,required"`
	Name       string                                                                 `json:"name,required"`
	Type       DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntry]
type dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntry) implementsDLPProfileGetResponseIntegrationSharedEntry() {
}

type DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                       `json:"available,required"`
	JSON      dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidence]
type dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryTypePredefined DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                       `json:"description,nullable"`
	JSON        dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON contains
// the JSON metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariant]
type dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntry struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                         `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntry]
type dlpProfileGetResponseIntegrationSharedEntriesIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntry) implementsDLPProfileGetResponseIntegrationSharedEntry() {
}

type DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntryType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                            `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                            `json:"enabled,required"`
	Name          string                                                          `json:"name,required"`
	Secret        bool                                                            `json:"secret,required"`
	Type          DLPProfileGetResponseIntegrationSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                       `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileGetResponseIntegrationSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesExactDataEntry]
type dlpProfileGetResponseIntegrationSharedEntriesExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationSharedEntriesExactDataEntry) implementsDLPProfileGetResponseIntegrationSharedEntry() {
}

type DLPProfileGetResponseIntegrationSharedEntriesExactDataEntryType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesExactDataEntryTypeExactData DLPProfileGetResponseIntegrationSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                      `json:"enabled,required"`
	Name      string                                                                    `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                 `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntry]
type dlpProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntry) implementsDLPProfileGetResponseIntegrationSharedEntry() {
}

type DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesWordListEntry struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfileGetResponseIntegrationSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                    `json:"word_list,required"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileGetResponseIntegrationSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileGetResponseIntegrationSharedEntriesWordListEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileGetResponseIntegrationSharedEntriesWordListEntry]
type dlpProfileGetResponseIntegrationSharedEntriesWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseIntegrationSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseIntegrationSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileGetResponseIntegrationSharedEntriesWordListEntry) implementsDLPProfileGetResponseIntegrationSharedEntry() {
}

type DLPProfileGetResponseIntegrationSharedEntriesWordListEntryType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesWordListEntryTypeWordList DLPProfileGetResponseIntegrationSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationSharedEntriesType string

const (
	DLPProfileGetResponseIntegrationSharedEntriesTypeCustom              DLPProfileGetResponseIntegrationSharedEntriesType = "custom"
	DLPProfileGetResponseIntegrationSharedEntriesTypePredefined          DLPProfileGetResponseIntegrationSharedEntriesType = "predefined"
	DLPProfileGetResponseIntegrationSharedEntriesTypeIntegration         DLPProfileGetResponseIntegrationSharedEntriesType = "integration"
	DLPProfileGetResponseIntegrationSharedEntriesTypeExactData           DLPProfileGetResponseIntegrationSharedEntriesType = "exact_data"
	DLPProfileGetResponseIntegrationSharedEntriesTypeDocumentFingerprint DLPProfileGetResponseIntegrationSharedEntriesType = "document_fingerprint"
	DLPProfileGetResponseIntegrationSharedEntriesTypeWordList            DLPProfileGetResponseIntegrationSharedEntriesType = "word_list"
)

func (r DLPProfileGetResponseIntegrationSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationSharedEntriesTypeCustom, DLPProfileGetResponseIntegrationSharedEntriesTypePredefined, DLPProfileGetResponseIntegrationSharedEntriesTypeIntegration, DLPProfileGetResponseIntegrationSharedEntriesTypeExactData, DLPProfileGetResponseIntegrationSharedEntriesTypeDocumentFingerprint, DLPProfileGetResponseIntegrationSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileGetResponseIntegrationType string

const (
	DLPProfileGetResponseIntegrationTypeIntegration DLPProfileGetResponseIntegrationType = "integration"
)

func (r DLPProfileGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseType string

const (
	DLPProfileGetResponseTypeCustom      DLPProfileGetResponseType = "custom"
	DLPProfileGetResponseTypePredefined  DLPProfileGetResponseType = "predefined"
	DLPProfileGetResponseTypeIntegration DLPProfileGetResponseType = "integration"
)

func (r DLPProfileGetResponseType) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseTypeCustom, DLPProfileGetResponseTypePredefined, DLPProfileGetResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileGetResponseConfidenceThreshold string

const (
	DLPProfileGetResponseConfidenceThresholdLow      DLPProfileGetResponseConfidenceThreshold = "low"
	DLPProfileGetResponseConfidenceThresholdMedium   DLPProfileGetResponseConfidenceThreshold = "medium"
	DLPProfileGetResponseConfidenceThresholdHigh     DLPProfileGetResponseConfidenceThreshold = "high"
	DLPProfileGetResponseConfidenceThresholdVeryHigh DLPProfileGetResponseConfidenceThreshold = "very_high"
)

func (r DLPProfileGetResponseConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseConfidenceThresholdLow, DLPProfileGetResponseConfidenceThresholdMedium, DLPProfileGetResponseConfidenceThresholdHigh, DLPProfileGetResponseConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Return all profiles, including those that current account does not have access
	// to.
	All param.Field[bool] `query:"all"`
}

// URLQuery serializes [DLPProfileListParams]'s query parameters as `url.Values`.
func (r DLPProfileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DLPProfileGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileGetResponseEnvelope struct {
	Errors   []DLPProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileGetResponse                `json:"result"`
	JSON    dlpProfileGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPProfileGetResponseEnvelope]
type dlpProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPProfileGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeErrors]
type dlpProfileGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpProfileGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeErrorsSource]
type dlpProfileGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPProfileGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPProfileGetResponseEnvelopeMessages]
type dlpProfileGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpProfileGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPProfileGetResponseEnvelopeMessagesSource]
type dlpProfileGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileGetResponseEnvelopeSuccess bool

const (
	DLPProfileGetResponseEnvelopeSuccessTrue DLPProfileGetResponseEnvelopeSuccess = true
)

func (r DLPProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
