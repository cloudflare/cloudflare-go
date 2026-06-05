// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// UsageAccountService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageAccountService] method instead.
type UsageAccountService struct {
	Options []option.RequestOption
}

// NewUsageAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUsageAccountService(opts ...option.RequestOption) (r *UsageAccountService) {
	r = &UsageAccountService{}
	r.Options = opts
	return
}

// Get the current DNS record usage and quota for an account. May include internal
// DNS usage and quota.
func (r *UsageAccountService) Get(ctx context.Context, query UsageAccountGetParams, opts ...option.RequestOption) (res *UsageAccountGetResponse, err error) {
	var env UsageAccountGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dns_records/usage", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type UsageAccountGetResponse struct {
	// Maximum number of DNS records allowed across all public zones in the account.
	// Null if using zone-level quota.
	RecordQuota int64 `json:"record_quota" api:"required,nullable"`
	// Current number of DNS records across all public zones in the account.
	RecordUsage int64 `json:"record_usage" api:"required"`
	// Maximum number of DNS records allowed across all internal zones in the account.
	// Only present if internal DNS is enabled.
	InternalRecordQuota int64 `json:"internal_record_quota"`
	// Current number of DNS records across all internal zones in the account. Only
	// present if internal DNS is enabled.
	InternalRecordUsage int64                       `json:"internal_record_usage"`
	JSON                usageAccountGetResponseJSON `json:"-"`
}

// usageAccountGetResponseJSON contains the JSON metadata for the struct
// [UsageAccountGetResponse]
type usageAccountGetResponseJSON struct {
	RecordQuota         apijson.Field
	RecordUsage         apijson.Field
	InternalRecordQuota apijson.Field
	InternalRecordUsage apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *UsageAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageAccountGetResponseJSON) RawJSON() string {
	return r.raw
}

type UsageAccountGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type UsageAccountGetResponseEnvelope struct {
	Errors   []UsageAccountGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []UsageAccountGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success UsageAccountGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  UsageAccountGetResponse                `json:"result"`
	JSON    usageAccountGetResponseEnvelopeJSON    `json:"-"`
}

// usageAccountGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UsageAccountGetResponseEnvelope]
type usageAccountGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageAccountGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageAccountGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UsageAccountGetResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           UsageAccountGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             usageAccountGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// usageAccountGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UsageAccountGetResponseEnvelopeErrors]
type usageAccountGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UsageAccountGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageAccountGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UsageAccountGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    usageAccountGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// usageAccountGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [UsageAccountGetResponseEnvelopeErrorsSource]
type usageAccountGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageAccountGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageAccountGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UsageAccountGetResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           UsageAccountGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             usageAccountGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// usageAccountGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UsageAccountGetResponseEnvelopeMessages]
type usageAccountGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UsageAccountGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageAccountGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UsageAccountGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    usageAccountGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// usageAccountGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [UsageAccountGetResponseEnvelopeMessagesSource]
type usageAccountGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageAccountGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageAccountGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UsageAccountGetResponseEnvelopeSuccess bool

const (
	UsageAccountGetResponseEnvelopeSuccessTrue UsageAccountGetResponseEnvelopeSuccess = true
)

func (r UsageAccountGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsageAccountGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
