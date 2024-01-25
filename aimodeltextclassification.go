// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AIModelTextClassificationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAIModelTextClassificationService] method instead.
type AIModelTextClassificationService struct {
	Options            []option.RequestOption
	DistilbertSst2Int8 *AIModelTextClassificationDistilbertSst2Int8Service
}

// NewAIModelTextClassificationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIModelTextClassificationService(opts ...option.RequestOption) (r *AIModelTextClassificationService) {
	r = &AIModelTextClassificationService{}
	r.Options = opts
	r.DistilbertSst2Int8 = NewAIModelTextClassificationDistilbertSst2Int8Service(opts...)
	return
}
