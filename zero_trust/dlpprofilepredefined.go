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

// DLPProfilePredefinedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPProfilePredefinedService] method instead.
type DLPProfilePredefinedService struct {
	Options []option.RequestOption
}

// NewDLPProfilePredefinedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPProfilePredefinedService(opts ...option.RequestOption) (r *DLPProfilePredefinedService) {
	r = &DLPProfilePredefinedService{}
	r.Options = opts
	return
}

// This is similar to `update_predefined` but only returns entries that are
// enabled. This is needed for our terraform API Updates a DLP predefined profile.
// Only supports enabling/disabling entries.
func (r *DLPProfilePredefinedService) Update(ctx context.Context, profileID string, params DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *PredefinedProfile, err error) {
	var env DLPProfilePredefinedUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s/config", params.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This is a no-op as predefined profiles can't be deleted but is needed for our
// generated terraform API.
func (r *DLPProfilePredefinedService) Delete(ctx context.Context, profileID string, body DLPProfilePredefinedDeleteParams, opts ...option.RequestOption) (res *DLPProfilePredefinedDeleteResponse, err error) {
	var env DLPProfilePredefinedDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s", body.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This is similar to `get_predefined` but only returns entries that are enabled.
// This is needed for our terraform API Fetches a predefined DLP profile by id.
func (r *DLPProfilePredefinedService) Get(ctx context.Context, profileID string, query DLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *PredefinedProfile, err error) {
	var env DLPProfilePredefinedGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if profileID == "" {
		err = errors.New("missing required profile_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/profiles/predefined/%s/config", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PredefinedProfile struct {
	// The id of the predefined profile (uuid).
	ID                  string `json:"id,required" format:"uuid"`
	AllowedMatchCount   int64  `json:"allowed_match_count,required"`
	ConfidenceThreshold string `json:"confidence_threshold,required,nullable"`
	// Entries to enable for this predefined profile. Any entries not provided will be
	// disabled.
	EnabledEntries []string `json:"enabled_entries,required" format:"uuid"`
	// This field has been deprecated for `enabled_entries`.
	//
	// Deprecated: deprecated
	Entries []PredefinedProfileEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name             string `json:"name,required"`
	AIContextEnabled bool   `json:"ai_context_enabled"`
	OCREnabled       bool   `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                  `json:"open_access"`
	JSON       predefinedProfileJSON `json:"-"`
}

// predefinedProfileJSON contains the JSON metadata for the struct
// [PredefinedProfile]
type predefinedProfileJSON struct {
	ID                  apijson.Field
	AllowedMatchCount   apijson.Field
	ConfidenceThreshold apijson.Field
	EnabledEntries      apijson.Field
	Entries             apijson.Field
	Name                apijson.Field
	AIContextEnabled    apijson.Field
	OCREnabled          apijson.Field
	OpenAccess          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PredefinedProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileJSON) RawJSON() string {
	return r.raw
}

type PredefinedProfileEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                         `json:"enabled,required"`
	Name    string                       `json:"name,required"`
	Type    PredefinedProfileEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [PredefinedProfileEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [PredefinedProfileEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                `json:"word_list"`
	JSON     predefinedProfileEntryJSON `json:"-"`
	union    PredefinedProfileEntriesUnion
}

// predefinedProfileEntryJSON contains the JSON metadata for the struct
// [PredefinedProfileEntry]
type predefinedProfileEntryJSON struct {
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

func (r predefinedProfileEntryJSON) RawJSON() string {
	return r.raw
}

func (r *PredefinedProfileEntry) UnmarshalJSON(data []byte) (err error) {
	*r = PredefinedProfileEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PredefinedProfileEntriesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [PredefinedProfileEntriesCustomEntry],
// [PredefinedProfileEntriesPredefinedEntry],
// [PredefinedProfileEntriesIntegrationEntry],
// [PredefinedProfileEntriesExactDataEntry],
// [PredefinedProfileEntriesDocumentFingerprintEntry],
// [PredefinedProfileEntriesWordListEntry].
func (r PredefinedProfileEntry) AsUnion() PredefinedProfileEntriesUnion {
	return r.union
}

// Union satisfied by [PredefinedProfileEntriesCustomEntry],
// [PredefinedProfileEntriesPredefinedEntry],
// [PredefinedProfileEntriesIntegrationEntry],
// [PredefinedProfileEntriesExactDataEntry],
// [PredefinedProfileEntriesDocumentFingerprintEntry] or
// [PredefinedProfileEntriesWordListEntry].
type PredefinedProfileEntriesUnion interface {
	implementsPredefinedProfileEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PredefinedProfileEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfileEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfileEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfileEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfileEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfileEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PredefinedProfileEntriesWordListEntry{}),
		},
	)
}

type PredefinedProfileEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                    `json:"enabled,required"`
	Name        string                                  `json:"name,required"`
	Pattern     Pattern                                 `json:"pattern,required"`
	Type        PredefinedProfileEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                               `json:"updated_at,required" format:"date-time"`
	Description string                                  `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      predefinedProfileEntriesCustomEntryJSON `json:"-"`
}

// predefinedProfileEntriesCustomEntryJSON contains the JSON metadata for the
// struct [PredefinedProfileEntriesCustomEntry]
type predefinedProfileEntriesCustomEntryJSON struct {
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

func (r *PredefinedProfileEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r PredefinedProfileEntriesCustomEntry) implementsPredefinedProfileEntry() {}

type PredefinedProfileEntriesCustomEntryType string

const (
	PredefinedProfileEntriesCustomEntryTypeCustom PredefinedProfileEntriesCustomEntryType = "custom"
)

func (r PredefinedProfileEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type PredefinedProfileEntriesPredefinedEntry struct {
	ID         string                                            `json:"id,required" format:"uuid"`
	Confidence PredefinedProfileEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                              `json:"enabled,required"`
	Name       string                                            `json:"name,required"`
	Type       PredefinedProfileEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                         `json:"profile_id,nullable" format:"uuid"`
	Variant   PredefinedProfileEntriesPredefinedEntryVariant `json:"variant"`
	JSON      predefinedProfileEntriesPredefinedEntryJSON    `json:"-"`
}

// predefinedProfileEntriesPredefinedEntryJSON contains the JSON metadata for the
// struct [PredefinedProfileEntriesPredefinedEntry]
type predefinedProfileEntriesPredefinedEntryJSON struct {
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

func (r *PredefinedProfileEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r PredefinedProfileEntriesPredefinedEntry) implementsPredefinedProfileEntry() {}

type PredefinedProfileEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                  `json:"available,required"`
	JSON      predefinedProfileEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// predefinedProfileEntriesPredefinedEntryConfidenceJSON contains the JSON metadata
// for the struct [PredefinedProfileEntriesPredefinedEntryConfidence]
type predefinedProfileEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PredefinedProfileEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type PredefinedProfileEntriesPredefinedEntryType string

const (
	PredefinedProfileEntriesPredefinedEntryTypePredefined PredefinedProfileEntriesPredefinedEntryType = "predefined"
)

func (r PredefinedProfileEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type PredefinedProfileEntriesPredefinedEntryVariant struct {
	TopicType   PredefinedProfileEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        PredefinedProfileEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                  `json:"description,nullable"`
	JSON        predefinedProfileEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// predefinedProfileEntriesPredefinedEntryVariantJSON contains the JSON metadata
// for the struct [PredefinedProfileEntriesPredefinedEntryVariant]
type predefinedProfileEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PredefinedProfileEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type PredefinedProfileEntriesPredefinedEntryVariantTopicType string

const (
	PredefinedProfileEntriesPredefinedEntryVariantTopicTypeIntent  PredefinedProfileEntriesPredefinedEntryVariantTopicType = "Intent"
	PredefinedProfileEntriesPredefinedEntryVariantTopicTypeContent PredefinedProfileEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r PredefinedProfileEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesPredefinedEntryVariantTopicTypeIntent, PredefinedProfileEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type PredefinedProfileEntriesPredefinedEntryVariantType string

const (
	PredefinedProfileEntriesPredefinedEntryVariantTypePromptTopic PredefinedProfileEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r PredefinedProfileEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type PredefinedProfileEntriesIntegrationEntry struct {
	ID        string                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                         `json:"enabled,required"`
	Name      string                                       `json:"name,required"`
	Type      PredefinedProfileEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      predefinedProfileEntriesIntegrationEntryJSON `json:"-"`
}

// predefinedProfileEntriesIntegrationEntryJSON contains the JSON metadata for the
// struct [PredefinedProfileEntriesIntegrationEntry]
type predefinedProfileEntriesIntegrationEntryJSON struct {
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

func (r *PredefinedProfileEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r PredefinedProfileEntriesIntegrationEntry) implementsPredefinedProfileEntry() {}

type PredefinedProfileEntriesIntegrationEntryType string

const (
	PredefinedProfileEntriesIntegrationEntryTypeIntegration PredefinedProfileEntriesIntegrationEntryType = "integration"
)

func (r PredefinedProfileEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type PredefinedProfileEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                       `json:"case_sensitive,required"`
	CreatedAt     time.Time                                  `json:"created_at,required" format:"date-time"`
	Enabled       bool                                       `json:"enabled,required"`
	Name          string                                     `json:"name,required"`
	Secret        bool                                       `json:"secret,required"`
	Type          PredefinedProfileEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                  `json:"updated_at,required" format:"date-time"`
	JSON          predefinedProfileEntriesExactDataEntryJSON `json:"-"`
}

// predefinedProfileEntriesExactDataEntryJSON contains the JSON metadata for the
// struct [PredefinedProfileEntriesExactDataEntry]
type predefinedProfileEntriesExactDataEntryJSON struct {
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

func (r *PredefinedProfileEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r PredefinedProfileEntriesExactDataEntry) implementsPredefinedProfileEntry() {}

type PredefinedProfileEntriesExactDataEntryType string

const (
	PredefinedProfileEntriesExactDataEntryTypeExactData PredefinedProfileEntriesExactDataEntryType = "exact_data"
)

func (r PredefinedProfileEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type PredefinedProfileEntriesDocumentFingerprintEntry struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      PredefinedProfileEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	JSON      predefinedProfileEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// predefinedProfileEntriesDocumentFingerprintEntryJSON contains the JSON metadata
// for the struct [PredefinedProfileEntriesDocumentFingerprintEntry]
type predefinedProfileEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PredefinedProfileEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r PredefinedProfileEntriesDocumentFingerprintEntry) implementsPredefinedProfileEntry() {}

type PredefinedProfileEntriesDocumentFingerprintEntryType string

const (
	PredefinedProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint PredefinedProfileEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r PredefinedProfileEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type PredefinedProfileEntriesWordListEntry struct {
	ID        string                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                      `json:"enabled,required"`
	Name      string                                    `json:"name,required"`
	Type      PredefinedProfileEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                               `json:"word_list,required"`
	ProfileID string                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      predefinedProfileEntriesWordListEntryJSON `json:"-"`
}

// predefinedProfileEntriesWordListEntryJSON contains the JSON metadata for the
// struct [PredefinedProfileEntriesWordListEntry]
type predefinedProfileEntriesWordListEntryJSON struct {
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

func (r *PredefinedProfileEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r predefinedProfileEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r PredefinedProfileEntriesWordListEntry) implementsPredefinedProfileEntry() {}

type PredefinedProfileEntriesWordListEntryType string

const (
	PredefinedProfileEntriesWordListEntryTypeWordList PredefinedProfileEntriesWordListEntryType = "word_list"
)

func (r PredefinedProfileEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type PredefinedProfileEntriesType string

const (
	PredefinedProfileEntriesTypeCustom              PredefinedProfileEntriesType = "custom"
	PredefinedProfileEntriesTypePredefined          PredefinedProfileEntriesType = "predefined"
	PredefinedProfileEntriesTypeIntegration         PredefinedProfileEntriesType = "integration"
	PredefinedProfileEntriesTypeExactData           PredefinedProfileEntriesType = "exact_data"
	PredefinedProfileEntriesTypeDocumentFingerprint PredefinedProfileEntriesType = "document_fingerprint"
	PredefinedProfileEntriesTypeWordList            PredefinedProfileEntriesType = "word_list"
)

func (r PredefinedProfileEntriesType) IsKnown() bool {
	switch r {
	case PredefinedProfileEntriesTypeCustom, PredefinedProfileEntriesTypePredefined, PredefinedProfileEntriesTypeIntegration, PredefinedProfileEntriesTypeExactData, PredefinedProfileEntriesTypeDocumentFingerprint, PredefinedProfileEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedDeleteResponse = interface{}

type DLPProfilePredefinedUpdateParams struct {
	AccountID           param.Field[string]                                  `path:"account_id,required"`
	AIContextEnabled    param.Field[bool]                                    `json:"ai_context_enabled"`
	AllowedMatchCount   param.Field[int64]                                   `json:"allowed_match_count"`
	ConfidenceThreshold param.Field[string]                                  `json:"confidence_threshold"`
	EnabledEntries      param.Field[[]string]                                `json:"enabled_entries" format:"uuid"`
	Entries             param.Field[[]DLPProfilePredefinedUpdateParamsEntry] `json:"entries"`
	OCREnabled          param.Field[bool]                                    `json:"ocr_enabled"`
}

func (r DLPProfilePredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedUpdateParamsEntry struct {
	ID      param.Field[string] `json:"id,required" format:"uuid"`
	Enabled param.Field[bool]   `json:"enabled,required"`
}

func (r DLPProfilePredefinedUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPProfilePredefinedUpdateResponseEnvelope struct {
	Errors   []DLPProfilePredefinedUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfilePredefinedUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  PredefinedProfile                                 `json:"result"`
	JSON    dlpProfilePredefinedUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedUpdateResponseEnvelope]
type dlpProfilePredefinedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedUpdateResponseEnvelopeErrors]
type dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource]
type dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseEnvelopeMessages]
type dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource]
type dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfilePredefinedUpdateResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedUpdateResponseEnvelopeSuccessTrue DLPProfilePredefinedUpdateResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfilePredefinedDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedDeleteResponseEnvelope struct {
	Errors   []DLPProfilePredefinedDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfilePredefinedDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPProfilePredefinedDeleteResponse                `json:"result,nullable"`
	JSON    dlpProfilePredefinedDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedDeleteResponseEnvelope]
type dlpProfilePredefinedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedDeleteResponseEnvelopeErrors]
type dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource]
type dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedDeleteResponseEnvelopeMessages]
type dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource]
type dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfilePredefinedDeleteResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedDeleteResponseEnvelopeSuccessTrue DLPProfilePredefinedDeleteResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPProfilePredefinedGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPProfilePredefinedGetResponseEnvelope struct {
	Errors   []DLPProfilePredefinedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPProfilePredefinedGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPProfilePredefinedGetResponseEnvelopeSuccess `json:"success,required"`
	Result  PredefinedProfile                              `json:"result"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseEnvelope]
type dlpProfilePredefinedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPProfilePredefinedGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpProfilePredefinedGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPProfilePredefinedGetResponseEnvelopeErrors]
type dlpProfilePredefinedGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEnvelopeErrorsSource]
type dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPProfilePredefinedGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpProfilePredefinedGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponseEnvelopeMessages]
type dlpProfilePredefinedGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEnvelopeMessagesSource]
type dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPProfilePredefinedGetResponseEnvelopeSuccess bool

const (
	DLPProfilePredefinedGetResponseEnvelopeSuccessTrue DLPProfilePredefinedGetResponseEnvelopeSuccess = true
)

func (r DLPProfilePredefinedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
