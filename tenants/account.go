// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AccountService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// List of accounts for the Tenant.
func (r *AccountService) List(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *pagination.SinglePage[TenantAccount], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/accounts", tenantID)
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

// List of accounts for the Tenant.
func (r *AccountService) ListAutoPaging(ctx context.Context, tenantID string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TenantAccount] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, tenantID, opts...))
}

type TenantAccount struct {
	ID        string                `json:"id" api:"required"`
	CreatedOn time.Time             `json:"created_on" api:"required" format:"date-time"`
	Name      string                `json:"name" api:"required,nullable"`
	Settings  TenantAccountSettings `json:"settings" api:"required"`
	Type      TenantAccountType     `json:"type" api:"required"`
	JSON      tenantAccountJSON     `json:"-"`
}

// tenantAccountJSON contains the JSON metadata for the struct [TenantAccount]
type tenantAccountJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Name        apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantAccountJSON) RawJSON() string {
	return r.raw
}

type TenantAccountSettings struct {
	AbuseContactEmail    string    `json:"abuse_contact_email" api:"required,nullable"`
	AccessApprovalExpiry time.Time `json:"access_approval_expiry" api:"required,nullable" format:"date-time"`
	APIAccessEnabled     bool      `json:"api_access_enabled" api:"required,nullable"`
	// Use
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-list-dns-settings)
	// instead. Deprecated.
	//
	// Deprecated: deprecated
	DefaultNameservers string `json:"default_nameservers" api:"required,nullable"`
	EnforceTwofactor   bool   `json:"enforce_twofactor" api:"required,nullable"`
	// Use
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-list-dns-settings)
	// instead. Deprecated.
	//
	// Deprecated: deprecated
	UseAccountCustomNSByDefault bool                      `json:"use_account_custom_ns_by_default" api:"required,nullable"`
	JSON                        tenantAccountSettingsJSON `json:"-"`
}

// tenantAccountSettingsJSON contains the JSON metadata for the struct
// [TenantAccountSettings]
type tenantAccountSettingsJSON struct {
	AbuseContactEmail           apijson.Field
	AccessApprovalExpiry        apijson.Field
	APIAccessEnabled            apijson.Field
	DefaultNameservers          apijson.Field
	EnforceTwofactor            apijson.Field
	UseAccountCustomNSByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TenantAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantAccountSettingsJSON) RawJSON() string {
	return r.raw
}

type TenantAccountType string

const (
	TenantAccountTypeStandard   TenantAccountType = "standard"
	TenantAccountTypeEnterprise TenantAccountType = "enterprise"
)

func (r TenantAccountType) IsKnown() bool {
	switch r {
	case TenantAccountTypeStandard, TenantAccountTypeEnterprise:
		return true
	}
	return false
}
