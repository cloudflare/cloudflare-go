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

// DNSFirewallAnalyticsReportBytimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDNSFirewallAnalyticsReportBytimeService] method instead.
type DNSFirewallAnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewDNSFirewallAnalyticsReportBytimeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDNSFirewallAnalyticsReportBytimeService(opts ...option.RequestOption) (r *DNSFirewallAnalyticsReportBytimeService) {
	r = &DNSFirewallAnalyticsReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSFirewallAnalyticsReportBytimeService) Get(ctx context.Context, accountIdentifier string, identifier string, query DNSFirewallAnalyticsReportBytimeGetParams, opts ...option.RequestOption) (res *DNSDNSAnalyticsAPIReportBytime, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallAnalyticsReportBytimeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSFirewallAnalyticsReportBytimeGetParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// Unit of time to group data by.
	TimeDelta param.Field[DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DNSFirewallAnalyticsReportBytimeGetParams]'s query
// parameters as `url.Values`.
func (r DNSFirewallAnalyticsReportBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta string

const (
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaAll        DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "all"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaAuto       DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "auto"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaYear       DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "year"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaQuarter    DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "quarter"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaMonth      DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "month"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaWeek       DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "week"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaDay        DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "day"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaHour       DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "hour"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaDekaminute DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "dekaminute"
	DNSFirewallAnalyticsReportBytimeGetParamsTimeDeltaMinute     DNSFirewallAnalyticsReportBytimeGetParamsTimeDelta = "minute"
)

type DNSFirewallAnalyticsReportBytimeGetResponseEnvelope struct {
	Errors   []DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSDNSAnalyticsAPIReportBytime                                `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [DNSFirewallAnalyticsReportBytimeGetResponseEnvelope]
type dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeErrors]
type dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeMessages]
type dnsFirewallAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess bool

const (
	DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeSuccessTrue DNSFirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess = true
)
