// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDeviceNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDeviceNetworkService]
// method instead.
type AccountDeviceNetworkService struct {
	Options []option.RequestOption
}

// NewAccountDeviceNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceNetworkService(opts ...option.RequestOption) (r *AccountDeviceNetworkService) {
	r = &AccountDeviceNetworkService{}
	r.Options = opts
	return
}

// Fetches details for a single managed network.
func (r *AccountDeviceNetworkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDeviceNetworkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured device managed network.
func (r *AccountDeviceNetworkService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDeviceNetworkUpdateParams, opts ...option.RequestOption) (res *AccountDeviceNetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a device managed network and fetches a list of the remaining device
// managed networks for an account.
func (r *AccountDeviceNetworkService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDeviceNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new device managed network.
func (r *AccountDeviceNetworkService) DeviceManagedNetworksNewDeviceManagedNetwork(ctx context.Context, identifier interface{}, body AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams, opts ...option.RequestOption) (res *AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a list of managed networks for an account.
func (r *AccountDeviceNetworkService) DeviceManagedNetworksListDeviceManagedNetworks(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDeviceNetworkGetResponse struct {
	Errors   []AccountDeviceNetworkGetResponseError   `json:"errors"`
	Messages []AccountDeviceNetworkGetResponseMessage `json:"messages"`
	Result   AccountDeviceNetworkGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceNetworkGetResponseSuccess `json:"success"`
	JSON    accountDeviceNetworkGetResponseJSON    `json:"-"`
}

// accountDeviceNetworkGetResponseJSON contains the JSON metadata for the struct
// [AccountDeviceNetworkGetResponse]
type accountDeviceNetworkGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountDeviceNetworkGetResponseErrorJSON `json:"-"`
}

// accountDeviceNetworkGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkGetResponseError]
type accountDeviceNetworkGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountDeviceNetworkGetResponseMessageJSON `json:"-"`
}

// accountDeviceNetworkGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkGetResponseMessage]
type accountDeviceNetworkGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkGetResponseResult struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config AccountDeviceNetworkGetResponseResultConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type AccountDeviceNetworkGetResponseResultType `json:"type"`
	JSON accountDeviceNetworkGetResponseResultJSON `json:"-"`
}

// accountDeviceNetworkGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkGetResponseResult]
type accountDeviceNetworkGetResponseResultJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type AccountDeviceNetworkGetResponseResultConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                          `json:"sha256"`
	JSON   accountDeviceNetworkGetResponseResultConfigJSON `json:"-"`
}

// accountDeviceNetworkGetResponseResultConfigJSON contains the JSON metadata for
// the struct [AccountDeviceNetworkGetResponseResultConfig]
type accountDeviceNetworkGetResponseResultConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkGetResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type AccountDeviceNetworkGetResponseResultType string

const (
	AccountDeviceNetworkGetResponseResultTypeTls AccountDeviceNetworkGetResponseResultType = "tls"
)

// Whether the API call was successful.
type AccountDeviceNetworkGetResponseSuccess bool

const (
	AccountDeviceNetworkGetResponseSuccessTrue AccountDeviceNetworkGetResponseSuccess = true
)

type AccountDeviceNetworkUpdateResponse struct {
	Errors   []AccountDeviceNetworkUpdateResponseError   `json:"errors"`
	Messages []AccountDeviceNetworkUpdateResponseMessage `json:"messages"`
	Result   AccountDeviceNetworkUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceNetworkUpdateResponseSuccess `json:"success"`
	JSON    accountDeviceNetworkUpdateResponseJSON    `json:"-"`
}

// accountDeviceNetworkUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDeviceNetworkUpdateResponse]
type accountDeviceNetworkUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDeviceNetworkUpdateResponseErrorJSON `json:"-"`
}

// accountDeviceNetworkUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkUpdateResponseError]
type accountDeviceNetworkUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountDeviceNetworkUpdateResponseMessageJSON `json:"-"`
}

// accountDeviceNetworkUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkUpdateResponseMessage]
type accountDeviceNetworkUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkUpdateResponseResult struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config AccountDeviceNetworkUpdateResponseResultConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type AccountDeviceNetworkUpdateResponseResultType `json:"type"`
	JSON accountDeviceNetworkUpdateResponseResultJSON `json:"-"`
}

// accountDeviceNetworkUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkUpdateResponseResult]
type accountDeviceNetworkUpdateResponseResultJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type AccountDeviceNetworkUpdateResponseResultConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                             `json:"sha256"`
	JSON   accountDeviceNetworkUpdateResponseResultConfigJSON `json:"-"`
}

// accountDeviceNetworkUpdateResponseResultConfigJSON contains the JSON metadata
// for the struct [AccountDeviceNetworkUpdateResponseResultConfig]
type accountDeviceNetworkUpdateResponseResultConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkUpdateResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type AccountDeviceNetworkUpdateResponseResultType string

const (
	AccountDeviceNetworkUpdateResponseResultTypeTls AccountDeviceNetworkUpdateResponseResultType = "tls"
)

// Whether the API call was successful.
type AccountDeviceNetworkUpdateResponseSuccess bool

const (
	AccountDeviceNetworkUpdateResponseSuccessTrue AccountDeviceNetworkUpdateResponseSuccess = true
)

type AccountDeviceNetworkDeleteResponse struct {
	Errors     []AccountDeviceNetworkDeleteResponseError    `json:"errors"`
	Messages   []AccountDeviceNetworkDeleteResponseMessage  `json:"messages"`
	Result     []AccountDeviceNetworkDeleteResponseResult   `json:"result"`
	ResultInfo AccountDeviceNetworkDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDeviceNetworkDeleteResponseSuccess `json:"success"`
	JSON    accountDeviceNetworkDeleteResponseJSON    `json:"-"`
}

// accountDeviceNetworkDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDeviceNetworkDeleteResponse]
type accountDeviceNetworkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDeviceNetworkDeleteResponseErrorJSON `json:"-"`
}

// accountDeviceNetworkDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkDeleteResponseError]
type accountDeviceNetworkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountDeviceNetworkDeleteResponseMessageJSON `json:"-"`
}

// accountDeviceNetworkDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkDeleteResponseMessage]
type accountDeviceNetworkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeleteResponseResult struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config AccountDeviceNetworkDeleteResponseResultConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type AccountDeviceNetworkDeleteResponseResultType `json:"type"`
	JSON accountDeviceNetworkDeleteResponseResultJSON `json:"-"`
}

// accountDeviceNetworkDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDeviceNetworkDeleteResponseResult]
type accountDeviceNetworkDeleteResponseResultJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type AccountDeviceNetworkDeleteResponseResultConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                             `json:"sha256"`
	JSON   accountDeviceNetworkDeleteResponseResultConfigJSON `json:"-"`
}

// accountDeviceNetworkDeleteResponseResultConfigJSON contains the JSON metadata
// for the struct [AccountDeviceNetworkDeleteResponseResultConfig]
type accountDeviceNetworkDeleteResponseResultConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeleteResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type AccountDeviceNetworkDeleteResponseResultType string

const (
	AccountDeviceNetworkDeleteResponseResultTypeTls AccountDeviceNetworkDeleteResponseResultType = "tls"
)

type AccountDeviceNetworkDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       accountDeviceNetworkDeleteResponseResultInfoJSON `json:"-"`
}

// accountDeviceNetworkDeleteResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountDeviceNetworkDeleteResponseResultInfo]
type accountDeviceNetworkDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceNetworkDeleteResponseSuccess bool

const (
	AccountDeviceNetworkDeleteResponseSuccessTrue AccountDeviceNetworkDeleteResponseSuccess = true
)

type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse struct {
	Errors   []AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseError   `json:"errors"`
	Messages []AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseMessage `json:"messages"`
	Result   AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseSuccess `json:"success"`
	JSON    accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseJSON    `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse]
type accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseErrorJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseError]
type accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseMessageJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseMessage]
type accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResult struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultType `json:"type"`
	JSON accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResult]
type accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                                                                   `json:"sha256"`
	JSON   accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultConfigJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultConfigJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultConfig]
type accountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultType string

const (
	AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultTypeTls AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseResultType = "tls"
)

// Whether the API call was successful.
type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseSuccess bool

const (
	AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseSuccessTrue AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseSuccess = true
)

type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse struct {
	Errors     []AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseError    `json:"errors"`
	Messages   []AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseMessage  `json:"messages"`
	Result     []AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResult   `json:"result"`
	ResultInfo AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseSuccess `json:"success"`
	JSON    accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseJSON    `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse]
type accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseErrorJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseError]
type accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseMessageJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseMessage]
type accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResult struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultType `json:"type"`
	JSON accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResult]
type accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                                                                     `json:"sha256"`
	JSON   accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultConfigJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultConfigJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultConfig]
type accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultType string

const (
	AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultTypeTls AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultType = "tls"
)

type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                  `json:"total_count"`
	JSON       accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultInfoJSON `json:"-"`
}

// accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultInfo]
type accountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseSuccess bool

const (
	AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseSuccessTrue AccountDeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseSuccess = true
)

type AccountDeviceNetworkUpdateParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[AccountDeviceNetworkUpdateParamsConfig] `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name"`
	// The type of device managed network.
	Type param.Field[AccountDeviceNetworkUpdateParamsType] `json:"type"`
}

func (r AccountDeviceNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type AccountDeviceNetworkUpdateParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r AccountDeviceNetworkUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type AccountDeviceNetworkUpdateParamsType string

const (
	AccountDeviceNetworkUpdateParamsTypeTls AccountDeviceNetworkUpdateParamsType = "tls"
)

type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsConfig] `json:"config,required"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name,required"`
	// The type of device managed network.
	Type param.Field[AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType] `json:"type,required"`
}

func (r AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType string

const (
	AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsTypeTls AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType = "tls"
)
