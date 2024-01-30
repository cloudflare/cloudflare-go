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

// ZoneSettingSecurityHeaderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingSecurityHeaderService] method instead.
type ZoneSettingSecurityHeaderService struct {
	Options []option.RequestOption
}

// NewZoneSettingSecurityHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSecurityHeaderService(opts ...option.RequestOption) (r *ZoneSettingSecurityHeaderService) {
	r = &ZoneSettingSecurityHeaderService{}
	r.Options = opts
	return
}

// Cloudflare security header for a zone.
func (r *ZoneSettingSecurityHeaderService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingSecurityHeaderUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSecurityHeaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_header", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Cloudflare security header for a zone.
func (r *ZoneSettingSecurityHeaderService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingSecurityHeaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_header", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingSecurityHeaderUpdateResponse struct {
	Errors   []ZoneSettingSecurityHeaderUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingSecurityHeaderUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                          `json:"success,required"`
	Result  ZoneSettingSecurityHeaderUpdateResponseResult `json:"result"`
	JSON    zoneSettingSecurityHeaderUpdateResponseJSON   `json:"-"`
}

// zoneSettingSecurityHeaderUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityHeaderUpdateResponse]
type zoneSettingSecurityHeaderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderUpdateResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingSecurityHeaderUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingSecurityHeaderUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderUpdateResponseError]
type zoneSettingSecurityHeaderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderUpdateResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingSecurityHeaderUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingSecurityHeaderUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingSecurityHeaderUpdateResponseMessage]
type zoneSettingSecurityHeaderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderUpdateResponseResult struct {
	// ID of the zone's security header.
	ID    ZoneSettingSecurityHeaderUpdateResponseResultID    `json:"id,required"`
	Value ZoneSettingSecurityHeaderUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityHeaderUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSecurityHeaderUpdateResponseResultJSON `json:"-"`
}

// zoneSettingSecurityHeaderUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderUpdateResponseResult]
type zoneSettingSecurityHeaderUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone's security header.
type ZoneSettingSecurityHeaderUpdateResponseResultID string

const (
	ZoneSettingSecurityHeaderUpdateResponseResultIDSecurityHeader ZoneSettingSecurityHeaderUpdateResponseResultID = "security_header"
)

type ZoneSettingSecurityHeaderUpdateResponseResultValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingSecurityHeaderUpdateResponseResultValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingSecurityHeaderUpdateResponseResultValueJSON                    `json:"-"`
}

// zoneSettingSecurityHeaderUpdateResponseResultValueJSON contains the JSON
// metadata for the struct [ZoneSettingSecurityHeaderUpdateResponseResultValue]
type zoneSettingSecurityHeaderUpdateResponseResultValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderUpdateResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZoneSettingSecurityHeaderUpdateResponseResultValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                          `json:"nosniff"`
	JSON    zoneSettingSecurityHeaderUpdateResponseResultValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingSecurityHeaderUpdateResponseResultValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingSecurityHeaderUpdateResponseResultValueStrictTransportSecurity]
type zoneSettingSecurityHeaderUpdateResponseResultValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderUpdateResponseResultValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityHeaderUpdateResponseResultEditable bool

const (
	ZoneSettingSecurityHeaderUpdateResponseResultEditableTrue  ZoneSettingSecurityHeaderUpdateResponseResultEditable = true
	ZoneSettingSecurityHeaderUpdateResponseResultEditableFalse ZoneSettingSecurityHeaderUpdateResponseResultEditable = false
)

type ZoneSettingSecurityHeaderListResponse struct {
	Errors   []ZoneSettingSecurityHeaderListResponseError   `json:"errors,required"`
	Messages []ZoneSettingSecurityHeaderListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                        `json:"success,required"`
	Result  ZoneSettingSecurityHeaderListResponseResult `json:"result"`
	JSON    zoneSettingSecurityHeaderListResponseJSON   `json:"-"`
}

// zoneSettingSecurityHeaderListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSecurityHeaderListResponse]
type zoneSettingSecurityHeaderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingSecurityHeaderListResponseErrorJSON `json:"-"`
}

// zoneSettingSecurityHeaderListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderListResponseError]
type zoneSettingSecurityHeaderListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingSecurityHeaderListResponseMessageJSON `json:"-"`
}

// zoneSettingSecurityHeaderListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderListResponseMessage]
type zoneSettingSecurityHeaderListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingSecurityHeaderListResponseResult struct {
	// ID of the zone's security header.
	ID    ZoneSettingSecurityHeaderListResponseResultID    `json:"id,required"`
	Value ZoneSettingSecurityHeaderListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingSecurityHeaderListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingSecurityHeaderListResponseResultJSON `json:"-"`
}

// zoneSettingSecurityHeaderListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingSecurityHeaderListResponseResult]
type zoneSettingSecurityHeaderListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone's security header.
type ZoneSettingSecurityHeaderListResponseResultID string

const (
	ZoneSettingSecurityHeaderListResponseResultIDSecurityHeader ZoneSettingSecurityHeaderListResponseResultID = "security_header"
)

type ZoneSettingSecurityHeaderListResponseResultValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingSecurityHeaderListResponseResultValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingSecurityHeaderListResponseResultValueJSON                    `json:"-"`
}

// zoneSettingSecurityHeaderListResponseResultValueJSON contains the JSON metadata
// for the struct [ZoneSettingSecurityHeaderListResponseResultValue]
type zoneSettingSecurityHeaderListResponseResultValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderListResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZoneSettingSecurityHeaderListResponseResultValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                        `json:"nosniff"`
	JSON    zoneSettingSecurityHeaderListResponseResultValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingSecurityHeaderListResponseResultValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingSecurityHeaderListResponseResultValueStrictTransportSecurity]
type zoneSettingSecurityHeaderListResponseResultValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingSecurityHeaderListResponseResultValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingSecurityHeaderListResponseResultEditable bool

const (
	ZoneSettingSecurityHeaderListResponseResultEditableTrue  ZoneSettingSecurityHeaderListResponseResultEditable = true
	ZoneSettingSecurityHeaderListResponseResultEditableFalse ZoneSettingSecurityHeaderListResponseResultEditable = false
)

type ZoneSettingSecurityHeaderUpdateParams struct {
	Value param.Field[ZoneSettingSecurityHeaderUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSecurityHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingSecurityHeaderUpdateParamsValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[ZoneSettingSecurityHeaderUpdateParamsValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r ZoneSettingSecurityHeaderUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type ZoneSettingSecurityHeaderUpdateParamsValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r ZoneSettingSecurityHeaderUpdateParamsValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
