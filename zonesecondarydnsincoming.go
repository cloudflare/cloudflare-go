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

// ZoneSecondaryDNSIncomingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSecondaryDNSIncomingService] method instead.
type ZoneSecondaryDNSIncomingService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSIncomingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSIncomingService(opts ...option.RequestOption) (r *ZoneSecondaryDNSIncomingService) {
	r = &ZoneSecondaryDNSIncomingService{}
	r.Options = opts
	return
}

// Delete secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) Delete(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSIncomingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) SecondaryDNSSecondaryZoneNewSecondaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationParams, opts ...option.RequestOption) (res *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) SecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetails(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) SecondaryDNSSecondaryZoneUpdateSecondaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationParams, opts ...option.RequestOption) (res *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneSecondaryDNSIncomingDeleteResponse struct {
	Errors   []ZoneSecondaryDNSIncomingDeleteResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSIncomingDeleteResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSIncomingDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSIncomingDeleteResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSIncomingDeleteResponseJSON    `json:"-"`
}

// zoneSecondaryDNSIncomingDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneSecondaryDNSIncomingDeleteResponse]
type zoneSecondaryDNSIncomingDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingDeleteResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingDeleteResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSIncomingDeleteResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSecondaryDNSIncomingDeleteResponseError]
type zoneSecondaryDNSIncomingDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingDeleteResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingDeleteResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSIncomingDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSecondaryDNSIncomingDeleteResponseMessage]
type zoneSecondaryDNSIncomingDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingDeleteResponseResult struct {
	ID   interface{}                                      `json:"id"`
	JSON zoneSecondaryDNSIncomingDeleteResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSIncomingDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneSecondaryDNSIncomingDeleteResponseResult]
type zoneSecondaryDNSIncomingDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSIncomingDeleteResponseSuccess bool

const (
	ZoneSecondaryDNSIncomingDeleteResponseSuccessTrue ZoneSecondaryDNSIncomingDeleteResponseSuccess = true
)

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse struct {
	Errors   []ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseJSON    `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseError]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseMessage]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseResult struct {
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
	SoaSerial float64                                                                                          `json:"soa_serial"`
	JSON      zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseResult]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseResultJSON struct {
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

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseSuccess bool

const (
	ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseSuccessTrue ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseSuccess = true
)

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse struct {
	Errors   []ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseJSON    `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseError]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseMessage]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseResult struct {
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
	SoaSerial float64                                                                                              `json:"soa_serial"`
	JSON      zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseResult]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseResultJSON struct {
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

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseSuccess bool

const (
	ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseSuccessTrue ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseSuccess = true
)

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse struct {
	Errors   []ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseMessage `json:"messages"`
	Result   ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseJSON    `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseError struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseError]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseMessage struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseMessage]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseResult struct {
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
	SoaSerial float64                                                                                             `json:"soa_serial"`
	JSON      zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseResultJSON `json:"-"`
}

// zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseResult]
type zoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseResultJSON struct {
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

func (r *ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseSuccess bool

const (
	ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseSuccessTrue ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseSuccess = true
)

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationParams struct {
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationParams struct {
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r ZoneSecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
