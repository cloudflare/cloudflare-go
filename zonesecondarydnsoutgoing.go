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
func (r *ZoneSecondaryDNSOutgoingService) SecondaryDNSPrimaryZoneNewPrimaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationParams, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) SecondaryDNSPrimaryZonePrimaryZoneConfigurationDetails(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) SecondaryDNSPrimaryZoneUpdatePrimaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationParams, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

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

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseError]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseMessage]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseResult struct {
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
	SoaSerial float64                                                                                      `json:"soa_serial"`
	JSON      zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseResult]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseResultJSON struct {
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

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseSuccessTrue ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseSuccess = true
)

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseError]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseMessage]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseResult struct {
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
	SoaSerial float64                                                                                          `json:"soa_serial"`
	JSON      zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseResult]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseResultJSON struct {
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

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseSuccessTrue ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseSuccess = true
)

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseError struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseError]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseMessage struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseMessage]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseResult struct {
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
	SoaSerial float64                                                                                         `json:"soa_serial"`
	JSON      zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseResult]
type zoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseResultJSON struct {
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

func (r *ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseSuccessTrue ZoneSecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseSuccess = true
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
