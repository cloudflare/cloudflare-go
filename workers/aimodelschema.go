// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AIModelSchemaService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIModelSchemaService] method instead.
type AIModelSchemaService struct {
	Options []option.RequestOption
}

// NewAIModelSchemaService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIModelSchemaService(opts ...option.RequestOption) (r *AIModelSchemaService) {
	r = &AIModelSchemaService{}
	r.Options = opts
	return
}

// Get Model Schema
func (r *AIModelSchemaService) Get(ctx context.Context, params AIModelSchemaGetParams, opts ...option.RequestOption) (res *AIModelSchemaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AIModelSchemaGetResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/models/schema", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIModelSchemaGetResponse = interface{}

type AIModelSchemaGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Model Name
	Model param.Field[string] `query:"model,required"`
}

// URLQuery serializes [AIModelSchemaGetParams]'s query parameters as `url.Values`.
func (r AIModelSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AIModelSchemaGetResponseEnvelope struct {
	Result  AIModelSchemaGetResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    aiModelSchemaGetResponseEnvelopeJSON `json:"-"`
}

// aiModelSchemaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIModelSchemaGetResponseEnvelope]
type aiModelSchemaGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIModelSchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiModelSchemaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
