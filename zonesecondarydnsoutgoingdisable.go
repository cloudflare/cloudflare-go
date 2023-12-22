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
func (r *ZoneSecondaryDNSOutgoingDisableService) SecondaryDNSPrimaryZoneDisableOutgoingZoneTransfers(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *DisableTransferResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/disable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type DisableTransferResponse struct {
	Errors   []DisableTransferResponseError   `json:"errors"`
	Messages []DisableTransferResponseMessage `json:"messages"`
	// The zone transfer status of a primary zone
	Result string `json:"result"`
	// Whether the API call was successful
	Success DisableTransferResponseSuccess `json:"success"`
	JSON    disableTransferResponseJSON    `json:"-"`
}

// disableTransferResponseJSON contains the JSON metadata for the struct
// [DisableTransferResponse]
type disableTransferResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisableTransferResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DisableTransferResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    disableTransferResponseErrorJSON `json:"-"`
}

// disableTransferResponseErrorJSON contains the JSON metadata for the struct
// [DisableTransferResponseError]
type disableTransferResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisableTransferResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DisableTransferResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    disableTransferResponseMessageJSON `json:"-"`
}

// disableTransferResponseMessageJSON contains the JSON metadata for the struct
// [DisableTransferResponseMessage]
type disableTransferResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DisableTransferResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DisableTransferResponseSuccess bool

const (
	DisableTransferResponseSuccessTrue DisableTransferResponseSuccess = true
)
