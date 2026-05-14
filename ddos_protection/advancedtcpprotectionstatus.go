// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ddos_protection

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

// AdvancedTCPProtectionStatusService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionStatusService] method instead.
type AdvancedTCPProtectionStatusService struct {
	Options []option.RequestOption
}

// NewAdvancedTCPProtectionStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAdvancedTCPProtectionStatusService(opts ...option.RequestOption) (r *AdvancedTCPProtectionStatusService) {
	r = &AdvancedTCPProtectionStatusService{}
	r.Options = opts
	return
}

// Update the protection status of the account.
func (r *AdvancedTCPProtectionStatusService) Edit(ctx context.Context, params AdvancedTCPProtectionStatusEditParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionStatusEditResponse, err error) {
	var env AdvancedTCPProtectionStatusEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_protection_status", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get the protection status of the account.
func (r *AdvancedTCPProtectionStatusService) Get(ctx context.Context, query AdvancedTCPProtectionStatusGetParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionStatusGetResponse, err error) {
	var env AdvancedTCPProtectionStatusGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_protection_status", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AdvancedTCPProtectionStatusEditResponse struct {
	Enabled bool                                        `json:"enabled" api:"required"`
	JSON    advancedTCPProtectionStatusEditResponseJSON `json:"-"`
}

// advancedTCPProtectionStatusEditResponseJSON contains the JSON metadata for the
// struct [AdvancedTCPProtectionStatusEditResponse]
type advancedTCPProtectionStatusEditResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusEditResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusGetResponse struct {
	Enabled bool                                       `json:"enabled" api:"required"`
	JSON    advancedTCPProtectionStatusGetResponseJSON `json:"-"`
}

// advancedTCPProtectionStatusGetResponseJSON contains the JSON metadata for the
// struct [AdvancedTCPProtectionStatusGetResponse]
type advancedTCPProtectionStatusGetResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Enables or disables protection.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r AdvancedTCPProtectionStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionStatusEditResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionStatusEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionStatusEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionStatusEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionStatusEditResponse                `json:"result"`
	JSON    advancedTCPProtectionStatusEditResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionStatusEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionStatusEditResponseEnvelope]
type advancedTCPProtectionStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusEditResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AdvancedTCPProtectionStatusEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionStatusEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionStatusEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionStatusEditResponseEnvelopeErrors]
type advancedTCPProtectionStatusEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    advancedTCPProtectionStatusEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionStatusEditResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionStatusEditResponseEnvelopeErrorsSource]
type advancedTCPProtectionStatusEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusEditResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           AdvancedTCPProtectionStatusEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionStatusEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionStatusEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionStatusEditResponseEnvelopeMessages]
type advancedTCPProtectionStatusEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    advancedTCPProtectionStatusEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionStatusEditResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionStatusEditResponseEnvelopeMessagesSource]
type advancedTCPProtectionStatusEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionStatusEditResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionStatusEditResponseEnvelopeSuccessTrue AdvancedTCPProtectionStatusEditResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionStatusEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionStatusEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionStatusGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionStatusGetResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionStatusGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionStatusGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionStatusGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionStatusGetResponse                `json:"result"`
	JSON    advancedTCPProtectionStatusGetResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionStatusGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionStatusGetResponseEnvelope]
type advancedTCPProtectionStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusGetResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AdvancedTCPProtectionStatusGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionStatusGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionStatusGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionStatusGetResponseEnvelopeErrors]
type advancedTCPProtectionStatusGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    advancedTCPProtectionStatusGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionStatusGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionStatusGetResponseEnvelopeErrorsSource]
type advancedTCPProtectionStatusGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusGetResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionStatusGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionStatusGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionStatusGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionStatusGetResponseEnvelopeMessages]
type advancedTCPProtectionStatusGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionStatusGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    advancedTCPProtectionStatusGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionStatusGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionStatusGetResponseEnvelopeMessagesSource]
type advancedTCPProtectionStatusGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionStatusGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionStatusGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionStatusGetResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionStatusGetResponseEnvelopeSuccessTrue AdvancedTCPProtectionStatusGetResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
