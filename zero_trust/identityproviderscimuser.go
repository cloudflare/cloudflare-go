// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// IdentityProviderSCIMUserService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityProviderSCIMUserService] method instead.
type IdentityProviderSCIMUserService struct {
	Options []option.RequestOption
}

// NewIdentityProviderSCIMUserService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIdentityProviderSCIMUserService(opts ...option.RequestOption) (r *IdentityProviderSCIMUserService) {
	r = &IdentityProviderSCIMUserService{}
	r.Options = opts
	return
}

// Lists SCIM User resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *IdentityProviderSCIMUserService) List(ctx context.Context, identityProviderID string, params IdentityProviderSCIMUserListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessUser], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/users", params.AccountID, identityProviderID)
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

// Lists SCIM User resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *IdentityProviderSCIMUserService) ListAutoPaging(ctx context.Context, identityProviderID string, params IdentityProviderSCIMUserListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessUser] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, identityProviderID, params, opts...))
}

// Deletes a SCIM User resource from the identity provider's SCIM population.
func (r *IdentityProviderSCIMUserService) Delete(ctx context.Context, identityProviderID string, userID string, body IdentityProviderSCIMUserDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/users/%s", body.AccountID, identityProviderID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns a SCIM User resource synced to Cloudflare via the System for
// Cross-domain Identity Management (SCIM).
func (r *IdentityProviderSCIMUserService) Get(ctx context.Context, identityProviderID string, userID string, query IdentityProviderSCIMUserGetParams, opts ...option.RequestOption) (res *AccessUser, err error) {
	var env IdentityProviderSCIMUserGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return nil, err
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/users/%s", query.AccountID, identityProviderID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type IdentityProviderSCIMUserListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The unique Cloudflare-generated Id of the SCIM User resource; also known as the
	// "Id". Pass once for a single lookup (`?cf_resource_id=A`) or repeat the
	// parameter (`?cf_resource_id=A&cf_resource_id=B`) to look up multiple users in
	// one request, up to 50 values. Mutually exclusive with `idp_resource_id`,
	// `username`, `email`, `name`, `search_contains`, and `search_starts_with`.
	CfResourceID param.Field[[]string] `query:"cf_resource_id"`
	// The email address of the SCIM User resource.
	Email param.Field[string] `query:"email"`
	// The IdP-generated Id of the SCIM User resource; also known as the "external Id".
	// Pass once for a single lookup (`?idp_resource_id=A`) or repeat the parameter
	// (`?idp_resource_id=A&idp_resource_id=B`) to look up multiple users in one
	// request, up to 50 values. Mutually exclusive with `cf_resource_id`, `username`,
	// `email`, `name`, `search_contains`, and `search_starts_with`.
	IdPResourceID param.Field[[]string] `query:"idp_resource_id"`
	// The name of the SCIM User resource.
	Name param.Field[string] `query:"name"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The username of the SCIM User resource.
	Username param.Field[string] `query:"username"`
}

// URLQuery serializes [IdentityProviderSCIMUserListParams]'s query parameters as
// `url.Values`.
func (r IdentityProviderSCIMUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IdentityProviderSCIMUserDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type IdentityProviderSCIMUserGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type IdentityProviderSCIMUserGetResponseEnvelope struct {
	Errors   []IdentityProviderSCIMUserGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []IdentityProviderSCIMUserGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success IdentityProviderSCIMUserGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUser                                         `json:"result"`
	JSON    identityProviderSCIMUserGetResponseEnvelopeJSON    `json:"-"`
}

// identityProviderSCIMUserGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [IdentityProviderSCIMUserGetResponseEnvelope]
type identityProviderSCIMUserGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMUserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMUserGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMUserGetResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           IdentityProviderSCIMUserGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             identityProviderSCIMUserGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// identityProviderSCIMUserGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [IdentityProviderSCIMUserGetResponseEnvelopeErrors]
type identityProviderSCIMUserGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityProviderSCIMUserGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMUserGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMUserGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    identityProviderSCIMUserGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// identityProviderSCIMUserGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [IdentityProviderSCIMUserGetResponseEnvelopeErrorsSource]
type identityProviderSCIMUserGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMUserGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMUserGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMUserGetResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           IdentityProviderSCIMUserGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             identityProviderSCIMUserGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// identityProviderSCIMUserGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IdentityProviderSCIMUserGetResponseEnvelopeMessages]
type identityProviderSCIMUserGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityProviderSCIMUserGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMUserGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMUserGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    identityProviderSCIMUserGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// identityProviderSCIMUserGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [IdentityProviderSCIMUserGetResponseEnvelopeMessagesSource]
type identityProviderSCIMUserGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMUserGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMUserGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IdentityProviderSCIMUserGetResponseEnvelopeSuccess bool

const (
	IdentityProviderSCIMUserGetResponseEnvelopeSuccessTrue IdentityProviderSCIMUserGetResponseEnvelopeSuccess = true
)

func (r IdentityProviderSCIMUserGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderSCIMUserGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
