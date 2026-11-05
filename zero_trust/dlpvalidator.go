// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// DLPValidatorService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPValidatorService] method instead.
type DLPValidatorService struct {
	Options        []option.RequestOption
	ProximityWords *DLPValidatorProximityWordService
}

// NewDLPValidatorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPValidatorService(opts ...option.RequestOption) (r *DLPValidatorService) {
	r = &DLPValidatorService{}
	r.Options = opts
	r.ProximityWords = NewDLPValidatorProximityWordService(opts...)
	return
}

// Lists the validators configured for the account, of every kind.
func (r *DLPValidatorService) List(ctx context.Context, params DLPValidatorListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Validator], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/validators", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists the validators configured for the account, of every kind.
func (r *DLPValidatorService) ListAutoPaging(ctx context.Context, params DLPValidatorListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Validator] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Fetches a validator of any kind by ID.
func (r *DLPValidatorService) Get(ctx context.Context, validatorID string, query DLPValidatorGetParams, opts ...option.RequestOption) (res *Validator, err error) {
	var env DLPValidatorGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if validatorID == "" {
		err = errors.New("missing required validator_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/validators/%s", query.AccountID, validatorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// What currently uses a validator. Deleting a validator that is still in use
// returns `409`, so check this first.
//
// Read-only. A validator is attached or detached by updating the detection entry
// that uses it, not by writing this back.
type References struct {
	// The detection entries using this validator, ordered by name.
	Entries []ReferencesEntry `json:"entries" api:"required"`
	// Number of detection entries this validator is attached to.
	EntryCount int64          `json:"entry_count" api:"required"`
	JSON       referencesJSON `json:"-"`
}

// referencesJSON contains the JSON metadata for the struct [References]
type referencesJSON struct {
	Entries     apijson.Field
	EntryCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *References) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesJSON) RawJSON() string {
	return r.raw
}

// A detection entry that uses a validator.
type ReferencesEntry struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// The type of a detection entry.
	Type ReferencesEntriesType `json:"type" api:"required"`
	JSON referencesEntryJSON   `json:"-"`
}

// referencesEntryJSON contains the JSON metadata for the struct [ReferencesEntry]
type referencesEntryJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesEntryJSON) RawJSON() string {
	return r.raw
}

// The type of a detection entry.
type ReferencesEntriesType string

const (
	ReferencesEntriesTypeCustom              ReferencesEntriesType = "custom"
	ReferencesEntriesTypeCustomPromptTopic   ReferencesEntriesType = "custom_prompt_topic"
	ReferencesEntriesTypePredefined          ReferencesEntriesType = "predefined"
	ReferencesEntriesTypeIntegration         ReferencesEntriesType = "integration"
	ReferencesEntriesTypeExactData           ReferencesEntriesType = "exact_data"
	ReferencesEntriesTypeDocumentFingerprint ReferencesEntriesType = "document_fingerprint"
	ReferencesEntriesTypeWordList            ReferencesEntriesType = "word_list"
)

func (r ReferencesEntriesType) IsKnown() bool {
	switch r {
	case ReferencesEntriesTypeCustom, ReferencesEntriesTypeCustomPromptTopic, ReferencesEntriesTypePredefined, ReferencesEntriesTypeIntegration, ReferencesEntriesTypeExactData, ReferencesEntriesTypeDocumentFingerprint, ReferencesEntriesTypeWordList:
		return true
	}
	return false
}

// Reusable logic that a detection entry can attach to in order to refine its
// matches. The `kind` field determines the rest of the body.
type Validator struct {
	ID        string        `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time     `json:"created_at" api:"required" format:"date-time"`
	Kind      ValidatorKind `json:"kind" api:"required"`
	// A name for this validator, at most 50 characters.
	Name string `json:"name" api:"required"`
	// What currently uses a validator. Deleting a validator that is still in use
	// returns `409`, so check this first.
	//
	// Read-only. A validator is attached or detached by updating the detection entry
	// that uses it, not by writing this back.
	ReferencedBy References `json:"referenced_by" api:"required"`
	UpdatedAt    time.Time  `json:"updated_at" api:"required" format:"date-time"`
	// This field can have the runtime type of [ValidatorWorkerValidatorDefaultUnion].
	Default interface{} `json:"default"`
	// The public script tag of the customer's Worker that scores matches.
	ScriptTag string `json:"script_tag"`
	// This field can have the runtime type of [[]Word].
	Words interface{}   `json:"words"`
	JSON  validatorJSON `json:"-"`
	union ValidatorUnion
}

// validatorJSON contains the JSON metadata for the struct [Validator]
type validatorJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Kind         apijson.Field
	Name         apijson.Field
	ReferencedBy apijson.Field
	UpdatedAt    apijson.Field
	Default      apijson.Field
	ScriptTag    apijson.Field
	Words        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r validatorJSON) RawJSON() string {
	return r.raw
}

func (r *Validator) UnmarshalJSON(data []byte) (err error) {
	*r = Validator{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ValidatorUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ValidatorProximityWordValidator],
// [ValidatorWorkerValidator].
func (r Validator) AsUnion() ValidatorUnion {
	return r.union
}

// Reusable logic that a detection entry can attach to in order to refine its
// matches. The `kind` field determines the rest of the body.
//
// Union satisfied by [ValidatorProximityWordValidator] or
// [ValidatorWorkerValidator].
type ValidatorUnion interface {
	implementsValidator()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ValidatorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ValidatorProximityWordValidator{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ValidatorWorkerValidator{}),
		},
	)
}

// A reusable list of weighted words. Attaching it to a detection entry raises the
// confidence of that entry's matches when one of the words appears nearby, and
// lowers it when none do.
type ValidatorProximityWordValidator struct {
	ID        string                              `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time                           `json:"created_at" api:"required" format:"date-time"`
	Kind      ValidatorProximityWordValidatorKind `json:"kind" api:"required"`
	// A name for this validator, at most 50 characters.
	Name string `json:"name" api:"required"`
	// What currently uses a validator. Deleting a validator that is still in use
	// returns `409`, so check this first.
	//
	// Read-only. A validator is attached or detached by updating the detection entry
	// that uses it, not by writing this back.
	ReferencedBy References `json:"referenced_by" api:"required"`
	UpdatedAt    time.Time  `json:"updated_at" api:"required" format:"date-time"`
	// The words that raise confidence when found near a match.
	Words []Word                              `json:"words" api:"required"`
	JSON  validatorProximityWordValidatorJSON `json:"-"`
}

// validatorProximityWordValidatorJSON contains the JSON metadata for the struct
// [ValidatorProximityWordValidator]
type validatorProximityWordValidatorJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Kind         apijson.Field
	Name         apijson.Field
	ReferencedBy apijson.Field
	UpdatedAt    apijson.Field
	Words        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ValidatorProximityWordValidator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validatorProximityWordValidatorJSON) RawJSON() string {
	return r.raw
}

func (r ValidatorProximityWordValidator) implementsValidator() {}

type ValidatorProximityWordValidatorKind string

const (
	ValidatorProximityWordValidatorKindProximityWords ValidatorProximityWordValidatorKind = "proximity_words"
)

func (r ValidatorProximityWordValidatorKind) IsKnown() bool {
	switch r {
	case ValidatorProximityWordValidatorKindProximityWords:
		return true
	}
	return false
}

// A customer Worker registered as a validator.
type ValidatorWorkerValidator struct {
	ID        string                       `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time                    `json:"created_at" api:"required" format:"date-time"`
	Kind      ValidatorWorkerValidatorKind `json:"kind" api:"required"`
	// A name for this validator, at most 50 characters.
	Name string `json:"name" api:"required"`
	// What currently uses a validator. Deleting a validator that is still in use
	// returns `409`, so check this first.
	//
	// Read-only. A validator is attached or detached by updating the detection entry
	// that uses it, not by writing this back.
	ReferencedBy References `json:"referenced_by" api:"required"`
	// The public script tag of the customer's Worker that scores matches.
	ScriptTag string                               `json:"script_tag" api:"required"`
	UpdatedAt time.Time                            `json:"updated_at" api:"required" format:"date-time"`
	Default   ValidatorWorkerValidatorDefaultUnion `json:"default"`
	JSON      validatorWorkerValidatorJSON         `json:"-"`
}

// validatorWorkerValidatorJSON contains the JSON metadata for the struct
// [ValidatorWorkerValidator]
type validatorWorkerValidatorJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Kind         apijson.Field
	Name         apijson.Field
	ReferencedBy apijson.Field
	ScriptTag    apijson.Field
	UpdatedAt    apijson.Field
	Default      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ValidatorWorkerValidator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validatorWorkerValidatorJSON) RawJSON() string {
	return r.raw
}

func (r ValidatorWorkerValidator) implementsValidator() {}

type ValidatorWorkerValidatorKind string

const (
	ValidatorWorkerValidatorKindWorker ValidatorWorkerValidatorKind = "worker"
)

func (r ValidatorWorkerValidatorKind) IsKnown() bool {
	switch r {
	case ValidatorWorkerValidatorKindWorker:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionFloat] or
// [ValidatorWorkerValidatorDefaultString].
type ValidatorWorkerValidatorDefaultUnion interface {
	ImplementsValidatorWorkerValidatorDefaultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ValidatorWorkerValidatorDefaultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(ValidatorWorkerValidatorDefaultString("")),
		},
	)
}

type ValidatorWorkerValidatorDefaultString string

const (
	ValidatorWorkerValidatorDefaultStringVeto ValidatorWorkerValidatorDefaultString = "veto"
)

func (r ValidatorWorkerValidatorDefaultString) IsKnown() bool {
	switch r {
	case ValidatorWorkerValidatorDefaultStringVeto:
		return true
	}
	return false
}

func (r ValidatorWorkerValidatorDefaultString) ImplementsValidatorWorkerValidatorDefaultUnion() {}

type ValidatorKind string

const (
	ValidatorKindProximityWords ValidatorKind = "proximity_words"
	ValidatorKindWorker         ValidatorKind = "worker"
)

func (r ValidatorKind) IsKnown() bool {
	switch r {
	case ValidatorKindProximityWords, ValidatorKindWorker:
		return true
	}
	return false
}

type ValidatorArray []Validator

type DLPValidatorListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Only return validators of this kind. Omit to return every kind.
	Kind param.Field[DLPValidatorListParamsKind] `query:"kind"`
}

// URLQuery serializes [DLPValidatorListParams]'s query parameters as `url.Values`.
func (r DLPValidatorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Only return validators of this kind. Omit to return every kind.
type DLPValidatorListParamsKind string

const (
	DLPValidatorListParamsKindProximityWords DLPValidatorListParamsKind = "proximity_words"
	DLPValidatorListParamsKindWorker         DLPValidatorListParamsKind = "worker"
)

func (r DLPValidatorListParamsKind) IsKnown() bool {
	switch r {
	case DLPValidatorListParamsKindProximityWords, DLPValidatorListParamsKindWorker:
		return true
	}
	return false
}

type DLPValidatorGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPValidatorGetResponseEnvelope struct {
	Errors   []DLPValidatorGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPValidatorGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPValidatorGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Reusable logic that a detection entry can attach to in order to refine its
	// matches. The `kind` field determines the rest of the body.
	Result Validator                           `json:"result"`
	JSON   dlpValidatorGetResponseEnvelopeJSON `json:"-"`
}

// dlpValidatorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPValidatorGetResponseEnvelope]
type dlpValidatorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorGetResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPValidatorGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpValidatorGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpValidatorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPValidatorGetResponseEnvelopeErrors]
type dlpValidatorGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpValidatorGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpValidatorGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPValidatorGetResponseEnvelopeErrorsSource]
type dlpValidatorGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorGetResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DLPValidatorGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpValidatorGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpValidatorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPValidatorGetResponseEnvelopeMessages]
type dlpValidatorGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dlpValidatorGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpValidatorGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPValidatorGetResponseEnvelopeMessagesSource]
type dlpValidatorGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPValidatorGetResponseEnvelopeSuccess bool

const (
	DLPValidatorGetResponseEnvelopeSuccessTrue DLPValidatorGetResponseEnvelopeSuccess = true
)

func (r DLPValidatorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPValidatorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
