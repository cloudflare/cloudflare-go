// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AIModelTextGenerationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAIModelTextGenerationService]
// method instead.
type AIModelTextGenerationService struct {
	Options               []option.RequestOption
	Llama2_7bChatInt8     *AIModelTextGenerationLlama2_7bChatInt8Service
	Llama2_7bChatFp16     *AIModelTextGenerationLlama2_7bChatFp16Service
	Mistral7bInstructV0_1 *AIModelTextGenerationMistral7bInstructV0_1Service
}

// NewAIModelTextGenerationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIModelTextGenerationService(opts ...option.RequestOption) (r *AIModelTextGenerationService) {
	r = &AIModelTextGenerationService{}
	r.Options = opts
	r.Llama2_7bChatInt8 = NewAIModelTextGenerationLlama2_7bChatInt8Service(opts...)
	r.Llama2_7bChatFp16 = NewAIModelTextGenerationLlama2_7bChatFp16Service(opts...)
	r.Mistral7bInstructV0_1 = NewAIModelTextGenerationMistral7bInstructV0_1Service(opts...)
	return
}
