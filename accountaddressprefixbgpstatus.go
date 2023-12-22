// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAddressPrefixBgpStatusService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAddressPrefixBgpStatusService] method instead.
type AccountAddressPrefixBgpStatusService struct {
	Options []option.RequestOption
}

// NewAccountAddressPrefixBgpStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressPrefixBgpStatusService(opts ...option.RequestOption) (r *AccountAddressPrefixBgpStatusService) {
	r = &AccountAddressPrefixBgpStatusService{}
	r.Options = opts
	return
}

// List the current advertisement state for a prefix.
func (r *AccountAddressPrefixBgpStatusService) IPAddressManagementDynamicAdvertisementGetAdvertisementStatus(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AdvertisedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Advertise or withdraw BGP route for a prefix.
func (r *AccountAddressPrefixBgpStatusService) IPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatus(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams, opts ...option.RequestOption) (res *AdvertisedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AdvertisedResponse struct {
	Errors   []AdvertisedResponseError   `json:"errors"`
	Messages []AdvertisedResponseMessage `json:"messages"`
	Result   AdvertisedResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AdvertisedResponseSuccess `json:"success"`
	JSON    advertisedResponseJSON    `json:"-"`
}

// advertisedResponseJSON contains the JSON metadata for the struct
// [AdvertisedResponse]
type advertisedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvertisedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AdvertisedResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    advertisedResponseErrorJSON `json:"-"`
}

// advertisedResponseErrorJSON contains the JSON metadata for the struct
// [AdvertisedResponseError]
type advertisedResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvertisedResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AdvertisedResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    advertisedResponseMessageJSON `json:"-"`
}

// advertisedResponseMessageJSON contains the JSON metadata for the struct
// [AdvertisedResponseMessage]
type advertisedResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvertisedResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AdvertisedResponseResult struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                    `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 advertisedResponseResultJSON `json:"-"`
}

// advertisedResponseResultJSON contains the JSON metadata for the struct
// [AdvertisedResponseResult]
type advertisedResponseResultJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AdvertisedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AdvertisedResponseSuccess bool

const (
	AdvertisedResponseSuccessTrue AdvertisedResponseSuccess = true
)

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
