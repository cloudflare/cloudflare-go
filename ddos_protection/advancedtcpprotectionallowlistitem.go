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

// AdvancedTCPProtectionAllowlistItemService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionAllowlistItemService] method instead.
type AdvancedTCPProtectionAllowlistItemService struct {
	Options []option.RequestOption
}

// NewAdvancedTCPProtectionAllowlistItemService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionAllowlistItemService(opts ...option.RequestOption) (r *AdvancedTCPProtectionAllowlistItemService) {
	r = &AdvancedTCPProtectionAllowlistItemService{}
	r.Options = opts
	return
}

// Delete the allowlist prefix for an account given a UUID.
func (r *AdvancedTCPProtectionAllowlistItemService) Delete(ctx context.Context, prefixID string, body AdvancedTCPProtectionAllowlistItemDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionAllowlistItemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/allowlist/%s", body.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Update an allowlist prefix specified by the given UUID.
func (r *AdvancedTCPProtectionAllowlistItemService) Edit(ctx context.Context, prefixID string, params AdvancedTCPProtectionAllowlistItemEditParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionAllowlistItemEditResponse, err error) {
	var env AdvancedTCPProtectionAllowlistItemEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/allowlist/%s", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get an allowlist prefix specified by the given UUID.
func (r *AdvancedTCPProtectionAllowlistItemService) Get(ctx context.Context, prefixID string, query AdvancedTCPProtectionAllowlistItemGetParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionAllowlistItemGetResponse, err error) {
	var env AdvancedTCPProtectionAllowlistItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/allowlist/%s", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AdvancedTCPProtectionAllowlistItemDeleteResponse struct {
	Errors   []AdvancedTCPProtectionAllowlistItemDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionAllowlistItemDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionAllowlistItemDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionAllowlistItemDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistItemDeleteResponseJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionAllowlistItemDeleteResponse]
type advancedTCPProtectionAllowlistItemDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemDeleteResponseError struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistItemDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistItemDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistItemDeleteResponseErrorJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionAllowlistItemDeleteResponseError]
type advancedTCPProtectionAllowlistItemDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemDeleteResponseErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistItemDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemDeleteResponseErrorsSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemDeleteResponseErrorsSource]
type advancedTCPProtectionAllowlistItemDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemDeleteResponseMessage struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistItemDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistItemDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistItemDeleteResponseMessageJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionAllowlistItemDeleteResponseMessage]
type advancedTCPProtectionAllowlistItemDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemDeleteResponseMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistItemDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemDeleteResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemDeleteResponseMessagesSource]
type advancedTCPProtectionAllowlistItemDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionAllowlistItemDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionAllowlistItemDeleteResponseSuccessTrue AdvancedTCPProtectionAllowlistItemDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionAllowlistItemDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionAllowlistItemDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionAllowlistItemEditResponse struct {
	// The unique ID of the allowlist prefix.
	ID string `json:"id" api:"required"`
	// An optional comment describing the allowlist prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the allowlist prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to enable the allowlist prefix into effect. Defaults to false.
	Enabled bool `json:"enabled" api:"required"`
	// The last modification timestamp of the allowlist prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The allowlist prefix in CIDR format.
	Prefix string                                             `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionAllowlistItemEditResponseJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemEditResponseJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionAllowlistItemEditResponse]
type advancedTCPProtectionAllowlistItemEditResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemEditResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemGetResponse struct {
	// The unique ID of the allowlist prefix.
	ID string `json:"id" api:"required"`
	// An optional comment describing the allowlist prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the allowlist prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to enable the allowlist prefix into effect. Defaults to false.
	Enabled bool `json:"enabled" api:"required"`
	// The last modification timestamp of the allowlist prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The allowlist prefix in CIDR format.
	Prefix string                                            `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionAllowlistItemGetResponseJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemGetResponseJSON contains the JSON metadata for
// the struct [AdvancedTCPProtectionAllowlistItemGetResponse]
type advancedTCPProtectionAllowlistItemGetResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionAllowlistItemEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A comment describing the allowlist prefix. Optional.
	Comment param.Field[string] `json:"comment"`
	// Whether to enable the allowlist prefix into effect. Optional.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AdvancedTCPProtectionAllowlistItemEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionAllowlistItemEditResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionAllowlistItemEditResponse                `json:"result"`
	JSON    advancedTCPProtectionAllowlistItemEditResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistItemEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionAllowlistItemEditResponseEnvelope]
type advancedTCPProtectionAllowlistItemEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrors struct {
	Code             int64                                                              `json:"code" api:"required"`
	Message          string                                                             `json:"message" api:"required"`
	DocumentationURL string                                                             `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrors]
type advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                                 `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSource]
type advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessages struct {
	Code             int64                                                                `json:"code" api:"required"`
	Message          string                                                               `json:"message" api:"required"`
	DocumentationURL string                                                               `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessages]
type advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                                   `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSource]
type advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeSuccessTrue AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionAllowlistItemEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionAllowlistItemGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionAllowlistItemGetResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionAllowlistItemGetResponse                `json:"result"`
	JSON    advancedTCPProtectionAllowlistItemGetResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistItemGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionAllowlistItemGetResponseEnvelope]
type advancedTCPProtectionAllowlistItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrors struct {
	Code             int64                                                             `json:"code" api:"required"`
	Message          string                                                            `json:"message" api:"required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrors]
type advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSource]
type advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessages struct {
	Code             int64                                                               `json:"code" api:"required"`
	Message          string                                                              `json:"message" api:"required"`
	DocumentationURL string                                                              `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessages]
type advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                                  `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSource]
type advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistItemGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeSuccessTrue AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionAllowlistItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
