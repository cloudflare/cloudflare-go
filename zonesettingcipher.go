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

// ZoneSettingCipherService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingCipherService] method
// instead.
type ZoneSettingCipherService struct {
	Options []option.RequestOption
}

// NewZoneSettingCipherService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingCipherService(opts ...option.RequestOption) (r *ZoneSettingCipherService) {
	r = &ZoneSettingCipherService{}
	r.Options = opts
	return
}

// Changes ciphers setting.
func (r *ZoneSettingCipherService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingCipherUpdateParams, opts ...option.RequestOption) (res *ZoneSettingCipherUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ciphers", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets ciphers setting.
func (r *ZoneSettingCipherService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingCipherListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ciphers", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingCipherUpdateResponse struct {
	Errors   []ZoneSettingCipherUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingCipherUpdateResponseMessage `json:"messages"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result ZoneSettingCipherUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                `json:"success"`
	JSON    zoneSettingCipherUpdateResponseJSON `json:"-"`
}

// zoneSettingCipherUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingCipherUpdateResponse]
type zoneSettingCipherUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingCipherUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingCipherUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingCipherUpdateResponseError]
type zoneSettingCipherUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSettingCipherUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingCipherUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingCipherUpdateResponseMessage]
type zoneSettingCipherUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingCipherUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingCipherUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingCipherUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value []string                                  `json:"value"`
	JSON  zoneSettingCipherUpdateResponseResultJSON `json:"-"`
}

// zoneSettingCipherUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingCipherUpdateResponseResult]
type zoneSettingCipherUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingCipherUpdateResponseResultID string

const (
	ZoneSettingCipherUpdateResponseResultIDCiphers ZoneSettingCipherUpdateResponseResultID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingCipherUpdateResponseResultEditable bool

const (
	ZoneSettingCipherUpdateResponseResultEditableTrue  ZoneSettingCipherUpdateResponseResultEditable = true
	ZoneSettingCipherUpdateResponseResultEditableFalse ZoneSettingCipherUpdateResponseResultEditable = false
)

type ZoneSettingCipherListResponse struct {
	Errors   []ZoneSettingCipherListResponseError   `json:"errors"`
	Messages []ZoneSettingCipherListResponseMessage `json:"messages"`
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Result ZoneSettingCipherListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    zoneSettingCipherListResponseJSON `json:"-"`
}

// zoneSettingCipherListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingCipherListResponse]
type zoneSettingCipherListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingCipherListResponseErrorJSON `json:"-"`
}

// zoneSettingCipherListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingCipherListResponseError]
type zoneSettingCipherListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingCipherListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingCipherListResponseMessageJSON `json:"-"`
}

// zoneSettingCipherListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingCipherListResponseMessage]
type zoneSettingCipherListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingCipherListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingCipherListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingCipherListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value []string                                `json:"value"`
	JSON  zoneSettingCipherListResponseResultJSON `json:"-"`
}

// zoneSettingCipherListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingCipherListResponseResult]
type zoneSettingCipherListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingCipherListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingCipherListResponseResultID string

const (
	ZoneSettingCipherListResponseResultIDCiphers ZoneSettingCipherListResponseResultID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingCipherListResponseResultEditable bool

const (
	ZoneSettingCipherListResponseResultEditableTrue  ZoneSettingCipherListResponseResultEditable = true
	ZoneSettingCipherListResponseResultEditableFalse ZoneSettingCipherListResponseResultEditable = false
)

type ZoneSettingCipherUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r ZoneSettingCipherUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
