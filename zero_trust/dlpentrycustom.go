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
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool      `json:"enabled,required"`
	Name        string    `json:"name,required"`
	Pattern     Pattern   `json:"pattern,required"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	Description string    `json:"description,nullable"`
	// Deprecated: deprecated
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
	Description apijson.Field
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
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool      `json:"enabled,required"`
	Name        string    `json:"name,required"`
	Pattern     Pattern   `json:"pattern,required"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	Description string    `json:"description,nullable"`
	// Deprecated: deprecated
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
	Description apijson.Field
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
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                           `json:"enabled,required"`
	Name    string                         `json:"name,required"`
	Type    DLPEntryCustomListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryCustomListResponsePredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID    string                                 `json:"profile_id,nullable" format:"uuid"`
	Secret       bool                                   `json:"secret"`
	UpdatedAt    time.Time                              `json:"updated_at" format:"date-time"`
	UploadStatus DLPEntryCustomListResponseUploadStatus `json:"upload_status"`
	// This field can have the runtime type of
	// [DLPEntryCustomListResponsePredefinedEntryVariant].
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
// Possible runtime types of the union are [DLPEntryCustomListResponseCustomEntry],
// [DLPEntryCustomListResponsePredefinedEntry],
// [DLPEntryCustomListResponseIntegrationEntry],
// [DLPEntryCustomListResponseExactDataEntry],
// [DLPEntryCustomListResponseDocumentFingerprintEntry],
// [DLPEntryCustomListResponseWordListEntry].
func (r DLPEntryCustomListResponse) AsUnion() DLPEntryCustomListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryCustomListResponseCustomEntry],
// [DLPEntryCustomListResponsePredefinedEntry],
// [DLPEntryCustomListResponseIntegrationEntry],
// [DLPEntryCustomListResponseExactDataEntry],
// [DLPEntryCustomListResponseDocumentFingerprintEntry] or
// [DLPEntryCustomListResponseWordListEntry].
type DLPEntryCustomListResponseUnion interface {
	implementsDLPEntryCustomListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryCustomListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomListResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomListResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomListResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomListResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomListResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomListResponseWordListEntry{}),
		},
	)
}

type DLPEntryCustomListResponseCustomEntry struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                      `json:"enabled,required"`
	Name        string                                    `json:"name,required"`
	Pattern     Pattern                                   `json:"pattern,required"`
	Type        DLPEntryCustomListResponseCustomEntryType `json:"type,required"`
	UpdatedAt   time.Time                                 `json:"updated_at,required" format:"date-time"`
	Description string                                    `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID    string                                            `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryCustomListResponseCustomEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryCustomListResponseCustomEntryJSON         `json:"-"`
}

// dlpEntryCustomListResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomListResponseCustomEntry]
type dlpEntryCustomListResponseCustomEntryJSON struct {
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

func (r *DLPEntryCustomListResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseCustomEntry) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseCustomEntryType string

const (
	DLPEntryCustomListResponseCustomEntryTypeCustom DLPEntryCustomListResponseCustomEntryType = "custom"
)

func (r DLPEntryCustomListResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryCustomListResponseCustomEntryUploadStatus string

const (
	DLPEntryCustomListResponseCustomEntryUploadStatusEmpty      DLPEntryCustomListResponseCustomEntryUploadStatus = "empty"
	DLPEntryCustomListResponseCustomEntryUploadStatusUploading  DLPEntryCustomListResponseCustomEntryUploadStatus = "uploading"
	DLPEntryCustomListResponseCustomEntryUploadStatusPending    DLPEntryCustomListResponseCustomEntryUploadStatus = "pending"
	DLPEntryCustomListResponseCustomEntryUploadStatusProcessing DLPEntryCustomListResponseCustomEntryUploadStatus = "processing"
	DLPEntryCustomListResponseCustomEntryUploadStatusFailed     DLPEntryCustomListResponseCustomEntryUploadStatus = "failed"
	DLPEntryCustomListResponseCustomEntryUploadStatusComplete   DLPEntryCustomListResponseCustomEntryUploadStatus = "complete"
)

func (r DLPEntryCustomListResponseCustomEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseCustomEntryUploadStatusEmpty, DLPEntryCustomListResponseCustomEntryUploadStatusUploading, DLPEntryCustomListResponseCustomEntryUploadStatusPending, DLPEntryCustomListResponseCustomEntryUploadStatusProcessing, DLPEntryCustomListResponseCustomEntryUploadStatusFailed, DLPEntryCustomListResponseCustomEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryCustomListResponsePredefinedEntry struct {
	ID         string                                              `json:"id,required" format:"uuid"`
	Confidence DLPEntryCustomListResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                `json:"enabled,required"`
	Name       string                                              `json:"name,required"`
	Type       DLPEntryCustomListResponsePredefinedEntryType       `json:"type,required"`
	// Deprecated: deprecated
	ProfileID    string                                                `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryCustomListResponsePredefinedEntryUploadStatus `json:"upload_status"`
	Variant      DLPEntryCustomListResponsePredefinedEntryVariant      `json:"variant"`
	JSON         dlpEntryCustomListResponsePredefinedEntryJSON         `json:"-"`
}

// dlpEntryCustomListResponsePredefinedEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomListResponsePredefinedEntry]
type dlpEntryCustomListResponsePredefinedEntryJSON struct {
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

func (r *DLPEntryCustomListResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponsePredefinedEntry) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                    `json:"available,required"`
	JSON      dlpEntryCustomListResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryCustomListResponsePredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct [DLPEntryCustomListResponsePredefinedEntryConfidence]
type dlpEntryCustomListResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryCustomListResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomListResponsePredefinedEntryType string

const (
	DLPEntryCustomListResponsePredefinedEntryTypePredefined DLPEntryCustomListResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryCustomListResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryCustomListResponsePredefinedEntryUploadStatus string

const (
	DLPEntryCustomListResponsePredefinedEntryUploadStatusEmpty      DLPEntryCustomListResponsePredefinedEntryUploadStatus = "empty"
	DLPEntryCustomListResponsePredefinedEntryUploadStatusUploading  DLPEntryCustomListResponsePredefinedEntryUploadStatus = "uploading"
	DLPEntryCustomListResponsePredefinedEntryUploadStatusPending    DLPEntryCustomListResponsePredefinedEntryUploadStatus = "pending"
	DLPEntryCustomListResponsePredefinedEntryUploadStatusProcessing DLPEntryCustomListResponsePredefinedEntryUploadStatus = "processing"
	DLPEntryCustomListResponsePredefinedEntryUploadStatusFailed     DLPEntryCustomListResponsePredefinedEntryUploadStatus = "failed"
	DLPEntryCustomListResponsePredefinedEntryUploadStatusComplete   DLPEntryCustomListResponsePredefinedEntryUploadStatus = "complete"
)

func (r DLPEntryCustomListResponsePredefinedEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponsePredefinedEntryUploadStatusEmpty, DLPEntryCustomListResponsePredefinedEntryUploadStatusUploading, DLPEntryCustomListResponsePredefinedEntryUploadStatusPending, DLPEntryCustomListResponsePredefinedEntryUploadStatusProcessing, DLPEntryCustomListResponsePredefinedEntryUploadStatusFailed, DLPEntryCustomListResponsePredefinedEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryCustomListResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryCustomListResponsePredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryCustomListResponsePredefinedEntryVariantType      `json:"type,required"`
	Description string                                                    `json:"description,nullable"`
	JSON        dlpEntryCustomListResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryCustomListResponsePredefinedEntryVariantJSON contains the JSON metadata
// for the struct [DLPEntryCustomListResponsePredefinedEntryVariant]
type dlpEntryCustomListResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomListResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomListResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryCustomListResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryCustomListResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryCustomListResponsePredefinedEntryVariantTopicTypeContent DLPEntryCustomListResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryCustomListResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryCustomListResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryCustomListResponsePredefinedEntryVariantType string

const (
	DLPEntryCustomListResponsePredefinedEntryVariantTypePromptTopic DLPEntryCustomListResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryCustomListResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryCustomListResponseIntegrationEntry struct {
	ID           string                                                 `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                              `json:"created_at,required" format:"date-time"`
	Enabled      bool                                                   `json:"enabled,required"`
	Name         string                                                 `json:"name,required"`
	Type         DLPEntryCustomListResponseIntegrationEntryType         `json:"type,required"`
	UpdatedAt    time.Time                                              `json:"updated_at,required" format:"date-time"`
	ProfileID    string                                                 `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryCustomListResponseIntegrationEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryCustomListResponseIntegrationEntryJSON         `json:"-"`
}

// dlpEntryCustomListResponseIntegrationEntryJSON contains the JSON metadata for
// the struct [DLPEntryCustomListResponseIntegrationEntry]
type dlpEntryCustomListResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryCustomListResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseIntegrationEntry) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseIntegrationEntryType string

const (
	DLPEntryCustomListResponseIntegrationEntryTypeIntegration DLPEntryCustomListResponseIntegrationEntryType = "integration"
)

func (r DLPEntryCustomListResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryCustomListResponseIntegrationEntryUploadStatus string

const (
	DLPEntryCustomListResponseIntegrationEntryUploadStatusEmpty      DLPEntryCustomListResponseIntegrationEntryUploadStatus = "empty"
	DLPEntryCustomListResponseIntegrationEntryUploadStatusUploading  DLPEntryCustomListResponseIntegrationEntryUploadStatus = "uploading"
	DLPEntryCustomListResponseIntegrationEntryUploadStatusPending    DLPEntryCustomListResponseIntegrationEntryUploadStatus = "pending"
	DLPEntryCustomListResponseIntegrationEntryUploadStatusProcessing DLPEntryCustomListResponseIntegrationEntryUploadStatus = "processing"
	DLPEntryCustomListResponseIntegrationEntryUploadStatusFailed     DLPEntryCustomListResponseIntegrationEntryUploadStatus = "failed"
	DLPEntryCustomListResponseIntegrationEntryUploadStatusComplete   DLPEntryCustomListResponseIntegrationEntryUploadStatus = "complete"
)

func (r DLPEntryCustomListResponseIntegrationEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseIntegrationEntryUploadStatusEmpty, DLPEntryCustomListResponseIntegrationEntryUploadStatusUploading, DLPEntryCustomListResponseIntegrationEntryUploadStatusPending, DLPEntryCustomListResponseIntegrationEntryUploadStatusProcessing, DLPEntryCustomListResponseIntegrationEntryUploadStatusFailed, DLPEntryCustomListResponseIntegrationEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryCustomListResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                                 `json:"case_sensitive,required"`
	CreatedAt     time.Time                                            `json:"created_at,required" format:"date-time"`
	Enabled       bool                                                 `json:"enabled,required"`
	Name          string                                               `json:"name,required"`
	Secret        bool                                                 `json:"secret,required"`
	Type          DLPEntryCustomListResponseExactDataEntryType         `json:"type,required"`
	UpdatedAt     time.Time                                            `json:"updated_at,required" format:"date-time"`
	UploadStatus  DLPEntryCustomListResponseExactDataEntryUploadStatus `json:"upload_status"`
	JSON          dlpEntryCustomListResponseExactDataEntryJSON         `json:"-"`
}

// dlpEntryCustomListResponseExactDataEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomListResponseExactDataEntry]
type dlpEntryCustomListResponseExactDataEntryJSON struct {
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

func (r *DLPEntryCustomListResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseExactDataEntry) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseExactDataEntryType string

const (
	DLPEntryCustomListResponseExactDataEntryTypeExactData DLPEntryCustomListResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryCustomListResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryCustomListResponseExactDataEntryUploadStatus string

const (
	DLPEntryCustomListResponseExactDataEntryUploadStatusEmpty      DLPEntryCustomListResponseExactDataEntryUploadStatus = "empty"
	DLPEntryCustomListResponseExactDataEntryUploadStatusUploading  DLPEntryCustomListResponseExactDataEntryUploadStatus = "uploading"
	DLPEntryCustomListResponseExactDataEntryUploadStatusPending    DLPEntryCustomListResponseExactDataEntryUploadStatus = "pending"
	DLPEntryCustomListResponseExactDataEntryUploadStatusProcessing DLPEntryCustomListResponseExactDataEntryUploadStatus = "processing"
	DLPEntryCustomListResponseExactDataEntryUploadStatusFailed     DLPEntryCustomListResponseExactDataEntryUploadStatus = "failed"
	DLPEntryCustomListResponseExactDataEntryUploadStatusComplete   DLPEntryCustomListResponseExactDataEntryUploadStatus = "complete"
)

func (r DLPEntryCustomListResponseExactDataEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseExactDataEntryUploadStatusEmpty, DLPEntryCustomListResponseExactDataEntryUploadStatusUploading, DLPEntryCustomListResponseExactDataEntryUploadStatusPending, DLPEntryCustomListResponseExactDataEntryUploadStatusProcessing, DLPEntryCustomListResponseExactDataEntryUploadStatusFailed, DLPEntryCustomListResponseExactDataEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryCustomListResponseDocumentFingerprintEntry struct {
	ID           string                                                         `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                                      `json:"created_at,required" format:"date-time"`
	Enabled      bool                                                           `json:"enabled,required"`
	Name         string                                                         `json:"name,required"`
	Type         DLPEntryCustomListResponseDocumentFingerprintEntryType         `json:"type,required"`
	UpdatedAt    time.Time                                                      `json:"updated_at,required" format:"date-time"`
	UploadStatus DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryCustomListResponseDocumentFingerprintEntryJSON         `json:"-"`
}

// dlpEntryCustomListResponseDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct [DLPEntryCustomListResponseDocumentFingerprintEntry]
type dlpEntryCustomListResponseDocumentFingerprintEntryJSON struct {
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

func (r *DLPEntryCustomListResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseDocumentFingerprintEntry) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseDocumentFingerprintEntryType string

const (
	DLPEntryCustomListResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryCustomListResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryCustomListResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus string

const (
	DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusEmpty      DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus = "empty"
	DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusUploading  DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus = "uploading"
	DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusPending    DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus = "pending"
	DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusProcessing DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus = "processing"
	DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusFailed     DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus = "failed"
	DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusComplete   DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus = "complete"
)

func (r DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusEmpty, DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusUploading, DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusPending, DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusProcessing, DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusFailed, DLPEntryCustomListResponseDocumentFingerprintEntryUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryCustomListResponseWordListEntry struct {
	ID           string                                              `json:"id,required" format:"uuid"`
	CreatedAt    time.Time                                           `json:"created_at,required" format:"date-time"`
	Enabled      bool                                                `json:"enabled,required"`
	Name         string                                              `json:"name,required"`
	Type         DLPEntryCustomListResponseWordListEntryType         `json:"type,required"`
	UpdatedAt    time.Time                                           `json:"updated_at,required" format:"date-time"`
	WordList     interface{}                                         `json:"word_list,required"`
	ProfileID    string                                              `json:"profile_id,nullable" format:"uuid"`
	UploadStatus DLPEntryCustomListResponseWordListEntryUploadStatus `json:"upload_status"`
	JSON         dlpEntryCustomListResponseWordListEntryJSON         `json:"-"`
}

// dlpEntryCustomListResponseWordListEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomListResponseWordListEntry]
type dlpEntryCustomListResponseWordListEntryJSON struct {
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

func (r *DLPEntryCustomListResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomListResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomListResponseWordListEntry) implementsDLPEntryCustomListResponse() {}

type DLPEntryCustomListResponseWordListEntryType string

const (
	DLPEntryCustomListResponseWordListEntryTypeWordList DLPEntryCustomListResponseWordListEntryType = "word_list"
)

func (r DLPEntryCustomListResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryCustomListResponseWordListEntryUploadStatus string

const (
	DLPEntryCustomListResponseWordListEntryUploadStatusEmpty      DLPEntryCustomListResponseWordListEntryUploadStatus = "empty"
	DLPEntryCustomListResponseWordListEntryUploadStatusUploading  DLPEntryCustomListResponseWordListEntryUploadStatus = "uploading"
	DLPEntryCustomListResponseWordListEntryUploadStatusPending    DLPEntryCustomListResponseWordListEntryUploadStatus = "pending"
	DLPEntryCustomListResponseWordListEntryUploadStatusProcessing DLPEntryCustomListResponseWordListEntryUploadStatus = "processing"
	DLPEntryCustomListResponseWordListEntryUploadStatusFailed     DLPEntryCustomListResponseWordListEntryUploadStatus = "failed"
	DLPEntryCustomListResponseWordListEntryUploadStatusComplete   DLPEntryCustomListResponseWordListEntryUploadStatus = "complete"
)

func (r DLPEntryCustomListResponseWordListEntryUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseWordListEntryUploadStatusEmpty, DLPEntryCustomListResponseWordListEntryUploadStatusUploading, DLPEntryCustomListResponseWordListEntryUploadStatusPending, DLPEntryCustomListResponseWordListEntryUploadStatusProcessing, DLPEntryCustomListResponseWordListEntryUploadStatusFailed, DLPEntryCustomListResponseWordListEntryUploadStatusComplete:
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

type DLPEntryCustomListResponseUploadStatus string

const (
	DLPEntryCustomListResponseUploadStatusEmpty      DLPEntryCustomListResponseUploadStatus = "empty"
	DLPEntryCustomListResponseUploadStatusUploading  DLPEntryCustomListResponseUploadStatus = "uploading"
	DLPEntryCustomListResponseUploadStatusPending    DLPEntryCustomListResponseUploadStatus = "pending"
	DLPEntryCustomListResponseUploadStatusProcessing DLPEntryCustomListResponseUploadStatus = "processing"
	DLPEntryCustomListResponseUploadStatusFailed     DLPEntryCustomListResponseUploadStatus = "failed"
	DLPEntryCustomListResponseUploadStatusComplete   DLPEntryCustomListResponseUploadStatus = "complete"
)

func (r DLPEntryCustomListResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomListResponseUploadStatusEmpty, DLPEntryCustomListResponseUploadStatusUploading, DLPEntryCustomListResponseUploadStatusPending, DLPEntryCustomListResponseUploadStatusProcessing, DLPEntryCustomListResponseUploadStatusFailed, DLPEntryCustomListResponseUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryCustomDeleteResponse = interface{}

type DLPEntryCustomGetResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                          `json:"enabled,required"`
	Name    string                        `json:"name,required"`
	Type    DLPEntryCustomGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryCustomGetResponseObjectConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description,nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string `json:"profile_id,nullable" format:"uuid"`
	// This field can have the runtime type of
	// [[]DLPEntryCustomGetResponseObjectProfile].
	Profiles     interface{}                           `json:"profiles"`
	Secret       bool                                  `json:"secret"`
	UpdatedAt    time.Time                             `json:"updated_at" format:"date-time"`
	UploadStatus DLPEntryCustomGetResponseUploadStatus `json:"upload_status"`
	// This field can have the runtime type of
	// [DLPEntryCustomGetResponseObjectVariant].
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
// Possible runtime types of the union are [DLPEntryCustomGetResponseObject],
// [DLPEntryCustomGetResponseObject], [DLPEntryCustomGetResponseObject],
// [DLPEntryCustomGetResponseObject], [DLPEntryCustomGetResponseObject],
// [DLPEntryCustomGetResponseObject].
func (r DLPEntryCustomGetResponse) AsUnion() DLPEntryCustomGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryCustomGetResponseObject],
// [DLPEntryCustomGetResponseObject], [DLPEntryCustomGetResponseObject],
// [DLPEntryCustomGetResponseObject], [DLPEntryCustomGetResponseObject] or
// [DLPEntryCustomGetResponseObject].
type DLPEntryCustomGetResponseUnion interface {
	implementsDLPEntryCustomGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryCustomGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomGetResponseObject{}),
		},
	)
}

type DLPEntryCustomGetResponseObject struct {
	ID        string    `json:"id,required" format:"uuid"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                `json:"enabled,required"`
	Name        string                              `json:"name,required"`
	Pattern     Pattern                             `json:"pattern,required"`
	Type        DLPEntryCustomGetResponseObjectType `json:"type,required"`
	UpdatedAt   time.Time                           `json:"updated_at,required" format:"date-time"`
	Description string                              `json:"description,nullable"`
	// Deprecated: deprecated
	ProfileID    string                                      `json:"profile_id,nullable" format:"uuid"`
	Profiles     []DLPEntryCustomGetResponseObjectProfile    `json:"profiles"`
	UploadStatus DLPEntryCustomGetResponseObjectUploadStatus `json:"upload_status"`
	JSON         dlpEntryCustomGetResponseObjectJSON         `json:"-"`
}

// dlpEntryCustomGetResponseObjectJSON contains the JSON metadata for the struct
// [DLPEntryCustomGetResponseObject]
type dlpEntryCustomGetResponseObjectJSON struct {
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

func (r *DLPEntryCustomGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomGetResponseObject) implementsDLPEntryCustomGetResponse() {}

type DLPEntryCustomGetResponseObjectType string

const (
	DLPEntryCustomGetResponseObjectTypeCustom DLPEntryCustomGetResponseObjectType = "custom"
)

func (r DLPEntryCustomGetResponseObjectType) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseObjectTypeCustom:
		return true
	}
	return false
}

// Computed entry field for a profile that an entry is shared into.
type DLPEntryCustomGetResponseObjectProfile struct {
	ID   string                                     `json:"id,required" format:"uuid"`
	Name string                                     `json:"name,required"`
	JSON dlpEntryCustomGetResponseObjectProfileJSON `json:"-"`
}

// dlpEntryCustomGetResponseObjectProfileJSON contains the JSON metadata for the
// struct [DLPEntryCustomGetResponseObjectProfile]
type dlpEntryCustomGetResponseObjectProfileJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomGetResponseObjectProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomGetResponseObjectProfileJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomGetResponseObjectUploadStatus string

const (
	DLPEntryCustomGetResponseObjectUploadStatusEmpty      DLPEntryCustomGetResponseObjectUploadStatus = "empty"
	DLPEntryCustomGetResponseObjectUploadStatusUploading  DLPEntryCustomGetResponseObjectUploadStatus = "uploading"
	DLPEntryCustomGetResponseObjectUploadStatusPending    DLPEntryCustomGetResponseObjectUploadStatus = "pending"
	DLPEntryCustomGetResponseObjectUploadStatusProcessing DLPEntryCustomGetResponseObjectUploadStatus = "processing"
	DLPEntryCustomGetResponseObjectUploadStatusFailed     DLPEntryCustomGetResponseObjectUploadStatus = "failed"
	DLPEntryCustomGetResponseObjectUploadStatusComplete   DLPEntryCustomGetResponseObjectUploadStatus = "complete"
)

func (r DLPEntryCustomGetResponseObjectUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseObjectUploadStatusEmpty, DLPEntryCustomGetResponseObjectUploadStatusUploading, DLPEntryCustomGetResponseObjectUploadStatusPending, DLPEntryCustomGetResponseObjectUploadStatusProcessing, DLPEntryCustomGetResponseObjectUploadStatusFailed, DLPEntryCustomGetResponseObjectUploadStatusComplete:
		return true
	}
	return false
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

type DLPEntryCustomGetResponseUploadStatus string

const (
	DLPEntryCustomGetResponseUploadStatusEmpty      DLPEntryCustomGetResponseUploadStatus = "empty"
	DLPEntryCustomGetResponseUploadStatusUploading  DLPEntryCustomGetResponseUploadStatus = "uploading"
	DLPEntryCustomGetResponseUploadStatusPending    DLPEntryCustomGetResponseUploadStatus = "pending"
	DLPEntryCustomGetResponseUploadStatusProcessing DLPEntryCustomGetResponseUploadStatus = "processing"
	DLPEntryCustomGetResponseUploadStatusFailed     DLPEntryCustomGetResponseUploadStatus = "failed"
	DLPEntryCustomGetResponseUploadStatusComplete   DLPEntryCustomGetResponseUploadStatus = "complete"
)

func (r DLPEntryCustomGetResponseUploadStatus) IsKnown() bool {
	switch r {
	case DLPEntryCustomGetResponseUploadStatusEmpty, DLPEntryCustomGetResponseUploadStatusUploading, DLPEntryCustomGetResponseUploadStatusPending, DLPEntryCustomGetResponseUploadStatusProcessing, DLPEntryCustomGetResponseUploadStatusFailed, DLPEntryCustomGetResponseUploadStatusComplete:
		return true
	}
	return false
}

type DLPEntryCustomNewParams struct {
	AccountID   param.Field[string]       `path:"account_id,required"`
	Enabled     param.Field[bool]         `json:"enabled,required"`
	Name        param.Field[string]       `json:"name,required"`
	Pattern     param.Field[PatternParam] `json:"pattern,required"`
	Description param.Field[string]       `json:"description"`
	ProfileID   param.Field[string]       `json:"profile_id" format:"uuid"`
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
	AccountID   param.Field[string]       `path:"account_id,required"`
	Enabled     param.Field[bool]         `json:"enabled,required"`
	Name        param.Field[string]       `json:"name,required"`
	Pattern     param.Field[PatternParam] `json:"pattern,required"`
	Description param.Field[string]       `json:"description"`
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
