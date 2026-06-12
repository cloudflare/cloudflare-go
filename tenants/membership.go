// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// MembershipService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMembershipService] method instead.
type MembershipService struct {
	Options []option.RequestOption
}

// NewMembershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMembershipService(opts ...option.RequestOption) (r *MembershipService) {
	r = &MembershipService{}
	r.Options = opts
	return
}

// List of active members (Cloudflare users) for the Tenant.
func (r *MembershipService) List(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *pagination.SinglePage[TenantMembership], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/memberships", tenantID)
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

// List of active members (Cloudflare users) for the Tenant.
func (r *MembershipService) ListAutoPaging(ctx context.Context, tenantID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TenantMembership] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, tenantID, opts...))
}

type TenantMembership struct {
	UserEmail string               `json:"user_email" api:"required"`
	UserName  string               `json:"user_name" api:"required"`
	UserTag   string               `json:"user_tag" api:"required"`
	JSON      tenantMembershipJSON `json:"-"`
}

// tenantMembershipJSON contains the JSON metadata for the struct
// [TenantMembership]
type tenantMembershipJSON struct {
	UserEmail   apijson.Field
	UserName    apijson.Field
	UserTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantMembershipJSON) RawJSON() string {
	return r.raw
}
