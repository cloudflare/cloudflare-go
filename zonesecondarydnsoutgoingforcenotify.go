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

// ZoneSecondaryDNSOutgoingForceNotifyService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSecondaryDNSOutgoingForceNotifyService] method instead.
type ZoneSecondaryDNSOutgoingForceNotifyService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSOutgoingForceNotifyService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSecondaryDNSOutgoingForceNotifyService(opts ...option.RequestOption) (r *ZoneSecondaryDNSOutgoingForceNotifyService) {
	r = &ZoneSecondaryDNSOutgoingForceNotifyService{}
	r.Options = opts
	return
}

// Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
func (r *ZoneSecondaryDNSOutgoingForceNotifyService) SecondaryDNSPrimaryZoneForceDNSNotify(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/force_notify", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponse struct {
	Errors   []ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseError   `json:"errors"`
	Messages []ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseMessage `json:"messages"`
	// When force_notify query parameter is set to true, the response is a simple
	// string
	Result string `json:"result"`
	// Whether the API call was successful
	Success ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseSuccess `json:"success"`
	JSON    zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseJSON    `json:"-"`
}

// zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponse]
type zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseError struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseErrorJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseError]
type zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseMessage struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseMessageJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseMessage]
type zoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseSuccessTrue ZoneSecondaryDNSOutgoingForceNotifySecondaryDNSPrimaryZoneForceDNSNotifyResponseSuccess = true
)
