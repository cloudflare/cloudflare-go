// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// DEXTestUniqueDeviceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXTestUniqueDeviceService] method instead.
type DEXTestUniqueDeviceService struct {
	Options []option.RequestOption
}

// NewDEXTestUniqueDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXTestUniqueDeviceService(opts ...option.RequestOption) (r *DEXTestUniqueDeviceService) {
	r = &DEXTestUniqueDeviceService{}
	r.Options = opts
	return
}

// Returns unique count of devices that have run synthetic application monitoring
// tests in the past 7 days.
func (r *DEXTestUniqueDeviceService) List(ctx context.Context, params DEXTestUniqueDeviceListParams, opts ...option.RequestOption) (res *UniqueDevices, err error) {
	var env DEXTestUniqueDeviceListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/tests/unique-devices", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UniqueDevices struct {
	// total number of unique devices
	UniqueDevicesTotal int64             `json:"uniqueDevicesTotal,required"`
	JSON               uniqueDevicesJSON `json:"-"`
}

// uniqueDevicesJSON contains the JSON metadata for the struct [UniqueDevices]
type uniqueDevicesJSON struct {
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UniqueDevices) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uniqueDevicesJSON) RawJSON() string {
	return r.raw
}

type DEXTestUniqueDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [DEXTestUniqueDeviceListParams]'s query parameters as
// `url.Values`.
func (r DEXTestUniqueDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DEXTestUniqueDeviceListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DEXTestUniqueDeviceListResponseEnvelopeSuccess `json:"success,required"`
	Result  UniqueDevices                                  `json:"result"`
	JSON    dexTestUniqueDeviceListResponseEnvelopeJSON    `json:"-"`
}

// dexTestUniqueDeviceListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DEXTestUniqueDeviceListResponseEnvelope]
type dexTestUniqueDeviceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTestUniqueDeviceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTestUniqueDeviceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DEXTestUniqueDeviceListResponseEnvelopeSuccess bool

const (
	DEXTestUniqueDeviceListResponseEnvelopeSuccessTrue DEXTestUniqueDeviceListResponseEnvelopeSuccess = true
)

func (r DEXTestUniqueDeviceListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXTestUniqueDeviceListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
