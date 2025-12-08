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

// SinkService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSinkService] method instead.
type SinkService struct {
	Options []option.RequestOption
}

// NewSinkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSinkService(opts ...option.RequestOption) (r *SinkService) {
	r = &SinkService{}
	r.Options = opts
	return
}

// Create a new Sink.
func (r *SinkService) New(ctx context.Context, params SinkNewParams, opts ...option.RequestOption) (res *SinkNewResponse, err error) {
	var env SinkNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/sinks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List/Filter Sinks in Account.
func (r *SinkService) List(ctx context.Context, params SinkListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SinkListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/sinks", params.AccountID)
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

// List/Filter Sinks in Account.
func (r *SinkService) ListAutoPaging(ctx context.Context, params SinkListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SinkListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete Pipeline in Account.
func (r *SinkService) Delete(ctx context.Context, sinkID string, params SinkDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if sinkID == "" {
		err = errors.New("missing required sink_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/sinks/%s", params.AccountID, sinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

// Get Sink Details.
func (r *SinkService) Get(ctx context.Context, sinkID string, query SinkGetParams, opts ...option.RequestOption) (res *SinkGetResponse, err error) {
	var env SinkGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if sinkID == "" {
		err = errors.New("missing required sink_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/sinks/%s", query.AccountID, sinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SinkNewResponse struct {
	// Indicates a unique identifier for this sink.
	ID         string    `json:"id,required"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Defines the name of the Sink.
	Name string `json:"name,required"`
	// Specifies the type of sink.
	Type SinkNewResponseType `json:"type,required"`
	// R2 Data Catalog Sink
	Config SinkNewResponseConfig `json:"config"`
	Format SinkNewResponseFormat `json:"format"`
	Schema SinkNewResponseSchema `json:"schema"`
	JSON   sinkNewResponseJSON   `json:"-"`
}

// sinkNewResponseJSON contains the JSON metadata for the struct [SinkNewResponse]
type sinkNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Config      apijson.Field
	Format      apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of sink.
type SinkNewResponseType string

const (
	SinkNewResponseTypeR2            SinkNewResponseType = "r2"
	SinkNewResponseTypeR2DataCatalog SinkNewResponseType = "r2_data_catalog"
)

func (r SinkNewResponseType) IsKnown() bool {
	switch r {
	case SinkNewResponseTypeR2, SinkNewResponseTypeR2DataCatalog:
		return true
	}
	return false
}

// R2 Data Catalog Sink
type SinkNewResponseConfig struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket,required"`
	// Authentication token
	Token string `json:"token" format:"var-str"`
	// This field can have the runtime type of
	// [SinkNewResponseConfigCloudflarePipelinesR2TableCredentials].
	Credentials interface{} `json:"credentials"`
	// This field can have the runtime type of
	// [SinkNewResponseConfigCloudflarePipelinesR2TableFileNaming].
	FileNaming interface{} `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Table namespace
	Namespace string `json:"namespace"`
	// This field can have the runtime type of
	// [SinkNewResponseConfigCloudflarePipelinesR2TablePartitioning].
	Partitioning interface{} `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// This field can have the runtime type of
	// [SinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicy],
	// [SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy].
	RollingPolicy interface{} `json:"rolling_policy"`
	// Table name
	TableName string                    `json:"table_name"`
	JSON      sinkNewResponseConfigJSON `json:"-"`
	union     SinkNewResponseConfigUnion
}

// sinkNewResponseConfigJSON contains the JSON metadata for the struct
// [SinkNewResponseConfig]
type sinkNewResponseConfigJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	Token         apijson.Field
	Credentials   apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Namespace     apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	TableName     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r sinkNewResponseConfigJSON) RawJSON() string {
	return r.raw
}

func (r *SinkNewResponseConfig) UnmarshalJSON(data []byte) (err error) {
	*r = SinkNewResponseConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkNewResponseConfigUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [SinkNewResponseConfigCloudflarePipelinesR2Table],
// [SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTable].
func (r SinkNewResponseConfig) AsUnion() SinkNewResponseConfigUnion {
	return r.union
}

// R2 Data Catalog Sink
//
// Union satisfied by [SinkNewResponseConfigCloudflarePipelinesR2Table] or
// [SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTable].
type SinkNewResponseConfigUnion interface {
	implementsSinkNewResponseConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkNewResponseConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkNewResponseConfigCloudflarePipelinesR2Table{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTable{}),
		},
	)
}

type SinkNewResponseConfigCloudflarePipelinesR2Table struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket      string                                                     `json:"bucket,required"`
	Credentials SinkNewResponseConfigCloudflarePipelinesR2TableCredentials `json:"credentials,required"`
	// Controls filename prefix/suffix and strategy.
	FileNaming SinkNewResponseConfigCloudflarePipelinesR2TableFileNaming `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Data-layout partitioning for sinks.
	Partitioning SinkNewResponseConfigCloudflarePipelinesR2TablePartitioning `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicy `json:"rolling_policy"`
	JSON          sinkNewResponseConfigCloudflarePipelinesR2TableJSON          `json:"-"`
}

// sinkNewResponseConfigCloudflarePipelinesR2TableJSON contains the JSON metadata
// for the struct [SinkNewResponseConfigCloudflarePipelinesR2Table]
type sinkNewResponseConfigCloudflarePipelinesR2TableJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	Credentials   apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkNewResponseConfigCloudflarePipelinesR2Table) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseConfigCloudflarePipelinesR2TableJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseConfigCloudflarePipelinesR2Table) implementsSinkNewResponseConfig() {}

type SinkNewResponseConfigCloudflarePipelinesR2TableCredentials struct {
	// Cloudflare Account ID for the bucket
	AccessKeyID string `json:"access_key_id,required" format:"var-str"`
	// Cloudflare Account ID for the bucket
	SecretAccessKey string                                                         `json:"secret_access_key,required" format:"var-str"`
	JSON            sinkNewResponseConfigCloudflarePipelinesR2TableCredentialsJSON `json:"-"`
}

// sinkNewResponseConfigCloudflarePipelinesR2TableCredentialsJSON contains the JSON
// metadata for the struct
// [SinkNewResponseConfigCloudflarePipelinesR2TableCredentials]
type sinkNewResponseConfigCloudflarePipelinesR2TableCredentialsJSON struct {
	AccessKeyID     apijson.Field
	SecretAccessKey apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkNewResponseConfigCloudflarePipelinesR2TableCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseConfigCloudflarePipelinesR2TableCredentialsJSON) RawJSON() string {
	return r.raw
}

// Controls filename prefix/suffix and strategy.
type SinkNewResponseConfigCloudflarePipelinesR2TableFileNaming struct {
	// The prefix to use in file name. i.e prefix-<uuid>.parquet
	Prefix string `json:"prefix"`
	// Filename generation strategy.
	Strategy SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategy `json:"strategy"`
	// This will overwrite the default file suffix. i.e .parquet, use with caution
	Suffix string                                                        `json:"suffix"`
	JSON   sinkNewResponseConfigCloudflarePipelinesR2TableFileNamingJSON `json:"-"`
}

// sinkNewResponseConfigCloudflarePipelinesR2TableFileNamingJSON contains the JSON
// metadata for the struct
// [SinkNewResponseConfigCloudflarePipelinesR2TableFileNaming]
type sinkNewResponseConfigCloudflarePipelinesR2TableFileNamingJSON struct {
	Prefix      apijson.Field
	Strategy    apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseConfigCloudflarePipelinesR2TableFileNaming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseConfigCloudflarePipelinesR2TableFileNamingJSON) RawJSON() string {
	return r.raw
}

// Filename generation strategy.
type SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategy string

const (
	SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategySerial SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "serial"
	SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUID   SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid"
	SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7 SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid_v7"
	SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUlid   SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "ulid"
)

func (r SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategy) IsKnown() bool {
	switch r {
	case SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategySerial, SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUID, SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7, SinkNewResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUlid:
		return true
	}
	return false
}

// Data-layout partitioning for sinks.
type SinkNewResponseConfigCloudflarePipelinesR2TablePartitioning struct {
	// The pattern of the date string
	TimePattern string                                                          `json:"time_pattern"`
	JSON        sinkNewResponseConfigCloudflarePipelinesR2TablePartitioningJSON `json:"-"`
}

// sinkNewResponseConfigCloudflarePipelinesR2TablePartitioningJSON contains the
// JSON metadata for the struct
// [SinkNewResponseConfigCloudflarePipelinesR2TablePartitioning]
type sinkNewResponseConfigCloudflarePipelinesR2TablePartitioningJSON struct {
	TimePattern apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseConfigCloudflarePipelinesR2TablePartitioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseConfigCloudflarePipelinesR2TablePartitioningJSON) RawJSON() string {
	return r.raw
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                            `json:"interval_seconds"`
	JSON            sinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON `json:"-"`
}

// sinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON contains the
// JSON metadata for the struct
// [SinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicy]
type sinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON) RawJSON() string {
	return r.raw
}

// R2 Data Catalog Sink
type SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTable struct {
	// Authentication token
	Token string `json:"token,required" format:"var-str"`
	// Cloudflare Account ID
	AccountID string `json:"account_id,required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket string `json:"bucket,required"`
	// Table name
	TableName string `json:"table_name,required"`
	// Table namespace
	Namespace string `json:"namespace"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy `json:"rolling_policy"`
	JSON          sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableJSON          `json:"-"`
}

// sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableJSON contains the JSON
// metadata for the struct
// [SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTable]
type sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableJSON struct {
	Token         apijson.Field
	AccountID     apijson.Field
	Bucket        apijson.Field
	TableName     apijson.Field
	Namespace     apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTable) implementsSinkNewResponseConfig() {
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                                       `json:"interval_seconds"`
	JSON            sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON `json:"-"`
}

// sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON
// contains the JSON metadata for the struct
// [SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy]
type sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON) RawJSON() string {
	return r.raw
}

type SinkNewResponseFormat struct {
	Type            SinkNewResponseFormatType            `json:"type,required"`
	Compression     SinkNewResponseFormatCompression     `json:"compression"`
	DecimalEncoding SinkNewResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                `json:"row_group_bytes,nullable"`
	TimestampFormat SinkNewResponseFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                 `json:"unstructured"`
	JSON            sinkNewResponseFormatJSON            `json:"-"`
	union           SinkNewResponseFormatUnion
}

// sinkNewResponseFormatJSON contains the JSON metadata for the struct
// [SinkNewResponseFormat]
type sinkNewResponseFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r sinkNewResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r *SinkNewResponseFormat) UnmarshalJSON(data []byte) (err error) {
	*r = SinkNewResponseFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkNewResponseFormatUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [SinkNewResponseFormatJson],
// [SinkNewResponseFormatParquet].
func (r SinkNewResponseFormat) AsUnion() SinkNewResponseFormatUnion {
	return r.union
}

// Union satisfied by [SinkNewResponseFormatJson] or
// [SinkNewResponseFormatParquet].
type SinkNewResponseFormatUnion interface {
	implementsSinkNewResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkNewResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkNewResponseFormatJson struct {
	Type            SinkNewResponseFormatJsonType            `json:"type,required"`
	DecimalEncoding SinkNewResponseFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkNewResponseFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                     `json:"unstructured"`
	JSON            sinkNewResponseFormatJsonJSON            `json:"-"`
}

// sinkNewResponseFormatJsonJSON contains the JSON metadata for the struct
// [SinkNewResponseFormatJson]
type sinkNewResponseFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkNewResponseFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseFormatJson) implementsSinkNewResponseFormat() {}

type SinkNewResponseFormatJsonType string

const (
	SinkNewResponseFormatJsonTypeJson SinkNewResponseFormatJsonType = "json"
)

func (r SinkNewResponseFormatJsonType) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkNewResponseFormatJsonDecimalEncoding string

const (
	SinkNewResponseFormatJsonDecimalEncodingNumber SinkNewResponseFormatJsonDecimalEncoding = "number"
	SinkNewResponseFormatJsonDecimalEncodingString SinkNewResponseFormatJsonDecimalEncoding = "string"
	SinkNewResponseFormatJsonDecimalEncodingBytes  SinkNewResponseFormatJsonDecimalEncoding = "bytes"
)

func (r SinkNewResponseFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatJsonDecimalEncodingNumber, SinkNewResponseFormatJsonDecimalEncodingString, SinkNewResponseFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewResponseFormatJsonTimestampFormat string

const (
	SinkNewResponseFormatJsonTimestampFormatRfc3339    SinkNewResponseFormatJsonTimestampFormat = "rfc3339"
	SinkNewResponseFormatJsonTimestampFormatUnixMillis SinkNewResponseFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkNewResponseFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatJsonTimestampFormatRfc3339, SinkNewResponseFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewResponseFormatParquet struct {
	Type          SinkNewResponseFormatParquetType        `json:"type,required"`
	Compression   SinkNewResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                   `json:"row_group_bytes,nullable"`
	JSON          sinkNewResponseFormatParquetJSON        `json:"-"`
}

// sinkNewResponseFormatParquetJSON contains the JSON metadata for the struct
// [SinkNewResponseFormatParquet]
type sinkNewResponseFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkNewResponseFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseFormatParquet) implementsSinkNewResponseFormat() {}

type SinkNewResponseFormatParquetType string

const (
	SinkNewResponseFormatParquetTypeParquet SinkNewResponseFormatParquetType = "parquet"
)

func (r SinkNewResponseFormatParquetType) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkNewResponseFormatParquetCompression string

const (
	SinkNewResponseFormatParquetCompressionUncompressed SinkNewResponseFormatParquetCompression = "uncompressed"
	SinkNewResponseFormatParquetCompressionSnappy       SinkNewResponseFormatParquetCompression = "snappy"
	SinkNewResponseFormatParquetCompressionGzip         SinkNewResponseFormatParquetCompression = "gzip"
	SinkNewResponseFormatParquetCompressionZstd         SinkNewResponseFormatParquetCompression = "zstd"
	SinkNewResponseFormatParquetCompressionLz4          SinkNewResponseFormatParquetCompression = "lz4"
)

func (r SinkNewResponseFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatParquetCompressionUncompressed, SinkNewResponseFormatParquetCompressionSnappy, SinkNewResponseFormatParquetCompressionGzip, SinkNewResponseFormatParquetCompressionZstd, SinkNewResponseFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkNewResponseFormatType string

const (
	SinkNewResponseFormatTypeJson    SinkNewResponseFormatType = "json"
	SinkNewResponseFormatTypeParquet SinkNewResponseFormatType = "parquet"
)

func (r SinkNewResponseFormatType) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatTypeJson, SinkNewResponseFormatTypeParquet:
		return true
	}
	return false
}

type SinkNewResponseFormatCompression string

const (
	SinkNewResponseFormatCompressionUncompressed SinkNewResponseFormatCompression = "uncompressed"
	SinkNewResponseFormatCompressionSnappy       SinkNewResponseFormatCompression = "snappy"
	SinkNewResponseFormatCompressionGzip         SinkNewResponseFormatCompression = "gzip"
	SinkNewResponseFormatCompressionZstd         SinkNewResponseFormatCompression = "zstd"
	SinkNewResponseFormatCompressionLz4          SinkNewResponseFormatCompression = "lz4"
)

func (r SinkNewResponseFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCompressionUncompressed, SinkNewResponseFormatCompressionSnappy, SinkNewResponseFormatCompressionGzip, SinkNewResponseFormatCompressionZstd, SinkNewResponseFormatCompressionLz4:
		return true
	}
	return false
}

type SinkNewResponseFormatDecimalEncoding string

const (
	SinkNewResponseFormatDecimalEncodingNumber SinkNewResponseFormatDecimalEncoding = "number"
	SinkNewResponseFormatDecimalEncodingString SinkNewResponseFormatDecimalEncoding = "string"
	SinkNewResponseFormatDecimalEncodingBytes  SinkNewResponseFormatDecimalEncoding = "bytes"
)

func (r SinkNewResponseFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatDecimalEncodingNumber, SinkNewResponseFormatDecimalEncodingString, SinkNewResponseFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewResponseFormatTimestampFormat string

const (
	SinkNewResponseFormatTimestampFormatRfc3339    SinkNewResponseFormatTimestampFormat = "rfc3339"
	SinkNewResponseFormatTimestampFormatUnixMillis SinkNewResponseFormatTimestampFormat = "unix_millis"
)

func (r SinkNewResponseFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatTimestampFormatRfc3339, SinkNewResponseFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewResponseSchema struct {
	Fields   []SinkNewResponseSchemaField `json:"fields"`
	Format   SinkNewResponseSchemaFormat  `json:"format"`
	Inferred bool                         `json:"inferred,nullable"`
	JSON     sinkNewResponseSchemaJSON    `json:"-"`
}

// sinkNewResponseSchemaJSON contains the JSON metadata for the struct
// [SinkNewResponseSchema]
type sinkNewResponseSchemaJSON struct {
	Fields      apijson.Field
	Format      apijson.Field
	Inferred    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type SinkNewResponseSchemaField struct {
	Type        SinkNewResponseSchemaFieldsType `json:"type,required"`
	MetadataKey string                          `json:"metadata_key,nullable"`
	Name        string                          `json:"name"`
	Required    bool                            `json:"required"`
	SqlName     string                          `json:"sql_name"`
	Unit        SinkNewResponseSchemaFieldsUnit `json:"unit"`
	JSON        sinkNewResponseSchemaFieldJSON  `json:"-"`
	union       SinkNewResponseSchemaFieldsUnion
}

// sinkNewResponseSchemaFieldJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaField]
type sinkNewResponseSchemaFieldJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r sinkNewResponseSchemaFieldJSON) RawJSON() string {
	return r.raw
}

func (r *SinkNewResponseSchemaField) UnmarshalJSON(data []byte) (err error) {
	*r = SinkNewResponseSchemaField{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkNewResponseSchemaFieldsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SinkNewResponseSchemaFieldsInt32],
// [SinkNewResponseSchemaFieldsInt64], [SinkNewResponseSchemaFieldsFloat32],
// [SinkNewResponseSchemaFieldsFloat64], [SinkNewResponseSchemaFieldsBool],
// [SinkNewResponseSchemaFieldsString], [SinkNewResponseSchemaFieldsBinary],
// [SinkNewResponseSchemaFieldsTimestamp], [SinkNewResponseSchemaFieldsJson],
// [SinkNewResponseSchemaFieldsStruct], [SinkNewResponseSchemaFieldsList].
func (r SinkNewResponseSchemaField) AsUnion() SinkNewResponseSchemaFieldsUnion {
	return r.union
}

// Union satisfied by [SinkNewResponseSchemaFieldsInt32],
// [SinkNewResponseSchemaFieldsInt64], [SinkNewResponseSchemaFieldsFloat32],
// [SinkNewResponseSchemaFieldsFloat64], [SinkNewResponseSchemaFieldsBool],
// [SinkNewResponseSchemaFieldsString], [SinkNewResponseSchemaFieldsBinary],
// [SinkNewResponseSchemaFieldsTimestamp], [SinkNewResponseSchemaFieldsJson],
// [SinkNewResponseSchemaFieldsStruct] or [SinkNewResponseSchemaFieldsList].
type SinkNewResponseSchemaFieldsUnion interface {
	implementsSinkNewResponseSchemaField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkNewResponseSchemaFieldsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsInt32{}),
			DiscriminatorValue: "int32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsInt64{}),
			DiscriminatorValue: "int64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsFloat32{}),
			DiscriminatorValue: "float32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsFloat64{}),
			DiscriminatorValue: "float64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsBool{}),
			DiscriminatorValue: "bool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsString{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsBinary{}),
			DiscriminatorValue: "binary",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsTimestamp{}),
			DiscriminatorValue: "timestamp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsStruct{}),
			DiscriminatorValue: "struct",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFieldsList{}),
			DiscriminatorValue: "list",
		},
	)
}

type SinkNewResponseSchemaFieldsInt32 struct {
	Type        SinkNewResponseSchemaFieldsInt32Type `json:"type,required"`
	MetadataKey string                               `json:"metadata_key,nullable"`
	Name        string                               `json:"name"`
	Required    bool                                 `json:"required"`
	SqlName     string                               `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsInt32JSON `json:"-"`
}

// sinkNewResponseSchemaFieldsInt32JSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsInt32]
type sinkNewResponseSchemaFieldsInt32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsInt32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsInt32JSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsInt32) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsInt32Type string

const (
	SinkNewResponseSchemaFieldsInt32TypeInt32 SinkNewResponseSchemaFieldsInt32Type = "int32"
)

func (r SinkNewResponseSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsInt64 struct {
	Type        SinkNewResponseSchemaFieldsInt64Type `json:"type,required"`
	MetadataKey string                               `json:"metadata_key,nullable"`
	Name        string                               `json:"name"`
	Required    bool                                 `json:"required"`
	SqlName     string                               `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsInt64JSON `json:"-"`
}

// sinkNewResponseSchemaFieldsInt64JSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsInt64]
type sinkNewResponseSchemaFieldsInt64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsInt64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsInt64JSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsInt64) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsInt64Type string

const (
	SinkNewResponseSchemaFieldsInt64TypeInt64 SinkNewResponseSchemaFieldsInt64Type = "int64"
)

func (r SinkNewResponseSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsFloat32 struct {
	Type        SinkNewResponseSchemaFieldsFloat32Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsFloat32JSON `json:"-"`
}

// sinkNewResponseSchemaFieldsFloat32JSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsFloat32]
type sinkNewResponseSchemaFieldsFloat32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsFloat32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsFloat32JSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsFloat32) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsFloat32Type string

const (
	SinkNewResponseSchemaFieldsFloat32TypeFloat32 SinkNewResponseSchemaFieldsFloat32Type = "float32"
)

func (r SinkNewResponseSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsFloat64 struct {
	Type        SinkNewResponseSchemaFieldsFloat64Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsFloat64JSON `json:"-"`
}

// sinkNewResponseSchemaFieldsFloat64JSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsFloat64]
type sinkNewResponseSchemaFieldsFloat64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsFloat64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsFloat64JSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsFloat64) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsFloat64Type string

const (
	SinkNewResponseSchemaFieldsFloat64TypeFloat64 SinkNewResponseSchemaFieldsFloat64Type = "float64"
)

func (r SinkNewResponseSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsBool struct {
	Type        SinkNewResponseSchemaFieldsBoolType `json:"type,required"`
	MetadataKey string                              `json:"metadata_key,nullable"`
	Name        string                              `json:"name"`
	Required    bool                                `json:"required"`
	SqlName     string                              `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsBoolJSON `json:"-"`
}

// sinkNewResponseSchemaFieldsBoolJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsBool]
type sinkNewResponseSchemaFieldsBoolJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsBoolJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsBool) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsBoolType string

const (
	SinkNewResponseSchemaFieldsBoolTypeBool SinkNewResponseSchemaFieldsBoolType = "bool"
)

func (r SinkNewResponseSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsString struct {
	Type        SinkNewResponseSchemaFieldsStringType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsStringJSON `json:"-"`
}

// sinkNewResponseSchemaFieldsStringJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsString]
type sinkNewResponseSchemaFieldsStringJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsStringJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsString) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsStringType string

const (
	SinkNewResponseSchemaFieldsStringTypeString SinkNewResponseSchemaFieldsStringType = "string"
)

func (r SinkNewResponseSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsBinary struct {
	Type        SinkNewResponseSchemaFieldsBinaryType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsBinaryJSON `json:"-"`
}

// sinkNewResponseSchemaFieldsBinaryJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsBinary]
type sinkNewResponseSchemaFieldsBinaryJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsBinary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsBinaryJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsBinary) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsBinaryType string

const (
	SinkNewResponseSchemaFieldsBinaryTypeBinary SinkNewResponseSchemaFieldsBinaryType = "binary"
)

func (r SinkNewResponseSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsTimestamp struct {
	Type        SinkNewResponseSchemaFieldsTimestampType `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	Unit        SinkNewResponseSchemaFieldsTimestampUnit `json:"unit"`
	JSON        sinkNewResponseSchemaFieldsTimestampJSON `json:"-"`
}

// sinkNewResponseSchemaFieldsTimestampJSON contains the JSON metadata for the
// struct [SinkNewResponseSchemaFieldsTimestamp]
type sinkNewResponseSchemaFieldsTimestampJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsTimestamp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsTimestampJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsTimestamp) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsTimestampType string

const (
	SinkNewResponseSchemaFieldsTimestampTypeTimestamp SinkNewResponseSchemaFieldsTimestampType = "timestamp"
)

func (r SinkNewResponseSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsTimestampUnit string

const (
	SinkNewResponseSchemaFieldsTimestampUnitSecond      SinkNewResponseSchemaFieldsTimestampUnit = "second"
	SinkNewResponseSchemaFieldsTimestampUnitMillisecond SinkNewResponseSchemaFieldsTimestampUnit = "millisecond"
	SinkNewResponseSchemaFieldsTimestampUnitMicrosecond SinkNewResponseSchemaFieldsTimestampUnit = "microsecond"
	SinkNewResponseSchemaFieldsTimestampUnitNanosecond  SinkNewResponseSchemaFieldsTimestampUnit = "nanosecond"
)

func (r SinkNewResponseSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsTimestampUnitSecond, SinkNewResponseSchemaFieldsTimestampUnitMillisecond, SinkNewResponseSchemaFieldsTimestampUnitMicrosecond, SinkNewResponseSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsJson struct {
	Type        SinkNewResponseSchemaFieldsJsonType `json:"type,required"`
	MetadataKey string                              `json:"metadata_key,nullable"`
	Name        string                              `json:"name"`
	Required    bool                                `json:"required"`
	SqlName     string                              `json:"sql_name"`
	JSON        sinkNewResponseSchemaFieldsJsonJSON `json:"-"`
}

// sinkNewResponseSchemaFieldsJsonJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsJson]
type sinkNewResponseSchemaFieldsJsonJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsJson) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsJsonType string

const (
	SinkNewResponseSchemaFieldsJsonTypeJson SinkNewResponseSchemaFieldsJsonType = "json"
)

func (r SinkNewResponseSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsStruct struct {
	JSON sinkNewResponseSchemaFieldsStructJSON `json:"-"`
}

// sinkNewResponseSchemaFieldsStructJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsStruct]
type sinkNewResponseSchemaFieldsStructJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsStruct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsStructJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsStruct) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsList struct {
	JSON sinkNewResponseSchemaFieldsListJSON `json:"-"`
}

// sinkNewResponseSchemaFieldsListJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFieldsList]
type sinkNewResponseSchemaFieldsListJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFieldsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFieldsListJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFieldsList) implementsSinkNewResponseSchemaField() {}

type SinkNewResponseSchemaFieldsType string

const (
	SinkNewResponseSchemaFieldsTypeInt32     SinkNewResponseSchemaFieldsType = "int32"
	SinkNewResponseSchemaFieldsTypeInt64     SinkNewResponseSchemaFieldsType = "int64"
	SinkNewResponseSchemaFieldsTypeFloat32   SinkNewResponseSchemaFieldsType = "float32"
	SinkNewResponseSchemaFieldsTypeFloat64   SinkNewResponseSchemaFieldsType = "float64"
	SinkNewResponseSchemaFieldsTypeBool      SinkNewResponseSchemaFieldsType = "bool"
	SinkNewResponseSchemaFieldsTypeString    SinkNewResponseSchemaFieldsType = "string"
	SinkNewResponseSchemaFieldsTypeBinary    SinkNewResponseSchemaFieldsType = "binary"
	SinkNewResponseSchemaFieldsTypeTimestamp SinkNewResponseSchemaFieldsType = "timestamp"
	SinkNewResponseSchemaFieldsTypeJson      SinkNewResponseSchemaFieldsType = "json"
	SinkNewResponseSchemaFieldsTypeStruct    SinkNewResponseSchemaFieldsType = "struct"
	SinkNewResponseSchemaFieldsTypeList      SinkNewResponseSchemaFieldsType = "list"
)

func (r SinkNewResponseSchemaFieldsType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsTypeInt32, SinkNewResponseSchemaFieldsTypeInt64, SinkNewResponseSchemaFieldsTypeFloat32, SinkNewResponseSchemaFieldsTypeFloat64, SinkNewResponseSchemaFieldsTypeBool, SinkNewResponseSchemaFieldsTypeString, SinkNewResponseSchemaFieldsTypeBinary, SinkNewResponseSchemaFieldsTypeTimestamp, SinkNewResponseSchemaFieldsTypeJson, SinkNewResponseSchemaFieldsTypeStruct, SinkNewResponseSchemaFieldsTypeList:
		return true
	}
	return false
}

type SinkNewResponseSchemaFieldsUnit string

const (
	SinkNewResponseSchemaFieldsUnitSecond      SinkNewResponseSchemaFieldsUnit = "second"
	SinkNewResponseSchemaFieldsUnitMillisecond SinkNewResponseSchemaFieldsUnit = "millisecond"
	SinkNewResponseSchemaFieldsUnitMicrosecond SinkNewResponseSchemaFieldsUnit = "microsecond"
	SinkNewResponseSchemaFieldsUnitNanosecond  SinkNewResponseSchemaFieldsUnit = "nanosecond"
)

func (r SinkNewResponseSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFieldsUnitSecond, SinkNewResponseSchemaFieldsUnitMillisecond, SinkNewResponseSchemaFieldsUnitMicrosecond, SinkNewResponseSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormat struct {
	Type            SinkNewResponseSchemaFormatType            `json:"type,required"`
	Compression     SinkNewResponseSchemaFormatCompression     `json:"compression"`
	DecimalEncoding SinkNewResponseSchemaFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                      `json:"row_group_bytes,nullable"`
	TimestampFormat SinkNewResponseSchemaFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                       `json:"unstructured"`
	JSON            sinkNewResponseSchemaFormatJSON            `json:"-"`
	union           SinkNewResponseSchemaFormatUnion
}

// sinkNewResponseSchemaFormatJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFormat]
type sinkNewResponseSchemaFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r sinkNewResponseSchemaFormatJSON) RawJSON() string {
	return r.raw
}

func (r *SinkNewResponseSchemaFormat) UnmarshalJSON(data []byte) (err error) {
	*r = SinkNewResponseSchemaFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkNewResponseSchemaFormatUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SinkNewResponseSchemaFormatJson],
// [SinkNewResponseSchemaFormatParquet].
func (r SinkNewResponseSchemaFormat) AsUnion() SinkNewResponseSchemaFormatUnion {
	return r.union
}

// Union satisfied by [SinkNewResponseSchemaFormatJson] or
// [SinkNewResponseSchemaFormatParquet].
type SinkNewResponseSchemaFormatUnion interface {
	implementsSinkNewResponseSchemaFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkNewResponseSchemaFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseSchemaFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkNewResponseSchemaFormatJson struct {
	Type            SinkNewResponseSchemaFormatJsonType            `json:"type,required"`
	DecimalEncoding SinkNewResponseSchemaFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkNewResponseSchemaFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                           `json:"unstructured"`
	JSON            sinkNewResponseSchemaFormatJsonJSON            `json:"-"`
}

// sinkNewResponseSchemaFormatJsonJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFormatJson]
type sinkNewResponseSchemaFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFormatJson) implementsSinkNewResponseSchemaFormat() {}

type SinkNewResponseSchemaFormatJsonType string

const (
	SinkNewResponseSchemaFormatJsonTypeJson SinkNewResponseSchemaFormatJsonType = "json"
)

func (r SinkNewResponseSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatJsonDecimalEncoding string

const (
	SinkNewResponseSchemaFormatJsonDecimalEncodingNumber SinkNewResponseSchemaFormatJsonDecimalEncoding = "number"
	SinkNewResponseSchemaFormatJsonDecimalEncodingString SinkNewResponseSchemaFormatJsonDecimalEncoding = "string"
	SinkNewResponseSchemaFormatJsonDecimalEncodingBytes  SinkNewResponseSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r SinkNewResponseSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatJsonDecimalEncodingNumber, SinkNewResponseSchemaFormatJsonDecimalEncodingString, SinkNewResponseSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatJsonTimestampFormat string

const (
	SinkNewResponseSchemaFormatJsonTimestampFormatRfc3339    SinkNewResponseSchemaFormatJsonTimestampFormat = "rfc3339"
	SinkNewResponseSchemaFormatJsonTimestampFormatUnixMillis SinkNewResponseSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkNewResponseSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatJsonTimestampFormatRfc3339, SinkNewResponseSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatParquet struct {
	Type          SinkNewResponseSchemaFormatParquetType        `json:"type,required"`
	Compression   SinkNewResponseSchemaFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                         `json:"row_group_bytes,nullable"`
	JSON          sinkNewResponseSchemaFormatParquetJSON        `json:"-"`
}

// sinkNewResponseSchemaFormatParquetJSON contains the JSON metadata for the struct
// [SinkNewResponseSchemaFormatParquet]
type sinkNewResponseSchemaFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkNewResponseSchemaFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseSchemaFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseSchemaFormatParquet) implementsSinkNewResponseSchemaFormat() {}

type SinkNewResponseSchemaFormatParquetType string

const (
	SinkNewResponseSchemaFormatParquetTypeParquet SinkNewResponseSchemaFormatParquetType = "parquet"
)

func (r SinkNewResponseSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatParquetCompression string

const (
	SinkNewResponseSchemaFormatParquetCompressionUncompressed SinkNewResponseSchemaFormatParquetCompression = "uncompressed"
	SinkNewResponseSchemaFormatParquetCompressionSnappy       SinkNewResponseSchemaFormatParquetCompression = "snappy"
	SinkNewResponseSchemaFormatParquetCompressionGzip         SinkNewResponseSchemaFormatParquetCompression = "gzip"
	SinkNewResponseSchemaFormatParquetCompressionZstd         SinkNewResponseSchemaFormatParquetCompression = "zstd"
	SinkNewResponseSchemaFormatParquetCompressionLz4          SinkNewResponseSchemaFormatParquetCompression = "lz4"
)

func (r SinkNewResponseSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatParquetCompressionUncompressed, SinkNewResponseSchemaFormatParquetCompressionSnappy, SinkNewResponseSchemaFormatParquetCompressionGzip, SinkNewResponseSchemaFormatParquetCompressionZstd, SinkNewResponseSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatType string

const (
	SinkNewResponseSchemaFormatTypeJson    SinkNewResponseSchemaFormatType = "json"
	SinkNewResponseSchemaFormatTypeParquet SinkNewResponseSchemaFormatType = "parquet"
)

func (r SinkNewResponseSchemaFormatType) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatTypeJson, SinkNewResponseSchemaFormatTypeParquet:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatCompression string

const (
	SinkNewResponseSchemaFormatCompressionUncompressed SinkNewResponseSchemaFormatCompression = "uncompressed"
	SinkNewResponseSchemaFormatCompressionSnappy       SinkNewResponseSchemaFormatCompression = "snappy"
	SinkNewResponseSchemaFormatCompressionGzip         SinkNewResponseSchemaFormatCompression = "gzip"
	SinkNewResponseSchemaFormatCompressionZstd         SinkNewResponseSchemaFormatCompression = "zstd"
	SinkNewResponseSchemaFormatCompressionLz4          SinkNewResponseSchemaFormatCompression = "lz4"
)

func (r SinkNewResponseSchemaFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatCompressionUncompressed, SinkNewResponseSchemaFormatCompressionSnappy, SinkNewResponseSchemaFormatCompressionGzip, SinkNewResponseSchemaFormatCompressionZstd, SinkNewResponseSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatDecimalEncoding string

const (
	SinkNewResponseSchemaFormatDecimalEncodingNumber SinkNewResponseSchemaFormatDecimalEncoding = "number"
	SinkNewResponseSchemaFormatDecimalEncodingString SinkNewResponseSchemaFormatDecimalEncoding = "string"
	SinkNewResponseSchemaFormatDecimalEncodingBytes  SinkNewResponseSchemaFormatDecimalEncoding = "bytes"
)

func (r SinkNewResponseSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatDecimalEncodingNumber, SinkNewResponseSchemaFormatDecimalEncodingString, SinkNewResponseSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewResponseSchemaFormatTimestampFormat string

const (
	SinkNewResponseSchemaFormatTimestampFormatRfc3339    SinkNewResponseSchemaFormatTimestampFormat = "rfc3339"
	SinkNewResponseSchemaFormatTimestampFormatUnixMillis SinkNewResponseSchemaFormatTimestampFormat = "unix_millis"
)

func (r SinkNewResponseSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewResponseSchemaFormatTimestampFormatRfc3339, SinkNewResponseSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkListResponse struct {
	// Indicates a unique identifier for this sink.
	ID         string    `json:"id,required"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Defines the name of the Sink.
	Name string `json:"name,required"`
	// Specifies the type of sink.
	Type SinkListResponseType `json:"type,required"`
	// Defines the configuration of the R2 Sink.
	Config SinkListResponseConfig `json:"config"`
	Format SinkListResponseFormat `json:"format"`
	Schema SinkListResponseSchema `json:"schema"`
	JSON   sinkListResponseJSON   `json:"-"`
}

// sinkListResponseJSON contains the JSON metadata for the struct
// [SinkListResponse]
type sinkListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Config      apijson.Field
	Format      apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of sink.
type SinkListResponseType string

const (
	SinkListResponseTypeR2            SinkListResponseType = "r2"
	SinkListResponseTypeR2DataCatalog SinkListResponseType = "r2_data_catalog"
)

func (r SinkListResponseType) IsKnown() bool {
	switch r {
	case SinkListResponseTypeR2, SinkListResponseTypeR2DataCatalog:
		return true
	}
	return false
}

// Defines the configuration of the R2 Sink.
type SinkListResponseConfig struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket,required"`
	// Authentication token
	Token string `json:"token" format:"var-str"`
	// This field can have the runtime type of
	// [SinkListResponseConfigCloudflarePipelinesR2TableCredentials].
	Credentials interface{} `json:"credentials"`
	// This field can have the runtime type of
	// [SinkListResponseConfigCloudflarePipelinesR2TableFileNaming].
	FileNaming interface{} `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Table namespace
	Namespace string `json:"namespace"`
	// This field can have the runtime type of
	// [SinkListResponseConfigCloudflarePipelinesR2TablePartitioning].
	Partitioning interface{} `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// This field can have the runtime type of
	// [SinkListResponseConfigCloudflarePipelinesR2TableRollingPolicy],
	// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy].
	RollingPolicy interface{} `json:"rolling_policy"`
	// Table name
	TableName string                     `json:"table_name"`
	JSON      sinkListResponseConfigJSON `json:"-"`
	union     SinkListResponseConfigUnion
}

// sinkListResponseConfigJSON contains the JSON metadata for the struct
// [SinkListResponseConfig]
type sinkListResponseConfigJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	Token         apijson.Field
	Credentials   apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Namespace     apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	TableName     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r sinkListResponseConfigJSON) RawJSON() string {
	return r.raw
}

func (r *SinkListResponseConfig) UnmarshalJSON(data []byte) (err error) {
	*r = SinkListResponseConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkListResponseConfigUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [SinkListResponseConfigCloudflarePipelinesR2Table],
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTable].
func (r SinkListResponseConfig) AsUnion() SinkListResponseConfigUnion {
	return r.union
}

// Defines the configuration of the R2 Sink.
//
// Union satisfied by [SinkListResponseConfigCloudflarePipelinesR2Table] or
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTable].
type SinkListResponseConfigUnion interface {
	implementsSinkListResponseConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkListResponseConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkListResponseConfigCloudflarePipelinesR2Table{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkListResponseConfigCloudflarePipelinesR2DataCatalogTable{}),
		},
	)
}

type SinkListResponseConfigCloudflarePipelinesR2Table struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket      string                                                      `json:"bucket,required"`
	Credentials SinkListResponseConfigCloudflarePipelinesR2TableCredentials `json:"credentials,required"`
	// Controls filename prefix/suffix and strategy.
	FileNaming SinkListResponseConfigCloudflarePipelinesR2TableFileNaming `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Data-layout partitioning for sinks.
	Partitioning SinkListResponseConfigCloudflarePipelinesR2TablePartitioning `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkListResponseConfigCloudflarePipelinesR2TableRollingPolicy `json:"rolling_policy"`
	JSON          sinkListResponseConfigCloudflarePipelinesR2TableJSON          `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TableJSON contains the JSON metadata
// for the struct [SinkListResponseConfigCloudflarePipelinesR2Table]
type sinkListResponseConfigCloudflarePipelinesR2TableJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	Credentials   apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2Table) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TableJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseConfigCloudflarePipelinesR2Table) implementsSinkListResponseConfig() {}

type SinkListResponseConfigCloudflarePipelinesR2TableCredentials struct {
	// Cloudflare Account ID for the bucket
	AccessKeyID string `json:"access_key_id,required" format:"var-str"`
	// Cloudflare Account ID for the bucket
	SecretAccessKey string                                                          `json:"secret_access_key,required" format:"var-str"`
	JSON            sinkListResponseConfigCloudflarePipelinesR2TableCredentialsJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TableCredentialsJSON contains the
// JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2TableCredentials]
type sinkListResponseConfigCloudflarePipelinesR2TableCredentialsJSON struct {
	AccessKeyID     apijson.Field
	SecretAccessKey apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TableCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TableCredentialsJSON) RawJSON() string {
	return r.raw
}

// Controls filename prefix/suffix and strategy.
type SinkListResponseConfigCloudflarePipelinesR2TableFileNaming struct {
	// The prefix to use in file name. i.e prefix-<uuid>.parquet
	Prefix string `json:"prefix"`
	// Filename generation strategy.
	Strategy SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategy `json:"strategy"`
	// This will overwrite the default file suffix. i.e .parquet, use with caution
	Suffix string                                                         `json:"suffix"`
	JSON   sinkListResponseConfigCloudflarePipelinesR2TableFileNamingJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TableFileNamingJSON contains the JSON
// metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2TableFileNaming]
type sinkListResponseConfigCloudflarePipelinesR2TableFileNamingJSON struct {
	Prefix      apijson.Field
	Strategy    apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TableFileNaming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TableFileNamingJSON) RawJSON() string {
	return r.raw
}

// Filename generation strategy.
type SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategy string

const (
	SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategySerial SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "serial"
	SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUID   SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid"
	SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7 SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid_v7"
	SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUlid   SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "ulid"
)

func (r SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategy) IsKnown() bool {
	switch r {
	case SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategySerial, SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUID, SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7, SinkListResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUlid:
		return true
	}
	return false
}

// Data-layout partitioning for sinks.
type SinkListResponseConfigCloudflarePipelinesR2TablePartitioning struct {
	// The pattern of the date string
	TimePattern string                                                           `json:"time_pattern"`
	JSON        sinkListResponseConfigCloudflarePipelinesR2TablePartitioningJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TablePartitioningJSON contains the
// JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2TablePartitioning]
type sinkListResponseConfigCloudflarePipelinesR2TablePartitioningJSON struct {
	TimePattern apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TablePartitioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TablePartitioningJSON) RawJSON() string {
	return r.raw
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkListResponseConfigCloudflarePipelinesR2TableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                             `json:"interval_seconds"`
	JSON            sinkListResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON contains the
// JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2TableRollingPolicy]
type sinkListResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TableRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON) RawJSON() string {
	return r.raw
}

// R2 Data Catalog Sink
type SinkListResponseConfigCloudflarePipelinesR2DataCatalogTable struct {
	// Authentication token
	Token string `json:"token,required" format:"var-str"`
	// Cloudflare Account ID
	AccountID string `json:"account_id,required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket string `json:"bucket,required"`
	// Table name
	TableName string `json:"table_name,required"`
	// Table namespace
	Namespace string `json:"namespace"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy `json:"rolling_policy"`
	JSON          sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableJSON          `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableJSON contains the
// JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTable]
type sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableJSON struct {
	Token         apijson.Field
	AccountID     apijson.Field
	Bucket        apijson.Field
	TableName     apijson.Field
	Namespace     apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2DataCatalogTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseConfigCloudflarePipelinesR2DataCatalogTable) implementsSinkListResponseConfig() {
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                                        `json:"interval_seconds"`
	JSON            sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON
// contains the JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy]
type sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON) RawJSON() string {
	return r.raw
}

type SinkListResponseFormat struct {
	Type            SinkListResponseFormatType            `json:"type,required"`
	Compression     SinkListResponseFormatCompression     `json:"compression"`
	DecimalEncoding SinkListResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                 `json:"row_group_bytes,nullable"`
	TimestampFormat SinkListResponseFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                  `json:"unstructured"`
	JSON            sinkListResponseFormatJSON            `json:"-"`
	union           SinkListResponseFormatUnion
}

// sinkListResponseFormatJSON contains the JSON metadata for the struct
// [SinkListResponseFormat]
type sinkListResponseFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r sinkListResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r *SinkListResponseFormat) UnmarshalJSON(data []byte) (err error) {
	*r = SinkListResponseFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkListResponseFormatUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [SinkListResponseFormatJson],
// [SinkListResponseFormatParquet].
func (r SinkListResponseFormat) AsUnion() SinkListResponseFormatUnion {
	return r.union
}

// Union satisfied by [SinkListResponseFormatJson] or
// [SinkListResponseFormatParquet].
type SinkListResponseFormatUnion interface {
	implementsSinkListResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkListResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkListResponseFormatJson struct {
	Type            SinkListResponseFormatJsonType            `json:"type,required"`
	DecimalEncoding SinkListResponseFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkListResponseFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                      `json:"unstructured"`
	JSON            sinkListResponseFormatJsonJSON            `json:"-"`
}

// sinkListResponseFormatJsonJSON contains the JSON metadata for the struct
// [SinkListResponseFormatJson]
type sinkListResponseFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkListResponseFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseFormatJson) implementsSinkListResponseFormat() {}

type SinkListResponseFormatJsonType string

const (
	SinkListResponseFormatJsonTypeJson SinkListResponseFormatJsonType = "json"
)

func (r SinkListResponseFormatJsonType) IsKnown() bool {
	switch r {
	case SinkListResponseFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkListResponseFormatJsonDecimalEncoding string

const (
	SinkListResponseFormatJsonDecimalEncodingNumber SinkListResponseFormatJsonDecimalEncoding = "number"
	SinkListResponseFormatJsonDecimalEncodingString SinkListResponseFormatJsonDecimalEncoding = "string"
	SinkListResponseFormatJsonDecimalEncodingBytes  SinkListResponseFormatJsonDecimalEncoding = "bytes"
)

func (r SinkListResponseFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkListResponseFormatJsonDecimalEncodingNumber, SinkListResponseFormatJsonDecimalEncodingString, SinkListResponseFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkListResponseFormatJsonTimestampFormat string

const (
	SinkListResponseFormatJsonTimestampFormatRfc3339    SinkListResponseFormatJsonTimestampFormat = "rfc3339"
	SinkListResponseFormatJsonTimestampFormatUnixMillis SinkListResponseFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkListResponseFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkListResponseFormatJsonTimestampFormatRfc3339, SinkListResponseFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkListResponseFormatParquet struct {
	Type          SinkListResponseFormatParquetType        `json:"type,required"`
	Compression   SinkListResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                    `json:"row_group_bytes,nullable"`
	JSON          sinkListResponseFormatParquetJSON        `json:"-"`
}

// sinkListResponseFormatParquetJSON contains the JSON metadata for the struct
// [SinkListResponseFormatParquet]
type sinkListResponseFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkListResponseFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseFormatParquet) implementsSinkListResponseFormat() {}

type SinkListResponseFormatParquetType string

const (
	SinkListResponseFormatParquetTypeParquet SinkListResponseFormatParquetType = "parquet"
)

func (r SinkListResponseFormatParquetType) IsKnown() bool {
	switch r {
	case SinkListResponseFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkListResponseFormatParquetCompression string

const (
	SinkListResponseFormatParquetCompressionUncompressed SinkListResponseFormatParquetCompression = "uncompressed"
	SinkListResponseFormatParquetCompressionSnappy       SinkListResponseFormatParquetCompression = "snappy"
	SinkListResponseFormatParquetCompressionGzip         SinkListResponseFormatParquetCompression = "gzip"
	SinkListResponseFormatParquetCompressionZstd         SinkListResponseFormatParquetCompression = "zstd"
	SinkListResponseFormatParquetCompressionLz4          SinkListResponseFormatParquetCompression = "lz4"
)

func (r SinkListResponseFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkListResponseFormatParquetCompressionUncompressed, SinkListResponseFormatParquetCompressionSnappy, SinkListResponseFormatParquetCompressionGzip, SinkListResponseFormatParquetCompressionZstd, SinkListResponseFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkListResponseFormatType string

const (
	SinkListResponseFormatTypeJson    SinkListResponseFormatType = "json"
	SinkListResponseFormatTypeParquet SinkListResponseFormatType = "parquet"
)

func (r SinkListResponseFormatType) IsKnown() bool {
	switch r {
	case SinkListResponseFormatTypeJson, SinkListResponseFormatTypeParquet:
		return true
	}
	return false
}

type SinkListResponseFormatCompression string

const (
	SinkListResponseFormatCompressionUncompressed SinkListResponseFormatCompression = "uncompressed"
	SinkListResponseFormatCompressionSnappy       SinkListResponseFormatCompression = "snappy"
	SinkListResponseFormatCompressionGzip         SinkListResponseFormatCompression = "gzip"
	SinkListResponseFormatCompressionZstd         SinkListResponseFormatCompression = "zstd"
	SinkListResponseFormatCompressionLz4          SinkListResponseFormatCompression = "lz4"
)

func (r SinkListResponseFormatCompression) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCompressionUncompressed, SinkListResponseFormatCompressionSnappy, SinkListResponseFormatCompressionGzip, SinkListResponseFormatCompressionZstd, SinkListResponseFormatCompressionLz4:
		return true
	}
	return false
}

type SinkListResponseFormatDecimalEncoding string

const (
	SinkListResponseFormatDecimalEncodingNumber SinkListResponseFormatDecimalEncoding = "number"
	SinkListResponseFormatDecimalEncodingString SinkListResponseFormatDecimalEncoding = "string"
	SinkListResponseFormatDecimalEncodingBytes  SinkListResponseFormatDecimalEncoding = "bytes"
)

func (r SinkListResponseFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkListResponseFormatDecimalEncodingNumber, SinkListResponseFormatDecimalEncodingString, SinkListResponseFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkListResponseFormatTimestampFormat string

const (
	SinkListResponseFormatTimestampFormatRfc3339    SinkListResponseFormatTimestampFormat = "rfc3339"
	SinkListResponseFormatTimestampFormatUnixMillis SinkListResponseFormatTimestampFormat = "unix_millis"
)

func (r SinkListResponseFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkListResponseFormatTimestampFormatRfc3339, SinkListResponseFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkListResponseSchema struct {
	Fields   []SinkListResponseSchemaField `json:"fields"`
	Format   SinkListResponseSchemaFormat  `json:"format"`
	Inferred bool                          `json:"inferred,nullable"`
	JSON     sinkListResponseSchemaJSON    `json:"-"`
}

// sinkListResponseSchemaJSON contains the JSON metadata for the struct
// [SinkListResponseSchema]
type sinkListResponseSchemaJSON struct {
	Fields      apijson.Field
	Format      apijson.Field
	Inferred    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type SinkListResponseSchemaField struct {
	Type        SinkListResponseSchemaFieldsType `json:"type,required"`
	MetadataKey string                           `json:"metadata_key,nullable"`
	Name        string                           `json:"name"`
	Required    bool                             `json:"required"`
	SqlName     string                           `json:"sql_name"`
	Unit        SinkListResponseSchemaFieldsUnit `json:"unit"`
	JSON        sinkListResponseSchemaFieldJSON  `json:"-"`
	union       SinkListResponseSchemaFieldsUnion
}

// sinkListResponseSchemaFieldJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaField]
type sinkListResponseSchemaFieldJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r sinkListResponseSchemaFieldJSON) RawJSON() string {
	return r.raw
}

func (r *SinkListResponseSchemaField) UnmarshalJSON(data []byte) (err error) {
	*r = SinkListResponseSchemaField{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkListResponseSchemaFieldsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SinkListResponseSchemaFieldsInt32],
// [SinkListResponseSchemaFieldsInt64], [SinkListResponseSchemaFieldsFloat32],
// [SinkListResponseSchemaFieldsFloat64], [SinkListResponseSchemaFieldsBool],
// [SinkListResponseSchemaFieldsString], [SinkListResponseSchemaFieldsBinary],
// [SinkListResponseSchemaFieldsTimestamp], [SinkListResponseSchemaFieldsJson],
// [SinkListResponseSchemaFieldsStruct], [SinkListResponseSchemaFieldsList].
func (r SinkListResponseSchemaField) AsUnion() SinkListResponseSchemaFieldsUnion {
	return r.union
}

// Union satisfied by [SinkListResponseSchemaFieldsInt32],
// [SinkListResponseSchemaFieldsInt64], [SinkListResponseSchemaFieldsFloat32],
// [SinkListResponseSchemaFieldsFloat64], [SinkListResponseSchemaFieldsBool],
// [SinkListResponseSchemaFieldsString], [SinkListResponseSchemaFieldsBinary],
// [SinkListResponseSchemaFieldsTimestamp], [SinkListResponseSchemaFieldsJson],
// [SinkListResponseSchemaFieldsStruct] or [SinkListResponseSchemaFieldsList].
type SinkListResponseSchemaFieldsUnion interface {
	implementsSinkListResponseSchemaField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkListResponseSchemaFieldsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsInt32{}),
			DiscriminatorValue: "int32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsInt64{}),
			DiscriminatorValue: "int64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsFloat32{}),
			DiscriminatorValue: "float32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsFloat64{}),
			DiscriminatorValue: "float64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsBool{}),
			DiscriminatorValue: "bool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsString{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsBinary{}),
			DiscriminatorValue: "binary",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsTimestamp{}),
			DiscriminatorValue: "timestamp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsStruct{}),
			DiscriminatorValue: "struct",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFieldsList{}),
			DiscriminatorValue: "list",
		},
	)
}

type SinkListResponseSchemaFieldsInt32 struct {
	Type        SinkListResponseSchemaFieldsInt32Type `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsInt32JSON `json:"-"`
}

// sinkListResponseSchemaFieldsInt32JSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsInt32]
type sinkListResponseSchemaFieldsInt32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsInt32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsInt32JSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsInt32) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsInt32Type string

const (
	SinkListResponseSchemaFieldsInt32TypeInt32 SinkListResponseSchemaFieldsInt32Type = "int32"
)

func (r SinkListResponseSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsInt64 struct {
	Type        SinkListResponseSchemaFieldsInt64Type `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsInt64JSON `json:"-"`
}

// sinkListResponseSchemaFieldsInt64JSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsInt64]
type sinkListResponseSchemaFieldsInt64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsInt64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsInt64JSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsInt64) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsInt64Type string

const (
	SinkListResponseSchemaFieldsInt64TypeInt64 SinkListResponseSchemaFieldsInt64Type = "int64"
)

func (r SinkListResponseSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsFloat32 struct {
	Type        SinkListResponseSchemaFieldsFloat32Type `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsFloat32JSON `json:"-"`
}

// sinkListResponseSchemaFieldsFloat32JSON contains the JSON metadata for the
// struct [SinkListResponseSchemaFieldsFloat32]
type sinkListResponseSchemaFieldsFloat32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsFloat32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsFloat32JSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsFloat32) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsFloat32Type string

const (
	SinkListResponseSchemaFieldsFloat32TypeFloat32 SinkListResponseSchemaFieldsFloat32Type = "float32"
)

func (r SinkListResponseSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsFloat64 struct {
	Type        SinkListResponseSchemaFieldsFloat64Type `json:"type,required"`
	MetadataKey string                                  `json:"metadata_key,nullable"`
	Name        string                                  `json:"name"`
	Required    bool                                    `json:"required"`
	SqlName     string                                  `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsFloat64JSON `json:"-"`
}

// sinkListResponseSchemaFieldsFloat64JSON contains the JSON metadata for the
// struct [SinkListResponseSchemaFieldsFloat64]
type sinkListResponseSchemaFieldsFloat64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsFloat64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsFloat64JSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsFloat64) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsFloat64Type string

const (
	SinkListResponseSchemaFieldsFloat64TypeFloat64 SinkListResponseSchemaFieldsFloat64Type = "float64"
)

func (r SinkListResponseSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsBool struct {
	Type        SinkListResponseSchemaFieldsBoolType `json:"type,required"`
	MetadataKey string                               `json:"metadata_key,nullable"`
	Name        string                               `json:"name"`
	Required    bool                                 `json:"required"`
	SqlName     string                               `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsBoolJSON `json:"-"`
}

// sinkListResponseSchemaFieldsBoolJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsBool]
type sinkListResponseSchemaFieldsBoolJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsBoolJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsBool) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsBoolType string

const (
	SinkListResponseSchemaFieldsBoolTypeBool SinkListResponseSchemaFieldsBoolType = "bool"
)

func (r SinkListResponseSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsString struct {
	Type        SinkListResponseSchemaFieldsStringType `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsStringJSON `json:"-"`
}

// sinkListResponseSchemaFieldsStringJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsString]
type sinkListResponseSchemaFieldsStringJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsStringJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsString) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsStringType string

const (
	SinkListResponseSchemaFieldsStringTypeString SinkListResponseSchemaFieldsStringType = "string"
)

func (r SinkListResponseSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsBinary struct {
	Type        SinkListResponseSchemaFieldsBinaryType `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsBinaryJSON `json:"-"`
}

// sinkListResponseSchemaFieldsBinaryJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsBinary]
type sinkListResponseSchemaFieldsBinaryJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsBinary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsBinaryJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsBinary) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsBinaryType string

const (
	SinkListResponseSchemaFieldsBinaryTypeBinary SinkListResponseSchemaFieldsBinaryType = "binary"
)

func (r SinkListResponseSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsTimestamp struct {
	Type        SinkListResponseSchemaFieldsTimestampType `json:"type,required"`
	MetadataKey string                                    `json:"metadata_key,nullable"`
	Name        string                                    `json:"name"`
	Required    bool                                      `json:"required"`
	SqlName     string                                    `json:"sql_name"`
	Unit        SinkListResponseSchemaFieldsTimestampUnit `json:"unit"`
	JSON        sinkListResponseSchemaFieldsTimestampJSON `json:"-"`
}

// sinkListResponseSchemaFieldsTimestampJSON contains the JSON metadata for the
// struct [SinkListResponseSchemaFieldsTimestamp]
type sinkListResponseSchemaFieldsTimestampJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsTimestamp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsTimestampJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsTimestamp) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsTimestampType string

const (
	SinkListResponseSchemaFieldsTimestampTypeTimestamp SinkListResponseSchemaFieldsTimestampType = "timestamp"
)

func (r SinkListResponseSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsTimestampUnit string

const (
	SinkListResponseSchemaFieldsTimestampUnitSecond      SinkListResponseSchemaFieldsTimestampUnit = "second"
	SinkListResponseSchemaFieldsTimestampUnitMillisecond SinkListResponseSchemaFieldsTimestampUnit = "millisecond"
	SinkListResponseSchemaFieldsTimestampUnitMicrosecond SinkListResponseSchemaFieldsTimestampUnit = "microsecond"
	SinkListResponseSchemaFieldsTimestampUnitNanosecond  SinkListResponseSchemaFieldsTimestampUnit = "nanosecond"
)

func (r SinkListResponseSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsTimestampUnitSecond, SinkListResponseSchemaFieldsTimestampUnitMillisecond, SinkListResponseSchemaFieldsTimestampUnitMicrosecond, SinkListResponseSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsJson struct {
	Type        SinkListResponseSchemaFieldsJsonType `json:"type,required"`
	MetadataKey string                               `json:"metadata_key,nullable"`
	Name        string                               `json:"name"`
	Required    bool                                 `json:"required"`
	SqlName     string                               `json:"sql_name"`
	JSON        sinkListResponseSchemaFieldsJsonJSON `json:"-"`
}

// sinkListResponseSchemaFieldsJsonJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsJson]
type sinkListResponseSchemaFieldsJsonJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsJson) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsJsonType string

const (
	SinkListResponseSchemaFieldsJsonTypeJson SinkListResponseSchemaFieldsJsonType = "json"
)

func (r SinkListResponseSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsStruct struct {
	JSON sinkListResponseSchemaFieldsStructJSON `json:"-"`
}

// sinkListResponseSchemaFieldsStructJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsStruct]
type sinkListResponseSchemaFieldsStructJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsStruct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsStructJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsStruct) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsList struct {
	JSON sinkListResponseSchemaFieldsListJSON `json:"-"`
}

// sinkListResponseSchemaFieldsListJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFieldsList]
type sinkListResponseSchemaFieldsListJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseSchemaFieldsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFieldsListJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFieldsList) implementsSinkListResponseSchemaField() {}

type SinkListResponseSchemaFieldsType string

const (
	SinkListResponseSchemaFieldsTypeInt32     SinkListResponseSchemaFieldsType = "int32"
	SinkListResponseSchemaFieldsTypeInt64     SinkListResponseSchemaFieldsType = "int64"
	SinkListResponseSchemaFieldsTypeFloat32   SinkListResponseSchemaFieldsType = "float32"
	SinkListResponseSchemaFieldsTypeFloat64   SinkListResponseSchemaFieldsType = "float64"
	SinkListResponseSchemaFieldsTypeBool      SinkListResponseSchemaFieldsType = "bool"
	SinkListResponseSchemaFieldsTypeString    SinkListResponseSchemaFieldsType = "string"
	SinkListResponseSchemaFieldsTypeBinary    SinkListResponseSchemaFieldsType = "binary"
	SinkListResponseSchemaFieldsTypeTimestamp SinkListResponseSchemaFieldsType = "timestamp"
	SinkListResponseSchemaFieldsTypeJson      SinkListResponseSchemaFieldsType = "json"
	SinkListResponseSchemaFieldsTypeStruct    SinkListResponseSchemaFieldsType = "struct"
	SinkListResponseSchemaFieldsTypeList      SinkListResponseSchemaFieldsType = "list"
)

func (r SinkListResponseSchemaFieldsType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsTypeInt32, SinkListResponseSchemaFieldsTypeInt64, SinkListResponseSchemaFieldsTypeFloat32, SinkListResponseSchemaFieldsTypeFloat64, SinkListResponseSchemaFieldsTypeBool, SinkListResponseSchemaFieldsTypeString, SinkListResponseSchemaFieldsTypeBinary, SinkListResponseSchemaFieldsTypeTimestamp, SinkListResponseSchemaFieldsTypeJson, SinkListResponseSchemaFieldsTypeStruct, SinkListResponseSchemaFieldsTypeList:
		return true
	}
	return false
}

type SinkListResponseSchemaFieldsUnit string

const (
	SinkListResponseSchemaFieldsUnitSecond      SinkListResponseSchemaFieldsUnit = "second"
	SinkListResponseSchemaFieldsUnitMillisecond SinkListResponseSchemaFieldsUnit = "millisecond"
	SinkListResponseSchemaFieldsUnitMicrosecond SinkListResponseSchemaFieldsUnit = "microsecond"
	SinkListResponseSchemaFieldsUnitNanosecond  SinkListResponseSchemaFieldsUnit = "nanosecond"
)

func (r SinkListResponseSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFieldsUnitSecond, SinkListResponseSchemaFieldsUnitMillisecond, SinkListResponseSchemaFieldsUnitMicrosecond, SinkListResponseSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type SinkListResponseSchemaFormat struct {
	Type            SinkListResponseSchemaFormatType            `json:"type,required"`
	Compression     SinkListResponseSchemaFormatCompression     `json:"compression"`
	DecimalEncoding SinkListResponseSchemaFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                       `json:"row_group_bytes,nullable"`
	TimestampFormat SinkListResponseSchemaFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                        `json:"unstructured"`
	JSON            sinkListResponseSchemaFormatJSON            `json:"-"`
	union           SinkListResponseSchemaFormatUnion
}

// sinkListResponseSchemaFormatJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFormat]
type sinkListResponseSchemaFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r sinkListResponseSchemaFormatJSON) RawJSON() string {
	return r.raw
}

func (r *SinkListResponseSchemaFormat) UnmarshalJSON(data []byte) (err error) {
	*r = SinkListResponseSchemaFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkListResponseSchemaFormatUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SinkListResponseSchemaFormatJson],
// [SinkListResponseSchemaFormatParquet].
func (r SinkListResponseSchemaFormat) AsUnion() SinkListResponseSchemaFormatUnion {
	return r.union
}

// Union satisfied by [SinkListResponseSchemaFormatJson] or
// [SinkListResponseSchemaFormatParquet].
type SinkListResponseSchemaFormatUnion interface {
	implementsSinkListResponseSchemaFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkListResponseSchemaFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseSchemaFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkListResponseSchemaFormatJson struct {
	Type            SinkListResponseSchemaFormatJsonType            `json:"type,required"`
	DecimalEncoding SinkListResponseSchemaFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkListResponseSchemaFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                            `json:"unstructured"`
	JSON            sinkListResponseSchemaFormatJsonJSON            `json:"-"`
}

// sinkListResponseSchemaFormatJsonJSON contains the JSON metadata for the struct
// [SinkListResponseSchemaFormatJson]
type sinkListResponseSchemaFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkListResponseSchemaFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFormatJson) implementsSinkListResponseSchemaFormat() {}

type SinkListResponseSchemaFormatJsonType string

const (
	SinkListResponseSchemaFormatJsonTypeJson SinkListResponseSchemaFormatJsonType = "json"
)

func (r SinkListResponseSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatJsonDecimalEncoding string

const (
	SinkListResponseSchemaFormatJsonDecimalEncodingNumber SinkListResponseSchemaFormatJsonDecimalEncoding = "number"
	SinkListResponseSchemaFormatJsonDecimalEncodingString SinkListResponseSchemaFormatJsonDecimalEncoding = "string"
	SinkListResponseSchemaFormatJsonDecimalEncodingBytes  SinkListResponseSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r SinkListResponseSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatJsonDecimalEncodingNumber, SinkListResponseSchemaFormatJsonDecimalEncodingString, SinkListResponseSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatJsonTimestampFormat string

const (
	SinkListResponseSchemaFormatJsonTimestampFormatRfc3339    SinkListResponseSchemaFormatJsonTimestampFormat = "rfc3339"
	SinkListResponseSchemaFormatJsonTimestampFormatUnixMillis SinkListResponseSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkListResponseSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatJsonTimestampFormatRfc3339, SinkListResponseSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatParquet struct {
	Type          SinkListResponseSchemaFormatParquetType        `json:"type,required"`
	Compression   SinkListResponseSchemaFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                          `json:"row_group_bytes,nullable"`
	JSON          sinkListResponseSchemaFormatParquetJSON        `json:"-"`
}

// sinkListResponseSchemaFormatParquetJSON contains the JSON metadata for the
// struct [SinkListResponseSchemaFormatParquet]
type sinkListResponseSchemaFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkListResponseSchemaFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseSchemaFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseSchemaFormatParquet) implementsSinkListResponseSchemaFormat() {}

type SinkListResponseSchemaFormatParquetType string

const (
	SinkListResponseSchemaFormatParquetTypeParquet SinkListResponseSchemaFormatParquetType = "parquet"
)

func (r SinkListResponseSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatParquetCompression string

const (
	SinkListResponseSchemaFormatParquetCompressionUncompressed SinkListResponseSchemaFormatParquetCompression = "uncompressed"
	SinkListResponseSchemaFormatParquetCompressionSnappy       SinkListResponseSchemaFormatParquetCompression = "snappy"
	SinkListResponseSchemaFormatParquetCompressionGzip         SinkListResponseSchemaFormatParquetCompression = "gzip"
	SinkListResponseSchemaFormatParquetCompressionZstd         SinkListResponseSchemaFormatParquetCompression = "zstd"
	SinkListResponseSchemaFormatParquetCompressionLz4          SinkListResponseSchemaFormatParquetCompression = "lz4"
)

func (r SinkListResponseSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatParquetCompressionUncompressed, SinkListResponseSchemaFormatParquetCompressionSnappy, SinkListResponseSchemaFormatParquetCompressionGzip, SinkListResponseSchemaFormatParquetCompressionZstd, SinkListResponseSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatType string

const (
	SinkListResponseSchemaFormatTypeJson    SinkListResponseSchemaFormatType = "json"
	SinkListResponseSchemaFormatTypeParquet SinkListResponseSchemaFormatType = "parquet"
)

func (r SinkListResponseSchemaFormatType) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatTypeJson, SinkListResponseSchemaFormatTypeParquet:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatCompression string

const (
	SinkListResponseSchemaFormatCompressionUncompressed SinkListResponseSchemaFormatCompression = "uncompressed"
	SinkListResponseSchemaFormatCompressionSnappy       SinkListResponseSchemaFormatCompression = "snappy"
	SinkListResponseSchemaFormatCompressionGzip         SinkListResponseSchemaFormatCompression = "gzip"
	SinkListResponseSchemaFormatCompressionZstd         SinkListResponseSchemaFormatCompression = "zstd"
	SinkListResponseSchemaFormatCompressionLz4          SinkListResponseSchemaFormatCompression = "lz4"
)

func (r SinkListResponseSchemaFormatCompression) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatCompressionUncompressed, SinkListResponseSchemaFormatCompressionSnappy, SinkListResponseSchemaFormatCompressionGzip, SinkListResponseSchemaFormatCompressionZstd, SinkListResponseSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatDecimalEncoding string

const (
	SinkListResponseSchemaFormatDecimalEncodingNumber SinkListResponseSchemaFormatDecimalEncoding = "number"
	SinkListResponseSchemaFormatDecimalEncodingString SinkListResponseSchemaFormatDecimalEncoding = "string"
	SinkListResponseSchemaFormatDecimalEncodingBytes  SinkListResponseSchemaFormatDecimalEncoding = "bytes"
)

func (r SinkListResponseSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatDecimalEncodingNumber, SinkListResponseSchemaFormatDecimalEncodingString, SinkListResponseSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkListResponseSchemaFormatTimestampFormat string

const (
	SinkListResponseSchemaFormatTimestampFormatRfc3339    SinkListResponseSchemaFormatTimestampFormat = "rfc3339"
	SinkListResponseSchemaFormatTimestampFormatUnixMillis SinkListResponseSchemaFormatTimestampFormat = "unix_millis"
)

func (r SinkListResponseSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkListResponseSchemaFormatTimestampFormatRfc3339, SinkListResponseSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkGetResponse struct {
	// Indicates a unique identifier for this sink.
	ID         string    `json:"id,required"`
	CreatedAt  time.Time `json:"created_at,required" format:"date-time"`
	ModifiedAt time.Time `json:"modified_at,required" format:"date-time"`
	// Defines the name of the Sink.
	Name string `json:"name,required"`
	// Specifies the type of sink.
	Type SinkGetResponseType `json:"type,required"`
	// Defines the configuration of the R2 Sink.
	Config SinkGetResponseConfig `json:"config"`
	Format SinkGetResponseFormat `json:"format"`
	Schema SinkGetResponseSchema `json:"schema"`
	JSON   sinkGetResponseJSON   `json:"-"`
}

// sinkGetResponseJSON contains the JSON metadata for the struct [SinkGetResponse]
type sinkGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Config      apijson.Field
	Format      apijson.Field
	Schema      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of sink.
type SinkGetResponseType string

const (
	SinkGetResponseTypeR2            SinkGetResponseType = "r2"
	SinkGetResponseTypeR2DataCatalog SinkGetResponseType = "r2_data_catalog"
)

func (r SinkGetResponseType) IsKnown() bool {
	switch r {
	case SinkGetResponseTypeR2, SinkGetResponseTypeR2DataCatalog:
		return true
	}
	return false
}

// Defines the configuration of the R2 Sink.
type SinkGetResponseConfig struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket,required"`
	// Authentication token
	Token string `json:"token" format:"var-str"`
	// This field can have the runtime type of
	// [SinkGetResponseConfigCloudflarePipelinesR2TableCredentials].
	Credentials interface{} `json:"credentials"`
	// This field can have the runtime type of
	// [SinkGetResponseConfigCloudflarePipelinesR2TableFileNaming].
	FileNaming interface{} `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Table namespace
	Namespace string `json:"namespace"`
	// This field can have the runtime type of
	// [SinkGetResponseConfigCloudflarePipelinesR2TablePartitioning].
	Partitioning interface{} `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// This field can have the runtime type of
	// [SinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicy],
	// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy].
	RollingPolicy interface{} `json:"rolling_policy"`
	// Table name
	TableName string                    `json:"table_name"`
	JSON      sinkGetResponseConfigJSON `json:"-"`
	union     SinkGetResponseConfigUnion
}

// sinkGetResponseConfigJSON contains the JSON metadata for the struct
// [SinkGetResponseConfig]
type sinkGetResponseConfigJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	Token         apijson.Field
	Credentials   apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Namespace     apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	TableName     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r sinkGetResponseConfigJSON) RawJSON() string {
	return r.raw
}

func (r *SinkGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	*r = SinkGetResponseConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkGetResponseConfigUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [SinkGetResponseConfigCloudflarePipelinesR2Table],
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTable].
func (r SinkGetResponseConfig) AsUnion() SinkGetResponseConfigUnion {
	return r.union
}

// Defines the configuration of the R2 Sink.
//
// Union satisfied by [SinkGetResponseConfigCloudflarePipelinesR2Table] or
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTable].
type SinkGetResponseConfigUnion interface {
	implementsSinkGetResponseConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkGetResponseConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkGetResponseConfigCloudflarePipelinesR2Table{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTable{}),
		},
	)
}

type SinkGetResponseConfigCloudflarePipelinesR2Table struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket      string                                                     `json:"bucket,required"`
	Credentials SinkGetResponseConfigCloudflarePipelinesR2TableCredentials `json:"credentials,required"`
	// Controls filename prefix/suffix and strategy.
	FileNaming SinkGetResponseConfigCloudflarePipelinesR2TableFileNaming `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Data-layout partitioning for sinks.
	Partitioning SinkGetResponseConfigCloudflarePipelinesR2TablePartitioning `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicy `json:"rolling_policy"`
	JSON          sinkGetResponseConfigCloudflarePipelinesR2TableJSON          `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TableJSON contains the JSON metadata
// for the struct [SinkGetResponseConfigCloudflarePipelinesR2Table]
type sinkGetResponseConfigCloudflarePipelinesR2TableJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	Credentials   apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2Table) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TableJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseConfigCloudflarePipelinesR2Table) implementsSinkGetResponseConfig() {}

type SinkGetResponseConfigCloudflarePipelinesR2TableCredentials struct {
	// Cloudflare Account ID for the bucket
	AccessKeyID string `json:"access_key_id,required" format:"var-str"`
	// Cloudflare Account ID for the bucket
	SecretAccessKey string                                                         `json:"secret_access_key,required" format:"var-str"`
	JSON            sinkGetResponseConfigCloudflarePipelinesR2TableCredentialsJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TableCredentialsJSON contains the JSON
// metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2TableCredentials]
type sinkGetResponseConfigCloudflarePipelinesR2TableCredentialsJSON struct {
	AccessKeyID     apijson.Field
	SecretAccessKey apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TableCredentials) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TableCredentialsJSON) RawJSON() string {
	return r.raw
}

// Controls filename prefix/suffix and strategy.
type SinkGetResponseConfigCloudflarePipelinesR2TableFileNaming struct {
	// The prefix to use in file name. i.e prefix-<uuid>.parquet
	Prefix string `json:"prefix"`
	// Filename generation strategy.
	Strategy SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategy `json:"strategy"`
	// This will overwrite the default file suffix. i.e .parquet, use with caution
	Suffix string                                                        `json:"suffix"`
	JSON   sinkGetResponseConfigCloudflarePipelinesR2TableFileNamingJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TableFileNamingJSON contains the JSON
// metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2TableFileNaming]
type sinkGetResponseConfigCloudflarePipelinesR2TableFileNamingJSON struct {
	Prefix      apijson.Field
	Strategy    apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TableFileNaming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TableFileNamingJSON) RawJSON() string {
	return r.raw
}

// Filename generation strategy.
type SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategy string

const (
	SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategySerial SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "serial"
	SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUID   SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid"
	SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7 SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid_v7"
	SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUlid   SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategy = "ulid"
)

func (r SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategy) IsKnown() bool {
	switch r {
	case SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategySerial, SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUID, SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7, SinkGetResponseConfigCloudflarePipelinesR2TableFileNamingStrategyUlid:
		return true
	}
	return false
}

// Data-layout partitioning for sinks.
type SinkGetResponseConfigCloudflarePipelinesR2TablePartitioning struct {
	// The pattern of the date string
	TimePattern string                                                          `json:"time_pattern"`
	JSON        sinkGetResponseConfigCloudflarePipelinesR2TablePartitioningJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TablePartitioningJSON contains the
// JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2TablePartitioning]
type sinkGetResponseConfigCloudflarePipelinesR2TablePartitioningJSON struct {
	TimePattern apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TablePartitioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TablePartitioningJSON) RawJSON() string {
	return r.raw
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                            `json:"interval_seconds"`
	JSON            sinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON contains the
// JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicy]
type sinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TableRollingPolicyJSON) RawJSON() string {
	return r.raw
}

// R2 Data Catalog Sink
type SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTable struct {
	// Authentication token
	Token string `json:"token,required" format:"var-str"`
	// Cloudflare Account ID
	AccountID string `json:"account_id,required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket string `json:"bucket,required"`
	// Table name
	TableName string `json:"table_name,required"`
	// Table namespace
	Namespace string `json:"namespace"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy `json:"rolling_policy"`
	JSON          sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableJSON          `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableJSON contains the JSON
// metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTable]
type sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableJSON struct {
	Token         apijson.Field
	AccountID     apijson.Field
	Bucket        apijson.Field
	TableName     apijson.Field
	Namespace     apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTable) implementsSinkGetResponseConfig() {
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                                       `json:"interval_seconds"`
	JSON            sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON
// contains the JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy]
type sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTableRollingPolicyJSON) RawJSON() string {
	return r.raw
}

type SinkGetResponseFormat struct {
	Type            SinkGetResponseFormatType            `json:"type,required"`
	Compression     SinkGetResponseFormatCompression     `json:"compression"`
	DecimalEncoding SinkGetResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                `json:"row_group_bytes,nullable"`
	TimestampFormat SinkGetResponseFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                 `json:"unstructured"`
	JSON            sinkGetResponseFormatJSON            `json:"-"`
	union           SinkGetResponseFormatUnion
}

// sinkGetResponseFormatJSON contains the JSON metadata for the struct
// [SinkGetResponseFormat]
type sinkGetResponseFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r sinkGetResponseFormatJSON) RawJSON() string {
	return r.raw
}

func (r *SinkGetResponseFormat) UnmarshalJSON(data []byte) (err error) {
	*r = SinkGetResponseFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkGetResponseFormatUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [SinkGetResponseFormatJson],
// [SinkGetResponseFormatParquet].
func (r SinkGetResponseFormat) AsUnion() SinkGetResponseFormatUnion {
	return r.union
}

// Union satisfied by [SinkGetResponseFormatJson] or
// [SinkGetResponseFormatParquet].
type SinkGetResponseFormatUnion interface {
	implementsSinkGetResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkGetResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkGetResponseFormatJson struct {
	Type            SinkGetResponseFormatJsonType            `json:"type,required"`
	DecimalEncoding SinkGetResponseFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkGetResponseFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                     `json:"unstructured"`
	JSON            sinkGetResponseFormatJsonJSON            `json:"-"`
}

// sinkGetResponseFormatJsonJSON contains the JSON metadata for the struct
// [SinkGetResponseFormatJson]
type sinkGetResponseFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkGetResponseFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseFormatJson) implementsSinkGetResponseFormat() {}

type SinkGetResponseFormatJsonType string

const (
	SinkGetResponseFormatJsonTypeJson SinkGetResponseFormatJsonType = "json"
)

func (r SinkGetResponseFormatJsonType) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkGetResponseFormatJsonDecimalEncoding string

const (
	SinkGetResponseFormatJsonDecimalEncodingNumber SinkGetResponseFormatJsonDecimalEncoding = "number"
	SinkGetResponseFormatJsonDecimalEncodingString SinkGetResponseFormatJsonDecimalEncoding = "string"
	SinkGetResponseFormatJsonDecimalEncodingBytes  SinkGetResponseFormatJsonDecimalEncoding = "bytes"
)

func (r SinkGetResponseFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatJsonDecimalEncodingNumber, SinkGetResponseFormatJsonDecimalEncodingString, SinkGetResponseFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkGetResponseFormatJsonTimestampFormat string

const (
	SinkGetResponseFormatJsonTimestampFormatRfc3339    SinkGetResponseFormatJsonTimestampFormat = "rfc3339"
	SinkGetResponseFormatJsonTimestampFormatUnixMillis SinkGetResponseFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkGetResponseFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatJsonTimestampFormatRfc3339, SinkGetResponseFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkGetResponseFormatParquet struct {
	Type          SinkGetResponseFormatParquetType        `json:"type,required"`
	Compression   SinkGetResponseFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                   `json:"row_group_bytes,nullable"`
	JSON          sinkGetResponseFormatParquetJSON        `json:"-"`
}

// sinkGetResponseFormatParquetJSON contains the JSON metadata for the struct
// [SinkGetResponseFormatParquet]
type sinkGetResponseFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkGetResponseFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseFormatParquet) implementsSinkGetResponseFormat() {}

type SinkGetResponseFormatParquetType string

const (
	SinkGetResponseFormatParquetTypeParquet SinkGetResponseFormatParquetType = "parquet"
)

func (r SinkGetResponseFormatParquetType) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkGetResponseFormatParquetCompression string

const (
	SinkGetResponseFormatParquetCompressionUncompressed SinkGetResponseFormatParquetCompression = "uncompressed"
	SinkGetResponseFormatParquetCompressionSnappy       SinkGetResponseFormatParquetCompression = "snappy"
	SinkGetResponseFormatParquetCompressionGzip         SinkGetResponseFormatParquetCompression = "gzip"
	SinkGetResponseFormatParquetCompressionZstd         SinkGetResponseFormatParquetCompression = "zstd"
	SinkGetResponseFormatParquetCompressionLz4          SinkGetResponseFormatParquetCompression = "lz4"
)

func (r SinkGetResponseFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatParquetCompressionUncompressed, SinkGetResponseFormatParquetCompressionSnappy, SinkGetResponseFormatParquetCompressionGzip, SinkGetResponseFormatParquetCompressionZstd, SinkGetResponseFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkGetResponseFormatType string

const (
	SinkGetResponseFormatTypeJson    SinkGetResponseFormatType = "json"
	SinkGetResponseFormatTypeParquet SinkGetResponseFormatType = "parquet"
)

func (r SinkGetResponseFormatType) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatTypeJson, SinkGetResponseFormatTypeParquet:
		return true
	}
	return false
}

type SinkGetResponseFormatCompression string

const (
	SinkGetResponseFormatCompressionUncompressed SinkGetResponseFormatCompression = "uncompressed"
	SinkGetResponseFormatCompressionSnappy       SinkGetResponseFormatCompression = "snappy"
	SinkGetResponseFormatCompressionGzip         SinkGetResponseFormatCompression = "gzip"
	SinkGetResponseFormatCompressionZstd         SinkGetResponseFormatCompression = "zstd"
	SinkGetResponseFormatCompressionLz4          SinkGetResponseFormatCompression = "lz4"
)

func (r SinkGetResponseFormatCompression) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCompressionUncompressed, SinkGetResponseFormatCompressionSnappy, SinkGetResponseFormatCompressionGzip, SinkGetResponseFormatCompressionZstd, SinkGetResponseFormatCompressionLz4:
		return true
	}
	return false
}

type SinkGetResponseFormatDecimalEncoding string

const (
	SinkGetResponseFormatDecimalEncodingNumber SinkGetResponseFormatDecimalEncoding = "number"
	SinkGetResponseFormatDecimalEncodingString SinkGetResponseFormatDecimalEncoding = "string"
	SinkGetResponseFormatDecimalEncodingBytes  SinkGetResponseFormatDecimalEncoding = "bytes"
)

func (r SinkGetResponseFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatDecimalEncodingNumber, SinkGetResponseFormatDecimalEncodingString, SinkGetResponseFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkGetResponseFormatTimestampFormat string

const (
	SinkGetResponseFormatTimestampFormatRfc3339    SinkGetResponseFormatTimestampFormat = "rfc3339"
	SinkGetResponseFormatTimestampFormatUnixMillis SinkGetResponseFormatTimestampFormat = "unix_millis"
)

func (r SinkGetResponseFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatTimestampFormatRfc3339, SinkGetResponseFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkGetResponseSchema struct {
	Fields   []SinkGetResponseSchemaField `json:"fields"`
	Format   SinkGetResponseSchemaFormat  `json:"format"`
	Inferred bool                         `json:"inferred,nullable"`
	JSON     sinkGetResponseSchemaJSON    `json:"-"`
}

// sinkGetResponseSchemaJSON contains the JSON metadata for the struct
// [SinkGetResponseSchema]
type sinkGetResponseSchemaJSON struct {
	Fields      apijson.Field
	Format      apijson.Field
	Inferred    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type SinkGetResponseSchemaField struct {
	Type        SinkGetResponseSchemaFieldsType `json:"type,required"`
	MetadataKey string                          `json:"metadata_key,nullable"`
	Name        string                          `json:"name"`
	Required    bool                            `json:"required"`
	SqlName     string                          `json:"sql_name"`
	Unit        SinkGetResponseSchemaFieldsUnit `json:"unit"`
	JSON        sinkGetResponseSchemaFieldJSON  `json:"-"`
	union       SinkGetResponseSchemaFieldsUnion
}

// sinkGetResponseSchemaFieldJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaField]
type sinkGetResponseSchemaFieldJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r sinkGetResponseSchemaFieldJSON) RawJSON() string {
	return r.raw
}

func (r *SinkGetResponseSchemaField) UnmarshalJSON(data []byte) (err error) {
	*r = SinkGetResponseSchemaField{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkGetResponseSchemaFieldsUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SinkGetResponseSchemaFieldsInt32],
// [SinkGetResponseSchemaFieldsInt64], [SinkGetResponseSchemaFieldsFloat32],
// [SinkGetResponseSchemaFieldsFloat64], [SinkGetResponseSchemaFieldsBool],
// [SinkGetResponseSchemaFieldsString], [SinkGetResponseSchemaFieldsBinary],
// [SinkGetResponseSchemaFieldsTimestamp], [SinkGetResponseSchemaFieldsJson],
// [SinkGetResponseSchemaFieldsStruct], [SinkGetResponseSchemaFieldsList].
func (r SinkGetResponseSchemaField) AsUnion() SinkGetResponseSchemaFieldsUnion {
	return r.union
}

// Union satisfied by [SinkGetResponseSchemaFieldsInt32],
// [SinkGetResponseSchemaFieldsInt64], [SinkGetResponseSchemaFieldsFloat32],
// [SinkGetResponseSchemaFieldsFloat64], [SinkGetResponseSchemaFieldsBool],
// [SinkGetResponseSchemaFieldsString], [SinkGetResponseSchemaFieldsBinary],
// [SinkGetResponseSchemaFieldsTimestamp], [SinkGetResponseSchemaFieldsJson],
// [SinkGetResponseSchemaFieldsStruct] or [SinkGetResponseSchemaFieldsList].
type SinkGetResponseSchemaFieldsUnion interface {
	implementsSinkGetResponseSchemaField()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkGetResponseSchemaFieldsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsInt32{}),
			DiscriminatorValue: "int32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsInt64{}),
			DiscriminatorValue: "int64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsFloat32{}),
			DiscriminatorValue: "float32",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsFloat64{}),
			DiscriminatorValue: "float64",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsBool{}),
			DiscriminatorValue: "bool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsString{}),
			DiscriminatorValue: "string",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsBinary{}),
			DiscriminatorValue: "binary",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsTimestamp{}),
			DiscriminatorValue: "timestamp",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsStruct{}),
			DiscriminatorValue: "struct",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFieldsList{}),
			DiscriminatorValue: "list",
		},
	)
}

type SinkGetResponseSchemaFieldsInt32 struct {
	Type        SinkGetResponseSchemaFieldsInt32Type `json:"type,required"`
	MetadataKey string                               `json:"metadata_key,nullable"`
	Name        string                               `json:"name"`
	Required    bool                                 `json:"required"`
	SqlName     string                               `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsInt32JSON `json:"-"`
}

// sinkGetResponseSchemaFieldsInt32JSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsInt32]
type sinkGetResponseSchemaFieldsInt32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsInt32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsInt32JSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsInt32) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsInt32Type string

const (
	SinkGetResponseSchemaFieldsInt32TypeInt32 SinkGetResponseSchemaFieldsInt32Type = "int32"
)

func (r SinkGetResponseSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsInt64 struct {
	Type        SinkGetResponseSchemaFieldsInt64Type `json:"type,required"`
	MetadataKey string                               `json:"metadata_key,nullable"`
	Name        string                               `json:"name"`
	Required    bool                                 `json:"required"`
	SqlName     string                               `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsInt64JSON `json:"-"`
}

// sinkGetResponseSchemaFieldsInt64JSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsInt64]
type sinkGetResponseSchemaFieldsInt64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsInt64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsInt64JSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsInt64) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsInt64Type string

const (
	SinkGetResponseSchemaFieldsInt64TypeInt64 SinkGetResponseSchemaFieldsInt64Type = "int64"
)

func (r SinkGetResponseSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsFloat32 struct {
	Type        SinkGetResponseSchemaFieldsFloat32Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsFloat32JSON `json:"-"`
}

// sinkGetResponseSchemaFieldsFloat32JSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsFloat32]
type sinkGetResponseSchemaFieldsFloat32JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsFloat32) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsFloat32JSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsFloat32) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsFloat32Type string

const (
	SinkGetResponseSchemaFieldsFloat32TypeFloat32 SinkGetResponseSchemaFieldsFloat32Type = "float32"
)

func (r SinkGetResponseSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsFloat64 struct {
	Type        SinkGetResponseSchemaFieldsFloat64Type `json:"type,required"`
	MetadataKey string                                 `json:"metadata_key,nullable"`
	Name        string                                 `json:"name"`
	Required    bool                                   `json:"required"`
	SqlName     string                                 `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsFloat64JSON `json:"-"`
}

// sinkGetResponseSchemaFieldsFloat64JSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsFloat64]
type sinkGetResponseSchemaFieldsFloat64JSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsFloat64) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsFloat64JSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsFloat64) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsFloat64Type string

const (
	SinkGetResponseSchemaFieldsFloat64TypeFloat64 SinkGetResponseSchemaFieldsFloat64Type = "float64"
)

func (r SinkGetResponseSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsBool struct {
	Type        SinkGetResponseSchemaFieldsBoolType `json:"type,required"`
	MetadataKey string                              `json:"metadata_key,nullable"`
	Name        string                              `json:"name"`
	Required    bool                                `json:"required"`
	SqlName     string                              `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsBoolJSON `json:"-"`
}

// sinkGetResponseSchemaFieldsBoolJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsBool]
type sinkGetResponseSchemaFieldsBoolJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsBool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsBoolJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsBool) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsBoolType string

const (
	SinkGetResponseSchemaFieldsBoolTypeBool SinkGetResponseSchemaFieldsBoolType = "bool"
)

func (r SinkGetResponseSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsString struct {
	Type        SinkGetResponseSchemaFieldsStringType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsStringJSON `json:"-"`
}

// sinkGetResponseSchemaFieldsStringJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsString]
type sinkGetResponseSchemaFieldsStringJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsStringJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsString) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsStringType string

const (
	SinkGetResponseSchemaFieldsStringTypeString SinkGetResponseSchemaFieldsStringType = "string"
)

func (r SinkGetResponseSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsBinary struct {
	Type        SinkGetResponseSchemaFieldsBinaryType `json:"type,required"`
	MetadataKey string                                `json:"metadata_key,nullable"`
	Name        string                                `json:"name"`
	Required    bool                                  `json:"required"`
	SqlName     string                                `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsBinaryJSON `json:"-"`
}

// sinkGetResponseSchemaFieldsBinaryJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsBinary]
type sinkGetResponseSchemaFieldsBinaryJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsBinary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsBinaryJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsBinary) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsBinaryType string

const (
	SinkGetResponseSchemaFieldsBinaryTypeBinary SinkGetResponseSchemaFieldsBinaryType = "binary"
)

func (r SinkGetResponseSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsTimestamp struct {
	Type        SinkGetResponseSchemaFieldsTimestampType `json:"type,required"`
	MetadataKey string                                   `json:"metadata_key,nullable"`
	Name        string                                   `json:"name"`
	Required    bool                                     `json:"required"`
	SqlName     string                                   `json:"sql_name"`
	Unit        SinkGetResponseSchemaFieldsTimestampUnit `json:"unit"`
	JSON        sinkGetResponseSchemaFieldsTimestampJSON `json:"-"`
}

// sinkGetResponseSchemaFieldsTimestampJSON contains the JSON metadata for the
// struct [SinkGetResponseSchemaFieldsTimestamp]
type sinkGetResponseSchemaFieldsTimestampJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsTimestamp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsTimestampJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsTimestamp) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsTimestampType string

const (
	SinkGetResponseSchemaFieldsTimestampTypeTimestamp SinkGetResponseSchemaFieldsTimestampType = "timestamp"
)

func (r SinkGetResponseSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsTimestampUnit string

const (
	SinkGetResponseSchemaFieldsTimestampUnitSecond      SinkGetResponseSchemaFieldsTimestampUnit = "second"
	SinkGetResponseSchemaFieldsTimestampUnitMillisecond SinkGetResponseSchemaFieldsTimestampUnit = "millisecond"
	SinkGetResponseSchemaFieldsTimestampUnitMicrosecond SinkGetResponseSchemaFieldsTimestampUnit = "microsecond"
	SinkGetResponseSchemaFieldsTimestampUnitNanosecond  SinkGetResponseSchemaFieldsTimestampUnit = "nanosecond"
)

func (r SinkGetResponseSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsTimestampUnitSecond, SinkGetResponseSchemaFieldsTimestampUnitMillisecond, SinkGetResponseSchemaFieldsTimestampUnitMicrosecond, SinkGetResponseSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsJson struct {
	Type        SinkGetResponseSchemaFieldsJsonType `json:"type,required"`
	MetadataKey string                              `json:"metadata_key,nullable"`
	Name        string                              `json:"name"`
	Required    bool                                `json:"required"`
	SqlName     string                              `json:"sql_name"`
	JSON        sinkGetResponseSchemaFieldsJsonJSON `json:"-"`
}

// sinkGetResponseSchemaFieldsJsonJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsJson]
type sinkGetResponseSchemaFieldsJsonJSON struct {
	Type        apijson.Field
	MetadataKey apijson.Field
	Name        apijson.Field
	Required    apijson.Field
	SqlName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsJson) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsJsonType string

const (
	SinkGetResponseSchemaFieldsJsonTypeJson SinkGetResponseSchemaFieldsJsonType = "json"
)

func (r SinkGetResponseSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsStruct struct {
	JSON sinkGetResponseSchemaFieldsStructJSON `json:"-"`
}

// sinkGetResponseSchemaFieldsStructJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsStruct]
type sinkGetResponseSchemaFieldsStructJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsStruct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsStructJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsStruct) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsList struct {
	JSON sinkGetResponseSchemaFieldsListJSON `json:"-"`
}

// sinkGetResponseSchemaFieldsListJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFieldsList]
type sinkGetResponseSchemaFieldsListJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFieldsList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFieldsListJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFieldsList) implementsSinkGetResponseSchemaField() {}

type SinkGetResponseSchemaFieldsType string

const (
	SinkGetResponseSchemaFieldsTypeInt32     SinkGetResponseSchemaFieldsType = "int32"
	SinkGetResponseSchemaFieldsTypeInt64     SinkGetResponseSchemaFieldsType = "int64"
	SinkGetResponseSchemaFieldsTypeFloat32   SinkGetResponseSchemaFieldsType = "float32"
	SinkGetResponseSchemaFieldsTypeFloat64   SinkGetResponseSchemaFieldsType = "float64"
	SinkGetResponseSchemaFieldsTypeBool      SinkGetResponseSchemaFieldsType = "bool"
	SinkGetResponseSchemaFieldsTypeString    SinkGetResponseSchemaFieldsType = "string"
	SinkGetResponseSchemaFieldsTypeBinary    SinkGetResponseSchemaFieldsType = "binary"
	SinkGetResponseSchemaFieldsTypeTimestamp SinkGetResponseSchemaFieldsType = "timestamp"
	SinkGetResponseSchemaFieldsTypeJson      SinkGetResponseSchemaFieldsType = "json"
	SinkGetResponseSchemaFieldsTypeStruct    SinkGetResponseSchemaFieldsType = "struct"
	SinkGetResponseSchemaFieldsTypeList      SinkGetResponseSchemaFieldsType = "list"
)

func (r SinkGetResponseSchemaFieldsType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsTypeInt32, SinkGetResponseSchemaFieldsTypeInt64, SinkGetResponseSchemaFieldsTypeFloat32, SinkGetResponseSchemaFieldsTypeFloat64, SinkGetResponseSchemaFieldsTypeBool, SinkGetResponseSchemaFieldsTypeString, SinkGetResponseSchemaFieldsTypeBinary, SinkGetResponseSchemaFieldsTypeTimestamp, SinkGetResponseSchemaFieldsTypeJson, SinkGetResponseSchemaFieldsTypeStruct, SinkGetResponseSchemaFieldsTypeList:
		return true
	}
	return false
}

type SinkGetResponseSchemaFieldsUnit string

const (
	SinkGetResponseSchemaFieldsUnitSecond      SinkGetResponseSchemaFieldsUnit = "second"
	SinkGetResponseSchemaFieldsUnitMillisecond SinkGetResponseSchemaFieldsUnit = "millisecond"
	SinkGetResponseSchemaFieldsUnitMicrosecond SinkGetResponseSchemaFieldsUnit = "microsecond"
	SinkGetResponseSchemaFieldsUnitNanosecond  SinkGetResponseSchemaFieldsUnit = "nanosecond"
)

func (r SinkGetResponseSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFieldsUnitSecond, SinkGetResponseSchemaFieldsUnitMillisecond, SinkGetResponseSchemaFieldsUnitMicrosecond, SinkGetResponseSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormat struct {
	Type            SinkGetResponseSchemaFormatType            `json:"type,required"`
	Compression     SinkGetResponseSchemaFormatCompression     `json:"compression"`
	DecimalEncoding SinkGetResponseSchemaFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                      `json:"row_group_bytes,nullable"`
	TimestampFormat SinkGetResponseSchemaFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                       `json:"unstructured"`
	JSON            sinkGetResponseSchemaFormatJSON            `json:"-"`
	union           SinkGetResponseSchemaFormatUnion
}

// sinkGetResponseSchemaFormatJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFormat]
type sinkGetResponseSchemaFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	RowGroupBytes   apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r sinkGetResponseSchemaFormatJSON) RawJSON() string {
	return r.raw
}

func (r *SinkGetResponseSchemaFormat) UnmarshalJSON(data []byte) (err error) {
	*r = SinkGetResponseSchemaFormat{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SinkGetResponseSchemaFormatUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [SinkGetResponseSchemaFormatJson],
// [SinkGetResponseSchemaFormatParquet].
func (r SinkGetResponseSchemaFormat) AsUnion() SinkGetResponseSchemaFormatUnion {
	return r.union
}

// Union satisfied by [SinkGetResponseSchemaFormatJson] or
// [SinkGetResponseSchemaFormatParquet].
type SinkGetResponseSchemaFormatUnion interface {
	implementsSinkGetResponseSchemaFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkGetResponseSchemaFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFormatJson{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseSchemaFormatParquet{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkGetResponseSchemaFormatJson struct {
	Type            SinkGetResponseSchemaFormatJsonType            `json:"type,required"`
	DecimalEncoding SinkGetResponseSchemaFormatJsonDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkGetResponseSchemaFormatJsonTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                           `json:"unstructured"`
	JSON            sinkGetResponseSchemaFormatJsonJSON            `json:"-"`
}

// sinkGetResponseSchemaFormatJsonJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFormatJson]
type sinkGetResponseSchemaFormatJsonJSON struct {
	Type            apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFormatJson) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFormatJsonJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFormatJson) implementsSinkGetResponseSchemaFormat() {}

type SinkGetResponseSchemaFormatJsonType string

const (
	SinkGetResponseSchemaFormatJsonTypeJson SinkGetResponseSchemaFormatJsonType = "json"
)

func (r SinkGetResponseSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatJsonDecimalEncoding string

const (
	SinkGetResponseSchemaFormatJsonDecimalEncodingNumber SinkGetResponseSchemaFormatJsonDecimalEncoding = "number"
	SinkGetResponseSchemaFormatJsonDecimalEncodingString SinkGetResponseSchemaFormatJsonDecimalEncoding = "string"
	SinkGetResponseSchemaFormatJsonDecimalEncodingBytes  SinkGetResponseSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r SinkGetResponseSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatJsonDecimalEncodingNumber, SinkGetResponseSchemaFormatJsonDecimalEncodingString, SinkGetResponseSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatJsonTimestampFormat string

const (
	SinkGetResponseSchemaFormatJsonTimestampFormatRfc3339    SinkGetResponseSchemaFormatJsonTimestampFormat = "rfc3339"
	SinkGetResponseSchemaFormatJsonTimestampFormatUnixMillis SinkGetResponseSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkGetResponseSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatJsonTimestampFormatRfc3339, SinkGetResponseSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatParquet struct {
	Type          SinkGetResponseSchemaFormatParquetType        `json:"type,required"`
	Compression   SinkGetResponseSchemaFormatParquetCompression `json:"compression"`
	RowGroupBytes int64                                         `json:"row_group_bytes,nullable"`
	JSON          sinkGetResponseSchemaFormatParquetJSON        `json:"-"`
}

// sinkGetResponseSchemaFormatParquetJSON contains the JSON metadata for the struct
// [SinkGetResponseSchemaFormatParquet]
type sinkGetResponseSchemaFormatParquetJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkGetResponseSchemaFormatParquet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseSchemaFormatParquetJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseSchemaFormatParquet) implementsSinkGetResponseSchemaFormat() {}

type SinkGetResponseSchemaFormatParquetType string

const (
	SinkGetResponseSchemaFormatParquetTypeParquet SinkGetResponseSchemaFormatParquetType = "parquet"
)

func (r SinkGetResponseSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatParquetCompression string

const (
	SinkGetResponseSchemaFormatParquetCompressionUncompressed SinkGetResponseSchemaFormatParquetCompression = "uncompressed"
	SinkGetResponseSchemaFormatParquetCompressionSnappy       SinkGetResponseSchemaFormatParquetCompression = "snappy"
	SinkGetResponseSchemaFormatParquetCompressionGzip         SinkGetResponseSchemaFormatParquetCompression = "gzip"
	SinkGetResponseSchemaFormatParquetCompressionZstd         SinkGetResponseSchemaFormatParquetCompression = "zstd"
	SinkGetResponseSchemaFormatParquetCompressionLz4          SinkGetResponseSchemaFormatParquetCompression = "lz4"
)

func (r SinkGetResponseSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatParquetCompressionUncompressed, SinkGetResponseSchemaFormatParquetCompressionSnappy, SinkGetResponseSchemaFormatParquetCompressionGzip, SinkGetResponseSchemaFormatParquetCompressionZstd, SinkGetResponseSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatType string

const (
	SinkGetResponseSchemaFormatTypeJson    SinkGetResponseSchemaFormatType = "json"
	SinkGetResponseSchemaFormatTypeParquet SinkGetResponseSchemaFormatType = "parquet"
)

func (r SinkGetResponseSchemaFormatType) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatTypeJson, SinkGetResponseSchemaFormatTypeParquet:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatCompression string

const (
	SinkGetResponseSchemaFormatCompressionUncompressed SinkGetResponseSchemaFormatCompression = "uncompressed"
	SinkGetResponseSchemaFormatCompressionSnappy       SinkGetResponseSchemaFormatCompression = "snappy"
	SinkGetResponseSchemaFormatCompressionGzip         SinkGetResponseSchemaFormatCompression = "gzip"
	SinkGetResponseSchemaFormatCompressionZstd         SinkGetResponseSchemaFormatCompression = "zstd"
	SinkGetResponseSchemaFormatCompressionLz4          SinkGetResponseSchemaFormatCompression = "lz4"
)

func (r SinkGetResponseSchemaFormatCompression) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatCompressionUncompressed, SinkGetResponseSchemaFormatCompressionSnappy, SinkGetResponseSchemaFormatCompressionGzip, SinkGetResponseSchemaFormatCompressionZstd, SinkGetResponseSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatDecimalEncoding string

const (
	SinkGetResponseSchemaFormatDecimalEncodingNumber SinkGetResponseSchemaFormatDecimalEncoding = "number"
	SinkGetResponseSchemaFormatDecimalEncodingString SinkGetResponseSchemaFormatDecimalEncoding = "string"
	SinkGetResponseSchemaFormatDecimalEncodingBytes  SinkGetResponseSchemaFormatDecimalEncoding = "bytes"
)

func (r SinkGetResponseSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatDecimalEncodingNumber, SinkGetResponseSchemaFormatDecimalEncodingString, SinkGetResponseSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkGetResponseSchemaFormatTimestampFormat string

const (
	SinkGetResponseSchemaFormatTimestampFormatRfc3339    SinkGetResponseSchemaFormatTimestampFormat = "rfc3339"
	SinkGetResponseSchemaFormatTimestampFormatUnixMillis SinkGetResponseSchemaFormatTimestampFormat = "unix_millis"
)

func (r SinkGetResponseSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkGetResponseSchemaFormatTimestampFormatRfc3339, SinkGetResponseSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Defines the name of the Sink.
	Name param.Field[string] `json:"name,required"`
	// Specifies the type of sink.
	Type param.Field[SinkNewParamsType] `json:"type,required"`
	// Defines the configuration of the R2 Sink.
	Config param.Field[SinkNewParamsConfigUnion] `json:"config"`
	Format param.Field[SinkNewParamsFormatUnion] `json:"format"`
	Schema param.Field[SinkNewParamsSchema]      `json:"schema"`
}

func (r SinkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of sink.
type SinkNewParamsType string

const (
	SinkNewParamsTypeR2            SinkNewParamsType = "r2"
	SinkNewParamsTypeR2DataCatalog SinkNewParamsType = "r2_data_catalog"
)

func (r SinkNewParamsType) IsKnown() bool {
	switch r {
	case SinkNewParamsTypeR2, SinkNewParamsTypeR2DataCatalog:
		return true
	}
	return false
}

// Defines the configuration of the R2 Sink.
type SinkNewParamsConfig struct {
	// Cloudflare Account ID for the bucket
	AccountID param.Field[string] `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket param.Field[string] `json:"bucket,required"`
	// Authentication token
	Token       param.Field[string]      `json:"token" format:"var-str"`
	Credentials param.Field[interface{}] `json:"credentials"`
	FileNaming  param.Field[interface{}] `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Table namespace
	Namespace    param.Field[string]      `json:"namespace"`
	Partitioning param.Field[interface{}] `json:"partitioning"`
	// Subpath within the bucket to write to
	Path          param.Field[string]      `json:"path"`
	RollingPolicy param.Field[interface{}] `json:"rolling_policy"`
	// Table name
	TableName param.Field[string] `json:"table_name"`
}

func (r SinkNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsConfig) implementsSinkNewParamsConfigUnion() {}

// Defines the configuration of the R2 Sink.
//
// Satisfied by [pipelines.SinkNewParamsConfigCloudflarePipelinesR2Table],
// [pipelines.SinkNewParamsConfigCloudflarePipelinesR2DataCatalogTable],
// [SinkNewParamsConfig].
type SinkNewParamsConfigUnion interface {
	implementsSinkNewParamsConfigUnion()
}

type SinkNewParamsConfigCloudflarePipelinesR2Table struct {
	// Cloudflare Account ID for the bucket
	AccountID param.Field[string] `json:"account_id,required"`
	// R2 Bucket to write to
	Bucket      param.Field[string]                                                   `json:"bucket,required"`
	Credentials param.Field[SinkNewParamsConfigCloudflarePipelinesR2TableCredentials] `json:"credentials,required"`
	// Controls filename prefix/suffix and strategy.
	FileNaming param.Field[SinkNewParamsConfigCloudflarePipelinesR2TableFileNaming] `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction param.Field[string] `json:"jurisdiction"`
	// Data-layout partitioning for sinks.
	Partitioning param.Field[SinkNewParamsConfigCloudflarePipelinesR2TablePartitioning] `json:"partitioning"`
	// Subpath within the bucket to write to
	Path param.Field[string] `json:"path"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy param.Field[SinkNewParamsConfigCloudflarePipelinesR2TableRollingPolicy] `json:"rolling_policy"`
}

func (r SinkNewParamsConfigCloudflarePipelinesR2Table) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsConfigCloudflarePipelinesR2Table) implementsSinkNewParamsConfigUnion() {}

type SinkNewParamsConfigCloudflarePipelinesR2TableCredentials struct {
	// Cloudflare Account ID for the bucket
	AccessKeyID param.Field[string] `json:"access_key_id,required" format:"var-str"`
	// Cloudflare Account ID for the bucket
	SecretAccessKey param.Field[string] `json:"secret_access_key,required" format:"var-str"`
}

func (r SinkNewParamsConfigCloudflarePipelinesR2TableCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls filename prefix/suffix and strategy.
type SinkNewParamsConfigCloudflarePipelinesR2TableFileNaming struct {
	// The prefix to use in file name. i.e prefix-<uuid>.parquet
	Prefix param.Field[string] `json:"prefix"`
	// Filename generation strategy.
	Strategy param.Field[SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategy] `json:"strategy"`
	// This will overwrite the default file suffix. i.e .parquet, use with caution
	Suffix param.Field[string] `json:"suffix"`
}

func (r SinkNewParamsConfigCloudflarePipelinesR2TableFileNaming) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filename generation strategy.
type SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategy string

const (
	SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategySerial SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategy = "serial"
	SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategyUUID   SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid"
	SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7 SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategy = "uuid_v7"
	SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategyUlid   SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategy = "ulid"
)

func (r SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategy) IsKnown() bool {
	switch r {
	case SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategySerial, SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategyUUID, SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategyUUIDV7, SinkNewParamsConfigCloudflarePipelinesR2TableFileNamingStrategyUlid:
		return true
	}
	return false
}

// Data-layout partitioning for sinks.
type SinkNewParamsConfigCloudflarePipelinesR2TablePartitioning struct {
	// The pattern of the date string
	TimePattern param.Field[string] `json:"time_pattern"`
}

func (r SinkNewParamsConfigCloudflarePipelinesR2TablePartitioning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkNewParamsConfigCloudflarePipelinesR2TableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes param.Field[int64] `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds param.Field[int64] `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds param.Field[int64] `json:"interval_seconds"`
}

func (r SinkNewParamsConfigCloudflarePipelinesR2TableRollingPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 Data Catalog Sink
type SinkNewParamsConfigCloudflarePipelinesR2DataCatalogTable struct {
	// Authentication token
	Token param.Field[string] `json:"token,required" format:"var-str"`
	// Cloudflare Account ID
	AccountID param.Field[string] `json:"account_id,required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket param.Field[string] `json:"bucket,required"`
	// Table name
	TableName param.Field[string] `json:"table_name,required"`
	// Table namespace
	Namespace param.Field[string] `json:"namespace"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy param.Field[SinkNewParamsConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy] `json:"rolling_policy"`
}

func (r SinkNewParamsConfigCloudflarePipelinesR2DataCatalogTable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsConfigCloudflarePipelinesR2DataCatalogTable) implementsSinkNewParamsConfigUnion() {
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkNewParamsConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes param.Field[int64] `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds param.Field[int64] `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds param.Field[int64] `json:"interval_seconds"`
}

func (r SinkNewParamsConfigCloudflarePipelinesR2DataCatalogTableRollingPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SinkNewParamsFormat struct {
	Type            param.Field[SinkNewParamsFormatType]            `json:"type,required"`
	Compression     param.Field[SinkNewParamsFormatCompression]     `json:"compression"`
	DecimalEncoding param.Field[SinkNewParamsFormatDecimalEncoding] `json:"decimal_encoding"`
	RowGroupBytes   param.Field[int64]                              `json:"row_group_bytes"`
	TimestampFormat param.Field[SinkNewParamsFormatTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                               `json:"unstructured"`
}

func (r SinkNewParamsFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsFormat) implementsSinkNewParamsFormatUnion() {}

// Satisfied by [pipelines.SinkNewParamsFormatJson],
// [pipelines.SinkNewParamsFormatParquet], [SinkNewParamsFormat].
type SinkNewParamsFormatUnion interface {
	implementsSinkNewParamsFormatUnion()
}

type SinkNewParamsFormatJson struct {
	Type            param.Field[SinkNewParamsFormatJsonType]            `json:"type,required"`
	DecimalEncoding param.Field[SinkNewParamsFormatJsonDecimalEncoding] `json:"decimal_encoding"`
	TimestampFormat param.Field[SinkNewParamsFormatJsonTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                   `json:"unstructured"`
}

func (r SinkNewParamsFormatJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsFormatJson) implementsSinkNewParamsFormatUnion() {}

type SinkNewParamsFormatJsonType string

const (
	SinkNewParamsFormatJsonTypeJson SinkNewParamsFormatJsonType = "json"
)

func (r SinkNewParamsFormatJsonType) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkNewParamsFormatJsonDecimalEncoding string

const (
	SinkNewParamsFormatJsonDecimalEncodingNumber SinkNewParamsFormatJsonDecimalEncoding = "number"
	SinkNewParamsFormatJsonDecimalEncodingString SinkNewParamsFormatJsonDecimalEncoding = "string"
	SinkNewParamsFormatJsonDecimalEncodingBytes  SinkNewParamsFormatJsonDecimalEncoding = "bytes"
)

func (r SinkNewParamsFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatJsonDecimalEncodingNumber, SinkNewParamsFormatJsonDecimalEncodingString, SinkNewParamsFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewParamsFormatJsonTimestampFormat string

const (
	SinkNewParamsFormatJsonTimestampFormatRfc3339    SinkNewParamsFormatJsonTimestampFormat = "rfc3339"
	SinkNewParamsFormatJsonTimestampFormatUnixMillis SinkNewParamsFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkNewParamsFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatJsonTimestampFormatRfc3339, SinkNewParamsFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewParamsFormatParquet struct {
	Type          param.Field[SinkNewParamsFormatParquetType]        `json:"type,required"`
	Compression   param.Field[SinkNewParamsFormatParquetCompression] `json:"compression"`
	RowGroupBytes param.Field[int64]                                 `json:"row_group_bytes"`
}

func (r SinkNewParamsFormatParquet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsFormatParquet) implementsSinkNewParamsFormatUnion() {}

type SinkNewParamsFormatParquetType string

const (
	SinkNewParamsFormatParquetTypeParquet SinkNewParamsFormatParquetType = "parquet"
)

func (r SinkNewParamsFormatParquetType) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkNewParamsFormatParquetCompression string

const (
	SinkNewParamsFormatParquetCompressionUncompressed SinkNewParamsFormatParquetCompression = "uncompressed"
	SinkNewParamsFormatParquetCompressionSnappy       SinkNewParamsFormatParquetCompression = "snappy"
	SinkNewParamsFormatParquetCompressionGzip         SinkNewParamsFormatParquetCompression = "gzip"
	SinkNewParamsFormatParquetCompressionZstd         SinkNewParamsFormatParquetCompression = "zstd"
	SinkNewParamsFormatParquetCompressionLz4          SinkNewParamsFormatParquetCompression = "lz4"
)

func (r SinkNewParamsFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatParquetCompressionUncompressed, SinkNewParamsFormatParquetCompressionSnappy, SinkNewParamsFormatParquetCompressionGzip, SinkNewParamsFormatParquetCompressionZstd, SinkNewParamsFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkNewParamsFormatType string

const (
	SinkNewParamsFormatTypeJson    SinkNewParamsFormatType = "json"
	SinkNewParamsFormatTypeParquet SinkNewParamsFormatType = "parquet"
)

func (r SinkNewParamsFormatType) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatTypeJson, SinkNewParamsFormatTypeParquet:
		return true
	}
	return false
}

type SinkNewParamsFormatCompression string

const (
	SinkNewParamsFormatCompressionUncompressed SinkNewParamsFormatCompression = "uncompressed"
	SinkNewParamsFormatCompressionSnappy       SinkNewParamsFormatCompression = "snappy"
	SinkNewParamsFormatCompressionGzip         SinkNewParamsFormatCompression = "gzip"
	SinkNewParamsFormatCompressionZstd         SinkNewParamsFormatCompression = "zstd"
	SinkNewParamsFormatCompressionLz4          SinkNewParamsFormatCompression = "lz4"
)

func (r SinkNewParamsFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCompressionUncompressed, SinkNewParamsFormatCompressionSnappy, SinkNewParamsFormatCompressionGzip, SinkNewParamsFormatCompressionZstd, SinkNewParamsFormatCompressionLz4:
		return true
	}
	return false
}

type SinkNewParamsFormatDecimalEncoding string

const (
	SinkNewParamsFormatDecimalEncodingNumber SinkNewParamsFormatDecimalEncoding = "number"
	SinkNewParamsFormatDecimalEncodingString SinkNewParamsFormatDecimalEncoding = "string"
	SinkNewParamsFormatDecimalEncodingBytes  SinkNewParamsFormatDecimalEncoding = "bytes"
)

func (r SinkNewParamsFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatDecimalEncodingNumber, SinkNewParamsFormatDecimalEncodingString, SinkNewParamsFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewParamsFormatTimestampFormat string

const (
	SinkNewParamsFormatTimestampFormatRfc3339    SinkNewParamsFormatTimestampFormat = "rfc3339"
	SinkNewParamsFormatTimestampFormatUnixMillis SinkNewParamsFormatTimestampFormat = "unix_millis"
)

func (r SinkNewParamsFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatTimestampFormatRfc3339, SinkNewParamsFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewParamsSchema struct {
	Fields   param.Field[[]SinkNewParamsSchemaFieldUnion] `json:"fields"`
	Format   param.Field[SinkNewParamsSchemaFormatUnion]  `json:"format"`
	Inferred param.Field[bool]                            `json:"inferred"`
}

func (r SinkNewParamsSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SinkNewParamsSchemaField struct {
	Type        param.Field[SinkNewParamsSchemaFieldsType] `json:"type,required"`
	MetadataKey param.Field[string]                        `json:"metadata_key"`
	Name        param.Field[string]                        `json:"name"`
	Required    param.Field[bool]                          `json:"required"`
	SqlName     param.Field[string]                        `json:"sql_name"`
	Unit        param.Field[SinkNewParamsSchemaFieldsUnit] `json:"unit"`
}

func (r SinkNewParamsSchemaField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaField) implementsSinkNewParamsSchemaFieldUnion() {}

// Satisfied by [pipelines.SinkNewParamsSchemaFieldsInt32],
// [pipelines.SinkNewParamsSchemaFieldsInt64],
// [pipelines.SinkNewParamsSchemaFieldsFloat32],
// [pipelines.SinkNewParamsSchemaFieldsFloat64],
// [pipelines.SinkNewParamsSchemaFieldsBool],
// [pipelines.SinkNewParamsSchemaFieldsString],
// [pipelines.SinkNewParamsSchemaFieldsBinary],
// [pipelines.SinkNewParamsSchemaFieldsTimestamp],
// [pipelines.SinkNewParamsSchemaFieldsJson],
// [pipelines.SinkNewParamsSchemaFieldsStruct],
// [pipelines.SinkNewParamsSchemaFieldsList], [SinkNewParamsSchemaField].
type SinkNewParamsSchemaFieldUnion interface {
	implementsSinkNewParamsSchemaFieldUnion()
}

type SinkNewParamsSchemaFieldsInt32 struct {
	Type        param.Field[SinkNewParamsSchemaFieldsInt32Type] `json:"type,required"`
	MetadataKey param.Field[string]                             `json:"metadata_key"`
	Name        param.Field[string]                             `json:"name"`
	Required    param.Field[bool]                               `json:"required"`
	SqlName     param.Field[string]                             `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsInt32) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsInt32) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsInt32Type string

const (
	SinkNewParamsSchemaFieldsInt32TypeInt32 SinkNewParamsSchemaFieldsInt32Type = "int32"
)

func (r SinkNewParamsSchemaFieldsInt32Type) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsInt32TypeInt32:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsInt64 struct {
	Type        param.Field[SinkNewParamsSchemaFieldsInt64Type] `json:"type,required"`
	MetadataKey param.Field[string]                             `json:"metadata_key"`
	Name        param.Field[string]                             `json:"name"`
	Required    param.Field[bool]                               `json:"required"`
	SqlName     param.Field[string]                             `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsInt64) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsInt64) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsInt64Type string

const (
	SinkNewParamsSchemaFieldsInt64TypeInt64 SinkNewParamsSchemaFieldsInt64Type = "int64"
)

func (r SinkNewParamsSchemaFieldsInt64Type) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsInt64TypeInt64:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsFloat32 struct {
	Type        param.Field[SinkNewParamsSchemaFieldsFloat32Type] `json:"type,required"`
	MetadataKey param.Field[string]                               `json:"metadata_key"`
	Name        param.Field[string]                               `json:"name"`
	Required    param.Field[bool]                                 `json:"required"`
	SqlName     param.Field[string]                               `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsFloat32) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsFloat32) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsFloat32Type string

const (
	SinkNewParamsSchemaFieldsFloat32TypeFloat32 SinkNewParamsSchemaFieldsFloat32Type = "float32"
)

func (r SinkNewParamsSchemaFieldsFloat32Type) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsFloat32TypeFloat32:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsFloat64 struct {
	Type        param.Field[SinkNewParamsSchemaFieldsFloat64Type] `json:"type,required"`
	MetadataKey param.Field[string]                               `json:"metadata_key"`
	Name        param.Field[string]                               `json:"name"`
	Required    param.Field[bool]                                 `json:"required"`
	SqlName     param.Field[string]                               `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsFloat64) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsFloat64) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsFloat64Type string

const (
	SinkNewParamsSchemaFieldsFloat64TypeFloat64 SinkNewParamsSchemaFieldsFloat64Type = "float64"
)

func (r SinkNewParamsSchemaFieldsFloat64Type) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsFloat64TypeFloat64:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsBool struct {
	Type        param.Field[SinkNewParamsSchemaFieldsBoolType] `json:"type,required"`
	MetadataKey param.Field[string]                            `json:"metadata_key"`
	Name        param.Field[string]                            `json:"name"`
	Required    param.Field[bool]                              `json:"required"`
	SqlName     param.Field[string]                            `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsBool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsBool) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsBoolType string

const (
	SinkNewParamsSchemaFieldsBoolTypeBool SinkNewParamsSchemaFieldsBoolType = "bool"
)

func (r SinkNewParamsSchemaFieldsBoolType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsBoolTypeBool:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsString struct {
	Type        param.Field[SinkNewParamsSchemaFieldsStringType] `json:"type,required"`
	MetadataKey param.Field[string]                              `json:"metadata_key"`
	Name        param.Field[string]                              `json:"name"`
	Required    param.Field[bool]                                `json:"required"`
	SqlName     param.Field[string]                              `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsString) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsStringType string

const (
	SinkNewParamsSchemaFieldsStringTypeString SinkNewParamsSchemaFieldsStringType = "string"
)

func (r SinkNewParamsSchemaFieldsStringType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsStringTypeString:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsBinary struct {
	Type        param.Field[SinkNewParamsSchemaFieldsBinaryType] `json:"type,required"`
	MetadataKey param.Field[string]                              `json:"metadata_key"`
	Name        param.Field[string]                              `json:"name"`
	Required    param.Field[bool]                                `json:"required"`
	SqlName     param.Field[string]                              `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsBinary) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsBinary) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsBinaryType string

const (
	SinkNewParamsSchemaFieldsBinaryTypeBinary SinkNewParamsSchemaFieldsBinaryType = "binary"
)

func (r SinkNewParamsSchemaFieldsBinaryType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsBinaryTypeBinary:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsTimestamp struct {
	Type        param.Field[SinkNewParamsSchemaFieldsTimestampType] `json:"type,required"`
	MetadataKey param.Field[string]                                 `json:"metadata_key"`
	Name        param.Field[string]                                 `json:"name"`
	Required    param.Field[bool]                                   `json:"required"`
	SqlName     param.Field[string]                                 `json:"sql_name"`
	Unit        param.Field[SinkNewParamsSchemaFieldsTimestampUnit] `json:"unit"`
}

func (r SinkNewParamsSchemaFieldsTimestamp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsTimestamp) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsTimestampType string

const (
	SinkNewParamsSchemaFieldsTimestampTypeTimestamp SinkNewParamsSchemaFieldsTimestampType = "timestamp"
)

func (r SinkNewParamsSchemaFieldsTimestampType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsTimestampTypeTimestamp:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsTimestampUnit string

const (
	SinkNewParamsSchemaFieldsTimestampUnitSecond      SinkNewParamsSchemaFieldsTimestampUnit = "second"
	SinkNewParamsSchemaFieldsTimestampUnitMillisecond SinkNewParamsSchemaFieldsTimestampUnit = "millisecond"
	SinkNewParamsSchemaFieldsTimestampUnitMicrosecond SinkNewParamsSchemaFieldsTimestampUnit = "microsecond"
	SinkNewParamsSchemaFieldsTimestampUnitNanosecond  SinkNewParamsSchemaFieldsTimestampUnit = "nanosecond"
)

func (r SinkNewParamsSchemaFieldsTimestampUnit) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsTimestampUnitSecond, SinkNewParamsSchemaFieldsTimestampUnitMillisecond, SinkNewParamsSchemaFieldsTimestampUnitMicrosecond, SinkNewParamsSchemaFieldsTimestampUnitNanosecond:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsJson struct {
	Type        param.Field[SinkNewParamsSchemaFieldsJsonType] `json:"type,required"`
	MetadataKey param.Field[string]                            `json:"metadata_key"`
	Name        param.Field[string]                            `json:"name"`
	Required    param.Field[bool]                              `json:"required"`
	SqlName     param.Field[string]                            `json:"sql_name"`
}

func (r SinkNewParamsSchemaFieldsJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsJson) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsJsonType string

const (
	SinkNewParamsSchemaFieldsJsonTypeJson SinkNewParamsSchemaFieldsJsonType = "json"
)

func (r SinkNewParamsSchemaFieldsJsonType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsJsonTypeJson:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsStruct struct {
}

func (r SinkNewParamsSchemaFieldsStruct) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsStruct) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsList struct {
}

func (r SinkNewParamsSchemaFieldsList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFieldsList) implementsSinkNewParamsSchemaFieldUnion() {}

type SinkNewParamsSchemaFieldsType string

const (
	SinkNewParamsSchemaFieldsTypeInt32     SinkNewParamsSchemaFieldsType = "int32"
	SinkNewParamsSchemaFieldsTypeInt64     SinkNewParamsSchemaFieldsType = "int64"
	SinkNewParamsSchemaFieldsTypeFloat32   SinkNewParamsSchemaFieldsType = "float32"
	SinkNewParamsSchemaFieldsTypeFloat64   SinkNewParamsSchemaFieldsType = "float64"
	SinkNewParamsSchemaFieldsTypeBool      SinkNewParamsSchemaFieldsType = "bool"
	SinkNewParamsSchemaFieldsTypeString    SinkNewParamsSchemaFieldsType = "string"
	SinkNewParamsSchemaFieldsTypeBinary    SinkNewParamsSchemaFieldsType = "binary"
	SinkNewParamsSchemaFieldsTypeTimestamp SinkNewParamsSchemaFieldsType = "timestamp"
	SinkNewParamsSchemaFieldsTypeJson      SinkNewParamsSchemaFieldsType = "json"
	SinkNewParamsSchemaFieldsTypeStruct    SinkNewParamsSchemaFieldsType = "struct"
	SinkNewParamsSchemaFieldsTypeList      SinkNewParamsSchemaFieldsType = "list"
)

func (r SinkNewParamsSchemaFieldsType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsTypeInt32, SinkNewParamsSchemaFieldsTypeInt64, SinkNewParamsSchemaFieldsTypeFloat32, SinkNewParamsSchemaFieldsTypeFloat64, SinkNewParamsSchemaFieldsTypeBool, SinkNewParamsSchemaFieldsTypeString, SinkNewParamsSchemaFieldsTypeBinary, SinkNewParamsSchemaFieldsTypeTimestamp, SinkNewParamsSchemaFieldsTypeJson, SinkNewParamsSchemaFieldsTypeStruct, SinkNewParamsSchemaFieldsTypeList:
		return true
	}
	return false
}

type SinkNewParamsSchemaFieldsUnit string

const (
	SinkNewParamsSchemaFieldsUnitSecond      SinkNewParamsSchemaFieldsUnit = "second"
	SinkNewParamsSchemaFieldsUnitMillisecond SinkNewParamsSchemaFieldsUnit = "millisecond"
	SinkNewParamsSchemaFieldsUnitMicrosecond SinkNewParamsSchemaFieldsUnit = "microsecond"
	SinkNewParamsSchemaFieldsUnitNanosecond  SinkNewParamsSchemaFieldsUnit = "nanosecond"
)

func (r SinkNewParamsSchemaFieldsUnit) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFieldsUnitSecond, SinkNewParamsSchemaFieldsUnitMillisecond, SinkNewParamsSchemaFieldsUnitMicrosecond, SinkNewParamsSchemaFieldsUnitNanosecond:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormat struct {
	Type            param.Field[SinkNewParamsSchemaFormatType]            `json:"type,required"`
	Compression     param.Field[SinkNewParamsSchemaFormatCompression]     `json:"compression"`
	DecimalEncoding param.Field[SinkNewParamsSchemaFormatDecimalEncoding] `json:"decimal_encoding"`
	RowGroupBytes   param.Field[int64]                                    `json:"row_group_bytes"`
	TimestampFormat param.Field[SinkNewParamsSchemaFormatTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                     `json:"unstructured"`
}

func (r SinkNewParamsSchemaFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFormat) implementsSinkNewParamsSchemaFormatUnion() {}

// Satisfied by [pipelines.SinkNewParamsSchemaFormatJson],
// [pipelines.SinkNewParamsSchemaFormatParquet], [SinkNewParamsSchemaFormat].
type SinkNewParamsSchemaFormatUnion interface {
	implementsSinkNewParamsSchemaFormatUnion()
}

type SinkNewParamsSchemaFormatJson struct {
	Type            param.Field[SinkNewParamsSchemaFormatJsonType]            `json:"type,required"`
	DecimalEncoding param.Field[SinkNewParamsSchemaFormatJsonDecimalEncoding] `json:"decimal_encoding"`
	TimestampFormat param.Field[SinkNewParamsSchemaFormatJsonTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                         `json:"unstructured"`
}

func (r SinkNewParamsSchemaFormatJson) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFormatJson) implementsSinkNewParamsSchemaFormatUnion() {}

type SinkNewParamsSchemaFormatJsonType string

const (
	SinkNewParamsSchemaFormatJsonTypeJson SinkNewParamsSchemaFormatJsonType = "json"
)

func (r SinkNewParamsSchemaFormatJsonType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatJsonTypeJson:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatJsonDecimalEncoding string

const (
	SinkNewParamsSchemaFormatJsonDecimalEncodingNumber SinkNewParamsSchemaFormatJsonDecimalEncoding = "number"
	SinkNewParamsSchemaFormatJsonDecimalEncodingString SinkNewParamsSchemaFormatJsonDecimalEncoding = "string"
	SinkNewParamsSchemaFormatJsonDecimalEncodingBytes  SinkNewParamsSchemaFormatJsonDecimalEncoding = "bytes"
)

func (r SinkNewParamsSchemaFormatJsonDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatJsonDecimalEncodingNumber, SinkNewParamsSchemaFormatJsonDecimalEncodingString, SinkNewParamsSchemaFormatJsonDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatJsonTimestampFormat string

const (
	SinkNewParamsSchemaFormatJsonTimestampFormatRfc3339    SinkNewParamsSchemaFormatJsonTimestampFormat = "rfc3339"
	SinkNewParamsSchemaFormatJsonTimestampFormatUnixMillis SinkNewParamsSchemaFormatJsonTimestampFormat = "unix_millis"
)

func (r SinkNewParamsSchemaFormatJsonTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatJsonTimestampFormatRfc3339, SinkNewParamsSchemaFormatJsonTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatParquet struct {
	Type          param.Field[SinkNewParamsSchemaFormatParquetType]        `json:"type,required"`
	Compression   param.Field[SinkNewParamsSchemaFormatParquetCompression] `json:"compression"`
	RowGroupBytes param.Field[int64]                                       `json:"row_group_bytes"`
}

func (r SinkNewParamsSchemaFormatParquet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsSchemaFormatParquet) implementsSinkNewParamsSchemaFormatUnion() {}

type SinkNewParamsSchemaFormatParquetType string

const (
	SinkNewParamsSchemaFormatParquetTypeParquet SinkNewParamsSchemaFormatParquetType = "parquet"
)

func (r SinkNewParamsSchemaFormatParquetType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatParquetTypeParquet:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatParquetCompression string

const (
	SinkNewParamsSchemaFormatParquetCompressionUncompressed SinkNewParamsSchemaFormatParquetCompression = "uncompressed"
	SinkNewParamsSchemaFormatParquetCompressionSnappy       SinkNewParamsSchemaFormatParquetCompression = "snappy"
	SinkNewParamsSchemaFormatParquetCompressionGzip         SinkNewParamsSchemaFormatParquetCompression = "gzip"
	SinkNewParamsSchemaFormatParquetCompressionZstd         SinkNewParamsSchemaFormatParquetCompression = "zstd"
	SinkNewParamsSchemaFormatParquetCompressionLz4          SinkNewParamsSchemaFormatParquetCompression = "lz4"
)

func (r SinkNewParamsSchemaFormatParquetCompression) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatParquetCompressionUncompressed, SinkNewParamsSchemaFormatParquetCompressionSnappy, SinkNewParamsSchemaFormatParquetCompressionGzip, SinkNewParamsSchemaFormatParquetCompressionZstd, SinkNewParamsSchemaFormatParquetCompressionLz4:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatType string

const (
	SinkNewParamsSchemaFormatTypeJson    SinkNewParamsSchemaFormatType = "json"
	SinkNewParamsSchemaFormatTypeParquet SinkNewParamsSchemaFormatType = "parquet"
)

func (r SinkNewParamsSchemaFormatType) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatTypeJson, SinkNewParamsSchemaFormatTypeParquet:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatCompression string

const (
	SinkNewParamsSchemaFormatCompressionUncompressed SinkNewParamsSchemaFormatCompression = "uncompressed"
	SinkNewParamsSchemaFormatCompressionSnappy       SinkNewParamsSchemaFormatCompression = "snappy"
	SinkNewParamsSchemaFormatCompressionGzip         SinkNewParamsSchemaFormatCompression = "gzip"
	SinkNewParamsSchemaFormatCompressionZstd         SinkNewParamsSchemaFormatCompression = "zstd"
	SinkNewParamsSchemaFormatCompressionLz4          SinkNewParamsSchemaFormatCompression = "lz4"
)

func (r SinkNewParamsSchemaFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatCompressionUncompressed, SinkNewParamsSchemaFormatCompressionSnappy, SinkNewParamsSchemaFormatCompressionGzip, SinkNewParamsSchemaFormatCompressionZstd, SinkNewParamsSchemaFormatCompressionLz4:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatDecimalEncoding string

const (
	SinkNewParamsSchemaFormatDecimalEncodingNumber SinkNewParamsSchemaFormatDecimalEncoding = "number"
	SinkNewParamsSchemaFormatDecimalEncodingString SinkNewParamsSchemaFormatDecimalEncoding = "string"
	SinkNewParamsSchemaFormatDecimalEncodingBytes  SinkNewParamsSchemaFormatDecimalEncoding = "bytes"
)

func (r SinkNewParamsSchemaFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatDecimalEncodingNumber, SinkNewParamsSchemaFormatDecimalEncodingString, SinkNewParamsSchemaFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewParamsSchemaFormatTimestampFormat string

const (
	SinkNewParamsSchemaFormatTimestampFormatRfc3339    SinkNewParamsSchemaFormatTimestampFormat = "rfc3339"
	SinkNewParamsSchemaFormatTimestampFormatUnixMillis SinkNewParamsSchemaFormatTimestampFormat = "unix_millis"
)

func (r SinkNewParamsSchemaFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewParamsSchemaFormatTimestampFormatRfc3339, SinkNewParamsSchemaFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewResponseEnvelope struct {
	Result SinkNewResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                        `json:"success,required"`
	JSON    sinkNewResponseEnvelopeJSON `json:"-"`
}

// sinkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkNewResponseEnvelope]
type sinkNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkListParams struct {
	// Specifies the public ID of the account.
	AccountID  param.Field[string]  `path:"account_id,required"`
	Page       param.Field[float64] `query:"page"`
	PerPage    param.Field[float64] `query:"per_page"`
	PipelineID param.Field[string]  `query:"pipeline_id"`
}

// URLQuery serializes [SinkListParams]'s query parameters as `url.Values`.
func (r SinkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SinkDeleteParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Delete sink forcefully, including deleting any dependent pipelines.
	Force param.Field[string] `query:"force"`
}

// URLQuery serializes [SinkDeleteParams]'s query parameters as `url.Values`.
func (r SinkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SinkGetParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SinkGetResponseEnvelope struct {
	Result SinkGetResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                        `json:"success,required"`
	JSON    sinkGetResponseEnvelopeJSON `json:"-"`
}

// sinkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkGetResponseEnvelope]
type sinkGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
