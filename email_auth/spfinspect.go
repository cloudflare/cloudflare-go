// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// SPFInspectService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSPFInspectService] method instead.
type SPFInspectService struct {
	Options []option.RequestOption
}

// NewSPFInspectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSPFInspectService(opts ...option.RequestOption) (r *SPFInspectService) {
	r = &SPFInspectService{}
	r.Options = opts
	return
}

// Inspects a specific SPF TXT record and returns a parsed tree structure in the
// spflimit-worker format.
//
// The record ID must be provided via the `id` query parameter.
//
// Returns a recursive tree showing:
//
// - Parsed components with their qualifiers and types
// - Nested includes recursively resolved within components
// - Per-component and total lookup counts
// - Detailed error information with context
func (r *SPFInspectService) Get(ctx context.Context, params SPFInspectGetParams, opts ...option.RequestOption) (res *SPFInspectGetResponse, err error) {
	var env SPFInspectGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/auth/spf/inspect", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Recursive SPF inspection tree
type SPFInspectGetResponse struct {
	// Parsed SPF components (mechanisms)
	Components []interface{} `json:"components" api:"required"`
	// Domain being inspected
	Domain string `json:"domain" api:"required"`
	// Raw SPF record content
	Record string `json:"record" api:"required"`
	// Total number of DNS lookups performed across all includes
	TotalLookups int64 `json:"total_lookups" api:"required"`
	// All errors encountered during inspection, collected from the entire tree. This
	// includes errors from nested includes at any depth, providing a quick overview of
	// all issues without needing to traverse the nested structure. Each error includes
	// a `domain` field to identify where it occurred. Empty array if no errors
	// (omitted from JSON when empty).
	Errors []SPFInspectGetResponseError `json:"errors"`
	JSON   spfInspectGetResponseJSON    `json:"-"`
}

// spfInspectGetResponseJSON contains the JSON metadata for the struct
// [SPFInspectGetResponse]
type spfInspectGetResponseJSON struct {
	Components   apijson.Field
	Domain       apijson.Field
	Record       apijson.Field
	TotalLookups apijson.Field
	Errors       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SPFInspectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spfInspectGetResponseJSON) RawJSON() string {
	return r.raw
}

// An error encountered during SPF inspection
type SPFInspectGetResponseError struct {
	// Error code. Known values:
	//
	//   - `lookup_failed` — DNS TXT lookup failed
	//   - `spf_not_found` — no SPF record found
	//   - `invalid_spf` — record does not start with `v=spf1`
	//   - `invalid_domain` — PSL validation failed
	//   - `loop_detected` — include/redirect cycle detected
	//   - `invalid_mechanism` — unrecognised or malformed mechanism
	//   - `resource_limit_exceeded` — internal resource protection limits exceeded
	//     (recursion depth or query budget)
	//   - `max_lookups` — RFC 7208 10-lookup limit exceeded
	Code string `json:"code" api:"required"`
	// Domain where the error occurred
	Domain string `json:"domain" api:"required"`
	// Human-readable error message
	Message string `json:"message" api:"required"`
	// Additional error-specific details (optional).
	//
	//   - For `invalid_domain` errors: the invalid domain string
	//   - For `invalid_mechanism` errors: the invalid mechanism text (e.g.,
	//     "invalidmech123")
	//   - For `loop_detected` errors: the domain that caused the loop
	//   - For other error types: not present
	Details string                         `json:"details"`
	JSON    spfInspectGetResponseErrorJSON `json:"-"`
}

// spfInspectGetResponseErrorJSON contains the JSON metadata for the struct
// [SPFInspectGetResponseError]
type spfInspectGetResponseErrorJSON struct {
	Code        apijson.Field
	Domain      apijson.Field
	Message     apijson.Field
	Details     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SPFInspectGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spfInspectGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type SPFInspectGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// DNS record ID (rec_tag) to inspect
	ID param.Field[string] `query:"id" api:"required"`
}

// URLQuery serializes [SPFInspectGetParams]'s query parameters as `url.Values`.
func (r SPFInspectGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SPFInspectGetResponseEnvelope struct {
	Errors   []SPFInspectGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SPFInspectGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SPFInspectGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Recursive SPF inspection tree
	Result SPFInspectGetResponse             `json:"result"`
	JSON   spfInspectGetResponseEnvelopeJSON `json:"-"`
}

// spfInspectGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SPFInspectGetResponseEnvelope]
type spfInspectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SPFInspectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spfInspectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SPFInspectGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           SPFInspectGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             spfInspectGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// spfInspectGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SPFInspectGetResponseEnvelopeErrors]
type spfInspectGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SPFInspectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spfInspectGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SPFInspectGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    spfInspectGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// spfInspectGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SPFInspectGetResponseEnvelopeErrorsSource]
type spfInspectGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SPFInspectGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spfInspectGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SPFInspectGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           SPFInspectGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             spfInspectGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// spfInspectGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SPFInspectGetResponseEnvelopeMessages]
type spfInspectGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SPFInspectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spfInspectGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SPFInspectGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    spfInspectGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// spfInspectGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SPFInspectGetResponseEnvelopeMessagesSource]
type spfInspectGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SPFInspectGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spfInspectGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SPFInspectGetResponseEnvelopeSuccess bool

const (
	SPFInspectGetResponseEnvelopeSuccessTrue SPFInspectGetResponseEnvelopeSuccess = true
)

func (r SPFInspectGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SPFInspectGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
