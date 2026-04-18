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
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a DLP entry.
func (r *DLPEntryCustomService) Update(ctx context.Context, entryID string, params DLPEntryCustomUpdateParams, opts ...option.RequestOption) (res *DLPEntryCustomUpdateResponse, err error) {
	var env DLPEntryCustomUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes a DLP custom entry.
func (r *DLPEntryCustomService) Delete(ctx context.Context, entryID string, body DLPEntryCustomDeleteParams, opts ...option.RequestOption) (res *DLPEntryCustomDeleteResponse, err error) {
	var env DLPEntryCustomDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/entries/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DLPEntryCustomNewResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool      `json:"enabled" api:"required"`
	Name        string    `json:"name" api:"required"`
	Pattern     Pattern   `json:"pattern" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"nullable"`
	// Deprecated: deprecated
	ProfileID string                        `json:"profile_id" api:"nullable" format:"uuid"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// Deprecated: deprecated
	Enabled bool                             `json:"enabled" api:"required"`
	Name    string                           `json:"name" api:"required"`
	Type    DLPEntryCustomUpdateResponseType `json:"type" api:"required"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool `json:"case_sensitive"`
	// This field can have the runtime type of
	// [DLPEntryCustomUpdateResponsePredefinedEntryConfidence].
	Confidence  interface{} `json:"confidence"`
	CreatedAt   time.Time   `json:"created_at" format:"date-time"`
	Description string      `json:"description" api:"nullable"`
	Pattern     Pattern     `json:"pattern"`
	// Deprecated: deprecated
	ProfileID string    `json:"profile_id" api:"nullable" format:"uuid"`
	Secret    bool      `json:"secret"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This field can have the runtime type of
	// [DLPEntryCustomUpdateResponsePredefinedEntryVariant].
	Variant interface{} `json:"variant"`
	// This field can have the runtime type of [interface{}].
	WordList interface{}                      `json:"word_list"`
	JSON     dlpEntryCustomUpdateResponseJSON `json:"-"`
	union    DLPEntryCustomUpdateResponseUnion
}

// dlpEntryCustomUpdateResponseJSON contains the JSON metadata for the struct
// [DLPEntryCustomUpdateResponse]
type dlpEntryCustomUpdateResponseJSON struct {
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

func (r dlpEntryCustomUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DLPEntryCustomUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DLPEntryCustomUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DLPEntryCustomUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DLPEntryCustomUpdateResponseCustomEntry],
// [DLPEntryCustomUpdateResponsePredefinedEntry],
// [DLPEntryCustomUpdateResponseIntegrationEntry],
// [DLPEntryCustomUpdateResponseExactDataEntry],
// [DLPEntryCustomUpdateResponseDocumentFingerprintEntry],
// [DLPEntryCustomUpdateResponseWordListEntry].
func (r DLPEntryCustomUpdateResponse) AsUnion() DLPEntryCustomUpdateResponseUnion {
	return r.union
}

// Union satisfied by [DLPEntryCustomUpdateResponseCustomEntry],
// [DLPEntryCustomUpdateResponsePredefinedEntry],
// [DLPEntryCustomUpdateResponseIntegrationEntry],
// [DLPEntryCustomUpdateResponseExactDataEntry],
// [DLPEntryCustomUpdateResponseDocumentFingerprintEntry] or
// [DLPEntryCustomUpdateResponseWordListEntry].
type DLPEntryCustomUpdateResponseUnion interface {
	implementsDLPEntryCustomUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DLPEntryCustomUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomUpdateResponseCustomEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomUpdateResponsePredefinedEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomUpdateResponseIntegrationEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomUpdateResponseExactDataEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomUpdateResponseDocumentFingerprintEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DLPEntryCustomUpdateResponseWordListEntry{}),
		},
	)
}

type DLPEntryCustomUpdateResponseCustomEntry struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool                                        `json:"enabled" api:"required"`
	Name        string                                      `json:"name" api:"required"`
	Pattern     Pattern                                     `json:"pattern" api:"required"`
	Type        DLPEntryCustomUpdateResponseCustomEntryType `json:"type" api:"required"`
	UpdatedAt   time.Time                                   `json:"updated_at" api:"required" format:"date-time"`
	Description string                                      `json:"description" api:"nullable"`
	// Deprecated: deprecated
	ProfileID string                                      `json:"profile_id" api:"nullable" format:"uuid"`
	JSON      dlpEntryCustomUpdateResponseCustomEntryJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseCustomEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomUpdateResponseCustomEntry]
type dlpEntryCustomUpdateResponseCustomEntryJSON struct {
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

func (r *DLPEntryCustomUpdateResponseCustomEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseCustomEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomUpdateResponseCustomEntry) implementsDLPEntryCustomUpdateResponse() {}

type DLPEntryCustomUpdateResponseCustomEntryType string

const (
	DLPEntryCustomUpdateResponseCustomEntryTypeCustom DLPEntryCustomUpdateResponseCustomEntryType = "custom"
)

func (r DLPEntryCustomUpdateResponseCustomEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponseCustomEntryTypeCustom:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponsePredefinedEntry struct {
	ID         string                                                `json:"id" api:"required" format:"uuid"`
	Confidence DLPEntryCustomUpdateResponsePredefinedEntryConfidence `json:"confidence" api:"required"`
	Enabled    bool                                                  `json:"enabled" api:"required"`
	Name       string                                                `json:"name" api:"required"`
	Type       DLPEntryCustomUpdateResponsePredefinedEntryType       `json:"type" api:"required"`
	// Deprecated: deprecated
	ProfileID string                                             `json:"profile_id" api:"nullable" format:"uuid"`
	Variant   DLPEntryCustomUpdateResponsePredefinedEntryVariant `json:"variant"`
	JSON      dlpEntryCustomUpdateResponsePredefinedEntryJSON    `json:"-"`
}

// dlpEntryCustomUpdateResponsePredefinedEntryJSON contains the JSON metadata for
// the struct [DLPEntryCustomUpdateResponsePredefinedEntry]
type dlpEntryCustomUpdateResponsePredefinedEntryJSON struct {
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

func (r *DLPEntryCustomUpdateResponsePredefinedEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponsePredefinedEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomUpdateResponsePredefinedEntry) implementsDLPEntryCustomUpdateResponse() {}

type DLPEntryCustomUpdateResponsePredefinedEntryConfidence struct {
	// Indicates whether this entry has AI remote service validation.
	AIContextAvailable bool `json:"ai_context_available" api:"required"`
	// Indicates whether this entry has any form of validation that is not an AI remote
	// service.
	Available bool                                                      `json:"available" api:"required"`
	JSON      dlpEntryCustomUpdateResponsePredefinedEntryConfidenceJSON `json:"-"`
}

// dlpEntryCustomUpdateResponsePredefinedEntryConfidenceJSON contains the JSON
// metadata for the struct [DLPEntryCustomUpdateResponsePredefinedEntryConfidence]
type dlpEntryCustomUpdateResponsePredefinedEntryConfidenceJSON struct {
	AIContextAvailable apijson.Field
	Available          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponsePredefinedEntryConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponsePredefinedEntryConfidenceJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomUpdateResponsePredefinedEntryType string

const (
	DLPEntryCustomUpdateResponsePredefinedEntryTypePredefined DLPEntryCustomUpdateResponsePredefinedEntryType = "predefined"
)

func (r DLPEntryCustomUpdateResponsePredefinedEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponsePredefinedEntryTypePredefined:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponsePredefinedEntryVariant struct {
	TopicType   DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicType `json:"topic_type" api:"required"`
	Type        DLPEntryCustomUpdateResponsePredefinedEntryVariantType      `json:"type" api:"required"`
	Description string                                                      `json:"description" api:"nullable"`
	JSON        dlpEntryCustomUpdateResponsePredefinedEntryVariantJSON      `json:"-"`
}

// dlpEntryCustomUpdateResponsePredefinedEntryVariantJSON contains the JSON
// metadata for the struct [DLPEntryCustomUpdateResponsePredefinedEntryVariant]
type dlpEntryCustomUpdateResponsePredefinedEntryVariantJSON struct {
	TopicType   apijson.Field
	Type        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponsePredefinedEntryVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponsePredefinedEntryVariantJSON) RawJSON() string {
	return r.raw
}

type DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicType string

const (
	DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicTypeIntent  DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicType = "Intent"
	DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicTypeContent DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicType = "Content"
)

func (r DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicTypeIntent, DLPEntryCustomUpdateResponsePredefinedEntryVariantTopicTypeContent:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponsePredefinedEntryVariantType string

const (
	DLPEntryCustomUpdateResponsePredefinedEntryVariantTypePromptTopic DLPEntryCustomUpdateResponsePredefinedEntryVariantType = "PromptTopic"
)

func (r DLPEntryCustomUpdateResponsePredefinedEntryVariantType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponsePredefinedEntryVariantTypePromptTopic:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponseIntegrationEntry struct {
	ID        string                                           `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time                                        `json:"created_at" api:"required" format:"date-time"`
	Enabled   bool                                             `json:"enabled" api:"required"`
	Name      string                                           `json:"name" api:"required"`
	Type      DLPEntryCustomUpdateResponseIntegrationEntryType `json:"type" api:"required"`
	UpdatedAt time.Time                                        `json:"updated_at" api:"required" format:"date-time"`
	ProfileID string                                           `json:"profile_id" api:"nullable" format:"uuid"`
	JSON      dlpEntryCustomUpdateResponseIntegrationEntryJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseIntegrationEntryJSON contains the JSON metadata for
// the struct [DLPEntryCustomUpdateResponseIntegrationEntry]
type dlpEntryCustomUpdateResponseIntegrationEntryJSON struct {
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

func (r *DLPEntryCustomUpdateResponseIntegrationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseIntegrationEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomUpdateResponseIntegrationEntry) implementsDLPEntryCustomUpdateResponse() {}

type DLPEntryCustomUpdateResponseIntegrationEntryType string

const (
	DLPEntryCustomUpdateResponseIntegrationEntryTypeIntegration DLPEntryCustomUpdateResponseIntegrationEntryType = "integration"
)

func (r DLPEntryCustomUpdateResponseIntegrationEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponseIntegrationEntryTypeIntegration:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponseExactDataEntry struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Only applies to custom word lists. Determines if the words should be matched in
	// a case-sensitive manner Cannot be set to false if secret is true
	CaseSensitive bool                                           `json:"case_sensitive" api:"required"`
	CreatedAt     time.Time                                      `json:"created_at" api:"required" format:"date-time"`
	Enabled       bool                                           `json:"enabled" api:"required"`
	Name          string                                         `json:"name" api:"required"`
	Secret        bool                                           `json:"secret" api:"required"`
	Type          DLPEntryCustomUpdateResponseExactDataEntryType `json:"type" api:"required"`
	UpdatedAt     time.Time                                      `json:"updated_at" api:"required" format:"date-time"`
	JSON          dlpEntryCustomUpdateResponseExactDataEntryJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseExactDataEntryJSON contains the JSON metadata for
// the struct [DLPEntryCustomUpdateResponseExactDataEntry]
type dlpEntryCustomUpdateResponseExactDataEntryJSON struct {
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

func (r *DLPEntryCustomUpdateResponseExactDataEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseExactDataEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomUpdateResponseExactDataEntry) implementsDLPEntryCustomUpdateResponse() {}

type DLPEntryCustomUpdateResponseExactDataEntryType string

const (
	DLPEntryCustomUpdateResponseExactDataEntryTypeExactData DLPEntryCustomUpdateResponseExactDataEntryType = "exact_data"
)

func (r DLPEntryCustomUpdateResponseExactDataEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponseExactDataEntryTypeExactData:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponseDocumentFingerprintEntry struct {
	ID        string                                                   `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time                                                `json:"created_at" api:"required" format:"date-time"`
	Enabled   bool                                                     `json:"enabled" api:"required"`
	Name      string                                                   `json:"name" api:"required"`
	Type      DLPEntryCustomUpdateResponseDocumentFingerprintEntryType `json:"type" api:"required"`
	UpdatedAt time.Time                                                `json:"updated_at" api:"required" format:"date-time"`
	JSON      dlpEntryCustomUpdateResponseDocumentFingerprintEntryJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseDocumentFingerprintEntryJSON contains the JSON
// metadata for the struct [DLPEntryCustomUpdateResponseDocumentFingerprintEntry]
type dlpEntryCustomUpdateResponseDocumentFingerprintEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPEntryCustomUpdateResponseDocumentFingerprintEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseDocumentFingerprintEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomUpdateResponseDocumentFingerprintEntry) implementsDLPEntryCustomUpdateResponse() {
}

type DLPEntryCustomUpdateResponseDocumentFingerprintEntryType string

const (
	DLPEntryCustomUpdateResponseDocumentFingerprintEntryTypeDocumentFingerprint DLPEntryCustomUpdateResponseDocumentFingerprintEntryType = "document_fingerprint"
)

func (r DLPEntryCustomUpdateResponseDocumentFingerprintEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponseDocumentFingerprintEntryTypeDocumentFingerprint:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponseWordListEntry struct {
	ID        string                                        `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	Enabled   bool                                          `json:"enabled" api:"required"`
	Name      string                                        `json:"name" api:"required"`
	Type      DLPEntryCustomUpdateResponseWordListEntryType `json:"type" api:"required"`
	UpdatedAt time.Time                                     `json:"updated_at" api:"required" format:"date-time"`
	WordList  interface{}                                   `json:"word_list" api:"required"`
	ProfileID string                                        `json:"profile_id" api:"nullable" format:"uuid"`
	JSON      dlpEntryCustomUpdateResponseWordListEntryJSON `json:"-"`
}

// dlpEntryCustomUpdateResponseWordListEntryJSON contains the JSON metadata for the
// struct [DLPEntryCustomUpdateResponseWordListEntry]
type dlpEntryCustomUpdateResponseWordListEntryJSON struct {
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

func (r *DLPEntryCustomUpdateResponseWordListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpEntryCustomUpdateResponseWordListEntryJSON) RawJSON() string {
	return r.raw
}

func (r DLPEntryCustomUpdateResponseWordListEntry) implementsDLPEntryCustomUpdateResponse() {}

type DLPEntryCustomUpdateResponseWordListEntryType string

const (
	DLPEntryCustomUpdateResponseWordListEntryTypeWordList DLPEntryCustomUpdateResponseWordListEntryType = "word_list"
)

func (r DLPEntryCustomUpdateResponseWordListEntryType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponseWordListEntryTypeWordList:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponseType string

const (
	DLPEntryCustomUpdateResponseTypeCustom              DLPEntryCustomUpdateResponseType = "custom"
	DLPEntryCustomUpdateResponseTypePredefined          DLPEntryCustomUpdateResponseType = "predefined"
	DLPEntryCustomUpdateResponseTypeIntegration         DLPEntryCustomUpdateResponseType = "integration"
	DLPEntryCustomUpdateResponseTypeExactData           DLPEntryCustomUpdateResponseType = "exact_data"
	DLPEntryCustomUpdateResponseTypeDocumentFingerprint DLPEntryCustomUpdateResponseType = "document_fingerprint"
	DLPEntryCustomUpdateResponseTypeWordList            DLPEntryCustomUpdateResponseType = "word_list"
)

func (r DLPEntryCustomUpdateResponseType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateResponseTypeCustom, DLPEntryCustomUpdateResponseTypePredefined, DLPEntryCustomUpdateResponseTypeIntegration, DLPEntryCustomUpdateResponseTypeExactData, DLPEntryCustomUpdateResponseTypeDocumentFingerprint, DLPEntryCustomUpdateResponseTypeWordList:
		return true
	}
	return false
}

type DLPEntryCustomDeleteResponse = interface{}

type DLPEntryCustomNewParams struct {
	AccountID   param.Field[string]       `path:"account_id" api:"required"`
	Enabled     param.Field[bool]         `json:"enabled" api:"required"`
	Name        param.Field[string]       `json:"name" api:"required"`
	Pattern     param.Field[PatternParam] `json:"pattern" api:"required"`
	Description param.Field[string]       `json:"description"`
	ProfileID   param.Field[string]       `json:"profile_id" format:"uuid"`
}

func (r DLPEntryCustomNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPEntryCustomNewResponseEnvelope struct {
	Errors   []DLPEntryCustomNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPEntryCustomNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPEntryCustomNewResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
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
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
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
	AccountID param.Field[string]                 `path:"account_id" api:"required"`
	Body      DLPEntryCustomUpdateParamsBodyUnion `json:"body" api:"required"`
}

func (r DLPEntryCustomUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DLPEntryCustomUpdateParamsBody struct {
	Type        param.Field[DLPEntryCustomUpdateParamsBodyType] `json:"type" api:"required"`
	Description param.Field[string]                             `json:"description"`
	Enabled     param.Field[bool]                               `json:"enabled"`
	Name        param.Field[string]                             `json:"name"`
	Pattern     param.Field[PatternParam]                       `json:"pattern"`
}

func (r DLPEntryCustomUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryCustomUpdateParamsBody) implementsDLPEntryCustomUpdateParamsBodyUnion() {}

// Satisfied by [zero_trust.DLPEntryCustomUpdateParamsBodyCustom],
// [zero_trust.DLPEntryCustomUpdateParamsBodyPredefined],
// [zero_trust.DLPEntryCustomUpdateParamsBodyIntegration],
// [DLPEntryCustomUpdateParamsBody].
type DLPEntryCustomUpdateParamsBodyUnion interface {
	implementsDLPEntryCustomUpdateParamsBodyUnion()
}

type DLPEntryCustomUpdateParamsBodyCustom struct {
	Name        param.Field[string]                                   `json:"name" api:"required"`
	Pattern     param.Field[PatternParam]                             `json:"pattern" api:"required"`
	Type        param.Field[DLPEntryCustomUpdateParamsBodyCustomType] `json:"type" api:"required"`
	Description param.Field[string]                                   `json:"description"`
	Enabled     param.Field[bool]                                     `json:"enabled"`
}

func (r DLPEntryCustomUpdateParamsBodyCustom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryCustomUpdateParamsBodyCustom) implementsDLPEntryCustomUpdateParamsBodyUnion() {}

type DLPEntryCustomUpdateParamsBodyCustomType string

const (
	DLPEntryCustomUpdateParamsBodyCustomTypeCustom DLPEntryCustomUpdateParamsBodyCustomType = "custom"
)

func (r DLPEntryCustomUpdateParamsBodyCustomType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateParamsBodyCustomTypeCustom:
		return true
	}
	return false
}

type DLPEntryCustomUpdateParamsBodyPredefined struct {
	Type    param.Field[DLPEntryCustomUpdateParamsBodyPredefinedType] `json:"type" api:"required"`
	Enabled param.Field[bool]                                         `json:"enabled"`
}

func (r DLPEntryCustomUpdateParamsBodyPredefined) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryCustomUpdateParamsBodyPredefined) implementsDLPEntryCustomUpdateParamsBodyUnion() {}

type DLPEntryCustomUpdateParamsBodyPredefinedType string

const (
	DLPEntryCustomUpdateParamsBodyPredefinedTypePredefined DLPEntryCustomUpdateParamsBodyPredefinedType = "predefined"
)

func (r DLPEntryCustomUpdateParamsBodyPredefinedType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateParamsBodyPredefinedTypePredefined:
		return true
	}
	return false
}

type DLPEntryCustomUpdateParamsBodyIntegration struct {
	Type    param.Field[DLPEntryCustomUpdateParamsBodyIntegrationType] `json:"type" api:"required"`
	Enabled param.Field[bool]                                          `json:"enabled"`
}

func (r DLPEntryCustomUpdateParamsBodyIntegration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DLPEntryCustomUpdateParamsBodyIntegration) implementsDLPEntryCustomUpdateParamsBodyUnion() {}

type DLPEntryCustomUpdateParamsBodyIntegrationType string

const (
	DLPEntryCustomUpdateParamsBodyIntegrationTypeIntegration DLPEntryCustomUpdateParamsBodyIntegrationType = "integration"
)

func (r DLPEntryCustomUpdateParamsBodyIntegrationType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateParamsBodyIntegrationTypeIntegration:
		return true
	}
	return false
}

type DLPEntryCustomUpdateParamsBodyType string

const (
	DLPEntryCustomUpdateParamsBodyTypeCustom      DLPEntryCustomUpdateParamsBodyType = "custom"
	DLPEntryCustomUpdateParamsBodyTypePredefined  DLPEntryCustomUpdateParamsBodyType = "predefined"
	DLPEntryCustomUpdateParamsBodyTypeIntegration DLPEntryCustomUpdateParamsBodyType = "integration"
)

func (r DLPEntryCustomUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case DLPEntryCustomUpdateParamsBodyTypeCustom, DLPEntryCustomUpdateParamsBodyTypePredefined, DLPEntryCustomUpdateParamsBodyTypeIntegration:
		return true
	}
	return false
}

type DLPEntryCustomUpdateResponseEnvelope struct {
	Errors   []DLPEntryCustomUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPEntryCustomUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPEntryCustomUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
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
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
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

type DLPEntryCustomDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPEntryCustomDeleteResponseEnvelope struct {
	Errors   []DLPEntryCustomDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPEntryCustomDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPEntryCustomDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPEntryCustomDeleteResponse                `json:"result" api:"nullable"`
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
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
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
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
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
