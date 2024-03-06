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

// ZoneSettingResponseBufferingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingResponseBufferingService] method instead.
type ZoneSettingResponseBufferingService struct {
	Options []option.RequestOption
}

// NewZoneSettingResponseBufferingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingResponseBufferingService(opts ...option.RequestOption) (r *ZoneSettingResponseBufferingService) {
	r = &ZoneSettingResponseBufferingService{}
	r.Options = opts
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *ZoneSettingResponseBufferingService) Edit(ctx context.Context, params ZoneSettingResponseBufferingEditParams, opts ...option.RequestOption) (res *ZoneSettingResponseBufferingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingResponseBufferingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *ZoneSettingResponseBufferingService) Get(ctx context.Context, query ZoneSettingResponseBufferingGetParams, opts ...option.RequestOption) (res *ZoneSettingResponseBufferingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingResponseBufferingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingResponseBufferingEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingResponseBufferingEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingResponseBufferingEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingResponseBufferingEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingResponseBufferingEditResponseJSON `json:"-"`
}

// zoneSettingResponseBufferingEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingResponseBufferingEditResponse]
type zoneSettingResponseBufferingEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingResponseBufferingEditResponseID string

const (
	ZoneSettingResponseBufferingEditResponseIDResponseBuffering ZoneSettingResponseBufferingEditResponseID = "response_buffering"
)

// Current value of the zone setting.
type ZoneSettingResponseBufferingEditResponseValue string

const (
	ZoneSettingResponseBufferingEditResponseValueOn  ZoneSettingResponseBufferingEditResponseValue = "on"
	ZoneSettingResponseBufferingEditResponseValueOff ZoneSettingResponseBufferingEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingResponseBufferingEditResponseEditable bool

const (
	ZoneSettingResponseBufferingEditResponseEditableTrue  ZoneSettingResponseBufferingEditResponseEditable = true
	ZoneSettingResponseBufferingEditResponseEditableFalse ZoneSettingResponseBufferingEditResponseEditable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingResponseBufferingGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingResponseBufferingGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingResponseBufferingGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingResponseBufferingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingResponseBufferingGetResponseJSON `json:"-"`
}

// zoneSettingResponseBufferingGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingResponseBufferingGetResponse]
type zoneSettingResponseBufferingGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingResponseBufferingGetResponseID string

const (
	ZoneSettingResponseBufferingGetResponseIDResponseBuffering ZoneSettingResponseBufferingGetResponseID = "response_buffering"
)

// Current value of the zone setting.
type ZoneSettingResponseBufferingGetResponseValue string

const (
	ZoneSettingResponseBufferingGetResponseValueOn  ZoneSettingResponseBufferingGetResponseValue = "on"
	ZoneSettingResponseBufferingGetResponseValueOff ZoneSettingResponseBufferingGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingResponseBufferingGetResponseEditable bool

const (
	ZoneSettingResponseBufferingGetResponseEditableTrue  ZoneSettingResponseBufferingGetResponseEditable = true
	ZoneSettingResponseBufferingGetResponseEditableFalse ZoneSettingResponseBufferingGetResponseEditable = false
)

type ZoneSettingResponseBufferingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingResponseBufferingEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingResponseBufferingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingResponseBufferingEditParamsValue string

const (
	ZoneSettingResponseBufferingEditParamsValueOn  ZoneSettingResponseBufferingEditParamsValue = "on"
	ZoneSettingResponseBufferingEditParamsValueOff ZoneSettingResponseBufferingEditParamsValue = "off"
)

type ZoneSettingResponseBufferingEditResponseEnvelope struct {
	Errors   []ZoneSettingResponseBufferingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingResponseBufferingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ZoneSettingResponseBufferingEditResponse             `json:"result"`
	JSON   zoneSettingResponseBufferingEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingResponseBufferingEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingResponseBufferingEditResponseEnvelope]
type zoneSettingResponseBufferingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingResponseBufferingEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingResponseBufferingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingResponseBufferingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingResponseBufferingEditResponseEnvelopeErrors]
type zoneSettingResponseBufferingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingResponseBufferingEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneSettingResponseBufferingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingResponseBufferingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingResponseBufferingEditResponseEnvelopeMessages]
type zoneSettingResponseBufferingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingResponseBufferingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingResponseBufferingGetResponseEnvelope struct {
	Errors   []ZoneSettingResponseBufferingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingResponseBufferingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result ZoneSettingResponseBufferingGetResponse             `json:"result"`
	JSON   zoneSettingResponseBufferingGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingResponseBufferingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingResponseBufferingGetResponseEnvelope]
type zoneSettingResponseBufferingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingResponseBufferingGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingResponseBufferingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingResponseBufferingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingResponseBufferingGetResponseEnvelopeErrors]
type zoneSettingResponseBufferingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingResponseBufferingGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingResponseBufferingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingResponseBufferingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingResponseBufferingGetResponseEnvelopeMessages]
type zoneSettingResponseBufferingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingResponseBufferingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingResponseBufferingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
