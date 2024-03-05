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

// ZoneSettingPseudoIPV4Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingPseudoIPV4Service]
// method instead.
type ZoneSettingPseudoIPV4Service struct {
	Options []option.RequestOption
}

// NewZoneSettingPseudoIPV4Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingPseudoIPV4Service(opts ...option.RequestOption) (r *ZoneSettingPseudoIPV4Service) {
	r = &ZoneSettingPseudoIPV4Service{}
	r.Options = opts
	return
}

// Value of the Pseudo IPv4 setting.
func (r *ZoneSettingPseudoIPV4Service) Edit(ctx context.Context, params ZoneSettingPseudoIPV4EditParams, opts ...option.RequestOption) (res *ZonesPseudoIPV4, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingPseudoIPV4EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the Pseudo IPv4 setting.
func (r *ZoneSettingPseudoIPV4Service) Get(ctx context.Context, query ZoneSettingPseudoIPV4GetParams, opts ...option.RequestOption) (res *ZonesPseudoIPV4, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingPseudoIPV4GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The value set for the Pseudo IPv4 setting.
type ZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZonesPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesPseudoIPV4JSON `json:"-"`
}

// zonesPseudoIPV4JSON contains the JSON metadata for the struct [ZonesPseudoIPV4]
type zonesPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesPseudoIPV4) implementsZoneSettingEditResponse() {}

func (r ZonesPseudoIPV4) implementsZoneSettingGetResponse() {}

// Value of the Pseudo IPv4 setting.
type ZonesPseudoIPV4ID string

const (
	ZonesPseudoIPV4IDPseudoIPV4 ZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type ZonesPseudoIPV4Value string

const (
	ZonesPseudoIPV4ValueOff             ZonesPseudoIPV4Value = "off"
	ZonesPseudoIPV4ValueAddHeader       ZonesPseudoIPV4Value = "add_header"
	ZonesPseudoIPV4ValueOverwriteHeader ZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesPseudoIPV4Editable bool

const (
	ZonesPseudoIPV4EditableTrue  ZonesPseudoIPV4Editable = true
	ZonesPseudoIPV4EditableFalse ZonesPseudoIPV4Editable = false
)

// The value set for the Pseudo IPv4 setting.
type ZonesPseudoIPV4Param struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[ZonesPseudoIPV4ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesPseudoIPV4Value] `json:"value,required"`
}

func (r ZonesPseudoIPV4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesPseudoIPV4Param) implementsZoneSettingEditParamsItem() {}

type ZoneSettingPseudoIPV4EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Pseudo IPv4 setting.
	Value param.Field[ZoneSettingPseudoIPV4EditParamsValue] `json:"value,required"`
}

func (r ZoneSettingPseudoIPV4EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Pseudo IPv4 setting.
type ZoneSettingPseudoIPV4EditParamsValue string

const (
	ZoneSettingPseudoIPV4EditParamsValueOff             ZoneSettingPseudoIPV4EditParamsValue = "off"
	ZoneSettingPseudoIPV4EditParamsValueAddHeader       ZoneSettingPseudoIPV4EditParamsValue = "add_header"
	ZoneSettingPseudoIPV4EditParamsValueOverwriteHeader ZoneSettingPseudoIPV4EditParamsValue = "overwrite_header"
)

type ZoneSettingPseudoIPV4EditResponseEnvelope struct {
	Errors   []ZoneSettingPseudoIPV4EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingPseudoIPV4EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result ZonesPseudoIPV4                               `json:"result"`
	JSON   zoneSettingPseudoIPV4EditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingPseudoIPV4EditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIPV4EditResponseEnvelope]
type zoneSettingPseudoIPV4EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIPV4EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIPV4EditResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingPseudoIPV4EditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingPseudoIPV4EditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingPseudoIPV4EditResponseEnvelopeErrors]
type zoneSettingPseudoIPV4EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIPV4EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIPV4EditResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingPseudoIPV4EditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingPseudoIPV4EditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingPseudoIPV4EditResponseEnvelopeMessages]
type zoneSettingPseudoIPV4EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIPV4EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIPV4GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingPseudoIPV4GetResponseEnvelope struct {
	Errors   []ZoneSettingPseudoIPV4GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingPseudoIPV4GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// The value set for the Pseudo IPv4 setting.
	Result ZonesPseudoIPV4                              `json:"result"`
	JSON   zoneSettingPseudoIPV4GetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingPseudoIPV4GetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIPV4GetResponseEnvelope]
type zoneSettingPseudoIPV4GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIPV4GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIPV4GetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingPseudoIPV4GetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingPseudoIPV4GetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingPseudoIPV4GetResponseEnvelopeErrors]
type zoneSettingPseudoIPV4GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIPV4GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIPV4GetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingPseudoIPV4GetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingPseudoIPV4GetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingPseudoIPV4GetResponseEnvelopeMessages]
type zoneSettingPseudoIPV4GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIPV4GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
