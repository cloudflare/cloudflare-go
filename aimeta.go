// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AIMetaService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAIMetaService] method instead.
type AIMetaService struct {
	Options []option.RequestOption
}

// NewAIMetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIMetaService(opts ...option.RequestOption) (r *AIMetaService) {
	r = &AIMetaService{}
	r.Options = opts
	return
}

// Execute @cf/meta/llama-2-7b-chat-fp16 model.
func (r *AIMetaService) Llama2_7bChatFp16(ctx context.Context, accountIdentifier string, body AIMetaLlama2_7bChatFp16Params, opts ...option.RequestOption) (res *AIMetaLlama2_7bChatFp16Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-fp16", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/meta/llama-2-7b-chat-int8 model.
func (r *AIMetaService) Llama2_7bChatInt8(ctx context.Context, accountIdentifier string, body AIMetaLlama2_7bChatInt8Params, opts ...option.RequestOption) (res *AIMetaLlama2_7bChatInt8Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-int8", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/meta/m2m100-1.2b model.
func (r *AIMetaService) M2m100_1_2b(ctx context.Context, accountIdentifier string, body AIMetaM2m100_1_2bParams, opts ...option.RequestOption) (res *AIMetaM2m100_1_2bResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/m2m100-1.2b", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Union satisfied by [AIMetaLlama2_7bChatFp16ResponseObject] or
// [shared.UnionString].
type AIMetaLlama2_7bChatFp16Response interface {
	ImplementsAIMetaLlama2_7bChatFp16Response()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIMetaLlama2_7bChatFp16Response)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AIMetaLlama2_7bChatFp16ResponseObject struct {
	Response string                                    `json:"response"`
	JSON     aiMetaLlama2_7bChatFp16ResponseObjectJSON `json:"-"`
}

// aiMetaLlama2_7bChatFp16ResponseObjectJSON contains the JSON metadata for the
// struct [AIMetaLlama2_7bChatFp16ResponseObject]
type aiMetaLlama2_7bChatFp16ResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMetaLlama2_7bChatFp16ResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AIMetaLlama2_7bChatFp16ResponseObject) ImplementsAIMetaLlama2_7bChatFp16Response() {}

// Union satisfied by [AIMetaLlama2_7bChatInt8ResponseObject] or
// [shared.UnionString].
type AIMetaLlama2_7bChatInt8Response interface {
	ImplementsAIMetaLlama2_7bChatInt8Response()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIMetaLlama2_7bChatInt8Response)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AIMetaLlama2_7bChatInt8ResponseObject struct {
	Response string                                    `json:"response"`
	JSON     aiMetaLlama2_7bChatInt8ResponseObjectJSON `json:"-"`
}

// aiMetaLlama2_7bChatInt8ResponseObjectJSON contains the JSON metadata for the
// struct [AIMetaLlama2_7bChatInt8ResponseObject]
type aiMetaLlama2_7bChatInt8ResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMetaLlama2_7bChatInt8ResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AIMetaLlama2_7bChatInt8ResponseObject) ImplementsAIMetaLlama2_7bChatInt8Response() {}

type AIMetaM2m100_1_2bResponse struct {
	TranslatedText string                        `json:"translated_text"`
	JSON           aiMetaM2m100_1_2bResponseJSON `json:"-"`
}

// aiMetaM2m100_1_2bResponseJSON contains the JSON metadata for the struct
// [AIMetaM2m100_1_2bResponse]
type aiMetaM2m100_1_2bResponseJSON struct {
	TranslatedText apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIMetaM2m100_1_2bResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This interface is a union satisfied by one of the following:
// [AIMetaLlama2_7bChatFp16ParamsVariant0],
// [AIMetaLlama2_7bChatFp16ParamsVariant1].
type AIMetaLlama2_7bChatFp16Params interface {
	ImplementsAIMetaLlama2_7bChatFp16Params()
}

type AIMetaLlama2_7bChatFp16ParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AIMetaLlama2_7bChatFp16ParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIMetaLlama2_7bChatFp16ParamsVariant0) ImplementsAIMetaLlama2_7bChatFp16Params() {

}

type AIMetaLlama2_7bChatFp16ParamsVariant1 struct {
	Messages param.Field[[]AIMetaLlama2_7bChatFp16ParamsVariant1Message] `json:"messages,required"`
}

func (r AIMetaLlama2_7bChatFp16ParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIMetaLlama2_7bChatFp16ParamsVariant1) ImplementsAIMetaLlama2_7bChatFp16Params() {

}

type AIMetaLlama2_7bChatFp16ParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIMetaLlama2_7bChatFp16ParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This interface is a union satisfied by one of the following:
// [AIMetaLlama2_7bChatInt8ParamsVariant0],
// [AIMetaLlama2_7bChatInt8ParamsVariant1].
type AIMetaLlama2_7bChatInt8Params interface {
	ImplementsAIMetaLlama2_7bChatInt8Params()
}

type AIMetaLlama2_7bChatInt8ParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AIMetaLlama2_7bChatInt8ParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIMetaLlama2_7bChatInt8ParamsVariant0) ImplementsAIMetaLlama2_7bChatInt8Params() {

}

type AIMetaLlama2_7bChatInt8ParamsVariant1 struct {
	Messages param.Field[[]AIMetaLlama2_7bChatInt8ParamsVariant1Message] `json:"messages,required"`
}

func (r AIMetaLlama2_7bChatInt8ParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIMetaLlama2_7bChatInt8ParamsVariant1) ImplementsAIMetaLlama2_7bChatInt8Params() {

}

type AIMetaLlama2_7bChatInt8ParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIMetaLlama2_7bChatInt8ParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIMetaM2m100_1_2bParams struct {
	TargetLang param.Field[string] `json:"target_lang,required"`
	Text       param.Field[string] `json:"text,required"`
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AIMetaM2m100_1_2bParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
