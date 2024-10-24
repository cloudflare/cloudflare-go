// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// DatasetJobService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetJobService] method instead.
type DatasetJobService struct {
	Options []option.RequestOption
}

// NewDatasetJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasetJobService(opts ...option.RequestOption) (r *DatasetJobService) {
	r = &DatasetJobService{}
	r.Options = opts
	return
}

// Lists Logpush jobs for an account or zone for a dataset.
func (r *DatasetJobService) Get(ctx context.Context, datasetID string, query DatasetJobGetParams, opts ...option.RequestOption) (res *[]LogpushJob, err error) {
	var env DatasetJobGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/logpush/datasets/%s/jobs", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DatasetJobGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type DatasetJobGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DatasetJobGetResponseEnvelopeSuccess `json:"success,required"`
	Result  []LogpushJob                         `json:"result"`
	JSON    datasetJobGetResponseEnvelopeJSON    `json:"-"`
}

// datasetJobGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetJobGetResponseEnvelope]
type datasetJobGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetJobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJobGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatasetJobGetResponseEnvelopeSuccess bool

const (
	DatasetJobGetResponseEnvelopeSuccessTrue DatasetJobGetResponseEnvelopeSuccess = true
)

func (r DatasetJobGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatasetJobGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
