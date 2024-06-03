// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SiteAppConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteAppConfigurationService] method instead.
type SiteAppConfigurationService struct {
	Options []option.RequestOption
}

// NewSiteAppConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSiteAppConfigurationService(opts ...option.RequestOption) (r *SiteAppConfigurationService) {
	r = &SiteAppConfigurationService{}
	r.Options = opts
	return
}

// Creates a new App Config for a site
func (r *SiteAppConfigurationService) New(ctx context.Context, siteID string, params SiteAppConfigurationNewParams, opts ...option.RequestOption) (res *SiteAppConfigurationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteAppConfigurationNewResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an App Config for a site
func (r *SiteAppConfigurationService) Update(ctx context.Context, siteID string, appConfigID string, params SiteAppConfigurationUpdateParams, opts ...option.RequestOption) (res *SiteAppConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteAppConfigurationUpdateResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if appConfigID == "" {
		err = errors.New("missing required app_config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs/%s", params.AccountID, siteID, appConfigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists App Configs associated with a site.
func (r *SiteAppConfigurationService) List(ctx context.Context, siteID string, query SiteAppConfigurationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SiteAppConfigurationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs", query.AccountID, siteID)
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

// Lists App Configs associated with a site.
func (r *SiteAppConfigurationService) ListAutoPaging(ctx context.Context, siteID string, query SiteAppConfigurationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SiteAppConfigurationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, siteID, query, opts...))
}

// Deletes specific App Config associated with a site.
func (r *SiteAppConfigurationService) Delete(ctx context.Context, siteID string, appConfigID string, body SiteAppConfigurationDeleteParams, opts ...option.RequestOption) (res *SiteAppConfigurationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteAppConfigurationDeleteResponseEnvelope
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if appConfigID == "" {
		err = errors.New("missing required app_config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs/%s", body.AccountID, siteID, appConfigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Traffic decision configuration for an app.
type SiteAppConfigurationNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                               `json:"priority"`
	JSON     siteAppConfigurationNewResponseJSON `json:"-"`
}

// siteAppConfigurationNewResponseJSON contains the JSON metadata for the struct
// [SiteAppConfigurationNewResponse]
type siteAppConfigurationNewResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteAppConfigurationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteAppConfigurationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Traffic decision configuration for an app.
type SiteAppConfigurationUpdateResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                                  `json:"priority"`
	JSON     siteAppConfigurationUpdateResponseJSON `json:"-"`
}

// siteAppConfigurationUpdateResponseJSON contains the JSON metadata for the struct
// [SiteAppConfigurationUpdateResponse]
type siteAppConfigurationUpdateResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteAppConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteAppConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Traffic decision configuration for an app.
type SiteAppConfigurationListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                                `json:"priority"`
	JSON     siteAppConfigurationListResponseJSON `json:"-"`
}

// siteAppConfigurationListResponseJSON contains the JSON metadata for the struct
// [SiteAppConfigurationListResponse]
type siteAppConfigurationListResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteAppConfigurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteAppConfigurationListResponseJSON) RawJSON() string {
	return r.raw
}

// Traffic decision configuration for an app.
type SiteAppConfigurationDeleteResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                                  `json:"priority"`
	JSON     siteAppConfigurationDeleteResponseJSON `json:"-"`
}

// siteAppConfigurationDeleteResponseJSON contains the JSON metadata for the struct
// [SiteAppConfigurationDeleteResponse]
type siteAppConfigurationDeleteResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteAppConfigurationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteAppConfigurationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteAppConfigurationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r SiteAppConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteAppConfigurationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Traffic decision configuration for an app.
	Result SiteAppConfigurationNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SiteAppConfigurationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteAppConfigurationNewResponseEnvelopeJSON    `json:"-"`
}

// siteAppConfigurationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SiteAppConfigurationNewResponseEnvelope]
type siteAppConfigurationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteAppConfigurationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteAppConfigurationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteAppConfigurationNewResponseEnvelopeSuccess bool

const (
	SiteAppConfigurationNewResponseEnvelopeSuccessTrue SiteAppConfigurationNewResponseEnvelopeSuccess = true
)

func (r SiteAppConfigurationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteAppConfigurationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteAppConfigurationUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r SiteAppConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteAppConfigurationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Traffic decision configuration for an app.
	Result SiteAppConfigurationUpdateResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SiteAppConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteAppConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteAppConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SiteAppConfigurationUpdateResponseEnvelope]
type siteAppConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteAppConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteAppConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteAppConfigurationUpdateResponseEnvelopeSuccess bool

const (
	SiteAppConfigurationUpdateResponseEnvelopeSuccessTrue SiteAppConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r SiteAppConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteAppConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteAppConfigurationListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteAppConfigurationDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteAppConfigurationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Traffic decision configuration for an app.
	Result SiteAppConfigurationDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SiteAppConfigurationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteAppConfigurationDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteAppConfigurationDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SiteAppConfigurationDeleteResponseEnvelope]
type siteAppConfigurationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteAppConfigurationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteAppConfigurationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteAppConfigurationDeleteResponseEnvelopeSuccess bool

const (
	SiteAppConfigurationDeleteResponseEnvelopeSuccessTrue SiteAppConfigurationDeleteResponseEnvelopeSuccess = true
)

func (r SiteAppConfigurationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteAppConfigurationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
