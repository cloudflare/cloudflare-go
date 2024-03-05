// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *ZoneSettingProxyReadTimeoutService) Edit(ctx context.Context, params ZoneSettingProxyReadTimeoutEditParams, opts ...option.RequestOption) (res *ZonesProxyReadTimeout, err error) {
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
func (r *ZoneSettingProxyReadTimeoutService) Get(ctx context.Context, query ZoneSettingProxyReadTimeoutGetParams, opts ...option.RequestOption) (res *ZonesProxyReadTimeout, err error) {
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
type ZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesProxyReadTimeoutJSON `json:"-"`
}

// zonesProxyReadTimeoutJSON contains the JSON metadata for the struct
// [ZonesProxyReadTimeout]
type zonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesProxyReadTimeout) implementsZoneSettingEditResponse() {}

func (r ZonesProxyReadTimeout) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesProxyReadTimeoutID string

const (
	ZonesProxyReadTimeoutIDProxyReadTimeout ZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesProxyReadTimeoutEditable bool

const (
	ZonesProxyReadTimeoutEditableTrue  ZonesProxyReadTimeoutEditable = true
	ZonesProxyReadTimeoutEditableFalse ZonesProxyReadTimeoutEditable = false
)

// Maximum time between two read operations from origin.
type ZonesProxyReadTimeoutParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesProxyReadTimeoutID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r ZonesProxyReadTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesProxyReadTimeoutParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingProxyReadTimeoutEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Maximum time between two read operations from origin.
	Value param.Field[ZonesProxyReadTimeoutParam] `json:"value,required"`
}

func (r ZoneSettingProxyReadTimeoutEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingProxyReadTimeoutEditResponseEnvelope struct {
	Errors   []ZoneSettingProxyReadTimeoutEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingProxyReadTimeoutEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Maximum time between two read operations from origin.
	Result ZonesProxyReadTimeout                               `json:"result"`
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
	Result ZonesProxyReadTimeout                              `json:"result"`
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
