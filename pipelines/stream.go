// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipelines

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// StreamService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStreamService] method instead.
type StreamService struct {
	Options []option.RequestOption
}

// NewStreamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStreamService(opts ...option.RequestOption) (r *StreamService) {
	r = &StreamService{}
	r.Options = opts
	return
}

// Create a new Stream.
func (r *StreamService) New(ctx context.Context, params StreamNewParams, opts ...option.RequestOption) (res *StreamNewResponse, err error) {
	var env StreamNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Stream.
func (r *StreamService) Update(ctx context.Context, streamID string, params StreamUpdateParams, opts ...option.RequestOption) (res *StreamUpdateResponse, err error) {
	var env StreamUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams/%s", params.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List/Filter Streams in Account.
func (r *StreamService) List(ctx context.Context, params StreamListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[StreamListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams", params.AccountID)
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

// List/Filter Streams in Account.
func (r *StreamService) ListAutoPaging(ctx context.Context, params StreamListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[StreamListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete Stream in Account.
func (r *StreamService) Delete(ctx context.Context, streamID string, params StreamDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams/%s", params.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Get Stream Details.
func (r *StreamService) Get(ctx context.Context, streamID string, query StreamGetParams, opts ...option.RequestOption) (res *StreamGetResponse, err error) {
	var env StreamGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams/%s", query.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamNewResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                `json:"id,required"`
	CreatedAt  time.Time             `json:"created_at,required" format:"date-time"`
	HTTP       StreamNewResponseHTTP `json:"http,required"`
	ModifiedAt time.Time             `json:"modified_at,required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name,required"`
	// Indicates the current version of this stream.
	Version       int64                          `json:"version,required"`
	WorkerBinding StreamNewResponseWorkerBinding `json:"worker_binding,required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string                  `json:"endpoint" format:"uri"`
	Format   StreamNewResponseFormat `json:"format"`
	Schema   StreamNewResponseSchema `json:"schema"`
	JSON     streamNewResponseJSON   `json:"-"`
}

// streamNewResponseJSON contains the JSON metadata for the struct
// [StreamNewResponse]
type streamNewResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	HTTP          apijson.Field
	ModifiedAt    apijson.Field
	Name          apijson.Field
	Version       apijson.Field
	WorkerBinding apijson.Field
	Endpoint      apijson.Field
	Format        apijson.Field
	Schema        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseHTTP struct {
	// Indicates that authentication is required for the HTTP endpoint.
	Authentication bool `json:"authentication,required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled,required"`
	// Specifies the CORS options for the HTTP endpoint.
	CORS StreamNewResponseHTTPCORS `json:"cors"`
	JSON streamNewResponseHTTPJSON `json:"-"`
}

// streamNewResponseHTTPJSON contains the JSON metadata for the struct
// [StreamNewResponseHTTP]
type streamNewResponseHTTPJSON struct {
	Authentication apijson.Field
	Enabled        apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamNewResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseHTTPJSON) RawJSON() string {
	return r.raw
}

// Specifies the CORS options for the HTTP endpoint.
type StreamNewResponseHTTPCORS struct {
	Origins []string                      `json:"origins"`
	JSON    streamNewResponseHttpcorsJSON `json:"-"`
}

// streamNewResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamNewResponseHTTPCORS]
type streamNewResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	Enabled bool                               `json:"enabled,required"`
	JSON    streamNewResponseWorkerBindingJSON `json:"-"`
}

// streamNewResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamNewResponseWorkerBinding]
type streamNewResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseFormat struct {
	Type            StreamNewResponseFormatType            `json:"type,required"`
	Compression     StreamNewResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamNewResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                  `json:"row_group_bytes,nullable"`
	TimestampFormat StreamNewResponseFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                   `json:"unstructured"`
	JSON            streamNewResponseFormatJSON            `json:"-"`
	union           StreamNewResponseFormatUnion
}

// streamNewResponseFormatJSON contains the JSON metadata for the struct
// [StreamNewResponseFormat]
type streamNewResponseFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r streamNewResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r *StreamNewResponseFormat) UnmarshalJSON(data []byte) (err error) {
	*r = StreamNewResponseFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamNewResponseFormatUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [StreamNewResponseFormatJson],
// [StreamNewResponseFormatParquet].
func (r StreamNewResponseFormat) AsUnion() StreamNewResponseFormatUnion {
	return r.union
}

// Union satisfied by [StreamNewResponseFormatJson] or
// [StreamNewResponseFormatParquet].
type StreamNewResponseFormatUnion interface {
	implementsStreamNewResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamNewResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type StreamNewResponseFormatJson struct {
	Type            StreamNewResponseFormatJsonType            `json:"type,required"`
	DecimalEncoding StreamNewResponseFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat StreamNewResponseFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                       `json:"unstructured"`
	JSON            streamNewResponseFormatJsonJSON            `json:"-"`
}

// streamNewResponseFormatJsonJSON contains the JSON metadata for the struct
// [StreamNewResponseFormatJson]
type streamNewResponseFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *StreamNewResponseFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseFormatJson) implementsStreamNewResponseFormat() {}

type StreamNewResponseFormatJsonType string

const (
	StreamNewResponseFormatJsonTypeJson StreamNewResponseFormatJsonType = "json"
)

func (r StreamNewResponseFormatJsonType) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamNewResponseFormatJsonDecimalEncoding string

const (
	StreamNewResponseFormatJsonDecimalEncodingNumber StreamNewResponseFormatJsonDecimalEncoding = "number"
	StreamNewResponseFormatJsonDecimalEncodingString StreamNewResponseFormatJsonDecimalEncoding = "string"
	StreamNewResponseFormatJsonDecimalEncodingBytes  StreamNewResponseFormatJsonDecimalEncoding = "bytes"
)

func (r StreamNewResponseFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatJsonDecimalEncodingNumber, StreamNewResponseFormatJsonDecimalEncodingString, StreamNewResponseFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewResponseFormatJsonTimestampFormat string

const (
	StreamNewResponseFormatJsonTimestampFormatRfc3339    StreamNewResponseFormatJsonTimestampFormat = "rfc3339"
	StreamNewResponseFormatJsonTimestampFormatUnixMillis StreamNewResponseFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamNewResponseFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatJsonTimestampFormatRfc3339, StreamNewResponseFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewResponseFormatParquet struct {
	Type          StreamNewResponseFormatParquetType        `json:"type,required"`
	Compression   StreamNewResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                     `json:"row_group_bytes,nullable"`
	JSON          streamNewResponseFormatParquetJSON        `json:"-"`
}

// streamNewResponseFormatParquetJSON contains the JSON metadata for the struct
// [StreamNewResponseFormatParquet]
type streamNewResponseFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamNewResponseFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseFormatParquet) implementsStreamNewResponseFormat() {}

type StreamNewResponseFormatParquetType string

const (
	StreamNewResponseFormatParquetTypeParquet StreamNewResponseFormatParquetType = "parquet"
)

func (r StreamNewResponseFormatParquetType) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamNewResponseFormatParquetCompression string

const (
	StreamNewResponseFormatParquetCompressionUncompressed StreamNewResponseFormatParquetCompression = "uncompressed"
	StreamNewResponseFormatParquetCompressionSnappy       StreamNewResponseFormatParquetCompression = "snappy"
	StreamNewResponseFormatParquetCompressionGzip         StreamNewResponseFormatParquetCompression = "gzip"
	StreamNewResponseFormatParquetCompressionZstd         StreamNewResponseFormatParquetCompression = "zstd"
	StreamNewResponseFormatParquetCompressionLz4          StreamNewResponseFormatParquetCompression = "lz4"
)

func (r StreamNewResponseFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatParquetCompressionUncompressed, StreamNewResponseFormatParquetCompressionSnappy, StreamNewResponseFormatParquetCompressionGzip, StreamNewResponseFormatParquetCompressionZstd, StreamNewResponseFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamNewResponseFormatType string

const (
	StreamNewResponseFormatTypeJson    StreamNewResponseFormatType = "json"
	StreamNewResponseFormatTypeParquet StreamNewResponseFormatType = "parquet"
)

func (r StreamNewResponseFormatType) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatTypeJson, StreamNewResponseFormatTypeParquet:
		return true
	}
	return false
}

type StreamNewResponseFormatCompression string

const (
	StreamNewResponseFormatCompressionUncompressed StreamNewResponseFormatCompression = "uncompressed"
	StreamNewResponseFormatCompressionSnappy       StreamNewResponseFormatCompression = "snappy"
	StreamNewResponseFormatCompressionGzip         StreamNewResponseFormatCompression = "gzip"
	StreamNewResponseFormatCompressionZstd         StreamNewResponseFormatCompression = "zstd"
	StreamNewResponseFormatCompressionLz4          StreamNewResponseFormatCompression = "lz4"
)

func (r StreamNewResponseFormatCompression) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatCompressionUncompressed, StreamNewResponseFormatCompressionSnappy, StreamNewResponseFormatCompressionGzip, StreamNewResponseFormatCompressionZstd, StreamNewResponseFormatCompressionLz4:
		return true
	}
	return false
}

type StreamNewResponseFormatDecimalEncoding string

const (
	StreamNewResponseFormatDecimalEncodingNumber StreamNewResponseFormatDecimalEncoding = "number"
	StreamNewResponseFormatDecimalEncodingString StreamNewResponseFormatDecimalEncoding = "string"
	StreamNewResponseFormatDecimalEncodingBytes  StreamNewResponseFormatDecimalEncoding = "bytes"
)

func (r StreamNewResponseFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatDecimalEncodingNumber, StreamNewResponseFormatDecimalEncodingString, StreamNewResponseFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewResponseFormatTimestampFormat string

const (
	StreamNewResponseFormatTimestampFormatRfc3339    StreamNewResponseFormatTimestampFormat = "rfc3339"
	StreamNewResponseFormatTimestampFormatUnixMillis StreamNewResponseFormatTimestampFormat = "unix_millis"
)

func (r StreamNewResponseFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewResponseFormatTimestampFormatRfc3339, StreamNewResponseFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewResponseSchema struct {
	Fields   []StreamNewResponseSchemaField `json:"fields"`
	Format   StreamNewResponseSchemaFormat  `json:"format"`
	Inferred bool                           `json:"inferred,nullable"`
	JSON     streamNewResponseSchemaJSON    `json:"-"`
}

// streamNewResponseSchemaJSON contains the JSON metadata for the struct
// [StreamNewResponseSchema]
type streamNewResponseSchemaJSON struct {
	Fields      apijson.Field
	Format      apijson.Field
	Inferred    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type StreamNewResponseSchemaField struct {
	Type        StreamNewResponseSchemaFieldsType `json:"type,required"`
	MetadataKey string                            `json:"metadata_key,nullable"`
	Name        string                            `json:"name"`
	Required    bool                              `json:"required"`
	SqlName     string                            `json:"sql_name"`
	Unit        StreamNewResponseSchemaFieldsUnit `json:"unit"`
	JSON        streamNewResponseSchemaFieldJSON  `json:"-"`
	union       StreamNewResponseSchemaFieldsUnion
}

// streamNewResponseSchemaFieldJSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaField]
type streamNewResponseSchemaFieldJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r streamNewResponseSchemaFieldJSON) RawJSON() string {
	return r.raw
}

func (r *StreamNewResponseSchemaField) UnmarshalJSON(data []byte) (err error) {
	*r = StreamNewResponseSchemaField{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamNewResponseSchemaFieldsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamNewResponseSchemaFieldsInt32],
// [StreamNewResponseSchemaFieldsInt64], [StreamNewResponseSchemaFieldsFloat32],
// [StreamNewResponseSchemaFieldsFloat64], [StreamNewResponseSchemaFieldsBool],
// [StreamNewResponseSchemaFieldsString], [StreamNewResponseSchemaFieldsBinary],
// [StreamNewResponseSchemaFieldsTimestamp], [StreamNewResponseSchemaFieldsJson],
// [StreamNewResponseSchemaFieldsStruct], [StreamNewResponseSchemaFieldsList].
func (r StreamNewResponseSchemaField) AsUnion() StreamNewResponseSchemaFieldsUnion {
	return r.union
}

// Union satisfied by [StreamNewResponseSchemaFieldsInt32],
// [StreamNewResponseSchemaFieldsInt64], [StreamNewResponseSchemaFieldsFloat32],
// [StreamNewResponseSchemaFieldsFloat64], [StreamNewResponseSchemaFieldsBool],
// [StreamNewResponseSchemaFieldsString], [StreamNewResponseSchemaFieldsBinary],
// [StreamNewResponseSchemaFieldsTimestamp], [StreamNewResponseSchemaFieldsJson],
// [StreamNewResponseSchemaFieldsStruct] or [StreamNewResponseSchemaFieldsList].
type StreamNewResponseSchemaFieldsUnion interface {
	implementsStreamNewResponseSchemaField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamNewResponseSchemaFieldsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsInt32{}),
			DiscriminatorValue: "int32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsInt64{}),
			DiscriminatorValue: "int64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsFloat32{}),
			DiscriminatorValue: "float32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsFloat64{}),
			DiscriminatorValue: "float64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsBool{}),
			DiscriminatorValue: "bool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsString{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsBinary{}),
			DiscriminatorValue: "binary",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsTimestamp{}),
			DiscriminatorValue: "timestamp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsStruct{}),
			DiscriminatorValue: "struct",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFieldsList{}),
			DiscriminatorValue: "list",
		},
	)
}

type StreamNewResponseSchemaFieldsInt32 struct {
	Type        StreamNewResponseSchemaFieldsInt32Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsInt32JSON `json:"-"`
}

// streamNewResponseSchemaFieldsInt32JSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaFieldsInt32]
type streamNewResponseSchemaFieldsInt32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsInt32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsInt32JSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsInt32) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsInt32Type string

const (
	StreamNewResponseSchemaFieldsInt32TypeInt32 StreamNewResponseSchemaFieldsInt32Type = "int32"
)

func (r StreamNewResponseSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsInt64 struct {
	Type        StreamNewResponseSchemaFieldsInt64Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsInt64JSON `json:"-"`
}

// streamNewResponseSchemaFieldsInt64JSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaFieldsInt64]
type streamNewResponseSchemaFieldsInt64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsInt64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsInt64JSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsInt64) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsInt64Type string

const (
	StreamNewResponseSchemaFieldsInt64TypeInt64 StreamNewResponseSchemaFieldsInt64Type = "int64"
)

func (r StreamNewResponseSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsFloat32 struct {
	Type        StreamNewResponseSchemaFieldsFloat32Type `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsFloat32JSON `json:"-"`
}

// streamNewResponseSchemaFieldsFloat32JSON contains the JSON metadata for the
// struct [StreamNewResponseSchemaFieldsFloat32]
type streamNewResponseSchemaFieldsFloat32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsFloat32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsFloat32JSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsFloat32) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsFloat32Type string

const (
	StreamNewResponseSchemaFieldsFloat32TypeFloat32 StreamNewResponseSchemaFieldsFloat32Type = "float32"
)

func (r StreamNewResponseSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsFloat64 struct {
	Type        StreamNewResponseSchemaFieldsFloat64Type `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsFloat64JSON `json:"-"`
}

// streamNewResponseSchemaFieldsFloat64JSON contains the JSON metadata for the
// struct [StreamNewResponseSchemaFieldsFloat64]
type streamNewResponseSchemaFieldsFloat64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsFloat64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsFloat64JSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsFloat64) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsFloat64Type string

const (
	StreamNewResponseSchemaFieldsFloat64TypeFloat64 StreamNewResponseSchemaFieldsFloat64Type = "float64"
)

func (r StreamNewResponseSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsBool struct {
	Type        StreamNewResponseSchemaFieldsBoolType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsBoolJSON `json:"-"`
}

// streamNewResponseSchemaFieldsBoolJSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaFieldsBool]
type streamNewResponseSchemaFieldsBoolJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsBoolJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsBool) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsBoolType string

const (
	StreamNewResponseSchemaFieldsBoolTypeBool StreamNewResponseSchemaFieldsBoolType = "bool"
)

func (r StreamNewResponseSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsString struct {
	Type        StreamNewResponseSchemaFieldsStringType `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsStringJSON `json:"-"`
}

// streamNewResponseSchemaFieldsStringJSON contains the JSON metadata for the
// struct [StreamNewResponseSchemaFieldsString]
type streamNewResponseSchemaFieldsStringJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsStringJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsString) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsStringType string

const (
	StreamNewResponseSchemaFieldsStringTypeString StreamNewResponseSchemaFieldsStringType = "string"
)

func (r StreamNewResponseSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsBinary struct {
	Type        StreamNewResponseSchemaFieldsBinaryType `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsBinaryJSON `json:"-"`
}

// streamNewResponseSchemaFieldsBinaryJSON contains the JSON metadata for the
// struct [StreamNewResponseSchemaFieldsBinary]
type streamNewResponseSchemaFieldsBinaryJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsBinary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsBinaryJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsBinary) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsBinaryType string

const (
	StreamNewResponseSchemaFieldsBinaryTypeBinary StreamNewResponseSchemaFieldsBinaryType = "binary"
)

func (r StreamNewResponseSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsTimestamp struct {
	Type        StreamNewResponseSchemaFieldsTimestampType `json:"type,required"`
	MetadataKey string                                     `json:"metadata_key,nullable"`
	Name        string                                     `json:"name"`
	Required    bool                                       `json:"required"`
	SqlName     string                                     `json:"sql_name"`
	Unit        StreamNewResponseSchemaFieldsTimestampUnit `json:"unit"`
	JSON        streamNewResponseSchemaFieldsTimestampJSON `json:"-"`
}

// streamNewResponseSchemaFieldsTimestampJSON contains the JSON metadata for the
// struct [StreamNewResponseSchemaFieldsTimestamp]
type streamNewResponseSchemaFieldsTimestampJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsTimestamp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsTimestampJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsTimestamp) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsTimestampType string

const (
	StreamNewResponseSchemaFieldsTimestampTypeTimestamp StreamNewResponseSchemaFieldsTimestampType = "timestamp"
)

func (r StreamNewResponseSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsTimestampUnit string

const (
	StreamNewResponseSchemaFieldsTimestampUnitSecond      StreamNewResponseSchemaFieldsTimestampUnit = "second"
	StreamNewResponseSchemaFieldsTimestampUnitMillisecond StreamNewResponseSchemaFieldsTimestampUnit = "millisecond"
	StreamNewResponseSchemaFieldsTimestampUnitMicrosecond StreamNewResponseSchemaFieldsTimestampUnit = "microsecond"
	StreamNewResponseSchemaFieldsTimestampUnitNanosecond  StreamNewResponseSchemaFieldsTimestampUnit = "nanosecond"
)

func (r StreamNewResponseSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsTimestampUnitSecond, StreamNewResponseSchemaFieldsTimestampUnitMillisecond, StreamNewResponseSchemaFieldsTimestampUnitMicrosecond, StreamNewResponseSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsJson struct {
	Type        StreamNewResponseSchemaFieldsJsonType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        streamNewResponseSchemaFieldsJsonJSON `json:"-"`
}

// streamNewResponseSchemaFieldsJsonJSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaFieldsJson]
type streamNewResponseSchemaFieldsJsonJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsJson) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsJsonType string

const (
	StreamNewResponseSchemaFieldsJsonTypeJson StreamNewResponseSchemaFieldsJsonType = "json"
)

func (r StreamNewResponseSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsStruct struct {
	JSON streamNewResponseSchemaFieldsStructJSON `json:"-"`
}

// streamNewResponseSchemaFieldsStructJSON contains the JSON metadata for the
// struct [StreamNewResponseSchemaFieldsStruct]
type streamNewResponseSchemaFieldsStructJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsStruct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsStructJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsStruct) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsList struct {
	JSON streamNewResponseSchemaFieldsListJSON `json:"-"`
}

// streamNewResponseSchemaFieldsListJSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaFieldsList]
type streamNewResponseSchemaFieldsListJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFieldsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFieldsListJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFieldsList) implementsStreamNewResponseSchemaField() {}

type StreamNewResponseSchemaFieldsType string

const (
	StreamNewResponseSchemaFieldsTypeInt32     StreamNewResponseSchemaFieldsType = "int32"
	StreamNewResponseSchemaFieldsTypeInt64     StreamNewResponseSchemaFieldsType = "int64"
	StreamNewResponseSchemaFieldsTypeFloat32   StreamNewResponseSchemaFieldsType = "float32"
	StreamNewResponseSchemaFieldsTypeFloat64   StreamNewResponseSchemaFieldsType = "float64"
	StreamNewResponseSchemaFieldsTypeBool      StreamNewResponseSchemaFieldsType = "bool"
	StreamNewResponseSchemaFieldsTypeString    StreamNewResponseSchemaFieldsType = "string"
	StreamNewResponseSchemaFieldsTypeBinary    StreamNewResponseSchemaFieldsType = "binary"
	StreamNewResponseSchemaFieldsTypeTimestamp StreamNewResponseSchemaFieldsType = "timestamp"
	StreamNewResponseSchemaFieldsTypeJson      StreamNewResponseSchemaFieldsType = "json"
	StreamNewResponseSchemaFieldsTypeStruct    StreamNewResponseSchemaFieldsType = "struct"
	StreamNewResponseSchemaFieldsTypeList      StreamNewResponseSchemaFieldsType = "list"
)

func (r StreamNewResponseSchemaFieldsType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsTypeInt32, StreamNewResponseSchemaFieldsTypeInt64, StreamNewResponseSchemaFieldsTypeFloat32, StreamNewResponseSchemaFieldsTypeFloat64, StreamNewResponseSchemaFieldsTypeBool, StreamNewResponseSchemaFieldsTypeString, StreamNewResponseSchemaFieldsTypeBinary, StreamNewResponseSchemaFieldsTypeTimestamp, StreamNewResponseSchemaFieldsTypeJson, StreamNewResponseSchemaFieldsTypeStruct, StreamNewResponseSchemaFieldsTypeList:
		return true
	}
	return false
}

type StreamNewResponseSchemaFieldsUnit string

const (
	StreamNewResponseSchemaFieldsUnitSecond      StreamNewResponseSchemaFieldsUnit = "second"
	StreamNewResponseSchemaFieldsUnitMillisecond StreamNewResponseSchemaFieldsUnit = "millisecond"
	StreamNewResponseSchemaFieldsUnitMicrosecond StreamNewResponseSchemaFieldsUnit = "microsecond"
	StreamNewResponseSchemaFieldsUnitNanosecond  StreamNewResponseSchemaFieldsUnit = "nanosecond"
)

func (r StreamNewResponseSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFieldsUnitSecond, StreamNewResponseSchemaFieldsUnitMillisecond, StreamNewResponseSchemaFieldsUnitMicrosecond, StreamNewResponseSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormat struct {
	Type            StreamNewResponseSchemaFormatType            `json:"type,required"`
	Compression     StreamNewResponseSchemaFormatCompression     `json:"compression"`
	DecimalEncoding StreamNewResponseSchemaFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                        `json:"row_group_bytes,nullable"`
	TimestampFormat StreamNewResponseSchemaFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                         `json:"unstructured"`
	JSON            streamNewResponseSchemaFormatJSON            `json:"-"`
	union           StreamNewResponseSchemaFormatUnion
}

// streamNewResponseSchemaFormatJSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaFormat]
type streamNewResponseSchemaFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r streamNewResponseSchemaFormatJSON) RawJSON() string {
	return r.raw
}

func (r *StreamNewResponseSchemaFormat) UnmarshalJSON(data []byte) (err error) {
	*r = StreamNewResponseSchemaFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamNewResponseSchemaFormatUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamNewResponseSchemaFormatJson],
// [StreamNewResponseSchemaFormatParquet].
func (r StreamNewResponseSchemaFormat) AsUnion() StreamNewResponseSchemaFormatUnion {
	return r.union
}

// Union satisfied by [StreamNewResponseSchemaFormatJson] or
// [StreamNewResponseSchemaFormatParquet].
type StreamNewResponseSchemaFormatUnion interface {
	implementsStreamNewResponseSchemaFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamNewResponseSchemaFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamNewResponseSchemaFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type StreamNewResponseSchemaFormatJson struct {
	Type            StreamNewResponseSchemaFormatJsonType            `json:"type,required"`
	DecimalEncoding StreamNewResponseSchemaFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat StreamNewResponseSchemaFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                             `json:"unstructured"`
	JSON            streamNewResponseSchemaFormatJsonJSON            `json:"-"`
}

// streamNewResponseSchemaFormatJsonJSON contains the JSON metadata for the struct
// [StreamNewResponseSchemaFormatJson]
type streamNewResponseSchemaFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFormatJson) implementsStreamNewResponseSchemaFormat() {}

type StreamNewResponseSchemaFormatJsonType string

const (
	StreamNewResponseSchemaFormatJsonTypeJson StreamNewResponseSchemaFormatJsonType = "json"
)

func (r StreamNewResponseSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatJsonDecimalEncoding string

const (
	StreamNewResponseSchemaFormatJsonDecimalEncodingNumber StreamNewResponseSchemaFormatJsonDecimalEncoding = "number"
	StreamNewResponseSchemaFormatJsonDecimalEncodingString StreamNewResponseSchemaFormatJsonDecimalEncoding = "string"
	StreamNewResponseSchemaFormatJsonDecimalEncodingBytes  StreamNewResponseSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r StreamNewResponseSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatJsonDecimalEncodingNumber, StreamNewResponseSchemaFormatJsonDecimalEncodingString, StreamNewResponseSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatJsonTimestampFormat string

const (
	StreamNewResponseSchemaFormatJsonTimestampFormatRfc3339    StreamNewResponseSchemaFormatJsonTimestampFormat = "rfc3339"
	StreamNewResponseSchemaFormatJsonTimestampFormatUnixMillis StreamNewResponseSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamNewResponseSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatJsonTimestampFormatRfc3339, StreamNewResponseSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatParquet struct {
	Type          StreamNewResponseSchemaFormatParquetType        `json:"type,required"`
	Compression   StreamNewResponseSchemaFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                           `json:"row_group_bytes,nullable"`
	JSON          streamNewResponseSchemaFormatParquetJSON        `json:"-"`
}

// streamNewResponseSchemaFormatParquetJSON contains the JSON metadata for the
// struct [StreamNewResponseSchemaFormatParquet]
type streamNewResponseSchemaFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamNewResponseSchemaFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseSchemaFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r StreamNewResponseSchemaFormatParquet) implementsStreamNewResponseSchemaFormat() {}

type StreamNewResponseSchemaFormatParquetType string

const (
	StreamNewResponseSchemaFormatParquetTypeParquet StreamNewResponseSchemaFormatParquetType = "parquet"
)

func (r StreamNewResponseSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatParquetCompression string

const (
	StreamNewResponseSchemaFormatParquetCompressionUncompressed StreamNewResponseSchemaFormatParquetCompression = "uncompressed"
	StreamNewResponseSchemaFormatParquetCompressionSnappy       StreamNewResponseSchemaFormatParquetCompression = "snappy"
	StreamNewResponseSchemaFormatParquetCompressionGzip         StreamNewResponseSchemaFormatParquetCompression = "gzip"
	StreamNewResponseSchemaFormatParquetCompressionZstd         StreamNewResponseSchemaFormatParquetCompression = "zstd"
	StreamNewResponseSchemaFormatParquetCompressionLz4          StreamNewResponseSchemaFormatParquetCompression = "lz4"
)

func (r StreamNewResponseSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatParquetCompressionUncompressed, StreamNewResponseSchemaFormatParquetCompressionSnappy, StreamNewResponseSchemaFormatParquetCompressionGzip, StreamNewResponseSchemaFormatParquetCompressionZstd, StreamNewResponseSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatType string

const (
	StreamNewResponseSchemaFormatTypeJson    StreamNewResponseSchemaFormatType = "json"
	StreamNewResponseSchemaFormatTypeParquet StreamNewResponseSchemaFormatType = "parquet"
)

func (r StreamNewResponseSchemaFormatType) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatTypeJson, StreamNewResponseSchemaFormatTypeParquet:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatCompression string

const (
	StreamNewResponseSchemaFormatCompressionUncompressed StreamNewResponseSchemaFormatCompression = "uncompressed"
	StreamNewResponseSchemaFormatCompressionSnappy       StreamNewResponseSchemaFormatCompression = "snappy"
	StreamNewResponseSchemaFormatCompressionGzip         StreamNewResponseSchemaFormatCompression = "gzip"
	StreamNewResponseSchemaFormatCompressionZstd         StreamNewResponseSchemaFormatCompression = "zstd"
	StreamNewResponseSchemaFormatCompressionLz4          StreamNewResponseSchemaFormatCompression = "lz4"
)

func (r StreamNewResponseSchemaFormatCompression) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatCompressionUncompressed, StreamNewResponseSchemaFormatCompressionSnappy, StreamNewResponseSchemaFormatCompressionGzip, StreamNewResponseSchemaFormatCompressionZstd, StreamNewResponseSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatDecimalEncoding string

const (
	StreamNewResponseSchemaFormatDecimalEncodingNumber StreamNewResponseSchemaFormatDecimalEncoding = "number"
	StreamNewResponseSchemaFormatDecimalEncodingString StreamNewResponseSchemaFormatDecimalEncoding = "string"
	StreamNewResponseSchemaFormatDecimalEncodingBytes  StreamNewResponseSchemaFormatDecimalEncoding = "bytes"
)

func (r StreamNewResponseSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatDecimalEncodingNumber, StreamNewResponseSchemaFormatDecimalEncodingString, StreamNewResponseSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewResponseSchemaFormatTimestampFormat string

const (
	StreamNewResponseSchemaFormatTimestampFormatRfc3339    StreamNewResponseSchemaFormatTimestampFormat = "rfc3339"
	StreamNewResponseSchemaFormatTimestampFormatUnixMillis StreamNewResponseSchemaFormatTimestampFormat = "unix_millis"
)

func (r StreamNewResponseSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewResponseSchemaFormatTimestampFormatRfc3339, StreamNewResponseSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamUpdateResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                   `json:"id,required"`
	CreatedAt  time.Time                `json:"created_at,required" format:"date-time"`
	HTTP       StreamUpdateResponseHTTP `json:"http,required"`
	ModifiedAt time.Time                `json:"modified_at,required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name,required"`
	// Indicates the current version of this stream.
	Version       int64                             `json:"version,required"`
	WorkerBinding StreamUpdateResponseWorkerBinding `json:"worker_binding,required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string                     `json:"endpoint" format:"uri"`
	Format   StreamUpdateResponseFormat `json:"format"`
	JSON     streamUpdateResponseJSON   `json:"-"`
}

// streamUpdateResponseJSON contains the JSON metadata for the struct
// [StreamUpdateResponse]
type streamUpdateResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	HTTP          apijson.Field
	ModifiedAt    apijson.Field
	Name          apijson.Field
	Version       apijson.Field
	WorkerBinding apijson.Field
	Endpoint      apijson.Field
	Format        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateResponseHTTP struct {
	// Indicates that authentication is required for the HTTP endpoint.
	Authentication bool `json:"authentication,required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled,required"`
	// Specifies the CORS options for the HTTP endpoint.
	CORS StreamUpdateResponseHTTPCORS `json:"cors"`
	JSON streamUpdateResponseHTTPJSON `json:"-"`
}

// streamUpdateResponseHTTPJSON contains the JSON metadata for the struct
// [StreamUpdateResponseHTTP]
type streamUpdateResponseHTTPJSON struct {
	Authentication apijson.Field
	Enabled        apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamUpdateResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseHTTPJSON) RawJSON() string {
	return r.raw
}

// Specifies the CORS options for the HTTP endpoint.
type StreamUpdateResponseHTTPCORS struct {
	Origins []string                         `json:"origins"`
	JSON    streamUpdateResponseHttpcorsJSON `json:"-"`
}

// streamUpdateResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamUpdateResponseHTTPCORS]
type streamUpdateResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateResponseWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	Enabled bool                                  `json:"enabled,required"`
	JSON    streamUpdateResponseWorkerBindingJSON `json:"-"`
}

// streamUpdateResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamUpdateResponseWorkerBinding]
type streamUpdateResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateResponseFormat struct {
	Type            StreamUpdateResponseFormatType            `json:"type,required"`
	Compression     StreamUpdateResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamUpdateResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                     `json:"row_group_bytes,nullable"`
	TimestampFormat StreamUpdateResponseFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                      `json:"unstructured"`
	JSON            streamUpdateResponseFormatJSON            `json:"-"`
	union           StreamUpdateResponseFormatUnion
}

// streamUpdateResponseFormatJSON contains the JSON metadata for the struct
// [StreamUpdateResponseFormat]
type streamUpdateResponseFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r streamUpdateResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r *StreamUpdateResponseFormat) UnmarshalJSON(data []byte) (err error) {
	*r = StreamUpdateResponseFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamUpdateResponseFormatUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamUpdateResponseFormatJson],
// [StreamUpdateResponseFormatParquet].
func (r StreamUpdateResponseFormat) AsUnion() StreamUpdateResponseFormatUnion {
	return r.union
}

// Union satisfied by [StreamUpdateResponseFormatJson] or
// [StreamUpdateResponseFormatParquet].
type StreamUpdateResponseFormatUnion interface {
	implementsStreamUpdateResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamUpdateResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamUpdateResponseFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamUpdateResponseFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type StreamUpdateResponseFormatJson struct {
	Type            StreamUpdateResponseFormatJsonType            `json:"type,required"`
	DecimalEncoding StreamUpdateResponseFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat StreamUpdateResponseFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                          `json:"unstructured"`
	JSON            streamUpdateResponseFormatJsonJSON            `json:"-"`
}

// streamUpdateResponseFormatJsonJSON contains the JSON metadata for the struct
// [StreamUpdateResponseFormatJson]
type streamUpdateResponseFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *StreamUpdateResponseFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamUpdateResponseFormatJson) implementsStreamUpdateResponseFormat() {}

type StreamUpdateResponseFormatJsonType string

const (
	StreamUpdateResponseFormatJsonTypeJson StreamUpdateResponseFormatJsonType = "json"
)

func (r StreamUpdateResponseFormatJsonType) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamUpdateResponseFormatJsonDecimalEncoding string

const (
	StreamUpdateResponseFormatJsonDecimalEncodingNumber StreamUpdateResponseFormatJsonDecimalEncoding = "number"
	StreamUpdateResponseFormatJsonDecimalEncodingString StreamUpdateResponseFormatJsonDecimalEncoding = "string"
	StreamUpdateResponseFormatJsonDecimalEncodingBytes  StreamUpdateResponseFormatJsonDecimalEncoding = "bytes"
)

func (r StreamUpdateResponseFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatJsonDecimalEncodingNumber, StreamUpdateResponseFormatJsonDecimalEncodingString, StreamUpdateResponseFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamUpdateResponseFormatJsonTimestampFormat string

const (
	StreamUpdateResponseFormatJsonTimestampFormatRfc3339    StreamUpdateResponseFormatJsonTimestampFormat = "rfc3339"
	StreamUpdateResponseFormatJsonTimestampFormatUnixMillis StreamUpdateResponseFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamUpdateResponseFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatJsonTimestampFormatRfc3339, StreamUpdateResponseFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamUpdateResponseFormatParquet struct {
	Type          StreamUpdateResponseFormatParquetType        `json:"type,required"`
	Compression   StreamUpdateResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                        `json:"row_group_bytes,nullable"`
	JSON          streamUpdateResponseFormatParquetJSON        `json:"-"`
}

// streamUpdateResponseFormatParquetJSON contains the JSON metadata for the struct
// [StreamUpdateResponseFormatParquet]
type streamUpdateResponseFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamUpdateResponseFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r StreamUpdateResponseFormatParquet) implementsStreamUpdateResponseFormat() {}

type StreamUpdateResponseFormatParquetType string

const (
	StreamUpdateResponseFormatParquetTypeParquet StreamUpdateResponseFormatParquetType = "parquet"
)

func (r StreamUpdateResponseFormatParquetType) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamUpdateResponseFormatParquetCompression string

const (
	StreamUpdateResponseFormatParquetCompressionUncompressed StreamUpdateResponseFormatParquetCompression = "uncompressed"
	StreamUpdateResponseFormatParquetCompressionSnappy       StreamUpdateResponseFormatParquetCompression = "snappy"
	StreamUpdateResponseFormatParquetCompressionGzip         StreamUpdateResponseFormatParquetCompression = "gzip"
	StreamUpdateResponseFormatParquetCompressionZstd         StreamUpdateResponseFormatParquetCompression = "zstd"
	StreamUpdateResponseFormatParquetCompressionLz4          StreamUpdateResponseFormatParquetCompression = "lz4"
)

func (r StreamUpdateResponseFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatParquetCompressionUncompressed, StreamUpdateResponseFormatParquetCompressionSnappy, StreamUpdateResponseFormatParquetCompressionGzip, StreamUpdateResponseFormatParquetCompressionZstd, StreamUpdateResponseFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamUpdateResponseFormatType string

const (
	StreamUpdateResponseFormatTypeJson    StreamUpdateResponseFormatType = "json"
	StreamUpdateResponseFormatTypeParquet StreamUpdateResponseFormatType = "parquet"
)

func (r StreamUpdateResponseFormatType) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatTypeJson, StreamUpdateResponseFormatTypeParquet:
		return true
	}
	return false
}

type StreamUpdateResponseFormatCompression string

const (
	StreamUpdateResponseFormatCompressionUncompressed StreamUpdateResponseFormatCompression = "uncompressed"
	StreamUpdateResponseFormatCompressionSnappy       StreamUpdateResponseFormatCompression = "snappy"
	StreamUpdateResponseFormatCompressionGzip         StreamUpdateResponseFormatCompression = "gzip"
	StreamUpdateResponseFormatCompressionZstd         StreamUpdateResponseFormatCompression = "zstd"
	StreamUpdateResponseFormatCompressionLz4          StreamUpdateResponseFormatCompression = "lz4"
)

func (r StreamUpdateResponseFormatCompression) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatCompressionUncompressed, StreamUpdateResponseFormatCompressionSnappy, StreamUpdateResponseFormatCompressionGzip, StreamUpdateResponseFormatCompressionZstd, StreamUpdateResponseFormatCompressionLz4:
		return true
	}
	return false
}

type StreamUpdateResponseFormatDecimalEncoding string

const (
	StreamUpdateResponseFormatDecimalEncodingNumber StreamUpdateResponseFormatDecimalEncoding = "number"
	StreamUpdateResponseFormatDecimalEncodingString StreamUpdateResponseFormatDecimalEncoding = "string"
	StreamUpdateResponseFormatDecimalEncodingBytes  StreamUpdateResponseFormatDecimalEncoding = "bytes"
)

func (r StreamUpdateResponseFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatDecimalEncodingNumber, StreamUpdateResponseFormatDecimalEncodingString, StreamUpdateResponseFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamUpdateResponseFormatTimestampFormat string

const (
	StreamUpdateResponseFormatTimestampFormatRfc3339    StreamUpdateResponseFormatTimestampFormat = "rfc3339"
	StreamUpdateResponseFormatTimestampFormatUnixMillis StreamUpdateResponseFormatTimestampFormat = "unix_millis"
)

func (r StreamUpdateResponseFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamUpdateResponseFormatTimestampFormatRfc3339, StreamUpdateResponseFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamListResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                 `json:"id,required"`
	CreatedAt  time.Time              `json:"created_at,required" format:"date-time"`
	HTTP       StreamListResponseHTTP `json:"http,required"`
	ModifiedAt time.Time              `json:"modified_at,required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name,required"`
	// Indicates the current version of this stream.
	Version       int64                           `json:"version,required"`
	WorkerBinding StreamListResponseWorkerBinding `json:"worker_binding,required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string                   `json:"endpoint" format:"uri"`
	Format   StreamListResponseFormat `json:"format"`
	Schema   StreamListResponseSchema `json:"schema"`
	JSON     streamListResponseJSON   `json:"-"`
}

// streamListResponseJSON contains the JSON metadata for the struct
// [StreamListResponse]
type streamListResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	HTTP          apijson.Field
	ModifiedAt    apijson.Field
	Name          apijson.Field
	Version       apijson.Field
	WorkerBinding apijson.Field
	Endpoint      apijson.Field
	Format        apijson.Field
	Schema        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseJSON) RawJSON() string {
	return r.raw
}

type StreamListResponseHTTP struct {
	// Indicates that authentication is required for the HTTP endpoint.
	Authentication bool `json:"authentication,required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled,required"`
	// Specifies the CORS options for the HTTP endpoint.
	CORS StreamListResponseHTTPCORS `json:"cors"`
	JSON streamListResponseHTTPJSON `json:"-"`
}

// streamListResponseHTTPJSON contains the JSON metadata for the struct
// [StreamListResponseHTTP]
type streamListResponseHTTPJSON struct {
	Authentication apijson.Field
	Enabled        apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamListResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseHTTPJSON) RawJSON() string {
	return r.raw
}

// Specifies the CORS options for the HTTP endpoint.
type StreamListResponseHTTPCORS struct {
	Origins []string                       `json:"origins"`
	JSON    streamListResponseHttpcorsJSON `json:"-"`
}

// streamListResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamListResponseHTTPCORS]
type streamListResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamListResponseWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	Enabled bool                                `json:"enabled,required"`
	JSON    streamListResponseWorkerBindingJSON `json:"-"`
}

// streamListResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamListResponseWorkerBinding]
type streamListResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

type StreamListResponseFormat struct {
	Type            StreamListResponseFormatType            `json:"type,required"`
	Compression     StreamListResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamListResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                   `json:"row_group_bytes,nullable"`
	TimestampFormat StreamListResponseFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                    `json:"unstructured"`
	JSON            streamListResponseFormatJSON            `json:"-"`
	union           StreamListResponseFormatUnion
}

// streamListResponseFormatJSON contains the JSON metadata for the struct
// [StreamListResponseFormat]
type streamListResponseFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r streamListResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r *StreamListResponseFormat) UnmarshalJSON(data []byte) (err error) {
	*r = StreamListResponseFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamListResponseFormatUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamListResponseFormatJson],
// [StreamListResponseFormatParquet].
func (r StreamListResponseFormat) AsUnion() StreamListResponseFormatUnion {
	return r.union
}

// Union satisfied by [StreamListResponseFormatJson] or
// [StreamListResponseFormatParquet].
type StreamListResponseFormatUnion interface {
	implementsStreamListResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamListResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type StreamListResponseFormatJson struct {
	Type            StreamListResponseFormatJsonType            `json:"type,required"`
	DecimalEncoding StreamListResponseFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat StreamListResponseFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                        `json:"unstructured"`
	JSON            streamListResponseFormatJsonJSON            `json:"-"`
}

// streamListResponseFormatJsonJSON contains the JSON metadata for the struct
// [StreamListResponseFormatJson]
type streamListResponseFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *StreamListResponseFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseFormatJson) implementsStreamListResponseFormat() {}

type StreamListResponseFormatJsonType string

const (
	StreamListResponseFormatJsonTypeJson StreamListResponseFormatJsonType = "json"
)

func (r StreamListResponseFormatJsonType) IsKnown() bool {
	switch r {
	case StreamListResponseFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamListResponseFormatJsonDecimalEncoding string

const (
	StreamListResponseFormatJsonDecimalEncodingNumber StreamListResponseFormatJsonDecimalEncoding = "number"
	StreamListResponseFormatJsonDecimalEncodingString StreamListResponseFormatJsonDecimalEncoding = "string"
	StreamListResponseFormatJsonDecimalEncodingBytes  StreamListResponseFormatJsonDecimalEncoding = "bytes"
)

func (r StreamListResponseFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamListResponseFormatJsonDecimalEncodingNumber, StreamListResponseFormatJsonDecimalEncodingString, StreamListResponseFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamListResponseFormatJsonTimestampFormat string

const (
	StreamListResponseFormatJsonTimestampFormatRfc3339    StreamListResponseFormatJsonTimestampFormat = "rfc3339"
	StreamListResponseFormatJsonTimestampFormatUnixMillis StreamListResponseFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamListResponseFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamListResponseFormatJsonTimestampFormatRfc3339, StreamListResponseFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamListResponseFormatParquet struct {
	Type          StreamListResponseFormatParquetType        `json:"type,required"`
	Compression   StreamListResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                      `json:"row_group_bytes,nullable"`
	JSON          streamListResponseFormatParquetJSON        `json:"-"`
}

// streamListResponseFormatParquetJSON contains the JSON metadata for the struct
// [StreamListResponseFormatParquet]
type streamListResponseFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamListResponseFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseFormatParquet) implementsStreamListResponseFormat() {}

type StreamListResponseFormatParquetType string

const (
	StreamListResponseFormatParquetTypeParquet StreamListResponseFormatParquetType = "parquet"
)

func (r StreamListResponseFormatParquetType) IsKnown() bool {
	switch r {
	case StreamListResponseFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamListResponseFormatParquetCompression string

const (
	StreamListResponseFormatParquetCompressionUncompressed StreamListResponseFormatParquetCompression = "uncompressed"
	StreamListResponseFormatParquetCompressionSnappy       StreamListResponseFormatParquetCompression = "snappy"
	StreamListResponseFormatParquetCompressionGzip         StreamListResponseFormatParquetCompression = "gzip"
	StreamListResponseFormatParquetCompressionZstd         StreamListResponseFormatParquetCompression = "zstd"
	StreamListResponseFormatParquetCompressionLz4          StreamListResponseFormatParquetCompression = "lz4"
)

func (r StreamListResponseFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamListResponseFormatParquetCompressionUncompressed, StreamListResponseFormatParquetCompressionSnappy, StreamListResponseFormatParquetCompressionGzip, StreamListResponseFormatParquetCompressionZstd, StreamListResponseFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamListResponseFormatType string

const (
	StreamListResponseFormatTypeJson    StreamListResponseFormatType = "json"
	StreamListResponseFormatTypeParquet StreamListResponseFormatType = "parquet"
)

func (r StreamListResponseFormatType) IsKnown() bool {
	switch r {
	case StreamListResponseFormatTypeJson, StreamListResponseFormatTypeParquet:
		return true
	}
	return false
}

type StreamListResponseFormatCompression string

const (
	StreamListResponseFormatCompressionUncompressed StreamListResponseFormatCompression = "uncompressed"
	StreamListResponseFormatCompressionSnappy       StreamListResponseFormatCompression = "snappy"
	StreamListResponseFormatCompressionGzip         StreamListResponseFormatCompression = "gzip"
	StreamListResponseFormatCompressionZstd         StreamListResponseFormatCompression = "zstd"
	StreamListResponseFormatCompressionLz4          StreamListResponseFormatCompression = "lz4"
)

func (r StreamListResponseFormatCompression) IsKnown() bool {
	switch r {
	case StreamListResponseFormatCompressionUncompressed, StreamListResponseFormatCompressionSnappy, StreamListResponseFormatCompressionGzip, StreamListResponseFormatCompressionZstd, StreamListResponseFormatCompressionLz4:
		return true
	}
	return false
}

type StreamListResponseFormatDecimalEncoding string

const (
	StreamListResponseFormatDecimalEncodingNumber StreamListResponseFormatDecimalEncoding = "number"
	StreamListResponseFormatDecimalEncodingString StreamListResponseFormatDecimalEncoding = "string"
	StreamListResponseFormatDecimalEncodingBytes  StreamListResponseFormatDecimalEncoding = "bytes"
)

func (r StreamListResponseFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamListResponseFormatDecimalEncodingNumber, StreamListResponseFormatDecimalEncodingString, StreamListResponseFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamListResponseFormatTimestampFormat string

const (
	StreamListResponseFormatTimestampFormatRfc3339    StreamListResponseFormatTimestampFormat = "rfc3339"
	StreamListResponseFormatTimestampFormatUnixMillis StreamListResponseFormatTimestampFormat = "unix_millis"
)

func (r StreamListResponseFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamListResponseFormatTimestampFormatRfc3339, StreamListResponseFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamListResponseSchema struct {
	Fields   []StreamListResponseSchemaField `json:"fields"`
	Format   StreamListResponseSchemaFormat  `json:"format"`
	Inferred bool                            `json:"inferred,nullable"`
	JSON     streamListResponseSchemaJSON    `json:"-"`
}

// streamListResponseSchemaJSON contains the JSON metadata for the struct
// [StreamListResponseSchema]
type streamListResponseSchemaJSON struct {
	Fields      apijson.Field
	Format      apijson.Field
	Inferred    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type StreamListResponseSchemaField struct {
	Type        StreamListResponseSchemaFieldsType `json:"type,required"`
	MetadataKey string                             `json:"metadata_key,nullable"`
	Name        string                             `json:"name"`
	Required    bool                               `json:"required"`
	SqlName     string                             `json:"sql_name"`
	Unit        StreamListResponseSchemaFieldsUnit `json:"unit"`
	JSON        streamListResponseSchemaFieldJSON  `json:"-"`
	union       StreamListResponseSchemaFieldsUnion
}

// streamListResponseSchemaFieldJSON contains the JSON metadata for the struct
// [StreamListResponseSchemaField]
type streamListResponseSchemaFieldJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r streamListResponseSchemaFieldJSON) RawJSON() string {
	return r.raw
}

func (r *StreamListResponseSchemaField) UnmarshalJSON(data []byte) (err error) {
	*r = StreamListResponseSchemaField{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamListResponseSchemaFieldsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamListResponseSchemaFieldsInt32],
// [StreamListResponseSchemaFieldsInt64], [StreamListResponseSchemaFieldsFloat32],
// [StreamListResponseSchemaFieldsFloat64], [StreamListResponseSchemaFieldsBool],
// [StreamListResponseSchemaFieldsString], [StreamListResponseSchemaFieldsBinary],
// [StreamListResponseSchemaFieldsTimestamp], [StreamListResponseSchemaFieldsJson],
// [StreamListResponseSchemaFieldsStruct], [StreamListResponseSchemaFieldsList].
func (r StreamListResponseSchemaField) AsUnion() StreamListResponseSchemaFieldsUnion {
	return r.union
}

// Union satisfied by [StreamListResponseSchemaFieldsInt32],
// [StreamListResponseSchemaFieldsInt64], [StreamListResponseSchemaFieldsFloat32],
// [StreamListResponseSchemaFieldsFloat64], [StreamListResponseSchemaFieldsBool],
// [StreamListResponseSchemaFieldsString], [StreamListResponseSchemaFieldsBinary],
// [StreamListResponseSchemaFieldsTimestamp], [StreamListResponseSchemaFieldsJson],
// [StreamListResponseSchemaFieldsStruct] or [StreamListResponseSchemaFieldsList].
type StreamListResponseSchemaFieldsUnion interface {
	implementsStreamListResponseSchemaField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamListResponseSchemaFieldsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsInt32{}),
			DiscriminatorValue: "int32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsInt64{}),
			DiscriminatorValue: "int64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsFloat32{}),
			DiscriminatorValue: "float32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsFloat64{}),
			DiscriminatorValue: "float64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsBool{}),
			DiscriminatorValue: "bool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsString{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsBinary{}),
			DiscriminatorValue: "binary",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsTimestamp{}),
			DiscriminatorValue: "timestamp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsStruct{}),
			DiscriminatorValue: "struct",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFieldsList{}),
			DiscriminatorValue: "list",
		},
	)
}

type StreamListResponseSchemaFieldsInt32 struct {
	Type        StreamListResponseSchemaFieldsInt32Type `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsInt32JSON `json:"-"`
}

// streamListResponseSchemaFieldsInt32JSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsInt32]
type streamListResponseSchemaFieldsInt32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsInt32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsInt32JSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsInt32) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsInt32Type string

const (
	StreamListResponseSchemaFieldsInt32TypeInt32 StreamListResponseSchemaFieldsInt32Type = "int32"
)

func (r StreamListResponseSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsInt64 struct {
	Type        StreamListResponseSchemaFieldsInt64Type `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsInt64JSON `json:"-"`
}

// streamListResponseSchemaFieldsInt64JSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsInt64]
type streamListResponseSchemaFieldsInt64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsInt64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsInt64JSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsInt64) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsInt64Type string

const (
	StreamListResponseSchemaFieldsInt64TypeInt64 StreamListResponseSchemaFieldsInt64Type = "int64"
)

func (r StreamListResponseSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsFloat32 struct {
	Type        StreamListResponseSchemaFieldsFloat32Type `json:"type,required"`
	MetadataKey string                                    `json:"metadata_key,nullable"`
	Name        string                                    `json:"name"`
	Required    bool                                      `json:"required"`
	SqlName     string                                    `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsFloat32JSON `json:"-"`
}

// streamListResponseSchemaFieldsFloat32JSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsFloat32]
type streamListResponseSchemaFieldsFloat32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsFloat32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsFloat32JSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsFloat32) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsFloat32Type string

const (
	StreamListResponseSchemaFieldsFloat32TypeFloat32 StreamListResponseSchemaFieldsFloat32Type = "float32"
)

func (r StreamListResponseSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsFloat64 struct {
	Type        StreamListResponseSchemaFieldsFloat64Type `json:"type,required"`
	MetadataKey string                                    `json:"metadata_key,nullable"`
	Name        string                                    `json:"name"`
	Required    bool                                      `json:"required"`
	SqlName     string                                    `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsFloat64JSON `json:"-"`
}

// streamListResponseSchemaFieldsFloat64JSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsFloat64]
type streamListResponseSchemaFieldsFloat64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsFloat64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsFloat64JSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsFloat64) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsFloat64Type string

const (
	StreamListResponseSchemaFieldsFloat64TypeFloat64 StreamListResponseSchemaFieldsFloat64Type = "float64"
)

func (r StreamListResponseSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsBool struct {
	Type        StreamListResponseSchemaFieldsBoolType `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsBoolJSON `json:"-"`
}

// streamListResponseSchemaFieldsBoolJSON contains the JSON metadata for the struct
// [StreamListResponseSchemaFieldsBool]
type streamListResponseSchemaFieldsBoolJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsBoolJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsBool) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsBoolType string

const (
	StreamListResponseSchemaFieldsBoolTypeBool StreamListResponseSchemaFieldsBoolType = "bool"
)

func (r StreamListResponseSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsString struct {
	Type        StreamListResponseSchemaFieldsStringType `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsStringJSON `json:"-"`
}

// streamListResponseSchemaFieldsStringJSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsString]
type streamListResponseSchemaFieldsStringJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsStringJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsString) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsStringType string

const (
	StreamListResponseSchemaFieldsStringTypeString StreamListResponseSchemaFieldsStringType = "string"
)

func (r StreamListResponseSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsBinary struct {
	Type        StreamListResponseSchemaFieldsBinaryType `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsBinaryJSON `json:"-"`
}

// streamListResponseSchemaFieldsBinaryJSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsBinary]
type streamListResponseSchemaFieldsBinaryJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsBinary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsBinaryJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsBinary) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsBinaryType string

const (
	StreamListResponseSchemaFieldsBinaryTypeBinary StreamListResponseSchemaFieldsBinaryType = "binary"
)

func (r StreamListResponseSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsTimestamp struct {
	Type        StreamListResponseSchemaFieldsTimestampType `json:"type,required"`
	MetadataKey string                                      `json:"metadata_key,nullable"`
	Name        string                                      `json:"name"`
	Required    bool                                        `json:"required"`
	SqlName     string                                      `json:"sql_name"`
	Unit        StreamListResponseSchemaFieldsTimestampUnit `json:"unit"`
	JSON        streamListResponseSchemaFieldsTimestampJSON `json:"-"`
}

// streamListResponseSchemaFieldsTimestampJSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsTimestamp]
type streamListResponseSchemaFieldsTimestampJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsTimestamp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsTimestampJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsTimestamp) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsTimestampType string

const (
	StreamListResponseSchemaFieldsTimestampTypeTimestamp StreamListResponseSchemaFieldsTimestampType = "timestamp"
)

func (r StreamListResponseSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsTimestampUnit string

const (
	StreamListResponseSchemaFieldsTimestampUnitSecond      StreamListResponseSchemaFieldsTimestampUnit = "second"
	StreamListResponseSchemaFieldsTimestampUnitMillisecond StreamListResponseSchemaFieldsTimestampUnit = "millisecond"
	StreamListResponseSchemaFieldsTimestampUnitMicrosecond StreamListResponseSchemaFieldsTimestampUnit = "microsecond"
	StreamListResponseSchemaFieldsTimestampUnitNanosecond  StreamListResponseSchemaFieldsTimestampUnit = "nanosecond"
)

func (r StreamListResponseSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsTimestampUnitSecond, StreamListResponseSchemaFieldsTimestampUnitMillisecond, StreamListResponseSchemaFieldsTimestampUnitMicrosecond, StreamListResponseSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsJson struct {
	Type        StreamListResponseSchemaFieldsJsonType `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        streamListResponseSchemaFieldsJsonJSON `json:"-"`
}

// streamListResponseSchemaFieldsJsonJSON contains the JSON metadata for the struct
// [StreamListResponseSchemaFieldsJson]
type streamListResponseSchemaFieldsJsonJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsJson) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsJsonType string

const (
	StreamListResponseSchemaFieldsJsonTypeJson StreamListResponseSchemaFieldsJsonType = "json"
)

func (r StreamListResponseSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsStruct struct {
	JSON streamListResponseSchemaFieldsStructJSON `json:"-"`
}

// streamListResponseSchemaFieldsStructJSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFieldsStruct]
type streamListResponseSchemaFieldsStructJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsStruct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsStructJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsStruct) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsList struct {
	JSON streamListResponseSchemaFieldsListJSON `json:"-"`
}

// streamListResponseSchemaFieldsListJSON contains the JSON metadata for the struct
// [StreamListResponseSchemaFieldsList]
type streamListResponseSchemaFieldsListJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamListResponseSchemaFieldsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFieldsListJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFieldsList) implementsStreamListResponseSchemaField() {}

type StreamListResponseSchemaFieldsType string

const (
	StreamListResponseSchemaFieldsTypeInt32     StreamListResponseSchemaFieldsType = "int32"
	StreamListResponseSchemaFieldsTypeInt64     StreamListResponseSchemaFieldsType = "int64"
	StreamListResponseSchemaFieldsTypeFloat32   StreamListResponseSchemaFieldsType = "float32"
	StreamListResponseSchemaFieldsTypeFloat64   StreamListResponseSchemaFieldsType = "float64"
	StreamListResponseSchemaFieldsTypeBool      StreamListResponseSchemaFieldsType = "bool"
	StreamListResponseSchemaFieldsTypeString    StreamListResponseSchemaFieldsType = "string"
	StreamListResponseSchemaFieldsTypeBinary    StreamListResponseSchemaFieldsType = "binary"
	StreamListResponseSchemaFieldsTypeTimestamp StreamListResponseSchemaFieldsType = "timestamp"
	StreamListResponseSchemaFieldsTypeJson      StreamListResponseSchemaFieldsType = "json"
	StreamListResponseSchemaFieldsTypeStruct    StreamListResponseSchemaFieldsType = "struct"
	StreamListResponseSchemaFieldsTypeList      StreamListResponseSchemaFieldsType = "list"
)

func (r StreamListResponseSchemaFieldsType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsTypeInt32, StreamListResponseSchemaFieldsTypeInt64, StreamListResponseSchemaFieldsTypeFloat32, StreamListResponseSchemaFieldsTypeFloat64, StreamListResponseSchemaFieldsTypeBool, StreamListResponseSchemaFieldsTypeString, StreamListResponseSchemaFieldsTypeBinary, StreamListResponseSchemaFieldsTypeTimestamp, StreamListResponseSchemaFieldsTypeJson, StreamListResponseSchemaFieldsTypeStruct, StreamListResponseSchemaFieldsTypeList:
		return true
	}
	return false
}

type StreamListResponseSchemaFieldsUnit string

const (
	StreamListResponseSchemaFieldsUnitSecond      StreamListResponseSchemaFieldsUnit = "second"
	StreamListResponseSchemaFieldsUnitMillisecond StreamListResponseSchemaFieldsUnit = "millisecond"
	StreamListResponseSchemaFieldsUnitMicrosecond StreamListResponseSchemaFieldsUnit = "microsecond"
	StreamListResponseSchemaFieldsUnitNanosecond  StreamListResponseSchemaFieldsUnit = "nanosecond"
)

func (r StreamListResponseSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFieldsUnitSecond, StreamListResponseSchemaFieldsUnitMillisecond, StreamListResponseSchemaFieldsUnitMicrosecond, StreamListResponseSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type StreamListResponseSchemaFormat struct {
	Type            StreamListResponseSchemaFormatType            `json:"type,required"`
	Compression     StreamListResponseSchemaFormatCompression     `json:"compression"`
	DecimalEncoding StreamListResponseSchemaFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                         `json:"row_group_bytes,nullable"`
	TimestampFormat StreamListResponseSchemaFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                          `json:"unstructured"`
	JSON            streamListResponseSchemaFormatJSON            `json:"-"`
	union           StreamListResponseSchemaFormatUnion
}

// streamListResponseSchemaFormatJSON contains the JSON metadata for the struct
// [StreamListResponseSchemaFormat]
type streamListResponseSchemaFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r streamListResponseSchemaFormatJSON) RawJSON() string {
	return r.raw
}

func (r *StreamListResponseSchemaFormat) UnmarshalJSON(data []byte) (err error) {
	*r = StreamListResponseSchemaFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamListResponseSchemaFormatUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamListResponseSchemaFormatJson],
// [StreamListResponseSchemaFormatParquet].
func (r StreamListResponseSchemaFormat) AsUnion() StreamListResponseSchemaFormatUnion {
	return r.union
}

// Union satisfied by [StreamListResponseSchemaFormatJson] or
// [StreamListResponseSchemaFormatParquet].
type StreamListResponseSchemaFormatUnion interface {
	implementsStreamListResponseSchemaFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamListResponseSchemaFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamListResponseSchemaFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type StreamListResponseSchemaFormatJson struct {
	Type            StreamListResponseSchemaFormatJsonType            `json:"type,required"`
	DecimalEncoding StreamListResponseSchemaFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat StreamListResponseSchemaFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                              `json:"unstructured"`
	JSON            streamListResponseSchemaFormatJsonJSON            `json:"-"`
}

// streamListResponseSchemaFormatJsonJSON contains the JSON metadata for the struct
// [StreamListResponseSchemaFormatJson]
type streamListResponseSchemaFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *StreamListResponseSchemaFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFormatJson) implementsStreamListResponseSchemaFormat() {}

type StreamListResponseSchemaFormatJsonType string

const (
	StreamListResponseSchemaFormatJsonTypeJson StreamListResponseSchemaFormatJsonType = "json"
)

func (r StreamListResponseSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatJsonDecimalEncoding string

const (
	StreamListResponseSchemaFormatJsonDecimalEncodingNumber StreamListResponseSchemaFormatJsonDecimalEncoding = "number"
	StreamListResponseSchemaFormatJsonDecimalEncodingString StreamListResponseSchemaFormatJsonDecimalEncoding = "string"
	StreamListResponseSchemaFormatJsonDecimalEncodingBytes  StreamListResponseSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r StreamListResponseSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatJsonDecimalEncodingNumber, StreamListResponseSchemaFormatJsonDecimalEncodingString, StreamListResponseSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatJsonTimestampFormat string

const (
	StreamListResponseSchemaFormatJsonTimestampFormatRfc3339    StreamListResponseSchemaFormatJsonTimestampFormat = "rfc3339"
	StreamListResponseSchemaFormatJsonTimestampFormatUnixMillis StreamListResponseSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamListResponseSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatJsonTimestampFormatRfc3339, StreamListResponseSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatParquet struct {
	Type          StreamListResponseSchemaFormatParquetType        `json:"type,required"`
	Compression   StreamListResponseSchemaFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                            `json:"row_group_bytes,nullable"`
	JSON          streamListResponseSchemaFormatParquetJSON        `json:"-"`
}

// streamListResponseSchemaFormatParquetJSON contains the JSON metadata for the
// struct [StreamListResponseSchemaFormatParquet]
type streamListResponseSchemaFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamListResponseSchemaFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamListResponseSchemaFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r StreamListResponseSchemaFormatParquet) implementsStreamListResponseSchemaFormat() {}

type StreamListResponseSchemaFormatParquetType string

const (
	StreamListResponseSchemaFormatParquetTypeParquet StreamListResponseSchemaFormatParquetType = "parquet"
)

func (r StreamListResponseSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatParquetCompression string

const (
	StreamListResponseSchemaFormatParquetCompressionUncompressed StreamListResponseSchemaFormatParquetCompression = "uncompressed"
	StreamListResponseSchemaFormatParquetCompressionSnappy       StreamListResponseSchemaFormatParquetCompression = "snappy"
	StreamListResponseSchemaFormatParquetCompressionGzip         StreamListResponseSchemaFormatParquetCompression = "gzip"
	StreamListResponseSchemaFormatParquetCompressionZstd         StreamListResponseSchemaFormatParquetCompression = "zstd"
	StreamListResponseSchemaFormatParquetCompressionLz4          StreamListResponseSchemaFormatParquetCompression = "lz4"
)

func (r StreamListResponseSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatParquetCompressionUncompressed, StreamListResponseSchemaFormatParquetCompressionSnappy, StreamListResponseSchemaFormatParquetCompressionGzip, StreamListResponseSchemaFormatParquetCompressionZstd, StreamListResponseSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatType string

const (
	StreamListResponseSchemaFormatTypeJson    StreamListResponseSchemaFormatType = "json"
	StreamListResponseSchemaFormatTypeParquet StreamListResponseSchemaFormatType = "parquet"
)

func (r StreamListResponseSchemaFormatType) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatTypeJson, StreamListResponseSchemaFormatTypeParquet:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatCompression string

const (
	StreamListResponseSchemaFormatCompressionUncompressed StreamListResponseSchemaFormatCompression = "uncompressed"
	StreamListResponseSchemaFormatCompressionSnappy       StreamListResponseSchemaFormatCompression = "snappy"
	StreamListResponseSchemaFormatCompressionGzip         StreamListResponseSchemaFormatCompression = "gzip"
	StreamListResponseSchemaFormatCompressionZstd         StreamListResponseSchemaFormatCompression = "zstd"
	StreamListResponseSchemaFormatCompressionLz4          StreamListResponseSchemaFormatCompression = "lz4"
)

func (r StreamListResponseSchemaFormatCompression) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatCompressionUncompressed, StreamListResponseSchemaFormatCompressionSnappy, StreamListResponseSchemaFormatCompressionGzip, StreamListResponseSchemaFormatCompressionZstd, StreamListResponseSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatDecimalEncoding string

const (
	StreamListResponseSchemaFormatDecimalEncodingNumber StreamListResponseSchemaFormatDecimalEncoding = "number"
	StreamListResponseSchemaFormatDecimalEncodingString StreamListResponseSchemaFormatDecimalEncoding = "string"
	StreamListResponseSchemaFormatDecimalEncodingBytes  StreamListResponseSchemaFormatDecimalEncoding = "bytes"
)

func (r StreamListResponseSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatDecimalEncodingNumber, StreamListResponseSchemaFormatDecimalEncodingString, StreamListResponseSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamListResponseSchemaFormatTimestampFormat string

const (
	StreamListResponseSchemaFormatTimestampFormatRfc3339    StreamListResponseSchemaFormatTimestampFormat = "rfc3339"
	StreamListResponseSchemaFormatTimestampFormatUnixMillis StreamListResponseSchemaFormatTimestampFormat = "unix_millis"
)

func (r StreamListResponseSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamListResponseSchemaFormatTimestampFormatRfc3339, StreamListResponseSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamGetResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                `json:"id,required"`
	CreatedAt  time.Time             `json:"created_at,required" format:"date-time"`
	HTTP       StreamGetResponseHTTP `json:"http,required"`
	ModifiedAt time.Time             `json:"modified_at,required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name,required"`
	// Indicates the current version of this stream.
	Version       int64                          `json:"version,required"`
	WorkerBinding StreamGetResponseWorkerBinding `json:"worker_binding,required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string                  `json:"endpoint" format:"uri"`
	Format   StreamGetResponseFormat `json:"format"`
	Schema   StreamGetResponseSchema `json:"schema"`
	JSON     streamGetResponseJSON   `json:"-"`
}

// streamGetResponseJSON contains the JSON metadata for the struct
// [StreamGetResponse]
type streamGetResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	HTTP          apijson.Field
	ModifiedAt    apijson.Field
	Name          apijson.Field
	Version       apijson.Field
	WorkerBinding apijson.Field
	Endpoint      apijson.Field
	Format        apijson.Field
	Schema        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseHTTP struct {
	// Indicates that authentication is required for the HTTP endpoint.
	Authentication bool `json:"authentication,required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled,required"`
	// Specifies the CORS options for the HTTP endpoint.
	CORS StreamGetResponseHTTPCORS `json:"cors"`
	JSON streamGetResponseHTTPJSON `json:"-"`
}

// streamGetResponseHTTPJSON contains the JSON metadata for the struct
// [StreamGetResponseHTTP]
type streamGetResponseHTTPJSON struct {
	Authentication apijson.Field
	Enabled        apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *StreamGetResponseHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseHTTPJSON) RawJSON() string {
	return r.raw
}

// Specifies the CORS options for the HTTP endpoint.
type StreamGetResponseHTTPCORS struct {
	Origins []string                      `json:"origins"`
	JSON    streamGetResponseHttpcorsJSON `json:"-"`
}

// streamGetResponseHttpcorsJSON contains the JSON metadata for the struct
// [StreamGetResponseHTTPCORS]
type streamGetResponseHttpcorsJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseHTTPCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseHttpcorsJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	Enabled bool                               `json:"enabled,required"`
	JSON    streamGetResponseWorkerBindingJSON `json:"-"`
}

// streamGetResponseWorkerBindingJSON contains the JSON metadata for the struct
// [StreamGetResponseWorkerBinding]
type streamGetResponseWorkerBindingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseWorkerBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseWorkerBindingJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseFormat struct {
	Type            StreamGetResponseFormatType            `json:"type,required"`
	Compression     StreamGetResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamGetResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                  `json:"row_group_bytes,nullable"`
	TimestampFormat StreamGetResponseFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                   `json:"unstructured"`
	JSON            streamGetResponseFormatJSON            `json:"-"`
	union           StreamGetResponseFormatUnion
}

// streamGetResponseFormatJSON contains the JSON metadata for the struct
// [StreamGetResponseFormat]
type streamGetResponseFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r streamGetResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r *StreamGetResponseFormat) UnmarshalJSON(data []byte) (err error) {
	*r = StreamGetResponseFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamGetResponseFormatUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [StreamGetResponseFormatJson],
// [StreamGetResponseFormatParquet].
func (r StreamGetResponseFormat) AsUnion() StreamGetResponseFormatUnion {
	return r.union
}

// Union satisfied by [StreamGetResponseFormatJson] or
// [StreamGetResponseFormatParquet].
type StreamGetResponseFormatUnion interface {
	implementsStreamGetResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamGetResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type StreamGetResponseFormatJson struct {
	Type            StreamGetResponseFormatJsonType            `json:"type,required"`
	DecimalEncoding StreamGetResponseFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat StreamGetResponseFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                       `json:"unstructured"`
	JSON            streamGetResponseFormatJsonJSON            `json:"-"`
}

// streamGetResponseFormatJsonJSON contains the JSON metadata for the struct
// [StreamGetResponseFormatJson]
type streamGetResponseFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *StreamGetResponseFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseFormatJson) implementsStreamGetResponseFormat() {}

type StreamGetResponseFormatJsonType string

const (
	StreamGetResponseFormatJsonTypeJson StreamGetResponseFormatJsonType = "json"
)

func (r StreamGetResponseFormatJsonType) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamGetResponseFormatJsonDecimalEncoding string

const (
	StreamGetResponseFormatJsonDecimalEncodingNumber StreamGetResponseFormatJsonDecimalEncoding = "number"
	StreamGetResponseFormatJsonDecimalEncodingString StreamGetResponseFormatJsonDecimalEncoding = "string"
	StreamGetResponseFormatJsonDecimalEncodingBytes  StreamGetResponseFormatJsonDecimalEncoding = "bytes"
)

func (r StreamGetResponseFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatJsonDecimalEncodingNumber, StreamGetResponseFormatJsonDecimalEncodingString, StreamGetResponseFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamGetResponseFormatJsonTimestampFormat string

const (
	StreamGetResponseFormatJsonTimestampFormatRfc3339    StreamGetResponseFormatJsonTimestampFormat = "rfc3339"
	StreamGetResponseFormatJsonTimestampFormatUnixMillis StreamGetResponseFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamGetResponseFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatJsonTimestampFormatRfc3339, StreamGetResponseFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamGetResponseFormatParquet struct {
	Type          StreamGetResponseFormatParquetType        `json:"type,required"`
	Compression   StreamGetResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                     `json:"row_group_bytes,nullable"`
	JSON          streamGetResponseFormatParquetJSON        `json:"-"`
}

// streamGetResponseFormatParquetJSON contains the JSON metadata for the struct
// [StreamGetResponseFormatParquet]
type streamGetResponseFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamGetResponseFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseFormatParquet) implementsStreamGetResponseFormat() {}

type StreamGetResponseFormatParquetType string

const (
	StreamGetResponseFormatParquetTypeParquet StreamGetResponseFormatParquetType = "parquet"
)

func (r StreamGetResponseFormatParquetType) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamGetResponseFormatParquetCompression string

const (
	StreamGetResponseFormatParquetCompressionUncompressed StreamGetResponseFormatParquetCompression = "uncompressed"
	StreamGetResponseFormatParquetCompressionSnappy       StreamGetResponseFormatParquetCompression = "snappy"
	StreamGetResponseFormatParquetCompressionGzip         StreamGetResponseFormatParquetCompression = "gzip"
	StreamGetResponseFormatParquetCompressionZstd         StreamGetResponseFormatParquetCompression = "zstd"
	StreamGetResponseFormatParquetCompressionLz4          StreamGetResponseFormatParquetCompression = "lz4"
)

func (r StreamGetResponseFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatParquetCompressionUncompressed, StreamGetResponseFormatParquetCompressionSnappy, StreamGetResponseFormatParquetCompressionGzip, StreamGetResponseFormatParquetCompressionZstd, StreamGetResponseFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamGetResponseFormatType string

const (
	StreamGetResponseFormatTypeJson    StreamGetResponseFormatType = "json"
	StreamGetResponseFormatTypeParquet StreamGetResponseFormatType = "parquet"
)

func (r StreamGetResponseFormatType) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatTypeJson, StreamGetResponseFormatTypeParquet:
		return true
	}
	return false
}

type StreamGetResponseFormatCompression string

const (
	StreamGetResponseFormatCompressionUncompressed StreamGetResponseFormatCompression = "uncompressed"
	StreamGetResponseFormatCompressionSnappy       StreamGetResponseFormatCompression = "snappy"
	StreamGetResponseFormatCompressionGzip         StreamGetResponseFormatCompression = "gzip"
	StreamGetResponseFormatCompressionZstd         StreamGetResponseFormatCompression = "zstd"
	StreamGetResponseFormatCompressionLz4          StreamGetResponseFormatCompression = "lz4"
)

func (r StreamGetResponseFormatCompression) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatCompressionUncompressed, StreamGetResponseFormatCompressionSnappy, StreamGetResponseFormatCompressionGzip, StreamGetResponseFormatCompressionZstd, StreamGetResponseFormatCompressionLz4:
		return true
	}
	return false
}

type StreamGetResponseFormatDecimalEncoding string

const (
	StreamGetResponseFormatDecimalEncodingNumber StreamGetResponseFormatDecimalEncoding = "number"
	StreamGetResponseFormatDecimalEncodingString StreamGetResponseFormatDecimalEncoding = "string"
	StreamGetResponseFormatDecimalEncodingBytes  StreamGetResponseFormatDecimalEncoding = "bytes"
)

func (r StreamGetResponseFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatDecimalEncodingNumber, StreamGetResponseFormatDecimalEncodingString, StreamGetResponseFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamGetResponseFormatTimestampFormat string

const (
	StreamGetResponseFormatTimestampFormatRfc3339    StreamGetResponseFormatTimestampFormat = "rfc3339"
	StreamGetResponseFormatTimestampFormatUnixMillis StreamGetResponseFormatTimestampFormat = "unix_millis"
)

func (r StreamGetResponseFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamGetResponseFormatTimestampFormatRfc3339, StreamGetResponseFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamGetResponseSchema struct {
	Fields   []StreamGetResponseSchemaField `json:"fields"`
	Format   StreamGetResponseSchemaFormat  `json:"format"`
	Inferred bool                           `json:"inferred,nullable"`
	JSON     streamGetResponseSchemaJSON    `json:"-"`
}

// streamGetResponseSchemaJSON contains the JSON metadata for the struct
// [StreamGetResponseSchema]
type streamGetResponseSchemaJSON struct {
	Fields      apijson.Field
	Format      apijson.Field
	Inferred    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type StreamGetResponseSchemaField struct {
	Type        StreamGetResponseSchemaFieldsType `json:"type,required"`
	MetadataKey string                            `json:"metadata_key,nullable"`
	Name        string                            `json:"name"`
	Required    bool                              `json:"required"`
	SqlName     string                            `json:"sql_name"`
	Unit        StreamGetResponseSchemaFieldsUnit `json:"unit"`
	JSON        streamGetResponseSchemaFieldJSON  `json:"-"`
	union       StreamGetResponseSchemaFieldsUnion
}

// streamGetResponseSchemaFieldJSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaField]
type streamGetResponseSchemaFieldJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r streamGetResponseSchemaFieldJSON) RawJSON() string {
	return r.raw
}

func (r *StreamGetResponseSchemaField) UnmarshalJSON(data []byte) (err error) {
	*r = StreamGetResponseSchemaField{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamGetResponseSchemaFieldsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamGetResponseSchemaFieldsInt32],
// [StreamGetResponseSchemaFieldsInt64], [StreamGetResponseSchemaFieldsFloat32],
// [StreamGetResponseSchemaFieldsFloat64], [StreamGetResponseSchemaFieldsBool],
// [StreamGetResponseSchemaFieldsString], [StreamGetResponseSchemaFieldsBinary],
// [StreamGetResponseSchemaFieldsTimestamp], [StreamGetResponseSchemaFieldsJson],
// [StreamGetResponseSchemaFieldsStruct], [StreamGetResponseSchemaFieldsList].
func (r StreamGetResponseSchemaField) AsUnion() StreamGetResponseSchemaFieldsUnion {
	return r.union
}

// Union satisfied by [StreamGetResponseSchemaFieldsInt32],
// [StreamGetResponseSchemaFieldsInt64], [StreamGetResponseSchemaFieldsFloat32],
// [StreamGetResponseSchemaFieldsFloat64], [StreamGetResponseSchemaFieldsBool],
// [StreamGetResponseSchemaFieldsString], [StreamGetResponseSchemaFieldsBinary],
// [StreamGetResponseSchemaFieldsTimestamp], [StreamGetResponseSchemaFieldsJson],
// [StreamGetResponseSchemaFieldsStruct] or [StreamGetResponseSchemaFieldsList].
type StreamGetResponseSchemaFieldsUnion interface {
	implementsStreamGetResponseSchemaField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamGetResponseSchemaFieldsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsInt32{}),
			DiscriminatorValue: "int32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsInt64{}),
			DiscriminatorValue: "int64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsFloat32{}),
			DiscriminatorValue: "float32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsFloat64{}),
			DiscriminatorValue: "float64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsBool{}),
			DiscriminatorValue: "bool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsString{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsBinary{}),
			DiscriminatorValue: "binary",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsTimestamp{}),
			DiscriminatorValue: "timestamp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsStruct{}),
			DiscriminatorValue: "struct",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFieldsList{}),
			DiscriminatorValue: "list",
		},
	)
}

type StreamGetResponseSchemaFieldsInt32 struct {
	Type        StreamGetResponseSchemaFieldsInt32Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsInt32JSON `json:"-"`
}

// streamGetResponseSchemaFieldsInt32JSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaFieldsInt32]
type streamGetResponseSchemaFieldsInt32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsInt32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsInt32JSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsInt32) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsInt32Type string

const (
	StreamGetResponseSchemaFieldsInt32TypeInt32 StreamGetResponseSchemaFieldsInt32Type = "int32"
)

func (r StreamGetResponseSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsInt64 struct {
	Type        StreamGetResponseSchemaFieldsInt64Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsInt64JSON `json:"-"`
}

// streamGetResponseSchemaFieldsInt64JSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaFieldsInt64]
type streamGetResponseSchemaFieldsInt64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsInt64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsInt64JSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsInt64) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsInt64Type string

const (
	StreamGetResponseSchemaFieldsInt64TypeInt64 StreamGetResponseSchemaFieldsInt64Type = "int64"
)

func (r StreamGetResponseSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsFloat32 struct {
	Type        StreamGetResponseSchemaFieldsFloat32Type `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsFloat32JSON `json:"-"`
}

// streamGetResponseSchemaFieldsFloat32JSON contains the JSON metadata for the
// struct [StreamGetResponseSchemaFieldsFloat32]
type streamGetResponseSchemaFieldsFloat32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsFloat32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsFloat32JSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsFloat32) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsFloat32Type string

const (
	StreamGetResponseSchemaFieldsFloat32TypeFloat32 StreamGetResponseSchemaFieldsFloat32Type = "float32"
)

func (r StreamGetResponseSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsFloat64 struct {
	Type        StreamGetResponseSchemaFieldsFloat64Type `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsFloat64JSON `json:"-"`
}

// streamGetResponseSchemaFieldsFloat64JSON contains the JSON metadata for the
// struct [StreamGetResponseSchemaFieldsFloat64]
type streamGetResponseSchemaFieldsFloat64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsFloat64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsFloat64JSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsFloat64) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsFloat64Type string

const (
	StreamGetResponseSchemaFieldsFloat64TypeFloat64 StreamGetResponseSchemaFieldsFloat64Type = "float64"
)

func (r StreamGetResponseSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsBool struct {
	Type        StreamGetResponseSchemaFieldsBoolType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsBoolJSON `json:"-"`
}

// streamGetResponseSchemaFieldsBoolJSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaFieldsBool]
type streamGetResponseSchemaFieldsBoolJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsBoolJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsBool) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsBoolType string

const (
	StreamGetResponseSchemaFieldsBoolTypeBool StreamGetResponseSchemaFieldsBoolType = "bool"
)

func (r StreamGetResponseSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsString struct {
	Type        StreamGetResponseSchemaFieldsStringType `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsStringJSON `json:"-"`
}

// streamGetResponseSchemaFieldsStringJSON contains the JSON metadata for the
// struct [StreamGetResponseSchemaFieldsString]
type streamGetResponseSchemaFieldsStringJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsStringJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsString) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsStringType string

const (
	StreamGetResponseSchemaFieldsStringTypeString StreamGetResponseSchemaFieldsStringType = "string"
)

func (r StreamGetResponseSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsBinary struct {
	Type        StreamGetResponseSchemaFieldsBinaryType `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsBinaryJSON `json:"-"`
}

// streamGetResponseSchemaFieldsBinaryJSON contains the JSON metadata for the
// struct [StreamGetResponseSchemaFieldsBinary]
type streamGetResponseSchemaFieldsBinaryJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsBinary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsBinaryJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsBinary) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsBinaryType string

const (
	StreamGetResponseSchemaFieldsBinaryTypeBinary StreamGetResponseSchemaFieldsBinaryType = "binary"
)

func (r StreamGetResponseSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsTimestamp struct {
	Type        StreamGetResponseSchemaFieldsTimestampType `json:"type,required"`
	MetadataKey string                                     `json:"metadata_key,nullable"`
	Name        string                                     `json:"name"`
	Required    bool                                       `json:"required"`
	SqlName     string                                     `json:"sql_name"`
	Unit        StreamGetResponseSchemaFieldsTimestampUnit `json:"unit"`
	JSON        streamGetResponseSchemaFieldsTimestampJSON `json:"-"`
}

// streamGetResponseSchemaFieldsTimestampJSON contains the JSON metadata for the
// struct [StreamGetResponseSchemaFieldsTimestamp]
type streamGetResponseSchemaFieldsTimestampJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsTimestamp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsTimestampJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsTimestamp) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsTimestampType string

const (
	StreamGetResponseSchemaFieldsTimestampTypeTimestamp StreamGetResponseSchemaFieldsTimestampType = "timestamp"
)

func (r StreamGetResponseSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsTimestampUnit string

const (
	StreamGetResponseSchemaFieldsTimestampUnitSecond      StreamGetResponseSchemaFieldsTimestampUnit = "second"
	StreamGetResponseSchemaFieldsTimestampUnitMillisecond StreamGetResponseSchemaFieldsTimestampUnit = "millisecond"
	StreamGetResponseSchemaFieldsTimestampUnitMicrosecond StreamGetResponseSchemaFieldsTimestampUnit = "microsecond"
	StreamGetResponseSchemaFieldsTimestampUnitNanosecond  StreamGetResponseSchemaFieldsTimestampUnit = "nanosecond"
)

func (r StreamGetResponseSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsTimestampUnitSecond, StreamGetResponseSchemaFieldsTimestampUnitMillisecond, StreamGetResponseSchemaFieldsTimestampUnitMicrosecond, StreamGetResponseSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsJson struct {
	Type        StreamGetResponseSchemaFieldsJsonType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        streamGetResponseSchemaFieldsJsonJSON `json:"-"`
}

// streamGetResponseSchemaFieldsJsonJSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaFieldsJson]
type streamGetResponseSchemaFieldsJsonJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsJson) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsJsonType string

const (
	StreamGetResponseSchemaFieldsJsonTypeJson StreamGetResponseSchemaFieldsJsonType = "json"
)

func (r StreamGetResponseSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsStruct struct {
	JSON streamGetResponseSchemaFieldsStructJSON `json:"-"`
}

// streamGetResponseSchemaFieldsStructJSON contains the JSON metadata for the
// struct [StreamGetResponseSchemaFieldsStruct]
type streamGetResponseSchemaFieldsStructJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsStruct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsStructJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsStruct) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsList struct {
	JSON streamGetResponseSchemaFieldsListJSON `json:"-"`
}

// streamGetResponseSchemaFieldsListJSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaFieldsList]
type streamGetResponseSchemaFieldsListJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFieldsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFieldsListJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFieldsList) implementsStreamGetResponseSchemaField() {}

type StreamGetResponseSchemaFieldsType string

const (
	StreamGetResponseSchemaFieldsTypeInt32     StreamGetResponseSchemaFieldsType = "int32"
	StreamGetResponseSchemaFieldsTypeInt64     StreamGetResponseSchemaFieldsType = "int64"
	StreamGetResponseSchemaFieldsTypeFloat32   StreamGetResponseSchemaFieldsType = "float32"
	StreamGetResponseSchemaFieldsTypeFloat64   StreamGetResponseSchemaFieldsType = "float64"
	StreamGetResponseSchemaFieldsTypeBool      StreamGetResponseSchemaFieldsType = "bool"
	StreamGetResponseSchemaFieldsTypeString    StreamGetResponseSchemaFieldsType = "string"
	StreamGetResponseSchemaFieldsTypeBinary    StreamGetResponseSchemaFieldsType = "binary"
	StreamGetResponseSchemaFieldsTypeTimestamp StreamGetResponseSchemaFieldsType = "timestamp"
	StreamGetResponseSchemaFieldsTypeJson      StreamGetResponseSchemaFieldsType = "json"
	StreamGetResponseSchemaFieldsTypeStruct    StreamGetResponseSchemaFieldsType = "struct"
	StreamGetResponseSchemaFieldsTypeList      StreamGetResponseSchemaFieldsType = "list"
)

func (r StreamGetResponseSchemaFieldsType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsTypeInt32, StreamGetResponseSchemaFieldsTypeInt64, StreamGetResponseSchemaFieldsTypeFloat32, StreamGetResponseSchemaFieldsTypeFloat64, StreamGetResponseSchemaFieldsTypeBool, StreamGetResponseSchemaFieldsTypeString, StreamGetResponseSchemaFieldsTypeBinary, StreamGetResponseSchemaFieldsTypeTimestamp, StreamGetResponseSchemaFieldsTypeJson, StreamGetResponseSchemaFieldsTypeStruct, StreamGetResponseSchemaFieldsTypeList:
		return true
	}
	return false
}

type StreamGetResponseSchemaFieldsUnit string

const (
	StreamGetResponseSchemaFieldsUnitSecond      StreamGetResponseSchemaFieldsUnit = "second"
	StreamGetResponseSchemaFieldsUnitMillisecond StreamGetResponseSchemaFieldsUnit = "millisecond"
	StreamGetResponseSchemaFieldsUnitMicrosecond StreamGetResponseSchemaFieldsUnit = "microsecond"
	StreamGetResponseSchemaFieldsUnitNanosecond  StreamGetResponseSchemaFieldsUnit = "nanosecond"
)

func (r StreamGetResponseSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFieldsUnitSecond, StreamGetResponseSchemaFieldsUnitMillisecond, StreamGetResponseSchemaFieldsUnitMicrosecond, StreamGetResponseSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormat struct {
	Type            StreamGetResponseSchemaFormatType            `json:"type,required"`
	Compression     StreamGetResponseSchemaFormatCompression     `json:"compression"`
	DecimalEncoding StreamGetResponseSchemaFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                        `json:"row_group_bytes,nullable"`
	TimestampFormat StreamGetResponseSchemaFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                         `json:"unstructured"`
	JSON            streamGetResponseSchemaFormatJSON            `json:"-"`
	union           StreamGetResponseSchemaFormatUnion
}

// streamGetResponseSchemaFormatJSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaFormat]
type streamGetResponseSchemaFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r streamGetResponseSchemaFormatJSON) RawJSON() string {
	return r.raw
}

func (r *StreamGetResponseSchemaFormat) UnmarshalJSON(data []byte) (err error) {
	*r = StreamGetResponseSchemaFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [StreamGetResponseSchemaFormatUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [StreamGetResponseSchemaFormatJson],
// [StreamGetResponseSchemaFormatParquet].
func (r StreamGetResponseSchemaFormat) AsUnion() StreamGetResponseSchemaFormatUnion {
	return r.union
}

// Union satisfied by [StreamGetResponseSchemaFormatJson] or
// [StreamGetResponseSchemaFormatParquet].
type StreamGetResponseSchemaFormatUnion interface {
	implementsStreamGetResponseSchemaFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamGetResponseSchemaFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(StreamGetResponseSchemaFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type StreamGetResponseSchemaFormatJson struct {
	Type            StreamGetResponseSchemaFormatJsonType            `json:"type,required"`
	DecimalEncoding StreamGetResponseSchemaFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat StreamGetResponseSchemaFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                             `json:"unstructured"`
	JSON            streamGetResponseSchemaFormatJsonJSON            `json:"-"`
}

// streamGetResponseSchemaFormatJsonJSON contains the JSON metadata for the struct
// [StreamGetResponseSchemaFormatJson]
type streamGetResponseSchemaFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFormatJson) implementsStreamGetResponseSchemaFormat() {}

type StreamGetResponseSchemaFormatJsonType string

const (
	StreamGetResponseSchemaFormatJsonTypeJson StreamGetResponseSchemaFormatJsonType = "json"
)

func (r StreamGetResponseSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatJsonDecimalEncoding string

const (
	StreamGetResponseSchemaFormatJsonDecimalEncodingNumber StreamGetResponseSchemaFormatJsonDecimalEncoding = "number"
	StreamGetResponseSchemaFormatJsonDecimalEncodingString StreamGetResponseSchemaFormatJsonDecimalEncoding = "string"
	StreamGetResponseSchemaFormatJsonDecimalEncodingBytes  StreamGetResponseSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r StreamGetResponseSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatJsonDecimalEncodingNumber, StreamGetResponseSchemaFormatJsonDecimalEncodingString, StreamGetResponseSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatJsonTimestampFormat string

const (
	StreamGetResponseSchemaFormatJsonTimestampFormatRfc3339    StreamGetResponseSchemaFormatJsonTimestampFormat = "rfc3339"
	StreamGetResponseSchemaFormatJsonTimestampFormatUnixMillis StreamGetResponseSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamGetResponseSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatJsonTimestampFormatRfc3339, StreamGetResponseSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatParquet struct {
	Type          StreamGetResponseSchemaFormatParquetType        `json:"type,required"`
	Compression   StreamGetResponseSchemaFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                           `json:"row_group_bytes,nullable"`
	JSON          streamGetResponseSchemaFormatParquetJSON        `json:"-"`
}

// streamGetResponseSchemaFormatParquetJSON contains the JSON metadata for the
// struct [StreamGetResponseSchemaFormatParquet]
type streamGetResponseSchemaFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StreamGetResponseSchemaFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseSchemaFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r StreamGetResponseSchemaFormatParquet) implementsStreamGetResponseSchemaFormat() {}

type StreamGetResponseSchemaFormatParquetType string

const (
	StreamGetResponseSchemaFormatParquetTypeParquet StreamGetResponseSchemaFormatParquetType = "parquet"
)

func (r StreamGetResponseSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatParquetCompression string

const (
	StreamGetResponseSchemaFormatParquetCompressionUncompressed StreamGetResponseSchemaFormatParquetCompression = "uncompressed"
	StreamGetResponseSchemaFormatParquetCompressionSnappy       StreamGetResponseSchemaFormatParquetCompression = "snappy"
	StreamGetResponseSchemaFormatParquetCompressionGzip         StreamGetResponseSchemaFormatParquetCompression = "gzip"
	StreamGetResponseSchemaFormatParquetCompressionZstd         StreamGetResponseSchemaFormatParquetCompression = "zstd"
	StreamGetResponseSchemaFormatParquetCompressionLz4          StreamGetResponseSchemaFormatParquetCompression = "lz4"
)

func (r StreamGetResponseSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatParquetCompressionUncompressed, StreamGetResponseSchemaFormatParquetCompressionSnappy, StreamGetResponseSchemaFormatParquetCompressionGzip, StreamGetResponseSchemaFormatParquetCompressionZstd, StreamGetResponseSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatType string

const (
	StreamGetResponseSchemaFormatTypeJson    StreamGetResponseSchemaFormatType = "json"
	StreamGetResponseSchemaFormatTypeParquet StreamGetResponseSchemaFormatType = "parquet"
)

func (r StreamGetResponseSchemaFormatType) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatTypeJson, StreamGetResponseSchemaFormatTypeParquet:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatCompression string

const (
	StreamGetResponseSchemaFormatCompressionUncompressed StreamGetResponseSchemaFormatCompression = "uncompressed"
	StreamGetResponseSchemaFormatCompressionSnappy       StreamGetResponseSchemaFormatCompression = "snappy"
	StreamGetResponseSchemaFormatCompressionGzip         StreamGetResponseSchemaFormatCompression = "gzip"
	StreamGetResponseSchemaFormatCompressionZstd         StreamGetResponseSchemaFormatCompression = "zstd"
	StreamGetResponseSchemaFormatCompressionLz4          StreamGetResponseSchemaFormatCompression = "lz4"
)

func (r StreamGetResponseSchemaFormatCompression) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatCompressionUncompressed, StreamGetResponseSchemaFormatCompressionSnappy, StreamGetResponseSchemaFormatCompressionGzip, StreamGetResponseSchemaFormatCompressionZstd, StreamGetResponseSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatDecimalEncoding string

const (
	StreamGetResponseSchemaFormatDecimalEncodingNumber StreamGetResponseSchemaFormatDecimalEncoding = "number"
	StreamGetResponseSchemaFormatDecimalEncodingString StreamGetResponseSchemaFormatDecimalEncoding = "string"
	StreamGetResponseSchemaFormatDecimalEncodingBytes  StreamGetResponseSchemaFormatDecimalEncoding = "bytes"
)

func (r StreamGetResponseSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatDecimalEncodingNumber, StreamGetResponseSchemaFormatDecimalEncodingString, StreamGetResponseSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamGetResponseSchemaFormatTimestampFormat string

const (
	StreamGetResponseSchemaFormatTimestampFormatRfc3339    StreamGetResponseSchemaFormatTimestampFormat = "rfc3339"
	StreamGetResponseSchemaFormatTimestampFormatUnixMillis StreamGetResponseSchemaFormatTimestampFormat = "unix_millis"
)

func (r StreamGetResponseSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamGetResponseSchemaFormatTimestampFormatRfc3339, StreamGetResponseSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies the name of the Stream.
	Name          param.Field[string]                       `json:"name,required"`
	Format        param.Field[StreamNewParamsFormatUnion]   `json:"format"`
	HTTP          param.Field[StreamNewParamsHTTP]          `json:"http"`
	Schema        param.Field[StreamNewParamsSchema]        `json:"schema"`
	WorkerBinding param.Field[StreamNewParamsWorkerBinding] `json:"worker_binding"`
}

func (r StreamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewParamsFormat struct {
	Type            param.Field[StreamNewParamsFormatType]            `json:"type,required"`
	Compression     param.Field[StreamNewParamsFormatCompression]     `json:"compression"`
	DecimalEncoding param.Field[StreamNewParamsFormatDecimalEncoding] `json:"decimal_encoding"`
	RowGroupBytes   param.Field[int64]                                `json:"row_group_bytes"`
	TimestampFormat param.Field[StreamNewParamsFormatTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                 `json:"unstructured"`
}

func (r StreamNewParamsFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsFormat) implementsStreamNewParamsFormatUnion() {}

// Satisfied by [pipelines.StreamNewParamsFormatJson],
// [pipelines.StreamNewParamsFormatParquet], [StreamNewParamsFormat].
type StreamNewParamsFormatUnion interface {
	implementsStreamNewParamsFormatUnion()
}

type StreamNewParamsFormatJson struct {
	Type            param.Field[StreamNewParamsFormatJsonType]            `json:"type,required"`
	DecimalEncoding param.Field[StreamNewParamsFormatJsonDecimalEncoding] `json:"decimal_encoding"`
	TimestampFormat param.Field[StreamNewParamsFormatJsonTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                     `json:"unstructured"`
}

func (r StreamNewParamsFormatJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsFormatJson) implementsStreamNewParamsFormatUnion() {}

type StreamNewParamsFormatJsonType string

const (
	StreamNewParamsFormatJsonTypeJson StreamNewParamsFormatJsonType = "json"
)

func (r StreamNewParamsFormatJsonType) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamNewParamsFormatJsonDecimalEncoding string

const (
	StreamNewParamsFormatJsonDecimalEncodingNumber StreamNewParamsFormatJsonDecimalEncoding = "number"
	StreamNewParamsFormatJsonDecimalEncodingString StreamNewParamsFormatJsonDecimalEncoding = "string"
	StreamNewParamsFormatJsonDecimalEncodingBytes  StreamNewParamsFormatJsonDecimalEncoding = "bytes"
)

func (r StreamNewParamsFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatJsonDecimalEncodingNumber, StreamNewParamsFormatJsonDecimalEncodingString, StreamNewParamsFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewParamsFormatJsonTimestampFormat string

const (
	StreamNewParamsFormatJsonTimestampFormatRfc3339    StreamNewParamsFormatJsonTimestampFormat = "rfc3339"
	StreamNewParamsFormatJsonTimestampFormatUnixMillis StreamNewParamsFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamNewParamsFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatJsonTimestampFormatRfc3339, StreamNewParamsFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewParamsFormatParquet struct {
	Type          param.Field[StreamNewParamsFormatParquetType]        `json:"type,required"`
	Compression   param.Field[StreamNewParamsFormatParquetCompression] `json:"compression"`
	RowGroupBytes param.Field[int64]                                   `json:"row_group_bytes"`
}

func (r StreamNewParamsFormatParquet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsFormatParquet) implementsStreamNewParamsFormatUnion() {}

type StreamNewParamsFormatParquetType string

const (
	StreamNewParamsFormatParquetTypeParquet StreamNewParamsFormatParquetType = "parquet"
)

func (r StreamNewParamsFormatParquetType) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamNewParamsFormatParquetCompression string

const (
	StreamNewParamsFormatParquetCompressionUncompressed StreamNewParamsFormatParquetCompression = "uncompressed"
	StreamNewParamsFormatParquetCompressionSnappy       StreamNewParamsFormatParquetCompression = "snappy"
	StreamNewParamsFormatParquetCompressionGzip         StreamNewParamsFormatParquetCompression = "gzip"
	StreamNewParamsFormatParquetCompressionZstd         StreamNewParamsFormatParquetCompression = "zstd"
	StreamNewParamsFormatParquetCompressionLz4          StreamNewParamsFormatParquetCompression = "lz4"
)

func (r StreamNewParamsFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatParquetCompressionUncompressed, StreamNewParamsFormatParquetCompressionSnappy, StreamNewParamsFormatParquetCompressionGzip, StreamNewParamsFormatParquetCompressionZstd, StreamNewParamsFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamNewParamsFormatType string

const (
	StreamNewParamsFormatTypeJson    StreamNewParamsFormatType = "json"
	StreamNewParamsFormatTypeParquet StreamNewParamsFormatType = "parquet"
)

func (r StreamNewParamsFormatType) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatTypeJson, StreamNewParamsFormatTypeParquet:
		return true
	}
	return false
}

type StreamNewParamsFormatCompression string

const (
	StreamNewParamsFormatCompressionUncompressed StreamNewParamsFormatCompression = "uncompressed"
	StreamNewParamsFormatCompressionSnappy       StreamNewParamsFormatCompression = "snappy"
	StreamNewParamsFormatCompressionGzip         StreamNewParamsFormatCompression = "gzip"
	StreamNewParamsFormatCompressionZstd         StreamNewParamsFormatCompression = "zstd"
	StreamNewParamsFormatCompressionLz4          StreamNewParamsFormatCompression = "lz4"
)

func (r StreamNewParamsFormatCompression) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatCompressionUncompressed, StreamNewParamsFormatCompressionSnappy, StreamNewParamsFormatCompressionGzip, StreamNewParamsFormatCompressionZstd, StreamNewParamsFormatCompressionLz4:
		return true
	}
	return false
}

type StreamNewParamsFormatDecimalEncoding string

const (
	StreamNewParamsFormatDecimalEncodingNumber StreamNewParamsFormatDecimalEncoding = "number"
	StreamNewParamsFormatDecimalEncodingString StreamNewParamsFormatDecimalEncoding = "string"
	StreamNewParamsFormatDecimalEncodingBytes  StreamNewParamsFormatDecimalEncoding = "bytes"
)

func (r StreamNewParamsFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatDecimalEncodingNumber, StreamNewParamsFormatDecimalEncodingString, StreamNewParamsFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewParamsFormatTimestampFormat string

const (
	StreamNewParamsFormatTimestampFormatRfc3339    StreamNewParamsFormatTimestampFormat = "rfc3339"
	StreamNewParamsFormatTimestampFormatUnixMillis StreamNewParamsFormatTimestampFormat = "unix_millis"
)

func (r StreamNewParamsFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewParamsFormatTimestampFormatRfc3339, StreamNewParamsFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewParamsHTTP struct {
	// Indicates that authentication is required for the HTTP endpoint.
	Authentication param.Field[bool] `json:"authentication,required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Specifies the CORS options for the HTTP endpoint.
	CORS param.Field[StreamNewParamsHTTPCORS] `json:"cors"`
}

func (r StreamNewParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the CORS options for the HTTP endpoint.
type StreamNewParamsHTTPCORS struct {
	Origins param.Field[[]string] `json:"origins"`
}

func (r StreamNewParamsHTTPCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewParamsSchema struct {
	Fields   param.Field[[]StreamNewParamsSchemaFieldUnion] `json:"fields"`
	Format   param.Field[StreamNewParamsSchemaFormatUnion]  `json:"format"`
	Inferred param.Field[bool]                              `json:"inferred"`
}

func (r StreamNewParamsSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewParamsSchemaField struct {
	Type        param.Field[StreamNewParamsSchemaFieldsType] `json:"type,required"`
	MetadataKey param.Field[string]                          `json:"metadata_key"`
	Name        param.Field[string]                          `json:"name"`
	Required    param.Field[bool]                            `json:"required"`
	SqlName     param.Field[string]                          `json:"sql_name"`
	Unit        param.Field[StreamNewParamsSchemaFieldsUnit] `json:"unit"`
}

func (r StreamNewParamsSchemaField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaField) implementsStreamNewParamsSchemaFieldUnion() {}

// Satisfied by [pipelines.StreamNewParamsSchemaFieldsInt32],
// [pipelines.StreamNewParamsSchemaFieldsInt64],
// [pipelines.StreamNewParamsSchemaFieldsFloat32],
// [pipelines.StreamNewParamsSchemaFieldsFloat64],
// [pipelines.StreamNewParamsSchemaFieldsBool],
// [pipelines.StreamNewParamsSchemaFieldsString],
// [pipelines.StreamNewParamsSchemaFieldsBinary],
// [pipelines.StreamNewParamsSchemaFieldsTimestamp],
// [pipelines.StreamNewParamsSchemaFieldsJson],
// [pipelines.StreamNewParamsSchemaFieldsStruct],
// [pipelines.StreamNewParamsSchemaFieldsList], [StreamNewParamsSchemaField].
type StreamNewParamsSchemaFieldUnion interface {
	implementsStreamNewParamsSchemaFieldUnion()
}

type StreamNewParamsSchemaFieldsInt32 struct {
	Type        param.Field[StreamNewParamsSchemaFieldsInt32Type] `json:"type,required"`
	MetadataKey param.Field[string]                               `json:"metadata_key"`
	Name        param.Field[string]                               `json:"name"`
	Required    param.Field[bool]                                 `json:"required"`
	SqlName     param.Field[string]                               `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsInt32) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsInt32) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsInt32Type string

const (
	StreamNewParamsSchemaFieldsInt32TypeInt32 StreamNewParamsSchemaFieldsInt32Type = "int32"
)

func (r StreamNewParamsSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsInt64 struct {
	Type        param.Field[StreamNewParamsSchemaFieldsInt64Type] `json:"type,required"`
	MetadataKey param.Field[string]                               `json:"metadata_key"`
	Name        param.Field[string]                               `json:"name"`
	Required    param.Field[bool]                                 `json:"required"`
	SqlName     param.Field[string]                               `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsInt64) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsInt64) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsInt64Type string

const (
	StreamNewParamsSchemaFieldsInt64TypeInt64 StreamNewParamsSchemaFieldsInt64Type = "int64"
)

func (r StreamNewParamsSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsFloat32 struct {
	Type        param.Field[StreamNewParamsSchemaFieldsFloat32Type] `json:"type,required"`
	MetadataKey param.Field[string]                                 `json:"metadata_key"`
	Name        param.Field[string]                                 `json:"name"`
	Required    param.Field[bool]                                   `json:"required"`
	SqlName     param.Field[string]                                 `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsFloat32) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsFloat32) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsFloat32Type string

const (
	StreamNewParamsSchemaFieldsFloat32TypeFloat32 StreamNewParamsSchemaFieldsFloat32Type = "float32"
)

func (r StreamNewParamsSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsFloat64 struct {
	Type        param.Field[StreamNewParamsSchemaFieldsFloat64Type] `json:"type,required"`
	MetadataKey param.Field[string]                                 `json:"metadata_key"`
	Name        param.Field[string]                                 `json:"name"`
	Required    param.Field[bool]                                   `json:"required"`
	SqlName     param.Field[string]                                 `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsFloat64) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsFloat64) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsFloat64Type string

const (
	StreamNewParamsSchemaFieldsFloat64TypeFloat64 StreamNewParamsSchemaFieldsFloat64Type = "float64"
)

func (r StreamNewParamsSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsBool struct {
	Type        param.Field[StreamNewParamsSchemaFieldsBoolType] `json:"type,required"`
	MetadataKey param.Field[string]                              `json:"metadata_key"`
	Name        param.Field[string]                              `json:"name"`
	Required    param.Field[bool]                                `json:"required"`
	SqlName     param.Field[string]                              `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsBool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsBool) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsBoolType string

const (
	StreamNewParamsSchemaFieldsBoolTypeBool StreamNewParamsSchemaFieldsBoolType = "bool"
)

func (r StreamNewParamsSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsString struct {
	Type        param.Field[StreamNewParamsSchemaFieldsStringType] `json:"type,required"`
	MetadataKey param.Field[string]                                `json:"metadata_key"`
	Name        param.Field[string]                                `json:"name"`
	Required    param.Field[bool]                                  `json:"required"`
	SqlName     param.Field[string]                                `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsString) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsStringType string

const (
	StreamNewParamsSchemaFieldsStringTypeString StreamNewParamsSchemaFieldsStringType = "string"
)

func (r StreamNewParamsSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsBinary struct {
	Type        param.Field[StreamNewParamsSchemaFieldsBinaryType] `json:"type,required"`
	MetadataKey param.Field[string]                                `json:"metadata_key"`
	Name        param.Field[string]                                `json:"name"`
	Required    param.Field[bool]                                  `json:"required"`
	SqlName     param.Field[string]                                `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsBinary) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsBinary) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsBinaryType string

const (
	StreamNewParamsSchemaFieldsBinaryTypeBinary StreamNewParamsSchemaFieldsBinaryType = "binary"
)

func (r StreamNewParamsSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsTimestamp struct {
	Type        param.Field[StreamNewParamsSchemaFieldsTimestampType] `json:"type,required"`
	MetadataKey param.Field[string]                                   `json:"metadata_key"`
	Name        param.Field[string]                                   `json:"name"`
	Required    param.Field[bool]                                     `json:"required"`
	SqlName     param.Field[string]                                   `json:"sql_name"`
	Unit        param.Field[StreamNewParamsSchemaFieldsTimestampUnit] `json:"unit"`
}

func (r StreamNewParamsSchemaFieldsTimestamp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsTimestamp) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsTimestampType string

const (
	StreamNewParamsSchemaFieldsTimestampTypeTimestamp StreamNewParamsSchemaFieldsTimestampType = "timestamp"
)

func (r StreamNewParamsSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsTimestampUnit string

const (
	StreamNewParamsSchemaFieldsTimestampUnitSecond      StreamNewParamsSchemaFieldsTimestampUnit = "second"
	StreamNewParamsSchemaFieldsTimestampUnitMillisecond StreamNewParamsSchemaFieldsTimestampUnit = "millisecond"
	StreamNewParamsSchemaFieldsTimestampUnitMicrosecond StreamNewParamsSchemaFieldsTimestampUnit = "microsecond"
	StreamNewParamsSchemaFieldsTimestampUnitNanosecond  StreamNewParamsSchemaFieldsTimestampUnit = "nanosecond"
)

func (r StreamNewParamsSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsTimestampUnitSecond, StreamNewParamsSchemaFieldsTimestampUnitMillisecond, StreamNewParamsSchemaFieldsTimestampUnitMicrosecond, StreamNewParamsSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsJson struct {
	Type        param.Field[StreamNewParamsSchemaFieldsJsonType] `json:"type,required"`
	MetadataKey param.Field[string]                              `json:"metadata_key"`
	Name        param.Field[string]                              `json:"name"`
	Required    param.Field[bool]                                `json:"required"`
	SqlName     param.Field[string]                              `json:"sql_name"`
}

func (r StreamNewParamsSchemaFieldsJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsJson) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsJsonType string

const (
	StreamNewParamsSchemaFieldsJsonTypeJson StreamNewParamsSchemaFieldsJsonType = "json"
)

func (r StreamNewParamsSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsStruct struct {
}

func (r StreamNewParamsSchemaFieldsStruct) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsStruct) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsList struct {
}

func (r StreamNewParamsSchemaFieldsList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFieldsList) implementsStreamNewParamsSchemaFieldUnion() {}

type StreamNewParamsSchemaFieldsType string

const (
	StreamNewParamsSchemaFieldsTypeInt32     StreamNewParamsSchemaFieldsType = "int32"
	StreamNewParamsSchemaFieldsTypeInt64     StreamNewParamsSchemaFieldsType = "int64"
	StreamNewParamsSchemaFieldsTypeFloat32   StreamNewParamsSchemaFieldsType = "float32"
	StreamNewParamsSchemaFieldsTypeFloat64   StreamNewParamsSchemaFieldsType = "float64"
	StreamNewParamsSchemaFieldsTypeBool      StreamNewParamsSchemaFieldsType = "bool"
	StreamNewParamsSchemaFieldsTypeString    StreamNewParamsSchemaFieldsType = "string"
	StreamNewParamsSchemaFieldsTypeBinary    StreamNewParamsSchemaFieldsType = "binary"
	StreamNewParamsSchemaFieldsTypeTimestamp StreamNewParamsSchemaFieldsType = "timestamp"
	StreamNewParamsSchemaFieldsTypeJson      StreamNewParamsSchemaFieldsType = "json"
	StreamNewParamsSchemaFieldsTypeStruct    StreamNewParamsSchemaFieldsType = "struct"
	StreamNewParamsSchemaFieldsTypeList      StreamNewParamsSchemaFieldsType = "list"
)

func (r StreamNewParamsSchemaFieldsType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsTypeInt32, StreamNewParamsSchemaFieldsTypeInt64, StreamNewParamsSchemaFieldsTypeFloat32, StreamNewParamsSchemaFieldsTypeFloat64, StreamNewParamsSchemaFieldsTypeBool, StreamNewParamsSchemaFieldsTypeString, StreamNewParamsSchemaFieldsTypeBinary, StreamNewParamsSchemaFieldsTypeTimestamp, StreamNewParamsSchemaFieldsTypeJson, StreamNewParamsSchemaFieldsTypeStruct, StreamNewParamsSchemaFieldsTypeList:
		return true
	}
	return false
}

type StreamNewParamsSchemaFieldsUnit string

const (
	StreamNewParamsSchemaFieldsUnitSecond      StreamNewParamsSchemaFieldsUnit = "second"
	StreamNewParamsSchemaFieldsUnitMillisecond StreamNewParamsSchemaFieldsUnit = "millisecond"
	StreamNewParamsSchemaFieldsUnitMicrosecond StreamNewParamsSchemaFieldsUnit = "microsecond"
	StreamNewParamsSchemaFieldsUnitNanosecond  StreamNewParamsSchemaFieldsUnit = "nanosecond"
)

func (r StreamNewParamsSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFieldsUnitSecond, StreamNewParamsSchemaFieldsUnitMillisecond, StreamNewParamsSchemaFieldsUnitMicrosecond, StreamNewParamsSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormat struct {
	Type            param.Field[StreamNewParamsSchemaFormatType]            `json:"type,required"`
	Compression     param.Field[StreamNewParamsSchemaFormatCompression]     `json:"compression"`
	DecimalEncoding param.Field[StreamNewParamsSchemaFormatDecimalEncoding] `json:"decimal_encoding"`
	RowGroupBytes   param.Field[int64]                                      `json:"row_group_bytes"`
	TimestampFormat param.Field[StreamNewParamsSchemaFormatTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                       `json:"unstructured"`
}

func (r StreamNewParamsSchemaFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFormat) implementsStreamNewParamsSchemaFormatUnion() {}

// Satisfied by [pipelines.StreamNewParamsSchemaFormatJson],
// [pipelines.StreamNewParamsSchemaFormatParquet], [StreamNewParamsSchemaFormat].
type StreamNewParamsSchemaFormatUnion interface {
	implementsStreamNewParamsSchemaFormatUnion()
}

type StreamNewParamsSchemaFormatJson struct {
	Type            param.Field[StreamNewParamsSchemaFormatJsonType]            `json:"type,required"`
	DecimalEncoding param.Field[StreamNewParamsSchemaFormatJsonDecimalEncoding] `json:"decimal_encoding"`
	TimestampFormat param.Field[StreamNewParamsSchemaFormatJsonTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                           `json:"unstructured"`
}

func (r StreamNewParamsSchemaFormatJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFormatJson) implementsStreamNewParamsSchemaFormatUnion() {}

type StreamNewParamsSchemaFormatJsonType string

const (
	StreamNewParamsSchemaFormatJsonTypeJson StreamNewParamsSchemaFormatJsonType = "json"
)

func (r StreamNewParamsSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatJsonDecimalEncoding string

const (
	StreamNewParamsSchemaFormatJsonDecimalEncodingNumber StreamNewParamsSchemaFormatJsonDecimalEncoding = "number"
	StreamNewParamsSchemaFormatJsonDecimalEncodingString StreamNewParamsSchemaFormatJsonDecimalEncoding = "string"
	StreamNewParamsSchemaFormatJsonDecimalEncodingBytes  StreamNewParamsSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r StreamNewParamsSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatJsonDecimalEncodingNumber, StreamNewParamsSchemaFormatJsonDecimalEncodingString, StreamNewParamsSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatJsonTimestampFormat string

const (
	StreamNewParamsSchemaFormatJsonTimestampFormatRfc3339    StreamNewParamsSchemaFormatJsonTimestampFormat = "rfc3339"
	StreamNewParamsSchemaFormatJsonTimestampFormatUnixMillis StreamNewParamsSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r StreamNewParamsSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatJsonTimestampFormatRfc3339, StreamNewParamsSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatParquet struct {
	Type          param.Field[StreamNewParamsSchemaFormatParquetType]        `json:"type,required"`
	Compression   param.Field[StreamNewParamsSchemaFormatParquetCompression] `json:"compression"`
	RowGroupBytes param.Field[int64]                                         `json:"row_group_bytes"`
}

func (r StreamNewParamsSchemaFormatParquet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r StreamNewParamsSchemaFormatParquet) implementsStreamNewParamsSchemaFormatUnion() {}

type StreamNewParamsSchemaFormatParquetType string

const (
	StreamNewParamsSchemaFormatParquetTypeParquet StreamNewParamsSchemaFormatParquetType = "parquet"
)

func (r StreamNewParamsSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatParquetCompression string

const (
	StreamNewParamsSchemaFormatParquetCompressionUncompressed StreamNewParamsSchemaFormatParquetCompression = "uncompressed"
	StreamNewParamsSchemaFormatParquetCompressionSnappy       StreamNewParamsSchemaFormatParquetCompression = "snappy"
	StreamNewParamsSchemaFormatParquetCompressionGzip         StreamNewParamsSchemaFormatParquetCompression = "gzip"
	StreamNewParamsSchemaFormatParquetCompressionZstd         StreamNewParamsSchemaFormatParquetCompression = "zstd"
	StreamNewParamsSchemaFormatParquetCompressionLz4          StreamNewParamsSchemaFormatParquetCompression = "lz4"
)

func (r StreamNewParamsSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatParquetCompressionUncompressed, StreamNewParamsSchemaFormatParquetCompressionSnappy, StreamNewParamsSchemaFormatParquetCompressionGzip, StreamNewParamsSchemaFormatParquetCompressionZstd, StreamNewParamsSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatType string

const (
	StreamNewParamsSchemaFormatTypeJson    StreamNewParamsSchemaFormatType = "json"
	StreamNewParamsSchemaFormatTypeParquet StreamNewParamsSchemaFormatType = "parquet"
)

func (r StreamNewParamsSchemaFormatType) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatTypeJson, StreamNewParamsSchemaFormatTypeParquet:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatCompression string

const (
	StreamNewParamsSchemaFormatCompressionUncompressed StreamNewParamsSchemaFormatCompression = "uncompressed"
	StreamNewParamsSchemaFormatCompressionSnappy       StreamNewParamsSchemaFormatCompression = "snappy"
	StreamNewParamsSchemaFormatCompressionGzip         StreamNewParamsSchemaFormatCompression = "gzip"
	StreamNewParamsSchemaFormatCompressionZstd         StreamNewParamsSchemaFormatCompression = "zstd"
	StreamNewParamsSchemaFormatCompressionLz4          StreamNewParamsSchemaFormatCompression = "lz4"
)

func (r StreamNewParamsSchemaFormatCompression) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatCompressionUncompressed, StreamNewParamsSchemaFormatCompressionSnappy, StreamNewParamsSchemaFormatCompressionGzip, StreamNewParamsSchemaFormatCompressionZstd, StreamNewParamsSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatDecimalEncoding string

const (
	StreamNewParamsSchemaFormatDecimalEncodingNumber StreamNewParamsSchemaFormatDecimalEncoding = "number"
	StreamNewParamsSchemaFormatDecimalEncodingString StreamNewParamsSchemaFormatDecimalEncoding = "string"
	StreamNewParamsSchemaFormatDecimalEncodingBytes  StreamNewParamsSchemaFormatDecimalEncoding = "bytes"
)

func (r StreamNewParamsSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatDecimalEncodingNumber, StreamNewParamsSchemaFormatDecimalEncodingString, StreamNewParamsSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type StreamNewParamsSchemaFormatTimestampFormat string

const (
	StreamNewParamsSchemaFormatTimestampFormatRfc3339    StreamNewParamsSchemaFormatTimestampFormat = "rfc3339"
	StreamNewParamsSchemaFormatTimestampFormatUnixMillis StreamNewParamsSchemaFormatTimestampFormat = "unix_millis"
)

func (r StreamNewParamsSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case StreamNewParamsSchemaFormatTimestampFormatRfc3339, StreamNewParamsSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type StreamNewParamsWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r StreamNewParamsWorkerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewResponseEnvelope struct {
	Result StreamNewResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                          `json:"success,required"`
	JSON    streamNewResponseEnvelopeJSON `json:"-"`
}

// streamNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamNewResponseEnvelope]
type streamNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamUpdateParams struct {
	// Specifies the public ID of the account.
	AccountID     param.Field[string]                          `path:"account_id,required"`
	HTTP          param.Field[StreamUpdateParamsHTTP]          `json:"http"`
	WorkerBinding param.Field[StreamUpdateParamsWorkerBinding] `json:"worker_binding"`
}

func (r StreamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamUpdateParamsHTTP struct {
	// Indicates that authentication is required for the HTTP endpoint.
	Authentication param.Field[bool] `json:"authentication,required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Specifies the CORS options for the HTTP endpoint.
	CORS param.Field[StreamUpdateParamsHTTPCORS] `json:"cors"`
}

func (r StreamUpdateParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the CORS options for the HTTP endpoint.
type StreamUpdateParamsHTTPCORS struct {
	Origins param.Field[[]string] `json:"origins"`
}

func (r StreamUpdateParamsHTTPCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamUpdateParamsWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r StreamUpdateParamsWorkerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamUpdateResponseEnvelope struct {
	Result StreamUpdateResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                             `json:"success,required"`
	JSON    streamUpdateResponseEnvelopeJSON `json:"-"`
}

// streamUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamUpdateResponseEnvelope]
type streamUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamListParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string]  `path:"account_id,required"`
	Page      param.Field[float64] `query:"page"`
	PerPage   param.Field[float64] `query:"per_page"`
	// Specifies the public ID of the pipeline.
	PipelineID param.Field[string] `query:"pipeline_id"`
}

// URLQuery serializes [StreamListParams]'s query parameters as `url.Values`.
func (r StreamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StreamDeleteParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Delete stream forcefully, including deleting any dependent pipelines.
	Force param.Field[string] `query:"force"`
}

// URLQuery serializes [StreamDeleteParams]'s query parameters as `url.Values`.
func (r StreamDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type StreamGetParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamGetResponseEnvelope struct {
	Result StreamGetResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                          `json:"success,required"`
	JSON    streamGetResponseEnvelopeJSON `json:"-"`
}

// streamGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamGetResponseEnvelope]
type streamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
