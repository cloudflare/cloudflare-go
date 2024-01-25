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

// AccountGatewayLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayLocationService]
// method instead.
type AccountGatewayLocationService struct {
	Options []option.RequestOption
}

// NewAccountGatewayLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayLocationService(opts ...option.RequestOption) (r *AccountGatewayLocationService) {
	r = &AccountGatewayLocationService{}
	r.Options = opts
	return
}

// Fetches a single Zero Trust Gateway location.
func (r *AccountGatewayLocationService) Get(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *AccountGatewayLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust Gateway location.
func (r *AccountGatewayLocationService) Update(ctx context.Context, identifier interface{}, uuid interface{}, body AccountGatewayLocationUpdateParams, opts ...option.RequestOption) (res *AccountGatewayLocationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a configured Zero Trust Gateway location.
func (r *AccountGatewayLocationService) Delete(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *AccountGatewayLocationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Zero Trust Gateway location.
func (r *AccountGatewayLocationService) ZeroTrustGatewayLocationsNewZeroTrustGatewayLocation(ctx context.Context, identifier interface{}, body AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams, opts ...option.RequestOption) (res *AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches Zero Trust Gateway locations for an account.
func (r *AccountGatewayLocationService) ZeroTrustGatewayLocationsListZeroTrustGatewayLocations(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/locations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayLocationGetResponse struct {
	Errors   []AccountGatewayLocationGetResponseError   `json:"errors"`
	Messages []AccountGatewayLocationGetResponseMessage `json:"messages"`
	Result   AccountGatewayLocationGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayLocationGetResponseSuccess `json:"success"`
	JSON    accountGatewayLocationGetResponseJSON    `json:"-"`
}

// accountGatewayLocationGetResponseJSON contains the JSON metadata for the struct
// [AccountGatewayLocationGetResponse]
type accountGatewayLocationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountGatewayLocationGetResponseErrorJSON `json:"-"`
}

// accountGatewayLocationGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayLocationGetResponseError]
type accountGatewayLocationGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountGatewayLocationGetResponseMessageJSON `json:"-"`
}

// accountGatewayLocationGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayLocationGetResponseMessage]
type accountGatewayLocationGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationGetResponseResult struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []AccountGatewayLocationGetResponseResultNetwork `json:"networks"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      accountGatewayLocationGetResponseResultJSON      `json:"-"`
}

// accountGatewayLocationGetResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayLocationGetResponseResult]
type accountGatewayLocationGetResponseResultJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayLocationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationGetResponseResultNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                             `json:"network,required"`
	JSON    accountGatewayLocationGetResponseResultNetworkJSON `json:"-"`
}

// accountGatewayLocationGetResponseResultNetworkJSON contains the JSON metadata
// for the struct [AccountGatewayLocationGetResponseResultNetwork]
type accountGatewayLocationGetResponseResultNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationGetResponseResultNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayLocationGetResponseSuccess bool

const (
	AccountGatewayLocationGetResponseSuccessTrue AccountGatewayLocationGetResponseSuccess = true
)

type AccountGatewayLocationUpdateResponse struct {
	Errors   []AccountGatewayLocationUpdateResponseError   `json:"errors"`
	Messages []AccountGatewayLocationUpdateResponseMessage `json:"messages"`
	Result   AccountGatewayLocationUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayLocationUpdateResponseSuccess `json:"success"`
	JSON    accountGatewayLocationUpdateResponseJSON    `json:"-"`
}

// accountGatewayLocationUpdateResponseJSON contains the JSON metadata for the
// struct [AccountGatewayLocationUpdateResponse]
type accountGatewayLocationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountGatewayLocationUpdateResponseErrorJSON `json:"-"`
}

// accountGatewayLocationUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayLocationUpdateResponseError]
type accountGatewayLocationUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountGatewayLocationUpdateResponseMessageJSON `json:"-"`
}

// accountGatewayLocationUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountGatewayLocationUpdateResponseMessage]
type accountGatewayLocationUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationUpdateResponseResult struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []AccountGatewayLocationUpdateResponseResultNetwork `json:"networks"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      accountGatewayLocationUpdateResponseResultJSON      `json:"-"`
}

// accountGatewayLocationUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountGatewayLocationUpdateResponseResult]
type accountGatewayLocationUpdateResponseResultJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayLocationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationUpdateResponseResultNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                                `json:"network,required"`
	JSON    accountGatewayLocationUpdateResponseResultNetworkJSON `json:"-"`
}

// accountGatewayLocationUpdateResponseResultNetworkJSON contains the JSON metadata
// for the struct [AccountGatewayLocationUpdateResponseResultNetwork]
type accountGatewayLocationUpdateResponseResultNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationUpdateResponseResultNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayLocationUpdateResponseSuccess bool

const (
	AccountGatewayLocationUpdateResponseSuccessTrue AccountGatewayLocationUpdateResponseSuccess = true
)

type AccountGatewayLocationDeleteResponse struct {
	Errors   []AccountGatewayLocationDeleteResponseError   `json:"errors"`
	Messages []AccountGatewayLocationDeleteResponseMessage `json:"messages"`
	Result   interface{}                                   `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayLocationDeleteResponseSuccess `json:"success"`
	JSON    accountGatewayLocationDeleteResponseJSON    `json:"-"`
}

// accountGatewayLocationDeleteResponseJSON contains the JSON metadata for the
// struct [AccountGatewayLocationDeleteResponse]
type accountGatewayLocationDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountGatewayLocationDeleteResponseErrorJSON `json:"-"`
}

// accountGatewayLocationDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayLocationDeleteResponseError]
type accountGatewayLocationDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountGatewayLocationDeleteResponseMessageJSON `json:"-"`
}

// accountGatewayLocationDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountGatewayLocationDeleteResponseMessage]
type accountGatewayLocationDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayLocationDeleteResponseSuccess bool

const (
	AccountGatewayLocationDeleteResponseSuccessTrue AccountGatewayLocationDeleteResponseSuccess = true
)

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse struct {
	Errors   []AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseError   `json:"errors"`
	Messages []AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseMessage `json:"messages"`
	Result   AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseSuccess `json:"success"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseJSON    `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse]
type accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseErrorJSON `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseError]
type accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseMessageJSON `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseMessage]
type accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResult struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultNetwork `json:"networks"`
	UpdatedAt time.Time                                                                                         `json:"updated_at" format:"date-time"`
	JSON      accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultJSON      `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResult]
type accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                                                                              `json:"network,required"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultNetworkJSON `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultNetworkJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultNetwork]
type accountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseResultNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseSuccess bool

const (
	AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseSuccessTrue AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationResponseSuccess = true
)

type AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse struct {
	Errors     []AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseError    `json:"errors"`
	Messages   []AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseMessage  `json:"messages"`
	Result     []AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResult   `json:"result"`
	ResultInfo AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseSuccess `json:"success"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseJSON    `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse]
type accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseError struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseErrorJSON `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseError]
type accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseMessage struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseMessageJSON `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseMessage]
type accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResult struct {
	ID interface{} `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	Networks  []AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultNetwork `json:"networks"`
	UpdatedAt time.Time                                                                                           `json:"updated_at" format:"date-time"`
	JSON      accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultJSON      `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResult]
type accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultJSON struct {
	ID            apijson.Field
	ClientDefault apijson.Field
	CreatedAt     apijson.Field
	DohSubdomain  apijson.Field
	EcsSupport    apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	Networks      apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string                                                                                                `json:"network,required"`
	JSON    accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultNetworkJSON `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultNetworkJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultNetwork]
type accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                            `json:"total_count"`
	JSON       accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultInfoJSON `json:"-"`
}

// accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultInfo]
type accountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseSuccess bool

const (
	AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseSuccessTrue AccountGatewayLocationZeroTrustGatewayLocationsListZeroTrustGatewayLocationsResponseSuccess = true
)

type AccountGatewayLocationUpdateParams struct {
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]AccountGatewayLocationUpdateParamsNetwork] `json:"networks"`
}

func (r AccountGatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayLocationUpdateParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r AccountGatewayLocationUpdateParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams struct {
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// A list of network ranges that requests from this location would originate from.
	Networks param.Field[[]AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParamsNetwork] `json:"networks"`
}

func (r AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParamsNetwork struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r AccountGatewayLocationZeroTrustGatewayLocationsNewZeroTrustGatewayLocationParamsNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
