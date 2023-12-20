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

// AIModelTextGenerationLlama2_7bChatInt8Service contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIModelTextGenerationLlama2_7bChatInt8Service] method instead.
type AIModelTextGenerationLlama2_7bChatInt8Service struct {
	Options []option.RequestOption
}

// NewAIModelTextGenerationLlama2_7bChatInt8Service generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAIModelTextGenerationLlama2_7bChatInt8Service(opts ...option.RequestOption) (r *AIModelTextGenerationLlama2_7bChatInt8Service) {
	r = &AIModelTextGenerationLlama2_7bChatInt8Service{}
	r.Options = opts
	return
}

// Execute @cf/meta/llama-2-7b-chat-int8 model.
func (r *AIModelTextGenerationLlama2_7bChatInt8Service) Run(ctx context.Context, accountIdentifier string, body AIModelTextGenerationLlama2_7bChatInt8RunParams, opts ...option.RequestOption) (res *AIModelTextGenerationLlama2_7bChatInt8RunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-int8", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Union satisfied by [AIModelTextGenerationLlama2_7bChatInt8RunResponseObject] or
// [shared.UnionString].
type AIModelTextGenerationLlama2_7bChatInt8RunResponse interface {
	ImplementsAIModelTextGenerationLlama2_7bChatInt8RunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIModelTextGenerationLlama2_7bChatInt8RunResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AIModelTextGenerationLlama2_7bChatInt8RunResponseObject struct {
	Response string                                                      `json:"response"`
	JSON     aiModelTextGenerationLlama2_7bChatInt8RunResponseObjectJSON `json:"-"`
}

// aiModelTextGenerationLlama2_7bChatInt8RunResponseObjectJSON contains the JSON
// metadata for the struct
// [AIModelTextGenerationLlama2_7bChatInt8RunResponseObject]
type aiModelTextGenerationLlama2_7bChatInt8RunResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelTextGenerationLlama2_7bChatInt8RunResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AIModelTextGenerationLlama2_7bChatInt8RunResponseObject) ImplementsAIModelTextGenerationLlama2_7bChatInt8RunResponse() {
}

// This interface is a union satisfied by one of the following:
// [AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant0],
// [AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant1].
type AIModelTextGenerationLlama2_7bChatInt8RunParams interface {
	ImplementsAIModelTextGenerationLlama2_7bChatInt8RunParams()
}

type AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant0) ImplementsAIModelTextGenerationLlama2_7bChatInt8RunParams() {

}

type AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant1 struct {
	Messages param.Field[[]AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant1Message] `json:"messages,required"`
}

func (r AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant1) ImplementsAIModelTextGenerationLlama2_7bChatInt8RunParams() {

}

type AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIModelTextGenerationLlama2_7bChatInt8RunParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
