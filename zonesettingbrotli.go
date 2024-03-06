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

// ZoneSettingBrotliService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingBrotliService] method
// instead.
type ZoneSettingBrotliService struct {
	Options []option.RequestOption
}

// NewZoneSettingBrotliService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingBrotliService(opts ...option.RequestOption) (r *ZoneSettingBrotliService) {
	r = &ZoneSettingBrotliService{}
	r.Options = opts
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *ZoneSettingBrotliService) Edit(ctx context.Context, params ZoneSettingBrotliEditParams, opts ...option.RequestOption) (res *ZoneSettingBrotliEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingBrotliEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/brotli", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *ZoneSettingBrotliService) Get(ctx context.Context, query ZoneSettingBrotliGetParams, opts ...option.RequestOption) (res *ZoneSettingBrotliGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingBrotliGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/brotli", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingBrotliEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingBrotliEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrotliEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrotliEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrotliEditResponseJSON `json:"-"`
}

// zoneSettingBrotliEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingBrotliEditResponse]
type zoneSettingBrotliEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingBrotliEditResponseID string

const (
	ZoneSettingBrotliEditResponseIDBrotli ZoneSettingBrotliEditResponseID = "brotli"
)

// Current value of the zone setting.
type ZoneSettingBrotliEditResponseValue string

const (
	ZoneSettingBrotliEditResponseValueOff ZoneSettingBrotliEditResponseValue = "off"
	ZoneSettingBrotliEditResponseValueOn  ZoneSettingBrotliEditResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrotliEditResponseEditable bool

const (
	ZoneSettingBrotliEditResponseEditableTrue  ZoneSettingBrotliEditResponseEditable = true
	ZoneSettingBrotliEditResponseEditableFalse ZoneSettingBrotliEditResponseEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingBrotliGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingBrotliGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrotliGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrotliGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrotliGetResponseJSON `json:"-"`
}

// zoneSettingBrotliGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingBrotliGetResponse]
type zoneSettingBrotliGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingBrotliGetResponseID string

const (
	ZoneSettingBrotliGetResponseIDBrotli ZoneSettingBrotliGetResponseID = "brotli"
)

// Current value of the zone setting.
type ZoneSettingBrotliGetResponseValue string

const (
	ZoneSettingBrotliGetResponseValueOff ZoneSettingBrotliGetResponseValue = "off"
	ZoneSettingBrotliGetResponseValueOn  ZoneSettingBrotliGetResponseValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrotliGetResponseEditable bool

const (
	ZoneSettingBrotliGetResponseEditableTrue  ZoneSettingBrotliGetResponseEditable = true
	ZoneSettingBrotliGetResponseEditableFalse ZoneSettingBrotliGetResponseEditable = false
)

type ZoneSettingBrotliEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingBrotliEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrotliEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingBrotliEditParamsValue string

const (
	ZoneSettingBrotliEditParamsValueOff ZoneSettingBrotliEditParamsValue = "off"
	ZoneSettingBrotliEditParamsValueOn  ZoneSettingBrotliEditParamsValue = "on"
)

type ZoneSettingBrotliEditResponseEnvelope struct {
	Errors   []ZoneSettingBrotliEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingBrotliEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZoneSettingBrotliEditResponse             `json:"result"`
	JSON   zoneSettingBrotliEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingBrotliEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingBrotliEditResponseEnvelope]
type zoneSettingBrotliEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingBrotliEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingBrotliEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingBrotliEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingBrotliEditResponseEnvelopeErrors]
type zoneSettingBrotliEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingBrotliEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingBrotliEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingBrotliEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingBrotliEditResponseEnvelopeMessages]
type zoneSettingBrotliEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingBrotliGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingBrotliGetResponseEnvelope struct {
	Errors   []ZoneSettingBrotliGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingBrotliGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZoneSettingBrotliGetResponse             `json:"result"`
	JSON   zoneSettingBrotliGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingBrotliGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingBrotliGetResponseEnvelope]
type zoneSettingBrotliGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingBrotliGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingBrotliGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingBrotliGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingBrotliGetResponseEnvelopeErrors]
type zoneSettingBrotliGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingBrotliGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingBrotliGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingBrotliGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingBrotliGetResponseEnvelopeMessages]
type zoneSettingBrotliGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrotliGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
