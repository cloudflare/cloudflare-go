// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package field_extractors

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// FieldExtractorService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFieldExtractorService] method instead.
type FieldExtractorService struct {
	Options []option.RequestOption
}

// NewFieldExtractorService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFieldExtractorService(opts ...option.RequestOption) (r *FieldExtractorService) {
	r = &FieldExtractorService{}
	r.Options = opts
	return
}

// Replaces all custom extraction rules for an extractor type. Omitted rules are
// deleted.
func (r *FieldExtractorService) Update(ctx context.Context, extractor string, params FieldExtractorUpdateParams, opts ...option.RequestOption) (res *FieldExtractorUpdateResponse, err error) {
	var env FieldExtractorUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if extractor == "" {
		err = errors.New("missing required extractor parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/field_extractors/%s", params.AccountID, extractor)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes all custom extraction rules for an extractor type.
func (r *FieldExtractorService) Delete(ctx context.Context, extractor string, body FieldExtractorDeleteParams, opts ...option.RequestOption) (res *FieldExtractorDeleteResponse, err error) {
	var env FieldExtractorDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if extractor == "" {
		err = errors.New("missing required extractor parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/field_extractors/%s", body.AccountID, extractor)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the custom extraction rules configured for a given extractor type.
func (r *FieldExtractorService) Get(ctx context.Context, extractor string, query FieldExtractorGetParams, opts ...option.RequestOption) (res *FieldExtractorGetResponse, err error) {
	var env FieldExtractorGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if extractor == "" {
		err = errors.New("missing required extractor parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/field_extractors/%s", query.AccountID, extractor)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type FieldExtractorUpdateResponse struct {
	// Extractor type.
	Extractor string                             `json:"extractor" api:"required"`
	Rules     []FieldExtractorUpdateResponseRule `json:"rules" api:"required"`
	JSON      fieldExtractorUpdateResponseJSON   `json:"-"`
}

// fieldExtractorUpdateResponseJSON contains the JSON metadata for the struct
// [FieldExtractorUpdateResponse]
type fieldExtractorUpdateResponseJSON struct {
	Extractor   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorUpdateResponseRule struct {
	Fields []FieldExtractorUpdateResponseRulesField `json:"fields" api:"required"`
	// Stable rule identifier.
	Ref string `json:"ref" api:"required"`
	// Human-readable rule description.
	Description string                               `json:"description"`
	JSON        fieldExtractorUpdateResponseRuleJSON `json:"-"`
}

// fieldExtractorUpdateResponseRuleJSON contains the JSON metadata for the struct
// [FieldExtractorUpdateResponseRule]
type fieldExtractorUpdateResponseRuleJSON struct {
	Fields      apijson.Field
	Ref         apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorUpdateResponseRulesField struct {
	// Wirefilter value expression.
	Expression string `json:"expression" api:"required"`
	// Field name.
	Name string                                     `json:"name" api:"required"`
	JSON fieldExtractorUpdateResponseRulesFieldJSON `json:"-"`
}

// fieldExtractorUpdateResponseRulesFieldJSON contains the JSON metadata for the
// struct [FieldExtractorUpdateResponseRulesField]
type fieldExtractorUpdateResponseRulesFieldJSON struct {
	Expression  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorUpdateResponseRulesField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorUpdateResponseRulesFieldJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorDeleteResponse = interface{}

type FieldExtractorGetResponse struct {
	// Extractor type.
	Extractor string                          `json:"extractor" api:"required"`
	Rules     []FieldExtractorGetResponseRule `json:"rules" api:"required"`
	JSON      fieldExtractorGetResponseJSON   `json:"-"`
}

// fieldExtractorGetResponseJSON contains the JSON metadata for the struct
// [FieldExtractorGetResponse]
type fieldExtractorGetResponseJSON struct {
	Extractor   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorGetResponseJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorGetResponseRule struct {
	Fields []FieldExtractorGetResponseRulesField `json:"fields" api:"required"`
	// Stable rule identifier.
	Ref string `json:"ref" api:"required"`
	// Human-readable rule description.
	Description string                            `json:"description"`
	JSON        fieldExtractorGetResponseRuleJSON `json:"-"`
}

// fieldExtractorGetResponseRuleJSON contains the JSON metadata for the struct
// [FieldExtractorGetResponseRule]
type fieldExtractorGetResponseRuleJSON struct {
	Fields      apijson.Field
	Ref         apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorGetResponseRulesField struct {
	// Wirefilter value expression.
	Expression string `json:"expression" api:"required"`
	// Field name.
	Name string                                  `json:"name" api:"required"`
	JSON fieldExtractorGetResponseRulesFieldJSON `json:"-"`
}

// fieldExtractorGetResponseRulesFieldJSON contains the JSON metadata for the
// struct [FieldExtractorGetResponseRulesField]
type fieldExtractorGetResponseRulesFieldJSON struct {
	Expression  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorGetResponseRulesField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorGetResponseRulesFieldJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorUpdateParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string]                           `path:"account_id" api:"required"`
	Rules     param.Field[[]FieldExtractorUpdateParamsRule] `json:"rules" api:"required"`
}

func (r FieldExtractorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FieldExtractorUpdateParamsRule struct {
	Fields      param.Field[[]FieldExtractorUpdateParamsRulesField] `json:"fields" api:"required"`
	Ref         param.Field[string]                                 `json:"ref" api:"required"`
	Description param.Field[string]                                 `json:"description"`
}

func (r FieldExtractorUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FieldExtractorUpdateParamsRulesField struct {
	Expression param.Field[string] `json:"expression" api:"required"`
	Name       param.Field[string] `json:"name" api:"required"`
}

func (r FieldExtractorUpdateParamsRulesField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FieldExtractorUpdateResponseEnvelope struct {
	Errors []FieldExtractorUpdateResponseEnvelopeErrors `json:"errors" api:"required"`
	// Additional informational messages.
	Messages []FieldExtractorUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   FieldExtractorUpdateResponse                   `json:"result" api:"required"`
	Success  FieldExtractorUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     fieldExtractorUpdateResponseEnvelopeJSON       `json:"-"`
}

// fieldExtractorUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FieldExtractorUpdateResponseEnvelope]
type fieldExtractorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorUpdateResponseEnvelopeErrors struct {
	Code    float64 `json:"code" api:"required"`
	Message string  `json:"message" api:"required"`
	// Points to the offending input as a JSON Pointer into the request body.
	Source FieldExtractorUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   fieldExtractorUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// fieldExtractorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FieldExtractorUpdateResponseEnvelopeErrors]
type fieldExtractorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Points to the offending input as a JSON Pointer into the request body.
type FieldExtractorUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    fieldExtractorUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fieldExtractorUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [FieldExtractorUpdateResponseEnvelopeErrorsSource]
type fieldExtractorUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorUpdateResponseEnvelopeMessages struct {
	Code    float64                                          `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	JSON    fieldExtractorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// fieldExtractorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FieldExtractorUpdateResponseEnvelopeMessages]
type fieldExtractorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorUpdateResponseEnvelopeSuccess bool

const (
	FieldExtractorUpdateResponseEnvelopeSuccessTrue FieldExtractorUpdateResponseEnvelopeSuccess = true
)

func (r FieldExtractorUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FieldExtractorUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FieldExtractorDeleteParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type FieldExtractorDeleteResponseEnvelope struct {
	Errors []FieldExtractorDeleteResponseEnvelopeErrors `json:"errors" api:"required"`
	// Additional informational messages.
	Messages []FieldExtractorDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Result is null.
	Result  FieldExtractorDeleteResponse                `json:"result" api:"required,nullable"`
	Success FieldExtractorDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    fieldExtractorDeleteResponseEnvelopeJSON    `json:"-"`
}

// fieldExtractorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FieldExtractorDeleteResponseEnvelope]
type fieldExtractorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorDeleteResponseEnvelopeErrors struct {
	Code    float64 `json:"code" api:"required"`
	Message string  `json:"message" api:"required"`
	// Points to the offending input as a JSON Pointer into the request body.
	Source FieldExtractorDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   fieldExtractorDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// fieldExtractorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FieldExtractorDeleteResponseEnvelopeErrors]
type fieldExtractorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Points to the offending input as a JSON Pointer into the request body.
type FieldExtractorDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer" api:"required"`
	JSON    fieldExtractorDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fieldExtractorDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [FieldExtractorDeleteResponseEnvelopeErrorsSource]
type fieldExtractorDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorDeleteResponseEnvelopeMessages struct {
	Code    float64                                          `json:"code" api:"required"`
	Message string                                           `json:"message" api:"required"`
	JSON    fieldExtractorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// fieldExtractorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FieldExtractorDeleteResponseEnvelopeMessages]
type fieldExtractorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorDeleteResponseEnvelopeSuccess bool

const (
	FieldExtractorDeleteResponseEnvelopeSuccessTrue FieldExtractorDeleteResponseEnvelopeSuccess = true
)

func (r FieldExtractorDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FieldExtractorDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FieldExtractorGetParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type FieldExtractorGetResponseEnvelope struct {
	Errors []FieldExtractorGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// Additional informational messages.
	Messages []FieldExtractorGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   FieldExtractorGetResponse                   `json:"result" api:"required"`
	Success  FieldExtractorGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     fieldExtractorGetResponseEnvelopeJSON       `json:"-"`
}

// fieldExtractorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FieldExtractorGetResponseEnvelope]
type fieldExtractorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorGetResponseEnvelopeErrors struct {
	Code    float64 `json:"code" api:"required"`
	Message string  `json:"message" api:"required"`
	// Points to the offending input as a JSON Pointer into the request body.
	Source FieldExtractorGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   fieldExtractorGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// fieldExtractorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FieldExtractorGetResponseEnvelopeErrors]
type fieldExtractorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Points to the offending input as a JSON Pointer into the request body.
type FieldExtractorGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    fieldExtractorGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fieldExtractorGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [FieldExtractorGetResponseEnvelopeErrorsSource]
type fieldExtractorGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorGetResponseEnvelopeMessages struct {
	Code    float64                                       `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	JSON    fieldExtractorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// fieldExtractorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FieldExtractorGetResponseEnvelopeMessages]
type fieldExtractorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FieldExtractorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fieldExtractorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FieldExtractorGetResponseEnvelopeSuccess bool

const (
	FieldExtractorGetResponseEnvelopeSuccessTrue FieldExtractorGetResponseEnvelopeSuccess = true
)

func (r FieldExtractorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FieldExtractorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
