// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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

// ScanResultService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanResultService] method instead.
type ScanResultService struct {
	Options []option.RequestOption
}

// NewScanResultService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScanResultService(opts ...option.RequestOption) (r *ScanResultService) {
	r = &ScanResultService{}
	r.Options = opts
	return
}

// Get the Latest Scan Result
func (r *ScanResultService) List(ctx context.Context, query ScanResultListParams, opts ...option.RequestOption) (res *ScanResultListResponse, err error) {
	var env ScanResultListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/results", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScanResultListResponse struct {
	Number1_1_1_1 []ScanResultListResponse1_1_1_1 `json:"1.1.1.1,required"`
	JSON          scanResultListResponseJSON      `json:"-"`
}

// scanResultListResponseJSON contains the JSON metadata for the struct
// [ScanResultListResponse]
type scanResultListResponseJSON struct {
	Number1_1_1_1 apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScanResultListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanResultListResponseJSON) RawJSON() string {
	return r.raw
}

type ScanResultListResponse1_1_1_1 struct {
	Number float64                           `json:"number"`
	Proto  string                            `json:"proto"`
	Status string                            `json:"status"`
	JSON   scanResultListResponse1_1_1_1JSON `json:"-"`
}

// scanResultListResponse1_1_1_1JSON contains the JSON metadata for the struct
// [ScanResultListResponse1_1_1_1]
type scanResultListResponse1_1_1_1JSON struct {
	Number      apijson.Field
	Proto       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanResultListResponse1_1_1_1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanResultListResponse1_1_1_1JSON) RawJSON() string {
	return r.raw
}

type ScanResultListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanResultListResponseEnvelope struct {
	Errors   []string                           `json:"errors,required"`
	Messages []string                           `json:"messages,required"`
	Result   ScanResultListResponse             `json:"result,required"`
	Success  bool                               `json:"success,required"`
	JSON     scanResultListResponseEnvelopeJSON `json:"-"`
}

// scanResultListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanResultListResponseEnvelope]
type scanResultListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanResultListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanResultListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
