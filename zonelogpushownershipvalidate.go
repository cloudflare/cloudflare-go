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

// ZoneLogpushOwnershipValidateService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneLogpushOwnershipValidateService] method instead.
type ZoneLogpushOwnershipValidateService struct {
	Options []option.RequestOption
}

// NewZoneLogpushOwnershipValidateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneLogpushOwnershipValidateService(opts ...option.RequestOption) (r *ZoneLogpushOwnershipValidateService) {
	r = &ZoneLogpushOwnershipValidateService{}
	r.Options = opts
	return
}

// Validates ownership challenge of the destination.
func (r *ZoneLogpushOwnershipValidateService) PostZonesZoneIdentifierLogpushOwnershipValidate(ctx context.Context, zoneIdentifier string, body ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateParams, opts ...option.RequestOption) (res *ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/ownership/validate", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponse struct {
	Errors   []ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseError   `json:"errors"`
	Messages []ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseMessage `json:"messages"`
	Result   ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseSuccess `json:"success"`
	JSON    zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseJSON    `json:"-"`
}

// zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseJSON
// contains the JSON metadata for the struct
// [ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponse]
type zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseError struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseErrorJSON `json:"-"`
}

// zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseError]
type zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseMessage struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseMessageJSON `json:"-"`
}

// zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseMessage]
type zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseResult struct {
	Valid bool                                                                                          `json:"valid"`
	JSON  zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseResultJSON `json:"-"`
}

// zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseResult]
type zoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseResultJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseSuccess bool

const (
	ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseSuccessTrue ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateResponseSuccess = true
)

type ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r ZoneLogpushOwnershipValidatePostZonesZoneIdentifierLogpushOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
