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

// Updates a specific interconnect associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicCfInterconnectService) Update(ctx context.Context, accountIdentifier string, tunnelIdentifier string, body MagicCfInterconnectUpdateParams, opts ...option.RequestOption) (res *MagicCfInterconnectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicCfInterconnectUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", accountIdentifier, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// Lists interconnects associated with an account.
func (r *MagicCfInterconnectService) MagicInterconnectsListInterconnects(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *MagicCfInterconnectMagicInterconnectsListInterconnectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates multiple interconnects associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *MagicCfInterconnectService) MagicInterconnectsUpdateMultipleInterconnects(ctx context.Context, accountIdentifier string, body MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsParams, opts ...option.RequestOption) (res *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicCfInterconnectUpdateResponse struct {
	Modified             bool                                  `json:"modified"`
	ModifiedInterconnect interface{}                           `json:"modified_interconnect"`
	JSON                 magicCfInterconnectUpdateResponseJSON `json:"-"`
}

// magicCfInterconnectUpdateResponseJSON contains the JSON metadata for the struct
// [MagicCfInterconnectUpdateResponse]
type magicCfInterconnectUpdateResponseJSON struct {
	Modified             apijson.Field
	ModifiedInterconnect apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MagicCfInterconnectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type MagicCfInterconnectMagicInterconnectsListInterconnectsResponse struct {
	Interconnects []MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnect `json:"interconnects"`
	JSON          magicCfInterconnectMagicInterconnectsListInterconnectsResponseJSON           `json:"-"`
}

// magicCfInterconnectMagicInterconnectsListInterconnectsResponseJSON contains the
// JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsListInterconnectsResponse]
type magicCfInterconnectMagicInterconnectsListInterconnectsResponseJSON struct {
	Interconnects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsListInterconnectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsGre         `json:"gre"`
	HealthCheck MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheck `json:"health_check"`
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
	Name string                                                                         `json:"name"`
	JSON magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnect]
type magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectJSON struct {
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

func (r *MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration specific to GRE interconnects.
type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                                             `json:"cloudflare_endpoint"`
	JSON               magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsGreJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsGreJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsGre]
type magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsGreJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsGre) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckType `json:"type"`
	JSON magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheck]
type magicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRate string

const (
	MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRateLow  MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRate = "low"
	MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRateMid  MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRate = "mid"
	MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRateHigh MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckType string

const (
	MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckTypeReply   MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckType = "reply"
	MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckTypeRequest MagicCfInterconnectMagicInterconnectsListInterconnectsResponseInterconnectsHealthCheckType = "request"
)

type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse struct {
	Modified              bool                                                                                           `json:"modified"`
	ModifiedInterconnects []MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnect `json:"modified_interconnects"`
	JSON                  magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseJSON                   `json:"-"`
}

// magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse]
type magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseJSON struct {
	Modified              apijson.Field
	ModifiedInterconnects apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsGre         `json:"gre"`
	HealthCheck MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheck `json:"health_check"`
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
	Name string                                                                                           `json:"name"`
	JSON magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnect]
type magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectJSON struct {
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

func (r *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration specific to GRE interconnects.
type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                                                               `json:"cloudflare_endpoint"`
	JSON               magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsGreJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsGreJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsGre]
type magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsGreJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsGre) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckType `json:"type"`
	JSON magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheck]
type magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRate string

const (
	MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRateLow  MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRate = "low"
	MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRateMid  MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRate = "mid"
	MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRateHigh MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckType string

const (
	MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckTypeReply   MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckType = "reply"
	MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckTypeRequest MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseModifiedInterconnectsHealthCheckType = "request"
)

type MagicCfInterconnectUpdateParams struct {
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// The configuration specific to GRE interconnects.
	Gre         param.Field[MagicCfInterconnectUpdateParamsGre]         `json:"gre"`
	HealthCheck param.Field[MagicCfInterconnectUpdateParamsHealthCheck] `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu param.Field[int64] `json:"mtu"`
}

func (r MagicCfInterconnectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration specific to GRE interconnects.
type MagicCfInterconnectUpdateParamsGre struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r MagicCfInterconnectUpdateParamsGre) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicCfInterconnectUpdateParamsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicCfInterconnectUpdateParamsHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicCfInterconnectUpdateParamsHealthCheckType] `json:"type"`
}

func (r MagicCfInterconnectUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How frequent the health check is run. The default value is `mid`.
type MagicCfInterconnectUpdateParamsHealthCheckRate string

const (
	MagicCfInterconnectUpdateParamsHealthCheckRateLow  MagicCfInterconnectUpdateParamsHealthCheckRate = "low"
	MagicCfInterconnectUpdateParamsHealthCheckRateMid  MagicCfInterconnectUpdateParamsHealthCheckRate = "mid"
	MagicCfInterconnectUpdateParamsHealthCheckRateHigh MagicCfInterconnectUpdateParamsHealthCheckRate = "high"
)

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicCfInterconnectUpdateParamsHealthCheckType string

const (
	MagicCfInterconnectUpdateParamsHealthCheckTypeReply   MagicCfInterconnectUpdateParamsHealthCheckType = "reply"
	MagicCfInterconnectUpdateParamsHealthCheckTypeRequest MagicCfInterconnectUpdateParamsHealthCheckType = "request"
)

type MagicCfInterconnectUpdateResponseEnvelope struct {
	Errors   []MagicCfInterconnectUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicCfInterconnectUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicCfInterconnectUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicCfInterconnectUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicCfInterconnectUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicCfInterconnectUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [MagicCfInterconnectUpdateResponseEnvelope]
type magicCfInterconnectUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    magicCfInterconnectUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicCfInterconnectUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [MagicCfInterconnectUpdateResponseEnvelopeErrors]
type magicCfInterconnectUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    magicCfInterconnectUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicCfInterconnectUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [MagicCfInterconnectUpdateResponseEnvelopeMessages]
type magicCfInterconnectUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicCfInterconnectUpdateResponseEnvelopeSuccess bool

const (
	MagicCfInterconnectUpdateResponseEnvelopeSuccessTrue MagicCfInterconnectUpdateResponseEnvelopeSuccess = true
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

type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelope struct {
	Errors   []MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicCfInterconnectMagicInterconnectsListInterconnectsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeJSON    `json:"-"`
}

// magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelope]
type magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeErrors struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeErrors]
type magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeMessages struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeMessages]
type magicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeSuccess bool

const (
	MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeSuccessTrue MagicCfInterconnectMagicInterconnectsListInterconnectsResponseEnvelopeSuccess = true
)

type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelope struct {
	Errors   []MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeJSON    `json:"-"`
}

// magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelope]
type magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeErrors struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeErrorsJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeErrors]
type magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeMessages struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeMessagesJSON `json:"-"`
}

// magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeMessages]
type magicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeSuccess bool

const (
	MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeSuccessTrue MagicCfInterconnectMagicInterconnectsUpdateMultipleInterconnectsResponseEnvelopeSuccess = true
)
