// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TokenPermissionGroupService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTokenPermissionGroupService]
// method instead.
type TokenPermissionGroupService struct {
	Options []option.RequestOption
}

// NewTokenPermissionGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTokenPermissionGroupService(opts ...option.RequestOption) (r *TokenPermissionGroupService) {
	r = &TokenPermissionGroupService{}
	r.Options = opts
	return
}

// Find all available permission groups.
func (r *TokenPermissionGroupService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[TokenPermissionGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/tokens/permission_groups"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Find all available permission groups.
func (r *TokenPermissionGroupService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TokenPermissionGroupListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

type Permission struct {
	Analytics    accounts.PermissionGrant `json:"analytics"`
	Billing      accounts.PermissionGrant `json:"billing"`
	CachePurge   accounts.PermissionGrant `json:"cache_purge"`
	DNS          accounts.PermissionGrant `json:"dns"`
	DNSRecords   accounts.PermissionGrant `json:"dns_records"`
	Lb           accounts.PermissionGrant `json:"lb"`
	Logs         accounts.PermissionGrant `json:"logs"`
	Organization accounts.PermissionGrant `json:"organization"`
	SSL          accounts.PermissionGrant `json:"ssl"`
	WAF          accounts.PermissionGrant `json:"waf"`
	ZoneSettings accounts.PermissionGrant `json:"zone_settings"`
	Zones        accounts.PermissionGrant `json:"zones"`
	JSON         permissionJSON           `json:"-"`
}

// permissionJSON contains the JSON metadata for the struct [Permission]
type permissionJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	Lb           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	SSL          apijson.Field
	WAF          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Permission) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionJSON) RawJSON() string {
	return r.raw
}

type PermissionParam struct {
	Analytics    param.Field[accounts.PermissionGrantParam] `json:"analytics"`
	Billing      param.Field[accounts.PermissionGrantParam] `json:"billing"`
	CachePurge   param.Field[accounts.PermissionGrantParam] `json:"cache_purge"`
	DNS          param.Field[accounts.PermissionGrantParam] `json:"dns"`
	DNSRecords   param.Field[accounts.PermissionGrantParam] `json:"dns_records"`
	Lb           param.Field[accounts.PermissionGrantParam] `json:"lb"`
	Logs         param.Field[accounts.PermissionGrantParam] `json:"logs"`
	Organization param.Field[accounts.PermissionGrantParam] `json:"organization"`
	SSL          param.Field[accounts.PermissionGrantParam] `json:"ssl"`
	WAF          param.Field[accounts.PermissionGrantParam] `json:"waf"`
	ZoneSettings param.Field[accounts.PermissionGrantParam] `json:"zone_settings"`
	Zones        param.Field[accounts.PermissionGrantParam] `json:"zones"`
}

func (r PermissionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PermissionItem = string

type TokenPermissionGroupListResponse = interface{}
