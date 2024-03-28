// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingAdvancedDDOSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAdvancedDDOSService]
// method instead.
type SettingAdvancedDDOSService struct {
	Options []option.RequestOption
}

// NewSettingAdvancedDDOSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAdvancedDDOSService(opts ...option.RequestOption) (r *SettingAdvancedDDOSService) {
	r = &SettingAdvancedDDOSService{}
	r.Options = opts
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
func (r *SettingAdvancedDDOSService) Get(ctx context.Context, query SettingAdvancedDDOSGetParams, opts ...option.RequestOption) (res *ZoneSettingAdvancedDDOS, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAdvancedDDOSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/advanced_ddos", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingAdvancedDDOS struct {
	// ID of the zone setting.
	ID ZoneSettingAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAdvancedDDOSJSON `json:"-"`
}

// zoneSettingAdvancedDDOSJSON contains the JSON metadata for the struct
// [ZoneSettingAdvancedDDOS]
type zoneSettingAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAdvancedDDOSJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingAdvancedDDOS) implementsZonesSettingEditResponse() {}

func (r ZoneSettingAdvancedDDOS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingAdvancedDDOSID string

const (
	ZoneSettingAdvancedDDOSIDAdvancedDDOS ZoneSettingAdvancedDDOSID = "advanced_ddos"
)

func (r ZoneSettingAdvancedDDOSID) IsKnown() bool {
	switch r {
	case ZoneSettingAdvancedDDOSIDAdvancedDDOS:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingAdvancedDDOSValue string

const (
	ZoneSettingAdvancedDDOSValueOn  ZoneSettingAdvancedDDOSValue = "on"
	ZoneSettingAdvancedDDOSValueOff ZoneSettingAdvancedDDOSValue = "off"
)

func (r ZoneSettingAdvancedDDOSValue) IsKnown() bool {
	switch r {
	case ZoneSettingAdvancedDDOSValueOn, ZoneSettingAdvancedDDOSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAdvancedDDOSEditable bool

const (
	ZoneSettingAdvancedDDOSEditableTrue  ZoneSettingAdvancedDDOSEditable = true
	ZoneSettingAdvancedDDOSEditableFalse ZoneSettingAdvancedDDOSEditable = false
)

func (r ZoneSettingAdvancedDDOSEditable) IsKnown() bool {
	switch r {
	case ZoneSettingAdvancedDDOSEditableTrue, ZoneSettingAdvancedDDOSEditableFalse:
		return true
	}
	return false
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingAdvancedDDOSParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingAdvancedDDOSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingAdvancedDDOSValue] `json:"value,required"`
}

func (r ZoneSettingAdvancedDDOSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingAdvancedDDOSParam) implementsZonesSettingEditParamsItem() {}

type SettingAdvancedDDOSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAdvancedDDOSGetResponseEnvelope struct {
	Errors   []SettingAdvancedDDOSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAdvancedDDOSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result ZoneSettingAdvancedDDOS                    `json:"result"`
	JSON   settingAdvancedDDOSGetResponseEnvelopeJSON `json:"-"`
}

// settingAdvancedDDOSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAdvancedDDOSGetResponseEnvelope]
type settingAdvancedDDOSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAdvancedDDOSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAdvancedDDOSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAdvancedDDOSGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingAdvancedDDOSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingAdvancedDDOSGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingAdvancedDDOSGetResponseEnvelopeErrors]
type settingAdvancedDDOSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAdvancedDDOSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAdvancedDDOSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAdvancedDDOSGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingAdvancedDDOSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingAdvancedDDOSGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingAdvancedDDOSGetResponseEnvelopeMessages]
type settingAdvancedDDOSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAdvancedDDOSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAdvancedDDOSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
