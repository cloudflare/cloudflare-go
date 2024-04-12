// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

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

// PreviewService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPreviewService] method instead.
type PreviewService struct {
	Options []option.RequestOption
}

// NewPreviewService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPreviewService(opts ...option.RequestOption) (r *PreviewService) {
	r = &PreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *PreviewService) Get(ctx context.Context, previewID string, query PreviewGetParams, opts ...option.RequestOption) (res *PreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PreviewGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/preview/%s", query.AccountID, previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PreviewGetResponse map[string]PreviewGetResponse

type PreviewGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PreviewGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Resulting health data from a preview operation.
	Result PreviewGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success PreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    previewGetResponseEnvelopeJSON    `json:"-"`
}

// previewGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PreviewGetResponseEnvelope]
type previewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewGetResponseEnvelopeSuccess bool

const (
	PreviewGetResponseEnvelopeSuccessTrue PreviewGetResponseEnvelopeSuccess = true
)

func (r PreviewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PreviewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
