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

// AIHuggingfaceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAIHuggingfaceService] method
// instead.
type AIHuggingfaceService struct {
	Options []option.RequestOption
}

// NewAIHuggingfaceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIHuggingfaceService(opts ...option.RequestOption) (r *AIHuggingfaceService) {
	r = &AIHuggingfaceService{}
	r.Options = opts
	return
}

// Execute @cf/huggingface/distilbert-sst-2-int8 model.
func (r *AIHuggingfaceService) DistilbertSst2Int8(ctx context.Context, accountIdentifier string, body AIHuggingfaceDistilbertSst2Int8Params, opts ...option.RequestOption) (res *[]AIHuggingfaceDistilbertSst2Int8Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/huggingface/distilbert-sst-2-int8", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIHuggingfaceDistilbertSst2Int8Response struct {
	Label string                                      `json:"label"`
	Score float64                                     `json:"score"`
	JSON  aiHuggingfaceDistilbertSst2Int8ResponseJSON `json:"-"`
}

// aiHuggingfaceDistilbertSst2Int8ResponseJSON contains the JSON metadata for the
// struct [AIHuggingfaceDistilbertSst2Int8Response]
type aiHuggingfaceDistilbertSst2Int8ResponseJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIHuggingfaceDistilbertSst2Int8Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIHuggingfaceDistilbertSst2Int8Params struct {
	Text param.Field[string] `json:"text,required"`
}

func (r AIHuggingfaceDistilbertSst2Int8Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
