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

// AIModelTextGenerationMistral7bInstructV0_1Service contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIModelTextGenerationMistral7bInstructV0_1Service] method instead.
type AIModelTextGenerationMistral7bInstructV0_1Service struct {
	Options []option.RequestOption
}

// NewAIModelTextGenerationMistral7bInstructV0_1Service generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAIModelTextGenerationMistral7bInstructV0_1Service(opts ...option.RequestOption) (r *AIModelTextGenerationMistral7bInstructV0_1Service) {
	r = &AIModelTextGenerationMistral7bInstructV0_1Service{}
	r.Options = opts
	return
}

// Execute @cf/mistral/mistral-7b-instruct-v0.1 model.
func (r *AIModelTextGenerationMistral7bInstructV0_1Service) Run(ctx context.Context, accountIdentifier string, body AIModelTextGenerationMistral7bInstructV0_1RunParams, opts ...option.RequestOption) (res *AIModelTextGenerationMistral7bInstructV0_1RunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/mistral/mistral-7b-instruct-v0.1", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Union satisfied by [AIModelTextGenerationMistral7bInstructV0_1RunResponseObject]
// or [shared.UnionString].
type AIModelTextGenerationMistral7bInstructV0_1RunResponse interface {
	ImplementsAIModelTextGenerationMistral7bInstructV0_1RunResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIModelTextGenerationMistral7bInstructV0_1RunResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AIModelTextGenerationMistral7bInstructV0_1RunResponseObject struct {
	Response string                                                          `json:"response"`
	JSON     aiModelTextGenerationMistral7bInstructV0_1RunResponseObjectJSON `json:"-"`
}

// aiModelTextGenerationMistral7bInstructV0_1RunResponseObjectJSON contains the
// JSON metadata for the struct
// [AIModelTextGenerationMistral7bInstructV0_1RunResponseObject]
type aiModelTextGenerationMistral7bInstructV0_1RunResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelTextGenerationMistral7bInstructV0_1RunResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AIModelTextGenerationMistral7bInstructV0_1RunResponseObject) ImplementsAIModelTextGenerationMistral7bInstructV0_1RunResponse() {
}

// This interface is a union satisfied by one of the following:
// [AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant0],
// [AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant1].
type AIModelTextGenerationMistral7bInstructV0_1RunParams interface {
	ImplementsAIModelTextGenerationMistral7bInstructV0_1RunParams()
}

type AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant0) ImplementsAIModelTextGenerationMistral7bInstructV0_1RunParams() {

}

type AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant1 struct {
	Messages param.Field[[]AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant1Message] `json:"messages,required"`
}

func (r AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant1) ImplementsAIModelTextGenerationMistral7bInstructV0_1RunParams() {

}

type AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIModelTextGenerationMistral7bInstructV0_1RunParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
