// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package port_scans

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ResultService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResultService] method instead.
type ResultService struct {
	Options []option.RequestOption
}

// NewResultService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewResultService(opts ...option.RequestOption) (r *ResultService) {
	r = &ResultService{}
	r.Options = opts
	return
}

// Get the Latest Scan Result
func (r *ResultService) List(ctx context.Context, query ResultListParams, opts ...option.RequestOption) (res *ResultListResponse, err error) {
	var env ResultListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("%s/scans/results", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ResultListResponse struct {
	Number1_1_1_1 []ResultListResponse1_1_1_1 `json:"1.1.1.1,required"`
	JSON          resultListResponseJSON      `json:"-"`
}

// resultListResponseJSON contains the JSON metadata for the struct
// [ResultListResponse]
type resultListResponseJSON struct {
	Number1_1_1_1 apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ResultListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultListResponseJSON) RawJSON() string {
	return r.raw
}

type ResultListResponse1_1_1_1 struct {
	Number float64                       `json:"number,required"`
	Proto  string                        `json:"proto,required"`
	Status string                        `json:"status,required"`
	JSON   resultListResponse1_1_1_1JSON `json:"-"`
}

// resultListResponse1_1_1_1JSON contains the JSON metadata for the struct
// [ResultListResponse1_1_1_1]
type resultListResponse1_1_1_1JSON struct {
	Number      apijson.Field
	Proto       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultListResponse1_1_1_1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultListResponse1_1_1_1JSON) RawJSON() string {
	return r.raw
}

type ResultListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ResultListResponseEnvelope struct {
	Errors   []string                       `json:"errors,required"`
	Messages []string                       `json:"messages,required"`
	Result   ResultListResponse             `json:"result,required"`
	Success  bool                           `json:"success,required"`
	JSON     resultListResponseEnvelopeJSON `json:"-"`
}

// resultListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ResultListResponseEnvelope]
type resultListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
