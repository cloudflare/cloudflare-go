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

// AIModelTranslationM2m100_1_2bService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAIModelTranslationM2m100_1_2bService] method instead.
type AIModelTranslationM2m100_1_2bService struct {
	Options []option.RequestOption
}

// NewAIModelTranslationM2m100_1_2bService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIModelTranslationM2m100_1_2bService(opts ...option.RequestOption) (r *AIModelTranslationM2m100_1_2bService) {
	r = &AIModelTranslationM2m100_1_2bService{}
	r.Options = opts
	return
}

// Execute @cf/meta/m2m100-1.2b model.
func (r *AIModelTranslationM2m100_1_2bService) Run(ctx context.Context, accountIdentifier string, body AIModelTranslationM2m100_1_2bRunParams, opts ...option.RequestOption) (res *AIModelTranslationM2m100_1_2bRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/m2m100-1.2b", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIModelTranslationM2m100_1_2bRunResponse struct {
	TranslatedText string                                       `json:"translated_text"`
	JSON           aiModelTranslationM2m100_1_2bRunResponseJSON `json:"-"`
}

// aiModelTranslationM2m100_1_2bRunResponseJSON contains the JSON metadata for the
// struct [AIModelTranslationM2m100_1_2bRunResponse]
type aiModelTranslationM2m100_1_2bRunResponseJSON struct {
	TranslatedText apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIModelTranslationM2m100_1_2bRunResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIModelTranslationM2m100_1_2bRunParams struct {
	TargetLang param.Field[string] `json:"target_lang,required"`
	Text       param.Field[string] `json:"text,required"`
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AIModelTranslationM2m100_1_2bRunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
