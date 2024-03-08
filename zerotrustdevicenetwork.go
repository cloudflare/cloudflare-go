// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustDeviceNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDeviceNetworkService]
// method instead.
type ZeroTrustDeviceNetworkService struct {
	Options []option.RequestOption
}

// NewZeroTrustDeviceNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDeviceNetworkService(opts ...option.RequestOption) (r *ZeroTrustDeviceNetworkService) {
	r = &ZeroTrustDeviceNetworkService{}
	r.Options = opts
	return
}

// Creates a new device managed network.
func (r *ZeroTrustDeviceNetworkService) New(ctx context.Context, params ZeroTrustDeviceNetworkNewParams, opts ...option.RequestOption) (res *ZeroTrustDeviceNetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceNetworkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device managed network.
func (r *ZeroTrustDeviceNetworkService) Update(ctx context.Context, networkID string, params ZeroTrustDeviceNetworkUpdateParams, opts ...option.RequestOption) (res *ZeroTrustDeviceNetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceNetworkUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", params.AccountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of managed networks for an account.
func (r *ZeroTrustDeviceNetworkService) List(ctx context.Context, query ZeroTrustDeviceNetworkListParams, opts ...option.RequestOption) (res *[]ZeroTrustDeviceNetworkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceNetworkListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a device managed network and fetches a list of the remaining device
// managed networks for an account.
func (r *ZeroTrustDeviceNetworkService) Delete(ctx context.Context, networkID string, body ZeroTrustDeviceNetworkDeleteParams, opts ...option.RequestOption) (res *[]ZeroTrustDeviceNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceNetworkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", body.AccountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single managed network.
func (r *ZeroTrustDeviceNetworkService) Get(ctx context.Context, networkID string, query ZeroTrustDeviceNetworkGetParams, opts ...option.RequestOption) (res *ZeroTrustDeviceNetworkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDeviceNetworkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/networks/%s", query.AccountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDeviceNetworkNewResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config ZeroTrustDeviceNetworkNewResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type ZeroTrustDeviceNetworkNewResponseType `json:"type"`
	JSON zeroTrustDeviceNetworkNewResponseJSON `json:"-"`
}

// zeroTrustDeviceNetworkNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceNetworkNewResponse]
type zeroTrustDeviceNetworkNewResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkNewResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ZeroTrustDeviceNetworkNewResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                      `json:"sha256"`
	JSON   zeroTrustDeviceNetworkNewResponseConfigJSON `json:"-"`
}

// zeroTrustDeviceNetworkNewResponseConfigJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceNetworkNewResponseConfig]
type zeroTrustDeviceNetworkNewResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkNewResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkNewResponseConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device managed network.
type ZeroTrustDeviceNetworkNewResponseType string

const (
	ZeroTrustDeviceNetworkNewResponseTypeTLS ZeroTrustDeviceNetworkNewResponseType = "tls"
)

type ZeroTrustDeviceNetworkUpdateResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config ZeroTrustDeviceNetworkUpdateResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type ZeroTrustDeviceNetworkUpdateResponseType `json:"type"`
	JSON zeroTrustDeviceNetworkUpdateResponseJSON `json:"-"`
}

// zeroTrustDeviceNetworkUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceNetworkUpdateResponse]
type zeroTrustDeviceNetworkUpdateResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ZeroTrustDeviceNetworkUpdateResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                         `json:"sha256"`
	JSON   zeroTrustDeviceNetworkUpdateResponseConfigJSON `json:"-"`
}

// zeroTrustDeviceNetworkUpdateResponseConfigJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceNetworkUpdateResponseConfig]
type zeroTrustDeviceNetworkUpdateResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkUpdateResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkUpdateResponseConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device managed network.
type ZeroTrustDeviceNetworkUpdateResponseType string

const (
	ZeroTrustDeviceNetworkUpdateResponseTypeTLS ZeroTrustDeviceNetworkUpdateResponseType = "tls"
)

type ZeroTrustDeviceNetworkListResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config ZeroTrustDeviceNetworkListResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type ZeroTrustDeviceNetworkListResponseType `json:"type"`
	JSON zeroTrustDeviceNetworkListResponseJSON `json:"-"`
}

// zeroTrustDeviceNetworkListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceNetworkListResponse]
type zeroTrustDeviceNetworkListResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkListResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ZeroTrustDeviceNetworkListResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                       `json:"sha256"`
	JSON   zeroTrustDeviceNetworkListResponseConfigJSON `json:"-"`
}

// zeroTrustDeviceNetworkListResponseConfigJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceNetworkListResponseConfig]
type zeroTrustDeviceNetworkListResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkListResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkListResponseConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device managed network.
type ZeroTrustDeviceNetworkListResponseType string

const (
	ZeroTrustDeviceNetworkListResponseTypeTLS ZeroTrustDeviceNetworkListResponseType = "tls"
)

type ZeroTrustDeviceNetworkDeleteResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config ZeroTrustDeviceNetworkDeleteResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type ZeroTrustDeviceNetworkDeleteResponseType `json:"type"`
	JSON zeroTrustDeviceNetworkDeleteResponseJSON `json:"-"`
}

// zeroTrustDeviceNetworkDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceNetworkDeleteResponse]
type zeroTrustDeviceNetworkDeleteResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ZeroTrustDeviceNetworkDeleteResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                         `json:"sha256"`
	JSON   zeroTrustDeviceNetworkDeleteResponseConfigJSON `json:"-"`
}

// zeroTrustDeviceNetworkDeleteResponseConfigJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceNetworkDeleteResponseConfig]
type zeroTrustDeviceNetworkDeleteResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkDeleteResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkDeleteResponseConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device managed network.
type ZeroTrustDeviceNetworkDeleteResponseType string

const (
	ZeroTrustDeviceNetworkDeleteResponseTypeTLS ZeroTrustDeviceNetworkDeleteResponseType = "tls"
)

type ZeroTrustDeviceNetworkGetResponse struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config ZeroTrustDeviceNetworkGetResponseConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type ZeroTrustDeviceNetworkGetResponseType `json:"type"`
	JSON zeroTrustDeviceNetworkGetResponseJSON `json:"-"`
}

// zeroTrustDeviceNetworkGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustDeviceNetworkGetResponse]
type zeroTrustDeviceNetworkGetResponseJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkGetResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ZeroTrustDeviceNetworkGetResponseConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                                      `json:"sha256"`
	JSON   zeroTrustDeviceNetworkGetResponseConfigJSON `json:"-"`
}

// zeroTrustDeviceNetworkGetResponseConfigJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceNetworkGetResponseConfig]
type zeroTrustDeviceNetworkGetResponseConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkGetResponseConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device managed network.
type ZeroTrustDeviceNetworkGetResponseType string

const (
	ZeroTrustDeviceNetworkGetResponseTypeTLS ZeroTrustDeviceNetworkGetResponseType = "tls"
)

type ZeroTrustDeviceNetworkNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[ZeroTrustDeviceNetworkNewParamsConfig] `json:"config,required"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name,required"`
	// The type of device managed network.
	Type param.Field[ZeroTrustDeviceNetworkNewParamsType] `json:"type,required"`
}

func (r ZeroTrustDeviceNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ZeroTrustDeviceNetworkNewParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r ZeroTrustDeviceNetworkNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type ZeroTrustDeviceNetworkNewParamsType string

const (
	ZeroTrustDeviceNetworkNewParamsTypeTLS ZeroTrustDeviceNetworkNewParamsType = "tls"
)

type ZeroTrustDeviceNetworkNewResponseEnvelope struct {
	Errors   []ZeroTrustDeviceNetworkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceNetworkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceNetworkNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceNetworkNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceNetworkNewResponseEnvelope]
type zeroTrustDeviceNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDeviceNetworkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceNetworkNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceNetworkNewResponseEnvelopeErrors]
type zeroTrustDeviceNetworkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDeviceNetworkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceNetworkNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceNetworkNewResponseEnvelopeMessages]
type zeroTrustDeviceNetworkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceNetworkNewResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceNetworkNewResponseEnvelopeSuccessTrue ZeroTrustDeviceNetworkNewResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceNetworkUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[ZeroTrustDeviceNetworkUpdateParamsConfig] `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name"`
	// The type of device managed network.
	Type param.Field[ZeroTrustDeviceNetworkUpdateParamsType] `json:"type"`
}

func (r ZeroTrustDeviceNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type ZeroTrustDeviceNetworkUpdateParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r ZeroTrustDeviceNetworkUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type ZeroTrustDeviceNetworkUpdateParamsType string

const (
	ZeroTrustDeviceNetworkUpdateParamsTypeTLS ZeroTrustDeviceNetworkUpdateParamsType = "tls"
)

type ZeroTrustDeviceNetworkUpdateResponseEnvelope struct {
	Errors   []ZeroTrustDeviceNetworkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceNetworkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceNetworkUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceNetworkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceNetworkUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceNetworkUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceNetworkUpdateResponseEnvelope]
type zeroTrustDeviceNetworkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceNetworkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceNetworkUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceNetworkUpdateResponseEnvelopeErrors]
type zeroTrustDeviceNetworkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDeviceNetworkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceNetworkUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceNetworkUpdateResponseEnvelopeMessages]
type zeroTrustDeviceNetworkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceNetworkUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceNetworkUpdateResponseEnvelopeSuccessTrue ZeroTrustDeviceNetworkUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceNetworkListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceNetworkListResponseEnvelope struct {
	Errors   []ZeroTrustDeviceNetworkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceNetworkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDeviceNetworkListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDeviceNetworkListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDeviceNetworkListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDeviceNetworkListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDeviceNetworkListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceNetworkListResponseEnvelope]
type zeroTrustDeviceNetworkListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkListResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustDeviceNetworkListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceNetworkListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceNetworkListResponseEnvelopeErrors]
type zeroTrustDeviceNetworkListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkListResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceNetworkListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceNetworkListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceNetworkListResponseEnvelopeMessages]
type zeroTrustDeviceNetworkListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceNetworkListResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceNetworkListResponseEnvelopeSuccessTrue ZeroTrustDeviceNetworkListResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceNetworkListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       zeroTrustDeviceNetworkListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDeviceNetworkListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceNetworkListResponseEnvelopeResultInfo]
type zeroTrustDeviceNetworkListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceNetworkDeleteResponseEnvelope struct {
	Errors   []ZeroTrustDeviceNetworkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceNetworkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDeviceNetworkDeleteResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDeviceNetworkDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDeviceNetworkDeleteResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDeviceNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustDeviceNetworkDeleteResponseEnvelope]
type zeroTrustDeviceNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustDeviceNetworkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceNetworkDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceNetworkDeleteResponseEnvelopeErrors]
type zeroTrustDeviceNetworkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustDeviceNetworkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceNetworkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceNetworkDeleteResponseEnvelopeMessages]
type zeroTrustDeviceNetworkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceNetworkDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceNetworkDeleteResponseEnvelopeSuccessTrue ZeroTrustDeviceNetworkDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       zeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfo]
type zeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDeviceNetworkGetResponseEnvelope struct {
	Errors   []ZeroTrustDeviceNetworkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDeviceNetworkGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDeviceNetworkGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZeroTrustDeviceNetworkGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDeviceNetworkGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDeviceNetworkGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDeviceNetworkGetResponseEnvelope]
type zeroTrustDeviceNetworkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDeviceNetworkGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDeviceNetworkGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceNetworkGetResponseEnvelopeErrors]
type zeroTrustDeviceNetworkGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDeviceNetworkGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustDeviceNetworkGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDeviceNetworkGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDeviceNetworkGetResponseEnvelopeMessages]
type zeroTrustDeviceNetworkGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDeviceNetworkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDeviceNetworkGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZeroTrustDeviceNetworkGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDeviceNetworkGetResponseEnvelopeSuccessTrue ZeroTrustDeviceNetworkGetResponseEnvelopeSuccess = true
)
