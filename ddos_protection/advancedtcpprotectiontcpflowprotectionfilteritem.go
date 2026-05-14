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

// AdvancedTCPProtectionTCPFlowProtectionFilterItemService contains methods and
// other services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionTCPFlowProtectionFilterItemService] method instead.
type AdvancedTCPProtectionTCPFlowProtectionFilterItemService struct {
	Options []option.RequestOption
}

// NewAdvancedTCPProtectionTCPFlowProtectionFilterItemService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewAdvancedTCPProtectionTCPFlowProtectionFilterItemService(opts ...option.RequestOption) (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemService) {
	r = &AdvancedTCPProtectionTCPFlowProtectionFilterItemService{}
	r.Options = opts
	return
}

// Delete a TCP Flow Protection filter specified by the given UUID.
func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemService) Delete(ctx context.Context, filterID string, body AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters/%s", body.AccountID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Update a TCP Flow Protection filter specified by the given UUID.
func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemService) Edit(ctx context.Context, filterID string, params AdvancedTCPProtectionTCPFlowProtectionFilterItemEditParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponse, err error) {
	var env AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters/%s", params.AccountID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get a TCP Flow Protection filter specified by the given UUID.
func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemService) Get(ctx context.Context, filterID string, query AdvancedTCPProtectionTCPFlowProtectionFilterItemGetParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponse, err error) {
	var env AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters/%s", query.AccountID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponse struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponse]
type advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseError struct {
	Code             int64                                                                      `json:"code" api:"required"`
	Message          string                                                                     `json:"message" api:"required"`
	DocumentationURL string                                                                     `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseError]
type advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSource struct {
	Pointer string                                                                         `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSource]
type advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessage struct {
	Code             int64                                                                        `json:"code" api:"required"`
	Message          string                                                                       `json:"message" api:"required"`
	DocumentationURL string                                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessageJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessage]
type advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSource struct {
	Pointer string                                                                           `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSource]
type advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseSuccessTrue AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                                        `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponse]
type advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                                       `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponse]
type advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The new filter expression. Optional.
	Expression param.Field[string] `json:"expression"`
	// The new mode for the filter. Optional. Must be one of 'enabled', 'disabled',
	// 'monitoring'.
	Mode param.Field[string] `json:"mode"`
}

func (r AdvancedTCPProtectionTCPFlowProtectionFilterItemEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponse                `json:"result"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelope]
type advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrors struct {
	Code             int64                                                                            `json:"code" api:"required"`
	Message          string                                                                           `json:"message" api:"required"`
	DocumentationURL string                                                                           `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrors]
type advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                                               `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSource]
type advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessages struct {
	Code             int64                                                                              `json:"code" api:"required"`
	Message          string                                                                             `json:"message" api:"required"`
	DocumentationURL string                                                                             `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessages]
type advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                                                 `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSource]
type advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeSuccessTrue AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionFilterItemEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponse                `json:"result"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelope]
type advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrors struct {
	Code             int64                                                                           `json:"code" api:"required"`
	Message          string                                                                          `json:"message" api:"required"`
	DocumentationURL string                                                                          `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrors]
type advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                                              `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSource]
type advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessages struct {
	Code             int64                                                                             `json:"code" api:"required"`
	Message          string                                                                            `json:"message" api:"required"`
	DocumentationURL string                                                                            `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessages]
type advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                                                `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSource]
type advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeSuccessTrue AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionFilterItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
