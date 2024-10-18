// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/tidwall/gjson"
)

// CfInterconnectService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCfInterconnectService] method instead.
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
func (r *CfInterconnectService) Update(ctx context.Context, cfInterconnectID string, params CfInterconnectUpdateParams, opts ...option.RequestOption) (res *CfInterconnectUpdateResponse, err error) {
	var env CfInterconnectUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", params.AccountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists interconnects associated with an account.
func (r *CfInterconnectService) List(ctx context.Context, params CfInterconnectListParams, opts ...option.RequestOption) (res *CfInterconnectListResponse, err error) {
	var env CfInterconnectListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific interconnect.
func (r *CfInterconnectService) Get(ctx context.Context, cfInterconnectID string, params CfInterconnectGetParams, opts ...option.RequestOption) (res *CfInterconnectGetResponse, err error) {
	var env CfInterconnectGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", params.AccountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CfInterconnectUpdateResponse struct {
	Modified             bool                                             `json:"modified"`
	ModifiedInterconnect CfInterconnectUpdateResponseModifiedInterconnect `json:"modified_interconnect"`
	JSON                 cfInterconnectUpdateResponseJSON                 `json:"-"`
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

type CfInterconnectUpdateResponseModifiedInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         CfInterconnectUpdateResponseModifiedInterconnectGRE         `json:"gre"`
	HealthCheck CfInterconnectUpdateResponseModifiedInterconnectHealthCheck `json:"health_check"`
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
	Name string                                               `json:"name"`
	JSON cfInterconnectUpdateResponseModifiedInterconnectJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectJSON contains the JSON metadata
// for the struct [CfInterconnectUpdateResponseModifiedInterconnect]
type cfInterconnectUpdateResponseModifiedInterconnectJSON struct {
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

func (r *CfInterconnectUpdateResponseModifiedInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type CfInterconnectUpdateResponseModifiedInterconnectGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                  `json:"cloudflare_endpoint"`
	JSON               cfInterconnectUpdateResponseModifiedInterconnectGREJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectGREJSON contains the JSON
// metadata for the struct [CfInterconnectUpdateResponseModifiedInterconnectGRE]
type cfInterconnectUpdateResponseModifiedInterconnectGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseModifiedInterconnectGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectUpdateResponseModifiedInterconnectHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type CfInterconnectUpdateResponseModifiedInterconnectHealthCheckType `json:"type"`
	JSON cfInterconnectUpdateResponseModifiedInterconnectHealthCheckJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectHealthCheckJSON contains the
// JSON metadata for the struct
// [CfInterconnectUpdateResponseModifiedInterconnectHealthCheck]
type cfInterconnectUpdateResponseModifiedInterconnectHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseModifiedInterconnectHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion interface {
	ImplementsMagicTransitCfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                                      `json:"saved"`
	JSON  cfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget]
type cfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitCfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion() {
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type CfInterconnectUpdateResponseModifiedInterconnectHealthCheckType string

const (
	CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTypeReply   CfInterconnectUpdateResponseModifiedInterconnectHealthCheckType = "reply"
	CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTypeRequest CfInterconnectUpdateResponseModifiedInterconnectHealthCheckType = "request"
)

func (r CfInterconnectUpdateResponseModifiedInterconnectHealthCheckType) IsKnown() bool {
	switch r {
	case CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTypeReply, CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTypeRequest:
		return true
	}
	return false
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
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target CfInterconnectListResponseInterconnectsHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type CfInterconnectListResponseInterconnectsHealthCheckType `json:"type"`
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

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type CfInterconnectListResponseInterconnectsHealthCheckTargetUnion interface {
	ImplementsMagicTransitCfInterconnectListResponseInterconnectsHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfInterconnectListResponseInterconnectsHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                             `json:"saved"`
	JSON  cfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget]
type cfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitCfInterconnectListResponseInterconnectsHealthCheckTargetUnion() {
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type CfInterconnectListResponseInterconnectsHealthCheckType string

const (
	CfInterconnectListResponseInterconnectsHealthCheckTypeReply   CfInterconnectListResponseInterconnectsHealthCheckType = "reply"
	CfInterconnectListResponseInterconnectsHealthCheckTypeRequest CfInterconnectListResponseInterconnectsHealthCheckType = "request"
)

func (r CfInterconnectListResponseInterconnectsHealthCheckType) IsKnown() bool {
	switch r {
	case CfInterconnectListResponseInterconnectsHealthCheckTypeReply, CfInterconnectListResponseInterconnectsHealthCheckTypeRequest:
		return true
	}
	return false
}

type CfInterconnectGetResponse struct {
	Interconnect CfInterconnectGetResponseInterconnect `json:"interconnect"`
	JSON         cfInterconnectGetResponseJSON         `json:"-"`
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

type CfInterconnectGetResponseInterconnect struct {
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         CfInterconnectGetResponseInterconnectGRE         `json:"gre"`
	HealthCheck CfInterconnectGetResponseInterconnectHealthCheck `json:"health_check"`
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
	Name string                                    `json:"name"`
	JSON cfInterconnectGetResponseInterconnectJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectJSON contains the JSON metadata for the
// struct [CfInterconnectGetResponseInterconnect]
type cfInterconnectGetResponseInterconnectJSON struct {
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

func (r *CfInterconnectGetResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type CfInterconnectGetResponseInterconnectGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                       `json:"cloudflare_endpoint"`
	JSON               cfInterconnectGetResponseInterconnectGREJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectGREJSON contains the JSON metadata for the
// struct [CfInterconnectGetResponseInterconnectGRE]
type cfInterconnectGetResponseInterconnectGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectGetResponseInterconnectGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectGetResponseInterconnectHealthCheck struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target CfInterconnectGetResponseInterconnectHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type CfInterconnectGetResponseInterconnectHealthCheckType `json:"type"`
	JSON cfInterconnectGetResponseInterconnectHealthCheckJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectHealthCheckJSON contains the JSON metadata
// for the struct [CfInterconnectGetResponseInterconnectHealthCheck]
type cfInterconnectGetResponseInterconnectHealthCheckJSON struct {
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectGetResponseInterconnectHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [magic_transit.CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type CfInterconnectGetResponseInterconnectHealthCheckTargetUnion interface {
	ImplementsMagicTransitCfInterconnectGetResponseInterconnectHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfInterconnectGetResponseInterconnectHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                           `json:"saved"`
	JSON  cfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget]
type cfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitCfInterconnectGetResponseInterconnectHealthCheckTargetUnion() {
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type CfInterconnectGetResponseInterconnectHealthCheckType string

const (
	CfInterconnectGetResponseInterconnectHealthCheckTypeReply   CfInterconnectGetResponseInterconnectHealthCheckType = "reply"
	CfInterconnectGetResponseInterconnectHealthCheckTypeRequest CfInterconnectGetResponseInterconnectHealthCheckType = "request"
)

func (r CfInterconnectGetResponseInterconnectHealthCheckType) IsKnown() bool {
	switch r {
	case CfInterconnectGetResponseInterconnectHealthCheckTypeReply, CfInterconnectGetResponseInterconnectHealthCheckTypeRequest:
		return true
	}
	return false
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
	Mtu               param.Field[int64] `json:"mtu"`
	XMagicNewHcTarget param.Field[bool]  `header:"x-magic-new-hc-target"`
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
	Rate param.Field[HealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[CfInterconnectUpdateParamsHealthCheckTargetUnion] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[CfInterconnectUpdateParamsHealthCheckType] `json:"type"`
}

func (r CfInterconnectUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Satisfied by
// [magic_transit.CfInterconnectUpdateParamsHealthCheckTargetMagicHealthCheckTarget],
// [shared.UnionString].
type CfInterconnectUpdateParamsHealthCheckTargetUnion interface {
	ImplementsMagicTransitCfInterconnectUpdateParamsHealthCheckTargetUnion()
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type CfInterconnectUpdateParamsHealthCheckTargetMagicHealthCheckTarget struct {
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved param.Field[string] `json:"saved"`
}

func (r CfInterconnectUpdateParamsHealthCheckTargetMagicHealthCheckTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CfInterconnectUpdateParamsHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTransitCfInterconnectUpdateParamsHealthCheckTargetUnion() {
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type CfInterconnectUpdateParamsHealthCheckType string

const (
	CfInterconnectUpdateParamsHealthCheckTypeReply   CfInterconnectUpdateParamsHealthCheckType = "reply"
	CfInterconnectUpdateParamsHealthCheckTypeRequest CfInterconnectUpdateParamsHealthCheckType = "request"
)

func (r CfInterconnectUpdateParamsHealthCheckType) IsKnown() bool {
	switch r {
	case CfInterconnectUpdateParamsHealthCheckTypeReply, CfInterconnectUpdateParamsHealthCheckTypeRequest:
		return true
	}
	return false
}

type CfInterconnectUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
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
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type CfInterconnectListResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   CfInterconnectListResponse `json:"result,required"`
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
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type CfInterconnectGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   CfInterconnectGetResponse `json:"result,required"`
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
