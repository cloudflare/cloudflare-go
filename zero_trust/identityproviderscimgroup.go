// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// IdentityProviderSCIMGroupService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityProviderSCIMGroupService] method instead.
type IdentityProviderSCIMGroupService struct {
	Options []option.RequestOption
}

// NewIdentityProviderSCIMGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIdentityProviderSCIMGroupService(opts ...option.RequestOption) (r *IdentityProviderSCIMGroupService) {
	r = &IdentityProviderSCIMGroupService{}
	r.Options = opts
	return
}

// Lists SCIM Group resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *IdentityProviderSCIMGroupService) List(ctx context.Context, identityProviderID string, params IdentityProviderSCIMGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ZeroTrustGroup], err error) {
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
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/groups", params.AccountID, identityProviderID)
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

// Lists SCIM Group resources synced to Cloudflare via the System for Cross-domain
// Identity Management (SCIM).
func (r *IdentityProviderSCIMGroupService) ListAutoPaging(ctx context.Context, identityProviderID string, params IdentityProviderSCIMGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ZeroTrustGroup] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, identityProviderID, params, opts...))
}

// Deletes a SCIM Group resource from the identity provider's SCIM population.
func (r *IdentityProviderSCIMGroupService) Delete(ctx context.Context, identityProviderID string, groupID string, body IdentityProviderSCIMGroupDeleteParams, opts ...option.RequestOption) (err error) {
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
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/groups/%s", body.AccountID, identityProviderID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns a SCIM Group resource, including its direct members, synced to
// Cloudflare via the System for Cross-domain Identity Management (SCIM).
func (r *IdentityProviderSCIMGroupService) Get(ctx context.Context, identityProviderID string, groupID string, query IdentityProviderSCIMGroupGetParams, opts ...option.RequestOption) (res *IdentityProviderSCIMGroupGetResponse, err error) {
	var env IdentityProviderSCIMGroupGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/identity_providers/%s/scim/groups/%s", query.AccountID, identityProviderID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type IdentityProviderSCIMGroupGetResponse struct {
	// The unique Cloudflare-generated Id of the SCIM resource.
	ID string `json:"id"`
	// The display name of the SCIM Group resource.
	DisplayName string `json:"displayName"`
	// The IdP-generated Id of the SCIM resource.
	ExternalID string `json:"externalId"`
	// The direct members of the SCIM Group resource.
	Members []IdentityProviderSCIMGroupGetResponseMember `json:"members"`
	// The metadata of the SCIM resource.
	Meta IdentityProviderSCIMGroupGetResponseMeta `json:"meta"`
	// The list of URIs which indicate the attributes contained within a SCIM resource.
	Schemas []string                                 `json:"schemas"`
	JSON    identityProviderSCIMGroupGetResponseJSON `json:"-"`
}

// identityProviderSCIMGroupGetResponseJSON contains the JSON metadata for the
// struct [IdentityProviderSCIMGroupGetResponse]
type identityProviderSCIMGroupGetResponseJSON struct {
	ID          apijson.Field
	DisplayName apijson.Field
	ExternalID  apijson.Field
	Members     apijson.Field
	Meta        apijson.Field
	Schemas     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMGroupGetResponseMember struct {
	// The unique Cloudflare-generated Id of the member resource.
	Value string `json:"value" api:"required"`
	// The URI of the member resource.
	Ref string `json:"$ref" format:"uri"`
	// The display name of the member resource.
	Display string `json:"display"`
	// The SCIM resource type of the member.
	Type IdentityProviderSCIMGroupGetResponseMembersType `json:"type"`
	JSON identityProviderSCIMGroupGetResponseMemberJSON  `json:"-"`
}

// identityProviderSCIMGroupGetResponseMemberJSON contains the JSON metadata for
// the struct [IdentityProviderSCIMGroupGetResponseMember]
type identityProviderSCIMGroupGetResponseMemberJSON struct {
	Value       apijson.Field
	Ref         apijson.Field
	Display     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseMemberJSON) RawJSON() string {
	return r.raw
}

// The SCIM resource type of the member.
type IdentityProviderSCIMGroupGetResponseMembersType string

const (
	IdentityProviderSCIMGroupGetResponseMembersTypeUser  IdentityProviderSCIMGroupGetResponseMembersType = "User"
	IdentityProviderSCIMGroupGetResponseMembersTypeGroup IdentityProviderSCIMGroupGetResponseMembersType = "Group"
)

func (r IdentityProviderSCIMGroupGetResponseMembersType) IsKnown() bool {
	switch r {
	case IdentityProviderSCIMGroupGetResponseMembersTypeUser, IdentityProviderSCIMGroupGetResponseMembersTypeGroup:
		return true
	}
	return false
}

// The metadata of the SCIM resource.
type IdentityProviderSCIMGroupGetResponseMeta struct {
	// The timestamp of when the SCIM resource was created.
	Created time.Time `json:"created" format:"date-time"`
	// The timestamp of when the SCIM resource was last modified.
	LastModified time.Time                                    `json:"lastModified" format:"date-time"`
	JSON         identityProviderSCIMGroupGetResponseMetaJSON `json:"-"`
}

// identityProviderSCIMGroupGetResponseMetaJSON contains the JSON metadata for the
// struct [IdentityProviderSCIMGroupGetResponseMeta]
type identityProviderSCIMGroupGetResponseMetaJSON struct {
	Created      apijson.Field
	LastModified apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMGroupListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The unique Cloudflare-generated Id of the SCIM Group resource; also known as the
	// "Id". Pass once for a single lookup (`?cf_resource_id=A`) or repeat the
	// parameter (`?cf_resource_id=A&cf_resource_id=B`) to look up multiple groups in
	// one request, up to 50 values. Mutually exclusive with `idp_resource_id`, `name`,
	// `search_contains`, and `search_starts_with`.
	CfResourceID param.Field[[]string] `query:"cf_resource_id"`
	// The IdP-generated Id of the SCIM Group resource; also known as the "external
	// Id". Pass once for a single lookup (`?idp_resource_id=A`) or repeat the
	// parameter (`?idp_resource_id=A&idp_resource_id=B`) to look up multiple groups in
	// one request, up to 50 values. Mutually exclusive with `cf_resource_id`, `name`,
	// `search_contains`, and `search_starts_with`.
	IdPResourceID param.Field[[]string] `query:"idp_resource_id"`
	// The display name of the SCIM Group resource.
	Name param.Field[string] `query:"name"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [IdentityProviderSCIMGroupListParams]'s query parameters as
// `url.Values`.
func (r IdentityProviderSCIMGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IdentityProviderSCIMGroupDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type IdentityProviderSCIMGroupGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type IdentityProviderSCIMGroupGetResponseEnvelope struct {
	Errors   []IdentityProviderSCIMGroupGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []IdentityProviderSCIMGroupGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success IdentityProviderSCIMGroupGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  IdentityProviderSCIMGroupGetResponse                `json:"result"`
	JSON    identityProviderSCIMGroupGetResponseEnvelopeJSON    `json:"-"`
}

// identityProviderSCIMGroupGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [IdentityProviderSCIMGroupGetResponseEnvelope]
type identityProviderSCIMGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMGroupGetResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           IdentityProviderSCIMGroupGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             identityProviderSCIMGroupGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// identityProviderSCIMGroupGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [IdentityProviderSCIMGroupGetResponseEnvelopeErrors]
type identityProviderSCIMGroupGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMGroupGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    identityProviderSCIMGroupGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// identityProviderSCIMGroupGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [IdentityProviderSCIMGroupGetResponseEnvelopeErrorsSource]
type identityProviderSCIMGroupGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMGroupGetResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           IdentityProviderSCIMGroupGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             identityProviderSCIMGroupGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// identityProviderSCIMGroupGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [IdentityProviderSCIMGroupGetResponseEnvelopeMessages]
type identityProviderSCIMGroupGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderSCIMGroupGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    identityProviderSCIMGroupGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// identityProviderSCIMGroupGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [IdentityProviderSCIMGroupGetResponseEnvelopeMessagesSource]
type identityProviderSCIMGroupGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderSCIMGroupGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderSCIMGroupGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IdentityProviderSCIMGroupGetResponseEnvelopeSuccess bool

const (
	IdentityProviderSCIMGroupGetResponseEnvelopeSuccessTrue IdentityProviderSCIMGroupGetResponseEnvelopeSuccess = true
)

func (r IdentityProviderSCIMGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderSCIMGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
