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

// ZoneSettingAdvancedDdoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingAdvancedDdoService]
// method instead.
type ZoneSettingAdvancedDdoService struct {
	Options []option.RequestOption
}

// NewZoneSettingAdvancedDdoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingAdvancedDdoService(opts ...option.RequestOption) (r *ZoneSettingAdvancedDdoService) {
	r = &ZoneSettingAdvancedDdoService{}
	r.Options = opts
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
func (r *ZoneSettingAdvancedDdoService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingAdvancedDdoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/advanced_ddos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingAdvancedDdoListResponse struct {
	Errors   []ZoneSettingAdvancedDdoListResponseError   `json:"errors,required"`
	Messages []ZoneSettingAdvancedDdoListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                     `json:"success,required"`
	Result  ZoneSettingAdvancedDdoListResponseResult `json:"result"`
	JSON    zoneSettingAdvancedDdoListResponseJSON   `json:"-"`
}

// zoneSettingAdvancedDdoListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAdvancedDdoListResponse]
type zoneSettingAdvancedDdoListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDdoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAdvancedDdoListResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingAdvancedDdoListResponseErrorJSON `json:"-"`
}

// zoneSettingAdvancedDdoListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingAdvancedDdoListResponseError]
type zoneSettingAdvancedDdoListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDdoListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAdvancedDdoListResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingAdvancedDdoListResponseMessageJSON `json:"-"`
}

// zoneSettingAdvancedDdoListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingAdvancedDdoListResponseMessage]
type zoneSettingAdvancedDdoListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDdoListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAdvancedDdoListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAdvancedDdoListResponseResultID `json:"id,required"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZoneSettingAdvancedDdoListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAdvancedDdoListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAdvancedDdoListResponseResultJSON `json:"-"`
}

// zoneSettingAdvancedDdoListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingAdvancedDdoListResponseResult]
type zoneSettingAdvancedDdoListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDdoListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAdvancedDdoListResponseResultID string

const (
	ZoneSettingAdvancedDdoListResponseResultIDAdvancedDdos ZoneSettingAdvancedDdoListResponseResultID = "advanced_ddos"
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingAdvancedDdoListResponseResultValue string

const (
	ZoneSettingAdvancedDdoListResponseResultValueOn  ZoneSettingAdvancedDdoListResponseResultValue = "on"
	ZoneSettingAdvancedDdoListResponseResultValueOff ZoneSettingAdvancedDdoListResponseResultValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAdvancedDdoListResponseResultEditable bool

const (
	ZoneSettingAdvancedDdoListResponseResultEditableTrue  ZoneSettingAdvancedDdoListResponseResultEditable = true
	ZoneSettingAdvancedDdoListResponseResultEditableFalse ZoneSettingAdvancedDdoListResponseResultEditable = false
)
