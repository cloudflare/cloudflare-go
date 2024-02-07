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

// LogpushDatasetFieldService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogpushDatasetFieldService]
// method instead.
type LogpushDatasetFieldService struct {
	Options []option.RequestOption
}

// NewLogpushDatasetFieldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogpushDatasetFieldService(opts ...option.RequestOption) (r *LogpushDatasetFieldService) {
	r = &LogpushDatasetFieldService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *LogpushDatasetFieldService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, datasetID string, opts ...option.RequestOption) (res *LogpushDatasetFieldListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/logpush/datasets/%s/fields", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogpushDatasetFieldListResponse struct {
	Errors   []LogpushDatasetFieldListResponseError   `json:"errors"`
	Messages []LogpushDatasetFieldListResponseMessage `json:"messages"`
	Result   interface{}                              `json:"result"`
	// Whether the API call was successful
	Success LogpushDatasetFieldListResponseSuccess `json:"success"`
	JSON    logpushDatasetFieldListResponseJSON    `json:"-"`
}

// logpushDatasetFieldListResponseJSON contains the JSON metadata for the struct
// [LogpushDatasetFieldListResponse]
type logpushDatasetFieldListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetFieldListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushDatasetFieldListResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    logpushDatasetFieldListResponseErrorJSON `json:"-"`
}

// logpushDatasetFieldListResponseErrorJSON contains the JSON metadata for the
// struct [LogpushDatasetFieldListResponseError]
type logpushDatasetFieldListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetFieldListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushDatasetFieldListResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    logpushDatasetFieldListResponseMessageJSON `json:"-"`
}

// logpushDatasetFieldListResponseMessageJSON contains the JSON metadata for the
// struct [LogpushDatasetFieldListResponseMessage]
type logpushDatasetFieldListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetFieldListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushDatasetFieldListResponseSuccess bool

const (
	LogpushDatasetFieldListResponseSuccessTrue LogpushDatasetFieldListResponseSuccess = true
)
