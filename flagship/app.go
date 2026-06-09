// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AppService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options  []option.RequestOption
	Flags    *AppFlagService
	Evaluate *AppEvaluateService
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	r.Flags = NewAppFlagService(opts...)
	r.Evaluate = NewAppEvaluateService(opts...)
	return
}

// Creates an app. The returned `id` is used in all subsequent flag, changelog, and
// evaluation requests.
func (r *AppService) New(ctx context.Context, params AppNewParams, opts ...option.RequestOption) (res *AppNewResponse, err error) {
	var env AppNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an app. Only `name` is mutable.
func (r *AppService) Update(ctx context.Context, appID string, params AppUpdateParams, opts ...option.RequestOption) (res *AppUpdateResponse, err error) {
	var env AppUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all apps in the account. Returns identity and audit fields only — flag
// definitions are not included.
func (r *AppService) List(ctx context.Context, query AppListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AppListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps", query.AccountID)
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

// Lists all apps in the account. Returns identity and audit fields only — flag
// definitions are not included.
func (r *AppService) ListAutoPaging(ctx context.Context, query AppListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AppListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an app and all its flags and changelog history. Returns 409 if any
// Worker still references this app via a Flagship binding.
func (r *AppService) Delete(ctx context.Context, appID string, body AppDeleteParams, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	var env AppDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s", body.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns an app's name and audit fields. Flag definitions are not included.
func (r *AppService) Get(ctx context.Context, appID string, query AppGetParams, opts ...option.RequestOption) (res *AppGetResponse, err error) {
	var env AppGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s", query.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AppNewResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"created_at" api:"required"`
	Name      string `json:"name" api:"required"`
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the actor who last modified the app, or `edge-gateway` for
	// gateway-authenticated changes.
	UpdatedBy string             `json:"updated_by" api:"required"`
	JSON      appNewResponseJSON `json:"-"`
}

// appNewResponseJSON contains the JSON metadata for the struct [AppNewResponse]
type appNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	UpdatedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"created_at" api:"required"`
	Name      string `json:"name" api:"required"`
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the actor who last modified the app, or `edge-gateway` for
	// gateway-authenticated changes.
	UpdatedBy string                `json:"updated_by" api:"required"`
	JSON      appUpdateResponseJSON `json:"-"`
}

// appUpdateResponseJSON contains the JSON metadata for the struct
// [AppUpdateResponse]
type appUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	UpdatedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AppListResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"created_at" api:"required"`
	Name      string `json:"name" api:"required"`
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the actor who last modified the app, or `edge-gateway` for
	// gateway-authenticated changes.
	UpdatedBy string              `json:"updated_by" api:"required"`
	JSON      appListResponseJSON `json:"-"`
}

// appListResponseJSON contains the JSON metadata for the struct [AppListResponse]
type appListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	UpdatedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appListResponseJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponse struct {
	ID   string                `json:"id" api:"required"`
	JSON appDeleteResponseJSON `json:"-"`
}

// appDeleteResponseJSON contains the JSON metadata for the struct
// [AppDeleteResponse]
type appDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AppGetResponse struct {
	ID        string `json:"id" api:"required"`
	CreatedAt string `json:"created_at" api:"required"`
	Name      string `json:"name" api:"required"`
	UpdatedAt string `json:"updated_at" api:"required"`
	// Email of the actor who last modified the app, or `edge-gateway` for
	// gateway-authenticated changes.
	UpdatedBy string             `json:"updated_by" api:"required"`
	JSON      appGetResponseJSON `json:"-"`
}

// appGetResponseJSON contains the JSON metadata for the struct [AppGetResponse]
type appGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	UpdatedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseJSON) RawJSON() string {
	return r.raw
}

type AppNewParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Name      param.Field[string] `json:"name" api:"required"`
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppNewResponseEnvelope struct {
	Errors   []AppNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppNewResponse                   `json:"result" api:"required"`
	Success  bool                             `json:"success" api:"required"`
	JSON     appNewResponseEnvelopeJSON       `json:"-"`
}

// appNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelope]
type appNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeErrors struct {
	Message string                           `json:"message" api:"required"`
	JSON    appNewResponseEnvelopeErrorsJSON `json:"-"`
}

// appNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelopeErrors]
type appNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeMessages struct {
	Message string                             `json:"message" api:"required"`
	JSON    appNewResponseEnvelopeMessagesJSON `json:"-"`
}

// appNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelopeMessages]
type appNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppUpdateParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Name      param.Field[string] `json:"name"`
}

func (r AppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppUpdateResponseEnvelope struct {
	Errors   []AppUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppUpdateResponse                   `json:"result" api:"required"`
	Success  bool                                `json:"success" api:"required"`
	JSON     appUpdateResponseEnvelopeJSON       `json:"-"`
}

// appUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelope]
type appUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeErrors struct {
	Message string                              `json:"message" api:"required"`
	JSON    appUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// appUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelopeErrors]
type appUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeMessages struct {
	Message string                                `json:"message" api:"required"`
	JSON    appUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// appUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelopeMessages]
type appUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppListParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AppDeleteParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AppDeleteResponseEnvelope struct {
	Errors   []AppDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                `json:"success" api:"required"`
	JSON     appDeleteResponseEnvelopeJSON       `json:"-"`
}

// appDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelope]
type appDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeErrors struct {
	Message string                              `json:"message" api:"required"`
	JSON    appDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// appDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelopeErrors]
type appDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeMessages struct {
	Message string                                `json:"message" api:"required"`
	JSON    appDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// appDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelopeMessages]
type appDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppGetParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AppGetResponseEnvelope struct {
	Errors   []AppGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppGetResponse                   `json:"result" api:"required"`
	Success  bool                             `json:"success" api:"required"`
	JSON     appGetResponseEnvelopeJSON       `json:"-"`
}

// appGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelope]
type appGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeErrors struct {
	Message string                           `json:"message" api:"required"`
	JSON    appGetResponseEnvelopeErrorsJSON `json:"-"`
}

// appGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelopeErrors]
type appGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeMessages struct {
	Message string                             `json:"message" api:"required"`
	JSON    appGetResponseEnvelopeMessagesJSON `json:"-"`
}

// appGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelopeMessages]
type appGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
