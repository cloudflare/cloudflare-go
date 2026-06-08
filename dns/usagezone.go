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

// UsageZoneService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageZoneService] method instead.
type UsageZoneService struct {
	Options []option.RequestOption
}

// NewUsageZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUsageZoneService(opts ...option.RequestOption) (r *UsageZoneService) {
	r = &UsageZoneService{}
	r.Options = opts
	return
}

// Get the current DNS record usage for a zone, including the number of records and
// the quota limit.
func (r *UsageZoneService) Get(ctx context.Context, query UsageZoneGetParams, opts ...option.RequestOption) (res *UsageZoneGetResponse, err error) {
	var env UsageZoneGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/dns_records/usage", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type UsageZoneGetResponse struct {
	// Maximum number of DNS records allowed for the zone. Null if using account-level
	// quota.
	RecordQuota int64 `json:"record_quota" api:"required,nullable"`
	// Current number of DNS records in the zone.
	RecordUsage int64                    `json:"record_usage" api:"required"`
	JSON        usageZoneGetResponseJSON `json:"-"`
}

// usageZoneGetResponseJSON contains the JSON metadata for the struct
// [UsageZoneGetResponse]
type usageZoneGetResponseJSON struct {
	RecordQuota apijson.Field
	RecordUsage apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageZoneGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageZoneGetResponseJSON) RawJSON() string {
	return r.raw
}

type UsageZoneGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type UsageZoneGetResponseEnvelope struct {
	Errors   []UsageZoneGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []UsageZoneGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success UsageZoneGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  UsageZoneGetResponse                `json:"result"`
	JSON    usageZoneGetResponseEnvelopeJSON    `json:"-"`
}

// usageZoneGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UsageZoneGetResponseEnvelope]
type usageZoneGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageZoneGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageZoneGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UsageZoneGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           UsageZoneGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             usageZoneGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// usageZoneGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UsageZoneGetResponseEnvelopeErrors]
type usageZoneGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UsageZoneGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageZoneGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UsageZoneGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    usageZoneGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// usageZoneGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [UsageZoneGetResponseEnvelopeErrorsSource]
type usageZoneGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageZoneGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageZoneGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UsageZoneGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           UsageZoneGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             usageZoneGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// usageZoneGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UsageZoneGetResponseEnvelopeMessages]
type usageZoneGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UsageZoneGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageZoneGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UsageZoneGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    usageZoneGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// usageZoneGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [UsageZoneGetResponseEnvelopeMessagesSource]
type usageZoneGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageZoneGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r usageZoneGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UsageZoneGetResponseEnvelopeSuccess bool

const (
	UsageZoneGetResponseEnvelopeSuccessTrue UsageZoneGetResponseEnvelopeSuccess = true
)

func (r UsageZoneGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UsageZoneGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
