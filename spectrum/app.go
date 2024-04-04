// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AppService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Creates a new Spectrum application from a configuration using a name for the
// origin.
func (r *AppService) New(ctx context.Context, zone string, body AppNewParams, opts ...option.RequestOption) (res *AppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AppNewResponseEnvelope
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
func (r *AppService) Update(ctx context.Context, zone string, appID string, body AppUpdateParams, opts ...option.RequestOption) (res *AppUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AppUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *AppService) List(ctx context.Context, zone string, query AppListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AppListResponse], err error) {
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
func (r *AppService) ListAutoPaging(ctx context.Context, zone string, query AppListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AppListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zone, query, opts...))
}

// Deletes a previously existing application.
func (r *AppService) Delete(ctx context.Context, zone string, appID string, body AppDeleteParams, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AppDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the application configuration of a specific application inside a zone.
func (r *AppService) Get(ctx context.Context, zone string, appID string, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef173Union, err error) {
	opts = append(r.Options[:], opts...)
	var env AppGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AppNewResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS AppNewResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs AppNewResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS AppNewResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort AppNewResponseOriginPortUnion `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppNewResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS AppNewResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppNewResponseTrafficType `json:"traffic_type"`
	JSON        appNewResponseJSON        `json:"-"`
}

// appNewResponseJSON contains the JSON metadata for the struct [AppNewResponse]
type appNewResponseJSON struct {
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

func (r *AppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseJSON) RawJSON() string {
	return r.raw
}

// The name and type of DNS record for the Spectrum application.
type AppNewResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type AppNewResponseDNSType `json:"type"`
	JSON appNewResponseDNSJSON `json:"-"`
}

// appNewResponseDNSJSON contains the JSON metadata for the struct
// [AppNewResponseDNS]
type appNewResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type AppNewResponseDNSType string

const (
	AppNewResponseDNSTypeCNAME   AppNewResponseDNSType = "CNAME"
	AppNewResponseDNSTypeAddress AppNewResponseDNSType = "ADDRESS"
)

func (r AppNewResponseDNSType) IsKnown() bool {
	switch r {
	case AppNewResponseDNSTypeCNAME, AppNewResponseDNSTypeAddress:
		return true
	}
	return false
}

// The anycast edge IP configuration for the hostname of this application.
type AppNewResponseEdgeIPs struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppNewResponseEdgeIPsConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type  AppNewResponseEdgeIPsType `json:"type"`
	IPs   interface{}               `json:"ips,required"`
	JSON  appNewResponseEdgeIPsJSON `json:"-"`
	union AppNewResponseEdgeIPsUnion
}

// appNewResponseEdgeIPsJSON contains the JSON metadata for the struct
// [AppNewResponseEdgeIPs]
type appNewResponseEdgeIPsJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	IPs          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r appNewResponseEdgeIPsJSON) RawJSON() string {
	return r.raw
}

func (r *AppNewResponseEdgeIPs) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r AppNewResponseEdgeIPs) AsUnion() AppNewResponseEdgeIPsUnion {
	return r.union
}

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by [spectrum.AppNewResponseEdgeIPsObject] or
// [spectrum.AppNewResponseEdgeIPsObject].
type AppNewResponseEdgeIPsUnion interface {
	implementsSpectrumAppNewResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppNewResponseEdgeIPsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppNewResponseEdgeIPsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppNewResponseEdgeIPsObject{}),
		},
	)
}

type AppNewResponseEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppNewResponseEdgeIPsObjectConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type AppNewResponseEdgeIPsObjectType `json:"type"`
	JSON appNewResponseEdgeIPsObjectJSON `json:"-"`
}

// appNewResponseEdgeIPsObjectJSON contains the JSON metadata for the struct
// [AppNewResponseEdgeIPsObject]
type appNewResponseEdgeIPsObjectJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppNewResponseEdgeIPsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEdgeIPsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppNewResponseEdgeIPsObject) implementsSpectrumAppNewResponseEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppNewResponseEdgeIPsObjectConnectivity string

const (
	AppNewResponseEdgeIPsObjectConnectivityAll  AppNewResponseEdgeIPsObjectConnectivity = "all"
	AppNewResponseEdgeIPsObjectConnectivityIPV4 AppNewResponseEdgeIPsObjectConnectivity = "ipv4"
	AppNewResponseEdgeIPsObjectConnectivityIPV6 AppNewResponseEdgeIPsObjectConnectivity = "ipv6"
)

func (r AppNewResponseEdgeIPsObjectConnectivity) IsKnown() bool {
	switch r {
	case AppNewResponseEdgeIPsObjectConnectivityAll, AppNewResponseEdgeIPsObjectConnectivityIPV4, AppNewResponseEdgeIPsObjectConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppNewResponseEdgeIPsObjectType string

const (
	AppNewResponseEdgeIPsObjectTypeDynamic AppNewResponseEdgeIPsObjectType = "dynamic"
)

func (r AppNewResponseEdgeIPsObjectType) IsKnown() bool {
	switch r {
	case AppNewResponseEdgeIPsObjectTypeDynamic:
		return true
	}
	return false
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppNewResponseEdgeIPsConnectivity string

const (
	AppNewResponseEdgeIPsConnectivityAll  AppNewResponseEdgeIPsConnectivity = "all"
	AppNewResponseEdgeIPsConnectivityIPV4 AppNewResponseEdgeIPsConnectivity = "ipv4"
	AppNewResponseEdgeIPsConnectivityIPV6 AppNewResponseEdgeIPsConnectivity = "ipv6"
)

func (r AppNewResponseEdgeIPsConnectivity) IsKnown() bool {
	switch r {
	case AppNewResponseEdgeIPsConnectivityAll, AppNewResponseEdgeIPsConnectivityIPV4, AppNewResponseEdgeIPsConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppNewResponseEdgeIPsType string

const (
	AppNewResponseEdgeIPsTypeDynamic AppNewResponseEdgeIPsType = "dynamic"
	AppNewResponseEdgeIPsTypeStatic  AppNewResponseEdgeIPsType = "static"
)

func (r AppNewResponseEdgeIPsType) IsKnown() bool {
	switch r {
	case AppNewResponseEdgeIPsTypeDynamic, AppNewResponseEdgeIPsTypeStatic:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type AppNewResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type AppNewResponseOriginDNSType `json:"type"`
	JSON appNewResponseOriginDNSJSON `json:"-"`
}

// appNewResponseOriginDNSJSON contains the JSON metadata for the struct
// [AppNewResponseOriginDNS]
type appNewResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseOriginDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppNewResponseOriginDNSType string

const (
	AppNewResponseOriginDNSTypeEmpty AppNewResponseOriginDNSType = ""
	AppNewResponseOriginDNSTypeA     AppNewResponseOriginDNSType = "A"
	AppNewResponseOriginDNSTypeAAAA  AppNewResponseOriginDNSType = "AAAA"
	AppNewResponseOriginDNSTypeSRV   AppNewResponseOriginDNSType = "SRV"
)

func (r AppNewResponseOriginDNSType) IsKnown() bool {
	switch r {
	case AppNewResponseOriginDNSTypeEmpty, AppNewResponseOriginDNSTypeA, AppNewResponseOriginDNSTypeAAAA, AppNewResponseOriginDNSTypeSRV:
		return true
	}
	return false
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type AppNewResponseOriginPortUnion interface {
	ImplementsSpectrumAppNewResponseOriginPortUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppNewResponseOriginPortUnion)(nil)).Elem(),
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
type AppNewResponseProxyProtocol string

const (
	AppNewResponseProxyProtocolOff    AppNewResponseProxyProtocol = "off"
	AppNewResponseProxyProtocolV1     AppNewResponseProxyProtocol = "v1"
	AppNewResponseProxyProtocolV2     AppNewResponseProxyProtocol = "v2"
	AppNewResponseProxyProtocolSimple AppNewResponseProxyProtocol = "simple"
)

func (r AppNewResponseProxyProtocol) IsKnown() bool {
	switch r {
	case AppNewResponseProxyProtocolOff, AppNewResponseProxyProtocolV1, AppNewResponseProxyProtocolV2, AppNewResponseProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppNewResponseTLS string

const (
	AppNewResponseTLSOff      AppNewResponseTLS = "off"
	AppNewResponseTLSFlexible AppNewResponseTLS = "flexible"
	AppNewResponseTLSFull     AppNewResponseTLS = "full"
	AppNewResponseTLSStrict   AppNewResponseTLS = "strict"
)

func (r AppNewResponseTLS) IsKnown() bool {
	switch r {
	case AppNewResponseTLSOff, AppNewResponseTLSFlexible, AppNewResponseTLSFull, AppNewResponseTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewResponseTrafficType string

const (
	AppNewResponseTrafficTypeDirect AppNewResponseTrafficType = "direct"
	AppNewResponseTrafficTypeHTTP   AppNewResponseTrafficType = "http"
	AppNewResponseTrafficTypeHTTPS  AppNewResponseTrafficType = "https"
)

func (r AppNewResponseTrafficType) IsKnown() bool {
	switch r {
	case AppNewResponseTrafficTypeDirect, AppNewResponseTrafficTypeHTTP, AppNewResponseTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppUpdateResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS AppUpdateResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs AppUpdateResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS AppUpdateResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort AppUpdateResponseOriginPortUnion `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppUpdateResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS AppUpdateResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppUpdateResponseTrafficType `json:"traffic_type"`
	JSON        appUpdateResponseJSON        `json:"-"`
}

// appUpdateResponseJSON contains the JSON metadata for the struct
// [AppUpdateResponse]
type appUpdateResponseJSON struct {
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

func (r *AppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The name and type of DNS record for the Spectrum application.
type AppUpdateResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type AppUpdateResponseDNSType `json:"type"`
	JSON appUpdateResponseDNSJSON `json:"-"`
}

// appUpdateResponseDNSJSON contains the JSON metadata for the struct
// [AppUpdateResponseDNS]
type appUpdateResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type AppUpdateResponseDNSType string

const (
	AppUpdateResponseDNSTypeCNAME   AppUpdateResponseDNSType = "CNAME"
	AppUpdateResponseDNSTypeAddress AppUpdateResponseDNSType = "ADDRESS"
)

func (r AppUpdateResponseDNSType) IsKnown() bool {
	switch r {
	case AppUpdateResponseDNSTypeCNAME, AppUpdateResponseDNSTypeAddress:
		return true
	}
	return false
}

// The anycast edge IP configuration for the hostname of this application.
type AppUpdateResponseEdgeIPs struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppUpdateResponseEdgeIPsConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type  AppUpdateResponseEdgeIPsType `json:"type"`
	IPs   interface{}                  `json:"ips,required"`
	JSON  appUpdateResponseEdgeIPsJSON `json:"-"`
	union AppUpdateResponseEdgeIPsUnion
}

// appUpdateResponseEdgeIPsJSON contains the JSON metadata for the struct
// [AppUpdateResponseEdgeIPs]
type appUpdateResponseEdgeIPsJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	IPs          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r appUpdateResponseEdgeIPsJSON) RawJSON() string {
	return r.raw
}

func (r *AppUpdateResponseEdgeIPs) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r AppUpdateResponseEdgeIPs) AsUnion() AppUpdateResponseEdgeIPsUnion {
	return r.union
}

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by [spectrum.AppUpdateResponseEdgeIPsObject] or
// [spectrum.AppUpdateResponseEdgeIPsObject].
type AppUpdateResponseEdgeIPsUnion interface {
	implementsSpectrumAppUpdateResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppUpdateResponseEdgeIPsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppUpdateResponseEdgeIPsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppUpdateResponseEdgeIPsObject{}),
		},
	)
}

type AppUpdateResponseEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppUpdateResponseEdgeIPsObjectConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type AppUpdateResponseEdgeIPsObjectType `json:"type"`
	JSON appUpdateResponseEdgeIPsObjectJSON `json:"-"`
}

// appUpdateResponseEdgeIPsObjectJSON contains the JSON metadata for the struct
// [AppUpdateResponseEdgeIPsObject]
type appUpdateResponseEdgeIPsObjectJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppUpdateResponseEdgeIPsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEdgeIPsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppUpdateResponseEdgeIPsObject) implementsSpectrumAppUpdateResponseEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppUpdateResponseEdgeIPsObjectConnectivity string

const (
	AppUpdateResponseEdgeIPsObjectConnectivityAll  AppUpdateResponseEdgeIPsObjectConnectivity = "all"
	AppUpdateResponseEdgeIPsObjectConnectivityIPV4 AppUpdateResponseEdgeIPsObjectConnectivity = "ipv4"
	AppUpdateResponseEdgeIPsObjectConnectivityIPV6 AppUpdateResponseEdgeIPsObjectConnectivity = "ipv6"
)

func (r AppUpdateResponseEdgeIPsObjectConnectivity) IsKnown() bool {
	switch r {
	case AppUpdateResponseEdgeIPsObjectConnectivityAll, AppUpdateResponseEdgeIPsObjectConnectivityIPV4, AppUpdateResponseEdgeIPsObjectConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppUpdateResponseEdgeIPsObjectType string

const (
	AppUpdateResponseEdgeIPsObjectTypeDynamic AppUpdateResponseEdgeIPsObjectType = "dynamic"
)

func (r AppUpdateResponseEdgeIPsObjectType) IsKnown() bool {
	switch r {
	case AppUpdateResponseEdgeIPsObjectTypeDynamic:
		return true
	}
	return false
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppUpdateResponseEdgeIPsConnectivity string

const (
	AppUpdateResponseEdgeIPsConnectivityAll  AppUpdateResponseEdgeIPsConnectivity = "all"
	AppUpdateResponseEdgeIPsConnectivityIPV4 AppUpdateResponseEdgeIPsConnectivity = "ipv4"
	AppUpdateResponseEdgeIPsConnectivityIPV6 AppUpdateResponseEdgeIPsConnectivity = "ipv6"
)

func (r AppUpdateResponseEdgeIPsConnectivity) IsKnown() bool {
	switch r {
	case AppUpdateResponseEdgeIPsConnectivityAll, AppUpdateResponseEdgeIPsConnectivityIPV4, AppUpdateResponseEdgeIPsConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppUpdateResponseEdgeIPsType string

const (
	AppUpdateResponseEdgeIPsTypeDynamic AppUpdateResponseEdgeIPsType = "dynamic"
	AppUpdateResponseEdgeIPsTypeStatic  AppUpdateResponseEdgeIPsType = "static"
)

func (r AppUpdateResponseEdgeIPsType) IsKnown() bool {
	switch r {
	case AppUpdateResponseEdgeIPsTypeDynamic, AppUpdateResponseEdgeIPsTypeStatic:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type AppUpdateResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type AppUpdateResponseOriginDNSType `json:"type"`
	JSON appUpdateResponseOriginDNSJSON `json:"-"`
}

// appUpdateResponseOriginDNSJSON contains the JSON metadata for the struct
// [AppUpdateResponseOriginDNS]
type appUpdateResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseOriginDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppUpdateResponseOriginDNSType string

const (
	AppUpdateResponseOriginDNSTypeEmpty AppUpdateResponseOriginDNSType = ""
	AppUpdateResponseOriginDNSTypeA     AppUpdateResponseOriginDNSType = "A"
	AppUpdateResponseOriginDNSTypeAAAA  AppUpdateResponseOriginDNSType = "AAAA"
	AppUpdateResponseOriginDNSTypeSRV   AppUpdateResponseOriginDNSType = "SRV"
)

func (r AppUpdateResponseOriginDNSType) IsKnown() bool {
	switch r {
	case AppUpdateResponseOriginDNSTypeEmpty, AppUpdateResponseOriginDNSTypeA, AppUpdateResponseOriginDNSTypeAAAA, AppUpdateResponseOriginDNSTypeSRV:
		return true
	}
	return false
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type AppUpdateResponseOriginPortUnion interface {
	ImplementsSpectrumAppUpdateResponseOriginPortUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppUpdateResponseOriginPortUnion)(nil)).Elem(),
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
type AppUpdateResponseProxyProtocol string

const (
	AppUpdateResponseProxyProtocolOff    AppUpdateResponseProxyProtocol = "off"
	AppUpdateResponseProxyProtocolV1     AppUpdateResponseProxyProtocol = "v1"
	AppUpdateResponseProxyProtocolV2     AppUpdateResponseProxyProtocol = "v2"
	AppUpdateResponseProxyProtocolSimple AppUpdateResponseProxyProtocol = "simple"
)

func (r AppUpdateResponseProxyProtocol) IsKnown() bool {
	switch r {
	case AppUpdateResponseProxyProtocolOff, AppUpdateResponseProxyProtocolV1, AppUpdateResponseProxyProtocolV2, AppUpdateResponseProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppUpdateResponseTLS string

const (
	AppUpdateResponseTLSOff      AppUpdateResponseTLS = "off"
	AppUpdateResponseTLSFlexible AppUpdateResponseTLS = "flexible"
	AppUpdateResponseTLSFull     AppUpdateResponseTLS = "full"
	AppUpdateResponseTLSStrict   AppUpdateResponseTLS = "strict"
)

func (r AppUpdateResponseTLS) IsKnown() bool {
	switch r {
	case AppUpdateResponseTLSOff, AppUpdateResponseTLSFlexible, AppUpdateResponseTLSFull, AppUpdateResponseTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateResponseTrafficType string

const (
	AppUpdateResponseTrafficTypeDirect AppUpdateResponseTrafficType = "direct"
	AppUpdateResponseTrafficTypeHTTP   AppUpdateResponseTrafficType = "http"
	AppUpdateResponseTrafficTypeHTTPS  AppUpdateResponseTrafficType = "https"
)

func (r AppUpdateResponseTrafficType) IsKnown() bool {
	switch r {
	case AppUpdateResponseTrafficTypeDirect, AppUpdateResponseTrafficTypeHTTP, AppUpdateResponseTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppListResponse = interface{}

type AppDeleteResponse struct {
	// Application identifier.
	ID   string                `json:"id"`
	JSON appDeleteResponseJSON `json:"-"`
}

// appDeleteResponseJSON contains the JSON metadata for the struct
// [AppDeleteResponse]
type appDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AppNewParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[AppNewParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[AppNewParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[AppNewParamsOriginPortUnion] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[AppNewParamsEdgeIPsUnion] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppNewParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppNewParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppNewParamsTrafficType] `json:"traffic_type"`
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type AppNewParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[AppNewParamsDNSType] `json:"type"`
}

func (r AppNewParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type AppNewParamsDNSType string

const (
	AppNewParamsDNSTypeCNAME   AppNewParamsDNSType = "CNAME"
	AppNewParamsDNSTypeAddress AppNewParamsDNSType = "ADDRESS"
)

func (r AppNewParamsDNSType) IsKnown() bool {
	switch r {
	case AppNewParamsDNSTypeCNAME, AppNewParamsDNSTypeAddress:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type AppNewParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[AppNewParamsOriginDNSType] `json:"type"`
}

func (r AppNewParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppNewParamsOriginDNSType string

const (
	AppNewParamsOriginDNSTypeEmpty AppNewParamsOriginDNSType = ""
	AppNewParamsOriginDNSTypeA     AppNewParamsOriginDNSType = "A"
	AppNewParamsOriginDNSTypeAAAA  AppNewParamsOriginDNSType = "AAAA"
	AppNewParamsOriginDNSTypeSRV   AppNewParamsOriginDNSType = "SRV"
)

func (r AppNewParamsOriginDNSType) IsKnown() bool {
	switch r {
	case AppNewParamsOriginDNSTypeEmpty, AppNewParamsOriginDNSTypeA, AppNewParamsOriginDNSTypeAAAA, AppNewParamsOriginDNSTypeSRV:
		return true
	}
	return false
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type AppNewParamsOriginPortUnion interface {
	ImplementsSpectrumAppNewParamsOriginPortUnion()
}

// The anycast edge IP configuration for the hostname of this application.
type AppNewParamsEdgeIPs struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppNewParamsEdgeIPsConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppNewParamsEdgeIPsType] `json:"type"`
	IPs  param.Field[interface{}]             `json:"ips,required"`
}

func (r AppNewParamsEdgeIPs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppNewParamsEdgeIPs) implementsSpectrumAppNewParamsEdgeIPsUnion() {}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [spectrum.AppNewParamsEdgeIPsObject],
// [spectrum.AppNewParamsEdgeIPsObject], [AppNewParamsEdgeIPs].
type AppNewParamsEdgeIPsUnion interface {
	implementsSpectrumAppNewParamsEdgeIPsUnion()
}

type AppNewParamsEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppNewParamsEdgeIPsObjectConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppNewParamsEdgeIPsObjectType] `json:"type"`
}

func (r AppNewParamsEdgeIPsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppNewParamsEdgeIPsObject) implementsSpectrumAppNewParamsEdgeIPsUnion() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppNewParamsEdgeIPsObjectConnectivity string

const (
	AppNewParamsEdgeIPsObjectConnectivityAll  AppNewParamsEdgeIPsObjectConnectivity = "all"
	AppNewParamsEdgeIPsObjectConnectivityIPV4 AppNewParamsEdgeIPsObjectConnectivity = "ipv4"
	AppNewParamsEdgeIPsObjectConnectivityIPV6 AppNewParamsEdgeIPsObjectConnectivity = "ipv6"
)

func (r AppNewParamsEdgeIPsObjectConnectivity) IsKnown() bool {
	switch r {
	case AppNewParamsEdgeIPsObjectConnectivityAll, AppNewParamsEdgeIPsObjectConnectivityIPV4, AppNewParamsEdgeIPsObjectConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppNewParamsEdgeIPsObjectType string

const (
	AppNewParamsEdgeIPsObjectTypeDynamic AppNewParamsEdgeIPsObjectType = "dynamic"
)

func (r AppNewParamsEdgeIPsObjectType) IsKnown() bool {
	switch r {
	case AppNewParamsEdgeIPsObjectTypeDynamic:
		return true
	}
	return false
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppNewParamsEdgeIPsConnectivity string

const (
	AppNewParamsEdgeIPsConnectivityAll  AppNewParamsEdgeIPsConnectivity = "all"
	AppNewParamsEdgeIPsConnectivityIPV4 AppNewParamsEdgeIPsConnectivity = "ipv4"
	AppNewParamsEdgeIPsConnectivityIPV6 AppNewParamsEdgeIPsConnectivity = "ipv6"
)

func (r AppNewParamsEdgeIPsConnectivity) IsKnown() bool {
	switch r {
	case AppNewParamsEdgeIPsConnectivityAll, AppNewParamsEdgeIPsConnectivityIPV4, AppNewParamsEdgeIPsConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppNewParamsEdgeIPsType string

const (
	AppNewParamsEdgeIPsTypeDynamic AppNewParamsEdgeIPsType = "dynamic"
	AppNewParamsEdgeIPsTypeStatic  AppNewParamsEdgeIPsType = "static"
)

func (r AppNewParamsEdgeIPsType) IsKnown() bool {
	switch r {
	case AppNewParamsEdgeIPsTypeDynamic, AppNewParamsEdgeIPsTypeStatic:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppNewParamsProxyProtocol string

const (
	AppNewParamsProxyProtocolOff    AppNewParamsProxyProtocol = "off"
	AppNewParamsProxyProtocolV1     AppNewParamsProxyProtocol = "v1"
	AppNewParamsProxyProtocolV2     AppNewParamsProxyProtocol = "v2"
	AppNewParamsProxyProtocolSimple AppNewParamsProxyProtocol = "simple"
)

func (r AppNewParamsProxyProtocol) IsKnown() bool {
	switch r {
	case AppNewParamsProxyProtocolOff, AppNewParamsProxyProtocolV1, AppNewParamsProxyProtocolV2, AppNewParamsProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppNewParamsTLS string

const (
	AppNewParamsTLSOff      AppNewParamsTLS = "off"
	AppNewParamsTLSFlexible AppNewParamsTLS = "flexible"
	AppNewParamsTLSFull     AppNewParamsTLS = "full"
	AppNewParamsTLSStrict   AppNewParamsTLS = "strict"
)

func (r AppNewParamsTLS) IsKnown() bool {
	switch r {
	case AppNewParamsTLSOff, AppNewParamsTLSFlexible, AppNewParamsTLSFull, AppNewParamsTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewParamsTrafficType string

const (
	AppNewParamsTrafficTypeDirect AppNewParamsTrafficType = "direct"
	AppNewParamsTrafficTypeHTTP   AppNewParamsTrafficType = "http"
	AppNewParamsTrafficTypeHTTPS  AppNewParamsTrafficType = "https"
)

func (r AppNewParamsTrafficType) IsKnown() bool {
	switch r {
	case AppNewParamsTrafficTypeDirect, AppNewParamsTrafficTypeHTTP, AppNewParamsTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AppNewResponse        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    appNewResponseEnvelopeJSON    `json:"-"`
}

// appNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelope]
type appNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppNewResponseEnvelopeSuccess bool

const (
	AppNewResponseEnvelopeSuccessTrue AppNewResponseEnvelopeSuccess = true
)

func (r AppNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppUpdateParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[AppUpdateParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[AppUpdateParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[AppUpdateParamsOriginPortUnion] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[AppUpdateParamsEdgeIPsUnion] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppUpdateParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppUpdateParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppUpdateParamsTrafficType] `json:"traffic_type"`
}

func (r AppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type AppUpdateParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[AppUpdateParamsDNSType] `json:"type"`
}

func (r AppUpdateParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type AppUpdateParamsDNSType string

const (
	AppUpdateParamsDNSTypeCNAME   AppUpdateParamsDNSType = "CNAME"
	AppUpdateParamsDNSTypeAddress AppUpdateParamsDNSType = "ADDRESS"
)

func (r AppUpdateParamsDNSType) IsKnown() bool {
	switch r {
	case AppUpdateParamsDNSTypeCNAME, AppUpdateParamsDNSTypeAddress:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type AppUpdateParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[AppUpdateParamsOriginDNSType] `json:"type"`
}

func (r AppUpdateParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppUpdateParamsOriginDNSType string

const (
	AppUpdateParamsOriginDNSTypeEmpty AppUpdateParamsOriginDNSType = ""
	AppUpdateParamsOriginDNSTypeA     AppUpdateParamsOriginDNSType = "A"
	AppUpdateParamsOriginDNSTypeAAAA  AppUpdateParamsOriginDNSType = "AAAA"
	AppUpdateParamsOriginDNSTypeSRV   AppUpdateParamsOriginDNSType = "SRV"
)

func (r AppUpdateParamsOriginDNSType) IsKnown() bool {
	switch r {
	case AppUpdateParamsOriginDNSTypeEmpty, AppUpdateParamsOriginDNSTypeA, AppUpdateParamsOriginDNSTypeAAAA, AppUpdateParamsOriginDNSTypeSRV:
		return true
	}
	return false
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type AppUpdateParamsOriginPortUnion interface {
	ImplementsSpectrumAppUpdateParamsOriginPortUnion()
}

// The anycast edge IP configuration for the hostname of this application.
type AppUpdateParamsEdgeIPs struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppUpdateParamsEdgeIPsConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppUpdateParamsEdgeIPsType] `json:"type"`
	IPs  param.Field[interface{}]                `json:"ips,required"`
}

func (r AppUpdateParamsEdgeIPs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppUpdateParamsEdgeIPs) implementsSpectrumAppUpdateParamsEdgeIPsUnion() {}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [spectrum.AppUpdateParamsEdgeIPsObject],
// [spectrum.AppUpdateParamsEdgeIPsObject], [AppUpdateParamsEdgeIPs].
type AppUpdateParamsEdgeIPsUnion interface {
	implementsSpectrumAppUpdateParamsEdgeIPsUnion()
}

type AppUpdateParamsEdgeIPsObject struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppUpdateParamsEdgeIPsObjectConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppUpdateParamsEdgeIPsObjectType] `json:"type"`
}

func (r AppUpdateParamsEdgeIPsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppUpdateParamsEdgeIPsObject) implementsSpectrumAppUpdateParamsEdgeIPsUnion() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppUpdateParamsEdgeIPsObjectConnectivity string

const (
	AppUpdateParamsEdgeIPsObjectConnectivityAll  AppUpdateParamsEdgeIPsObjectConnectivity = "all"
	AppUpdateParamsEdgeIPsObjectConnectivityIPV4 AppUpdateParamsEdgeIPsObjectConnectivity = "ipv4"
	AppUpdateParamsEdgeIPsObjectConnectivityIPV6 AppUpdateParamsEdgeIPsObjectConnectivity = "ipv6"
)

func (r AppUpdateParamsEdgeIPsObjectConnectivity) IsKnown() bool {
	switch r {
	case AppUpdateParamsEdgeIPsObjectConnectivityAll, AppUpdateParamsEdgeIPsObjectConnectivityIPV4, AppUpdateParamsEdgeIPsObjectConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppUpdateParamsEdgeIPsObjectType string

const (
	AppUpdateParamsEdgeIPsObjectTypeDynamic AppUpdateParamsEdgeIPsObjectType = "dynamic"
)

func (r AppUpdateParamsEdgeIPsObjectType) IsKnown() bool {
	switch r {
	case AppUpdateParamsEdgeIPsObjectTypeDynamic:
		return true
	}
	return false
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppUpdateParamsEdgeIPsConnectivity string

const (
	AppUpdateParamsEdgeIPsConnectivityAll  AppUpdateParamsEdgeIPsConnectivity = "all"
	AppUpdateParamsEdgeIPsConnectivityIPV4 AppUpdateParamsEdgeIPsConnectivity = "ipv4"
	AppUpdateParamsEdgeIPsConnectivityIPV6 AppUpdateParamsEdgeIPsConnectivity = "ipv6"
)

func (r AppUpdateParamsEdgeIPsConnectivity) IsKnown() bool {
	switch r {
	case AppUpdateParamsEdgeIPsConnectivityAll, AppUpdateParamsEdgeIPsConnectivityIPV4, AppUpdateParamsEdgeIPsConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppUpdateParamsEdgeIPsType string

const (
	AppUpdateParamsEdgeIPsTypeDynamic AppUpdateParamsEdgeIPsType = "dynamic"
	AppUpdateParamsEdgeIPsTypeStatic  AppUpdateParamsEdgeIPsType = "static"
)

func (r AppUpdateParamsEdgeIPsType) IsKnown() bool {
	switch r {
	case AppUpdateParamsEdgeIPsTypeDynamic, AppUpdateParamsEdgeIPsTypeStatic:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppUpdateParamsProxyProtocol string

const (
	AppUpdateParamsProxyProtocolOff    AppUpdateParamsProxyProtocol = "off"
	AppUpdateParamsProxyProtocolV1     AppUpdateParamsProxyProtocol = "v1"
	AppUpdateParamsProxyProtocolV2     AppUpdateParamsProxyProtocol = "v2"
	AppUpdateParamsProxyProtocolSimple AppUpdateParamsProxyProtocol = "simple"
)

func (r AppUpdateParamsProxyProtocol) IsKnown() bool {
	switch r {
	case AppUpdateParamsProxyProtocolOff, AppUpdateParamsProxyProtocolV1, AppUpdateParamsProxyProtocolV2, AppUpdateParamsProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppUpdateParamsTLS string

const (
	AppUpdateParamsTLSOff      AppUpdateParamsTLS = "off"
	AppUpdateParamsTLSFlexible AppUpdateParamsTLS = "flexible"
	AppUpdateParamsTLSFull     AppUpdateParamsTLS = "full"
	AppUpdateParamsTLSStrict   AppUpdateParamsTLS = "strict"
)

func (r AppUpdateParamsTLS) IsKnown() bool {
	switch r {
	case AppUpdateParamsTLSOff, AppUpdateParamsTLSFlexible, AppUpdateParamsTLSFull, AppUpdateParamsTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateParamsTrafficType string

const (
	AppUpdateParamsTrafficTypeDirect AppUpdateParamsTrafficType = "direct"
	AppUpdateParamsTrafficTypeHTTP   AppUpdateParamsTrafficType = "http"
	AppUpdateParamsTrafficTypeHTTPS  AppUpdateParamsTrafficType = "https"
)

func (r AppUpdateParamsTrafficType) IsKnown() bool {
	switch r {
	case AppUpdateParamsTrafficTypeDirect, AppUpdateParamsTrafficTypeHTTP, AppUpdateParamsTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AppUpdateResponse     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    appUpdateResponseEnvelopeJSON    `json:"-"`
}

// appUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelope]
type appUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppUpdateResponseEnvelopeSuccess bool

const (
	AppUpdateResponseEnvelopeSuccessTrue AppUpdateResponseEnvelopeSuccess = true
)

func (r AppUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppListParams struct {
	// Sets the direction by which results are ordered.
	Direction param.Field[AppListParamsDirection] `query:"direction"`
	// Application field by which results are ordered.
	Order param.Field[AppListParamsOrder] `query:"order"`
	// Page number of paginated results. This parameter is required in order to use
	// other pagination parameters. If included in the query, `result_info` will be
	// present in the response.
	Page param.Field[float64] `query:"page"`
	// Sets the maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AppListParams]'s query parameters as `url.Values`.
func (r AppListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sets the direction by which results are ordered.
type AppListParamsDirection string

const (
	AppListParamsDirectionAsc  AppListParamsDirection = "asc"
	AppListParamsDirectionDesc AppListParamsDirection = "desc"
)

func (r AppListParamsDirection) IsKnown() bool {
	switch r {
	case AppListParamsDirectionAsc, AppListParamsDirectionDesc:
		return true
	}
	return false
}

// Application field by which results are ordered.
type AppListParamsOrder string

const (
	AppListParamsOrderProtocol   AppListParamsOrder = "protocol"
	AppListParamsOrderAppID      AppListParamsOrder = "app_id"
	AppListParamsOrderCreatedOn  AppListParamsOrder = "created_on"
	AppListParamsOrderModifiedOn AppListParamsOrder = "modified_on"
	AppListParamsOrderDNS        AppListParamsOrder = "dns"
)

func (r AppListParamsOrder) IsKnown() bool {
	switch r {
	case AppListParamsOrderProtocol, AppListParamsOrderAppID, AppListParamsOrderCreatedOn, AppListParamsOrderModifiedOn, AppListParamsOrderDNS:
		return true
	}
	return false
}

type AppDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AppDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AppDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AppDeleteResponse     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    appDeleteResponseEnvelopeJSON    `json:"-"`
}

// appDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelope]
type appDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppDeleteResponseEnvelopeSuccess bool

const (
	AppDeleteResponseEnvelopeSuccessTrue AppDeleteResponseEnvelopeSuccess = true
)

func (r AppDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   shared.UnnamedSchemaRef173Union `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    appGetResponseEnvelopeJSON    `json:"-"`
}

// appGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelope]
type appGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppGetResponseEnvelopeSuccess bool

const (
	AppGetResponseEnvelopeSuccessTrue AppGetResponseEnvelopeSuccess = true
)

func (r AppGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
