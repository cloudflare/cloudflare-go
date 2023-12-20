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

// AIModelTextEmbeddingBgeLargeEnV1_5Service contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAIModelTextEmbeddingBgeLargeEnV1_5Service] method instead.
type AIModelTextEmbeddingBgeLargeEnV1_5Service struct {
	Options []option.RequestOption
}

// NewAIModelTextEmbeddingBgeLargeEnV1_5Service generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAIModelTextEmbeddingBgeLargeEnV1_5Service(opts ...option.RequestOption) (r *AIModelTextEmbeddingBgeLargeEnV1_5Service) {
	r = &AIModelTextEmbeddingBgeLargeEnV1_5Service{}
	r.Options = opts
	return
}

// Execute @cf/baai/bge-large-en-v1.5 model.
func (r *AIModelTextEmbeddingBgeLargeEnV1_5Service) Run(ctx context.Context, accountIdentifier string, body AIModelTextEmbeddingBgeLargeEnV1_5RunParams, opts ...option.RequestOption) (res *AIModelTextEmbeddingBgeLargeEnV1_5RunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-large-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIModelTextEmbeddingBgeLargeEnV1_5RunResponse struct {
	Data  [][]float64                                       `json:"data"`
	Shape []float64                                         `json:"shape"`
	JSON  aiModelTextEmbeddingBgeLargeEnV1_5RunResponseJSON `json:"-"`
}

// aiModelTextEmbeddingBgeLargeEnV1_5RunResponseJSON contains the JSON metadata for
// the struct [AIModelTextEmbeddingBgeLargeEnV1_5RunResponse]
type aiModelTextEmbeddingBgeLargeEnV1_5RunResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelTextEmbeddingBgeLargeEnV1_5RunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIModelTextEmbeddingBgeLargeEnV1_5RunParams struct {
	Text param.Field[AIModelTextEmbeddingBgeLargeEnV1_5RunParamsText] `json:"text,required"`
}

func (r AIModelTextEmbeddingBgeLargeEnV1_5RunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [AIModelTextEmbeddingBgeLargeEnV1_5RunParamsTextObject].
type AIModelTextEmbeddingBgeLargeEnV1_5RunParamsText interface {
	ImplementsAIModelTextEmbeddingBgeLargeEnV1_5RunParamsText()
}
