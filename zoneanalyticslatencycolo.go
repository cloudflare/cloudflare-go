// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAnalyticsLatencyColoService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneAnalyticsLatencyColoService] method instead.
type ZoneAnalyticsLatencyColoService struct {
	Options []option.RequestOption
}

// NewZoneAnalyticsLatencyColoService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAnalyticsLatencyColoService(opts ...option.RequestOption) (r *ZoneAnalyticsLatencyColoService) {
	r = &ZoneAnalyticsLatencyColoService{}
	r.Options = opts
	return
}

// Argo Analytics for a zone at different PoPs
func (r *ZoneAnalyticsLatencyColoService) ArgoAnalyticsForGeolocationArgoAnalyticsForAZoneAtDifferentPoPs(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/latency/colos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
