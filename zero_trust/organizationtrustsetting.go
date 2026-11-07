// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// OrganizationTrustSettingService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationTrustSettingService] method instead.
type OrganizationTrustSettingService struct {
	Options []option.RequestOption
}

// NewOrganizationTrustSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOrganizationTrustSettingService(opts ...option.RequestOption) (r *OrganizationTrustSettingService) {
	r = &OrganizationTrustSettingService{}
	r.Options = opts
	return
}

// Updates the trust settings for your Zero Trust organization.
func (r *OrganizationTrustSettingService) Update(ctx context.Context, params OrganizationTrustSettingUpdateParams, opts ...option.RequestOption) (res *OrganizationTrustSettingUpdateResponse, err error) {
	var env OrganizationTrustSettingUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/trust_settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Resets the trust settings for your Zero Trust organization.
func (r *OrganizationTrustSettingService) Delete(ctx context.Context, body OrganizationTrustSettingDeleteParams, opts ...option.RequestOption) (res *OrganizationTrustSettingDeleteResponse, err error) {
	var env OrganizationTrustSettingDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/trust_settings", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the trust settings for your Zero Trust organization.
func (r *OrganizationTrustSettingService) Get(ctx context.Context, query OrganizationTrustSettingGetParams, opts ...option.RequestOption) (res *OrganizationTrustSettingGetResponse, err error) {
	var env OrganizationTrustSettingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/trust_settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the trust settings for your Zero Trust organization.
func (r *OrganizationTrustSettingService) Replace(ctx context.Context, params OrganizationTrustSettingReplaceParams, opts ...option.RequestOption) (res *OrganizationTrustSettingReplaceResponse, err error) {
	var env OrganizationTrustSettingReplaceResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/organizations/trust_settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type OrganizationTrustSettingUpdateResponse struct {
	// Whether all accounts in the organization are trusted to use this account's
	// cross-account Access resources.
	GrantTrustToAllAccountsInOrg bool `json:"grant_trust_to_all_accounts_in_org" api:"required"`
	// Explicit account IDs trusted to use this account's cross-account Access
	// resources. At most 100 unique valid account IDs are accepted. Omitted from
	// responses when empty.
	TrustingAccounts []string                                   `json:"trusting_accounts"`
	JSON             organizationTrustSettingUpdateResponseJSON `json:"-"`
}

// organizationTrustSettingUpdateResponseJSON contains the JSON metadata for the
// struct [OrganizationTrustSettingUpdateResponse]
type organizationTrustSettingUpdateResponseJSON struct {
	GrantTrustToAllAccountsInOrg apijson.Field
	TrustingAccounts             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *OrganizationTrustSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingDeleteResponse struct {
	// Whether all accounts in the organization are trusted to use this account's
	// cross-account Access resources.
	GrantTrustToAllAccountsInOrg bool `json:"grant_trust_to_all_accounts_in_org" api:"required"`
	// Explicit account IDs trusted to use this account's cross-account Access
	// resources. At most 100 unique valid account IDs are accepted. Omitted from
	// responses when empty.
	TrustingAccounts []string                                   `json:"trusting_accounts"`
	JSON             organizationTrustSettingDeleteResponseJSON `json:"-"`
}

// organizationTrustSettingDeleteResponseJSON contains the JSON metadata for the
// struct [OrganizationTrustSettingDeleteResponse]
type organizationTrustSettingDeleteResponseJSON struct {
	GrantTrustToAllAccountsInOrg apijson.Field
	TrustingAccounts             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *OrganizationTrustSettingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingGetResponse struct {
	// Whether all accounts in the organization are trusted to use this account's
	// cross-account Access resources.
	GrantTrustToAllAccountsInOrg bool `json:"grant_trust_to_all_accounts_in_org" api:"required"`
	// Explicit account IDs trusted to use this account's cross-account Access
	// resources. At most 100 unique valid account IDs are accepted. Omitted from
	// responses when empty.
	TrustingAccounts []string                                `json:"trusting_accounts"`
	JSON             organizationTrustSettingGetResponseJSON `json:"-"`
}

// organizationTrustSettingGetResponseJSON contains the JSON metadata for the
// struct [OrganizationTrustSettingGetResponse]
type organizationTrustSettingGetResponseJSON struct {
	GrantTrustToAllAccountsInOrg apijson.Field
	TrustingAccounts             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *OrganizationTrustSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingReplaceResponse struct {
	// Whether all accounts in the organization are trusted to use this account's
	// cross-account Access resources.
	GrantTrustToAllAccountsInOrg bool `json:"grant_trust_to_all_accounts_in_org" api:"required"`
	// Explicit account IDs trusted to use this account's cross-account Access
	// resources. At most 100 unique valid account IDs are accepted. Omitted from
	// responses when empty.
	TrustingAccounts []string                                    `json:"trusting_accounts"`
	JSON             organizationTrustSettingReplaceResponseJSON `json:"-"`
}

// organizationTrustSettingReplaceResponseJSON contains the JSON metadata for the
// struct [OrganizationTrustSettingReplaceResponse]
type organizationTrustSettingReplaceResponseJSON struct {
	GrantTrustToAllAccountsInOrg apijson.Field
	TrustingAccounts             apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *OrganizationTrustSettingReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingReplaceResponseJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Whether all accounts in the organization are trusted to use this account's
	// cross-account Access resources.
	GrantTrustToAllAccountsInOrg param.Field[bool] `json:"grant_trust_to_all_accounts_in_org"`
	// Explicit account IDs trusted to use this account's cross-account Access
	// resources. At most 100 unique valid account IDs are accepted. Omitted from
	// responses when empty.
	TrustingAccounts param.Field[[]string] `json:"trusting_accounts"`
}

func (r OrganizationTrustSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationTrustSettingUpdateResponseEnvelope struct {
	Errors   []OrganizationTrustSettingUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OrganizationTrustSettingUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OrganizationTrustSettingUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OrganizationTrustSettingUpdateResponse                `json:"result"`
	JSON    organizationTrustSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// organizationTrustSettingUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [OrganizationTrustSettingUpdateResponseEnvelope]
type organizationTrustSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingUpdateResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           OrganizationTrustSettingUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationTrustSettingUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationTrustSettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OrganizationTrustSettingUpdateResponseEnvelopeErrors]
type organizationTrustSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    organizationTrustSettingUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationTrustSettingUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [OrganizationTrustSettingUpdateResponseEnvelopeErrorsSource]
type organizationTrustSettingUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingUpdateResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           OrganizationTrustSettingUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationTrustSettingUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationTrustSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OrganizationTrustSettingUpdateResponseEnvelopeMessages]
type organizationTrustSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    organizationTrustSettingUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationTrustSettingUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [OrganizationTrustSettingUpdateResponseEnvelopeMessagesSource]
type organizationTrustSettingUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationTrustSettingUpdateResponseEnvelopeSuccess bool

const (
	OrganizationTrustSettingUpdateResponseEnvelopeSuccessTrue OrganizationTrustSettingUpdateResponseEnvelopeSuccess = true
)

func (r OrganizationTrustSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationTrustSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationTrustSettingDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type OrganizationTrustSettingDeleteResponseEnvelope struct {
	Errors   []OrganizationTrustSettingDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OrganizationTrustSettingDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OrganizationTrustSettingDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OrganizationTrustSettingDeleteResponse                `json:"result"`
	JSON    organizationTrustSettingDeleteResponseEnvelopeJSON    `json:"-"`
}

// organizationTrustSettingDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [OrganizationTrustSettingDeleteResponseEnvelope]
type organizationTrustSettingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingDeleteResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           OrganizationTrustSettingDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationTrustSettingDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationTrustSettingDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OrganizationTrustSettingDeleteResponseEnvelopeErrors]
type organizationTrustSettingDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    organizationTrustSettingDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationTrustSettingDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [OrganizationTrustSettingDeleteResponseEnvelopeErrorsSource]
type organizationTrustSettingDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingDeleteResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           OrganizationTrustSettingDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationTrustSettingDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationTrustSettingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OrganizationTrustSettingDeleteResponseEnvelopeMessages]
type organizationTrustSettingDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    organizationTrustSettingDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationTrustSettingDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [OrganizationTrustSettingDeleteResponseEnvelopeMessagesSource]
type organizationTrustSettingDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationTrustSettingDeleteResponseEnvelopeSuccess bool

const (
	OrganizationTrustSettingDeleteResponseEnvelopeSuccessTrue OrganizationTrustSettingDeleteResponseEnvelopeSuccess = true
)

func (r OrganizationTrustSettingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationTrustSettingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationTrustSettingGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type OrganizationTrustSettingGetResponseEnvelope struct {
	Errors   []OrganizationTrustSettingGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OrganizationTrustSettingGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OrganizationTrustSettingGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OrganizationTrustSettingGetResponse                `json:"result"`
	JSON    organizationTrustSettingGetResponseEnvelopeJSON    `json:"-"`
}

// organizationTrustSettingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [OrganizationTrustSettingGetResponseEnvelope]
type organizationTrustSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingGetResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           OrganizationTrustSettingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationTrustSettingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationTrustSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [OrganizationTrustSettingGetResponseEnvelopeErrors]
type organizationTrustSettingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    organizationTrustSettingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationTrustSettingGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [OrganizationTrustSettingGetResponseEnvelopeErrorsSource]
type organizationTrustSettingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingGetResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           OrganizationTrustSettingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationTrustSettingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationTrustSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OrganizationTrustSettingGetResponseEnvelopeMessages]
type organizationTrustSettingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    organizationTrustSettingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationTrustSettingGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [OrganizationTrustSettingGetResponseEnvelopeMessagesSource]
type organizationTrustSettingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationTrustSettingGetResponseEnvelopeSuccess bool

const (
	OrganizationTrustSettingGetResponseEnvelopeSuccessTrue OrganizationTrustSettingGetResponseEnvelopeSuccess = true
)

func (r OrganizationTrustSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationTrustSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationTrustSettingReplaceParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Whether all accounts in the organization are trusted to use this account's
	// cross-account Access resources.
	GrantTrustToAllAccountsInOrg param.Field[bool] `json:"grant_trust_to_all_accounts_in_org" api:"required"`
	// Explicit account IDs trusted to use this account's cross-account Access
	// resources. At most 100 unique valid account IDs are accepted. Omitted from
	// responses when empty.
	TrustingAccounts param.Field[[]string] `json:"trusting_accounts"`
}

func (r OrganizationTrustSettingReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationTrustSettingReplaceResponseEnvelope struct {
	Errors   []OrganizationTrustSettingReplaceResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []OrganizationTrustSettingReplaceResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OrganizationTrustSettingReplaceResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  OrganizationTrustSettingReplaceResponse                `json:"result"`
	JSON    organizationTrustSettingReplaceResponseEnvelopeJSON    `json:"-"`
}

// organizationTrustSettingReplaceResponseEnvelopeJSON contains the JSON metadata
// for the struct [OrganizationTrustSettingReplaceResponseEnvelope]
type organizationTrustSettingReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingReplaceResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingReplaceResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           OrganizationTrustSettingReplaceResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationTrustSettingReplaceResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationTrustSettingReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OrganizationTrustSettingReplaceResponseEnvelopeErrors]
type organizationTrustSettingReplaceResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingReplaceResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingReplaceResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    organizationTrustSettingReplaceResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationTrustSettingReplaceResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [OrganizationTrustSettingReplaceResponseEnvelopeErrorsSource]
type organizationTrustSettingReplaceResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingReplaceResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingReplaceResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingReplaceResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           OrganizationTrustSettingReplaceResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationTrustSettingReplaceResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationTrustSettingReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [OrganizationTrustSettingReplaceResponseEnvelopeMessages]
type organizationTrustSettingReplaceResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationTrustSettingReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingReplaceResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationTrustSettingReplaceResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    organizationTrustSettingReplaceResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationTrustSettingReplaceResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [OrganizationTrustSettingReplaceResponseEnvelopeMessagesSource]
type organizationTrustSettingReplaceResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationTrustSettingReplaceResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationTrustSettingReplaceResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationTrustSettingReplaceResponseEnvelopeSuccess bool

const (
	OrganizationTrustSettingReplaceResponseEnvelopeSuccessTrue OrganizationTrustSettingReplaceResponseEnvelopeSuccess = true
)

func (r OrganizationTrustSettingReplaceResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationTrustSettingReplaceResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
