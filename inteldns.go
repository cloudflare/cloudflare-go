// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IntelDNSService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewIntelDNSService] method instead.
type IntelDNSService struct {
	Options []option.RequestOption
}

// NewIntelDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelDNSService(opts ...option.RequestOption) (r *IntelDNSService) {
	r = &IntelDNSService{}
	r.Options = opts
	return
}

// Get Passive DNS by IP
func (r *IntelDNSService) PassiveDNSByIPGetPassiveDNSByIP(ctx context.Context, accountID string, query IntelDNSPassiveDNSByIPGetPassiveDNSByIPParams, opts ...option.RequestOption) (res *IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/dns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords []IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseReverseRecord `json:"reverse_records"`
	JSON           intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseJSON            `json:"-"`
}

// intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseJSON contains the JSON metadata
// for the struct [IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse]
type intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseJSON struct {
	Count          apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	ReverseRecords apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseReverseRecord struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname interface{} `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time                                                        `json:"last_seen" format:"date"`
	JSON     intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseReverseRecordJSON `json:"-"`
}

// intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseReverseRecordJSON contains the
// JSON metadata for the struct
// [IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseReverseRecord]
type intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseReverseRecordJSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseReverseRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDNSPassiveDNSByIPGetPassiveDNSByIPParams struct {
	IPV4 param.Field[string] `query:"ipv4"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage        param.Field[float64]                                                     `query:"per_page"`
	StartEndParams param.Field[IntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams] `query:"start_end_params"`
}

// URLQuery serializes [IntelDNSPassiveDNSByIPGetPassiveDNSByIPParams]'s query
// parameters as `url.Values`.
func (r IntelDNSPassiveDNSByIPGetPassiveDNSByIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams struct {
	// Defaults to the current date.
	End param.Field[time.Time] `query:"end" format:"date"`
	// Defaults to 30 days before the end parameter value.
	Start param.Field[time.Time] `query:"start" format:"date"`
}

// URLQuery serializes
// [IntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams]'s query parameters
// as `url.Values`.
func (r IntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelope struct {
	Errors   []IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeJSON    `json:"-"`
}

// intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelope]
type intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeErrors]
type intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeMessages]
type intelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeSuccess bool

const (
	IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeSuccessTrue IntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseEnvelopeSuccess = true
)
