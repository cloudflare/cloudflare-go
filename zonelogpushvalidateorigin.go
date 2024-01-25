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

// ZoneLogpushValidateOriginService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneLogpushValidateOriginService] method instead.
type ZoneLogpushValidateOriginService struct {
	Options []option.RequestOption
}

// NewZoneLogpushValidateOriginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneLogpushValidateOriginService(opts ...option.RequestOption) (r *ZoneLogpushValidateOriginService) {
	r = &ZoneLogpushValidateOriginService{}
	r.Options = opts
	return
}

// Validates logpull origin with logpull_options.
func (r *ZoneLogpushValidateOriginService) PostZonesZoneIdentifierLogpushValidateOrigin(ctx context.Context, zoneIdentifier string, body ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginParams, opts ...option.RequestOption) (res *ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/validate/origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponse struct {
	Errors   []ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseError   `json:"errors"`
	Messages []ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseMessage `json:"messages"`
	Result   ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseSuccess `json:"success"`
	JSON    zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseJSON    `json:"-"`
}

// zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponse]
type zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseError struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseErrorJSON `json:"-"`
}

// zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseError]
type zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseMessage struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseMessageJSON `json:"-"`
}

// zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseMessage]
type zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseResult struct {
	Message string                                                                                  `json:"message"`
	Valid   bool                                                                                    `json:"valid"`
	JSON    zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseResultJSON `json:"-"`
}

// zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseResult]
type zoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseResultJSON struct {
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseSuccess bool

const (
	ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseSuccessTrue ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginResponseSuccess = true
)

type ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginParams struct {
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options,required" format:"uri-reference"`
}

func (r ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
