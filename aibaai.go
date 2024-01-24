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

// AIBaaiService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAIBaaiService] method instead.
type AIBaaiService struct {
	Options []option.RequestOption
}

// NewAIBaaiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIBaaiService(opts ...option.RequestOption) (r *AIBaaiService) {
	r = &AIBaaiService{}
	r.Options = opts
	return
}

// Execute @cf/baai/bge-base-en-v1.5 model.
func (r *AIBaaiService) BgeBaseEnV1_5(ctx context.Context, accountIdentifier string, body AIBaaiBgeBaseEnV1_5Params, opts ...option.RequestOption) (res *AIBaaiBgeBaseEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-base-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/baai/bge-large-en-v1.5 model.
func (r *AIBaaiService) BgeLargeEnV1_5(ctx context.Context, accountIdentifier string, body AIBaaiBgeLargeEnV1_5Params, opts ...option.RequestOption) (res *AIBaaiBgeLargeEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-large-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Execute @cf/baai/bge-small-en-v1.5 model.
func (r *AIBaaiService) BgeSmallEnV1_5(ctx context.Context, accountIdentifier string, body AIBaaiBgeSmallEnV1_5Params, opts ...option.RequestOption) (res *AIBaaiBgeSmallEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apiv4/accounts/%s/ai/run/@cf/baai/bge-small-en-v1.5", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AIBaaiBgeBaseEnV1_5Response struct {
	Data  [][]float64                     `json:"data"`
	Shape []float64                       `json:"shape"`
	JSON  aiBaaiBgeBaseEnV1_5ResponseJSON `json:"-"`
}

// aiBaaiBgeBaseEnV1_5ResponseJSON contains the JSON metadata for the struct
// [AIBaaiBgeBaseEnV1_5Response]
type aiBaaiBgeBaseEnV1_5ResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBaaiBgeBaseEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIBaaiBgeLargeEnV1_5Response struct {
	Data  [][]float64                      `json:"data"`
	Shape []float64                        `json:"shape"`
	JSON  aiBaaiBgeLargeEnV1_5ResponseJSON `json:"-"`
}

// aiBaaiBgeLargeEnV1_5ResponseJSON contains the JSON metadata for the struct
// [AIBaaiBgeLargeEnV1_5Response]
type aiBaaiBgeLargeEnV1_5ResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBaaiBgeLargeEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIBaaiBgeSmallEnV1_5Response struct {
	Data  [][]float64                      `json:"data"`
	Shape []float64                        `json:"shape"`
	JSON  aiBaaiBgeSmallEnV1_5ResponseJSON `json:"-"`
}

// aiBaaiBgeSmallEnV1_5ResponseJSON contains the JSON metadata for the struct
// [AIBaaiBgeSmallEnV1_5Response]
type aiBaaiBgeSmallEnV1_5ResponseJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBaaiBgeSmallEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AIBaaiBgeBaseEnV1_5Params struct {
	Text param.Field[AIBaaiBgeBaseEnV1_5ParamsText] `json:"text,required"`
}

func (r AIBaaiBgeBaseEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [AIBaaiBgeBaseEnV1_5ParamsTextArray].
type AIBaaiBgeBaseEnV1_5ParamsText interface {
	ImplementsAIBaaiBgeBaseEnV1_5ParamsText()
}

type AIBaaiBgeBaseEnV1_5ParamsTextArray []string

func (r AIBaaiBgeBaseEnV1_5ParamsTextArray) ImplementsAIBaaiBgeBaseEnV1_5ParamsText() {}

type AIBaaiBgeLargeEnV1_5Params struct {
	Text param.Field[AIBaaiBgeLargeEnV1_5ParamsText] `json:"text,required"`
}

func (r AIBaaiBgeLargeEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [AIBaaiBgeLargeEnV1_5ParamsTextArray].
type AIBaaiBgeLargeEnV1_5ParamsText interface {
	ImplementsAIBaaiBgeLargeEnV1_5ParamsText()
}

type AIBaaiBgeLargeEnV1_5ParamsTextArray []string

func (r AIBaaiBgeLargeEnV1_5ParamsTextArray) ImplementsAIBaaiBgeLargeEnV1_5ParamsText() {}

type AIBaaiBgeSmallEnV1_5Params struct {
	Text param.Field[AIBaaiBgeSmallEnV1_5ParamsText] `json:"text,required"`
}

func (r AIBaaiBgeSmallEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [AIBaaiBgeSmallEnV1_5ParamsTextArray].
type AIBaaiBgeSmallEnV1_5ParamsText interface {
	ImplementsAIBaaiBgeSmallEnV1_5ParamsText()
}

type AIBaaiBgeSmallEnV1_5ParamsTextArray []string

func (r AIBaaiBgeSmallEnV1_5ParamsTextArray) ImplementsAIBaaiBgeSmallEnV1_5ParamsText() {}
