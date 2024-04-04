// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// CfInterconnectService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCfInterconnectService] method
// instead.
type CfInterconnectService struct {
	Options []option.RequestOption
}

// NewCfInterconnectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCfInterconnectService(opts ...option.RequestOption) (r *CfInterconnectService) {
	r = &CfInterconnectService{}
	r.Options = opts
	return
}

// Updates a specific interconnect associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *CfInterconnectService) Update(ctx context.Context, tunnelIdentifier string, params CfInterconnectUpdateParams, opts ...option.RequestOption) (res *CfInterconnectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfInterconnectUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", params.AccountID, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists interconnects associated with an account.
func (r *CfInterconnectService) List(ctx context.Context, query CfInterconnectListParams, opts ...option.RequestOption) (res *CfInterconnectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfInterconnectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific interconnect.
func (r *CfInterconnectService) Get(ctx context.Context, tunnelIdentifier string, query CfInterconnectGetParams, opts ...option.RequestOption) (res *CfInterconnectGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CfInterconnectGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", query.AccountID, tunnelIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CfInterconnectUpdateResponse struct {
	Modified             bool                             `json:"modified"`
	ModifiedInterconnect interface{}                      `json:"modified_interconnect"`
	JSON                 cfInterconnectUpdateResponseJSON `json:"-"`
}

// cfInterconnectUpdateResponseJSON contains the JSON metadata for the struct
// [CfInterconnectUpdateResponse]
type cfInterconnectUpdateResponseJSON struct {
	Modified             apijson.Field
	ModifiedInterconnect apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectListResponse struct {
	Interconnects []CfInterconnectListResponseInterconnect `json:"interconnects"`
	JSON          cfInterconnectListResponseJSON           `json:"-"`
}

// cfInterconnectListResponseJSON contains the JSON metadata for the struct
// [CfInterconnectListResponse]
type cfInterconnectListResponseJSON struct {
	Interconnects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CfInterconnectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectListResponseInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         CfInterconnectListResponseInterconnectsGRE         `json:"gre"`
	HealthCheck CfInterconnectListResponseInterconnectsHealthCheck `json:"health_check"`
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
	Name string                                     `json:"name"`
	JSON cfInterconnectListResponseInterconnectJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectJSON contains the JSON metadata for the
// struct [CfInterconnectListResponseInterconnect]
type cfInterconnectListResponseInterconnectJSON struct {
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

func (r *CfInterconnectListResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type CfInterconnectListResponseInterconnectsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                         `json:"cloudflare_endpoint"`
	JSON               cfInterconnectListResponseInterconnectsGREJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectsGREJSON contains the JSON metadata for
// the struct [CfInterconnectListResponseInterconnectsGRE]
type cfInterconnectListResponseInterconnectsGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectListResponseInterconnectsGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectsGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectListResponseInterconnectsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate shared.UnnamedSchemaRef83 `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type shared.UnnamedSchemaRef84                              `json:"type"`
	JSON cfInterconnectListResponseInterconnectsHealthCheckJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectsHealthCheckJSON contains the JSON
// metadata for the struct [CfInterconnectListResponseInterconnectsHealthCheck]
type cfInterconnectListResponseInterconnectsHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectListResponseInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectsHealthCheckJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectGetResponse struct {
	Interconnect interface{}                   `json:"interconnect"`
	JSON         cfInterconnectGetResponseJSON `json:"-"`
}

// cfInterconnectGetResponseJSON contains the JSON metadata for the struct
// [CfInterconnectGetResponse]
type cfInterconnectGetResponseJSON struct {
	Interconnect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CfInterconnectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         param.Field[CfInterconnectUpdateParamsGRE]         `json:"gre"`
	HealthCheck param.Field[CfInterconnectUpdateParamsHealthCheck] `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu param.Field[int64] `json:"mtu"`
}

func (r CfInterconnectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration specific to GRE interconnects.
type CfInterconnectUpdateParamsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r CfInterconnectUpdateParamsGRE) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CfInterconnectUpdateParamsHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[shared.UnnamedSchemaRef83] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[shared.UnnamedSchemaRef84] `json:"type"`
}

func (r CfInterconnectUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CfInterconnectUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   CfInterconnectUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfInterconnectUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfInterconnectUpdateResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CfInterconnectUpdateResponseEnvelope]
type cfInterconnectUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectUpdateResponseEnvelopeSuccess bool

const (
	CfInterconnectUpdateResponseEnvelopeSuccessTrue CfInterconnectUpdateResponseEnvelopeSuccess = true
)

func (r CfInterconnectUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CfInterconnectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type CfInterconnectListResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   CfInterconnectListResponse   `json:"result,required"`
	// Whether the API call was successful
	Success CfInterconnectListResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfInterconnectListResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CfInterconnectListResponseEnvelope]
type cfInterconnectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectListResponseEnvelopeSuccess bool

const (
	CfInterconnectListResponseEnvelopeSuccessTrue CfInterconnectListResponseEnvelopeSuccess = true
)

func (r CfInterconnectListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CfInterconnectGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type CfInterconnectGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   CfInterconnectGetResponse    `json:"result,required"`
	// Whether the API call was successful
	Success CfInterconnectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfInterconnectGetResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CfInterconnectGetResponseEnvelope]
type cfInterconnectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectGetResponseEnvelopeSuccess bool

const (
	CfInterconnectGetResponseEnvelopeSuccessTrue CfInterconnectGetResponseEnvelopeSuccess = true
)

func (r CfInterconnectGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
