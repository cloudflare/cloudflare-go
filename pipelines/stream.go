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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update a Stream.
func (r *StreamService) Update(ctx context.Context, streamID string, params StreamUpdateParams, opts ...option.RequestOption) (res *StreamUpdateResponse, err error) {
	var env StreamUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams/%s", params.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List/Filter Streams in Account.
func (r *StreamService) List(ctx context.Context, params StreamListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[StreamListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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
func (r *StreamService) Delete(ctx context.Context, streamID string, body StreamDeleteParams, opts ...option.RequestOption) (res *StreamDeleteResponse, err error) {
	var env StreamDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams/%s", body.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get Stream Details.
func (r *StreamService) Get(ctx context.Context, streamID string, query StreamGetParams, opts ...option.RequestOption) (res *StreamGetResponse, err error) {
	var env StreamGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if streamID == "" {
		err = errors.New("missing required stream_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/streams/%s", query.AccountID, streamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type StreamNewResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                `json:"id" api:"required"`
	CreatedAt  time.Time             `json:"created_at" api:"required" format:"date-time"`
	HTTP       StreamNewResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time             `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name" api:"required"`
	// Indicates the current version of this stream.
	Version       int64                          `json:"version" api:"required"`
	WorkerBinding StreamNewResponseWorkerBinding `json:"worker_binding" api:"required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string `json:"endpoint" format:"uri"`
	// Defines the data format of the events.
	Format StreamNewResponseFormat `json:"format"`
	// Defines the schema of the events in the data stream.
	Schema StreamNewResponseSchema `json:"schema"`
	JSON   streamNewResponseJSON   `json:"-"`
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
	Authentication bool `json:"authentication" api:"required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled" api:"required"`
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
	Enabled bool                               `json:"enabled" api:"required"`
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

// Defines the data format of the events.
type StreamNewResponseFormat struct {
	Type            StreamNewResponseFormatType            `json:"type" api:"required"`
	Compression     StreamNewResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamNewResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                  `json:"row_group_bytes" api:"nullable"`
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

// Defines the data format of the events.
//
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
	Type            StreamNewResponseFormatJsonType            `json:"type" api:"required"`
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
	Type          StreamNewResponseFormatParquetType        `json:"type" api:"required"`
	Compression   StreamNewResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                     `json:"row_group_bytes" api:"nullable"`
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

// Defines the schema of the events in the data stream.
type StreamNewResponseSchema struct {
	Fields   []SourceField               `json:"fields"`
	Inferred bool                        `json:"inferred" api:"nullable"`
	JSON     streamNewResponseSchemaJSON `json:"-"`
}

// streamNewResponseSchemaJSON contains the JSON metadata for the struct
// [StreamNewResponseSchema]
type streamNewResponseSchemaJSON struct {
	Fields      apijson.Field
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

type StreamUpdateResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                   `json:"id" api:"required"`
	CreatedAt  time.Time                `json:"created_at" api:"required" format:"date-time"`
	HTTP       StreamUpdateResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time                `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name" api:"required"`
	// Indicates the current version of this stream.
	Version       int64                             `json:"version" api:"required"`
	WorkerBinding StreamUpdateResponseWorkerBinding `json:"worker_binding" api:"required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string `json:"endpoint" format:"uri"`
	// Defines the data format of the events.
	Format StreamUpdateResponseFormat `json:"format"`
	// Defines the schema of the events in the data stream.
	Schema StreamUpdateResponseSchema `json:"schema"`
	JSON   streamUpdateResponseJSON   `json:"-"`
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
	Schema        apijson.Field
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
	Authentication bool `json:"authentication" api:"required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled" api:"required"`
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
	Enabled bool                                  `json:"enabled" api:"required"`
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

// Defines the data format of the events.
type StreamUpdateResponseFormat struct {
	Type            StreamUpdateResponseFormatType            `json:"type" api:"required"`
	Compression     StreamUpdateResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamUpdateResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                     `json:"row_group_bytes" api:"nullable"`
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

// Defines the data format of the events.
//
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
	Type            StreamUpdateResponseFormatJsonType            `json:"type" api:"required"`
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
	Type          StreamUpdateResponseFormatParquetType        `json:"type" api:"required"`
	Compression   StreamUpdateResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                        `json:"row_group_bytes" api:"nullable"`
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

// Defines the schema of the events in the data stream.
type StreamUpdateResponseSchema struct {
	Fields   []SourceField                  `json:"fields"`
	Inferred bool                           `json:"inferred" api:"nullable"`
	JSON     streamUpdateResponseSchemaJSON `json:"-"`
}

// streamUpdateResponseSchemaJSON contains the JSON metadata for the struct
// [StreamUpdateResponseSchema]
type streamUpdateResponseSchemaJSON struct {
	Fields      apijson.Field
	Inferred    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamUpdateResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamUpdateResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type StreamListResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                 `json:"id" api:"required"`
	CreatedAt  time.Time              `json:"created_at" api:"required" format:"date-time"`
	HTTP       StreamListResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time              `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name" api:"required"`
	// Indicates the current version of this stream.
	Version       int64                           `json:"version" api:"required"`
	WorkerBinding StreamListResponseWorkerBinding `json:"worker_binding" api:"required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string `json:"endpoint" format:"uri"`
	// Defines the data format of the events.
	Format StreamListResponseFormat `json:"format"`
	// Defines the schema of the events in the data stream.
	Schema StreamListResponseSchema `json:"schema"`
	JSON   streamListResponseJSON   `json:"-"`
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
	Authentication bool `json:"authentication" api:"required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled" api:"required"`
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
	Enabled bool                                `json:"enabled" api:"required"`
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

// Defines the data format of the events.
type StreamListResponseFormat struct {
	Type            StreamListResponseFormatType            `json:"type" api:"required"`
	Compression     StreamListResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamListResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                   `json:"row_group_bytes" api:"nullable"`
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

// Defines the data format of the events.
//
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
	Type            StreamListResponseFormatJsonType            `json:"type" api:"required"`
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
	Type          StreamListResponseFormatParquetType        `json:"type" api:"required"`
	Compression   StreamListResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                      `json:"row_group_bytes" api:"nullable"`
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

// Defines the schema of the events in the data stream.
type StreamListResponseSchema struct {
	Fields   []SourceField                `json:"fields"`
	Inferred bool                         `json:"inferred" api:"nullable"`
	JSON     streamListResponseSchemaJSON `json:"-"`
}

// streamListResponseSchemaJSON contains the JSON metadata for the struct
// [StreamListResponseSchema]
type streamListResponseSchemaJSON struct {
	Fields      apijson.Field
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

type StreamDeleteResponse = interface{}

type StreamGetResponse struct {
	// Indicates a unique identifier for this stream.
	ID         string                `json:"id" api:"required"`
	CreatedAt  time.Time             `json:"created_at" api:"required" format:"date-time"`
	HTTP       StreamGetResponseHTTP `json:"http" api:"required"`
	ModifiedAt time.Time             `json:"modified_at" api:"required" format:"date-time"`
	// Indicates the name of the Stream.
	Name string `json:"name" api:"required"`
	// Indicates the current version of this stream.
	Version       int64                          `json:"version" api:"required"`
	WorkerBinding StreamGetResponseWorkerBinding `json:"worker_binding" api:"required"`
	// Indicates the endpoint URL of this stream.
	Endpoint string `json:"endpoint" format:"uri"`
	// Defines the data format of the events.
	Format StreamGetResponseFormat `json:"format"`
	// Defines the schema of the events in the data stream.
	Schema StreamGetResponseSchema `json:"schema"`
	JSON   streamGetResponseJSON   `json:"-"`
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
	Authentication bool `json:"authentication" api:"required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled bool `json:"enabled" api:"required"`
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
	Enabled bool                               `json:"enabled" api:"required"`
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

// Defines the data format of the events.
type StreamGetResponseFormat struct {
	Type            StreamGetResponseFormatType            `json:"type" api:"required"`
	Compression     StreamGetResponseFormatCompression     `json:"compression"`
	DecimalEncoding StreamGetResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                  `json:"row_group_bytes" api:"nullable"`
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

// Defines the data format of the events.
//
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
	Type            StreamGetResponseFormatJsonType            `json:"type" api:"required"`
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
	Type          StreamGetResponseFormatParquetType        `json:"type" api:"required"`
	Compression   StreamGetResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                     `json:"row_group_bytes" api:"nullable"`
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

// Defines the schema of the events in the data stream.
type StreamGetResponseSchema struct {
	Fields   []SourceField               `json:"fields"`
	Inferred bool                        `json:"inferred" api:"nullable"`
	JSON     streamGetResponseSchemaJSON `json:"-"`
}

// streamGetResponseSchemaJSON contains the JSON metadata for the struct
// [StreamGetResponseSchema]
type streamGetResponseSchemaJSON struct {
	Fields      apijson.Field
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

type StreamNewParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Specifies the name of the Stream.
	Name param.Field[string] `json:"name" api:"required"`
	// Defines the data format of the events.
	Format param.Field[StreamNewParamsFormatUnion] `json:"format"`
	HTTP   param.Field[StreamNewParamsHTTP]        `json:"http"`
	// Defines the schema of the events in the data stream.
	Schema        param.Field[StreamNewParamsSchema]        `json:"schema"`
	WorkerBinding param.Field[StreamNewParamsWorkerBinding] `json:"worker_binding"`
}

func (r StreamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the data format of the events.
type StreamNewParamsFormat struct {
	Type            param.Field[StreamNewParamsFormatType]            `json:"type" api:"required"`
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

// Defines the data format of the events.
//
// Satisfied by [pipelines.StreamNewParamsFormatJson],
// [pipelines.StreamNewParamsFormatParquet], [StreamNewParamsFormat].
type StreamNewParamsFormatUnion interface {
	implementsStreamNewParamsFormatUnion()
}

type StreamNewParamsFormatJson struct {
	Type            param.Field[StreamNewParamsFormatJsonType]            `json:"type" api:"required"`
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
	Type          param.Field[StreamNewParamsFormatParquetType]        `json:"type" api:"required"`
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
	Authentication param.Field[bool] `json:"authentication" api:"required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
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

// Defines the schema of the events in the data stream.
type StreamNewParamsSchema struct {
	Fields   param.Field[[]SourceFieldUnionParam] `json:"fields"`
	Inferred param.Field[bool]                    `json:"inferred"`
}

func (r StreamNewParamsSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewParamsWorkerBinding struct {
	// Indicates that the worker binding is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r StreamNewParamsWorkerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamNewResponseEnvelope struct {
	Result StreamNewResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                          `json:"success" api:"required"`
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
	AccountID     param.Field[string]                          `path:"account_id" api:"required"`
	HTTP          param.Field[StreamUpdateParamsHTTP]          `json:"http"`
	WorkerBinding param.Field[StreamUpdateParamsWorkerBinding] `json:"worker_binding"`
}

func (r StreamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamUpdateParamsHTTP struct {
	// Indicates that authentication is required for the HTTP endpoint.
	Authentication param.Field[bool] `json:"authentication" api:"required"`
	// Indicates that the HTTP endpoint is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
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
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r StreamUpdateParamsWorkerBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamUpdateResponseEnvelope struct {
	Result StreamUpdateResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                             `json:"success" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters streams by name (case-insensitive substring).
	Name    param.Field[string]  `query:"name"`
	Page    param.Field[float64] `query:"page"`
	PerPage param.Field[float64] `query:"per_page"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type StreamDeleteResponseEnvelope struct {
	Result StreamDeleteResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                             `json:"success" api:"required"`
	JSON    streamDeleteResponseEnvelopeJSON `json:"-"`
}

// streamDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamDeleteResponseEnvelope]
type streamDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamGetParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type StreamGetResponseEnvelope struct {
	Result StreamGetResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                          `json:"success" api:"required"`
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
