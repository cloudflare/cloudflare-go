// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/cloudflare/cloudflare-go/v7/tenants"
	"github.com/tidwall/gjson"
)

// OrganizationAccountService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationAccountService] method instead.
type OrganizationAccountService struct {
	Options []option.RequestOption
}

// NewOrganizationAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationAccountService(opts ...option.RequestOption) (r *OrganizationAccountService) {
	r = &OrganizationAccountService{}
	r.Options = opts
	return
}

// Retrieve the accounts immediately attached to a specific organization. Accounts
// attached to sub-organizations are not included. Use `page_token` and `page_size`
// for cursor pagination. The default response uses Account fields;
// `Accept: application/json; legacy=true` requests the LegacyAccount
// representation.
func (r *OrganizationAccountService) Get(ctx context.Context, organizationID string, params OrganizationAccountGetParams, opts ...option.RequestOption) (res *[]OrganizationAccountGetResponse, err error) {
	var env OrganizationAccountGetResponseEnvelope
	if params.Accept.Present {
		opts = append(opts, option.WithHeader("Accept", fmt.Sprintf("%v", params.Accept)))
	}
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/accounts", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type OrganizationAccountGetResponse struct {
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// This field can have the runtime type of [tenants.TenantAccountSettings],
	// [OrganizationAccountGetResponseOrganizationsAPILegacyAccountSettings].
	Settings       interface{}                               `json:"settings" api:"required"`
	ID             string                                    `json:"id"`
	AccountPubname string                                    `json:"account_pubname" api:"nullable"`
	AccountTag     string                                    `json:"account_tag"`
	AccountType    OrganizationAccountGetResponseAccountType `json:"account_type"`
	Name           string                                    `json:"name" api:"nullable"`
	// This field can have the runtime type of [map[string]string].
	Tags  interface{}                        `json:"tags"`
	Type  OrganizationAccountGetResponseType `json:"type"`
	JSON  organizationAccountGetResponseJSON `json:"-"`
	union OrganizationAccountGetResponseUnion
}

// organizationAccountGetResponseJSON contains the JSON metadata for the struct
// [OrganizationAccountGetResponse]
type organizationAccountGetResponseJSON struct {
	CreatedOn      apijson.Field
	Settings       apijson.Field
	ID             apijson.Field
	AccountPubname apijson.Field
	AccountTag     apijson.Field
	AccountType    apijson.Field
	Name           apijson.Field
	Tags           apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r organizationAccountGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *OrganizationAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = OrganizationAccountGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OrganizationAccountGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [tenants.TenantAccount],
// [OrganizationAccountGetResponseOrganizationsAPILegacyAccount].
func (r OrganizationAccountGetResponse) AsUnion() OrganizationAccountGetResponseUnion {
	return r.union
}

// Union satisfied by [tenants.TenantAccount] or
// [OrganizationAccountGetResponseOrganizationsAPILegacyAccount].
type OrganizationAccountGetResponseUnion interface {
	ImplementsOrganizationAccountGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OrganizationAccountGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(tenants.TenantAccount{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OrganizationAccountGetResponseOrganizationsAPILegacyAccount{}),
		},
	)
}

type OrganizationAccountGetResponseOrganizationsAPILegacyAccount struct {
	AccountPubname string                                                                 `json:"account_pubname" api:"required,nullable"`
	AccountTag     string                                                                 `json:"account_tag" api:"required"`
	AccountType    OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountType `json:"account_type" api:"required"`
	CreatedOn      time.Time                                                              `json:"created_on" api:"required" format:"date-time"`
	Settings       OrganizationAccountGetResponseOrganizationsAPILegacyAccountSettings    `json:"settings" api:"required"`
	JSON           organizationAccountGetResponseOrganizationsAPILegacyAccountJSON        `json:"-"`
}

// organizationAccountGetResponseOrganizationsAPILegacyAccountJSON contains the
// JSON metadata for the struct
// [OrganizationAccountGetResponseOrganizationsAPILegacyAccount]
type organizationAccountGetResponseOrganizationsAPILegacyAccountJSON struct {
	AccountPubname apijson.Field
	AccountTag     apijson.Field
	AccountType    apijson.Field
	CreatedOn      apijson.Field
	Settings       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrganizationAccountGetResponseOrganizationsAPILegacyAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAccountGetResponseOrganizationsAPILegacyAccountJSON) RawJSON() string {
	return r.raw
}

func (r OrganizationAccountGetResponseOrganizationsAPILegacyAccount) ImplementsOrganizationAccountGetResponse() {
}

type OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountType string

const (
	OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountTypeStandard   OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountType = "standard"
	OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountTypeEnterprise OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountType = "enterprise"
)

func (r OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountType) IsKnown() bool {
	switch r {
	case OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountTypeStandard, OrganizationAccountGetResponseOrganizationsAPILegacyAccountAccountTypeEnterprise:
		return true
	}
	return false
}

type OrganizationAccountGetResponseOrganizationsAPILegacyAccountSettings struct {
	AccessApprovalExpiry        time.Time                                                               `json:"access_approval_expiry" api:"required,nullable" format:"date-time"`
	APIAccessEnabled            bool                                                                    `json:"api_access_enabled" api:"required"`
	EnforceTwoFactor            bool                                                                    `json:"enforce_two_factor" api:"required"`
	UseAccountCustomNSByDefault bool                                                                    `json:"use_account_custom_ns_by_default" api:"required"`
	JSON                        organizationAccountGetResponseOrganizationsAPILegacyAccountSettingsJSON `json:"-"`
}

// organizationAccountGetResponseOrganizationsAPILegacyAccountSettingsJSON contains
// the JSON metadata for the struct
// [OrganizationAccountGetResponseOrganizationsAPILegacyAccountSettings]
type organizationAccountGetResponseOrganizationsAPILegacyAccountSettingsJSON struct {
	AccessApprovalExpiry        apijson.Field
	APIAccessEnabled            apijson.Field
	EnforceTwoFactor            apijson.Field
	UseAccountCustomNSByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrganizationAccountGetResponseOrganizationsAPILegacyAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAccountGetResponseOrganizationsAPILegacyAccountSettingsJSON) RawJSON() string {
	return r.raw
}

type OrganizationAccountGetResponseAccountType string

const (
	OrganizationAccountGetResponseAccountTypeStandard   OrganizationAccountGetResponseAccountType = "standard"
	OrganizationAccountGetResponseAccountTypeEnterprise OrganizationAccountGetResponseAccountType = "enterprise"
)

func (r OrganizationAccountGetResponseAccountType) IsKnown() bool {
	switch r {
	case OrganizationAccountGetResponseAccountTypeStandard, OrganizationAccountGetResponseAccountTypeEnterprise:
		return true
	}
	return false
}

type OrganizationAccountGetResponseType string

const (
	OrganizationAccountGetResponseTypeStandard   OrganizationAccountGetResponseType = "standard"
	OrganizationAccountGetResponseTypeEnterprise OrganizationAccountGetResponseType = "enterprise"
)

func (r OrganizationAccountGetResponseType) IsKnown() bool {
	switch r {
	case OrganizationAccountGetResponseTypeStandard, OrganizationAccountGetResponseTypeEnterprise:
		return true
	}
	return false
}

type OrganizationAccountGetParams struct {
	AccountPubname param.Field[OrganizationAccountGetParamsAccountPubname] `query:"account_pubname"`
	// Sort direction for the order_by field. Valid values: `asc`, `desc`. Defaults to
	// `asc` when order_by is specified.
	Direction param.Field[OrganizationAccountGetParamsDirection] `query:"direction"`
	// Include Account tags from the resource tag mirror. Omit this parameter to
	// preserve the existing Account response shape.
	IncludeTags param.Field[bool] `query:"include_tags"`
	// Whether to calculate and return the exact result_info.total_size for cursor
	// pagination. Defaults to true. When false, total_size is omitted. page_size and
	// include_total may change between pages; next_page_token remains the
	// authoritative continuation signal.
	IncludeTotal param.Field[bool]                             `query:"include_total"`
	Name         param.Field[OrganizationAccountGetParamsName] `query:"name"`
	// Field to order results by. Currently supported values: `account_name`.
	OrderBy param.Field[OrganizationAccountGetParamsOrderBy] `query:"order_by"`
	// The amount of items to return. Defaults to 10.
	PageSize param.Field[int64] `query:"page_size"`
	// An opaque token returned from the last list response that when provided will
	// retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent requests
	// with a page token.
	PageToken param.Field[string] `query:"page_token"`
	// Filter accounts by lifecycle status. Defaults to `active`.
	Status param.Field[OrganizationAccountGetParamsStatus] `query:"status"`
	Accept param.Field[string]                             `header:"Accept"`
}

// URLQuery serializes [OrganizationAccountGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationAccountGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OrganizationAccountGetParamsAccountPubname struct {
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

// URLQuery serializes [OrganizationAccountGetParamsAccountPubname]'s query
// parameters as `url.Values`.
func (r OrganizationAccountGetParamsAccountPubname) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction for the order_by field. Valid values: `asc`, `desc`. Defaults to
// `asc` when order_by is specified.
type OrganizationAccountGetParamsDirection string

const (
	OrganizationAccountGetParamsDirectionAsc  OrganizationAccountGetParamsDirection = "asc"
	OrganizationAccountGetParamsDirectionDesc OrganizationAccountGetParamsDirection = "desc"
)

func (r OrganizationAccountGetParamsDirection) IsKnown() bool {
	switch r {
	case OrganizationAccountGetParamsDirectionAsc, OrganizationAccountGetParamsDirectionDesc:
		return true
	}
	return false
}

type OrganizationAccountGetParamsName struct {
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

// URLQuery serializes [OrganizationAccountGetParamsName]'s query parameters as
// `url.Values`.
func (r OrganizationAccountGetParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order results by. Currently supported values: `account_name`.
type OrganizationAccountGetParamsOrderBy string

const (
	OrganizationAccountGetParamsOrderByAccountName OrganizationAccountGetParamsOrderBy = "account_name"
)

func (r OrganizationAccountGetParamsOrderBy) IsKnown() bool {
	switch r {
	case OrganizationAccountGetParamsOrderByAccountName:
		return true
	}
	return false
}

// Filter accounts by lifecycle status. Defaults to `active`.
type OrganizationAccountGetParamsStatus string

const (
	OrganizationAccountGetParamsStatusActive  OrganizationAccountGetParamsStatus = "active"
	OrganizationAccountGetParamsStatusDeleted OrganizationAccountGetParamsStatus = "deleted"
)

func (r OrganizationAccountGetParamsStatus) IsKnown() bool {
	switch r {
	case OrganizationAccountGetParamsStatusActive, OrganizationAccountGetParamsStatusDeleted:
		return true
	}
	return false
}

type OrganizationAccountGetResponseEnvelope struct {
	Errors     []interface{}                                    `json:"errors" api:"required"`
	Messages   []shared.ResponseInfo                            `json:"messages" api:"required"`
	Result     []OrganizationAccountGetResponse                 `json:"result" api:"required"`
	ResultInfo OrganizationAccountGetResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	Success    OrganizationAccountGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON       organizationAccountGetResponseEnvelopeJSON       `json:"-"`
}

// organizationAccountGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OrganizationAccountGetResponseEnvelope]
type organizationAccountGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationAccountGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAccountGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationAccountGetResponseEnvelopeResultInfo struct {
	// Use this opaque token in the next request to retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent requests
	// with a page token.
	NextPageToken string `json:"next_page_token"`
	// Counts the total amount of items in a list with the applied filters. The API
	// omits next_page_token to indicate no more items in a particular list.
	TotalSize int64                                                `json:"total_size"`
	JSON      organizationAccountGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// organizationAccountGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [OrganizationAccountGetResponseEnvelopeResultInfo]
type organizationAccountGetResponseEnvelopeResultInfoJSON struct {
	NextPageToken apijson.Field
	TotalSize     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrganizationAccountGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAccountGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type OrganizationAccountGetResponseEnvelopeSuccess bool

const (
	OrganizationAccountGetResponseEnvelopeSuccessTrue OrganizationAccountGetResponseEnvelopeSuccess = true
)

func (r OrganizationAccountGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationAccountGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
