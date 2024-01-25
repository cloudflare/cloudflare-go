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

// ZoneSecondaryDNSOutgoingEnableService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSecondaryDNSOutgoingEnableService] method instead.
type ZoneSecondaryDNSOutgoingEnableService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSOutgoingEnableService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSOutgoingEnableService(opts ...option.RequestOption) (r *ZoneSecondaryDNSOutgoingEnableService) {
	r = &ZoneSecondaryDNSOutgoingEnableService{}
	r.Options = opts
	return
}

// Enable outgoing zone transfers for primary zone.
func (r *ZoneSecondaryDNSOutgoingEnableService) SecondaryDNSPrimaryZoneEnableOutgoingZoneTransfers(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseMessage `json:"messages"`
	// The zone transfer status of a primary zone
	Result string `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponse]
type zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseError]
type zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseMessage]
type zoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseSuccessTrue ZoneSecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseSuccess = true
)
