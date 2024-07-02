// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// DatabaseService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabaseService] method instead.
type DatabaseService struct {
	Options []option.RequestOption
}

// NewDatabaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatabaseService(opts ...option.RequestOption) (r *DatabaseService) {
	r = &DatabaseService{}
	r.Options = opts
	return
}

// Returns the created D1 database.
func (r *DatabaseService) New(ctx context.Context, params DatabaseNewParams, opts ...option.RequestOption) (res *DatabaseNewResponse, err error) {
	var env DatabaseNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of D1 databases.
func (r *DatabaseService) List(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DatabaseListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
func (r *DatabaseService) ListAutoPaging(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DatabaseListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the specified D1 database.
func (r *DatabaseService) Delete(ctx context.Context, databaseID string, body DatabaseDeleteParams, opts ...option.RequestOption) (res *DatabaseDeleteResponseUnion, err error) {
	var env DatabaseDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", body.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified D1 database.
func (r *DatabaseService) Get(ctx context.Context, databaseID string, query DatabaseGetParams, opts ...option.RequestOption) (res *D1, err error) {
	var env DatabaseGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", query.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the query result as an object.
func (r *DatabaseService) Query(ctx context.Context, databaseID string, params DatabaseQueryParams, opts ...option.RequestOption) (res *[]QueryResult, err error) {
	var env DatabaseQueryResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/query", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the query result rows as arrays rather than objects. This is a
// performance-optimized version of the /query endpoint.
func (r *DatabaseService) Raw(ctx context.Context, databaseID string, params DatabaseRawParams, opts ...option.RequestOption) (res *[]DatabaseRawResponse, err error) {
	var env DatabaseRawResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/raw", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type QueryResult struct {
	Meta    QueryResultMeta `json:"meta"`
	Results []interface{}   `json:"results"`
	Success bool            `json:"success"`
	JSON    queryResultJSON `json:"-"`
}

// queryResultJSON contains the JSON metadata for the struct [QueryResult]
type queryResultJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResultJSON) RawJSON() string {
	return r.raw
}

type QueryResultMeta struct {
	ChangedDB   bool                `json:"changed_db"`
	Changes     float64             `json:"changes"`
	Duration    float64             `json:"duration"`
	LastRowID   float64             `json:"last_row_id"`
	RowsRead    float64             `json:"rows_read"`
	RowsWritten float64             `json:"rows_written"`
	SizeAfter   float64             `json:"size_after"`
	JSON        queryResultMetaJSON `json:"-"`
}

// queryResultMetaJSON contains the JSON metadata for the struct [QueryResultMeta]
type queryResultMetaJSON struct {
	ChangedDB   apijson.Field
	Changes     apijson.Field
	Duration    apijson.Field
	LastRowID   apijson.Field
	RowsRead    apijson.Field
	RowsWritten apijson.Field
	SizeAfter   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResultMetaJSON) RawJSON() string {
	return r.raw
}

type DatabaseNewResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt string                  `json:"created_at"`
	Name      string                  `json:"name"`
	UUID      string                  `json:"uuid"`
	Version   string                  `json:"version"`
	JSON      databaseNewResponseJSON `json:"-"`
}

// databaseNewResponseJSON contains the JSON metadata for the struct
// [DatabaseNewResponse]
type databaseNewResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseNewResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseListResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt string                   `json:"created_at"`
	Name      string                   `json:"name"`
	UUID      string                   `json:"uuid"`
	Version   string                   `json:"version"`
	JSON      databaseListResponseJSON `json:"-"`
}

// databaseListResponseJSON contains the JSON metadata for the struct
// [DatabaseListResponse]
type databaseListResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseListResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [d1.DatabaseDeleteResponseUnknown] or [shared.UnionString].
type DatabaseDeleteResponseUnion interface {
	ImplementsD1DatabaseDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatabaseDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DatabaseRawResponse struct {
	Meta    DatabaseRawResponseMeta    `json:"meta"`
	Results DatabaseRawResponseResults `json:"results"`
	Success bool                       `json:"success"`
	JSON    databaseRawResponseJSON    `json:"-"`
}

// databaseRawResponseJSON contains the JSON metadata for the struct
// [DatabaseRawResponse]
type databaseRawResponseJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseRawResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseRawResponseMeta struct {
	ChangedDB   bool                        `json:"changed_db"`
	Changes     float64                     `json:"changes"`
	Duration    float64                     `json:"duration"`
	LastRowID   float64                     `json:"last_row_id"`
	RowsRead    float64                     `json:"rows_read"`
	RowsWritten float64                     `json:"rows_written"`
	SizeAfter   float64                     `json:"size_after"`
	JSON        databaseRawResponseMetaJSON `json:"-"`
}

// databaseRawResponseMetaJSON contains the JSON metadata for the struct
// [DatabaseRawResponseMeta]
type databaseRawResponseMetaJSON struct {
	ChangedDB   apijson.Field
	Changes     apijson.Field
	Duration    apijson.Field
	LastRowID   apijson.Field
	RowsRead    apijson.Field
	RowsWritten apijson.Field
	SizeAfter   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseRawResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DatabaseRawResponseResults struct {
	Columns []string                                `json:"columns"`
	Rows    [][]DatabaseRawResponseResultsRowsUnion `json:"rows"`
	JSON    databaseRawResponseResultsJSON          `json:"-"`
}

// databaseRawResponseResultsJSON contains the JSON metadata for the struct
// [DatabaseRawResponseResults]
type databaseRawResponseResultsJSON struct {
	Columns     apijson.Field
	Rows        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseRawResponseResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseResultsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [d1.DatabaseRawResponseResultsRowsUnknown].
type DatabaseRawResponseResultsRowsUnion interface {
	ImplementsD1DatabaseRawResponseResultsRowsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatabaseRawResponseResultsRowsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DatabaseNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name,required"`
}

func (r DatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DatabaseNewResponse   `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseNewResponseEnvelopeJSON    `json:"-"`
}

// databaseNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseNewResponseEnvelope]
type databaseNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseNewResponseEnvelopeSuccess bool

const (
	DatabaseNewResponseEnvelopeSuccessTrue DatabaseNewResponseEnvelopeSuccess = true
)

func (r DatabaseNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// a database name to search for.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [DatabaseListParams]'s query parameters as `url.Values`.
func (r DatabaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DatabaseDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DatabaseDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   DatabaseDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseDeleteResponseEnvelopeJSON    `json:"-"`
}

// databaseDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseDeleteResponseEnvelope]
type databaseDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseDeleteResponseEnvelopeSuccess bool

const (
	DatabaseDeleteResponseEnvelopeSuccessTrue DatabaseDeleteResponseEnvelopeSuccess = true
)

func (r DatabaseDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DatabaseGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   D1                    `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseGetResponseEnvelopeJSON    `json:"-"`
}

// databaseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseGetResponseEnvelope]
type databaseGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseGetResponseEnvelopeSuccess bool

const (
	DatabaseGetResponseEnvelopeSuccessTrue DatabaseGetResponseEnvelopeSuccess = true
)

func (r DatabaseGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseQueryParams struct {
	// Account identifier tag.
	AccountID param.Field[string]   `path:"account_id,required"`
	Sql       param.Field[string]   `json:"sql,required"`
	Params    param.Field[[]string] `json:"params"`
}

func (r DatabaseQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseQueryResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []QueryResult         `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseQueryResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseQueryResponseEnvelopeJSON    `json:"-"`
}

// databaseQueryResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseQueryResponseEnvelope]
type databaseQueryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseQueryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseQueryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseQueryResponseEnvelopeSuccess bool

const (
	DatabaseQueryResponseEnvelopeSuccessTrue DatabaseQueryResponseEnvelopeSuccess = true
)

func (r DatabaseQueryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseQueryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseRawParams struct {
	// Account identifier tag.
	AccountID param.Field[string]   `path:"account_id,required"`
	Sql       param.Field[string]   `json:"sql,required"`
	Params    param.Field[[]string] `json:"params"`
}

func (r DatabaseRawParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseRawResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []DatabaseRawResponse `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseRawResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseRawResponseEnvelopeJSON    `json:"-"`
}

// databaseRawResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseRawResponseEnvelope]
type databaseRawResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseRawResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseRawResponseEnvelopeSuccess bool

const (
	DatabaseRawResponseEnvelopeSuccessTrue DatabaseRawResponseEnvelopeSuccess = true
)

func (r DatabaseRawResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseRawResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
