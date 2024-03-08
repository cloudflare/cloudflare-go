// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustDeviceOverrideCodeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustDeviceOverrideCodeService] method instead.
type ZeroTrustDeviceOverrideCodeService struct {
	Options []option.RequestOption
}

// NewZeroTrustDeviceOverrideCodeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDeviceOverrideCodeService(opts ...option.RequestOption) (r *ZeroTrustDeviceOverrideCodeService) {
	r = &ZeroTrustDeviceOverrideCodeService{}
	r.Options = opts
	return
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration.
func (r *ZeroTrustDeviceOverrideCodeService) List(ctx context.Context, deviceID string, query ZeroTrustDeviceOverrideCodeListParams, opts ...option.RequestOption) (res *ZeroTrustDeviceOverrideCodeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceOverrideCodeListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/%s/override_codes", query.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDeviceOverrideCodeListResponse struct {
	DisableForTime ZeroTrustDeviceOverrideCodeListResponseDisableForTime `json:"disable_for_time"`
	JSON           zeroTrustDeviceOverrideCodeListResponseJSON           `json:"-"`
}

// zeroTrustDeviceOverrideCodeListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceOverrideCodeListResponse]
type zeroTrustDeviceOverrideCodeListResponseJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustDeviceOverrideCodeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceOverrideCodeListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceOverrideCodeListResponseDisableForTime struct {
	// Override code that is valid for 1 hour.
	Number1 interface{} `json:"1"`
	// Override code that is valid for 12 hour2.
	Number12 interface{} `json:"12"`
	// Override code that is valid for 24 hour.2.
	Number24 interface{} `json:"24"`
	// Override code that is valid for 3 hours.
	Number3 interface{} `json:"3"`
	// Override code that is valid for 6 hours.
	Number6 interface{}                                               `json:"6"`
	JSON    zeroTrustDeviceOverrideCodeListResponseDisableForTimeJSON `json:"-"`
}

// zeroTrustDeviceOverrideCodeListResponseDisableForTimeJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceOverrideCodeListResponseDisableForTime]
type zeroTrustDeviceOverrideCodeListResponseDisableForTimeJSON struct {
	Number1     apijson.Field
	Number12    apijson.Field
	Number24    apijson.Field
	Number3     apijson.Field
	Number6     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceOverrideCodeListResponseDisableForTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceOverrideCodeListResponseDisableForTimeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceOverrideCodeListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceOverrideCodeListResponseEnvelope struct {
	Errors   []ZeroTrustDeviceOverrideCodeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceOverrideCodeListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceOverrideCodeListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDeviceOverrideCodeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDeviceOverrideCodeListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDeviceOverrideCodeListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceOverrideCodeListResponseEnvelope]
type zeroTrustDeviceOverrideCodeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceOverrideCodeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceOverrideCodeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceOverrideCodeListResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustDeviceOverrideCodeListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceOverrideCodeListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceOverrideCodeListResponseEnvelopeErrors]
type zeroTrustDeviceOverrideCodeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceOverrideCodeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceOverrideCodeListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceOverrideCodeListResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustDeviceOverrideCodeListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceOverrideCodeListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustDeviceOverrideCodeListResponseEnvelopeMessages]
type zeroTrustDeviceOverrideCodeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceOverrideCodeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceOverrideCodeListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceOverrideCodeListResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceOverrideCodeListResponseEnvelopeSuccessTrue ZeroTrustDeviceOverrideCodeListResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       zeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfo]
type zeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceOverrideCodeListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
