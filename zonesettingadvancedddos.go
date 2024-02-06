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
func (r *ZoneSettingAdvancedDDOSService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingAdvancedDDOSListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/advanced_ddos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingAdvancedDDOSListResponse struct {
	Errors   []ZoneSettingAdvancedDDOSListResponseError   `json:"errors"`
	Messages []ZoneSettingAdvancedDDOSListResponseMessage `json:"messages"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result ZoneSettingAdvancedDDOSListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                    `json:"success"`
	JSON    zoneSettingAdvancedDDOSListResponseJSON `json:"-"`
}

// zoneSettingAdvancedDDOSListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingAdvancedDDOSListResponse]
type zoneSettingAdvancedDDOSListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAdvancedDDOSListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingAdvancedDDOSListResponseErrorJSON `json:"-"`
}

// zoneSettingAdvancedDDOSListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingAdvancedDDOSListResponseError]
type zoneSettingAdvancedDDOSListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingAdvancedDDOSListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingAdvancedDDOSListResponseMessageJSON `json:"-"`
}

// zoneSettingAdvancedDDOSListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingAdvancedDDOSListResponseMessage]
type zoneSettingAdvancedDDOSListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingAdvancedDDOSListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAdvancedDDOSListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingAdvancedDDOSListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZoneSettingAdvancedDDOSListResponseResultValue `json:"value"`
	JSON  zoneSettingAdvancedDDOSListResponseResultJSON  `json:"-"`
}

// zoneSettingAdvancedDDOSListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingAdvancedDDOSListResponseResult]
type zoneSettingAdvancedDDOSListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAdvancedDDOSListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingAdvancedDDOSListResponseResultID string

const (
	ZoneSettingAdvancedDDOSListResponseResultIDAdvancedDDOS ZoneSettingAdvancedDDOSListResponseResultID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingAdvancedDDOSListResponseResultEditable bool

const (
	ZoneSettingAdvancedDDOSListResponseResultEditableTrue  ZoneSettingAdvancedDDOSListResponseResultEditable = true
	ZoneSettingAdvancedDDOSListResponseResultEditableFalse ZoneSettingAdvancedDDOSListResponseResultEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingAdvancedDDOSListResponseResultValue string

const (
	ZoneSettingAdvancedDDOSListResponseResultValueOn  ZoneSettingAdvancedDDOSListResponseResultValue = "on"
	ZoneSettingAdvancedDDOSListResponseResultValueOff ZoneSettingAdvancedDDOSListResponseResultValue = "off"
)
