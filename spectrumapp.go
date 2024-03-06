// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// SpectrumAppService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSpectrumAppService] method
// instead.
type SpectrumAppService struct {
	Options []option.RequestOption
}

// NewSpectrumAppService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpectrumAppService(opts ...option.RequestOption) (r *SpectrumAppService) {
	r = &SpectrumAppService{}
	r.Options = opts
	return
}

// Creates a new Spectrum application from a configuration using a name for the
// origin.
func (r *SpectrumAppService) New(ctx context.Context, zone string, body SpectrumAppNewParams, opts ...option.RequestOption) (res *SpectrumAppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAppNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a previously existing application's configuration that uses a name for
// the origin.
func (r *SpectrumAppService) Update(ctx context.Context, zone string, appID string, body SpectrumAppUpdateParams, opts ...option.RequestOption) (res *SpectrumAppUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAppUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *SpectrumAppService) List(ctx context.Context, zone string, query SpectrumAppListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[SpectrumAppListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/spectrum/apps", zone)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *SpectrumAppService) ListAutoPaging(ctx context.Context, zone string, query SpectrumAppListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[SpectrumAppListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zone, query, opts...))
}

// Deletes a previously existing application.
func (r *SpectrumAppService) Delete(ctx context.Context, zone string, appID string, opts ...option.RequestOption) (res *SpectrumAppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAppDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the application configuration of a specific application inside a zone.
func (r *SpectrumAppService) Get(ctx context.Context, zone string, appID string, opts ...option.RequestOption) (res *SpectrumAppGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAppGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SpectrumAppNewResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS SpectrumAppNewResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs SpectrumAppNewResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS SpectrumAppNewResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort SpectrumAppNewResponseOriginPort `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol SpectrumAppNewResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS SpectrumAppNewResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType SpectrumAppNewResponseTrafficType `json:"traffic_type"`
	JSON        spectrumAppNewResponseJSON        `json:"-"`
}

// spectrumAppNewResponseJSON contains the JSON metadata for the struct
// [SpectrumAppNewResponse]
type spectrumAppNewResponseJSON struct {
	ID               apijson.Field
	ArgoSmartRouting apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	ModifiedOn       apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	Protocol         apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SpectrumAppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseJSON) RawJSON() string {
	return r.raw
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppNewResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type SpectrumAppNewResponseDNSType `json:"type"`
	JSON spectrumAppNewResponseDNSJSON `json:"-"`
}

// spectrumAppNewResponseDNSJSON contains the JSON metadata for the struct
// [SpectrumAppNewResponseDNS]
type spectrumAppNewResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppNewResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type SpectrumAppNewResponseDNSType string

const (
	SpectrumAppNewResponseDNSTypeCNAME   SpectrumAppNewResponseDNSType = "CNAME"
	SpectrumAppNewResponseDNSTypeAddress SpectrumAppNewResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by
// [SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable] or
// [SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type SpectrumAppNewResponseEdgeIPs interface {
	implementsSpectrumAppNewResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SpectrumAppNewResponseEdgeIPs)(nil)).Elem(), "")
}

type SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType `json:"type"`
	JSON spectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON `json:"-"`
}

// spectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON contains the
// JSON metadata for the struct
// [SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable]
type spectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppNewResponseEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs []string `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType `json:"type"`
	JSON spectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON `json:"-"`
}

// spectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON contains
// the JSON metadata for the struct
// [SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable]
type spectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON struct {
	IPs         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppNewResponseEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic SpectrumAppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppNewResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type SpectrumAppNewResponseOriginDNSType `json:"type"`
	JSON spectrumAppNewResponseOriginDNSJSON `json:"-"`
}

// spectrumAppNewResponseOriginDNSJSON contains the JSON metadata for the struct
// [SpectrumAppNewResponseOriginDNS]
type spectrumAppNewResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppNewResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseOriginDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppNewResponseOriginDNSType string

const (
	SpectrumAppNewResponseOriginDNSTypeEmpty SpectrumAppNewResponseOriginDNSType = ""
	SpectrumAppNewResponseOriginDNSTypeA     SpectrumAppNewResponseOriginDNSType = "A"
	SpectrumAppNewResponseOriginDNSTypeAAAA  SpectrumAppNewResponseOriginDNSType = "AAAA"
	SpectrumAppNewResponseOriginDNSTypeSRV   SpectrumAppNewResponseOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type SpectrumAppNewResponseOriginPort interface {
	ImplementsSpectrumAppNewResponseOriginPort()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpectrumAppNewResponseOriginPort)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type SpectrumAppNewResponseProxyProtocol string

const (
	SpectrumAppNewResponseProxyProtocolOff    SpectrumAppNewResponseProxyProtocol = "off"
	SpectrumAppNewResponseProxyProtocolV1     SpectrumAppNewResponseProxyProtocol = "v1"
	SpectrumAppNewResponseProxyProtocolV2     SpectrumAppNewResponseProxyProtocol = "v2"
	SpectrumAppNewResponseProxyProtocolSimple SpectrumAppNewResponseProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppNewResponseTLS string

const (
	SpectrumAppNewResponseTLSOff      SpectrumAppNewResponseTLS = "off"
	SpectrumAppNewResponseTLSFlexible SpectrumAppNewResponseTLS = "flexible"
	SpectrumAppNewResponseTLSFull     SpectrumAppNewResponseTLS = "full"
	SpectrumAppNewResponseTLSStrict   SpectrumAppNewResponseTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppNewResponseTrafficType string

const (
	SpectrumAppNewResponseTrafficTypeDirect SpectrumAppNewResponseTrafficType = "direct"
	SpectrumAppNewResponseTrafficTypeHTTP   SpectrumAppNewResponseTrafficType = "http"
	SpectrumAppNewResponseTrafficTypeHTTPS  SpectrumAppNewResponseTrafficType = "https"
)

type SpectrumAppUpdateResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS SpectrumAppUpdateResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs SpectrumAppUpdateResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS SpectrumAppUpdateResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort SpectrumAppUpdateResponseOriginPort `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol SpectrumAppUpdateResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS SpectrumAppUpdateResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType SpectrumAppUpdateResponseTrafficType `json:"traffic_type"`
	JSON        spectrumAppUpdateResponseJSON        `json:"-"`
}

// spectrumAppUpdateResponseJSON contains the JSON metadata for the struct
// [SpectrumAppUpdateResponse]
type spectrumAppUpdateResponseJSON struct {
	ID               apijson.Field
	ArgoSmartRouting apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	ModifiedOn       apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	Protocol         apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppUpdateResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type SpectrumAppUpdateResponseDNSType `json:"type"`
	JSON spectrumAppUpdateResponseDNSJSON `json:"-"`
}

// spectrumAppUpdateResponseDNSJSON contains the JSON metadata for the struct
// [SpectrumAppUpdateResponseDNS]
type spectrumAppUpdateResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type SpectrumAppUpdateResponseDNSType string

const (
	SpectrumAppUpdateResponseDNSTypeCNAME   SpectrumAppUpdateResponseDNSType = "CNAME"
	SpectrumAppUpdateResponseDNSTypeAddress SpectrumAppUpdateResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by
// [SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable] or
// [SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type SpectrumAppUpdateResponseEdgeIPs interface {
	implementsSpectrumAppUpdateResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SpectrumAppUpdateResponseEdgeIPs)(nil)).Elem(), "")
}

type SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType `json:"type"`
	JSON spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON `json:"-"`
}

// spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON contains
// the JSON metadata for the struct
// [SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable]
type spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppUpdateResponseEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs []string `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType `json:"type"`
	JSON spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON `json:"-"`
}

// spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON
// contains the JSON metadata for the struct
// [SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable]
type spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON struct {
	IPs         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppUpdateResponseEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic SpectrumAppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppUpdateResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type SpectrumAppUpdateResponseOriginDNSType `json:"type"`
	JSON spectrumAppUpdateResponseOriginDNSJSON `json:"-"`
}

// spectrumAppUpdateResponseOriginDNSJSON contains the JSON metadata for the struct
// [SpectrumAppUpdateResponseOriginDNS]
type spectrumAppUpdateResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseOriginDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppUpdateResponseOriginDNSType string

const (
	SpectrumAppUpdateResponseOriginDNSTypeEmpty SpectrumAppUpdateResponseOriginDNSType = ""
	SpectrumAppUpdateResponseOriginDNSTypeA     SpectrumAppUpdateResponseOriginDNSType = "A"
	SpectrumAppUpdateResponseOriginDNSTypeAAAA  SpectrumAppUpdateResponseOriginDNSType = "AAAA"
	SpectrumAppUpdateResponseOriginDNSTypeSRV   SpectrumAppUpdateResponseOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type SpectrumAppUpdateResponseOriginPort interface {
	ImplementsSpectrumAppUpdateResponseOriginPort()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpectrumAppUpdateResponseOriginPort)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type SpectrumAppUpdateResponseProxyProtocol string

const (
	SpectrumAppUpdateResponseProxyProtocolOff    SpectrumAppUpdateResponseProxyProtocol = "off"
	SpectrumAppUpdateResponseProxyProtocolV1     SpectrumAppUpdateResponseProxyProtocol = "v1"
	SpectrumAppUpdateResponseProxyProtocolV2     SpectrumAppUpdateResponseProxyProtocol = "v2"
	SpectrumAppUpdateResponseProxyProtocolSimple SpectrumAppUpdateResponseProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppUpdateResponseTLS string

const (
	SpectrumAppUpdateResponseTLSOff      SpectrumAppUpdateResponseTLS = "off"
	SpectrumAppUpdateResponseTLSFlexible SpectrumAppUpdateResponseTLS = "flexible"
	SpectrumAppUpdateResponseTLSFull     SpectrumAppUpdateResponseTLS = "full"
	SpectrumAppUpdateResponseTLSStrict   SpectrumAppUpdateResponseTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppUpdateResponseTrafficType string

const (
	SpectrumAppUpdateResponseTrafficTypeDirect SpectrumAppUpdateResponseTrafficType = "direct"
	SpectrumAppUpdateResponseTrafficTypeHTTP   SpectrumAppUpdateResponseTrafficType = "http"
	SpectrumAppUpdateResponseTrafficTypeHTTPS  SpectrumAppUpdateResponseTrafficType = "https"
)

type SpectrumAppListResponse = interface{}

type SpectrumAppDeleteResponse struct {
	// Application identifier.
	ID   string                        `json:"id"`
	JSON spectrumAppDeleteResponseJSON `json:"-"`
}

// spectrumAppDeleteResponseJSON contains the JSON metadata for the struct
// [SpectrumAppDeleteResponse]
type spectrumAppDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [SpectrumAppGetResponseUnknown] or [shared.UnionString].
type SpectrumAppGetResponse interface {
	ImplementsSpectrumAppGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpectrumAppGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SpectrumAppNewParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[SpectrumAppNewParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[SpectrumAppNewParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[SpectrumAppNewParamsOriginPort] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[SpectrumAppNewParamsEdgeIPs] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[SpectrumAppNewParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[SpectrumAppNewParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[SpectrumAppNewParamsTrafficType] `json:"traffic_type"`
}

func (r SpectrumAppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppNewParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[SpectrumAppNewParamsDNSType] `json:"type"`
}

func (r SpectrumAppNewParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type SpectrumAppNewParamsDNSType string

const (
	SpectrumAppNewParamsDNSTypeCNAME   SpectrumAppNewParamsDNSType = "CNAME"
	SpectrumAppNewParamsDNSTypeAddress SpectrumAppNewParamsDNSType = "ADDRESS"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppNewParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[SpectrumAppNewParamsOriginDNSType] `json:"type"`
}

func (r SpectrumAppNewParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppNewParamsOriginDNSType string

const (
	SpectrumAppNewParamsOriginDNSTypeEmpty SpectrumAppNewParamsOriginDNSType = ""
	SpectrumAppNewParamsOriginDNSTypeA     SpectrumAppNewParamsOriginDNSType = "A"
	SpectrumAppNewParamsOriginDNSTypeAAAA  SpectrumAppNewParamsOriginDNSType = "AAAA"
	SpectrumAppNewParamsOriginDNSTypeSRV   SpectrumAppNewParamsOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type SpectrumAppNewParamsOriginPort interface {
	ImplementsSpectrumAppNewParamsOriginPort()
}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable],
// [SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type SpectrumAppNewParamsEdgeIPs interface {
	implementsSpectrumAppNewParamsEdgeIPs()
}

type SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType] `json:"type"`
}

func (r SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppNewParamsEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs param.Field[[]string] `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type param.Field[SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType] `json:"type"`
}

func (r SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppNewParamsEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic SpectrumAppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type SpectrumAppNewParamsProxyProtocol string

const (
	SpectrumAppNewParamsProxyProtocolOff    SpectrumAppNewParamsProxyProtocol = "off"
	SpectrumAppNewParamsProxyProtocolV1     SpectrumAppNewParamsProxyProtocol = "v1"
	SpectrumAppNewParamsProxyProtocolV2     SpectrumAppNewParamsProxyProtocol = "v2"
	SpectrumAppNewParamsProxyProtocolSimple SpectrumAppNewParamsProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppNewParamsTLS string

const (
	SpectrumAppNewParamsTLSOff      SpectrumAppNewParamsTLS = "off"
	SpectrumAppNewParamsTLSFlexible SpectrumAppNewParamsTLS = "flexible"
	SpectrumAppNewParamsTLSFull     SpectrumAppNewParamsTLS = "full"
	SpectrumAppNewParamsTLSStrict   SpectrumAppNewParamsTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppNewParamsTrafficType string

const (
	SpectrumAppNewParamsTrafficTypeDirect SpectrumAppNewParamsTrafficType = "direct"
	SpectrumAppNewParamsTrafficTypeHTTP   SpectrumAppNewParamsTrafficType = "http"
	SpectrumAppNewParamsTrafficTypeHTTPS  SpectrumAppNewParamsTrafficType = "https"
)

type SpectrumAppNewResponseEnvelope struct {
	Errors   []SpectrumAppNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAppNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SpectrumAppNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SpectrumAppNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAppNewResponseEnvelopeJSON    `json:"-"`
}

// spectrumAppNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpectrumAppNewResponseEnvelope]
type spectrumAppNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    spectrumAppNewResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAppNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpectrumAppNewResponseEnvelopeErrors]
type spectrumAppNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    spectrumAppNewResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAppNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpectrumAppNewResponseEnvelopeMessages]
type spectrumAppNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SpectrumAppNewResponseEnvelopeSuccess bool

const (
	SpectrumAppNewResponseEnvelopeSuccessTrue SpectrumAppNewResponseEnvelopeSuccess = true
)

type SpectrumAppUpdateParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[SpectrumAppUpdateParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[SpectrumAppUpdateParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[SpectrumAppUpdateParamsOriginPort] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[SpectrumAppUpdateParamsEdgeIPs] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[SpectrumAppUpdateParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[SpectrumAppUpdateParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[SpectrumAppUpdateParamsTrafficType] `json:"traffic_type"`
}

func (r SpectrumAppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppUpdateParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[SpectrumAppUpdateParamsDNSType] `json:"type"`
}

func (r SpectrumAppUpdateParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type SpectrumAppUpdateParamsDNSType string

const (
	SpectrumAppUpdateParamsDNSTypeCNAME   SpectrumAppUpdateParamsDNSType = "CNAME"
	SpectrumAppUpdateParamsDNSTypeAddress SpectrumAppUpdateParamsDNSType = "ADDRESS"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppUpdateParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[SpectrumAppUpdateParamsOriginDNSType] `json:"type"`
}

func (r SpectrumAppUpdateParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppUpdateParamsOriginDNSType string

const (
	SpectrumAppUpdateParamsOriginDNSTypeEmpty SpectrumAppUpdateParamsOriginDNSType = ""
	SpectrumAppUpdateParamsOriginDNSTypeA     SpectrumAppUpdateParamsOriginDNSType = "A"
	SpectrumAppUpdateParamsOriginDNSTypeAAAA  SpectrumAppUpdateParamsOriginDNSType = "AAAA"
	SpectrumAppUpdateParamsOriginDNSTypeSRV   SpectrumAppUpdateParamsOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type SpectrumAppUpdateParamsOriginPort interface {
	ImplementsSpectrumAppUpdateParamsOriginPort()
}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable],
// [SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type SpectrumAppUpdateParamsEdgeIPs interface {
	implementsSpectrumAppUpdateParamsEdgeIPs()
}

type SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType] `json:"type"`
}

func (r SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppUpdateParamsEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs param.Field[[]string] `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type param.Field[SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType] `json:"type"`
}

func (r SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppUpdateParamsEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic SpectrumAppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type SpectrumAppUpdateParamsProxyProtocol string

const (
	SpectrumAppUpdateParamsProxyProtocolOff    SpectrumAppUpdateParamsProxyProtocol = "off"
	SpectrumAppUpdateParamsProxyProtocolV1     SpectrumAppUpdateParamsProxyProtocol = "v1"
	SpectrumAppUpdateParamsProxyProtocolV2     SpectrumAppUpdateParamsProxyProtocol = "v2"
	SpectrumAppUpdateParamsProxyProtocolSimple SpectrumAppUpdateParamsProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppUpdateParamsTLS string

const (
	SpectrumAppUpdateParamsTLSOff      SpectrumAppUpdateParamsTLS = "off"
	SpectrumAppUpdateParamsTLSFlexible SpectrumAppUpdateParamsTLS = "flexible"
	SpectrumAppUpdateParamsTLSFull     SpectrumAppUpdateParamsTLS = "full"
	SpectrumAppUpdateParamsTLSStrict   SpectrumAppUpdateParamsTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppUpdateParamsTrafficType string

const (
	SpectrumAppUpdateParamsTrafficTypeDirect SpectrumAppUpdateParamsTrafficType = "direct"
	SpectrumAppUpdateParamsTrafficTypeHTTP   SpectrumAppUpdateParamsTrafficType = "http"
	SpectrumAppUpdateParamsTrafficTypeHTTPS  SpectrumAppUpdateParamsTrafficType = "https"
)

type SpectrumAppUpdateResponseEnvelope struct {
	Errors   []SpectrumAppUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAppUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SpectrumAppUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SpectrumAppUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAppUpdateResponseEnvelopeJSON    `json:"-"`
}

// spectrumAppUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpectrumAppUpdateResponseEnvelope]
type spectrumAppUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    spectrumAppUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAppUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpectrumAppUpdateResponseEnvelopeErrors]
type spectrumAppUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    spectrumAppUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAppUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpectrumAppUpdateResponseEnvelopeMessages]
type spectrumAppUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SpectrumAppUpdateResponseEnvelopeSuccess bool

const (
	SpectrumAppUpdateResponseEnvelopeSuccessTrue SpectrumAppUpdateResponseEnvelopeSuccess = true
)

type SpectrumAppListParams struct {
	// Sets the direction by which results are ordered.
	Direction param.Field[SpectrumAppListParamsDirection] `query:"direction"`
	// Application field by which results are ordered.
	Order param.Field[SpectrumAppListParamsOrder] `query:"order"`
	// Page number of paginated results. This parameter is required in order to use
	// other pagination parameters. If included in the query, `result_info` will be
	// present in the response.
	Page param.Field[float64] `query:"page"`
	// Sets the maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [SpectrumAppListParams]'s query parameters as `url.Values`.
func (r SpectrumAppListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sets the direction by which results are ordered.
type SpectrumAppListParamsDirection string

const (
	SpectrumAppListParamsDirectionAsc  SpectrumAppListParamsDirection = "asc"
	SpectrumAppListParamsDirectionDesc SpectrumAppListParamsDirection = "desc"
)

// Application field by which results are ordered.
type SpectrumAppListParamsOrder string

const (
	SpectrumAppListParamsOrderProtocol   SpectrumAppListParamsOrder = "protocol"
	SpectrumAppListParamsOrderAppID      SpectrumAppListParamsOrder = "app_id"
	SpectrumAppListParamsOrderCreatedOn  SpectrumAppListParamsOrder = "created_on"
	SpectrumAppListParamsOrderModifiedOn SpectrumAppListParamsOrder = "modified_on"
	SpectrumAppListParamsOrderDNS        SpectrumAppListParamsOrder = "dns"
)

type SpectrumAppDeleteResponseEnvelope struct {
	Errors   []SpectrumAppDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAppDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SpectrumAppDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SpectrumAppDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAppDeleteResponseEnvelopeJSON    `json:"-"`
}

// spectrumAppDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpectrumAppDeleteResponseEnvelope]
type spectrumAppDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    spectrumAppDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAppDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpectrumAppDeleteResponseEnvelopeErrors]
type spectrumAppDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    spectrumAppDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAppDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpectrumAppDeleteResponseEnvelopeMessages]
type spectrumAppDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SpectrumAppDeleteResponseEnvelopeSuccess bool

const (
	SpectrumAppDeleteResponseEnvelopeSuccessTrue SpectrumAppDeleteResponseEnvelopeSuccess = true
)

type SpectrumAppGetResponseEnvelope struct {
	Errors   []SpectrumAppGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAppGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SpectrumAppGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SpectrumAppGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAppGetResponseEnvelopeJSON    `json:"-"`
}

// spectrumAppGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpectrumAppGetResponseEnvelope]
type spectrumAppGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    spectrumAppGetResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAppGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpectrumAppGetResponseEnvelopeErrors]
type spectrumAppGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SpectrumAppGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    spectrumAppGetResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAppGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SpectrumAppGetResponseEnvelopeMessages]
type spectrumAppGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAppGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SpectrumAppGetResponseEnvelopeSuccess bool

const (
	SpectrumAppGetResponseEnvelopeSuccessTrue SpectrumAppGetResponseEnvelopeSuccess = true
)
