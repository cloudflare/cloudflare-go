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

// DLPEntryIntegrationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEntryIntegrationService] method instead.
type DLPEntryIntegrationService struct {
	Options []option.RequestOption
}

// NewDLPEntryIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPEntryIntegrationService(opts ...option.RequestOption) (r *DLPEntryIntegrationService) {
	r = &DLPEntryIntegrationService{}
	r.Options = opts
	return
}

// Integration entries can't be created, this will update an existing integration
// entry. This is needed for our generated terraform API.
func (r *DLPEntryIntegrationService) New(ctx context.Context, params DLPEntryIntegrationNewParams, opts ...option.RequestOption) (res *DLPEntryIntegrationNewResponse, err error) {
	var env DLPEntryIntegrationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/integration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP entry.
func (r *DLPEntryIntegrationService) Update(ctx context.Context, entryID string, params DLPEntryIntegrationUpdateParams, opts ...option.RequestOption) (res *DLPEntryIntegrationUpdateResponse, err error) {
	var env DLPEntryIntegrationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/integration/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all DLP entries in an account.
func (r *DLPEntryIntegrationService) List(ctx context.Context, query DLPEntryIntegrationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPEntryIntegrationListResponse], err error) {
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
func (r *DLPEntryIntegrationService) ListAutoPaging(ctx context.Context, query DLPEntryIntegrationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPEntryIntegrationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// This is a no-op as integration entires can't be deleted but is needed for our
// generated terraform API.
func (r *DLPEntryIntegrationService) Delete(ctx context.Context, entryID string, body DLPEntryIntegrationDeleteParams, opts ...option.RequestOption) (res *DLPEntryIntegrationDeleteResponse, err error) {
	var env DLPEntryIntegrationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/integration/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a DLP entry by ID.
func (r *DLPEntryIntegrationService) Get(ctx context.Context, entryID string, query DLPEntryIntegrationGetParams, opts ...option.RequestOption) (res *DLPEntryIntegrationGetResponse, err error) {
	var env DLPEntryIntegrationGetResponseEnvelope
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

type DLPEntryIntegrationNewResponse struct {
	ID        string                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                               `json:"enabled,required"`
	Name      string                             `json:"name,required"`
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationNewResponseJSON `json:"-"`
}

// dlpEntryIntegrationNewResponseJSON contains the JSON metadata for the struct
// [DLPEntryIntegrationNewResponse]
type dlpEntryIntegrationNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponse struct {
	ID        string                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                  `json:"enabled,required"`
	Name      string                                `json:"name,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationUpdateResponseJSON `json:"-"`
}

// dlpEntryIntegrationUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryIntegrationUpdateResponse]
type dlpEntryIntegrationUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationListResponse struct {
	ID      string                              `json:"id,required" format:"uuid"`
	Enabled bool                                `json:"enabled,required"`
	Name    string                              `json:"name,required"`
	Type    DLPEntryIntegrationListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationListResponsePredefinedConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationListResponsePredefinedVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                         `json:"word_list"`
	JSON     dlpEntryIntegrationListResponseJSON `json:"-"`
	union    DLPEntryIntegrationListResponseUnion
}

// dlpEntryIntegrationListResponseJSON contains the JSON metadata for the struct
// [DLPEntryIntegrationListResponse]
type dlpEntryIntegrationListResponseJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpEntryIntegrationListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryIntegrationListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryIntegrationListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryIntegrationListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryIntegrationListResponseCustom],
// [DLPEntryIntegrationListResponsePredefined],
// [DLPEntryIntegrationListResponseIntegration],
// [DLPEntryIntegrationListResponseExactData],
// [DLPEntryIntegrationListResponseDocumentFingerprint],
// [DLPEntryIntegrationListResponseWordList].
func (r DLPEntryIntegrationListResponse) AsUnion() DLPEntryIntegrationListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryIntegrationListResponseCustom],
// [DLPEntryIntegrationListResponsePredefined],
// [DLPEntryIntegrationListResponseIntegration],
// [DLPEntryIntegrationListResponseExactData],
// [DLPEntryIntegrationListResponseDocumentFingerprint] or
// [DLPEntryIntegrationListResponseWordList].
type DLPEntryIntegrationListResponseUnion interface {
	implementsDLPEntryIntegrationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryIntegrationListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationListResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationListResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationListResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationListResponseExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationListResponseDocumentFingerprint{}),
			DiscriminatorValue: "document_fingerprint",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationListResponseWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPEntryIntegrationListResponseCustom struct {
	ID        string                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                      `json:"enabled,required"`
	Name      string                                    `json:"name,required"`
	Pattern   Pattern                                   `json:"pattern,required"`
	Type      DLPEntryIntegrationListResponseCustomType `json:"type,required"`
	UpdatedAt time.Time                                 `json:"updated_at,required" format:"date-time"`
	ProfileID string                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationListResponseCustomJSON `json:"-"`
}

// dlpEntryIntegrationListResponseCustomJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationListResponseCustom]
type dlpEntryIntegrationListResponseCustomJSON struct {
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

func (r *DLPEntryIntegrationListResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseCustom) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponseCustomType string

const (
	DLPEntryIntegrationListResponseCustomTypeCustom DLPEntryIntegrationListResponseCustomType = "custom"
)

func (r DLPEntryIntegrationListResponseCustomType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponsePredefined struct {
	ID         string                                              `json:"id,required" format:"uuid"`
	Confidence DLPEntryIntegrationListResponsePredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                                `json:"enabled,required"`
	Name       string                                              `json:"name,required"`
	Type       DLPEntryIntegrationListResponsePredefinedType       `json:"type,required"`
	ProfileID  string                                              `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryIntegrationListResponsePredefinedVariant    `json:"variant"`
	JSON       dlpEntryIntegrationListResponsePredefinedJSON       `json:"-"`
}

// dlpEntryIntegrationListResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationListResponsePredefined]
type dlpEntryIntegrationListResponsePredefinedJSON struct {
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

func (r *DLPEntryIntegrationListResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponsePredefined) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponsePredefinedConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                    `json:"available,required"`
	JSON      dlpEntryIntegrationListResponsePredefinedConfidenceJSON `json:"-"`
}

// dlpEntryIntegrationListResponsePredefinedConfidenceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationListResponsePredefinedConfidence]
type dlpEntryIntegrationListResponsePredefinedConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponsePredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponsePredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationListResponsePredefinedType string

const (
	DLPEntryIntegrationListResponsePredefinedTypePredefined DLPEntryIntegrationListResponsePredefinedType = "predefined"
)

func (r DLPEntryIntegrationListResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponsePredefinedVariant struct {
	TopicType   DLPEntryIntegrationListResponsePredefinedVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryIntegrationListResponsePredefinedVariantType      `json:"type,required"`
	Description string                                                    `json:"description,nullable"`
	JSON        dlpEntryIntegrationListResponsePredefinedVariantJSON      `json:"-"`
}

// dlpEntryIntegrationListResponsePredefinedVariantJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationListResponsePredefinedVariant]
type dlpEntryIntegrationListResponsePredefinedVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponsePredefinedVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponsePredefinedVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationListResponsePredefinedVariantTopicType string

const (
	DLPEntryIntegrationListResponsePredefinedVariantTopicTypeIntent  DLPEntryIntegrationListResponsePredefinedVariantTopicType = "Intent"
	DLPEntryIntegrationListResponsePredefinedVariantTopicTypeContent DLPEntryIntegrationListResponsePredefinedVariantTopicType = "Content"
)

func (r DLPEntryIntegrationListResponsePredefinedVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponsePredefinedVariantTopicTypeIntent, DLPEntryIntegrationListResponsePredefinedVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponsePredefinedVariantType string

const (
	DLPEntryIntegrationListResponsePredefinedVariantTypePromptTopic DLPEntryIntegrationListResponsePredefinedVariantType = "PromptTopic"
)

func (r DLPEntryIntegrationListResponsePredefinedVariantType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponsePredefinedVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseIntegration struct {
	ID        string                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                           `json:"enabled,required"`
	Name      string                                         `json:"name,required"`
	Type      DLPEntryIntegrationListResponseIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationListResponseIntegrationJSON `json:"-"`
}

// dlpEntryIntegrationListResponseIntegrationJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationListResponseIntegration]
type dlpEntryIntegrationListResponseIntegrationJSON struct {
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

func (r *DLPEntryIntegrationListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseIntegration) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponseIntegrationType string

const (
	DLPEntryIntegrationListResponseIntegrationTypeIntegration DLPEntryIntegrationListResponseIntegrationType = "integration"
)

func (r DLPEntryIntegrationListResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseExactData struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                         `json:"case_sensitive,required"`
	CreatedAt     time.Time                                    `json:"created_at,required" format:"date-time"`
	Enabled       bool                                         `json:"enabled,required"`
	Name          string                                       `json:"name,required"`
	Secret        bool                                         `json:"secret,required"`
	Type          DLPEntryIntegrationListResponseExactDataType `json:"type,required"`
	UpdatedAt     time.Time                                    `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryIntegrationListResponseExactDataJSON `json:"-"`
}

// dlpEntryIntegrationListResponseExactDataJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationListResponseExactData]
type dlpEntryIntegrationListResponseExactDataJSON struct {
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

func (r *DLPEntryIntegrationListResponseExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseExactData) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponseExactDataType string

const (
	DLPEntryIntegrationListResponseExactDataTypeExactData DLPEntryIntegrationListResponseExactDataType = "exact_data"
)

func (r DLPEntryIntegrationListResponseExactDataType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseExactDataTypeExactData:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseDocumentFingerprint struct {
	ID        string                                                 `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                              `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                   `json:"enabled,required"`
	Name      string                                                 `json:"name,required"`
	Type      DLPEntryIntegrationListResponseDocumentFingerprintType `json:"type,required"`
	UpdatedAt time.Time                                              `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryIntegrationListResponseDocumentFingerprintJSON `json:"-"`
}

// dlpEntryIntegrationListResponseDocumentFingerprintJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationListResponseDocumentFingerprint]
type dlpEntryIntegrationListResponseDocumentFingerprintJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponseDocumentFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseDocumentFingerprintJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseDocumentFingerprint) implementsDLPEntryIntegrationListResponse() {
}

type DLPEntryIntegrationListResponseDocumentFingerprintType string

const (
	DLPEntryIntegrationListResponseDocumentFingerprintTypeDocumentFingerprint DLPEntryIntegrationListResponseDocumentFingerprintType = "document_fingerprint"
)

func (r DLPEntryIntegrationListResponseDocumentFingerprintType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseDocumentFingerprintTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseWordList struct {
	ID        string                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                        `json:"enabled,required"`
	Name      string                                      `json:"name,required"`
	Type      DLPEntryIntegrationListResponseWordListType `json:"type,required"`
	UpdatedAt time.Time                                   `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                 `json:"word_list,required"`
	ProfileID string                                      `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationListResponseWordListJSON `json:"-"`
}

// dlpEntryIntegrationListResponseWordListJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationListResponseWordList]
type dlpEntryIntegrationListResponseWordListJSON struct {
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

func (r *DLPEntryIntegrationListResponseWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseWordList) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponseWordListType string

const (
	DLPEntryIntegrationListResponseWordListTypeWordList DLPEntryIntegrationListResponseWordListType = "word_list"
)

func (r DLPEntryIntegrationListResponseWordListType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseWordListTypeWordList:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseType string

const (
	DLPEntryIntegrationListResponseTypeCustom              DLPEntryIntegrationListResponseType = "custom"
	DLPEntryIntegrationListResponseTypePredefined          DLPEntryIntegrationListResponseType = "predefined"
	DLPEntryIntegrationListResponseTypeIntegration         DLPEntryIntegrationListResponseType = "integration"
	DLPEntryIntegrationListResponseTypeExactData           DLPEntryIntegrationListResponseType = "exact_data"
	DLPEntryIntegrationListResponseTypeDocumentFingerprint DLPEntryIntegrationListResponseType = "document_fingerprint"
	DLPEntryIntegrationListResponseTypeWordList            DLPEntryIntegrationListResponseType = "word_list"
)

func (r DLPEntryIntegrationListResponseType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseTypeCustom, DLPEntryIntegrationListResponseTypePredefined, DLPEntryIntegrationListResponseTypeIntegration, DLPEntryIntegrationListResponseTypeExactData, DLPEntryIntegrationListResponseTypeDocumentFingerprint, DLPEntryIntegrationListResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryIntegrationDeleteResponse = interface{}

type DLPEntryIntegrationGetResponse struct {
	ID      string                             `json:"id,required" format:"uuid"`
	Enabled bool                               `json:"enabled,required"`
	Name    string                             `json:"name,required"`
	Type    DLPEntryIntegrationGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationGetResponsePredefinedConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationGetResponsePredefinedVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                        `json:"word_list"`
	JSON     dlpEntryIntegrationGetResponseJSON `json:"-"`
	union    DLPEntryIntegrationGetResponseUnion
}

// dlpEntryIntegrationGetResponseJSON contains the JSON metadata for the struct
// [DLPEntryIntegrationGetResponse]
type dlpEntryIntegrationGetResponseJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpEntryIntegrationGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryIntegrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryIntegrationGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryIntegrationGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryIntegrationGetResponseCustom],
// [DLPEntryIntegrationGetResponsePredefined],
// [DLPEntryIntegrationGetResponseIntegration],
// [DLPEntryIntegrationGetResponseExactData],
// [DLPEntryIntegrationGetResponseDocumentFingerprint],
// [DLPEntryIntegrationGetResponseWordList].
func (r DLPEntryIntegrationGetResponse) AsUnion() DLPEntryIntegrationGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryIntegrationGetResponseCustom],
// [DLPEntryIntegrationGetResponsePredefined],
// [DLPEntryIntegrationGetResponseIntegration],
// [DLPEntryIntegrationGetResponseExactData],
// [DLPEntryIntegrationGetResponseDocumentFingerprint] or
// [DLPEntryIntegrationGetResponseWordList].
type DLPEntryIntegrationGetResponseUnion interface {
	implementsDLPEntryIntegrationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryIntegrationGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationGetResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationGetResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationGetResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationGetResponseExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationGetResponseDocumentFingerprint{}),
			DiscriminatorValue: "document_fingerprint",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryIntegrationGetResponseWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPEntryIntegrationGetResponseCustom struct {
	ID        string                                   `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                `json:"created_at,required" format:"date-time"`
	Enabled   bool                                     `json:"enabled,required"`
	Name      string                                   `json:"name,required"`
	Pattern   Pattern                                  `json:"pattern,required"`
	Type      DLPEntryIntegrationGetResponseCustomType `json:"type,required"`
	UpdatedAt time.Time                                `json:"updated_at,required" format:"date-time"`
	ProfileID string                                   `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationGetResponseCustomJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseCustomJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationGetResponseCustom]
type dlpEntryIntegrationGetResponseCustomJSON struct {
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

func (r *DLPEntryIntegrationGetResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationGetResponseCustom) implementsDLPEntryIntegrationGetResponse() {}

type DLPEntryIntegrationGetResponseCustomType string

const (
	DLPEntryIntegrationGetResponseCustomTypeCustom DLPEntryIntegrationGetResponseCustomType = "custom"
)

func (r DLPEntryIntegrationGetResponseCustomType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponsePredefined struct {
	ID         string                                             `json:"id,required" format:"uuid"`
	Confidence DLPEntryIntegrationGetResponsePredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                               `json:"enabled,required"`
	Name       string                                             `json:"name,required"`
	Type       DLPEntryIntegrationGetResponsePredefinedType       `json:"type,required"`
	ProfileID  string                                             `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryIntegrationGetResponsePredefinedVariant    `json:"variant"`
	JSON       dlpEntryIntegrationGetResponsePredefinedJSON       `json:"-"`
}

// dlpEntryIntegrationGetResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationGetResponsePredefined]
type dlpEntryIntegrationGetResponsePredefinedJSON struct {
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

func (r *DLPEntryIntegrationGetResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationGetResponsePredefined) implementsDLPEntryIntegrationGetResponse() {}

type DLPEntryIntegrationGetResponsePredefinedConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                   `json:"available,required"`
	JSON      dlpEntryIntegrationGetResponsePredefinedConfidenceJSON `json:"-"`
}

// dlpEntryIntegrationGetResponsePredefinedConfidenceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationGetResponsePredefinedConfidence]
type dlpEntryIntegrationGetResponsePredefinedConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponsePredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponsePredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationGetResponsePredefinedType string

const (
	DLPEntryIntegrationGetResponsePredefinedTypePredefined DLPEntryIntegrationGetResponsePredefinedType = "predefined"
)

func (r DLPEntryIntegrationGetResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponsePredefinedVariant struct {
	TopicType   DLPEntryIntegrationGetResponsePredefinedVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryIntegrationGetResponsePredefinedVariantType      `json:"type,required"`
	Description string                                                   `json:"description,nullable"`
	JSON        dlpEntryIntegrationGetResponsePredefinedVariantJSON      `json:"-"`
}

// dlpEntryIntegrationGetResponsePredefinedVariantJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationGetResponsePredefinedVariant]
type dlpEntryIntegrationGetResponsePredefinedVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponsePredefinedVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponsePredefinedVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationGetResponsePredefinedVariantTopicType string

const (
	DLPEntryIntegrationGetResponsePredefinedVariantTopicTypeIntent  DLPEntryIntegrationGetResponsePredefinedVariantTopicType = "Intent"
	DLPEntryIntegrationGetResponsePredefinedVariantTopicTypeContent DLPEntryIntegrationGetResponsePredefinedVariantTopicType = "Content"
)

func (r DLPEntryIntegrationGetResponsePredefinedVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponsePredefinedVariantTopicTypeIntent, DLPEntryIntegrationGetResponsePredefinedVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponsePredefinedVariantType string

const (
	DLPEntryIntegrationGetResponsePredefinedVariantTypePromptTopic DLPEntryIntegrationGetResponsePredefinedVariantType = "PromptTopic"
)

func (r DLPEntryIntegrationGetResponsePredefinedVariantType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponsePredefinedVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponseIntegration struct {
	ID        string                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                          `json:"enabled,required"`
	Name      string                                        `json:"name,required"`
	Type      DLPEntryIntegrationGetResponseIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationGetResponseIntegrationJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationGetResponseIntegration]
type dlpEntryIntegrationGetResponseIntegrationJSON struct {
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

func (r *DLPEntryIntegrationGetResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationGetResponseIntegration) implementsDLPEntryIntegrationGetResponse() {}

type DLPEntryIntegrationGetResponseIntegrationType string

const (
	DLPEntryIntegrationGetResponseIntegrationTypeIntegration DLPEntryIntegrationGetResponseIntegrationType = "integration"
)

func (r DLPEntryIntegrationGetResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponseExactData struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                        `json:"case_sensitive,required"`
	CreatedAt     time.Time                                   `json:"created_at,required" format:"date-time"`
	Enabled       bool                                        `json:"enabled,required"`
	Name          string                                      `json:"name,required"`
	Secret        bool                                        `json:"secret,required"`
	Type          DLPEntryIntegrationGetResponseExactDataType `json:"type,required"`
	UpdatedAt     time.Time                                   `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryIntegrationGetResponseExactDataJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseExactDataJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationGetResponseExactData]
type dlpEntryIntegrationGetResponseExactDataJSON struct {
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

func (r *DLPEntryIntegrationGetResponseExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationGetResponseExactData) implementsDLPEntryIntegrationGetResponse() {}

type DLPEntryIntegrationGetResponseExactDataType string

const (
	DLPEntryIntegrationGetResponseExactDataTypeExactData DLPEntryIntegrationGetResponseExactDataType = "exact_data"
)

func (r DLPEntryIntegrationGetResponseExactDataType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseExactDataTypeExactData:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponseDocumentFingerprint struct {
	ID        string                                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                  `json:"enabled,required"`
	Name      string                                                `json:"name,required"`
	Type      DLPEntryIntegrationGetResponseDocumentFingerprintType `json:"type,required"`
	UpdatedAt time.Time                                             `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryIntegrationGetResponseDocumentFingerprintJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseDocumentFingerprintJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationGetResponseDocumentFingerprint]
type dlpEntryIntegrationGetResponseDocumentFingerprintJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponseDocumentFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseDocumentFingerprintJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationGetResponseDocumentFingerprint) implementsDLPEntryIntegrationGetResponse() {
}

type DLPEntryIntegrationGetResponseDocumentFingerprintType string

const (
	DLPEntryIntegrationGetResponseDocumentFingerprintTypeDocumentFingerprint DLPEntryIntegrationGetResponseDocumentFingerprintType = "document_fingerprint"
)

func (r DLPEntryIntegrationGetResponseDocumentFingerprintType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseDocumentFingerprintTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponseWordList struct {
	ID        string                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                       `json:"enabled,required"`
	Name      string                                     `json:"name,required"`
	Type      DLPEntryIntegrationGetResponseWordListType `json:"type,required"`
	UpdatedAt time.Time                                  `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                `json:"word_list,required"`
	ProfileID string                                     `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryIntegrationGetResponseWordListJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseWordListJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationGetResponseWordList]
type dlpEntryIntegrationGetResponseWordListJSON struct {
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

func (r *DLPEntryIntegrationGetResponseWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationGetResponseWordList) implementsDLPEntryIntegrationGetResponse() {}

type DLPEntryIntegrationGetResponseWordListType string

const (
	DLPEntryIntegrationGetResponseWordListTypeWordList DLPEntryIntegrationGetResponseWordListType = "word_list"
)

func (r DLPEntryIntegrationGetResponseWordListType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseWordListTypeWordList:
		return true
	}
	return false
}

type DLPEntryIntegrationGetResponseType string

const (
	DLPEntryIntegrationGetResponseTypeCustom              DLPEntryIntegrationGetResponseType = "custom"
	DLPEntryIntegrationGetResponseTypePredefined          DLPEntryIntegrationGetResponseType = "predefined"
	DLPEntryIntegrationGetResponseTypeIntegration         DLPEntryIntegrationGetResponseType = "integration"
	DLPEntryIntegrationGetResponseTypeExactData           DLPEntryIntegrationGetResponseType = "exact_data"
	DLPEntryIntegrationGetResponseTypeDocumentFingerprint DLPEntryIntegrationGetResponseType = "document_fingerprint"
	DLPEntryIntegrationGetResponseTypeWordList            DLPEntryIntegrationGetResponseType = "word_list"
)

func (r DLPEntryIntegrationGetResponseType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseTypeCustom, DLPEntryIntegrationGetResponseTypePredefined, DLPEntryIntegrationGetResponseTypeIntegration, DLPEntryIntegrationGetResponseTypeExactData, DLPEntryIntegrationGetResponseTypeDocumentFingerprint, DLPEntryIntegrationGetResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryIntegrationNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
	EntryID   param.Field[string] `json:"entry_id,required" format:"uuid"`
	// This field is not used as the owning profile. For predefined entries it is
	// already set to a predefined profile.
	ProfileID param.Field[string] `json:"profile_id" format:"uuid"`
}

func (r DLPEntryIntegrationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryIntegrationNewResponseEnvelope struct {
	Errors   []DLPEntryIntegrationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryIntegrationNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryIntegrationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryIntegrationNewResponse                `json:"result"`
	JSON    dlpEntryIntegrationNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationNewResponseEnvelope]
type dlpEntryIntegrationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPEntryIntegrationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryIntegrationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationNewResponseEnvelopeErrors]
type dlpEntryIntegrationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationNewResponseEnvelopeErrorsSource]
type dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPEntryIntegrationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryIntegrationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationNewResponseEnvelopeMessages]
type dlpEntryIntegrationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationNewResponseEnvelopeMessagesSource]
type dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryIntegrationNewResponseEnvelopeSuccess bool

const (
	DLPEntryIntegrationNewResponseEnvelopeSuccessTrue DLPEntryIntegrationNewResponseEnvelopeSuccess = true
)

func (r DLPEntryIntegrationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryIntegrationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
}

func (r DLPEntryIntegrationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryIntegrationUpdateResponseEnvelope struct {
	Errors   []DLPEntryIntegrationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryIntegrationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryIntegrationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryIntegrationUpdateResponse                `json:"result"`
	JSON    dlpEntryIntegrationUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationUpdateResponseEnvelope]
type dlpEntryIntegrationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationUpdateResponseEnvelopeErrors]
type dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource]
type dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationUpdateResponseEnvelopeMessages]
type dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource]
type dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryIntegrationUpdateResponseEnvelopeSuccess bool

const (
	DLPEntryIntegrationUpdateResponseEnvelopeSuccessTrue DLPEntryIntegrationUpdateResponseEnvelopeSuccess = true
)

func (r DLPEntryIntegrationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryIntegrationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryIntegrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryIntegrationDeleteResponseEnvelope struct {
	Errors   []DLPEntryIntegrationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryIntegrationDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryIntegrationDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryIntegrationDeleteResponse                `json:"result,nullable"`
	JSON    dlpEntryIntegrationDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationDeleteResponseEnvelope]
type dlpEntryIntegrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationDeleteResponseEnvelopeErrors]
type dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource]
type dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationDeleteResponseEnvelopeMessages]
type dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource]
type dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryIntegrationDeleteResponseEnvelopeSuccess bool

const (
	DLPEntryIntegrationDeleteResponseEnvelopeSuccessTrue DLPEntryIntegrationDeleteResponseEnvelopeSuccess = true
)

func (r DLPEntryIntegrationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryIntegrationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryIntegrationGetResponseEnvelope struct {
	Errors   []DLPEntryIntegrationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryIntegrationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryIntegrationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryIntegrationGetResponse                `json:"result"`
	JSON    dlpEntryIntegrationGetResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryIntegrationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationGetResponseEnvelope]
type dlpEntryIntegrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPEntryIntegrationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryIntegrationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryIntegrationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationGetResponseEnvelopeErrors]
type dlpEntryIntegrationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpEntryIntegrationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationGetResponseEnvelopeErrorsSource]
type dlpEntryIntegrationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPEntryIntegrationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryIntegrationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryIntegrationGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationGetResponseEnvelopeMessages]
type dlpEntryIntegrationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpEntryIntegrationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationGetResponseEnvelopeMessagesSource]
type dlpEntryIntegrationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryIntegrationGetResponseEnvelopeSuccess bool

const (
	DLPEntryIntegrationGetResponseEnvelopeSuccessTrue DLPEntryIntegrationGetResponseEnvelopeSuccess = true
)

func (r DLPEntryIntegrationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
