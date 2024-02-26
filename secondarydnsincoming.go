// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SecondaryDNSIncomingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSIncomingService]
// method instead.
type SecondaryDNSIncomingService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSIncomingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSIncomingService(opts ...option.RequestOption) (r *SecondaryDNSIncomingService) {
	r = &SecondaryDNSIncomingService{}
	r.Options = opts
	return
}

// Create secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) New(ctx context.Context, params SecondaryDNSIncomingNewParams, opts ...option.RequestOption) (res *SecondaryDNSIncomingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingNewResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) Update(ctx context.Context, params SecondaryDNSIncomingUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSIncomingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) Delete(ctx context.Context, body SecondaryDNSIncomingDeleteParams, opts ...option.RequestOption) (res *SecondaryDNSIncomingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) Get(ctx context.Context, query SecondaryDNSIncomingGetParams, opts ...option.RequestOption) (res *SecondaryDNSIncomingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingGetResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSIncomingNewResponse struct {
	ID interface{} `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                             `json:"soa_serial"`
	JSON      secondaryDNSIncomingNewResponseJSON `json:"-"`
}

// secondaryDNSIncomingNewResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSIncomingNewResponse]
type secondaryDNSIncomingNewResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SoaSerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecondaryDNSIncomingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingUpdateResponse struct {
	ID interface{} `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                `json:"soa_serial"`
	JSON      secondaryDNSIncomingUpdateResponseJSON `json:"-"`
}

// secondaryDNSIncomingUpdateResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSIncomingUpdateResponse]
type secondaryDNSIncomingUpdateResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SoaSerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecondaryDNSIncomingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingDeleteResponse struct {
	ID   interface{}                            `json:"id"`
	JSON secondaryDNSIncomingDeleteResponseJSON `json:"-"`
}

// secondaryDNSIncomingDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSIncomingDeleteResponse]
type secondaryDNSIncomingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingGetResponse struct {
	ID interface{} `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                             `json:"soa_serial"`
	JSON      secondaryDNSIncomingGetResponseJSON `json:"-"`
}

// secondaryDNSIncomingGetResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSIncomingGetResponse]
type secondaryDNSIncomingGetResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SoaSerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecondaryDNSIncomingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingNewParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSIncomingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSIncomingNewResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingNewResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSIncomingNewResponseEnvelope]
type secondaryDNSIncomingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    secondaryDNSIncomingNewResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSIncomingNewResponseEnvelopeErrors]
type secondaryDNSIncomingNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    secondaryDNSIncomingNewResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSIncomingNewResponseEnvelopeMessages]
type secondaryDNSIncomingNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingNewResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingNewResponseEnvelopeSuccessTrue SecondaryDNSIncomingNewResponseEnvelopeSuccess = true
)

type SecondaryDNSIncomingUpdateParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSIncomingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSIncomingUpdateResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingUpdateResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSIncomingUpdateResponseEnvelope]
type secondaryDNSIncomingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSIncomingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSIncomingUpdateResponseEnvelopeErrors]
type secondaryDNSIncomingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    secondaryDNSIncomingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSIncomingUpdateResponseEnvelopeMessages]
type secondaryDNSIncomingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingUpdateResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingUpdateResponseEnvelopeSuccessTrue SecondaryDNSIncomingUpdateResponseEnvelopeSuccess = true
)

type SecondaryDNSIncomingDeleteParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSIncomingDeleteResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSIncomingDeleteResponseEnvelope]
type secondaryDNSIncomingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSIncomingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSIncomingDeleteResponseEnvelopeErrors]
type secondaryDNSIncomingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    secondaryDNSIncomingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSIncomingDeleteResponseEnvelopeMessages]
type secondaryDNSIncomingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingDeleteResponseEnvelopeSuccessTrue SecondaryDNSIncomingDeleteResponseEnvelopeSuccess = true
)

type SecondaryDNSIncomingGetParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSIncomingGetResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingGetResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSIncomingGetResponseEnvelope]
type secondaryDNSIncomingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    secondaryDNSIncomingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSIncomingGetResponseEnvelopeErrors]
type secondaryDNSIncomingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    secondaryDNSIncomingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSIncomingGetResponseEnvelopeMessages]
type secondaryDNSIncomingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingGetResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingGetResponseEnvelopeSuccessTrue SecondaryDNSIncomingGetResponseEnvelopeSuccess = true
)
