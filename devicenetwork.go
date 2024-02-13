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

// DeviceNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceNetworkService] method
// instead.
type DeviceNetworkService struct {
	Options []option.RequestOption
}

// NewDeviceNetworkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceNetworkService(opts ...option.RequestOption) (r *DeviceNetworkService) {
	r = &DeviceNetworkService{}
	r.Options = opts
	return
}

// Updates a configured device managed network.
func (r *DeviceNetworkService) Update(ctx context.Context, identifier interface{}, uuid string, body DeviceNetworkUpdateParams, opts ...option.RequestOption) (res *DeviceNetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceNetworkUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device managed network and fetches a list of the remaining device
// managed networks for an account.
func (r *DeviceNetworkService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *[]DeviceNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceNetworkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new device managed network.
func (r *DeviceNetworkService) DeviceManagedNetworksNewDeviceManagedNetwork(ctx context.Context, identifier interface{}, body DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams, opts ...option.RequestOption) (res *DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of managed networks for an account.
func (r *DeviceNetworkService) DeviceManagedNetworksListDeviceManagedNetworks(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single managed network.
func (r *DeviceNetworkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DeviceNetworkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceNetworkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceNetworkUpdateResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkUpdateResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkUpdateResponseType `json:"type"`
	JSON deviceNetworkUpdateResponseJSON `json:"-"`
}

// deviceNetworkUpdateResponseJSON contains the JSON metadata for the struct
// [DeviceNetworkUpdateResponse]
type deviceNetworkUpdateResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkUpdateResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                `json:"sha256"`
	JSON   deviceNetworkUpdateResponseConfigJSON `json:"-"`
}

// deviceNetworkUpdateResponseConfigJSON contains the JSON metadata for the struct
// [DeviceNetworkUpdateResponseConfig]
type deviceNetworkUpdateResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkUpdateResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type DeviceNetworkUpdateResponseType string

const (
	DeviceNetworkUpdateResponseTypeTLS DeviceNetworkUpdateResponseType = "tls"
)

type DeviceNetworkDeleteResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkDeleteResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkDeleteResponseType `json:"type"`
	JSON deviceNetworkDeleteResponseJSON `json:"-"`
}

// deviceNetworkDeleteResponseJSON contains the JSON metadata for the struct
// [DeviceNetworkDeleteResponse]
type deviceNetworkDeleteResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkDeleteResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                `json:"sha256"`
	JSON   deviceNetworkDeleteResponseConfigJSON `json:"-"`
}

// deviceNetworkDeleteResponseConfigJSON contains the JSON metadata for the struct
// [DeviceNetworkDeleteResponseConfig]
type deviceNetworkDeleteResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeleteResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type DeviceNetworkDeleteResponseType string

const (
	DeviceNetworkDeleteResponseTypeTLS DeviceNetworkDeleteResponseType = "tls"
)

type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseType `json:"type"`
	JSON deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseJSON contains
// the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse]
type deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                                                      `json:"sha256"`
	JSON   deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseConfigJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseConfigJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseConfig]
type deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseType string

const (
	DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseTypeTLS DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseType = "tls"
)

type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseType `json:"type"`
	JSON deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseJSON contains
// the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse]
type deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                                                        `json:"sha256"`
	JSON   deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseConfigJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseConfigJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseConfig]
type deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseType string

const (
	DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseTypeTLS DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseType = "tls"
)

type DeviceNetworkGetResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkGetResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkGetResponseType `json:"type"`
	JSON deviceNetworkGetResponseJSON `json:"-"`
}

// deviceNetworkGetResponseJSON contains the JSON metadata for the struct
// [DeviceNetworkGetResponse]
type deviceNetworkGetResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkGetResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                             `json:"sha256"`
	JSON   deviceNetworkGetResponseConfigJSON `json:"-"`
}

// deviceNetworkGetResponseConfigJSON contains the JSON metadata for the struct
// [DeviceNetworkGetResponseConfig]
type deviceNetworkGetResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type DeviceNetworkGetResponseType string

const (
	DeviceNetworkGetResponseTypeTLS DeviceNetworkGetResponseType = "tls"
)

type DeviceNetworkUpdateParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[DeviceNetworkUpdateParamsConfig] `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name"`
	// The type of device managed network.
	Type param.Field[DeviceNetworkUpdateParamsType] `json:"type"`
}

func (r DeviceNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkUpdateParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r DeviceNetworkUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type DeviceNetworkUpdateParamsType string

const (
	DeviceNetworkUpdateParamsTypeTLS DeviceNetworkUpdateParamsType = "tls"
)

type DeviceNetworkUpdateResponseEnvelope struct {
	Errors   []DeviceNetworkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceNetworkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceNetworkUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceNetworkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceNetworkUpdateResponseEnvelopeJSON    `json:"-"`
}

// deviceNetworkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceNetworkUpdateResponseEnvelope]
type deviceNetworkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceNetworkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceNetworkUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceNetworkUpdateResponseEnvelopeErrors]
type deviceNetworkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceNetworkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceNetworkUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceNetworkUpdateResponseEnvelopeMessages]
type deviceNetworkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceNetworkUpdateResponseEnvelopeSuccess bool

const (
	DeviceNetworkUpdateResponseEnvelopeSuccessTrue DeviceNetworkUpdateResponseEnvelopeSuccess = true
)

type DeviceNetworkDeleteResponseEnvelope struct {
	Errors   []DeviceNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []DeviceNetworkDeleteResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceNetworkDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceNetworkDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceNetworkDeleteResponseEnvelopeJSON       `json:"-"`
}

// deviceNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceNetworkDeleteResponseEnvelope]
type deviceNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceNetworkDeleteResponseEnvelopeErrors]
type deviceNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceNetworkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceNetworkDeleteResponseEnvelopeMessages]
type deviceNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceNetworkDeleteResponseEnvelopeSuccess bool

const (
	DeviceNetworkDeleteResponseEnvelopeSuccessTrue DeviceNetworkDeleteResponseEnvelopeSuccess = true
)

type DeviceNetworkDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       deviceNetworkDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceNetworkDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DeviceNetworkDeleteResponseEnvelopeResultInfo]
type deviceNetworkDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsConfig] `json:"config,required"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name,required"`
	// The type of device managed network.
	Type param.Field[DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType] `json:"type,required"`
}

func (r DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType string

const (
	DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsTypeTLS DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkParamsType = "tls"
)

type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelope struct {
	Errors   []DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeJSON    `json:"-"`
}

// deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelope]
type deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeErrors struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeErrors]
type deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeMessages struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeMessages]
type deviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeSuccess bool

const (
	DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeSuccessTrue DeviceNetworkDeviceManagedNetworksNewDeviceManagedNetworkResponseEnvelopeSuccess = true
)

type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelope struct {
	Errors   []DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeMessages `json:"messages,required"`
	Result   []DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeJSON       `json:"-"`
}

// deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelope]
type deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeErrors]
type deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeMessages]
type deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeSuccess bool

const (
	DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeSuccessTrue DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeSuccess = true
)

type DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                   `json:"total_count"`
	JSON       deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeResultInfo]
type deviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkDeviceManagedNetworksListDeviceManagedNetworksResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkGetResponseEnvelope struct {
	Errors   []DeviceNetworkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceNetworkGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceNetworkGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceNetworkGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceNetworkGetResponseEnvelopeJSON    `json:"-"`
}

// deviceNetworkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceNetworkGetResponseEnvelope]
type deviceNetworkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    deviceNetworkGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceNetworkGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceNetworkGetResponseEnvelopeErrors]
type deviceNetworkGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    deviceNetworkGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceNetworkGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceNetworkGetResponseEnvelopeMessages]
type deviceNetworkGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceNetworkGetResponseEnvelopeSuccess bool

const (
	DeviceNetworkGetResponseEnvelopeSuccessTrue DeviceNetworkGetResponseEnvelopeSuccess = true
)
