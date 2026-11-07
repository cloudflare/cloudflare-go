// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
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

// Lists all Cloudflare accounts across this tenant's hierarchy using the legacy
// Account response shape. Use `GET /organizations/{organization_id}/accounts` to
// list only accounts immediately attached to a specific Organization using the
// modern Account response shape. Use `page_token` and `page_size` for cursor
// pagination.
func (r *AccountService) List(ctx context.Context, tenantID string, query AccountListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccountListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/accounts", tenantID)
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

// Lists all Cloudflare accounts across this tenant's hierarchy using the legacy
// Account response shape. Use `GET /organizations/{organization_id}/accounts` to
// list only accounts immediately attached to a specific Organization using the
// modern Account response shape. Use `page_token` and `page_size` for cursor
// pagination.
func (r *AccountService) ListAutoPaging(ctx context.Context, tenantID string, query AccountListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccountListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, tenantID, query, opts...))
}

type TenantAccount struct {
	ID        string                `json:"id" api:"required"`
	CreatedOn time.Time             `json:"created_on" api:"required" format:"date-time"`
	Name      string                `json:"name" api:"required,nullable"`
	Settings  TenantAccountSettings `json:"settings" api:"required"`
	Type      TenantAccountType     `json:"type" api:"required"`
	// Account tags, present only when `include_tags=true` is requested.
	Tags map[string]string `json:"tags"`
	JSON tenantAccountJSON `json:"-"`
}

// tenantAccountJSON contains the JSON metadata for the struct [TenantAccount]
type tenantAccountJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Name        apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantAccountJSON) RawJSON() string {
	return r.raw
}

func (r TenantAccount) ImplementsOrganizationAccountGetResponse() {}

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

type AccountListResponse struct {
	AccountPubname string                         `json:"account_pubname" api:"required,nullable"`
	AccountTag     string                         `json:"account_tag" api:"required"`
	AccountType    AccountListResponseAccountType `json:"account_type" api:"required"`
	CreatedOn      time.Time                      `json:"created_on" api:"required" format:"date-time"`
	Settings       AccountListResponseSettings    `json:"settings" api:"required"`
	JSON           accountListResponseJSON        `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	AccountPubname apijson.Field
	AccountTag     apijson.Field
	AccountType    apijson.Field
	CreatedOn      apijson.Field
	Settings       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListResponseAccountType string

const (
	AccountListResponseAccountTypeStandard   AccountListResponseAccountType = "standard"
	AccountListResponseAccountTypeEnterprise AccountListResponseAccountType = "enterprise"
)

func (r AccountListResponseAccountType) IsKnown() bool {
	switch r {
	case AccountListResponseAccountTypeStandard, AccountListResponseAccountTypeEnterprise:
		return true
	}
	return false
}

type AccountListResponseSettings struct {
	AccessApprovalExpiry        time.Time                       `json:"access_approval_expiry" api:"required,nullable" format:"date-time"`
	APIAccessEnabled            bool                            `json:"api_access_enabled" api:"required"`
	EnforceTwoFactor            bool                            `json:"enforce_two_factor" api:"required"`
	UseAccountCustomNSByDefault bool                            `json:"use_account_custom_ns_by_default" api:"required"`
	JSON                        accountListResponseSettingsJSON `json:"-"`
}

// accountListResponseSettingsJSON contains the JSON metadata for the struct
// [AccountListResponseSettings]
type accountListResponseSettingsJSON struct {
	AccessApprovalExpiry        apijson.Field
	APIAccessEnabled            apijson.Field
	EnforceTwoFactor            apijson.Field
	UseAccountCustomNSByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccountListResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountListParams struct {
	AccountPubname param.Field[AccountListParamsAccountPubname] `query:"account_pubname"`
	// Sort direction for the order_by field. Valid values: `asc`, `desc`. Defaults to
	// `asc` when order_by is specified.
	Direction param.Field[AccountListParamsDirection] `query:"direction"`
	// Whether to calculate and return the exact `result_info.total_size` for cursor
	// pagination. Defaults to true. When false, `total_size` is omitted.
	IncludeTotal param.Field[bool]                  `query:"include_total"`
	Name         param.Field[AccountListParamsName] `query:"name"`
	// Field to order results by. Currently supported values: `account_name`.
	OrderBy param.Field[AccountListParamsOrderBy] `query:"order_by"`
	// The amount of items to return. Defaults to 10.
	PageSize param.Field[int64] `query:"page_size"`
	// An opaque token returned from the last list response that when provided will
	// retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent requests
	// with a page token.
	PageToken param.Field[string] `query:"page_token"`
	// Filter accounts by lifecycle status. Defaults to `active`.
	Status param.Field[AccountListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccountListParamsAccountPubname struct {
	// (case-insensitive) Filter the list of accounts to where the account_pubname
	// contains a particular string.
	Contains param.Field[string] `query:"contains"`
	// (case-insensitive) Filter the list of accounts to where the account_pubname ends
	// with a particular string.
	EndsWith param.Field[string] `query:"endsWith"`
	// (case-insensitive) Filter the list of accounts to where the account_pubname
	// starts with a particular string.
	StartsWith param.Field[string] `query:"startsWith"`
}

// URLQuery serializes [AccountListParamsAccountPubname]'s query parameters as
// `url.Values`.
func (r AccountListParamsAccountPubname) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction for the order_by field. Valid values: `asc`, `desc`. Defaults to
// `asc` when order_by is specified.
type AccountListParamsDirection string

const (
	AccountListParamsDirectionAsc  AccountListParamsDirection = "asc"
	AccountListParamsDirectionDesc AccountListParamsDirection = "desc"
)

func (r AccountListParamsDirection) IsKnown() bool {
	switch r {
	case AccountListParamsDirectionAsc, AccountListParamsDirectionDesc:
		return true
	}
	return false
}

type AccountListParamsName struct {
	// (case-insensitive) Filter the list of accounts to where the name contains a
	// particular string.
	Contains param.Field[string] `query:"contains"`
	// (case-insensitive) Filter the list of accounts to where the name ends with a
	// particular string.
	EndsWith param.Field[string] `query:"endsWith"`
	// (case-insensitive) Filter the list of accounts to where the name starts with a
	// particular string.
	StartsWith param.Field[string] `query:"startsWith"`
}

// URLQuery serializes [AccountListParamsName]'s query parameters as `url.Values`.
func (r AccountListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order results by. Currently supported values: `account_name`.
type AccountListParamsOrderBy string

const (
	AccountListParamsOrderByAccountName AccountListParamsOrderBy = "account_name"
)

func (r AccountListParamsOrderBy) IsKnown() bool {
	switch r {
	case AccountListParamsOrderByAccountName:
		return true
	}
	return false
}

// Filter accounts by lifecycle status. Defaults to `active`.
type AccountListParamsStatus string

const (
	AccountListParamsStatusActive  AccountListParamsStatus = "active"
	AccountListParamsStatusDeleted AccountListParamsStatus = "deleted"
)

func (r AccountListParamsStatus) IsKnown() bool {
	switch r {
	case AccountListParamsStatusActive, AccountListParamsStatusDeleted:
		return true
	}
	return false
}
