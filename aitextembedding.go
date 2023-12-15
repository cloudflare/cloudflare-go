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

// AITextEmbeddingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAITextEmbeddingService] method
// instead.
type AITextEmbeddingService struct {
	Options []option.RequestOption
}

// NewAITextEmbeddingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAITextEmbeddingService(opts ...option.RequestOption) (r *AITextEmbeddingService) {
	r = &AITextEmbeddingService{}
	r.Options = opts
	return
}

// Execute @cf/baai/bge-base-en-v1.5 model.
func (r *AITextEmbeddingService) BgeBaseEnV1_5(ctx context.Context, accountIdentifier string, body AITextEmbeddingBgeBaseEnV1_5Params, opts ...option.RequestOption) (res *AITextEmbeddingBgeBaseEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-base-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/baai/bge-large-en-v1.5 model.
func (r *AITextEmbeddingService) BgeLargeEnV1_5(ctx context.Context, accountIdentifier string, body AITextEmbeddingBgeLargeEnV1_5Params, opts ...option.RequestOption) (res *AITextEmbeddingBgeLargeEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-large-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/baai/bge-small-en-v1.5 model.
func (r *AITextEmbeddingService) BgeSmallEnV1_5(ctx context.Context, accountIdentifier string, body AITextEmbeddingBgeSmallEnV1_5Params, opts ...option.RequestOption) (res *AITextEmbeddingBgeSmallEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-small-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AITextEmbeddingBgeBaseEnV1_5Response struct {
	Data  [][]float64                              `json:"data"`
	Shape []float64                                `json:"shape"`
	JSON  aiTextEmbeddingBgeBaseEnV1_5ResponseJSON `json:"-"`
}

// aiTextEmbeddingBgeBaseEnV1_5ResponseJSON contains the JSON metadata for the
// struct [AITextEmbeddingBgeBaseEnV1_5Response]
type aiTextEmbeddingBgeBaseEnV1_5ResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITextEmbeddingBgeBaseEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AITextEmbeddingBgeLargeEnV1_5Response struct {
	Data  [][]float64                               `json:"data"`
	Shape []float64                                 `json:"shape"`
	JSON  aiTextEmbeddingBgeLargeEnV1_5ResponseJSON `json:"-"`
}

// aiTextEmbeddingBgeLargeEnV1_5ResponseJSON contains the JSON metadata for the
// struct [AITextEmbeddingBgeLargeEnV1_5Response]
type aiTextEmbeddingBgeLargeEnV1_5ResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITextEmbeddingBgeLargeEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AITextEmbeddingBgeSmallEnV1_5Response struct {
	Data  [][]float64                               `json:"data"`
	Shape []float64                                 `json:"shape"`
	JSON  aiTextEmbeddingBgeSmallEnV1_5ResponseJSON `json:"-"`
}

// aiTextEmbeddingBgeSmallEnV1_5ResponseJSON contains the JSON metadata for the
// struct [AITextEmbeddingBgeSmallEnV1_5Response]
type aiTextEmbeddingBgeSmallEnV1_5ResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITextEmbeddingBgeSmallEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AITextEmbeddingBgeBaseEnV1_5Params struct {
	Text param.Field[AITextEmbeddingBgeBaseEnV1_5ParamsText] `json:"text,required"`
}

func (r AITextEmbeddingBgeBaseEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [AITextEmbeddingBgeBaseEnV1_5ParamsTextObject].
type AITextEmbeddingBgeBaseEnV1_5ParamsText interface {
	ImplementsAITextEmbeddingBgeBaseEnV1_5ParamsText()
}

type AITextEmbeddingBgeLargeEnV1_5Params struct {
	Text param.Field[AITextEmbeddingBgeLargeEnV1_5ParamsText] `json:"text,required"`
}

func (r AITextEmbeddingBgeLargeEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [AITextEmbeddingBgeLargeEnV1_5ParamsTextObject].
type AITextEmbeddingBgeLargeEnV1_5ParamsText interface {
	ImplementsAITextEmbeddingBgeLargeEnV1_5ParamsText()
}

type AITextEmbeddingBgeSmallEnV1_5Params struct {
	Text param.Field[AITextEmbeddingBgeSmallEnV1_5ParamsText] `json:"text,required"`
}

func (r AITextEmbeddingBgeSmallEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [AITextEmbeddingBgeSmallEnV1_5ParamsTextObject].
type AITextEmbeddingBgeSmallEnV1_5ParamsText interface {
	ImplementsAITextEmbeddingBgeSmallEnV1_5ParamsText()
}
