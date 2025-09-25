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

// NetworkHostnameRouteService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkHostnameRouteService] method instead.
type NetworkHostnameRouteService struct {
	Options []option.RequestOption
}

// NewNetworkHostnameRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkHostnameRouteService(opts ...option.RequestOption) (r *NetworkHostnameRouteService) {
	r = &NetworkHostnameRouteService{}
	r.Options = opts
	return
}

// Create a hostname route.
func (r *NetworkHostnameRouteService) New(ctx context.Context, params NetworkHostnameRouteNewParams, opts ...option.RequestOption) (res *HostnameRoute, err error) {
	var env NetworkHostnameRouteNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/routes/hostname", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters hostname routes in an account.
func (r *NetworkHostnameRouteService) List(ctx context.Context, params NetworkHostnameRouteListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[HostnameRoute], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/routes/hostname", params.AccountID)
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

// Lists and filters hostname routes in an account.
func (r *NetworkHostnameRouteService) ListAutoPaging(ctx context.Context, params NetworkHostnameRouteListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[HostnameRoute] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a hostname route.
func (r *NetworkHostnameRouteService) Delete(ctx context.Context, hostnameRouteID string, body NetworkHostnameRouteDeleteParams, opts ...option.RequestOption) (res *HostnameRoute, err error) {
	var env NetworkHostnameRouteDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hostnameRouteID == "" {
		err = errors.New("missing required hostname_route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/routes/hostname/%s", body.AccountID, hostnameRouteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a hostname route.
func (r *NetworkHostnameRouteService) Edit(ctx context.Context, hostnameRouteID string, params NetworkHostnameRouteEditParams, opts ...option.RequestOption) (res *HostnameRoute, err error) {
	var env NetworkHostnameRouteEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hostnameRouteID == "" {
		err = errors.New("missing required hostname_route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/routes/hostname/%s", params.AccountID, hostnameRouteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a hostname route.
func (r *NetworkHostnameRouteService) Get(ctx context.Context, hostnameRouteID string, query NetworkHostnameRouteGetParams, opts ...option.RequestOption) (res *HostnameRoute, err error) {
	var env NetworkHostnameRouteGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if hostnameRouteID == "" {
		err = errors.New("missing required hostname_route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/routes/hostname/%s", query.AccountID, hostnameRouteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HostnameRoute struct {
	// The hostname route ID.
	ID string `json:"id" format:"uuid"`
	// An optional description of the hostname route.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// The hostname of the route.
	Hostname string `json:"hostname"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// A user-friendly name for a tunnel.
	TunnelName string            `json:"tunnel_name"`
	JSON       hostnameRouteJSON `json:"-"`
}

// hostnameRouteJSON contains the JSON metadata for the struct [HostnameRoute]
type hostnameRouteJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedAt   apijson.Field
	DeletedAt   apijson.Field
	Hostname    apijson.Field
	TunnelID    apijson.Field
	TunnelName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameRouteJSON) RawJSON() string {
	return r.raw
}

type NetworkHostnameRouteNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description of the hostname route.
	Comment param.Field[string] `json:"comment"`
	// The hostname of the route.
	Hostname param.Field[string] `json:"hostname"`
	// UUID of the tunnel.
	TunnelID param.Field[string] `json:"tunnel_id" format:"uuid"`
}

func (r NetworkHostnameRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkHostnameRouteNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   HostnameRoute         `json:"result,required"`
	// Whether the API call was successful
	Success NetworkHostnameRouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkHostnameRouteNewResponseEnvelopeJSON    `json:"-"`
}

// networkHostnameRouteNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkHostnameRouteNewResponseEnvelope]
type networkHostnameRouteNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkHostnameRouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkHostnameRouteNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkHostnameRouteNewResponseEnvelopeSuccess bool

const (
	NetworkHostnameRouteNewResponseEnvelopeSuccessTrue NetworkHostnameRouteNewResponseEnvelopeSuccess = true
)

func (r NetworkHostnameRouteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkHostnameRouteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkHostnameRouteListParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The hostname route ID.
	ID param.Field[string] `query:"id" format:"uuid"`
	// If set, only list hostname routes with the given comment.
	Comment param.Field[string] `query:"comment"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	// If set, only list hostname routes that contain a substring of the given value,
	// the filter is case-insensitive.
	Hostname param.Field[string] `query:"hostname"`
	// If `true`, only return deleted hostname routes. If `false`, exclude deleted
	// hostname routes.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// If set, only list hostname routes that point to a specific tunnel.
	TunnelID param.Field[string] `query:"tunnel_id" format:"uuid"`
}

// URLQuery serializes [NetworkHostnameRouteListParams]'s query parameters as
// `url.Values`.
func (r NetworkHostnameRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NetworkHostnameRouteDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkHostnameRouteDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   HostnameRoute         `json:"result,required"`
	// Whether the API call was successful
	Success NetworkHostnameRouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkHostnameRouteDeleteResponseEnvelopeJSON    `json:"-"`
}

// networkHostnameRouteDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [NetworkHostnameRouteDeleteResponseEnvelope]
type networkHostnameRouteDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkHostnameRouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkHostnameRouteDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkHostnameRouteDeleteResponseEnvelopeSuccess bool

const (
	NetworkHostnameRouteDeleteResponseEnvelopeSuccessTrue NetworkHostnameRouteDeleteResponseEnvelopeSuccess = true
)

func (r NetworkHostnameRouteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkHostnameRouteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkHostnameRouteEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description of the hostname route.
	Comment param.Field[string] `json:"comment"`
	// The hostname of the route.
	Hostname param.Field[string] `json:"hostname"`
	// UUID of the tunnel.
	TunnelID param.Field[string] `json:"tunnel_id" format:"uuid"`
}

func (r NetworkHostnameRouteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkHostnameRouteEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   HostnameRoute         `json:"result,required"`
	// Whether the API call was successful
	Success NetworkHostnameRouteEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkHostnameRouteEditResponseEnvelopeJSON    `json:"-"`
}

// networkHostnameRouteEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkHostnameRouteEditResponseEnvelope]
type networkHostnameRouteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkHostnameRouteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkHostnameRouteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkHostnameRouteEditResponseEnvelopeSuccess bool

const (
	NetworkHostnameRouteEditResponseEnvelopeSuccessTrue NetworkHostnameRouteEditResponseEnvelopeSuccess = true
)

func (r NetworkHostnameRouteEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkHostnameRouteEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkHostnameRouteGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkHostnameRouteGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   HostnameRoute         `json:"result,required"`
	// Whether the API call was successful
	Success NetworkHostnameRouteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkHostnameRouteGetResponseEnvelopeJSON    `json:"-"`
}

// networkHostnameRouteGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkHostnameRouteGetResponseEnvelope]
type networkHostnameRouteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkHostnameRouteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkHostnameRouteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkHostnameRouteGetResponseEnvelopeSuccess bool

const (
	NetworkHostnameRouteGetResponseEnvelopeSuccessTrue NetworkHostnameRouteGetResponseEnvelopeSuccess = true
)

func (r NetworkHostnameRouteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkHostnameRouteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
