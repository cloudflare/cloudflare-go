// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

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

// QuotaService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQuotaService] method instead.
type QuotaService struct {
	Options []option.RequestOption
}

// NewQuotaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQuotaService(opts ...option.RequestOption) (r *QuotaService) {
	r = &QuotaService{}
	r.Options = opts
	return
}

// Returns custom hostname quota usage for a zone. The allocated quota is a soft
// limit; creating custom hostnames after usage exceeds this limit can still
// succeed until the hard cap is reached. Use the exceeded and hard_cap fields to
// track when usage is above the soft limit and when new custom hostname creation
// will be rejected.
func (r *QuotaService) Get(ctx context.Context, query QuotaGetParams, opts ...option.RequestOption) (res *QuotaGetResponse, err error) {
	var env QuotaGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/quota", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type QuotaGetResponse struct {
	// The allocated custom hostname quota.
	Allocated int64 `json:"allocated" api:"required"`
	// Whether the current usage has exceeded the allocated quota.
	Exceeded bool `json:"exceeded" api:"required"`
	// The maximum number of custom hostnames allowed before create requests are
	// rejected.
	HardCap int64 `json:"hard_cap" api:"required"`
	// The number of custom hostnames currently in use.
	Used int64                `json:"used" api:"required"`
	JSON quotaGetResponseJSON `json:"-"`
}

// quotaGetResponseJSON contains the JSON metadata for the struct
// [QuotaGetResponse]
type quotaGetResponseJSON struct {
	Allocated   apijson.Field
	Exceeded    apijson.Field
	HardCap     apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseJSON) RawJSON() string {
	return r.raw
}

type QuotaGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type QuotaGetResponseEnvelope struct {
	Errors []QuotaGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// Informational messages returned by the custom hostname API.
	Messages []string `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success QuotaGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  QuotaGetResponse                `json:"result"`
	JSON    quotaGetResponseEnvelopeJSON    `json:"-"`
}

// quotaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [QuotaGetResponseEnvelope]
type quotaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type QuotaGetResponseEnvelopeErrors struct {
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           QuotaGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             quotaGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// quotaGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [QuotaGetResponseEnvelopeErrors]
type quotaGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type QuotaGetResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    quotaGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// quotaGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [QuotaGetResponseEnvelopeErrorsSource]
type quotaGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QuotaGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r quotaGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type QuotaGetResponseEnvelopeSuccess bool

const (
	QuotaGetResponseEnvelopeSuccessTrue QuotaGetResponseEnvelopeSuccess = true
)

func (r QuotaGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case QuotaGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
