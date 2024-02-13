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

// SettingServerSideExcludeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingServerSideExcludeService] method instead.
type SettingServerSideExcludeService struct {
	Options []option.RequestOption
}

// NewSettingServerSideExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingServerSideExcludeService(opts ...option.RequestOption) (r *SettingServerSideExcludeService) {
	r = &SettingServerSideExcludeService{}
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
func (r *SettingServerSideExcludeService) Update(ctx context.Context, zoneID string, body SettingServerSideExcludeUpdateParams, opts ...option.RequestOption) (res *SettingServerSideExcludeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingServerSideExcludeUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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
func (r *SettingServerSideExcludeService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingServerSideExcludeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingServerSideExcludeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type SettingServerSideExcludeUpdateResponse struct {
	// ID of the zone setting.
	ID SettingServerSideExcludeUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingServerSideExcludeUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingServerSideExcludeUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingServerSideExcludeUpdateResponseJSON `json:"-"`
}

// settingServerSideExcludeUpdateResponseJSON contains the JSON metadata for the
// struct [SettingServerSideExcludeUpdateResponse]
type settingServerSideExcludeUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingServerSideExcludeUpdateResponseID string

const (
	SettingServerSideExcludeUpdateResponseIDServerSideExclude SettingServerSideExcludeUpdateResponseID = "server_side_exclude"
)

// Current value of the zone setting.
type SettingServerSideExcludeUpdateResponseValue string

const (
	SettingServerSideExcludeUpdateResponseValueOn  SettingServerSideExcludeUpdateResponseValue = "on"
	SettingServerSideExcludeUpdateResponseValueOff SettingServerSideExcludeUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingServerSideExcludeUpdateResponseEditable bool

const (
	SettingServerSideExcludeUpdateResponseEditableTrue  SettingServerSideExcludeUpdateResponseEditable = true
	SettingServerSideExcludeUpdateResponseEditableFalse SettingServerSideExcludeUpdateResponseEditable = false
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
type SettingServerSideExcludeGetResponse struct {
	// ID of the zone setting.
	ID SettingServerSideExcludeGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingServerSideExcludeGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingServerSideExcludeGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingServerSideExcludeGetResponseJSON `json:"-"`
}

// settingServerSideExcludeGetResponseJSON contains the JSON metadata for the
// struct [SettingServerSideExcludeGetResponse]
type settingServerSideExcludeGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingServerSideExcludeGetResponseID string

const (
	SettingServerSideExcludeGetResponseIDServerSideExclude SettingServerSideExcludeGetResponseID = "server_side_exclude"
)

// Current value of the zone setting.
type SettingServerSideExcludeGetResponseValue string

const (
	SettingServerSideExcludeGetResponseValueOn  SettingServerSideExcludeGetResponseValue = "on"
	SettingServerSideExcludeGetResponseValueOff SettingServerSideExcludeGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingServerSideExcludeGetResponseEditable bool

const (
	SettingServerSideExcludeGetResponseEditableTrue  SettingServerSideExcludeGetResponseEditable = true
	SettingServerSideExcludeGetResponseEditableFalse SettingServerSideExcludeGetResponseEditable = false
)

type SettingServerSideExcludeUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingServerSideExcludeUpdateParamsValue] `json:"value,required"`
}

func (r SettingServerSideExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingServerSideExcludeUpdateParamsValue string

const (
	SettingServerSideExcludeUpdateParamsValueOn  SettingServerSideExcludeUpdateParamsValue = "on"
	SettingServerSideExcludeUpdateParamsValueOff SettingServerSideExcludeUpdateParamsValue = "off"
)

type SettingServerSideExcludeUpdateResponseEnvelope struct {
	Errors   []SettingServerSideExcludeUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingServerSideExcludeUpdateResponseEnvelopeMessages `json:"messages,required"`
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
	Result SettingServerSideExcludeUpdateResponse             `json:"result"`
	JSON   settingServerSideExcludeUpdateResponseEnvelopeJSON `json:"-"`
}

// settingServerSideExcludeUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingServerSideExcludeUpdateResponseEnvelope]
type settingServerSideExcludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingServerSideExcludeUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingServerSideExcludeUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingServerSideExcludeUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingServerSideExcludeUpdateResponseEnvelopeErrors]
type settingServerSideExcludeUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingServerSideExcludeUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    settingServerSideExcludeUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingServerSideExcludeUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingServerSideExcludeUpdateResponseEnvelopeMessages]
type settingServerSideExcludeUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingServerSideExcludeGetResponseEnvelope struct {
	Errors   []SettingServerSideExcludeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingServerSideExcludeGetResponseEnvelopeMessages `json:"messages,required"`
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
	Result SettingServerSideExcludeGetResponse             `json:"result"`
	JSON   settingServerSideExcludeGetResponseEnvelopeJSON `json:"-"`
}

// settingServerSideExcludeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingServerSideExcludeGetResponseEnvelope]
type settingServerSideExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingServerSideExcludeGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingServerSideExcludeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingServerSideExcludeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingServerSideExcludeGetResponseEnvelopeErrors]
type settingServerSideExcludeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingServerSideExcludeGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingServerSideExcludeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingServerSideExcludeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingServerSideExcludeGetResponseEnvelopeMessages]
type settingServerSideExcludeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingServerSideExcludeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
