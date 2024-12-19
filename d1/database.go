// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1

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
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
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
func (r *DatabaseService) New(ctx context.Context, params DatabaseNewParams, opts ...option.RequestOption) (res *D1, err error) {
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
func (r *DatabaseService) Delete(ctx context.Context, databaseID string, body DatabaseDeleteParams, opts ...option.RequestOption) (res *DatabaseDeleteResponse, err error) {
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

// Returns a URL where the SQL contents of your D1 can be downloaded. Note: this
// process may take some time for larger DBs, during which your D1 will be
// unavailable to serve queries. To avoid blocking your DB unnecessarily, an
// in-progress export must be continually polled or will automatically cancel.
func (r *DatabaseService) Export(ctx context.Context, databaseID string, params DatabaseExportParams, opts ...option.RequestOption) (res *DatabaseExportResponse, err error) {
	var env DatabaseExportResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/export", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
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

// Generates a temporary URL for uploading an SQL file to, then instructing the D1
// to import it and polling it for status updates. Imports block the D1 for their
// duration.
func (r *DatabaseService) Import(ctx context.Context, databaseID string, params DatabaseImportParams, opts ...option.RequestOption) (res *DatabaseImportResponse, err error) {
	var env DatabaseImportResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/import", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
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

type DatabaseListResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt time.Time                `json:"created_at" format:"date-time"`
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

type DatabaseDeleteResponse = interface{}

type DatabaseExportResponse struct {
	// The current time-travel bookmark for your D1, used to poll for updates. Will not
	// change for the duration of the export task.
	AtBookmark string `json:"at_bookmark"`
	// Only present when status = 'error'. Contains the error message.
	Error string `json:"error"`
	// Logs since the last time you polled
	Messages []string `json:"messages"`
	// Only present when status = 'complete'
	Result  DatabaseExportResponseResult `json:"result"`
	Status  DatabaseExportResponseStatus `json:"status"`
	Success bool                         `json:"success"`
	Type    DatabaseExportResponseType   `json:"type"`
	JSON    databaseExportResponseJSON   `json:"-"`
}

// databaseExportResponseJSON contains the JSON metadata for the struct
// [DatabaseExportResponse]
type databaseExportResponseJSON struct {
	AtBookmark  apijson.Field
	Error       apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseExportResponseJSON) RawJSON() string {
	return r.raw
}

// Only present when status = 'complete'
type DatabaseExportResponseResult struct {
	// The generated SQL filename.
	Filename string `json:"filename"`
	// The URL to download the exported SQL. Available for one hour.
	SignedURL string                           `json:"signed_url"`
	JSON      databaseExportResponseResultJSON `json:"-"`
}

// databaseExportResponseResultJSON contains the JSON metadata for the struct
// [DatabaseExportResponseResult]
type databaseExportResponseResultJSON struct {
	Filename    apijson.Field
	SignedURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseExportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseExportResponseResultJSON) RawJSON() string {
	return r.raw
}

type DatabaseExportResponseStatus string

const (
	DatabaseExportResponseStatusComplete DatabaseExportResponseStatus = "complete"
	DatabaseExportResponseStatusError    DatabaseExportResponseStatus = "error"
)

func (r DatabaseExportResponseStatus) IsKnown() bool {
	switch r {
	case DatabaseExportResponseStatusComplete, DatabaseExportResponseStatusError:
		return true
	}
	return false
}

type DatabaseExportResponseType string

const (
	DatabaseExportResponseTypeExport DatabaseExportResponseType = "export"
)

func (r DatabaseExportResponseType) IsKnown() bool {
	switch r {
	case DatabaseExportResponseTypeExport:
		return true
	}
	return false
}

type DatabaseImportResponse struct {
	// The current time-travel bookmark for your D1, used to poll for updates. Will not
	// change for the duration of the import. Only returned if an import process is
	// currently running or recently finished.
	AtBookmark string `json:"at_bookmark"`
	// Only present when status = 'error'. Contains the error message that prevented
	// the import from succeeding.
	Error string `json:"error"`
	// Derived from the database ID and etag, to use in avoiding repeated uploads. Only
	// returned when for the 'init' action.
	Filename string `json:"filename"`
	// Logs since the last time you polled
	Messages []string `json:"messages"`
	// Only present when status = 'complete'
	Result  DatabaseImportResponseResult `json:"result"`
	Status  DatabaseImportResponseStatus `json:"status"`
	Success bool                         `json:"success"`
	Type    DatabaseImportResponseType   `json:"type"`
	// The R2 presigned URL to use for uploading. Only returned when for the 'init'
	// action.
	UploadURL string                     `json:"upload_url"`
	JSON      databaseImportResponseJSON `json:"-"`
}

// databaseImportResponseJSON contains the JSON metadata for the struct
// [DatabaseImportResponse]
type databaseImportResponseJSON struct {
	AtBookmark  apijson.Field
	Error       apijson.Field
	Filename    apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseJSON) RawJSON() string {
	return r.raw
}

// Only present when status = 'complete'
type DatabaseImportResponseResult struct {
	// The time-travel bookmark if you need restore your D1 to directly after the
	// import succeeded.
	FinalBookmark string                           `json:"final_bookmark"`
	Meta          DatabaseImportResponseResultMeta `json:"meta"`
	// The total number of queries that were executed during the import.
	NumQueries float64                          `json:"num_queries"`
	JSON       databaseImportResponseResultJSON `json:"-"`
}

// databaseImportResponseResultJSON contains the JSON metadata for the struct
// [DatabaseImportResponseResult]
type databaseImportResponseResultJSON struct {
	FinalBookmark apijson.Field
	Meta          apijson.Field
	NumQueries    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DatabaseImportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseResultJSON) RawJSON() string {
	return r.raw
}

type DatabaseImportResponseResultMeta struct {
	ChangedDB   bool                                 `json:"changed_db"`
	Changes     float64                              `json:"changes"`
	Duration    float64                              `json:"duration"`
	LastRowID   float64                              `json:"last_row_id"`
	RowsRead    float64                              `json:"rows_read"`
	RowsWritten float64                              `json:"rows_written"`
	SizeAfter   float64                              `json:"size_after"`
	JSON        databaseImportResponseResultMetaJSON `json:"-"`
}

// databaseImportResponseResultMetaJSON contains the JSON metadata for the struct
// [DatabaseImportResponseResultMeta]
type databaseImportResponseResultMetaJSON struct {
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

func (r *DatabaseImportResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type DatabaseImportResponseStatus string

const (
	DatabaseImportResponseStatusComplete DatabaseImportResponseStatus = "complete"
	DatabaseImportResponseStatusError    DatabaseImportResponseStatus = "error"
)

func (r DatabaseImportResponseStatus) IsKnown() bool {
	switch r {
	case DatabaseImportResponseStatusComplete, DatabaseImportResponseStatusError:
		return true
	}
	return false
}

type DatabaseImportResponseType string

const (
	DatabaseImportResponseTypeImport DatabaseImportResponseType = "import"
)

func (r DatabaseImportResponseType) IsKnown() bool {
	switch r {
	case DatabaseImportResponseTypeImport:
		return true
	}
	return false
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
	Columns []string                       `json:"columns"`
	Rows    [][]interface{}                `json:"rows"`
	JSON    databaseRawResponseResultsJSON `json:"-"`
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

type DatabaseNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name,required"`
	// Specify the region to create the D1 primary, if available. If this option is
	// omitted, the D1 will be created as close as possible to the current user.
	PrimaryLocationHint param.Field[DatabaseNewParamsPrimaryLocationHint] `json:"primary_location_hint"`
}

func (r DatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the region to create the D1 primary, if available. If this option is
// omitted, the D1 will be created as close as possible to the current user.
type DatabaseNewParamsPrimaryLocationHint string

const (
	DatabaseNewParamsPrimaryLocationHintWnam DatabaseNewParamsPrimaryLocationHint = "wnam"
	DatabaseNewParamsPrimaryLocationHintEnam DatabaseNewParamsPrimaryLocationHint = "enam"
	DatabaseNewParamsPrimaryLocationHintWeur DatabaseNewParamsPrimaryLocationHint = "weur"
	DatabaseNewParamsPrimaryLocationHintEeur DatabaseNewParamsPrimaryLocationHint = "eeur"
	DatabaseNewParamsPrimaryLocationHintApac DatabaseNewParamsPrimaryLocationHint = "apac"
	DatabaseNewParamsPrimaryLocationHintOc   DatabaseNewParamsPrimaryLocationHint = "oc"
)

func (r DatabaseNewParamsPrimaryLocationHint) IsKnown() bool {
	switch r {
	case DatabaseNewParamsPrimaryLocationHintWnam, DatabaseNewParamsPrimaryLocationHintEnam, DatabaseNewParamsPrimaryLocationHintWeur, DatabaseNewParamsPrimaryLocationHintEeur, DatabaseNewParamsPrimaryLocationHintApac, DatabaseNewParamsPrimaryLocationHintOc:
		return true
	}
	return false
}

type DatabaseNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   D1                    `json:"result,required"`
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
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   DatabaseDeleteResponse `json:"result,required,nullable"`
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

type DatabaseExportParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies that you will poll this endpoint until the export completes
	OutputFormat param.Field[DatabaseExportParamsOutputFormat] `json:"output_format,required"`
	// To poll an in-progress export, provide the current bookmark (returned by your
	// first polling response)
	CurrentBookmark param.Field[string]                          `json:"current_bookmark"`
	DumpOptions     param.Field[DatabaseExportParamsDumpOptions] `json:"dump_options"`
}

func (r DatabaseExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies that you will poll this endpoint until the export completes
type DatabaseExportParamsOutputFormat string

const (
	DatabaseExportParamsOutputFormatPolling DatabaseExportParamsOutputFormat = "polling"
)

func (r DatabaseExportParamsOutputFormat) IsKnown() bool {
	switch r {
	case DatabaseExportParamsOutputFormatPolling:
		return true
	}
	return false
}

type DatabaseExportParamsDumpOptions struct {
	// Export only the table definitions, not their contents
	NoData param.Field[bool] `json:"no_data"`
	// Export only each table's contents, not its definition
	NoSchema param.Field[bool] `json:"no_schema"`
	// Filter the export to just one or more tables. Passing an empty array is the same
	// as not passing anything and means: export all tables.
	Tables param.Field[[]string] `json:"tables"`
}

func (r DatabaseExportParamsDumpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseExportResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   DatabaseExportResponse `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseExportResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseExportResponseEnvelopeJSON    `json:"-"`
}

// databaseExportResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseExportResponseEnvelope]
type databaseExportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseExportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseExportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseExportResponseEnvelopeSuccess bool

const (
	DatabaseExportResponseEnvelopeSuccessTrue DatabaseExportResponseEnvelopeSuccess = true
)

func (r DatabaseExportResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseExportResponseEnvelopeSuccessTrue:
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

type DatabaseImportParams struct {
	// Account identifier tag.
	AccountID param.Field[string]           `path:"account_id,required"`
	Body      DatabaseImportParamsBodyUnion `json:"body,required"`
}

func (r DatabaseImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DatabaseImportParamsBody struct {
	// Indicates you have a new SQL file to upload.
	Action param.Field[DatabaseImportParamsBodyAction] `json:"action,required"`
	// This identifies the currently-running import, checking its status.
	CurrentBookmark param.Field[string] `json:"current_bookmark"`
	// Required when action is 'init' or 'ingest'. An md5 hash of the file you're
	// uploading. Used to check if it already exists, and validate its contents before
	// ingesting.
	Etag param.Field[string] `json:"etag"`
	// The filename you have successfully uploaded.
	Filename param.Field[string] `json:"filename"`
}

func (r DatabaseImportParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatabaseImportParamsBody) implementsD1DatabaseImportParamsBodyUnion() {}

// Satisfied by [d1.DatabaseImportParamsBodyObject],
// [d1.DatabaseImportParamsBodyObject], [d1.DatabaseImportParamsBodyObject],
// [DatabaseImportParamsBody].
type DatabaseImportParamsBodyUnion interface {
	implementsD1DatabaseImportParamsBodyUnion()
}

type DatabaseImportParamsBodyObject struct {
	// Indicates you have a new SQL file to upload.
	Action param.Field[DatabaseImportParamsBodyObjectAction] `json:"action,required"`
	// Required when action is 'init' or 'ingest'. An md5 hash of the file you're
	// uploading. Used to check if it already exists, and validate its contents before
	// ingesting.
	Etag param.Field[string] `json:"etag,required"`
}

func (r DatabaseImportParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatabaseImportParamsBodyObject) implementsD1DatabaseImportParamsBodyUnion() {}

// Indicates you have a new SQL file to upload.
type DatabaseImportParamsBodyObjectAction string

const (
	DatabaseImportParamsBodyObjectActionInit DatabaseImportParamsBodyObjectAction = "init"
)

func (r DatabaseImportParamsBodyObjectAction) IsKnown() bool {
	switch r {
	case DatabaseImportParamsBodyObjectActionInit:
		return true
	}
	return false
}

// Indicates you have a new SQL file to upload.
type DatabaseImportParamsBodyAction string

const (
	DatabaseImportParamsBodyActionInit   DatabaseImportParamsBodyAction = "init"
	DatabaseImportParamsBodyActionIngest DatabaseImportParamsBodyAction = "ingest"
	DatabaseImportParamsBodyActionPoll   DatabaseImportParamsBodyAction = "poll"
)

func (r DatabaseImportParamsBodyAction) IsKnown() bool {
	switch r {
	case DatabaseImportParamsBodyActionInit, DatabaseImportParamsBodyActionIngest, DatabaseImportParamsBodyActionPoll:
		return true
	}
	return false
}

type DatabaseImportResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   DatabaseImportResponse `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseImportResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseImportResponseEnvelopeJSON    `json:"-"`
}

// databaseImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseImportResponseEnvelope]
type databaseImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseImportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseImportResponseEnvelopeSuccess bool

const (
	DatabaseImportResponseEnvelopeSuccessTrue DatabaseImportResponseEnvelopeSuccess = true
)

func (r DatabaseImportResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseImportResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseQueryParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Your SQL query. Supports multiple statements, joined by semicolons, which will
	// be executed as a batch.
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
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
	AccountID param.Field[string] `path:"account_id,required"`
	// Your SQL query. Supports multiple statements, joined by semicolons, which will
	// be executed as a batch.
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
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
