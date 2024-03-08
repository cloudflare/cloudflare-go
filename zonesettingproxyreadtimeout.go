// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZoneSettingProxyReadTimeoutService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingProxyReadTimeoutService] method instead.
type ZoneSettingProxyReadTimeoutService struct {
	Options []option.RequestOption
}

// NewZoneSettingProxyReadTimeoutService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingProxyReadTimeoutService(opts ...option.RequestOption) (r *ZoneSettingProxyReadTimeoutService) {
	r = &ZoneSettingProxyReadTimeoutService{}
	r.Options = opts
	return
}

// Maximum time between two read operations from origin.
func (r *ZoneSettingProxyReadTimeoutService) Edit(ctx context.Context, params ZoneSettingProxyReadTimeoutEditParams, opts ...option.RequestOption) (res *ZoneSettingProxyReadTimeoutEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingProxyReadTimeoutEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
func (r *ZoneSettingProxyReadTimeoutService) Get(ctx context.Context, query ZoneSettingProxyReadTimeoutGetParams, opts ...option.RequestOption) (res *ZoneSettingProxyReadTimeoutGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingProxyReadTimeoutGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Maximum time between two read operations from origin.
type ZoneSettingProxyReadTimeoutEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingProxyReadTimeoutEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingProxyReadTimeoutEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingProxyReadTimeoutEditResponseJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingProxyReadTimeoutEditResponse]
type zoneSettingProxyReadTimeoutEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingProxyReadTimeoutEditResponseID string

const (
	ZoneSettingProxyReadTimeoutEditResponseIDProxyReadTimeout ZoneSettingProxyReadTimeoutEditResponseID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingProxyReadTimeoutEditResponseEditable bool

const (
	ZoneSettingProxyReadTimeoutEditResponseEditableTrue  ZoneSettingProxyReadTimeoutEditResponseEditable = true
	ZoneSettingProxyReadTimeoutEditResponseEditableFalse ZoneSettingProxyReadTimeoutEditResponseEditable = false
)

// Maximum time between two read operations from origin.
type ZoneSettingProxyReadTimeoutGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingProxyReadTimeoutGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingProxyReadTimeoutGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingProxyReadTimeoutGetResponseJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingProxyReadTimeoutGetResponse]
type zoneSettingProxyReadTimeoutGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingProxyReadTimeoutGetResponseID string

const (
	ZoneSettingProxyReadTimeoutGetResponseIDProxyReadTimeout ZoneSettingProxyReadTimeoutGetResponseID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingProxyReadTimeoutGetResponseEditable bool

const (
	ZoneSettingProxyReadTimeoutGetResponseEditableTrue  ZoneSettingProxyReadTimeoutGetResponseEditable = true
	ZoneSettingProxyReadTimeoutGetResponseEditableFalse ZoneSettingProxyReadTimeoutGetResponseEditable = false
)

type ZoneSettingProxyReadTimeoutEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Maximum time between two read operations from origin.
	Value param.Field[ZoneSettingProxyReadTimeoutEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingProxyReadTimeoutEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Maximum time between two read operations from origin.
type ZoneSettingProxyReadTimeoutEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingProxyReadTimeoutEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r ZoneSettingProxyReadTimeoutEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingProxyReadTimeoutEditParamsValueID string

const (
	ZoneSettingProxyReadTimeoutEditParamsValueIDProxyReadTimeout ZoneSettingProxyReadTimeoutEditParamsValueID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingProxyReadTimeoutEditParamsValueEditable bool

const (
	ZoneSettingProxyReadTimeoutEditParamsValueEditableTrue  ZoneSettingProxyReadTimeoutEditParamsValueEditable = true
	ZoneSettingProxyReadTimeoutEditParamsValueEditableFalse ZoneSettingProxyReadTimeoutEditParamsValueEditable = false
)

type ZoneSettingProxyReadTimeoutEditResponseEnvelope struct {
	Errors   []ZoneSettingProxyReadTimeoutEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingProxyReadTimeoutEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ZoneSettingProxyReadTimeoutEditResponse             `json:"result"`
	JSON   zoneSettingProxyReadTimeoutEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingProxyReadTimeoutEditResponseEnvelope]
type zoneSettingProxyReadTimeoutEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingProxyReadTimeoutEditResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingProxyReadTimeoutEditResponseEnvelopeErrors]
type zoneSettingProxyReadTimeoutEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingProxyReadTimeoutEditResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingProxyReadTimeoutEditResponseEnvelopeMessages]
type zoneSettingProxyReadTimeoutEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingProxyReadTimeoutGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingProxyReadTimeoutGetResponseEnvelope struct {
	Errors   []ZoneSettingProxyReadTimeoutGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingProxyReadTimeoutGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ZoneSettingProxyReadTimeoutGetResponse             `json:"result"`
	JSON   zoneSettingProxyReadTimeoutGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingProxyReadTimeoutGetResponseEnvelope]
type zoneSettingProxyReadTimeoutGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingProxyReadTimeoutGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingProxyReadTimeoutGetResponseEnvelopeErrors]
type zoneSettingProxyReadTimeoutGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingProxyReadTimeoutGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingProxyReadTimeoutGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingProxyReadTimeoutGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingProxyReadTimeoutGetResponseEnvelopeMessages]
type zoneSettingProxyReadTimeoutGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingProxyReadTimeoutGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingProxyReadTimeoutGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
