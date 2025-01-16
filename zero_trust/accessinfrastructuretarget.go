// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccessInfrastructureTargetService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessInfrastructureTargetService] method instead.
type AccessInfrastructureTargetService struct {
	Options []option.RequestOption
}

// NewAccessInfrastructureTargetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessInfrastructureTargetService(opts ...option.RequestOption) (r *AccessInfrastructureTargetService) {
	r = &AccessInfrastructureTargetService{}
	r.Options = opts
	return
}

// Create new target
func (r *AccessInfrastructureTargetService) New(ctx context.Context, params AccessInfrastructureTargetNewParams, opts ...option.RequestOption) (res *AccessInfrastructureTargetNewResponse, err error) {
	var env AccessInfrastructureTargetNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update target
func (r *AccessInfrastructureTargetService) Update(ctx context.Context, targetID string, params AccessInfrastructureTargetUpdateParams, opts ...option.RequestOption) (res *AccessInfrastructureTargetUpdateResponse, err error) {
	var env AccessInfrastructureTargetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", params.AccountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all targets
func (r *AccessInfrastructureTargetService) List(ctx context.Context, params AccessInfrastructureTargetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessInfrastructureTargetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets", params.AccountID)
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

// List all targets
func (r *AccessInfrastructureTargetService) ListAutoPaging(ctx context.Context, params AccessInfrastructureTargetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessInfrastructureTargetListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete target
func (r *AccessInfrastructureTargetService) Delete(ctx context.Context, targetID string, body AccessInfrastructureTargetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", body.AccountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Removes one or more targets.
func (r *AccessInfrastructureTargetService) BulkDelete(ctx context.Context, body AccessInfrastructureTargetBulkDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Adds one or more targets.
func (r *AccessInfrastructureTargetService) BulkUpdate(ctx context.Context, params AccessInfrastructureTargetBulkUpdateParams, opts ...option.RequestOption) (res *[]AccessInfrastructureTargetBulkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/batch", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Get target
func (r *AccessInfrastructureTargetService) Get(ctx context.Context, targetID string, query AccessInfrastructureTargetGetParams, opts ...option.RequestOption) (res *AccessInfrastructureTargetGetResponse, err error) {
	var env AccessInfrastructureTargetGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", query.AccountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessInfrastructureTargetNewResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetNewResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                 `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetNewResponseJSON `json:"-"`
}

// accessInfrastructureTargetNewResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetNewResponse]
type accessInfrastructureTargetNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetNewResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetNewResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetNewResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetNewResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetNewResponseIPJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetNewResponseIP]
type accessInfrastructureTargetNewResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetNewResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetNewResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetNewResponseIpipv4JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetNewResponseIPIPV4]
type accessInfrastructureTargetNewResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetNewResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetNewResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetNewResponseIpipv6JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetNewResponseIPIPV6]
type accessInfrastructureTargetNewResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetUpdateResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetUpdateResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                    `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetUpdateResponseJSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetUpdateResponse]
type accessInfrastructureTargetUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetUpdateResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetUpdateResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetUpdateResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetUpdateResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseIPJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetUpdateResponseIP]
type accessInfrastructureTargetUpdateResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetUpdateResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                             `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetUpdateResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseIpipv4JSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetUpdateResponseIPIPV4]
type accessInfrastructureTargetUpdateResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetUpdateResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                             `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetUpdateResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetUpdateResponseIpipv6JSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetUpdateResponseIPIPV6]
type accessInfrastructureTargetUpdateResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetListResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetListResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                  `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetListResponseJSON `json:"-"`
}

// accessInfrastructureTargetListResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetListResponse]
type accessInfrastructureTargetListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetListResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetListResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetListResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetListResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetListResponseIPJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetListResponseIP]
type accessInfrastructureTargetListResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetListResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                           `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetListResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetListResponseIpipv4JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetListResponseIPIPV4]
type accessInfrastructureTargetListResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetListResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                           `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetListResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetListResponseIpipv6JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetListResponseIPIPV6]
type accessInfrastructureTargetListResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetListResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetListResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetBulkUpdateResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetBulkUpdateResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                        `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetBulkUpdateResponseJSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetBulkUpdateResponse]
type accessInfrastructureTargetBulkUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetBulkUpdateResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetBulkUpdateResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetBulkUpdateResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetBulkUpdateResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseIPJSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetBulkUpdateResponseIP]
type accessInfrastructureTargetBulkUpdateResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetBulkUpdateResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                                 `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetBulkUpdateResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseIpipv4JSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetBulkUpdateResponseIPIPV4]
type accessInfrastructureTargetBulkUpdateResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetBulkUpdateResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                                 `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetBulkUpdateResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetBulkUpdateResponseIpipv6JSON contains the JSON
// metadata for the struct [AccessInfrastructureTargetBulkUpdateResponseIPIPV6]
type accessInfrastructureTargetBulkUpdateResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetBulkUpdateResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetBulkUpdateResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetGetResponse struct {
	// Target identifier
	ID string `json:"id,required" format:"uuid"`
	// Date and time at which the target was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A non-unique field that refers to a target
	Hostname string `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP AccessInfrastructureTargetGetResponseIP `json:"ip,required"`
	// Date and time at which the target was modified
	ModifiedAt time.Time                                 `json:"modified_at,required" format:"date-time"`
	JSON       accessInfrastructureTargetGetResponseJSON `json:"-"`
}

// accessInfrastructureTargetGetResponseJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetGetResponse]
type accessInfrastructureTargetGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseJSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetGetResponseIP struct {
	// The target's IPv4 address
	IPV4 AccessInfrastructureTargetGetResponseIPIPV4 `json:"ipv4"`
	// The target's IPv6 address
	IPV6 AccessInfrastructureTargetGetResponseIPIPV6 `json:"ipv6"`
	JSON accessInfrastructureTargetGetResponseIPJSON `json:"-"`
}

// accessInfrastructureTargetGetResponseIPJSON contains the JSON metadata for the
// struct [AccessInfrastructureTargetGetResponseIP]
type accessInfrastructureTargetGetResponseIPJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseIPJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type AccessInfrastructureTargetGetResponseIPIPV4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetGetResponseIpipv4JSON `json:"-"`
}

// accessInfrastructureTargetGetResponseIpipv4JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetGetResponseIPIPV4]
type accessInfrastructureTargetGetResponseIpipv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseIPIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseIpipv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type AccessInfrastructureTargetGetResponseIPIPV6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string                                          `json:"virtual_network_id" format:"uuid"`
	JSON             accessInfrastructureTargetGetResponseIpipv6JSON `json:"-"`
}

// accessInfrastructureTargetGetResponseIpipv6JSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetGetResponseIPIPV6]
type accessInfrastructureTargetGetResponseIpipv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseIPIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseIpipv6JSON) RawJSON() string {
	return r.raw
}

type AccessInfrastructureTargetNewParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[AccessInfrastructureTargetNewParamsIP] `json:"ip,required"`
}

func (r AccessInfrastructureTargetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetNewParamsIP struct {
	// The target's IPv4 address
	IPV4 param.Field[AccessInfrastructureTargetNewParamsIPIPV4] `json:"ipv4"`
	// The target's IPv6 address
	IPV6 param.Field[AccessInfrastructureTargetNewParamsIPIPV6] `json:"ipv6"`
}

func (r AccessInfrastructureTargetNewParamsIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv4 address
type AccessInfrastructureTargetNewParamsIPIPV4 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetNewParamsIPIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv6 address
type AccessInfrastructureTargetNewParamsIPIPV6 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetNewParamsIPIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessInfrastructureTargetNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessInfrastructureTargetNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessInfrastructureTargetNewResponse                `json:"result"`
	JSON    accessInfrastructureTargetNewResponseEnvelopeJSON    `json:"-"`
}

// accessInfrastructureTargetNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetNewResponseEnvelope]
type accessInfrastructureTargetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessInfrastructureTargetNewResponseEnvelopeSuccess bool

const (
	AccessInfrastructureTargetNewResponseEnvelopeSuccessTrue AccessInfrastructureTargetNewResponseEnvelopeSuccess = true
)

func (r AccessInfrastructureTargetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessInfrastructureTargetUpdateParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[AccessInfrastructureTargetUpdateParamsIP] `json:"ip,required"`
}

func (r AccessInfrastructureTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetUpdateParamsIP struct {
	// The target's IPv4 address
	IPV4 param.Field[AccessInfrastructureTargetUpdateParamsIPIPV4] `json:"ipv4"`
	// The target's IPv6 address
	IPV6 param.Field[AccessInfrastructureTargetUpdateParamsIPIPV6] `json:"ipv6"`
}

func (r AccessInfrastructureTargetUpdateParamsIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv4 address
type AccessInfrastructureTargetUpdateParamsIPIPV4 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetUpdateParamsIPIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv6 address
type AccessInfrastructureTargetUpdateParamsIPIPV6 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetUpdateParamsIPIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessInfrastructureTargetUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessInfrastructureTargetUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessInfrastructureTargetUpdateResponse                `json:"result"`
	JSON    accessInfrastructureTargetUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessInfrastructureTargetUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessInfrastructureTargetUpdateResponseEnvelope]
type accessInfrastructureTargetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessInfrastructureTargetUpdateResponseEnvelopeSuccess bool

const (
	AccessInfrastructureTargetUpdateResponseEnvelopeSuccessTrue AccessInfrastructureTargetUpdateResponseEnvelopeSuccess = true
)

func (r AccessInfrastructureTargetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessInfrastructureTargetListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Date and time at which the target was created
	CreatedAfter param.Field[time.Time] `query:"created_after" format:"date-time"`
	// Hostname of a target
	Hostname param.Field[string] `query:"hostname"`
	// Partial match to the hostname of a target
	HostnameContains param.Field[string] `query:"hostname_contains"`
	// IPv4 address of the target
	IPV4 param.Field[string] `query:"ip_v4"`
	// IPv6 address of the target
	IPV6 param.Field[string] `query:"ip_v6"`
	// Date and time at which the target was modified
	ModifiedAfter param.Field[time.Time] `query:"modified_after" format:"date-time"`
	// Current page in the response
	Page param.Field[int64] `query:"page"`
	// Max amount of entries returned per page
	PerPage param.Field[int64] `query:"per_page"`
	// Private virtual network identifier of the target
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [AccessInfrastructureTargetListParams]'s query parameters as
// `url.Values`.
func (r AccessInfrastructureTargetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessInfrastructureTargetDeleteParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessInfrastructureTargetBulkDeleteParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessInfrastructureTargetBulkUpdateParams struct {
	// Account identifier
	AccountID param.Field[string]                              `path:"account_id,required"`
	Body      []AccessInfrastructureTargetBulkUpdateParamsBody `json:"body,required"`
}

func (r AccessInfrastructureTargetBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessInfrastructureTargetBulkUpdateParamsBody struct {
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[AccessInfrastructureTargetBulkUpdateParamsBodyIP] `json:"ip,required"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The IPv4/IPv6 address that identifies where to reach a target
type AccessInfrastructureTargetBulkUpdateParamsBodyIP struct {
	// The target's IPv4 address
	IPV4 param.Field[AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV4] `json:"ipv4"`
	// The target's IPv6 address
	IPV6 param.Field[AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV6] `json:"ipv6"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBodyIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv4 address
type AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV4 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv6 address
type AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV6 struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccessInfrastructureTargetBulkUpdateParamsBodyIPIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessInfrastructureTargetGetParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessInfrastructureTargetGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessInfrastructureTargetGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessInfrastructureTargetGetResponse                `json:"result"`
	JSON    accessInfrastructureTargetGetResponseEnvelopeJSON    `json:"-"`
}

// accessInfrastructureTargetGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessInfrastructureTargetGetResponseEnvelope]
type accessInfrastructureTargetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessInfrastructureTargetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessInfrastructureTargetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessInfrastructureTargetGetResponseEnvelopeSuccess bool

const (
	AccessInfrastructureTargetGetResponseEnvelopeSuccessTrue AccessInfrastructureTargetGetResponseEnvelopeSuccess = true
)

func (r AccessInfrastructureTargetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessInfrastructureTargetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
