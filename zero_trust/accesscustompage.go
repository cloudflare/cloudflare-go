// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccessCustomPageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessCustomPageService] method instead.
type AccessCustomPageService struct {
	Options []option.RequestOption
}

// NewAccessCustomPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessCustomPageService(opts ...option.RequestOption) (r *AccessCustomPageService) {
	r = &AccessCustomPageService{}
	r.Options = opts
	return
}

// Create a custom page
func (r *AccessCustomPageService) New(ctx context.Context, params AccessCustomPageNewParams, opts ...option.RequestOption) (res *CustomPageWithoutHTML, err error) {
	var env AccessCustomPageNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a custom page
func (r *AccessCustomPageService) Update(ctx context.Context, customPageID string, params AccessCustomPageUpdateParams, opts ...option.RequestOption) (res *CustomPageWithoutHTML, err error) {
	var env AccessCustomPageUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customPageID == "" {
		err = errors.New("missing required custom_page_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", params.AccountID, customPageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List custom pages
func (r *AccessCustomPageService) List(ctx context.Context, query AccessCustomPageListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CustomPageWithoutHTML], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// List custom pages
func (r *AccessCustomPageService) ListAutoPaging(ctx context.Context, query AccessCustomPageListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CustomPageWithoutHTML] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a custom page
func (r *AccessCustomPageService) Delete(ctx context.Context, customPageID string, body AccessCustomPageDeleteParams, opts ...option.RequestOption) (res *AccessCustomPageDeleteResponse, err error) {
	var env AccessCustomPageDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customPageID == "" {
		err = errors.New("missing required custom_page_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", body.AccountID, customPageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom page and also returns its HTML.
func (r *AccessCustomPageService) Get(ctx context.Context, customPageID string, query AccessCustomPageGetParams, opts ...option.RequestOption) (res *CustomPage, err error) {
	var env AccessCustomPageGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customPageID == "" {
		err = errors.New("missing required custom_page_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", query.AccountID, customPageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomPage struct {
	// Custom page HTML.
	CustomHTML string `json:"custom_html,required"`
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type CustomPageType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	UID       string         `json:"uid"`
	UpdatedAt time.Time      `json:"updated_at" format:"date-time"`
	JSON      customPageJSON `json:"-"`
}

// customPageJSON contains the JSON metadata for the struct [CustomPage]
type customPageJSON struct {
	CustomHTML  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UID         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type CustomPageType string

const (
	CustomPageTypeIdentityDenied CustomPageType = "identity_denied"
	CustomPageTypeForbidden      CustomPageType = "forbidden"
)

func (r CustomPageType) IsKnown() bool {
	switch r {
	case CustomPageTypeIdentityDenied, CustomPageTypeForbidden:
		return true
	}
	return false
}

type CustomPageParam struct {
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[CustomPageType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r CustomPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomPageWithoutHTML struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type CustomPageWithoutHTMLType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	UID       string                    `json:"uid"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      customPageWithoutHTMLJSON `json:"-"`
}

// customPageWithoutHTMLJSON contains the JSON metadata for the struct
// [CustomPageWithoutHTML]
type customPageWithoutHTMLJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	UID         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPageWithoutHTML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPageWithoutHTMLJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type CustomPageWithoutHTMLType string

const (
	CustomPageWithoutHTMLTypeIdentityDenied CustomPageWithoutHTMLType = "identity_denied"
	CustomPageWithoutHTMLTypeForbidden      CustomPageWithoutHTMLType = "forbidden"
)

func (r CustomPageWithoutHTMLType) IsKnown() bool {
	switch r {
	case CustomPageWithoutHTMLTypeIdentityDenied, CustomPageWithoutHTMLTypeForbidden:
		return true
	}
	return false
}

type AccessCustomPageDeleteResponse struct {
	// UUID
	ID   string                             `json:"id"`
	JSON accessCustomPageDeleteResponseJSON `json:"-"`
}

// accessCustomPageDeleteResponseJSON contains the JSON metadata for the struct
// [AccessCustomPageDeleteResponse]
type accessCustomPageDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCustomPageDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessCustomPageNewParams struct {
	// Identifier
	AccountID  param.Field[string] `path:"account_id,required"`
	CustomPage CustomPageParam     `json:"custom_page,required"`
}

func (r AccessCustomPageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomPage)
}

type AccessCustomPageNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCustomPageNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomPageWithoutHTML                      `json:"result"`
	JSON    accessCustomPageNewResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageNewResponseEnvelope]
type accessCustomPageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCustomPageNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCustomPageNewResponseEnvelopeSuccess bool

const (
	AccessCustomPageNewResponseEnvelopeSuccessTrue AccessCustomPageNewResponseEnvelopeSuccess = true
)

func (r AccessCustomPageNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCustomPageNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCustomPageUpdateParams struct {
	// Identifier
	AccountID  param.Field[string] `path:"account_id,required"`
	CustomPage CustomPageParam     `json:"custom_page,required"`
}

func (r AccessCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CustomPage)
}

type AccessCustomPageUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCustomPageUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomPageWithoutHTML                         `json:"result"`
	JSON    accessCustomPageUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageUpdateResponseEnvelope]
type accessCustomPageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCustomPageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCustomPageUpdateResponseEnvelopeSuccess bool

const (
	AccessCustomPageUpdateResponseEnvelopeSuccessTrue AccessCustomPageUpdateResponseEnvelopeSuccess = true
)

func (r AccessCustomPageUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCustomPageUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCustomPageListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessCustomPageDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessCustomPageDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCustomPageDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessCustomPageDeleteResponse                `json:"result"`
	JSON    accessCustomPageDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageDeleteResponseEnvelope]
type accessCustomPageDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCustomPageDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCustomPageDeleteResponseEnvelopeSuccess bool

const (
	AccessCustomPageDeleteResponseEnvelopeSuccessTrue AccessCustomPageDeleteResponseEnvelopeSuccess = true
)

func (r AccessCustomPageDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCustomPageDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCustomPageGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessCustomPageGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessCustomPageGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CustomPage                                 `json:"result"`
	JSON    accessCustomPageGetResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageGetResponseEnvelope]
type accessCustomPageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCustomPageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCustomPageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCustomPageGetResponseEnvelopeSuccess bool

const (
	AccessCustomPageGetResponseEnvelopeSuccessTrue AccessCustomPageGetResponseEnvelopeSuccess = true
)

func (r AccessCustomPageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCustomPageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
