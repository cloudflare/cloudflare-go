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

// ZoneLogpushValidateDestinationExistService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneLogpushValidateDestinationExistService] method instead.
type ZoneLogpushValidateDestinationExistService struct {
	Options []option.RequestOption
}

// NewZoneLogpushValidateDestinationExistService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneLogpushValidateDestinationExistService(opts ...option.RequestOption) (r *ZoneLogpushValidateDestinationExistService) {
	r = &ZoneLogpushValidateDestinationExistService{}
	r.Options = opts
	return
}

// Checks if there is an existing job with a destination.
func (r *ZoneLogpushValidateDestinationExistService) PostZonesZoneIdentifierLogpushValidateDestinationExists(ctx context.Context, zoneIdentifier string, body ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsParams, opts ...option.RequestOption) (res *ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/validate/destination/exists", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponse struct {
	Errors   []ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseError   `json:"errors"`
	Messages []ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseMessage `json:"messages"`
	Result   ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseSuccess `json:"success"`
	JSON    zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseJSON    `json:"-"`
}

// zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponse]
type zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseError struct {
	Code    int64                                                                                                       `json:"code,required"`
	Message string                                                                                                      `json:"message,required"`
	JSON    zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseErrorJSON `json:"-"`
}

// zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseError]
type zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseMessage struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseMessageJSON `json:"-"`
}

// zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseMessage]
type zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseResult struct {
	Exists bool                                                                                                         `json:"exists"`
	JSON   zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseResultJSON `json:"-"`
}

// zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseResult]
type zoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseResultJSON struct {
	Exists      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseSuccess bool

const (
	ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseSuccessTrue ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsResponseSuccess = true
)

type ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
