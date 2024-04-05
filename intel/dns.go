// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDNSService] method instead.
type DNSService struct {
	Options []option.RequestOption
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	return
}

// Get Passive DNS by IP
func (r *DNSService) Get(ctx context.Context, params DNSGetParams, opts ...option.RequestOption) (res *DNS, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/dns", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNS struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords []UnnamedSchemaRefB5e16cee4f32382c294201aedb9fc050 `json:"reverse_records"`
	JSON           dnsJSON                                            `json:"-"`
}

// dnsJSON contains the JSON metadata for the struct [DNS]
type dnsJSON struct {
	Count          apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	ReverseRecords apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsJSON) RawJSON() string {
	return r.raw
}

type DNSParam struct {
	// Total results returned based on your search parameters.
	Count param.Field[float64] `json:"count"`
	// Current page within paginated list of results.
	Page param.Field[float64] `json:"page"`
	// Number of results per page of results.
	PerPage param.Field[float64] `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords param.Field[[]UnnamedSchemaRefB5e16cee4f32382c294201aedb9fc050Param] `json:"reverse_records"`
}

func (r DNSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UnnamedSchemaRefB5e16cee4f32382c294201aedb9fc050 struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname interface{} `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time                                            `json:"last_seen" format:"date"`
	JSON     unnamedSchemaRefB5e16cee4f32382c294201aedb9fc050JSON `json:"-"`
}

// unnamedSchemaRefB5e16cee4f32382c294201aedb9fc050JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefB5e16cee4f32382c294201aedb9fc050]
type unnamedSchemaRefB5e16cee4f32382c294201aedb9fc050JSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRefB5e16cee4f32382c294201aedb9fc050) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefB5e16cee4f32382c294201aedb9fc050JSON) RawJSON() string {
	return r.raw
}

type DNSGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	IPV4      param.Field[string] `query:"ipv4"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage        param.Field[float64]                    `query:"per_page"`
	StartEndParams param.Field[DNSGetParamsStartEndParams] `query:"start_end_params"`
}

// URLQuery serializes [DNSGetParams]'s query parameters as `url.Values`.
func (r DNSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSGetParamsStartEndParams struct {
	// Defaults to the current date.
	End param.Field[time.Time] `query:"end" format:"date"`
	// Defaults to 30 days before the end parameter value.
	Start param.Field[time.Time] `query:"start" format:"date"`
}

// URLQuery serializes [DNSGetParamsStartEndParams]'s query parameters as
// `url.Values`.
func (r DNSGetParamsStartEndParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DNS                                                       `json:"result,required"`
	// Whether the API call was successful
	Success DNSGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsGetResponseEnvelopeJSON    `json:"-"`
}

// dnsGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelope]
type dnsGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSGetResponseEnvelopeSuccess bool

const (
	DNSGetResponseEnvelopeSuccessTrue DNSGetResponseEnvelopeSuccess = true
)

func (r DNSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
