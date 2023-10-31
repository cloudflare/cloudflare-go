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

// ZoneSettingOpportunisticEncryptionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingOpportunisticEncryptionService] method instead.
type ZoneSettingOpportunisticEncryptionService struct {
	Options []option.RequestOption
}

// NewZoneSettingOpportunisticEncryptionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSettingOpportunisticEncryptionService(opts ...option.RequestOption) (r *ZoneSettingOpportunisticEncryptionService) {
	r = &ZoneSettingOpportunisticEncryptionService{}
	r.Options = opts
	return
}

// Changes Opportunistic Encryption setting.
func (r *ZoneSettingOpportunisticEncryptionService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingOpportunisticEncryptionUpdateParams, opts ...option.RequestOption) (res *ZoneSettingOpportunisticEncryptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets Opportunistic Encryption setting.
func (r *ZoneSettingOpportunisticEncryptionService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingOpportunisticEncryptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingOpportunisticEncryptionUpdateResponse struct {
	Errors   []ZoneSettingOpportunisticEncryptionUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingOpportunisticEncryptionUpdateResponseMessage `json:"messages"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZoneSettingOpportunisticEncryptionUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingOpportunisticEncryptionUpdateResponseJSON
}

// zoneSettingOpportunisticEncryptionUpdateResponseJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticEncryptionUpdateResponse]
type zoneSettingOpportunisticEncryptionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionUpdateResponseErrorJSON
}

// zoneSettingOpportunisticEncryptionUpdateResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticEncryptionUpdateResponseError]
type zoneSettingOpportunisticEncryptionUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionUpdateResponseMessageJSON
}

// zoneSettingOpportunisticEncryptionUpdateResponseMessageJSON contains the JSON
// metadata for the struct
// [ZoneSettingOpportunisticEncryptionUpdateResponseMessage]
type zoneSettingOpportunisticEncryptionUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingOpportunisticEncryptionUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticEncryptionUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticEncryptionUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingOpportunisticEncryptionUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingOpportunisticEncryptionUpdateResponseResultJSON
}

// zoneSettingOpportunisticEncryptionUpdateResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticEncryptionUpdateResponseResult]
type zoneSettingOpportunisticEncryptionUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOpportunisticEncryptionUpdateResponseResultID string

const (
	ZoneSettingOpportunisticEncryptionUpdateResponseResultIDOpportunisticEncryption ZoneSettingOpportunisticEncryptionUpdateResponseResultID = "opportunistic_encryption"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticEncryptionUpdateResponseResultEditable bool

const (
	ZoneSettingOpportunisticEncryptionUpdateResponseResultEditableTrue  ZoneSettingOpportunisticEncryptionUpdateResponseResultEditable = true
	ZoneSettingOpportunisticEncryptionUpdateResponseResultEditableFalse ZoneSettingOpportunisticEncryptionUpdateResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticEncryptionUpdateResponseResultValue string

const (
	ZoneSettingOpportunisticEncryptionUpdateResponseResultValueOn  ZoneSettingOpportunisticEncryptionUpdateResponseResultValue = "on"
	ZoneSettingOpportunisticEncryptionUpdateResponseResultValueOff ZoneSettingOpportunisticEncryptionUpdateResponseResultValue = "off"
)

type ZoneSettingOpportunisticEncryptionListResponse struct {
	Errors   []ZoneSettingOpportunisticEncryptionListResponseError   `json:"errors"`
	Messages []ZoneSettingOpportunisticEncryptionListResponseMessage `json:"messages"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZoneSettingOpportunisticEncryptionListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingOpportunisticEncryptionListResponseJSON
}

// zoneSettingOpportunisticEncryptionListResponseJSON contains the JSON metadata
// for the struct [ZoneSettingOpportunisticEncryptionListResponse]
type zoneSettingOpportunisticEncryptionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionListResponseErrorJSON
}

// zoneSettingOpportunisticEncryptionListResponseErrorJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticEncryptionListResponseError]
type zoneSettingOpportunisticEncryptionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionListResponseMessageJSON
}

// zoneSettingOpportunisticEncryptionListResponseMessageJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticEncryptionListResponseMessage]
type zoneSettingOpportunisticEncryptionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingOpportunisticEncryptionListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingOpportunisticEncryptionListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingOpportunisticEncryptionListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingOpportunisticEncryptionListResponseResultValue `json:"value"`
	JSON  zoneSettingOpportunisticEncryptionListResponseResultJSON
}

// zoneSettingOpportunisticEncryptionListResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticEncryptionListResponseResult]
type zoneSettingOpportunisticEncryptionListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingOpportunisticEncryptionListResponseResultID string

const (
	ZoneSettingOpportunisticEncryptionListResponseResultIDOpportunisticEncryption ZoneSettingOpportunisticEncryptionListResponseResultID = "opportunistic_encryption"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingOpportunisticEncryptionListResponseResultEditable bool

const (
	ZoneSettingOpportunisticEncryptionListResponseResultEditableTrue  ZoneSettingOpportunisticEncryptionListResponseResultEditable = true
	ZoneSettingOpportunisticEncryptionListResponseResultEditableFalse ZoneSettingOpportunisticEncryptionListResponseResultEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticEncryptionListResponseResultValue string

const (
	ZoneSettingOpportunisticEncryptionListResponseResultValueOn  ZoneSettingOpportunisticEncryptionListResponseResultValue = "on"
	ZoneSettingOpportunisticEncryptionListResponseResultValueOff ZoneSettingOpportunisticEncryptionListResponseResultValue = "off"
)

type ZoneSettingOpportunisticEncryptionUpdateParams struct {
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingOpportunisticEncryptionUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingOpportunisticEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticEncryptionUpdateParamsValue string

const (
	ZoneSettingOpportunisticEncryptionUpdateParamsValueOn  ZoneSettingOpportunisticEncryptionUpdateParamsValue = "on"
	ZoneSettingOpportunisticEncryptionUpdateParamsValueOff ZoneSettingOpportunisticEncryptionUpdateParamsValue = "off"
)
