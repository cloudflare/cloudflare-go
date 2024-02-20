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

// MagicCfInterconnectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicCfInterconnectService]
// method instead.
type MagicCfInterconnectService struct {
	Options []option.RequestOption
}

// NewMagicCfInterconnectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMagicCfInterconnectService(opts ...option.RequestOption) (r *MagicCfInterconnectService) {
	r = &MagicCfInterconnectService{}
	r.Options = opts
	return
}

// Lists interconnects associated with an account.
func (r *MagicCfInterconnectService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicCfInterconnectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicCfInterconnectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific interconnect.
func (r *MagicCfInterconnectService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicCfInterconnectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicCfInterconnectGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specific interconnect associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicCfInterconnectService) Replace(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicCfInterconnectReplaceParams, opts ...option.RequestOption) (res *MagicCfInterconnectReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicCfInterconnectReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicCfInterconnectListResponse struct {
	Interconnects []MagicCfInterconnectListResponseInterconnect `json:"interconnects"`
	JSON          magicCfInterconnectListResponseJSON           `json:"-"`
}

// magicCfInterconnectListResponseJSON contains the JSON metadata for the struct
// [MagicCfInterconnectListResponse]
type magicCfInterconnectListResponseJSON struct {
	Interconnects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicCfInterconnectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectListResponseInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         MagicCfInterconnectListResponseInterconnectsGre         `json:"gre"`
	HealthCheck MagicCfInterconnectListResponseInterconnectsHealthCheck `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                                          `json:"name"`
	JSON magicCfInterconnectListResponseInterconnectJSON `json:"-"`
}

// magicCfInterconnectListResponseInterconnectJSON contains the JSON metadata for
// the struct [MagicCfInterconnectListResponseInterconnect]
type magicCfInterconnectListResponseInterconnectJSON struct {
	ID               apijson.Field
	ColoName         apijson.Field
	CreatedOn        apijson.Field
	Description      apijson.Field
	Gre              apijson.Field
	HealthCheck      apijson.Field
	InterfaceAddress apijson.Field
	ModifiedOn       apijson.Field
	Mtu              apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicCfInterconnectListResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration specific to GRE interconnects.
type MagicCfInterconnectListResponseInterconnectsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                              `json:"cloudflare_endpoint"`
	JSON               magicCfInterconnectListResponseInterconnectsGreJSON `json:"-"`
}

// magicCfInterconnectListResponseInterconnectsGreJSON contains the JSON metadata
// for the struct [MagicCfInterconnectListResponseInterconnectsGre]
type magicCfInterconnectListResponseInterconnectsGreJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicCfInterconnectListResponseInterconnectsGre) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectListResponseInterconnectsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicCfInterconnectListResponseInterconnectsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicCfInterconnectListResponseInterconnectsHealthCheckType `json:"type"`
	JSON magicCfInterconnectListResponseInterconnectsHealthCheckJSON `json:"-"`
}

// magicCfInterconnectListResponseInterconnectsHealthCheckJSON contains the JSON
// metadata for the struct
// [MagicCfInterconnectListResponseInterconnectsHealthCheck]
type magicCfInterconnectListResponseInterconnectsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectListResponseInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicCfInterconnectListResponseInterconnectsHealthCheckRate string

const (
	MagicCfInterconnectListResponseInterconnectsHealthCheckRateLow  MagicCfInterconnectListResponseInterconnectsHealthCheckRate = "low"
	MagicCfInterconnectListResponseInterconnectsHealthCheckRateMid  MagicCfInterconnectListResponseInterconnectsHealthCheckRate = "mid"
	MagicCfInterconnectListResponseInterconnectsHealthCheckRateHigh MagicCfInterconnectListResponseInterconnectsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicCfInterconnectListResponseInterconnectsHealthCheckType string

const (
	MagicCfInterconnectListResponseInterconnectsHealthCheckTypeReply   MagicCfInterconnectListResponseInterconnectsHealthCheckType = "reply"
	MagicCfInterconnectListResponseInterconnectsHealthCheckTypeRequest MagicCfInterconnectListResponseInterconnectsHealthCheckType = "request"
)

type MagicCfInterconnectGetResponse struct {
	Interconnect interface{}                        `json:"interconnect"`
	JSON         magicCfInterconnectGetResponseJSON `json:"-"`
}

// magicCfInterconnectGetResponseJSON contains the JSON metadata for the struct
// [MagicCfInterconnectGetResponse]
type magicCfInterconnectGetResponseJSON struct {
	Interconnect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicCfInterconnectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectReplaceResponse struct {
	Modified             bool                                   `json:"modified"`
	ModifiedInterconnect interface{}                            `json:"modified_interconnect"`
	JSON                 magicCfInterconnectReplaceResponseJSON `json:"-"`
}

// magicCfInterconnectReplaceResponseJSON contains the JSON metadata for the struct
// [MagicCfInterconnectReplaceResponse]
type magicCfInterconnectReplaceResponseJSON struct {
	Modified             apijson.Field
	ModifiedInterconnect apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MagicCfInterconnectReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectListResponseEnvelope struct {
	Errors   []MagicCfInterconnectListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicCfInterconnectListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicCfInterconnectListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicCfInterconnectListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicCfInterconnectListResponseEnvelopeJSON    `json:"-"`
}

// magicCfInterconnectListResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicCfInterconnectListResponseEnvelope]
type magicCfInterconnectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    magicCfInterconnectListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicCfInterconnectListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicCfInterconnectListResponseEnvelopeErrors]
type magicCfInterconnectListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    magicCfInterconnectListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicCfInterconnectListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicCfInterconnectListResponseEnvelopeMessages]
type magicCfInterconnectListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicCfInterconnectListResponseEnvelopeSuccess bool

const (
	MagicCfInterconnectListResponseEnvelopeSuccessTrue MagicCfInterconnectListResponseEnvelopeSuccess = true
)

type MagicCfInterconnectGetResponseEnvelope struct {
	Errors   []MagicCfInterconnectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicCfInterconnectGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicCfInterconnectGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicCfInterconnectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicCfInterconnectGetResponseEnvelopeJSON    `json:"-"`
}

// magicCfInterconnectGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicCfInterconnectGetResponseEnvelope]
type magicCfInterconnectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    magicCfInterconnectGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicCfInterconnectGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [MagicCfInterconnectGetResponseEnvelopeErrors]
type magicCfInterconnectGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    magicCfInterconnectGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicCfInterconnectGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicCfInterconnectGetResponseEnvelopeMessages]
type magicCfInterconnectGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicCfInterconnectGetResponseEnvelopeSuccess bool

const (
	MagicCfInterconnectGetResponseEnvelopeSuccessTrue MagicCfInterconnectGetResponseEnvelopeSuccess = true
)

type MagicCfInterconnectReplaceParams struct {
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         param.Field[MagicCfInterconnectReplaceParamsGre]         `json:"gre"`
	HealthCheck param.Field[MagicCfInterconnectReplaceParamsHealthCheck] `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu param.Field[int64] `json:"mtu"`
}

func (r MagicCfInterconnectReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration specific to GRE interconnects.
type MagicCfInterconnectReplaceParamsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r MagicCfInterconnectReplaceParamsGre) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicCfInterconnectReplaceParamsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicCfInterconnectReplaceParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicCfInterconnectReplaceParamsHealthCheckType] `json:"type"`
}

func (r MagicCfInterconnectReplaceParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicCfInterconnectReplaceParamsHealthCheckRate string

const (
	MagicCfInterconnectReplaceParamsHealthCheckRateLow  MagicCfInterconnectReplaceParamsHealthCheckRate = "low"
	MagicCfInterconnectReplaceParamsHealthCheckRateMid  MagicCfInterconnectReplaceParamsHealthCheckRate = "mid"
	MagicCfInterconnectReplaceParamsHealthCheckRateHigh MagicCfInterconnectReplaceParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicCfInterconnectReplaceParamsHealthCheckType string

const (
	MagicCfInterconnectReplaceParamsHealthCheckTypeReply   MagicCfInterconnectReplaceParamsHealthCheckType = "reply"
	MagicCfInterconnectReplaceParamsHealthCheckTypeRequest MagicCfInterconnectReplaceParamsHealthCheckType = "request"
)

type MagicCfInterconnectReplaceResponseEnvelope struct {
	Errors   []MagicCfInterconnectReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicCfInterconnectReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicCfInterconnectReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicCfInterconnectReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicCfInterconnectReplaceResponseEnvelopeJSON    `json:"-"`
}

// magicCfInterconnectReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicCfInterconnectReplaceResponseEnvelope]
type magicCfInterconnectReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectReplaceResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    magicCfInterconnectReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// magicCfInterconnectReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicCfInterconnectReplaceResponseEnvelopeErrors]
type magicCfInterconnectReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectReplaceResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    magicCfInterconnectReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// magicCfInterconnectReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicCfInterconnectReplaceResponseEnvelopeMessages]
type magicCfInterconnectReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicCfInterconnectReplaceResponseEnvelopeSuccess bool

const (
	MagicCfInterconnectReplaceResponseEnvelopeSuccessTrue MagicCfInterconnectReplaceResponseEnvelopeSuccess = true
)
