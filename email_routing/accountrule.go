// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AccountRuleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRuleService] method instead.
type AccountRuleService struct {
	Options []option.RequestOption
}

// NewAccountRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRuleService(opts ...option.RequestOption) (r *AccountRuleService) {
	r = &AccountRuleService{}
	r.Options = opts
	return
}

// Lists existing routing rules across all zones in the account or zone.
func (r *AccountRuleService) List(ctx context.Context, params AccountRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccountRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/email/routing/rules", accountOrZone, accountOrZoneID)
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

// Lists existing routing rules across all zones in the account or zone.
func (r *AccountRuleService) ListAutoPaging(ctx context.Context, params AccountRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccountRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type AccountRule struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []Action `json:"actions"`
	// Routing rule status.
	Enabled AccountRuleEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []Matcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Who manages the rule. `api` covers dashboard, generic API, and Terraform;
	// `wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults to
	// `api` when omitted on write.
	Source AccountRuleSource `json:"source"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	//
	// Deprecated: deprecated
	Tag string `json:"tag"`
	// Zone information for the routing rule.
	Zone AccountRuleZone `json:"zone"`
	JSON accountRuleJSON `json:"-"`
}

// accountRuleJSON contains the JSON metadata for the struct [AccountRule]
type accountRuleJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Source      apijson.Field
	Tag         apijson.Field
	Zone        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleJSON) RawJSON() string {
	return r.raw
}

// Routing rule status.
type AccountRuleEnabled bool

const (
	AccountRuleEnabledTrue  AccountRuleEnabled = true
	AccountRuleEnabledFalse AccountRuleEnabled = false
)

func (r AccountRuleEnabled) IsKnown() bool {
	switch r {
	case AccountRuleEnabledTrue, AccountRuleEnabledFalse:
		return true
	}
	return false
}

// Who manages the rule. `api` covers dashboard, generic API, and Terraform;
// `wrangler` means the rule is managed by a Worker's wrangler.jsonc. Defaults to
// `api` when omitted on write.
type AccountRuleSource string

const (
	AccountRuleSourceAPI      AccountRuleSource = "api"
	AccountRuleSourceWrangler AccountRuleSource = "wrangler"
)

func (r AccountRuleSource) IsKnown() bool {
	switch r {
	case AccountRuleSourceAPI, AccountRuleSourceWrangler:
		return true
	}
	return false
}

// Zone information for the routing rule.
type AccountRuleZone struct {
	// Zone name.
	Name string `json:"name"`
	// Zone tag.
	Tag  string              `json:"tag"`
	JSON accountRuleZoneJSON `json:"-"`
}

// accountRuleZoneJSON contains the JSON metadata for the struct [AccountRuleZone]
type accountRuleZoneJSON struct {
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleZoneJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Filter by enabled routing rules.
	Enabled param.Field[AccountRuleListParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountRuleListParams]'s query parameters as `url.Values`.
func (r AccountRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by enabled routing rules.
type AccountRuleListParamsEnabled bool

const (
	AccountRuleListParamsEnabledTrue  AccountRuleListParamsEnabled = true
	AccountRuleListParamsEnabledFalse AccountRuleListParamsEnabled = false
)

func (r AccountRuleListParamsEnabled) IsKnown() bool {
	switch r {
	case AccountRuleListParamsEnabledTrue, AccountRuleListParamsEnabledFalse:
		return true
	}
	return false
}
