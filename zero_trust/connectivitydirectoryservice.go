// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// ConnectivityDirectoryServiceService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectivityDirectoryServiceService] method instead.
type ConnectivityDirectoryServiceService struct {
	Options []option.RequestOption
}

// NewConnectivityDirectoryServiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectivityDirectoryServiceService(opts ...option.RequestOption) (r *ConnectivityDirectoryServiceService) {
	r = &ConnectivityDirectoryServiceService{}
	r.Options = opts
	return
}

// Create connectivity service
func (r *ConnectivityDirectoryServiceService) New(ctx context.Context, params ConnectivityDirectoryServiceNewParams, opts ...option.RequestOption) (res *ConnectivityDirectoryServiceNewResponse, err error) {
	var env ConnectivityDirectoryServiceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update connectivity service
func (r *ConnectivityDirectoryServiceService) Update(ctx context.Context, serviceID string, params ConnectivityDirectoryServiceUpdateParams, opts ...option.RequestOption) (res *ConnectivityDirectoryServiceUpdateResponse, err error) {
	var env ConnectivityDirectoryServiceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", params.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List connectivity services
func (r *ConnectivityDirectoryServiceService) List(ctx context.Context, params ConnectivityDirectoryServiceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ConnectivityDirectoryServiceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
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

// List connectivity services
func (r *ConnectivityDirectoryServiceService) ListAutoPaging(ctx context.Context, params ConnectivityDirectoryServiceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ConnectivityDirectoryServiceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete connectivity service
func (r *ConnectivityDirectoryServiceService) Delete(ctx context.Context, serviceID string, body ConnectivityDirectoryServiceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", body.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get connectivity service
func (r *ConnectivityDirectoryServiceService) Get(ctx context.Context, serviceID string, query ConnectivityDirectoryServiceGetParams, opts ...option.RequestOption) (res *ConnectivityDirectoryServiceGetResponse, err error) {
	var env ConnectivityDirectoryServiceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", query.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectivityDirectoryServiceNewResponse struct {
	Host      ConnectivityDirectoryServiceNewResponseHost `json:"host,required"`
	Name      string                                      `json:"name,required"`
	Type      ConnectivityDirectoryServiceNewResponseType `json:"type,required"`
	CreatedAt time.Time                                   `json:"created_at" format:"date-time"`
	HTTPPort  int64                                       `json:"http_port,nullable"`
	HTTPSPort int64                                       `json:"https_port,nullable"`
	ServiceID string                                      `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      connectivityDirectoryServiceNewResponseJSON `json:"-"`
}

// connectivityDirectoryServiceNewResponseJSON contains the JSON metadata for the
// struct [ConnectivityDirectoryServiceNewResponse]
type connectivityDirectoryServiceNewResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceNewResponseHostInfraIPv4HostNetwork],
	// [ConnectivityDirectoryServiceNewResponseHostInfraIPv6HostNetwork],
	// [ConnectivityDirectoryServiceNewResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                     `json:"resolver_network"`
	JSON            connectivityDirectoryServiceNewResponseHostJSON `json:"-"`
	union           ConnectivityDirectoryServiceNewResponseHostUnion
}

// connectivityDirectoryServiceNewResponseHostJSON contains the JSON metadata for
// the struct [ConnectivityDirectoryServiceNewResponseHost]
type connectivityDirectoryServiceNewResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r connectivityDirectoryServiceNewResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectivityDirectoryServiceNewResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectivityDirectoryServiceNewResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectivityDirectoryServiceNewResponseHostUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConnectivityDirectoryServiceNewResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceNewResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceNewResponseHostInfraDualStackHost],
// [ConnectivityDirectoryServiceNewResponseHostInfraHostnameHost].
func (r ConnectivityDirectoryServiceNewResponseHost) AsUnion() ConnectivityDirectoryServiceNewResponseHostUnion {
	return r.union
}

// Union satisfied by [ConnectivityDirectoryServiceNewResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceNewResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceNewResponseHostInfraDualStackHost] or
// [ConnectivityDirectoryServiceNewResponseHostInfraHostnameHost].
type ConnectivityDirectoryServiceNewResponseHostUnion interface {
	implementsConnectivityDirectoryServiceNewResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectivityDirectoryServiceNewResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceNewResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceNewResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceNewResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceNewResponseHostInfraHostnameHost{}),
		},
	)
}

type ConnectivityDirectoryServiceNewResponseHostInfraIPv4Host struct {
	IPV4    string                                                          `json:"ipv4,required"`
	Network ConnectivityDirectoryServiceNewResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceNewResponseHostInfraIPv4HostJSON    `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraIPv4HostJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraIPv4Host]
type connectivityDirectoryServiceNewResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceNewResponseHostInfraIPv4Host) implementsConnectivityDirectoryServiceNewResponseHost() {
}

type ConnectivityDirectoryServiceNewResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                              `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceNewResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraIPv4HostNetworkJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraIPv4HostNetwork]
type connectivityDirectoryServiceNewResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseHostInfraIPv6Host struct {
	IPV6    string                                                          `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceNewResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceNewResponseHostInfraIPv6HostJSON    `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraIPv6HostJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraIPv6Host]
type connectivityDirectoryServiceNewResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceNewResponseHostInfraIPv6Host) implementsConnectivityDirectoryServiceNewResponseHost() {
}

type ConnectivityDirectoryServiceNewResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                              `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceNewResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraIPv6HostNetworkJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraIPv6HostNetwork]
type connectivityDirectoryServiceNewResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseHostInfraDualStackHost struct {
	IPV4    string                                                               `json:"ipv4,required"`
	IPV6    string                                                               `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceNewResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceNewResponseHostInfraDualStackHostJSON    `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraDualStackHostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraDualStackHost]
type connectivityDirectoryServiceNewResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceNewResponseHostInfraDualStackHost) implementsConnectivityDirectoryServiceNewResponseHost() {
}

type ConnectivityDirectoryServiceNewResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                   `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceNewResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraDualStackHostNetwork]
type connectivityDirectoryServiceNewResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseHostInfraHostnameHost struct {
	Hostname        string                                                                      `json:"hostname,required"`
	ResolverNetwork ConnectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            connectivityDirectoryServiceNewResponseHostInfraHostnameHostJSON            `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraHostnameHostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraHostnameHost]
type connectivityDirectoryServiceNewResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceNewResponseHostInfraHostnameHost) implementsConnectivityDirectoryServiceNewResponseHost() {
}

type ConnectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                          `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                                        `json:"resolver_ips,nullable"`
	JSON        connectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork]
type connectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseType string

const (
	ConnectivityDirectoryServiceNewResponseTypeHTTP ConnectivityDirectoryServiceNewResponseType = "http"
)

func (r ConnectivityDirectoryServiceNewResponseType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceNewResponseTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceUpdateResponse struct {
	Host      ConnectivityDirectoryServiceUpdateResponseHost `json:"host,required"`
	Name      string                                         `json:"name,required"`
	Type      ConnectivityDirectoryServiceUpdateResponseType `json:"type,required"`
	CreatedAt time.Time                                      `json:"created_at" format:"date-time"`
	HTTPPort  int64                                          `json:"http_port,nullable"`
	HTTPSPort int64                                          `json:"https_port,nullable"`
	ServiceID string                                         `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      connectivityDirectoryServiceUpdateResponseJSON `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseJSON contains the JSON metadata for
// the struct [ConnectivityDirectoryServiceUpdateResponse]
type connectivityDirectoryServiceUpdateResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetwork],
	// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetwork],
	// [ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                        `json:"resolver_network"`
	JSON            connectivityDirectoryServiceUpdateResponseHostJSON `json:"-"`
	union           ConnectivityDirectoryServiceUpdateResponseHostUnion
}

// connectivityDirectoryServiceUpdateResponseHostJSON contains the JSON metadata
// for the struct [ConnectivityDirectoryServiceUpdateResponseHost]
type connectivityDirectoryServiceUpdateResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r connectivityDirectoryServiceUpdateResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectivityDirectoryServiceUpdateResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectivityDirectoryServiceUpdateResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectivityDirectoryServiceUpdateResponseHostUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHost],
// [ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHost].
func (r ConnectivityDirectoryServiceUpdateResponseHost) AsUnion() ConnectivityDirectoryServiceUpdateResponseHostUnion {
	return r.union
}

// Union satisfied by
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHost] or
// [ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHost].
type ConnectivityDirectoryServiceUpdateResponseHostUnion interface {
	implementsConnectivityDirectoryServiceUpdateResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectivityDirectoryServiceUpdateResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHost{}),
		},
	)
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4Host struct {
	IPV4    string                                                             `json:"ipv4,required"`
	Network ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostJSON    `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4Host]
type connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4Host) implementsConnectivityDirectoryServiceUpdateResponseHost() {
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                                 `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON contains
// the JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetwork]
type connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6Host struct {
	IPV6    string                                                             `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostJSON    `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6Host]
type connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6Host) implementsConnectivityDirectoryServiceUpdateResponseHost() {
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                                 `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON contains
// the JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetwork]
type connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHost struct {
	IPV4    string                                                                  `json:"ipv4,required"`
	IPV6    string                                                                  `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostJSON    `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostJSON contains
// the JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHost]
type connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHost) implementsConnectivityDirectoryServiceUpdateResponseHost() {
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                      `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetwork]
type connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHost struct {
	Hostname        string                                                                         `json:"hostname,required"`
	ResolverNetwork ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostJSON            `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHost]
type connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHost) implementsConnectivityDirectoryServiceUpdateResponseHost() {
}

type ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                             `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                                           `json:"resolver_ips,nullable"`
	JSON        connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork]
type connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseType string

const (
	ConnectivityDirectoryServiceUpdateResponseTypeHTTP ConnectivityDirectoryServiceUpdateResponseType = "http"
)

func (r ConnectivityDirectoryServiceUpdateResponseType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceUpdateResponseTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceListResponse struct {
	Host      ConnectivityDirectoryServiceListResponseHost `json:"host,required"`
	Name      string                                       `json:"name,required"`
	Type      ConnectivityDirectoryServiceListResponseType `json:"type,required"`
	CreatedAt time.Time                                    `json:"created_at" format:"date-time"`
	HTTPPort  int64                                        `json:"http_port,nullable"`
	HTTPSPort int64                                        `json:"https_port,nullable"`
	ServiceID string                                       `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      connectivityDirectoryServiceListResponseJSON `json:"-"`
}

// connectivityDirectoryServiceListResponseJSON contains the JSON metadata for the
// struct [ConnectivityDirectoryServiceListResponse]
type connectivityDirectoryServiceListResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceListResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceListResponseHostInfraIPv4HostNetwork],
	// [ConnectivityDirectoryServiceListResponseHostInfraIPv6HostNetwork],
	// [ConnectivityDirectoryServiceListResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                      `json:"resolver_network"`
	JSON            connectivityDirectoryServiceListResponseHostJSON `json:"-"`
	union           ConnectivityDirectoryServiceListResponseHostUnion
}

// connectivityDirectoryServiceListResponseHostJSON contains the JSON metadata for
// the struct [ConnectivityDirectoryServiceListResponseHost]
type connectivityDirectoryServiceListResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r connectivityDirectoryServiceListResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectivityDirectoryServiceListResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectivityDirectoryServiceListResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectivityDirectoryServiceListResponseHostUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConnectivityDirectoryServiceListResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceListResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceListResponseHostInfraDualStackHost],
// [ConnectivityDirectoryServiceListResponseHostInfraHostnameHost].
func (r ConnectivityDirectoryServiceListResponseHost) AsUnion() ConnectivityDirectoryServiceListResponseHostUnion {
	return r.union
}

// Union satisfied by [ConnectivityDirectoryServiceListResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceListResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceListResponseHostInfraDualStackHost] or
// [ConnectivityDirectoryServiceListResponseHostInfraHostnameHost].
type ConnectivityDirectoryServiceListResponseHostUnion interface {
	implementsConnectivityDirectoryServiceListResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectivityDirectoryServiceListResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceListResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceListResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceListResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceListResponseHostInfraHostnameHost{}),
		},
	)
}

type ConnectivityDirectoryServiceListResponseHostInfraIPv4Host struct {
	IPV4    string                                                           `json:"ipv4,required"`
	Network ConnectivityDirectoryServiceListResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceListResponseHostInfraIPv4HostJSON    `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraIPv4HostJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraIPv4Host]
type connectivityDirectoryServiceListResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceListResponseHostInfraIPv4Host) implementsConnectivityDirectoryServiceListResponseHost() {
}

type ConnectivityDirectoryServiceListResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                               `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceListResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraIPv4HostNetworkJSON contains
// the JSON metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraIPv4HostNetwork]
type connectivityDirectoryServiceListResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceListResponseHostInfraIPv6Host struct {
	IPV6    string                                                           `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceListResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceListResponseHostInfraIPv6HostJSON    `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraIPv6HostJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraIPv6Host]
type connectivityDirectoryServiceListResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceListResponseHostInfraIPv6Host) implementsConnectivityDirectoryServiceListResponseHost() {
}

type ConnectivityDirectoryServiceListResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                               `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceListResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraIPv6HostNetworkJSON contains
// the JSON metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraIPv6HostNetwork]
type connectivityDirectoryServiceListResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceListResponseHostInfraDualStackHost struct {
	IPV4    string                                                                `json:"ipv4,required"`
	IPV6    string                                                                `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceListResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceListResponseHostInfraDualStackHostJSON    `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraDualStackHostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraDualStackHost]
type connectivityDirectoryServiceListResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceListResponseHostInfraDualStackHost) implementsConnectivityDirectoryServiceListResponseHost() {
}

type ConnectivityDirectoryServiceListResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                    `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceListResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraDualStackHostNetwork]
type connectivityDirectoryServiceListResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceListResponseHostInfraHostnameHost struct {
	Hostname        string                                                                       `json:"hostname,required"`
	ResolverNetwork ConnectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            connectivityDirectoryServiceListResponseHostInfraHostnameHostJSON            `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraHostnameHostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraHostnameHost]
type connectivityDirectoryServiceListResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceListResponseHostInfraHostnameHost) implementsConnectivityDirectoryServiceListResponseHost() {
}

type ConnectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                           `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                                         `json:"resolver_ips,nullable"`
	JSON        connectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetwork]
type connectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceListResponseType string

const (
	ConnectivityDirectoryServiceListResponseTypeHTTP ConnectivityDirectoryServiceListResponseType = "http"
)

func (r ConnectivityDirectoryServiceListResponseType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceListResponseTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceGetResponse struct {
	Host      ConnectivityDirectoryServiceGetResponseHost `json:"host,required"`
	Name      string                                      `json:"name,required"`
	Type      ConnectivityDirectoryServiceGetResponseType `json:"type,required"`
	CreatedAt time.Time                                   `json:"created_at" format:"date-time"`
	HTTPPort  int64                                       `json:"http_port,nullable"`
	HTTPSPort int64                                       `json:"https_port,nullable"`
	ServiceID string                                      `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      connectivityDirectoryServiceGetResponseJSON `json:"-"`
}

// connectivityDirectoryServiceGetResponseJSON contains the JSON metadata for the
// struct [ConnectivityDirectoryServiceGetResponse]
type connectivityDirectoryServiceGetResponseJSON struct {
	Host        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedAt   apijson.Field
	HTTPPort    apijson.Field
	HTTPSPort   apijson.Field
	ServiceID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceGetResponseHostInfraIPv4HostNetwork],
	// [ConnectivityDirectoryServiceGetResponseHostInfraIPv6HostNetwork],
	// [ConnectivityDirectoryServiceGetResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [ConnectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                                     `json:"resolver_network"`
	JSON            connectivityDirectoryServiceGetResponseHostJSON `json:"-"`
	union           ConnectivityDirectoryServiceGetResponseHostUnion
}

// connectivityDirectoryServiceGetResponseHostJSON contains the JSON metadata for
// the struct [ConnectivityDirectoryServiceGetResponseHost]
type connectivityDirectoryServiceGetResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r connectivityDirectoryServiceGetResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectivityDirectoryServiceGetResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectivityDirectoryServiceGetResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectivityDirectoryServiceGetResponseHostUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConnectivityDirectoryServiceGetResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceGetResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceGetResponseHostInfraDualStackHost],
// [ConnectivityDirectoryServiceGetResponseHostInfraHostnameHost].
func (r ConnectivityDirectoryServiceGetResponseHost) AsUnion() ConnectivityDirectoryServiceGetResponseHostUnion {
	return r.union
}

// Union satisfied by [ConnectivityDirectoryServiceGetResponseHostInfraIPv4Host],
// [ConnectivityDirectoryServiceGetResponseHostInfraIPv6Host],
// [ConnectivityDirectoryServiceGetResponseHostInfraDualStackHost] or
// [ConnectivityDirectoryServiceGetResponseHostInfraHostnameHost].
type ConnectivityDirectoryServiceGetResponseHostUnion interface {
	implementsConnectivityDirectoryServiceGetResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectivityDirectoryServiceGetResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceGetResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceGetResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceGetResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConnectivityDirectoryServiceGetResponseHostInfraHostnameHost{}),
		},
	)
}

type ConnectivityDirectoryServiceGetResponseHostInfraIPv4Host struct {
	IPV4    string                                                          `json:"ipv4,required"`
	Network ConnectivityDirectoryServiceGetResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceGetResponseHostInfraIPv4HostJSON    `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraIPv4HostJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraIPv4Host]
type connectivityDirectoryServiceGetResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceGetResponseHostInfraIPv4Host) implementsConnectivityDirectoryServiceGetResponseHost() {
}

type ConnectivityDirectoryServiceGetResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                              `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceGetResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraIPv4HostNetworkJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraIPv4HostNetwork]
type connectivityDirectoryServiceGetResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseHostInfraIPv6Host struct {
	IPV6    string                                                          `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceGetResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceGetResponseHostInfraIPv6HostJSON    `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraIPv6HostJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraIPv6Host]
type connectivityDirectoryServiceGetResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceGetResponseHostInfraIPv6Host) implementsConnectivityDirectoryServiceGetResponseHost() {
}

type ConnectivityDirectoryServiceGetResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                              `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceGetResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraIPv6HostNetworkJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraIPv6HostNetwork]
type connectivityDirectoryServiceGetResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseHostInfraDualStackHost struct {
	IPV4    string                                                               `json:"ipv4,required"`
	IPV6    string                                                               `json:"ipv6,required"`
	Network ConnectivityDirectoryServiceGetResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    connectivityDirectoryServiceGetResponseHostInfraDualStackHostJSON    `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraDualStackHostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraDualStackHost]
type connectivityDirectoryServiceGetResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceGetResponseHostInfraDualStackHost) implementsConnectivityDirectoryServiceGetResponseHost() {
}

type ConnectivityDirectoryServiceGetResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                                   `json:"tunnel_id,required" format:"uuid"`
	JSON     connectivityDirectoryServiceGetResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraDualStackHostNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraDualStackHostNetwork]
type connectivityDirectoryServiceGetResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseHostInfraHostnameHost struct {
	Hostname        string                                                                      `json:"hostname,required"`
	ResolverNetwork ConnectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            connectivityDirectoryServiceGetResponseHostInfraHostnameHostJSON            `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraHostnameHostJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraHostnameHost]
type connectivityDirectoryServiceGetResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r ConnectivityDirectoryServiceGetResponseHostInfraHostnameHost) implementsConnectivityDirectoryServiceGetResponseHost() {
}

type ConnectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                          `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                                        `json:"resolver_ips,nullable"`
	JSON        connectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// connectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON
// contains the JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork]
type connectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseType string

const (
	ConnectivityDirectoryServiceGetResponseTypeHTTP ConnectivityDirectoryServiceGetResponseType = "http"
)

func (r ConnectivityDirectoryServiceGetResponseType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceGetResponseTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceNewParams struct {
	// Account identifier
	AccountID param.Field[string]                                         `path:"account_id,required"`
	Host      param.Field[ConnectivityDirectoryServiceNewParamsHostUnion] `json:"host,required"`
	Name      param.Field[string]                                         `json:"name,required"`
	Type      param.Field[ConnectivityDirectoryServiceNewParamsType]      `json:"type,required"`
	HTTPPort  param.Field[int64]                                          `json:"http_port"`
	HTTPSPort param.Field[int64]                                          `json:"https_port"`
}

func (r ConnectivityDirectoryServiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceNewParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r ConnectivityDirectoryServiceNewParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceNewParamsHost) implementsConnectivityDirectoryServiceNewParamsHostUnion() {
}

// Satisfied by
// [zero_trust.ConnectivityDirectoryServiceNewParamsHostInfraIPv4Host],
// [zero_trust.ConnectivityDirectoryServiceNewParamsHostInfraIPv6Host],
// [zero_trust.ConnectivityDirectoryServiceNewParamsHostInfraDualStackHost],
// [zero_trust.ConnectivityDirectoryServiceNewParamsHostInfraHostnameHost],
// [ConnectivityDirectoryServiceNewParamsHost].
type ConnectivityDirectoryServiceNewParamsHostUnion interface {
	implementsConnectivityDirectoryServiceNewParamsHostUnion()
}

type ConnectivityDirectoryServiceNewParamsHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                                        `json:"ipv4,required"`
	Network param.Field[ConnectivityDirectoryServiceNewParamsHostInfraIPv4HostNetwork] `json:"network,required"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraIPv4Host) implementsConnectivityDirectoryServiceNewParamsHostUnion() {
}

type ConnectivityDirectoryServiceNewParamsHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceNewParamsHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                                        `json:"ipv6,required"`
	Network param.Field[ConnectivityDirectoryServiceNewParamsHostInfraIPv6HostNetwork] `json:"network,required"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraIPv6Host) implementsConnectivityDirectoryServiceNewParamsHostUnion() {
}

type ConnectivityDirectoryServiceNewParamsHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceNewParamsHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                             `json:"ipv4,required"`
	IPV6    param.Field[string]                                                             `json:"ipv6,required"`
	Network param.Field[ConnectivityDirectoryServiceNewParamsHostInfraDualStackHostNetwork] `json:"network,required"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraDualStackHost) implementsConnectivityDirectoryServiceNewParamsHostUnion() {
}

type ConnectivityDirectoryServiceNewParamsHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceNewParamsHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                                    `json:"hostname,required"`
	ResolverNetwork param.Field[ConnectivityDirectoryServiceNewParamsHostInfraHostnameHostResolverNetwork] `json:"resolver_network,required"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraHostnameHost) implementsConnectivityDirectoryServiceNewParamsHostUnion() {
}

type ConnectivityDirectoryServiceNewParamsHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r ConnectivityDirectoryServiceNewParamsHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceNewParamsType string

const (
	ConnectivityDirectoryServiceNewParamsTypeHTTP ConnectivityDirectoryServiceNewParamsType = "http"
)

func (r ConnectivityDirectoryServiceNewParamsType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceNewParamsTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceNewResponseEnvelope struct {
	Errors   []ConnectivityDirectoryServiceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectivityDirectoryServiceNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ConnectivityDirectoryServiceNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ConnectivityDirectoryServiceNewResponse                `json:"result"`
	JSON    connectivityDirectoryServiceNewResponseEnvelopeJSON    `json:"-"`
}

// connectivityDirectoryServiceNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ConnectivityDirectoryServiceNewResponseEnvelope]
type connectivityDirectoryServiceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           ConnectivityDirectoryServiceNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             connectivityDirectoryServiceNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// connectivityDirectoryServiceNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ConnectivityDirectoryServiceNewResponseEnvelopeErrors]
type connectivityDirectoryServiceNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    connectivityDirectoryServiceNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// connectivityDirectoryServiceNewResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseEnvelopeErrorsSource]
type connectivityDirectoryServiceNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           ConnectivityDirectoryServiceNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             connectivityDirectoryServiceNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// connectivityDirectoryServiceNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceNewResponseEnvelopeMessages]
type connectivityDirectoryServiceNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    connectivityDirectoryServiceNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// connectivityDirectoryServiceNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceNewResponseEnvelopeMessagesSource]
type connectivityDirectoryServiceNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConnectivityDirectoryServiceNewResponseEnvelopeSuccess bool

const (
	ConnectivityDirectoryServiceNewResponseEnvelopeSuccessTrue ConnectivityDirectoryServiceNewResponseEnvelopeSuccess = true
)

func (r ConnectivityDirectoryServiceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceUpdateParams struct {
	AccountID param.Field[string]                                            `path:"account_id,required"`
	Host      param.Field[ConnectivityDirectoryServiceUpdateParamsHostUnion] `json:"host,required"`
	Name      param.Field[string]                                            `json:"name,required"`
	Type      param.Field[ConnectivityDirectoryServiceUpdateParamsType]      `json:"type,required"`
	HTTPPort  param.Field[int64]                                             `json:"http_port"`
	HTTPSPort param.Field[int64]                                             `json:"https_port"`
}

func (r ConnectivityDirectoryServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceUpdateParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceUpdateParamsHost) implementsConnectivityDirectoryServiceUpdateParamsHostUnion() {
}

// Satisfied by
// [zero_trust.ConnectivityDirectoryServiceUpdateParamsHostInfraIPv4Host],
// [zero_trust.ConnectivityDirectoryServiceUpdateParamsHostInfraIPv6Host],
// [zero_trust.ConnectivityDirectoryServiceUpdateParamsHostInfraDualStackHost],
// [zero_trust.ConnectivityDirectoryServiceUpdateParamsHostInfraHostnameHost],
// [ConnectivityDirectoryServiceUpdateParamsHost].
type ConnectivityDirectoryServiceUpdateParamsHostUnion interface {
	implementsConnectivityDirectoryServiceUpdateParamsHostUnion()
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                                           `json:"ipv4,required"`
	Network param.Field[ConnectivityDirectoryServiceUpdateParamsHostInfraIPv4HostNetwork] `json:"network,required"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraIPv4Host) implementsConnectivityDirectoryServiceUpdateParamsHostUnion() {
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                                           `json:"ipv6,required"`
	Network param.Field[ConnectivityDirectoryServiceUpdateParamsHostInfraIPv6HostNetwork] `json:"network,required"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraIPv6Host) implementsConnectivityDirectoryServiceUpdateParamsHostUnion() {
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                                `json:"ipv4,required"`
	IPV6    param.Field[string]                                                                `json:"ipv6,required"`
	Network param.Field[ConnectivityDirectoryServiceUpdateParamsHostInfraDualStackHostNetwork] `json:"network,required"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraDualStackHost) implementsConnectivityDirectoryServiceUpdateParamsHostUnion() {
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                                       `json:"hostname,required"`
	ResolverNetwork param.Field[ConnectivityDirectoryServiceUpdateParamsHostInfraHostnameHostResolverNetwork] `json:"resolver_network,required"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraHostnameHost) implementsConnectivityDirectoryServiceUpdateParamsHostUnion() {
}

type ConnectivityDirectoryServiceUpdateParamsHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceUpdateParamsType string

const (
	ConnectivityDirectoryServiceUpdateParamsTypeHTTP ConnectivityDirectoryServiceUpdateParamsType = "http"
)

func (r ConnectivityDirectoryServiceUpdateParamsType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceUpdateParamsTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceUpdateResponseEnvelope struct {
	Errors   []ConnectivityDirectoryServiceUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectivityDirectoryServiceUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ConnectivityDirectoryServiceUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ConnectivityDirectoryServiceUpdateResponse                `json:"result"`
	JSON    connectivityDirectoryServiceUpdateResponseEnvelopeJSON    `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ConnectivityDirectoryServiceUpdateResponseEnvelope]
type connectivityDirectoryServiceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           ConnectivityDirectoryServiceUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             connectivityDirectoryServiceUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseEnvelopeErrors]
type connectivityDirectoryServiceUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    connectivityDirectoryServiceUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseEnvelopeErrorsSource]
type connectivityDirectoryServiceUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code,required"`
	Message          string                                                           `json:"message,required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           ConnectivityDirectoryServiceUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             connectivityDirectoryServiceUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseEnvelopeMessages]
type connectivityDirectoryServiceUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    connectivityDirectoryServiceUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// connectivityDirectoryServiceUpdateResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [ConnectivityDirectoryServiceUpdateResponseEnvelopeMessagesSource]
type connectivityDirectoryServiceUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConnectivityDirectoryServiceUpdateResponseEnvelopeSuccess bool

const (
	ConnectivityDirectoryServiceUpdateResponseEnvelopeSuccessTrue ConnectivityDirectoryServiceUpdateResponseEnvelopeSuccess = true
)

func (r ConnectivityDirectoryServiceUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Current page in the response
	Page param.Field[int64] `query:"page"`
	// Max amount of entries returned per page
	PerPage param.Field[int64]                                      `query:"per_page"`
	Type    param.Field[ConnectivityDirectoryServiceListParamsType] `query:"type"`
}

// URLQuery serializes [ConnectivityDirectoryServiceListParams]'s query parameters
// as `url.Values`.
func (r ConnectivityDirectoryServiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConnectivityDirectoryServiceListParamsType string

const (
	ConnectivityDirectoryServiceListParamsTypeHTTP ConnectivityDirectoryServiceListParamsType = "http"
)

func (r ConnectivityDirectoryServiceListParamsType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceListParamsTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectivityDirectoryServiceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectivityDirectoryServiceGetResponseEnvelope struct {
	Errors   []ConnectivityDirectoryServiceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectivityDirectoryServiceGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ConnectivityDirectoryServiceGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ConnectivityDirectoryServiceGetResponse                `json:"result"`
	JSON    connectivityDirectoryServiceGetResponseEnvelopeJSON    `json:"-"`
}

// connectivityDirectoryServiceGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ConnectivityDirectoryServiceGetResponseEnvelope]
type connectivityDirectoryServiceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           ConnectivityDirectoryServiceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             connectivityDirectoryServiceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// connectivityDirectoryServiceGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ConnectivityDirectoryServiceGetResponseEnvelopeErrors]
type connectivityDirectoryServiceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    connectivityDirectoryServiceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// connectivityDirectoryServiceGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseEnvelopeErrorsSource]
type connectivityDirectoryServiceGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           ConnectivityDirectoryServiceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             connectivityDirectoryServiceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// connectivityDirectoryServiceGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ConnectivityDirectoryServiceGetResponseEnvelopeMessages]
type connectivityDirectoryServiceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectivityDirectoryServiceGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    connectivityDirectoryServiceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// connectivityDirectoryServiceGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ConnectivityDirectoryServiceGetResponseEnvelopeMessagesSource]
type connectivityDirectoryServiceGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivityDirectoryServiceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivityDirectoryServiceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConnectivityDirectoryServiceGetResponseEnvelopeSuccess bool

const (
	ConnectivityDirectoryServiceGetResponseEnvelopeSuccessTrue ConnectivityDirectoryServiceGetResponseEnvelopeSuccess = true
)

func (r ConnectivityDirectoryServiceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
