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

// Updates a previously existing application's configuration that uses a name for
// the origin.
func (r *SpectrumAppService) Replace(ctx context.Context, zone string, appID string, body SpectrumAppReplaceParams, opts ...option.RequestOption) (res *SpectrumAppReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAppReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// The type of DNS record associated with the application.
type SpectrumAppNewResponseDNSType string

const (
	SpectrumAppNewResponseDNSTypeCname   SpectrumAppNewResponseDNSType = "CNAME"
	SpectrumAppNewResponseDNSTypeAddress SpectrumAppNewResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by [SpectrumAppNewResponseEdgeIPsObject] or
// [SpectrumAppNewResponseEdgeIPsObject].
type SpectrumAppNewResponseEdgeIPs interface {
	implementsSpectrumAppNewResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SpectrumAppNewResponseEdgeIPs)(nil)).Elem(), "")
}

type SpectrumAppNewResponseEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity SpectrumAppNewResponseEdgeIPsObjectConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type SpectrumAppNewResponseEdgeIPsObjectType `json:"type"`
	JSON spectrumAppNewResponseEdgeIPsObjectJSON `json:"-"`
}

// spectrumAppNewResponseEdgeIPsObjectJSON contains the JSON metadata for the
// struct [SpectrumAppNewResponseEdgeIPsObject]
type spectrumAppNewResponseEdgeIPsObjectJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SpectrumAppNewResponseEdgeIPsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SpectrumAppNewResponseEdgeIPsObject) implementsSpectrumAppNewResponseEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppNewResponseEdgeIPsObjectConnectivity string

const (
	SpectrumAppNewResponseEdgeIPsObjectConnectivityAll  SpectrumAppNewResponseEdgeIPsObjectConnectivity = "all"
	SpectrumAppNewResponseEdgeIPsObjectConnectivityIPV4 SpectrumAppNewResponseEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppNewResponseEdgeIPsObjectConnectivityIPV6 SpectrumAppNewResponseEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppNewResponseEdgeIPsObjectType string

const (
	SpectrumAppNewResponseEdgeIPsObjectTypeDynamic SpectrumAppNewResponseEdgeIPsObjectType = "dynamic"
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

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppNewResponseOriginDNSType string

const (
	SpectrumAppNewResponseOriginDNSTypeEmpty SpectrumAppNewResponseOriginDNSType = ""
	SpectrumAppNewResponseOriginDNSTypeA     SpectrumAppNewResponseOriginDNSType = "A"
	SpectrumAppNewResponseOriginDNSTypeAaaa  SpectrumAppNewResponseOriginDNSType = "AAAA"
	SpectrumAppNewResponseOriginDNSTypeSrv   SpectrumAppNewResponseOriginDNSType = "SRV"
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

type SpectrumAppReplaceResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS SpectrumAppReplaceResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs SpectrumAppReplaceResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS SpectrumAppReplaceResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort SpectrumAppReplaceResponseOriginPort `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol SpectrumAppReplaceResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS SpectrumAppReplaceResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType SpectrumAppReplaceResponseTrafficType `json:"traffic_type"`
	JSON        spectrumAppReplaceResponseJSON        `json:"-"`
}

// spectrumAppReplaceResponseJSON contains the JSON metadata for the struct
// [SpectrumAppReplaceResponse]
type spectrumAppReplaceResponseJSON struct {
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

func (r *SpectrumAppReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppReplaceResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type SpectrumAppReplaceResponseDNSType `json:"type"`
	JSON spectrumAppReplaceResponseDNSJSON `json:"-"`
}

// spectrumAppReplaceResponseDNSJSON contains the JSON metadata for the struct
// [SpectrumAppReplaceResponseDNS]
type spectrumAppReplaceResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppReplaceResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of DNS record associated with the application.
type SpectrumAppReplaceResponseDNSType string

const (
	SpectrumAppReplaceResponseDNSTypeCname   SpectrumAppReplaceResponseDNSType = "CNAME"
	SpectrumAppReplaceResponseDNSTypeAddress SpectrumAppReplaceResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by [SpectrumAppReplaceResponseEdgeIPsObject] or
// [SpectrumAppReplaceResponseEdgeIPsObject].
type SpectrumAppReplaceResponseEdgeIPs interface {
	implementsSpectrumAppReplaceResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SpectrumAppReplaceResponseEdgeIPs)(nil)).Elem(), "")
}

type SpectrumAppReplaceResponseEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity SpectrumAppReplaceResponseEdgeIPsObjectConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type SpectrumAppReplaceResponseEdgeIPsObjectType `json:"type"`
	JSON spectrumAppReplaceResponseEdgeIPsObjectJSON `json:"-"`
}

// spectrumAppReplaceResponseEdgeIPsObjectJSON contains the JSON metadata for the
// struct [SpectrumAppReplaceResponseEdgeIPsObject]
type spectrumAppReplaceResponseEdgeIPsObjectJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SpectrumAppReplaceResponseEdgeIPsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SpectrumAppReplaceResponseEdgeIPsObject) implementsSpectrumAppReplaceResponseEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppReplaceResponseEdgeIPsObjectConnectivity string

const (
	SpectrumAppReplaceResponseEdgeIPsObjectConnectivityAll  SpectrumAppReplaceResponseEdgeIPsObjectConnectivity = "all"
	SpectrumAppReplaceResponseEdgeIPsObjectConnectivityIPV4 SpectrumAppReplaceResponseEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppReplaceResponseEdgeIPsObjectConnectivityIPV6 SpectrumAppReplaceResponseEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppReplaceResponseEdgeIPsObjectType string

const (
	SpectrumAppReplaceResponseEdgeIPsObjectTypeDynamic SpectrumAppReplaceResponseEdgeIPsObjectType = "dynamic"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppReplaceResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type SpectrumAppReplaceResponseOriginDNSType `json:"type"`
	JSON spectrumAppReplaceResponseOriginDNSJSON `json:"-"`
}

// spectrumAppReplaceResponseOriginDNSJSON contains the JSON metadata for the
// struct [SpectrumAppReplaceResponseOriginDNS]
type spectrumAppReplaceResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppReplaceResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppReplaceResponseOriginDNSType string

const (
	SpectrumAppReplaceResponseOriginDNSTypeEmpty SpectrumAppReplaceResponseOriginDNSType = ""
	SpectrumAppReplaceResponseOriginDNSTypeA     SpectrumAppReplaceResponseOriginDNSType = "A"
	SpectrumAppReplaceResponseOriginDNSTypeAaaa  SpectrumAppReplaceResponseOriginDNSType = "AAAA"
	SpectrumAppReplaceResponseOriginDNSTypeSrv   SpectrumAppReplaceResponseOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type SpectrumAppReplaceResponseOriginPort interface {
	ImplementsSpectrumAppReplaceResponseOriginPort()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpectrumAppReplaceResponseOriginPort)(nil)).Elem(),
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
type SpectrumAppReplaceResponseProxyProtocol string

const (
	SpectrumAppReplaceResponseProxyProtocolOff    SpectrumAppReplaceResponseProxyProtocol = "off"
	SpectrumAppReplaceResponseProxyProtocolV1     SpectrumAppReplaceResponseProxyProtocol = "v1"
	SpectrumAppReplaceResponseProxyProtocolV2     SpectrumAppReplaceResponseProxyProtocol = "v2"
	SpectrumAppReplaceResponseProxyProtocolSimple SpectrumAppReplaceResponseProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppReplaceResponseTLS string

const (
	SpectrumAppReplaceResponseTLSOff      SpectrumAppReplaceResponseTLS = "off"
	SpectrumAppReplaceResponseTLSFlexible SpectrumAppReplaceResponseTLS = "flexible"
	SpectrumAppReplaceResponseTLSFull     SpectrumAppReplaceResponseTLS = "full"
	SpectrumAppReplaceResponseTLSStrict   SpectrumAppReplaceResponseTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppReplaceResponseTrafficType string

const (
	SpectrumAppReplaceResponseTrafficTypeDirect SpectrumAppReplaceResponseTrafficType = "direct"
	SpectrumAppReplaceResponseTrafficTypeHTTP   SpectrumAppReplaceResponseTrafficType = "http"
	SpectrumAppReplaceResponseTrafficTypeHTTPS  SpectrumAppReplaceResponseTrafficType = "https"
)

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
	SpectrumAppNewParamsDNSTypeCname   SpectrumAppNewParamsDNSType = "CNAME"
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
	SpectrumAppNewParamsOriginDNSTypeAaaa  SpectrumAppNewParamsOriginDNSType = "AAAA"
	SpectrumAppNewParamsOriginDNSTypeSrv   SpectrumAppNewParamsOriginDNSType = "SRV"
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
// Satisfied by [SpectrumAppNewParamsEdgeIPsObject],
// [SpectrumAppNewParamsEdgeIPsObject].
type SpectrumAppNewParamsEdgeIPs interface {
	implementsSpectrumAppNewParamsEdgeIPs()
}

type SpectrumAppNewParamsEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[SpectrumAppNewParamsEdgeIPsObjectConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[SpectrumAppNewParamsEdgeIPsObjectType] `json:"type"`
}

func (r SpectrumAppNewParamsEdgeIPsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppNewParamsEdgeIPsObject) implementsSpectrumAppNewParamsEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppNewParamsEdgeIPsObjectConnectivity string

const (
	SpectrumAppNewParamsEdgeIPsObjectConnectivityAll  SpectrumAppNewParamsEdgeIPsObjectConnectivity = "all"
	SpectrumAppNewParamsEdgeIPsObjectConnectivityIPV4 SpectrumAppNewParamsEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppNewParamsEdgeIPsObjectConnectivityIPV6 SpectrumAppNewParamsEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppNewParamsEdgeIPsObjectType string

const (
	SpectrumAppNewParamsEdgeIPsObjectTypeDynamic SpectrumAppNewParamsEdgeIPsObjectType = "dynamic"
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

// Whether the API call was successful
type SpectrumAppNewResponseEnvelopeSuccess bool

const (
	SpectrumAppNewResponseEnvelopeSuccessTrue SpectrumAppNewResponseEnvelopeSuccess = true
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

// Whether the API call was successful
type SpectrumAppGetResponseEnvelopeSuccess bool

const (
	SpectrumAppGetResponseEnvelopeSuccessTrue SpectrumAppGetResponseEnvelopeSuccess = true
)

type SpectrumAppReplaceParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[SpectrumAppReplaceParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[SpectrumAppReplaceParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[SpectrumAppReplaceParamsOriginPort] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[SpectrumAppReplaceParamsEdgeIPs] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[SpectrumAppReplaceParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[SpectrumAppReplaceParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[SpectrumAppReplaceParamsTrafficType] `json:"traffic_type"`
}

func (r SpectrumAppReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppReplaceParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[SpectrumAppReplaceParamsDNSType] `json:"type"`
}

func (r SpectrumAppReplaceParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type SpectrumAppReplaceParamsDNSType string

const (
	SpectrumAppReplaceParamsDNSTypeCname   SpectrumAppReplaceParamsDNSType = "CNAME"
	SpectrumAppReplaceParamsDNSTypeAddress SpectrumAppReplaceParamsDNSType = "ADDRESS"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppReplaceParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[SpectrumAppReplaceParamsOriginDNSType] `json:"type"`
}

func (r SpectrumAppReplaceParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppReplaceParamsOriginDNSType string

const (
	SpectrumAppReplaceParamsOriginDNSTypeEmpty SpectrumAppReplaceParamsOriginDNSType = ""
	SpectrumAppReplaceParamsOriginDNSTypeA     SpectrumAppReplaceParamsOriginDNSType = "A"
	SpectrumAppReplaceParamsOriginDNSTypeAaaa  SpectrumAppReplaceParamsOriginDNSType = "AAAA"
	SpectrumAppReplaceParamsOriginDNSTypeSrv   SpectrumAppReplaceParamsOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type SpectrumAppReplaceParamsOriginPort interface {
	ImplementsSpectrumAppReplaceParamsOriginPort()
}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [SpectrumAppReplaceParamsEdgeIPsObject],
// [SpectrumAppReplaceParamsEdgeIPsObject].
type SpectrumAppReplaceParamsEdgeIPs interface {
	implementsSpectrumAppReplaceParamsEdgeIPs()
}

type SpectrumAppReplaceParamsEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[SpectrumAppReplaceParamsEdgeIPsObjectConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[SpectrumAppReplaceParamsEdgeIPsObjectType] `json:"type"`
}

func (r SpectrumAppReplaceParamsEdgeIPsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppReplaceParamsEdgeIPsObject) implementsSpectrumAppReplaceParamsEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppReplaceParamsEdgeIPsObjectConnectivity string

const (
	SpectrumAppReplaceParamsEdgeIPsObjectConnectivityAll  SpectrumAppReplaceParamsEdgeIPsObjectConnectivity = "all"
	SpectrumAppReplaceParamsEdgeIPsObjectConnectivityIPV4 SpectrumAppReplaceParamsEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppReplaceParamsEdgeIPsObjectConnectivityIPV6 SpectrumAppReplaceParamsEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppReplaceParamsEdgeIPsObjectType string

const (
	SpectrumAppReplaceParamsEdgeIPsObjectTypeDynamic SpectrumAppReplaceParamsEdgeIPsObjectType = "dynamic"
)

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type SpectrumAppReplaceParamsProxyProtocol string

const (
	SpectrumAppReplaceParamsProxyProtocolOff    SpectrumAppReplaceParamsProxyProtocol = "off"
	SpectrumAppReplaceParamsProxyProtocolV1     SpectrumAppReplaceParamsProxyProtocol = "v1"
	SpectrumAppReplaceParamsProxyProtocolV2     SpectrumAppReplaceParamsProxyProtocol = "v2"
	SpectrumAppReplaceParamsProxyProtocolSimple SpectrumAppReplaceParamsProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppReplaceParamsTLS string

const (
	SpectrumAppReplaceParamsTLSOff      SpectrumAppReplaceParamsTLS = "off"
	SpectrumAppReplaceParamsTLSFlexible SpectrumAppReplaceParamsTLS = "flexible"
	SpectrumAppReplaceParamsTLSFull     SpectrumAppReplaceParamsTLS = "full"
	SpectrumAppReplaceParamsTLSStrict   SpectrumAppReplaceParamsTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppReplaceParamsTrafficType string

const (
	SpectrumAppReplaceParamsTrafficTypeDirect SpectrumAppReplaceParamsTrafficType = "direct"
	SpectrumAppReplaceParamsTrafficTypeHTTP   SpectrumAppReplaceParamsTrafficType = "http"
	SpectrumAppReplaceParamsTrafficTypeHTTPS  SpectrumAppReplaceParamsTrafficType = "https"
)

type SpectrumAppReplaceResponseEnvelope struct {
	Errors   []SpectrumAppReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAppReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   SpectrumAppReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SpectrumAppReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAppReplaceResponseEnvelopeJSON    `json:"-"`
}

// spectrumAppReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [SpectrumAppReplaceResponseEnvelope]
type spectrumAppReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAppReplaceResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    spectrumAppReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAppReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SpectrumAppReplaceResponseEnvelopeErrors]
type spectrumAppReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAppReplaceResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    spectrumAppReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAppReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SpectrumAppReplaceResponseEnvelopeMessages]
type spectrumAppReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SpectrumAppReplaceResponseEnvelopeSuccess bool

const (
	SpectrumAppReplaceResponseEnvelopeSuccessTrue SpectrumAppReplaceResponseEnvelopeSuccess = true
)
