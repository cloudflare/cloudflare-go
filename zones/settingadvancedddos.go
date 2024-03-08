// File generated from our OpenAPI spec by Stainless.

package zones

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
func (r *SettingAdvancedDDOSService) Get(ctx context.Context, query SettingAdvancedDDOSGetParams, opts ...option.RequestOption) (res *ZonesAdvancedDDOS, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAdvancedDDOSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/advanced_ddos", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID ZonesAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesAdvancedDDOSJSON `json:"-"`
}

// zonesAdvancedDDOSJSON contains the JSON metadata for the struct
// [ZonesAdvancedDDOS]
type zonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesAdvancedDDOSJSON) RawJSON() string {
	return r.raw
}

func (r ZonesAdvancedDDOS) implementsZonesSettingEditResponse() {}

func (r ZonesAdvancedDDOS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesAdvancedDDOSID string

const (
	ZonesAdvancedDDOSIDAdvancedDDOS ZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type ZonesAdvancedDDOSValue string

const (
	ZonesAdvancedDDOSValueOn  ZonesAdvancedDDOSValue = "on"
	ZonesAdvancedDDOSValueOff ZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesAdvancedDDOSEditable bool

const (
	ZonesAdvancedDDOSEditableTrue  ZonesAdvancedDDOSEditable = true
	ZonesAdvancedDDOSEditableFalse ZonesAdvancedDDOSEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZonesAdvancedDDOSParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesAdvancedDDOSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesAdvancedDDOSValue] `json:"value,required"`
}

func (r ZonesAdvancedDDOSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesAdvancedDDOSParam) implementsZonesSettingEditParamsItem() {}

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
	Result ZonesAdvancedDDOS                          `json:"result"`
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
