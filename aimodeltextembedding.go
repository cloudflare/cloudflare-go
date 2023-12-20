// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AIModelTextEmbeddingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAIModelTextEmbeddingService]
// method instead.
type AIModelTextEmbeddingService struct {
	Options        []option.RequestOption
	BgeSmallEnV1_5 *AIModelTextEmbeddingBgeSmallEnV1_5Service
	BgeBaseEnV1_5  *AIModelTextEmbeddingBgeBaseEnV1_5Service
	BgeLargeEnV1_5 *AIModelTextEmbeddingBgeLargeEnV1_5Service
}

// NewAIModelTextEmbeddingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIModelTextEmbeddingService(opts ...option.RequestOption) (r *AIModelTextEmbeddingService) {
	r = &AIModelTextEmbeddingService{}
	r.Options = opts
	r.BgeSmallEnV1_5 = NewAIModelTextEmbeddingBgeSmallEnV1_5Service(opts...)
	r.BgeBaseEnV1_5 = NewAIModelTextEmbeddingBgeBaseEnV1_5Service(opts...)
	r.BgeLargeEnV1_5 = NewAIModelTextEmbeddingBgeLargeEnV1_5Service(opts...)
	return
}
