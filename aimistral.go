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

// AIMistralService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAIMistralService] method instead.
type AIMistralService struct {
	Options []option.RequestOption
}

// NewAIMistralService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIMistralService(opts ...option.RequestOption) (r *AIMistralService) {
	r = &AIMistralService{}
	r.Options = opts
	return
}

// Execute @cf/mistral/mistral-7b-instruct-v0.1 model.
func (r *AIMistralService) Mistral7bInstructV0_1(ctx context.Context, accountIdentifier string, body AIMistralMistral7bInstructV0_1Params, opts ...option.RequestOption) (res *AIMistralMistral7bInstructV0_1Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/mistral/mistral-7b-instruct-v0.1", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Union satisfied by [AIMistralMistral7bInstructV0_1ResponseObject] or
// [shared.UnionString].
type AIMistralMistral7bInstructV0_1Response interface {
	ImplementsAIMistralMistral7bInstructV0_1Response()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AIMistralMistral7bInstructV0_1Response)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AIMistralMistral7bInstructV0_1ResponseObject struct {
	Response string                                           `json:"response"`
	JSON     aiMistralMistral7bInstructV0_1ResponseObjectJSON `json:"-"`
}

// aiMistralMistral7bInstructV0_1ResponseObjectJSON contains the JSON metadata for
// the struct [AIMistralMistral7bInstructV0_1ResponseObject]
type aiMistralMistral7bInstructV0_1ResponseObjectJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIMistralMistral7bInstructV0_1ResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AIMistralMistral7bInstructV0_1ResponseObject) ImplementsAIMistralMistral7bInstructV0_1Response() {
}

// This interface is a union satisfied by one of the following:
// [AIMistralMistral7bInstructV0_1ParamsVariant0],
// [AIMistralMistral7bInstructV0_1ParamsVariant1].
type AIMistralMistral7bInstructV0_1Params interface {
	ImplementsAIMistralMistral7bInstructV0_1Params()
}

type AIMistralMistral7bInstructV0_1ParamsVariant0 struct {
	Prompt param.Field[string] `json:"prompt,required"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r AIMistralMistral7bInstructV0_1ParamsVariant0) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIMistralMistral7bInstructV0_1ParamsVariant0) ImplementsAIMistralMistral7bInstructV0_1Params() {

}

type AIMistralMistral7bInstructV0_1ParamsVariant1 struct {
	Messages param.Field[[]AIMistralMistral7bInstructV0_1ParamsVariant1Message] `json:"messages,required"`
}

func (r AIMistralMistral7bInstructV0_1ParamsVariant1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (AIMistralMistral7bInstructV0_1ParamsVariant1) ImplementsAIMistralMistral7bInstructV0_1Params() {

}

type AIMistralMistral7bInstructV0_1ParamsVariant1Message struct {
	Content param.Field[string] `json:"content,required"`
	Role    param.Field[string] `json:"role,required"`
}

func (r AIMistralMistral7bInstructV0_1ParamsVariant1Message) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
