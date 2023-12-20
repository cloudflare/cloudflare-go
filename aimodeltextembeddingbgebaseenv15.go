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

// AIModelTextEmbeddingBgeBaseEnV1_5Service contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAIModelTextEmbeddingBgeBaseEnV1_5Service] method instead.
type AIModelTextEmbeddingBgeBaseEnV1_5Service struct {
	Options []option.RequestOption
}

// NewAIModelTextEmbeddingBgeBaseEnV1_5Service generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIModelTextEmbeddingBgeBaseEnV1_5Service(opts ...option.RequestOption) (r *AIModelTextEmbeddingBgeBaseEnV1_5Service) {
	r = &AIModelTextEmbeddingBgeBaseEnV1_5Service{}
	r.Options = opts
	return
}

// Execute @cf/baai/bge-base-en-v1.5 model.
func (r *AIModelTextEmbeddingBgeBaseEnV1_5Service) Run(ctx context.Context, accountIdentifier string, body AIModelTextEmbeddingBgeBaseEnV1_5RunParams, opts ...option.RequestOption) (res *AIModelTextEmbeddingBgeBaseEnV1_5RunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-base-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIModelTextEmbeddingBgeBaseEnV1_5RunResponse struct {
	Data  [][]float64                                      `json:"data"`
	Shape []float64                                        `json:"shape"`
	JSON  aiModelTextEmbeddingBgeBaseEnV1_5RunResponseJSON `json:"-"`
}

// aiModelTextEmbeddingBgeBaseEnV1_5RunResponseJSON contains the JSON metadata for
// the struct [AIModelTextEmbeddingBgeBaseEnV1_5RunResponse]
type aiModelTextEmbeddingBgeBaseEnV1_5RunResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelTextEmbeddingBgeBaseEnV1_5RunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIModelTextEmbeddingBgeBaseEnV1_5RunParams struct {
	Text param.Field[AIModelTextEmbeddingBgeBaseEnV1_5RunParamsText] `json:"text,required"`
}

func (r AIModelTextEmbeddingBgeBaseEnV1_5RunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [AIModelTextEmbeddingBgeBaseEnV1_5RunParamsTextObject].
type AIModelTextEmbeddingBgeBaseEnV1_5RunParamsText interface {
	ImplementsAIModelTextEmbeddingBgeBaseEnV1_5RunParamsText()
}
