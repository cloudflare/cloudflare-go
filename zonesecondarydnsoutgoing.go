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

// ZoneSecondaryDNSOutgoingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSecondaryDNSOutgoingService] method instead.
type ZoneSecondaryDNSOutgoingService struct {
	Options       []option.RequestOption
	Disables      *ZoneSecondaryDNSOutgoingDisableService
	Enables       *ZoneSecondaryDNSOutgoingEnableService
	ForceNotifies *ZoneSecondaryDNSOutgoingForceNotifyService
	Statuses      *ZoneSecondaryDNSOutgoingStatusService
}

// NewZoneSecondaryDNSOutgoingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSOutgoingService(opts ...option.RequestOption) (r *ZoneSecondaryDNSOutgoingService) {
	r = &ZoneSecondaryDNSOutgoingService{}
	r.Options = opts
	r.Disables = NewZoneSecondaryDNSOutgoingDisableService(opts...)
	r.Enables = NewZoneSecondaryDNSOutgoingEnableService(opts...)
	r.ForceNotifies = NewZoneSecondaryDNSOutgoingForceNotifyService(opts...)
	r.Statuses = NewZoneSecondaryDNSOutgoingStatusService(opts...)
	return
}

// Delete primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) Delete(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) SecondaryDNSPrimaryZoneNewPrimaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationParams, opts ...option.RequestOption) (res *SingleResponseOutgoing, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) SecondaryDNSPrimaryZonePrimaryZoneConfigurationDetails(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *SingleResponseOutgoing, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) SecondaryDNSPrimaryZoneUpdatePrimaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationParams, opts ...option.RequestOption) (res *SingleResponseOutgoing, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SingleResponseOutgoing struct {
	Errors   []SingleResponseOutgoingError   `json:"errors"`
	Messages []SingleResponseOutgoingMessage `json:"messages"`
	Result   SingleResponseOutgoingResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseOutgoingSuccess `json:"success"`
	JSON    singleResponseOutgoingJSON    `json:"-"`
}

// singleResponseOutgoingJSON contains the JSON metadata for the struct
// [SingleResponseOutgoing]
type singleResponseOutgoingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseOutgoing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseOutgoingError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseOutgoingErrorJSON `json:"-"`
}

// singleResponseOutgoingErrorJSON contains the JSON metadata for the struct
// [SingleResponseOutgoingError]
type singleResponseOutgoingErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseOutgoingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseOutgoingMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseOutgoingMessageJSON `json:"-"`
}

// singleResponseOutgoingMessageJSON contains the JSON metadata for the struct
// [SingleResponseOutgoingMessage]
type singleResponseOutgoingMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseOutgoingMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseOutgoingResult struct {
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
	SoaSerial float64                          `json:"soa_serial"`
	JSON      singleResponseOutgoingResultJSON `json:"-"`
}

// singleResponseOutgoingResultJSON contains the JSON metadata for the struct
// [SingleResponseOutgoingResult]
type singleResponseOutgoingResultJSON struct {
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

func (r *SingleResponseOutgoingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponseOutgoingSuccess bool

const (
	SingleResponseOutgoingSuccessTrue SingleResponseOutgoingSuccess = true
)

type ZoneSecondaryDNSOutgoingDeleteResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingDeleteResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingDeleteResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSOutgoingDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingDeleteResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingDeleteResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneSecondaryDNSOutgoingDeleteResponse]
type zoneSecondaryDNSOutgoingDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingDeleteResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingDeleteResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingDeleteResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSecondaryDNSOutgoingDeleteResponseError]
type zoneSecondaryDNSOutgoingDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingDeleteResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingDeleteResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSecondaryDNSOutgoingDeleteResponseMessage]
type zoneSecondaryDNSOutgoingDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingDeleteResponseResult struct {
	ID   interface{}                                      `json:"id"`
	JSON zoneSecondaryDNSOutgoingDeleteResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneSecondaryDNSOutgoingDeleteResponseResult]
type zoneSecondaryDNSOutgoingDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingDeleteResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingDeleteResponseSuccessTrue ZoneSecondaryDNSOutgoingDeleteResponseSuccess = true
)

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationParams struct {
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationParams struct {
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
