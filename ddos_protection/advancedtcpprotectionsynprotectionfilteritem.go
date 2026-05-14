// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ddos_protection

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
)

// AdvancedTCPProtectionSynProtectionFilterItemService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionSynProtectionFilterItemService] method instead.
type AdvancedTCPProtectionSynProtectionFilterItemService struct {
	Options []option.RequestOption
}

// NewAdvancedTCPProtectionSynProtectionFilterItemService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionSynProtectionFilterItemService(opts ...option.RequestOption) (r *AdvancedTCPProtectionSynProtectionFilterItemService) {
	r = &AdvancedTCPProtectionSynProtectionFilterItemService{}
	r.Options = opts
	return
}

// Delete a SYN Protection filter specified by the given UUID.
func (r *AdvancedTCPProtectionSynProtectionFilterItemService) Delete(ctx context.Context, filterID string, body AdvancedTCPProtectionSynProtectionFilterItemDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionFilterItemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/filters/%s", body.AccountID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Update a SYN Protection filter specified by the given UUID.
func (r *AdvancedTCPProtectionSynProtectionFilterItemService) Edit(ctx context.Context, filterID string, params AdvancedTCPProtectionSynProtectionFilterItemEditParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionFilterItemEditResponse, err error) {
	var env AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/filters/%s", params.AccountID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get a SYN Protection filter specified by the given UUID.
func (r *AdvancedTCPProtectionSynProtectionFilterItemService) Get(ctx context.Context, filterID string, query AdvancedTCPProtectionSynProtectionFilterItemGetParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionFilterItemGetResponse, err error) {
	var env AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/filters/%s", query.AccountID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AdvancedTCPProtectionSynProtectionFilterItemDeleteResponse struct {
	Errors   []AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionSynProtectionFilterItemDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemDeleteResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemDeleteResponse]
type advancedTCPProtectionSynProtectionFilterItemDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseError struct {
	Code             int64                                                                  `json:"code" api:"required"`
	Message          string                                                                 `json:"message" api:"required"`
	DocumentationURL string                                                                 `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseError]
type advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSource struct {
	Pointer string                                                                     `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSource]
type advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessage struct {
	Code             int64                                                                    `json:"code" api:"required"`
	Message          string                                                                   `json:"message" api:"required"`
	DocumentationURL string                                                                   `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessageJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessage]
type advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSource struct {
	Pointer string                                                                       `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSource]
type advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseSuccessTrue AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionFilterItemDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionFilterItemEditResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                                    `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionSynProtectionFilterItemEditResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemEditResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemEditResponse]
type advancedTCPProtectionSynProtectionFilterItemEditResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemEditResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemGetResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                                   `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionSynProtectionFilterItemGetResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemGetResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemGetResponse]
type advancedTCPProtectionSynProtectionFilterItemGetResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionSynProtectionFilterItemEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The new filter expression. Optional.
	Expression param.Field[string] `json:"expression"`
	// The new mode for the filter. Optional. Must be one of 'enabled', 'disabled',
	// 'monitoring'.
	Mode param.Field[string] `json:"mode"`
}

func (r AdvancedTCPProtectionSynProtectionFilterItemEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionSynProtectionFilterItemEditResponse                `json:"result"`
	JSON    advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelope]
type advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrors struct {
	Code             int64                                                                        `json:"code" api:"required"`
	Message          string                                                                       `json:"message" api:"required"`
	DocumentationURL string                                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrors]
type advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                                           `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSource]
type advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessages struct {
	Code             int64                                                                          `json:"code" api:"required"`
	Message          string                                                                         `json:"message" api:"required"`
	DocumentationURL string                                                                         `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessages]
type advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                                             `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSource]
type advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeSuccessTrue AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionFilterItemEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionFilterItemGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionSynProtectionFilterItemGetResponse                `json:"result"`
	JSON    advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelope]
type advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrors struct {
	Code             int64                                                                       `json:"code" api:"required"`
	Message          string                                                                      `json:"message" api:"required"`
	DocumentationURL string                                                                      `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrors]
type advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                                          `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSource]
type advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessages struct {
	Code             int64                                                                         `json:"code" api:"required"`
	Message          string                                                                        `json:"message" api:"required"`
	DocumentationURL string                                                                        `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessages]
type advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                                            `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSource]
type advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeSuccessTrue AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionFilterItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
