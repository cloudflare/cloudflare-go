// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// DeviceOverrideCodeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceOverrideCodeService] method instead.
type DeviceOverrideCodeService struct {
	Options []option.RequestOption
}

// NewDeviceOverrideCodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceOverrideCodeService(opts ...option.RequestOption) (r *DeviceOverrideCodeService) {
	r = &DeviceOverrideCodeService{}
	r.Options = opts
	return
}

// Fetches a one-time use admin override code for a registration. This relies on
// the **Admin Override** setting being enabled in your device configuration.
//
// **Deprecated:** please use GET
// /accounts/{account_id}/devices/registrations/{registration_id}/override_codes
// instead.
func (r *DeviceOverrideCodeService) List(ctx context.Context, deviceID string, query DeviceOverrideCodeListParams, opts ...option.RequestOption) (res *DeviceOverrideCodeListResponse, err error) {
	var env DeviceOverrideCodeListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/%s/override_codes", query.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches one-time use admin override codes for a registration. This relies on the
// **Admin Override** setting being enabled in your device configuration.
func (r *DeviceOverrideCodeService) Get(ctx context.Context, registrationID string, query DeviceOverrideCodeGetParams, opts ...option.RequestOption) (res *DeviceOverrideCodeGetResponse, err error) {
	var env DeviceOverrideCodeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if registrationID == "" {
		err = errors.New("missing required registration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations/%s/override_codes", query.AccountID, registrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceOverrideCodeListResponse struct {
	DisableForTime DeviceOverrideCodeListResponseDisableForTime `json:"disable_for_time"`
	JSON           deviceOverrideCodeListResponseJSON           `json:"-"`
}

// deviceOverrideCodeListResponseJSON contains the JSON metadata for the struct
// [DeviceOverrideCodeListResponse]
type deviceOverrideCodeListResponseJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceOverrideCodeListResponseDisableForTime struct {
	// Override code that is valid for 1 hour.
	One string `json:"1"`
	// Override code that is valid for 12 hour2.
	Twelve string `json:"12"`
	// Override code that is valid for 24 hour.2.
	TwentyFour string `json:"24"`
	// Override code that is valid for 3 hours.
	Three string `json:"3"`
	// Override code that is valid for 6 hours.
	Six  string                                           `json:"6"`
	JSON deviceOverrideCodeListResponseDisableForTimeJSON `json:"-"`
}

// deviceOverrideCodeListResponseDisableForTimeJSON contains the JSON metadata for
// the struct [DeviceOverrideCodeListResponseDisableForTime]
type deviceOverrideCodeListResponseDisableForTimeJSON struct {
	One         apijson.Field
	Twelve      apijson.Field
	TwentyFour  apijson.Field
	Three       apijson.Field
	Six         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponseDisableForTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseDisableForTimeJSON) RawJSON() string {
	return r.raw
}

type DeviceOverrideCodeGetResponse struct {
	DisableForTime map[string]string                 `json:"disable_for_time"`
	JSON           deviceOverrideCodeGetResponseJSON `json:"-"`
}

// deviceOverrideCodeGetResponseJSON contains the JSON metadata for the struct
// [DeviceOverrideCodeGetResponse]
type deviceOverrideCodeGetResponseJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceOverrideCodeListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceOverrideCodeListResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   DeviceOverrideCodeListResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceOverrideCodeListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceOverrideCodeListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceOverrideCodeListResponseEnvelopeJSON       `json:"-"`
}

// deviceOverrideCodeListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceOverrideCodeListResponseEnvelope]
type deviceOverrideCodeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceOverrideCodeListResponseEnvelopeSuccess bool

const (
	DeviceOverrideCodeListResponseEnvelopeSuccessTrue DeviceOverrideCodeListResponseEnvelopeSuccess = true
)

func (r DeviceOverrideCodeListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceOverrideCodeListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceOverrideCodeListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       deviceOverrideCodeListResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceOverrideCodeListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DeviceOverrideCodeListResponseEnvelopeResultInfo]
type deviceOverrideCodeListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DeviceOverrideCodeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceOverrideCodeGetResponseEnvelope struct {
	Errors   []DeviceOverrideCodeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceOverrideCodeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceOverrideCodeGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success,required"`
	JSON    deviceOverrideCodeGetResponseEnvelopeJSON `json:"-"`
}

// deviceOverrideCodeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceOverrideCodeGetResponseEnvelope]
type deviceOverrideCodeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceOverrideCodeGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceOverrideCodeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceOverrideCodeGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeviceOverrideCodeGetResponseEnvelopeErrors]
type deviceOverrideCodeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceOverrideCodeGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    deviceOverrideCodeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceOverrideCodeGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceOverrideCodeGetResponseEnvelopeMessages]
type deviceOverrideCodeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
