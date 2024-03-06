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

// ZoneSettingAdvancedDDOSService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingAdvancedDDOSService] method instead.
type ZoneSettingAdvancedDDOSService struct {
	Options []option.RequestOption
}

// NewZoneSettingAdvancedDDOSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingAdvancedDDOSService(opts ...option.RequestOption) (r *ZoneSettingAdvancedDDOSService) {
	r = &ZoneSettingAdvancedDDOSService{}
	r.Options = opts
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
func (r *ZoneSettingAdvancedDDOSService) Get(ctx context.Context, query ZoneSettingAdvancedDDOSGetParams, opts ...option.RequestOption) (res *ZoneSettingAdvancedDDOSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingAdvancedDDOSGetResponseEnvelope
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
type ZoneSettingAdvancedDDOSGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingAdvancedDDOSGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAdvancedDDOSGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAdvancedDDOSGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAdvancedDDOSGetResponseJSON `json:"-"`
}

// zoneSettingAdvancedDDOSGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAdvancedDDOSGetResponse]
type zoneSettingAdvancedDDOSGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAdvancedDDOSGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingAdvancedDDOSGetResponseID string

const (
	ZoneSettingAdvancedDDOSGetResponseIDAdvancedDDOS ZoneSettingAdvancedDDOSGetResponseID = "advanced_ddos"
)

// Current value of the zone setting.
type ZoneSettingAdvancedDDOSGetResponseValue string

const (
	ZoneSettingAdvancedDDOSGetResponseValueOn  ZoneSettingAdvancedDDOSGetResponseValue = "on"
	ZoneSettingAdvancedDDOSGetResponseValueOff ZoneSettingAdvancedDDOSGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAdvancedDDOSGetResponseEditable bool

const (
	ZoneSettingAdvancedDDOSGetResponseEditableTrue  ZoneSettingAdvancedDDOSGetResponseEditable = true
	ZoneSettingAdvancedDDOSGetResponseEditableFalse ZoneSettingAdvancedDDOSGetResponseEditable = false
)

type ZoneSettingAdvancedDDOSGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingAdvancedDDOSGetResponseEnvelope struct {
	Errors   []ZoneSettingAdvancedDDOSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingAdvancedDDOSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result ZoneSettingAdvancedDDOSGetResponse             `json:"result"`
	JSON   zoneSettingAdvancedDDOSGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingAdvancedDDOSGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingAdvancedDDOSGetResponseEnvelope]
type zoneSettingAdvancedDDOSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAdvancedDDOSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAdvancedDDOSGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingAdvancedDDOSGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingAdvancedDDOSGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingAdvancedDDOSGetResponseEnvelopeErrors]
type zoneSettingAdvancedDDOSGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAdvancedDDOSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAdvancedDDOSGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingAdvancedDDOSGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingAdvancedDDOSGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingAdvancedDDOSGetResponseEnvelopeMessages]
type zoneSettingAdvancedDDOSGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAdvancedDDOSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
