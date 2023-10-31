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
func (r *ZoneSettingServerSideExcludeService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingServerSideExcludeUpdateParams, opts ...option.RequestOption) (res *ZoneSettingServerSideExcludeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
func (r *ZoneSettingServerSideExcludeService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingServerSideExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingServerSideExcludeUpdateResponse struct {
	Errors   []ZoneSettingServerSideExcludeUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingServerSideExcludeUpdateResponseMessage `json:"messages"`
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
	Result ZoneSettingServerSideExcludeUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingServerSideExcludeUpdateResponseJSON
}

// zoneSettingServerSideExcludeUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingServerSideExcludeUpdateResponse]
type zoneSettingServerSideExcludeUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingServerSideExcludeUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingServerSideExcludeUpdateResponseErrorJSON
}

// zoneSettingServerSideExcludeUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingServerSideExcludeUpdateResponseError]
type zoneSettingServerSideExcludeUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingServerSideExcludeUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingServerSideExcludeUpdateResponseMessageJSON
}

// zoneSettingServerSideExcludeUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingServerSideExcludeUpdateResponseMessage]
type zoneSettingServerSideExcludeUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type ZoneSettingServerSideExcludeUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingServerSideExcludeUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingServerSideExcludeUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingServerSideExcludeUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingServerSideExcludeUpdateResponseResultJSON
}

// zoneSettingServerSideExcludeUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingServerSideExcludeUpdateResponseResult]
type zoneSettingServerSideExcludeUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingServerSideExcludeUpdateResponseResultID string

const (
	ZoneSettingServerSideExcludeUpdateResponseResultIDServerSideExclude ZoneSettingServerSideExcludeUpdateResponseResultID = "server_side_exclude"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingServerSideExcludeUpdateResponseResultEditable bool

const (
	ZoneSettingServerSideExcludeUpdateResponseResultEditableTrue  ZoneSettingServerSideExcludeUpdateResponseResultEditable = true
	ZoneSettingServerSideExcludeUpdateResponseResultEditableFalse ZoneSettingServerSideExcludeUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingServerSideExcludeUpdateResponseResultValue string

const (
	ZoneSettingServerSideExcludeUpdateResponseResultValueOn  ZoneSettingServerSideExcludeUpdateResponseResultValue = "on"
	ZoneSettingServerSideExcludeUpdateResponseResultValueOff ZoneSettingServerSideExcludeUpdateResponseResultValue = "off"
)

type ZoneSettingServerSideExcludeListResponse struct {
	Errors   []ZoneSettingServerSideExcludeListResponseError   `json:"errors"`
	Messages []ZoneSettingServerSideExcludeListResponseMessage `json:"messages"`
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
	Result ZoneSettingServerSideExcludeListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingServerSideExcludeListResponseJSON
}

// zoneSettingServerSideExcludeListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingServerSideExcludeListResponse]
type zoneSettingServerSideExcludeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingServerSideExcludeListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingServerSideExcludeListResponseErrorJSON
}

// zoneSettingServerSideExcludeListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingServerSideExcludeListResponseError]
type zoneSettingServerSideExcludeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingServerSideExcludeListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingServerSideExcludeListResponseMessageJSON
}

// zoneSettingServerSideExcludeListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingServerSideExcludeListResponseMessage]
type zoneSettingServerSideExcludeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type ZoneSettingServerSideExcludeListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingServerSideExcludeListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingServerSideExcludeListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingServerSideExcludeListResponseResultValue `json:"value"`
	JSON  zoneSettingServerSideExcludeListResponseResultJSON
}

// zoneSettingServerSideExcludeListResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingServerSideExcludeListResponseResult]
type zoneSettingServerSideExcludeListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingServerSideExcludeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingServerSideExcludeListResponseResultID string

const (
	ZoneSettingServerSideExcludeListResponseResultIDServerSideExclude ZoneSettingServerSideExcludeListResponseResultID = "server_side_exclude"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingServerSideExcludeListResponseResultEditable bool

const (
	ZoneSettingServerSideExcludeListResponseResultEditableTrue  ZoneSettingServerSideExcludeListResponseResultEditable = true
	ZoneSettingServerSideExcludeListResponseResultEditableFalse ZoneSettingServerSideExcludeListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingServerSideExcludeListResponseResultValue string

const (
	ZoneSettingServerSideExcludeListResponseResultValueOn  ZoneSettingServerSideExcludeListResponseResultValue = "on"
	ZoneSettingServerSideExcludeListResponseResultValueOff ZoneSettingServerSideExcludeListResponseResultValue = "off"
)

type ZoneSettingServerSideExcludeUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingServerSideExcludeUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingServerSideExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingServerSideExcludeUpdateParamsValue string

const (
	ZoneSettingServerSideExcludeUpdateParamsValueOn  ZoneSettingServerSideExcludeUpdateParamsValue = "on"
	ZoneSettingServerSideExcludeUpdateParamsValueOff ZoneSettingServerSideExcludeUpdateParamsValue = "off"
)
