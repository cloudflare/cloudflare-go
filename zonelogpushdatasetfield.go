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

// ZoneLogpushDatasetFieldService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneLogpushDatasetFieldService] method instead.
type ZoneLogpushDatasetFieldService struct {
	Options []option.RequestOption
}

// NewZoneLogpushDatasetFieldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneLogpushDatasetFieldService(opts ...option.RequestOption) (r *ZoneLogpushDatasetFieldService) {
	r = &ZoneLogpushDatasetFieldService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *ZoneLogpushDatasetFieldService) GetZonesZoneIdentifierLogpushDatasetsDatasetFields(ctx context.Context, zoneIdentifier string, dataset string, opts ...option.RequestOption) (res *ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/datasets/%s/fields", zoneIdentifier, dataset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponse struct {
	Errors   []ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseError   `json:"errors"`
	Messages []ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseMessage `json:"messages"`
	Result   interface{}                                                                                `json:"result"`
	// Whether the API call was successful
	Success ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseSuccess `json:"success"`
	JSON    zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseJSON    `json:"-"`
}

// zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseJSON
// contains the JSON metadata for the struct
// [ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponse]
type zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseError struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseErrorJSON `json:"-"`
}

// zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseError]
type zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseMessage struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseMessageJSON `json:"-"`
}

// zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseMessage]
type zoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseSuccess bool

const (
	ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseSuccessTrue ZoneLogpushDatasetFieldGetZonesZoneIdentifierLogpushDatasetsDatasetFieldsResponseSuccess = true
)
