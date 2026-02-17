// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// DLPPayloadLogService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPPayloadLogService] method instead.
type DLPPayloadLogService struct {
	Options []option.RequestOption
}

// NewDLPPayloadLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPPayloadLogService(opts ...option.RequestOption) (r *DLPPayloadLogService) {
	r = &DLPPayloadLogService{}
	r.Options = opts
	return
}

// Set payload log settings
func (r *DLPPayloadLogService) Update(ctx context.Context, params DLPPayloadLogUpdateParams, opts ...option.RequestOption) (res *DLPPayloadLogUpdateResponse, err error) {
	var env DLPPayloadLogUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get payload log settings
func (r *DLPPayloadLogService) Get(ctx context.Context, query DLPPayloadLogGetParams, opts ...option.RequestOption) (res *DLPPayloadLogGetResponse, err error) {
	var env DLPPayloadLogGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPayloadLogUpdateResponse struct {
	// Masking level for payload logs.
	//
	// - `full`: The entire payload is masked.
	// - `partial`: Only partial payload content is masked.
	// - `clear`: No masking is applied to the payload content.
	// - `default`: DLP uses its default masking behavior.
	MaskingLevel DLPPayloadLogUpdateResponseMaskingLevel `json:"masking_level,required"`
	UpdatedAt    time.Time                               `json:"updated_at,required" format:"date-time"`
	// Base64-encoded public key for encrypting payload logs. Null when payload logging
	// is disabled.
	PublicKey string                          `json:"public_key,nullable"`
	JSON      dlpPayloadLogUpdateResponseJSON `json:"-"`
}

// dlpPayloadLogUpdateResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogUpdateResponse]
type dlpPayloadLogUpdateResponseJSON struct {
	MaskingLevel apijson.Field
	UpdatedAt    apijson.Field
	PublicKey    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Masking level for payload logs.
//
// - `full`: The entire payload is masked.
// - `partial`: Only partial payload content is masked.
// - `clear`: No masking is applied to the payload content.
// - `default`: DLP uses its default masking behavior.
type DLPPayloadLogUpdateResponseMaskingLevel string

const (
	DLPPayloadLogUpdateResponseMaskingLevelFull    DLPPayloadLogUpdateResponseMaskingLevel = "full"
	DLPPayloadLogUpdateResponseMaskingLevelPartial DLPPayloadLogUpdateResponseMaskingLevel = "partial"
	DLPPayloadLogUpdateResponseMaskingLevelClear   DLPPayloadLogUpdateResponseMaskingLevel = "clear"
	DLPPayloadLogUpdateResponseMaskingLevelDefault DLPPayloadLogUpdateResponseMaskingLevel = "default"
)

func (r DLPPayloadLogUpdateResponseMaskingLevel) IsKnown() bool {
	switch r {
	case DLPPayloadLogUpdateResponseMaskingLevelFull, DLPPayloadLogUpdateResponseMaskingLevelPartial, DLPPayloadLogUpdateResponseMaskingLevelClear, DLPPayloadLogUpdateResponseMaskingLevelDefault:
		return true
	}
	return false
}

type DLPPayloadLogGetResponse struct {
	// Masking level for payload logs.
	//
	// - `full`: The entire payload is masked.
	// - `partial`: Only partial payload content is masked.
	// - `clear`: No masking is applied to the payload content.
	// - `default`: DLP uses its default masking behavior.
	MaskingLevel DLPPayloadLogGetResponseMaskingLevel `json:"masking_level,required"`
	UpdatedAt    time.Time                            `json:"updated_at,required" format:"date-time"`
	// Base64-encoded public key for encrypting payload logs. Null when payload logging
	// is disabled.
	PublicKey string                       `json:"public_key,nullable"`
	JSON      dlpPayloadLogGetResponseJSON `json:"-"`
}

// dlpPayloadLogGetResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponse]
type dlpPayloadLogGetResponseJSON struct {
	MaskingLevel apijson.Field
	UpdatedAt    apijson.Field
	PublicKey    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseJSON) RawJSON() string {
	return r.raw
}

// Masking level for payload logs.
//
// - `full`: The entire payload is masked.
// - `partial`: Only partial payload content is masked.
// - `clear`: No masking is applied to the payload content.
// - `default`: DLP uses its default masking behavior.
type DLPPayloadLogGetResponseMaskingLevel string

const (
	DLPPayloadLogGetResponseMaskingLevelFull    DLPPayloadLogGetResponseMaskingLevel = "full"
	DLPPayloadLogGetResponseMaskingLevelPartial DLPPayloadLogGetResponseMaskingLevel = "partial"
	DLPPayloadLogGetResponseMaskingLevelClear   DLPPayloadLogGetResponseMaskingLevel = "clear"
	DLPPayloadLogGetResponseMaskingLevelDefault DLPPayloadLogGetResponseMaskingLevel = "default"
)

func (r DLPPayloadLogGetResponseMaskingLevel) IsKnown() bool {
	switch r {
	case DLPPayloadLogGetResponseMaskingLevelFull, DLPPayloadLogGetResponseMaskingLevelPartial, DLPPayloadLogGetResponseMaskingLevelClear, DLPPayloadLogGetResponseMaskingLevelDefault:
		return true
	}
	return false
}

type DLPPayloadLogUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Masking level for payload logs.
	//
	// - `full`: The entire payload is masked.
	// - `partial`: Only partial payload content is masked.
	// - `clear`: No masking is applied to the payload content.
	// - `default`: DLP uses its default masking behavior.
	MaskingLevel param.Field[DLPPayloadLogUpdateParamsMaskingLevel] `json:"masking_level"`
	// Base64-encoded public key for encrypting payload logs.
	//
	// - Set to null or empty string to disable payload logging.
	// - Set to a non-empty base64 string to enable payload logging with the given key.
	//
	// For customers with configurable payload masking feature rolled out:
	//
	//   - If the field is missing, the existing setting will be kept. Note that this is
	//     different from setting to null or empty string.
	//
	// For all other customers:
	//
	// - If the field is missing, the existing setting will be cleared.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DLPPayloadLogUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Masking level for payload logs.
//
// - `full`: The entire payload is masked.
// - `partial`: Only partial payload content is masked.
// - `clear`: No masking is applied to the payload content.
// - `default`: DLP uses its default masking behavior.
type DLPPayloadLogUpdateParamsMaskingLevel string

const (
	DLPPayloadLogUpdateParamsMaskingLevelFull    DLPPayloadLogUpdateParamsMaskingLevel = "full"
	DLPPayloadLogUpdateParamsMaskingLevelPartial DLPPayloadLogUpdateParamsMaskingLevel = "partial"
	DLPPayloadLogUpdateParamsMaskingLevelClear   DLPPayloadLogUpdateParamsMaskingLevel = "clear"
	DLPPayloadLogUpdateParamsMaskingLevelDefault DLPPayloadLogUpdateParamsMaskingLevel = "default"
)

func (r DLPPayloadLogUpdateParamsMaskingLevel) IsKnown() bool {
	switch r {
	case DLPPayloadLogUpdateParamsMaskingLevelFull, DLPPayloadLogUpdateParamsMaskingLevelPartial, DLPPayloadLogUpdateParamsMaskingLevelClear, DLPPayloadLogUpdateParamsMaskingLevelDefault:
		return true
	}
	return false
}

type DLPPayloadLogUpdateResponseEnvelope struct {
	Errors   []DLPPayloadLogUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPPayloadLogUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPPayloadLogUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPPayloadLogUpdateResponse                `json:"result"`
	JSON    dlpPayloadLogUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPPayloadLogUpdateResponseEnvelope]
type dlpPayloadLogUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogUpdateResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DLPPayloadLogUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpPayloadLogUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPPayloadLogUpdateResponseEnvelopeErrors]
type dlpPayloadLogUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dlpPayloadLogUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPPayloadLogUpdateResponseEnvelopeErrorsSource]
type dlpPayloadLogUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogUpdateResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DLPPayloadLogUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpPayloadLogUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPPayloadLogUpdateResponseEnvelopeMessages]
type dlpPayloadLogUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    dlpPayloadLogUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPPayloadLogUpdateResponseEnvelopeMessagesSource]
type dlpPayloadLogUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPPayloadLogUpdateResponseEnvelopeSuccess bool

const (
	DLPPayloadLogUpdateResponseEnvelopeSuccessTrue DLPPayloadLogUpdateResponseEnvelopeSuccess = true
)

func (r DLPPayloadLogUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPPayloadLogUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPPayloadLogGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPPayloadLogGetResponseEnvelope struct {
	Errors   []DLPPayloadLogGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DLPPayloadLogGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DLPPayloadLogGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPPayloadLogGetResponse                `json:"result"`
	JSON    dlpPayloadLogGetResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponseEnvelope]
type dlpPayloadLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPPayloadLogGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpPayloadLogGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPPayloadLogGetResponseEnvelopeErrors]
type dlpPayloadLogGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpPayloadLogGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPPayloadLogGetResponseEnvelopeErrorsSource]
type dlpPayloadLogGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPPayloadLogGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpPayloadLogGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPPayloadLogGetResponseEnvelopeMessages]
type dlpPayloadLogGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpPayloadLogGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPPayloadLogGetResponseEnvelopeMessagesSource]
type dlpPayloadLogGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPPayloadLogGetResponseEnvelopeSuccess bool

const (
	DLPPayloadLogGetResponseEnvelopeSuccessTrue DLPPayloadLogGetResponseEnvelopeSuccess = true
)

func (r DLPPayloadLogGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPPayloadLogGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
