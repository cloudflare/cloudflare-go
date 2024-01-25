// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AIModelTextClassificationDistilbertSst2Int8Service contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIModelTextClassificationDistilbertSst2Int8Service] method instead.
type AIModelTextClassificationDistilbertSst2Int8Service struct {
	Options []option.RequestOption
}

// NewAIModelTextClassificationDistilbertSst2Int8Service generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAIModelTextClassificationDistilbertSst2Int8Service(opts ...option.RequestOption) (r *AIModelTextClassificationDistilbertSst2Int8Service) {
	r = &AIModelTextClassificationDistilbertSst2Int8Service{}
	r.Options = opts
	return
}

// Execute @cf/huggingface/distilbert-sst-2-int8 model.
func (r *AIModelTextClassificationDistilbertSst2Int8Service) Run(ctx context.Context, accountIdentifier string, body AIModelTextClassificationDistilbertSst2Int8RunParams, opts ...option.RequestOption) (res *[]AIModelTextClassificationDistilbertSst2Int8RunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/huggingface/distilbert-sst-2-int8", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIModelTextClassificationDistilbertSst2Int8RunResponse struct {
	Label string                                                     `json:"label"`
	Score float64                                                    `json:"score"`
	JSON  aiModelTextClassificationDistilbertSst2Int8RunResponseJSON `json:"-"`
}

// aiModelTextClassificationDistilbertSst2Int8RunResponseJSON contains the JSON
// metadata for the struct [AIModelTextClassificationDistilbertSst2Int8RunResponse]
type aiModelTextClassificationDistilbertSst2Int8RunResponseJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelTextClassificationDistilbertSst2Int8RunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIModelTextClassificationDistilbertSst2Int8RunParams struct {
	Text param.Field[string] `json:"text,required"`
}

func (r AIModelTextClassificationDistilbertSst2Int8RunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
