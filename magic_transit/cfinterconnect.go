// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
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
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%v", params.XMagicNewHcTarget)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", params.AccountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists interconnects associated with an account.
func (r *CfInterconnectService) List(ctx context.Context, params CfInterconnectListParams, opts ...option.RequestOption) (res *CfInterconnectListResponse, err error) {
	var env CfInterconnectListResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%v", params.XMagicNewHcTarget)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates multiple interconnects associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *CfInterconnectService) BulkUpdate(ctx context.Context, params CfInterconnectBulkUpdateParams, opts ...option.RequestOption) (res *CfInterconnectBulkUpdateResponse, err error) {
	var env CfInterconnectBulkUpdateResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%v", params.XMagicNewHcTarget)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists details for a specific interconnect.
func (r *CfInterconnectService) Get(ctx context.Context, cfInterconnectID string, params CfInterconnectGetParams, opts ...option.RequestOption) (res *CfInterconnectGetResponse, err error) {
	var env CfInterconnectGetResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%v", params.XMagicNewHcTarget)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", params.AccountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
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
	// Identifier
	ID string `json:"id"`
	// True if automatic stateful return routing should be enabled for a tunnel, false
	// otherwise. Requires the `coupler_integration` account flag to be enabled;
	// requests setting this to `true` without that flag will be rejected.
	AutomaticReturnRouting bool                                                `json:"automatic_return_routing"`
	BGP                    CfInterconnectUpdateResponseModifiedInterconnectBGP `json:"bgp"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// Omitted in responses for version 1.5 interconnects.
	GRE         CfInterconnectUpdateResponseModifiedInterconnectGRE         `json:"gre"`
	HealthCheck CfInterconnectUpdateResponseModifiedInterconnectHealthCheck `json:"health_check"`
	// The IPv4 interface address for the interconnect. For MPLS Interconnects, use a
	// /30 or /31 prefix. For GRE Interconnects, a /30 or /31 prefix may be used.
	// Version 1.5 interconnects require a /31 prefix and may also use a prefix from
	// the account's authorized prefixes; otherwise, select the subnet from RFC 1918 or
	// the approved link-local ranges.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string `json:"name"`
	// Immutable interconnect version configured at creation time. One of:
	//
	// - "1"
	// - "1.5"
	// - "2"
	Version string `json:"version"`
	// An identifier that correlates this interconnect with the corresponding V2 CNI
	// interconnect resource.
	VirtualPortReservationID string                                               `json:"virtual_port_reservation_id"`
	JSON                     cfInterconnectUpdateResponseModifiedInterconnectJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectJSON contains the JSON metadata
// for the struct [CfInterconnectUpdateResponseModifiedInterconnect]
type cfInterconnectUpdateResponseModifiedInterconnectJSON struct {
	ID                       apijson.Field
	AutomaticReturnRouting   apijson.Field
	BGP                      apijson.Field
	ColoName                 apijson.Field
	CreatedOn                apijson.Field
	Description              apijson.Field
	GRE                      apijson.Field
	HealthCheck              apijson.Field
	InterfaceAddress         apijson.Field
	InterfaceAddress6        apijson.Field
	ModifiedOn               apijson.Field
	Mtu                      apijson.Field
	Name                     apijson.Field
	Version                  apijson.Field
	VirtualPortReservationID apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseModifiedInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectUpdateResponseModifiedInterconnectBGP struct {
	// Deprecated. Use customer_asn.
	//
	// Deprecated: deprecated
	AsNo int64 `json:"as_no"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CloudflareEndpoint string `json:"cloudflare_endpoint" format:"ipv4"`
	// ASN used on the customer end of the BGP session.
	CustomerASN int64 `json:"customer_asn"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CustomerEndpoint string `json:"customer_endpoint" format:"ipv4"`
	// ID of the BGP filter profile applied to routes advertised to the customer.
	ExportFilterID string `json:"export_filter_id"`
	// Prefixes in this list will be advertised to the customer device, in addition to
	// the routes in the Magic routing table.
	ExtraPrefixes []string `json:"extra_prefixes" format:"cidr"`
	// ID of the BGP filter profile applied to routes received from the customer.
	ImportFilterID string `json:"import_filter_id"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                                                  `json:"md5_key"`
	JSON   cfInterconnectUpdateResponseModifiedInterconnectBGPJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectBGPJSON contains the JSON
// metadata for the struct [CfInterconnectUpdateResponseModifiedInterconnectBGP]
type cfInterconnectUpdateResponseModifiedInterconnectBGPJSON struct {
	AsNo               apijson.Field
	CloudflareEndpoint apijson.Field
	CustomerASN        apijson.Field
	CustomerEndpoint   apijson.Field
	ExportFilterID     apijson.Field
	ExtraPrefixes      apijson.Field
	ImportFilterID     apijson.Field
	Md5Key             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseModifiedInterconnectBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectBGPJSON) RawJSON() string {
	return r.raw
}

// Omitted in responses for version 1.5 interconnects.
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
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the interconnect and the result comes back to Cloudflare
	// via the open Internet, or bidirectional where both the probe and result come and
	// go via the interconnect.
	Direction CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The source IPv4 address used for bidirectional health checks. Supported only for
	// version 1.5 interconnects. It is required when `direction` is `bidirectional`
	// and must be omitted (and is cleared) when `direction` is `unidirectional`. The
	// address must be within RFC1918 space, the approved link-local range
	// 169.254.240.0/20, or the Cloudflare reserved range 198.41.199.224/27.
	Source string `json:"source"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                                 `json:"type"`
	JSON cfInterconnectUpdateResponseModifiedInterconnectHealthCheckJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectHealthCheckJSON contains the
// JSON metadata for the struct
// [CfInterconnectUpdateResponseModifiedInterconnectHealthCheck]
type cfInterconnectUpdateResponseModifiedInterconnectHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Source      apijson.Field
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

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the interconnect and the result comes back to Cloudflare
// via the open Internet, or bidirectional where both the probe and result come and
// go via the interconnect.
type CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirection string

const (
	CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirectionUnidirectional CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirection = "unidirectional"
	CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirectionBidirectional  CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirection = "bidirectional"
)

func (r CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirection) IsKnown() bool {
	switch r {
	case CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirectionUnidirectional, CfInterconnectUpdateResponseModifiedInterconnectHealthCheckDirectionBidirectional:
		return true
	}
	return false
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
// [CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion interface {
	ImplementsCfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion()
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

func (r CfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetMagicHealthCheckTarget) ImplementsCfInterconnectUpdateResponseModifiedInterconnectHealthCheckTargetUnion() {
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
	// Identifier
	ID string `json:"id"`
	// True if automatic stateful return routing should be enabled for a tunnel, false
	// otherwise. Requires the `coupler_integration` account flag to be enabled;
	// requests setting this to `true` without that flag will be rejected.
	AutomaticReturnRouting bool                                       `json:"automatic_return_routing"`
	BGP                    CfInterconnectListResponseInterconnectsBGP `json:"bgp"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// Omitted in responses for version 1.5 interconnects.
	GRE         CfInterconnectListResponseInterconnectsGRE         `json:"gre"`
	HealthCheck CfInterconnectListResponseInterconnectsHealthCheck `json:"health_check"`
	// The IPv4 interface address for the interconnect. For MPLS Interconnects, use a
	// /30 or /31 prefix. For GRE Interconnects, a /30 or /31 prefix may be used.
	// Version 1.5 interconnects require a /31 prefix and may also use a prefix from
	// the account's authorized prefixes; otherwise, select the subnet from RFC 1918 or
	// the approved link-local ranges.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string `json:"name"`
	// Immutable interconnect version configured at creation time. One of:
	//
	// - "1"
	// - "1.5"
	// - "2"
	Version string `json:"version"`
	// An identifier that correlates this interconnect with the corresponding V2 CNI
	// interconnect resource.
	VirtualPortReservationID string                                     `json:"virtual_port_reservation_id"`
	JSON                     cfInterconnectListResponseInterconnectJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectJSON contains the JSON metadata for the
// struct [CfInterconnectListResponseInterconnect]
type cfInterconnectListResponseInterconnectJSON struct {
	ID                       apijson.Field
	AutomaticReturnRouting   apijson.Field
	BGP                      apijson.Field
	ColoName                 apijson.Field
	CreatedOn                apijson.Field
	Description              apijson.Field
	GRE                      apijson.Field
	HealthCheck              apijson.Field
	InterfaceAddress         apijson.Field
	InterfaceAddress6        apijson.Field
	ModifiedOn               apijson.Field
	Mtu                      apijson.Field
	Name                     apijson.Field
	Version                  apijson.Field
	VirtualPortReservationID apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CfInterconnectListResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectListResponseInterconnectsBGP struct {
	// Deprecated. Use customer_asn.
	//
	// Deprecated: deprecated
	AsNo int64 `json:"as_no"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CloudflareEndpoint string `json:"cloudflare_endpoint" format:"ipv4"`
	// ASN used on the customer end of the BGP session.
	CustomerASN int64 `json:"customer_asn"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CustomerEndpoint string `json:"customer_endpoint" format:"ipv4"`
	// ID of the BGP filter profile applied to routes advertised to the customer.
	ExportFilterID string `json:"export_filter_id"`
	// Prefixes in this list will be advertised to the customer device, in addition to
	// the routes in the Magic routing table.
	ExtraPrefixes []string `json:"extra_prefixes" format:"cidr"`
	// ID of the BGP filter profile applied to routes received from the customer.
	ImportFilterID string `json:"import_filter_id"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                                         `json:"md5_key"`
	JSON   cfInterconnectListResponseInterconnectsBGPJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectsBGPJSON contains the JSON metadata for
// the struct [CfInterconnectListResponseInterconnectsBGP]
type cfInterconnectListResponseInterconnectsBGPJSON struct {
	AsNo               apijson.Field
	CloudflareEndpoint apijson.Field
	CustomerASN        apijson.Field
	CustomerEndpoint   apijson.Field
	ExportFilterID     apijson.Field
	ExtraPrefixes      apijson.Field
	ImportFilterID     apijson.Field
	Md5Key             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectListResponseInterconnectsBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectsBGPJSON) RawJSON() string {
	return r.raw
}

// Omitted in responses for version 1.5 interconnects.
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
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the interconnect and the result comes back to Cloudflare
	// via the open Internet, or bidirectional where both the probe and result come and
	// go via the interconnect.
	Direction CfInterconnectListResponseInterconnectsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The source IPv4 address used for bidirectional health checks. Supported only for
	// version 1.5 interconnects. It is required when `direction` is `bidirectional`
	// and must be omitted (and is cleared) when `direction` is `unidirectional`. The
	// address must be within RFC1918 space, the approved link-local range
	// 169.254.240.0/20, or the Cloudflare reserved range 198.41.199.224/27.
	Source string `json:"source"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target CfInterconnectListResponseInterconnectsHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                        `json:"type"`
	JSON cfInterconnectListResponseInterconnectsHealthCheckJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectsHealthCheckJSON contains the JSON
// metadata for the struct [CfInterconnectListResponseInterconnectsHealthCheck]
type cfInterconnectListResponseInterconnectsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Source      apijson.Field
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

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the interconnect and the result comes back to Cloudflare
// via the open Internet, or bidirectional where both the probe and result come and
// go via the interconnect.
type CfInterconnectListResponseInterconnectsHealthCheckDirection string

const (
	CfInterconnectListResponseInterconnectsHealthCheckDirectionUnidirectional CfInterconnectListResponseInterconnectsHealthCheckDirection = "unidirectional"
	CfInterconnectListResponseInterconnectsHealthCheckDirectionBidirectional  CfInterconnectListResponseInterconnectsHealthCheckDirection = "bidirectional"
)

func (r CfInterconnectListResponseInterconnectsHealthCheckDirection) IsKnown() bool {
	switch r {
	case CfInterconnectListResponseInterconnectsHealthCheckDirectionUnidirectional, CfInterconnectListResponseInterconnectsHealthCheckDirectionBidirectional:
		return true
	}
	return false
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
// [CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type CfInterconnectListResponseInterconnectsHealthCheckTargetUnion interface {
	ImplementsCfInterconnectListResponseInterconnectsHealthCheckTargetUnion()
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

func (r CfInterconnectListResponseInterconnectsHealthCheckTargetMagicHealthCheckTarget) ImplementsCfInterconnectListResponseInterconnectsHealthCheckTargetUnion() {
}

type CfInterconnectBulkUpdateResponse struct {
	Modified              bool                                                   `json:"modified"`
	ModifiedInterconnects []CfInterconnectBulkUpdateResponseModifiedInterconnect `json:"modified_interconnects"`
	JSON                  cfInterconnectBulkUpdateResponseJSON                   `json:"-"`
}

// cfInterconnectBulkUpdateResponseJSON contains the JSON metadata for the struct
// [CfInterconnectBulkUpdateResponse]
type cfInterconnectBulkUpdateResponseJSON struct {
	Modified              apijson.Field
	ModifiedInterconnects apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectBulkUpdateResponseModifiedInterconnect struct {
	// Identifier
	ID string `json:"id"`
	// True if automatic stateful return routing should be enabled for a tunnel, false
	// otherwise. Requires the `coupler_integration` account flag to be enabled;
	// requests setting this to `true` without that flag will be rejected.
	AutomaticReturnRouting bool                                                     `json:"automatic_return_routing"`
	BGP                    CfInterconnectBulkUpdateResponseModifiedInterconnectsBGP `json:"bgp"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// Omitted in responses for version 1.5 interconnects.
	GRE         CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE         `json:"gre"`
	HealthCheck CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheck `json:"health_check"`
	// The IPv4 interface address for the interconnect. For MPLS Interconnects, use a
	// /30 or /31 prefix. For GRE Interconnects, a /30 or /31 prefix may be used.
	// Version 1.5 interconnects require a /31 prefix and may also use a prefix from
	// the account's authorized prefixes; otherwise, select the subnet from RFC 1918 or
	// the approved link-local ranges.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string `json:"name"`
	// Immutable interconnect version configured at creation time. One of:
	//
	// - "1"
	// - "1.5"
	// - "2"
	Version string `json:"version"`
	// An identifier that correlates this interconnect with the corresponding V2 CNI
	// interconnect resource.
	VirtualPortReservationID string                                                   `json:"virtual_port_reservation_id"`
	JSON                     cfInterconnectBulkUpdateResponseModifiedInterconnectJSON `json:"-"`
}

// cfInterconnectBulkUpdateResponseModifiedInterconnectJSON contains the JSON
// metadata for the struct [CfInterconnectBulkUpdateResponseModifiedInterconnect]
type cfInterconnectBulkUpdateResponseModifiedInterconnectJSON struct {
	ID                       apijson.Field
	AutomaticReturnRouting   apijson.Field
	BGP                      apijson.Field
	ColoName                 apijson.Field
	CreatedOn                apijson.Field
	Description              apijson.Field
	GRE                      apijson.Field
	HealthCheck              apijson.Field
	InterfaceAddress         apijson.Field
	InterfaceAddress6        apijson.Field
	ModifiedOn               apijson.Field
	Mtu                      apijson.Field
	Name                     apijson.Field
	Version                  apijson.Field
	VirtualPortReservationID apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseModifiedInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseModifiedInterconnectJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectBulkUpdateResponseModifiedInterconnectsBGP struct {
	// Deprecated. Use customer_asn.
	//
	// Deprecated: deprecated
	AsNo int64 `json:"as_no"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CloudflareEndpoint string `json:"cloudflare_endpoint" format:"ipv4"`
	// ASN used on the customer end of the BGP session.
	CustomerASN int64 `json:"customer_asn"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CustomerEndpoint string `json:"customer_endpoint" format:"ipv4"`
	// ID of the BGP filter profile applied to routes advertised to the customer.
	ExportFilterID string `json:"export_filter_id"`
	// Prefixes in this list will be advertised to the customer device, in addition to
	// the routes in the Magic routing table.
	ExtraPrefixes []string `json:"extra_prefixes" format:"cidr"`
	// ID of the BGP filter profile applied to routes received from the customer.
	ImportFilterID string `json:"import_filter_id"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                                                       `json:"md5_key"`
	JSON   cfInterconnectBulkUpdateResponseModifiedInterconnectsBGPJSON `json:"-"`
}

// cfInterconnectBulkUpdateResponseModifiedInterconnectsBGPJSON contains the JSON
// metadata for the struct
// [CfInterconnectBulkUpdateResponseModifiedInterconnectsBGP]
type cfInterconnectBulkUpdateResponseModifiedInterconnectsBGPJSON struct {
	AsNo               apijson.Field
	CloudflareEndpoint apijson.Field
	CustomerASN        apijson.Field
	CustomerEndpoint   apijson.Field
	ExportFilterID     apijson.Field
	ExtraPrefixes      apijson.Field
	ImportFilterID     apijson.Field
	Md5Key             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseModifiedInterconnectsBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseModifiedInterconnectsBGPJSON) RawJSON() string {
	return r.raw
}

// Omitted in responses for version 1.5 interconnects.
type CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                       `json:"cloudflare_endpoint"`
	JSON               cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON `json:"-"`
}

// cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON contains the JSON
// metadata for the struct
// [CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE]
type cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the interconnect and the result comes back to Cloudflare
	// via the open Internet, or bidirectional where both the probe and result come and
	// go via the interconnect.
	Direction CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The source IPv4 address used for bidirectional health checks. Supported only for
	// version 1.5 interconnects. It is required when `direction` is `bidirectional`
	// and must be omitted (and is cleared) when `direction` is `unidirectional`. The
	// address must be within RFC1918 space, the approved link-local range
	// 169.254.240.0/20, or the Cloudflare reserved range 198.41.199.224/27.
	Source string `json:"source"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                                      `json:"type"`
	JSON cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckJSON `json:"-"`
}

// cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckJSON contains
// the JSON metadata for the struct
// [CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheck]
type cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Source      apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the interconnect and the result comes back to Cloudflare
// via the open Internet, or bidirectional where both the probe and result come and
// go via the interconnect.
type CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirection string

const (
	CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirectionUnidirectional CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirection = "unidirectional"
	CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirectionBidirectional  CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirection = "bidirectional"
)

func (r CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirection) IsKnown() bool {
	switch r {
	case CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirectionUnidirectional, CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckDirectionBidirectional:
		return true
	}
	return false
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
// [CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetUnion interface {
	ImplementsCfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTarget{}),
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
type CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                                           `json:"saved"`
	JSON  cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTarget]
type cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r CfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetMagicHealthCheckTarget) ImplementsCfInterconnectBulkUpdateResponseModifiedInterconnectsHealthCheckTargetUnion() {
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
	// Identifier
	ID string `json:"id"`
	// True if automatic stateful return routing should be enabled for a tunnel, false
	// otherwise. Requires the `coupler_integration` account flag to be enabled;
	// requests setting this to `true` without that flag will be rejected.
	AutomaticReturnRouting bool                                     `json:"automatic_return_routing"`
	BGP                    CfInterconnectGetResponseInterconnectBGP `json:"bgp"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// Omitted in responses for version 1.5 interconnects.
	GRE         CfInterconnectGetResponseInterconnectGRE         `json:"gre"`
	HealthCheck CfInterconnectGetResponseInterconnectHealthCheck `json:"health_check"`
	// The IPv4 interface address for the interconnect. For MPLS Interconnects, use a
	// /30 or /31 prefix. For GRE Interconnects, a /30 or /31 prefix may be used.
	// Version 1.5 interconnects require a /31 prefix and may also use a prefix from
	// the account's authorized prefixes; otherwise, select the subnet from RFC 1918 or
	// the approved link-local ranges.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string `json:"name"`
	// Immutable interconnect version configured at creation time. One of:
	//
	// - "1"
	// - "1.5"
	// - "2"
	Version string `json:"version"`
	// An identifier that correlates this interconnect with the corresponding V2 CNI
	// interconnect resource.
	VirtualPortReservationID string                                    `json:"virtual_port_reservation_id"`
	JSON                     cfInterconnectGetResponseInterconnectJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectJSON contains the JSON metadata for the
// struct [CfInterconnectGetResponseInterconnect]
type cfInterconnectGetResponseInterconnectJSON struct {
	ID                       apijson.Field
	AutomaticReturnRouting   apijson.Field
	BGP                      apijson.Field
	ColoName                 apijson.Field
	CreatedOn                apijson.Field
	Description              apijson.Field
	GRE                      apijson.Field
	HealthCheck              apijson.Field
	InterfaceAddress         apijson.Field
	InterfaceAddress6        apijson.Field
	ModifiedOn               apijson.Field
	Mtu                      apijson.Field
	Name                     apijson.Field
	Version                  apijson.Field
	VirtualPortReservationID apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *CfInterconnectGetResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectGetResponseInterconnectBGP struct {
	// Deprecated. Use customer_asn.
	//
	// Deprecated: deprecated
	AsNo int64 `json:"as_no"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CloudflareEndpoint string `json:"cloudflare_endpoint" format:"ipv4"`
	// ASN used on the customer end of the BGP session.
	CustomerASN int64 `json:"customer_asn"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CustomerEndpoint string `json:"customer_endpoint" format:"ipv4"`
	// ID of the BGP filter profile applied to routes advertised to the customer.
	ExportFilterID string `json:"export_filter_id"`
	// Prefixes in this list will be advertised to the customer device, in addition to
	// the routes in the Magic routing table.
	ExtraPrefixes []string `json:"extra_prefixes" format:"cidr"`
	// ID of the BGP filter profile applied to routes received from the customer.
	ImportFilterID string `json:"import_filter_id"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                                       `json:"md5_key"`
	JSON   cfInterconnectGetResponseInterconnectBGPJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectBGPJSON contains the JSON metadata for the
// struct [CfInterconnectGetResponseInterconnectBGP]
type cfInterconnectGetResponseInterconnectBGPJSON struct {
	AsNo               apijson.Field
	CloudflareEndpoint apijson.Field
	CustomerASN        apijson.Field
	CustomerEndpoint   apijson.Field
	ExportFilterID     apijson.Field
	ExtraPrefixes      apijson.Field
	ImportFilterID     apijson.Field
	Md5Key             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectGetResponseInterconnectBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectBGPJSON) RawJSON() string {
	return r.raw
}

// Omitted in responses for version 1.5 interconnects.
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
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the interconnect and the result comes back to Cloudflare
	// via the open Internet, or bidirectional where both the probe and result come and
	// go via the interconnect.
	Direction CfInterconnectGetResponseInterconnectHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The source IPv4 address used for bidirectional health checks. Supported only for
	// version 1.5 interconnects. It is required when `direction` is `bidirectional`
	// and must be omitted (and is cleared) when `direction` is `unidirectional`. The
	// address must be within RFC1918 space, the approved link-local range
	// 169.254.240.0/20, or the Cloudflare reserved range 198.41.199.224/27.
	Source string `json:"source"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target CfInterconnectGetResponseInterconnectHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                      `json:"type"`
	JSON cfInterconnectGetResponseInterconnectHealthCheckJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectHealthCheckJSON contains the JSON metadata
// for the struct [CfInterconnectGetResponseInterconnectHealthCheck]
type cfInterconnectGetResponseInterconnectHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Source      apijson.Field
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

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the interconnect and the result comes back to Cloudflare
// via the open Internet, or bidirectional where both the probe and result come and
// go via the interconnect.
type CfInterconnectGetResponseInterconnectHealthCheckDirection string

const (
	CfInterconnectGetResponseInterconnectHealthCheckDirectionUnidirectional CfInterconnectGetResponseInterconnectHealthCheckDirection = "unidirectional"
	CfInterconnectGetResponseInterconnectHealthCheckDirectionBidirectional  CfInterconnectGetResponseInterconnectHealthCheckDirection = "bidirectional"
)

func (r CfInterconnectGetResponseInterconnectHealthCheckDirection) IsKnown() bool {
	switch r {
	case CfInterconnectGetResponseInterconnectHealthCheckDirectionUnidirectional, CfInterconnectGetResponseInterconnectHealthCheckDirectionBidirectional:
		return true
	}
	return false
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
// [CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type CfInterconnectGetResponseInterconnectHealthCheckTargetUnion interface {
	ImplementsCfInterconnectGetResponseInterconnectHealthCheckTargetUnion()
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

func (r CfInterconnectGetResponseInterconnectHealthCheckTargetMagicHealthCheckTarget) ImplementsCfInterconnectGetResponseInterconnectHealthCheckTargetUnion() {
}

type CfInterconnectUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// True if automatic stateful return routing should be enabled for a tunnel, false
	// otherwise. Requires the `coupler_integration` account flag to be enabled;
	// requests setting this to `true` without that flag will be rejected.
	AutomaticReturnRouting param.Field[bool]                          `json:"automatic_return_routing"`
	BGP                    param.Field[CfInterconnectUpdateParamsBGP] `json:"bgp"`
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// Not configurable for version 1.5 interconnects; supplying it returns an error.
	GRE         param.Field[CfInterconnectUpdateParamsGRE]         `json:"gre"`
	HealthCheck param.Field[CfInterconnectUpdateParamsHealthCheck] `json:"health_check"`
	// The IPv4 interface address for the interconnect. For MPLS Interconnects, use a
	// /30 or /31 prefix. For GRE Interconnects, a /30 or /31 prefix may be used.
	// Version 1.5 interconnects require a /31 prefix and may also use a prefix from
	// the account's authorized prefixes; otherwise, select the subnet from RFC 1918 or
	// the approved link-local ranges.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 param.Field[string] `json:"interface_address6"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name              param.Field[string] `json:"name"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

func (r CfInterconnectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CfInterconnectUpdateParamsBGP struct {
	// Deprecated. Use customer_asn.
	//
	// Deprecated: deprecated
	AsNo param.Field[int64] `json:"as_no"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint" format:"ipv4"`
	// ASN used on the customer end of the BGP session.
	CustomerASN param.Field[int64] `json:"customer_asn"`
	// Read-only for v1.5; derived from interface_address.
	//
	// Deprecated: deprecated
	CustomerEndpoint param.Field[string] `json:"customer_endpoint" format:"ipv4"`
	// ID of the BGP filter profile applied to routes advertised to the customer.
	ExportFilterID param.Field[string] `json:"export_filter_id"`
	// Prefixes in this list will be advertised to the customer device, in addition to
	// the routes in the Magic routing table.
	ExtraPrefixes param.Field[[]string] `json:"extra_prefixes" format:"cidr"`
	// ID of the BGP filter profile applied to routes received from the customer.
	ImportFilterID param.Field[string] `json:"import_filter_id"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key param.Field[string] `json:"md5_key"`
}

func (r CfInterconnectUpdateParamsBGP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Not configurable for version 1.5 interconnects; supplying it returns an error.
type CfInterconnectUpdateParamsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r CfInterconnectUpdateParamsGRE) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CfInterconnectUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the interconnect and the result comes back to Cloudflare
	// via the open Internet, or bidirectional where both the probe and result come and
	// go via the interconnect.
	Direction param.Field[CfInterconnectUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[HealthCheckRate] `json:"rate"`
	// The source IPv4 address used for bidirectional health checks. Supported only for
	// version 1.5 interconnects. It is required when `direction` is `bidirectional`
	// and must be omitted (and is cleared) when `direction` is `unidirectional`. The
	// address must be within RFC1918 space, the approved link-local range
	// 169.254.240.0/20, or the Cloudflare reserved range 198.41.199.224/27.
	Source param.Field[string] `json:"source"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[CfInterconnectUpdateParamsHealthCheckTargetUnion] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[HealthCheckType] `json:"type"`
}

func (r CfInterconnectUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the interconnect and the result comes back to Cloudflare
// via the open Internet, or bidirectional where both the probe and result come and
// go via the interconnect.
type CfInterconnectUpdateParamsHealthCheckDirection string

const (
	CfInterconnectUpdateParamsHealthCheckDirectionUnidirectional CfInterconnectUpdateParamsHealthCheckDirection = "unidirectional"
	CfInterconnectUpdateParamsHealthCheckDirectionBidirectional  CfInterconnectUpdateParamsHealthCheckDirection = "bidirectional"
)

func (r CfInterconnectUpdateParamsHealthCheckDirection) IsKnown() bool {
	switch r {
	case CfInterconnectUpdateParamsHealthCheckDirectionUnidirectional, CfInterconnectUpdateParamsHealthCheckDirectionBidirectional:
		return true
	}
	return false
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
	ImplementsCfInterconnectUpdateParamsHealthCheckTargetUnion()
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

func (r CfInterconnectUpdateParamsHealthCheckTargetMagicHealthCheckTarget) ImplementsCfInterconnectUpdateParamsHealthCheckTargetUnion() {
}

type CfInterconnectUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors" api:"required"`
	Messages []shared.ResponseInfo        `json:"messages" api:"required"`
	Result   CfInterconnectUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success CfInterconnectUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
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
	AccountID         param.Field[string] `path:"account_id" api:"required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type CfInterconnectListResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors" api:"required"`
	Messages []shared.ResponseInfo      `json:"messages" api:"required"`
	Result   CfInterconnectListResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success CfInterconnectListResponseEnvelopeSuccess `json:"success" api:"required"`
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

type CfInterconnectBulkUpdateParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id" api:"required"`
	Body              interface{}         `json:"body" api:"required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

func (r CfInterconnectBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CfInterconnectBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors" api:"required"`
	Messages []shared.ResponseInfo            `json:"messages" api:"required"`
	Result   CfInterconnectBulkUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success CfInterconnectBulkUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    cfInterconnectBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CfInterconnectBulkUpdateResponseEnvelope]
type cfInterconnectBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectBulkUpdateResponseEnvelopeSuccess bool

const (
	CfInterconnectBulkUpdateResponseEnvelopeSuccessTrue CfInterconnectBulkUpdateResponseEnvelopeSuccess = true
)

func (r CfInterconnectBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CfInterconnectGetParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id" api:"required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type CfInterconnectGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors" api:"required"`
	Messages []shared.ResponseInfo     `json:"messages" api:"required"`
	Result   CfInterconnectGetResponse `json:"result" api:"required"`
	// Whether the API call was successful
	Success CfInterconnectGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
