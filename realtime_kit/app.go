// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

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
)

// AppService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Fetch all apps for your account
func (r *AppService) Get(ctx context.Context, params AppGetParams, opts ...option.RequestOption) (res *AppGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/apps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Fetch details for an app in your account.
func (r *AppService) GetAppByID(ctx context.Context, appID string, query AppGetAppByIDParams, opts ...option.RequestOption) (res *AppGetAppByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/apps/%s", query.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create new app for your account
func (r *AppService) Post(ctx context.Context, params AppPostParams, opts ...option.RequestOption) (res *AppPostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/apps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update the name of an app in your account.
func (r *AppService) UpdateAppByID(ctx context.Context, appID string, params AppUpdateAppByIDParams, opts ...option.RequestOption) (res *AppUpdateAppByIDResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/apps/%s", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type AppGetResponse struct {
	Data    []AppGetResponseData `json:"data"`
	Paging  AppGetResponsePaging `json:"paging"`
	Success bool                 `json:"success"`
	JSON    appGetResponseJSON   `json:"-"`
}

// appGetResponseJSON contains the JSON metadata for the struct [AppGetResponse]
type appGetResponseJSON struct {
	Data        apijson.Field
	Paging      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseData struct {
	ID        string                 `json:"id" format:"uuid"`
	CreatedAt time.Time              `json:"created_at" format:"date-time"`
	Name      string                 `json:"name"`
	JSON      appGetResponseDataJSON `json:"-"`
}

// appGetResponseDataJSON contains the JSON metadata for the struct
// [AppGetResponseData]
type appGetResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type AppGetResponsePaging struct {
	EndOffset   float64                  `json:"end_offset"`
	StartOffset float64                  `json:"start_offset"`
	TotalCount  float64                  `json:"total_count"`
	JSON        appGetResponsePagingJSON `json:"-"`
}

// appGetResponsePagingJSON contains the JSON metadata for the struct
// [AppGetResponsePaging]
type appGetResponsePagingJSON struct {
	EndOffset   apijson.Field
	StartOffset apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponsePaging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponsePagingJSON) RawJSON() string {
	return r.raw
}

type AppGetAppByIDResponse struct {
	Data    AppGetAppByIDResponseData `json:"data"`
	Success bool                      `json:"success"`
	JSON    appGetAppByIDResponseJSON `json:"-"`
}

// appGetAppByIDResponseJSON contains the JSON metadata for the struct
// [AppGetAppByIDResponse]
type appGetAppByIDResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetAppByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetAppByIDResponseJSON) RawJSON() string {
	return r.raw
}

type AppGetAppByIDResponseData struct {
	ID        string                        `json:"id" format:"uuid"`
	CreatedAt time.Time                     `json:"created_at" format:"date-time"`
	Name      string                        `json:"name" api:"nullable"`
	JSON      appGetAppByIDResponseDataJSON `json:"-"`
}

// appGetAppByIDResponseDataJSON contains the JSON metadata for the struct
// [AppGetAppByIDResponseData]
type appGetAppByIDResponseDataJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetAppByIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetAppByIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type AppPostResponse struct {
	Data    AppPostResponseData `json:"data"`
	Success bool                `json:"success"`
	JSON    appPostResponseJSON `json:"-"`
}

// appPostResponseJSON contains the JSON metadata for the struct [AppPostResponse]
type appPostResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppPostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appPostResponseJSON) RawJSON() string {
	return r.raw
}

type AppPostResponseData struct {
	App  AppPostResponseDataApp  `json:"app"`
	JSON appPostResponseDataJSON `json:"-"`
}

// appPostResponseDataJSON contains the JSON metadata for the struct
// [AppPostResponseData]
type appPostResponseDataJSON struct {
	App         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppPostResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appPostResponseDataJSON) RawJSON() string {
	return r.raw
}

type AppPostResponseDataApp struct {
	ID        string                     `json:"id" format:"uuid"`
	CreatedAt time.Time                  `json:"created_at" format:"date-time"`
	Name      string                     `json:"name"`
	JSON      appPostResponseDataAppJSON `json:"-"`
}

// appPostResponseDataAppJSON contains the JSON metadata for the struct
// [AppPostResponseDataApp]
type appPostResponseDataAppJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppPostResponseDataApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appPostResponseDataAppJSON) RawJSON() string {
	return r.raw
}

type AppUpdateAppByIDResponse struct {
	Data    AppUpdateAppByIDResponseData    `json:"data" api:"required"`
	Success AppUpdateAppByIDResponseSuccess `json:"success" api:"required"`
	JSON    appUpdateAppByIDResponseJSON    `json:"-"`
}

// appUpdateAppByIDResponseJSON contains the JSON metadata for the struct
// [AppUpdateAppByIDResponse]
type appUpdateAppByIDResponseJSON struct {
	Data        apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateAppByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateAppByIDResponseJSON) RawJSON() string {
	return r.raw
}

type AppUpdateAppByIDResponseData struct {
	ID           string                           `json:"id" format:"uuid"`
	CfAccountTag string                           `json:"cf_account_tag" api:"nullable"`
	CreatedAt    time.Time                        `json:"created_at" format:"date-time"`
	IsOrgOnKV    bool                             `json:"is_org_on_kv"`
	Name         string                           `json:"name" api:"nullable"`
	JSON         appUpdateAppByIDResponseDataJSON `json:"-"`
}

// appUpdateAppByIDResponseDataJSON contains the JSON metadata for the struct
// [AppUpdateAppByIDResponseData]
type appUpdateAppByIDResponseDataJSON struct {
	ID           apijson.Field
	CfAccountTag apijson.Field
	CreatedAt    apijson.Field
	IsOrgOnKV    apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppUpdateAppByIDResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateAppByIDResponseDataJSON) RawJSON() string {
	return r.raw
}

type AppUpdateAppByIDResponseSuccess bool

const (
	AppUpdateAppByIDResponseSuccessTrue AppUpdateAppByIDResponseSuccess = true
)

func (r AppUpdateAppByIDResponseSuccess) IsKnown() bool {
	switch r {
	case AppUpdateAppByIDResponseSuccessTrue:
		return true
	}
	return false
}

type AppGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The page number from which you want your page search results to be displayed.
	PageNo param.Field[int64] `query:"page_no"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Search string that matches apps by name.
	Search param.Field[string] `query:"search"`
	// Sort order for apps by creation time.
	SortOrder param.Field[AppGetParamsSortOrder] `query:"sort_order"`
}

// URLQuery serializes [AppGetParams]'s query parameters as `url.Values`.
func (r AppGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort order for apps by creation time.
type AppGetParamsSortOrder string

const (
	AppGetParamsSortOrderAsc  AppGetParamsSortOrder = "ASC"
	AppGetParamsSortOrderDesc AppGetParamsSortOrder = "DESC"
)

func (r AppGetParamsSortOrder) IsKnown() bool {
	switch r {
	case AppGetParamsSortOrderAsc, AppGetParamsSortOrderDesc:
		return true
	}
	return false
}

type AppGetAppByIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AppPostParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Name      param.Field[string] `json:"name" api:"required"`
}

func (r AppPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppUpdateAppByIDParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The app name. Leading and trailing whitespace is removed, and the name must
	// contain at least one non-whitespace character.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r AppUpdateAppByIDParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
