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

// ZoneSettingOriginMaxHTTPVersionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingOriginMaxHTTPVersionService] method instead.
type ZoneSettingOriginMaxHTTPVersionService struct {
	Options []option.RequestOption
}

// NewZoneSettingOriginMaxHTTPVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingOriginMaxHTTPVersionService(opts ...option.RequestOption) (r *ZoneSettingOriginMaxHTTPVersionService) {
	r = &ZoneSettingOriginMaxHTTPVersionService{}
	r.Options = opts
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
func (r *ZoneSettingOriginMaxHTTPVersionService) Edit(ctx context.Context, params ZoneSettingOriginMaxHTTPVersionEditParams, opts ...option.RequestOption) (res *ZoneSettingOriginMaxHTTPVersionEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOriginMaxHTTPVersionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
func (r *ZoneSettingOriginMaxHTTPVersionService) Get(ctx context.Context, query ZoneSettingOriginMaxHTTPVersionGetParams, opts ...option.RequestOption) (res *ZoneSettingOriginMaxHTTPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOriginMaxHTTPVersionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
type ZoneSettingOriginMaxHTTPVersionEditResponse struct {
	// Value of the zone setting.
	ID ZoneSettingOriginMaxHTTPVersionEditResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Origin Max HTTP Version Setting.
	Value ZoneSettingOriginMaxHTTPVersionEditResponseValue `json:"value,required"`
	JSON  zoneSettingOriginMaxHTTPVersionEditResponseJSON  `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionEditResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionEditResponse]
type zoneSettingOriginMaxHTTPVersionEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Value of the zone setting.
type ZoneSettingOriginMaxHTTPVersionEditResponseID string

const (
	ZoneSettingOriginMaxHTTPVersionEditResponseIDOriginMaxHTTPVersion ZoneSettingOriginMaxHTTPVersionEditResponseID = "origin_max_http_version"
)

// Value of the Origin Max HTTP Version Setting.
type ZoneSettingOriginMaxHTTPVersionEditResponseValue string

const (
	ZoneSettingOriginMaxHTTPVersionEditResponseValue2 ZoneSettingOriginMaxHTTPVersionEditResponseValue = "2"
	ZoneSettingOriginMaxHTTPVersionEditResponseValue1 ZoneSettingOriginMaxHTTPVersionEditResponseValue = "1"
)

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except ENT
// where it is "1"
type ZoneSettingOriginMaxHTTPVersionGetResponse struct {
	// Value of the zone setting.
	ID ZoneSettingOriginMaxHTTPVersionGetResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Origin Max HTTP Version Setting.
	Value ZoneSettingOriginMaxHTTPVersionGetResponseValue `json:"value,required"`
	JSON  zoneSettingOriginMaxHTTPVersionGetResponseJSON  `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionGetResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionGetResponse]
type zoneSettingOriginMaxHTTPVersionGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Value of the zone setting.
type ZoneSettingOriginMaxHTTPVersionGetResponseID string

const (
	ZoneSettingOriginMaxHTTPVersionGetResponseIDOriginMaxHTTPVersion ZoneSettingOriginMaxHTTPVersionGetResponseID = "origin_max_http_version"
)

// Value of the Origin Max HTTP Version Setting.
type ZoneSettingOriginMaxHTTPVersionGetResponseValue string

const (
	ZoneSettingOriginMaxHTTPVersionGetResponseValue2 ZoneSettingOriginMaxHTTPVersionGetResponseValue = "2"
	ZoneSettingOriginMaxHTTPVersionGetResponseValue1 ZoneSettingOriginMaxHTTPVersionGetResponseValue = "1"
)

type ZoneSettingOriginMaxHTTPVersionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Origin Max HTTP Version Setting.
	Value param.Field[ZoneSettingOriginMaxHTTPVersionEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingOriginMaxHTTPVersionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Max HTTP Version Setting.
type ZoneSettingOriginMaxHTTPVersionEditParamsValue string

const (
	ZoneSettingOriginMaxHTTPVersionEditParamsValue2 ZoneSettingOriginMaxHTTPVersionEditParamsValue = "2"
	ZoneSettingOriginMaxHTTPVersionEditParamsValue1 ZoneSettingOriginMaxHTTPVersionEditParamsValue = "1"
)

type ZoneSettingOriginMaxHTTPVersionEditResponseEnvelope struct {
	Errors   []ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeMessages `json:"messages,required"`
	// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
	// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
	// requests to your origin. (Refer to
	// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
	// for more information.). The default value is "2" for all plan types except ENT
	// where it is "1"
	Result ZoneSettingOriginMaxHTTPVersionEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeJSON    `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingOriginMaxHTTPVersionEditResponseEnvelope]
type zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeErrors]
type zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeMessages]
type zoneSettingOriginMaxHTTPVersionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeSuccess bool

const (
	ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeSuccessTrue ZoneSettingOriginMaxHTTPVersionEditResponseEnvelopeSuccess = true
)

type ZoneSettingOriginMaxHTTPVersionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingOriginMaxHTTPVersionGetResponseEnvelope struct {
	Errors   []ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
	// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
	// requests to your origin. (Refer to
	// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
	// for more information.). The default value is "2" for all plan types except ENT
	// where it is "1"
	Result ZoneSettingOriginMaxHTTPVersionGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeJSON    `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingOriginMaxHTTPVersionGetResponseEnvelope]
type zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeErrors]
type zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeMessages]
type zoneSettingOriginMaxHTTPVersionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeSuccess bool

const (
	ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeSuccessTrue ZoneSettingOriginMaxHTTPVersionGetResponseEnvelopeSuccess = true
)
