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

// AccountIntelDNSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelDNSService] method
// instead.
type AccountIntelDNSService struct {
	Options []option.RequestOption
}

// NewAccountIntelDNSService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountIntelDNSService(opts ...option.RequestOption) (r *AccountIntelDNSService) {
	r = &AccountIntelDNSService{}
	r.Options = opts
	return
}

// Get Passive DNS by IP
func (r *AccountIntelDNSService) PassiveDNSByIPGetPassiveDNSByIP(ctx context.Context, accountIdentifier string, query AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParams, opts ...option.RequestOption) (res *AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/dns", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse struct {
	Errors   []AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseError   `json:"errors"`
	Messages []AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseMessage `json:"messages"`
	Result   AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseSuccess `json:"success"`
	JSON    accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseJSON    `json:"-"`
}

// accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseJSON contains the JSON
// metadata for the struct [AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse]
type accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseErrorJSON `json:"-"`
}

// accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseError]
type accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseMessageJSON `json:"-"`
}

// accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseMessage]
type accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResult struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords []AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultReverseRecord `json:"reverse_records"`
	JSON           accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultJSON            `json:"-"`
}

// accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultJSON contains the
// JSON metadata for the struct
// [AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResult]
type accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultJSON struct {
	Count          apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	ReverseRecords apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultReverseRecord struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname interface{} `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time                                                                     `json:"last_seen" format:"date"`
	JSON     accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultReverseRecordJSON `json:"-"`
}

// accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultReverseRecordJSON
// contains the JSON metadata for the struct
// [AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultReverseRecord]
type accountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultReverseRecordJSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseResultReverseRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseSuccess bool

const (
	AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseSuccessTrue AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPResponseSuccess = true
)

type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParams struct {
	Ipv4 param.Field[string] `query:"ipv4"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage        param.Field[float64]                                                            `query:"per_page"`
	StartEndParams param.Field[AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams] `query:"start_end_params"`
}

// URLQuery serializes [AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParams]'s
// query parameters as `url.Values`.
func (r AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams struct {
	// Defaults to the current date.
	End param.Field[time.Time] `query:"end" format:"date"`
	// Defaults to 30 days before the end parameter value.
	Start param.Field[time.Time] `query:"start" format:"date"`
}

// URLQuery serializes
// [AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams]'s query
// parameters as `url.Values`.
func (r AccountIntelDNSPassiveDNSByIPGetPassiveDNSByIPParamsStartEndParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
