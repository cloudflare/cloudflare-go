// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPLimitService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPLimitService] method instead.
type DLPLimitService struct {
	Options []option.RequestOption
}

// NewDLPLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPLimitService(opts ...option.RequestOption) (r *DLPLimitService) {
	r = &DLPLimitService{}
	r.Options = opts
	return
}

// Fetch limits associated with DLP for account
func (r *DLPLimitService) List(ctx context.Context, query DLPLimitListParams, opts ...option.RequestOption) (res *DLPLimitListResponse, err error) {
	var env DLPLimitListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/limits", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPLimitListResponse struct {
	MaxDatasetCells int64                    `json:"max_dataset_cells,required"`
	JSON            dlpLimitListResponseJSON `json:"-"`
}

// dlpLimitListResponseJSON contains the JSON metadata for the struct
// [DLPLimitListResponse]
type dlpLimitListResponseJSON struct {
	MaxDatasetCells apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DLPLimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseJSON) RawJSON() string {
	return r.raw
}

type DLPLimitListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPLimitListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPLimitListResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPLimitListResponse                `json:"result"`
	JSON    dlpLimitListResponseEnvelopeJSON    `json:"-"`
}

// dlpLimitListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPLimitListResponseEnvelope]
type dlpLimitListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPLimitListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPLimitListResponseEnvelopeSuccess bool

const (
	DLPLimitListResponseEnvelopeSuccessTrue DLPLimitListResponseEnvelopeSuccess = true
)

func (r DLPLimitListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPLimitListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
