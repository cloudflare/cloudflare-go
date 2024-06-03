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

// SiteSiteAppConfigurationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteSiteAppConfigurationService] method instead.
type SiteSiteAppConfigurationService struct {
	Options []option.RequestOption
}

// NewSiteSiteAppConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSiteSiteAppConfigurationService(opts ...option.RequestOption) (r *SiteSiteAppConfigurationService) {
	r = &SiteSiteAppConfigurationService{}
	r.Options = opts
	return
}

// Creates a new App Config for a site
func (r *SiteSiteAppConfigurationService) New(ctx context.Context, siteID string, params SiteSiteAppConfigurationNewParams, opts ...option.RequestOption) (res *SiteSiteAppConfigurationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteSiteAppConfigurationNewResponseEnvelope
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
func (r *SiteSiteAppConfigurationService) Update(ctx context.Context, siteID string, appConfigID string, params SiteSiteAppConfigurationUpdateParams, opts ...option.RequestOption) (res *SiteSiteAppConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteSiteAppConfigurationUpdateResponseEnvelope
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
func (r *SiteSiteAppConfigurationService) List(ctx context.Context, siteID string, query SiteSiteAppConfigurationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SiteSiteAppConfigurationListResponse], err error) {
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
func (r *SiteSiteAppConfigurationService) ListAutoPaging(ctx context.Context, siteID string, query SiteSiteAppConfigurationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SiteSiteAppConfigurationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, siteID, query, opts...))
}

// Deletes specific App Config associated with a site.
func (r *SiteSiteAppConfigurationService) Delete(ctx context.Context, siteID string, appConfigID string, body SiteSiteAppConfigurationDeleteParams, opts ...option.RequestOption) (res *SiteSiteAppConfigurationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteSiteAppConfigurationDeleteResponseEnvelope
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
type SiteSiteAppConfigurationNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                                   `json:"priority"`
	JSON     siteSiteAppConfigurationNewResponseJSON `json:"-"`
}

// siteSiteAppConfigurationNewResponseJSON contains the JSON metadata for the
// struct [SiteSiteAppConfigurationNewResponse]
type siteSiteAppConfigurationNewResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteSiteAppConfigurationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteSiteAppConfigurationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Traffic decision configuration for an app.
type SiteSiteAppConfigurationUpdateResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                                      `json:"priority"`
	JSON     siteSiteAppConfigurationUpdateResponseJSON `json:"-"`
}

// siteSiteAppConfigurationUpdateResponseJSON contains the JSON metadata for the
// struct [SiteSiteAppConfigurationUpdateResponse]
type siteSiteAppConfigurationUpdateResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteSiteAppConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteSiteAppConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Traffic decision configuration for an app.
type SiteSiteAppConfigurationListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                                    `json:"priority"`
	JSON     siteSiteAppConfigurationListResponseJSON `json:"-"`
}

// siteSiteAppConfigurationListResponseJSON contains the JSON metadata for the
// struct [SiteSiteAppConfigurationListResponse]
type siteSiteAppConfigurationListResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteSiteAppConfigurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteSiteAppConfigurationListResponseJSON) RawJSON() string {
	return r.raw
}

// Traffic decision configuration for an app.
type SiteSiteAppConfigurationDeleteResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Identifier
	SiteID string `json:"site_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64                                      `json:"priority"`
	JSON     siteSiteAppConfigurationDeleteResponseJSON `json:"-"`
}

// siteSiteAppConfigurationDeleteResponseJSON contains the JSON metadata for the
// struct [SiteSiteAppConfigurationDeleteResponse]
type siteSiteAppConfigurationDeleteResponseJSON struct {
	ID          apijson.Field
	SiteID      apijson.Field
	Breakout    apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteSiteAppConfigurationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteSiteAppConfigurationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteSiteAppConfigurationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r SiteSiteAppConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteSiteAppConfigurationNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Traffic decision configuration for an app.
	Result SiteSiteAppConfigurationNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SiteSiteAppConfigurationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteSiteAppConfigurationNewResponseEnvelopeJSON    `json:"-"`
}

// siteSiteAppConfigurationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [SiteSiteAppConfigurationNewResponseEnvelope]
type siteSiteAppConfigurationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteSiteAppConfigurationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteSiteAppConfigurationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteSiteAppConfigurationNewResponseEnvelopeSuccess bool

const (
	SiteSiteAppConfigurationNewResponseEnvelopeSuccessTrue SiteSiteAppConfigurationNewResponseEnvelopeSuccess = true
)

func (r SiteSiteAppConfigurationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteSiteAppConfigurationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteSiteAppConfigurationUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r SiteSiteAppConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteSiteAppConfigurationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Traffic decision configuration for an app.
	Result SiteSiteAppConfigurationUpdateResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SiteSiteAppConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteSiteAppConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteSiteAppConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [SiteSiteAppConfigurationUpdateResponseEnvelope]
type siteSiteAppConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteSiteAppConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteSiteAppConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteSiteAppConfigurationUpdateResponseEnvelopeSuccess bool

const (
	SiteSiteAppConfigurationUpdateResponseEnvelopeSuccessTrue SiteSiteAppConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r SiteSiteAppConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteSiteAppConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteSiteAppConfigurationListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteSiteAppConfigurationDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteSiteAppConfigurationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Traffic decision configuration for an app.
	Result SiteSiteAppConfigurationDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SiteSiteAppConfigurationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteSiteAppConfigurationDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteSiteAppConfigurationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [SiteSiteAppConfigurationDeleteResponseEnvelope]
type siteSiteAppConfigurationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteSiteAppConfigurationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteSiteAppConfigurationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteSiteAppConfigurationDeleteResponseEnvelopeSuccess bool

const (
	SiteSiteAppConfigurationDeleteResponseEnvelopeSuccessTrue SiteSiteAppConfigurationDeleteResponseEnvelopeSuccess = true
)

func (r SiteSiteAppConfigurationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteSiteAppConfigurationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
