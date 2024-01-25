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
func (r *AccountAddressPrefixBgpStatusService) IPAddressManagementDynamicAdvertisementGetAdvertisementStatus(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Advertise or withdraw BGP route for a prefix.
func (r *AccountAddressPrefixBgpStatusService) IPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatus(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams, opts ...option.RequestOption) (res *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse struct {
	Errors   []AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseError   `json:"errors"`
	Messages []AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseMessage `json:"messages"`
	Result   AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseSuccess `json:"success"`
	JSON    accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseJSON    `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseError struct {
	Code    int64                                                                                                       `json:"code,required"`
	Message string                                                                                                      `json:"message,required"`
	JSON    accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseErrorJSON `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseError]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseMessage struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseMessageJSON `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseMessage]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseResult struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                                                                                    `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseResultJSON `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseResult]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseResultJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseSuccess bool

const (
	AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseSuccessTrue AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementGetAdvertisementStatusResponseSuccess = true
)

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse struct {
	Errors   []AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseError   `json:"errors"`
	Messages []AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseMessage `json:"messages"`
	Result   AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseSuccess `json:"success"`
	JSON    accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseJSON    `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseError struct {
	Code    int64                                                                                                                       `json:"code,required"`
	Message string                                                                                                                      `json:"message,required"`
	JSON    accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseErrorJSON `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseError]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseMessage struct {
	Code    int64                                                                                                                         `json:"code,required"`
	Message string                                                                                                                        `json:"message,required"`
	JSON    accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseMessageJSON `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseMessage]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseResult struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                                                                                                    `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseResultJSON `json:"-"`
}

// accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseResult]
type accountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseResultJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseSuccess bool

const (
	AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseSuccessTrue AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusResponseSuccess = true
)

type AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams struct {
	// Enablement of prefix advertisement to the Internet.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r AccountAddressPrefixBgpStatusIPAddressManagementDynamicAdvertisementUpdatePrefixDynamicAdvertisementStatusParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
