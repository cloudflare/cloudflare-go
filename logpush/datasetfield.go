// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   DatasetFieldGetResponse      `json:"result,required"`
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

// Whether the API call was successful
type DatasetFieldGetResponseEnvelopeSuccess bool

const (
	DatasetFieldGetResponseEnvelopeSuccessTrue DatasetFieldGetResponseEnvelopeSuccess = true
)

func (r DatasetFieldGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatasetFieldGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
