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

// SettingZoneViewService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingZoneViewService] method instead.
type SettingZoneViewService struct {
	Options []option.RequestOption
}

// NewSettingZoneViewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingZoneViewService(opts ...option.RequestOption) (r *SettingZoneViewService) {
	r = &SettingZoneViewService{}
	r.Options = opts
	return
}

// Create Internal DNS View for an account
func (r *SettingZoneViewService) New(ctx context.Context, params SettingZoneViewNewParams, opts ...option.RequestOption) (res *SettingZoneViewNewResponse, err error) {
	var env SettingZoneViewNewResponseEnvelope
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
func (r *SettingZoneViewService) List(ctx context.Context, params SettingZoneViewListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingZoneViewListResponse], err error) {
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
func (r *SettingZoneViewService) ListAutoPaging(ctx context.Context, params SettingZoneViewListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingZoneViewListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an existing Internal DNS View
func (r *SettingZoneViewService) Delete(ctx context.Context, viewID string, body SettingZoneViewDeleteParams, opts ...option.RequestOption) (res *SettingZoneViewDeleteResponse, err error) {
	var env SettingZoneViewDeleteResponseEnvelope
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
func (r *SettingZoneViewService) Edit(ctx context.Context, viewID string, params SettingZoneViewEditParams, opts ...option.RequestOption) (res *SettingZoneViewEditResponse, err error) {
	var env SettingZoneViewEditResponseEnvelope
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
func (r *SettingZoneViewService) Get(ctx context.Context, viewID string, query SettingZoneViewGetParams, opts ...option.RequestOption) (res *SettingZoneViewGetResponse, err error) {
	var env SettingZoneViewGetResponseEnvelope
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

type SettingZoneViewNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                       `json:"zones,required"`
	JSON  settingZoneViewNewResponseJSON `json:"-"`
}

// settingZoneViewNewResponseJSON contains the JSON metadata for the struct
// [SettingZoneViewNewResponse]
type settingZoneViewNewResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingZoneViewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewNewResponseJSON) RawJSON() string {
	return r.raw
}

type SettingZoneViewListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                        `json:"zones,required"`
	JSON  settingZoneViewListResponseJSON `json:"-"`
}

// settingZoneViewListResponseJSON contains the JSON metadata for the struct
// [SettingZoneViewListResponse]
type settingZoneViewListResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingZoneViewListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingZoneViewDeleteResponse struct {
	// Identifier
	ID   string                            `json:"id"`
	JSON settingZoneViewDeleteResponseJSON `json:"-"`
}

// settingZoneViewDeleteResponseJSON contains the JSON metadata for the struct
// [SettingZoneViewDeleteResponse]
type settingZoneViewDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneViewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingZoneViewEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                        `json:"zones,required"`
	JSON  settingZoneViewEditResponseJSON `json:"-"`
}

// settingZoneViewEditResponseJSON contains the JSON metadata for the struct
// [SettingZoneViewEditResponse]
type settingZoneViewEditResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingZoneViewEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingZoneViewGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string                       `json:"zones,required"`
	JSON  settingZoneViewGetResponseJSON `json:"-"`
}

// settingZoneViewGetResponseJSON contains the JSON metadata for the struct
// [SettingZoneViewGetResponse]
type settingZoneViewGetResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingZoneViewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingZoneViewNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the view.
	Name param.Field[string] `json:"name,required"`
	// The list of zones linked to this view.
	Zones param.Field[[]string] `json:"zones,required"`
}

func (r SettingZoneViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingZoneViewNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingZoneViewNewResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingZoneViewNewResponse                `json:"result"`
	JSON    settingZoneViewNewResponseEnvelopeJSON    `json:"-"`
}

// settingZoneViewNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZoneViewNewResponseEnvelope]
type settingZoneViewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneViewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingZoneViewNewResponseEnvelopeSuccess bool

const (
	SettingZoneViewNewResponseEnvelopeSuccessTrue SettingZoneViewNewResponseEnvelopeSuccess = true
)

func (r SettingZoneViewNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingZoneViewNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingZoneViewListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order DNS views in.
	Direction param.Field[SettingZoneViewListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead.
	Match param.Field[SettingZoneViewListParamsMatch] `query:"match"`
	Name  param.Field[SettingZoneViewListParamsName]  `query:"name"`
	// Field to order DNS views by.
	Order param.Field[SettingZoneViewListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of DNS views per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A zone ID that exists in the zones list for the view.
	ZoneID param.Field[string] `query:"zone_id"`
	// A zone name that exists in the zones list for the view.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [SettingZoneViewListParams]'s query parameters as
// `url.Values`.
func (r SettingZoneViewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order DNS views in.
type SettingZoneViewListParamsDirection string

const (
	SettingZoneViewListParamsDirectionAsc  SettingZoneViewListParamsDirection = "asc"
	SettingZoneViewListParamsDirectionDesc SettingZoneViewListParamsDirection = "desc"
)

func (r SettingZoneViewListParamsDirection) IsKnown() bool {
	switch r {
	case SettingZoneViewListParamsDirectionAsc, SettingZoneViewListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead.
type SettingZoneViewListParamsMatch string

const (
	SettingZoneViewListParamsMatchAny SettingZoneViewListParamsMatch = "any"
	SettingZoneViewListParamsMatchAll SettingZoneViewListParamsMatch = "all"
)

func (r SettingZoneViewListParamsMatch) IsKnown() bool {
	switch r {
	case SettingZoneViewListParamsMatchAny, SettingZoneViewListParamsMatchAll:
		return true
	}
	return false
}

type SettingZoneViewListParamsName struct {
	// Substring of the DNS view name.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS view name.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS view name.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS view name.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [SettingZoneViewListParamsName]'s query parameters as
// `url.Values`.
func (r SettingZoneViewListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order DNS views by.
type SettingZoneViewListParamsOrder string

const (
	SettingZoneViewListParamsOrderName       SettingZoneViewListParamsOrder = "name"
	SettingZoneViewListParamsOrderCreatedOn  SettingZoneViewListParamsOrder = "created_on"
	SettingZoneViewListParamsOrderModifiedOn SettingZoneViewListParamsOrder = "modified_on"
)

func (r SettingZoneViewListParamsOrder) IsKnown() bool {
	switch r {
	case SettingZoneViewListParamsOrderName, SettingZoneViewListParamsOrderCreatedOn, SettingZoneViewListParamsOrderModifiedOn:
		return true
	}
	return false
}

type SettingZoneViewDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingZoneViewDeleteResponseEnvelope struct {
	Result SettingZoneViewDeleteResponse             `json:"result"`
	JSON   settingZoneViewDeleteResponseEnvelopeJSON `json:"-"`
}

// settingZoneViewDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingZoneViewDeleteResponseEnvelope]
type settingZoneViewDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneViewDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZoneViewEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the view.
	Name param.Field[string] `json:"name"`
	// The list of zones linked to this view.
	Zones param.Field[[]string] `json:"zones"`
}

func (r SettingZoneViewEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingZoneViewEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingZoneViewEditResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingZoneViewEditResponse                `json:"result"`
	JSON    settingZoneViewEditResponseEnvelopeJSON    `json:"-"`
}

// settingZoneViewEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingZoneViewEditResponseEnvelope]
type settingZoneViewEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneViewEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingZoneViewEditResponseEnvelopeSuccess bool

const (
	SettingZoneViewEditResponseEnvelopeSuccessTrue SettingZoneViewEditResponseEnvelopeSuccess = true
)

func (r SettingZoneViewEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingZoneViewEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingZoneViewGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingZoneViewGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SettingZoneViewGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingZoneViewGetResponse                `json:"result"`
	JSON    settingZoneViewGetResponseEnvelopeJSON    `json:"-"`
}

// settingZoneViewGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZoneViewGetResponseEnvelope]
type settingZoneViewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneViewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneViewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SettingZoneViewGetResponseEnvelopeSuccess bool

const (
	SettingZoneViewGetResponseEnvelopeSuccessTrue SettingZoneViewGetResponseEnvelopeSuccess = true
)

func (r SettingZoneViewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingZoneViewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
