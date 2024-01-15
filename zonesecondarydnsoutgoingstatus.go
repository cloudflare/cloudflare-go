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

// ZoneSecondaryDNSOutgoingStatusService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSecondaryDNSOutgoingStatusService] method instead.
type ZoneSecondaryDNSOutgoingStatusService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSOutgoingStatusService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSOutgoingStatusService(opts ...option.RequestOption) (r *ZoneSecondaryDNSOutgoingStatusService) {
	r = &ZoneSecondaryDNSOutgoingStatusService{}
	r.Options = opts
	return
}

// Get primary zone transfer status.
func (r *ZoneSecondaryDNSOutgoingStatusService) SecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatus(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/status", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseMessage `json:"messages"`
	// The zone transfer status of a primary zone
	Result string `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponse]
type zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseError]
type zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseMessage]
type zoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseSuccessTrue ZoneSecondaryDNSOutgoingStatusSecondaryDNSPrimaryZoneGetOutgoingZoneTransferStatusResponseSuccess = true
)
