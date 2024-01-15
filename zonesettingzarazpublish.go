// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingZarazPublishService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingZarazPublishService] method instead.
type ZoneSettingZarazPublishService struct {
	Options []option.RequestOption
}

// NewZoneSettingZarazPublishService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZarazPublishService(opts ...option.RequestOption) (r *ZoneSettingZarazPublishService) {
	r = &ZoneSettingZarazPublishService{}
	r.Options = opts
	return
}

// Publish current Zaraz preview configuration for a zone.
func (r *ZoneSettingZarazPublishService) New(ctx context.Context, zoneIdentifier string, body ZoneSettingZarazPublishNewParams, opts ...option.RequestOption) (res *ZoneSettingZarazPublishNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/publish", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneSettingZarazPublishNewResponse struct {
	Errors   []ZoneSettingZarazPublishNewResponseError   `json:"errors"`
	Messages []ZoneSettingZarazPublishNewResponseMessage `json:"messages"`
	Result   string                                      `json:"result"`
	// Whether the API call was successful
	Success bool                                   `json:"success"`
	JSON    zoneSettingZarazPublishNewResponseJSON `json:"-"`
}

// zoneSettingZarazPublishNewResponseJSON contains the JSON metadata for the struct
// [ZoneSettingZarazPublishNewResponse]
type zoneSettingZarazPublishNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazPublishNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazPublishNewResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingZarazPublishNewResponseErrorJSON `json:"-"`
}

// zoneSettingZarazPublishNewResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingZarazPublishNewResponseError]
type zoneSettingZarazPublishNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazPublishNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazPublishNewResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneSettingZarazPublishNewResponseMessageJSON `json:"-"`
}

// zoneSettingZarazPublishNewResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingZarazPublishNewResponseMessage]
type zoneSettingZarazPublishNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZarazPublishNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZarazPublishNewParams struct {
	// Zaraz configuration description.
	Body param.Field[string] `json:"body,required"`
}

func (r ZoneSettingZarazPublishNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
