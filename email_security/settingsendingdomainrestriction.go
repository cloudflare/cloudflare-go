// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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

// SettingSendingDomainRestrictionService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingSendingDomainRestrictionService] method instead.
type SettingSendingDomainRestrictionService struct {
	Options []option.RequestOption
}

// NewSettingSendingDomainRestrictionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingSendingDomainRestrictionService(opts ...option.RequestOption) (r *SettingSendingDomainRestrictionService) {
	r = &SettingSendingDomainRestrictionService{}
	r.Options = opts
	return
}

// Creates a new sending domain restriction to enforce TLS requirements for a
// domain. Emails without TLS from this domain will be dropped unless the subdomain
// is in the exclude list.
func (r *SettingSendingDomainRestrictionService) New(ctx context.Context, params SettingSendingDomainRestrictionNewParams, opts ...option.RequestOption) (res *SettingSendingDomainRestrictionNewResponse, err error) {
	var env SettingSendingDomainRestrictionNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/sending_domain_restrictions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of sending domain restrictions. These restrictions
// enforce TLS requirements for emails from specific domains. Mail without TLS from
// restricted domains will be dropped unless the subdomain is in the exclude list.
// Supports sorting and searching.
func (r *SettingSendingDomainRestrictionService) List(ctx context.Context, params SettingSendingDomainRestrictionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingSendingDomainRestrictionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/sending_domain_restrictions", params.AccountID)
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

// Returns a paginated list of sending domain restrictions. These restrictions
// enforce TLS requirements for emails from specific domains. Mail without TLS from
// restricted domains will be dropped unless the subdomain is in the exclude list.
// Supports sorting and searching.
func (r *SettingSendingDomainRestrictionService) ListAutoPaging(ctx context.Context, params SettingSendingDomainRestrictionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingSendingDomainRestrictionListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes a sending domain restriction. After deletion, TLS will no longer be
// enforced for emails from this domain.
func (r *SettingSendingDomainRestrictionService) Delete(ctx context.Context, sendingDomainRestrictionID string, body SettingSendingDomainRestrictionDeleteParams, opts ...option.RequestOption) (res *SettingSendingDomainRestrictionDeleteResponse, err error) {
	var env SettingSendingDomainRestrictionDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sendingDomainRestrictionID == "" {
		err = errors.New("missing required sending_domain_restriction_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/sending_domain_restrictions/%s", body.AccountID, sendingDomainRestrictionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing sending domain restriction. Only provided fields will be
// modified. Changes affect which domains require TLS and which subdomains are
// excluded.
func (r *SettingSendingDomainRestrictionService) Edit(ctx context.Context, sendingDomainRestrictionID string, params SettingSendingDomainRestrictionEditParams, opts ...option.RequestOption) (res *SettingSendingDomainRestrictionEditResponse, err error) {
	var env SettingSendingDomainRestrictionEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sendingDomainRestrictionID == "" {
		err = errors.New("missing required sending_domain_restriction_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/sending_domain_restrictions/%s", params.AccountID, sendingDomainRestrictionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific sending domain restriction including the domain
// requiring TLS and any excluded subdomains exempt from the TLS requirement.
func (r *SettingSendingDomainRestrictionService) Get(ctx context.Context, sendingDomainRestrictionID string, query SettingSendingDomainRestrictionGetParams, opts ...option.RequestOption) (res *SettingSendingDomainRestrictionGetResponse, err error) {
	var env SettingSendingDomainRestrictionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sendingDomainRestrictionID == "" {
		err = errors.New("missing required sending_domain_restriction_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/sending_domain_restrictions/%s", query.AccountID, sendingDomainRestrictionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A sending domain restriction that enforces TLS (Transport Layer Security)
// requirements for emails from specific domains. If TLS is required, mail without
// TLS from the specified domain will be dropped.
type SettingSendingDomainRestrictionNewResponse struct {
	// Sending domain restriction identifier.
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Domain that requires TLS enforcement.
	Domain string `json:"domain"`
	// Excluded subdomains that are exempt from TLS requirements.
	Exclude []string `json:"exclude"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                                      `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                                      `json:"modified_at" format:"date-time"`
	JSON         settingSendingDomainRestrictionNewResponseJSON `json:"-"`
}

// settingSendingDomainRestrictionNewResponseJSON contains the JSON metadata for
// the struct [SettingSendingDomainRestrictionNewResponse]
type settingSendingDomainRestrictionNewResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	Domain       apijson.Field
	Exclude      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionNewResponseJSON) RawJSON() string {
	return r.raw
}

// A sending domain restriction that enforces TLS (Transport Layer Security)
// requirements for emails from specific domains. If TLS is required, mail without
// TLS from the specified domain will be dropped.
type SettingSendingDomainRestrictionListResponse struct {
	// Sending domain restriction identifier.
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Domain that requires TLS enforcement.
	Domain string `json:"domain"`
	// Excluded subdomains that are exempt from TLS requirements.
	Exclude []string `json:"exclude"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                                       `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                                       `json:"modified_at" format:"date-time"`
	JSON         settingSendingDomainRestrictionListResponseJSON `json:"-"`
}

// settingSendingDomainRestrictionListResponseJSON contains the JSON metadata for
// the struct [SettingSendingDomainRestrictionListResponse]
type settingSendingDomainRestrictionListResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	Domain       apijson.Field
	Exclude      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionDeleteResponse struct {
	// Sending domain restriction identifier.
	ID   string                                            `json:"id" api:"required" format:"uuid"`
	JSON settingSendingDomainRestrictionDeleteResponseJSON `json:"-"`
}

// settingSendingDomainRestrictionDeleteResponseJSON contains the JSON metadata for
// the struct [SettingSendingDomainRestrictionDeleteResponse]
type settingSendingDomainRestrictionDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A sending domain restriction that enforces TLS (Transport Layer Security)
// requirements for emails from specific domains. If TLS is required, mail without
// TLS from the specified domain will be dropped.
type SettingSendingDomainRestrictionEditResponse struct {
	// Sending domain restriction identifier.
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Domain that requires TLS enforcement.
	Domain string `json:"domain"`
	// Excluded subdomains that are exempt from TLS requirements.
	Exclude []string `json:"exclude"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                                       `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                                       `json:"modified_at" format:"date-time"`
	JSON         settingSendingDomainRestrictionEditResponseJSON `json:"-"`
}

// settingSendingDomainRestrictionEditResponseJSON contains the JSON metadata for
// the struct [SettingSendingDomainRestrictionEditResponse]
type settingSendingDomainRestrictionEditResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	Domain       apijson.Field
	Exclude      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionEditResponseJSON) RawJSON() string {
	return r.raw
}

// A sending domain restriction that enforces TLS (Transport Layer Security)
// requirements for emails from specific domains. If TLS is required, mail without
// TLS from the specified domain will be dropped.
type SettingSendingDomainRestrictionGetResponse struct {
	// Sending domain restriction identifier.
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Domain that requires TLS enforcement.
	Domain string `json:"domain"`
	// Excluded subdomains that are exempt from TLS requirements.
	Exclude []string `json:"exclude"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                                      `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                                      `json:"modified_at" format:"date-time"`
	JSON         settingSendingDomainRestrictionGetResponseJSON `json:"-"`
}

// settingSendingDomainRestrictionGetResponseJSON contains the JSON metadata for
// the struct [SettingSendingDomainRestrictionGetResponse]
type settingSendingDomainRestrictionGetResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	Domain       apijson.Field
	Exclude      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Domain that requires TLS enforcement.
	Domain param.Field[string] `json:"domain" api:"required"`
	// Excluded subdomains that are exempt from TLS requirements.
	Exclude  param.Field[[]string] `json:"exclude" api:"required"`
	Comments param.Field[string]   `json:"comments"`
}

func (r SettingSendingDomainRestrictionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSendingDomainRestrictionNewResponseEnvelope struct {
	Errors   []SettingSendingDomainRestrictionNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingSendingDomainRestrictionNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingSendingDomainRestrictionNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A sending domain restriction that enforces TLS (Transport Layer Security)
	// requirements for emails from specific domains. If TLS is required, mail without
	// TLS from the specified domain will be dropped.
	Result SettingSendingDomainRestrictionNewResponse             `json:"result"`
	JSON   settingSendingDomainRestrictionNewResponseEnvelopeJSON `json:"-"`
}

// settingSendingDomainRestrictionNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingSendingDomainRestrictionNewResponseEnvelope]
type settingSendingDomainRestrictionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionNewResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingSendingDomainRestrictionNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingSendingDomainRestrictionNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingSendingDomainRestrictionNewResponseEnvelopeErrors]
type settingSendingDomainRestrictionNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    settingSendingDomainRestrictionNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionNewResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingSendingDomainRestrictionNewResponseEnvelopeErrorsSource]
type settingSendingDomainRestrictionNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionNewResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code" api:"required"`
	Message          string                                                           `json:"message" api:"required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingSendingDomainRestrictionNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingSendingDomainRestrictionNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingSendingDomainRestrictionNewResponseEnvelopeMessages]
type settingSendingDomainRestrictionNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    settingSendingDomainRestrictionNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionNewResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [SettingSendingDomainRestrictionNewResponseEnvelopeMessagesSource]
type settingSendingDomainRestrictionNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingSendingDomainRestrictionNewResponseEnvelopeSuccess bool

const (
	SettingSendingDomainRestrictionNewResponseEnvelopeSuccessTrue SettingSendingDomainRestrictionNewResponseEnvelopeSuccess = true
)

func (r SettingSendingDomainRestrictionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingSendingDomainRestrictionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingSendingDomainRestrictionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The sorting direction.
	Direction param.Field[SettingSendingDomainRestrictionListParamsDirection] `query:"direction"`
	// Field to sort by.
	Order param.Field[SettingSendingDomainRestrictionListParamsOrder] `query:"order"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
	// Search term for filtering records. Behavior may change.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingSendingDomainRestrictionListParams]'s query
// parameters as `url.Values`.
func (r SettingSendingDomainRestrictionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingSendingDomainRestrictionListParamsDirection string

const (
	SettingSendingDomainRestrictionListParamsDirectionAsc  SettingSendingDomainRestrictionListParamsDirection = "asc"
	SettingSendingDomainRestrictionListParamsDirectionDesc SettingSendingDomainRestrictionListParamsDirection = "desc"
)

func (r SettingSendingDomainRestrictionListParamsDirection) IsKnown() bool {
	switch r {
	case SettingSendingDomainRestrictionListParamsDirectionAsc, SettingSendingDomainRestrictionListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by.
type SettingSendingDomainRestrictionListParamsOrder string

const (
	SettingSendingDomainRestrictionListParamsOrderDomain    SettingSendingDomainRestrictionListParamsOrder = "domain"
	SettingSendingDomainRestrictionListParamsOrderCreatedAt SettingSendingDomainRestrictionListParamsOrder = "created_at"
)

func (r SettingSendingDomainRestrictionListParamsOrder) IsKnown() bool {
	switch r {
	case SettingSendingDomainRestrictionListParamsOrderDomain, SettingSendingDomainRestrictionListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingSendingDomainRestrictionDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingSendingDomainRestrictionDeleteResponseEnvelope struct {
	Errors   []SettingSendingDomainRestrictionDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingSendingDomainRestrictionDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingSendingDomainRestrictionDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingSendingDomainRestrictionDeleteResponse                `json:"result"`
	JSON    settingSendingDomainRestrictionDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingSendingDomainRestrictionDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingSendingDomainRestrictionDeleteResponseEnvelope]
type settingSendingDomainRestrictionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionDeleteResponseEnvelopeErrors struct {
	Code             int64                                                             `json:"code" api:"required"`
	Message          string                                                            `json:"message" api:"required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [SettingSendingDomainRestrictionDeleteResponseEnvelopeErrors]
type settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [SettingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSource]
type settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionDeleteResponseEnvelopeMessages struct {
	Code             int64                                                               `json:"code" api:"required"`
	Message          string                                                              `json:"message" api:"required"`
	DocumentationURL string                                                              `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SettingSendingDomainRestrictionDeleteResponseEnvelopeMessages]
type settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                                  `json:"pointer"`
	JSON    settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [SettingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSource]
type settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingSendingDomainRestrictionDeleteResponseEnvelopeSuccess bool

const (
	SettingSendingDomainRestrictionDeleteResponseEnvelopeSuccessTrue SettingSendingDomainRestrictionDeleteResponseEnvelopeSuccess = true
)

func (r SettingSendingDomainRestrictionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingSendingDomainRestrictionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingSendingDomainRestrictionEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Comments  param.Field[string] `json:"comments"`
	// Domain that requires TLS enforcement.
	Domain param.Field[string] `json:"domain"`
	// Excluded subdomains that are exempt from TLS requirements.
	Exclude param.Field[[]string] `json:"exclude"`
}

func (r SettingSendingDomainRestrictionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingSendingDomainRestrictionEditResponseEnvelope struct {
	Errors   []SettingSendingDomainRestrictionEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingSendingDomainRestrictionEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingSendingDomainRestrictionEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// A sending domain restriction that enforces TLS (Transport Layer Security)
	// requirements for emails from specific domains. If TLS is required, mail without
	// TLS from the specified domain will be dropped.
	Result SettingSendingDomainRestrictionEditResponse             `json:"result"`
	JSON   settingSendingDomainRestrictionEditResponseEnvelopeJSON `json:"-"`
}

// settingSendingDomainRestrictionEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingSendingDomainRestrictionEditResponseEnvelope]
type settingSendingDomainRestrictionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionEditResponseEnvelopeErrors struct {
	Code             int64                                                           `json:"code" api:"required"`
	Message          string                                                          `json:"message" api:"required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingSendingDomainRestrictionEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingSendingDomainRestrictionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingSendingDomainRestrictionEditResponseEnvelopeErrors]
type settingSendingDomainRestrictionEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    settingSendingDomainRestrictionEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionEditResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingSendingDomainRestrictionEditResponseEnvelopeErrorsSource]
type settingSendingDomainRestrictionEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionEditResponseEnvelopeMessages struct {
	Code             int64                                                             `json:"code" api:"required"`
	Message          string                                                            `json:"message" api:"required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingSendingDomainRestrictionEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingSendingDomainRestrictionEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [SettingSendingDomainRestrictionEditResponseEnvelopeMessages]
type settingSendingDomainRestrictionEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    settingSendingDomainRestrictionEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionEditResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [SettingSendingDomainRestrictionEditResponseEnvelopeMessagesSource]
type settingSendingDomainRestrictionEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingSendingDomainRestrictionEditResponseEnvelopeSuccess bool

const (
	SettingSendingDomainRestrictionEditResponseEnvelopeSuccessTrue SettingSendingDomainRestrictionEditResponseEnvelopeSuccess = true
)

func (r SettingSendingDomainRestrictionEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingSendingDomainRestrictionEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingSendingDomainRestrictionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingSendingDomainRestrictionGetResponseEnvelope struct {
	Errors   []SettingSendingDomainRestrictionGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingSendingDomainRestrictionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingSendingDomainRestrictionGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A sending domain restriction that enforces TLS (Transport Layer Security)
	// requirements for emails from specific domains. If TLS is required, mail without
	// TLS from the specified domain will be dropped.
	Result SettingSendingDomainRestrictionGetResponse             `json:"result"`
	JSON   settingSendingDomainRestrictionGetResponseEnvelopeJSON `json:"-"`
}

// settingSendingDomainRestrictionGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [SettingSendingDomainRestrictionGetResponseEnvelope]
type settingSendingDomainRestrictionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionGetResponseEnvelopeErrors struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingSendingDomainRestrictionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingSendingDomainRestrictionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [SettingSendingDomainRestrictionGetResponseEnvelopeErrors]
type settingSendingDomainRestrictionGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    settingSendingDomainRestrictionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [SettingSendingDomainRestrictionGetResponseEnvelopeErrorsSource]
type settingSendingDomainRestrictionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionGetResponseEnvelopeMessages struct {
	Code             int64                                                            `json:"code" api:"required"`
	Message          string                                                           `json:"message" api:"required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           SettingSendingDomainRestrictionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingSendingDomainRestrictionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingSendingDomainRestrictionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [SettingSendingDomainRestrictionGetResponseEnvelopeMessages]
type settingSendingDomainRestrictionGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingSendingDomainRestrictionGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    settingSendingDomainRestrictionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingSendingDomainRestrictionGetResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [SettingSendingDomainRestrictionGetResponseEnvelopeMessagesSource]
type settingSendingDomainRestrictionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingSendingDomainRestrictionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingSendingDomainRestrictionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingSendingDomainRestrictionGetResponseEnvelopeSuccess bool

const (
	SettingSendingDomainRestrictionGetResponseEnvelopeSuccessTrue SettingSendingDomainRestrictionGetResponseEnvelopeSuccess = true
)

func (r SettingSendingDomainRestrictionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingSendingDomainRestrictionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
