// File generated from our OpenAPI spec by Stainless.

package logpush

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// DatasetFieldService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDatasetFieldService] method
// instead.
type DatasetFieldService struct {
	Options []option.RequestOption
}

// NewDatasetFieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasetFieldService(opts ...option.RequestOption) (r *DatasetFieldService) {
	r = &DatasetFieldService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *DatasetFieldService) Get(ctx context.Context, datasetID string, query DatasetFieldGetParams, opts ...option.RequestOption) (res *DatasetFieldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DatasetFieldGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/datasets/%s/fields", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DatasetFieldGetResponse = interface{}

type DatasetFieldGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type DatasetFieldGetResponseEnvelope struct {
	Errors   []DatasetFieldGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DatasetFieldGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DatasetFieldGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DatasetFieldGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    datasetFieldGetResponseEnvelopeJSON    `json:"-"`
}

// datasetFieldGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetFieldGetResponseEnvelope]
type datasetFieldGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DatasetFieldGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    datasetFieldGetResponseEnvelopeErrorsJSON `json:"-"`
}

// datasetFieldGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DatasetFieldGetResponseEnvelopeErrors]
type datasetFieldGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DatasetFieldGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    datasetFieldGetResponseEnvelopeMessagesJSON `json:"-"`
}

// datasetFieldGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DatasetFieldGetResponseEnvelopeMessages]
type datasetFieldGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatasetFieldGetResponseEnvelopeSuccess bool

const (
	DatasetFieldGetResponseEnvelopeSuccessTrue DatasetFieldGetResponseEnvelopeSuccess = true
)
