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

// AISpeechRecognitionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAISpeechRecognitionService]
// method instead.
type AISpeechRecognitionService struct {
	Options []option.RequestOption
}

// NewAISpeechRecognitionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAISpeechRecognitionService(opts ...option.RequestOption) (r *AISpeechRecognitionService) {
	r = &AISpeechRecognitionService{}
	r.Options = opts
	return
}

// Execute @cf/openai/whisper model.
func (r *AISpeechRecognitionService) Whisper(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AISpeechRecognitionWhisperResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/openai/whisper", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AISpeechRecognitionWhisperResponse struct {
	Text string                                 `json:"text"`
	JSON aiSpeechRecognitionWhisperResponseJSON `json:"-"`
}

// aiSpeechRecognitionWhisperResponseJSON contains the JSON metadata for the struct
// [AISpeechRecognitionWhisperResponse]
type aiSpeechRecognitionWhisperResponseJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AISpeechRecognitionWhisperResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
