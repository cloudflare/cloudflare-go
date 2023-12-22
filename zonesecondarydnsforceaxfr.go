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
func (r *ZoneSecondaryDNSForceAxfrService) SecondaryDNSSecondaryZoneForceAxfr(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *ForceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/secondary_dns/force_axfr", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ForceResponse struct {
	Errors   []ForceResponseError   `json:"errors"`
	Messages []ForceResponseMessage `json:"messages"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result string `json:"result"`
	// Whether the API call was successful
	Success ForceResponseSuccess `json:"success"`
	JSON    forceResponseJSON    `json:"-"`
}

// forceResponseJSON contains the JSON metadata for the struct [ForceResponse]
type forceResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ForceResponseError struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    forceResponseErrorJSON `json:"-"`
}

// forceResponseErrorJSON contains the JSON metadata for the struct
// [ForceResponseError]
type forceResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ForceResponseMessage struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    forceResponseMessageJSON `json:"-"`
}

// forceResponseMessageJSON contains the JSON metadata for the struct
// [ForceResponseMessage]
type forceResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ForceResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ForceResponseSuccess bool

const (
	ForceResponseSuccessTrue ForceResponseSuccess = true
)
