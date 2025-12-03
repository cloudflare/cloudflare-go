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
	ID        string                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time               `json:"created_at,required" format:"date-time"`
	Enabled   bool                    `json:"enabled,required"`
	Name      string                  `json:"name,required"`
	Pattern   Pattern                 `json:"pattern,required"`
	UpdatedAt time.Time               `json:"updated_at,required" format:"date-time"`
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
	ID      string                     `json:"id,required" format:"uuid"`
	Enabled bool                       `json:"enabled,required"`
	Name    string                     `json:"name,required"`
	Type    DLPEntryUpdateResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryUpdateResponsePredefinedConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryUpdateResponsePredefinedVariant].
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
// Possible runtime types of the union are [DLPEntryUpdateResponseCustom],
// [DLPEntryUpdateResponsePredefined], [DLPEntryUpdateResponseIntegration],
// [DLPEntryUpdateResponseExactData], [DLPEntryUpdateResponseDocumentFingerprint],
// [DLPEntryUpdateResponseWordList].
func (r DLPEntryUpdateResponse) AsUnion() DLPEntryUpdateResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryUpdateResponseCustom],
// [DLPEntryUpdateResponsePredefined], [DLPEntryUpdateResponseIntegration],
// [DLPEntryUpdateResponseExactData], [DLPEntryUpdateResponseDocumentFingerprint]
// or [DLPEntryUpdateResponseWordList].
type DLPEntryUpdateResponseUnion interface {
	implementsDLPEntryUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryUpdateResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryUpdateResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryUpdateResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryUpdateResponseExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryUpdateResponseDocumentFingerprint{}),
			DiscriminatorValue: "document_fingerprint",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryUpdateResponseWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPEntryUpdateResponseCustom struct {
	ID        string                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                             `json:"enabled,required"`
	Name      string                           `json:"name,required"`
	Pattern   Pattern                          `json:"pattern,required"`
	Type      DLPEntryUpdateResponseCustomType `json:"type,required"`
	UpdatedAt time.Time                        `json:"updated_at,required" format:"date-time"`
	ProfileID string                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseCustomJSON `json:"-"`
}

// dlpEntryUpdateResponseCustomJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseCustom]
type dlpEntryUpdateResponseCustomJSON struct {
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

func (r *DLPEntryUpdateResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseCustom) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseCustomType string

const (
	DLPEntryUpdateResponseCustomTypeCustom DLPEntryUpdateResponseCustomType = "custom"
)

func (r DLPEntryUpdateResponseCustomType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryUpdateResponsePredefined struct {
	ID         string                                     `json:"id,required" format:"uuid"`
	Confidence DLPEntryUpdateResponsePredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                       `json:"enabled,required"`
	Name       string                                     `json:"name,required"`
	Type       DLPEntryUpdateResponsePredefinedType       `json:"type,required"`
	ProfileID  string                                     `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryUpdateResponsePredefinedVariant    `json:"variant"`
	JSON       dlpEntryUpdateResponsePredefinedJSON       `json:"-"`
}

// dlpEntryUpdateResponsePredefinedJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponsePredefined]
type dlpEntryUpdateResponsePredefinedJSON struct {
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

func (r *DLPEntryUpdateResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponsePredefined) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponsePredefinedConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                           `json:"available,required"`
	JSON      dlpEntryUpdateResponsePredefinedConfidenceJSON `json:"-"`
}

// dlpEntryUpdateResponsePredefinedConfidenceJSON contains the JSON metadata for
// the struct [DLPEntryUpdateResponsePredefinedConfidence]
type dlpEntryUpdateResponsePredefinedConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryUpdateResponsePredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponsePredefinedType string

const (
	DLPEntryUpdateResponsePredefinedTypePredefined DLPEntryUpdateResponsePredefinedType = "predefined"
)

func (r DLPEntryUpdateResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryUpdateResponsePredefinedVariant struct {
	TopicType   DLPEntryUpdateResponsePredefinedVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryUpdateResponsePredefinedVariantType      `json:"type,required"`
	Description string                                           `json:"description,nullable"`
	JSON        dlpEntryUpdateResponsePredefinedVariantJSON      `json:"-"`
}

// dlpEntryUpdateResponsePredefinedVariantJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponsePredefinedVariant]
type dlpEntryUpdateResponsePredefinedVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponsePredefinedVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponsePredefinedVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryUpdateResponsePredefinedVariantTopicType string

const (
	DLPEntryUpdateResponsePredefinedVariantTopicTypeIntent  DLPEntryUpdateResponsePredefinedVariantTopicType = "Intent"
	DLPEntryUpdateResponsePredefinedVariantTopicTypeContent DLPEntryUpdateResponsePredefinedVariantTopicType = "Content"
)

func (r DLPEntryUpdateResponsePredefinedVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponsePredefinedVariantTopicTypeIntent, DLPEntryUpdateResponsePredefinedVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryUpdateResponsePredefinedVariantType string

const (
	DLPEntryUpdateResponsePredefinedVariantTypePromptTopic DLPEntryUpdateResponsePredefinedVariantType = "PromptTopic"
)

func (r DLPEntryUpdateResponsePredefinedVariantType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponsePredefinedVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryUpdateResponseIntegration struct {
	ID        string                                `json:"id,required" format:"uuid"`
	CreatedAt time.Time                             `json:"created_at,required" format:"date-time"`
	Enabled   bool                                  `json:"enabled,required"`
	Name      string                                `json:"name,required"`
	Type      DLPEntryUpdateResponseIntegrationType `json:"type,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	ProfileID string                                `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseIntegrationJSON `json:"-"`
}

// dlpEntryUpdateResponseIntegrationJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseIntegration]
type dlpEntryUpdateResponseIntegrationJSON struct {
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

func (r *DLPEntryUpdateResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseIntegration) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseIntegrationType string

const (
	DLPEntryUpdateResponseIntegrationTypeIntegration DLPEntryUpdateResponseIntegrationType = "integration"
)

func (r DLPEntryUpdateResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryUpdateResponseExactData struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                `json:"case_sensitive,required"`
	CreatedAt     time.Time                           `json:"created_at,required" format:"date-time"`
	Enabled       bool                                `json:"enabled,required"`
	Name          string                              `json:"name,required"`
	Secret        bool                                `json:"secret,required"`
	Type          DLPEntryUpdateResponseExactDataType `json:"type,required"`
	UpdatedAt     time.Time                           `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryUpdateResponseExactDataJSON `json:"-"`
}

// dlpEntryUpdateResponseExactDataJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseExactData]
type dlpEntryUpdateResponseExactDataJSON struct {
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

func (r *DLPEntryUpdateResponseExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseExactData) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseExactDataType string

const (
	DLPEntryUpdateResponseExactDataTypeExactData DLPEntryUpdateResponseExactDataType = "exact_data"
)

func (r DLPEntryUpdateResponseExactDataType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseExactDataTypeExactData:
		return true
	}
	return false
}

type DLPEntryUpdateResponseDocumentFingerprint struct {
	ID        string                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                          `json:"enabled,required"`
	Name      string                                        `json:"name,required"`
	Type      DLPEntryUpdateResponseDocumentFingerprintType `json:"type,required"`
	UpdatedAt time.Time                                     `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryUpdateResponseDocumentFingerprintJSON `json:"-"`
}

// dlpEntryUpdateResponseDocumentFingerprintJSON contains the JSON metadata for the
// struct [DLPEntryUpdateResponseDocumentFingerprint]
type dlpEntryUpdateResponseDocumentFingerprintJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryUpdateResponseDocumentFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseDocumentFingerprintJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseDocumentFingerprint) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseDocumentFingerprintType string

const (
	DLPEntryUpdateResponseDocumentFingerprintTypeDocumentFingerprint DLPEntryUpdateResponseDocumentFingerprintType = "document_fingerprint"
)

func (r DLPEntryUpdateResponseDocumentFingerprintType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseDocumentFingerprintTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryUpdateResponseWordList struct {
	ID        string                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                               `json:"enabled,required"`
	Name      string                             `json:"name,required"`
	Type      DLPEntryUpdateResponseWordListType `json:"type,required"`
	UpdatedAt time.Time                          `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                        `json:"word_list,required"`
	ProfileID string                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryUpdateResponseWordListJSON `json:"-"`
}

// dlpEntryUpdateResponseWordListJSON contains the JSON metadata for the struct
// [DLPEntryUpdateResponseWordList]
type dlpEntryUpdateResponseWordListJSON struct {
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

func (r *DLPEntryUpdateResponseWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryUpdateResponseWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryUpdateResponseWordList) implementsDLPEntryUpdateResponse() {}

type DLPEntryUpdateResponseWordListType string

const (
	DLPEntryUpdateResponseWordListTypeWordList DLPEntryUpdateResponseWordListType = "word_list"
)

func (r DLPEntryUpdateResponseWordListType) IsKnown() bool {
	switch r {
	case DLPEntryUpdateResponseWordListTypeWordList:
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
	ID      string                   `json:"id,required" format:"uuid"`
	Enabled bool                     `json:"enabled,required"`
	Name    string                   `json:"name,required"`
	Type    DLPEntryListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryListResponsePredefinedConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of [DLPEntryListResponsePredefinedVariant].
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
	Pattern       apijson.Field
	ProfileID     apijson.Field
	Secret        apijson.Field
	UpdatedAt     apijson.Field
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
// Possible runtime types of the union are [DLPEntryListResponseCustom],
// [DLPEntryListResponsePredefined], [DLPEntryListResponseIntegration],
// [DLPEntryListResponseExactData], [DLPEntryListResponseDocumentFingerprint],
// [DLPEntryListResponseWordList].
func (r DLPEntryListResponse) AsUnion() DLPEntryListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryListResponseCustom],
// [DLPEntryListResponsePredefined], [DLPEntryListResponseIntegration],
// [DLPEntryListResponseExactData], [DLPEntryListResponseDocumentFingerprint] or
// [DLPEntryListResponseWordList].
type DLPEntryListResponseUnion interface {
	implementsDLPEntryListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryListResponseCustom{}),
			DiscriminatorValue: "custom",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryListResponsePredefined{}),
			DiscriminatorValue: "predefined",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryListResponseIntegration{}),
			DiscriminatorValue: "integration",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryListResponseExactData{}),
			DiscriminatorValue: "exact_data",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryListResponseDocumentFingerprint{}),
			DiscriminatorValue: "document_fingerprint",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DLPEntryListResponseWordList{}),
			DiscriminatorValue: "word_list",
		},
	)
}

type DLPEntryListResponseCustom struct {
	ID        string                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                           `json:"enabled,required"`
	Name      string                         `json:"name,required"`
	Pattern   Pattern                        `json:"pattern,required"`
	Type      DLPEntryListResponseCustomType `json:"type,required"`
	UpdatedAt time.Time                      `json:"updated_at,required" format:"date-time"`
	ProfileID string                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryListResponseCustomJSON `json:"-"`
}

// dlpEntryListResponseCustomJSON contains the JSON metadata for the struct
// [DLPEntryListResponseCustom]
type dlpEntryListResponseCustomJSON struct {
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

func (r *DLPEntryListResponseCustom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseCustomJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseCustom) implementsDLPEntryListResponse() {}

type DLPEntryListResponseCustomType string

const (
	DLPEntryListResponseCustomTypeCustom DLPEntryListResponseCustomType = "custom"
)

func (r DLPEntryListResponseCustomType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryListResponsePredefined struct {
	ID         string                                   `json:"id,required" format:"uuid"`
	Confidence DLPEntryListResponsePredefinedConfidence `json:"confidence,required"`
	Enabled    bool                                     `json:"enabled,required"`
	Name       string                                   `json:"name,required"`
	Type       DLPEntryListResponsePredefinedType       `json:"type,required"`
	ProfileID  string                                   `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryListResponsePredefinedVariant    `json:"variant"`
	JSON       dlpEntryListResponsePredefinedJSON       `json:"-"`
}

// dlpEntryListResponsePredefinedJSON contains the JSON metadata for the struct
// [DLPEntryListResponsePredefined]
type dlpEntryListResponsePredefinedJSON struct {
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

func (r *DLPEntryListResponsePredefined) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponsePredefinedJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponsePredefined) implementsDLPEntryListResponse() {}

type DLPEntryListResponsePredefinedConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                         `json:"available,required"`
	JSON      dlpEntryListResponsePredefinedConfidenceJSON `json:"-"`
}

// dlpEntryListResponsePredefinedConfidenceJSON contains the JSON metadata for the
// struct [DLPEntryListResponsePredefinedConfidence]
type dlpEntryListResponsePredefinedConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryListResponsePredefinedConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponsePredefinedConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryListResponsePredefinedType string

const (
	DLPEntryListResponsePredefinedTypePredefined DLPEntryListResponsePredefinedType = "predefined"
)

func (r DLPEntryListResponsePredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryListResponsePredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryListResponsePredefinedVariant struct {
	TopicType   DLPEntryListResponsePredefinedVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryListResponsePredefinedVariantType      `json:"type,required"`
	Description string                                         `json:"description,nullable"`
	JSON        dlpEntryListResponsePredefinedVariantJSON      `json:"-"`
}

// dlpEntryListResponsePredefinedVariantJSON contains the JSON metadata for the
// struct [DLPEntryListResponsePredefinedVariant]
type dlpEntryListResponsePredefinedVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryListResponsePredefinedVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponsePredefinedVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryListResponsePredefinedVariantTopicType string

const (
	DLPEntryListResponsePredefinedVariantTopicTypeIntent  DLPEntryListResponsePredefinedVariantTopicType = "Intent"
	DLPEntryListResponsePredefinedVariantTopicTypeContent DLPEntryListResponsePredefinedVariantTopicType = "Content"
)

func (r DLPEntryListResponsePredefinedVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryListResponsePredefinedVariantTopicTypeIntent, DLPEntryListResponsePredefinedVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryListResponsePredefinedVariantType string

const (
	DLPEntryListResponsePredefinedVariantTypePromptTopic DLPEntryListResponsePredefinedVariantType = "PromptTopic"
)

func (r DLPEntryListResponsePredefinedVariantType) IsKnown() bool {
	switch r {
	case DLPEntryListResponsePredefinedVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryListResponseIntegration struct {
	ID        string                              `json:"id,required" format:"uuid"`
	CreatedAt time.Time                           `json:"created_at,required" format:"date-time"`
	Enabled   bool                                `json:"enabled,required"`
	Name      string                              `json:"name,required"`
	Type      DLPEntryListResponseIntegrationType `json:"type,required"`
	UpdatedAt time.Time                           `json:"updated_at,required" format:"date-time"`
	ProfileID string                              `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryListResponseIntegrationJSON `json:"-"`
}

// dlpEntryListResponseIntegrationJSON contains the JSON metadata for the struct
// [DLPEntryListResponseIntegration]
type dlpEntryListResponseIntegrationJSON struct {
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

func (r *DLPEntryListResponseIntegration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseIntegrationJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseIntegration) implementsDLPEntryListResponse() {}

type DLPEntryListResponseIntegrationType string

const (
	DLPEntryListResponseIntegrationTypeIntegration DLPEntryListResponseIntegrationType = "integration"
)

func (r DLPEntryListResponseIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryListResponseExactData struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                              `json:"case_sensitive,required"`
	CreatedAt     time.Time                         `json:"created_at,required" format:"date-time"`
	Enabled       bool                              `json:"enabled,required"`
	Name          string                            `json:"name,required"`
	Secret        bool                              `json:"secret,required"`
	Type          DLPEntryListResponseExactDataType `json:"type,required"`
	UpdatedAt     time.Time                         `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryListResponseExactDataJSON `json:"-"`
}

// dlpEntryListResponseExactDataJSON contains the JSON metadata for the struct
// [DLPEntryListResponseExactData]
type dlpEntryListResponseExactDataJSON struct {
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

func (r *DLPEntryListResponseExactData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseExactDataJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseExactData) implementsDLPEntryListResponse() {}

type DLPEntryListResponseExactDataType string

const (
	DLPEntryListResponseExactDataTypeExactData DLPEntryListResponseExactDataType = "exact_data"
)

func (r DLPEntryListResponseExactDataType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseExactDataTypeExactData:
		return true
	}
	return false
}

type DLPEntryListResponseDocumentFingerprint struct {
	ID        string                                      `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                   `json:"created_at,required" format:"date-time"`
	Enabled   bool                                        `json:"enabled,required"`
	Name      string                                      `json:"name,required"`
	Type      DLPEntryListResponseDocumentFingerprintType `json:"type,required"`
	UpdatedAt time.Time                                   `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryListResponseDocumentFingerprintJSON `json:"-"`
}

// dlpEntryListResponseDocumentFingerprintJSON contains the JSON metadata for the
// struct [DLPEntryListResponseDocumentFingerprint]
type dlpEntryListResponseDocumentFingerprintJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryListResponseDocumentFingerprint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseDocumentFingerprintJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseDocumentFingerprint) implementsDLPEntryListResponse() {}

type DLPEntryListResponseDocumentFingerprintType string

const (
	DLPEntryListResponseDocumentFingerprintTypeDocumentFingerprint DLPEntryListResponseDocumentFingerprintType = "document_fingerprint"
)

func (r DLPEntryListResponseDocumentFingerprintType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseDocumentFingerprintTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryListResponseWordList struct {
	ID        string                           `json:"id,required" format:"uuid"`
	CreatedAt time.Time                        `json:"created_at,required" format:"date-time"`
	Enabled   bool                             `json:"enabled,required"`
	Name      string                           `json:"name,required"`
	Type      DLPEntryListResponseWordListType `json:"type,required"`
	UpdatedAt time.Time                        `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                      `json:"word_list,required"`
	ProfileID string                           `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryListResponseWordListJSON `json:"-"`
}

// dlpEntryListResponseWordListJSON contains the JSON metadata for the struct
// [DLPEntryListResponseWordList]
type dlpEntryListResponseWordListJSON struct {
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

func (r *DLPEntryListResponseWordList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryListResponseWordListJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryListResponseWordList) implementsDLPEntryListResponse() {}

type DLPEntryListResponseWordListType string

const (
	DLPEntryListResponseWordListTypeWordList DLPEntryListResponseWordListType = "word_list"
)

func (r DLPEntryListResponseWordListType) IsKnown() bool {
	switch r {
	case DLPEntryListResponseWordListTypeWordList:
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

type DLPEntryDeleteResponse = interface{}

type DLPEntryGetResponse struct {
	ID      string                  `json:"id,required" format:"uuid"`
	Enabled bool                    `json:"enabled,required"`
	Name    string                  `json:"name,required"`
	Type    DLPEntryGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryGetResponsePredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	// This field can have the runtime type of
	// [[]DLPEntryGetResponseCustomEntryProfile],
	// [[]DLPEntryGetResponsePredefinedEntryProfile],
	// [[]DLPEntryGetResponseIntegrationEntryProfile],
	// [[]DLPEntryGetResponseExactDataEntryProfile],
	// [[]DLPEntryGetResponseDocumentFingerprintEntryProfile],
	// [[]DLPEntryGetResponseWordListEntryProfile].
	Profiles  interface{} `json:"profiles"`
	Secret    bool        `json:"secret"`
	UpdatedAt time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryGetResponsePredefinedEntryVariant].
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
// Possible runtime types of the union are [DLPEntryGetResponseCustomEntry],
// [DLPEntryGetResponsePredefinedEntry], [DLPEntryGetResponseIntegrationEntry],
// [DLPEntryGetResponseExactDataEntry],
// [DLPEntryGetResponseDocumentFingerprintEntry],
// [DLPEntryGetResponseWordListEntry].
func (r DLPEntryGetResponse) AsUnion() DLPEntryGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryGetResponseCustomEntry],
// [DLPEntryGetResponsePredefinedEntry], [DLPEntryGetResponseIntegrationEntry],
// [DLPEntryGetResponseExactDataEntry],
// [DLPEntryGetResponseDocumentFingerprintEntry] or
// [DLPEntryGetResponseWordListEntry].
type DLPEntryGetResponseUnion interface {
	implementsDLPEntryGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryGetResponseWordListEntry{}),
		},
	)
}

type DLPEntryGetResponseCustomEntry struct {
	ID        string                                  `json:"id,required" format:"uuid"`
	CreatedAt time.Time                               `json:"created_at,required" format:"date-time"`
	Enabled   bool                                    `json:"enabled,required"`
	Name      string                                  `json:"name,required"`
	Pattern   Pattern                                 `json:"pattern,required"`
	Type      DLPEntryGetResponseCustomEntryType      `json:"type,required"`
	UpdatedAt time.Time                               `json:"updated_at,required" format:"date-time"`
	ProfileID string                                  `json:"profile_id,nullable" format:"uuid"`
	Profiles  []DLPEntryGetResponseCustomEntryProfile `json:"profiles"`
	JSON      dlpEntryGetResponseCustomEntryJSON      `json:"-"`
}

// dlpEntryGetResponseCustomEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseCustomEntry]
type dlpEntryGetResponseCustomEntryJSON struct {
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

func (r *DLPEntryGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseCustomEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseCustomEntryType string

const (
	DLPEntryGetResponseCustomEntryTypeCustom DLPEntryGetResponseCustomEntryType = "custom"
)

func (r DLPEntryGetResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryGetResponseCustomEntryProfile struct {
	ID   string                                    `json:"id,required" format:"uuid"`
	Name string                                    `json:"name,required"`
	JSON dlpEntryGetResponseCustomEntryProfileJSON `json:"-"`
}

// dlpEntryGetResponseCustomEntryProfileJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseCustomEntryProfile]
type dlpEntryGetResponseCustomEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseCustomEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseCustomEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponsePredefinedEntry struct {
	ID         string                                       `json:"id,required" format:"uuid"`
	Confidence DLPEntryGetResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                         `json:"enabled,required"`
	Name       string                                       `json:"name,required"`
	Type       DLPEntryGetResponsePredefinedEntryType       `json:"type,required"`
	ProfileID  string                                       `json:"profile_id,nullable" format:"uuid"`
	Profiles   []DLPEntryGetResponsePredefinedEntryProfile  `json:"profiles"`
	Variant    DLPEntryGetResponsePredefinedEntryVariant    `json:"variant"`
	JSON       dlpEntryGetResponsePredefinedEntryJSON       `json:"-"`
}

// dlpEntryGetResponsePredefinedEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponsePredefinedEntry]
type dlpEntryGetResponsePredefinedEntryJSON struct {
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

func (r *DLPEntryGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponsePredefinedEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                             `json:"available,required"`
	JSON      dlpEntryGetResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryGetResponsePredefinedEntryConfidenceJSON contains the JSON metadata for
// the struct [DLPEntryGetResponsePredefinedEntryConfidence]
type dlpEntryGetResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryGetResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponsePredefinedEntryType string

const (
	DLPEntryGetResponsePredefinedEntryTypePredefined DLPEntryGetResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryGetResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryGetResponsePredefinedEntryProfile struct {
	ID   string                                        `json:"id,required" format:"uuid"`
	Name string                                        `json:"name,required"`
	JSON dlpEntryGetResponsePredefinedEntryProfileJSON `json:"-"`
}

// dlpEntryGetResponsePredefinedEntryProfileJSON contains the JSON metadata for the
// struct [DLPEntryGetResponsePredefinedEntryProfile]
type dlpEntryGetResponsePredefinedEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponsePredefinedEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponsePredefinedEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryGetResponsePredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryGetResponsePredefinedEntryVariantType      `json:"type,required"`
	Description string                                             `json:"description,nullable"`
	JSON        dlpEntryGetResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryGetResponsePredefinedEntryVariantJSON contains the JSON metadata for the
// struct [DLPEntryGetResponsePredefinedEntryVariant]
type dlpEntryGetResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryGetResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryGetResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryGetResponsePredefinedEntryVariantTopicTypeContent DLPEntryGetResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryGetResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryGetResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryGetResponsePredefinedEntryVariantType string

const (
	DLPEntryGetResponsePredefinedEntryVariantTypePromptTopic DLPEntryGetResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryGetResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryGetResponseIntegrationEntry struct {
	ID        string                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                         `json:"enabled,required"`
	Name      string                                       `json:"name,required"`
	Type      DLPEntryGetResponseIntegrationEntryType      `json:"type,required"`
	UpdatedAt time.Time                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	Profiles  []DLPEntryGetResponseIntegrationEntryProfile `json:"profiles"`
	JSON      dlpEntryGetResponseIntegrationEntryJSON      `json:"-"`
}

// dlpEntryGetResponseIntegrationEntryJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseIntegrationEntry]
type dlpEntryGetResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseIntegrationEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseIntegrationEntryType string

const (
	DLPEntryGetResponseIntegrationEntryTypeIntegration DLPEntryGetResponseIntegrationEntryType = "integration"
)

func (r DLPEntryGetResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryGetResponseIntegrationEntryProfile struct {
	ID   string                                         `json:"id,required" format:"uuid"`
	Name string                                         `json:"name,required"`
	JSON dlpEntryGetResponseIntegrationEntryProfileJSON `json:"-"`
}

// dlpEntryGetResponseIntegrationEntryProfileJSON contains the JSON metadata for
// the struct [DLPEntryGetResponseIntegrationEntryProfile]
type dlpEntryGetResponseIntegrationEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseIntegrationEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseIntegrationEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                       `json:"case_sensitive,required"`
	CreatedAt     time.Time                                  `json:"created_at,required" format:"date-time"`
	Enabled       bool                                       `json:"enabled,required"`
	Name          string                                     `json:"name,required"`
	Secret        bool                                       `json:"secret,required"`
	Type          DLPEntryGetResponseExactDataEntryType      `json:"type,required"`
	UpdatedAt     time.Time                                  `json:"updated_at,required" format:"date-time"`
	Profiles      []DLPEntryGetResponseExactDataEntryProfile `json:"profiles"`
	JSON          dlpEntryGetResponseExactDataEntryJSON      `json:"-"`
}

// dlpEntryGetResponseExactDataEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseExactDataEntry]
type dlpEntryGetResponseExactDataEntryJSON struct {
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

func (r *DLPEntryGetResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseExactDataEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseExactDataEntryType string

const (
	DLPEntryGetResponseExactDataEntryTypeExactData DLPEntryGetResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryGetResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryGetResponseExactDataEntryProfile struct {
	ID   string                                       `json:"id,required" format:"uuid"`
	Name string                                       `json:"name,required"`
	JSON dlpEntryGetResponseExactDataEntryProfileJSON `json:"-"`
}

// dlpEntryGetResponseExactDataEntryProfileJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseExactDataEntryProfile]
type dlpEntryGetResponseExactDataEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseExactDataEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseExactDataEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseDocumentFingerprintEntry struct {
	ID        string                                               `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                 `json:"enabled,required"`
	Name      string                                               `json:"name,required"`
	Type      DLPEntryGetResponseDocumentFingerprintEntryType      `json:"type,required"`
	UpdatedAt time.Time                                            `json:"updated_at,required" format:"date-time"`
	Profiles  []DLPEntryGetResponseDocumentFingerprintEntryProfile `json:"profiles"`
	JSON      dlpEntryGetResponseDocumentFingerprintEntryJSON      `json:"-"`
}

// dlpEntryGetResponseDocumentFingerprintEntryJSON contains the JSON metadata for
// the struct [DLPEntryGetResponseDocumentFingerprintEntry]
type dlpEntryGetResponseDocumentFingerprintEntryJSON struct {
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

func (r *DLPEntryGetResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseDocumentFingerprintEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseDocumentFingerprintEntryType string

const (
	DLPEntryGetResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryGetResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryGetResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryGetResponseDocumentFingerprintEntryProfile struct {
	ID   string                                                 `json:"id,required" format:"uuid"`
	Name string                                                 `json:"name,required"`
	JSON dlpEntryGetResponseDocumentFingerprintEntryProfileJSON `json:"-"`
}

// dlpEntryGetResponseDocumentFingerprintEntryProfileJSON contains the JSON
// metadata for the struct [DLPEntryGetResponseDocumentFingerprintEntryProfile]
type dlpEntryGetResponseDocumentFingerprintEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseDocumentFingerprintEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseDocumentFingerprintEntryProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryGetResponseWordListEntry struct {
	ID        string                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                      `json:"enabled,required"`
	Name      string                                    `json:"name,required"`
	Type      DLPEntryGetResponseWordListEntryType      `json:"type,required"`
	UpdatedAt time.Time                                 `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                               `json:"word_list,required"`
	ProfileID string                                    `json:"profile_id,nullable" format:"uuid"`
	Profiles  []DLPEntryGetResponseWordListEntryProfile `json:"profiles"`
	JSON      dlpEntryGetResponseWordListEntryJSON      `json:"-"`
}

// dlpEntryGetResponseWordListEntryJSON contains the JSON metadata for the struct
// [DLPEntryGetResponseWordListEntry]
type dlpEntryGetResponseWordListEntryJSON struct {
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

func (r *DLPEntryGetResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryGetResponseWordListEntry) implementsDLPEntryGetResponse() {}

type DLPEntryGetResponseWordListEntryType string

const (
	DLPEntryGetResponseWordListEntryTypeWordList DLPEntryGetResponseWordListEntryType = "word_list"
)

func (r DLPEntryGetResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryGetResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryGetResponseWordListEntryProfile struct {
	ID   string                                      `json:"id,required" format:"uuid"`
	Name string                                      `json:"name,required"`
	JSON dlpEntryGetResponseWordListEntryProfileJSON `json:"-"`
}

// dlpEntryGetResponseWordListEntryProfileJSON contains the JSON metadata for the
// struct [DLPEntryGetResponseWordListEntryProfile]
type dlpEntryGetResponseWordListEntryProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryGetResponseWordListEntryProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryGetResponseWordListEntryProfileJSON) RawJSON() string {
	return r.raw
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

type DLPEntryNewParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Enabled   param.Field[bool]         `json:"enabled,required"`
	Name      param.Field[string]       `json:"name,required"`
	Pattern   param.Field[PatternParam] `json:"pattern,required"`
	ProfileID param.Field[string]       `json:"profile_id" format:"uuid"`
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
	Type    param.Field[DLPEntryUpdateParamsBodyType] `json:"type,required"`
	Enabled param.Field[bool]                         `json:"enabled"`
	Name    param.Field[string]                       `json:"name"`
	Pattern param.Field[PatternParam]                 `json:"pattern"`
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
	Name    param.Field[string]                             `json:"name,required"`
	Pattern param.Field[PatternParam]                       `json:"pattern,required"`
	Type    param.Field[DLPEntryUpdateParamsBodyCustomType] `json:"type,required"`
	Enabled param.Field[bool]                               `json:"enabled"`
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
