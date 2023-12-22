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

// Fetch a single Managed Network.
func (r *AccountDeviceNetworkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SingleResponse8ngPd7Kr, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Device Managed Network.
func (r *AccountDeviceNetworkService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountDeviceNetworkUpdateParams, opts ...option.RequestOption) (res *SingleResponse8ngPd7Kr, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Device Managed Network. Returns the remaining Device Managed Networks
// for the account.
func (r *AccountDeviceNetworkService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *ResponseCollectionYsH63Sfm, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Device Managed Network.
func (r *AccountDeviceNetworkService) DeviceManagedNetworksNewDeviceManagedNetwork(ctx context.Context, identifier interface{}, body AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams, opts ...option.RequestOption) (res *SingleResponse8ngPd7Kr, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Managed Networks for an account.
func (r *AccountDeviceNetworkService) DeviceManagedNetworksListDeviceManagedNetworks(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ResponseCollectionYsH63Sfm, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseCollectionYsH63Sfm struct {
	Errors     []ResponseCollectionYsH63SfmError    `json:"errors"`
	Messages   []ResponseCollectionYsH63SfmMessage  `json:"messages"`
	Result     []ResponseCollectionYsH63SfmResult   `json:"result"`
	ResultInfo ResponseCollectionYsH63SfmResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionYsH63SfmSuccess `json:"success"`
	JSON    responseCollectionYsH63SfmJSON    `json:"-"`
}

// responseCollectionYsH63SfmJSON contains the JSON metadata for the struct
// [ResponseCollectionYsH63Sfm]
type responseCollectionYsH63SfmJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYsH63Sfm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionYsH63SfmError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionYsH63SfmErrorJSON `json:"-"`
}

// responseCollectionYsH63SfmErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionYsH63SfmError]
type responseCollectionYsH63SfmErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYsH63SfmError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionYsH63SfmMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionYsH63SfmMessageJSON `json:"-"`
}

// responseCollectionYsH63SfmMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionYsH63SfmMessage]
type responseCollectionYsH63SfmMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYsH63SfmMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionYsH63SfmResult struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config ResponseCollectionYsH63SfmResultConfig `json:"config"`
	// The name of the Device Managed Network. Must be unique.
	Name string `json:"name"`
	// API uuid tag.
	NetworkID string `json:"network_id"`
	// The type of Device Managed Network.
	Type ResponseCollectionYsH63SfmResultType `json:"type"`
	JSON responseCollectionYsH63SfmResultJSON `json:"-"`
}

// responseCollectionYsH63SfmResultJSON contains the JSON metadata for the struct
// [ResponseCollectionYsH63SfmResult]
type responseCollectionYsH63SfmResultJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYsH63SfmResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ResponseCollectionYsH63SfmResultConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                     `json:"sha256"`
	JSON   responseCollectionYsH63SfmResultConfigJSON `json:"-"`
}

// responseCollectionYsH63SfmResultConfigJSON contains the JSON metadata for the
// struct [ResponseCollectionYsH63SfmResultConfig]
type responseCollectionYsH63SfmResultConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYsH63SfmResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of Device Managed Network.
type ResponseCollectionYsH63SfmResultType string

const (
	ResponseCollectionYsH63SfmResultTypeTls ResponseCollectionYsH63SfmResultType = "tls"
)

type ResponseCollectionYsH63SfmResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionYsH63SfmResultInfoJSON `json:"-"`
}

// responseCollectionYsH63SfmResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionYsH63SfmResultInfo]
type responseCollectionYsH63SfmResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYsH63SfmResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionYsH63SfmSuccess bool

const (
	ResponseCollectionYsH63SfmSuccessTrue ResponseCollectionYsH63SfmSuccess = true
)

type SingleResponse8ngPd7Kr struct {
	Errors   []SingleResponse8ngPd7KrError   `json:"errors"`
	Messages []SingleResponse8ngPd7KrMessage `json:"messages"`
	Result   SingleResponse8ngPd7KrResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponse8ngPd7KrSuccess `json:"success"`
	JSON    singleResponse8ngPd7KrJSON    `json:"-"`
}

// singleResponse8ngPd7KrJSON contains the JSON metadata for the struct
// [SingleResponse8ngPd7Kr]
type singleResponse8ngPd7KrJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponse8ngPd7Kr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponse8ngPd7KrError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponse8ngPd7KrErrorJSON `json:"-"`
}

// singleResponse8ngPd7KrErrorJSON contains the JSON metadata for the struct
// [SingleResponse8ngPd7KrError]
type singleResponse8ngPd7KrErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponse8ngPd7KrError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponse8ngPd7KrMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponse8ngPd7KrMessageJSON `json:"-"`
}

// singleResponse8ngPd7KrMessageJSON contains the JSON metadata for the struct
// [SingleResponse8ngPd7KrMessage]
type singleResponse8ngPd7KrMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponse8ngPd7KrMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponse8ngPd7KrResult struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config SingleResponse8ngPd7KrResultConfig `json:"config"`
	// The name of the Device Managed Network. Must be unique.
	Name string `json:"name"`
	// API uuid tag.
	NetworkID string `json:"network_id"`
	// The type of Device Managed Network.
	Type SingleResponse8ngPd7KrResultType `json:"type"`
	JSON singleResponse8ngPd7KrResultJSON `json:"-"`
}

// singleResponse8ngPd7KrResultJSON contains the JSON metadata for the struct
// [SingleResponse8ngPd7KrResult]
type singleResponse8ngPd7KrResultJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponse8ngPd7KrResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type SingleResponse8ngPd7KrResultConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TlsSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                 `json:"sha256"`
	JSON   singleResponse8ngPd7KrResultConfigJSON `json:"-"`
}

// singleResponse8ngPd7KrResultConfigJSON contains the JSON metadata for the struct
// [SingleResponse8ngPd7KrResultConfig]
type singleResponse8ngPd7KrResultConfigJSON struct {
	TlsSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponse8ngPd7KrResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of Device Managed Network.
type SingleResponse8ngPd7KrResultType string

const (
	SingleResponse8ngPd7KrResultTypeTls SingleResponse8ngPd7KrResultType = "tls"
)

// Whether the API call was successful
type SingleResponse8ngPd7KrSuccess bool

const (
	SingleResponse8ngPd7KrSuccessTrue SingleResponse8ngPd7KrSuccess = true
)

type AccountDeviceNetworkUpdateParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[AccountDeviceNetworkUpdateParamsConfig] `json:"config"`
	// The name of the Device Managed Network. Must be unique.
	Name param.Field[string] `json:"name"`
	// The type of Device Managed Network.
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

// The type of Device Managed Network.
type AccountDeviceNetworkUpdateParamsType string

const (
	AccountDeviceNetworkUpdateParamsTypeTls AccountDeviceNetworkUpdateParamsType = "tls"
)

type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsConfig] `json:"config,required"`
	// The name of the Device Managed Network. Must be unique.
	Name param.Field[string] `json:"name,required"`
	// The type of Device Managed Network.
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

// The type of Device Managed Network.
type AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType string

const (
	AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsTypeTls AccountDeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType = "tls"
)
