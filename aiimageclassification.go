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

// AIImageClassificationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAIImageClassificationService]
// method instead.
type AIImageClassificationService struct {
	Options []option.RequestOption
}

// NewAIImageClassificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIImageClassificationService(opts ...option.RequestOption) (r *AIImageClassificationService) {
	r = &AIImageClassificationService{}
	r.Options = opts
	return
}

// Execute @cf/microsoft/resnet-50 model.
func (r *AIImageClassificationService) Resnet50(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]AIImageClassificationResnet50Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/microsoft/resnet-50", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AIImageClassificationResnet50Response struct {
	Label string                                    `json:"label"`
	Score float64                                   `json:"score"`
	JSON  aiImageClassificationResnet50ResponseJSON `json:"-"`
}

// aiImageClassificationResnet50ResponseJSON contains the JSON metadata for the
// struct [AIImageClassificationResnet50Response]
type aiImageClassificationResnet50ResponseJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIImageClassificationResnet50Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
