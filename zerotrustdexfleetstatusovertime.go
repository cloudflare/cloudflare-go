// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDEXFleetStatusOverTimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDEXFleetStatusOverTimeService] method instead.
type ZeroTrustDEXFleetStatusOverTimeService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXFleetStatusOverTimeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDEXFleetStatusOverTimeService(opts ...option.RequestOption) (r *ZeroTrustDEXFleetStatusOverTimeService) {
	r = &ZeroTrustDEXFleetStatusOverTimeService{}
	r.Options = opts
	return
}

// List details for devices using WARP, up to 7 days
func (r *ZeroTrustDEXFleetStatusOverTimeService) List(ctx context.Context, params ZeroTrustDEXFleetStatusOverTimeListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/over-time", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type ZeroTrustDEXFleetStatusOverTimeListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Timestamp in ISO format
	TimeEnd param.Field[string] `query:"time_end,required"`
	// Timestamp in ISO format
	TimeStart param.Field[string] `query:"time_start,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
}

// URLQuery serializes [ZeroTrustDEXFleetStatusOverTimeListParams]'s query
// parameters as `url.Values`.
func (r ZeroTrustDEXFleetStatusOverTimeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
