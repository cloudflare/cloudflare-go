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
	var env LogpushDatasetFieldListResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/datasets/%s/fields", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushDatasetFieldListResponse = interface{}

type LogpushDatasetFieldListResponseEnvelope struct {
	Errors   []LogpushDatasetFieldListResponseEnvelopeErrors   `json:"errors"`
	Messages []LogpushDatasetFieldListResponseEnvelopeMessages `json:"messages"`
	Result   LogpushDatasetFieldListResponse                   `json:"result"`
	// Whether the API call was successful
	Success LogpushDatasetFieldListResponseEnvelopeSuccess `json:"success"`
	JSON    logpushDatasetFieldListResponseEnvelopeJSON    `json:"-"`
}

// logpushDatasetFieldListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogpushDatasetFieldListResponseEnvelope]
type logpushDatasetFieldListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetFieldListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushDatasetFieldListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    logpushDatasetFieldListResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushDatasetFieldListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LogpushDatasetFieldListResponseEnvelopeErrors]
type logpushDatasetFieldListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetFieldListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushDatasetFieldListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    logpushDatasetFieldListResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushDatasetFieldListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LogpushDatasetFieldListResponseEnvelopeMessages]
type logpushDatasetFieldListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetFieldListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushDatasetFieldListResponseEnvelopeSuccess bool

const (
	LogpushDatasetFieldListResponseEnvelopeSuccessTrue LogpushDatasetFieldListResponseEnvelopeSuccess = true
)
