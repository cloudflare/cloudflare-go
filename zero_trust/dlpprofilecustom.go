// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/tidwall/gjson"
)

// DLPProfileCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfileCustomService] method instead.
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

// Creates a DLP custom profile.
func (r *DLPProfileCustomService) New(ctx context.Context, params DLPProfileCustomNewParams, opts ...option.RequestOption) (res *DLPProfileCustomNewResponse, err error) {
	var env DLPProfileCustomNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
	var env DLPProfileCustomUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a DLP custom profile.
func (r *DLPProfileCustomService) Delete(ctx context.Context, profileID string, body DLPProfileCustomDeleteParams, opts ...option.RequestOption) (res *DLPProfileCustomDeleteResponse, err error) {
	var env DLPProfileCustomDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom DLP profile by id.
func (r *DLPProfileCustomService) Get(ctx context.Context, profileID string, query DLPProfileCustomGetParams, opts ...option.RequestOption) (res *DLPProfileCustomGetResponse, err error) {
	var env DLPProfileCustomGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/custom/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Pattern struct {
	Regex string `json:"regex,required"`
	// Deprecated: deprecated
	Validation PatternValidation `json:"validation"`
	JSON       patternJSON       `json:"-"`
}

// patternJSON contains the JSON metadata for the struct [Pattern]
type patternJSON struct {
	Regex       apijson.Field
	Validation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r patternJSON) RawJSON() string {
	return r.raw
}

type PatternValidation string

const (
	PatternValidationLuhn PatternValidation = "luhn"
)

func (r PatternValidation) IsKnown() bool {
	switch r {
	case PatternValidationLuhn:
		return true
	}
	return false
}

type PatternParam struct {
	Regex param.Field[string] `json:"regex,required"`
	// Deprecated: deprecated
	Validation param.Field[PatternValidation] `json:"validation"`
}

func (r PatternParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewResponse struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile.
	Name             string                          `json:"name,required"`
	Type             DLPProfileCustomNewResponseType `json:"type,required"`
	AIContextEnabled bool                            `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64                                          `json:"allowed_match_count"`
	ConfidenceThreshold DLPProfileCustomNewResponseConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomNewResponseCustomEntry],
	// [[]DLPProfileCustomNewResponsePredefinedEntry],
	// [[]DLPProfileCustomNewResponseIntegrationEntry].
	Entries    interface{} `json:"entries"`
	OCREnabled bool        `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool `json:"open_access"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomNewResponseCustomSharedEntry],
	// [[]DLPProfileCustomNewResponseIntegrationSharedEntry].
	SharedEntries interface{} `json:"shared_entries"`
	// When the profile was lasted updated.
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomNewResponseJSON `json:"-"`
	union     DLPProfileCustomNewResponseUnion
}

// dlpProfileCustomNewResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomNewResponse]
type dlpProfileCustomNewResponseJSON struct {
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

func (r dlpProfileCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [DLPProfileCustomNewResponseCustom],
// [DLPProfileCustomNewResponsePredefined],
// [DLPProfileCustomNewResponseIntegration].
func (r DLPProfileCustomNewResponse) AsUnion() DLPProfileCustomNewResponseUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomNewResponseCustom],
// [DLPProfileCustomNewResponsePredefined] or
// [DLPProfileCustomNewResponseIntegration].
type DLPProfileCustomNewResponseUnion interface {
	implementsDLPProfileCustomNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomNewResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileCustomNewResponseCustom struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the profile.
	Name       string                                `json:"name,required"`
	OCREnabled bool                                  `json:"ocr_enabled,required"`
	Type       DLPProfileCustomNewResponseCustomType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time                                            `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool                                                 `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomNewResponseCustomConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// Deprecated: deprecated
	Entries       []DLPProfileCustomNewResponseCustomEntry       `json:"entries"`
	SharedEntries []DLPProfileCustomNewResponseCustomSharedEntry `json:"shared_entries"`
	JSON          dlpProfileCustomNewResponseCustomJSON          `json:"-"`
}

// dlpProfileCustomNewResponseCustomJSON contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustom]
type dlpProfileCustomNewResponseCustomJSON struct {
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

func (r *DLPProfileCustomNewResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustom) implementsDLPProfileCustomNewResponse() {}

type DLPProfileCustomNewResponseCustomType string

const (
	DLPProfileCustomNewResponseCustomTypeCustom DLPProfileCustomNewResponseCustomType = "custom"
)

func (r DLPProfileCustomNewResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomConfidenceThreshold string

const (
	DLPProfileCustomNewResponseCustomConfidenceThresholdLow      DLPProfileCustomNewResponseCustomConfidenceThreshold = "low"
	DLPProfileCustomNewResponseCustomConfidenceThresholdMedium   DLPProfileCustomNewResponseCustomConfidenceThreshold = "medium"
	DLPProfileCustomNewResponseCustomConfidenceThresholdHigh     DLPProfileCustomNewResponseCustomConfidenceThreshold = "high"
	DLPProfileCustomNewResponseCustomConfidenceThresholdVeryHigh DLPProfileCustomNewResponseCustomConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomNewResponseCustomConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomConfidenceThresholdLow, DLPProfileCustomNewResponseCustomConfidenceThresholdMedium, DLPProfileCustomNewResponseCustomConfidenceThresholdHigh, DLPProfileCustomNewResponseCustomConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                         `json:"enabled,required"`
	Name    string                                       `json:"name,required"`
	Type    DLPProfileCustomNewResponseCustomEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseCustomEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                `json:"word_list"`
	JSON     dlpProfileCustomNewResponseCustomEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseCustomEntriesUnion
}

// dlpProfileCustomNewResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseCustomEntry]
type dlpProfileCustomNewResponseCustomEntryJSON struct {
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

func (r dlpProfileCustomNewResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseCustomEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponseCustomEntriesCustomEntry],
// [DLPProfileCustomNewResponseCustomEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseCustomEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseCustomEntriesExactDataEntry],
// [DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntry],
// [DLPProfileCustomNewResponseCustomEntriesWordListEntry].
func (r DLPProfileCustomNewResponseCustomEntry) AsUnion() DLPProfileCustomNewResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomNewResponseCustomEntriesCustomEntry],
// [DLPProfileCustomNewResponseCustomEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseCustomEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseCustomEntriesExactDataEntry],
// [DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomNewResponseCustomEntriesWordListEntry].
type DLPProfileCustomNewResponseCustomEntriesUnion interface {
	implementsDLPProfileCustomNewResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseCustomEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponseCustomEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                    `json:"enabled,required"`
	Name        string                                                  `json:"name,required"`
	Pattern     Pattern                                                 `json:"pattern,required"`
	Type        DLPProfileCustomNewResponseCustomEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                               `json:"updated_at,required" format:"date-time"`
	Description string                                                  `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesCustomEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseCustomEntriesCustomEntry]
type dlpProfileCustomNewResponseCustomEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesCustomEntry) implementsDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponseCustomEntriesCustomEntryTypeCustom DLPProfileCustomNewResponseCustomEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponseCustomEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesPredefinedEntry struct {
	ID         string                                                            `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseCustomEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                              `json:"enabled,required"`
	Name       string                                                            `json:"name,required"`
	Type       DLPProfileCustomNewResponseCustomEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomNewResponseCustomEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomEntriesPredefinedEntry]
type dlpProfileCustomNewResponseCustomEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesPredefinedEntry) implementsDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                  `json:"available,required"`
	JSON      dlpProfileCustomNewResponseCustomEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesPredefinedEntryConfidenceJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponseCustomEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseCustomEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponseCustomEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponseCustomEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponseCustomEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                  `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseCustomEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariant]
type dlpProfileCustomNewResponseCustomEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesIntegrationEntry struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomEntriesIntegrationEntry]
type dlpProfileCustomNewResponseCustomEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesIntegrationEntry) implementsDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponseCustomEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponseCustomEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponseCustomEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                       `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                       `json:"enabled,required"`
	Name          string                                                     `json:"name,required"`
	Secret        bool                                                       `json:"secret,required"`
	Type          DLPProfileCustomNewResponseCustomEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                  `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponseCustomEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesExactDataEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseCustomEntriesExactDataEntry]
type dlpProfileCustomNewResponseCustomEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesExactDataEntry) implementsDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponseCustomEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponseCustomEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponseCustomEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntry struct {
	ID        string                                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                 `json:"enabled,required"`
	Name      string                                                               `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                            `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntry]
type dlpProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntry) implementsDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesWordListEntry struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                               `json:"word_list,required"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomEntriesWordListEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomNewResponseCustomEntriesWordListEntry]
type dlpProfileCustomNewResponseCustomEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomEntriesWordListEntry) implementsDLPProfileCustomNewResponseCustomEntry() {
}

type DLPProfileCustomNewResponseCustomEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponseCustomEntriesWordListEntryTypeWordList DLPProfileCustomNewResponseCustomEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomEntriesType string

const (
	DLPProfileCustomNewResponseCustomEntriesTypeCustom              DLPProfileCustomNewResponseCustomEntriesType = "custom"
	DLPProfileCustomNewResponseCustomEntriesTypePredefined          DLPProfileCustomNewResponseCustomEntriesType = "predefined"
	DLPProfileCustomNewResponseCustomEntriesTypeIntegration         DLPProfileCustomNewResponseCustomEntriesType = "integration"
	DLPProfileCustomNewResponseCustomEntriesTypeExactData           DLPProfileCustomNewResponseCustomEntriesType = "exact_data"
	DLPProfileCustomNewResponseCustomEntriesTypeDocumentFingerprint DLPProfileCustomNewResponseCustomEntriesType = "document_fingerprint"
	DLPProfileCustomNewResponseCustomEntriesTypeWordList            DLPProfileCustomNewResponseCustomEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomEntriesTypeCustom, DLPProfileCustomNewResponseCustomEntriesTypePredefined, DLPProfileCustomNewResponseCustomEntriesTypeIntegration, DLPProfileCustomNewResponseCustomEntriesTypeExactData, DLPProfileCustomNewResponseCustomEntriesTypeDocumentFingerprint, DLPProfileCustomNewResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                               `json:"enabled,required"`
	Name    string                                             `json:"name,required"`
	Type    DLPProfileCustomNewResponseCustomSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                      `json:"word_list"`
	JSON     dlpProfileCustomNewResponseCustomSharedEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseCustomSharedEntriesUnion
}

// dlpProfileCustomNewResponseCustomSharedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseCustomSharedEntry]
type dlpProfileCustomNewResponseCustomSharedEntryJSON struct {
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

func (r dlpProfileCustomNewResponseCustomSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseCustomSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseCustomSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseCustomSharedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponseCustomSharedEntriesCustomEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesWordListEntry].
func (r DLPProfileCustomNewResponseCustomSharedEntry) AsUnion() DLPProfileCustomNewResponseCustomSharedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomNewResponseCustomSharedEntriesCustomEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomNewResponseCustomSharedEntriesWordListEntry].
type DLPProfileCustomNewResponseCustomSharedEntriesUnion interface {
	implementsDLPProfileCustomNewResponseCustomSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseCustomSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseCustomSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponseCustomSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                          `json:"enabled,required"`
	Name        string                                                        `json:"name,required"`
	Pattern     Pattern                                                       `json:"pattern,required"`
	Type        DLPProfileCustomNewResponseCustomSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                     `json:"updated_at,required" format:"date-time"`
	Description string                                                        `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesCustomEntry]
type dlpProfileCustomNewResponseCustomSharedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomSharedEntriesCustomEntry) implementsDLPProfileCustomNewResponseCustomSharedEntry() {
}

type DLPProfileCustomNewResponseCustomSharedEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesCustomEntryTypeCustom DLPProfileCustomNewResponseCustomSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntry struct {
	ID         string                                                                  `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                    `json:"enabled,required"`
	Name       string                                                                  `json:"name,required"`
	Type       DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                               `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntry]
type dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntry) implementsDLPProfileCustomNewResponseCustomSharedEntry() {
}

type DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                        `json:"available,required"`
	JSON      dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                        `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariant]
type dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntry struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntry]
type dlpProfileCustomNewResponseCustomSharedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntry) implementsDLPProfileCustomNewResponseCustomSharedEntry() {
}

type DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                             `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                             `json:"enabled,required"`
	Name          string                                                           `json:"name,required"`
	Secret        bool                                                             `json:"secret,required"`
	Type          DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                        `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponseCustomSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntry]
type dlpProfileCustomNewResponseCustomSharedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntry) implementsDLPProfileCustomNewResponseCustomSharedEntry() {
}

type DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                       `json:"enabled,required"`
	Name      string                                                                     `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                  `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntry]
type dlpProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomNewResponseCustomSharedEntry() {
}

type DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesWordListEntry struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfileCustomNewResponseCustomSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                       `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                     `json:"word_list,required"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseCustomSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseCustomSharedEntriesWordListEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseCustomSharedEntriesWordListEntry]
type dlpProfileCustomNewResponseCustomSharedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseCustomSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseCustomSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseCustomSharedEntriesWordListEntry) implementsDLPProfileCustomNewResponseCustomSharedEntry() {
}

type DLPProfileCustomNewResponseCustomSharedEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesWordListEntryTypeWordList DLPProfileCustomNewResponseCustomSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseCustomSharedEntriesType string

const (
	DLPProfileCustomNewResponseCustomSharedEntriesTypeCustom              DLPProfileCustomNewResponseCustomSharedEntriesType = "custom"
	DLPProfileCustomNewResponseCustomSharedEntriesTypePredefined          DLPProfileCustomNewResponseCustomSharedEntriesType = "predefined"
	DLPProfileCustomNewResponseCustomSharedEntriesTypeIntegration         DLPProfileCustomNewResponseCustomSharedEntriesType = "integration"
	DLPProfileCustomNewResponseCustomSharedEntriesTypeExactData           DLPProfileCustomNewResponseCustomSharedEntriesType = "exact_data"
	DLPProfileCustomNewResponseCustomSharedEntriesTypeDocumentFingerprint DLPProfileCustomNewResponseCustomSharedEntriesType = "document_fingerprint"
	DLPProfileCustomNewResponseCustomSharedEntriesTypeWordList            DLPProfileCustomNewResponseCustomSharedEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseCustomSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseCustomSharedEntriesTypeCustom, DLPProfileCustomNewResponseCustomSharedEntriesTypePredefined, DLPProfileCustomNewResponseCustomSharedEntriesTypeIntegration, DLPProfileCustomNewResponseCustomSharedEntriesTypeExactData, DLPProfileCustomNewResponseCustomSharedEntriesTypeDocumentFingerprint, DLPProfileCustomNewResponseCustomSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefined struct {
	// The id of the predefined profile (uuid).
	ID                string `json:"id,required" format:"uuid"`
	AllowedMatchCount int64  `json:"allowed_match_count,required"`
	// Deprecated: deprecated
	Entries []DLPProfileCustomNewResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                                                   `json:"name,required"`
	Type                DLPProfileCustomNewResponsePredefinedType                `json:"type,required"`
	AIContextEnabled    bool                                                     `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomNewResponsePredefinedConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                                      `json:"open_access"`
	JSON       dlpProfileCustomNewResponsePredefinedJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponsePredefined]
type dlpProfileCustomNewResponsePredefinedJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefined) implementsDLPProfileCustomNewResponse() {}

type DLPProfileCustomNewResponsePredefinedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                             `json:"enabled,required"`
	Name    string                                           `json:"name,required"`
	Type    DLPProfileCustomNewResponsePredefinedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                    `json:"word_list"`
	JSON     dlpProfileCustomNewResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponsePredefinedEntriesUnion
}

// dlpProfileCustomNewResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponsePredefinedEntry]
type dlpProfileCustomNewResponsePredefinedEntryJSON struct {
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

func (r dlpProfileCustomNewResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponsePredefinedEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponsePredefinedEntriesCustomEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesExactDataEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesWordListEntry].
func (r DLPProfileCustomNewResponsePredefinedEntry) AsUnion() DLPProfileCustomNewResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomNewResponsePredefinedEntriesCustomEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesExactDataEntry],
// [DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomNewResponsePredefinedEntriesWordListEntry].
type DLPProfileCustomNewResponsePredefinedEntriesUnion interface {
	implementsDLPProfileCustomNewResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponsePredefinedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponsePredefinedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponsePredefinedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                        `json:"enabled,required"`
	Name        string                                                      `json:"name,required"`
	Pattern     Pattern                                                     `json:"pattern,required"`
	Type        DLPProfileCustomNewResponsePredefinedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                   `json:"updated_at,required" format:"date-time"`
	Description string                                                      `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesCustomEntry]
type dlpProfileCustomNewResponsePredefinedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesCustomEntry) implementsDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesCustomEntryTypeCustom DLPProfileCustomNewResponsePredefinedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntry struct {
	ID         string                                                                `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                  `json:"enabled,required"`
	Name       string                                                                `json:"name,required"`
	Type       DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntry]
type dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntry) implementsDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                      `json:"available,required"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                      `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariant]
type dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntry struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntry]
type dlpProfileCustomNewResponsePredefinedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntry) implementsDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                           `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                           `json:"enabled,required"`
	Name          string                                                         `json:"name,required"`
	Secret        bool                                                           `json:"secret,required"`
	Type          DLPProfileCustomNewResponsePredefinedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                      `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponsePredefinedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesExactDataEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesExactDataEntry]
type dlpProfileCustomNewResponsePredefinedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesExactDataEntry) implementsDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponsePredefinedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                     `json:"enabled,required"`
	Name      string                                                                   `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntry]
type dlpProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesWordListEntry struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Type      DLPProfileCustomNewResponsePredefinedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                   `json:"word_list,required"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponsePredefinedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponsePredefinedEntriesWordListEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponsePredefinedEntriesWordListEntry]
type dlpProfileCustomNewResponsePredefinedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomNewResponsePredefinedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponsePredefinedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponsePredefinedEntriesWordListEntry) implementsDLPProfileCustomNewResponsePredefinedEntry() {
}

type DLPProfileCustomNewResponsePredefinedEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesWordListEntryTypeWordList DLPProfileCustomNewResponsePredefinedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedEntriesType string

const (
	DLPProfileCustomNewResponsePredefinedEntriesTypeCustom              DLPProfileCustomNewResponsePredefinedEntriesType = "custom"
	DLPProfileCustomNewResponsePredefinedEntriesTypePredefined          DLPProfileCustomNewResponsePredefinedEntriesType = "predefined"
	DLPProfileCustomNewResponsePredefinedEntriesTypeIntegration         DLPProfileCustomNewResponsePredefinedEntriesType = "integration"
	DLPProfileCustomNewResponsePredefinedEntriesTypeExactData           DLPProfileCustomNewResponsePredefinedEntriesType = "exact_data"
	DLPProfileCustomNewResponsePredefinedEntriesTypeDocumentFingerprint DLPProfileCustomNewResponsePredefinedEntriesType = "document_fingerprint"
	DLPProfileCustomNewResponsePredefinedEntriesTypeWordList            DLPProfileCustomNewResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedEntriesTypeCustom, DLPProfileCustomNewResponsePredefinedEntriesTypePredefined, DLPProfileCustomNewResponsePredefinedEntriesTypeIntegration, DLPProfileCustomNewResponsePredefinedEntriesTypeExactData, DLPProfileCustomNewResponsePredefinedEntriesTypeDocumentFingerprint, DLPProfileCustomNewResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedType string

const (
	DLPProfileCustomNewResponsePredefinedTypePredefined DLPProfileCustomNewResponsePredefinedType = "predefined"
)

func (r DLPProfileCustomNewResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponsePredefinedConfidenceThreshold string

const (
	DLPProfileCustomNewResponsePredefinedConfidenceThresholdLow      DLPProfileCustomNewResponsePredefinedConfidenceThreshold = "low"
	DLPProfileCustomNewResponsePredefinedConfidenceThresholdMedium   DLPProfileCustomNewResponsePredefinedConfidenceThreshold = "medium"
	DLPProfileCustomNewResponsePredefinedConfidenceThresholdHigh     DLPProfileCustomNewResponsePredefinedConfidenceThreshold = "high"
	DLPProfileCustomNewResponsePredefinedConfidenceThresholdVeryHigh DLPProfileCustomNewResponsePredefinedConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomNewResponsePredefinedConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponsePredefinedConfidenceThresholdLow, DLPProfileCustomNewResponsePredefinedConfidenceThresholdMedium, DLPProfileCustomNewResponsePredefinedConfidenceThresholdHigh, DLPProfileCustomNewResponsePredefinedConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegration struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Entries       []DLPProfileCustomNewResponseIntegrationEntry       `json:"entries,required"`
	Name          string                                              `json:"name,required"`
	SharedEntries []DLPProfileCustomNewResponseIntegrationSharedEntry `json:"shared_entries,required"`
	Type          DLPProfileCustomNewResponseIntegrationType          `json:"type,required"`
	UpdatedAt     time.Time                                           `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                                     `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseIntegrationJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseIntegration]
type dlpProfileCustomNewResponseIntegrationJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegration) implementsDLPProfileCustomNewResponse() {}

type DLPProfileCustomNewResponseIntegrationEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                              `json:"enabled,required"`
	Name    string                                            `json:"name,required"`
	Type    DLPProfileCustomNewResponseIntegrationEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                     `json:"word_list"`
	JSON     dlpProfileCustomNewResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseIntegrationEntriesUnion
}

// dlpProfileCustomNewResponseIntegrationEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseIntegrationEntry]
type dlpProfileCustomNewResponseIntegrationEntryJSON struct {
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

func (r dlpProfileCustomNewResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseIntegrationEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponseIntegrationEntriesCustomEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesExactDataEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesWordListEntry].
func (r DLPProfileCustomNewResponseIntegrationEntry) AsUnion() DLPProfileCustomNewResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomNewResponseIntegrationEntriesCustomEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesExactDataEntry],
// [DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomNewResponseIntegrationEntriesWordListEntry].
type DLPProfileCustomNewResponseIntegrationEntriesUnion interface {
	implementsDLPProfileCustomNewResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseIntegrationEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponseIntegrationEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                         `json:"enabled,required"`
	Name        string                                                       `json:"name,required"`
	Pattern     Pattern                                                      `json:"pattern,required"`
	Type        DLPProfileCustomNewResponseIntegrationEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                    `json:"updated_at,required" format:"date-time"`
	Description string                                                       `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesCustomEntry]
type dlpProfileCustomNewResponseIntegrationEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesCustomEntry) implementsDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesCustomEntryTypeCustom DLPProfileCustomNewResponseIntegrationEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntry struct {
	ID         string                                                                 `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                   `json:"enabled,required"`
	Name       string                                                                 `json:"name,required"`
	Type       DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntry]
type dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntry) implementsDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                       `json:"available,required"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                       `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariant]
type dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntry struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                         `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntry]
type dlpProfileCustomNewResponseIntegrationEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntry) implementsDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                            `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                            `json:"enabled,required"`
	Name          string                                                          `json:"name,required"`
	Secret        bool                                                            `json:"secret,required"`
	Type          DLPProfileCustomNewResponseIntegrationEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                       `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponseIntegrationEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesExactDataEntry]
type dlpProfileCustomNewResponseIntegrationEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesExactDataEntry) implementsDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponseIntegrationEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntry struct {
	ID        string                                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                      `json:"enabled,required"`
	Name      string                                                                    `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                 `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntry]
type dlpProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntry) implementsDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesWordListEntry struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                    `json:"word_list,required"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationEntriesWordListEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomNewResponseIntegrationEntriesWordListEntry]
type dlpProfileCustomNewResponseIntegrationEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationEntriesWordListEntry) implementsDLPProfileCustomNewResponseIntegrationEntry() {
}

type DLPProfileCustomNewResponseIntegrationEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesWordListEntryTypeWordList DLPProfileCustomNewResponseIntegrationEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationEntriesType string

const (
	DLPProfileCustomNewResponseIntegrationEntriesTypeCustom              DLPProfileCustomNewResponseIntegrationEntriesType = "custom"
	DLPProfileCustomNewResponseIntegrationEntriesTypePredefined          DLPProfileCustomNewResponseIntegrationEntriesType = "predefined"
	DLPProfileCustomNewResponseIntegrationEntriesTypeIntegration         DLPProfileCustomNewResponseIntegrationEntriesType = "integration"
	DLPProfileCustomNewResponseIntegrationEntriesTypeExactData           DLPProfileCustomNewResponseIntegrationEntriesType = "exact_data"
	DLPProfileCustomNewResponseIntegrationEntriesTypeDocumentFingerprint DLPProfileCustomNewResponseIntegrationEntriesType = "document_fingerprint"
	DLPProfileCustomNewResponseIntegrationEntriesTypeWordList            DLPProfileCustomNewResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationEntriesTypeCustom, DLPProfileCustomNewResponseIntegrationEntriesTypePredefined, DLPProfileCustomNewResponseIntegrationEntriesTypeIntegration, DLPProfileCustomNewResponseIntegrationEntriesTypeExactData, DLPProfileCustomNewResponseIntegrationEntriesTypeDocumentFingerprint, DLPProfileCustomNewResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                                    `json:"enabled,required"`
	Name    string                                                  `json:"name,required"`
	Type    DLPProfileCustomNewResponseIntegrationSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                           `json:"word_list"`
	JSON     dlpProfileCustomNewResponseIntegrationSharedEntryJSON `json:"-"`
	union    DLPProfileCustomNewResponseIntegrationSharedEntriesUnion
}

// dlpProfileCustomNewResponseIntegrationSharedEntryJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseIntegrationSharedEntry]
type dlpProfileCustomNewResponseIntegrationSharedEntryJSON struct {
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

func (r dlpProfileCustomNewResponseIntegrationSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomNewResponseIntegrationSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomNewResponseIntegrationSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomNewResponseIntegrationSharedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntry].
func (r DLPProfileCustomNewResponseIntegrationSharedEntry) AsUnion() DLPProfileCustomNewResponseIntegrationSharedEntriesUnion {
	return r.union
}

// Union satisfied by
// [DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntry].
type DLPProfileCustomNewResponseIntegrationSharedEntriesUnion interface {
	implementsDLPProfileCustomNewResponseIntegrationSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomNewResponseIntegrationSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                               `json:"enabled,required"`
	Name        string                                                             `json:"name,required"`
	Pattern     Pattern                                                            `json:"pattern,required"`
	Type        DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                          `json:"updated_at,required" format:"date-time"`
	Description string                                                             `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesCustomEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntry]
type dlpProfileCustomNewResponseIntegrationSharedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntry) implementsDLPProfileCustomNewResponseIntegrationSharedEntry() {
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntryType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntryTypeCustom DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntry struct {
	ID         string                                                                       `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                         `json:"enabled,required"`
	Name       string                                                                       `json:"name,required"`
	Type       DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                                    `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntry]
type dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntry) implementsDLPProfileCustomNewResponseIntegrationSharedEntry() {
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                             `json:"available,required"`
	JSON      dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidence]
type dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                             `json:"description,nullable"`
	JSON        dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariant]
type dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntry struct {
	ID        string                                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                    `json:"enabled,required"`
	Name      string                                                                  `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntry]
type dlpProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntry) implementsDLPProfileCustomNewResponseIntegrationSharedEntry() {
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                                  `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                             `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                                  `json:"enabled,required"`
	Name          string                                                                `json:"name,required"`
	Secret        bool                                                                  `json:"secret,required"`
	Type          DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                             `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntry]
type dlpProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntry) implementsDLPProfileCustomNewResponseIntegrationSharedEntry() {
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryTypeExactData DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                            `json:"enabled,required"`
	Name      string                                                                          `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                       `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntry]
type dlpProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomNewResponseIntegrationSharedEntry() {
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntry struct {
	ID        string                                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                 `json:"enabled,required"`
	Name      string                                                               `json:"name,required"`
	Type      DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                          `json:"word_list,required"`
	ProfileID string                                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomNewResponseIntegrationSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomNewResponseIntegrationSharedEntriesWordListEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntry]
type dlpProfileCustomNewResponseIntegrationSharedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseIntegrationSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntry) implementsDLPProfileCustomNewResponseIntegrationSharedEntry() {
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntryType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntryTypeWordList DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationSharedEntriesType string

const (
	DLPProfileCustomNewResponseIntegrationSharedEntriesTypeCustom              DLPProfileCustomNewResponseIntegrationSharedEntriesType = "custom"
	DLPProfileCustomNewResponseIntegrationSharedEntriesTypePredefined          DLPProfileCustomNewResponseIntegrationSharedEntriesType = "predefined"
	DLPProfileCustomNewResponseIntegrationSharedEntriesTypeIntegration         DLPProfileCustomNewResponseIntegrationSharedEntriesType = "integration"
	DLPProfileCustomNewResponseIntegrationSharedEntriesTypeExactData           DLPProfileCustomNewResponseIntegrationSharedEntriesType = "exact_data"
	DLPProfileCustomNewResponseIntegrationSharedEntriesTypeDocumentFingerprint DLPProfileCustomNewResponseIntegrationSharedEntriesType = "document_fingerprint"
	DLPProfileCustomNewResponseIntegrationSharedEntriesTypeWordList            DLPProfileCustomNewResponseIntegrationSharedEntriesType = "word_list"
)

func (r DLPProfileCustomNewResponseIntegrationSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationSharedEntriesTypeCustom, DLPProfileCustomNewResponseIntegrationSharedEntriesTypePredefined, DLPProfileCustomNewResponseIntegrationSharedEntriesTypeIntegration, DLPProfileCustomNewResponseIntegrationSharedEntriesTypeExactData, DLPProfileCustomNewResponseIntegrationSharedEntriesTypeDocumentFingerprint, DLPProfileCustomNewResponseIntegrationSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseIntegrationType string

const (
	DLPProfileCustomNewResponseIntegrationTypeIntegration DLPProfileCustomNewResponseIntegrationType = "integration"
)

func (r DLPProfileCustomNewResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseType string

const (
	DLPProfileCustomNewResponseTypeCustom      DLPProfileCustomNewResponseType = "custom"
	DLPProfileCustomNewResponseTypePredefined  DLPProfileCustomNewResponseType = "predefined"
	DLPProfileCustomNewResponseTypeIntegration DLPProfileCustomNewResponseType = "integration"
)

func (r DLPProfileCustomNewResponseType) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseTypeCustom, DLPProfileCustomNewResponseTypePredefined, DLPProfileCustomNewResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomNewResponseConfidenceThreshold string

const (
	DLPProfileCustomNewResponseConfidenceThresholdLow      DLPProfileCustomNewResponseConfidenceThreshold = "low"
	DLPProfileCustomNewResponseConfidenceThresholdMedium   DLPProfileCustomNewResponseConfidenceThreshold = "medium"
	DLPProfileCustomNewResponseConfidenceThresholdHigh     DLPProfileCustomNewResponseConfidenceThreshold = "high"
	DLPProfileCustomNewResponseConfidenceThresholdVeryHigh DLPProfileCustomNewResponseConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomNewResponseConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomNewResponseConfidenceThresholdLow, DLPProfileCustomNewResponseConfidenceThresholdMedium, DLPProfileCustomNewResponseConfidenceThresholdHigh, DLPProfileCustomNewResponseConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponse struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile.
	Name             string                             `json:"name,required"`
	Type             DLPProfileCustomUpdateResponseType `json:"type,required"`
	AIContextEnabled bool                               `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64                                             `json:"allowed_match_count"`
	ConfidenceThreshold DLPProfileCustomUpdateResponseConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomUpdateResponseCustomEntry],
	// [[]DLPProfileCustomUpdateResponsePredefinedEntry],
	// [[]DLPProfileCustomUpdateResponseIntegrationEntry].
	Entries    interface{} `json:"entries"`
	OCREnabled bool        `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool `json:"open_access"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomUpdateResponseCustomSharedEntry],
	// [[]DLPProfileCustomUpdateResponseIntegrationSharedEntry].
	SharedEntries interface{} `json:"shared_entries"`
	// When the profile was lasted updated.
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseJSON `json:"-"`
	union     DLPProfileCustomUpdateResponseUnion
}

// dlpProfileCustomUpdateResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponse]
type dlpProfileCustomUpdateResponseJSON struct {
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

func (r dlpProfileCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [DLPProfileCustomUpdateResponseCustom],
// [DLPProfileCustomUpdateResponsePredefined],
// [DLPProfileCustomUpdateResponseIntegration].
func (r DLPProfileCustomUpdateResponse) AsUnion() DLPProfileCustomUpdateResponseUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomUpdateResponseCustom],
// [DLPProfileCustomUpdateResponsePredefined] or
// [DLPProfileCustomUpdateResponseIntegration].
type DLPProfileCustomUpdateResponseUnion interface {
	implementsDLPProfileCustomUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomUpdateResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileCustomUpdateResponseCustom struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the profile.
	Name       string                                   `json:"name,required"`
	OCREnabled bool                                     `json:"ocr_enabled,required"`
	Type       DLPProfileCustomUpdateResponseCustomType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time                                               `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool                                                    `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomUpdateResponseCustomConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// Deprecated: deprecated
	Entries       []DLPProfileCustomUpdateResponseCustomEntry       `json:"entries"`
	SharedEntries []DLPProfileCustomUpdateResponseCustomSharedEntry `json:"shared_entries"`
	JSON          dlpProfileCustomUpdateResponseCustomJSON          `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseCustom]
type dlpProfileCustomUpdateResponseCustomJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustom) implementsDLPProfileCustomUpdateResponse() {}

type DLPProfileCustomUpdateResponseCustomType string

const (
	DLPProfileCustomUpdateResponseCustomTypeCustom DLPProfileCustomUpdateResponseCustomType = "custom"
)

func (r DLPProfileCustomUpdateResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomConfidenceThreshold string

const (
	DLPProfileCustomUpdateResponseCustomConfidenceThresholdLow      DLPProfileCustomUpdateResponseCustomConfidenceThreshold = "low"
	DLPProfileCustomUpdateResponseCustomConfidenceThresholdMedium   DLPProfileCustomUpdateResponseCustomConfidenceThreshold = "medium"
	DLPProfileCustomUpdateResponseCustomConfidenceThresholdHigh     DLPProfileCustomUpdateResponseCustomConfidenceThreshold = "high"
	DLPProfileCustomUpdateResponseCustomConfidenceThresholdVeryHigh DLPProfileCustomUpdateResponseCustomConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomUpdateResponseCustomConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomConfidenceThresholdLow, DLPProfileCustomUpdateResponseCustomConfidenceThresholdMedium, DLPProfileCustomUpdateResponseCustomConfidenceThresholdHigh, DLPProfileCustomUpdateResponseCustomConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                            `json:"enabled,required"`
	Name    string                                          `json:"name,required"`
	Type    DLPProfileCustomUpdateResponseCustomEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                   `json:"word_list"`
	JSON     dlpProfileCustomUpdateResponseCustomEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponseCustomEntriesUnion
}

// dlpProfileCustomUpdateResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseCustomEntry]
type dlpProfileCustomUpdateResponseCustomEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseCustomEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomUpdateResponseCustomEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesWordListEntry].
func (r DLPProfileCustomUpdateResponseCustomEntry) AsUnion() DLPProfileCustomUpdateResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomUpdateResponseCustomEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomUpdateResponseCustomEntriesWordListEntry].
type DLPProfileCustomUpdateResponseCustomEntriesUnion interface {
	implementsDLPProfileCustomUpdateResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseCustomEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomUpdateResponseCustomEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                       `json:"enabled,required"`
	Name        string                                                     `json:"name,required"`
	Pattern     Pattern                                                    `json:"pattern,required"`
	Type        DLPProfileCustomUpdateResponseCustomEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                  `json:"updated_at,required" format:"date-time"`
	Description string                                                     `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesCustomEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseCustomEntriesCustomEntry]
type dlpProfileCustomUpdateResponseCustomEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesCustomEntry) implementsDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesCustomEntryType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesCustomEntryTypeCustom DLPProfileCustomUpdateResponseCustomEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntry struct {
	ID         string                                                               `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                 `json:"enabled,required"`
	Name       string                                                               `json:"name,required"`
	Type       DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                            `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntry]
type dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntry) implementsDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                     `json:"available,required"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidence]
type dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryTypePredefined DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                     `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariant]
type dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntry struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                       `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntry]
type dlpProfileCustomUpdateResponseCustomEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntry) implementsDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntryType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntryTypeIntegration DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                          `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                          `json:"enabled,required"`
	Name          string                                                        `json:"name,required"`
	Secret        bool                                                          `json:"secret,required"`
	Type          DLPProfileCustomUpdateResponseCustomEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                     `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomUpdateResponseCustomEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesExactDataEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesExactDataEntry]
type dlpProfileCustomUpdateResponseCustomEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesExactDataEntry) implementsDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesExactDataEntryType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesExactDataEntryTypeExactData DLPProfileCustomUpdateResponseCustomEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntry struct {
	ID        string                                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                    `json:"enabled,required"`
	Name      string                                                                  `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                               `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntry]
type dlpProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntry) implementsDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesWordListEntry struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                  `json:"word_list,required"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomEntriesWordListEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponseCustomEntriesWordListEntry]
type dlpProfileCustomUpdateResponseCustomEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomEntriesWordListEntry) implementsDLPProfileCustomUpdateResponseCustomEntry() {
}

type DLPProfileCustomUpdateResponseCustomEntriesWordListEntryType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesWordListEntryTypeWordList DLPProfileCustomUpdateResponseCustomEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomEntriesType string

const (
	DLPProfileCustomUpdateResponseCustomEntriesTypeCustom              DLPProfileCustomUpdateResponseCustomEntriesType = "custom"
	DLPProfileCustomUpdateResponseCustomEntriesTypePredefined          DLPProfileCustomUpdateResponseCustomEntriesType = "predefined"
	DLPProfileCustomUpdateResponseCustomEntriesTypeIntegration         DLPProfileCustomUpdateResponseCustomEntriesType = "integration"
	DLPProfileCustomUpdateResponseCustomEntriesTypeExactData           DLPProfileCustomUpdateResponseCustomEntriesType = "exact_data"
	DLPProfileCustomUpdateResponseCustomEntriesTypeDocumentFingerprint DLPProfileCustomUpdateResponseCustomEntriesType = "document_fingerprint"
	DLPProfileCustomUpdateResponseCustomEntriesTypeWordList            DLPProfileCustomUpdateResponseCustomEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomEntriesTypeCustom, DLPProfileCustomUpdateResponseCustomEntriesTypePredefined, DLPProfileCustomUpdateResponseCustomEntriesTypeIntegration, DLPProfileCustomUpdateResponseCustomEntriesTypeExactData, DLPProfileCustomUpdateResponseCustomEntriesTypeDocumentFingerprint, DLPProfileCustomUpdateResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                                  `json:"enabled,required"`
	Name    string                                                `json:"name,required"`
	Type    DLPProfileCustomUpdateResponseCustomSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                         `json:"word_list"`
	JSON     dlpProfileCustomUpdateResponseCustomSharedEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponseCustomSharedEntriesUnion
}

// dlpProfileCustomUpdateResponseCustomSharedEntryJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseCustomSharedEntry]
type dlpProfileCustomUpdateResponseCustomSharedEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponseCustomSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponseCustomSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponseCustomSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseCustomSharedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntry].
func (r DLPProfileCustomUpdateResponseCustomSharedEntry) AsUnion() DLPProfileCustomUpdateResponseCustomSharedEntriesUnion {
	return r.union
}

// Union satisfied by
// [DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntry].
type DLPProfileCustomUpdateResponseCustomSharedEntriesUnion interface {
	implementsDLPProfileCustomUpdateResponseCustomSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseCustomSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                             `json:"enabled,required"`
	Name        string                                                           `json:"name,required"`
	Pattern     Pattern                                                          `json:"pattern,required"`
	Type        DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                        `json:"updated_at,required" format:"date-time"`
	Description string                                                           `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesCustomEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntry]
type dlpProfileCustomUpdateResponseCustomSharedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntry) implementsDLPProfileCustomUpdateResponseCustomSharedEntry() {
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntryType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntryTypeCustom DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntry struct {
	ID         string                                                                     `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                       `json:"enabled,required"`
	Name       string                                                                     `json:"name,required"`
	Type       DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                                  `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntry]
type dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntry) implementsDLPProfileCustomUpdateResponseCustomSharedEntry() {
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                           `json:"available,required"`
	JSON      dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidence]
type dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                           `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariant]
type dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntry struct {
	ID        string                                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                  `json:"enabled,required"`
	Name      string                                                                `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntry]
type dlpProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntry) implementsDLPProfileCustomUpdateResponseCustomSharedEntry() {
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                                `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                                `json:"enabled,required"`
	Name          string                                                              `json:"name,required"`
	Secret        bool                                                                `json:"secret,required"`
	Type          DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                           `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntry]
type dlpProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntry) implementsDLPProfileCustomUpdateResponseCustomSharedEntry() {
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryTypeExactData DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                          `json:"enabled,required"`
	Name      string                                                                        `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                     `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntry]
type dlpProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomUpdateResponseCustomSharedEntry() {
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntry struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                        `json:"word_list,required"`
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseCustomSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseCustomSharedEntriesWordListEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntry]
type dlpProfileCustomUpdateResponseCustomSharedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseCustomSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntry) implementsDLPProfileCustomUpdateResponseCustomSharedEntry() {
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntryType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntryTypeWordList DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseCustomSharedEntriesType string

const (
	DLPProfileCustomUpdateResponseCustomSharedEntriesTypeCustom              DLPProfileCustomUpdateResponseCustomSharedEntriesType = "custom"
	DLPProfileCustomUpdateResponseCustomSharedEntriesTypePredefined          DLPProfileCustomUpdateResponseCustomSharedEntriesType = "predefined"
	DLPProfileCustomUpdateResponseCustomSharedEntriesTypeIntegration         DLPProfileCustomUpdateResponseCustomSharedEntriesType = "integration"
	DLPProfileCustomUpdateResponseCustomSharedEntriesTypeExactData           DLPProfileCustomUpdateResponseCustomSharedEntriesType = "exact_data"
	DLPProfileCustomUpdateResponseCustomSharedEntriesTypeDocumentFingerprint DLPProfileCustomUpdateResponseCustomSharedEntriesType = "document_fingerprint"
	DLPProfileCustomUpdateResponseCustomSharedEntriesTypeWordList            DLPProfileCustomUpdateResponseCustomSharedEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponseCustomSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseCustomSharedEntriesTypeCustom, DLPProfileCustomUpdateResponseCustomSharedEntriesTypePredefined, DLPProfileCustomUpdateResponseCustomSharedEntriesTypeIntegration, DLPProfileCustomUpdateResponseCustomSharedEntriesTypeExactData, DLPProfileCustomUpdateResponseCustomSharedEntriesTypeDocumentFingerprint, DLPProfileCustomUpdateResponseCustomSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefined struct {
	// The id of the predefined profile (uuid).
	ID                string `json:"id,required" format:"uuid"`
	AllowedMatchCount int64  `json:"allowed_match_count,required"`
	// Deprecated: deprecated
	Entries []DLPProfileCustomUpdateResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                                                      `json:"name,required"`
	Type                DLPProfileCustomUpdateResponsePredefinedType                `json:"type,required"`
	AIContextEnabled    bool                                                        `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomUpdateResponsePredefinedConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                                         `json:"open_access"`
	JSON       dlpProfileCustomUpdateResponsePredefinedJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponsePredefined]
type dlpProfileCustomUpdateResponsePredefinedJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefined) implementsDLPProfileCustomUpdateResponse() {}

type DLPProfileCustomUpdateResponsePredefinedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                                `json:"enabled,required"`
	Name    string                                              `json:"name,required"`
	Type    DLPProfileCustomUpdateResponsePredefinedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                       `json:"word_list"`
	JSON     dlpProfileCustomUpdateResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponsePredefinedEntriesUnion
}

// dlpProfileCustomUpdateResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomUpdateResponsePredefinedEntry]
type dlpProfileCustomUpdateResponsePredefinedEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponsePredefinedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntry].
func (r DLPProfileCustomUpdateResponsePredefinedEntry) AsUnion() DLPProfileCustomUpdateResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntry].
type DLPProfileCustomUpdateResponsePredefinedEntriesUnion interface {
	implementsDLPProfileCustomUpdateResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponsePredefinedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                           `json:"enabled,required"`
	Name        string                                                         `json:"name,required"`
	Pattern     Pattern                                                        `json:"pattern,required"`
	Type        DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                      `json:"updated_at,required" format:"date-time"`
	Description string                                                         `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntry]
type dlpProfileCustomUpdateResponsePredefinedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntry) implementsDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntryType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntryTypeCustom DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntry struct {
	ID         string                                                                   `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                     `json:"enabled,required"`
	Name       string                                                                   `json:"name,required"`
	Type       DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                                `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntry]
type dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntry) implementsDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                         `json:"available,required"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidence]
type dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryTypePredefined DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                         `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariant]
type dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntry struct {
	ID        string                                                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                `json:"enabled,required"`
	Name      string                                                              `json:"name,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntry]
type dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntry) implementsDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryTypeIntegration DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                              `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                              `json:"enabled,required"`
	Name          string                                                            `json:"name,required"`
	Secret        bool                                                              `json:"secret,required"`
	Type          DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                         `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomUpdateResponsePredefinedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntry]
type dlpProfileCustomUpdateResponsePredefinedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntry) implementsDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntryType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntryTypeExactData DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                        `json:"enabled,required"`
	Name      string                                                                      `json:"name,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                   `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntry]
type dlpProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntry struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Type      DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                      `json:"word_list,required"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponsePredefinedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponsePredefinedEntriesWordListEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntry]
type dlpProfileCustomUpdateResponsePredefinedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponsePredefinedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntry) implementsDLPProfileCustomUpdateResponsePredefinedEntry() {
}

type DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntryType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntryTypeWordList DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedEntriesType string

const (
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeCustom              DLPProfileCustomUpdateResponsePredefinedEntriesType = "custom"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypePredefined          DLPProfileCustomUpdateResponsePredefinedEntriesType = "predefined"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeIntegration         DLPProfileCustomUpdateResponsePredefinedEntriesType = "integration"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeExactData           DLPProfileCustomUpdateResponsePredefinedEntriesType = "exact_data"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeDocumentFingerprint DLPProfileCustomUpdateResponsePredefinedEntriesType = "document_fingerprint"
	DLPProfileCustomUpdateResponsePredefinedEntriesTypeWordList            DLPProfileCustomUpdateResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedEntriesTypeCustom, DLPProfileCustomUpdateResponsePredefinedEntriesTypePredefined, DLPProfileCustomUpdateResponsePredefinedEntriesTypeIntegration, DLPProfileCustomUpdateResponsePredefinedEntriesTypeExactData, DLPProfileCustomUpdateResponsePredefinedEntriesTypeDocumentFingerprint, DLPProfileCustomUpdateResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedType string

const (
	DLPProfileCustomUpdateResponsePredefinedTypePredefined DLPProfileCustomUpdateResponsePredefinedType = "predefined"
)

func (r DLPProfileCustomUpdateResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponsePredefinedConfidenceThreshold string

const (
	DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdLow      DLPProfileCustomUpdateResponsePredefinedConfidenceThreshold = "low"
	DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdMedium   DLPProfileCustomUpdateResponsePredefinedConfidenceThreshold = "medium"
	DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdHigh     DLPProfileCustomUpdateResponsePredefinedConfidenceThreshold = "high"
	DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdVeryHigh DLPProfileCustomUpdateResponsePredefinedConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomUpdateResponsePredefinedConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdLow, DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdMedium, DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdHigh, DLPProfileCustomUpdateResponsePredefinedConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegration struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Entries       []DLPProfileCustomUpdateResponseIntegrationEntry       `json:"entries,required"`
	Name          string                                                 `json:"name,required"`
	SharedEntries []DLPProfileCustomUpdateResponseIntegrationSharedEntry `json:"shared_entries,required"`
	Type          DLPProfileCustomUpdateResponseIntegrationType          `json:"type,required"`
	UpdatedAt     time.Time                                              `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                                        `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponseIntegrationJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseIntegration]
type dlpProfileCustomUpdateResponseIntegrationJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegration) implementsDLPProfileCustomUpdateResponse() {}

type DLPProfileCustomUpdateResponseIntegrationEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                                 `json:"enabled,required"`
	Name    string                                               `json:"name,required"`
	Type    DLPProfileCustomUpdateResponseIntegrationEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                        `json:"word_list"`
	JSON     dlpProfileCustomUpdateResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponseIntegrationEntriesUnion
}

// dlpProfileCustomUpdateResponseIntegrationEntryJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseIntegrationEntry]
type dlpProfileCustomUpdateResponseIntegrationEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseIntegrationEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntry].
func (r DLPProfileCustomUpdateResponseIntegrationEntry) AsUnion() DLPProfileCustomUpdateResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by
// [DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntry].
type DLPProfileCustomUpdateResponseIntegrationEntriesUnion interface {
	implementsDLPProfileCustomUpdateResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseIntegrationEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                            `json:"enabled,required"`
	Name        string                                                          `json:"name,required"`
	Pattern     Pattern                                                         `json:"pattern,required"`
	Type        DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                       `json:"updated_at,required" format:"date-time"`
	Description string                                                          `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesCustomEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntry]
type dlpProfileCustomUpdateResponseIntegrationEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntry) implementsDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntryTypeCustom DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntry struct {
	ID         string                                                                    `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                      `json:"enabled,required"`
	Name       string                                                                    `json:"name,required"`
	Type       DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                                 `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntry]
type dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntry) implementsDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                          `json:"available,required"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidence]
type dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryTypePredefined DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                          `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariant]
type dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntry struct {
	ID        string                                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                 `json:"enabled,required"`
	Name      string                                                               `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                            `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntry]
type dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntry) implementsDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryTypeIntegration DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                               `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                               `json:"enabled,required"`
	Name          string                                                             `json:"name,required"`
	Secret        bool                                                               `json:"secret,required"`
	Type          DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                          `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomUpdateResponseIntegrationEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntry]
type dlpProfileCustomUpdateResponseIntegrationEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntry) implementsDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntryTypeExactData DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntry struct {
	ID        string                                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                         `json:"enabled,required"`
	Name      string                                                                       `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                    `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntry]
type dlpProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntry) implementsDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntry struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                         `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                       `json:"word_list,required"`
	ProfileID string                                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationEntriesWordListEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntry]
type dlpProfileCustomUpdateResponseIntegrationEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntry) implementsDLPProfileCustomUpdateResponseIntegrationEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntryTypeWordList DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationEntriesType string

const (
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeCustom              DLPProfileCustomUpdateResponseIntegrationEntriesType = "custom"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypePredefined          DLPProfileCustomUpdateResponseIntegrationEntriesType = "predefined"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeIntegration         DLPProfileCustomUpdateResponseIntegrationEntriesType = "integration"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeExactData           DLPProfileCustomUpdateResponseIntegrationEntriesType = "exact_data"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeDocumentFingerprint DLPProfileCustomUpdateResponseIntegrationEntriesType = "document_fingerprint"
	DLPProfileCustomUpdateResponseIntegrationEntriesTypeWordList            DLPProfileCustomUpdateResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationEntriesTypeCustom, DLPProfileCustomUpdateResponseIntegrationEntriesTypePredefined, DLPProfileCustomUpdateResponseIntegrationEntriesTypeIntegration, DLPProfileCustomUpdateResponseIntegrationEntriesTypeExactData, DLPProfileCustomUpdateResponseIntegrationEntriesTypeDocumentFingerprint, DLPProfileCustomUpdateResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                                       `json:"enabled,required"`
	Name    string                                                     `json:"name,required"`
	Type    DLPProfileCustomUpdateResponseIntegrationSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                              `json:"word_list"`
	JSON     dlpProfileCustomUpdateResponseIntegrationSharedEntryJSON `json:"-"`
	union    DLPProfileCustomUpdateResponseIntegrationSharedEntriesUnion
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseIntegrationSharedEntry]
type dlpProfileCustomUpdateResponseIntegrationSharedEntryJSON struct {
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

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomUpdateResponseIntegrationSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomUpdateResponseIntegrationSharedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntry].
func (r DLPProfileCustomUpdateResponseIntegrationSharedEntry) AsUnion() DLPProfileCustomUpdateResponseIntegrationSharedEntriesUnion {
	return r.union
}

// Union satisfied by
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntry]
// or [DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntry].
type DLPProfileCustomUpdateResponseIntegrationSharedEntriesUnion interface {
	implementsDLPProfileCustomUpdateResponseIntegrationSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomUpdateResponseIntegrationSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                                  `json:"enabled,required"`
	Name        string                                                                `json:"name,required"`
	Pattern     Pattern                                                               `json:"pattern,required"`
	Type        DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                             `json:"updated_at,required" format:"date-time"`
	Description string                                                                `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntry]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntry) implementsDLPProfileCustomUpdateResponseIntegrationSharedEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryTypeCustom DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntry struct {
	ID         string                                                                          `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                            `json:"enabled,required"`
	Name       string                                                                          `json:"name,required"`
	Type       DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                                       `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntry]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntry) implementsDLPProfileCustomUpdateResponseIntegrationSharedEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                                `json:"available,required"`
	JSON      dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidence]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                                `json:"description,nullable"`
	JSON        dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariant]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntry struct {
	ID        string                                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                       `json:"enabled,required"`
	Name      string                                                                     `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntry]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntry) implementsDLPProfileCustomUpdateResponseIntegrationSharedEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                                     `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                                `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                                     `json:"enabled,required"`
	Name          string                                                                   `json:"name,required"`
	Secret        bool                                                                     `json:"secret,required"`
	Type          DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                                `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntry]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntry) implementsDLPProfileCustomUpdateResponseIntegrationSharedEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryTypeExactData DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                               `json:"enabled,required"`
	Name      string                                                                             `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                          `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntry]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomUpdateResponseIntegrationSharedEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntry struct {
	ID        string                                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                    `json:"enabled,required"`
	Name      string                                                                  `json:"name,required"`
	Type      DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                               `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                             `json:"word_list,required"`
	ProfileID string                                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntry]
type dlpProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntry) implementsDLPProfileCustomUpdateResponseIntegrationSharedEntry() {
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryTypeWordList DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationSharedEntriesType string

const (
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeCustom              DLPProfileCustomUpdateResponseIntegrationSharedEntriesType = "custom"
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypePredefined          DLPProfileCustomUpdateResponseIntegrationSharedEntriesType = "predefined"
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeIntegration         DLPProfileCustomUpdateResponseIntegrationSharedEntriesType = "integration"
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeExactData           DLPProfileCustomUpdateResponseIntegrationSharedEntriesType = "exact_data"
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeDocumentFingerprint DLPProfileCustomUpdateResponseIntegrationSharedEntriesType = "document_fingerprint"
	DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeWordList            DLPProfileCustomUpdateResponseIntegrationSharedEntriesType = "word_list"
)

func (r DLPProfileCustomUpdateResponseIntegrationSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeCustom, DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypePredefined, DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeIntegration, DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeExactData, DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeDocumentFingerprint, DLPProfileCustomUpdateResponseIntegrationSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseIntegrationType string

const (
	DLPProfileCustomUpdateResponseIntegrationTypeIntegration DLPProfileCustomUpdateResponseIntegrationType = "integration"
)

func (r DLPProfileCustomUpdateResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseType string

const (
	DLPProfileCustomUpdateResponseTypeCustom      DLPProfileCustomUpdateResponseType = "custom"
	DLPProfileCustomUpdateResponseTypePredefined  DLPProfileCustomUpdateResponseType = "predefined"
	DLPProfileCustomUpdateResponseTypeIntegration DLPProfileCustomUpdateResponseType = "integration"
)

func (r DLPProfileCustomUpdateResponseType) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseTypeCustom, DLPProfileCustomUpdateResponseTypePredefined, DLPProfileCustomUpdateResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomUpdateResponseConfidenceThreshold string

const (
	DLPProfileCustomUpdateResponseConfidenceThresholdLow      DLPProfileCustomUpdateResponseConfidenceThreshold = "low"
	DLPProfileCustomUpdateResponseConfidenceThresholdMedium   DLPProfileCustomUpdateResponseConfidenceThreshold = "medium"
	DLPProfileCustomUpdateResponseConfidenceThresholdHigh     DLPProfileCustomUpdateResponseConfidenceThreshold = "high"
	DLPProfileCustomUpdateResponseConfidenceThresholdVeryHigh DLPProfileCustomUpdateResponseConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomUpdateResponseConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseConfidenceThresholdLow, DLPProfileCustomUpdateResponseConfidenceThresholdMedium, DLPProfileCustomUpdateResponseConfidenceThresholdHigh, DLPProfileCustomUpdateResponseConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomDeleteResponse = interface{}

type DLPProfileCustomGetResponse struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// The name of the profile.
	Name             string                          `json:"name,required"`
	Type             DLPProfileCustomGetResponseType `json:"type,required"`
	AIContextEnabled bool                            `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   int64                                          `json:"allowed_match_count"`
	ConfidenceThreshold DLPProfileCustomGetResponseConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomGetResponseCustomEntry],
	// [[]DLPProfileCustomGetResponsePredefinedEntry],
	// [[]DLPProfileCustomGetResponseIntegrationEntry].
	Entries    interface{} `json:"entries"`
	OCREnabled bool        `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool `json:"open_access"`
	// This field can have the runtime type of
	// [[]DLPProfileCustomGetResponseCustomSharedEntry],
	// [[]DLPProfileCustomGetResponseIntegrationSharedEntry].
	SharedEntries interface{} `json:"shared_entries"`
	// When the profile was lasted updated.
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      dlpProfileCustomGetResponseJSON `json:"-"`
	union     DLPProfileCustomGetResponseUnion
}

// dlpProfileCustomGetResponseJSON contains the JSON metadata for the struct
// [DLPProfileCustomGetResponse]
type dlpProfileCustomGetResponseJSON struct {
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

func (r dlpProfileCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [DLPProfileCustomGetResponseCustom],
// [DLPProfileCustomGetResponsePredefined],
// [DLPProfileCustomGetResponseIntegration].
func (r DLPProfileCustomGetResponse) AsUnion() DLPProfileCustomGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomGetResponseCustom],
// [DLPProfileCustomGetResponsePredefined] or
// [DLPProfileCustomGetResponseIntegration].
type DLPProfileCustomGetResponseUnion interface {
	implementsDLPProfileCustomGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPProfileCustomGetResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
	)
}

type DLPProfileCustomGetResponseCustom struct {
	// The id of the profile (uuid).
	ID string `json:"id,required" format:"uuid"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount int64 `json:"allowed_match_count,required"`
	// When the profile was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The name of the profile.
	Name       string                                `json:"name,required"`
	OCREnabled bool                                  `json:"ocr_enabled,required"`
	Type       DLPProfileCustomGetResponseCustomType `json:"type,required"`
	// When the profile was lasted updated.
	UpdatedAt           time.Time                                            `json:"updated_at,required" format:"date-time"`
	AIContextEnabled    bool                                                 `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomGetResponseCustomConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	// The description of the profile.
	Description string `json:"description,nullable"`
	// Deprecated: deprecated
	Entries       []DLPProfileCustomGetResponseCustomEntry       `json:"entries"`
	SharedEntries []DLPProfileCustomGetResponseCustomSharedEntry `json:"shared_entries"`
	JSON          dlpProfileCustomGetResponseCustomJSON          `json:"-"`
}

// dlpProfileCustomGetResponseCustomJSON contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseCustom]
type dlpProfileCustomGetResponseCustomJSON struct {
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

func (r *DLPProfileCustomGetResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustom) implementsDLPProfileCustomGetResponse() {}

type DLPProfileCustomGetResponseCustomType string

const (
	DLPProfileCustomGetResponseCustomTypeCustom DLPProfileCustomGetResponseCustomType = "custom"
)

func (r DLPProfileCustomGetResponseCustomType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomConfidenceThreshold string

const (
	DLPProfileCustomGetResponseCustomConfidenceThresholdLow      DLPProfileCustomGetResponseCustomConfidenceThreshold = "low"
	DLPProfileCustomGetResponseCustomConfidenceThresholdMedium   DLPProfileCustomGetResponseCustomConfidenceThreshold = "medium"
	DLPProfileCustomGetResponseCustomConfidenceThresholdHigh     DLPProfileCustomGetResponseCustomConfidenceThreshold = "high"
	DLPProfileCustomGetResponseCustomConfidenceThresholdVeryHigh DLPProfileCustomGetResponseCustomConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomGetResponseCustomConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomConfidenceThresholdLow, DLPProfileCustomGetResponseCustomConfidenceThresholdMedium, DLPProfileCustomGetResponseCustomConfidenceThresholdHigh, DLPProfileCustomGetResponseCustomConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                         `json:"enabled,required"`
	Name    string                                       `json:"name,required"`
	Type    DLPProfileCustomGetResponseCustomEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseCustomEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                `json:"word_list"`
	JSON     dlpProfileCustomGetResponseCustomEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponseCustomEntriesUnion
}

// dlpProfileCustomGetResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseCustomEntry]
type dlpProfileCustomGetResponseCustomEntryJSON struct {
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

func (r dlpProfileCustomGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponseCustomEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseCustomEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomGetResponseCustomEntriesCustomEntry],
// [DLPProfileCustomGetResponseCustomEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseCustomEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseCustomEntriesExactDataEntry],
// [DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntry],
// [DLPProfileCustomGetResponseCustomEntriesWordListEntry].
func (r DLPProfileCustomGetResponseCustomEntry) AsUnion() DLPProfileCustomGetResponseCustomEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomGetResponseCustomEntriesCustomEntry],
// [DLPProfileCustomGetResponseCustomEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseCustomEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseCustomEntriesExactDataEntry],
// [DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomGetResponseCustomEntriesWordListEntry].
type DLPProfileCustomGetResponseCustomEntriesUnion interface {
	implementsDLPProfileCustomGetResponseCustomEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseCustomEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomGetResponseCustomEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                    `json:"enabled,required"`
	Name        string                                                  `json:"name,required"`
	Pattern     Pattern                                                 `json:"pattern,required"`
	Type        DLPProfileCustomGetResponseCustomEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                               `json:"updated_at,required" format:"date-time"`
	Description string                                                  `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesCustomEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseCustomEntriesCustomEntry]
type dlpProfileCustomGetResponseCustomEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesCustomEntry) implementsDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesCustomEntryType string

const (
	DLPProfileCustomGetResponseCustomEntriesCustomEntryTypeCustom DLPProfileCustomGetResponseCustomEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomGetResponseCustomEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesPredefinedEntry struct {
	ID         string                                                            `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomGetResponseCustomEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                              `json:"enabled,required"`
	Name       string                                                            `json:"name,required"`
	Type       DLPProfileCustomGetResponseCustomEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomGetResponseCustomEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseCustomEntriesPredefinedEntry]
type dlpProfileCustomGetResponseCustomEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesPredefinedEntry) implementsDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                  `json:"available,required"`
	JSON      dlpProfileCustomGetResponseCustomEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesPredefinedEntryConfidenceJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomEntriesPredefinedEntryConfidence]
type dlpProfileCustomGetResponseCustomEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseCustomEntriesPredefinedEntryType string

const (
	DLPProfileCustomGetResponseCustomEntriesPredefinedEntryTypePredefined DLPProfileCustomGetResponseCustomEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomGetResponseCustomEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                  `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponseCustomEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariant]
type dlpProfileCustomGetResponseCustomEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesIntegrationEntry struct {
	ID        string                                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                         `json:"enabled,required"`
	Name      string                                                       `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseCustomEntriesIntegrationEntry]
type dlpProfileCustomGetResponseCustomEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesIntegrationEntry) implementsDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesIntegrationEntryType string

const (
	DLPProfileCustomGetResponseCustomEntriesIntegrationEntryTypeIntegration DLPProfileCustomGetResponseCustomEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomGetResponseCustomEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                       `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                       `json:"enabled,required"`
	Name          string                                                     `json:"name,required"`
	Secret        bool                                                       `json:"secret,required"`
	Type          DLPProfileCustomGetResponseCustomEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                  `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomGetResponseCustomEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesExactDataEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseCustomEntriesExactDataEntry]
type dlpProfileCustomGetResponseCustomEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesExactDataEntry) implementsDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesExactDataEntryType string

const (
	DLPProfileCustomGetResponseCustomEntriesExactDataEntryTypeExactData DLPProfileCustomGetResponseCustomEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomGetResponseCustomEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntry struct {
	ID        string                                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                 `json:"enabled,required"`
	Name      string                                                               `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                            `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntry]
type dlpProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntry) implementsDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesWordListEntry struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                               `json:"word_list,required"`
	ProfileID string                                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomEntriesWordListEntryJSON contains the JSON
// metadata for the struct [DLPProfileCustomGetResponseCustomEntriesWordListEntry]
type dlpProfileCustomGetResponseCustomEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomEntriesWordListEntry) implementsDLPProfileCustomGetResponseCustomEntry() {
}

type DLPProfileCustomGetResponseCustomEntriesWordListEntryType string

const (
	DLPProfileCustomGetResponseCustomEntriesWordListEntryTypeWordList DLPProfileCustomGetResponseCustomEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomGetResponseCustomEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomEntriesType string

const (
	DLPProfileCustomGetResponseCustomEntriesTypeCustom              DLPProfileCustomGetResponseCustomEntriesType = "custom"
	DLPProfileCustomGetResponseCustomEntriesTypePredefined          DLPProfileCustomGetResponseCustomEntriesType = "predefined"
	DLPProfileCustomGetResponseCustomEntriesTypeIntegration         DLPProfileCustomGetResponseCustomEntriesType = "integration"
	DLPProfileCustomGetResponseCustomEntriesTypeExactData           DLPProfileCustomGetResponseCustomEntriesType = "exact_data"
	DLPProfileCustomGetResponseCustomEntriesTypeDocumentFingerprint DLPProfileCustomGetResponseCustomEntriesType = "document_fingerprint"
	DLPProfileCustomGetResponseCustomEntriesTypeWordList            DLPProfileCustomGetResponseCustomEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponseCustomEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomEntriesTypeCustom, DLPProfileCustomGetResponseCustomEntriesTypePredefined, DLPProfileCustomGetResponseCustomEntriesTypeIntegration, DLPProfileCustomGetResponseCustomEntriesTypeExactData, DLPProfileCustomGetResponseCustomEntriesTypeDocumentFingerprint, DLPProfileCustomGetResponseCustomEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                               `json:"enabled,required"`
	Name    string                                             `json:"name,required"`
	Type    DLPProfileCustomGetResponseCustomSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                      `json:"word_list"`
	JSON     dlpProfileCustomGetResponseCustomSharedEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponseCustomSharedEntriesUnion
}

// dlpProfileCustomGetResponseCustomSharedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseCustomSharedEntry]
type dlpProfileCustomGetResponseCustomSharedEntryJSON struct {
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

func (r dlpProfileCustomGetResponseCustomSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponseCustomSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponseCustomSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseCustomSharedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomGetResponseCustomSharedEntriesCustomEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesWordListEntry].
func (r DLPProfileCustomGetResponseCustomSharedEntry) AsUnion() DLPProfileCustomGetResponseCustomSharedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomGetResponseCustomSharedEntriesCustomEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntry],
// [DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomGetResponseCustomSharedEntriesWordListEntry].
type DLPProfileCustomGetResponseCustomSharedEntriesUnion interface {
	implementsDLPProfileCustomGetResponseCustomSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseCustomSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseCustomSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomGetResponseCustomSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                          `json:"enabled,required"`
	Name        string                                                        `json:"name,required"`
	Pattern     Pattern                                                       `json:"pattern,required"`
	Type        DLPProfileCustomGetResponseCustomSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                     `json:"updated_at,required" format:"date-time"`
	Description string                                                        `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesCustomEntry]
type dlpProfileCustomGetResponseCustomSharedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomSharedEntriesCustomEntry) implementsDLPProfileCustomGetResponseCustomSharedEntry() {
}

type DLPProfileCustomGetResponseCustomSharedEntriesCustomEntryType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesCustomEntryTypeCustom DLPProfileCustomGetResponseCustomSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntry struct {
	ID         string                                                                  `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                    `json:"enabled,required"`
	Name       string                                                                  `json:"name,required"`
	Type       DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                               `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntry]
type dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntry) implementsDLPProfileCustomGetResponseCustomSharedEntry() {
}

type DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                        `json:"available,required"`
	JSON      dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidence]
type dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                        `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariant]
type dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntry struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntry]
type dlpProfileCustomGetResponseCustomSharedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntry) implementsDLPProfileCustomGetResponseCustomSharedEntry() {
}

type DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                             `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                             `json:"enabled,required"`
	Name          string                                                           `json:"name,required"`
	Secret        bool                                                             `json:"secret,required"`
	Type          DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                        `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomGetResponseCustomSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntry]
type dlpProfileCustomGetResponseCustomSharedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntry) implementsDLPProfileCustomGetResponseCustomSharedEntry() {
}

type DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntryTypeExactData DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                       `json:"enabled,required"`
	Name      string                                                                     `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                  `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntry]
type dlpProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomGetResponseCustomSharedEntry() {
}

type DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesWordListEntry struct {
	ID        string                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                            `json:"enabled,required"`
	Name      string                                                          `json:"name,required"`
	Type      DLPProfileCustomGetResponseCustomSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                       `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                     `json:"word_list,required"`
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseCustomSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseCustomSharedEntriesWordListEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseCustomSharedEntriesWordListEntry]
type dlpProfileCustomGetResponseCustomSharedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseCustomSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseCustomSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseCustomSharedEntriesWordListEntry) implementsDLPProfileCustomGetResponseCustomSharedEntry() {
}

type DLPProfileCustomGetResponseCustomSharedEntriesWordListEntryType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesWordListEntryTypeWordList DLPProfileCustomGetResponseCustomSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseCustomSharedEntriesType string

const (
	DLPProfileCustomGetResponseCustomSharedEntriesTypeCustom              DLPProfileCustomGetResponseCustomSharedEntriesType = "custom"
	DLPProfileCustomGetResponseCustomSharedEntriesTypePredefined          DLPProfileCustomGetResponseCustomSharedEntriesType = "predefined"
	DLPProfileCustomGetResponseCustomSharedEntriesTypeIntegration         DLPProfileCustomGetResponseCustomSharedEntriesType = "integration"
	DLPProfileCustomGetResponseCustomSharedEntriesTypeExactData           DLPProfileCustomGetResponseCustomSharedEntriesType = "exact_data"
	DLPProfileCustomGetResponseCustomSharedEntriesTypeDocumentFingerprint DLPProfileCustomGetResponseCustomSharedEntriesType = "document_fingerprint"
	DLPProfileCustomGetResponseCustomSharedEntriesTypeWordList            DLPProfileCustomGetResponseCustomSharedEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponseCustomSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseCustomSharedEntriesTypeCustom, DLPProfileCustomGetResponseCustomSharedEntriesTypePredefined, DLPProfileCustomGetResponseCustomSharedEntriesTypeIntegration, DLPProfileCustomGetResponseCustomSharedEntriesTypeExactData, DLPProfileCustomGetResponseCustomSharedEntriesTypeDocumentFingerprint, DLPProfileCustomGetResponseCustomSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefined struct {
	// The id of the predefined profile (uuid).
	ID                string `json:"id,required" format:"uuid"`
	AllowedMatchCount int64  `json:"allowed_match_count,required"`
	// Deprecated: deprecated
	Entries []DLPProfileCustomGetResponsePredefinedEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name                string                                                   `json:"name,required"`
	Type                DLPProfileCustomGetResponsePredefinedType                `json:"type,required"`
	AIContextEnabled    bool                                                     `json:"ai_context_enabled"`
	ConfidenceThreshold DLPProfileCustomGetResponsePredefinedConfidenceThreshold `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	//
	// Deprecated: deprecated
	ContextAwareness ContextAwareness `json:"context_awareness"`
	OCREnabled       bool             `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                                      `json:"open_access"`
	JSON       dlpProfileCustomGetResponsePredefinedJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponsePredefined]
type dlpProfileCustomGetResponsePredefinedJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefined) implementsDLPProfileCustomGetResponse() {}

type DLPProfileCustomGetResponsePredefinedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                             `json:"enabled,required"`
	Name    string                                           `json:"name,required"`
	Type    DLPProfileCustomGetResponsePredefinedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                    `json:"word_list"`
	JSON     dlpProfileCustomGetResponsePredefinedEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponsePredefinedEntriesUnion
}

// dlpProfileCustomGetResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponsePredefinedEntry]
type dlpProfileCustomGetResponsePredefinedEntryJSON struct {
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

func (r dlpProfileCustomGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponsePredefinedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponsePredefinedEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomGetResponsePredefinedEntriesCustomEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesExactDataEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesWordListEntry].
func (r DLPProfileCustomGetResponsePredefinedEntry) AsUnion() DLPProfileCustomGetResponsePredefinedEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomGetResponsePredefinedEntriesCustomEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesExactDataEntry],
// [DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomGetResponsePredefinedEntriesWordListEntry].
type DLPProfileCustomGetResponsePredefinedEntriesUnion interface {
	implementsDLPProfileCustomGetResponsePredefinedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponsePredefinedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponsePredefinedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomGetResponsePredefinedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                        `json:"enabled,required"`
	Name        string                                                      `json:"name,required"`
	Pattern     Pattern                                                     `json:"pattern,required"`
	Type        DLPProfileCustomGetResponsePredefinedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                   `json:"updated_at,required" format:"date-time"`
	Description string                                                      `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesCustomEntry]
type dlpProfileCustomGetResponsePredefinedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesCustomEntry) implementsDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesCustomEntryType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesCustomEntryTypeCustom DLPProfileCustomGetResponsePredefinedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntry struct {
	ID         string                                                                `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                  `json:"enabled,required"`
	Name       string                                                                `json:"name,required"`
	Type       DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntry]
type dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntry) implementsDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                      `json:"available,required"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidence]
type dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryTypePredefined DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                      `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariant]
type dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntry struct {
	ID        string                                                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                             `json:"enabled,required"`
	Name      string                                                           `json:"name,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                        `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntry]
type dlpProfileCustomGetResponsePredefinedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntry) implementsDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntryType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntryTypeIntegration DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                           `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                           `json:"enabled,required"`
	Name          string                                                         `json:"name,required"`
	Secret        bool                                                           `json:"secret,required"`
	Type          DLPProfileCustomGetResponsePredefinedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                      `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomGetResponsePredefinedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesExactDataEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesExactDataEntry]
type dlpProfileCustomGetResponsePredefinedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesExactDataEntry) implementsDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesExactDataEntryType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesExactDataEntryTypeExactData DLPProfileCustomGetResponsePredefinedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                     `json:"enabled,required"`
	Name      string                                                                   `json:"name,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntry]
type dlpProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesWordListEntry struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Type      DLPProfileCustomGetResponsePredefinedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                   `json:"word_list,required"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponsePredefinedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponsePredefinedEntriesWordListEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponsePredefinedEntriesWordListEntry]
type dlpProfileCustomGetResponsePredefinedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomGetResponsePredefinedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponsePredefinedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponsePredefinedEntriesWordListEntry) implementsDLPProfileCustomGetResponsePredefinedEntry() {
}

type DLPProfileCustomGetResponsePredefinedEntriesWordListEntryType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesWordListEntryTypeWordList DLPProfileCustomGetResponsePredefinedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedEntriesType string

const (
	DLPProfileCustomGetResponsePredefinedEntriesTypeCustom              DLPProfileCustomGetResponsePredefinedEntriesType = "custom"
	DLPProfileCustomGetResponsePredefinedEntriesTypePredefined          DLPProfileCustomGetResponsePredefinedEntriesType = "predefined"
	DLPProfileCustomGetResponsePredefinedEntriesTypeIntegration         DLPProfileCustomGetResponsePredefinedEntriesType = "integration"
	DLPProfileCustomGetResponsePredefinedEntriesTypeExactData           DLPProfileCustomGetResponsePredefinedEntriesType = "exact_data"
	DLPProfileCustomGetResponsePredefinedEntriesTypeDocumentFingerprint DLPProfileCustomGetResponsePredefinedEntriesType = "document_fingerprint"
	DLPProfileCustomGetResponsePredefinedEntriesTypeWordList            DLPProfileCustomGetResponsePredefinedEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponsePredefinedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedEntriesTypeCustom, DLPProfileCustomGetResponsePredefinedEntriesTypePredefined, DLPProfileCustomGetResponsePredefinedEntriesTypeIntegration, DLPProfileCustomGetResponsePredefinedEntriesTypeExactData, DLPProfileCustomGetResponsePredefinedEntriesTypeDocumentFingerprint, DLPProfileCustomGetResponsePredefinedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedType string

const (
	DLPProfileCustomGetResponsePredefinedTypePredefined DLPProfileCustomGetResponsePredefinedType = "predefined"
)

func (r DLPProfileCustomGetResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponsePredefinedConfidenceThreshold string

const (
	DLPProfileCustomGetResponsePredefinedConfidenceThresholdLow      DLPProfileCustomGetResponsePredefinedConfidenceThreshold = "low"
	DLPProfileCustomGetResponsePredefinedConfidenceThresholdMedium   DLPProfileCustomGetResponsePredefinedConfidenceThreshold = "medium"
	DLPProfileCustomGetResponsePredefinedConfidenceThresholdHigh     DLPProfileCustomGetResponsePredefinedConfidenceThreshold = "high"
	DLPProfileCustomGetResponsePredefinedConfidenceThresholdVeryHigh DLPProfileCustomGetResponsePredefinedConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomGetResponsePredefinedConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponsePredefinedConfidenceThresholdLow, DLPProfileCustomGetResponsePredefinedConfidenceThresholdMedium, DLPProfileCustomGetResponsePredefinedConfidenceThresholdHigh, DLPProfileCustomGetResponsePredefinedConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegration struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Entries       []DLPProfileCustomGetResponseIntegrationEntry       `json:"entries,required"`
	Name          string                                              `json:"name,required"`
	SharedEntries []DLPProfileCustomGetResponseIntegrationSharedEntry `json:"shared_entries,required"`
	Type          DLPProfileCustomGetResponseIntegrationType          `json:"type,required"`
	UpdatedAt     time.Time                                           `json:"updated_at,required" format:"date-time"`
	// The description of the profile.
	Description string                                     `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponseIntegrationJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseIntegration]
type dlpProfileCustomGetResponseIntegrationJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegration) implementsDLPProfileCustomGetResponse() {}

type DLPProfileCustomGetResponseIntegrationEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                              `json:"enabled,required"`
	Name    string                                            `json:"name,required"`
	Type    DLPProfileCustomGetResponseIntegrationEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                     `json:"word_list"`
	JSON     dlpProfileCustomGetResponseIntegrationEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponseIntegrationEntriesUnion
}

// dlpProfileCustomGetResponseIntegrationEntryJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseIntegrationEntry]
type dlpProfileCustomGetResponseIntegrationEntryJSON struct {
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

func (r dlpProfileCustomGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponseIntegrationEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseIntegrationEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomGetResponseIntegrationEntriesCustomEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesExactDataEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesWordListEntry].
func (r DLPProfileCustomGetResponseIntegrationEntry) AsUnion() DLPProfileCustomGetResponseIntegrationEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfileCustomGetResponseIntegrationEntriesCustomEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesExactDataEntry],
// [DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomGetResponseIntegrationEntriesWordListEntry].
type DLPProfileCustomGetResponseIntegrationEntriesUnion interface {
	implementsDLPProfileCustomGetResponseIntegrationEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseIntegrationEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomGetResponseIntegrationEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                         `json:"enabled,required"`
	Name        string                                                       `json:"name,required"`
	Pattern     Pattern                                                      `json:"pattern,required"`
	Type        DLPProfileCustomGetResponseIntegrationEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                    `json:"updated_at,required" format:"date-time"`
	Description string                                                       `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesCustomEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesCustomEntry]
type dlpProfileCustomGetResponseIntegrationEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesCustomEntry) implementsDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesCustomEntryType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesCustomEntryTypeCustom DLPProfileCustomGetResponseIntegrationEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntry struct {
	ID         string                                                                 `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                   `json:"enabled,required"`
	Name       string                                                                 `json:"name,required"`
	Type       DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                              `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntry]
type dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntry) implementsDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                       `json:"available,required"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidence]
type dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryTypePredefined DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                       `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariant]
type dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntry struct {
	ID        string                                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                              `json:"enabled,required"`
	Name      string                                                            `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                         `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesIntegrationEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntry]
type dlpProfileCustomGetResponseIntegrationEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntry) implementsDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntryType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntryTypeIntegration DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                            `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                       `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                            `json:"enabled,required"`
	Name          string                                                          `json:"name,required"`
	Secret        bool                                                            `json:"secret,required"`
	Type          DLPProfileCustomGetResponseIntegrationEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                       `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomGetResponseIntegrationEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesExactDataEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesExactDataEntry]
type dlpProfileCustomGetResponseIntegrationEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesExactDataEntry) implementsDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesExactDataEntryType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesExactDataEntryTypeExactData DLPProfileCustomGetResponseIntegrationEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntry struct {
	ID        string                                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                      `json:"enabled,required"`
	Name      string                                                                    `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                 `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntry]
type dlpProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntry) implementsDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesWordListEntry struct {
	ID        string                                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                           `json:"enabled,required"`
	Name      string                                                         `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                      `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                    `json:"word_list,required"`
	ProfileID string                                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationEntriesWordListEntryJSON contains the JSON
// metadata for the struct
// [DLPProfileCustomGetResponseIntegrationEntriesWordListEntry]
type dlpProfileCustomGetResponseIntegrationEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationEntriesWordListEntry) implementsDLPProfileCustomGetResponseIntegrationEntry() {
}

type DLPProfileCustomGetResponseIntegrationEntriesWordListEntryType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesWordListEntryTypeWordList DLPProfileCustomGetResponseIntegrationEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationEntriesType string

const (
	DLPProfileCustomGetResponseIntegrationEntriesTypeCustom              DLPProfileCustomGetResponseIntegrationEntriesType = "custom"
	DLPProfileCustomGetResponseIntegrationEntriesTypePredefined          DLPProfileCustomGetResponseIntegrationEntriesType = "predefined"
	DLPProfileCustomGetResponseIntegrationEntriesTypeIntegration         DLPProfileCustomGetResponseIntegrationEntriesType = "integration"
	DLPProfileCustomGetResponseIntegrationEntriesTypeExactData           DLPProfileCustomGetResponseIntegrationEntriesType = "exact_data"
	DLPProfileCustomGetResponseIntegrationEntriesTypeDocumentFingerprint DLPProfileCustomGetResponseIntegrationEntriesType = "document_fingerprint"
	DLPProfileCustomGetResponseIntegrationEntriesTypeWordList            DLPProfileCustomGetResponseIntegrationEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponseIntegrationEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationEntriesTypeCustom, DLPProfileCustomGetResponseIntegrationEntriesTypePredefined, DLPProfileCustomGetResponseIntegrationEntriesTypeIntegration, DLPProfileCustomGetResponseIntegrationEntriesTypeExactData, DLPProfileCustomGetResponseIntegrationEntriesTypeDocumentFingerprint, DLPProfileCustomGetResponseIntegrationEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                                    `json:"enabled,required"`
	Name    string                                                  `json:"name,required"`
	Type    DLPProfileCustomGetResponseIntegrationSharedEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                           `json:"word_list"`
	JSON     dlpProfileCustomGetResponseIntegrationSharedEntryJSON `json:"-"`
	union    DLPProfileCustomGetResponseIntegrationSharedEntriesUnion
}

// dlpProfileCustomGetResponseIntegrationSharedEntryJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseIntegrationSharedEntry]
type dlpProfileCustomGetResponseIntegrationSharedEntryJSON struct {
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

func (r dlpProfileCustomGetResponseIntegrationSharedEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfileCustomGetResponseIntegrationSharedEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfileCustomGetResponseIntegrationSharedEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfileCustomGetResponseIntegrationSharedEntriesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntry].
func (r DLPProfileCustomGetResponseIntegrationSharedEntry) AsUnion() DLPProfileCustomGetResponseIntegrationSharedEntriesUnion {
	return r.union
}

// Union satisfied by
// [DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntry],
// [DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntry] or
// [DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntry].
type DLPProfileCustomGetResponseIntegrationSharedEntriesUnion interface {
	implementsDLPProfileCustomGetResponseIntegrationSharedEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfileCustomGetResponseIntegrationSharedEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntry{}),
		},
	)
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                               `json:"enabled,required"`
	Name        string                                                             `json:"name,required"`
	Pattern     Pattern                                                            `json:"pattern,required"`
	Type        DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                          `json:"updated_at,required" format:"date-time"`
	Description string                                                             `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationSharedEntriesCustomEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesCustomEntryJSON contains the
// JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntry]
type dlpProfileCustomGetResponseIntegrationSharedEntriesCustomEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntry) implementsDLPProfileCustomGetResponseIntegrationSharedEntry() {
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntryType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntryTypeCustom DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntryType = "custom"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntry struct {
	ID         string                                                                       `json:"id,required" format:"uuid"`
	Confidence DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                                         `json:"enabled,required"`
	Name       string                                                                       `json:"name,required"`
	Type       DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                                    `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntry]
type dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntry) implementsDLPProfileCustomGetResponseIntegrationSharedEntry() {
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                             `json:"available,required"`
	JSON      dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidence]
type dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryTypePredefined DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                             `json:"description,nullable"`
	JSON        dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariant]
type dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntry struct {
	ID        string                                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                    `json:"enabled,required"`
	Name      string                                                                  `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntry]
type dlpProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntry) implementsDLPProfileCustomGetResponseIntegrationSharedEntry() {
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryType = "integration"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                                  `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                             `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                                  `json:"enabled,required"`
	Name          string                                                                `json:"name,required"`
	Secret        bool                                                                  `json:"secret,required"`
	Type          DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                             `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntry]
type dlpProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntry) implementsDLPProfileCustomGetResponseIntegrationSharedEntry() {
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryTypeExactData DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntry struct {
	ID        string                                                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                            `json:"enabled,required"`
	Name      string                                                                          `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                                       `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON
// contains the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntry]
type dlpProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntry) implementsDLPProfileCustomGetResponseIntegrationSharedEntry() {
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntry struct {
	ID        string                                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                 `json:"enabled,required"`
	Name      string                                                               `json:"name,required"`
	Type      DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                            `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                          `json:"word_list,required"`
	ProfileID string                                                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfileCustomGetResponseIntegrationSharedEntriesWordListEntryJSON `json:"-"`
}

// dlpProfileCustomGetResponseIntegrationSharedEntriesWordListEntryJSON contains
// the JSON metadata for the struct
// [DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntry]
type dlpProfileCustomGetResponseIntegrationSharedEntriesWordListEntryJSON struct {
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

func (r *DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseIntegrationSharedEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntry) implementsDLPProfileCustomGetResponseIntegrationSharedEntry() {
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntryType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntryTypeWordList DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntryType = "word_list"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationSharedEntriesType string

const (
	DLPProfileCustomGetResponseIntegrationSharedEntriesTypeCustom              DLPProfileCustomGetResponseIntegrationSharedEntriesType = "custom"
	DLPProfileCustomGetResponseIntegrationSharedEntriesTypePredefined          DLPProfileCustomGetResponseIntegrationSharedEntriesType = "predefined"
	DLPProfileCustomGetResponseIntegrationSharedEntriesTypeIntegration         DLPProfileCustomGetResponseIntegrationSharedEntriesType = "integration"
	DLPProfileCustomGetResponseIntegrationSharedEntriesTypeExactData           DLPProfileCustomGetResponseIntegrationSharedEntriesType = "exact_data"
	DLPProfileCustomGetResponseIntegrationSharedEntriesTypeDocumentFingerprint DLPProfileCustomGetResponseIntegrationSharedEntriesType = "document_fingerprint"
	DLPProfileCustomGetResponseIntegrationSharedEntriesTypeWordList            DLPProfileCustomGetResponseIntegrationSharedEntriesType = "word_list"
)

func (r DLPProfileCustomGetResponseIntegrationSharedEntriesType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationSharedEntriesTypeCustom, DLPProfileCustomGetResponseIntegrationSharedEntriesTypePredefined, DLPProfileCustomGetResponseIntegrationSharedEntriesTypeIntegration, DLPProfileCustomGetResponseIntegrationSharedEntriesTypeExactData, DLPProfileCustomGetResponseIntegrationSharedEntriesTypeDocumentFingerprint, DLPProfileCustomGetResponseIntegrationSharedEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseIntegrationType string

const (
	DLPProfileCustomGetResponseIntegrationTypeIntegration DLPProfileCustomGetResponseIntegrationType = "integration"
)

func (r DLPProfileCustomGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseType string

const (
	DLPProfileCustomGetResponseTypeCustom      DLPProfileCustomGetResponseType = "custom"
	DLPProfileCustomGetResponseTypePredefined  DLPProfileCustomGetResponseType = "predefined"
	DLPProfileCustomGetResponseTypeIntegration DLPProfileCustomGetResponseType = "integration"
)

func (r DLPProfileCustomGetResponseType) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseTypeCustom, DLPProfileCustomGetResponseTypePredefined, DLPProfileCustomGetResponseTypeIntegration:
		return true
	}
	return false
}

type DLPProfileCustomGetResponseConfidenceThreshold string

const (
	DLPProfileCustomGetResponseConfidenceThresholdLow      DLPProfileCustomGetResponseConfidenceThreshold = "low"
	DLPProfileCustomGetResponseConfidenceThresholdMedium   DLPProfileCustomGetResponseConfidenceThreshold = "medium"
	DLPProfileCustomGetResponseConfidenceThresholdHigh     DLPProfileCustomGetResponseConfidenceThreshold = "high"
	DLPProfileCustomGetResponseConfidenceThresholdVeryHigh DLPProfileCustomGetResponseConfidenceThreshold = "very_high"
)

func (r DLPProfileCustomGetResponseConfidenceThreshold) IsKnown() bool {
	switch r {
	case DLPProfileCustomGetResponseConfidenceThresholdLow, DLPProfileCustomGetResponseConfidenceThresholdMedium, DLPProfileCustomGetResponseConfidenceThresholdHigh, DLPProfileCustomGetResponseConfidenceThresholdVeryHigh:
		return true
	}
	return false
}

type DLPProfileCustomNewParams struct {
	AccountID        param.Field[string] `path:"account_id,required"`
	Name             param.Field[string] `json:"name,required"`
	AIContextEnabled param.Field[bool]   `json:"ai_context_enabled"`
	// Related DLP policies will trigger when the match count exceeds the number set.
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string]                                `json:"description"`
	Entries     param.Field[[]DLPProfileCustomNewParamsEntryUnion] `json:"entries"`
	OCREnabled  param.Field[bool]                                  `json:"ocr_enabled"`
	// Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your
	// Microsoft Information Protection profiles).
	SharedEntries param.Field[[]DLPProfileCustomNewParamsSharedEntry] `json:"shared_entries"`
}

func (r DLPProfileCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewParamsEntry struct {
	Enabled     param.Field[bool]         `json:"enabled,required"`
	Name        param.Field[string]       `json:"name,required"`
	Description param.Field[string]       `json:"description"`
	Pattern     param.Field[PatternParam] `json:"pattern"`
	Words       param.Field[interface{}]  `json:"words"`
}

func (r DLPProfileCustomNewParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsEntry) implementsDLPProfileCustomNewParamsEntryUnion() {}

// Satisfied by [zero_trust.DLPProfileCustomNewParamsEntriesDLPNewCustomEntry],
// [zero_trust.DLPProfileCustomNewParamsEntriesDLPNewWordListEntry],
// [DLPProfileCustomNewParamsEntry].
type DLPProfileCustomNewParamsEntryUnion interface {
	implementsDLPProfileCustomNewParamsEntryUnion()
}

type DLPProfileCustomNewParamsEntriesDLPNewCustomEntry struct {
	Enabled     param.Field[bool]         `json:"enabled,required"`
	Name        param.Field[string]       `json:"name,required"`
	Pattern     param.Field[PatternParam] `json:"pattern,required"`
	Description param.Field[string]       `json:"description"`
}

func (r DLPProfileCustomNewParamsEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsEntriesDLPNewCustomEntry) implementsDLPProfileCustomNewParamsEntryUnion() {
}

type DLPProfileCustomNewParamsEntriesDLPNewWordListEntry struct {
	Enabled param.Field[bool]     `json:"enabled,required"`
	Name    param.Field[string]   `json:"name,required"`
	Words   param.Field[[]string] `json:"words,required"`
}

func (r DLPProfileCustomNewParamsEntriesDLPNewWordListEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomNewParamsEntriesDLPNewWordListEntry) implementsDLPProfileCustomNewParamsEntryUnion() {
}

type DLPProfileCustomNewParamsSharedEntry struct {
	Enabled param.Field[bool]   `json:"enabled,required"`
	EntryID param.Field[string] `json:"entry_id,required" format:"uuid"`
}

func (r DLPProfileCustomNewParamsSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomNewResponseEnvelope struct {
	Errors   []DLPProfileCustomNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomNewResponse                `json:"result"`
	JSON    dlpProfileCustomNewResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelope]
type dlpProfileCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPProfileCustomNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomNewResponseEnvelopeErrors]
type dlpProfileCustomNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseEnvelopeErrorsSource]
type dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPProfileCustomNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomNewResponseEnvelopeMessages]
type dlpProfileCustomNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomNewResponseEnvelopeMessagesSource]
type dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

type DLPProfileCustomUpdateParams struct {
	AccountID           param.Field[string] `path:"account_id,required"`
	Name                param.Field[string] `json:"name,required"`
	AIContextEnabled    param.Field[bool]   `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]  `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string] `json:"confidence_threshold"`
	// Scan the context of predefined entries to only return matches surrounded by
	// keywords.
	ContextAwareness param.Field[ContextAwarenessParam] `json:"context_awareness"`
	// The description of the profile.
	Description param.Field[string] `json:"description"`
	// Custom entries from this profile. If this field is omitted, entries owned by
	// this profile will not be changed.
	Entries    param.Field[[]DLPProfileCustomUpdateParamsEntryUnion] `json:"entries"`
	OCREnabled param.Field[bool]                                     `json:"ocr_enabled"`
	// Other entries, e.g. predefined or integration.
	SharedEntries param.Field[[]DLPProfileCustomUpdateParamsSharedEntry] `json:"shared_entries"`
}

func (r DLPProfileCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomUpdateParamsEntry struct {
	Enabled     param.Field[bool]         `json:"enabled,required"`
	Name        param.Field[string]       `json:"name,required"`
	Pattern     param.Field[PatternParam] `json:"pattern,required"`
	Description param.Field[string]       `json:"description"`
	EntryID     param.Field[string]       `json:"entry_id" format:"uuid"`
}

func (r DLPProfileCustomUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntry) implementsDLPProfileCustomUpdateParamsEntryUnion() {}

// Satisfied by
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID],
// [zero_trust.DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry],
// [DLPProfileCustomUpdateParamsEntry].
type DLPProfileCustomUpdateParamsEntryUnion interface {
	implementsDLPProfileCustomUpdateParamsEntryUnion()
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID struct {
	Enabled     param.Field[bool]         `json:"enabled,required"`
	EntryID     param.Field[string]       `json:"entry_id,required" format:"uuid"`
	Name        param.Field[string]       `json:"name,required"`
	Pattern     param.Field[PatternParam] `json:"pattern,required"`
	Description param.Field[string]       `json:"description"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntryWithID) implementsDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry struct {
	Enabled     param.Field[bool]         `json:"enabled,required"`
	Name        param.Field[string]       `json:"name,required"`
	Pattern     param.Field[PatternParam] `json:"pattern,required"`
	Description param.Field[string]       `json:"description"`
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPProfileCustomUpdateParamsEntriesDLPNewCustomEntry) implementsDLPProfileCustomUpdateParamsEntryUnion() {
}

type DLPProfileCustomUpdateParamsSharedEntry struct {
	Enabled param.Field[bool]   `json:"enabled,required"`
	EntryID param.Field[string] `json:"entry_id,required" format:"uuid"`
}

func (r DLPProfileCustomUpdateParamsSharedEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfileCustomUpdateResponseEnvelope struct {
	Errors   []DLPProfileCustomUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomUpdateResponse                `json:"result"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomUpdateResponseEnvelope]
type dlpProfileCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPProfileCustomUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfileCustomUpdateResponseEnvelopeErrors]
type dlpProfileCustomUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseEnvelopeErrorsSource]
type dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPProfileCustomUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfileCustomUpdateResponseEnvelopeMessages]
type dlpProfileCustomUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomUpdateResponseEnvelopeMessagesSource]
type dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfileCustomUpdateResponseEnvelopeSuccess bool

const (
	DLPProfileCustomUpdateResponseEnvelopeSuccessTrue DLPProfileCustomUpdateResponseEnvelopeSuccess = true
)

func (r DLPProfileCustomUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfileCustomUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfileCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomDeleteResponseEnvelope struct {
	Errors   []DLPProfileCustomDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomDeleteResponse                `json:"result,nullable"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomDeleteResponseEnvelope]
type dlpProfileCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPProfileCustomDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfileCustomDeleteResponseEnvelopeErrors]
type dlpProfileCustomDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomDeleteResponseEnvelopeErrorsSource]
type dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPProfileCustomDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfileCustomDeleteResponseEnvelopeMessages]
type dlpProfileCustomDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfileCustomDeleteResponseEnvelopeMessagesSource]
type dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfileCustomGetResponseEnvelope struct {
	Errors   []DLPProfileCustomGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfileCustomGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfileCustomGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfileCustomGetResponse                `json:"result"`
	JSON    dlpProfileCustomGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelope]
type dlpProfileCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPProfileCustomGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfileCustomGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPProfileCustomGetResponseEnvelopeErrors]
type dlpProfileCustomGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseEnvelopeErrorsSource]
type dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPProfileCustomGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfileCustomGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPProfileCustomGetResponseEnvelopeMessages]
type dlpProfileCustomGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfileCustomGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPProfileCustomGetResponseEnvelopeMessagesSource]
type dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfileCustomGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfileCustomGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
