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

// Creates a new device managed network.
func (r *DeviceNetworkService) New(ctx context.Context, identifier interface{}, body DeviceNetworkNewParams, opts ...option.RequestOption) (res *DeviceNetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceNetworkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// Fetches a list of managed networks for an account.
func (r *DeviceNetworkService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]DeviceNetworkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceNetworkListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

type DeviceNetworkNewResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkNewResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkNewResponseType `json:"type"`
	JSON deviceNetworkNewResponseJSON `json:"-"`
}

// deviceNetworkNewResponseJSON contains the JSON metadata for the struct
// [DeviceNetworkNewResponse]
type deviceNetworkNewResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkNewResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                             `json:"sha256"`
	JSON   deviceNetworkNewResponseConfigJSON `json:"-"`
}

// deviceNetworkNewResponseConfigJSON contains the JSON metadata for the struct
// [DeviceNetworkNewResponseConfig]
type deviceNetworkNewResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkNewResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type DeviceNetworkNewResponseType string

const (
	DeviceNetworkNewResponseTypeTLS DeviceNetworkNewResponseType = "tls"
)

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

type DeviceNetworkListResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkListResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkListResponseType `json:"type"`
	JSON deviceNetworkListResponseJSON `json:"-"`
}

// deviceNetworkListResponseJSON contains the JSON metadata for the struct
// [DeviceNetworkListResponse]
type deviceNetworkListResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkListResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                              `json:"sha256"`
	JSON   deviceNetworkListResponseConfigJSON `json:"-"`
}

// deviceNetworkListResponseConfigJSON contains the JSON metadata for the struct
// [DeviceNetworkListResponseConfig]
type deviceNetworkListResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkListResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of device managed network.
type DeviceNetworkListResponseType string

const (
	DeviceNetworkListResponseTypeTLS DeviceNetworkListResponseType = "tls"
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

type DeviceNetworkNewParams struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[DeviceNetworkNewParamsConfig] `json:"config,required"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name,required"`
	// The type of device managed network.
	Type param.Field[DeviceNetworkNewParamsType] `json:"type,required"`
}

func (r DeviceNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkNewParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r DeviceNetworkNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type DeviceNetworkNewParamsType string

const (
	DeviceNetworkNewParamsTypeTLS DeviceNetworkNewParamsType = "tls"
)

type DeviceNetworkNewResponseEnvelope struct {
	Errors   []DeviceNetworkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceNetworkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceNetworkNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// deviceNetworkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceNetworkNewResponseEnvelope]
type deviceNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    deviceNetworkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceNetworkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceNetworkNewResponseEnvelopeErrors]
type deviceNetworkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    deviceNetworkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceNetworkNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceNetworkNewResponseEnvelopeMessages]
type deviceNetworkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceNetworkNewResponseEnvelopeSuccess bool

const (
	DeviceNetworkNewResponseEnvelopeSuccessTrue DeviceNetworkNewResponseEnvelopeSuccess = true
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

type DeviceNetworkListResponseEnvelope struct {
	Errors   []DeviceNetworkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceNetworkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DeviceNetworkListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceNetworkListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceNetworkListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceNetworkListResponseEnvelopeJSON       `json:"-"`
}

// deviceNetworkListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceNetworkListResponseEnvelope]
type deviceNetworkListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    deviceNetworkListResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceNetworkListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceNetworkListResponseEnvelopeErrors]
type deviceNetworkListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceNetworkListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceNetworkListResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceNetworkListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceNetworkListResponseEnvelopeMessages]
type deviceNetworkListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceNetworkListResponseEnvelopeSuccess bool

const (
	DeviceNetworkListResponseEnvelopeSuccessTrue DeviceNetworkListResponseEnvelopeSuccess = true
)

type DeviceNetworkListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       deviceNetworkListResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceNetworkListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DeviceNetworkListResponseEnvelopeResultInfo]
type deviceNetworkListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
