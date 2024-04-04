// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AccessCustomPageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessCustomPageService] method
// instead.
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
func (r *AccessCustomPageService) New(ctx context.Context, identifier string, body AccessCustomPageNewParams, opts ...option.RequestOption) (res *ZeroTrustCustomPageWithoutHTML, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a custom page
func (r *AccessCustomPageService) Update(ctx context.Context, identifier string, uuid string, body AccessCustomPageUpdateParams, opts ...option.RequestOption) (res *ZeroTrustCustomPageWithoutHTML, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List custom pages
func (r *AccessCustomPageService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustCustomPageWithoutHTML], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/access/custom_pages", identifier)
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
func (r *AccessCustomPageService) ListAutoPaging(ctx context.Context, identifier string, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustCustomPageWithoutHTML] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, identifier, opts...))
}

// Delete a custom page
func (r *AccessCustomPageService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccessCustomPageDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a custom page and also returns its HTML.
func (r *AccessCustomPageService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZeroTrustCustomPage, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCustomPageGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/custom_pages/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustCustomPage struct {
	// Custom page HTML.
	CustomHTML string `json:"custom_html,required"`
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type ZeroTrustCustomPageType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                  `json:"uid"`
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	JSON      zeroTrustCustomPageJSON `json:"-"`
}

// zeroTrustCustomPageJSON contains the JSON metadata for the struct
// [ZeroTrustCustomPage]
type zeroTrustCustomPageJSON struct {
	CustomHTML  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustCustomPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustCustomPageJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type ZeroTrustCustomPageType string

const (
	ZeroTrustCustomPageTypeIdentityDenied ZeroTrustCustomPageType = "identity_denied"
	ZeroTrustCustomPageTypeForbidden      ZeroTrustCustomPageType = "forbidden"
)

func (r ZeroTrustCustomPageType) IsKnown() bool {
	switch r {
	case ZeroTrustCustomPageTypeIdentityDenied, ZeroTrustCustomPageTypeForbidden:
		return true
	}
	return false
}

type ZeroTrustCustomPageWithoutHTML struct {
	// Custom page name.
	Name string `json:"name,required"`
	// Custom page type.
	Type ZeroTrustCustomPageWithoutHTMLType `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount  int64     `json:"app_count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// UUID
	Uid       string                             `json:"uid"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      zeroTrustCustomPageWithoutHTMLJSON `json:"-"`
}

// zeroTrustCustomPageWithoutHTMLJSON contains the JSON metadata for the struct
// [ZeroTrustCustomPageWithoutHTML]
type zeroTrustCustomPageWithoutHTMLJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	AppCount    apijson.Field
	CreatedAt   apijson.Field
	Uid         apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustCustomPageWithoutHTML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustCustomPageWithoutHTMLJSON) RawJSON() string {
	return r.raw
}

// Custom page type.
type ZeroTrustCustomPageWithoutHTMLType string

const (
	ZeroTrustCustomPageWithoutHTMLTypeIdentityDenied ZeroTrustCustomPageWithoutHTMLType = "identity_denied"
	ZeroTrustCustomPageWithoutHTMLTypeForbidden      ZeroTrustCustomPageWithoutHTMLType = "forbidden"
)

func (r ZeroTrustCustomPageWithoutHTMLType) IsKnown() bool {
	switch r {
	case ZeroTrustCustomPageWithoutHTMLTypeIdentityDenied, ZeroTrustCustomPageWithoutHTMLTypeForbidden:
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
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[AccessCustomPageNewParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r AccessCustomPageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type AccessCustomPageNewParamsType string

const (
	AccessCustomPageNewParamsTypeIdentityDenied AccessCustomPageNewParamsType = "identity_denied"
	AccessCustomPageNewParamsTypeForbidden      AccessCustomPageNewParamsType = "forbidden"
)

func (r AccessCustomPageNewParamsType) IsKnown() bool {
	switch r {
	case AccessCustomPageNewParamsTypeIdentityDenied, AccessCustomPageNewParamsTypeForbidden:
		return true
	}
	return false
}

type AccessCustomPageNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   ZeroTrustCustomPageWithoutHTML `json:"result,required"`
	// Whether the API call was successful
	Success AccessCustomPageNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCustomPageNewResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageNewResponseEnvelope]
type accessCustomPageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	// Custom page HTML.
	CustomHTML param.Field[string] `json:"custom_html,required"`
	// Custom page name.
	Name param.Field[string] `json:"name,required"`
	// Custom page type.
	Type param.Field[AccessCustomPageUpdateParamsType] `json:"type,required"`
	// Number of apps the custom page is assigned to.
	AppCount param.Field[int64] `json:"app_count"`
}

func (r AccessCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom page type.
type AccessCustomPageUpdateParamsType string

const (
	AccessCustomPageUpdateParamsTypeIdentityDenied AccessCustomPageUpdateParamsType = "identity_denied"
	AccessCustomPageUpdateParamsTypeForbidden      AccessCustomPageUpdateParamsType = "forbidden"
)

func (r AccessCustomPageUpdateParamsType) IsKnown() bool {
	switch r {
	case AccessCustomPageUpdateParamsTypeIdentityDenied, AccessCustomPageUpdateParamsTypeForbidden:
		return true
	}
	return false
}

type AccessCustomPageUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   ZeroTrustCustomPageWithoutHTML `json:"result,required"`
	// Whether the API call was successful
	Success AccessCustomPageUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCustomPageUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageUpdateResponseEnvelope]
type accessCustomPageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type AccessCustomPageDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   AccessCustomPageDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success AccessCustomPageDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCustomPageDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageDeleteResponseEnvelope]
type accessCustomPageDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type AccessCustomPageGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZeroTrustCustomPage   `json:"result,required"`
	// Whether the API call was successful
	Success AccessCustomPageGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCustomPageGetResponseEnvelopeJSON    `json:"-"`
}

// accessCustomPageGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCustomPageGetResponseEnvelope]
type accessCustomPageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
