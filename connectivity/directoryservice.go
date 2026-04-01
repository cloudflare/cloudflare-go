// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// DirectoryServiceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirectoryServiceService] method instead.
type DirectoryServiceService struct {
	Options []option.RequestOption
}

// NewDirectoryServiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDirectoryServiceService(opts ...option.RequestOption) (r *DirectoryServiceService) {
	r = &DirectoryServiceService{}
	r.Options = opts
	return
}

// Create Workers VPC connectivity service
func (r *DirectoryServiceService) New(ctx context.Context, params DirectoryServiceNewParams, opts ...option.RequestOption) (res *DirectoryServiceNewResponse, err error) {
	var env DirectoryServiceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update Workers VPC connectivity service
func (r *DirectoryServiceService) Update(ctx context.Context, serviceID string, params DirectoryServiceUpdateParams, opts ...option.RequestOption) (res *DirectoryServiceUpdateResponse, err error) {
	var env DirectoryServiceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", params.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List Workers VPC connectivity services
func (r *DirectoryServiceService) List(ctx context.Context, params DirectoryServiceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DirectoryServiceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List Workers VPC connectivity services
func (r *DirectoryServiceService) ListAutoPaging(ctx context.Context, params DirectoryServiceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DirectoryServiceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete Workers VPC connectivity service
func (r *DirectoryServiceService) Delete(ctx context.Context, serviceID string, body DirectoryServiceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", body.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get Workers VPC connectivity service
func (r *DirectoryServiceService) Get(ctx context.Context, serviceID string, query DirectoryServiceGetParams, opts ...option.RequestOption) (res *DirectoryServiceGetResponse, err error) {
	var env DirectoryServiceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", query.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DirectoryServiceNewResponse struct {
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseInfraHTTPServiceConfigHost],
	// [DirectoryServiceNewResponseInfraTCPServiceConfigHost].
	Host        interface{}                            `json:"host" api:"required"`
	Name        string                                 `json:"name" api:"required"`
	Type        DirectoryServiceNewResponseType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceNewResponseAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                              `json:"created_at" format:"date-time"`
	HTTPPort    int64                                  `json:"http_port" api:"nullable"`
	HTTPSPort   int64                                  `json:"https_port" api:"nullable"`
	ServiceID   string                                 `json:"service_id" format:"uuid"`
	TCPPort     int64                                  `json:"tcp_port" api:"nullable"`
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseInfraHTTPServiceConfigTLSSettings],
	// [DirectoryServiceNewResponseInfraTCPServiceConfigTLSSettings].
	TLSSettings interface{}                     `json:"tls_settings"`
	UpdatedAt   time.Time                       `json:"updated_at" format:"date-time"`
	JSON        directoryServiceNewResponseJSON `json:"-"`
	union       DirectoryServiceNewResponseUnion
}

// directoryServiceNewResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceNewResponse]
type directoryServiceNewResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r directoryServiceNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceNewResponseInfraHTTPServiceConfig],
// [DirectoryServiceNewResponseInfraTCPServiceConfig].
func (r DirectoryServiceNewResponse) AsUnion() DirectoryServiceNewResponseUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceNewResponseInfraHTTPServiceConfig] or
// [DirectoryServiceNewResponseInfraTCPServiceConfig].
type DirectoryServiceNewResponseUnion interface {
	implementsDirectoryServiceNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceNewResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceNewResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceNewResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "http",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceNewResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceNewResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "http",
		},
	)
}

type DirectoryServiceNewResponseInfraHTTPServiceConfig struct {
	Host      DirectoryServiceNewResponseInfraHTTPServiceConfigHost `json:"host" api:"required"`
	Name      string                                                `json:"name" api:"required"`
	Type      DirectoryServiceNewResponseInfraHTTPServiceConfigType `json:"type" api:"required"`
	CreatedAt time.Time                                             `json:"created_at" format:"date-time"`
	HTTPPort  int64                                                 `json:"http_port" api:"nullable"`
	HTTPSPort int64                                                 `json:"https_port" api:"nullable"`
	ServiceID string                                                `json:"service_id" format:"uuid"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceNewResponseInfraHTTPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON        directoryServiceNewResponseInfraHTTPServiceConfigJSON        `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigJSON contains the JSON metadata
// for the struct [DirectoryServiceNewResponseInfraHTTPServiceConfig]
type directoryServiceNewResponseInfraHTTPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraHTTPServiceConfig) implementsDirectoryServiceNewResponse() {}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                               `json:"resolver_network"`
	JSON            directoryServiceNewResponseInfraHTTPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceNewResponseInfraHTTPServiceConfigHostUnion
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostJSON contains the JSON
// metadata for the struct [DirectoryServiceNewResponseInfraHTTPServiceConfigHost]
type directoryServiceNewResponseInfraHTTPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceNewResponseInfraHTTPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceNewResponseInfraHTTPServiceConfigHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceNewResponseInfraHTTPServiceConfigHost) AsUnion() DirectoryServiceNewResponseInfraHTTPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHost].
type DirectoryServiceNewResponseInfraHTTPServiceConfigHostUnion interface {
	implementsDirectoryServiceNewResponseInfraHTTPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceNewResponseInfraHTTPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                    `json:"ipv4" api:"required"`
	Network DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4Host]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceNewResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                        `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                    `json:"ipv6" api:"required"`
	Network DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6Host]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceNewResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                        `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                         `json:"ipv4" api:"required"`
	IPV6    string                                                                         `json:"ipv6" api:"required"`
	Network DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHost]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceNewResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                             `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                                `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHost]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceNewResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                    `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                  `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraHTTPServiceConfigType string

const (
	DirectoryServiceNewResponseInfraHTTPServiceConfigTypeTCP  DirectoryServiceNewResponseInfraHTTPServiceConfigType = "tcp"
	DirectoryServiceNewResponseInfraHTTPServiceConfigTypeHTTP DirectoryServiceNewResponseInfraHTTPServiceConfigType = "http"
)

func (r DirectoryServiceNewResponseInfraHTTPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewResponseInfraHTTPServiceConfigTypeTCP, DirectoryServiceNewResponseInfraHTTPServiceConfigTypeHTTP:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceNewResponseInfraHTTPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                           `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceNewResponseInfraHTTPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceNewResponseInfraHTTPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceNewResponseInfraHTTPServiceConfigTLSSettings]
type directoryServiceNewResponseInfraHTTPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraHTTPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraHTTPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraTCPServiceConfig struct {
	Host        DirectoryServiceNewResponseInfraTCPServiceConfigHost        `json:"host" api:"required"`
	Name        string                                                      `json:"name" api:"required"`
	Type        DirectoryServiceNewResponseInfraTCPServiceConfigType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                                                   `json:"created_at" format:"date-time"`
	ServiceID   string                                                      `json:"service_id" format:"uuid"`
	TCPPort     int64                                                       `json:"tcp_port" api:"nullable"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceNewResponseInfraTCPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                   `json:"updated_at" format:"date-time"`
	JSON        directoryServiceNewResponseInfraTCPServiceConfigJSON        `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigJSON contains the JSON metadata
// for the struct [DirectoryServiceNewResponseInfraTCPServiceConfig]
type directoryServiceNewResponseInfraTCPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraTCPServiceConfig) implementsDirectoryServiceNewResponse() {}

type DirectoryServiceNewResponseInfraTCPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                              `json:"resolver_network"`
	JSON            directoryServiceNewResponseInfraTCPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceNewResponseInfraTCPServiceConfigHostUnion
}

// directoryServiceNewResponseInfraTCPServiceConfigHostJSON contains the JSON
// metadata for the struct [DirectoryServiceNewResponseInfraTCPServiceConfigHost]
type directoryServiceNewResponseInfraTCPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceNewResponseInfraTCPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceNewResponseInfraTCPServiceConfigHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceNewResponseInfraTCPServiceConfigHost) AsUnion() DirectoryServiceNewResponseInfraTCPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHost].
type DirectoryServiceNewResponseInfraTCPServiceConfigHostUnion interface {
	implementsDirectoryServiceNewResponseInfraTCPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceNewResponseInfraTCPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                   `json:"ipv4" api:"required"`
	Network DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4Host]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceNewResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                       `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                   `json:"ipv6" api:"required"`
	Network DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6Host]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceNewResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                       `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                        `json:"ipv4" api:"required"`
	IPV6    string                                                                        `json:"ipv6" api:"required"`
	Network DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHost]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceNewResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                            `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                               `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHost]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceNewResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                   `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                 `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseInfraTCPServiceConfigType string

const (
	DirectoryServiceNewResponseInfraTCPServiceConfigTypeTCP  DirectoryServiceNewResponseInfraTCPServiceConfigType = "tcp"
	DirectoryServiceNewResponseInfraTCPServiceConfigTypeHTTP DirectoryServiceNewResponseInfraTCPServiceConfigType = "http"
)

func (r DirectoryServiceNewResponseInfraTCPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewResponseInfraTCPServiceConfigTypeTCP, DirectoryServiceNewResponseInfraTCPServiceConfigTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocol string

const (
	DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocolPostgresql DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocol = "postgresql"
	DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocolMysql      DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocol = "mysql"
)

func (r DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocolPostgresql, DirectoryServiceNewResponseInfraTCPServiceConfigAppProtocolMysql:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceNewResponseInfraTCPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                          `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceNewResponseInfraTCPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceNewResponseInfraTCPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceNewResponseInfraTCPServiceConfigTLSSettings]
type directoryServiceNewResponseInfraTCPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseInfraTCPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseInfraTCPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseType string

const (
	DirectoryServiceNewResponseTypeTCP  DirectoryServiceNewResponseType = "tcp"
	DirectoryServiceNewResponseTypeHTTP DirectoryServiceNewResponseType = "http"
)

func (r DirectoryServiceNewResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewResponseTypeTCP, DirectoryServiceNewResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceNewResponseAppProtocol string

const (
	DirectoryServiceNewResponseAppProtocolPostgresql DirectoryServiceNewResponseAppProtocol = "postgresql"
	DirectoryServiceNewResponseAppProtocolMysql      DirectoryServiceNewResponseAppProtocol = "mysql"
)

func (r DirectoryServiceNewResponseAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceNewResponseAppProtocolPostgresql, DirectoryServiceNewResponseAppProtocolMysql:
		return true
	}
	return false
}

type DirectoryServiceUpdateResponse struct {
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHost],
	// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHost].
	Host        interface{}                               `json:"host" api:"required"`
	Name        string                                    `json:"name" api:"required"`
	Type        DirectoryServiceUpdateResponseType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceUpdateResponseAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                                 `json:"created_at" format:"date-time"`
	HTTPPort    int64                                     `json:"http_port" api:"nullable"`
	HTTPSPort   int64                                     `json:"https_port" api:"nullable"`
	ServiceID   string                                    `json:"service_id" format:"uuid"`
	TCPPort     int64                                     `json:"tcp_port" api:"nullable"`
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettings],
	// [DirectoryServiceUpdateResponseInfraTCPServiceConfigTLSSettings].
	TLSSettings interface{}                        `json:"tls_settings"`
	UpdatedAt   time.Time                          `json:"updated_at" format:"date-time"`
	JSON        directoryServiceUpdateResponseJSON `json:"-"`
	union       DirectoryServiceUpdateResponseUnion
}

// directoryServiceUpdateResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponse]
type directoryServiceUpdateResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r directoryServiceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceUpdateResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfig],
// [DirectoryServiceUpdateResponseInfraTCPServiceConfig].
func (r DirectoryServiceUpdateResponse) AsUnion() DirectoryServiceUpdateResponseUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceUpdateResponseInfraHTTPServiceConfig] or
// [DirectoryServiceUpdateResponseInfraTCPServiceConfig].
type DirectoryServiceUpdateResponseUnion interface {
	implementsDirectoryServiceUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceUpdateResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceUpdateResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "http",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceUpdateResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceUpdateResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "http",
		},
	)
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfig struct {
	Host      DirectoryServiceUpdateResponseInfraHTTPServiceConfigHost `json:"host" api:"required"`
	Name      string                                                   `json:"name" api:"required"`
	Type      DirectoryServiceUpdateResponseInfraHTTPServiceConfigType `json:"type" api:"required"`
	CreatedAt time.Time                                                `json:"created_at" format:"date-time"`
	HTTPPort  int64                                                    `json:"http_port" api:"nullable"`
	HTTPSPort int64                                                    `json:"https_port" api:"nullable"`
	ServiceID string                                                   `json:"service_id" format:"uuid"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON        directoryServiceUpdateResponseInfraHTTPServiceConfigJSON        `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseInfraHTTPServiceConfig]
type directoryServiceUpdateResponseInfraHTTPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraHTTPServiceConfig) implementsDirectoryServiceUpdateResponse() {
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                                  `json:"resolver_network"`
	JSON            directoryServiceUpdateResponseInfraHTTPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostUnion
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostJSON contains the JSON
// metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHost]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceUpdateResponseInfraHTTPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceUpdateResponseInfraHTTPServiceConfigHost) AsUnion() DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHost].
type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostUnion interface {
	implementsDirectoryServiceUpdateResponseInfraHTTPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                       `json:"ipv4" api:"required"`
	Network DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4Host]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceUpdateResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                           `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                       `json:"ipv6" api:"required"`
	Network DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6Host]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceUpdateResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                           `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                            `json:"ipv4" api:"required"`
	IPV6    string                                                                            `json:"ipv6" api:"required"`
	Network DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHost]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceUpdateResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                                `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                                   `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHost]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceUpdateResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                       `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                     `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraHTTPServiceConfigType string

const (
	DirectoryServiceUpdateResponseInfraHTTPServiceConfigTypeTCP  DirectoryServiceUpdateResponseInfraHTTPServiceConfigType = "tcp"
	DirectoryServiceUpdateResponseInfraHTTPServiceConfigTypeHTTP DirectoryServiceUpdateResponseInfraHTTPServiceConfigType = "http"
)

func (r DirectoryServiceUpdateResponseInfraHTTPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateResponseInfraHTTPServiceConfigTypeTCP, DirectoryServiceUpdateResponseInfraHTTPServiceConfigTypeHTTP:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                              `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettings]
type directoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraHTTPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfig struct {
	Host        DirectoryServiceUpdateResponseInfraTCPServiceConfigHost        `json:"host" api:"required"`
	Name        string                                                         `json:"name" api:"required"`
	Type        DirectoryServiceUpdateResponseInfraTCPServiceConfigType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                                                      `json:"created_at" format:"date-time"`
	ServiceID   string                                                         `json:"service_id" format:"uuid"`
	TCPPort     int64                                                          `json:"tcp_port" api:"nullable"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceUpdateResponseInfraTCPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON        directoryServiceUpdateResponseInfraTCPServiceConfigJSON        `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseInfraTCPServiceConfig]
type directoryServiceUpdateResponseInfraTCPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraTCPServiceConfig) implementsDirectoryServiceUpdateResponse() {
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                                 `json:"resolver_network"`
	JSON            directoryServiceUpdateResponseInfraTCPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceUpdateResponseInfraTCPServiceConfigHostUnion
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostJSON contains the JSON
// metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHost]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceUpdateResponseInfraTCPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceUpdateResponseInfraTCPServiceConfigHost) AsUnion() DirectoryServiceUpdateResponseInfraTCPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHost].
type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostUnion interface {
	implementsDirectoryServiceUpdateResponseInfraTCPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceUpdateResponseInfraTCPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                      `json:"ipv4" api:"required"`
	Network DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4Host]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceUpdateResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                          `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                      `json:"ipv6" api:"required"`
	Network DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6Host]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceUpdateResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                          `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                           `json:"ipv4" api:"required"`
	IPV6    string                                                                           `json:"ipv6" api:"required"`
	Network DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHost]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceUpdateResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                               `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                                  `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHost]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceUpdateResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                      `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                    `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigType string

const (
	DirectoryServiceUpdateResponseInfraTCPServiceConfigTypeTCP  DirectoryServiceUpdateResponseInfraTCPServiceConfigType = "tcp"
	DirectoryServiceUpdateResponseInfraTCPServiceConfigTypeHTTP DirectoryServiceUpdateResponseInfraTCPServiceConfigType = "http"
)

func (r DirectoryServiceUpdateResponseInfraTCPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateResponseInfraTCPServiceConfigTypeTCP, DirectoryServiceUpdateResponseInfraTCPServiceConfigTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocol string

const (
	DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocolPostgresql DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocol = "postgresql"
	DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocolMysql      DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocol = "mysql"
)

func (r DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocolPostgresql, DirectoryServiceUpdateResponseInfraTCPServiceConfigAppProtocolMysql:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceUpdateResponseInfraTCPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                             `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceUpdateResponseInfraTCPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceUpdateResponseInfraTCPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceUpdateResponseInfraTCPServiceConfigTLSSettings]
type directoryServiceUpdateResponseInfraTCPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseInfraTCPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseInfraTCPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseType string

const (
	DirectoryServiceUpdateResponseTypeTCP  DirectoryServiceUpdateResponseType = "tcp"
	DirectoryServiceUpdateResponseTypeHTTP DirectoryServiceUpdateResponseType = "http"
)

func (r DirectoryServiceUpdateResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateResponseTypeTCP, DirectoryServiceUpdateResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceUpdateResponseAppProtocol string

const (
	DirectoryServiceUpdateResponseAppProtocolPostgresql DirectoryServiceUpdateResponseAppProtocol = "postgresql"
	DirectoryServiceUpdateResponseAppProtocolMysql      DirectoryServiceUpdateResponseAppProtocol = "mysql"
)

func (r DirectoryServiceUpdateResponseAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateResponseAppProtocolPostgresql, DirectoryServiceUpdateResponseAppProtocolMysql:
		return true
	}
	return false
}

type DirectoryServiceListResponse struct {
	// This field can have the runtime type of
	// [DirectoryServiceListResponseInfraHTTPServiceConfigHost],
	// [DirectoryServiceListResponseInfraTCPServiceConfigHost].
	Host        interface{}                             `json:"host" api:"required"`
	Name        string                                  `json:"name" api:"required"`
	Type        DirectoryServiceListResponseType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceListResponseAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                               `json:"created_at" format:"date-time"`
	HTTPPort    int64                                   `json:"http_port" api:"nullable"`
	HTTPSPort   int64                                   `json:"https_port" api:"nullable"`
	ServiceID   string                                  `json:"service_id" format:"uuid"`
	TCPPort     int64                                   `json:"tcp_port" api:"nullable"`
	// This field can have the runtime type of
	// [DirectoryServiceListResponseInfraHTTPServiceConfigTLSSettings],
	// [DirectoryServiceListResponseInfraTCPServiceConfigTLSSettings].
	TLSSettings interface{}                      `json:"tls_settings"`
	UpdatedAt   time.Time                        `json:"updated_at" format:"date-time"`
	JSON        directoryServiceListResponseJSON `json:"-"`
	union       DirectoryServiceListResponseUnion
}

// directoryServiceListResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceListResponse]
type directoryServiceListResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r directoryServiceListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceListResponseInfraHTTPServiceConfig],
// [DirectoryServiceListResponseInfraTCPServiceConfig].
func (r DirectoryServiceListResponse) AsUnion() DirectoryServiceListResponseUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceListResponseInfraHTTPServiceConfig] or
// [DirectoryServiceListResponseInfraTCPServiceConfig].
type DirectoryServiceListResponseUnion interface {
	implementsDirectoryServiceListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceListResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceListResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "http",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceListResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceListResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "http",
		},
	)
}

type DirectoryServiceListResponseInfraHTTPServiceConfig struct {
	Host      DirectoryServiceListResponseInfraHTTPServiceConfigHost `json:"host" api:"required"`
	Name      string                                                 `json:"name" api:"required"`
	Type      DirectoryServiceListResponseInfraHTTPServiceConfigType `json:"type" api:"required"`
	CreatedAt time.Time                                              `json:"created_at" format:"date-time"`
	HTTPPort  int64                                                  `json:"http_port" api:"nullable"`
	HTTPSPort int64                                                  `json:"https_port" api:"nullable"`
	ServiceID string                                                 `json:"service_id" format:"uuid"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceListResponseInfraHTTPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                     `json:"updated_at" format:"date-time"`
	JSON        directoryServiceListResponseInfraHTTPServiceConfigJSON        `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigJSON contains the JSON
// metadata for the struct [DirectoryServiceListResponseInfraHTTPServiceConfig]
type directoryServiceListResponseInfraHTTPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraHTTPServiceConfig) implementsDirectoryServiceListResponse() {
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                                `json:"resolver_network"`
	JSON            directoryServiceListResponseInfraHTTPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceListResponseInfraHTTPServiceConfigHostUnion
}

// directoryServiceListResponseInfraHTTPServiceConfigHostJSON contains the JSON
// metadata for the struct [DirectoryServiceListResponseInfraHTTPServiceConfigHost]
type directoryServiceListResponseInfraHTTPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceListResponseInfraHTTPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceListResponseInfraHTTPServiceConfigHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceListResponseInfraHTTPServiceConfigHost) AsUnion() DirectoryServiceListResponseInfraHTTPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHost].
type DirectoryServiceListResponseInfraHTTPServiceConfigHostUnion interface {
	implementsDirectoryServiceListResponseInfraHTTPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceListResponseInfraHTTPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                     `json:"ipv4" api:"required"`
	Network DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4Host]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceListResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                         `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                     `json:"ipv6" api:"required"`
	Network DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6Host]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceListResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                         `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                          `json:"ipv4" api:"required"`
	IPV6    string                                                                          `json:"ipv6" api:"required"`
	Network DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHost]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceListResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                              `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                                 `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHost]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceListResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                     `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                   `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraHTTPServiceConfigType string

const (
	DirectoryServiceListResponseInfraHTTPServiceConfigTypeTCP  DirectoryServiceListResponseInfraHTTPServiceConfigType = "tcp"
	DirectoryServiceListResponseInfraHTTPServiceConfigTypeHTTP DirectoryServiceListResponseInfraHTTPServiceConfigType = "http"
)

func (r DirectoryServiceListResponseInfraHTTPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceListResponseInfraHTTPServiceConfigTypeTCP, DirectoryServiceListResponseInfraHTTPServiceConfigTypeHTTP:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceListResponseInfraHTTPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                            `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceListResponseInfraHTTPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceListResponseInfraHTTPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceListResponseInfraHTTPServiceConfigTLSSettings]
type directoryServiceListResponseInfraHTTPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraHTTPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraHTTPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraTCPServiceConfig struct {
	Host        DirectoryServiceListResponseInfraTCPServiceConfigHost        `json:"host" api:"required"`
	Name        string                                                       `json:"name" api:"required"`
	Type        DirectoryServiceListResponseInfraTCPServiceConfigType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceListResponseInfraTCPServiceConfigAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                                                    `json:"created_at" format:"date-time"`
	ServiceID   string                                                       `json:"service_id" format:"uuid"`
	TCPPort     int64                                                        `json:"tcp_port" api:"nullable"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceListResponseInfraTCPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON        directoryServiceListResponseInfraTCPServiceConfigJSON        `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigJSON contains the JSON metadata
// for the struct [DirectoryServiceListResponseInfraTCPServiceConfig]
type directoryServiceListResponseInfraTCPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraTCPServiceConfig) implementsDirectoryServiceListResponse() {}

type DirectoryServiceListResponseInfraTCPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                               `json:"resolver_network"`
	JSON            directoryServiceListResponseInfraTCPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceListResponseInfraTCPServiceConfigHostUnion
}

// directoryServiceListResponseInfraTCPServiceConfigHostJSON contains the JSON
// metadata for the struct [DirectoryServiceListResponseInfraTCPServiceConfigHost]
type directoryServiceListResponseInfraTCPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceListResponseInfraTCPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceListResponseInfraTCPServiceConfigHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceListResponseInfraTCPServiceConfigHost) AsUnion() DirectoryServiceListResponseInfraTCPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHost].
type DirectoryServiceListResponseInfraTCPServiceConfigHostUnion interface {
	implementsDirectoryServiceListResponseInfraTCPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceListResponseInfraTCPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                    `json:"ipv4" api:"required"`
	Network DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4Host]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceListResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                        `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                    `json:"ipv6" api:"required"`
	Network DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6Host]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceListResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                        `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                         `json:"ipv4" api:"required"`
	IPV6    string                                                                         `json:"ipv6" api:"required"`
	Network DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHost]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceListResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                             `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                                `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHost]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceListResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                    `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                  `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseInfraTCPServiceConfigType string

const (
	DirectoryServiceListResponseInfraTCPServiceConfigTypeTCP  DirectoryServiceListResponseInfraTCPServiceConfigType = "tcp"
	DirectoryServiceListResponseInfraTCPServiceConfigTypeHTTP DirectoryServiceListResponseInfraTCPServiceConfigType = "http"
)

func (r DirectoryServiceListResponseInfraTCPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceListResponseInfraTCPServiceConfigTypeTCP, DirectoryServiceListResponseInfraTCPServiceConfigTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceListResponseInfraTCPServiceConfigAppProtocol string

const (
	DirectoryServiceListResponseInfraTCPServiceConfigAppProtocolPostgresql DirectoryServiceListResponseInfraTCPServiceConfigAppProtocol = "postgresql"
	DirectoryServiceListResponseInfraTCPServiceConfigAppProtocolMysql      DirectoryServiceListResponseInfraTCPServiceConfigAppProtocol = "mysql"
)

func (r DirectoryServiceListResponseInfraTCPServiceConfigAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceListResponseInfraTCPServiceConfigAppProtocolPostgresql, DirectoryServiceListResponseInfraTCPServiceConfigAppProtocolMysql:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceListResponseInfraTCPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                           `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceListResponseInfraTCPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceListResponseInfraTCPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceListResponseInfraTCPServiceConfigTLSSettings]
type directoryServiceListResponseInfraTCPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceListResponseInfraTCPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseInfraTCPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseType string

const (
	DirectoryServiceListResponseTypeTCP  DirectoryServiceListResponseType = "tcp"
	DirectoryServiceListResponseTypeHTTP DirectoryServiceListResponseType = "http"
)

func (r DirectoryServiceListResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceListResponseTypeTCP, DirectoryServiceListResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceListResponseAppProtocol string

const (
	DirectoryServiceListResponseAppProtocolPostgresql DirectoryServiceListResponseAppProtocol = "postgresql"
	DirectoryServiceListResponseAppProtocolMysql      DirectoryServiceListResponseAppProtocol = "mysql"
)

func (r DirectoryServiceListResponseAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceListResponseAppProtocolPostgresql, DirectoryServiceListResponseAppProtocolMysql:
		return true
	}
	return false
}

type DirectoryServiceGetResponse struct {
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseInfraHTTPServiceConfigHost],
	// [DirectoryServiceGetResponseInfraTCPServiceConfigHost].
	Host        interface{}                            `json:"host" api:"required"`
	Name        string                                 `json:"name" api:"required"`
	Type        DirectoryServiceGetResponseType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceGetResponseAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                              `json:"created_at" format:"date-time"`
	HTTPPort    int64                                  `json:"http_port" api:"nullable"`
	HTTPSPort   int64                                  `json:"https_port" api:"nullable"`
	ServiceID   string                                 `json:"service_id" format:"uuid"`
	TCPPort     int64                                  `json:"tcp_port" api:"nullable"`
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseInfraHTTPServiceConfigTLSSettings],
	// [DirectoryServiceGetResponseInfraTCPServiceConfigTLSSettings].
	TLSSettings interface{}                     `json:"tls_settings"`
	UpdatedAt   time.Time                       `json:"updated_at" format:"date-time"`
	JSON        directoryServiceGetResponseJSON `json:"-"`
	union       DirectoryServiceGetResponseUnion
}

// directoryServiceGetResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceGetResponse]
type directoryServiceGetResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r directoryServiceGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceGetResponseInfraHTTPServiceConfig],
// [DirectoryServiceGetResponseInfraTCPServiceConfig].
func (r DirectoryServiceGetResponse) AsUnion() DirectoryServiceGetResponseUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceGetResponseInfraHTTPServiceConfig] or
// [DirectoryServiceGetResponseInfraTCPServiceConfig].
type DirectoryServiceGetResponseUnion interface {
	implementsDirectoryServiceGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceGetResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceGetResponseInfraHTTPServiceConfig{}),
			DiscriminatorValue: "http",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceGetResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "tcp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DirectoryServiceGetResponseInfraTCPServiceConfig{}),
			DiscriminatorValue: "http",
		},
	)
}

type DirectoryServiceGetResponseInfraHTTPServiceConfig struct {
	Host      DirectoryServiceGetResponseInfraHTTPServiceConfigHost `json:"host" api:"required"`
	Name      string                                                `json:"name" api:"required"`
	Type      DirectoryServiceGetResponseInfraHTTPServiceConfigType `json:"type" api:"required"`
	CreatedAt time.Time                                             `json:"created_at" format:"date-time"`
	HTTPPort  int64                                                 `json:"http_port" api:"nullable"`
	HTTPSPort int64                                                 `json:"https_port" api:"nullable"`
	ServiceID string                                                `json:"service_id" format:"uuid"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceGetResponseInfraHTTPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON        directoryServiceGetResponseInfraHTTPServiceConfigJSON        `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigJSON contains the JSON metadata
// for the struct [DirectoryServiceGetResponseInfraHTTPServiceConfig]
type directoryServiceGetResponseInfraHTTPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraHTTPServiceConfig) implementsDirectoryServiceGetResponse() {}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                               `json:"resolver_network"`
	JSON            directoryServiceGetResponseInfraHTTPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceGetResponseInfraHTTPServiceConfigHostUnion
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostJSON contains the JSON
// metadata for the struct [DirectoryServiceGetResponseInfraHTTPServiceConfigHost]
type directoryServiceGetResponseInfraHTTPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceGetResponseInfraHTTPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceGetResponseInfraHTTPServiceConfigHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceGetResponseInfraHTTPServiceConfigHost) AsUnion() DirectoryServiceGetResponseInfraHTTPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHost].
type DirectoryServiceGetResponseInfraHTTPServiceConfigHostUnion interface {
	implementsDirectoryServiceGetResponseInfraHTTPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceGetResponseInfraHTTPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                    `json:"ipv4" api:"required"`
	Network DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4Host]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceGetResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                        `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                    `json:"ipv6" api:"required"`
	Network DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6Host]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceGetResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                        `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                         `json:"ipv4" api:"required"`
	IPV6    string                                                                         `json:"ipv6" api:"required"`
	Network DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHost]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceGetResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                             `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                                `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHost]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceGetResponseInfraHTTPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                    `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                  `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraHTTPServiceConfigType string

const (
	DirectoryServiceGetResponseInfraHTTPServiceConfigTypeTCP  DirectoryServiceGetResponseInfraHTTPServiceConfigType = "tcp"
	DirectoryServiceGetResponseInfraHTTPServiceConfigTypeHTTP DirectoryServiceGetResponseInfraHTTPServiceConfigType = "http"
)

func (r DirectoryServiceGetResponseInfraHTTPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceGetResponseInfraHTTPServiceConfigTypeTCP, DirectoryServiceGetResponseInfraHTTPServiceConfigTypeHTTP:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceGetResponseInfraHTTPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                           `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceGetResponseInfraHTTPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceGetResponseInfraHTTPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceGetResponseInfraHTTPServiceConfigTLSSettings]
type directoryServiceGetResponseInfraHTTPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraHTTPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraHTTPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraTCPServiceConfig struct {
	Host        DirectoryServiceGetResponseInfraTCPServiceConfigHost        `json:"host" api:"required"`
	Name        string                                                      `json:"name" api:"required"`
	Type        DirectoryServiceGetResponseInfraTCPServiceConfigType        `json:"type" api:"required"`
	AppProtocol DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocol `json:"app_protocol" api:"nullable"`
	CreatedAt   time.Time                                                   `json:"created_at" format:"date-time"`
	ServiceID   string                                                      `json:"service_id" format:"uuid"`
	TCPPort     int64                                                       `json:"tcp_port" api:"nullable"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings DirectoryServiceGetResponseInfraTCPServiceConfigTLSSettings `json:"tls_settings" api:"nullable"`
	UpdatedAt   time.Time                                                   `json:"updated_at" format:"date-time"`
	JSON        directoryServiceGetResponseInfraTCPServiceConfigJSON        `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigJSON contains the JSON metadata
// for the struct [DirectoryServiceGetResponseInfraTCPServiceConfig]
type directoryServiceGetResponseInfraTCPServiceConfigJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppProtocol apijson.Field
	CreatedAt   apijson.Field
	ServiceID   apijson.Field
	TCPPort     apijson.Field
	TLSSettings apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraTCPServiceConfig) implementsDirectoryServiceGetResponse() {}

type DirectoryServiceGetResponseInfraTCPServiceConfigHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork],
	// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork],
	// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                              `json:"resolver_network"`
	JSON            directoryServiceGetResponseInfraTCPServiceConfigHostJSON `json:"-"`
	union           DirectoryServiceGetResponseInfraTCPServiceConfigHostUnion
}

// directoryServiceGetResponseInfraTCPServiceConfigHostJSON contains the JSON
// metadata for the struct [DirectoryServiceGetResponseInfraTCPServiceConfigHost]
type directoryServiceGetResponseInfraTCPServiceConfigHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceGetResponseInfraTCPServiceConfigHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceGetResponseInfraTCPServiceConfigHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHost],
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHost].
func (r DirectoryServiceGetResponseInfraTCPServiceConfigHost) AsUnion() DirectoryServiceGetResponseInfraTCPServiceConfigHostUnion {
	return r.union
}

// Union satisfied by
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4Host],
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6Host],
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHost] or
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHost].
type DirectoryServiceGetResponseInfraTCPServiceConfigHostUnion interface {
	implementsDirectoryServiceGetResponseInfraTCPServiceConfigHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceGetResponseInfraTCPServiceConfigHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4Host struct {
	IPV4    string                                                                   `json:"ipv4" api:"required"`
	Network DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4Host]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceGetResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                       `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6Host struct {
	IPV6    string                                                                   `json:"ipv6" api:"required"`
	Network DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork `json:"network" api:"required"`
	JSON    directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostJSON contains
// the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6Host]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceGetResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                       `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHost struct {
	IPV4    string                                                                        `json:"ipv4" api:"required"`
	IPV6    string                                                                        `json:"ipv6" api:"required"`
	Network DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork `json:"network" api:"required"`
	JSON    directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHost]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceGetResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                            `json:"tunnel_id" api:"required" format:"uuid"`
	JSON     directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHost struct {
	Hostname        string                                                                               `json:"hostname" api:"required"`
	ResolverNetwork DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork `json:"resolver_network" api:"required"`
	JSON            directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHost]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceGetResponseInfraTCPServiceConfigHost() {
}

type DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                                   `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs []string                                                                                 `json:"resolver_ips" api:"nullable"`
	JSON        directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork]
type directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseInfraTCPServiceConfigType string

const (
	DirectoryServiceGetResponseInfraTCPServiceConfigTypeTCP  DirectoryServiceGetResponseInfraTCPServiceConfigType = "tcp"
	DirectoryServiceGetResponseInfraTCPServiceConfigTypeHTTP DirectoryServiceGetResponseInfraTCPServiceConfigType = "http"
)

func (r DirectoryServiceGetResponseInfraTCPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceGetResponseInfraTCPServiceConfigTypeTCP, DirectoryServiceGetResponseInfraTCPServiceConfigTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocol string

const (
	DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocolPostgresql DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocol = "postgresql"
	DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocolMysql      DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocol = "mysql"
)

func (r DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocolPostgresql, DirectoryServiceGetResponseInfraTCPServiceConfigAppProtocolMysql:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceGetResponseInfraTCPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode string                                                          `json:"cert_verification_mode" api:"required"`
	JSON                 directoryServiceGetResponseInfraTCPServiceConfigTLSSettingsJSON `json:"-"`
}

// directoryServiceGetResponseInfraTCPServiceConfigTLSSettingsJSON contains the
// JSON metadata for the struct
// [DirectoryServiceGetResponseInfraTCPServiceConfigTLSSettings]
type directoryServiceGetResponseInfraTCPServiceConfigTLSSettingsJSON struct {
	CERTVerificationMode apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseInfraTCPServiceConfigTLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseInfraTCPServiceConfigTLSSettingsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseType string

const (
	DirectoryServiceGetResponseTypeTCP  DirectoryServiceGetResponseType = "tcp"
	DirectoryServiceGetResponseTypeHTTP DirectoryServiceGetResponseType = "http"
)

func (r DirectoryServiceGetResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceGetResponseTypeTCP, DirectoryServiceGetResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceGetResponseAppProtocol string

const (
	DirectoryServiceGetResponseAppProtocolPostgresql DirectoryServiceGetResponseAppProtocol = "postgresql"
	DirectoryServiceGetResponseAppProtocolMysql      DirectoryServiceGetResponseAppProtocol = "mysql"
)

func (r DirectoryServiceGetResponseAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceGetResponseAppProtocolPostgresql, DirectoryServiceGetResponseAppProtocolMysql:
		return true
	}
	return false
}

type DirectoryServiceNewParams struct {
	// Account identifier
	AccountID param.Field[string]                `path:"account_id" api:"required"`
	Body      DirectoryServiceNewParamsBodyUnion `json:"body" api:"required"`
}

func (r DirectoryServiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DirectoryServiceNewParamsBody struct {
	Host        param.Field[interface{}]                              `json:"host" api:"required"`
	Name        param.Field[string]                                   `json:"name" api:"required"`
	Type        param.Field[DirectoryServiceNewParamsBodyType]        `json:"type" api:"required"`
	AppProtocol param.Field[DirectoryServiceNewParamsBodyAppProtocol] `json:"app_protocol"`
	HTTPPort    param.Field[int64]                                    `json:"http_port"`
	HTTPSPort   param.Field[int64]                                    `json:"https_port"`
	TCPPort     param.Field[int64]                                    `json:"tcp_port"`
	TLSSettings param.Field[interface{}]                              `json:"tls_settings"`
}

func (r DirectoryServiceNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBody) implementsDirectoryServiceNewParamsBodyUnion() {}

// Satisfied by [connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfig],
// [connectivity.DirectoryServiceNewParamsBodyInfraTCPServiceConfig],
// [DirectoryServiceNewParamsBody].
type DirectoryServiceNewParamsBodyUnion interface {
	implementsDirectoryServiceNewParamsBodyUnion()
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfig struct {
	Host      param.Field[DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion] `json:"host" api:"required"`
	Name      param.Field[string]                                                       `json:"name" api:"required"`
	Type      param.Field[DirectoryServiceNewParamsBodyInfraHTTPServiceConfigType]      `json:"type" api:"required"`
	HTTPPort  param.Field[int64]                                                        `json:"http_port"`
	HTTPSPort param.Field[int64]                                                        `json:"https_port"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings param.Field[DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTLSSettings] `json:"tls_settings"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfig) implementsDirectoryServiceNewParamsBodyUnion() {
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHost) implementsDirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion() {
}

// Satisfied by
// [connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host],
// [connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host],
// [connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost],
// [connectivity.DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost],
// [DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHost].
type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion interface {
	implementsDirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion()
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                                                      `json:"ipv4" api:"required"`
	Network param.Field[DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv4HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                                                      `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv6HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                                           `json:"ipv4" api:"required"`
	IPV6    param.Field[string]                                                                           `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraDualStackHostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                                                  `json:"hostname" api:"required"`
	ResolverNetwork param.Field[DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork] `json:"resolver_network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigType string

const (
	DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTypeTCP  DirectoryServiceNewParamsBodyInfraHTTPServiceConfigType = "tcp"
	DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTypeHTTP DirectoryServiceNewParamsBodyInfraHTTPServiceConfigType = "http"
)

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTypeTCP, DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTypeHTTP:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode param.Field[string] `json:"cert_verification_mode" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraHTTPServiceConfigTLSSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfig struct {
	Host        param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion]   `json:"host" api:"required"`
	Name        param.Field[string]                                                        `json:"name" api:"required"`
	Type        param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigType]        `json:"type" api:"required"`
	AppProtocol param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocol] `json:"app_protocol"`
	TCPPort     param.Field[int64]                                                         `json:"tcp_port"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigTLSSettings] `json:"tls_settings"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfig) implementsDirectoryServiceNewParamsBodyUnion() {
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHost) implementsDirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion() {
}

// Satisfied by
// [connectivity.DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv4Host],
// [connectivity.DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv6Host],
// [connectivity.DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraDualStackHost],
// [connectivity.DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraHostnameHost],
// [DirectoryServiceNewParamsBodyInfraTCPServiceConfigHost].
type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion interface {
	implementsDirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion()
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                                                     `json:"ipv4" api:"required"`
	Network param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv4HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                                                     `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv6HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                                          `json:"ipv4" api:"required"`
	IPV6    param.Field[string]                                                                          `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraDualStackHostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                                                 `json:"hostname" api:"required"`
	ResolverNetwork param.Field[DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork] `json:"resolver_network" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceNewParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigType string

const (
	DirectoryServiceNewParamsBodyInfraTCPServiceConfigTypeTCP  DirectoryServiceNewParamsBodyInfraTCPServiceConfigType = "tcp"
	DirectoryServiceNewParamsBodyInfraTCPServiceConfigTypeHTTP DirectoryServiceNewParamsBodyInfraTCPServiceConfigType = "http"
)

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewParamsBodyInfraTCPServiceConfigTypeTCP, DirectoryServiceNewParamsBodyInfraTCPServiceConfigTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocol string

const (
	DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocolPostgresql DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocol = "postgresql"
	DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocolMysql      DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocol = "mysql"
)

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocolPostgresql, DirectoryServiceNewParamsBodyInfraTCPServiceConfigAppProtocolMysql:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceNewParamsBodyInfraTCPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode param.Field[string] `json:"cert_verification_mode" api:"required"`
}

func (r DirectoryServiceNewParamsBodyInfraTCPServiceConfigTLSSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsBodyType string

const (
	DirectoryServiceNewParamsBodyTypeTCP  DirectoryServiceNewParamsBodyType = "tcp"
	DirectoryServiceNewParamsBodyTypeHTTP DirectoryServiceNewParamsBodyType = "http"
)

func (r DirectoryServiceNewParamsBodyType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewParamsBodyTypeTCP, DirectoryServiceNewParamsBodyTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceNewParamsBodyAppProtocol string

const (
	DirectoryServiceNewParamsBodyAppProtocolPostgresql DirectoryServiceNewParamsBodyAppProtocol = "postgresql"
	DirectoryServiceNewParamsBodyAppProtocolMysql      DirectoryServiceNewParamsBodyAppProtocol = "mysql"
)

func (r DirectoryServiceNewParamsBodyAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceNewParamsBodyAppProtocolPostgresql, DirectoryServiceNewParamsBodyAppProtocolMysql:
		return true
	}
	return false
}

type DirectoryServiceNewResponseEnvelope struct {
	Errors   []DirectoryServiceNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DirectoryServiceNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DirectoryServiceNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DirectoryServiceNewResponse                `json:"result"`
	JSON    directoryServiceNewResponseEnvelopeJSON    `json:"-"`
}

// directoryServiceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DirectoryServiceNewResponseEnvelope]
type directoryServiceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DirectoryServiceNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             directoryServiceNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// directoryServiceNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DirectoryServiceNewResponseEnvelopeErrors]
type directoryServiceNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    directoryServiceNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// directoryServiceNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DirectoryServiceNewResponseEnvelopeErrorsSource]
type directoryServiceNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DirectoryServiceNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             directoryServiceNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// directoryServiceNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DirectoryServiceNewResponseEnvelopeMessages]
type directoryServiceNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    directoryServiceNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// directoryServiceNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DirectoryServiceNewResponseEnvelopeMessagesSource]
type directoryServiceNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DirectoryServiceNewResponseEnvelopeSuccess bool

const (
	DirectoryServiceNewResponseEnvelopeSuccessTrue DirectoryServiceNewResponseEnvelopeSuccess = true
)

func (r DirectoryServiceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DirectoryServiceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DirectoryServiceUpdateParams struct {
	AccountID param.Field[string]                   `path:"account_id" api:"required"`
	Body      DirectoryServiceUpdateParamsBodyUnion `json:"body" api:"required"`
}

func (r DirectoryServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DirectoryServiceUpdateParamsBody struct {
	Host        param.Field[interface{}]                                 `json:"host" api:"required"`
	Name        param.Field[string]                                      `json:"name" api:"required"`
	Type        param.Field[DirectoryServiceUpdateParamsBodyType]        `json:"type" api:"required"`
	AppProtocol param.Field[DirectoryServiceUpdateParamsBodyAppProtocol] `json:"app_protocol"`
	HTTPPort    param.Field[int64]                                       `json:"http_port"`
	HTTPSPort   param.Field[int64]                                       `json:"https_port"`
	TCPPort     param.Field[int64]                                       `json:"tcp_port"`
	TLSSettings param.Field[interface{}]                                 `json:"tls_settings"`
}

func (r DirectoryServiceUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBody) implementsDirectoryServiceUpdateParamsBodyUnion() {}

// Satisfied by
// [connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfig],
// [connectivity.DirectoryServiceUpdateParamsBodyInfraTCPServiceConfig],
// [DirectoryServiceUpdateParamsBody].
type DirectoryServiceUpdateParamsBodyUnion interface {
	implementsDirectoryServiceUpdateParamsBodyUnion()
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfig struct {
	Host      param.Field[DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion] `json:"host" api:"required"`
	Name      param.Field[string]                                                          `json:"name" api:"required"`
	Type      param.Field[DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigType]      `json:"type" api:"required"`
	HTTPPort  param.Field[int64]                                                           `json:"http_port"`
	HTTPSPort param.Field[int64]                                                           `json:"https_port"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings param.Field[DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTLSSettings] `json:"tls_settings"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfig) implementsDirectoryServiceUpdateParamsBodyUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHost) implementsDirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion() {
}

// Satisfied by
// [connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host],
// [connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host],
// [connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost],
// [connectivity.DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost],
// [DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHost].
type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion interface {
	implementsDirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion()
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                                                         `json:"ipv4" api:"required"`
	Network param.Field[DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                                                         `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv6HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                                              `json:"ipv4" api:"required"`
	IPV6    param.Field[string]                                                                              `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraDualStackHostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                                                     `json:"hostname" api:"required"`
	ResolverNetwork param.Field[DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork] `json:"resolver_network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigType string

const (
	DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTypeTCP  DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigType = "tcp"
	DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTypeHTTP DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigType = "http"
)

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTypeTCP, DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTypeHTTP:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode param.Field[string] `json:"cert_verification_mode" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraHTTPServiceConfigTLSSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfig struct {
	Host        param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion]   `json:"host" api:"required"`
	Name        param.Field[string]                                                           `json:"name" api:"required"`
	Type        param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigType]        `json:"type" api:"required"`
	AppProtocol param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocol] `json:"app_protocol"`
	TCPPort     param.Field[int64]                                                            `json:"tcp_port"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	TLSSettings param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigTLSSettings] `json:"tls_settings"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfig) implementsDirectoryServiceUpdateParamsBodyUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHost) implementsDirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion() {
}

// Satisfied by
// [connectivity.DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv4Host],
// [connectivity.DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv6Host],
// [connectivity.DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraDualStackHost],
// [connectivity.DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraHostnameHost],
// [DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHost].
type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion interface {
	implementsDirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion()
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                                                        `json:"ipv4" api:"required"`
	Network param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv4HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv4Host) implementsDirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                                                        `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv6HostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv6Host) implementsDirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                                             `json:"ipv4" api:"required"`
	IPV6    param.Field[string]                                                                             `json:"ipv6" api:"required"`
	Network param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraDualStackHostNetwork] `json:"network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraDualStackHost) implementsDirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id" api:"required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                                                    `json:"hostname" api:"required"`
	ResolverNetwork param.Field[DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork] `json:"resolver_network" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraHostnameHost) implementsDirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostUnion() {
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id" api:"required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigType string

const (
	DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigTypeTCP  DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigType = "tcp"
	DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigTypeHTTP DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigType = "http"
)

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigTypeTCP, DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocol string

const (
	DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocolPostgresql DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocol = "postgresql"
	DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocolMysql      DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocol = "mysql"
)

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocolPostgresql, DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigAppProtocolMysql:
		return true
	}
	return false
}

// TLS settings for a connectivity service.
//
// If omitted, the default mode (`verify_full`) is used.
type DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigTLSSettings struct {
	// TLS certificate verification mode for the connection to the origin.
	//
	// - `"verify_full"` — verify certificate chain and hostname (default)
	// - `"verify_ca"` — verify certificate chain only, skip hostname check
	// - `"disabled"` — do not verify the server certificate at all
	CERTVerificationMode param.Field[string] `json:"cert_verification_mode" api:"required"`
}

func (r DirectoryServiceUpdateParamsBodyInfraTCPServiceConfigTLSSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsBodyType string

const (
	DirectoryServiceUpdateParamsBodyTypeTCP  DirectoryServiceUpdateParamsBodyType = "tcp"
	DirectoryServiceUpdateParamsBodyTypeHTTP DirectoryServiceUpdateParamsBodyType = "http"
)

func (r DirectoryServiceUpdateParamsBodyType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateParamsBodyTypeTCP, DirectoryServiceUpdateParamsBodyTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceUpdateParamsBodyAppProtocol string

const (
	DirectoryServiceUpdateParamsBodyAppProtocolPostgresql DirectoryServiceUpdateParamsBodyAppProtocol = "postgresql"
	DirectoryServiceUpdateParamsBodyAppProtocolMysql      DirectoryServiceUpdateParamsBodyAppProtocol = "mysql"
)

func (r DirectoryServiceUpdateParamsBodyAppProtocol) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateParamsBodyAppProtocolPostgresql, DirectoryServiceUpdateParamsBodyAppProtocolMysql:
		return true
	}
	return false
}

type DirectoryServiceUpdateResponseEnvelope struct {
	Errors   []DirectoryServiceUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DirectoryServiceUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DirectoryServiceUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DirectoryServiceUpdateResponse                `json:"result"`
	JSON    directoryServiceUpdateResponseEnvelopeJSON    `json:"-"`
}

// directoryServiceUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DirectoryServiceUpdateResponseEnvelope]
type directoryServiceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DirectoryServiceUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             directoryServiceUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// directoryServiceUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DirectoryServiceUpdateResponseEnvelopeErrors]
type directoryServiceUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    directoryServiceUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// directoryServiceUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseEnvelopeErrorsSource]
type directoryServiceUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DirectoryServiceUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             directoryServiceUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// directoryServiceUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DirectoryServiceUpdateResponseEnvelopeMessages]
type directoryServiceUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    directoryServiceUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// directoryServiceUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseEnvelopeMessagesSource]
type directoryServiceUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DirectoryServiceUpdateResponseEnvelopeSuccess bool

const (
	DirectoryServiceUpdateResponseEnvelopeSuccessTrue DirectoryServiceUpdateResponseEnvelopeSuccess = true
)

func (r DirectoryServiceUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DirectoryServiceListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Current page in the response
	Page param.Field[int64] `query:"page"`
	// Max amount of entries returned per page
	PerPage param.Field[int64]                          `query:"per_page"`
	Type    param.Field[DirectoryServiceListParamsType] `query:"type"`
}

// URLQuery serializes [DirectoryServiceListParams]'s query parameters as
// `url.Values`.
func (r DirectoryServiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DirectoryServiceListParamsType string

const (
	DirectoryServiceListParamsTypeTCP  DirectoryServiceListParamsType = "tcp"
	DirectoryServiceListParamsTypeHTTP DirectoryServiceListParamsType = "http"
)

func (r DirectoryServiceListParamsType) IsKnown() bool {
	switch r {
	case DirectoryServiceListParamsTypeTCP, DirectoryServiceListParamsTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DirectoryServiceGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DirectoryServiceGetResponseEnvelope struct {
	Errors   []DirectoryServiceGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DirectoryServiceGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DirectoryServiceGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DirectoryServiceGetResponse                `json:"result"`
	JSON    directoryServiceGetResponseEnvelopeJSON    `json:"-"`
}

// directoryServiceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DirectoryServiceGetResponseEnvelope]
type directoryServiceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DirectoryServiceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             directoryServiceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// directoryServiceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DirectoryServiceGetResponseEnvelopeErrors]
type directoryServiceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    directoryServiceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// directoryServiceGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DirectoryServiceGetResponseEnvelopeErrorsSource]
type directoryServiceGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           DirectoryServiceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             directoryServiceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// directoryServiceGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DirectoryServiceGetResponseEnvelopeMessages]
type directoryServiceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    directoryServiceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// directoryServiceGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DirectoryServiceGetResponseEnvelopeMessagesSource]
type directoryServiceGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DirectoryServiceGetResponseEnvelopeSuccess bool

const (
	DirectoryServiceGetResponseEnvelopeSuccessTrue DirectoryServiceGetResponseEnvelopeSuccess = true
)

func (r DirectoryServiceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DirectoryServiceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
