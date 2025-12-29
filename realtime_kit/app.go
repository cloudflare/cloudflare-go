// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package realtime_kit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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
func (r *AppService) Get(ctx context.Context, query AppGetParams, opts ...option.RequestOption) (res *AppGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/apps", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create new app for your account
func (r *AppService) Post(ctx context.Context, params AppPostParams, opts ...option.RequestOption) (res *AppPostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/realtime/kit/apps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AppGetResponse struct {
	Data    []AppGetResponseData `json:"data"`
	Success bool                 `json:"success"`
	JSON    appGetResponseJSON   `json:"-"`
}

// appGetResponseJSON contains the JSON metadata for the struct [AppGetResponse]
type appGetResponseJSON struct {
	Data        apijson.Field
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
	ID        string                 `json:"id"`
	CreatedAt string                 `json:"created_at"`
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
	ID        string                     `json:"id"`
	CreatedAt string                     `json:"created_at"`
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

type AppGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AppPostParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r AppPostParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
