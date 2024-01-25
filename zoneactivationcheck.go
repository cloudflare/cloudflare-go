// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneActivationCheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneActivationCheckService]
// method instead.
type ZoneActivationCheckService struct {
	Options []option.RequestOption
}

// NewZoneActivationCheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneActivationCheckService(opts ...option.RequestOption) (r *ZoneActivationCheckService) {
	r = &ZoneActivationCheckService{}
	r.Options = opts
	return
}

// Triggeres a new activation check for a PENDING Zone. This can be triggered every
// 5 min for paygo/ent customers, every hour for FREE Zones.
func (r *ZoneActivationCheckService) PutZonesZoneIDActivationCheck(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneActivationCheckPutZonesZoneIDActivationCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/activation_check", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type ZoneActivationCheckPutZonesZoneIDActivationCheckResponse struct {
	Errors   []ZoneActivationCheckPutZonesZoneIDActivationCheckResponseError   `json:"errors"`
	Messages []ZoneActivationCheckPutZonesZoneIDActivationCheckResponseMessage `json:"messages"`
	Result   ZoneActivationCheckPutZonesZoneIDActivationCheckResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneActivationCheckPutZonesZoneIDActivationCheckResponseSuccess `json:"success"`
	JSON    zoneActivationCheckPutZonesZoneIDActivationCheckResponseJSON    `json:"-"`
}

// zoneActivationCheckPutZonesZoneIDActivationCheckResponseJSON contains the JSON
// metadata for the struct
// [ZoneActivationCheckPutZonesZoneIDActivationCheckResponse]
type zoneActivationCheckPutZonesZoneIDActivationCheckResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckPutZonesZoneIDActivationCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneActivationCheckPutZonesZoneIDActivationCheckResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneActivationCheckPutZonesZoneIDActivationCheckResponseErrorJSON `json:"-"`
}

// zoneActivationCheckPutZonesZoneIDActivationCheckResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneActivationCheckPutZonesZoneIDActivationCheckResponseError]
type zoneActivationCheckPutZonesZoneIDActivationCheckResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckPutZonesZoneIDActivationCheckResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneActivationCheckPutZonesZoneIDActivationCheckResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneActivationCheckPutZonesZoneIDActivationCheckResponseMessageJSON `json:"-"`
}

// zoneActivationCheckPutZonesZoneIDActivationCheckResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneActivationCheckPutZonesZoneIDActivationCheckResponseMessage]
type zoneActivationCheckPutZonesZoneIDActivationCheckResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckPutZonesZoneIDActivationCheckResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneActivationCheckPutZonesZoneIDActivationCheckResponseResult struct {
	// Identifier
	ID   string                                                             `json:"id"`
	JSON zoneActivationCheckPutZonesZoneIDActivationCheckResponseResultJSON `json:"-"`
}

// zoneActivationCheckPutZonesZoneIDActivationCheckResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneActivationCheckPutZonesZoneIDActivationCheckResponseResult]
type zoneActivationCheckPutZonesZoneIDActivationCheckResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneActivationCheckPutZonesZoneIDActivationCheckResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneActivationCheckPutZonesZoneIDActivationCheckResponseSuccess bool

const (
	ZoneActivationCheckPutZonesZoneIDActivationCheckResponseSuccessTrue ZoneActivationCheckPutZonesZoneIDActivationCheckResponseSuccess = true
)
