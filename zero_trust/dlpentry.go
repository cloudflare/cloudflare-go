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
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// DLPEntryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEntryService] method instead.
type DLPEntryService struct {
	Options     []option.RequestOption
	Custom      *DLPEntryCustomService
	Predefined  *DLPEntryPredefinedService
	Integration *DLPEntryIntegrationService
}

// NewDLPEntryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPEntryService(opts ...option.RequestOption) (r *DLPEntryService) {
	r = &DLPEntryService{}
	r.Options = opts
	r.Custom = NewDLPEntryCustomService(opts...)
	r.Predefined = NewDLPEntryPredefinedService(opts...)
	r.Integration = NewDLPEntryIntegrationService(opts...)
	return
}

// Creates a DLP custom entry.
func (r *DLPEntryService) New(ctx context.Context, params DLPEntryNewParams, opts ...option.RequestOption) (res *DLPEntryNewResponse, err error) {
	var env DLPEntryNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP entry.
func (r *DLPEntryService) Update(ctx context.Context, entryID string, params DLPEntryUpdateParams, opts ...option.RequestOption) (res *DLPEntryUpdateResponse, err error) {
	var env DLPEntryUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all DLP entries in an account.
func (r *DLPEntryService) List(ctx context.Context, query DLPEntryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPEntryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all DLP entries in an account.
func (r *DLPEntryService) ListAutoPaging(ctx context.Context, query DLPEntryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPEntryListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a DLP custom entry.
func (r *DLPEntryService) Delete(ctx context.Context, entryID string, body DLPEntryDeleteParams, opts ...option.RequestOption) (res *DLPEntryDeleteResponse, err error) {
	var env DLPEntryDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a DLP entry by ID.
func (r *DLPEntryService) Get(ctx context.Context, entryID string, query DLPEntryGetParams, opts ...option.RequestOption) (res *DLPEntryGetResponse, err error) {
	var env DLPEntryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", query.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPEntryNewResponse struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool      `json:"enabled,required"`
	Name        string    `json:"name,required"`
	Pattern     Pattern   `json:"pattern,required"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	Description string    `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryNewResponseJSON `json:"-"`
}

// dlpEntryNewResponseJSON contains the JSON metadata for the struct
// [DLPEntryNewResponse]
type dlpEntryNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                       `json:"enabled,required"`
	Name    string                     `json:"name,required"`
	Type    DLPEntryUpdateResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryUpdateResponsePredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id,nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryUpdateResponsePredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                `json:"word_list"`
	JSON     dlpEntryUpdateResponseJSON `json:"-"`
	union    DLPEntryUpdateResponseUnion
}

// dlpEntryUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponse]
type dlpEntryUpdateResponseJSON struct {
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

func (r dlpEntryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryUpdateResponseCustomEntry],
// [DLPEntryUpdateResponsePredefinedEntry],
// [DLPEntryUpdateResponseIntegrationEntry],
// [DLPEntryUpdateResponseExactDataEntry],
// [DLPEntryUpdateResponseDocumentFingerprintEntry],
// [DLPEntryUpdateResponseWordListEntry].
func (r DLPEntryUpdateResponse) AsUnion() DLPEntryUpdateResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryUpdateResponseCustomEntry],
// [DLPEntryUpdateResponsePredefinedEntry],
// [DLPEntryUpdateResponseIntegrationEntry],
// [DLPEntryUpdateResponseExactDataEntry],
// [DLPEntryUpdateResponseDocumentFingerprintEntry] or
// [DLPEntryUpdateResponseWordListEntry].
type DLPEntryUpdateResponseUnion interface {
	implementsDLPEntryUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryUpdateResponseWordListEntry{}),
		},
	)
}

type DLPEntryUpdateResponseCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                  `json:"enabled,required"`
	Name        string                                `json:"name,required"`
	Pattern     Pattern                               `json:"pattern,required"`
	Type        DLPEntryUpdateResponseCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                             `json:"updated_at,required" format:"date-time"`
	Description string                                `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID string                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseCustomEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseCustomEntry]
type dlpEntryUpdateResponseCustomEntryJSON struct {
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

func (r *DLPEntryUpdateResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseCustomEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseCustomEntryType string

const (
	DLPEntryUpdateResponseCustomEntryTypeCustom DLPEntryUpdateResponseCustomEntryType = "custom"
)

func (r DLPEntryUpdateResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryUpdateResponsePredefinedEntry struct {
	ID         string                                          `json:"id,required" format:"uuid"`
	Confidence DLPEntryUpdateResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                            `json:"enabled,required"`
	Name       string                                          `json:"name,required"`
	Type       DLPEntryUpdateResponsePredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	Variant   DLPEntryUpdateResponsePredefinedEntryVariant `json:"variant"`
	JSON      dlpEntryUpdateResponsePredefinedEntryJSON    `json:"-"`
}

// dlpEntryUpdateResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponsePredefinedEntry]
type dlpEntryUpdateResponsePredefinedEntryJSON struct {
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

func (r *DLPEntryUpdateResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponsePredefinedEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                `json:"available,required"`
	JSON      dlpEntryUpdateResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryUpdateResponsePredefinedEntryConfidenceJSON contains the JSON metadata
// for the struct [DLPEntryUpdateResponsePredefinedEntryConfidence]
type dlpEntryUpdateResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryUpdateResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponsePredefinedEntryType string

const (
	DLPEntryUpdateResponsePredefinedEntryTypePredefined DLPEntryUpdateResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryUpdateResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryUpdateResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryUpdateResponsePredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryUpdateResponsePredefinedEntryVariantType      `json:"type,required"`
	Description string                                                `json:"description,nullable"`
	JSON        dlpEntryUpdateResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryUpdateResponsePredefinedEntryVariantJSON contains the JSON metadata for
// the struct [DLPEntryUpdateResponsePredefinedEntryVariant]
type dlpEntryUpdateResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryUpdateResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryUpdateResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryUpdateResponsePredefinedEntryVariantTopicTypeContent DLPEntryUpdateResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryUpdateResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryUpdateResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryUpdateResponsePredefinedEntryVariantType string

const (
	DLPEntryUpdateResponsePredefinedEntryVariantTypePromptTopic DLPEntryUpdateResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryUpdateResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryUpdateResponseIntegrationEntry struct {
	ID        string                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                       `json:"enabled,required"`
	Name      string                                     `json:"name,required"`
	Type      DLPEntryUpdateResponseIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                  `json:"updated_at,required" format:"date-time"`
	ProfileID string                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseIntegrationEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseIntegrationEntry]
type dlpEntryUpdateResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryUpdateResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseIntegrationEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseIntegrationEntryType string

const (
	DLPEntryUpdateResponseIntegrationEntryTypeIntegration DLPEntryUpdateResponseIntegrationEntryType = "integration"
)

func (r DLPEntryUpdateResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryUpdateResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                     `json:"case_sensitive,required"`
	CreatedAt     time.Time                                `json:"created_at,required" format:"date-time"`
	Enabled       bool                                     `json:"enabled,required"`
	Name          string                                   `json:"name,required"`
	Secret        bool                                     `json:"secret,required"`
	Type          DLPEntryUpdateResponseExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryUpdateResponseExactDataEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseExactDataEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseExactDataEntry]
type dlpEntryUpdateResponseExactDataEntryJSON struct {
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

func (r *DLPEntryUpdateResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseExactDataEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseExactDataEntryType string

const (
	DLPEntryUpdateResponseExactDataEntryTypeExactData DLPEntryUpdateResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryUpdateResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryUpdateResponseDocumentFingerprintEntry struct {
	ID        string                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                               `json:"enabled,required"`
	Name      string                                             `json:"name,required"`
	Type      DLPEntryUpdateResponseDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                          `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryUpdateResponseDocumentFingerprintEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseDocumentFingerprintEntryJSON contains the JSON metadata
// for the struct [DLPEntryUpdateResponseDocumentFingerprintEntry]
type dlpEntryUpdateResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseDocumentFingerprintEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseDocumentFingerprintEntryType string

const (
	DLPEntryUpdateResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryUpdateResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryUpdateResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryUpdateResponseWordListEntry struct {
	ID        string                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                    `json:"enabled,required"`
	Name      string                                  `json:"name,required"`
	Type      DLPEntryUpdateResponseWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                               `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                             `json:"word_list,required"`
	ProfileID string                                  `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseWordListEntryJSON `json:"-"`
}

// dlpEntryUpdateResponseWordListEntryJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseWordListEntry]
type dlpEntryUpdateResponseWordListEntryJSON struct {
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

func (r *DLPEntryUpdateResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseWordListEntry) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseWordListEntryType string

const (
	DLPEntryUpdateResponseWordListEntryTypeWordList DLPEntryUpdateResponseWordListEntryType = "word_list"
)

func (r DLPEntryUpdateResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryUpdateResponseType string

const (
	DLPEntryUpdateResponseTypeCustom              DLPEntryUpdateResponseType = "custom"
	DLPEntryUpdateResponseTypePredefined          DLPEntryUpdateResponseType = "predefined"
	DLPEntryUpdateResponseTypeIntegration         DLPEntryUpdateResponseType = "integration"
	DLPEntryUpdateResponseTypeExactData           DLPEntryUpdateResponseType = "exact_data"
	DLPEntryUpdateResponseTypeDocumentFingerprint DLPEntryUpdateResponseType = "document_fingerprint"
	DLPEntryUpdateResponseTypeWordList            DLPEntryUpdateResponseType = "word_list"
)

func (r DLPEntryUpdateResponseType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseTypeCustom, DLPEntryUpdateResponseTypePredefined, DLPEntryUpdateResponseTypeIntegration, DLPEntryUpdateResponseTypeExactData, DLPEntryUpdateResponseTypeDocumentFingerprint, DLPEntryUpdateResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryListResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                     `json:"enabled,required"`
	Name    string                   `json:"name,required"`
	Type    DLPEntryListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of [DLPEntryListResponseObjectConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID    string                           `json:"profile_id,nullable" format:"uuid"`
	Secret       bool                             `json:"secret"`
	UpdatedAt    time.Time                        `json:"updated_at" format:"date-time"`
	UploadStatus DLPEntryListResponseUploadStatus `json:"upload_status"`
	// This field can have the runtime type of [DLPEntryListResponseObjectVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}              `json:"word_list"`
	JSON     dlpEntryListResponseJSON `json:"-"`
	union    DLPEntryListResponseUnion
}

// dlpEntryListResponseJSON contains the JSON metadata for the struct
// [DLPEntryListResponse]
type dlpEntryListResponseJSON struct {
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
	UploadStatus  apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpEntryListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryListResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryListResponseObject],
// [DLPEntryListResponseObject], [DLPEntryListResponseObject],
// [DLPEntryListResponseObject], [DLPEntryListResponseObject],
// [DLPEntryListResponseObject].
func (r DLPEntryListResponse) AsUnion() DLPEntryListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryListResponseObject], [DLPEntryListResponseObject],
// [DLPEntryListResponseObject], [DLPEntryListResponseObject],
// [DLPEntryListResponseObject] or [DLPEntryListResponseObject].
type DLPEntryListResponseUnion interface {
	implementsDLPEntryListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryListResponseObject{}),
		},
	)
}

type DLPEntryListResponseObject struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                           `json:"enabled,required"`
	Name        string                         `json:"name,required"`
	Pattern     Pattern                        `json:"pattern,required"`
	Type        DLPEntryListResponseObjectType `json:"type,required"`
	UpdatedAt   time.Time                      `json:"updated_at,required" format:"date-time"`
	Description string                         `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID    string                                 `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryListResponseObjectUploadStatus `json:"upload_status"`
	JSON         dlpEntryListResponseObjectJSON         `json:"-"`
}

// dlpEntryListResponseObjectJSON contains the JSON metadata for the struct
// [DLPEntryListResponseObject]
type dlpEntryListResponseObjectJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	Pattern      apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	Description  apijson.Field
	ProfileID    apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPEntryListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseObject) implementsDLPEntryListResponse() {}

type DLPEntryListResponseObjectType string

const (
	DLPEntryListResponseObjectTypeCustom DLPEntryListResponseObjectType = "custom"
)

func (r DLPEntryListResponseObjectType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseObjectTypeCustom:
		return true
	}
	return false
}

type DLPEntryListResponseObjectUploadStatus string

const (
	DLPEntryListResponseObjectUploadStatusEmpty      DLPEntryListResponseObjectUploadStatus = "empty"
	DLPEntryListResponseObjectUploadStatusUploading  DLPEntryListResponseObjectUploadStatus = "uploading"
	DLPEntryListResponseObjectUploadStatusPending    DLPEntryListResponseObjectUploadStatus = "pending"
	DLPEntryListResponseObjectUploadStatusProcessing DLPEntryListResponseObjectUploadStatus = "processing"
	DLPEntryListResponseObjectUploadStatusFailed     DLPEntryListResponseObjectUploadStatus = "failed"
	DLPEntryListResponseObjectUploadStatusComplete   DLPEntryListResponseObjectUploadStatus = "complete"
)

func (r DLPEntryListResponseObjectUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryListResponseObjectUploadStatusEmpty, DLPEntryListResponseObjectUploadStatusUploading, DLPEntryListResponseObjectUploadStatusPending, DLPEntryListResponseObjectUploadStatusProcessing, DLPEntryListResponseObjectUploadStatusFailed, DLPEntryListResponseObjectUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryListResponseType string

const (
	DLPEntryListResponseTypeCustom              DLPEntryListResponseType = "custom"
	DLPEntryListResponseTypePredefined          DLPEntryListResponseType = "predefined"
	DLPEntryListResponseTypeIntegration         DLPEntryListResponseType = "integration"
	DLPEntryListResponseTypeExactData           DLPEntryListResponseType = "exact_data"
	DLPEntryListResponseTypeDocumentFingerprint DLPEntryListResponseType = "document_fingerprint"
	DLPEntryListResponseTypeWordList            DLPEntryListResponseType = "word_list"
)

func (r DLPEntryListResponseType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseTypeCustom, DLPEntryListResponseTypePredefined, DLPEntryListResponseTypeIntegration, DLPEntryListResponseTypeExactData, DLPEntryListResponseTypeDocumentFingerprint, DLPEntryListResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryListResponseUploadStatus string

const (
	DLPEntryListResponseUploadStatusEmpty      DLPEntryListResponseUploadStatus = "empty"
	DLPEntryListResponseUploadStatusUploading  DLPEntryListResponseUploadStatus = "uploading"
	DLPEntryListResponseUploadStatusPending    DLPEntryListResponseUploadStatus = "pending"
	DLPEntryListResponseUploadStatusProcessing DLPEntryListResponseUploadStatus = "processing"
	DLPEntryListResponseUploadStatusFailed     DLPEntryListResponseUploadStatus = "failed"
	DLPEntryListResponseUploadStatusComplete   DLPEntryListResponseUploadStatus = "complete"
)

func (r DLPEntryListResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryListResponseUploadStatusEmpty, DLPEntryListResponseUploadStatusUploading, DLPEntryListResponseUploadStatusPending, DLPEntryListResponseUploadStatusProcessing, DLPEntryListResponseUploadStatusFailed, DLPEntryListResponseUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryDeleteResponse = interface{}

type DLPEntryGetResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                    `json:"enabled,required"`
	Name    string                  `json:"name,required"`
	Type    DLPEntryGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of [DLPEntryGetResponseObjectConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string `json:"profile_id,nullable" format:"uuid"`
	// This field can have the runtime type of [[]DLPEntryGetResponseObjectProfile].
	Profiles     interface{}                     `json:"profiles"`
	Secret       bool                            `json:"secret"`
	UpdatedAt    time.Time                       `json:"updated_at" format:"date-time"`
	UploadStatus DLPEntryGetResponseUploadStatus `json:"upload_status"`
	// This field can have the runtime type of [DLPEntryGetResponseObjectVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}             `json:"word_list"`
	JSON     dlpEntryGetResponseJSON `json:"-"`
	union    DLPEntryGetResponseUnion
}

// dlpEntryGetResponseJSON contains the JSON metadata for the struct
// [DLPEntryGetResponse]
type dlpEntryGetResponseJSON struct {
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
	Profiles      apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	UploadStatus  apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpEntryGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryGetResponseObject],
// [DLPEntryGetResponseObject], [DLPEntryGetResponseObject],
// [DLPEntryGetResponseObject], [DLPEntryGetResponseObject],
// [DLPEntryGetResponseObject].
func (r DLPEntryGetResponse) AsUnion() DLPEntryGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryGetResponseObject], [DLPEntryGetResponseObject],
// [DLPEntryGetResponseObject], [DLPEntryGetResponseObject],
// [DLPEntryGetResponseObject] or [DLPEntryGetResponseObject].
type DLPEntryGetResponseUnion interface {
	implementsDLPEntryGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseObject{}),
		},
	)
}

type DLPEntryGetResponseObject struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                          `json:"enabled,required"`
	Name        string                        `json:"name,required"`
	Pattern     Pattern                       `json:"pattern,required"`
	Type        DLPEntryGetResponseObjectType `json:"type,required"`
	UpdatedAt   time.Time                     `json:"updated_at,required" format:"date-time"`
	Description string                        `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID    string                                `json:"profile_id,nullable" format:"uuid"`
	Profiles     []DLPEntryGetResponseObjectProfile    `json:"profiles"`
	UploadStatus DLPEntryGetResponseObjectUploadStatus `json:"upload_status"`
	JSON         dlpEntryGetResponseObjectJSON         `json:"-"`
}

// dlpEntryGetResponseObjectJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseObject]
type dlpEntryGetResponseObjectJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	Pattern      apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	Description  apijson.Field
	ProfileID    apijson.Field
	Profiles     apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPEntryGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseObject) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseObjectType string

const (
	DLPEntryGetResponseObjectTypeCustom DLPEntryGetResponseObjectType = "custom"
)

func (r DLPEntryGetResponseObjectType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseObjectTypeCustom:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryGetResponseObjectProfile struct {
	ID   string                               `json:"id,required" format:"uuid"`
	Name string                               `json:"name,required"`
	JSON dlpEntryGetResponseObjectProfileJSON `json:"-"`
}

// dlpEntryGetResponseObjectProfileJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseObjectProfile]
type dlpEntryGetResponseObjectProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseObjectProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseObjectProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseObjectUploadStatus string

const (
	DLPEntryGetResponseObjectUploadStatusEmpty      DLPEntryGetResponseObjectUploadStatus = "empty"
	DLPEntryGetResponseObjectUploadStatusUploading  DLPEntryGetResponseObjectUploadStatus = "uploading"
	DLPEntryGetResponseObjectUploadStatusPending    DLPEntryGetResponseObjectUploadStatus = "pending"
	DLPEntryGetResponseObjectUploadStatusProcessing DLPEntryGetResponseObjectUploadStatus = "processing"
	DLPEntryGetResponseObjectUploadStatusFailed     DLPEntryGetResponseObjectUploadStatus = "failed"
	DLPEntryGetResponseObjectUploadStatusComplete   DLPEntryGetResponseObjectUploadStatus = "complete"
)

func (r DLPEntryGetResponseObjectUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseObjectUploadStatusEmpty, DLPEntryGetResponseObjectUploadStatusUploading, DLPEntryGetResponseObjectUploadStatusPending, DLPEntryGetResponseObjectUploadStatusProcessing, DLPEntryGetResponseObjectUploadStatusFailed, DLPEntryGetResponseObjectUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryGetResponseType string

const (
	DLPEntryGetResponseTypeCustom              DLPEntryGetResponseType = "custom"
	DLPEntryGetResponseTypePredefined          DLPEntryGetResponseType = "predefined"
	DLPEntryGetResponseTypeIntegration         DLPEntryGetResponseType = "integration"
	DLPEntryGetResponseTypeExactData           DLPEntryGetResponseType = "exact_data"
	DLPEntryGetResponseTypeDocumentFingerprint DLPEntryGetResponseType = "document_fingerprint"
	DLPEntryGetResponseTypeWordList            DLPEntryGetResponseType = "word_list"
)

func (r DLPEntryGetResponseType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseTypeCustom, DLPEntryGetResponseTypePredefined, DLPEntryGetResponseTypeIntegration, DLPEntryGetResponseTypeExactData, DLPEntryGetResponseTypeDocumentFingerprint, DLPEntryGetResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryGetResponseUploadStatus string

const (
	DLPEntryGetResponseUploadStatusEmpty      DLPEntryGetResponseUploadStatus = "empty"
	DLPEntryGetResponseUploadStatusUploading  DLPEntryGetResponseUploadStatus = "uploading"
	DLPEntryGetResponseUploadStatusPending    DLPEntryGetResponseUploadStatus = "pending"
	DLPEntryGetResponseUploadStatusProcessing DLPEntryGetResponseUploadStatus = "processing"
	DLPEntryGetResponseUploadStatusFailed     DLPEntryGetResponseUploadStatus = "failed"
	DLPEntryGetResponseUploadStatusComplete   DLPEntryGetResponseUploadStatus = "complete"
)

func (r DLPEntryGetResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseUploadStatusEmpty, DLPEntryGetResponseUploadStatusUploading, DLPEntryGetResponseUploadStatusPending, DLPEntryGetResponseUploadStatusProcessing, DLPEntryGetResponseUploadStatusFailed, DLPEntryGetResponseUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryNewParams struct {
	AccountID   param.Field[string]       `path:"account_id,required"`
	Enabled     param.Field[bool]         `json:"enabled,required"`
	Name        param.Field[string]       `json:"name,required"`
	Pattern     param.Field[PatternParam] `json:"pattern,required"`
	Description param.Field[string]       `json:"description"`
	ProfileID   param.Field[string]       `json:"profile_id" format:"uuid"`
}

func (r DLPEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryNewResponseEnvelope struct {
	Errors   []DLPEntryNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryNewResponse                `json:"result"`
	JSON    dlpEntryNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryNewResponseEnvelope]
type dlpEntryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DLPEntryNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DLPEntryNewResponseEnvelopeErrors]
type dlpEntryNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    dlpEntryNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPEntryNewResponseEnvelopeErrorsSource]
type dlpEntryNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPEntryNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryNewResponseEnvelopeMessages]
type dlpEntryNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryNewResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpEntryNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DLPEntryNewResponseEnvelopeMessagesSource]
type dlpEntryNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryNewResponseEnvelopeSuccess bool

const (
	DLPEntryNewResponseEnvelopeSuccessTrue DLPEntryNewResponseEnvelopeSuccess = true
)

func (r DLPEntryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryUpdateParams struct {
	AccountID param.Field[string]           `path:"account_id,required"`
	Body      DLPEntryUpdateParamsBodyUnion `json:"body,required"`
}

func (r DLPEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPEntryUpdateParamsBody struct {
	Type        param.Field[DLPEntryUpdateParamsBodyType] `json:"type,required"`
	Description param.Field[string]                       `json:"description"`
	Enabled     param.Field[bool]                         `json:"enabled"`
	Name        param.Field[string]                       `json:"name"`
	Pattern     param.Field[PatternParam]                 `json:"pattern"`
}

func (r DLPEntryUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBody) implementsDLPEntryUpdateParamsBodyUnion() {}

// Satisfied by [zero_trust.DLPEntryUpdateParamsBodyCustom],
// [zero_trust.DLPEntryUpdateParamsBodyPredefined],
// [zero_trust.DLPEntryUpdateParamsBodyIntegration], [DLPEntryUpdateParamsBody].
type DLPEntryUpdateParamsBodyUnion interface {
	implementsDLPEntryUpdateParamsBodyUnion()
}

type DLPEntryUpdateParamsBodyCustom struct {
	Name        param.Field[string]                             `json:"name,required"`
	Pattern     param.Field[PatternParam]                       `json:"pattern,required"`
	Type        param.Field[DLPEntryUpdateParamsBodyCustomType] `json:"type,required"`
	Description param.Field[string]                             `json:"description"`
	Enabled     param.Field[bool]                               `json:"enabled"`
}

func (r DLPEntryUpdateParamsBodyCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBodyCustom) implementsDLPEntryUpdateParamsBodyUnion() {}

type DLPEntryUpdateParamsBodyCustomType string

const (
	DLPEntryUpdateParamsBodyCustomTypeCustom DLPEntryUpdateParamsBodyCustomType = "custom"
)

func (r DLPEntryUpdateParamsBodyCustomType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryUpdateParamsBodyPredefined struct {
	Type    param.Field[DLPEntryUpdateParamsBodyPredefinedType] `json:"type,required"`
	Enabled param.Field[bool]                                   `json:"enabled"`
}

func (r DLPEntryUpdateParamsBodyPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBodyPredefined) implementsDLPEntryUpdateParamsBodyUnion() {}

type DLPEntryUpdateParamsBodyPredefinedType string

const (
	DLPEntryUpdateParamsBodyPredefinedTypePredefined DLPEntryUpdateParamsBodyPredefinedType = "predefined"
)

func (r DLPEntryUpdateParamsBodyPredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryUpdateParamsBodyIntegration struct {
	Type    param.Field[DLPEntryUpdateParamsBodyIntegrationType] `json:"type,required"`
	Enabled param.Field[bool]                                    `json:"enabled"`
}

func (r DLPEntryUpdateParamsBodyIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryUpdateParamsBodyIntegration) implementsDLPEntryUpdateParamsBodyUnion() {}

type DLPEntryUpdateParamsBodyIntegrationType string

const (
	DLPEntryUpdateParamsBodyIntegrationTypeIntegration DLPEntryUpdateParamsBodyIntegrationType = "integration"
)

func (r DLPEntryUpdateParamsBodyIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryUpdateParamsBodyType string

const (
	DLPEntryUpdateParamsBodyTypeCustom      DLPEntryUpdateParamsBodyType = "custom"
	DLPEntryUpdateParamsBodyTypePredefined  DLPEntryUpdateParamsBodyType = "predefined"
	DLPEntryUpdateParamsBodyTypeIntegration DLPEntryUpdateParamsBodyType = "integration"
)

func (r DLPEntryUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateParamsBodyTypeCustom, DLPEntryUpdateParamsBodyTypePredefined, DLPEntryUpdateParamsBodyTypeIntegration:
		return true
	}
	return false
}

type DLPEntryUpdateResponseEnvelope struct {
	Errors   []DLPEntryUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryUpdateResponse                `json:"result"`
	JSON    dlpEntryUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseEnvelope]
type dlpEntryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DLPEntryUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseEnvelopeErrors]
type dlpEntryUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dlpEntryUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEntryUpdateResponseEnvelopeErrorsSource]
type dlpEntryUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPEntryUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseEnvelopeMessages]
type dlpEntryUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpEntryUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPEntryUpdateResponseEnvelopeMessagesSource]
type dlpEntryUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryUpdateResponseEnvelopeSuccess bool

const (
	DLPEntryUpdateResponseEnvelopeSuccessTrue DLPEntryUpdateResponseEnvelopeSuccess = true
)

func (r DLPEntryUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryDeleteResponseEnvelope struct {
	Errors   []DLPEntryDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryDeleteResponse                `json:"result,nullable"`
	JSON    dlpEntryDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryDeleteResponseEnvelope]
type dlpEntryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DLPEntryDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEntryDeleteResponseEnvelopeErrors]
type dlpEntryDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dlpEntryDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEntryDeleteResponseEnvelopeErrorsSource]
type dlpEntryDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPEntryDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryDeleteResponseEnvelopeMessages]
type dlpEntryDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpEntryDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPEntryDeleteResponseEnvelopeMessagesSource]
type dlpEntryDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryDeleteResponseEnvelopeSuccess bool

const (
	DLPEntryDeleteResponseEnvelopeSuccessTrue DLPEntryDeleteResponseEnvelopeSuccess = true
)

func (r DLPEntryDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryGetResponseEnvelope struct {
	Errors   []DLPEntryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryGetResponse                `json:"result"`
	JSON    dlpEntryGetResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseEnvelope]
type dlpEntryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DLPEntryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseEnvelopeErrors]
type dlpEntryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    dlpEntryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseEnvelopeErrorsSource]
type dlpEntryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPEntryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseEnvelopeMessages]
type dlpEntryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpEntryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseEnvelopeMessagesSource]
type dlpEntryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryGetResponseEnvelopeSuccess bool

const (
	DLPEntryGetResponseEnvelopeSuccessTrue DLPEntryGetResponseEnvelopeSuccess = true
)

func (r DLPEntryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
