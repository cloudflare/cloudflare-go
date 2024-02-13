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

// Creates a new Spectrum application from a configuration using a name for the
// origin.
func (r *SpectrumAppService) SpectrumApplicationsNewSpectrumApplicationUsingANameForTheOrigin(ctx context.Context, zone string, body SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParams, opts ...option.RequestOption) (res *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *SpectrumAppService) SpectrumApplicationsListSpectrumApplications(ctx context.Context, zone string, query SpectrumAppSpectrumApplicationsListSpectrumApplicationsParams, opts ...option.RequestOption) (res *[]SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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

// The type of DNS record associated with the application.
type SpectrumAppUpdateResponseDNSType string

const (
	SpectrumAppUpdateResponseDNSTypeCname   SpectrumAppUpdateResponseDNSType = "CNAME"
	SpectrumAppUpdateResponseDNSTypeAddress SpectrumAppUpdateResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by [SpectrumAppUpdateResponseEdgeIPsObject] or
// [SpectrumAppUpdateResponseEdgeIPsObject].
type SpectrumAppUpdateResponseEdgeIPs interface {
	implementsSpectrumAppUpdateResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SpectrumAppUpdateResponseEdgeIPs)(nil)).Elem(), "")
}

type SpectrumAppUpdateResponseEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity SpectrumAppUpdateResponseEdgeIPsObjectConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type SpectrumAppUpdateResponseEdgeIPsObjectType `json:"type"`
	JSON spectrumAppUpdateResponseEdgeIPsObjectJSON `json:"-"`
}

// spectrumAppUpdateResponseEdgeIPsObjectJSON contains the JSON metadata for the
// struct [SpectrumAppUpdateResponseEdgeIPsObject]
type spectrumAppUpdateResponseEdgeIPsObjectJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SpectrumAppUpdateResponseEdgeIPsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SpectrumAppUpdateResponseEdgeIPsObject) implementsSpectrumAppUpdateResponseEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppUpdateResponseEdgeIPsObjectConnectivity string

const (
	SpectrumAppUpdateResponseEdgeIPsObjectConnectivityAll  SpectrumAppUpdateResponseEdgeIPsObjectConnectivity = "all"
	SpectrumAppUpdateResponseEdgeIPsObjectConnectivityIPV4 SpectrumAppUpdateResponseEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppUpdateResponseEdgeIPsObjectConnectivityIPV6 SpectrumAppUpdateResponseEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppUpdateResponseEdgeIPsObjectType string

const (
	SpectrumAppUpdateResponseEdgeIPsObjectTypeDynamic SpectrumAppUpdateResponseEdgeIPsObjectType = "dynamic"
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

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppUpdateResponseOriginDNSType string

const (
	SpectrumAppUpdateResponseOriginDNSTypeEmpty SpectrumAppUpdateResponseOriginDNSType = ""
	SpectrumAppUpdateResponseOriginDNSTypeA     SpectrumAppUpdateResponseOriginDNSType = "A"
	SpectrumAppUpdateResponseOriginDNSTypeAaaa  SpectrumAppUpdateResponseOriginDNSType = "AAAA"
	SpectrumAppUpdateResponseOriginDNSTypeSrv   SpectrumAppUpdateResponseOriginDNSType = "SRV"
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
	SpectrumAppUpdateResponseTrafficTypeHTTPs  SpectrumAppUpdateResponseTrafficType = "https"
)

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

type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginPort `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficType `json:"traffic_type"`
	JSON        spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseJSON        `json:"-"`
}

// spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponse]
type spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseJSON struct {
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

func (r *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSType `json:"type"`
	JSON spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNS]
type spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of DNS record associated with the application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSTypeCname   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSType = "CNAME"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSTypeAddress SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObject]
// or
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObject].
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPs interface {
	implementsSpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPs)(nil)).Elem(), "")
}

type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectType `json:"type"`
	JSON spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObject]
type spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObject) implementsSpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivity string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivityAll  SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivity = "all"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivityIPV4 SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivityIPV6 SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectTypeDynamic SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEdgeIPsObjectType = "dynamic"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSType `json:"type"`
	JSON spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNS]
type spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSTypeEmpty SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSType = ""
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSTypeA     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSType = "A"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSTypeAaaa  SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSType = "AAAA"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSTypeSrv   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginPort interface {
	ImplementsSpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginPort()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseOriginPort)(nil)).Elem(),
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
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocol string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocolOff    SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocol = "off"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocolV1     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocol = "v1"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocolV2     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocol = "v2"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocolSimple SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLS string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLSOff      SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLS = "off"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLSFlexible SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLS = "flexible"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLSFull     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLS = "full"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLSStrict   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficTypeDirect SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficType = "direct"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficTypeHTTP   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficType = "http"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficTypeHTTPs  SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseTrafficType = "https"
)

type SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponse = interface{}

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
	SpectrumAppUpdateParamsDNSTypeCname   SpectrumAppUpdateParamsDNSType = "CNAME"
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
	SpectrumAppUpdateParamsOriginDNSTypeAaaa  SpectrumAppUpdateParamsOriginDNSType = "AAAA"
	SpectrumAppUpdateParamsOriginDNSTypeSrv   SpectrumAppUpdateParamsOriginDNSType = "SRV"
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
// Satisfied by [SpectrumAppUpdateParamsEdgeIPsObject],
// [SpectrumAppUpdateParamsEdgeIPsObject].
type SpectrumAppUpdateParamsEdgeIPs interface {
	implementsSpectrumAppUpdateParamsEdgeIPs()
}

type SpectrumAppUpdateParamsEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[SpectrumAppUpdateParamsEdgeIPsObjectConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[SpectrumAppUpdateParamsEdgeIPsObjectType] `json:"type"`
}

func (r SpectrumAppUpdateParamsEdgeIPsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppUpdateParamsEdgeIPsObject) implementsSpectrumAppUpdateParamsEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppUpdateParamsEdgeIPsObjectConnectivity string

const (
	SpectrumAppUpdateParamsEdgeIPsObjectConnectivityAll  SpectrumAppUpdateParamsEdgeIPsObjectConnectivity = "all"
	SpectrumAppUpdateParamsEdgeIPsObjectConnectivityIPV4 SpectrumAppUpdateParamsEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppUpdateParamsEdgeIPsObjectConnectivityIPV6 SpectrumAppUpdateParamsEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppUpdateParamsEdgeIPsObjectType string

const (
	SpectrumAppUpdateParamsEdgeIPsObjectTypeDynamic SpectrumAppUpdateParamsEdgeIPsObjectType = "dynamic"
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
	SpectrumAppUpdateParamsTrafficTypeHTTPs  SpectrumAppUpdateParamsTrafficType = "https"
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

// Whether the API call was successful
type SpectrumAppUpdateResponseEnvelopeSuccess bool

const (
	SpectrumAppUpdateResponseEnvelopeSuccessTrue SpectrumAppUpdateResponseEnvelopeSuccess = true
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

type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginPort] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPs] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficType] `json:"traffic_type"`
}

func (r SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNSType] `json:"type"`
}

func (r SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNSType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNSTypeCname   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNSType = "CNAME"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNSTypeAddress SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsDNSType = "ADDRESS"
)

// The name and type of DNS record for the Spectrum application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSType] `json:"type"`
}

func (r SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSTypeEmpty SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSType = ""
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSTypeA     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSType = "A"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSTypeAaaa  SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSType = "AAAA"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSTypeSrv   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginPort interface {
	ImplementsSpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsOriginPort()
}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObject],
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObject].
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPs interface {
	implementsSpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPs()
}

type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectType] `json:"type"`
}

func (r SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObject) implementsSpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivity string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivityAll  SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivity = "all"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivityIPV4 SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivity = "ipv4"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivityIPV6 SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectTypeDynamic SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsEdgeIPsObjectType = "dynamic"
)

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocol string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocolOff    SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocol = "off"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocolV1     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocol = "v1"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocolV2     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocol = "v2"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocolSimple SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLS string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLSOff      SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLS = "off"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLSFlexible SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLS = "flexible"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLSFull     SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLS = "full"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLSStrict   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficType string

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficTypeDirect SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficType = "direct"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficTypeHTTP   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficType = "http"
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficTypeHTTPs  SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginParamsTrafficType = "https"
)

type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelope struct {
	Errors   []SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeMessages `json:"messages,required"`
	Result   SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeJSON    `json:"-"`
}

// spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelope]
type spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeErrors struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeErrors]
type spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeMessages struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeMessages]
type spectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeSuccess bool

const (
	SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeSuccessTrue SpectrumAppSpectrumApplicationsNewSpectrumApplicationUsingANameForTheOriginResponseEnvelopeSuccess = true
)

type SpectrumAppSpectrumApplicationsListSpectrumApplicationsParams struct {
	// Sets the direction by which results are ordered.
	Direction param.Field[SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsDirection] `query:"direction"`
	// Application field by which results are ordered.
	Order param.Field[SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrder] `query:"order"`
	// Page number of paginated results. This parameter is required in order to use
	// other pagination parameters. If included in the query, `result_info` will be
	// present in the response.
	Page param.Field[float64] `query:"page"`
	// Sets the maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [SpectrumAppSpectrumApplicationsListSpectrumApplicationsParams]'s query
// parameters as `url.Values`.
func (r SpectrumAppSpectrumApplicationsListSpectrumApplicationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sets the direction by which results are ordered.
type SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsDirection string

const (
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsDirectionAsc  SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsDirection = "asc"
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsDirectionDesc SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsDirection = "desc"
)

// Application field by which results are ordered.
type SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrder string

const (
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrderProtocol   SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrder = "protocol"
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrderAppID      SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrder = "app_id"
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrderCreatedOn  SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrder = "created_on"
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrderModifiedOn SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrder = "modified_on"
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrderDNS        SpectrumAppSpectrumApplicationsListSpectrumApplicationsParamsOrder = "dns"
)

type SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelope struct {
	Errors   []SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeMessages `json:"messages,required"`
	Result   []SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeJSON       `json:"-"`
}

// spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelope]
type spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeErrors]
type spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeMessages]
type spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeSuccess bool

const (
	SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeSuccessTrue SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeSuccess = true
)

type SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                               `json:"total_count"`
	JSON       spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeResultInfoJSON `json:"-"`
}

// spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeResultInfo]
type spectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAppSpectrumApplicationsListSpectrumApplicationsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
