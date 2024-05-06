// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AIGatewayLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAIGatewayLogService] method
// instead.
type AIGatewayLogService struct {
	Options []option.RequestOption
}

// NewAIGatewayLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIGatewayLogService(opts ...option.RequestOption) (r *AIGatewayLogService) {
	r = &AIGatewayLogService{}
	r.Options = opts
	return
}

// List Gateway Logs
func (r *AIGatewayLogService) Get(ctx context.Context, accountTag string, id string, query AIGatewayLogGetParams, opts ...option.RequestOption) (res *[]AIGatewayLogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AIGatewayLogGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs", accountTag, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIGatewayLogGetResponse struct {
	ID        string                      `json:"id,required" format:"uuid"`
	Cached    bool                        `json:"cached,required"`
	CreatedAt time.Time                   `json:"created_at,required" format:"date-time"`
	Duration  int64                       `json:"duration,required"`
	Model     string                      `json:"model,required"`
	Path      string                      `json:"path,required"`
	Provider  string                      `json:"provider,required"`
	Request   string                      `json:"request,required"`
	Response  string                      `json:"response,required"`
	Success   bool                        `json:"success,required"`
	TokensIn  int64                       `json:"tokens_in,required"`
	TokensOut int64                       `json:"tokens_out,required"`
	JSON      aiGatewayLogGetResponseJSON `json:"-"`
}

// aiGatewayLogGetResponseJSON contains the JSON metadata for the struct
// [AIGatewayLogGetResponse]
type aiGatewayLogGetResponseJSON struct {
	ID          apijson.Field
	Cached      apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Model       apijson.Field
	Path        apijson.Field
	Provider    apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	Success     apijson.Field
	TokensIn    apijson.Field
	TokensOut   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayLogGetResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayLogGetParams struct {
	Cached    param.Field[bool]                           `query:"cached"`
	Direction param.Field[AIGatewayLogGetParamsDirection] `query:"direction"`
	EndDate   param.Field[time.Time]                      `query:"end_date" format:"date-time"`
	OrderBy   param.Field[AIGatewayLogGetParamsOrderBy]   `query:"order_by"`
	Page      param.Field[int64]                          `query:"page"`
	PerPage   param.Field[int64]                          `query:"per_page"`
	Search    param.Field[string]                         `query:"search"`
	StartDate param.Field[time.Time]                      `query:"start_date" format:"date-time"`
	Success   param.Field[bool]                           `query:"success"`
}

// URLQuery serializes [AIGatewayLogGetParams]'s query parameters as `url.Values`.
func (r AIGatewayLogGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AIGatewayLogGetParamsDirection string

const (
	AIGatewayLogGetParamsDirectionAsc  AIGatewayLogGetParamsDirection = "asc"
	AIGatewayLogGetParamsDirectionDesc AIGatewayLogGetParamsDirection = "desc"
)

func (r AIGatewayLogGetParamsDirection) IsKnown() bool {
	switch r {
	case AIGatewayLogGetParamsDirectionAsc, AIGatewayLogGetParamsDirectionDesc:
		return true
	}
	return false
}

type AIGatewayLogGetParamsOrderBy string

const (
	AIGatewayLogGetParamsOrderByCreatedAt AIGatewayLogGetParamsOrderBy = "created_at"
	AIGatewayLogGetParamsOrderByProvider  AIGatewayLogGetParamsOrderBy = "provider"
)

func (r AIGatewayLogGetParamsOrderBy) IsKnown() bool {
	switch r {
	case AIGatewayLogGetParamsOrderByCreatedAt, AIGatewayLogGetParamsOrderByProvider:
		return true
	}
	return false
}

type AIGatewayLogGetResponseEnvelope struct {
	Result  []AIGatewayLogGetResponse           `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiGatewayLogGetResponseEnvelopeJSON `json:"-"`
}

// aiGatewayLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayLogGetResponseEnvelope]
type aiGatewayLogGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
