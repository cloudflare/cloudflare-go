// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AnalyticsLatencyColoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsLatencyColoService]
// method instead.
type AnalyticsLatencyColoService struct {
	Options []option.RequestOption
}

// NewAnalyticsLatencyColoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsLatencyColoService(opts ...option.RequestOption) (r *AnalyticsLatencyColoService) {
	r = &AnalyticsLatencyColoService{}
	r.Options = opts
	return
}

// Argo Analytics for a zone at different PoPs
func (r *AnalyticsLatencyColoService) ArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPs(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelope
	path := fmt.Sprintf("zones/%s/analytics/latency/colos", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseUnknown]
// or [shared.UnionString].
type AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse interface {
	ImplementsAnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelope struct {
	Errors   []AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeMessages `json:"messages,required"`
	Result   AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeJSON    `json:"-"`
}

// analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelope]
type analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeErrors struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeErrors]
type analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeMessages struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeMessages]
type analyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeSuccess bool

const (
	AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeSuccessTrue AnalyticsLatencyColoArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPsResponseEnvelopeSuccess = true
)
