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

// ZoneSecondaryDNSOutgoingDisableService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSecondaryDNSOutgoingDisableService] method instead.
type ZoneSecondaryDNSOutgoingDisableService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSOutgoingDisableService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSOutgoingDisableService(opts ...option.RequestOption) (r *ZoneSecondaryDNSOutgoingDisableService) {
	r = &ZoneSecondaryDNSOutgoingDisableService{}
	r.Options = opts
	return
}

// Disable outgoing zone transfers for primary zone and clears IXFR backlog of
// primary zone.
func (r *ZoneSecondaryDNSOutgoingDisableService) SecondaryDNSPrimaryZoneDisableOutgoingZoneTransfers(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseMessage `json:"messages"`
	// The zone transfer status of a primary zone
	Result string `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponse]
type zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseError]
type zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseMessage]
type zoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseSuccessTrue ZoneSecondaryDNSOutgoingDisableSecondaryDNSPrimaryZoneDisableOutgoingZoneTransfersResponseSuccess = true
)
