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
func (r *DLPProfilePredefinedService) Update(ctx context.Context, profileID string, params DLPProfilePredefinedUpdateParams, opts ...option.RequestOption) (res *DLPProfilePredefinedUpdateResponse, err error) {
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
func (r *DLPProfilePredefinedService) Get(ctx context.Context, profileID string, query DLPProfilePredefinedGetParams, opts ...option.RequestOption) (res *DLPProfilePredefinedGetResponse, err error) {
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

type DLPProfilePredefinedUpdateResponse struct {
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
	Entries []DLPProfilePredefinedUpdateResponseEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name             string `json:"name,required"`
	AIContextEnabled bool   `json:"ai_context_enabled"`
	OCREnabled       bool   `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                                   `json:"open_access"`
	JSON       dlpProfilePredefinedUpdateResponseJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseJSON contains the JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponse]
type dlpProfilePredefinedUpdateResponseJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                          `json:"enabled,required"`
	Name    string                                        `json:"name,required"`
	Type    DLPProfilePredefinedUpdateResponseEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                                 `json:"word_list"`
	JSON     dlpProfilePredefinedUpdateResponseEntryJSON `json:"-"`
	union    DLPProfilePredefinedUpdateResponseEntriesUnion
}

// dlpProfilePredefinedUpdateResponseEntryJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedUpdateResponseEntry]
type dlpProfilePredefinedUpdateResponseEntryJSON struct {
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

func (r dlpProfilePredefinedUpdateResponseEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedUpdateResponseEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedUpdateResponseEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedUpdateResponseEntriesUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfilePredefinedUpdateResponseEntriesCustomEntry],
// [DLPProfilePredefinedUpdateResponseEntriesPredefinedEntry],
// [DLPProfilePredefinedUpdateResponseEntriesIntegrationEntry],
// [DLPProfilePredefinedUpdateResponseEntriesExactDataEntry],
// [DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntry],
// [DLPProfilePredefinedUpdateResponseEntriesWordListEntry].
func (r DLPProfilePredefinedUpdateResponseEntry) AsUnion() DLPProfilePredefinedUpdateResponseEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfilePredefinedUpdateResponseEntriesCustomEntry],
// [DLPProfilePredefinedUpdateResponseEntriesPredefinedEntry],
// [DLPProfilePredefinedUpdateResponseEntriesIntegrationEntry],
// [DLPProfilePredefinedUpdateResponseEntriesExactDataEntry],
// [DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntry] or
// [DLPProfilePredefinedUpdateResponseEntriesWordListEntry].
type DLPProfilePredefinedUpdateResponseEntriesUnion interface {
	implementsDLPProfilePredefinedUpdateResponseEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedUpdateResponseEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedUpdateResponseEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedUpdateResponseEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedUpdateResponseEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedUpdateResponseEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedUpdateResponseEntriesWordListEntry{}),
		},
	)
}

type DLPProfilePredefinedUpdateResponseEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                     `json:"enabled,required"`
	Name        string                                                   `json:"name,required"`
	Pattern     Pattern                                                  `json:"pattern,required"`
	Type        DLPProfilePredefinedUpdateResponseEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                                `json:"updated_at,required" format:"date-time"`
	Description string                                                   `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseEntriesCustomEntryJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesCustomEntryJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseEntriesCustomEntry]
type dlpProfilePredefinedUpdateResponseEntriesCustomEntryJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseEntriesCustomEntry) implementsDLPProfilePredefinedUpdateResponseEntry() {
}

type DLPProfilePredefinedUpdateResponseEntriesCustomEntryType string

const (
	DLPProfilePredefinedUpdateResponseEntriesCustomEntryTypeCustom DLPProfilePredefinedUpdateResponseEntriesCustomEntryType = "custom"
)

func (r DLPProfilePredefinedUpdateResponseEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesPredefinedEntry struct {
	ID         string                                                             `json:"id,required" format:"uuid"`
	Confidence DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                               `json:"enabled,required"`
	Name       string                                                             `json:"name,required"`
	Type       DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                          `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseEntriesPredefinedEntry]
type dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseEntriesPredefinedEntry) implementsDLPProfilePredefinedUpdateResponseEntry() {
}

type DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                   `json:"available,required"`
	JSON      dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidenceJSON contains
// the JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidence]
type dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryType string

const (
	DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryTypePredefined DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                   `json:"description,nullable"`
	JSON        dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariant]
type dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicTypeContent DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantType string

const (
	DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTypePromptTopic DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesIntegrationEntry struct {
	ID        string                                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                          `json:"enabled,required"`
	Name      string                                                        `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseEntriesIntegrationEntry]
type dlpProfilePredefinedUpdateResponseEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseEntriesIntegrationEntry) implementsDLPProfilePredefinedUpdateResponseEntry() {
}

type DLPProfilePredefinedUpdateResponseEntriesIntegrationEntryType string

const (
	DLPProfilePredefinedUpdateResponseEntriesIntegrationEntryTypeIntegration DLPProfilePredefinedUpdateResponseEntriesIntegrationEntryType = "integration"
)

func (r DLPProfilePredefinedUpdateResponseEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                        `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                        `json:"enabled,required"`
	Name          string                                                      `json:"name,required"`
	Secret        bool                                                        `json:"secret,required"`
	Type          DLPProfilePredefinedUpdateResponseEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                   `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfilePredefinedUpdateResponseEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesExactDataEntryJSON contains the JSON
// metadata for the struct
// [DLPProfilePredefinedUpdateResponseEntriesExactDataEntry]
type dlpProfilePredefinedUpdateResponseEntriesExactDataEntryJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseEntriesExactDataEntry) implementsDLPProfilePredefinedUpdateResponseEntry() {
}

type DLPProfilePredefinedUpdateResponseEntriesExactDataEntryType string

const (
	DLPProfilePredefinedUpdateResponseEntriesExactDataEntryTypeExactData DLPProfilePredefinedUpdateResponseEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfilePredefinedUpdateResponseEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntry struct {
	ID        string                                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                                  `json:"enabled,required"`
	Name      string                                                                `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                             `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryJSON contains
// the JSON metadata for the struct
// [DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntry]
type dlpProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntry) implementsDLPProfilePredefinedUpdateResponseEntry() {
}

type DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryType string

const (
	DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesWordListEntry struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPProfilePredefinedUpdateResponseEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                                `json:"word_list,required"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedUpdateResponseEntriesWordListEntryJSON `json:"-"`
}

// dlpProfilePredefinedUpdateResponseEntriesWordListEntryJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedUpdateResponseEntriesWordListEntry]
type dlpProfilePredefinedUpdateResponseEntriesWordListEntryJSON struct {
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

func (r *DLPProfilePredefinedUpdateResponseEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedUpdateResponseEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedUpdateResponseEntriesWordListEntry) implementsDLPProfilePredefinedUpdateResponseEntry() {
}

type DLPProfilePredefinedUpdateResponseEntriesWordListEntryType string

const (
	DLPProfilePredefinedUpdateResponseEntriesWordListEntryTypeWordList DLPProfilePredefinedUpdateResponseEntriesWordListEntryType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponseEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedUpdateResponseEntriesType string

const (
	DLPProfilePredefinedUpdateResponseEntriesTypeCustom              DLPProfilePredefinedUpdateResponseEntriesType = "custom"
	DLPProfilePredefinedUpdateResponseEntriesTypePredefined          DLPProfilePredefinedUpdateResponseEntriesType = "predefined"
	DLPProfilePredefinedUpdateResponseEntriesTypeIntegration         DLPProfilePredefinedUpdateResponseEntriesType = "integration"
	DLPProfilePredefinedUpdateResponseEntriesTypeExactData           DLPProfilePredefinedUpdateResponseEntriesType = "exact_data"
	DLPProfilePredefinedUpdateResponseEntriesTypeDocumentFingerprint DLPProfilePredefinedUpdateResponseEntriesType = "document_fingerprint"
	DLPProfilePredefinedUpdateResponseEntriesTypeWordList            DLPProfilePredefinedUpdateResponseEntriesType = "word_list"
)

func (r DLPProfilePredefinedUpdateResponseEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedUpdateResponseEntriesTypeCustom, DLPProfilePredefinedUpdateResponseEntriesTypePredefined, DLPProfilePredefinedUpdateResponseEntriesTypeIntegration, DLPProfilePredefinedUpdateResponseEntriesTypeExactData, DLPProfilePredefinedUpdateResponseEntriesTypeDocumentFingerprint, DLPProfilePredefinedUpdateResponseEntriesTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedDeleteResponse = interface{}

type DLPProfilePredefinedGetResponse struct {
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
	Entries []DLPProfilePredefinedGetResponseEntry `json:"entries,required"`
	// The name of the predefined profile.
	Name             string `json:"name,required"`
	AIContextEnabled bool   `json:"ai_context_enabled"`
	OCREnabled       bool   `json:"ocr_enabled"`
	// Whether this profile can be accessed by anyone.
	OpenAccess bool                                `json:"open_access"`
	JSON       dlpProfilePredefinedGetResponseJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseJSON contains the JSON metadata for the struct
// [DLPProfilePredefinedGetResponse]
type dlpProfilePredefinedGetResponseJSON struct {
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

func (r *DLPProfilePredefinedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                       `json:"enabled,required"`
	Name    string                                     `json:"name,required"`
	Type    DLPProfilePredefinedGetResponseEntriesType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedGetResponseEntriesPredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                              `json:"word_list"`
	JSON     dlpProfilePredefinedGetResponseEntryJSON `json:"-"`
	union    DLPProfilePredefinedGetResponseEntriesUnion
}

// dlpProfilePredefinedGetResponseEntryJSON contains the JSON metadata for the
// struct [DLPProfilePredefinedGetResponseEntry]
type dlpProfilePredefinedGetResponseEntryJSON struct {
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

func (r dlpProfilePredefinedGetResponseEntryJSON) RawJSON() string {
	return r.raw
}

func (r *DLPProfilePredefinedGetResponseEntry) UnmarshalJSON(data []byte) (err error) {
	*r = DLPProfilePredefinedGetResponseEntry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPProfilePredefinedGetResponseEntriesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPProfilePredefinedGetResponseEntriesCustomEntry],
// [DLPProfilePredefinedGetResponseEntriesPredefinedEntry],
// [DLPProfilePredefinedGetResponseEntriesIntegrationEntry],
// [DLPProfilePredefinedGetResponseEntriesExactDataEntry],
// [DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntry],
// [DLPProfilePredefinedGetResponseEntriesWordListEntry].
func (r DLPProfilePredefinedGetResponseEntry) AsUnion() DLPProfilePredefinedGetResponseEntriesUnion {
	return r.union
}

// Union satisfied by [DLPProfilePredefinedGetResponseEntriesCustomEntry],
// [DLPProfilePredefinedGetResponseEntriesPredefinedEntry],
// [DLPProfilePredefinedGetResponseEntriesIntegrationEntry],
// [DLPProfilePredefinedGetResponseEntriesExactDataEntry],
// [DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntry] or
// [DLPProfilePredefinedGetResponseEntriesWordListEntry].
type DLPProfilePredefinedGetResponseEntriesUnion interface {
	implementsDLPProfilePredefinedGetResponseEntry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPProfilePredefinedGetResponseEntriesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedGetResponseEntriesCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedGetResponseEntriesPredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedGetResponseEntriesIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedGetResponseEntriesExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPProfilePredefinedGetResponseEntriesWordListEntry{}),
		},
	)
}

type DLPProfilePredefinedGetResponseEntriesCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                                  `json:"enabled,required"`
	Name        string                                                `json:"name,required"`
	Pattern     Pattern                                               `json:"pattern,required"`
	Type        DLPProfilePredefinedGetResponseEntriesCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                             `json:"updated_at,required" format:"date-time"`
	Description string                                                `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseEntriesCustomEntryJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesCustomEntryJSON contains the JSON metadata
// for the struct [DLPProfilePredefinedGetResponseEntriesCustomEntry]
type dlpProfilePredefinedGetResponseEntriesCustomEntryJSON struct {
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

func (r *DLPProfilePredefinedGetResponseEntriesCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseEntriesCustomEntry) implementsDLPProfilePredefinedGetResponseEntry() {
}

type DLPProfilePredefinedGetResponseEntriesCustomEntryType string

const (
	DLPProfilePredefinedGetResponseEntriesCustomEntryTypeCustom DLPProfilePredefinedGetResponseEntriesCustomEntryType = "custom"
)

func (r DLPProfilePredefinedGetResponseEntriesCustomEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesPredefinedEntry struct {
	ID         string                                                          `json:"id,required" format:"uuid"`
	Confidence DLPProfilePredefinedGetResponseEntriesPredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                            `json:"enabled,required"`
	Name       string                                                          `json:"name,required"`
	Type       DLPProfilePredefinedGetResponseEntriesPredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                                       `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariant `json:"variant"`
	JSON      dlpProfilePredefinedGetResponseEntriesPredefinedEntryJSON    `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesPredefinedEntryJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEntriesPredefinedEntry]
type dlpProfilePredefinedGetResponseEntriesPredefinedEntryJSON struct {
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

func (r *DLPProfilePredefinedGetResponseEntriesPredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesPredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseEntriesPredefinedEntry) implementsDLPProfilePredefinedGetResponseEntry() {
}

type DLPProfilePredefinedGetResponseEntriesPredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                                `json:"available,required"`
	JSON      dlpProfilePredefinedGetResponseEntriesPredefinedEntryConfidenceJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesPredefinedEntryConfidenceJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponseEntriesPredefinedEntryConfidence]
type dlpProfilePredefinedGetResponseEntriesPredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEntriesPredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesPredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEntriesPredefinedEntryType string

const (
	DLPProfilePredefinedGetResponseEntriesPredefinedEntryTypePredefined DLPProfilePredefinedGetResponseEntriesPredefinedEntryType = "predefined"
)

func (r DLPProfilePredefinedGetResponseEntriesPredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesPredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariant struct {
	TopicType   DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantType      `json:"type,required"`
	Description string                                                                `json:"description,nullable"`
	JSON        dlpProfilePredefinedGetResponseEntriesPredefinedEntryVariantJSON      `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesPredefinedEntryVariantJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariant]
type dlpProfilePredefinedGetResponseEntriesPredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesPredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicType string

const (
	DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicTypeIntent  DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicType = "Intent"
	DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicTypeContent DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicType = "Content"
)

func (r DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicTypeIntent, DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantType string

const (
	DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTypePromptTopic DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantType = "PromptTopic"
)

func (r DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesPredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesIntegrationEntry struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseEntriesIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseEntriesIntegrationEntryJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesIntegrationEntryJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEntriesIntegrationEntry]
type dlpProfilePredefinedGetResponseEntriesIntegrationEntryJSON struct {
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

func (r *DLPProfilePredefinedGetResponseEntriesIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseEntriesIntegrationEntry) implementsDLPProfilePredefinedGetResponseEntry() {
}

type DLPProfilePredefinedGetResponseEntriesIntegrationEntryType string

const (
	DLPProfilePredefinedGetResponseEntriesIntegrationEntryTypeIntegration DLPProfilePredefinedGetResponseEntriesIntegrationEntryType = "integration"
)

func (r DLPProfilePredefinedGetResponseEntriesIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                     `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                     `json:"enabled,required"`
	Name          string                                                   `json:"name,required"`
	Secret        bool                                                     `json:"secret,required"`
	Type          DLPProfilePredefinedGetResponseEntriesExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                                `json:"updated_at,required" format:"date-time"`
	JSON          dlpProfilePredefinedGetResponseEntriesExactDataEntryJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesExactDataEntryJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEntriesExactDataEntry]
type dlpProfilePredefinedGetResponseEntriesExactDataEntryJSON struct {
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

func (r *DLPProfilePredefinedGetResponseEntriesExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseEntriesExactDataEntry) implementsDLPProfilePredefinedGetResponseEntry() {
}

type DLPProfilePredefinedGetResponseEntriesExactDataEntryType string

const (
	DLPProfilePredefinedGetResponseEntriesExactDataEntryTypeExactData DLPProfilePredefinedGetResponseEntriesExactDataEntryType = "exact_data"
)

func (r DLPProfilePredefinedGetResponseEntriesExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntry struct {
	ID        string                                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                               `json:"enabled,required"`
	Name      string                                                             `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                          `json:"updated_at,required" format:"date-time"`
	JSON      dlpProfilePredefinedGetResponseEntriesDocumentFingerprintEntryJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesDocumentFingerprintEntryJSON contains the
// JSON metadata for the struct
// [DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntry]
type dlpProfilePredefinedGetResponseEntriesDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntry) implementsDLPProfilePredefinedGetResponseEntry() {
}

type DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntryType string

const (
	DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntryTypeDocumentFingerprint DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesWordListEntry struct {
	ID        string                                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                    `json:"enabled,required"`
	Name      string                                                  `json:"name,required"`
	Type      DLPProfilePredefinedGetResponseEntriesWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                               `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                             `json:"word_list,required"`
	ProfileID string                                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpProfilePredefinedGetResponseEntriesWordListEntryJSON `json:"-"`
}

// dlpProfilePredefinedGetResponseEntriesWordListEntryJSON contains the JSON
// metadata for the struct [DLPProfilePredefinedGetResponseEntriesWordListEntry]
type dlpProfilePredefinedGetResponseEntriesWordListEntryJSON struct {
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

func (r *DLPProfilePredefinedGetResponseEntriesWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpProfilePredefinedGetResponseEntriesWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPProfilePredefinedGetResponseEntriesWordListEntry) implementsDLPProfilePredefinedGetResponseEntry() {
}

type DLPProfilePredefinedGetResponseEntriesWordListEntryType string

const (
	DLPProfilePredefinedGetResponseEntriesWordListEntryTypeWordList DLPProfilePredefinedGetResponseEntriesWordListEntryType = "word_list"
)

func (r DLPProfilePredefinedGetResponseEntriesWordListEntryType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPProfilePredefinedGetResponseEntriesType string

const (
	DLPProfilePredefinedGetResponseEntriesTypeCustom              DLPProfilePredefinedGetResponseEntriesType = "custom"
	DLPProfilePredefinedGetResponseEntriesTypePredefined          DLPProfilePredefinedGetResponseEntriesType = "predefined"
	DLPProfilePredefinedGetResponseEntriesTypeIntegration         DLPProfilePredefinedGetResponseEntriesType = "integration"
	DLPProfilePredefinedGetResponseEntriesTypeExactData           DLPProfilePredefinedGetResponseEntriesType = "exact_data"
	DLPProfilePredefinedGetResponseEntriesTypeDocumentFingerprint DLPProfilePredefinedGetResponseEntriesType = "document_fingerprint"
	DLPProfilePredefinedGetResponseEntriesTypeWordList            DLPProfilePredefinedGetResponseEntriesType = "word_list"
)

func (r DLPProfilePredefinedGetResponseEntriesType) IsKnown() bool {
	switch r {
	case DLPProfilePredefinedGetResponseEntriesTypeCustom, DLPProfilePredefinedGetResponseEntriesTypePredefined, DLPProfilePredefinedGetResponseEntriesTypeIntegration, DLPProfilePredefinedGetResponseEntriesTypeExactData, DLPProfilePredefinedGetResponseEntriesTypeDocumentFingerprint, DLPProfilePredefinedGetResponseEntriesTypeWordList:
		return true
	}
	return false
}

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
	Result  DLPProfilePredefinedUpdateResponse                `json:"result"`
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
	Result  DLPProfilePredefinedGetResponse                `json:"result"`
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
