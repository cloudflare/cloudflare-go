// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AnalyticsLatencyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsLatencyService] method
// instead.
type AnalyticsLatencyService struct {
	Options []option.RequestOption
	Colos   *AnalyticsLatencyColoService
}

// NewAnalyticsLatencyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsLatencyService(opts ...option.RequestOption) (r *AnalyticsLatencyService) {
	r = &AnalyticsLatencyService{}
	r.Options = opts
	r.Colos = NewAnalyticsLatencyColoService(opts...)
	return
}

// Argo Analytics for a zone
func (r *AnalyticsLatencyService) ArgoAnalyticsForZoneArgoAnalyticsForAZone(ctx context.Context, zoneID string, query AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams, opts ...option.RequestOption) (res *AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelope
	path := fmt.Sprintf("zones/%s/analytics/latency", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseUnknown] or
// [shared.UnionString].
type AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse interface {
	ImplementsAnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams struct {
	Bins param.Field[string] `query:"bins"`
}

// URLQuery serializes
// [AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams]'s query
// parameters as `url.Values`.
func (r AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelope struct {
	Errors   []AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeMessages `json:"messages,required"`
	Result   AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeJSON    `json:"-"`
}

// analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelope]
type analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeErrors struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeErrors]
type analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeMessages struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeMessages]
type analyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeSuccess bool

const (
	AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeSuccessTrue AnalyticsLatencyArgoAnalyticsForZoneArgoAnalyticsForAZoneResponseEnvelopeSuccess = true
)
