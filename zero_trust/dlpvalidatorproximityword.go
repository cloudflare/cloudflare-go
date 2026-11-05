// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// DLPValidatorProximityWordService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPValidatorProximityWordService] method instead.
type DLPValidatorProximityWordService struct {
	Options []option.RequestOption
}

// NewDLPValidatorProximityWordService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDLPValidatorProximityWordService(opts ...option.RequestOption) (r *DLPValidatorProximityWordService) {
	r = &DLPValidatorProximityWordService{}
	r.Options = opts
	return
}

// Creates a reusable list of weighted words that detection entries can attach to
// in order to refine their matches. An account may own at most 100 validators
// across all kinds.
func (r *DLPValidatorProximityWordService) New(ctx context.Context, params DLPValidatorProximityWordNewParams, opts ...option.RequestOption) (res *ProximityWordValidator, err error) {
	var env DLPValidatorProximityWordNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/validators/proximity_words", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the name and word list of a `proximity_words` validator. Attachments to
// detection entries are owned from the entry side and are not affected.
func (r *DLPValidatorProximityWordService) Update(ctx context.Context, validatorID string, params DLPValidatorProximityWordUpdateParams, opts ...option.RequestOption) (res *ProximityWordValidator, err error) {
	var env DLPValidatorProximityWordUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if validatorID == "" {
		err = errors.New("missing required validator_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/validators/proximity_words/%s", params.AccountID, validatorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists the `proximity_words` validators configured for the account.
func (r *DLPValidatorProximityWordService) List(ctx context.Context, query DLPValidatorProximityWordListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProximityWordValidator], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/validators/proximity_words", query.AccountID)
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

// Lists the `proximity_words` validators configured for the account.
func (r *DLPValidatorProximityWordService) ListAutoPaging(ctx context.Context, query DLPValidatorProximityWordListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProximityWordValidator] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a `proximity_words` validator. Fails with `409` while any detection
// entry is still attached to it; detach it from those entries first.
func (r *DLPValidatorProximityWordService) Delete(ctx context.Context, validatorID string, body DLPValidatorProximityWordDeleteParams, opts ...option.RequestOption) (res *DLPValidatorProximityWordDeleteResponse, err error) {
	var env DLPValidatorProximityWordDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if validatorID == "" {
		err = errors.New("missing required validator_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/validators/proximity_words/%s", body.AccountID, validatorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches a `proximity_words` validator by ID.
func (r *DLPValidatorProximityWordService) Get(ctx context.Context, validatorID string, query DLPValidatorProximityWordGetParams, opts ...option.RequestOption) (res *ProximityWordValidator, err error) {
	var env DLPValidatorProximityWordGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if validatorID == "" {
		err = errors.New("missing required validator_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/validators/proximity_words/%s", query.AccountID, validatorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A reusable list of weighted words. Attaching it to a detection entry raises the
// confidence of that entry's matches when one of the words appears nearby, and
// lowers it when none do.
type ProximityWordValidator struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
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
	Words []Word                     `json:"words" api:"required"`
	JSON  proximityWordValidatorJSON `json:"-"`
}

// proximityWordValidatorJSON contains the JSON metadata for the struct
// [ProximityWordValidator]
type proximityWordValidatorJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Name         apijson.Field
	ReferencedBy apijson.Field
	UpdatedAt    apijson.Field
	Words        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProximityWordValidator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proximityWordValidatorJSON) RawJSON() string {
	return r.raw
}

type ProximityWordValidatorArray []ProximityWordValidator

// How much a word contributes to a match's confidence when it is found nearby.
//
// Separate from a profile's high/medium/low confidence threshold: this describes
// how telling the word is, not how confident the resulting match is.
type Weight string

const (
	WeightStrong   Weight = "strong"
	WeightModerate Weight = "moderate"
	WeightWeak     Weight = "weak"
)

func (r Weight) IsKnown() bool {
	switch r {
	case WeightStrong, WeightModerate, WeightWeak:
		return true
	}
	return false
}

// A word that raises a match's confidence when it appears near the match.
type Word struct {
	// How much a word contributes to a match's confidence when it is found nearby.
	//
	// Separate from a profile's high/medium/low confidence threshold: this describes
	// how telling the word is, not how confident the resulting match is.
	Weight Weight `json:"weight" api:"required"`
	// The word or phrase to look for. Must contain 1–128 characters and is matched
	// case-insensitively.
	Word string   `json:"word" api:"required"`
	JSON wordJSON `json:"-"`
}

// wordJSON contains the JSON metadata for the struct [Word]
type wordJSON struct {
	Weight      apijson.Field
	Word        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Word) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wordJSON) RawJSON() string {
	return r.raw
}

// A word that raises a match's confidence when it appears near the match.
type WordParam struct {
	// How much a word contributes to a match's confidence when it is found nearby.
	//
	// Separate from a profile's high/medium/low confidence threshold: this describes
	// how telling the word is, not how confident the resulting match is.
	Weight param.Field[Weight] `json:"weight" api:"required"`
	// The word or phrase to look for. Must contain 1–128 characters and is matched
	// case-insensitively.
	Word param.Field[string] `json:"word" api:"required"`
}

func (r WordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPValidatorProximityWordDeleteResponse = interface{}

type DLPValidatorProximityWordNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A name for this validator, at most 50 characters.
	Name param.Field[string] `json:"name" api:"required"`
	// The words that raise confidence when found near a match, at most 100 per
	// validator. Matching ignores case, so the list may not contain the same word
	// twice with different capitalisation. Words are stored as sent.
	//
	// For example, a validator for medical record numbers might use
	// `[{"word": "patient", "weight": "strong"}, {"word": "diagnosis", "weight": "moderate"}, {"word": "mrn", "weight": "weak"}]`.
	Words param.Field[[]WordParam] `json:"words" api:"required"`
}

func (r DLPValidatorProximityWordNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPValidatorProximityWordNewResponseEnvelope struct {
	Errors   []DLPValidatorProximityWordNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPValidatorProximityWordNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPValidatorProximityWordNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A reusable list of weighted words. Attaching it to a detection entry raises the
	// confidence of that entry's matches when one of the words appears nearby, and
	// lowers it when none do.
	Result ProximityWordValidator                           `json:"result"`
	JSON   dlpValidatorProximityWordNewResponseEnvelopeJSON `json:"-"`
}

// dlpValidatorProximityWordNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPValidatorProximityWordNewResponseEnvelope]
type dlpValidatorProximityWordNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordNewResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPValidatorProximityWordNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpValidatorProximityWordNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpValidatorProximityWordNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPValidatorProximityWordNewResponseEnvelopeErrors]
type dlpValidatorProximityWordNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpValidatorProximityWordNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpValidatorProximityWordNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPValidatorProximityWordNewResponseEnvelopeErrorsSource]
type dlpValidatorProximityWordNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordNewResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           DLPValidatorProximityWordNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpValidatorProximityWordNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpValidatorProximityWordNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPValidatorProximityWordNewResponseEnvelopeMessages]
type dlpValidatorProximityWordNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    dlpValidatorProximityWordNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpValidatorProximityWordNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPValidatorProximityWordNewResponseEnvelopeMessagesSource]
type dlpValidatorProximityWordNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPValidatorProximityWordNewResponseEnvelopeSuccess bool

const (
	DLPValidatorProximityWordNewResponseEnvelopeSuccessTrue DLPValidatorProximityWordNewResponseEnvelopeSuccess = true
)

func (r DLPValidatorProximityWordNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPValidatorProximityWordNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPValidatorProximityWordUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A name for this validator, at most 50 characters.
	Name param.Field[string] `json:"name" api:"required"`
	// The words that raise confidence when found near a match, at most 100 per
	// validator. Matching ignores case, so the list may not contain the same word
	// twice with different capitalisation. Words are stored as sent.
	//
	// For example, a validator for medical record numbers might use
	// `[{"word": "patient", "weight": "strong"}, {"word": "diagnosis", "weight": "moderate"}, {"word": "mrn", "weight": "weak"}]`.
	Words param.Field[[]WordParam] `json:"words" api:"required"`
}

func (r DLPValidatorProximityWordUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPValidatorProximityWordUpdateResponseEnvelope struct {
	Errors   []DLPValidatorProximityWordUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPValidatorProximityWordUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPValidatorProximityWordUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// A reusable list of weighted words. Attaching it to a detection entry raises the
	// confidence of that entry's matches when one of the words appears nearby, and
	// lowers it when none do.
	Result ProximityWordValidator                              `json:"result"`
	JSON   dlpValidatorProximityWordUpdateResponseEnvelopeJSON `json:"-"`
}

// dlpValidatorProximityWordUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DLPValidatorProximityWordUpdateResponseEnvelope]
type dlpValidatorProximityWordUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordUpdateResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DLPValidatorProximityWordUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpValidatorProximityWordUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpValidatorProximityWordUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPValidatorProximityWordUpdateResponseEnvelopeErrors]
type dlpValidatorProximityWordUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dlpValidatorProximityWordUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpValidatorProximityWordUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DLPValidatorProximityWordUpdateResponseEnvelopeErrorsSource]
type dlpValidatorProximityWordUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordUpdateResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           DLPValidatorProximityWordUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpValidatorProximityWordUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpValidatorProximityWordUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DLPValidatorProximityWordUpdateResponseEnvelopeMessages]
type dlpValidatorProximityWordUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    dlpValidatorProximityWordUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpValidatorProximityWordUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DLPValidatorProximityWordUpdateResponseEnvelopeMessagesSource]
type dlpValidatorProximityWordUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPValidatorProximityWordUpdateResponseEnvelopeSuccess bool

const (
	DLPValidatorProximityWordUpdateResponseEnvelopeSuccessTrue DLPValidatorProximityWordUpdateResponseEnvelopeSuccess = true
)

func (r DLPValidatorProximityWordUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPValidatorProximityWordUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPValidatorProximityWordListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPValidatorProximityWordDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPValidatorProximityWordDeleteResponseEnvelope struct {
	Errors   []DLPValidatorProximityWordDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPValidatorProximityWordDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPValidatorProximityWordDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPValidatorProximityWordDeleteResponse                `json:"result" api:"nullable"`
	JSON    dlpValidatorProximityWordDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpValidatorProximityWordDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [DLPValidatorProximityWordDeleteResponseEnvelope]
type dlpValidatorProximityWordDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordDeleteResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DLPValidatorProximityWordDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpValidatorProximityWordDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpValidatorProximityWordDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPValidatorProximityWordDeleteResponseEnvelopeErrors]
type dlpValidatorProximityWordDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dlpValidatorProximityWordDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpValidatorProximityWordDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DLPValidatorProximityWordDeleteResponseEnvelopeErrorsSource]
type dlpValidatorProximityWordDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordDeleteResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           DLPValidatorProximityWordDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpValidatorProximityWordDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpValidatorProximityWordDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DLPValidatorProximityWordDeleteResponseEnvelopeMessages]
type dlpValidatorProximityWordDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    dlpValidatorProximityWordDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpValidatorProximityWordDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DLPValidatorProximityWordDeleteResponseEnvelopeMessagesSource]
type dlpValidatorProximityWordDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPValidatorProximityWordDeleteResponseEnvelopeSuccess bool

const (
	DLPValidatorProximityWordDeleteResponseEnvelopeSuccessTrue DLPValidatorProximityWordDeleteResponseEnvelopeSuccess = true
)

func (r DLPValidatorProximityWordDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPValidatorProximityWordDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPValidatorProximityWordGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPValidatorProximityWordGetResponseEnvelope struct {
	Errors   []DLPValidatorProximityWordGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPValidatorProximityWordGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPValidatorProximityWordGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A reusable list of weighted words. Attaching it to a detection entry raises the
	// confidence of that entry's matches when one of the words appears nearby, and
	// lowers it when none do.
	Result ProximityWordValidator                           `json:"result"`
	JSON   dlpValidatorProximityWordGetResponseEnvelopeJSON `json:"-"`
}

// dlpValidatorProximityWordGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPValidatorProximityWordGetResponseEnvelope]
type dlpValidatorProximityWordGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordGetResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPValidatorProximityWordGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpValidatorProximityWordGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpValidatorProximityWordGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPValidatorProximityWordGetResponseEnvelopeErrors]
type dlpValidatorProximityWordGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpValidatorProximityWordGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpValidatorProximityWordGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPValidatorProximityWordGetResponseEnvelopeErrorsSource]
type dlpValidatorProximityWordGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordGetResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           DLPValidatorProximityWordGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpValidatorProximityWordGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpValidatorProximityWordGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPValidatorProximityWordGetResponseEnvelopeMessages]
type dlpValidatorProximityWordGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPValidatorProximityWordGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPValidatorProximityWordGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    dlpValidatorProximityWordGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpValidatorProximityWordGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPValidatorProximityWordGetResponseEnvelopeMessagesSource]
type dlpValidatorProximityWordGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPValidatorProximityWordGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpValidatorProximityWordGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPValidatorProximityWordGetResponseEnvelopeSuccess bool

const (
	DLPValidatorProximityWordGetResponseEnvelopeSuccessTrue DLPValidatorProximityWordGetResponseEnvelopeSuccess = true
)

func (r DLPValidatorProximityWordGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPValidatorProximityWordGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
