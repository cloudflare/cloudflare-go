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

// AITranslationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAITranslationService] method
// instead.
type AITranslationService struct {
	Options []option.RequestOption
}

// NewAITranslationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAITranslationService(opts ...option.RequestOption) (r *AITranslationService) {
	r = &AITranslationService{}
	r.Options = opts
	return
}

// Execute @cf/meta/m2m100-1.2b model.
func (r *AITranslationService) M2m100_1_2b(ctx context.Context, accountIdentifier string, body AITranslationM2m100_1_2bParams, opts ...option.RequestOption) (res *AITranslationM2m100_1_2bResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/m2m100-1.2b", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AITranslationM2m100_1_2bResponse struct {
	TranslatedText string                               `json:"translated_text"`
	JSON           aiTranslationM2m100_1_2bResponseJSON `json:"-"`
}

// aiTranslationM2m100_1_2bResponseJSON contains the JSON metadata for the struct
// [AITranslationM2m100_1_2bResponse]
type aiTranslationM2m100_1_2bResponseJSON struct {
	TranslatedText apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AITranslationM2m100_1_2bResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AITranslationM2m100_1_2bParams struct {
	TargetLang param.Field[string] `json:"target_lang,required"`
	Text       param.Field[string] `json:"text,required"`
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AITranslationM2m100_1_2bParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
