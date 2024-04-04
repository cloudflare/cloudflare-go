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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingAdvancedDDoSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingAdvancedDDoSService]
// method instead.
type SettingAdvancedDDoSService struct {
	Options []option.RequestOption
}

// NewSettingAdvancedDDoSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingAdvancedDDoSService(opts ...option.RequestOption) (r *SettingAdvancedDDoSService) {
	r = &SettingAdvancedDDoSService{}
	r.Options = opts
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
func (r *SettingAdvancedDDoSService) Get(ctx context.Context, query SettingAdvancedDDoSGetParams, opts ...option.RequestOption) (res *ZoneSettingAdvancedDDoS, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAdvancedDDoSGetResponseEnvelope
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
type ZoneSettingAdvancedDDoS struct {
	// ID of the zone setting.
	ID ZoneSettingAdvancedDDoSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAdvancedDDoSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAdvancedDDoSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAdvancedDDoSJSON `json:"-"`
}

// zoneSettingAdvancedDDoSJSON contains the JSON metadata for the struct
// [ZoneSettingAdvancedDDoS]
type zoneSettingAdvancedDDoSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDoS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAdvancedDDoSJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingAdvancedDDoS) implementsZonesSettingEditResponse() {}

func (r ZoneSettingAdvancedDDoS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingAdvancedDDoSID string

const (
	ZoneSettingAdvancedDDoSIDAdvancedDDoS ZoneSettingAdvancedDDoSID = "advanced_ddos"
)

func (r ZoneSettingAdvancedDDoSID) IsKnown() bool {
	switch r {
	case ZoneSettingAdvancedDDoSIDAdvancedDDoS:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingAdvancedDDoSValue string

const (
	ZoneSettingAdvancedDDoSValueOn  ZoneSettingAdvancedDDoSValue = "on"
	ZoneSettingAdvancedDDoSValueOff ZoneSettingAdvancedDDoSValue = "off"
)

func (r ZoneSettingAdvancedDDoSValue) IsKnown() bool {
	switch r {
	case ZoneSettingAdvancedDDoSValueOn, ZoneSettingAdvancedDDoSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAdvancedDDoSEditable bool

const (
	ZoneSettingAdvancedDDoSEditableTrue  ZoneSettingAdvancedDDoSEditable = true
	ZoneSettingAdvancedDDoSEditableFalse ZoneSettingAdvancedDDoSEditable = false
)

func (r ZoneSettingAdvancedDDoSEditable) IsKnown() bool {
	switch r {
	case ZoneSettingAdvancedDDoSEditableTrue, ZoneSettingAdvancedDDoSEditableFalse:
		return true
	}
	return false
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingAdvancedDDoSParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingAdvancedDDoSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingAdvancedDDoSValue] `json:"value,required"`
}

func (r ZoneSettingAdvancedDDoSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingAdvancedDDoSParam) implementsZonesSettingEditParamsItem() {}

type SettingAdvancedDDoSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAdvancedDDoSGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result ZoneSettingAdvancedDDoS                    `json:"result"`
	JSON   settingAdvancedDDoSGetResponseEnvelopeJSON `json:"-"`
}

// settingAdvancedDDoSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingAdvancedDDoSGetResponseEnvelope]
type settingAdvancedDDoSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAdvancedDDoSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAdvancedDDoSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
