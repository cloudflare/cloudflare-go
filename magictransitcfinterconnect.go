// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// MagicTransitCfInterconnectService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewMagicTransitCfInterconnectService] method instead.
type MagicTransitCfInterconnectService struct {
	Options []option.RequestOption
}

// NewMagicTransitCfInterconnectService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMagicTransitCfInterconnectService(opts ...option.RequestOption) (r *MagicTransitCfInterconnectService) {
	r = &MagicTransitCfInterconnectService{}
	r.Options = opts
	return
}

// Updates a specific interconnect associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicTransitCfInterconnectService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicTransitCfInterconnectUpdateParams, opts ...option.RequestOption) (res *MagicTransitCfInterconnectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitCfInterconnectUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists interconnects associated with an account.
func (r *MagicTransitCfInterconnectService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicTransitCfInterconnectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitCfInterconnectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific interconnect.
func (r *MagicTransitCfInterconnectService) Get(ctx context.Context, accountIdentifier string, tunnelIdentifier string, opts ...option.RequestOption) (res *MagicTransitCfInterconnectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicTransitCfInterconnectGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicTransitCfInterconnectUpdateResponse struct {
	Modified             bool                                         `json:"modified"`
	ModifiedInterconnect interface{}                                  `json:"modified_interconnect"`
	JSON                 magicTransitCfInterconnectUpdateResponseJSON `json:"-"`
}

// magicTransitCfInterconnectUpdateResponseJSON contains the JSON metadata for the
// struct [MagicTransitCfInterconnectUpdateResponse]
type magicTransitCfInterconnectUpdateResponseJSON struct {
	Modified             apijson.Field
	ModifiedInterconnect apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectListResponse struct {
	Interconnects []MagicTransitCfInterconnectListResponseInterconnect `json:"interconnects"`
	JSON          magicTransitCfInterconnectListResponseJSON           `json:"-"`
}

// magicTransitCfInterconnectListResponseJSON contains the JSON metadata for the
// struct [MagicTransitCfInterconnectListResponse]
type magicTransitCfInterconnectListResponseJSON struct {
	Interconnects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectListResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectListResponseInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         MagicTransitCfInterconnectListResponseInterconnectsGRE         `json:"gre"`
	HealthCheck MagicTransitCfInterconnectListResponseInterconnectsHealthCheck `json:"health_check"`
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
	Name string                                                 `json:"name"`
	JSON magicTransitCfInterconnectListResponseInterconnectJSON `json:"-"`
}

// magicTransitCfInterconnectListResponseInterconnectJSON contains the JSON
// metadata for the struct [MagicTransitCfInterconnectListResponseInterconnect]
type magicTransitCfInterconnectListResponseInterconnectJSON struct {
	ID               apijson.Field
	ColoName         apijson.Field
	CreatedOn        apijson.Field
	Description      apijson.Field
	GRE              apijson.Field
	HealthCheck      apijson.Field
	InterfaceAddress apijson.Field
	ModifiedOn       apijson.Field
	Mtu              apijson.Field
	Name             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectListResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectListResponseInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type MagicTransitCfInterconnectListResponseInterconnectsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                     `json:"cloudflare_endpoint"`
	JSON               magicTransitCfInterconnectListResponseInterconnectsGREJSON `json:"-"`
}

// magicTransitCfInterconnectListResponseInterconnectsGREJSON contains the JSON
// metadata for the struct [MagicTransitCfInterconnectListResponseInterconnectsGRE]
type magicTransitCfInterconnectListResponseInterconnectsGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectListResponseInterconnectsGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectListResponseInterconnectsGREJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectListResponseInterconnectsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicTransitCfInterconnectListResponseInterconnectsHealthCheckType `json:"type"`
	JSON magicTransitCfInterconnectListResponseInterconnectsHealthCheckJSON `json:"-"`
}

// magicTransitCfInterconnectListResponseInterconnectsHealthCheckJSON contains the
// JSON metadata for the struct
// [MagicTransitCfInterconnectListResponseInterconnectsHealthCheck]
type magicTransitCfInterconnectListResponseInterconnectsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectListResponseInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectListResponseInterconnectsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// How frequent the health check is run. The default value is `mid`.
type MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRate string

const (
	MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRateLow  MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRate = "low"
	MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRateMid  MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRate = "mid"
	MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRateHigh MagicTransitCfInterconnectListResponseInterconnectsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitCfInterconnectListResponseInterconnectsHealthCheckType string

const (
	MagicTransitCfInterconnectListResponseInterconnectsHealthCheckTypeReply   MagicTransitCfInterconnectListResponseInterconnectsHealthCheckType = "reply"
	MagicTransitCfInterconnectListResponseInterconnectsHealthCheckTypeRequest MagicTransitCfInterconnectListResponseInterconnectsHealthCheckType = "request"
)

type MagicTransitCfInterconnectGetResponse struct {
	Interconnect interface{}                               `json:"interconnect"`
	JSON         magicTransitCfInterconnectGetResponseJSON `json:"-"`
}

// magicTransitCfInterconnectGetResponseJSON contains the JSON metadata for the
// struct [MagicTransitCfInterconnectGetResponse]
type magicTransitCfInterconnectGetResponseJSON struct {
	Interconnect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectGetResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectUpdateParams struct {
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         param.Field[MagicTransitCfInterconnectUpdateParamsGRE]         `json:"gre"`
	HealthCheck param.Field[MagicTransitCfInterconnectUpdateParamsHealthCheck] `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu param.Field[int64] `json:"mtu"`
}

func (r MagicTransitCfInterconnectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration specific to GRE interconnects.
type MagicTransitCfInterconnectUpdateParamsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r MagicTransitCfInterconnectUpdateParamsGRE) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicTransitCfInterconnectUpdateParamsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicTransitCfInterconnectUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicTransitCfInterconnectUpdateParamsHealthCheckType] `json:"type"`
}

func (r MagicTransitCfInterconnectUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicTransitCfInterconnectUpdateParamsHealthCheckRate string

const (
	MagicTransitCfInterconnectUpdateParamsHealthCheckRateLow  MagicTransitCfInterconnectUpdateParamsHealthCheckRate = "low"
	MagicTransitCfInterconnectUpdateParamsHealthCheckRateMid  MagicTransitCfInterconnectUpdateParamsHealthCheckRate = "mid"
	MagicTransitCfInterconnectUpdateParamsHealthCheckRateHigh MagicTransitCfInterconnectUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTransitCfInterconnectUpdateParamsHealthCheckType string

const (
	MagicTransitCfInterconnectUpdateParamsHealthCheckTypeReply   MagicTransitCfInterconnectUpdateParamsHealthCheckType = "reply"
	MagicTransitCfInterconnectUpdateParamsHealthCheckTypeRequest MagicTransitCfInterconnectUpdateParamsHealthCheckType = "request"
)

type MagicTransitCfInterconnectUpdateResponseEnvelope struct {
	Errors   []MagicTransitCfInterconnectUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitCfInterconnectUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitCfInterconnectUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitCfInterconnectUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitCfInterconnectUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicTransitCfInterconnectUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicTransitCfInterconnectUpdateResponseEnvelope]
type magicTransitCfInterconnectUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectUpdateResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicTransitCfInterconnectUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitCfInterconnectUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicTransitCfInterconnectUpdateResponseEnvelopeErrors]
type magicTransitCfInterconnectUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectUpdateResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    magicTransitCfInterconnectUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitCfInterconnectUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicTransitCfInterconnectUpdateResponseEnvelopeMessages]
type magicTransitCfInterconnectUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitCfInterconnectUpdateResponseEnvelopeSuccess bool

const (
	MagicTransitCfInterconnectUpdateResponseEnvelopeSuccessTrue MagicTransitCfInterconnectUpdateResponseEnvelopeSuccess = true
)

type MagicTransitCfInterconnectListResponseEnvelope struct {
	Errors   []MagicTransitCfInterconnectListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitCfInterconnectListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitCfInterconnectListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitCfInterconnectListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitCfInterconnectListResponseEnvelopeJSON    `json:"-"`
}

// magicTransitCfInterconnectListResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicTransitCfInterconnectListResponseEnvelope]
type magicTransitCfInterconnectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    magicTransitCfInterconnectListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitCfInterconnectListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicTransitCfInterconnectListResponseEnvelopeErrors]
type magicTransitCfInterconnectListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicTransitCfInterconnectListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitCfInterconnectListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitCfInterconnectListResponseEnvelopeMessages]
type magicTransitCfInterconnectListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitCfInterconnectListResponseEnvelopeSuccess bool

const (
	MagicTransitCfInterconnectListResponseEnvelopeSuccessTrue MagicTransitCfInterconnectListResponseEnvelopeSuccess = true
)

type MagicTransitCfInterconnectGetResponseEnvelope struct {
	Errors   []MagicTransitCfInterconnectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicTransitCfInterconnectGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicTransitCfInterconnectGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicTransitCfInterconnectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicTransitCfInterconnectGetResponseEnvelopeJSON    `json:"-"`
}

// magicTransitCfInterconnectGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicTransitCfInterconnectGetResponseEnvelope]
type magicTransitCfInterconnectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicTransitCfInterconnectGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicTransitCfInterconnectGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicTransitCfInterconnectGetResponseEnvelopeErrors]
type magicTransitCfInterconnectGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicTransitCfInterconnectGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    magicTransitCfInterconnectGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicTransitCfInterconnectGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicTransitCfInterconnectGetResponseEnvelopeMessages]
type magicTransitCfInterconnectGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTransitCfInterconnectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTransitCfInterconnectGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicTransitCfInterconnectGetResponseEnvelopeSuccess bool

const (
	MagicTransitCfInterconnectGetResponseEnvelopeSuccessTrue MagicTransitCfInterconnectGetResponseEnvelopeSuccess = true
)
