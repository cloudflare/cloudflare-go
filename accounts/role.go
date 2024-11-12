// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// RoleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoleService] method instead.
type RoleService struct {
	Options []option.RequestOption
}

// NewRoleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoleService(opts ...option.RequestOption) (r *RoleService) {
	r = &RoleService{}
	r.Options = opts
	return
}

// Get all available roles for an account.
func (r *RoleService) List(ctx context.Context, query RoleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RoleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/roles", query.AccountID)
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

// Get all available roles for an account.
func (r *RoleService) ListAutoPaging(ctx context.Context, query RoleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RoleListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Get information about a specific role for an account.
func (r *RoleService) Get(ctx context.Context, roleID string, query RoleGetParams, opts ...option.RequestOption) (res *RoleGetResponse, err error) {
	var env RoleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if roleID == "" {
		err = errors.New("missing required role_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/roles/%s", query.AccountID, roleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoleListResponse struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                      `json:"name,required"`
	Permissions RoleListResponsePermissions `json:"permissions,required"`
	JSON        roleListResponseJSON        `json:"-"`
}

// roleListResponseJSON contains the JSON metadata for the struct
// [RoleListResponse]
type roleListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleListResponseJSON) RawJSON() string {
	return r.raw
}

type RoleListResponsePermissions struct {
	Analytics    shared.PermissionGrant          `json:"analytics"`
	Billing      shared.PermissionGrant          `json:"billing"`
	CachePurge   shared.PermissionGrant          `json:"cache_purge"`
	DNS          shared.PermissionGrant          `json:"dns"`
	DNSRecords   shared.PermissionGrant          `json:"dns_records"`
	LB           shared.PermissionGrant          `json:"lb"`
	Logs         shared.PermissionGrant          `json:"logs"`
	Organization shared.PermissionGrant          `json:"organization"`
	SSL          shared.PermissionGrant          `json:"ssl"`
	WAF          shared.PermissionGrant          `json:"waf"`
	ZoneSettings shared.PermissionGrant          `json:"zone_settings"`
	Zones        shared.PermissionGrant          `json:"zones"`
	JSON         roleListResponsePermissionsJSON `json:"-"`
}

// roleListResponsePermissionsJSON contains the JSON metadata for the struct
// [RoleListResponsePermissions]
type roleListResponsePermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	LB           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	SSL          apijson.Field
	WAF          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RoleListResponsePermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleListResponsePermissionsJSON) RawJSON() string {
	return r.raw
}

type RoleGetResponse struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                     `json:"name,required"`
	Permissions RoleGetResponsePermissions `json:"permissions,required"`
	JSON        roleGetResponseJSON        `json:"-"`
}

// roleGetResponseJSON contains the JSON metadata for the struct [RoleGetResponse]
type roleGetResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleGetResponseJSON) RawJSON() string {
	return r.raw
}

type RoleGetResponsePermissions struct {
	Analytics    shared.PermissionGrant         `json:"analytics"`
	Billing      shared.PermissionGrant         `json:"billing"`
	CachePurge   shared.PermissionGrant         `json:"cache_purge"`
	DNS          shared.PermissionGrant         `json:"dns"`
	DNSRecords   shared.PermissionGrant         `json:"dns_records"`
	LB           shared.PermissionGrant         `json:"lb"`
	Logs         shared.PermissionGrant         `json:"logs"`
	Organization shared.PermissionGrant         `json:"organization"`
	SSL          shared.PermissionGrant         `json:"ssl"`
	WAF          shared.PermissionGrant         `json:"waf"`
	ZoneSettings shared.PermissionGrant         `json:"zone_settings"`
	Zones        shared.PermissionGrant         `json:"zones"`
	JSON         roleGetResponsePermissionsJSON `json:"-"`
}

// roleGetResponsePermissionsJSON contains the JSON metadata for the struct
// [RoleGetResponsePermissions]
type roleGetResponsePermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	LB           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	SSL          apijson.Field
	WAF          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RoleGetResponsePermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleGetResponsePermissionsJSON) RawJSON() string {
	return r.raw
}

type RoleListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RoleGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type RoleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RoleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RoleGetResponse                `json:"result"`
	JSON    roleGetResponseEnvelopeJSON    `json:"-"`
}

// roleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoleGetResponseEnvelope]
type roleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoleGetResponseEnvelopeSuccess bool

const (
	RoleGetResponseEnvelopeSuccessTrue RoleGetResponseEnvelopeSuccess = true
)

func (r RoleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RoleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
