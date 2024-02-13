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

// SettingResponseBufferingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingResponseBufferingService] method instead.
type SettingResponseBufferingService struct {
	Options []option.RequestOption
}

// NewSettingResponseBufferingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingResponseBufferingService(opts ...option.RequestOption) (r *SettingResponseBufferingService) {
	r = &SettingResponseBufferingService{}
	r.Options = opts
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *SettingResponseBufferingService) Update(ctx context.Context, zoneID string, body SettingResponseBufferingUpdateParams, opts ...option.RequestOption) (res *SettingResponseBufferingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingResponseBufferingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *SettingResponseBufferingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingResponseBufferingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingResponseBufferingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/response_buffering", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingResponseBufferingUpdateResponse struct {
	// ID of the zone setting.
	ID SettingResponseBufferingUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingResponseBufferingUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingResponseBufferingUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingResponseBufferingUpdateResponseJSON `json:"-"`
}

// settingResponseBufferingUpdateResponseJSON contains the JSON metadata for the
// struct [SettingResponseBufferingUpdateResponse]
type settingResponseBufferingUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingResponseBufferingUpdateResponseID string

const (
	SettingResponseBufferingUpdateResponseIDResponseBuffering SettingResponseBufferingUpdateResponseID = "response_buffering"
)

// Current value of the zone setting.
type SettingResponseBufferingUpdateResponseValue string

const (
	SettingResponseBufferingUpdateResponseValueOn  SettingResponseBufferingUpdateResponseValue = "on"
	SettingResponseBufferingUpdateResponseValueOff SettingResponseBufferingUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingResponseBufferingUpdateResponseEditable bool

const (
	SettingResponseBufferingUpdateResponseEditableTrue  SettingResponseBufferingUpdateResponseEditable = true
	SettingResponseBufferingUpdateResponseEditableFalse SettingResponseBufferingUpdateResponseEditable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingResponseBufferingGetResponse struct {
	// ID of the zone setting.
	ID SettingResponseBufferingGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingResponseBufferingGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingResponseBufferingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingResponseBufferingGetResponseJSON `json:"-"`
}

// settingResponseBufferingGetResponseJSON contains the JSON metadata for the
// struct [SettingResponseBufferingGetResponse]
type settingResponseBufferingGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingResponseBufferingGetResponseID string

const (
	SettingResponseBufferingGetResponseIDResponseBuffering SettingResponseBufferingGetResponseID = "response_buffering"
)

// Current value of the zone setting.
type SettingResponseBufferingGetResponseValue string

const (
	SettingResponseBufferingGetResponseValueOn  SettingResponseBufferingGetResponseValue = "on"
	SettingResponseBufferingGetResponseValueOff SettingResponseBufferingGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingResponseBufferingGetResponseEditable bool

const (
	SettingResponseBufferingGetResponseEditableTrue  SettingResponseBufferingGetResponseEditable = true
	SettingResponseBufferingGetResponseEditableFalse SettingResponseBufferingGetResponseEditable = false
)

type SettingResponseBufferingUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingResponseBufferingUpdateParamsValue] `json:"value,required"`
}

func (r SettingResponseBufferingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingResponseBufferingUpdateParamsValue string

const (
	SettingResponseBufferingUpdateParamsValueOn  SettingResponseBufferingUpdateParamsValue = "on"
	SettingResponseBufferingUpdateParamsValueOff SettingResponseBufferingUpdateParamsValue = "off"
)

type SettingResponseBufferingUpdateResponseEnvelope struct {
	Errors   []SettingResponseBufferingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingResponseBufferingUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result SettingResponseBufferingUpdateResponse             `json:"result"`
	JSON   settingResponseBufferingUpdateResponseEnvelopeJSON `json:"-"`
}

// settingResponseBufferingUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [SettingResponseBufferingUpdateResponseEnvelope]
type settingResponseBufferingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingResponseBufferingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    settingResponseBufferingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingResponseBufferingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingResponseBufferingUpdateResponseEnvelopeErrors]
type settingResponseBufferingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingResponseBufferingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    settingResponseBufferingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingResponseBufferingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingResponseBufferingUpdateResponseEnvelopeMessages]
type settingResponseBufferingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingResponseBufferingGetResponseEnvelope struct {
	Errors   []SettingResponseBufferingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingResponseBufferingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result SettingResponseBufferingGetResponse             `json:"result"`
	JSON   settingResponseBufferingGetResponseEnvelopeJSON `json:"-"`
}

// settingResponseBufferingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingResponseBufferingGetResponseEnvelope]
type settingResponseBufferingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingResponseBufferingGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingResponseBufferingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingResponseBufferingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingResponseBufferingGetResponseEnvelopeErrors]
type settingResponseBufferingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingResponseBufferingGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingResponseBufferingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingResponseBufferingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingResponseBufferingGetResponseEnvelopeMessages]
type settingResponseBufferingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingResponseBufferingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
