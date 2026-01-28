// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
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

// List all accounts you have ownership or verified access to.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccountListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
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

// List all accounts you have ownership or verified access to.
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccountListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

type AccountListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Account name
	Name string                  `json:"name,required"`
	Type AccountListResponseType `json:"type,required"`
	// Timestamp for the creation of the account
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Parent container details
	ManagedBy AccountListResponseManagedBy `json:"managed_by"`
	// Account settings
	Settings AccountListResponseSettings `json:"settings"`
	JSON     accountListResponseJSON     `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	CreatedOn   apijson.Field
	ManagedBy   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountListResponseType string

const (
	AccountListResponseTypeStandard   AccountListResponseType = "standard"
	AccountListResponseTypeEnterprise AccountListResponseType = "enterprise"
)

func (r AccountListResponseType) IsKnown() bool {
	switch r {
	case AccountListResponseTypeStandard, AccountListResponseTypeEnterprise:
		return true
	}
	return false
}

// Parent container details
type AccountListResponseManagedBy struct {
	// ID of the parent Organization, if one exists
	ParentOrgID string `json:"parent_org_id"`
	// Name of the parent Organization, if one exists
	ParentOrgName string                           `json:"parent_org_name"`
	JSON          accountListResponseManagedByJSON `json:"-"`
}

// accountListResponseManagedByJSON contains the JSON metadata for the struct
// [AccountListResponseManagedBy]
type accountListResponseManagedByJSON struct {
	ParentOrgID   apijson.Field
	ParentOrgName apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountListResponseManagedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseManagedByJSON) RawJSON() string {
	return r.raw
}

// Account settings
type AccountListResponseSettings struct {
	// Sets an abuse contact email to notify for abuse reports.
	AbuseContactEmail string `json:"abuse_contact_email"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool                            `json:"enforce_twofactor"`
	JSON             accountListResponseSettingsJSON `json:"-"`
}

// accountListResponseSettingsJSON contains the JSON metadata for the struct
// [AccountListResponseSettings]
type accountListResponseSettingsJSON struct {
	AbuseContactEmail apijson.Field
	EnforceTwofactor  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountListResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountListParams struct {
	// Direction to order results.
	Direction param.Field[AccountListParamsDirection] `query:"direction"`
	// Name of the account.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
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
