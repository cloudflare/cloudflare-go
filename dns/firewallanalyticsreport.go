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

// FirewallAnalyticsReportService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewFirewallAnalyticsReportService] method instead.
type FirewallAnalyticsReportService struct {
	Options []option.RequestOption
	Bytimes *FirewallAnalyticsReportBytimeService
}

// NewFirewallAnalyticsReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallAnalyticsReportService(opts ...option.RequestOption) (r *FirewallAnalyticsReportService) {
	r = &FirewallAnalyticsReportService{}
	r.Options = opts
	r.Bytimes = NewFirewallAnalyticsReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *FirewallAnalyticsReportService) Get(ctx context.Context, dnsFirewallID string, params FirewallAnalyticsReportGetParams, opts ...option.RequestOption) (res *Report, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAnalyticsReportGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report", params.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallAnalyticsReportGetParams struct {
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
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [FirewallAnalyticsReportGetParams]'s query parameters as
// `url.Values`.
func (r FirewallAnalyticsReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallAnalyticsReportGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FirewallAnalyticsReportGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Report                                            `json:"result"`
	JSON    firewallAnalyticsReportGetResponseEnvelopeJSON    `json:"-"`
}

// firewallAnalyticsReportGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallAnalyticsReportGetResponseEnvelope]
type firewallAnalyticsReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAnalyticsReportGetResponseEnvelopeSuccess bool

const (
	FirewallAnalyticsReportGetResponseEnvelopeSuccessTrue FirewallAnalyticsReportGetResponseEnvelopeSuccess = true
)

func (r FirewallAnalyticsReportGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallAnalyticsReportGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
