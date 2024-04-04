// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// FirewallAnalyticsReportBytimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewFirewallAnalyticsReportBytimeService] method instead.
type FirewallAnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewFirewallAnalyticsReportBytimeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFirewallAnalyticsReportBytimeService(opts ...option.RequestOption) (r *FirewallAnalyticsReportBytimeService) {
	r = &FirewallAnalyticsReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *FirewallAnalyticsReportBytimeService) Get(ctx context.Context, dnsFirewallID string, params FirewallAnalyticsReportBytimeGetParams, opts ...option.RequestOption) (res *DNSAnalyticsReportByTime, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAnalyticsReportBytimeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", params.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallAnalyticsReportBytimeGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	TimeDelta param.Field[FirewallAnalyticsReportBytimeGetParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [FirewallAnalyticsReportBytimeGetParams]'s query parameters
// as `url.Values`.
func (r FirewallAnalyticsReportBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type FirewallAnalyticsReportBytimeGetParamsTimeDelta string

const (
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaAll        FirewallAnalyticsReportBytimeGetParamsTimeDelta = "all"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaAuto       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "auto"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaYear       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "year"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaQuarter    FirewallAnalyticsReportBytimeGetParamsTimeDelta = "quarter"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaMonth      FirewallAnalyticsReportBytimeGetParamsTimeDelta = "month"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaWeek       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "week"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaDay        FirewallAnalyticsReportBytimeGetParamsTimeDelta = "day"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaHour       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "hour"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaDekaminute FirewallAnalyticsReportBytimeGetParamsTimeDelta = "dekaminute"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaMinute     FirewallAnalyticsReportBytimeGetParamsTimeDelta = "minute"
)

func (r FirewallAnalyticsReportBytimeGetParamsTimeDelta) IsKnown() bool {
	switch r {
	case FirewallAnalyticsReportBytimeGetParamsTimeDeltaAll, FirewallAnalyticsReportBytimeGetParamsTimeDeltaAuto, FirewallAnalyticsReportBytimeGetParamsTimeDeltaYear, FirewallAnalyticsReportBytimeGetParamsTimeDeltaQuarter, FirewallAnalyticsReportBytimeGetParamsTimeDeltaMonth, FirewallAnalyticsReportBytimeGetParamsTimeDeltaWeek, FirewallAnalyticsReportBytimeGetParamsTimeDeltaDay, FirewallAnalyticsReportBytimeGetParamsTimeDeltaHour, FirewallAnalyticsReportBytimeGetParamsTimeDeltaDekaminute, FirewallAnalyticsReportBytimeGetParamsTimeDeltaMinute:
		return true
	}
	return false
}

type FirewallAnalyticsReportBytimeGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DNSAnalyticsReportByTime                                  `json:"result,required"`
	// Whether the API call was successful
	Success FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAnalyticsReportBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// firewallAnalyticsReportBytimeGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [FirewallAnalyticsReportBytimeGetResponseEnvelope]
type firewallAnalyticsReportBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportBytimeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess bool

const (
	FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccessTrue FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess = true
)

func (r FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
