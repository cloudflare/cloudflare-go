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

// Create connectivity service
func (r *DirectoryServiceService) New(ctx context.Context, params DirectoryServiceNewParams, opts ...option.RequestOption) (res *DirectoryServiceNewResponse, err error) {
	var env DirectoryServiceNewResponseEnvelope
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
func (r *DirectoryServiceService) Update(ctx context.Context, serviceID string, params DirectoryServiceUpdateParams, opts ...option.RequestOption) (res *DirectoryServiceUpdateResponse, err error) {
	var env DirectoryServiceUpdateResponseEnvelope
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
func (r *DirectoryServiceService) List(ctx context.Context, params DirectoryServiceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DirectoryServiceListResponse], err error) {
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
func (r *DirectoryServiceService) ListAutoPaging(ctx context.Context, params DirectoryServiceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DirectoryServiceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete connectivity service
func (r *DirectoryServiceService) Delete(ctx context.Context, serviceID string, body DirectoryServiceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
func (r *DirectoryServiceService) Get(ctx context.Context, serviceID string, query DirectoryServiceGetParams, opts ...option.RequestOption) (res *DirectoryServiceGetResponse, err error) {
	var env DirectoryServiceGetResponseEnvelope
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

type DirectoryServiceNewResponse struct {
	Host      DirectoryServiceNewResponseHost `json:"host,required"`
	Name      string                          `json:"name,required"`
	Type      DirectoryServiceNewResponseType `json:"type,required"`
	CreatedAt time.Time                       `json:"created_at" format:"date-time"`
	HTTPPort  int64                           `json:"http_port,nullable"`
	HTTPSPort int64                           `json:"https_port,nullable"`
	ServiceID string                          `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      directoryServiceNewResponseJSON `json:"-"`
}

// directoryServiceNewResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceNewResponse]
type directoryServiceNewResponseJSON struct {
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

func (r *DirectoryServiceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseHostInfraIPv4HostNetwork],
	// [DirectoryServiceNewResponseHostInfraIPv6HostNetwork],
	// [DirectoryServiceNewResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                         `json:"resolver_network"`
	JSON            directoryServiceNewResponseHostJSON `json:"-"`
	union           DirectoryServiceNewResponseHostUnion
}

// directoryServiceNewResponseHostJSON contains the JSON metadata for the struct
// [DirectoryServiceNewResponseHost]
type directoryServiceNewResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceNewResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceNewResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceNewResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceNewResponseHostUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceNewResponseHostInfraIPv4Host],
// [DirectoryServiceNewResponseHostInfraIPv6Host],
// [DirectoryServiceNewResponseHostInfraDualStackHost],
// [DirectoryServiceNewResponseHostInfraHostnameHost].
func (r DirectoryServiceNewResponseHost) AsUnion() DirectoryServiceNewResponseHostUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceNewResponseHostInfraIPv4Host],
// [DirectoryServiceNewResponseHostInfraIPv6Host],
// [DirectoryServiceNewResponseHostInfraDualStackHost] or
// [DirectoryServiceNewResponseHostInfraHostnameHost].
type DirectoryServiceNewResponseHostUnion interface {
	implementsDirectoryServiceNewResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceNewResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceNewResponseHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceNewResponseHostInfraIPv4Host struct {
	IPV4    string                                              `json:"ipv4,required"`
	Network DirectoryServiceNewResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    directoryServiceNewResponseHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceNewResponseHostInfraIPv4HostJSON contains the JSON metadata for
// the struct [DirectoryServiceNewResponseHostInfraIPv4Host]
type directoryServiceNewResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseHostInfraIPv4Host) implementsDirectoryServiceNewResponseHost() {}

type DirectoryServiceNewResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                  `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceNewResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseHostInfraIPv4HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceNewResponseHostInfraIPv4HostNetwork]
type directoryServiceNewResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseHostInfraIPv6Host struct {
	IPV6    string                                              `json:"ipv6,required"`
	Network DirectoryServiceNewResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    directoryServiceNewResponseHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceNewResponseHostInfraIPv6HostJSON contains the JSON metadata for
// the struct [DirectoryServiceNewResponseHostInfraIPv6Host]
type directoryServiceNewResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseHostInfraIPv6Host) implementsDirectoryServiceNewResponseHost() {}

type DirectoryServiceNewResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                  `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceNewResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseHostInfraIPv6HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceNewResponseHostInfraIPv6HostNetwork]
type directoryServiceNewResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseHostInfraDualStackHost struct {
	IPV4    string                                                   `json:"ipv4,required"`
	IPV6    string                                                   `json:"ipv6,required"`
	Network DirectoryServiceNewResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    directoryServiceNewResponseHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceNewResponseHostInfraDualStackHostJSON contains the JSON metadata
// for the struct [DirectoryServiceNewResponseHostInfraDualStackHost]
type directoryServiceNewResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseHostInfraDualStackHost) implementsDirectoryServiceNewResponseHost() {
}

type DirectoryServiceNewResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                       `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceNewResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceNewResponseHostInfraDualStackHostNetworkJSON contains the JSON
// metadata for the struct
// [DirectoryServiceNewResponseHostInfraDualStackHostNetwork]
type directoryServiceNewResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseHostInfraHostnameHost struct {
	Hostname        string                                                          `json:"hostname,required"`
	ResolverNetwork DirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            directoryServiceNewResponseHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceNewResponseHostInfraHostnameHostJSON contains the JSON metadata
// for the struct [DirectoryServiceNewResponseHostInfraHostnameHost]
type directoryServiceNewResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceNewResponseHostInfraHostnameHost) implementsDirectoryServiceNewResponseHost() {
}

type DirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                              `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                            `json:"resolver_ips,nullable"`
	JSON        directoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON contains the
// JSON metadata for the struct
// [DirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork]
type directoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceNewResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceNewResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceNewResponseType string

const (
	DirectoryServiceNewResponseTypeHTTP DirectoryServiceNewResponseType = "http"
)

func (r DirectoryServiceNewResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceUpdateResponse struct {
	Host      DirectoryServiceUpdateResponseHost `json:"host,required"`
	Name      string                             `json:"name,required"`
	Type      DirectoryServiceUpdateResponseType `json:"type,required"`
	CreatedAt time.Time                          `json:"created_at" format:"date-time"`
	HTTPPort  int64                              `json:"http_port,nullable"`
	HTTPSPort int64                              `json:"https_port,nullable"`
	ServiceID string                             `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      directoryServiceUpdateResponseJSON `json:"-"`
}

// directoryServiceUpdateResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponse]
type directoryServiceUpdateResponseJSON struct {
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

func (r *DirectoryServiceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseHostInfraIPv4HostNetwork],
	// [DirectoryServiceUpdateResponseHostInfraIPv6HostNetwork],
	// [DirectoryServiceUpdateResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                            `json:"resolver_network"`
	JSON            directoryServiceUpdateResponseHostJSON `json:"-"`
	union           DirectoryServiceUpdateResponseHostUnion
}

// directoryServiceUpdateResponseHostJSON contains the JSON metadata for the struct
// [DirectoryServiceUpdateResponseHost]
type directoryServiceUpdateResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceUpdateResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceUpdateResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceUpdateResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceUpdateResponseHostUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceUpdateResponseHostInfraIPv4Host],
// [DirectoryServiceUpdateResponseHostInfraIPv6Host],
// [DirectoryServiceUpdateResponseHostInfraDualStackHost],
// [DirectoryServiceUpdateResponseHostInfraHostnameHost].
func (r DirectoryServiceUpdateResponseHost) AsUnion() DirectoryServiceUpdateResponseHostUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceUpdateResponseHostInfraIPv4Host],
// [DirectoryServiceUpdateResponseHostInfraIPv6Host],
// [DirectoryServiceUpdateResponseHostInfraDualStackHost] or
// [DirectoryServiceUpdateResponseHostInfraHostnameHost].
type DirectoryServiceUpdateResponseHostUnion interface {
	implementsDirectoryServiceUpdateResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceUpdateResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceUpdateResponseHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceUpdateResponseHostInfraIPv4Host struct {
	IPV4    string                                                 `json:"ipv4,required"`
	Network DirectoryServiceUpdateResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    directoryServiceUpdateResponseHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceUpdateResponseHostInfraIPv4HostJSON contains the JSON metadata
// for the struct [DirectoryServiceUpdateResponseHostInfraIPv4Host]
type directoryServiceUpdateResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseHostInfraIPv4Host) implementsDirectoryServiceUpdateResponseHost() {
}

type DirectoryServiceUpdateResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                     `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseHostInfraIPv4HostNetwork]
type directoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseHostInfraIPv6Host struct {
	IPV6    string                                                 `json:"ipv6,required"`
	Network DirectoryServiceUpdateResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    directoryServiceUpdateResponseHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceUpdateResponseHostInfraIPv6HostJSON contains the JSON metadata
// for the struct [DirectoryServiceUpdateResponseHostInfraIPv6Host]
type directoryServiceUpdateResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseHostInfraIPv6Host) implementsDirectoryServiceUpdateResponseHost() {
}

type DirectoryServiceUpdateResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                     `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseHostInfraIPv6HostNetwork]
type directoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseHostInfraDualStackHost struct {
	IPV4    string                                                      `json:"ipv4,required"`
	IPV6    string                                                      `json:"ipv6,required"`
	Network DirectoryServiceUpdateResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    directoryServiceUpdateResponseHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceUpdateResponseHostInfraDualStackHostJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseHostInfraDualStackHost]
type directoryServiceUpdateResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseHostInfraDualStackHost) implementsDirectoryServiceUpdateResponseHost() {
}

type DirectoryServiceUpdateResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                          `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON contains the
// JSON metadata for the struct
// [DirectoryServiceUpdateResponseHostInfraDualStackHostNetwork]
type directoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseHostInfraHostnameHost struct {
	Hostname        string                                                             `json:"hostname,required"`
	ResolverNetwork DirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            directoryServiceUpdateResponseHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceUpdateResponseHostInfraHostnameHostJSON contains the JSON
// metadata for the struct [DirectoryServiceUpdateResponseHostInfraHostnameHost]
type directoryServiceUpdateResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceUpdateResponseHostInfraHostnameHost) implementsDirectoryServiceUpdateResponseHost() {
}

type DirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                                 `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                               `json:"resolver_ips,nullable"`
	JSON        directoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON contains
// the JSON metadata for the struct
// [DirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork]
type directoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceUpdateResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceUpdateResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceUpdateResponseType string

const (
	DirectoryServiceUpdateResponseTypeHTTP DirectoryServiceUpdateResponseType = "http"
)

func (r DirectoryServiceUpdateResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceListResponse struct {
	Host      DirectoryServiceListResponseHost `json:"host,required"`
	Name      string                           `json:"name,required"`
	Type      DirectoryServiceListResponseType `json:"type,required"`
	CreatedAt time.Time                        `json:"created_at" format:"date-time"`
	HTTPPort  int64                            `json:"http_port,nullable"`
	HTTPSPort int64                            `json:"https_port,nullable"`
	ServiceID string                           `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      directoryServiceListResponseJSON `json:"-"`
}

// directoryServiceListResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceListResponse]
type directoryServiceListResponseJSON struct {
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

func (r *DirectoryServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceListResponseHostInfraIPv4HostNetwork],
	// [DirectoryServiceListResponseHostInfraIPv6HostNetwork],
	// [DirectoryServiceListResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceListResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                          `json:"resolver_network"`
	JSON            directoryServiceListResponseHostJSON `json:"-"`
	union           DirectoryServiceListResponseHostUnion
}

// directoryServiceListResponseHostJSON contains the JSON metadata for the struct
// [DirectoryServiceListResponseHost]
type directoryServiceListResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceListResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceListResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceListResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceListResponseHostUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceListResponseHostInfraIPv4Host],
// [DirectoryServiceListResponseHostInfraIPv6Host],
// [DirectoryServiceListResponseHostInfraDualStackHost],
// [DirectoryServiceListResponseHostInfraHostnameHost].
func (r DirectoryServiceListResponseHost) AsUnion() DirectoryServiceListResponseHostUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceListResponseHostInfraIPv4Host],
// [DirectoryServiceListResponseHostInfraIPv6Host],
// [DirectoryServiceListResponseHostInfraDualStackHost] or
// [DirectoryServiceListResponseHostInfraHostnameHost].
type DirectoryServiceListResponseHostUnion interface {
	implementsDirectoryServiceListResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceListResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceListResponseHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceListResponseHostInfraIPv4Host struct {
	IPV4    string                                               `json:"ipv4,required"`
	Network DirectoryServiceListResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    directoryServiceListResponseHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceListResponseHostInfraIPv4HostJSON contains the JSON metadata for
// the struct [DirectoryServiceListResponseHostInfraIPv4Host]
type directoryServiceListResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseHostInfraIPv4Host) implementsDirectoryServiceListResponseHost() {}

type DirectoryServiceListResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                   `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceListResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceListResponseHostInfraIPv4HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceListResponseHostInfraIPv4HostNetwork]
type directoryServiceListResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseHostInfraIPv6Host struct {
	IPV6    string                                               `json:"ipv6,required"`
	Network DirectoryServiceListResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    directoryServiceListResponseHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceListResponseHostInfraIPv6HostJSON contains the JSON metadata for
// the struct [DirectoryServiceListResponseHostInfraIPv6Host]
type directoryServiceListResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseHostInfraIPv6Host) implementsDirectoryServiceListResponseHost() {}

type DirectoryServiceListResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                   `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceListResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceListResponseHostInfraIPv6HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceListResponseHostInfraIPv6HostNetwork]
type directoryServiceListResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseHostInfraDualStackHost struct {
	IPV4    string                                                    `json:"ipv4,required"`
	IPV6    string                                                    `json:"ipv6,required"`
	Network DirectoryServiceListResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    directoryServiceListResponseHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceListResponseHostInfraDualStackHostJSON contains the JSON
// metadata for the struct [DirectoryServiceListResponseHostInfraDualStackHost]
type directoryServiceListResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseHostInfraDualStackHost) implementsDirectoryServiceListResponseHost() {
}

type DirectoryServiceListResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                        `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceListResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceListResponseHostInfraDualStackHostNetworkJSON contains the JSON
// metadata for the struct
// [DirectoryServiceListResponseHostInfraDualStackHostNetwork]
type directoryServiceListResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseHostInfraHostnameHost struct {
	Hostname        string                                                           `json:"hostname,required"`
	ResolverNetwork DirectoryServiceListResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            directoryServiceListResponseHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceListResponseHostInfraHostnameHostJSON contains the JSON metadata
// for the struct [DirectoryServiceListResponseHostInfraHostnameHost]
type directoryServiceListResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceListResponseHostInfraHostnameHost) implementsDirectoryServiceListResponseHost() {
}

type DirectoryServiceListResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                               `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                             `json:"resolver_ips,nullable"`
	JSON        directoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON contains
// the JSON metadata for the struct
// [DirectoryServiceListResponseHostInfraHostnameHostResolverNetwork]
type directoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceListResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceListResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceListResponseType string

const (
	DirectoryServiceListResponseTypeHTTP DirectoryServiceListResponseType = "http"
)

func (r DirectoryServiceListResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceListResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceGetResponse struct {
	Host      DirectoryServiceGetResponseHost `json:"host,required"`
	Name      string                          `json:"name,required"`
	Type      DirectoryServiceGetResponseType `json:"type,required"`
	CreatedAt time.Time                       `json:"created_at" format:"date-time"`
	HTTPPort  int64                           `json:"http_port,nullable"`
	HTTPSPort int64                           `json:"https_port,nullable"`
	ServiceID string                          `json:"service_id" format:"uuid"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      directoryServiceGetResponseJSON `json:"-"`
}

// directoryServiceGetResponseJSON contains the JSON metadata for the struct
// [DirectoryServiceGetResponse]
type directoryServiceGetResponseJSON struct {
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

func (r *DirectoryServiceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseHost struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	IPV6     string `json:"ipv6"`
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseHostInfraIPv4HostNetwork],
	// [DirectoryServiceGetResponseHostInfraIPv6HostNetwork],
	// [DirectoryServiceGetResponseHostInfraDualStackHostNetwork].
	Network interface{} `json:"network"`
	// This field can have the runtime type of
	// [DirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork].
	ResolverNetwork interface{}                         `json:"resolver_network"`
	JSON            directoryServiceGetResponseHostJSON `json:"-"`
	union           DirectoryServiceGetResponseHostUnion
}

// directoryServiceGetResponseHostJSON contains the JSON metadata for the struct
// [DirectoryServiceGetResponseHost]
type directoryServiceGetResponseHostJSON struct {
	Hostname        apijson.Field
	IPV4            apijson.Field
	IPV6            apijson.Field
	Network         apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r directoryServiceGetResponseHostJSON) RawJSON() string {
	return r.raw
}

func (r *DirectoryServiceGetResponseHost) UnmarshalJSON(data []byte) (err error) {
	*r = DirectoryServiceGetResponseHost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DirectoryServiceGetResponseHostUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DirectoryServiceGetResponseHostInfraIPv4Host],
// [DirectoryServiceGetResponseHostInfraIPv6Host],
// [DirectoryServiceGetResponseHostInfraDualStackHost],
// [DirectoryServiceGetResponseHostInfraHostnameHost].
func (r DirectoryServiceGetResponseHost) AsUnion() DirectoryServiceGetResponseHostUnion {
	return r.union
}

// Union satisfied by [DirectoryServiceGetResponseHostInfraIPv4Host],
// [DirectoryServiceGetResponseHostInfraIPv6Host],
// [DirectoryServiceGetResponseHostInfraDualStackHost] or
// [DirectoryServiceGetResponseHostInfraHostnameHost].
type DirectoryServiceGetResponseHostUnion interface {
	implementsDirectoryServiceGetResponseHost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DirectoryServiceGetResponseHostUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseHostInfraIPv4Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseHostInfraIPv6Host{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseHostInfraDualStackHost{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DirectoryServiceGetResponseHostInfraHostnameHost{}),
		},
	)
}

type DirectoryServiceGetResponseHostInfraIPv4Host struct {
	IPV4    string                                              `json:"ipv4,required"`
	Network DirectoryServiceGetResponseHostInfraIPv4HostNetwork `json:"network,required"`
	JSON    directoryServiceGetResponseHostInfraIPv4HostJSON    `json:"-"`
}

// directoryServiceGetResponseHostInfraIPv4HostJSON contains the JSON metadata for
// the struct [DirectoryServiceGetResponseHostInfraIPv4Host]
type directoryServiceGetResponseHostInfraIPv4HostJSON struct {
	IPV4        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraIPv4Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraIPv4HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseHostInfraIPv4Host) implementsDirectoryServiceGetResponseHost() {}

type DirectoryServiceGetResponseHostInfraIPv4HostNetwork struct {
	TunnelID string                                                  `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceGetResponseHostInfraIPv4HostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseHostInfraIPv4HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceGetResponseHostInfraIPv4HostNetwork]
type directoryServiceGetResponseHostInfraIPv4HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraIPv4HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraIPv4HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseHostInfraIPv6Host struct {
	IPV6    string                                              `json:"ipv6,required"`
	Network DirectoryServiceGetResponseHostInfraIPv6HostNetwork `json:"network,required"`
	JSON    directoryServiceGetResponseHostInfraIPv6HostJSON    `json:"-"`
}

// directoryServiceGetResponseHostInfraIPv6HostJSON contains the JSON metadata for
// the struct [DirectoryServiceGetResponseHostInfraIPv6Host]
type directoryServiceGetResponseHostInfraIPv6HostJSON struct {
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraIPv6Host) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraIPv6HostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseHostInfraIPv6Host) implementsDirectoryServiceGetResponseHost() {}

type DirectoryServiceGetResponseHostInfraIPv6HostNetwork struct {
	TunnelID string                                                  `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceGetResponseHostInfraIPv6HostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseHostInfraIPv6HostNetworkJSON contains the JSON
// metadata for the struct [DirectoryServiceGetResponseHostInfraIPv6HostNetwork]
type directoryServiceGetResponseHostInfraIPv6HostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraIPv6HostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraIPv6HostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseHostInfraDualStackHost struct {
	IPV4    string                                                   `json:"ipv4,required"`
	IPV6    string                                                   `json:"ipv6,required"`
	Network DirectoryServiceGetResponseHostInfraDualStackHostNetwork `json:"network,required"`
	JSON    directoryServiceGetResponseHostInfraDualStackHostJSON    `json:"-"`
}

// directoryServiceGetResponseHostInfraDualStackHostJSON contains the JSON metadata
// for the struct [DirectoryServiceGetResponseHostInfraDualStackHost]
type directoryServiceGetResponseHostInfraDualStackHostJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraDualStackHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraDualStackHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseHostInfraDualStackHost) implementsDirectoryServiceGetResponseHost() {
}

type DirectoryServiceGetResponseHostInfraDualStackHostNetwork struct {
	TunnelID string                                                       `json:"tunnel_id,required" format:"uuid"`
	JSON     directoryServiceGetResponseHostInfraDualStackHostNetworkJSON `json:"-"`
}

// directoryServiceGetResponseHostInfraDualStackHostNetworkJSON contains the JSON
// metadata for the struct
// [DirectoryServiceGetResponseHostInfraDualStackHostNetwork]
type directoryServiceGetResponseHostInfraDualStackHostNetworkJSON struct {
	TunnelID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraDualStackHostNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraDualStackHostNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseHostInfraHostnameHost struct {
	Hostname        string                                                          `json:"hostname,required"`
	ResolverNetwork DirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork `json:"resolver_network,required"`
	JSON            directoryServiceGetResponseHostInfraHostnameHostJSON            `json:"-"`
}

// directoryServiceGetResponseHostInfraHostnameHostJSON contains the JSON metadata
// for the struct [DirectoryServiceGetResponseHostInfraHostnameHost]
type directoryServiceGetResponseHostInfraHostnameHostJSON struct {
	Hostname        apijson.Field
	ResolverNetwork apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraHostnameHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraHostnameHostJSON) RawJSON() string {
	return r.raw
}

func (r DirectoryServiceGetResponseHostInfraHostnameHost) implementsDirectoryServiceGetResponseHost() {
}

type DirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork struct {
	TunnelID    string                                                              `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs []string                                                            `json:"resolver_ips,nullable"`
	JSON        directoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON `json:"-"`
}

// directoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON contains the
// JSON metadata for the struct
// [DirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork]
type directoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON struct {
	TunnelID    apijson.Field
	ResolverIPs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DirectoryServiceGetResponseHostInfraHostnameHostResolverNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r directoryServiceGetResponseHostInfraHostnameHostResolverNetworkJSON) RawJSON() string {
	return r.raw
}

type DirectoryServiceGetResponseType string

const (
	DirectoryServiceGetResponseTypeHTTP DirectoryServiceGetResponseType = "http"
)

func (r DirectoryServiceGetResponseType) IsKnown() bool {
	switch r {
	case DirectoryServiceGetResponseTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceNewParams struct {
	// Account identifier
	AccountID param.Field[string]                             `path:"account_id,required"`
	Host      param.Field[DirectoryServiceNewParamsHostUnion] `json:"host,required"`
	Name      param.Field[string]                             `json:"name,required"`
	Type      param.Field[DirectoryServiceNewParamsType]      `json:"type,required"`
	HTTPPort  param.Field[int64]                              `json:"http_port"`
	HTTPSPort param.Field[int64]                              `json:"https_port"`
}

func (r DirectoryServiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceNewParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsHost) implementsDirectoryServiceNewParamsHostUnion() {}

// Satisfied by [connectivity.DirectoryServiceNewParamsHostInfraIPv4Host],
// [connectivity.DirectoryServiceNewParamsHostInfraIPv6Host],
// [connectivity.DirectoryServiceNewParamsHostInfraDualStackHost],
// [connectivity.DirectoryServiceNewParamsHostInfraHostnameHost],
// [DirectoryServiceNewParamsHost].
type DirectoryServiceNewParamsHostUnion interface {
	implementsDirectoryServiceNewParamsHostUnion()
}

type DirectoryServiceNewParamsHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                            `json:"ipv4,required"`
	Network param.Field[DirectoryServiceNewParamsHostInfraIPv4HostNetwork] `json:"network,required"`
}

func (r DirectoryServiceNewParamsHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsHostInfraIPv4Host) implementsDirectoryServiceNewParamsHostUnion() {}

type DirectoryServiceNewParamsHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r DirectoryServiceNewParamsHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                            `json:"ipv6,required"`
	Network param.Field[DirectoryServiceNewParamsHostInfraIPv6HostNetwork] `json:"network,required"`
}

func (r DirectoryServiceNewParamsHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsHostInfraIPv6Host) implementsDirectoryServiceNewParamsHostUnion() {}

type DirectoryServiceNewParamsHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r DirectoryServiceNewParamsHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                 `json:"ipv4,required"`
	IPV6    param.Field[string]                                                 `json:"ipv6,required"`
	Network param.Field[DirectoryServiceNewParamsHostInfraDualStackHostNetwork] `json:"network,required"`
}

func (r DirectoryServiceNewParamsHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsHostInfraDualStackHost) implementsDirectoryServiceNewParamsHostUnion() {
}

type DirectoryServiceNewParamsHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r DirectoryServiceNewParamsHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                        `json:"hostname,required"`
	ResolverNetwork param.Field[DirectoryServiceNewParamsHostInfraHostnameHostResolverNetwork] `json:"resolver_network,required"`
}

func (r DirectoryServiceNewParamsHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceNewParamsHostInfraHostnameHost) implementsDirectoryServiceNewParamsHostUnion() {
}

type DirectoryServiceNewParamsHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r DirectoryServiceNewParamsHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsType string

const (
	DirectoryServiceNewParamsTypeHTTP DirectoryServiceNewParamsType = "http"
)

func (r DirectoryServiceNewParamsType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewParamsTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceNewResponseEnvelope struct {
	Errors   []DirectoryServiceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DirectoryServiceNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DirectoryServiceNewResponseEnvelopeSuccess `json:"success,required"`
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
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
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
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
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
	AccountID param.Field[string]                                `path:"account_id,required"`
	Host      param.Field[DirectoryServiceUpdateParamsHostUnion] `json:"host,required"`
	Name      param.Field[string]                                `json:"name,required"`
	Type      param.Field[DirectoryServiceUpdateParamsType]      `json:"type,required"`
	HTTPPort  param.Field[int64]                                 `json:"http_port"`
	HTTPSPort param.Field[int64]                                 `json:"https_port"`
}

func (r DirectoryServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceUpdateParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsHost) implementsDirectoryServiceUpdateParamsHostUnion() {}

// Satisfied by [connectivity.DirectoryServiceUpdateParamsHostInfraIPv4Host],
// [connectivity.DirectoryServiceUpdateParamsHostInfraIPv6Host],
// [connectivity.DirectoryServiceUpdateParamsHostInfraDualStackHost],
// [connectivity.DirectoryServiceUpdateParamsHostInfraHostnameHost],
// [DirectoryServiceUpdateParamsHost].
type DirectoryServiceUpdateParamsHostUnion interface {
	implementsDirectoryServiceUpdateParamsHostUnion()
}

type DirectoryServiceUpdateParamsHostInfraIPv4Host struct {
	IPV4    param.Field[string]                                               `json:"ipv4,required"`
	Network param.Field[DirectoryServiceUpdateParamsHostInfraIPv4HostNetwork] `json:"network,required"`
}

func (r DirectoryServiceUpdateParamsHostInfraIPv4Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsHostInfraIPv4Host) implementsDirectoryServiceUpdateParamsHostUnion() {
}

type DirectoryServiceUpdateParamsHostInfraIPv4HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsHostInfraIPv4HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsHostInfraIPv6Host struct {
	IPV6    param.Field[string]                                               `json:"ipv6,required"`
	Network param.Field[DirectoryServiceUpdateParamsHostInfraIPv6HostNetwork] `json:"network,required"`
}

func (r DirectoryServiceUpdateParamsHostInfraIPv6Host) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsHostInfraIPv6Host) implementsDirectoryServiceUpdateParamsHostUnion() {
}

type DirectoryServiceUpdateParamsHostInfraIPv6HostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsHostInfraIPv6HostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsHostInfraDualStackHost struct {
	IPV4    param.Field[string]                                                    `json:"ipv4,required"`
	IPV6    param.Field[string]                                                    `json:"ipv6,required"`
	Network param.Field[DirectoryServiceUpdateParamsHostInfraDualStackHostNetwork] `json:"network,required"`
}

func (r DirectoryServiceUpdateParamsHostInfraDualStackHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsHostInfraDualStackHost) implementsDirectoryServiceUpdateParamsHostUnion() {
}

type DirectoryServiceUpdateParamsHostInfraDualStackHostNetwork struct {
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
}

func (r DirectoryServiceUpdateParamsHostInfraDualStackHostNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsHostInfraHostnameHost struct {
	Hostname        param.Field[string]                                                           `json:"hostname,required"`
	ResolverNetwork param.Field[DirectoryServiceUpdateParamsHostInfraHostnameHostResolverNetwork] `json:"resolver_network,required"`
}

func (r DirectoryServiceUpdateParamsHostInfraHostnameHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DirectoryServiceUpdateParamsHostInfraHostnameHost) implementsDirectoryServiceUpdateParamsHostUnion() {
}

type DirectoryServiceUpdateParamsHostInfraHostnameHostResolverNetwork struct {
	TunnelID    param.Field[string]   `json:"tunnel_id,required" format:"uuid"`
	ResolverIPs param.Field[[]string] `json:"resolver_ips"`
}

func (r DirectoryServiceUpdateParamsHostInfraHostnameHostResolverNetwork) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsType string

const (
	DirectoryServiceUpdateParamsTypeHTTP DirectoryServiceUpdateParamsType = "http"
)

func (r DirectoryServiceUpdateParamsType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateParamsTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceUpdateResponseEnvelope struct {
	Errors   []DirectoryServiceUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DirectoryServiceUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DirectoryServiceUpdateResponseEnvelopeSuccess `json:"success,required"`
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
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
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
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
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
	DirectoryServiceListParamsTypeHTTP DirectoryServiceListParamsType = "http"
)

func (r DirectoryServiceListParamsType) IsKnown() bool {
	switch r {
	case DirectoryServiceListParamsTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DirectoryServiceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DirectoryServiceGetResponseEnvelope struct {
	Errors   []DirectoryServiceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DirectoryServiceGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DirectoryServiceGetResponseEnvelopeSuccess `json:"success,required"`
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
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
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
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
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
