// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

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
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// WAFPackageService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFPackageService] method instead.
type WAFPackageService struct {
	Options []option.RequestOption
	Groups  *WAFPackageGroupService
	Rules   *WAFPackageRuleService
}

// NewWAFPackageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFPackageService(opts ...option.RequestOption) (r *WAFPackageService) {
	r = &WAFPackageService{}
	r.Options = opts
	r.Groups = NewWAFPackageGroupService(opts...)
	r.Rules = NewWAFPackageRuleService(opts...)
	return
}

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageService) List(ctx context.Context, params WAFPackageListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WAFPackageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages", params.ZoneID)
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

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageService) ListAutoPaging(ctx context.Context, params WAFPackageListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WAFPackageListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Fetches the details of a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageService) Get(ctx context.Context, packageID string, query WAFPackageGetParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env WAFPackageGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", query.ZoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type WAFPackageListResponse = interface{}

type WAFPackageListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The direction used to sort returned packages.
	Direction param.Field[WAFPackageListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[WAFPackageListParamsMatch] `query:"match"`
	// The name of the WAF package.
	Name param.Field[string] `query:"name"`
	// The field used to sort returned packages.
	Order param.Field[WAFPackageListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of packages per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WAFPackageListParams]'s query parameters as `url.Values`.
func (r WAFPackageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned packages.
type WAFPackageListParamsDirection string

const (
	WAFPackageListParamsDirectionAsc  WAFPackageListParamsDirection = "asc"
	WAFPackageListParamsDirectionDesc WAFPackageListParamsDirection = "desc"
)

func (r WAFPackageListParamsDirection) IsKnown() bool {
	switch r {
	case WAFPackageListParamsDirectionAsc, WAFPackageListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type WAFPackageListParamsMatch string

const (
	WAFPackageListParamsMatchAny WAFPackageListParamsMatch = "any"
	WAFPackageListParamsMatchAll WAFPackageListParamsMatch = "all"
)

func (r WAFPackageListParamsMatch) IsKnown() bool {
	switch r {
	case WAFPackageListParamsMatchAny, WAFPackageListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned packages.
type WAFPackageListParamsOrder string

const (
	WAFPackageListParamsOrderName WAFPackageListParamsOrder = "name"
)

func (r WAFPackageListParamsOrder) IsKnown() bool {
	switch r {
	case WAFPackageListParamsOrderName:
		return true
	}
	return false
}

type WAFPackageGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type WAFPackageGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Result   interface{}           `json:"result" api:"required"`
	// Defines whether the API call was successful.
	Success WAFPackageGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    wafPackageGetResponseEnvelopeJSON    `json:"-"`
}

// wafPackageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFPackageGetResponseEnvelope]
type wafPackageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFPackageGetResponseEnvelopeSuccess bool

const (
	WAFPackageGetResponseEnvelopeSuccessTrue WAFPackageGetResponseEnvelopeSuccess = true
)

func (r WAFPackageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
