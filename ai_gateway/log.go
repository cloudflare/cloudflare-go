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
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
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
func (r *LogService) List(ctx context.Context, id string, params LogListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[LogListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs", params.AccountID, id)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Gateway Logs
func (r *LogService) ListAutoPaging(ctx context.Context, id string, params LogListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[LogListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, id, params, opts...))
}

type LogListResponse struct {
	ID                  string              `json:"id,required"`
	Cached              bool                `json:"cached,required"`
	CreatedAt           time.Time           `json:"created_at,required" format:"date-time"`
	Duration            int64               `json:"duration,required"`
	Model               string              `json:"model,required"`
	Path                string              `json:"path,required"`
	Provider            string              `json:"provider,required"`
	Request             string              `json:"request,required"`
	Response            string              `json:"response,required"`
	Success             bool                `json:"success,required"`
	TokensIn            int64               `json:"tokens_in,required"`
	TokensOut           int64               `json:"tokens_out,required"`
	Cost                float64             `json:"cost"`
	Metadata            string              `json:"metadata"`
	RequestContentType  string              `json:"request_content_type"`
	RequestType         string              `json:"request_type"`
	ResponseContentType string              `json:"response_content_type"`
	StatusCode          int64               `json:"status_code"`
	Step                int64               `json:"step"`
	JSON                logListResponseJSON `json:"-"`
}

// logListResponseJSON contains the JSON metadata for the struct [LogListResponse]
type logListResponseJSON struct {
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
	Cost                apijson.Field
	Metadata            apijson.Field
	RequestContentType  apijson.Field
	RequestType         apijson.Field
	ResponseContentType apijson.Field
	StatusCode          apijson.Field
	Step                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logListResponseJSON) RawJSON() string {
	return r.raw
}

type LogListParams struct {
	AccountID param.Field[string]                 `path:"account_id,required"`
	Cached    param.Field[bool]                   `query:"cached"`
	Direction param.Field[LogListParamsDirection] `query:"direction"`
	EndDate   param.Field[time.Time]              `query:"end_date" format:"date-time"`
	OrderBy   param.Field[LogListParamsOrderBy]   `query:"order_by"`
	Page      param.Field[int64]                  `query:"page"`
	PerPage   param.Field[int64]                  `query:"per_page"`
	Search    param.Field[string]                 `query:"search"`
	StartDate param.Field[time.Time]              `query:"start_date" format:"date-time"`
	Success   param.Field[bool]                   `query:"success"`
}

// URLQuery serializes [LogListParams]'s query parameters as `url.Values`.
func (r LogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogListParamsDirection string

const (
	LogListParamsDirectionAsc  LogListParamsDirection = "asc"
	LogListParamsDirectionDesc LogListParamsDirection = "desc"
)

func (r LogListParamsDirection) IsKnown() bool {
	switch r {
	case LogListParamsDirectionAsc, LogListParamsDirectionDesc:
		return true
	}
	return false
}

type LogListParamsOrderBy string

const (
	LogListParamsOrderByCreatedAt LogListParamsOrderBy = "created_at"
	LogListParamsOrderByProvider  LogListParamsOrderBy = "provider"
)

func (r LogListParamsOrderBy) IsKnown() bool {
	switch r {
	case LogListParamsOrderByCreatedAt, LogListParamsOrderByProvider:
		return true
	}
	return false
}
