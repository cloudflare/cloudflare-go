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

// ZoneSettingAlwaysUseHTTPSService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingAlwaysUseHTTPSService] method instead.
type ZoneSettingAlwaysUseHTTPSService struct {
	Options []option.RequestOption
}

// NewZoneSettingAlwaysUseHTTPSService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingAlwaysUseHTTPSService(opts ...option.RequestOption) (r *ZoneSettingAlwaysUseHTTPSService) {
	r = &ZoneSettingAlwaysUseHTTPSService{}
	r.Options = opts
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *ZoneSettingAlwaysUseHTTPSService) Edit(ctx context.Context, params ZoneSettingAlwaysUseHTTPSEditParams, opts ...option.RequestOption) (res *ZonesAlwaysUseHTTPS, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAlwaysUseHTTPSEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *ZoneSettingAlwaysUseHTTPSService) Get(ctx context.Context, query ZoneSettingAlwaysUseHTTPSGetParams, opts ...option.RequestOption) (res *ZonesAlwaysUseHTTPS, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAlwaysUseHTTPSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/always_use_https", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID ZonesAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time               `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesAlwaysUseHTTPSJSON `json:"-"`
}

// zonesAlwaysUseHTTPSJSON contains the JSON metadata for the struct
// [ZonesAlwaysUseHTTPS]
type zonesAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesAlwaysUseHTTPS) implementsZoneSettingEditResponse() {}

func (r ZonesAlwaysUseHTTPS) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesAlwaysUseHTTPSID string

const (
	ZonesAlwaysUseHTTPSIDAlwaysUseHTTPS ZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type ZonesAlwaysUseHTTPSValue string

const (
	ZonesAlwaysUseHTTPSValueOn  ZonesAlwaysUseHTTPSValue = "on"
	ZonesAlwaysUseHTTPSValueOff ZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesAlwaysUseHTTPSEditable bool

const (
	ZonesAlwaysUseHTTPSEditableTrue  ZonesAlwaysUseHTTPSEditable = true
	ZonesAlwaysUseHTTPSEditableFalse ZonesAlwaysUseHTTPSEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZonesAlwaysUseHTTPSParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesAlwaysUseHTTPSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesAlwaysUseHTTPSValue] `json:"value,required"`
}

func (r ZonesAlwaysUseHTTPSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesAlwaysUseHTTPSParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingAlwaysUseHTTPSEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingAlwaysUseHTTPSEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingAlwaysUseHTTPSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingAlwaysUseHTTPSEditParamsValue string

const (
	ZoneSettingAlwaysUseHTTPSEditParamsValueOn  ZoneSettingAlwaysUseHTTPSEditParamsValue = "on"
	ZoneSettingAlwaysUseHTTPSEditParamsValueOff ZoneSettingAlwaysUseHTTPSEditParamsValue = "off"
)

type ZoneSettingAlwaysUseHTTPSEditResponseEnvelope struct {
	Errors   []ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result ZonesAlwaysUseHTTPS                               `json:"result"`
	JSON   zoneSettingAlwaysUseHTTPSEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPSEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysUseHTTPSEditResponseEnvelope]
type zoneSettingAlwaysUseHTTPSEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeErrors]
type zoneSettingAlwaysUseHTTPSEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeMessages]
type zoneSettingAlwaysUseHTTPSEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPSEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingAlwaysUseHTTPSGetResponseEnvelope struct {
	Errors   []ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result ZonesAlwaysUseHTTPS                              `json:"result"`
	JSON   zoneSettingAlwaysUseHTTPSGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPSGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingAlwaysUseHTTPSGetResponseEnvelope]
type zoneSettingAlwaysUseHTTPSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeErrors]
type zoneSettingAlwaysUseHTTPSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeMessages]
type zoneSettingAlwaysUseHTTPSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAlwaysUseHTTPSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
