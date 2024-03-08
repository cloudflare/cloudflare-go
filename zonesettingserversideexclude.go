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

// ZoneSettingServerSideExcludeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingServerSideExcludeService] method instead.
type ZoneSettingServerSideExcludeService struct {
	Options []option.RequestOption
}

// NewZoneSettingServerSideExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingServerSideExcludeService(opts ...option.RequestOption) (r *ZoneSettingServerSideExcludeService) {
	r = &ZoneSettingServerSideExcludeService{}
	r.Options = opts
	return
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
func (r *ZoneSettingServerSideExcludeService) Edit(ctx context.Context, params ZoneSettingServerSideExcludeEditParams, opts ...option.RequestOption) (res *ZoneSettingServerSideExcludeEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingServerSideExcludeEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
func (r *ZoneSettingServerSideExcludeService) Get(ctx context.Context, query ZoneSettingServerSideExcludeGetParams, opts ...option.RequestOption) (res *ZoneSettingServerSideExcludeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingServerSideExcludeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ZoneSettingServerSideExcludeEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingServerSideExcludeEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingServerSideExcludeEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingServerSideExcludeEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingServerSideExcludeEditResponseJSON `json:"-"`
}

// zoneSettingServerSideExcludeEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingServerSideExcludeEditResponse]
type zoneSettingServerSideExcludeEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingServerSideExcludeEditResponseID string

const (
	ZoneSettingServerSideExcludeEditResponseIDServerSideExclude ZoneSettingServerSideExcludeEditResponseID = "server_side_exclude"
)

// Current value of the zone setting.
type ZoneSettingServerSideExcludeEditResponseValue string

const (
	ZoneSettingServerSideExcludeEditResponseValueOn  ZoneSettingServerSideExcludeEditResponseValue = "on"
	ZoneSettingServerSideExcludeEditResponseValueOff ZoneSettingServerSideExcludeEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingServerSideExcludeEditResponseEditable bool

const (
	ZoneSettingServerSideExcludeEditResponseEditableTrue  ZoneSettingServerSideExcludeEditResponseEditable = true
	ZoneSettingServerSideExcludeEditResponseEditableFalse ZoneSettingServerSideExcludeEditResponseEditable = false
)

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ZoneSettingServerSideExcludeGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingServerSideExcludeGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingServerSideExcludeGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingServerSideExcludeGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingServerSideExcludeGetResponseJSON `json:"-"`
}

// zoneSettingServerSideExcludeGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingServerSideExcludeGetResponse]
type zoneSettingServerSideExcludeGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingServerSideExcludeGetResponseID string

const (
	ZoneSettingServerSideExcludeGetResponseIDServerSideExclude ZoneSettingServerSideExcludeGetResponseID = "server_side_exclude"
)

// Current value of the zone setting.
type ZoneSettingServerSideExcludeGetResponseValue string

const (
	ZoneSettingServerSideExcludeGetResponseValueOn  ZoneSettingServerSideExcludeGetResponseValue = "on"
	ZoneSettingServerSideExcludeGetResponseValueOff ZoneSettingServerSideExcludeGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingServerSideExcludeGetResponseEditable bool

const (
	ZoneSettingServerSideExcludeGetResponseEditableTrue  ZoneSettingServerSideExcludeGetResponseEditable = true
	ZoneSettingServerSideExcludeGetResponseEditableFalse ZoneSettingServerSideExcludeGetResponseEditable = false
)

type ZoneSettingServerSideExcludeEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingServerSideExcludeEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingServerSideExcludeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingServerSideExcludeEditParamsValue string

const (
	ZoneSettingServerSideExcludeEditParamsValueOn  ZoneSettingServerSideExcludeEditParamsValue = "on"
	ZoneSettingServerSideExcludeEditParamsValueOff ZoneSettingServerSideExcludeEditParamsValue = "off"
)

type ZoneSettingServerSideExcludeEditResponseEnvelope struct {
	Errors   []ZoneSettingServerSideExcludeEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingServerSideExcludeEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// If there is sensitive content on your website that you want visible to real
	// visitors, but that you want to hide from suspicious visitors, all you have to do
	// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
	// be excluded from suspicious visitors in the following SSE tags:
	// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
	// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
	// have HTML minification enabled, you won't see the SSE tags in your HTML source
	// when it's served through Cloudflare. SSE will still function in this case, as
	// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
	// resource moves through our network to the visitor's computer.
	// (https://support.cloudflare.com/hc/en-us/articles/200170036).
	Result ZoneSettingServerSideExcludeEditResponse             `json:"result"`
	JSON   zoneSettingServerSideExcludeEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingServerSideExcludeEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingServerSideExcludeEditResponseEnvelope]
type zoneSettingServerSideExcludeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingServerSideExcludeEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingServerSideExcludeEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingServerSideExcludeEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingServerSideExcludeEditResponseEnvelopeErrors]
type zoneSettingServerSideExcludeEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingServerSideExcludeEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneSettingServerSideExcludeEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingServerSideExcludeEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingServerSideExcludeEditResponseEnvelopeMessages]
type zoneSettingServerSideExcludeEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingServerSideExcludeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingServerSideExcludeGetResponseEnvelope struct {
	Errors   []ZoneSettingServerSideExcludeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingServerSideExcludeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// If there is sensitive content on your website that you want visible to real
	// visitors, but that you want to hide from suspicious visitors, all you have to do
	// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
	// be excluded from suspicious visitors in the following SSE tags:
	// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
	// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
	// have HTML minification enabled, you won't see the SSE tags in your HTML source
	// when it's served through Cloudflare. SSE will still function in this case, as
	// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
	// resource moves through our network to the visitor's computer.
	// (https://support.cloudflare.com/hc/en-us/articles/200170036).
	Result ZoneSettingServerSideExcludeGetResponse             `json:"result"`
	JSON   zoneSettingServerSideExcludeGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingServerSideExcludeGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingServerSideExcludeGetResponseEnvelope]
type zoneSettingServerSideExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingServerSideExcludeGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingServerSideExcludeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingServerSideExcludeGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingServerSideExcludeGetResponseEnvelopeErrors]
type zoneSettingServerSideExcludeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingServerSideExcludeGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingServerSideExcludeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingServerSideExcludeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingServerSideExcludeGetResponseEnvelopeMessages]
type zoneSettingServerSideExcludeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingServerSideExcludeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
