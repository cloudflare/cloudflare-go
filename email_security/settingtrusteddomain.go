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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// SettingTrustedDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingTrustedDomainService] method instead.
type SettingTrustedDomainService struct {
	Options []option.RequestOption
}

// NewSettingTrustedDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingTrustedDomainService(opts ...option.RequestOption) (r *SettingTrustedDomainService) {
	r = &SettingTrustedDomainService{}
	r.Options = opts
	return
}

// Creates a new trusted domain pattern. Use for partner domains or approved
// senders that should bypass recent domain registration and similarity checks.
// Configure whether it prevents recent domain or spoof dispositions.
func (r *SettingTrustedDomainService) New(ctx context.Context, params SettingTrustedDomainNewParams, opts ...option.RequestOption) (res *SettingTrustedDomainNewResponse, err error) {
	var env SettingTrustedDomainNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of trusted domain patterns. Trusted domains prevent
// false positives for recently registered domains and lookalike domain detections.
// Patterns can use regular expressions for flexible matching.
func (r *SettingTrustedDomainService) List(ctx context.Context, params SettingTrustedDomainListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingTrustedDomainListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains", params.AccountID)
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

// Returns a paginated list of trusted domain patterns. Trusted domains prevent
// false positives for recently registered domains and lookalike domain detections.
// Patterns can use regular expressions for flexible matching.
func (r *SettingTrustedDomainService) ListAutoPaging(ctx context.Context, params SettingTrustedDomainListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingTrustedDomainListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes a trusted domain pattern. After deletion, emails from this domain will
// be subject to normal recent domain and similarity checks.
func (r *SettingTrustedDomainService) Delete(ctx context.Context, trustedDomainID string, body SettingTrustedDomainDeleteParams, opts ...option.RequestOption) (res *SettingTrustedDomainDeleteResponse, err error) {
	var env SettingTrustedDomainDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if trustedDomainID == "" {
		err = errors.New("missing required trusted_domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%s", body.AccountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing trusted domain pattern. Only provided fields will be
// modified. Changes take effect for new emails matching the pattern.
func (r *SettingTrustedDomainService) Edit(ctx context.Context, trustedDomainID string, params SettingTrustedDomainEditParams, opts ...option.RequestOption) (res *SettingTrustedDomainEditResponse, err error) {
	var env SettingTrustedDomainEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if trustedDomainID == "" {
		err = errors.New("missing required trusted_domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%s", params.AccountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific trusted domain pattern including its pattern
// value, whether it uses regex matching, and which detection types it affects.
func (r *SettingTrustedDomainService) Get(ctx context.Context, trustedDomainID string, query SettingTrustedDomainGetParams, opts ...option.RequestOption) (res *SettingTrustedDomainGetResponse, err error) {
	var env SettingTrustedDomainGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if trustedDomainID == "" {
		err = errors.New("missing required trusted_domain_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/trusted_domains/%s", query.AccountID, trustedDomainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A trusted email domain
type SettingTrustedDomainNewResponse struct {
	// Trusted domain identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent"`
	IsRegex  bool `json:"is_regex"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool `json:"is_similarity"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                           `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                           `json:"modified_at" format:"date-time"`
	Pattern      string                              `json:"pattern"`
	JSON         settingTrustedDomainNewResponseJSON `json:"-"`
}

// settingTrustedDomainNewResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainNewResponse]
type settingTrustedDomainNewResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseJSON) RawJSON() string {
	return r.raw
}

// A trusted email domain
type SettingTrustedDomainListResponse struct {
	// Trusted domain identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent"`
	IsRegex  bool `json:"is_regex"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool `json:"is_similarity"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                            `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                            `json:"modified_at" format:"date-time"`
	Pattern      string                               `json:"pattern"`
	JSON         settingTrustedDomainListResponseJSON `json:"-"`
}

// settingTrustedDomainListResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainListResponse]
type settingTrustedDomainListResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainDeleteResponse struct {
	// Trusted domain identifier
	ID   string                                 `json:"id" api:"required" format:"uuid"`
	JSON settingTrustedDomainDeleteResponseJSON `json:"-"`
}

// settingTrustedDomainDeleteResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainDeleteResponse]
type settingTrustedDomainDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A trusted email domain
type SettingTrustedDomainEditResponse struct {
	// Trusted domain identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent"`
	IsRegex  bool `json:"is_regex"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool `json:"is_similarity"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                            `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                            `json:"modified_at" format:"date-time"`
	Pattern      string                               `json:"pattern"`
	JSON         settingTrustedDomainEditResponseJSON `json:"-"`
}

// settingTrustedDomainEditResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainEditResponse]
type settingTrustedDomainEditResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseJSON) RawJSON() string {
	return r.raw
}

// A trusted email domain
type SettingTrustedDomainGetResponse struct {
	// Trusted domain identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent bool `json:"is_recent"`
	IsRegex  bool `json:"is_regex"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity bool `json:"is_similarity"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                           `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                           `json:"modified_at" format:"date-time"`
	Pattern      string                              `json:"pattern"`
	JSON         settingTrustedDomainGetResponseJSON `json:"-"`
}

// settingTrustedDomainGetResponseJSON contains the JSON metadata for the struct
// [SettingTrustedDomainGetResponse]
type settingTrustedDomainGetResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRecent     apijson.Field
	IsRegex      apijson.Field
	IsSimilarity apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent param.Field[bool] `json:"is_recent" api:"required"`
	IsRegex  param.Field[bool] `json:"is_regex" api:"required"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity param.Field[bool]   `json:"is_similarity" api:"required"`
	Pattern      param.Field[string] `json:"pattern" api:"required"`
	Comments     param.Field[string] `json:"comments"`
}

func (r SettingTrustedDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingTrustedDomainNewResponseEnvelope struct {
	Errors   []SettingTrustedDomainNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingTrustedDomainNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingTrustedDomainNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A trusted email domain
	Result SettingTrustedDomainNewResponse             `json:"result"`
	JSON   settingTrustedDomainNewResponseEnvelopeJSON `json:"-"`
}

// settingTrustedDomainNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTrustedDomainNewResponseEnvelope]
type settingTrustedDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainNewResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingTrustedDomainNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingTrustedDomainNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingTrustedDomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingTrustedDomainNewResponseEnvelopeErrors]
type settingTrustedDomainNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingTrustedDomainNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingTrustedDomainNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingTrustedDomainNewResponseEnvelopeErrorsSource]
type settingTrustedDomainNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainNewResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           SettingTrustedDomainNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingTrustedDomainNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingTrustedDomainNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingTrustedDomainNewResponseEnvelopeMessages]
type settingTrustedDomainNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    settingTrustedDomainNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingTrustedDomainNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingTrustedDomainNewResponseEnvelopeMessagesSource]
type settingTrustedDomainNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingTrustedDomainNewResponseEnvelopeSuccess bool

const (
	SettingTrustedDomainNewResponseEnvelopeSuccessTrue SettingTrustedDomainNewResponseEnvelopeSuccess = true
)

func (r SettingTrustedDomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTrustedDomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingTrustedDomainListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The sorting direction.
	Direction param.Field[SettingTrustedDomainListParamsDirection] `query:"direction"`
	// Filter to show only recently registered domains that are trusted to prevent
	// triggering Suspicious or Malicious dispositions.
	IsRecent param.Field[bool] `query:"is_recent"`
	// Filter to show only proximity domains (partner or approved domains with similar
	// spelling to connected domains) that prevent Spoof dispositions.
	IsSimilarity param.Field[bool] `query:"is_similarity"`
	// Field to sort by.
	Order param.Field[SettingTrustedDomainListParamsOrder] `query:"order"`
	// Current page within paginated list of results.
	Page    param.Field[int64]  `query:"page"`
	Pattern param.Field[string] `query:"pattern"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
	// Search term for filtering records. Behavior may change.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingTrustedDomainListParams]'s query parameters as
// `url.Values`.
func (r SettingTrustedDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingTrustedDomainListParamsDirection string

const (
	SettingTrustedDomainListParamsDirectionAsc  SettingTrustedDomainListParamsDirection = "asc"
	SettingTrustedDomainListParamsDirectionDesc SettingTrustedDomainListParamsDirection = "desc"
)

func (r SettingTrustedDomainListParamsDirection) IsKnown() bool {
	switch r {
	case SettingTrustedDomainListParamsDirectionAsc, SettingTrustedDomainListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by.
type SettingTrustedDomainListParamsOrder string

const (
	SettingTrustedDomainListParamsOrderPattern   SettingTrustedDomainListParamsOrder = "pattern"
	SettingTrustedDomainListParamsOrderCreatedAt SettingTrustedDomainListParamsOrder = "created_at"
)

func (r SettingTrustedDomainListParamsOrder) IsKnown() bool {
	switch r {
	case SettingTrustedDomainListParamsOrderPattern, SettingTrustedDomainListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingTrustedDomainDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingTrustedDomainDeleteResponseEnvelope struct {
	Errors   []SettingTrustedDomainDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingTrustedDomainDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingTrustedDomainDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingTrustedDomainDeleteResponse                `json:"result"`
	JSON    settingTrustedDomainDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingTrustedDomainDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingTrustedDomainDeleteResponseEnvelope]
type settingTrustedDomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingTrustedDomainDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingTrustedDomainDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingTrustedDomainDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingTrustedDomainDeleteResponseEnvelopeErrors]
type settingTrustedDomainDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingTrustedDomainDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingTrustedDomainDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingTrustedDomainDeleteResponseEnvelopeErrorsSource]
type settingTrustedDomainDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           SettingTrustedDomainDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingTrustedDomainDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingTrustedDomainDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingTrustedDomainDeleteResponseEnvelopeMessages]
type settingTrustedDomainDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    settingTrustedDomainDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingTrustedDomainDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [SettingTrustedDomainDeleteResponseEnvelopeMessagesSource]
type settingTrustedDomainDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingTrustedDomainDeleteResponseEnvelopeSuccess bool

const (
	SettingTrustedDomainDeleteResponseEnvelopeSuccessTrue SettingTrustedDomainDeleteResponseEnvelopeSuccess = true
)

func (r SettingTrustedDomainDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTrustedDomainDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingTrustedDomainEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Comments  param.Field[string] `json:"comments"`
	// Select to prevent recently registered domains from triggering a Suspicious or
	// Malicious disposition.
	IsRecent param.Field[bool] `json:"is_recent"`
	IsRegex  param.Field[bool] `json:"is_regex"`
	// Select for partner or other approved domains that have similar spelling to your
	// connected domains. Prevents listed domains from triggering a Spoof disposition.
	IsSimilarity param.Field[bool]   `json:"is_similarity"`
	Pattern      param.Field[string] `json:"pattern"`
}

func (r SettingTrustedDomainEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingTrustedDomainEditResponseEnvelope struct {
	Errors   []SettingTrustedDomainEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingTrustedDomainEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingTrustedDomainEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// A trusted email domain
	Result SettingTrustedDomainEditResponse             `json:"result"`
	JSON   settingTrustedDomainEditResponseEnvelopeJSON `json:"-"`
}

// settingTrustedDomainEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTrustedDomainEditResponseEnvelope]
type settingTrustedDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainEditResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           SettingTrustedDomainEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingTrustedDomainEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingTrustedDomainEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingTrustedDomainEditResponseEnvelopeErrors]
type settingTrustedDomainEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    settingTrustedDomainEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingTrustedDomainEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingTrustedDomainEditResponseEnvelopeErrorsSource]
type settingTrustedDomainEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainEditResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingTrustedDomainEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingTrustedDomainEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingTrustedDomainEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingTrustedDomainEditResponseEnvelopeMessages]
type settingTrustedDomainEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingTrustedDomainEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingTrustedDomainEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingTrustedDomainEditResponseEnvelopeMessagesSource]
type settingTrustedDomainEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingTrustedDomainEditResponseEnvelopeSuccess bool

const (
	SettingTrustedDomainEditResponseEnvelopeSuccessTrue SettingTrustedDomainEditResponseEnvelopeSuccess = true
)

func (r SettingTrustedDomainEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTrustedDomainEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingTrustedDomainGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingTrustedDomainGetResponseEnvelope struct {
	Errors   []SettingTrustedDomainGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingTrustedDomainGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingTrustedDomainGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A trusted email domain
	Result SettingTrustedDomainGetResponse             `json:"result"`
	JSON   settingTrustedDomainGetResponseEnvelopeJSON `json:"-"`
}

// settingTrustedDomainGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingTrustedDomainGetResponseEnvelope]
type settingTrustedDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingTrustedDomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingTrustedDomainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingTrustedDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingTrustedDomainGetResponseEnvelopeErrors]
type settingTrustedDomainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingTrustedDomainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingTrustedDomainGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingTrustedDomainGetResponseEnvelopeErrorsSource]
type settingTrustedDomainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           SettingTrustedDomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingTrustedDomainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingTrustedDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingTrustedDomainGetResponseEnvelopeMessages]
type settingTrustedDomainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTrustedDomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    settingTrustedDomainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingTrustedDomainGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingTrustedDomainGetResponseEnvelopeMessagesSource]
type settingTrustedDomainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTrustedDomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTrustedDomainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingTrustedDomainGetResponseEnvelopeSuccess bool

const (
	SettingTrustedDomainGetResponseEnvelopeSuccessTrue SettingTrustedDomainGetResponseEnvelopeSuccess = true
)

func (r SettingTrustedDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTrustedDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
