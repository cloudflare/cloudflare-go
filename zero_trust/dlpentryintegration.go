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
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                                `json:"enabled,required"`
	Name    string                              `json:"name,required"`
	Type    DLPEntryIntegrationListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationListResponsePredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID    string                                      `json:"profile_id,nullable" format:"uuid"`
	Secret       bool                                        `json:"secret"`
	UpdatedAt    time.Time                                   `json:"updated_at" format:"date-time"`
	UploadStatus DLPEntryIntegrationListResponseUploadStatus `json:"upload_status"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationListResponsePredefinedEntryVariant].
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
// Possible runtime types of the union are
// [DLPEntryIntegrationListResponseCustomEntry],
// [DLPEntryIntegrationListResponsePredefinedEntry],
// [DLPEntryIntegrationListResponseIntegrationEntry],
// [DLPEntryIntegrationListResponseExactDataEntry],
// [DLPEntryIntegrationListResponseDocumentFingerprintEntry],
// [DLPEntryIntegrationListResponseWordListEntry].
func (r DLPEntryIntegrationListResponse) AsUnion() DLPEntryIntegrationListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryIntegrationListResponseCustomEntry],
// [DLPEntryIntegrationListResponsePredefinedEntry],
// [DLPEntryIntegrationListResponseIntegrationEntry],
// [DLPEntryIntegrationListResponseExactDataEntry],
// [DLPEntryIntegrationListResponseDocumentFingerprintEntry] or
// [DLPEntryIntegrationListResponseWordListEntry].
type DLPEntryIntegrationListResponseUnion interface {
	implementsDLPEntryIntegrationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryIntegrationListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationListResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationListResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationListResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationListResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationListResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationListResponseWordListEntry{}),
		},
	)
}

type DLPEntryIntegrationListResponseCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                           `json:"enabled,required"`
	Name        string                                         `json:"name,required"`
	Pattern     Pattern                                        `json:"pattern,required"`
	Type        DLPEntryIntegrationListResponseCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                      `json:"updated_at,required" format:"date-time"`
	Description string                                         `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID    string                                                 `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryIntegrationListResponseCustomEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryIntegrationListResponseCustomEntryJSON         `json:"-"`
}

// dlpEntryIntegrationListResponseCustomEntryJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationListResponseCustomEntry]
type dlpEntryIntegrationListResponseCustomEntryJSON struct {
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

func (r *DLPEntryIntegrationListResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseCustomEntry) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponseCustomEntryType string

const (
	DLPEntryIntegrationListResponseCustomEntryTypeCustom DLPEntryIntegrationListResponseCustomEntryType = "custom"
)

func (r DLPEntryIntegrationListResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseCustomEntryUploadStatus string

const (
	DLPEntryIntegrationListResponseCustomEntryUploadStatusEmpty      DLPEntryIntegrationListResponseCustomEntryUploadStatus = "empty"
	DLPEntryIntegrationListResponseCustomEntryUploadStatusUploading  DLPEntryIntegrationListResponseCustomEntryUploadStatus = "uploading"
	DLPEntryIntegrationListResponseCustomEntryUploadStatusPending    DLPEntryIntegrationListResponseCustomEntryUploadStatus = "pending"
	DLPEntryIntegrationListResponseCustomEntryUploadStatusProcessing DLPEntryIntegrationListResponseCustomEntryUploadStatus = "processing"
	DLPEntryIntegrationListResponseCustomEntryUploadStatusFailed     DLPEntryIntegrationListResponseCustomEntryUploadStatus = "failed"
	DLPEntryIntegrationListResponseCustomEntryUploadStatusComplete   DLPEntryIntegrationListResponseCustomEntryUploadStatus = "complete"
)

func (r DLPEntryIntegrationListResponseCustomEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseCustomEntryUploadStatusEmpty, DLPEntryIntegrationListResponseCustomEntryUploadStatusUploading, DLPEntryIntegrationListResponseCustomEntryUploadStatusPending, DLPEntryIntegrationListResponseCustomEntryUploadStatusProcessing, DLPEntryIntegrationListResponseCustomEntryUploadStatusFailed, DLPEntryIntegrationListResponseCustomEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponsePredefinedEntry struct {
	ID         string                                                   `json:"id,required" format:"uuid"`
	Confidence DLPEntryIntegrationListResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                     `json:"enabled,required"`
	Name       string                                                   `json:"name,required"`
	Type       DLPEntryIntegrationListResponsePredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID    string                                                     `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryIntegrationListResponsePredefinedEntryUploadStatus `json:"upload_status"`
	Variant      DLPEntryIntegrationListResponsePredefinedEntryVariant      `json:"variant"`
	JSON         dlpEntryIntegrationListResponsePredefinedEntryJSON         `json:"-"`
}

// dlpEntryIntegrationListResponsePredefinedEntryJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationListResponsePredefinedEntry]
type dlpEntryIntegrationListResponsePredefinedEntryJSON struct {
	ID           apijson.Field
	Confidence   apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	ProfileID    apijson.Field
	UploadStatus apijson.Field
	Variant      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponsePredefinedEntry) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                         `json:"available,required"`
	JSON      dlpEntryIntegrationListResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryIntegrationListResponsePredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct
// [DLPEntryIntegrationListResponsePredefinedEntryConfidence]
type dlpEntryIntegrationListResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationListResponsePredefinedEntryType string

const (
	DLPEntryIntegrationListResponsePredefinedEntryTypePredefined DLPEntryIntegrationListResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryIntegrationListResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponsePredefinedEntryUploadStatus string

const (
	DLPEntryIntegrationListResponsePredefinedEntryUploadStatusEmpty      DLPEntryIntegrationListResponsePredefinedEntryUploadStatus = "empty"
	DLPEntryIntegrationListResponsePredefinedEntryUploadStatusUploading  DLPEntryIntegrationListResponsePredefinedEntryUploadStatus = "uploading"
	DLPEntryIntegrationListResponsePredefinedEntryUploadStatusPending    DLPEntryIntegrationListResponsePredefinedEntryUploadStatus = "pending"
	DLPEntryIntegrationListResponsePredefinedEntryUploadStatusProcessing DLPEntryIntegrationListResponsePredefinedEntryUploadStatus = "processing"
	DLPEntryIntegrationListResponsePredefinedEntryUploadStatusFailed     DLPEntryIntegrationListResponsePredefinedEntryUploadStatus = "failed"
	DLPEntryIntegrationListResponsePredefinedEntryUploadStatusComplete   DLPEntryIntegrationListResponsePredefinedEntryUploadStatus = "complete"
)

func (r DLPEntryIntegrationListResponsePredefinedEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponsePredefinedEntryUploadStatusEmpty, DLPEntryIntegrationListResponsePredefinedEntryUploadStatusUploading, DLPEntryIntegrationListResponsePredefinedEntryUploadStatusPending, DLPEntryIntegrationListResponsePredefinedEntryUploadStatusProcessing, DLPEntryIntegrationListResponsePredefinedEntryUploadStatusFailed, DLPEntryIntegrationListResponsePredefinedEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryIntegrationListResponsePredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryIntegrationListResponsePredefinedEntryVariantType      `json:"type,required"`
	Description string                                                         `json:"description,nullable"`
	JSON        dlpEntryIntegrationListResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryIntegrationListResponsePredefinedEntryVariantJSON contains the JSON
// metadata for the struct [DLPEntryIntegrationListResponsePredefinedEntryVariant]
type dlpEntryIntegrationListResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationListResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryIntegrationListResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryIntegrationListResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryIntegrationListResponsePredefinedEntryVariantTopicTypeContent DLPEntryIntegrationListResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryIntegrationListResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryIntegrationListResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponsePredefinedEntryVariantType string

const (
	DLPEntryIntegrationListResponsePredefinedEntryVariantTypePromptTopic DLPEntryIntegrationListResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryIntegrationListResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseIntegrationEntry struct {
	ID           string                                                      `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                                   `json:"created_at,required" format:"date-time"`
	Enabled      bool                                                        `json:"enabled,required"`
	Name         string                                                      `json:"name,required"`
	Type         DLPEntryIntegrationListResponseIntegrationEntryType         `json:"type,required"`
	UpdatedAt    time.Time                                                   `json:"updated_at,required" format:"date-time"`
	ProfileID    string                                                      `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryIntegrationListResponseIntegrationEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryIntegrationListResponseIntegrationEntryJSON         `json:"-"`
}

// dlpEntryIntegrationListResponseIntegrationEntryJSON contains the JSON metadata
// for the struct [DLPEntryIntegrationListResponseIntegrationEntry]
type dlpEntryIntegrationListResponseIntegrationEntryJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	ProfileID    apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseIntegrationEntry) implementsDLPEntryIntegrationListResponse() {
}

type DLPEntryIntegrationListResponseIntegrationEntryType string

const (
	DLPEntryIntegrationListResponseIntegrationEntryTypeIntegration DLPEntryIntegrationListResponseIntegrationEntryType = "integration"
)

func (r DLPEntryIntegrationListResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseIntegrationEntryUploadStatus string

const (
	DLPEntryIntegrationListResponseIntegrationEntryUploadStatusEmpty      DLPEntryIntegrationListResponseIntegrationEntryUploadStatus = "empty"
	DLPEntryIntegrationListResponseIntegrationEntryUploadStatusUploading  DLPEntryIntegrationListResponseIntegrationEntryUploadStatus = "uploading"
	DLPEntryIntegrationListResponseIntegrationEntryUploadStatusPending    DLPEntryIntegrationListResponseIntegrationEntryUploadStatus = "pending"
	DLPEntryIntegrationListResponseIntegrationEntryUploadStatusProcessing DLPEntryIntegrationListResponseIntegrationEntryUploadStatus = "processing"
	DLPEntryIntegrationListResponseIntegrationEntryUploadStatusFailed     DLPEntryIntegrationListResponseIntegrationEntryUploadStatus = "failed"
	DLPEntryIntegrationListResponseIntegrationEntryUploadStatusComplete   DLPEntryIntegrationListResponseIntegrationEntryUploadStatus = "complete"
)

func (r DLPEntryIntegrationListResponseIntegrationEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseIntegrationEntryUploadStatusEmpty, DLPEntryIntegrationListResponseIntegrationEntryUploadStatusUploading, DLPEntryIntegrationListResponseIntegrationEntryUploadStatusPending, DLPEntryIntegrationListResponseIntegrationEntryUploadStatusProcessing, DLPEntryIntegrationListResponseIntegrationEntryUploadStatusFailed, DLPEntryIntegrationListResponseIntegrationEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                      `json:"case_sensitive,required"`
	CreatedAt     time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                      `json:"enabled,required"`
	Name          string                                                    `json:"name,required"`
	Secret        bool                                                      `json:"secret,required"`
	Type          DLPEntryIntegrationListResponseExactDataEntryType         `json:"type,required"`
	UpdatedAt     time.Time                                                 `json:"updated_at,required" format:"date-time"`
	UploadStatus  DLPEntryIntegrationListResponseExactDataEntryUploadStatus `json:"upload_status"`
	JSON          dlpEntryIntegrationListResponseExactDataEntryJSON         `json:"-"`
}

// dlpEntryIntegrationListResponseExactDataEntryJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationListResponseExactDataEntry]
type dlpEntryIntegrationListResponseExactDataEntryJSON struct {
	ID            apijson.Field
	CaseSensitive apijson.Field
	CreatedAt     apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Secret        apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadStatus  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseExactDataEntry) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponseExactDataEntryType string

const (
	DLPEntryIntegrationListResponseExactDataEntryTypeExactData DLPEntryIntegrationListResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryIntegrationListResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseExactDataEntryUploadStatus string

const (
	DLPEntryIntegrationListResponseExactDataEntryUploadStatusEmpty      DLPEntryIntegrationListResponseExactDataEntryUploadStatus = "empty"
	DLPEntryIntegrationListResponseExactDataEntryUploadStatusUploading  DLPEntryIntegrationListResponseExactDataEntryUploadStatus = "uploading"
	DLPEntryIntegrationListResponseExactDataEntryUploadStatusPending    DLPEntryIntegrationListResponseExactDataEntryUploadStatus = "pending"
	DLPEntryIntegrationListResponseExactDataEntryUploadStatusProcessing DLPEntryIntegrationListResponseExactDataEntryUploadStatus = "processing"
	DLPEntryIntegrationListResponseExactDataEntryUploadStatusFailed     DLPEntryIntegrationListResponseExactDataEntryUploadStatus = "failed"
	DLPEntryIntegrationListResponseExactDataEntryUploadStatusComplete   DLPEntryIntegrationListResponseExactDataEntryUploadStatus = "complete"
)

func (r DLPEntryIntegrationListResponseExactDataEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseExactDataEntryUploadStatusEmpty, DLPEntryIntegrationListResponseExactDataEntryUploadStatusUploading, DLPEntryIntegrationListResponseExactDataEntryUploadStatusPending, DLPEntryIntegrationListResponseExactDataEntryUploadStatusProcessing, DLPEntryIntegrationListResponseExactDataEntryUploadStatusFailed, DLPEntryIntegrationListResponseExactDataEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseDocumentFingerprintEntry struct {
	ID           string                                                              `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                                           `json:"created_at,required" format:"date-time"`
	Enabled      bool                                                                `json:"enabled,required"`
	Name         string                                                              `json:"name,required"`
	Type         DLPEntryIntegrationListResponseDocumentFingerprintEntryType         `json:"type,required"`
	UpdatedAt    time.Time                                                           `json:"updated_at,required" format:"date-time"`
	UploadStatus DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryIntegrationListResponseDocumentFingerprintEntryJSON         `json:"-"`
}

// dlpEntryIntegrationListResponseDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct
// [DLPEntryIntegrationListResponseDocumentFingerprintEntry]
type dlpEntryIntegrationListResponseDocumentFingerprintEntryJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseDocumentFingerprintEntry) implementsDLPEntryIntegrationListResponse() {
}

type DLPEntryIntegrationListResponseDocumentFingerprintEntryType string

const (
	DLPEntryIntegrationListResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryIntegrationListResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryIntegrationListResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus string

const (
	DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusEmpty      DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus = "empty"
	DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusUploading  DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus = "uploading"
	DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusPending    DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus = "pending"
	DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusProcessing DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus = "processing"
	DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusFailed     DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus = "failed"
	DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusComplete   DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus = "complete"
)

func (r DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusEmpty, DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusUploading, DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusPending, DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusProcessing, DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusFailed, DLPEntryIntegrationListResponseDocumentFingerprintEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseWordListEntry struct {
	ID           string                                                   `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                                `json:"created_at,required" format:"date-time"`
	Enabled      bool                                                     `json:"enabled,required"`
	Name         string                                                   `json:"name,required"`
	Type         DLPEntryIntegrationListResponseWordListEntryType         `json:"type,required"`
	UpdatedAt    time.Time                                                `json:"updated_at,required" format:"date-time"`
	WordList     interface{}                                              `json:"word_list,required"`
	ProfileID    string                                                   `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryIntegrationListResponseWordListEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryIntegrationListResponseWordListEntryJSON         `json:"-"`
}

// dlpEntryIntegrationListResponseWordListEntryJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationListResponseWordListEntry]
type dlpEntryIntegrationListResponseWordListEntryJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Enabled      apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	WordList     apijson.Field
	ProfileID    apijson.Field
	UploadStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPEntryIntegrationListResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationListResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationListResponseWordListEntry) implementsDLPEntryIntegrationListResponse() {}

type DLPEntryIntegrationListResponseWordListEntryType string

const (
	DLPEntryIntegrationListResponseWordListEntryTypeWordList DLPEntryIntegrationListResponseWordListEntryType = "word_list"
)

func (r DLPEntryIntegrationListResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryIntegrationListResponseWordListEntryUploadStatus string

const (
	DLPEntryIntegrationListResponseWordListEntryUploadStatusEmpty      DLPEntryIntegrationListResponseWordListEntryUploadStatus = "empty"
	DLPEntryIntegrationListResponseWordListEntryUploadStatusUploading  DLPEntryIntegrationListResponseWordListEntryUploadStatus = "uploading"
	DLPEntryIntegrationListResponseWordListEntryUploadStatusPending    DLPEntryIntegrationListResponseWordListEntryUploadStatus = "pending"
	DLPEntryIntegrationListResponseWordListEntryUploadStatusProcessing DLPEntryIntegrationListResponseWordListEntryUploadStatus = "processing"
	DLPEntryIntegrationListResponseWordListEntryUploadStatusFailed     DLPEntryIntegrationListResponseWordListEntryUploadStatus = "failed"
	DLPEntryIntegrationListResponseWordListEntryUploadStatusComplete   DLPEntryIntegrationListResponseWordListEntryUploadStatus = "complete"
)

func (r DLPEntryIntegrationListResponseWordListEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseWordListEntryUploadStatusEmpty, DLPEntryIntegrationListResponseWordListEntryUploadStatusUploading, DLPEntryIntegrationListResponseWordListEntryUploadStatusPending, DLPEntryIntegrationListResponseWordListEntryUploadStatusProcessing, DLPEntryIntegrationListResponseWordListEntryUploadStatusFailed, DLPEntryIntegrationListResponseWordListEntryUploadStatusComplete:
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

type DLPEntryIntegrationListResponseUploadStatus string

const (
	DLPEntryIntegrationListResponseUploadStatusEmpty      DLPEntryIntegrationListResponseUploadStatus = "empty"
	DLPEntryIntegrationListResponseUploadStatusUploading  DLPEntryIntegrationListResponseUploadStatus = "uploading"
	DLPEntryIntegrationListResponseUploadStatusPending    DLPEntryIntegrationListResponseUploadStatus = "pending"
	DLPEntryIntegrationListResponseUploadStatusProcessing DLPEntryIntegrationListResponseUploadStatus = "processing"
	DLPEntryIntegrationListResponseUploadStatusFailed     DLPEntryIntegrationListResponseUploadStatus = "failed"
	DLPEntryIntegrationListResponseUploadStatusComplete   DLPEntryIntegrationListResponseUploadStatus = "complete"
)

func (r DLPEntryIntegrationListResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationListResponseUploadStatusEmpty, DLPEntryIntegrationListResponseUploadStatusUploading, DLPEntryIntegrationListResponseUploadStatusPending, DLPEntryIntegrationListResponseUploadStatusProcessing, DLPEntryIntegrationListResponseUploadStatusFailed, DLPEntryIntegrationListResponseUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryIntegrationDeleteResponse = interface{}

type DLPEntryIntegrationGetResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                               `json:"enabled,required"`
	Name    string                             `json:"name,required"`
	Type    DLPEntryIntegrationGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationGetResponseObjectConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string `json:"profile_id,nullable" format:"uuid"`
	// This field can have the runtime type of
	// [[]DLPEntryIntegrationGetResponseObjectProfile].
	Profiles     interface{}                                `json:"profiles"`
	Secret       bool                                       `json:"secret"`
	UpdatedAt    time.Time                                  `json:"updated_at" format:"date-time"`
	UploadStatus DLPEntryIntegrationGetResponseUploadStatus `json:"upload_status"`
	// This field can have the runtime type of
	// [DLPEntryIntegrationGetResponseObjectVariant].
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
// Possible runtime types of the union are [DLPEntryIntegrationGetResponseObject],
// [DLPEntryIntegrationGetResponseObject], [DLPEntryIntegrationGetResponseObject],
// [DLPEntryIntegrationGetResponseObject], [DLPEntryIntegrationGetResponseObject],
// [DLPEntryIntegrationGetResponseObject].
func (r DLPEntryIntegrationGetResponse) AsUnion() DLPEntryIntegrationGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryIntegrationGetResponseObject],
// [DLPEntryIntegrationGetResponseObject], [DLPEntryIntegrationGetResponseObject],
// [DLPEntryIntegrationGetResponseObject], [DLPEntryIntegrationGetResponseObject]
// or [DLPEntryIntegrationGetResponseObject].
type DLPEntryIntegrationGetResponseUnion interface {
	implementsDLPEntryIntegrationGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryIntegrationGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryIntegrationGetResponseObject{}),
		},
	)
}

type DLPEntryIntegrationGetResponseObject struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                     `json:"enabled,required"`
	Name        string                                   `json:"name,required"`
	Pattern     Pattern                                  `json:"pattern,required"`
	Type        DLPEntryIntegrationGetResponseObjectType `json:"type,required"`
	UpdatedAt   time.Time                                `json:"updated_at,required" format:"date-time"`
	Description string                                   `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID    string                                           `json:"profile_id,nullable" format:"uuid"`
	Profiles     []DLPEntryIntegrationGetResponseObjectProfile    `json:"profiles"`
	UploadStatus DLPEntryIntegrationGetResponseObjectUploadStatus `json:"upload_status"`
	JSON         dlpEntryIntegrationGetResponseObjectJSON         `json:"-"`
}

// dlpEntryIntegrationGetResponseObjectJSON contains the JSON metadata for the
// struct [DLPEntryIntegrationGetResponseObject]
type dlpEntryIntegrationGetResponseObjectJSON struct {
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

func (r *DLPEntryIntegrationGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryIntegrationGetResponseObject) implementsDLPEntryIntegrationGetResponse() {}

type DLPEntryIntegrationGetResponseObjectType string

const (
	DLPEntryIntegrationGetResponseObjectTypeCustom DLPEntryIntegrationGetResponseObjectType = "custom"
)

func (r DLPEntryIntegrationGetResponseObjectType) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseObjectTypeCustom:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryIntegrationGetResponseObjectProfile struct {
	ID   string                                          `json:"id,required" format:"uuid"`
	Name string                                          `json:"name,required"`
	JSON dlpEntryIntegrationGetResponseObjectProfileJSON `json:"-"`
}

// dlpEntryIntegrationGetResponseObjectProfileJSON contains the JSON metadata for
// the struct [DLPEntryIntegrationGetResponseObjectProfile]
type dlpEntryIntegrationGetResponseObjectProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryIntegrationGetResponseObjectProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryIntegrationGetResponseObjectProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryIntegrationGetResponseObjectUploadStatus string

const (
	DLPEntryIntegrationGetResponseObjectUploadStatusEmpty      DLPEntryIntegrationGetResponseObjectUploadStatus = "empty"
	DLPEntryIntegrationGetResponseObjectUploadStatusUploading  DLPEntryIntegrationGetResponseObjectUploadStatus = "uploading"
	DLPEntryIntegrationGetResponseObjectUploadStatusPending    DLPEntryIntegrationGetResponseObjectUploadStatus = "pending"
	DLPEntryIntegrationGetResponseObjectUploadStatusProcessing DLPEntryIntegrationGetResponseObjectUploadStatus = "processing"
	DLPEntryIntegrationGetResponseObjectUploadStatusFailed     DLPEntryIntegrationGetResponseObjectUploadStatus = "failed"
	DLPEntryIntegrationGetResponseObjectUploadStatusComplete   DLPEntryIntegrationGetResponseObjectUploadStatus = "complete"
)

func (r DLPEntryIntegrationGetResponseObjectUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseObjectUploadStatusEmpty, DLPEntryIntegrationGetResponseObjectUploadStatusUploading, DLPEntryIntegrationGetResponseObjectUploadStatusPending, DLPEntryIntegrationGetResponseObjectUploadStatusProcessing, DLPEntryIntegrationGetResponseObjectUploadStatusFailed, DLPEntryIntegrationGetResponseObjectUploadStatusComplete:
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

type DLPEntryIntegrationGetResponseUploadStatus string

const (
	DLPEntryIntegrationGetResponseUploadStatusEmpty      DLPEntryIntegrationGetResponseUploadStatus = "empty"
	DLPEntryIntegrationGetResponseUploadStatusUploading  DLPEntryIntegrationGetResponseUploadStatus = "uploading"
	DLPEntryIntegrationGetResponseUploadStatusPending    DLPEntryIntegrationGetResponseUploadStatus = "pending"
	DLPEntryIntegrationGetResponseUploadStatusProcessing DLPEntryIntegrationGetResponseUploadStatus = "processing"
	DLPEntryIntegrationGetResponseUploadStatusFailed     DLPEntryIntegrationGetResponseUploadStatus = "failed"
	DLPEntryIntegrationGetResponseUploadStatusComplete   DLPEntryIntegrationGetResponseUploadStatus = "complete"
)

func (r DLPEntryIntegrationGetResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryIntegrationGetResponseUploadStatusEmpty, DLPEntryIntegrationGetResponseUploadStatusUploading, DLPEntryIntegrationGetResponseUploadStatusPending, DLPEntryIntegrationGetResponseUploadStatusProcessing, DLPEntryIntegrationGetResponseUploadStatusFailed, DLPEntryIntegrationGetResponseUploadStatusComplete:
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
