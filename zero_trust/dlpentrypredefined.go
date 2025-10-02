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

// DLPEntryPredefinedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPEntryPredefinedService] method instead.
type DLPEntryPredefinedService struct {
	Options []option.RequestOption
}

// NewDLPEntryPredefinedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPEntryPredefinedService(opts ...option.RequestOption) (r *DLPEntryPredefinedService) {
	r = &DLPEntryPredefinedService{}
	r.Options = opts
	return
}

// Predefined entries can't be created, this will update an existing predefined
// entry This is needed for our generated terraform API
func (r *DLPEntryPredefinedService) New(ctx context.Context, params DLPEntryPredefinedNewParams, opts ...option.RequestOption) (res *DLPEntryPredefinedNewResponse, err error) {
	var env DLPEntryPredefinedNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/predefined", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a DLP entry.
func (r *DLPEntryPredefinedService) Update(ctx context.Context, entryID string, params DLPEntryPredefinedUpdateParams, opts ...option.RequestOption) (res *DLPEntryPredefinedUpdateResponse, err error) {
	var env DLPEntryPredefinedUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/predefined/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all DLP entries in an account.
func (r *DLPEntryPredefinedService) List(ctx context.Context, query DLPEntryPredefinedListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPEntryPredefinedListResponse], err error) {
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
func (r *DLPEntryPredefinedService) ListAutoPaging(ctx context.Context, query DLPEntryPredefinedListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPEntryPredefinedListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// This is a no-op as predefined entires can't be deleted but is needed for our
// generated terraform API
func (r *DLPEntryPredefinedService) Delete(ctx context.Context, entryID string, body DLPEntryPredefinedDeleteParams, opts ...option.RequestOption) (res *DLPEntryPredefinedDeleteResponse, err error) {
	var env DLPEntryPredefinedDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/predefined/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a DLP entry by ID.
func (r *DLPEntryPredefinedService) Get(ctx context.Context, entryID string, query DLPEntryPredefinedGetParams, opts ...option.RequestOption) (res *DLPEntryPredefinedGetResponse, err error) {
	var env DLPEntryPredefinedGetResponseEnvelope
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

type DLPEntryPredefinedNewResponse struct {
	ID         string                                  `json:"id,required" format:"uuid"`
	Confidence DLPEntryPredefinedNewResponseConfidence `json:"confidence,required"`
	Enabled    bool                                    `json:"enabled,required"`
	Name       string                                  `json:"name,required"`
	ProfileID  string                                  `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryPredefinedNewResponseVariant    `json:"variant"`
	JSON       dlpEntryPredefinedNewResponseJSON       `json:"-"`
}

// dlpEntryPredefinedNewResponseJSON contains the JSON metadata for the struct
// [DLPEntryPredefinedNewResponse]
type dlpEntryPredefinedNewResponseJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                        `json:"available,required"`
	JSON      dlpEntryPredefinedNewResponseConfidenceJSON `json:"-"`
}

// dlpEntryPredefinedNewResponseConfidenceJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedNewResponseConfidence]
type dlpEntryPredefinedNewResponseConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseVariant struct {
	TopicType   DLPEntryPredefinedNewResponseVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryPredefinedNewResponseVariantType      `json:"type,required"`
	Description string                                        `json:"description,nullable"`
	JSON        dlpEntryPredefinedNewResponseVariantJSON      `json:"-"`
}

// dlpEntryPredefinedNewResponseVariantJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedNewResponseVariant]
type dlpEntryPredefinedNewResponseVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseVariantTopicType string

const (
	DLPEntryPredefinedNewResponseVariantTopicTypeIntent  DLPEntryPredefinedNewResponseVariantTopicType = "Intent"
	DLPEntryPredefinedNewResponseVariantTopicTypeContent DLPEntryPredefinedNewResponseVariantTopicType = "Content"
)

func (r DLPEntryPredefinedNewResponseVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedNewResponseVariantTopicTypeIntent, DLPEntryPredefinedNewResponseVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryPredefinedNewResponseVariantType string

const (
	DLPEntryPredefinedNewResponseVariantTypePromptTopic DLPEntryPredefinedNewResponseVariantType = "PromptTopic"
)

func (r DLPEntryPredefinedNewResponseVariantType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedNewResponseVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryPredefinedUpdateResponse struct {
	ID         string                                     `json:"id,required" format:"uuid"`
	Confidence DLPEntryPredefinedUpdateResponseConfidence `json:"confidence,required"`
	Enabled    bool                                       `json:"enabled,required"`
	Name       string                                     `json:"name,required"`
	ProfileID  string                                     `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryPredefinedUpdateResponseVariant    `json:"variant"`
	JSON       dlpEntryPredefinedUpdateResponseJSON       `json:"-"`
}

// dlpEntryPredefinedUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryPredefinedUpdateResponse]
type dlpEntryPredefinedUpdateResponseJSON struct {
	ID          apijson.Field
	Confidence  apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                           `json:"available,required"`
	JSON      dlpEntryPredefinedUpdateResponseConfidenceJSON `json:"-"`
}

// dlpEntryPredefinedUpdateResponseConfidenceJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedUpdateResponseConfidence]
type dlpEntryPredefinedUpdateResponseConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseVariant struct {
	TopicType   DLPEntryPredefinedUpdateResponseVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryPredefinedUpdateResponseVariantType      `json:"type,required"`
	Description string                                           `json:"description,nullable"`
	JSON        dlpEntryPredefinedUpdateResponseVariantJSON      `json:"-"`
}

// dlpEntryPredefinedUpdateResponseVariantJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedUpdateResponseVariant]
type dlpEntryPredefinedUpdateResponseVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseVariantTopicType string

const (
	DLPEntryPredefinedUpdateResponseVariantTopicTypeIntent  DLPEntryPredefinedUpdateResponseVariantTopicType = "Intent"
	DLPEntryPredefinedUpdateResponseVariantTopicTypeContent DLPEntryPredefinedUpdateResponseVariantTopicType = "Content"
)

func (r DLPEntryPredefinedUpdateResponseVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedUpdateResponseVariantTopicTypeIntent, DLPEntryPredefinedUpdateResponseVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryPredefinedUpdateResponseVariantType string

const (
	DLPEntryPredefinedUpdateResponseVariantTypePromptTopic DLPEntryPredefinedUpdateResponseVariantType = "PromptTopic"
)

func (r DLPEntryPredefinedUpdateResponseVariantType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedUpdateResponseVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponse struct {
	ID      string                             `json:"id,required" format:"uuid"`
	Enabled bool                               `json:"enabled,required"`
	Name    string                             `json:"name,required"`
	Type    DLPEntryPredefinedListResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryPredefinedListResponsePredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryPredefinedListResponsePredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                        `json:"word_list"`
	JSON     dlpEntryPredefinedListResponseJSON `json:"-"`
	union    DLPEntryPredefinedListResponseUnion
}

// dlpEntryPredefinedListResponseJSON contains the JSON metadata for the struct
// [DLPEntryPredefinedListResponse]
type dlpEntryPredefinedListResponseJSON struct {
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

func (r dlpEntryPredefinedListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryPredefinedListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryPredefinedListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryPredefinedListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPEntryPredefinedListResponseCustomEntry],
// [DLPEntryPredefinedListResponsePredefinedEntry],
// [DLPEntryPredefinedListResponseIntegrationEntry],
// [DLPEntryPredefinedListResponseExactDataEntry],
// [DLPEntryPredefinedListResponseDocumentFingerprintEntry],
// [DLPEntryPredefinedListResponseWordListEntry].
func (r DLPEntryPredefinedListResponse) AsUnion() DLPEntryPredefinedListResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryPredefinedListResponseCustomEntry],
// [DLPEntryPredefinedListResponsePredefinedEntry],
// [DLPEntryPredefinedListResponseIntegrationEntry],
// [DLPEntryPredefinedListResponseExactDataEntry],
// [DLPEntryPredefinedListResponseDocumentFingerprintEntry] or
// [DLPEntryPredefinedListResponseWordListEntry].
type DLPEntryPredefinedListResponseUnion interface {
	implementsDLPEntryPredefinedListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryPredefinedListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedListResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedListResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedListResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedListResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedListResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedListResponseWordListEntry{}),
		},
	)
}

type DLPEntryPredefinedListResponseCustomEntry struct {
	ID        string                                        `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                     `json:"created_at,required" format:"date-time"`
	Enabled   bool                                          `json:"enabled,required"`
	Name      string                                        `json:"name,required"`
	Pattern   Pattern                                       `json:"pattern,required"`
	Type      DLPEntryPredefinedListResponseCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                     `json:"updated_at,required" format:"date-time"`
	ProfileID string                                        `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryPredefinedListResponseCustomEntryJSON `json:"-"`
}

// dlpEntryPredefinedListResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedListResponseCustomEntry]
type dlpEntryPredefinedListResponseCustomEntryJSON struct {
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

func (r *DLPEntryPredefinedListResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedListResponseCustomEntry) implementsDLPEntryPredefinedListResponse() {}

type DLPEntryPredefinedListResponseCustomEntryType string

const (
	DLPEntryPredefinedListResponseCustomEntryTypeCustom DLPEntryPredefinedListResponseCustomEntryType = "custom"
)

func (r DLPEntryPredefinedListResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponsePredefinedEntry struct {
	ID         string                                                  `json:"id,required" format:"uuid"`
	Confidence DLPEntryPredefinedListResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                    `json:"enabled,required"`
	Name       string                                                  `json:"name,required"`
	Type       DLPEntryPredefinedListResponsePredefinedEntryType       `json:"type,required"`
	ProfileID  string                                                  `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryPredefinedListResponsePredefinedEntryVariant    `json:"variant"`
	JSON       dlpEntryPredefinedListResponsePredefinedEntryJSON       `json:"-"`
}

// dlpEntryPredefinedListResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedListResponsePredefinedEntry]
type dlpEntryPredefinedListResponsePredefinedEntryJSON struct {
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

func (r *DLPEntryPredefinedListResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedListResponsePredefinedEntry) implementsDLPEntryPredefinedListResponse() {}

type DLPEntryPredefinedListResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                        `json:"available,required"`
	JSON      dlpEntryPredefinedListResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryPredefinedListResponsePredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct
// [DLPEntryPredefinedListResponsePredefinedEntryConfidence]
type dlpEntryPredefinedListResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryPredefinedListResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedListResponsePredefinedEntryType string

const (
	DLPEntryPredefinedListResponsePredefinedEntryTypePredefined DLPEntryPredefinedListResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryPredefinedListResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryPredefinedListResponsePredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryPredefinedListResponsePredefinedEntryVariantType      `json:"type,required"`
	Description string                                                        `json:"description,nullable"`
	JSON        dlpEntryPredefinedListResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryPredefinedListResponsePredefinedEntryVariantJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedListResponsePredefinedEntryVariant]
type dlpEntryPredefinedListResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedListResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedListResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryPredefinedListResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryPredefinedListResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryPredefinedListResponsePredefinedEntryVariantTopicTypeContent DLPEntryPredefinedListResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryPredefinedListResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryPredefinedListResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponsePredefinedEntryVariantType string

const (
	DLPEntryPredefinedListResponsePredefinedEntryVariantTypePromptTopic DLPEntryPredefinedListResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryPredefinedListResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponseIntegrationEntry struct {
	ID        string                                             `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                          `json:"created_at,required" format:"date-time"`
	Enabled   bool                                               `json:"enabled,required"`
	Name      string                                             `json:"name,required"`
	Type      DLPEntryPredefinedListResponseIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                          `json:"updated_at,required" format:"date-time"`
	ProfileID string                                             `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryPredefinedListResponseIntegrationEntryJSON `json:"-"`
}

// dlpEntryPredefinedListResponseIntegrationEntryJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedListResponseIntegrationEntry]
type dlpEntryPredefinedListResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryPredefinedListResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedListResponseIntegrationEntry) implementsDLPEntryPredefinedListResponse() {}

type DLPEntryPredefinedListResponseIntegrationEntryType string

const (
	DLPEntryPredefinedListResponseIntegrationEntryTypeIntegration DLPEntryPredefinedListResponseIntegrationEntryType = "integration"
)

func (r DLPEntryPredefinedListResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                             `json:"case_sensitive,required"`
	CreatedAt     time.Time                                        `json:"created_at,required" format:"date-time"`
	Enabled       bool                                             `json:"enabled,required"`
	Name          string                                           `json:"name,required"`
	Secret        bool                                             `json:"secret,required"`
	Type          DLPEntryPredefinedListResponseExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                        `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryPredefinedListResponseExactDataEntryJSON `json:"-"`
}

// dlpEntryPredefinedListResponseExactDataEntryJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedListResponseExactDataEntry]
type dlpEntryPredefinedListResponseExactDataEntryJSON struct {
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

func (r *DLPEntryPredefinedListResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedListResponseExactDataEntry) implementsDLPEntryPredefinedListResponse() {}

type DLPEntryPredefinedListResponseExactDataEntryType string

const (
	DLPEntryPredefinedListResponseExactDataEntryTypeExactData DLPEntryPredefinedListResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryPredefinedListResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponseDocumentFingerprintEntry struct {
	ID        string                                                     `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                  `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                       `json:"enabled,required"`
	Name      string                                                     `json:"name,required"`
	Type      DLPEntryPredefinedListResponseDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                  `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryPredefinedListResponseDocumentFingerprintEntryJSON `json:"-"`
}

// dlpEntryPredefinedListResponseDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedListResponseDocumentFingerprintEntry]
type dlpEntryPredefinedListResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedListResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedListResponseDocumentFingerprintEntry) implementsDLPEntryPredefinedListResponse() {
}

type DLPEntryPredefinedListResponseDocumentFingerprintEntryType string

const (
	DLPEntryPredefinedListResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryPredefinedListResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryPredefinedListResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponseWordListEntry struct {
	ID        string                                          `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled   bool                                            `json:"enabled,required"`
	Name      string                                          `json:"name,required"`
	Type      DLPEntryPredefinedListResponseWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                       `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                     `json:"word_list,required"`
	ProfileID string                                          `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryPredefinedListResponseWordListEntryJSON `json:"-"`
}

// dlpEntryPredefinedListResponseWordListEntryJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedListResponseWordListEntry]
type dlpEntryPredefinedListResponseWordListEntryJSON struct {
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

func (r *DLPEntryPredefinedListResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedListResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedListResponseWordListEntry) implementsDLPEntryPredefinedListResponse() {}

type DLPEntryPredefinedListResponseWordListEntryType string

const (
	DLPEntryPredefinedListResponseWordListEntryTypeWordList DLPEntryPredefinedListResponseWordListEntryType = "word_list"
)

func (r DLPEntryPredefinedListResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryPredefinedListResponseType string

const (
	DLPEntryPredefinedListResponseTypeCustom              DLPEntryPredefinedListResponseType = "custom"
	DLPEntryPredefinedListResponseTypePredefined          DLPEntryPredefinedListResponseType = "predefined"
	DLPEntryPredefinedListResponseTypeIntegration         DLPEntryPredefinedListResponseType = "integration"
	DLPEntryPredefinedListResponseTypeExactData           DLPEntryPredefinedListResponseType = "exact_data"
	DLPEntryPredefinedListResponseTypeDocumentFingerprint DLPEntryPredefinedListResponseType = "document_fingerprint"
	DLPEntryPredefinedListResponseTypeWordList            DLPEntryPredefinedListResponseType = "word_list"
)

func (r DLPEntryPredefinedListResponseType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedListResponseTypeCustom, DLPEntryPredefinedListResponseTypePredefined, DLPEntryPredefinedListResponseTypeIntegration, DLPEntryPredefinedListResponseTypeExactData, DLPEntryPredefinedListResponseTypeDocumentFingerprint, DLPEntryPredefinedListResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryPredefinedDeleteResponse = interface{}

type DLPEntryPredefinedGetResponse struct {
	ID      string                            `json:"id,required" format:"uuid"`
	Enabled bool                              `json:"enabled,required"`
	Name    string                            `json:"name,required"`
	Type    DLPEntryPredefinedGetResponseType `json:"type,required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryPredefinedGetResponsePredefinedEntryConfidence].
	Confidence interface{} `json:"confidence"`
	CreatedAt  time.Time   `json:"created_at" format:"date-time"`
	Pattern    Pattern     `json:"pattern"`
	ProfileID  string      `json:"profile_id,nullable" format:"uuid"`
	Secret     bool        `json:"secret"`
	UpdatedAt  time.Time   `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryPredefinedGetResponsePredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                       `json:"word_list"`
	JSON     dlpEntryPredefinedGetResponseJSON `json:"-"`
	union    DLPEntryPredefinedGetResponseUnion
}

// dlpEntryPredefinedGetResponseJSON contains the JSON metadata for the struct
// [DLPEntryPredefinedGetResponse]
type dlpEntryPredefinedGetResponseJSON struct {
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

func (r dlpEntryPredefinedGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryPredefinedGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryPredefinedGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryPredefinedGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPEntryPredefinedGetResponseCustomEntry],
// [DLPEntryPredefinedGetResponsePredefinedEntry],
// [DLPEntryPredefinedGetResponseIntegrationEntry],
// [DLPEntryPredefinedGetResponseExactDataEntry],
// [DLPEntryPredefinedGetResponseDocumentFingerprintEntry],
// [DLPEntryPredefinedGetResponseWordListEntry].
func (r DLPEntryPredefinedGetResponse) AsUnion() DLPEntryPredefinedGetResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryPredefinedGetResponseCustomEntry],
// [DLPEntryPredefinedGetResponsePredefinedEntry],
// [DLPEntryPredefinedGetResponseIntegrationEntry],
// [DLPEntryPredefinedGetResponseExactDataEntry],
// [DLPEntryPredefinedGetResponseDocumentFingerprintEntry] or
// [DLPEntryPredefinedGetResponseWordListEntry].
type DLPEntryPredefinedGetResponseUnion interface {
	implementsDLPEntryPredefinedGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryPredefinedGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedGetResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedGetResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedGetResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedGetResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedGetResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryPredefinedGetResponseWordListEntry{}),
		},
	)
}

type DLPEntryPredefinedGetResponseCustomEntry struct {
	ID        string                                       `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                    `json:"created_at,required" format:"date-time"`
	Enabled   bool                                         `json:"enabled,required"`
	Name      string                                       `json:"name,required"`
	Pattern   Pattern                                      `json:"pattern,required"`
	Type      DLPEntryPredefinedGetResponseCustomEntryType `json:"type,required"`
	UpdatedAt time.Time                                    `json:"updated_at,required" format:"date-time"`
	ProfileID string                                       `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryPredefinedGetResponseCustomEntryJSON `json:"-"`
}

// dlpEntryPredefinedGetResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedGetResponseCustomEntry]
type dlpEntryPredefinedGetResponseCustomEntryJSON struct {
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

func (r *DLPEntryPredefinedGetResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedGetResponseCustomEntry) implementsDLPEntryPredefinedGetResponse() {}

type DLPEntryPredefinedGetResponseCustomEntryType string

const (
	DLPEntryPredefinedGetResponseCustomEntryTypeCustom DLPEntryPredefinedGetResponseCustomEntryType = "custom"
)

func (r DLPEntryPredefinedGetResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponsePredefinedEntry struct {
	ID         string                                                 `json:"id,required" format:"uuid"`
	Confidence DLPEntryPredefinedGetResponsePredefinedEntryConfidence `json:"confidence,required"`
	Enabled    bool                                                   `json:"enabled,required"`
	Name       string                                                 `json:"name,required"`
	Type       DLPEntryPredefinedGetResponsePredefinedEntryType       `json:"type,required"`
	ProfileID  string                                                 `json:"profile_id,nullable" format:"uuid"`
	Variant    DLPEntryPredefinedGetResponsePredefinedEntryVariant    `json:"variant"`
	JSON       dlpEntryPredefinedGetResponsePredefinedEntryJSON       `json:"-"`
}

// dlpEntryPredefinedGetResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedGetResponsePredefinedEntry]
type dlpEntryPredefinedGetResponsePredefinedEntryJSON struct {
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

func (r *DLPEntryPredefinedGetResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedGetResponsePredefinedEntry) implementsDLPEntryPredefinedGetResponse() {}

type DLPEntryPredefinedGetResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available,required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                       `json:"available,required"`
	JSON      dlpEntryPredefinedGetResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryPredefinedGetResponsePredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedGetResponsePredefinedEntryConfidence]
type dlpEntryPredefinedGetResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedGetResponsePredefinedEntryType string

const (
	DLPEntryPredefinedGetResponsePredefinedEntryTypePredefined DLPEntryPredefinedGetResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryPredefinedGetResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicType `json:"topic_type,required"`
	Type        DLPEntryPredefinedGetResponsePredefinedEntryVariantType      `json:"type,required"`
	Description string                                                       `json:"description,nullable"`
	JSON        dlpEntryPredefinedGetResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryPredefinedGetResponsePredefinedEntryVariantJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedGetResponsePredefinedEntryVariant]
type dlpEntryPredefinedGetResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicTypeContent DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryPredefinedGetResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponsePredefinedEntryVariantType string

const (
	DLPEntryPredefinedGetResponsePredefinedEntryVariantTypePromptTopic DLPEntryPredefinedGetResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryPredefinedGetResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponseIntegrationEntry struct {
	ID        string                                            `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                         `json:"created_at,required" format:"date-time"`
	Enabled   bool                                              `json:"enabled,required"`
	Name      string                                            `json:"name,required"`
	Type      DLPEntryPredefinedGetResponseIntegrationEntryType `json:"type,required"`
	UpdatedAt time.Time                                         `json:"updated_at,required" format:"date-time"`
	ProfileID string                                            `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryPredefinedGetResponseIntegrationEntryJSON `json:"-"`
}

// dlpEntryPredefinedGetResponseIntegrationEntryJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedGetResponseIntegrationEntry]
type dlpEntryPredefinedGetResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryPredefinedGetResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedGetResponseIntegrationEntry) implementsDLPEntryPredefinedGetResponse() {}

type DLPEntryPredefinedGetResponseIntegrationEntryType string

const (
	DLPEntryPredefinedGetResponseIntegrationEntryTypeIntegration DLPEntryPredefinedGetResponseIntegrationEntryType = "integration"
)

func (r DLPEntryPredefinedGetResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponseExactDataEntry struct {
	ID string `json:"id,required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                            `json:"case_sensitive,required"`
	CreatedAt     time.Time                                       `json:"created_at,required" format:"date-time"`
	Enabled       bool                                            `json:"enabled,required"`
	Name          string                                          `json:"name,required"`
	Secret        bool                                            `json:"secret,required"`
	Type          DLPEntryPredefinedGetResponseExactDataEntryType `json:"type,required"`
	UpdatedAt     time.Time                                       `json:"updated_at,required" format:"date-time"`
	JSON          dlpEntryPredefinedGetResponseExactDataEntryJSON `json:"-"`
}

// dlpEntryPredefinedGetResponseExactDataEntryJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedGetResponseExactDataEntry]
type dlpEntryPredefinedGetResponseExactDataEntryJSON struct {
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

func (r *DLPEntryPredefinedGetResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedGetResponseExactDataEntry) implementsDLPEntryPredefinedGetResponse() {}

type DLPEntryPredefinedGetResponseExactDataEntryType string

const (
	DLPEntryPredefinedGetResponseExactDataEntryTypeExactData DLPEntryPredefinedGetResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryPredefinedGetResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponseDocumentFingerprintEntry struct {
	ID        string                                                    `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                                 `json:"created_at,required" format:"date-time"`
	Enabled   bool                                                      `json:"enabled,required"`
	Name      string                                                    `json:"name,required"`
	Type      DLPEntryPredefinedGetResponseDocumentFingerprintEntryType `json:"type,required"`
	UpdatedAt time.Time                                                 `json:"updated_at,required" format:"date-time"`
	JSON      dlpEntryPredefinedGetResponseDocumentFingerprintEntryJSON `json:"-"`
}

// dlpEntryPredefinedGetResponseDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedGetResponseDocumentFingerprintEntry]
type dlpEntryPredefinedGetResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedGetResponseDocumentFingerprintEntry) implementsDLPEntryPredefinedGetResponse() {
}

type DLPEntryPredefinedGetResponseDocumentFingerprintEntryType string

const (
	DLPEntryPredefinedGetResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryPredefinedGetResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryPredefinedGetResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponseWordListEntry struct {
	ID        string                                         `json:"id,required" format:"uuid"`
	CreatedAt time.Time                                      `json:"created_at,required" format:"date-time"`
	Enabled   bool                                           `json:"enabled,required"`
	Name      string                                         `json:"name,required"`
	Type      DLPEntryPredefinedGetResponseWordListEntryType `json:"type,required"`
	UpdatedAt time.Time                                      `json:"updated_at,required" format:"date-time"`
	WordList  interface{}                                    `json:"word_list,required"`
	ProfileID string                                         `json:"profile_id,nullable" format:"uuid"`
	JSON      dlpEntryPredefinedGetResponseWordListEntryJSON `json:"-"`
}

// dlpEntryPredefinedGetResponseWordListEntryJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedGetResponseWordListEntry]
type dlpEntryPredefinedGetResponseWordListEntryJSON struct {
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

func (r *DLPEntryPredefinedGetResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryPredefinedGetResponseWordListEntry) implementsDLPEntryPredefinedGetResponse() {}

type DLPEntryPredefinedGetResponseWordListEntryType string

const (
	DLPEntryPredefinedGetResponseWordListEntryTypeWordList DLPEntryPredefinedGetResponseWordListEntryType = "word_list"
)

func (r DLPEntryPredefinedGetResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryPredefinedGetResponseType string

const (
	DLPEntryPredefinedGetResponseTypeCustom              DLPEntryPredefinedGetResponseType = "custom"
	DLPEntryPredefinedGetResponseTypePredefined          DLPEntryPredefinedGetResponseType = "predefined"
	DLPEntryPredefinedGetResponseTypeIntegration         DLPEntryPredefinedGetResponseType = "integration"
	DLPEntryPredefinedGetResponseTypeExactData           DLPEntryPredefinedGetResponseType = "exact_data"
	DLPEntryPredefinedGetResponseTypeDocumentFingerprint DLPEntryPredefinedGetResponseType = "document_fingerprint"
	DLPEntryPredefinedGetResponseTypeWordList            DLPEntryPredefinedGetResponseType = "word_list"
)

func (r DLPEntryPredefinedGetResponseType) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponseTypeCustom, DLPEntryPredefinedGetResponseTypePredefined, DLPEntryPredefinedGetResponseTypeIntegration, DLPEntryPredefinedGetResponseTypeExactData, DLPEntryPredefinedGetResponseTypeDocumentFingerprint, DLPEntryPredefinedGetResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryPredefinedNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
	EntryID   param.Field[string] `json:"entry_id,required" format:"uuid"`
	// This field is not actually used as the owning profile for a predefined entry is
	// already set to a predefined profile
	ProfileID param.Field[string] `json:"profile_id" format:"uuid"`
}

func (r DLPEntryPredefinedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryPredefinedNewResponseEnvelope struct {
	Errors   []DLPEntryPredefinedNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryPredefinedNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryPredefinedNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryPredefinedNewResponse                `json:"result"`
	JSON    dlpEntryPredefinedNewResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedNewResponseEnvelope]
type dlpEntryPredefinedNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPEntryPredefinedNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryPredefinedNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedNewResponseEnvelopeErrors]
type dlpEntryPredefinedNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedNewResponseEnvelopeErrorsSource]
type dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPEntryPredefinedNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryPredefinedNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedNewResponseEnvelopeMessages]
type dlpEntryPredefinedNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedNewResponseEnvelopeMessagesSource]
type dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryPredefinedNewResponseEnvelopeSuccess bool

const (
	DLPEntryPredefinedNewResponseEnvelopeSuccessTrue DLPEntryPredefinedNewResponseEnvelopeSuccess = true
)

func (r DLPEntryPredefinedNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryPredefinedUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enabled   param.Field[bool]   `json:"enabled,required"`
}

func (r DLPEntryPredefinedUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryPredefinedUpdateResponseEnvelope struct {
	Errors   []DLPEntryPredefinedUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryPredefinedUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryPredefinedUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryPredefinedUpdateResponse                `json:"result"`
	JSON    dlpEntryPredefinedUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedUpdateResponseEnvelope]
type dlpEntryPredefinedUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedUpdateResponseEnvelopeErrors]
type dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource]
type dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedUpdateResponseEnvelopeMessages]
type dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource]
type dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryPredefinedUpdateResponseEnvelopeSuccess bool

const (
	DLPEntryPredefinedUpdateResponseEnvelopeSuccessTrue DLPEntryPredefinedUpdateResponseEnvelopeSuccess = true
)

func (r DLPEntryPredefinedUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryPredefinedListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryPredefinedDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryPredefinedDeleteResponseEnvelope struct {
	Errors   []DLPEntryPredefinedDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryPredefinedDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryPredefinedDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryPredefinedDeleteResponse                `json:"result,nullable"`
	JSON    dlpEntryPredefinedDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedDeleteResponseEnvelope]
type dlpEntryPredefinedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedDeleteResponseEnvelopeErrors]
type dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource]
type dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedDeleteResponseEnvelopeMessages]
type dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource]
type dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryPredefinedDeleteResponseEnvelopeSuccess bool

const (
	DLPEntryPredefinedDeleteResponseEnvelopeSuccessTrue DLPEntryPredefinedDeleteResponseEnvelopeSuccess = true
)

func (r DLPEntryPredefinedDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPEntryPredefinedGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPEntryPredefinedGetResponseEnvelope struct {
	Errors   []DLPEntryPredefinedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPEntryPredefinedGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPEntryPredefinedGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPEntryPredefinedGetResponse                `json:"result"`
	JSON    dlpEntryPredefinedGetResponseEnvelopeJSON    `json:"-"`
}

// dlpEntryPredefinedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPEntryPredefinedGetResponseEnvelope]
type dlpEntryPredefinedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPEntryPredefinedGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpEntryPredefinedGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpEntryPredefinedGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedGetResponseEnvelopeErrors]
type dlpEntryPredefinedGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpEntryPredefinedGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpEntryPredefinedGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPEntryPredefinedGetResponseEnvelopeErrorsSource]
type dlpEntryPredefinedGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPEntryPredefinedGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpEntryPredefinedGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpEntryPredefinedGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPEntryPredefinedGetResponseEnvelopeMessages]
type dlpEntryPredefinedGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPEntryPredefinedGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpEntryPredefinedGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpEntryPredefinedGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPEntryPredefinedGetResponseEnvelopeMessagesSource]
type dlpEntryPredefinedGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryPredefinedGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryPredefinedGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPEntryPredefinedGetResponseEnvelopeSuccess bool

const (
	DLPEntryPredefinedGetResponseEnvelopeSuccessTrue DLPEntryPredefinedGetResponseEnvelopeSuccess = true
)

func (r DLPEntryPredefinedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPEntryPredefinedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
