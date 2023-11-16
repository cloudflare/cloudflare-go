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

// AIOpenAIService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAIOpenAIService] method instead.
type AIOpenAIService struct {
	Options []option.RequestOption
}

// NewAIOpenAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIOpenAIService(opts ...option.RequestOption) (r *AIOpenAIService) {
	r = &AIOpenAIService{}
	r.Options = opts
	return
}

// Execute @cf/openai/whisper model.
func (r *AIOpenAIService) Whisper(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AIOpenAIWhisperResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/openai/whisper", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AIOpenAIWhisperResponse struct {
	Text string                      `json:"text"`
	JSON aiOpenAIWhisperResponseJSON `json:"-"`
}

// aiOpenAIWhisperResponseJSON contains the JSON metadata for the struct
// [AIOpenAIWhisperResponse]
type aiOpenAIWhisperResponseJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIOpenAIWhisperResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
