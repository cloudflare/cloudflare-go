// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// EvaluationTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvaluationTypeService] method instead.
type EvaluationTypeService struct {
	Options []option.RequestOption
}

// NewEvaluationTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvaluationTypeService(opts ...option.RequestOption) (r *EvaluationTypeService) {
	r = &EvaluationTypeService{}
	r.Options = opts
	return
}

// List Evaluators
func (r *EvaluationTypeService) Get(ctx context.Context, params EvaluationTypeGetParams, opts ...option.RequestOption) (res *[]EvaluationTypeGetResponse, err error) {
	var env EvaluationTypeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/evaluation-types", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EvaluationTypeGetResponse struct {
	ID          string                        `json:"id,required"`
	CreatedAt   time.Time                     `json:"created_at,required" format:"date-time"`
	Description string                        `json:"description,required"`
	Enable      bool                          `json:"enable,required"`
	Mandatory   bool                          `json:"mandatory,required"`
	ModifiedAt  time.Time                     `json:"modified_at,required" format:"date-time"`
	Name        string                        `json:"name,required"`
	Type        string                        `json:"type,required"`
	JSON        evaluationTypeGetResponseJSON `json:"-"`
}

// evaluationTypeGetResponseJSON contains the JSON metadata for the struct
// [EvaluationTypeGetResponse]
type evaluationTypeGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Enable      apijson.Field
	Mandatory   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationTypeGetResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationTypeGetParams struct {
	AccountID        param.Field[string]                                  `path:"account_id,required"`
	OrderBy          param.Field[string]                                  `query:"order_by"`
	OrderByDirection param.Field[EvaluationTypeGetParamsOrderByDirection] `query:"order_by_direction"`
	Page             param.Field[int64]                                   `query:"page"`
	PerPage          param.Field[int64]                                   `query:"per_page"`
}

// URLQuery serializes [EvaluationTypeGetParams]'s query parameters as
// `url.Values`.
func (r EvaluationTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EvaluationTypeGetParamsOrderByDirection string

const (
	EvaluationTypeGetParamsOrderByDirectionAsc  EvaluationTypeGetParamsOrderByDirection = "asc"
	EvaluationTypeGetParamsOrderByDirectionDesc EvaluationTypeGetParamsOrderByDirection = "desc"
)

func (r EvaluationTypeGetParamsOrderByDirection) IsKnown() bool {
	switch r {
	case EvaluationTypeGetParamsOrderByDirectionAsc, EvaluationTypeGetParamsOrderByDirectionDesc:
		return true
	}
	return false
}

type EvaluationTypeGetResponseEnvelope struct {
	Result     []EvaluationTypeGetResponse                 `json:"result,required"`
	ResultInfo EvaluationTypeGetResponseEnvelopeResultInfo `json:"result_info,required"`
	Success    bool                                        `json:"success,required"`
	JSON       evaluationTypeGetResponseEnvelopeJSON       `json:"-"`
}

// evaluationTypeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EvaluationTypeGetResponseEnvelope]
type evaluationTypeGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationTypeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EvaluationTypeGetResponseEnvelopeResultInfo struct {
	Count      float64                                         `json:"count,required"`
	Page       float64                                         `json:"page,required"`
	PerPage    float64                                         `json:"per_page,required"`
	TotalCount float64                                         `json:"total_count,required"`
	JSON       evaluationTypeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// evaluationTypeGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [EvaluationTypeGetResponseEnvelopeResultInfo]
type evaluationTypeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationTypeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationTypeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
