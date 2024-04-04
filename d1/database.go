// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DatabaseService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDatabaseService] method instead.
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
func (r *DatabaseService) New(ctx context.Context, params DatabaseNewParams, opts ...option.RequestOption) (res *D1CreateDatabase, err error) {
	opts = append(r.Options[:], opts...)
	var env DatabaseNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of D1 databases.
func (r *DatabaseService) List(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[D1CreateDatabase], err error) {
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
func (r *DatabaseService) ListAutoPaging(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[D1CreateDatabase] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the specified D1 database.
func (r *DatabaseService) Delete(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef173, err error) {
	opts = append(r.Options[:], opts...)
	var env DatabaseDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified D1 database.
func (r *DatabaseService) Get(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *D1DatabaseDetails, err error) {
	opts = append(r.Options[:], opts...)
	var env DatabaseGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the query result.
func (r *DatabaseService) Query(ctx context.Context, accountIdentifier string, databaseIdentifier string, body DatabaseQueryParams, opts ...option.RequestOption) (res *[]D1QueryResult, err error) {
	opts = append(r.Options[:], opts...)
	var env DatabaseQueryResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database/%s/query", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type D1CreateDatabase struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt string               `json:"created_at"`
	Name      string               `json:"name"`
	UUID      string               `json:"uuid"`
	Version   string               `json:"version"`
	JSON      d1CreateDatabaseJSON `json:"-"`
}

// d1CreateDatabaseJSON contains the JSON metadata for the struct
// [D1CreateDatabase]
type d1CreateDatabaseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1CreateDatabase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1CreateDatabaseJSON) RawJSON() string {
	return r.raw
}

type D1DatabaseDetails struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt string `json:"created_at"`
	// The D1 database's size, in bytes.
	FileSize  float64               `json:"file_size"`
	Name      string                `json:"name"`
	NumTables float64               `json:"num_tables"`
	UUID      string                `json:"uuid"`
	Version   string                `json:"version"`
	JSON      d1DatabaseDetailsJSON `json:"-"`
}

// d1DatabaseDetailsJSON contains the JSON metadata for the struct
// [D1DatabaseDetails]
type d1DatabaseDetailsJSON struct {
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Name        apijson.Field
	NumTables   apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1DatabaseDetailsJSON) RawJSON() string {
	return r.raw
}

type D1QueryResult struct {
	Meta    D1QueryResultMeta `json:"meta"`
	Results []interface{}     `json:"results"`
	Success bool              `json:"success"`
	JSON    d1QueryResultJSON `json:"-"`
}

// d1QueryResultJSON contains the JSON metadata for the struct [D1QueryResult]
type d1QueryResultJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1QueryResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1QueryResultJSON) RawJSON() string {
	return r.raw
}

type D1QueryResultMeta struct {
	ChangedDB   bool                  `json:"changed_db"`
	Changes     float64               `json:"changes"`
	Duration    float64               `json:"duration"`
	LastRowID   float64               `json:"last_row_id"`
	RowsRead    float64               `json:"rows_read"`
	RowsWritten float64               `json:"rows_written"`
	SizeAfter   float64               `json:"size_after"`
	JSON        d1QueryResultMetaJSON `json:"-"`
}

// d1QueryResultMetaJSON contains the JSON metadata for the struct
// [D1QueryResultMeta]
type d1QueryResultMetaJSON struct {
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

func (r *D1QueryResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1QueryResultMetaJSON) RawJSON() string {
	return r.raw
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   D1CreateDatabase             `json:"result,required"`
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
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DatabaseDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   shared.UnnamedSchemaRef173   `json:"result,required,nullable"`
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

type DatabaseGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   D1DatabaseDetails            `json:"result,required"`
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
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r DatabaseQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseQueryResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   []D1QueryResult              `json:"result,required"`
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
