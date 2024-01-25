// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSecondaryDNSForceAxfrService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSecondaryDNSForceAxfrService] method instead.
type ZoneSecondaryDNSForceAxfrService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSForceAxfrService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSForceAxfrService(opts ...option.RequestOption) (r *ZoneSecondaryDNSForceAxfrService) {
	r = &ZoneSecondaryDNSForceAxfrService{}
	r.Options = opts
	return
}

// Sends AXFR zone transfer request to primary nameserver(s).
func (r *ZoneSecondaryDNSForceAxfrService) SecondaryDNSSecondaryZoneForceAxfr(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/force_axfr", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponse struct {
	Errors   []ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseMessage `json:"messages"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result string `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseJSON    `json:"-"`
}

// zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseJSON contains
// the JSON metadata for the struct
// [ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponse]
type zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseError]
type zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseMessage]
type zoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseSuccess bool

const (
	ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseSuccessTrue ZoneSecondaryDNSForceAxfrSecondaryDNSSecondaryZoneForceAxfrResponseSuccess = true
)
