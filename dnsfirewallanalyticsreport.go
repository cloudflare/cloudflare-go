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

// DNSFirewallAnalyticsReportService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDNSFirewallAnalyticsReportService] method instead.
type DNSFirewallAnalyticsReportService struct {
	Options []option.RequestOption
	Bytimes *DNSFirewallAnalyticsReportBytimeService
}

// NewDNSFirewallAnalyticsReportService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDNSFirewallAnalyticsReportService(opts ...option.RequestOption) (r *DNSFirewallAnalyticsReportService) {
	r = &DNSFirewallAnalyticsReportService{}
	r.Options = opts
	r.Bytimes = NewDNSFirewallAnalyticsReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSFirewallAnalyticsReportService) Get(ctx context.Context, accountIdentifier string, identifier string, query DNSFirewallAnalyticsReportGetParams, opts ...option.RequestOption) (res *DNSDNSAnalyticsAPIReport, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallAnalyticsReportGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSFirewallAnalyticsReportGetParams struct {
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
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DNSFirewallAnalyticsReportGetParams]'s query parameters as
// `url.Values`.
func (r DNSFirewallAnalyticsReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSFirewallAnalyticsReportGetResponseEnvelope struct {
	Errors   []DNSFirewallAnalyticsReportGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallAnalyticsReportGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSDNSAnalyticsAPIReport                                `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallAnalyticsReportGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallAnalyticsReportGetResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallAnalyticsReportGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSFirewallAnalyticsReportGetResponseEnvelope]
type dnsFirewallAnalyticsReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    dnsFirewallAnalyticsReportGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallAnalyticsReportGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DNSFirewallAnalyticsReportGetResponseEnvelopeErrors]
type dnsFirewallAnalyticsReportGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    dnsFirewallAnalyticsReportGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallAnalyticsReportGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DNSFirewallAnalyticsReportGetResponseEnvelopeMessages]
type dnsFirewallAnalyticsReportGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallAnalyticsReportGetResponseEnvelopeSuccess bool

const (
	DNSFirewallAnalyticsReportGetResponseEnvelopeSuccessTrue DNSFirewallAnalyticsReportGetResponseEnvelopeSuccess = true
)
