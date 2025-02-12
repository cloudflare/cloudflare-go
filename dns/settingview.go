// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// SettingViewService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingViewService] method instead.
type SettingViewService struct {
	Options []option.RequestOption
}

// NewSettingViewService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingViewService(opts ...option.RequestOption) (r *SettingViewService) {
	r = &SettingViewService{}
	r.Options = opts
	return
}

// Create Internal DNS View for an account
func (r *SettingViewService) New(ctx context.Context, params SettingViewNewParams, opts ...option.RequestOption) (res *SettingViewNewResponse, err error) {
	var env SettingViewNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List DNS Internal Views for an Account
func (r *SettingViewService) List(ctx context.Context, params SettingViewListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingViewListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views", params.AccountID)
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

// List DNS Internal Views for an Account
func (r *SettingViewService) ListAutoPaging(ctx context.Context, params SettingViewListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingViewListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an existing Internal DNS View
func (r *SettingViewService) Delete(ctx context.Context, viewID string, body SettingViewDeleteParams, opts ...option.RequestOption) (res *SettingViewDeleteResponse, err error) {
	var env SettingViewDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", body.AccountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing Internal DNS View
func (r *SettingViewService) Edit(ctx context.Context, viewID string, params SettingViewEditParams, opts ...option.RequestOption) (res *SettingViewEditResponse, err error) {
	var env SettingViewEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", params.AccountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get DNS Internal View
func (r *SettingViewService) Get(ctx context.Context, viewID string, query SettingViewGetParams, opts ...option.RequestOption) (res *SettingViewGetResponse, err error) {
	var env SettingViewGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", query.AccountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingViewNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                   `json:"zones,required"`
	JSON  settingViewNewResponseJSON `json:"-"`
}

// settingViewNewResponseJSON contains the JSON metadata for the struct
// [SettingViewNewResponse]
type settingViewNewResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingViewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewNewResponseJSON) RawJSON() string {
	return r.raw
}

type SettingViewListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                    `json:"zones,required"`
	JSON  settingViewListResponseJSON `json:"-"`
}

// settingViewListResponseJSON contains the JSON metadata for the struct
// [SettingViewListResponse]
type settingViewListResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingViewListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingViewDeleteResponse struct {
	// Identifier
	ID   string                        `json:"id"`
	JSON settingViewDeleteResponseJSON `json:"-"`
}

// settingViewDeleteResponseJSON contains the JSON metadata for the struct
// [SettingViewDeleteResponse]
type settingViewDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingViewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingViewEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                    `json:"zones,required"`
	JSON  settingViewEditResponseJSON `json:"-"`
}

// settingViewEditResponseJSON contains the JSON metadata for the struct
// [SettingViewEditResponse]
type settingViewEditResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingViewEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingViewGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                   `json:"zones,required"`
	JSON  settingViewGetResponseJSON `json:"-"`
}

// settingViewGetResponseJSON contains the JSON metadata for the struct
// [SettingViewGetResponse]
type settingViewGetResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingViewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingViewNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the view.
	Name param.Field[string] `json:"name,required"`
	// The list of zones linked to this view.
	Zones param.Field[[]string] `json:"zones,required"`
}

func (r SettingViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingViewNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingViewNewResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingViewNewResponse                `json:"result"`
	JSON    settingViewNewResponseEnvelopeJSON    `json:"-"`
}

// settingViewNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingViewNewResponseEnvelope]
type settingViewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingViewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingViewNewResponseEnvelopeSuccess bool

const (
	SettingViewNewResponseEnvelopeSuccessTrue SettingViewNewResponseEnvelopeSuccess = true
)

func (r SettingViewNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingViewNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingViewListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order DNS views in.
	Direction param.Field[SettingViewListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead.
	Match param.Field[SettingViewListParamsMatch] `query:"match"`
	Name  param.Field[SettingViewListParamsName]  `query:"name"`
	// Field to order DNS views by.
	Order param.Field[SettingViewListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of DNS views per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A zone ID that exists in the zones list for the view.
	ZoneID param.Field[string] `query:"zone_id"`
	// A zone name that exists in the zones list for the view.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [SettingViewListParams]'s query parameters as `url.Values`.
func (r SettingViewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order DNS views in.
type SettingViewListParamsDirection string

const (
	SettingViewListParamsDirectionAsc  SettingViewListParamsDirection = "asc"
	SettingViewListParamsDirectionDesc SettingViewListParamsDirection = "desc"
)

func (r SettingViewListParamsDirection) IsKnown() bool {
	switch r {
	case SettingViewListParamsDirectionAsc, SettingViewListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead.
type SettingViewListParamsMatch string

const (
	SettingViewListParamsMatchAny SettingViewListParamsMatch = "any"
	SettingViewListParamsMatchAll SettingViewListParamsMatch = "all"
)

func (r SettingViewListParamsMatch) IsKnown() bool {
	switch r {
	case SettingViewListParamsMatchAny, SettingViewListParamsMatchAll:
		return true
	}
	return false
}

type SettingViewListParamsName struct {
	// Substring of the DNS view name.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS view name.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS view name.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS view name.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [SettingViewListParamsName]'s query parameters as
// `url.Values`.
func (r SettingViewListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order DNS views by.
type SettingViewListParamsOrder string

const (
	SettingViewListParamsOrderName       SettingViewListParamsOrder = "name"
	SettingViewListParamsOrderCreatedOn  SettingViewListParamsOrder = "created_on"
	SettingViewListParamsOrderModifiedOn SettingViewListParamsOrder = "modified_on"
)

func (r SettingViewListParamsOrder) IsKnown() bool {
	switch r {
	case SettingViewListParamsOrderName, SettingViewListParamsOrderCreatedOn, SettingViewListParamsOrderModifiedOn:
		return true
	}
	return false
}

type SettingViewDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingViewDeleteResponseEnvelope struct {
	Result SettingViewDeleteResponse             `json:"result"`
	JSON   settingViewDeleteResponseEnvelopeJSON `json:"-"`
}

// settingViewDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingViewDeleteResponseEnvelope]
type settingViewDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingViewDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingViewEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the view.
	Name param.Field[string] `json:"name"`
	// The list of zones linked to this view.
	Zones param.Field[[]string] `json:"zones"`
}

func (r SettingViewEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingViewEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingViewEditResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingViewEditResponse                `json:"result"`
	JSON    settingViewEditResponseEnvelopeJSON    `json:"-"`
}

// settingViewEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingViewEditResponseEnvelope]
type settingViewEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingViewEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingViewEditResponseEnvelopeSuccess bool

const (
	SettingViewEditResponseEnvelopeSuccessTrue SettingViewEditResponseEnvelopeSuccess = true
)

func (r SettingViewEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingViewEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingViewGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingViewGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingViewGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingViewGetResponse                `json:"result"`
	JSON    settingViewGetResponseEnvelopeJSON    `json:"-"`
}

// settingViewGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingViewGetResponseEnvelope]
type settingViewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingViewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingViewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingViewGetResponseEnvelopeSuccess bool

const (
	SettingViewGetResponseEnvelopeSuccessTrue SettingViewGetResponseEnvelopeSuccess = true
)

func (r SettingViewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingViewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
