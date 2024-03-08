// File generated from our OpenAPI spec by Stainless.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *DNSService) Get(ctx context.Context, params DNSGetParams, opts ...option.RequestOption) (res *IntelPassiveDNSByIP, err error) {
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

type IntelPassiveDNSByIP struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords []IntelPassiveDNSByIPReverseRecord `json:"reverse_records"`
	JSON           intelPassiveDNSByIPJSON            `json:"-"`
}

// intelPassiveDNSByIPJSON contains the JSON metadata for the struct
// [IntelPassiveDNSByIP]
type intelPassiveDNSByIPJSON struct {
	Count          apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	ReverseRecords apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IntelPassiveDNSByIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPassiveDNSByIPJSON) RawJSON() string {
	return r.raw
}

type IntelPassiveDNSByIPReverseRecord struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname interface{} `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time                            `json:"last_seen" format:"date"`
	JSON     intelPassiveDNSByIPReverseRecordJSON `json:"-"`
}

// intelPassiveDNSByIPReverseRecordJSON contains the JSON metadata for the struct
// [IntelPassiveDNSByIPReverseRecord]
type intelPassiveDNSByIPReverseRecordJSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPassiveDNSByIPReverseRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPassiveDNSByIPReverseRecordJSON) RawJSON() string {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSGetResponseEnvelope struct {
	Errors   []DNSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelPassiveDNSByIP              `json:"result,required"`
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

type DNSGetResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dnsGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeErrors]
type dnsGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSGetResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    dnsGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DNSGetResponseEnvelopeMessages]
type dnsGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSGetResponseEnvelopeSuccess bool

const (
	DNSGetResponseEnvelopeSuccessTrue DNSGetResponseEnvelopeSuccess = true
)
