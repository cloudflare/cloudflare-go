// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

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

// IndicatorFeedDownloadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIndicatorFeedDownloadService] method instead.
type IndicatorFeedDownloadService struct {
	Options []option.RequestOption
}

// NewIndicatorFeedDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIndicatorFeedDownloadService(opts ...option.RequestOption) (r *IndicatorFeedDownloadService) {
	r = &IndicatorFeedDownloadService{}
	r.Options = opts
	return
}

// Download indicator feed data
func (r *IndicatorFeedDownloadService) Get(ctx context.Context, feedID int64, query IndicatorFeedDownloadGetParams, opts ...option.RequestOption) (res *IndicatorFeedDownloadGetResponse, err error) {
	var env IndicatorFeedDownloadGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/indicator_feeds/%v/download", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IndicatorFeedDownloadGetResponse struct {
	// Feed id
	FileID int64 `json:"file_id"`
	// Name of the file unified in our system
	Filename string `json:"filename"`
	// Current status of upload, should be unified
	Status string                               `json:"status"`
	JSON   indicatorFeedDownloadGetResponseJSON `json:"-"`
}

// indicatorFeedDownloadGetResponseJSON contains the JSON metadata for the struct
// [IndicatorFeedDownloadGetResponse]
type indicatorFeedDownloadGetResponseJSON struct {
	FileID      apijson.Field
	Filename    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseJSON) RawJSON() string {
	return r.raw
}

type IndicatorFeedDownloadGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type IndicatorFeedDownloadGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IndicatorFeedDownloadGetResponseEnvelopeSuccess `json:"success,required"`
	Result  IndicatorFeedDownloadGetResponse                `json:"result"`
	JSON    indicatorFeedDownloadGetResponseEnvelopeJSON    `json:"-"`
}

// indicatorFeedDownloadGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [IndicatorFeedDownloadGetResponseEnvelope]
type indicatorFeedDownloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IndicatorFeedDownloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r indicatorFeedDownloadGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IndicatorFeedDownloadGetResponseEnvelopeSuccess bool

const (
	IndicatorFeedDownloadGetResponseEnvelopeSuccessTrue IndicatorFeedDownloadGetResponseEnvelopeSuccess = true
)

func (r IndicatorFeedDownloadGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IndicatorFeedDownloadGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
