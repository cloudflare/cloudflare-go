// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
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

// Retrieve a list of accounts that belong to a specific organization. (Currently
// in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
func (r *OrganizationAccountService) Get(ctx context.Context, organizationID string, query OrganizationAccountGetParams, opts ...option.RequestOption) (res *[]OrganizationAccountGetResponse, err error) {
	var env OrganizationAccountGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/accounts", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OrganizationAccountGetResponse struct {
	ID        string                                 `json:"id,required"`
	CreatedOn time.Time                              `json:"created_on,required" format:"date-time"`
	Name      string                                 `json:"name,required,nullable"`
	Settings  OrganizationAccountGetResponseSettings `json:"settings,required"`
	Type      OrganizationAccountGetResponseType     `json:"type,required"`
	JSON      organizationAccountGetResponseJSON     `json:"-"`
}

// organizationAccountGetResponseJSON contains the JSON metadata for the struct
// [OrganizationAccountGetResponse]
type organizationAccountGetResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Name        apijson.Field
	Settings    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAccountGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationAccountGetResponseSettings struct {
	AbuseContactEmail    string    `json:"abuse_contact_email,required,nullable"`
	AccessApprovalExpiry time.Time `json:"access_approval_expiry,required,nullable" format:"date-time"`
	APIAccessEnabled     bool      `json:"api_access_enabled,required,nullable"`
	// Use
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-list-dns-settings)
	// instead. Deprecated.
	//
	// Deprecated: deprecated
	DefaultNameservers string `json:"default_nameservers,required,nullable"`
	EnforceTwofactor   bool   `json:"enforce_twofactor,required,nullable"`
	// Use
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-list-dns-settings)
	// instead. Deprecated.
	//
	// Deprecated: deprecated
	UseAccountCustomNSByDefault bool                                       `json:"use_account_custom_ns_by_default,required,nullable"`
	JSON                        organizationAccountGetResponseSettingsJSON `json:"-"`
}

// organizationAccountGetResponseSettingsJSON contains the JSON metadata for the
// struct [OrganizationAccountGetResponseSettings]
type organizationAccountGetResponseSettingsJSON struct {
	AbuseContactEmail           apijson.Field
	AccessApprovalExpiry        apijson.Field
	APIAccessEnabled            apijson.Field
	DefaultNameservers          apijson.Field
	EnforceTwofactor            apijson.Field
	UseAccountCustomNSByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrganizationAccountGetResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAccountGetResponseSettingsJSON) RawJSON() string {
	return r.raw
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
	Name           param.Field[OrganizationAccountGetParamsName]           `query:"name"`
	// The amount of items to return. Defaults to 10.
	PageSize param.Field[int64] `query:"page_size"`
	// An opaque token returned from the last list response that when provided will
	// retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent requests
	// with a page token.
	PageToken param.Field[string] `query:"page_token"`
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

type OrganizationAccountGetResponseEnvelope struct {
	Errors     []interface{}                                    `json:"errors,required"`
	Messages   []shared.ResponseInfo                            `json:"messages,required"`
	Result     []OrganizationAccountGetResponse                 `json:"result,required"`
	ResultInfo OrganizationAccountGetResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    OrganizationAccountGetResponseEnvelopeSuccess    `json:"success,required"`
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
