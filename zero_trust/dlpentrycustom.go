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

// DLPEntryCustomService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEntryCustomService] method instead.
type DLPEntryCustomService struct {
	Options []option.RequestOption
}

// NewDLPEntryCustomService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPEntryCustomService(opts ...option.RequestOption) (r *DLPEntryCustomService) {
	r = &DLPEntryCustomService{}
	r.Options = opts
	return
}

// Creates a DLP custom entry.
func (r *DLPEntryCustomService) New(ctx context.Context, params DLPEntryCustomNewParams, opts ...option.RequestOption) (res *DLPEntryCustomNewResponse, err error) {
	var env DLPEntryCustomNewResponseEnvelope
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

// Updates a DLP custom entry.
func (r *DLPEntryCustomService) Update(ctx context.Context, entryID string, params DLPEntryCustomUpdateParams, opts ...option.RequestOption) (res *DLPEntryCustomUpdateResponse, err error) {
	var env DLPEntryCustomUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/custom/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all DLP entries in an account.
func (r *DLPEntryCustomService) List(ctx context.Context, query DLPEntryCustomListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPEntryCustomListResponse], err error) {
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
func (r *DLPEntryCustomService) ListAutoPaging(ctx context.Context, query DLPEntryCustomListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPEntryCustomListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a DLP custom entry.
func (r *DLPEntryCustomService) Delete(ctx context.Context, entryID string, body DLPEntryCustomDeleteParams, opts ...option.RequestOption) (res *DLPEntryCustomDeleteResponse, err error) {
	var env DLPEntryCustomDeleteResponseEnvelope
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
func (r *DLPEntryCustomService) Get(ctx context.Context, entryID string, query DLPEntryCustomGetParams, opts ...option.RequestOption) (res *DLPEntryCustomGetResponse, err error) {
	var env DLPEntryCustomGetResponseEnvelope
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

type DLPEntryCustomNewResponse struct {
	ID        string                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                          `json:"enabled,required"`
	Name      string                        `json:"name,required"`
	Pattern   Pattern                       `json:"pattern,required"`
	UpdatedAt time.Time                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryCustomNewResponseJSON `json:"-"`
}

// dlpEntryCustomNewResponseJSON contains the JSON metadata for the struct
// [DLPEntryCustomNewResponse]
type dlpEntryCustomNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomUpdateResponse struct {
	ID        string                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                             `json:"enabled,required"`
	Name      string                           `json:"name,required"`
	Pattern   Pattern                          `json:"pattern,required"`
	UpdatedAt time.Time                        `json:"updated_at,required" format:"date-time"`
	ProfileID string                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryCustomUpdateResponseJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryCustomUpdateResponse]
type dlpEntryCustomUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomListResponse struct {
	ID      string                         `json:"id,required" format:"uuid"`
	Enabled bool                           `json:"enabled,required"`
	Name    string                         `json:"name,required"`
	Type    DLPEntryCustomListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryCustomListResponsePredefinedConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryCustomListResponsePredefinedVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                    `json:"word_list"`
	JSON     dlpEntryCustomListResponseJSON `json:"-"`
	union    DLPEntryCustomListResponseUnion
}

// dlpEntryCustomListResponseJSON contains the JSON metadata for the struct
// [DLPEntryCustomListResponse]
type dlpEntryCustomListResponseJSON struct {
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

func (r dlpEntryCustomListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryCustomListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryCustomListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryCustomListResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryCustomListResponseCustom],
// [DLPEntryCustomListResponsePredefined], [DLPEntryCustomListResponseIntegration],
// [DLPEntryCustomListResponseExactData],
// [DLPEntryCustomListResponseDocumentFingerprint],
// [DLPEntryCustomListResponseWordList].
func (r DLPEntryCustomListResponse) AsUnion() DLPEntryCustomListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryCustomListResponseCustom],
// [DLPEntryCustomListResponsePredefined], [DLPEntryCustomListResponseIntegration],
// [DLPEntryCustomListResponseExactData],
// [DLPEntryCustomListResponseDocumentFingerprint] or
// [DLPEntryCustomListResponseWordList].
type DLPEntryCustomListResponseUnion interface {
	implementsDLPEntryCustomListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryCustomListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryCustomListResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryCustomListResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryCustomListResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryCustomListResponseExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryCustomListResponseDocumentFingerprint{}),
			DiscriminatorValue: "document_fingerprint",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryCustomListResponseWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPEntryCustomListResponseCustom struct {
	ID        string                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                 `json:"enabled,required"`
	Name      string                               `json:"name,required"`
	Pattern   Pattern                              `json:"pattern,required"`
	Type      DLPEntryCustomListResponseCustomType `json:"type,required"`
	UpdatedAt time.Time                            `json:"updated_at,required" format:"date-time"`
	ProfileID string                               `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryCustomListResponseCustomJSON `json:"-"`
}

// dlpEntryCustomListResponseCustomJSON contains the JSON metadata for the struct
// [DLPEntryCustomListResponseCustom]
type dlpEntryCustomListResponseCustomJSON struct {
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

func (r *DLPEntryCustomListResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseCustom) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseCustomType string

const (
	DLPEntryCustomListResponseCustomTypeCustom DLPEntryCustomListResponseCustomType = "custom"
)

func (r DLPEntryCustomListResponseCustomType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryCustomListResponsePredefined struct {
	ID         string                                         `json:"id,required" format:"uuid"`
	Confidence DLPEntryCustomListResponsePredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                           `json:"enabled,required"`
	Name       string                                         `json:"name,required"`
	Type       DLPEntryCustomListResponsePredefinedType       `json:"type,required"`
	ProfileID  string                                         `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryCustomListResponsePredefinedVariant    `json:"variant"`
	JSON       dlpEntryCustomListResponsePredefinedJSON       `json:"-"`
}

// dlpEntryCustomListResponsePredefinedJSON contains the JSON metadata for the
// struct [DLPEntryCustomListResponsePredefined]
type dlpEntryCustomListResponsePredefinedJSON struct {
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

func (r *DLPEntryCustomListResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponsePredefined) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponsePredefinedConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                               `json:"available,required"`
	JSON      dlpEntryCustomListResponsePredefinedConfidenceJSON `json:"-"`
}

// dlpEntryCustomListResponsePredefinedConfidenceJSON contains the JSON metadata
// for the struct [DLPEntryCustomListResponsePredefinedConfidence]
type dlpEntryCustomListResponsePredefinedConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryCustomListResponsePredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponsePredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomListResponsePredefinedType string

const (
	DLPEntryCustomListResponsePredefinedTypePredefined DLPEntryCustomListResponsePredefinedType = "predefined"
)

func (r DLPEntryCustomListResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryCustomListResponsePredefinedVariant struct {
	TopicType   DLPEntryCustomListResponsePredefinedVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryCustomListResponsePredefinedVariantType      `json:"type,required"`
	Description string                                               `json:"description,nullable"`
	JSON        dlpEntryCustomListResponsePredefinedVariantJSON      `json:"-"`
}

// dlpEntryCustomListResponsePredefinedVariantJSON contains the JSON metadata for
// the struct [DLPEntryCustomListResponsePredefinedVariant]
type dlpEntryCustomListResponsePredefinedVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomListResponsePredefinedVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponsePredefinedVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomListResponsePredefinedVariantTopicType string

const (
	DLPEntryCustomListResponsePredefinedVariantTopicTypeIntent  DLPEntryCustomListResponsePredefinedVariantTopicType = "Intent"
	DLPEntryCustomListResponsePredefinedVariantTopicTypeContent DLPEntryCustomListResponsePredefinedVariantTopicType = "Content"
)

func (r DLPEntryCustomListResponsePredefinedVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponsePredefinedVariantTopicTypeIntent, DLPEntryCustomListResponsePredefinedVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryCustomListResponsePredefinedVariantType string

const (
	DLPEntryCustomListResponsePredefinedVariantTypePromptTopic DLPEntryCustomListResponsePredefinedVariantType = "PromptTopic"
)

func (r DLPEntryCustomListResponsePredefinedVariantType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponsePredefinedVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryCustomListResponseIntegration struct {
	ID        string                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                      `json:"enabled,required"`
	Name      string                                    `json:"name,required"`
	Type      DLPEntryCustomListResponseIntegrationType `json:"type,required"`
	UpdatedAt time.Time                                 `json:"updated_at,required" format:"date-time"`
	ProfileID string                                    `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryCustomListResponseIntegrationJSON `json:"-"`
}

// dlpEntryCustomListResponseIntegrationJSON contains the JSON metadata for the
// struct [DLPEntryCustomListResponseIntegration]
type dlpEntryCustomListResponseIntegrationJSON struct {
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

func (r *DLPEntryCustomListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseIntegration) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseIntegrationType string

const (
	DLPEntryCustomListResponseIntegrationTypeIntegration DLPEntryCustomListResponseIntegrationType = "integration"
)

func (r DLPEntryCustomListResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryCustomListResponseExactData struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                    `json:"case_sensitive,required"`
	CreatedAt     time.Time                               `json:"created_at,required" format:"date-time"`
	Enabled       bool                                    `json:"enabled,required"`
	Name          string                                  `json:"name,required"`
	Secret        bool                                    `json:"secret,required"`
	Type          DLPEntryCustomListResponseExactDataType `json:"type,required"`
	UpdatedAt     time.Time                               `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryCustomListResponseExactDataJSON `json:"-"`
}

// dlpEntryCustomListResponseExactDataJSON contains the JSON metadata for the
// struct [DLPEntryCustomListResponseExactData]
type dlpEntryCustomListResponseExactDataJSON struct {
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

func (r *DLPEntryCustomListResponseExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseExactData) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseExactDataType string

const (
	DLPEntryCustomListResponseExactDataTypeExactData DLPEntryCustomListResponseExactDataType = "exact_data"
)

func (r DLPEntryCustomListResponseExactDataType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseExactDataTypeExactData:
		return true
	}
	return false
}

type DLPEntryCustomListResponseDocumentFingerprint struct {
	ID        string                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                              `json:"enabled,required"`
	Name      string                                            `json:"name,required"`
	Type      DLPEntryCustomListResponseDocumentFingerprintType `json:"type,required"`
	UpdatedAt time.Time                                         `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryCustomListResponseDocumentFingerprintJSON `json:"-"`
}

// dlpEntryCustomListResponseDocumentFingerprintJSON contains the JSON metadata for
// the struct [DLPEntryCustomListResponseDocumentFingerprint]
type dlpEntryCustomListResponseDocumentFingerprintJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomListResponseDocumentFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseDocumentFingerprintJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseDocumentFingerprint) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseDocumentFingerprintType string

const (
	DLPEntryCustomListResponseDocumentFingerprintTypeDocumentFingerprint DLPEntryCustomListResponseDocumentFingerprintType = "document_fingerprint"
)

func (r DLPEntryCustomListResponseDocumentFingerprintType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseDocumentFingerprintTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryCustomListResponseWordList struct {
	ID        string                                 `json:"id,required" format:"uuid"`
	CreatedAt time.Time                              `json:"created_at,required" format:"date-time"`
	Enabled   bool                                   `json:"enabled,required"`
	Name      string                                 `json:"name,required"`
	Type      DLPEntryCustomListResponseWordListType `json:"type,required"`
	UpdatedAt time.Time                              `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                            `json:"word_list,required"`
	ProfileID string                                 `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryCustomListResponseWordListJSON `json:"-"`
}

// dlpEntryCustomListResponseWordListJSON contains the JSON metadata for the struct
// [DLPEntryCustomListResponseWordList]
type dlpEntryCustomListResponseWordListJSON struct {
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

func (r *DLPEntryCustomListResponseWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseWordList) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseWordListType string

const (
	DLPEntryCustomListResponseWordListTypeWordList DLPEntryCustomListResponseWordListType = "word_list"
)

func (r DLPEntryCustomListResponseWordListType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseWordListTypeWordList:
		return true
	}
	return false
}

type DLPEntryCustomListResponseType string

const (
	DLPEntryCustomListResponseTypeCustom              DLPEntryCustomListResponseType = "custom"
	DLPEntryCustomListResponseTypePredefined          DLPEntryCustomListResponseType = "predefined"
	DLPEntryCustomListResponseTypeIntegration         DLPEntryCustomListResponseType = "integration"
	DLPEntryCustomListResponseTypeExactData           DLPEntryCustomListResponseType = "exact_data"
	DLPEntryCustomListResponseTypeDocumentFingerprint DLPEntryCustomListResponseType = "document_fingerprint"
	DLPEntryCustomListResponseTypeWordList            DLPEntryCustomListResponseType = "word_list"
)

func (r DLPEntryCustomListResponseType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseTypeCustom, DLPEntryCustomListResponseTypePredefined, DLPEntryCustomListResponseTypeIntegration, DLPEntryCustomListResponseTypeExactData, DLPEntryCustomListResponseTypeDocumentFingerprint, DLPEntryCustomListResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryCustomDeleteResponse = interface{}

type DLPEntryCustomGetResponse struct {
	ID      string                        `json:"id,required" format:"uuid"`
	Enabled bool                          `json:"enabled,required"`
	Name    string                        `json:"name,required"`
	Type    DLPEntryCustomGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryCustomGetResponsePredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	// This field can have the runtime type of
	// [[]DLPEntryCustomGetResponseCustomEntryProfile],
	// [[]DLPEntryCustomGetResponsePredefinedEntryProfile],
	// [[]DLPEntryCustomGetResponseIntegrationEntryProfile],
	// [[]DLPEntryCustomGetResponseExactDataEntryProfile],
	// [[]DLPEntryCustomGetResponseDocumentFingerprintEntryProfile],
	// [[]DLPEntryCustomGetResponseWordListEntryProfile].
	Profiles  interface{} `json:"profiles"`
	Secret    bool        `json:"secret"`
	UpdatedAt time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryCustomGetResponsePredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                   `json:"word_list"`
	JSON     dlpEntryCustomGetResponseJSON `json:"-"`
	union    DLPEntryCustomGetResponseUnion
}

// dlpEntryCustomGetResponseJSON contains the JSON metadata for the struct
// [DLPEntryCustomGetResponse]
type dlpEntryCustomGetResponseJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CaseSensitive apijson.Field
	Confidence    apijson.Field
	CreatedAt     apijson.Field
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Profiles      apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
	Variant       apijson.Field
	WordList      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dlpEntryCustomGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryCustomGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryCustomGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryCustomGetResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [DLPEntryCustomGetResponseCustomEntry],
// [DLPEntryCustomGetResponsePredefinedEntry],
// [DLPEntryCustomGetResponseIntegrationEntry],
// [DLPEntryCustomGetResponseExactDataEntry],
// [DLPEntryCustomGetResponseDocumentFingerprintEntry],
// [DLPEntryCustomGetResponseWordListEntry].
func (r DLPEntryCustomGetResponse) AsUnion() DLPEntryCustomGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryCustomGetResponseCustomEntry],
// [DLPEntryCustomGetResponsePredefinedEntry],
// [DLPEntryCustomGetResponseIntegrationEntry],
// [DLPEntryCustomGetResponseExactDataEntry],
// [DLPEntryCustomGetResponseDocumentFingerprintEntry] or
// [DLPEntryCustomGetResponseWordListEntry].
type DLPEntryCustomGetResponseUnion interface {
	implementsDLPEntryCustomGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryCustomGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseWordListEntry{}),
		},
	)
}

type DLPEntryCustomGetResponseCustomEntry struct {
	ID        string                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                          `json:"enabled,required"`
	Name      string                                        `json:"name,required"`
	Pattern   Pattern                                       `json:"pattern,required"`
	Type      DLPEntryCustomGetResponseCustomEntryType      `json:"type,required"`
	UpdatedAt time.Time                                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                                        `json:"profile_id,nullable" format:"uuid"`
	Profiles  []DLPEntryCustomGetResponseCustomEntryProfile `json:"profiles"`
	JSON      dlpEntryCustomGetResponseCustomEntryJSON      `json:"-"`
}

// dlpEntryCustomGetResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponseCustomEntry]
type dlpEntryCustomGetResponseCustomEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomGetResponseCustomEntry) implementsDLPEntryCustomGetResponse() {}

type DLPEntryCustomGetResponseCustomEntryType string

const (
	DLPEntryCustomGetResponseCustomEntryTypeCustom DLPEntryCustomGetResponseCustomEntryType = "custom"
)

func (r DLPEntryCustomGetResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryCustomGetResponseCustomEntryProfile struct {
	ID   string                                          `json:"id,required" format:"uuid"`
	Name string                                          `json:"name,required"`
	JSON dlpEntryCustomGetResponseCustomEntryProfileJSON `json:"-"`
}

// dlpEntryCustomGetResponseCustomEntryProfileJSON contains the JSON metadata for
// the struct [DLPEntryCustomGetResponseCustomEntryProfile]
type dlpEntryCustomGetResponseCustomEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseCustomEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseCustomEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponsePredefinedEntry struct {
	ID         string                                             `json:"id,required" format:"uuid"`
	Confidence DLPEntryCustomGetResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                               `json:"enabled,required"`
	Name       string                                             `json:"name,required"`
	Type       DLPEntryCustomGetResponsePredefinedEntryType       `json:"type,required"`
	ProfileID  string                                             `json:"profile_id,nullable" format:"uuid"`
	Profiles   []DLPEntryCustomGetResponsePredefinedEntryProfile  `json:"profiles"`
	Variant    DLPEntryCustomGetResponsePredefinedEntryVariant    `json:"variant"`
	JSON       dlpEntryCustomGetResponsePredefinedEntryJSON       `json:"-"`
}

// dlpEntryCustomGetResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponsePredefinedEntry]
type dlpEntryCustomGetResponsePredefinedEntryJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ProfileID   apijson.Field
	Profiles    apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomGetResponsePredefinedEntry) implementsDLPEntryCustomGetResponse() {}

type DLPEntryCustomGetResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                   `json:"available,required"`
	JSON      dlpEntryCustomGetResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryCustomGetResponsePredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct [DLPEntryCustomGetResponsePredefinedEntryConfidence]
type dlpEntryCustomGetResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponsePredefinedEntryType string

const (
	DLPEntryCustomGetResponsePredefinedEntryTypePredefined DLPEntryCustomGetResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryCustomGetResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryCustomGetResponsePredefinedEntryProfile struct {
	ID   string                                              `json:"id,required" format:"uuid"`
	Name string                                              `json:"name,required"`
	JSON dlpEntryCustomGetResponsePredefinedEntryProfileJSON `json:"-"`
}

// dlpEntryCustomGetResponsePredefinedEntryProfileJSON contains the JSON metadata
// for the struct [DLPEntryCustomGetResponsePredefinedEntryProfile]
type dlpEntryCustomGetResponsePredefinedEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponsePredefinedEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponsePredefinedEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryCustomGetResponsePredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryCustomGetResponsePredefinedEntryVariantType      `json:"type,required"`
	Description string                                                   `json:"description,nullable"`
	JSON        dlpEntryCustomGetResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryCustomGetResponsePredefinedEntryVariantJSON contains the JSON metadata
// for the struct [DLPEntryCustomGetResponsePredefinedEntryVariant]
type dlpEntryCustomGetResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryCustomGetResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryCustomGetResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryCustomGetResponsePredefinedEntryVariantTopicTypeContent DLPEntryCustomGetResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryCustomGetResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryCustomGetResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryCustomGetResponsePredefinedEntryVariantType string

const (
	DLPEntryCustomGetResponsePredefinedEntryVariantTypePromptTopic DLPEntryCustomGetResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryCustomGetResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryCustomGetResponseIntegrationEntry struct {
	ID        string                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                               `json:"enabled,required"`
	Name      string                                             `json:"name,required"`
	Type      DLPEntryCustomGetResponseIntegrationEntryType      `json:"type,required"`
	UpdatedAt time.Time                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                             `json:"profile_id,nullable" format:"uuid"`
	Profiles  []DLPEntryCustomGetResponseIntegrationEntryProfile `json:"profiles"`
	JSON      dlpEntryCustomGetResponseIntegrationEntryJSON      `json:"-"`
}

// dlpEntryCustomGetResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponseIntegrationEntry]
type dlpEntryCustomGetResponseIntegrationEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	ProfileID   apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomGetResponseIntegrationEntry) implementsDLPEntryCustomGetResponse() {}

type DLPEntryCustomGetResponseIntegrationEntryType string

const (
	DLPEntryCustomGetResponseIntegrationEntryTypeIntegration DLPEntryCustomGetResponseIntegrationEntryType = "integration"
)

func (r DLPEntryCustomGetResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryCustomGetResponseIntegrationEntryProfile struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON dlpEntryCustomGetResponseIntegrationEntryProfileJSON `json:"-"`
}

// dlpEntryCustomGetResponseIntegrationEntryProfileJSON contains the JSON metadata
// for the struct [DLPEntryCustomGetResponseIntegrationEntryProfile]
type dlpEntryCustomGetResponseIntegrationEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseIntegrationEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseIntegrationEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                             `json:"case_sensitive,required"`
	CreatedAt     time.Time                                        `json:"created_at,required" format:"date-time"`
	Enabled       bool                                             `json:"enabled,required"`
	Name          string                                           `json:"name,required"`
	Secret        bool                                             `json:"secret,required"`
	Type          DLPEntryCustomGetResponseExactDataEntryType      `json:"type,required"`
	UpdatedAt     time.Time                                        `json:"updated_at,required" format:"date-time"`
	Profiles      []DLPEntryCustomGetResponseExactDataEntryProfile `json:"profiles"`
	JSON          dlpEntryCustomGetResponseExactDataEntryJSON      `json:"-"`
}

// dlpEntryCustomGetResponseExactDataEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponseExactDataEntry]
type dlpEntryCustomGetResponseExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	Profiles      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomGetResponseExactDataEntry) implementsDLPEntryCustomGetResponse() {}

type DLPEntryCustomGetResponseExactDataEntryType string

const (
	DLPEntryCustomGetResponseExactDataEntryTypeExactData DLPEntryCustomGetResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryCustomGetResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryCustomGetResponseExactDataEntryProfile struct {
	ID   string                                             `json:"id,required" format:"uuid"`
	Name string                                             `json:"name,required"`
	JSON dlpEntryCustomGetResponseExactDataEntryProfileJSON `json:"-"`
}

// dlpEntryCustomGetResponseExactDataEntryProfileJSON contains the JSON metadata
// for the struct [DLPEntryCustomGetResponseExactDataEntryProfile]
type dlpEntryCustomGetResponseExactDataEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseExactDataEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseExactDataEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseDocumentFingerprintEntry struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPEntryCustomGetResponseDocumentFingerprintEntryType      `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	Profiles  []DLPEntryCustomGetResponseDocumentFingerprintEntryProfile `json:"profiles"`
	JSON      dlpEntryCustomGetResponseDocumentFingerprintEntryJSON      `json:"-"`
}

// dlpEntryCustomGetResponseDocumentFingerprintEntryJSON contains the JSON metadata
// for the struct [DLPEntryCustomGetResponseDocumentFingerprintEntry]
type dlpEntryCustomGetResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomGetResponseDocumentFingerprintEntry) implementsDLPEntryCustomGetResponse() {}

type DLPEntryCustomGetResponseDocumentFingerprintEntryType string

const (
	DLPEntryCustomGetResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryCustomGetResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryCustomGetResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryCustomGetResponseDocumentFingerprintEntryProfile struct {
	ID   string                                                       `json:"id,required" format:"uuid"`
	Name string                                                       `json:"name,required"`
	JSON dlpEntryCustomGetResponseDocumentFingerprintEntryProfileJSON `json:"-"`
}

// dlpEntryCustomGetResponseDocumentFingerprintEntryProfileJSON contains the JSON
// metadata for the struct
// [DLPEntryCustomGetResponseDocumentFingerprintEntryProfile]
type dlpEntryCustomGetResponseDocumentFingerprintEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseDocumentFingerprintEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseDocumentFingerprintEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseWordListEntry struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Type      DLPEntryCustomGetResponseWordListEntryType      `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                     `json:"word_list,required"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	Profiles  []DLPEntryCustomGetResponseWordListEntryProfile `json:"profiles"`
	JSON      dlpEntryCustomGetResponseWordListEntryJSON      `json:"-"`
}

// dlpEntryCustomGetResponseWordListEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponseWordListEntry]
type dlpEntryCustomGetResponseWordListEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	WordList    apijson.Field
	ProfileID   apijson.Field
	Profiles    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomGetResponseWordListEntry) implementsDLPEntryCustomGetResponse() {}

type DLPEntryCustomGetResponseWordListEntryType string

const (
	DLPEntryCustomGetResponseWordListEntryTypeWordList DLPEntryCustomGetResponseWordListEntryType = "word_list"
)

func (r DLPEntryCustomGetResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryCustomGetResponseWordListEntryProfile struct {
	ID   string                                            `json:"id,required" format:"uuid"`
	Name string                                            `json:"name,required"`
	JSON dlpEntryCustomGetResponseWordListEntryProfileJSON `json:"-"`
}

// dlpEntryCustomGetResponseWordListEntryProfileJSON contains the JSON metadata for
// the struct [DLPEntryCustomGetResponseWordListEntryProfile]
type dlpEntryCustomGetResponseWordListEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseWordListEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseWordListEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseType string

const (
	DLPEntryCustomGetResponseTypeCustom              DLPEntryCustomGetResponseType = "custom"
	DLPEntryCustomGetResponseTypePredefined          DLPEntryCustomGetResponseType = "predefined"
	DLPEntryCustomGetResponseTypeIntegration         DLPEntryCustomGetResponseType = "integration"
	DLPEntryCustomGetResponseTypeExactData           DLPEntryCustomGetResponseType = "exact_data"
	DLPEntryCustomGetResponseTypeDocumentFingerprint DLPEntryCustomGetResponseType = "document_fingerprint"
	DLPEntryCustomGetResponseTypeWordList            DLPEntryCustomGetResponseType = "word_list"
)

func (r DLPEntryCustomGetResponseType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseTypeCustom, DLPEntryCustomGetResponseTypePredefined, DLPEntryCustomGetResponseTypeIntegration, DLPEntryCustomGetResponseTypeExactData, DLPEntryCustomGetResponseTypeDocumentFingerprint, DLPEntryCustomGetResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryCustomNewParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Enabled   param.Field[bool]         `json:"enabled,required"`
	Name      param.Field[string]       `json:"name,required"`
	Pattern   param.Field[PatternParam] `json:"pattern,required"`
	ProfileID param.Field[string]       `json:"profile_id" format:"uuid"`
}

func (r DLPEntryCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryCustomNewResponseEnvelope struct {
	Errors   []DLPEntryCustomNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryCustomNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryCustomNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryCustomNewResponse                `json:"result"`
	JSON    dlpEntryCustomNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryCustomNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryCustomNewResponseEnvelope]
type dlpEntryCustomNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomNewResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DLPEntryCustomNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryCustomNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryCustomNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEntryCustomNewResponseEnvelopeErrors]
type dlpEntryCustomNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomNewResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dlpEntryCustomNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryCustomNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEntryCustomNewResponseEnvelopeErrorsSource]
type dlpEntryCustomNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomNewResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPEntryCustomNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryCustomNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryCustomNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryCustomNewResponseEnvelopeMessages]
type dlpEntryCustomNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomNewResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpEntryCustomNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryCustomNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPEntryCustomNewResponseEnvelopeMessagesSource]
type dlpEntryCustomNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryCustomNewResponseEnvelopeSuccess bool

const (
	DLPEntryCustomNewResponseEnvelopeSuccessTrue DLPEntryCustomNewResponseEnvelopeSuccess = true
)

func (r DLPEntryCustomNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryCustomNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryCustomUpdateParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Enabled   param.Field[bool]         `json:"enabled,required"`
	Name      param.Field[string]       `json:"name,required"`
	Pattern   param.Field[PatternParam] `json:"pattern,required"`
}

func (r DLPEntryCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryCustomUpdateResponseEnvelope struct {
	Errors   []DLPEntryCustomUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryCustomUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryCustomUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryCustomUpdateResponse                `json:"result"`
	JSON    dlpEntryCustomUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryCustomUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryCustomUpdateResponseEnvelope]
type dlpEntryCustomUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomUpdateResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPEntryCustomUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryCustomUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryCustomUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryCustomUpdateResponseEnvelopeErrors]
type dlpEntryCustomUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpEntryCustomUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEntryCustomUpdateResponseEnvelopeErrorsSource]
type dlpEntryCustomUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomUpdateResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPEntryCustomUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryCustomUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryCustomUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEntryCustomUpdateResponseEnvelopeMessages]
type dlpEntryCustomUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpEntryCustomUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryCustomUpdateResponseEnvelopeMessagesSource]
type dlpEntryCustomUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryCustomUpdateResponseEnvelopeSuccess bool

const (
	DLPEntryCustomUpdateResponseEnvelopeSuccessTrue DLPEntryCustomUpdateResponseEnvelopeSuccess = true
)

func (r DLPEntryCustomUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryCustomListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryCustomDeleteResponseEnvelope struct {
	Errors   []DLPEntryCustomDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryCustomDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryCustomDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryCustomDeleteResponse                `json:"result,nullable"`
	JSON    dlpEntryCustomDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryCustomDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryCustomDeleteResponseEnvelope]
type dlpEntryCustomDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomDeleteResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPEntryCustomDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryCustomDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryCustomDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryCustomDeleteResponseEnvelopeErrors]
type dlpEntryCustomDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpEntryCustomDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryCustomDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEntryCustomDeleteResponseEnvelopeErrorsSource]
type dlpEntryCustomDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomDeleteResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPEntryCustomDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryCustomDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryCustomDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEntryCustomDeleteResponseEnvelopeMessages]
type dlpEntryCustomDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpEntryCustomDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryCustomDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryCustomDeleteResponseEnvelopeMessagesSource]
type dlpEntryCustomDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryCustomDeleteResponseEnvelopeSuccess bool

const (
	DLPEntryCustomDeleteResponseEnvelopeSuccessTrue DLPEntryCustomDeleteResponseEnvelopeSuccess = true
)

func (r DLPEntryCustomDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryCustomDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryCustomGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryCustomGetResponseEnvelope struct {
	Errors   []DLPEntryCustomGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryCustomGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryCustomGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryCustomGetResponse                `json:"result"`
	JSON    dlpEntryCustomGetResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryCustomGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPEntryCustomGetResponseEnvelope]
type dlpEntryCustomGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DLPEntryCustomGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryCustomGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryCustomGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponseEnvelopeErrors]
type dlpEntryCustomGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dlpEntryCustomGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryCustomGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPEntryCustomGetResponseEnvelopeErrorsSource]
type dlpEntryCustomGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPEntryCustomGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryCustomGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryCustomGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponseEnvelopeMessages]
type dlpEntryCustomGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpEntryCustomGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryCustomGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPEntryCustomGetResponseEnvelopeMessagesSource]
type dlpEntryCustomGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryCustomGetResponseEnvelopeSuccess bool

const (
	DLPEntryCustomGetResponseEnvelopeSuccessTrue DLPEntryCustomGetResponseEnvelopeSuccess = true
)

func (r DLPEntryCustomGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
