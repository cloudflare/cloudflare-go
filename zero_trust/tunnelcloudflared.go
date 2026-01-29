// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// TunnelCloudflaredService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelCloudflaredService] method instead.
type TunnelCloudflaredService struct {
	Options        []option.RequestOption
	Configurations *TunnelCloudflaredConfigurationService
	Connections    *TunnelCloudflaredConnectionService
	Token          *TunnelCloudflaredTokenService
	Connectors     *TunnelCloudflaredConnectorService
	Management     *TunnelCloudflaredManagementService
}

// NewTunnelCloudflaredService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelCloudflaredService(opts ...option.RequestOption) (r *TunnelCloudflaredService) {
	r = &TunnelCloudflaredService{}
	r.Options = opts
	r.Configurations = NewTunnelCloudflaredConfigurationService(opts...)
	r.Connections = NewTunnelCloudflaredConnectionService(opts...)
	r.Token = NewTunnelCloudflaredTokenService(opts...)
	r.Connectors = NewTunnelCloudflaredConnectorService(opts...)
	r.Management = NewTunnelCloudflaredManagementService(opts...)
	return
}

// Creates a new Cloudflare Tunnel in an account.
func (r *TunnelCloudflaredService) New(ctx context.Context, params TunnelCloudflaredNewParams, opts ...option.RequestOption) (res *shared.CloudflareTunnel, err error) {
	var env TunnelCloudflaredNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters Cloudflare Tunnels in an account.
func (r *TunnelCloudflaredService) List(ctx context.Context, params TunnelCloudflaredListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[shared.CloudflareTunnel], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", params.AccountID)
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

// Lists and filters Cloudflare Tunnels in an account.
func (r *TunnelCloudflaredService) ListAutoPaging(ctx context.Context, params TunnelCloudflaredListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[shared.CloudflareTunnel] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Cloudflare Tunnel from an account.
func (r *TunnelCloudflaredService) Delete(ctx context.Context, tunnelID string, body TunnelCloudflaredDeleteParams, opts ...option.RequestOption) (res *shared.CloudflareTunnel, err error) {
	var env TunnelCloudflaredDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", body.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *TunnelCloudflaredService) Edit(ctx context.Context, tunnelID string, params TunnelCloudflaredEditParams, opts ...option.RequestOption) (res *shared.CloudflareTunnel, err error) {
	var env TunnelCloudflaredEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Cloudflare Tunnel.
func (r *TunnelCloudflaredService) Get(ctx context.Context, tunnelID string, query TunnelCloudflaredGetParams, opts ...option.RequestOption) (res *shared.CloudflareTunnel, err error) {
	var env TunnelCloudflaredGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelCloudflaredNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name,required"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc param.Field[TunnelCloudflaredNewParamsConfigSrc] `json:"config_src"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r TunnelCloudflaredNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelCloudflaredNewParamsConfigSrc string

const (
	TunnelCloudflaredNewParamsConfigSrcLocal      TunnelCloudflaredNewParamsConfigSrc = "local"
	TunnelCloudflaredNewParamsConfigSrcCloudflare TunnelCloudflaredNewParamsConfigSrc = "cloudflare"
)

func (r TunnelCloudflaredNewParamsConfigSrc) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewParamsConfigSrcLocal, TunnelCloudflaredNewParamsConfigSrcCloudflare:
		return true
	}
	return false
}

type TunnelCloudflaredNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result shared.CloudflareTunnel `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredNewResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredNewResponseEnvelope]
type tunnelCloudflaredNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredNewResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredNewResponseEnvelopeSuccessTrue TunnelCloudflaredNewResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredListParams struct {
	// Cloudflare account ID
	AccountID     param.Field[string] `path:"account_id,required"`
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt     param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	IncludePrefix param.Field[string] `query:"include_prefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status param.Field[TunnelCloudflaredListParamsStatus] `query:"status"`
	// UUID of the tunnel.
	UUID          param.Field[string]    `query:"uuid" format:"uuid"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [TunnelCloudflaredListParams]'s query parameters as
// `url.Values`.
func (r TunnelCloudflaredListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredListParamsStatus string

const (
	TunnelCloudflaredListParamsStatusInactive TunnelCloudflaredListParamsStatus = "inactive"
	TunnelCloudflaredListParamsStatusDegraded TunnelCloudflaredListParamsStatus = "degraded"
	TunnelCloudflaredListParamsStatusHealthy  TunnelCloudflaredListParamsStatus = "healthy"
	TunnelCloudflaredListParamsStatusDown     TunnelCloudflaredListParamsStatus = "down"
)

func (r TunnelCloudflaredListParamsStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredListParamsStatusInactive, TunnelCloudflaredListParamsStatusDegraded, TunnelCloudflaredListParamsStatusHealthy, TunnelCloudflaredListParamsStatusDown:
		return true
	}
	return false
}

type TunnelCloudflaredDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelCloudflaredDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result shared.CloudflareTunnel `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredDeleteResponseEnvelope]
type tunnelCloudflaredDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredDeleteResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredDeleteResponseEnvelopeSuccessTrue TunnelCloudflaredDeleteResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r TunnelCloudflaredEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelCloudflaredEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result shared.CloudflareTunnel `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredEditResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredEditResponseEnvelope]
type tunnelCloudflaredEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredEditResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredEditResponseEnvelopeSuccessTrue TunnelCloudflaredEditResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelCloudflaredGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result shared.CloudflareTunnel `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredGetResponseEnvelope]
type tunnelCloudflaredGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredGetResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredGetResponseEnvelopeSuccessTrue TunnelCloudflaredGetResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
