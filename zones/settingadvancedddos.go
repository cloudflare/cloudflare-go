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
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingAdvancedDDoSService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAdvancedDDoSService] method instead.
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
func (r *SettingAdvancedDDoSService) Get(ctx context.Context, query SettingAdvancedDDoSGetParams, opts ...option.RequestOption) (res *AdvancedDDoS, err error) {
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
type AdvancedDDoS struct {
	// ID of the zone setting.
	ID AdvancedDDoSID `json:"id,required"`
	// Current value of the zone setting.
	Value AdvancedDDoSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AdvancedDDoSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,nullable" format:"date-time"`
	JSON       advancedDDoSJSON `json:"-"`
}

// advancedDDoSJSON contains the JSON metadata for the struct [AdvancedDDoS]
type advancedDDoSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedDDoS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedDDoSJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type AdvancedDDoSID string

const (
	AdvancedDDoSIDAdvancedDDoS AdvancedDDoSID = "advanced_ddos"
)

func (r AdvancedDDoSID) IsKnown() bool {
	switch r {
	case AdvancedDDoSIDAdvancedDDoS:
		return true
	}
	return false
}

// Current value of the zone setting.
type AdvancedDDoSValue string

const (
	AdvancedDDoSValueOn  AdvancedDDoSValue = "on"
	AdvancedDDoSValueOff AdvancedDDoSValue = "off"
)

func (r AdvancedDDoSValue) IsKnown() bool {
	switch r {
	case AdvancedDDoSValueOn, AdvancedDDoSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type AdvancedDDoSEditable bool

const (
	AdvancedDDoSEditableTrue  AdvancedDDoSEditable = true
	AdvancedDDoSEditableFalse AdvancedDDoSEditable = false
)

func (r AdvancedDDoSEditable) IsKnown() bool {
	switch r {
	case AdvancedDDoSEditableTrue, AdvancedDDoSEditableFalse:
		return true
	}
	return false
}

type SettingAdvancedDDoSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingAdvancedDDoSGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result AdvancedDDoS                               `json:"result"`
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
