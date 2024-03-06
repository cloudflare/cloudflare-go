// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Deletes the specified D1 database.
func (r *D1DatabaseService) Delete(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *D1DatabaseDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env D1DatabaseDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified D1 database.
func (r *D1DatabaseService) Get(ctx context.Context, accountIdentifier string, databaseIdentifier string, opts ...option.RequestOption) (res *D1DatabaseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env D1DatabaseGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database/%s", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the query result.
func (r *D1DatabaseService) Query(ctx context.Context, accountIdentifier string, databaseIdentifier string, body D1DatabaseQueryParams, opts ...option.RequestOption) (res *[]D1DatabaseQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env D1DatabaseQueryResponseEnvelope
	path := fmt.Sprintf("accounts/%s/d1/database/%s/query", accountIdentifier, databaseIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

// Union satisfied by [D1DatabaseDeleteResponseUnknown] or [shared.UnionString].
type D1DatabaseDeleteResponse interface {
	ImplementsD1DatabaseDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*D1DatabaseDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type D1DatabaseGetResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt interface{} `json:"created_at"`
	// The D1 database's size, in bytes.
	FileSize  float64                   `json:"file_size"`
	Name      string                    `json:"name"`
	NumTables float64                   `json:"num_tables"`
	UUID      string                    `json:"uuid"`
	Version   string                    `json:"version"`
	JSON      d1DatabaseGetResponseJSON `json:"-"`
}

// d1DatabaseGetResponseJSON contains the JSON metadata for the struct
// [D1DatabaseGetResponse]
type d1DatabaseGetResponseJSON struct {
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Name        apijson.Field
	NumTables   apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseQueryResponse struct {
	Meta    D1DatabaseQueryResponseMeta `json:"meta"`
	Results []interface{}               `json:"results"`
	Success bool                        `json:"success"`
	JSON    d1DatabaseQueryResponseJSON `json:"-"`
}

// d1DatabaseQueryResponseJSON contains the JSON metadata for the struct
// [D1DatabaseQueryResponse]
type d1DatabaseQueryResponseJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseQueryResponseMeta struct {
	ChangedDB   bool                            `json:"changed_db"`
	Changes     float64                         `json:"changes"`
	Duration    float64                         `json:"duration"`
	LastRowID   float64                         `json:"last_row_id"`
	RowsRead    float64                         `json:"rows_read"`
	RowsWritten float64                         `json:"rows_written"`
	SizeAfter   float64                         `json:"size_after"`
	JSON        d1DatabaseQueryResponseMetaJSON `json:"-"`
}

// d1DatabaseQueryResponseMetaJSON contains the JSON metadata for the struct
// [D1DatabaseQueryResponseMeta]
type d1DatabaseQueryResponseMetaJSON struct {
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

func (r *D1DatabaseQueryResponseMeta) UnmarshalJSON(data []byte) (err error) {
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

type D1DatabaseDeleteResponseEnvelope struct {
	Errors   []D1DatabaseDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []D1DatabaseDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   D1DatabaseDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success D1DatabaseDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    d1DatabaseDeleteResponseEnvelopeJSON    `json:"-"`
}

// d1DatabaseDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [D1DatabaseDeleteResponseEnvelope]
type d1DatabaseDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseDeleteResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    d1DatabaseDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// d1DatabaseDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [D1DatabaseDeleteResponseEnvelopeErrors]
type d1DatabaseDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseDeleteResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    d1DatabaseDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// d1DatabaseDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [D1DatabaseDeleteResponseEnvelopeMessages]
type d1DatabaseDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type D1DatabaseDeleteResponseEnvelopeSuccess bool

const (
	D1DatabaseDeleteResponseEnvelopeSuccessTrue D1DatabaseDeleteResponseEnvelopeSuccess = true
)

type D1DatabaseGetResponseEnvelope struct {
	Errors   []D1DatabaseGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []D1DatabaseGetResponseEnvelopeMessages `json:"messages,required"`
	Result   D1DatabaseGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success D1DatabaseGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    d1DatabaseGetResponseEnvelopeJSON    `json:"-"`
}

// d1DatabaseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [D1DatabaseGetResponseEnvelope]
type d1DatabaseGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    d1DatabaseGetResponseEnvelopeErrorsJSON `json:"-"`
}

// d1DatabaseGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [D1DatabaseGetResponseEnvelopeErrors]
type d1DatabaseGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    d1DatabaseGetResponseEnvelopeMessagesJSON `json:"-"`
}

// d1DatabaseGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [D1DatabaseGetResponseEnvelopeMessages]
type d1DatabaseGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type D1DatabaseGetResponseEnvelopeSuccess bool

const (
	D1DatabaseGetResponseEnvelopeSuccessTrue D1DatabaseGetResponseEnvelopeSuccess = true
)

type D1DatabaseQueryParams struct {
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r D1DatabaseQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type D1DatabaseQueryResponseEnvelope struct {
	Errors   []D1DatabaseQueryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []D1DatabaseQueryResponseEnvelopeMessages `json:"messages,required"`
	Result   []D1DatabaseQueryResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success D1DatabaseQueryResponseEnvelopeSuccess `json:"success,required"`
	JSON    d1DatabaseQueryResponseEnvelopeJSON    `json:"-"`
}

// d1DatabaseQueryResponseEnvelopeJSON contains the JSON metadata for the struct
// [D1DatabaseQueryResponseEnvelope]
type d1DatabaseQueryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseQueryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseQueryResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    d1DatabaseQueryResponseEnvelopeErrorsJSON `json:"-"`
}

// d1DatabaseQueryResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [D1DatabaseQueryResponseEnvelopeErrors]
type d1DatabaseQueryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseQueryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type D1DatabaseQueryResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    d1DatabaseQueryResponseEnvelopeMessagesJSON `json:"-"`
}

// d1DatabaseQueryResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [D1DatabaseQueryResponseEnvelopeMessages]
type d1DatabaseQueryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1DatabaseQueryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type D1DatabaseQueryResponseEnvelopeSuccess bool

const (
	D1DatabaseQueryResponseEnvelopeSuccessTrue D1DatabaseQueryResponseEnvelopeSuccess = true
)
