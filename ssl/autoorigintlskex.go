// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// AutoOriginTLSKexService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutoOriginTLSKexService] method instead.
type AutoOriginTLSKexService struct {
	Options []option.RequestOption
}

// NewAutoOriginTLSKexService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAutoOriginTLSKexService(opts ...option.RequestOption) (r *AutoOriginTLSKexService) {
	r = &AutoOriginTLSKexService{}
	r.Options = opts
	return
}

// Enable or disable Auto-Origin TLS KEX selection for the zone by sending
// `{"enabled": true}` or `{"enabled": false}`. When enabled, Cloudflare runs a
// periodic scan of the zone's origins to determine the preferred key-exchange
// algorithm and writes that preference to the edge so it is sent first in the TLS
// ClientHello to the origin.
func (r *AutoOriginTLSKexService) Edit(ctx context.Context, params AutoOriginTLSKexEditParams, opts ...option.RequestOption) (res *AutoOriginTLSKexEditResponse, err error) {
	var env AutoOriginTLSKexEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/auto_origin_tls_kex", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// When enabled, Cloudflare automatically selects the preferred TLS key-exchange
// algorithm to use when establishing the TLS connection to the zone's origin,
// picking from the algorithms permitted by the zone's
// `origin_tls_compliance_modes` setting. When disabled, the default key-exchange
// ordering is used.
func (r *AutoOriginTLSKexService) Get(ctx context.Context, query AutoOriginTLSKexGetParams, opts ...option.RequestOption) (res *AutoOriginTLSKexGetResponse, err error) {
	var env AutoOriginTLSKexGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/auto_origin_tls_kex", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AutoOriginTLSKexEditResponse struct {
	ID string `json:"id" api:"required"`
	// Whether Auto-Origin TLS KEX selection is enabled for the zone.
	Enabled bool `json:"enabled" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on" api:"required" format:"date-time"`
	JSON       autoOriginTLSKexEditResponseJSON `json:"-"`
}

// autoOriginTLSKexEditResponseJSON contains the JSON metadata for the struct
// [AutoOriginTLSKexEditResponse]
type autoOriginTLSKexEditResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutoOriginTLSKexEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r autoOriginTLSKexEditResponseJSON) RawJSON() string {
	return r.raw
}

type AutoOriginTLSKexGetResponse struct {
	ID string `json:"id" api:"required"`
	// Whether Auto-Origin TLS KEX selection is enabled for the zone.
	Enabled bool `json:"enabled" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on" api:"required" format:"date-time"`
	JSON       autoOriginTLSKexGetResponseJSON `json:"-"`
}

// autoOriginTLSKexGetResponseJSON contains the JSON metadata for the struct
// [AutoOriginTLSKexGetResponse]
type autoOriginTLSKexGetResponseJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutoOriginTLSKexGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r autoOriginTLSKexGetResponseJSON) RawJSON() string {
	return r.raw
}

type AutoOriginTLSKexEditParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Controls enablement of Auto-Origin TLS KEX selection for the zone.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r AutoOriginTLSKexEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AutoOriginTLSKexEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors" api:"required"`
	Messages []shared.ResponseInfo        `json:"messages" api:"required"`
	Result   AutoOriginTLSKexEditResponse `json:"result" api:"required"`
	// Indicates the API call's success or failure.
	Success bool                                     `json:"success" api:"required"`
	JSON    autoOriginTLSKexEditResponseEnvelopeJSON `json:"-"`
}

// autoOriginTLSKexEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [AutoOriginTLSKexEditResponseEnvelope]
type autoOriginTLSKexEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutoOriginTLSKexEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r autoOriginTLSKexEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AutoOriginTLSKexGetParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type AutoOriginTLSKexGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors" api:"required"`
	Messages []shared.ResponseInfo       `json:"messages" api:"required"`
	Result   AutoOriginTLSKexGetResponse `json:"result" api:"required"`
	// Indicates the API call's success or failure.
	Success bool                                    `json:"success" api:"required"`
	JSON    autoOriginTLSKexGetResponseEnvelopeJSON `json:"-"`
}

// autoOriginTLSKexGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AutoOriginTLSKexGetResponseEnvelope]
type autoOriginTLSKexGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutoOriginTLSKexGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r autoOriginTLSKexGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
