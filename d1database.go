// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// D1DatabaseService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewD1DatabaseService] method instead.
type D1DatabaseService struct {
	Options []option.RequestOption
}

// NewD1DatabaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewD1DatabaseService(opts ...option.RequestOption) (r *D1DatabaseService) {
	r = &D1DatabaseService{}
	r.Options = opts
	return
}

// Returns the created D1 database.
func (r *D1DatabaseService) New(ctx context.Context, params D1DatabaseNewParams, opts ...option.RequestOption) (res *D1DatabaseNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env D1DatabaseNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of D1 databases.
func (r *D1DatabaseService) List(ctx context.Context, params D1DatabaseListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[D1DatabaseListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/d1/database", params.AccountID)
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

// Returns a list of D1 databases.
func (r *D1DatabaseService) ListAutoPaging(ctx context.Context, params D1DatabaseListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[D1DatabaseListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type D1DatabaseNewResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{}               `json:"created_at"`
	Name      string                    `json:"name"`
	UUID      string                    `json:"uuid"`
	Version   string                    `json:"version"`
	JSON      d1DatabaseNewResponseJSON `json:"-"`
}

// d1DatabaseNewResponseJSON contains the JSON metadata for the struct
// [D1DatabaseNewResponse]
type d1DatabaseNewResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseListResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{}                `json:"created_at"`
	Name      string                     `json:"name"`
	UUID      string                     `json:"uuid"`
	Version   string                     `json:"version"`
	JSON      d1DatabaseListResponseJSON `json:"-"`
}

// d1DatabaseListResponseJSON contains the JSON metadata for the struct
// [D1DatabaseListResponse]
type d1DatabaseListResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r D1DatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type D1DatabaseNewResponseEnvelope struct {
	Errors   []D1DatabaseNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []D1DatabaseNewResponseEnvelopeMessages `json:"messages,required"`
	Result   D1DatabaseNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success D1DatabaseNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    d1DatabaseNewResponseEnvelopeJSON    `json:"-"`
}

// d1DatabaseNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [D1DatabaseNewResponseEnvelope]
type d1DatabaseNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    d1DatabaseNewResponseEnvelopeErrorsJSON `json:"-"`
}

// d1DatabaseNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [D1DatabaseNewResponseEnvelopeErrors]
type d1DatabaseNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    d1DatabaseNewResponseEnvelopeMessagesJSON `json:"-"`
}

// d1DatabaseNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [D1DatabaseNewResponseEnvelopeMessages]
type d1DatabaseNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type D1DatabaseNewResponseEnvelopeSuccess bool

const (
	D1DatabaseNewResponseEnvelopeSuccessTrue D1DatabaseNewResponseEnvelopeSuccess = true
)

type D1DatabaseListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// a database name to search for.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [D1DatabaseListParams]'s query parameters as `url.Values`.
func (r D1DatabaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
