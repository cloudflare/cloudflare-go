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

// AdvancedTCPProtectionPrefixItemService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionPrefixItemService] method instead.
type AdvancedTCPProtectionPrefixItemService struct {
	Options []option.RequestOption
}

// NewAdvancedTCPProtectionPrefixItemService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAdvancedTCPProtectionPrefixItemService(opts ...option.RequestOption) (r *AdvancedTCPProtectionPrefixItemService) {
	r = &AdvancedTCPProtectionPrefixItemService{}
	r.Options = opts
	return
}

// Delete the prefix for an account given a UUID.
func (r *AdvancedTCPProtectionPrefixItemService) Delete(ctx context.Context, prefixID string, body AdvancedTCPProtectionPrefixItemDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionPrefixItemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/prefixes/%s", body.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Update a prefix specified by the given UUID.
func (r *AdvancedTCPProtectionPrefixItemService) Edit(ctx context.Context, prefixID string, params AdvancedTCPProtectionPrefixItemEditParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionPrefixItemEditResponse, err error) {
	var env AdvancedTCPProtectionPrefixItemEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/prefixes/%s", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get a prefix specified by the given UUID.
func (r *AdvancedTCPProtectionPrefixItemService) Get(ctx context.Context, prefixID string, query AdvancedTCPProtectionPrefixItemGetParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionPrefixItemGetResponse, err error) {
	var env AdvancedTCPProtectionPrefixItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/prefixes/%s", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AdvancedTCPProtectionPrefixItemDeleteResponse struct {
	Errors   []AdvancedTCPProtectionPrefixItemDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionPrefixItemDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionPrefixItemDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionPrefixItemDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionPrefixItemDeleteResponseJSON contains the JSON metadata for
// the struct [AdvancedTCPProtectionPrefixItemDeleteResponse]
type advancedTCPProtectionPrefixItemDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemDeleteResponseError struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixItemDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionPrefixItemDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionPrefixItemDeleteResponseErrorJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixItemDeleteResponseError]
type advancedTCPProtectionPrefixItemDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemDeleteResponseErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    advancedTCPProtectionPrefixItemDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemDeleteResponseErrorsSourceJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionPrefixItemDeleteResponseErrorsSource]
type advancedTCPProtectionPrefixItemDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemDeleteResponseMessage struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixItemDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionPrefixItemDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionPrefixItemDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixItemDeleteResponseMessage]
type advancedTCPProtectionPrefixItemDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemDeleteResponseMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    advancedTCPProtectionPrefixItemDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemDeleteResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionPrefixItemDeleteResponseMessagesSource]
type advancedTCPProtectionPrefixItemDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionPrefixItemDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionPrefixItemDeleteResponseSuccessTrue AdvancedTCPProtectionPrefixItemDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionPrefixItemDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionPrefixItemDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionPrefixItemEditResponse struct {
	// The unique ID of the prefix.
	ID string `json:"id" api:"required"`
	// A comment describing the prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to exclude the prefix from protection.
	Excluded bool `json:"excluded" api:"required"`
	// The last modification timestamp of the prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The prefix in CIDR format.
	Prefix string                                          `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionPrefixItemEditResponseJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemEditResponseJSON contains the JSON metadata for
// the struct [AdvancedTCPProtectionPrefixItemEditResponse]
type advancedTCPProtectionPrefixItemEditResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Excluded    apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemEditResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemGetResponse struct {
	// The unique ID of the prefix.
	ID string `json:"id" api:"required"`
	// A comment describing the prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to exclude the prefix from protection.
	Excluded bool `json:"excluded" api:"required"`
	// The last modification timestamp of the prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The prefix in CIDR format.
	Prefix string                                         `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionPrefixItemGetResponseJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemGetResponseJSON contains the JSON metadata for
// the struct [AdvancedTCPProtectionPrefixItemGetResponse]
type advancedTCPProtectionPrefixItemGetResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Excluded    apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionPrefixItemEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A new comment for the prefix. Optional.
	Comment param.Field[string] `json:"comment"`
	// Whether to exclude the prefix from protection. Optional.
	Excluded param.Field[bool] `json:"excluded"`
}

func (r AdvancedTCPProtectionPrefixItemEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionPrefixItemEditResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionPrefixItemEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionPrefixItemEditResponse                `json:"result"`
	JSON    advancedTCPProtectionPrefixItemEditResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionPrefixItemEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixItemEditResponseEnvelope]
type advancedTCPProtectionPrefixItemEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrors struct {
	Code             int64                                                           `json:"code" api:"required"`
	Message          string                                                          `json:"message" api:"required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrors]
type advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSource]
type advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessages struct {
	Code             int64                                                             `json:"code" api:"required"`
	Message          string                                                            `json:"message" api:"required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessages]
type advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSource]
type advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionPrefixItemEditResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionPrefixItemEditResponseEnvelopeSuccessTrue AdvancedTCPProtectionPrefixItemEditResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionPrefixItemEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionPrefixItemEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionPrefixItemGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionPrefixItemGetResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionPrefixItemGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionPrefixItemGetResponse                `json:"result"`
	JSON    advancedTCPProtectionPrefixItemGetResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionPrefixItemGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixItemGetResponseEnvelope]
type advancedTCPProtectionPrefixItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrors]
type advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSource]
type advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code" api:"required"`
	Message          string                                                           `json:"message" api:"required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessages]
type advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSource]
type advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixItemGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionPrefixItemGetResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionPrefixItemGetResponseEnvelopeSuccessTrue AdvancedTCPProtectionPrefixItemGetResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionPrefixItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionPrefixItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
