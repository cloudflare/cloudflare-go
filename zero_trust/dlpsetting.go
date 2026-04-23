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

// DLPSettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPSettingService] method instead.
type DLPSettingService struct {
	Options []option.RequestOption
}

// NewDLPSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPSettingService(opts ...option.RequestOption) (r *DLPSettingService) {
	r = &DLPSettingService{}
	r.Options = opts
	return
}

// Missing fields are reset to initial (unconfigured) values.
func (r *DLPSettingService) Update(ctx context.Context, params DLPSettingUpdateParams, opts ...option.RequestOption) (res *DLPSettings, err error) {
	var env DLPSettingUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete (reset) DLP account-level settings to initial values.
func (r *DLPSettingService) Delete(ctx context.Context, body DLPSettingDeleteParams, opts ...option.RequestOption) (res *DLPSettings, err error) {
	var env DLPSettingDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/settings", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Missing fields keep their existing values.
func (r *DLPSettingService) Edit(ctx context.Context, params DLPSettingEditParams, opts ...option.RequestOption) (res *DLPSettings, err error) {
	var env DLPSettingEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get DLP account-level settings.
func (r *DLPSettingService) Get(ctx context.Context, query DLPSettingGetParams, opts ...option.RequestOption) (res *DLPSettings, err error) {
	var env DLPSettingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// DLP account-level settings response.
type DLPSettings struct {
	// Whether AI context analysis is enabled at the account level.
	AIContextAnalysis bool `json:"ai_context_analysis" api:"required"`
	// Whether OCR is enabled at the account level.
	OCR            bool                      `json:"ocr" api:"required"`
	PayloadLogging DLPSettingsPayloadLogging `json:"payload_logging" api:"required"`
	JSON           dlpSettingsJSON           `json:"-"`
}

// dlpSettingsJSON contains the JSON metadata for the struct [DLPSettings]
type dlpSettingsJSON struct {
	AIContextAnalysis apijson.Field
	OCR               apijson.Field
	PayloadLogging    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingsJSON) RawJSON() string {
	return r.raw
}

type DLPSettingsPayloadLogging struct {
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Masking level for payload logs.
	//
	// - `full`: The entire payload is masked.
	// - `partial`: Only partial payload content is masked.
	// - `clear`: No masking is applied to the payload content.
	// - `default`: DLP uses its default masking behavior.
	MaskingLevel DLPSettingsPayloadLoggingMaskingLevel `json:"masking_level"`
	// Base64-encoded public key for encrypting payload logs. Null when payload logging
	// is disabled.
	PublicKey string                        `json:"public_key" api:"nullable"`
	JSON      dlpSettingsPayloadLoggingJSON `json:"-"`
}

// dlpSettingsPayloadLoggingJSON contains the JSON metadata for the struct
// [DLPSettingsPayloadLogging]
type dlpSettingsPayloadLoggingJSON struct {
	UpdatedAt    apijson.Field
	MaskingLevel apijson.Field
	PublicKey    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DLPSettingsPayloadLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingsPayloadLoggingJSON) RawJSON() string {
	return r.raw
}

// Masking level for payload logs.
//
// - `full`: The entire payload is masked.
// - `partial`: Only partial payload content is masked.
// - `clear`: No masking is applied to the payload content.
// - `default`: DLP uses its default masking behavior.
type DLPSettingsPayloadLoggingMaskingLevel string

const (
	DLPSettingsPayloadLoggingMaskingLevelFull    DLPSettingsPayloadLoggingMaskingLevel = "full"
	DLPSettingsPayloadLoggingMaskingLevelPartial DLPSettingsPayloadLoggingMaskingLevel = "partial"
	DLPSettingsPayloadLoggingMaskingLevelClear   DLPSettingsPayloadLoggingMaskingLevel = "clear"
	DLPSettingsPayloadLoggingMaskingLevelDefault DLPSettingsPayloadLoggingMaskingLevel = "default"
)

func (r DLPSettingsPayloadLoggingMaskingLevel) IsKnown() bool {
	switch r {
	case DLPSettingsPayloadLoggingMaskingLevelFull, DLPSettingsPayloadLoggingMaskingLevelPartial, DLPSettingsPayloadLoggingMaskingLevelClear, DLPSettingsPayloadLoggingMaskingLevelDefault:
		return true
	}
	return false
}

type DLPSettingUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Whether AI context analysis is enabled at the account level.
	AIContextAnalysis param.Field[bool] `json:"ai_context_analysis"`
	// Whether OCR is enabled at the account level.
	OCR param.Field[bool] `json:"ocr"`
	// Request model for payload log settings within the DLP settings endpoint. Unlike
	// the legacy endpoint, null and missing are treated identically here (both mean
	// "not provided" for PATCH, "reset to default" for PUT).
	PayloadLogging param.Field[DLPSettingUpdateParamsPayloadLogging] `json:"payload_logging"`
}

func (r DLPSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Request model for payload log settings within the DLP settings endpoint. Unlike
// the legacy endpoint, null and missing are treated identically here (both mean
// "not provided" for PATCH, "reset to default" for PUT).
type DLPSettingUpdateParamsPayloadLogging struct {
	// Masking level for payload logs.
	//
	// - `full`: The entire payload is masked.
	// - `partial`: Only partial payload content is masked.
	// - `clear`: No masking is applied to the payload content.
	// - `default`: DLP uses its default masking behavior.
	MaskingLevel param.Field[DLPSettingUpdateParamsPayloadLoggingMaskingLevel] `json:"masking_level"`
	// Base64-encoded public key for encrypting payload logs.
	//
	// - Set to a non-empty base64 string to enable payload logging with the given key.
	// - Set to an empty string to disable payload logging.
	// - Omit or set to null to leave unchanged (PATCH) or reset to disabled (PUT).
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DLPSettingUpdateParamsPayloadLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Masking level for payload logs.
//
// - `full`: The entire payload is masked.
// - `partial`: Only partial payload content is masked.
// - `clear`: No masking is applied to the payload content.
// - `default`: DLP uses its default masking behavior.
type DLPSettingUpdateParamsPayloadLoggingMaskingLevel string

const (
	DLPSettingUpdateParamsPayloadLoggingMaskingLevelFull    DLPSettingUpdateParamsPayloadLoggingMaskingLevel = "full"
	DLPSettingUpdateParamsPayloadLoggingMaskingLevelPartial DLPSettingUpdateParamsPayloadLoggingMaskingLevel = "partial"
	DLPSettingUpdateParamsPayloadLoggingMaskingLevelClear   DLPSettingUpdateParamsPayloadLoggingMaskingLevel = "clear"
	DLPSettingUpdateParamsPayloadLoggingMaskingLevelDefault DLPSettingUpdateParamsPayloadLoggingMaskingLevel = "default"
)

func (r DLPSettingUpdateParamsPayloadLoggingMaskingLevel) IsKnown() bool {
	switch r {
	case DLPSettingUpdateParamsPayloadLoggingMaskingLevelFull, DLPSettingUpdateParamsPayloadLoggingMaskingLevelPartial, DLPSettingUpdateParamsPayloadLoggingMaskingLevelClear, DLPSettingUpdateParamsPayloadLoggingMaskingLevelDefault:
		return true
	}
	return false
}

type DLPSettingUpdateResponseEnvelope struct {
	Errors   []DLPSettingUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSettingUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSettingUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// DLP account-level settings response.
	Result DLPSettings                          `json:"result"`
	JSON   dlpSettingUpdateResponseEnvelopeJSON `json:"-"`
}

// dlpSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPSettingUpdateResponseEnvelope]
type dlpSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSettingUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPSettingUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSettingUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSettingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPSettingUpdateResponseEnvelopeErrors]
type dlpSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSettingUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpSettingUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSettingUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPSettingUpdateResponseEnvelopeErrorsSource]
type dlpSettingUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSettingUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPSettingUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSettingUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSettingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPSettingUpdateResponseEnvelopeMessages]
type dlpSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSettingUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpSettingUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSettingUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPSettingUpdateResponseEnvelopeMessagesSource]
type dlpSettingUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPSettingUpdateResponseEnvelopeSuccess bool

const (
	DLPSettingUpdateResponseEnvelopeSuccessTrue DLPSettingUpdateResponseEnvelopeSuccess = true
)

func (r DLPSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSettingDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSettingDeleteResponseEnvelope struct {
	Errors   []DLPSettingDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSettingDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSettingDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	// DLP account-level settings response.
	Result DLPSettings                          `json:"result"`
	JSON   dlpSettingDeleteResponseEnvelopeJSON `json:"-"`
}

// dlpSettingDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPSettingDeleteResponseEnvelope]
type dlpSettingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSettingDeleteResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPSettingDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSettingDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSettingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPSettingDeleteResponseEnvelopeErrors]
type dlpSettingDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSettingDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpSettingDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSettingDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPSettingDeleteResponseEnvelopeErrorsSource]
type dlpSettingDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSettingDeleteResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPSettingDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSettingDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSettingDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPSettingDeleteResponseEnvelopeMessages]
type dlpSettingDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSettingDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpSettingDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSettingDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPSettingDeleteResponseEnvelopeMessagesSource]
type dlpSettingDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPSettingDeleteResponseEnvelopeSuccess bool

const (
	DLPSettingDeleteResponseEnvelopeSuccessTrue DLPSettingDeleteResponseEnvelopeSuccess = true
)

func (r DLPSettingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSettingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSettingEditParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Whether AI context analysis is enabled at the account level.
	AIContextAnalysis param.Field[bool] `json:"ai_context_analysis"`
	// Whether OCR is enabled at the account level.
	OCR param.Field[bool] `json:"ocr"`
	// Request model for payload log settings within the DLP settings endpoint. Unlike
	// the legacy endpoint, null and missing are treated identically here (both mean
	// "not provided" for PATCH, "reset to default" for PUT).
	PayloadLogging param.Field[DLPSettingEditParamsPayloadLogging] `json:"payload_logging"`
}

func (r DLPSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Request model for payload log settings within the DLP settings endpoint. Unlike
// the legacy endpoint, null and missing are treated identically here (both mean
// "not provided" for PATCH, "reset to default" for PUT).
type DLPSettingEditParamsPayloadLogging struct {
	// Masking level for payload logs.
	//
	// - `full`: The entire payload is masked.
	// - `partial`: Only partial payload content is masked.
	// - `clear`: No masking is applied to the payload content.
	// - `default`: DLP uses its default masking behavior.
	MaskingLevel param.Field[DLPSettingEditParamsPayloadLoggingMaskingLevel] `json:"masking_level"`
	// Base64-encoded public key for encrypting payload logs.
	//
	// - Set to a non-empty base64 string to enable payload logging with the given key.
	// - Set to an empty string to disable payload logging.
	// - Omit or set to null to leave unchanged (PATCH) or reset to disabled (PUT).
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DLPSettingEditParamsPayloadLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Masking level for payload logs.
//
// - `full`: The entire payload is masked.
// - `partial`: Only partial payload content is masked.
// - `clear`: No masking is applied to the payload content.
// - `default`: DLP uses its default masking behavior.
type DLPSettingEditParamsPayloadLoggingMaskingLevel string

const (
	DLPSettingEditParamsPayloadLoggingMaskingLevelFull    DLPSettingEditParamsPayloadLoggingMaskingLevel = "full"
	DLPSettingEditParamsPayloadLoggingMaskingLevelPartial DLPSettingEditParamsPayloadLoggingMaskingLevel = "partial"
	DLPSettingEditParamsPayloadLoggingMaskingLevelClear   DLPSettingEditParamsPayloadLoggingMaskingLevel = "clear"
	DLPSettingEditParamsPayloadLoggingMaskingLevelDefault DLPSettingEditParamsPayloadLoggingMaskingLevel = "default"
)

func (r DLPSettingEditParamsPayloadLoggingMaskingLevel) IsKnown() bool {
	switch r {
	case DLPSettingEditParamsPayloadLoggingMaskingLevelFull, DLPSettingEditParamsPayloadLoggingMaskingLevelPartial, DLPSettingEditParamsPayloadLoggingMaskingLevelClear, DLPSettingEditParamsPayloadLoggingMaskingLevelDefault:
		return true
	}
	return false
}

type DLPSettingEditResponseEnvelope struct {
	Errors   []DLPSettingEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSettingEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSettingEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// DLP account-level settings response.
	Result DLPSettings                        `json:"result"`
	JSON   dlpSettingEditResponseEnvelopeJSON `json:"-"`
}

// dlpSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPSettingEditResponseEnvelope]
type dlpSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSettingEditResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DLPSettingEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSettingEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPSettingEditResponseEnvelopeErrors]
type dlpSettingEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSettingEditResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dlpSettingEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSettingEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPSettingEditResponseEnvelopeErrorsSource]
type dlpSettingEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSettingEditResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DLPSettingEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSettingEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPSettingEditResponseEnvelopeMessages]
type dlpSettingEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSettingEditResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dlpSettingEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSettingEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPSettingEditResponseEnvelopeMessagesSource]
type dlpSettingEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPSettingEditResponseEnvelopeSuccess bool

const (
	DLPSettingEditResponseEnvelopeSuccessTrue DLPSettingEditResponseEnvelopeSuccess = true
)

func (r DLPSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSettingGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSettingGetResponseEnvelope struct {
	Errors   []DLPSettingGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSettingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSettingGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// DLP account-level settings response.
	Result DLPSettings                       `json:"result"`
	JSON   dlpSettingGetResponseEnvelopeJSON `json:"-"`
}

// dlpSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPSettingGetResponseEnvelope]
type dlpSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSettingGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DLPSettingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSettingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPSettingGetResponseEnvelopeErrors]
type dlpSettingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSettingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dlpSettingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSettingGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPSettingGetResponseEnvelopeErrorsSource]
type dlpSettingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSettingGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPSettingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSettingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPSettingGetResponseEnvelopeMessages]
type dlpSettingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSettingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpSettingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSettingGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPSettingGetResponseEnvelopeMessagesSource]
type dlpSettingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSettingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSettingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPSettingGetResponseEnvelopeSuccess bool

const (
	DLPSettingGetResponseEnvelopeSuccessTrue DLPSettingGetResponseEnvelopeSuccess = true
)

func (r DLPSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
