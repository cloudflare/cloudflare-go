// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *SettingAdvancedDDOSService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingAdvancedDDOSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingAdvancedDDOSGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/advanced_ddos", zoneID)
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
type SettingAdvancedDDOSGetResponse struct {
	// ID of the zone setting.
	ID SettingAdvancedDDOSGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingAdvancedDDOSGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingAdvancedDDOSGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingAdvancedDDOSGetResponseJSON `json:"-"`
}

// settingAdvancedDDOSGetResponseJSON contains the JSON metadata for the struct
// [SettingAdvancedDDOSGetResponse]
type settingAdvancedDDOSGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAdvancedDDOSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingAdvancedDDOSGetResponseID string

const (
	SettingAdvancedDDOSGetResponseIDAdvancedDDOS SettingAdvancedDDOSGetResponseID = "advanced_ddos"
)

// Current value of the zone setting.
type SettingAdvancedDDOSGetResponseValue string

const (
	SettingAdvancedDDOSGetResponseValueOn  SettingAdvancedDDOSGetResponseValue = "on"
	SettingAdvancedDDOSGetResponseValueOff SettingAdvancedDDOSGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingAdvancedDDOSGetResponseEditable bool

const (
	SettingAdvancedDDOSGetResponseEditableTrue  SettingAdvancedDDOSGetResponseEditable = true
	SettingAdvancedDDOSGetResponseEditableFalse SettingAdvancedDDOSGetResponseEditable = false
)

type SettingAdvancedDDOSGetResponseEnvelope struct {
	Errors   []SettingAdvancedDDOSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAdvancedDDOSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result SettingAdvancedDDOSGetResponse             `json:"result"`
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
