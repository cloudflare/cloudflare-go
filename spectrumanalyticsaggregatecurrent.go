// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SpectrumAnalyticsAggregateCurrentService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSpectrumAnalyticsAggregateCurrentService] method instead.
type SpectrumAnalyticsAggregateCurrentService struct {
	Options []option.RequestOption
}

// NewSpectrumAnalyticsAggregateCurrentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSpectrumAnalyticsAggregateCurrentService(opts ...option.RequestOption) (r *SpectrumAnalyticsAggregateCurrentService) {
	r = &SpectrumAnalyticsAggregateCurrentService{}
	r.Options = opts
	return
}

// Retrieves analytics aggregated from the last minute of usage on Spectrum
// applications underneath a given zone.
func (r *SpectrumAnalyticsAggregateCurrentService) SpectrumAggregateAnalyticsGetCurrentAggregatedAnalytics(ctx context.Context, zone string, query SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams, opts ...option.RequestOption) (res *[]SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/analytics/aggregate/current", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponse = interface{}

type SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams struct {
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppIDParam param.Field[string] `query:"app_id_param"`
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppID param.Field[string] `query:"appID"`
	// Co-location identifier.
	ColoName param.Field[string] `query:"colo_name"`
}

// URLQuery serializes
// [SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams]'s
// query parameters as `url.Values`.
func (r SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelope struct {
	Errors   []SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeMessages `json:"messages,required"`
	Result   []SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeJSON    `json:"-"`
}

// spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelope]
type spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeErrors struct {
	Code    int64                                                                                                              `json:"code,required"`
	Message string                                                                                                             `json:"message,required"`
	JSON    spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeErrors]
type spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeMessages struct {
	Code    int64                                                                                                                `json:"code,required"`
	Message string                                                                                                               `json:"message,required"`
	JSON    spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeMessages]
type spectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeSuccess bool

const (
	SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeSuccessTrue SpectrumAnalyticsAggregateCurrentSpectrumAggregateAnalyticsGetCurrentAggregatedAnalyticsResponseEnvelopeSuccess = true
)
