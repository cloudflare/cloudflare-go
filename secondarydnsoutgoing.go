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

// SecondaryDNSOutgoingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSOutgoingService]
// method instead.
type SecondaryDNSOutgoingService struct {
	Options []option.RequestOption
	Status  *SecondaryDNSOutgoingStatusService
}

// NewSecondaryDNSOutgoingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingService) {
	r = &SecondaryDNSOutgoingService{}
	r.Options = opts
	r.Status = NewSecondaryDNSOutgoingStatusService(opts...)
	return
}

// Create primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) New(ctx context.Context, params SecondaryDNSOutgoingNewParams, opts ...option.RequestOption) (res *SecondaryDNSOutgoingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingNewResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) Update(ctx context.Context, params SecondaryDNSOutgoingUpdateParams, opts ...option.RequestOption) (res *SecondaryDNSOutgoingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) Delete(ctx context.Context, body SecondaryDNSOutgoingDeleteParams, opts ...option.RequestOption) (res *SecondaryDNSOutgoingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable outgoing zone transfers for primary zone and clears IXFR backlog of
// primary zone.
func (r *SecondaryDNSOutgoingService) Disable(ctx context.Context, body SecondaryDNSOutgoingDisableParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingDisableResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/disable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable outgoing zone transfers for primary zone.
func (r *SecondaryDNSOutgoingService) Enable(ctx context.Context, body SecondaryDNSOutgoingEnableParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingEnableResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/enable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
func (r *SecondaryDNSOutgoingService) ForceNotify(ctx context.Context, body SecondaryDNSOutgoingForceNotifyParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingForceNotifyResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/force_notify", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) Get(ctx context.Context, query SecondaryDNSOutgoingGetParams, opts ...option.RequestOption) (res *SecondaryDNSOutgoingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingGetResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSOutgoingNewResponse struct {
	ID interface{} `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                             `json:"soa_serial"`
	JSON      secondaryDNSOutgoingNewResponseJSON `json:"-"`
}

// secondaryDNSOutgoingNewResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSOutgoingNewResponse]
type secondaryDNSOutgoingNewResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SoaSerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingUpdateResponse struct {
	ID interface{} `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                `json:"soa_serial"`
	JSON      secondaryDNSOutgoingUpdateResponseJSON `json:"-"`
}

// secondaryDNSOutgoingUpdateResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSOutgoingUpdateResponse]
type secondaryDNSOutgoingUpdateResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SoaSerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDeleteResponse struct {
	ID   interface{}                            `json:"id"`
	JSON secondaryDNSOutgoingDeleteResponseJSON `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSOutgoingDeleteResponse]
type secondaryDNSOutgoingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingGetResponse struct {
	ID interface{} `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                             `json:"soa_serial"`
	JSON      secondaryDNSOutgoingGetResponseJSON `json:"-"`
}

// secondaryDNSOutgoingGetResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSOutgoingGetResponse]
type secondaryDNSOutgoingGetResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SoaSerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingNewParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSOutgoingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSOutgoingNewResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingNewResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSOutgoingNewResponseEnvelope]
type secondaryDNSOutgoingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    secondaryDNSOutgoingNewResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingNewResponseEnvelopeErrors]
type secondaryDNSOutgoingNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    secondaryDNSOutgoingNewResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingNewResponseEnvelopeMessages]
type secondaryDNSOutgoingNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingNewResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingNewResponseEnvelopeSuccessTrue SecondaryDNSOutgoingNewResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingUpdateParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSOutgoingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSOutgoingUpdateResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingUpdateResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingUpdateResponseEnvelope]
type secondaryDNSOutgoingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSOutgoingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingUpdateResponseEnvelopeErrors]
type secondaryDNSOutgoingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    secondaryDNSOutgoingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingUpdateResponseEnvelopeMessages]
type secondaryDNSOutgoingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingUpdateResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingUpdateResponseEnvelopeSuccessTrue SecondaryDNSOutgoingUpdateResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingDeleteParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSOutgoingDeleteResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingDeleteResponseEnvelope]
type secondaryDNSOutgoingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSOutgoingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingDeleteResponseEnvelopeErrors]
type secondaryDNSOutgoingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    secondaryDNSOutgoingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingDeleteResponseEnvelopeMessages]
type secondaryDNSOutgoingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingDeleteResponseEnvelopeSuccessTrue SecondaryDNSOutgoingDeleteResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingDisableParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSOutgoingDisableResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingDisableResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingDisableResponseEnvelopeMessages `json:"messages,required"`
	// The zone transfer status of a primary zone
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingDisableResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingDisableResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingDisableResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingDisableResponseEnvelope]
type secondaryDNSOutgoingDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDisableResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    secondaryDNSOutgoingDisableResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingDisableResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingDisableResponseEnvelopeErrors]
type secondaryDNSOutgoingDisableResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDisableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDisableResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    secondaryDNSOutgoingDisableResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingDisableResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingDisableResponseEnvelopeMessages]
type secondaryDNSOutgoingDisableResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDisableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingDisableResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingDisableResponseEnvelopeSuccessTrue SecondaryDNSOutgoingDisableResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingEnableParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSOutgoingEnableResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingEnableResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingEnableResponseEnvelopeMessages `json:"messages,required"`
	// The zone transfer status of a primary zone
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingEnableResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingEnableResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingEnableResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingEnableResponseEnvelope]
type secondaryDNSOutgoingEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingEnableResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSOutgoingEnableResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingEnableResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingEnableResponseEnvelopeErrors]
type secondaryDNSOutgoingEnableResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingEnableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingEnableResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    secondaryDNSOutgoingEnableResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingEnableResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingEnableResponseEnvelopeMessages]
type secondaryDNSOutgoingEnableResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingEnableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingEnableResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingEnableResponseEnvelopeSuccessTrue SecondaryDNSOutgoingEnableResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingForceNotifyParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSOutgoingForceNotifyResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingForceNotifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingForceNotifyResponseEnvelopeMessages `json:"messages,required"`
	// When force_notify query parameter is set to true, the response is a simple
	// string
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingForceNotifyResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingForceNotifyResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingForceNotifyResponseEnvelopeJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingForceNotifyResponseEnvelope]
type secondaryDNSOutgoingForceNotifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingForceNotifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingForceNotifyResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    secondaryDNSOutgoingForceNotifyResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingForceNotifyResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingForceNotifyResponseEnvelopeErrors]
type secondaryDNSOutgoingForceNotifyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingForceNotifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingForceNotifyResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    secondaryDNSOutgoingForceNotifyResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingForceNotifyResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SecondaryDNSOutgoingForceNotifyResponseEnvelopeMessages]
type secondaryDNSOutgoingForceNotifyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingForceNotifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingForceNotifyResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingForceNotifyResponseEnvelopeSuccessTrue SecondaryDNSOutgoingForceNotifyResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingGetParams struct {
	ZoneID param.Field[interface{}] `path:"zone_id,required"`
}

type SecondaryDNSOutgoingGetResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingGetResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SecondaryDNSOutgoingGetResponseEnvelope]
type secondaryDNSOutgoingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    secondaryDNSOutgoingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingGetResponseEnvelopeErrors]
type secondaryDNSOutgoingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    secondaryDNSOutgoingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingGetResponseEnvelopeMessages]
type secondaryDNSOutgoingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingGetResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingGetResponseEnvelopeSuccessTrue SecondaryDNSOutgoingGetResponseEnvelopeSuccess = true
)
