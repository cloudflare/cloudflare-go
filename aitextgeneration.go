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

// AITextGenerationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAITextGenerationService] method
// instead.
type AITextGenerationService struct {
	Options []option.RequestOption
}

// NewAITextGenerationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITextGenerationService(opts ...option.RequestOption) (r *AITextGenerationService) {
	r = &AITextGenerationService{}
	r.Options = opts
	return
}

// Execute @cf/meta/llama-2-7b-chat-fp16 model.
func (r *AITextGenerationService) Llama2_7bChatFp16(ctx context.Context, accountIdentifier string, body AITextGenerationLlama2_7bChatFp16Params, opts ...option.RequestOption) (res *AITextGenerationLlama2_7bChatFp16Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-fp16", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/meta/llama-2-7b-chat-int8 model.
func (r *AITextGenerationService) Llama2_7bChatInt8(ctx context.Context, accountIdentifier string, body AITextGenerationLlama2_7bChatInt8Params, opts ...option.RequestOption) (res *AITextGenerationLlama2_7bChatInt8Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-int8", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/mistral/mistral-7b-instruct-v0.1 model.
func (r *AITextGenerationService) Mistral7bInstructV0_1(ctx context.Context, accountIdentifier string, body AITextGenerationMistral7bInstructV0_1Params, opts ...option.RequestOption) (res *AITextGenerationMistral7bInstructV0_1Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/mistral/mistral-7b-instruct-v0.1", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Union satisfied by [AITextGenerationLlama2_7bChatFp16ResponseObject] or
// [shared.UnionString].
type AITextGenerationLlama2_7bChatFp16Response interface {
	ImplementsAITextGenerationLlama2_7bChatFp16Response()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AITextGenerationLlama2_7bChatFp16Response)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AITextGenerationLlama2_7bChatFp16ResponseObject struct {
	Response string                                              `json:"response"`
	JSON     aiTextGenerationLlama2_7bChatFp16ResponseObjectJSON `json:"-"`
}

// aiTextGenerationLlama2_7bChatFp16ResponseObjectJSON contains the JSON metadata
// for the struct [AITextGenerationLlama2_7bChatFp16ResponseObject]
type aiTextGenerationLlama2_7bChatFp16ResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITextGenerationLlama2_7bChatFp16ResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AITextGenerationLlama2_7bChatFp16ResponseObject) ImplementsAITextGenerationLlama2_7bChatFp16Response() {
}

// Union satisfied by [AITextGenerationLlama2_7bChatInt8ResponseObject] or
// [shared.UnionString].
type AITextGenerationLlama2_7bChatInt8Response interface {
	ImplementsAITextGenerationLlama2_7bChatInt8Response()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AITextGenerationLlama2_7bChatInt8Response)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AITextGenerationLlama2_7bChatInt8ResponseObject struct {
	Response string                                              `json:"response"`
	JSON     aiTextGenerationLlama2_7bChatInt8ResponseObjectJSON `json:"-"`
}

// aiTextGenerationLlama2_7bChatInt8ResponseObjectJSON contains the JSON metadata
// for the struct [AITextGenerationLlama2_7bChatInt8ResponseObject]
type aiTextGenerationLlama2_7bChatInt8ResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITextGenerationLlama2_7bChatInt8ResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AITextGenerationLlama2_7bChatInt8ResponseObject) ImplementsAITextGenerationLlama2_7bChatInt8Response() {
}

// Union satisfied by [AITextGenerationMistral7bInstructV0_1ResponseObject] or
// [shared.UnionString].
type AITextGenerationMistral7bInstructV0_1Response interface {
	ImplementsAITextGenerationMistral7bInstructV0_1Response()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AITextGenerationMistral7bInstructV0_1Response)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AITextGenerationMistral7bInstructV0_1ResponseObject struct {
	Response string                                                  `json:"response"`
	JSON     aiTextGenerationMistral7bInstructV0_1ResponseObjectJSON `json:"-"`
}

// aiTextGenerationMistral7bInstructV0_1ResponseObjectJSON contains the JSON
// metadata for the struct [AITextGenerationMistral7bInstructV0_1ResponseObject]
type aiTextGenerationMistral7bInstructV0_1ResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITextGenerationMistral7bInstructV0_1ResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AITextGenerationMistral7bInstructV0_1ResponseObject) ImplementsAITextGenerationMistral7bInstructV0_1Response() {
}

// This interface is a union satisfied by one of the following:
// [AITextGenerationLlama2_7bChatFp16ParamsVariant0],
// [AITextGenerationLlama2_7bChatFp16ParamsVariant1].
type AITextGenerationLlama2_7bChatFp16Params interface {
	ImplementsAITextGenerationLlama2_7bChatFp16Params()
}

type AITextGenerationLlama2_7bChatFp16ParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AITextGenerationLlama2_7bChatFp16ParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AITextGenerationLlama2_7bChatFp16ParamsVariant0) ImplementsAITextGenerationLlama2_7bChatFp16Params() {

}

type AITextGenerationLlama2_7bChatFp16ParamsVariant1 struct {
	Messages param.Field[[]AITextGenerationLlama2_7bChatFp16ParamsVariant1Message] `json:"messages,required"`
}

func (r AITextGenerationLlama2_7bChatFp16ParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AITextGenerationLlama2_7bChatFp16ParamsVariant1) ImplementsAITextGenerationLlama2_7bChatFp16Params() {

}

type AITextGenerationLlama2_7bChatFp16ParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AITextGenerationLlama2_7bChatFp16ParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This interface is a union satisfied by one of the following:
// [AITextGenerationLlama2_7bChatInt8ParamsVariant0],
// [AITextGenerationLlama2_7bChatInt8ParamsVariant1].
type AITextGenerationLlama2_7bChatInt8Params interface {
	ImplementsAITextGenerationLlama2_7bChatInt8Params()
}

type AITextGenerationLlama2_7bChatInt8ParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AITextGenerationLlama2_7bChatInt8ParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AITextGenerationLlama2_7bChatInt8ParamsVariant0) ImplementsAITextGenerationLlama2_7bChatInt8Params() {

}

type AITextGenerationLlama2_7bChatInt8ParamsVariant1 struct {
	Messages param.Field[[]AITextGenerationLlama2_7bChatInt8ParamsVariant1Message] `json:"messages,required"`
}

func (r AITextGenerationLlama2_7bChatInt8ParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AITextGenerationLlama2_7bChatInt8ParamsVariant1) ImplementsAITextGenerationLlama2_7bChatInt8Params() {

}

type AITextGenerationLlama2_7bChatInt8ParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AITextGenerationLlama2_7bChatInt8ParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This interface is a union satisfied by one of the following:
// [AITextGenerationMistral7bInstructV0_1ParamsVariant0],
// [AITextGenerationMistral7bInstructV0_1ParamsVariant1].
type AITextGenerationMistral7bInstructV0_1Params interface {
	ImplementsAITextGenerationMistral7bInstructV0_1Params()
}

type AITextGenerationMistral7bInstructV0_1ParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AITextGenerationMistral7bInstructV0_1ParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AITextGenerationMistral7bInstructV0_1ParamsVariant0) ImplementsAITextGenerationMistral7bInstructV0_1Params() {

}

type AITextGenerationMistral7bInstructV0_1ParamsVariant1 struct {
	Messages param.Field[[]AITextGenerationMistral7bInstructV0_1ParamsVariant1Message] `json:"messages,required"`
}

func (r AITextGenerationMistral7bInstructV0_1ParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AITextGenerationMistral7bInstructV0_1ParamsVariant1) ImplementsAITextGenerationMistral7bInstructV0_1Params() {

}

type AITextGenerationMistral7bInstructV0_1ParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AITextGenerationMistral7bInstructV0_1ParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
