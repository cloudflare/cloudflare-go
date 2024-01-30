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

// ZoneSettingPseudoIpv4Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingPseudoIpv4Service]
// method instead.
type ZoneSettingPseudoIpv4Service struct {
	Options []option.RequestOption
}

// NewZoneSettingPseudoIpv4Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingPseudoIpv4Service(opts ...option.RequestOption) (r *ZoneSettingPseudoIpv4Service) {
	r = &ZoneSettingPseudoIpv4Service{}
	r.Options = opts
	return
}

// Value of the Pseudo IPv4 setting.
func (r *ZoneSettingPseudoIpv4Service) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingPseudoIpv4UpdateParams, opts ...option.RequestOption) (res *ZoneSettingPseudoIpv4UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the Pseudo IPv4 setting.
func (r *ZoneSettingPseudoIpv4Service) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingPseudoIpv4ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingPseudoIpv4UpdateResponse struct {
	Errors   []ZoneSettingPseudoIpv4UpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingPseudoIpv4UpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                      `json:"success,required"`
	Result  ZoneSettingPseudoIpv4UpdateResponseResult `json:"result"`
	JSON    zoneSettingPseudoIpv4UpdateResponseJSON   `json:"-"`
}

// zoneSettingPseudoIpv4UpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIpv4UpdateResponse]
type zoneSettingPseudoIpv4UpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIpv4UpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingPseudoIpv4UpdateResponseErrorJSON `json:"-"`
}

// zoneSettingPseudoIpv4UpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIpv4UpdateResponseError]
type zoneSettingPseudoIpv4UpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4UpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIpv4UpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingPseudoIpv4UpdateResponseMessageJSON `json:"-"`
}

// zoneSettingPseudoIpv4UpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingPseudoIpv4UpdateResponseMessage]
type zoneSettingPseudoIpv4UpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4UpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIpv4UpdateResponseResult struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingPseudoIpv4UpdateResponseResultID `json:"id,required"`
	// Value of the Pseudo IPv4 setting.
	Value ZoneSettingPseudoIpv4UpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPseudoIpv4UpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPseudoIpv4UpdateResponseResultJSON `json:"-"`
}

// zoneSettingPseudoIpv4UpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIpv4UpdateResponseResult]
type zoneSettingPseudoIpv4UpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4UpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Value of the Pseudo IPv4 setting.
type ZoneSettingPseudoIpv4UpdateResponseResultID string

const (
	ZoneSettingPseudoIpv4UpdateResponseResultIDPseudoIpv4 ZoneSettingPseudoIpv4UpdateResponseResultID = "pseudo_ipv4"
)

// Value of the Pseudo IPv4 setting.
type ZoneSettingPseudoIpv4UpdateResponseResultValue string

const (
	ZoneSettingPseudoIpv4UpdateResponseResultValueOff             ZoneSettingPseudoIpv4UpdateResponseResultValue = "off"
	ZoneSettingPseudoIpv4UpdateResponseResultValueAddHeader       ZoneSettingPseudoIpv4UpdateResponseResultValue = "add_header"
	ZoneSettingPseudoIpv4UpdateResponseResultValueOverwriteHeader ZoneSettingPseudoIpv4UpdateResponseResultValue = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPseudoIpv4UpdateResponseResultEditable bool

const (
	ZoneSettingPseudoIpv4UpdateResponseResultEditableTrue  ZoneSettingPseudoIpv4UpdateResponseResultEditable = true
	ZoneSettingPseudoIpv4UpdateResponseResultEditableFalse ZoneSettingPseudoIpv4UpdateResponseResultEditable = false
)

type ZoneSettingPseudoIpv4ListResponse struct {
	Errors   []ZoneSettingPseudoIpv4ListResponseError   `json:"errors,required"`
	Messages []ZoneSettingPseudoIpv4ListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                    `json:"success,required"`
	Result  ZoneSettingPseudoIpv4ListResponseResult `json:"result"`
	JSON    zoneSettingPseudoIpv4ListResponseJSON   `json:"-"`
}

// zoneSettingPseudoIpv4ListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingPseudoIpv4ListResponse]
type zoneSettingPseudoIpv4ListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIpv4ListResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSettingPseudoIpv4ListResponseErrorJSON `json:"-"`
}

// zoneSettingPseudoIpv4ListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIpv4ListResponseError]
type zoneSettingPseudoIpv4ListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4ListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIpv4ListResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneSettingPseudoIpv4ListResponseMessageJSON `json:"-"`
}

// zoneSettingPseudoIpv4ListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIpv4ListResponseMessage]
type zoneSettingPseudoIpv4ListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4ListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPseudoIpv4ListResponseResult struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingPseudoIpv4ListResponseResultID `json:"id,required"`
	// Value of the Pseudo IPv4 setting.
	Value ZoneSettingPseudoIpv4ListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPseudoIpv4ListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPseudoIpv4ListResponseResultJSON `json:"-"`
}

// zoneSettingPseudoIpv4ListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingPseudoIpv4ListResponseResult]
type zoneSettingPseudoIpv4ListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPseudoIpv4ListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Value of the Pseudo IPv4 setting.
type ZoneSettingPseudoIpv4ListResponseResultID string

const (
	ZoneSettingPseudoIpv4ListResponseResultIDPseudoIpv4 ZoneSettingPseudoIpv4ListResponseResultID = "pseudo_ipv4"
)

// Value of the Pseudo IPv4 setting.
type ZoneSettingPseudoIpv4ListResponseResultValue string

const (
	ZoneSettingPseudoIpv4ListResponseResultValueOff             ZoneSettingPseudoIpv4ListResponseResultValue = "off"
	ZoneSettingPseudoIpv4ListResponseResultValueAddHeader       ZoneSettingPseudoIpv4ListResponseResultValue = "add_header"
	ZoneSettingPseudoIpv4ListResponseResultValueOverwriteHeader ZoneSettingPseudoIpv4ListResponseResultValue = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPseudoIpv4ListResponseResultEditable bool

const (
	ZoneSettingPseudoIpv4ListResponseResultEditableTrue  ZoneSettingPseudoIpv4ListResponseResultEditable = true
	ZoneSettingPseudoIpv4ListResponseResultEditableFalse ZoneSettingPseudoIpv4ListResponseResultEditable = false
)

type ZoneSettingPseudoIpv4UpdateParams struct {
	// Value of the Pseudo IPv4 setting.
	Value param.Field[ZoneSettingPseudoIpv4UpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingPseudoIpv4UpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Pseudo IPv4 setting.
type ZoneSettingPseudoIpv4UpdateParamsValue string

const (
	ZoneSettingPseudoIpv4UpdateParamsValueOff             ZoneSettingPseudoIpv4UpdateParamsValue = "off"
	ZoneSettingPseudoIpv4UpdateParamsValueAddHeader       ZoneSettingPseudoIpv4UpdateParamsValue = "add_header"
	ZoneSettingPseudoIpv4UpdateParamsValueOverwriteHeader ZoneSettingPseudoIpv4UpdateParamsValue = "overwrite_header"
)
