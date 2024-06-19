// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
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

// LogService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r *LogService) {
	r = &LogService{}
	r.Options = opts
	return
}

// List Gateway Logs
func (r *LogService) Get(ctx context.Context, id string, params LogGetParams, opts ...option.RequestOption) (res *[]LogGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogGetResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogGetResponse struct {
	ID                  string             `json:"id,required"`
	Cached              bool               `json:"cached,required"`
	CreatedAt           time.Time          `json:"created_at,required" format:"date-time"`
	Duration            int64              `json:"duration,required"`
	Model               string             `json:"model,required"`
	Path                string             `json:"path,required"`
	Provider            string             `json:"provider,required"`
	Request             string             `json:"request,required"`
	Response            string             `json:"response,required"`
	Success             bool               `json:"success,required"`
	TokensIn            int64              `json:"tokens_in,required"`
	TokensOut           int64              `json:"tokens_out,required"`
	Metadata            string             `json:"metadata"`
	RequestContentType  string             `json:"request_content_type"`
	RequestType         string             `json:"request_type"`
	ResponseContentType string             `json:"response_content_type"`
	StatusCode          int64              `json:"status_code"`
	Step                int64              `json:"step"`
	JSON                logGetResponseJSON `json:"-"`
}

// logGetResponseJSON contains the JSON metadata for the struct [LogGetResponse]
type logGetResponseJSON struct {
	ID                  apijson.Field
	Cached              apijson.Field
	CreatedAt           apijson.Field
	Duration            apijson.Field
	Model               apijson.Field
	Path                apijson.Field
	Provider            apijson.Field
	Request             apijson.Field
	Response            apijson.Field
	Success             apijson.Field
	TokensIn            apijson.Field
	TokensOut           apijson.Field
	Metadata            apijson.Field
	RequestContentType  apijson.Field
	RequestType         apijson.Field
	ResponseContentType apijson.Field
	StatusCode          apijson.Field
	Step                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logGetResponseJSON) RawJSON() string {
	return r.raw
}

type LogGetParams struct {
	AccountID param.Field[string]                `path:"account_id,required"`
	Cached    param.Field[bool]                  `query:"cached"`
	Direction param.Field[LogGetParamsDirection] `query:"direction"`
	EndDate   param.Field[time.Time]             `query:"end_date" format:"date-time"`
	OrderBy   param.Field[LogGetParamsOrderBy]   `query:"order_by"`
	Page      param.Field[int64]                 `query:"page"`
	PerPage   param.Field[int64]                 `query:"per_page"`
	Search    param.Field[string]                `query:"search"`
	StartDate param.Field[time.Time]             `query:"start_date" format:"date-time"`
	Success   param.Field[bool]                  `query:"success"`
}

// URLQuery serializes [LogGetParams]'s query parameters as `url.Values`.
func (r LogGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LogGetParamsDirection string

const (
	LogGetParamsDirectionAsc  LogGetParamsDirection = "asc"
	LogGetParamsDirectionDesc LogGetParamsDirection = "desc"
)

func (r LogGetParamsDirection) IsKnown() bool {
	switch r {
	case LogGetParamsDirectionAsc, LogGetParamsDirectionDesc:
		return true
	}
	return false
}

type LogGetParamsOrderBy string

const (
	LogGetParamsOrderByCreatedAt LogGetParamsOrderBy = "created_at"
	LogGetParamsOrderByProvider  LogGetParamsOrderBy = "provider"
)

func (r LogGetParamsOrderBy) IsKnown() bool {
	switch r {
	case LogGetParamsOrderByCreatedAt, LogGetParamsOrderByProvider:
		return true
	}
	return false
}

type LogGetResponseEnvelope struct {
	Result  []LogGetResponse           `json:"result,required"`
	Success bool                       `json:"success,required"`
	JSON    logGetResponseEnvelopeJSON `json:"-"`
}

// logGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogGetResponseEnvelope]
type logGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
