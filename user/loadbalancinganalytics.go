// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// LoadBalancingAnalyticsService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancingAnalyticsService]
// method instead.
type LoadBalancingAnalyticsService struct {
	Options []option.RequestOption
	Events  *LoadBalancingAnalyticsEventService
}

// NewLoadBalancingAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancingAnalyticsService(opts ...option.RequestOption) (r *LoadBalancingAnalyticsService) {
	r = &LoadBalancingAnalyticsService{}
	r.Options = opts
	r.Events = NewLoadBalancingAnalyticsEventService(opts...)
	return
}

type Analytics struct {
	ID        int64         `json:"id"`
	Origins   []interface{} `json:"origins"`
	Pool      interface{}   `json:"pool"`
	Timestamp time.Time     `json:"timestamp" format:"date-time"`
	JSON      analyticsJSON `json:"-"`
}

// analyticsJSON contains the JSON metadata for the struct [Analytics]
type analyticsJSON struct {
	ID          apijson.Field
	Origins     apijson.Field
	Pool        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Analytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsJSON) RawJSON() string {
	return r.raw
}
