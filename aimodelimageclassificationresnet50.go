// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AIModelImageClassificationResnet50Service contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAIModelImageClassificationResnet50Service] method instead.
type AIModelImageClassificationResnet50Service struct {
	Options []option.RequestOption
}

// NewAIModelImageClassificationResnet50Service generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAIModelImageClassificationResnet50Service(opts ...option.RequestOption) (r *AIModelImageClassificationResnet50Service) {
	r = &AIModelImageClassificationResnet50Service{}
	r.Options = opts
	return
}

// Execute @cf/microsoft/resnet-50 model.
func (r *AIModelImageClassificationResnet50Service) Run(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]AIModelImageClassificationResnet50RunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/microsoft/resnet-50", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AIModelImageClassificationResnet50RunResponse struct {
	Label string                                            `json:"label"`
	Score float64                                           `json:"score"`
	JSON  aiModelImageClassificationResnet50RunResponseJSON `json:"-"`
}

// aiModelImageClassificationResnet50RunResponseJSON contains the JSON metadata for
// the struct [AIModelImageClassificationResnet50RunResponse]
type aiModelImageClassificationResnet50RunResponseJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelImageClassificationResnet50RunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
