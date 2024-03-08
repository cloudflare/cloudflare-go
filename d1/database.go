// File generated from our OpenAPI spec by Stainless.

package d1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
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
func (r *DatabaseService) New(ctx context.Context, params DatabaseNewParams, opts ...option.RequestOption) (res *DatabaseNewResponse, err error) {
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
func (r *DatabaseService) List(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[DatabaseListResponse], err error) {
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
func (r *DatabaseService) ListAutoPaging(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[DatabaseListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the specified D1 database.
func (r *DatabaseService) Delete(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *DatabaseDeleteResponse, err error) {
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
func (r *DatabaseService) Get(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *DatabaseGetResponse, err error) {
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
func (r *DatabaseService) Query(ctx context.Context, accountIdentifier string, databaseIdentifier string, body DatabaseQueryParams, opts ...option.RequestOption) (res *[]DatabaseQueryResponse, err error) {
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

type DatabaseNewResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{}             `json:"created_at"`
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
	CreatedAt interface{}              `json:"created_at"`
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
type DatabaseDeleteResponse interface {
	ImplementsD1DatabaseDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatabaseDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DatabaseGetResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{} `json:"created_at"`
	// The D1 database's size, in bytes.
	FileSize  float64                 `json:"file_size"`
	Name      string                  `json:"name"`
	NumTables float64                 `json:"num_tables"`
	UUID      string                  `json:"uuid"`
	Version   string                  `json:"version"`
	JSON      databaseGetResponseJSON `json:"-"`
}

// databaseGetResponseJSON contains the JSON metadata for the struct
// [DatabaseGetResponse]
type databaseGetResponseJSON struct {
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Name        apijson.Field
	NumTables   apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseGetResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseQueryResponse struct {
	Meta    DatabaseQueryResponseMeta `json:"meta"`
	Results []interface{}             `json:"results"`
	Success bool                      `json:"success"`
	JSON    databaseQueryResponseJSON `json:"-"`
}

// databaseQueryResponseJSON contains the JSON metadata for the struct
// [DatabaseQueryResponse]
type databaseQueryResponseJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseQueryResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseQueryResponseMeta struct {
	ChangedDB   bool                          `json:"changed_db"`
	Changes     float64                       `json:"changes"`
	Duration    float64                       `json:"duration"`
	LastRowID   float64                       `json:"last_row_id"`
	RowsRead    float64                       `json:"rows_read"`
	RowsWritten float64                       `json:"rows_written"`
	SizeAfter   float64                       `json:"size_after"`
	JSON        databaseQueryResponseMetaJSON `json:"-"`
}

// databaseQueryResponseMetaJSON contains the JSON metadata for the struct
// [DatabaseQueryResponseMeta]
type databaseQueryResponseMetaJSON struct {
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

func (r *DatabaseQueryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseQueryResponseMetaJSON) RawJSON() string {
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
	Errors   []DatabaseNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DatabaseNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DatabaseNewResponse                   `json:"result,required"`
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

type DatabaseNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    databaseNewResponseEnvelopeErrorsJSON `json:"-"`
}

// databaseNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DatabaseNewResponseEnvelopeErrors]
type databaseNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DatabaseNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    databaseNewResponseEnvelopeMessagesJSON `json:"-"`
}

// databaseNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DatabaseNewResponseEnvelopeMessages]
type databaseNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseNewResponseEnvelopeSuccess bool

const (
	DatabaseNewResponseEnvelopeSuccessTrue DatabaseNewResponseEnvelopeSuccess = true
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DatabaseDeleteResponseEnvelope struct {
	Errors   []DatabaseDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DatabaseDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DatabaseDeleteResponse                   `json:"result,required,nullable"`
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

type DatabaseDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    databaseDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// databaseDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DatabaseDeleteResponseEnvelopeErrors]
type databaseDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DatabaseDeleteResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    databaseDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// databaseDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DatabaseDeleteResponseEnvelopeMessages]
type databaseDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseDeleteResponseEnvelopeSuccess bool

const (
	DatabaseDeleteResponseEnvelopeSuccessTrue DatabaseDeleteResponseEnvelopeSuccess = true
)

type DatabaseGetResponseEnvelope struct {
	Errors   []DatabaseGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DatabaseGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DatabaseGetResponse                   `json:"result,required"`
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

type DatabaseGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    databaseGetResponseEnvelopeErrorsJSON `json:"-"`
}

// databaseGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DatabaseGetResponseEnvelopeErrors]
type databaseGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DatabaseGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    databaseGetResponseEnvelopeMessagesJSON `json:"-"`
}

// databaseGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DatabaseGetResponseEnvelopeMessages]
type databaseGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseGetResponseEnvelopeSuccess bool

const (
	DatabaseGetResponseEnvelopeSuccessTrue DatabaseGetResponseEnvelopeSuccess = true
)

type DatabaseQueryParams struct {
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r DatabaseQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseQueryResponseEnvelope struct {
	Errors   []DatabaseQueryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DatabaseQueryResponseEnvelopeMessages `json:"messages,required"`
	Result   []DatabaseQueryResponse                 `json:"result,required"`
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

type DatabaseQueryResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    databaseQueryResponseEnvelopeErrorsJSON `json:"-"`
}

// databaseQueryResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DatabaseQueryResponseEnvelopeErrors]
type databaseQueryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseQueryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseQueryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DatabaseQueryResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    databaseQueryResponseEnvelopeMessagesJSON `json:"-"`
}

// databaseQueryResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DatabaseQueryResponseEnvelopeMessages]
type databaseQueryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseQueryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseQueryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseQueryResponseEnvelopeSuccess bool

const (
	DatabaseQueryResponseEnvelopeSuccessTrue DatabaseQueryResponseEnvelopeSuccess = true
)
