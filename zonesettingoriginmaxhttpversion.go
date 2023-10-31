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

// The highest HTTP version Cloudflare will attempt to use with your origin. This
// setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.).
func (r *ZoneSettingOriginMaxHTTPVersionService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingOriginMaxHTTPVersionUpdateParams, opts ...option.RequestOption) (res *ZoneSettingOriginMaxHTTPVersionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// The highest HTTP version Cloudflare will attempt to use with your origin. This
// setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.).
func (r *ZoneSettingOriginMaxHTTPVersionService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingOriginMaxHTTPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingOriginMaxHTTPVersionUpdateResponse struct {
	Errors   []ZoneSettingOriginMaxHTTPVersionUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingOriginMaxHTTPVersionUpdateResponseMessage `json:"messages"`
	Result   ZoneSettingOriginMaxHTTPVersionUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingOriginMaxHTTPVersionUpdateResponseJSON
}

// zoneSettingOriginMaxHTTPVersionUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionUpdateResponse]
type zoneSettingOriginMaxHTTPVersionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionUpdateResponseErrorJSON
}

// zoneSettingOriginMaxHTTPVersionUpdateResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingOriginMaxHTTPVersionUpdateResponseError]
type zoneSettingOriginMaxHTTPVersionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionUpdateResponseMessageJSON
}

// zoneSettingOriginMaxHTTPVersionUpdateResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingOriginMaxHTTPVersionUpdateResponseMessage]
type zoneSettingOriginMaxHTTPVersionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionUpdateResponseResult struct {
	// Identifier of the zone setting.
	ID ZoneSettingOriginMaxHTTPVersionUpdateResponseResultID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginMaxHTTPVersionUpdateResponseResultJSON
}

// zoneSettingOriginMaxHTTPVersionUpdateResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingOriginMaxHTTPVersionUpdateResponseResult]
type zoneSettingOriginMaxHTTPVersionUpdateResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type ZoneSettingOriginMaxHTTPVersionUpdateResponseResultID string

const (
	ZoneSettingOriginMaxHTTPVersionUpdateResponseResultIDOriginMaxHTTPVersion ZoneSettingOriginMaxHTTPVersionUpdateResponseResultID = "origin_max_http_version"
)

type ZoneSettingOriginMaxHTTPVersionListResponse struct {
	Errors   []ZoneSettingOriginMaxHTTPVersionListResponseError   `json:"errors"`
	Messages []ZoneSettingOriginMaxHTTPVersionListResponseMessage `json:"messages"`
	Result   ZoneSettingOriginMaxHTTPVersionListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingOriginMaxHTTPVersionListResponseJSON
}

// zoneSettingOriginMaxHTTPVersionListResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionListResponse]
type zoneSettingOriginMaxHTTPVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionListResponseErrorJSON
}

// zoneSettingOriginMaxHTTPVersionListResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingOriginMaxHTTPVersionListResponseError]
type zoneSettingOriginMaxHTTPVersionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOriginMaxHTTPVersionListResponseMessageJSON
}

// zoneSettingOriginMaxHTTPVersionListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingOriginMaxHTTPVersionListResponseMessage]
type zoneSettingOriginMaxHTTPVersionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOriginMaxHTTPVersionListResponseResult struct {
	// Identifier of the zone setting.
	ID ZoneSettingOriginMaxHTTPVersionListResponseResultID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginMaxHTTPVersionListResponseResultJSON
}

// zoneSettingOriginMaxHTTPVersionListResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingOriginMaxHTTPVersionListResponseResult]
type zoneSettingOriginMaxHTTPVersionListResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of the zone setting.
type ZoneSettingOriginMaxHTTPVersionListResponseResultID string

const (
	ZoneSettingOriginMaxHTTPVersionListResponseResultIDOriginMaxHTTPVersion ZoneSettingOriginMaxHTTPVersionListResponseResultID = "origin_max_http_version"
)

type ZoneSettingOriginMaxHTTPVersionUpdateParams struct {
	Value param.Field[ZoneSettingOriginMaxHTTPVersionUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingOriginMaxHTTPVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingOriginMaxHTTPVersionUpdateParamsValue struct {
	// Identifier of the zone setting.
	ID param.Field[ZoneSettingOriginMaxHTTPVersionUpdateParamsValueID] `json:"id,required"`
}

func (r ZoneSettingOriginMaxHTTPVersionUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Identifier of the zone setting.
type ZoneSettingOriginMaxHTTPVersionUpdateParamsValueID string

const (
	ZoneSettingOriginMaxHTTPVersionUpdateParamsValueIDOriginMaxHTTPVersion ZoneSettingOriginMaxHTTPVersionUpdateParamsValueID = "origin_max_http_version"
)
