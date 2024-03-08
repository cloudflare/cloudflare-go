// File generated from our OpenAPI spec by Stainless.

package cloudflare

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
func (r *IntelDNSService) Get(ctx context.Context, params IntelDNSGetParams, opts ...option.RequestOption) (res *IntelDNSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDNSGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/dns", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDNSGetResponse struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords []IntelDNSGetResponseReverseRecord `json:"reverse_records"`
	JSON           intelDNSGetResponseJSON            `json:"-"`
}

// intelDNSGetResponseJSON contains the JSON metadata for the struct
// [IntelDNSGetResponse]
type intelDNSGetResponseJSON struct {
	Count          apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	ReverseRecords apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IntelDNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDNSGetResponseJSON) RawJSON() string {
	return r.raw
}

type IntelDNSGetResponseReverseRecord struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname interface{} `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time                            `json:"last_seen" format:"date"`
	JSON     intelDNSGetResponseReverseRecordJSON `json:"-"`
}

// intelDNSGetResponseReverseRecordJSON contains the JSON metadata for the struct
// [IntelDNSGetResponseReverseRecord]
type intelDNSGetResponseReverseRecordJSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSGetResponseReverseRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDNSGetResponseReverseRecordJSON) RawJSON() string {
	return r.raw
}

type IntelDNSGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	IPV4      param.Field[string] `query:"ipv4"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage        param.Field[float64]                         `query:"per_page"`
	StartEndParams param.Field[IntelDNSGetParamsStartEndParams] `query:"start_end_params"`
}

// URLQuery serializes [IntelDNSGetParams]'s query parameters as `url.Values`.
func (r IntelDNSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDNSGetParamsStartEndParams struct {
	// Defaults to the current date.
	End param.Field[time.Time] `query:"end" format:"date"`
	// Defaults to 30 days before the end parameter value.
	Start param.Field[time.Time] `query:"start" format:"date"`
}

// URLQuery serializes [IntelDNSGetParamsStartEndParams]'s query parameters as
// `url.Values`.
func (r IntelDNSGetParamsStartEndParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDNSGetResponseEnvelope struct {
	Errors   []IntelDNSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDNSGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelDNSGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelDNSGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelDNSGetResponseEnvelopeJSON    `json:"-"`
}

// intelDNSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelDNSGetResponseEnvelope]
type intelDNSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDNSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelDNSGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    intelDNSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDNSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IntelDNSGetResponseEnvelopeErrors]
type intelDNSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDNSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelDNSGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    intelDNSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDNSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IntelDNSGetResponseEnvelopeMessages]
type intelDNSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDNSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelDNSGetResponseEnvelopeSuccess bool

const (
	IntelDNSGetResponseEnvelopeSuccessTrue IntelDNSGetResponseEnvelopeSuccess = true
)
