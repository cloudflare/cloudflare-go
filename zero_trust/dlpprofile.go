// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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
	opts = append(r.Options[:], opts...)
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

// Fetches a DLP profile by ID
func (r *DLPProfileService) Get(ctx context.Context, profileID string, query DLPProfileGetParams, opts ...option.RequestOption) (res *Profile, err error) {
	var env DLPProfileGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile
	Name string      `json:"name,required"`
	Type ProfileType `json:"type,required"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64                      `json:"allowed_match_count"`
	ConfidenceThreshold ProfileConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile
	Description string `json:"description,nullable"`
	// This field can have the runtime type of [[]ProfileCustomEntry],
	// [[]ProfilePredefinedEntry], [[]ProfileIntegrationEntry].
	Entries    interface{} `json:"entries"`
	OCREnabled bool        `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone
	OpenAccess bool `json:"open_access"`
	// When the profile was lasted updated
	UpdatedAt time.Time   `json:"updated_at" format:"date-time"`
	JSON      profileJSON `json:"-"`
	union     ProfileUnion
}

// profileJSON contains the JSON metadata for the struct [Profile]
type profileJSON struct {
	ID                  apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	AllowedMatchCount   apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Description         apijson.Field
	Entries             apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
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
// Possible runtime types of the union are [zero_trust.ProfileCustom],
// [zero_trust.ProfilePredefined], [zero_trust.ProfileIntegration].
func (r Profile) AsUnion() ProfileUnion {
	return r.union
}

// Union satisfied by [zero_trust.ProfileCustom], [zero_trust.ProfilePredefined] or
// [zero_trust.ProfileIntegration].
type ProfileUnion interface {
	implementsZeroTrustProfile()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfilePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type ProfileCustom struct {
	// The id of the profile (uuid)
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness,required"`
	// When the profile was created
	CreatedAt time.Time            `json:"created_at,required" format:"date-time"`
	Entries   []ProfileCustomEntry `json:"entries,required"`
	// The name of the profile
	Name       string            `json:"name,required"`
	OCREnabled bool              `json:"ocr_enabled,required"`
	Type       ProfileCustomType `json:"type,required"`
	// When the profile was lasted updated
	UpdatedAt           time.Time                        `json:"updated_at,required" format:"date-time"`
	ConfidenceThreshold ProfileCustomConfidenceThreshold `json:"confidence_threshold"`
	// The description of the profile
	Description string            `json:"description,nullable"`
	JSON        profileCustomJSON `json:"-"`
}

// profileCustomJSON contains the JSON metadata for the struct [ProfileCustom]
type profileCustomJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	ContextAwareness    apijson.Field
	CreatedAt           apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	OCREnabled          apijson.Field
	Type                apijson.Field
	UpdatedAt           apijson.Field
	ConfidenceThreshold apijson.Field
	Description         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfileCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustom) implementsZeroTrustProfile() {}

type ProfileCustomEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of
	// [ProfileCustomEntriesPredefinedConfidence].
	Confidence interface{}              `json:"confidence,required"`
	Enabled    bool                     `json:"enabled,required"`
	Name       string                   `json:"name,required"`
	Type       ProfileCustomEntriesType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	WordList  interface{}            `json:"word_list,required"`
	CreatedAt time.Time              `json:"created_at" format:"date-time"`
	Pattern   Pattern                `json:"pattern"`
	ProfileID string                 `json:"profile_id,nullable" format:"uuid"`
	Secret    bool                   `json:"secret"`
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	JSON      profileCustomEntryJSON `json:"-"`
	union     ProfileCustomEntriesUnion
}

// profileCustomEntryJSON contains the JSON metadata for the struct
// [ProfileCustomEntry]
type profileCustomEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	WordList    apijson.Field
	CreatedAt   apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	Secret      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r profileCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileCustomEntriesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.ProfileCustomEntriesCustom],
// [zero_trust.ProfileCustomEntriesPredefined],
// [zero_trust.ProfileCustomEntriesIntegration],
// [zero_trust.ProfileCustomEntriesExactData],
// [zero_trust.ProfileCustomEntriesWordList].
func (r ProfileCustomEntry) AsUnion() ProfileCustomEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.ProfileCustomEntriesCustom],
// [zero_trust.ProfileCustomEntriesPredefined],
// [zero_trust.ProfileCustomEntriesIntegration],
// [zero_trust.ProfileCustomEntriesExactData] or
// [zero_trust.ProfileCustomEntriesWordList].
type ProfileCustomEntriesUnion interface {
	implementsZeroTrustProfileCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileCustomEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileCustomEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileCustomEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileCustomEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileCustomEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileCustomEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type ProfileCustomEntriesCustom struct {
	ID        string                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                           `json:"enabled,required"`
	Name      string                         `json:"name,required"`
	Pattern   Pattern                        `json:"pattern,required"`
	Type      ProfileCustomEntriesCustomType `json:"type,required"`
	UpdatedAt time.Time                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                         `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomEntriesCustomJSON `json:"-"`
}

// profileCustomEntriesCustomJSON contains the JSON metadata for the struct
// [ProfileCustomEntriesCustom]
type profileCustomEntriesCustomJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomEntriesCustom) implementsZeroTrustProfileCustomEntry() {}

type ProfileCustomEntriesCustomType string

const (
	ProfileCustomEntriesCustomTypeCustom ProfileCustomEntriesCustomType = "custom"
)

func (r ProfileCustomEntriesCustomType) IsKnown() bool {
	switch r {
	case ProfileCustomEntriesCustomTypeCustom:
		return true
	}
	return false
}

type ProfileCustomEntriesPredefined struct {
	ID         string                                   `json:"id,required" format:"uuid"`
	Confidence ProfileCustomEntriesPredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                     `json:"enabled,required"`
	Name       string                                   `json:"name,required"`
	Type       ProfileCustomEntriesPredefinedType       `json:"type,required"`
	ProfileID  string                                   `json:"profile_id,nullable" format:"uuid"`
	JSON       profileCustomEntriesPredefinedJSON       `json:"-"`
}

// profileCustomEntriesPredefinedJSON contains the JSON metadata for the struct
// [ProfileCustomEntriesPredefined]
type profileCustomEntriesPredefinedJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomEntriesPredefined) implementsZeroTrustProfileCustomEntry() {}

type ProfileCustomEntriesPredefinedConfidence struct {
	// Indicates whether this entry can be made more or less sensitive by setting a
	// confidence threshold. Profiles that use an entry with `available` set to true
	// can use confidence thresholds
	Available bool                                         `json:"available,required"`
	JSON      profileCustomEntriesPredefinedConfidenceJSON `json:"-"`
}

// profileCustomEntriesPredefinedConfidenceJSON contains the JSON metadata for the
// struct [ProfileCustomEntriesPredefinedConfidence]
type profileCustomEntriesPredefinedConfidenceJSON struct {
	Available   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomEntriesPredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomEntriesPredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileCustomEntriesPredefinedType string

const (
	ProfileCustomEntriesPredefinedTypePredefined ProfileCustomEntriesPredefinedType = "predefined"
)

func (r ProfileCustomEntriesPredefinedType) IsKnown() bool {
	switch r {
	case ProfileCustomEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type ProfileCustomEntriesIntegration struct {
	ID        string                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                `json:"enabled,required"`
	Name      string                              `json:"name,required"`
	Type      ProfileCustomEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                              `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomEntriesIntegrationJSON `json:"-"`
}

// profileCustomEntriesIntegrationJSON contains the JSON metadata for the struct
// [ProfileCustomEntriesIntegration]
type profileCustomEntriesIntegrationJSON struct {
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

func (r *ProfileCustomEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomEntriesIntegration) implementsZeroTrustProfileCustomEntry() {}

type ProfileCustomEntriesIntegrationType string

const (
	ProfileCustomEntriesIntegrationTypeIntegration ProfileCustomEntriesIntegrationType = "integration"
)

func (r ProfileCustomEntriesIntegrationType) IsKnown() bool {
	switch r {
	case ProfileCustomEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type ProfileCustomEntriesExactData struct {
	ID        string                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                              `json:"enabled,required"`
	Name      string                            `json:"name,required"`
	Secret    bool                              `json:"secret,required"`
	Type      ProfileCustomEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                         `json:"updated_at,required" format:"date-time"`
	JSON      profileCustomEntriesExactDataJSON `json:"-"`
}

// profileCustomEntriesExactDataJSON contains the JSON metadata for the struct
// [ProfileCustomEntriesExactData]
type profileCustomEntriesExactDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileCustomEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomEntriesExactData) implementsZeroTrustProfileCustomEntry() {}

type ProfileCustomEntriesExactDataType string

const (
	ProfileCustomEntriesExactDataTypeExactData ProfileCustomEntriesExactDataType = "exact_data"
)

func (r ProfileCustomEntriesExactDataType) IsKnown() bool {
	switch r {
	case ProfileCustomEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type ProfileCustomEntriesWordList struct {
	ID        string                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                             `json:"enabled,required"`
	Name      string                           `json:"name,required"`
	Type      ProfileCustomEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                      `json:"word_list,required"`
	ProfileID string                           `json:"profile_id,nullable" format:"uuid"`
	JSON      profileCustomEntriesWordListJSON `json:"-"`
}

// profileCustomEntriesWordListJSON contains the JSON metadata for the struct
// [ProfileCustomEntriesWordList]
type profileCustomEntriesWordListJSON struct {
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

func (r *ProfileCustomEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileCustomEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r ProfileCustomEntriesWordList) implementsZeroTrustProfileCustomEntry() {}

type ProfileCustomEntriesWordListType string

const (
	ProfileCustomEntriesWordListTypeWordList ProfileCustomEntriesWordListType = "word_list"
)

func (r ProfileCustomEntriesWordListType) IsKnown() bool {
	switch r {
	case ProfileCustomEntriesWordListTypeWordList:
		return true
	}
	return false
}

type ProfileCustomEntriesType string

const (
	ProfileCustomEntriesTypeCustom      ProfileCustomEntriesType = "custom"
	ProfileCustomEntriesTypePredefined  ProfileCustomEntriesType = "predefined"
	ProfileCustomEntriesTypeIntegration ProfileCustomEntriesType = "integration"
	ProfileCustomEntriesTypeExactData   ProfileCustomEntriesType = "exact_data"
	ProfileCustomEntriesTypeWordList    ProfileCustomEntriesType = "word_list"
)

func (r ProfileCustomEntriesType) IsKnown() bool {
	switch r {
	case ProfileCustomEntriesTypeCustom, ProfileCustomEntriesTypePredefined, ProfileCustomEntriesTypeIntegration, ProfileCustomEntriesTypeExactData, ProfileCustomEntriesTypeWordList:
		return true
	}
	return false
}

type ProfileCustomType string

const (
	ProfileCustomTypeCustom ProfileCustomType = "custom"
)

func (r ProfileCustomType) IsKnown() bool {
	switch r {
	case ProfileCustomTypeCustom:
		return true
	}
	return false
}

type ProfileCustomConfidenceThreshold string

const (
	ProfileCustomConfidenceThresholdLow      ProfileCustomConfidenceThreshold = "low"
	ProfileCustomConfidenceThresholdMedium   ProfileCustomConfidenceThreshold = "medium"
	ProfileCustomConfidenceThresholdHigh     ProfileCustomConfidenceThreshold = "high"
	ProfileCustomConfidenceThresholdVeryHigh ProfileCustomConfidenceThreshold = "very_high"
)

func (r ProfileCustomConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfileCustomConfidenceThresholdLow, ProfileCustomConfidenceThresholdMedium, ProfileCustomConfidenceThresholdHigh, ProfileCustomConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type ProfilePredefined struct {
	// The id of the predefined profile (uuid)
	ID                string                   `json:"id,required" format:"uuid"`
	AllowedMatchCount int64                    `json:"allowed_match_count,required"`
	Entries           []ProfilePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile
	Name                string                               `json:"name,required"`
	Type                ProfilePredefinedType                `json:"type,required"`
	ConfidenceThreshold ProfilePredefinedConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone
	OpenAccess bool                  `json:"open_access"`
	JSON       profilePredefinedJSON `json:"-"`
}

// profilePredefinedJSON contains the JSON metadata for the struct
// [ProfilePredefined]
type profilePredefinedJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	Type                apijson.Field
	ConfidenceThreshold apijson.Field
	ContextAwareness    apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ProfilePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefined) implementsZeroTrustProfile() {}

type ProfilePredefinedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of
	// [ProfilePredefinedEntriesPredefinedConfidence].
	Confidence interface{}                  `json:"confidence,required"`
	Enabled    bool                         `json:"enabled,required"`
	Name       string                       `json:"name,required"`
	Type       ProfilePredefinedEntriesType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	WordList  interface{}                `json:"word_list,required"`
	CreatedAt time.Time                  `json:"created_at" format:"date-time"`
	Pattern   Pattern                    `json:"pattern"`
	ProfileID string                     `json:"profile_id,nullable" format:"uuid"`
	Secret    bool                       `json:"secret"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      profilePredefinedEntryJSON `json:"-"`
	union     ProfilePredefinedEntriesUnion
}

// profilePredefinedEntryJSON contains the JSON metadata for the struct
// [ProfilePredefinedEntry]
type profilePredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	WordList    apijson.Field
	CreatedAt   apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	Secret      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r profilePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfilePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfilePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfilePredefinedEntriesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.ProfilePredefinedEntriesCustom],
// [zero_trust.ProfilePredefinedEntriesPredefined],
// [zero_trust.ProfilePredefinedEntriesIntegration],
// [zero_trust.ProfilePredefinedEntriesExactData],
// [zero_trust.ProfilePredefinedEntriesWordList].
func (r ProfilePredefinedEntry) AsUnion() ProfilePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.ProfilePredefinedEntriesCustom],
// [zero_trust.ProfilePredefinedEntriesPredefined],
// [zero_trust.ProfilePredefinedEntriesIntegration],
// [zero_trust.ProfilePredefinedEntriesExactData] or
// [zero_trust.ProfilePredefinedEntriesWordList].
type ProfilePredefinedEntriesUnion interface {
	implementsZeroTrustProfilePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfilePredefinedEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfilePredefinedEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfilePredefinedEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfilePredefinedEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfilePredefinedEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfilePredefinedEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type ProfilePredefinedEntriesCustom struct {
	ID        string                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                               `json:"enabled,required"`
	Name      string                             `json:"name,required"`
	Pattern   Pattern                            `json:"pattern,required"`
	Type      ProfilePredefinedEntriesCustomType `json:"type,required"`
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                             `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedEntriesCustomJSON `json:"-"`
}

// profilePredefinedEntriesCustomJSON contains the JSON metadata for the struct
// [ProfilePredefinedEntriesCustom]
type profilePredefinedEntriesCustomJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedEntriesCustom) implementsZeroTrustProfilePredefinedEntry() {}

type ProfilePredefinedEntriesCustomType string

const (
	ProfilePredefinedEntriesCustomTypeCustom ProfilePredefinedEntriesCustomType = "custom"
)

func (r ProfilePredefinedEntriesCustomType) IsKnown() bool {
	switch r {
	case ProfilePredefinedEntriesCustomTypeCustom:
		return true
	}
	return false
}

type ProfilePredefinedEntriesPredefined struct {
	ID         string                                       `json:"id,required" format:"uuid"`
	Confidence ProfilePredefinedEntriesPredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                         `json:"enabled,required"`
	Name       string                                       `json:"name,required"`
	Type       ProfilePredefinedEntriesPredefinedType       `json:"type,required"`
	ProfileID  string                                       `json:"profile_id,nullable" format:"uuid"`
	JSON       profilePredefinedEntriesPredefinedJSON       `json:"-"`
}

// profilePredefinedEntriesPredefinedJSON contains the JSON metadata for the struct
// [ProfilePredefinedEntriesPredefined]
type profilePredefinedEntriesPredefinedJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedEntriesPredefined) implementsZeroTrustProfilePredefinedEntry() {}

type ProfilePredefinedEntriesPredefinedConfidence struct {
	// Indicates whether this entry can be made more or less sensitive by setting a
	// confidence threshold. Profiles that use an entry with `available` set to true
	// can use confidence thresholds
	Available bool                                             `json:"available,required"`
	JSON      profilePredefinedEntriesPredefinedConfidenceJSON `json:"-"`
}

// profilePredefinedEntriesPredefinedConfidenceJSON contains the JSON metadata for
// the struct [ProfilePredefinedEntriesPredefinedConfidence]
type profilePredefinedEntriesPredefinedConfidenceJSON struct {
	Available   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedEntriesPredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedEntriesPredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfilePredefinedEntriesPredefinedType string

const (
	ProfilePredefinedEntriesPredefinedTypePredefined ProfilePredefinedEntriesPredefinedType = "predefined"
)

func (r ProfilePredefinedEntriesPredefinedType) IsKnown() bool {
	switch r {
	case ProfilePredefinedEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type ProfilePredefinedEntriesIntegration struct {
	ID        string                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                    `json:"enabled,required"`
	Name      string                                  `json:"name,required"`
	Type      ProfilePredefinedEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedEntriesIntegrationJSON `json:"-"`
}

// profilePredefinedEntriesIntegrationJSON contains the JSON metadata for the
// struct [ProfilePredefinedEntriesIntegration]
type profilePredefinedEntriesIntegrationJSON struct {
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

func (r *ProfilePredefinedEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedEntriesIntegration) implementsZeroTrustProfilePredefinedEntry() {}

type ProfilePredefinedEntriesIntegrationType string

const (
	ProfilePredefinedEntriesIntegrationTypeIntegration ProfilePredefinedEntriesIntegrationType = "integration"
)

func (r ProfilePredefinedEntriesIntegrationType) IsKnown() bool {
	switch r {
	case ProfilePredefinedEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type ProfilePredefinedEntriesExactData struct {
	ID        string                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                  `json:"enabled,required"`
	Name      string                                `json:"name,required"`
	Secret    bool                                  `json:"secret,required"`
	Type      ProfilePredefinedEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	JSON      profilePredefinedEntriesExactDataJSON `json:"-"`
}

// profilePredefinedEntriesExactDataJSON contains the JSON metadata for the struct
// [ProfilePredefinedEntriesExactData]
type profilePredefinedEntriesExactDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilePredefinedEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedEntriesExactData) implementsZeroTrustProfilePredefinedEntry() {}

type ProfilePredefinedEntriesExactDataType string

const (
	ProfilePredefinedEntriesExactDataTypeExactData ProfilePredefinedEntriesExactDataType = "exact_data"
)

func (r ProfilePredefinedEntriesExactDataType) IsKnown() bool {
	switch r {
	case ProfilePredefinedEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type ProfilePredefinedEntriesWordList struct {
	ID        string                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                 `json:"enabled,required"`
	Name      string                               `json:"name,required"`
	Type      ProfilePredefinedEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                          `json:"word_list,required"`
	ProfileID string                               `json:"profile_id,nullable" format:"uuid"`
	JSON      profilePredefinedEntriesWordListJSON `json:"-"`
}

// profilePredefinedEntriesWordListJSON contains the JSON metadata for the struct
// [ProfilePredefinedEntriesWordList]
type profilePredefinedEntriesWordListJSON struct {
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

func (r *ProfilePredefinedEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profilePredefinedEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r ProfilePredefinedEntriesWordList) implementsZeroTrustProfilePredefinedEntry() {}

type ProfilePredefinedEntriesWordListType string

const (
	ProfilePredefinedEntriesWordListTypeWordList ProfilePredefinedEntriesWordListType = "word_list"
)

func (r ProfilePredefinedEntriesWordListType) IsKnown() bool {
	switch r {
	case ProfilePredefinedEntriesWordListTypeWordList:
		return true
	}
	return false
}

type ProfilePredefinedEntriesType string

const (
	ProfilePredefinedEntriesTypeCustom      ProfilePredefinedEntriesType = "custom"
	ProfilePredefinedEntriesTypePredefined  ProfilePredefinedEntriesType = "predefined"
	ProfilePredefinedEntriesTypeIntegration ProfilePredefinedEntriesType = "integration"
	ProfilePredefinedEntriesTypeExactData   ProfilePredefinedEntriesType = "exact_data"
	ProfilePredefinedEntriesTypeWordList    ProfilePredefinedEntriesType = "word_list"
)

func (r ProfilePredefinedEntriesType) IsKnown() bool {
	switch r {
	case ProfilePredefinedEntriesTypeCustom, ProfilePredefinedEntriesTypePredefined, ProfilePredefinedEntriesTypeIntegration, ProfilePredefinedEntriesTypeExactData, ProfilePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type ProfilePredefinedType string

const (
	ProfilePredefinedTypePredefined ProfilePredefinedType = "predefined"
)

func (r ProfilePredefinedType) IsKnown() bool {
	switch r {
	case ProfilePredefinedTypePredefined:
		return true
	}
	return false
}

type ProfilePredefinedConfidenceThreshold string

const (
	ProfilePredefinedConfidenceThresholdLow      ProfilePredefinedConfidenceThreshold = "low"
	ProfilePredefinedConfidenceThresholdMedium   ProfilePredefinedConfidenceThreshold = "medium"
	ProfilePredefinedConfidenceThresholdHigh     ProfilePredefinedConfidenceThreshold = "high"
	ProfilePredefinedConfidenceThresholdVeryHigh ProfilePredefinedConfidenceThreshold = "very_high"
)

func (r ProfilePredefinedConfidenceThreshold) IsKnown() bool {
	switch r {
	case ProfilePredefinedConfidenceThresholdLow, ProfilePredefinedConfidenceThresholdMedium, ProfilePredefinedConfidenceThresholdHigh, ProfilePredefinedConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type ProfileIntegration struct {
	ID        string                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                 `json:"created_at,required" format:"date-time"`
	Entries   []ProfileIntegrationEntry `json:"entries,required"`
	Name      string                    `json:"name,required"`
	Type      ProfileIntegrationType    `json:"type,required"`
	UpdatedAt time.Time                 `json:"updated_at,required" format:"date-time"`
	// The description of the profile
	Description string                 `json:"description,nullable"`
	JSON        profileIntegrationJSON `json:"-"`
}

// profileIntegrationJSON contains the JSON metadata for the struct
// [ProfileIntegration]
type profileIntegrationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegration) implementsZeroTrustProfile() {}

type ProfileIntegrationEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of
	// [ProfileIntegrationEntriesPredefinedConfidence].
	Confidence interface{}                   `json:"confidence,required"`
	Enabled    bool                          `json:"enabled,required"`
	Name       string                        `json:"name,required"`
	Type       ProfileIntegrationEntriesType `json:"type,required"`
	// This field can have the runtime type of [interface{}].
	WordList  interface{}                 `json:"word_list,required"`
	CreatedAt time.Time                   `json:"created_at" format:"date-time"`
	Pattern   Pattern                     `json:"pattern"`
	ProfileID string                      `json:"profile_id,nullable" format:"uuid"`
	Secret    bool                        `json:"secret"`
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      profileIntegrationEntryJSON `json:"-"`
	union     ProfileIntegrationEntriesUnion
}

// profileIntegrationEntryJSON contains the JSON metadata for the struct
// [ProfileIntegrationEntry]
type profileIntegrationEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	WordList    apijson.Field
	CreatedAt   apijson.Field
	Pattern     apijson.Field
	ProfileID   apijson.Field
	Secret      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r profileIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *ProfileIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = ProfileIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProfileIntegrationEntriesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [zero_trust.ProfileIntegrationEntriesCustom],
// [zero_trust.ProfileIntegrationEntriesPredefined],
// [zero_trust.ProfileIntegrationEntriesIntegration],
// [zero_trust.ProfileIntegrationEntriesExactData],
// [zero_trust.ProfileIntegrationEntriesWordList].
func (r ProfileIntegrationEntry) AsUnion() ProfileIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by [zero_trust.ProfileIntegrationEntriesCustom],
// [zero_trust.ProfileIntegrationEntriesPredefined],
// [zero_trust.ProfileIntegrationEntriesIntegration],
// [zero_trust.ProfileIntegrationEntriesExactData] or
// [zero_trust.ProfileIntegrationEntriesWordList].
type ProfileIntegrationEntriesUnion interface {
	implementsZeroTrustProfileIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProfileIntegrationEntriesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileIntegrationEntriesCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileIntegrationEntriesPredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileIntegrationEntriesIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileIntegrationEntriesExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProfileIntegrationEntriesWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type ProfileIntegrationEntriesCustom struct {
	ID        string                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                `json:"enabled,required"`
	Name      string                              `json:"name,required"`
	Pattern   Pattern                             `json:"pattern,required"`
	Type      ProfileIntegrationEntriesCustomType `json:"type,required"`
	UpdatedAt time.Time                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                              `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationEntriesCustomJSON `json:"-"`
}

// profileIntegrationEntriesCustomJSON contains the JSON metadata for the struct
// [ProfileIntegrationEntriesCustom]
type profileIntegrationEntriesCustomJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationEntriesCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationEntriesCustomJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationEntriesCustom) implementsZeroTrustProfileIntegrationEntry() {}

type ProfileIntegrationEntriesCustomType string

const (
	ProfileIntegrationEntriesCustomTypeCustom ProfileIntegrationEntriesCustomType = "custom"
)

func (r ProfileIntegrationEntriesCustomType) IsKnown() bool {
	switch r {
	case ProfileIntegrationEntriesCustomTypeCustom:
		return true
	}
	return false
}

type ProfileIntegrationEntriesPredefined struct {
	ID         string                                        `json:"id,required" format:"uuid"`
	Confidence ProfileIntegrationEntriesPredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                          `json:"enabled,required"`
	Name       string                                        `json:"name,required"`
	Type       ProfileIntegrationEntriesPredefinedType       `json:"type,required"`
	ProfileID  string                                        `json:"profile_id,nullable" format:"uuid"`
	JSON       profileIntegrationEntriesPredefinedJSON       `json:"-"`
}

// profileIntegrationEntriesPredefinedJSON contains the JSON metadata for the
// struct [ProfileIntegrationEntriesPredefined]
type profileIntegrationEntriesPredefinedJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationEntriesPredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationEntriesPredefinedJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationEntriesPredefined) implementsZeroTrustProfileIntegrationEntry() {}

type ProfileIntegrationEntriesPredefinedConfidence struct {
	// Indicates whether this entry can be made more or less sensitive by setting a
	// confidence threshold. Profiles that use an entry with `available` set to true
	// can use confidence thresholds
	Available bool                                              `json:"available,required"`
	JSON      profileIntegrationEntriesPredefinedConfidenceJSON `json:"-"`
}

// profileIntegrationEntriesPredefinedConfidenceJSON contains the JSON metadata for
// the struct [ProfileIntegrationEntriesPredefinedConfidence]
type profileIntegrationEntriesPredefinedConfidenceJSON struct {
	Available   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationEntriesPredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationEntriesPredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type ProfileIntegrationEntriesPredefinedType string

const (
	ProfileIntegrationEntriesPredefinedTypePredefined ProfileIntegrationEntriesPredefinedType = "predefined"
)

func (r ProfileIntegrationEntriesPredefinedType) IsKnown() bool {
	switch r {
	case ProfileIntegrationEntriesPredefinedTypePredefined:
		return true
	}
	return false
}

type ProfileIntegrationEntriesIntegration struct {
	ID        string                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                     `json:"enabled,required"`
	Name      string                                   `json:"name,required"`
	Type      ProfileIntegrationEntriesIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                `json:"updated_at,required" format:"date-time"`
	ProfileID string                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationEntriesIntegrationJSON `json:"-"`
}

// profileIntegrationEntriesIntegrationJSON contains the JSON metadata for the
// struct [ProfileIntegrationEntriesIntegration]
type profileIntegrationEntriesIntegrationJSON struct {
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

func (r *ProfileIntegrationEntriesIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationEntriesIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationEntriesIntegration) implementsZeroTrustProfileIntegrationEntry() {}

type ProfileIntegrationEntriesIntegrationType string

const (
	ProfileIntegrationEntriesIntegrationTypeIntegration ProfileIntegrationEntriesIntegrationType = "integration"
)

func (r ProfileIntegrationEntriesIntegrationType) IsKnown() bool {
	switch r {
	case ProfileIntegrationEntriesIntegrationTypeIntegration:
		return true
	}
	return false
}

type ProfileIntegrationEntriesExactData struct {
	ID        string                                 `json:"id,required" format:"uuid"`
	CreatedAt time.Time                              `json:"created_at,required" format:"date-time"`
	Enabled   bool                                   `json:"enabled,required"`
	Name      string                                 `json:"name,required"`
	Secret    bool                                   `json:"secret,required"`
	Type      ProfileIntegrationEntriesExactDataType `json:"type,required"`
	UpdatedAt time.Time                              `json:"updated_at,required" format:"date-time"`
	JSON      profileIntegrationEntriesExactDataJSON `json:"-"`
}

// profileIntegrationEntriesExactDataJSON contains the JSON metadata for the struct
// [ProfileIntegrationEntriesExactData]
type profileIntegrationEntriesExactDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfileIntegrationEntriesExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationEntriesExactDataJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationEntriesExactData) implementsZeroTrustProfileIntegrationEntry() {}

type ProfileIntegrationEntriesExactDataType string

const (
	ProfileIntegrationEntriesExactDataTypeExactData ProfileIntegrationEntriesExactDataType = "exact_data"
)

func (r ProfileIntegrationEntriesExactDataType) IsKnown() bool {
	switch r {
	case ProfileIntegrationEntriesExactDataTypeExactData:
		return true
	}
	return false
}

type ProfileIntegrationEntriesWordList struct {
	ID        string                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                  `json:"enabled,required"`
	Name      string                                `json:"name,required"`
	Type      ProfileIntegrationEntriesWordListType `json:"type,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                           `json:"word_list,required"`
	ProfileID string                                `json:"profile_id,nullable" format:"uuid"`
	JSON      profileIntegrationEntriesWordListJSON `json:"-"`
}

// profileIntegrationEntriesWordListJSON contains the JSON metadata for the struct
// [ProfileIntegrationEntriesWordList]
type profileIntegrationEntriesWordListJSON struct {
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

func (r *ProfileIntegrationEntriesWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r profileIntegrationEntriesWordListJSON) RawJSON() string {
	return r.raw
}

func (r ProfileIntegrationEntriesWordList) implementsZeroTrustProfileIntegrationEntry() {}

type ProfileIntegrationEntriesWordListType string

const (
	ProfileIntegrationEntriesWordListTypeWordList ProfileIntegrationEntriesWordListType = "word_list"
)

func (r ProfileIntegrationEntriesWordListType) IsKnown() bool {
	switch r {
	case ProfileIntegrationEntriesWordListTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationEntriesType string

const (
	ProfileIntegrationEntriesTypeCustom      ProfileIntegrationEntriesType = "custom"
	ProfileIntegrationEntriesTypePredefined  ProfileIntegrationEntriesType = "predefined"
	ProfileIntegrationEntriesTypeIntegration ProfileIntegrationEntriesType = "integration"
	ProfileIntegrationEntriesTypeExactData   ProfileIntegrationEntriesType = "exact_data"
	ProfileIntegrationEntriesTypeWordList    ProfileIntegrationEntriesType = "word_list"
)

func (r ProfileIntegrationEntriesType) IsKnown() bool {
	switch r {
	case ProfileIntegrationEntriesTypeCustom, ProfileIntegrationEntriesTypePredefined, ProfileIntegrationEntriesTypeIntegration, ProfileIntegrationEntriesTypeExactData, ProfileIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type ProfileIntegrationType string

const (
	ProfileIntegrationTypeIntegration ProfileIntegrationType = "integration"
)

func (r ProfileIntegrationType) IsKnown() bool {
	switch r {
	case ProfileIntegrationTypeIntegration:
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Profile                              `json:"result"`
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

// Whether the API call was successful
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
