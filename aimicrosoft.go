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

// AIMicrosoftService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAIMicrosoftService] method
// instead.
type AIMicrosoftService struct {
	Options []option.RequestOption
}

// NewAIMicrosoftService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIMicrosoftService(opts ...option.RequestOption) (r *AIMicrosoftService) {
	r = &AIMicrosoftService{}
	r.Options = opts
	return
}

// Execute @cf/microsoft/resnet-50 model.
func (r *AIMicrosoftService) Resnet50(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]AIMicrosoftResnet50Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/microsoft/resnet-50", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AIMicrosoftResnet50Response struct {
	Label string                          `json:"label"`
	Score float64                         `json:"score"`
	JSON  aiMicrosoftResnet50ResponseJSON `json:"-"`
}

// aiMicrosoftResnet50ResponseJSON contains the JSON metadata for the struct
// [AIMicrosoftResnet50Response]
type aiMicrosoftResnet50ResponseJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMicrosoftResnet50Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
