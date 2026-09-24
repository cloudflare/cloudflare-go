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
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/sinks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List/Filter Sinks in Account.
func (r *SinkService) List(ctx context.Context, params SinkListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SinkListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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

// Delete Sink in Account.
func (r *SinkService) Delete(ctx context.Context, sinkID string, body SinkDeleteParams, opts ...option.RequestOption) (res *SinkDeleteResponse, err error) {
	var env SinkDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sinkID == "" {
		err = errors.New("missing required sink_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/sinks/%s", body.AccountID, sinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get Sink Details.
func (r *SinkService) Get(ctx context.Context, sinkID string, query SinkGetParams, opts ...option.RequestOption) (res *SinkGetResponse, err error) {
	var env SinkGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sinkID == "" {
		err = errors.New("missing required sink_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/pipelines/v1/sinks/%s", query.AccountID, sinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SinkNewResponse struct {
	// Indicates a unique identifier for this sink.
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt time.Time `json:"modified_at" api:"required" format:"date-time"`
	// Defines the name of the Sink.
	Name string `json:"name" api:"required"`
	// Specifies the type of sink.
	Type SinkNewResponseType `json:"type" api:"required"`
	// R2 Data Catalog Sink
	Config SinkNewResponseConfig `json:"config"`
	// Defines the output data format of a sink.
	Format SinkNewResponseFormat `json:"format"`
	// Defines the schema of the events in the data stream.
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
	AccountID string `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket" api:"required"`
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
	AccountID string `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket      string                                                     `json:"bucket" api:"required"`
	Credentials SinkNewResponseConfigCloudflarePipelinesR2TableCredentials `json:"credentials" api:"required"`
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
	AccessKeyID string `json:"access_key_id" api:"required" format:"var-str"`
	// Cloudflare Account ID for the bucket
	SecretAccessKey string                                                         `json:"secret_access_key" api:"required" format:"var-str"`
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
	Token string `json:"token" api:"required" format:"var-str"`
	// Cloudflare Account ID
	AccountID string `json:"account_id" api:"required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket string `json:"bucket" api:"required"`
	// Table name
	TableName string `json:"table_name" api:"required"`
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

// Defines the output data format of a sink.
type SinkNewResponseFormat struct {
	Type SinkNewResponseFormatType `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
	Compression     SinkNewResponseFormatCompression     `json:"compression"`
	DecimalEncoding SinkNewResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                `json:"row_group_bytes" api:"nullable"`
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
// Possible runtime types of the union are
// [SinkNewResponseFormatCloudflarePipelinesSinkJsonFormat],
// [SinkNewResponseFormatCloudflarePipelinesSinkParquetFormat].
func (r SinkNewResponseFormat) AsUnion() SinkNewResponseFormatUnion {
	return r.union
}

// Defines the output data format of a sink.
//
// Union satisfied by [SinkNewResponseFormatCloudflarePipelinesSinkJsonFormat] or
// [SinkNewResponseFormatCloudflarePipelinesSinkParquetFormat].
type SinkNewResponseFormatUnion interface {
	implementsSinkNewResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkNewResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseFormatCloudflarePipelinesSinkJsonFormat{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkNewResponseFormatCloudflarePipelinesSinkParquetFormat{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkNewResponseFormatCloudflarePipelinesSinkJsonFormat struct {
	Type SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatType `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
	Compression     SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompression     `json:"compression"`
	DecimalEncoding SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                                                  `json:"unstructured"`
	JSON            sinkNewResponseFormatCloudflarePipelinesSinkJsonFormatJSON            `json:"-"`
}

// sinkNewResponseFormatCloudflarePipelinesSinkJsonFormatJSON contains the JSON
// metadata for the struct [SinkNewResponseFormatCloudflarePipelinesSinkJsonFormat]
type sinkNewResponseFormatCloudflarePipelinesSinkJsonFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkNewResponseFormatCloudflarePipelinesSinkJsonFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseFormatCloudflarePipelinesSinkJsonFormatJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseFormatCloudflarePipelinesSinkJsonFormat) implementsSinkNewResponseFormat() {}

type SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatType string

const (
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTypeJson SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatType = "json"
)

func (r SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatType) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTypeJson:
		return true
	}
	return false
}

// Specifies the compression applied to JSON sink output.
type SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompression string

const (
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompression = "uncompressed"
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompressionGzip         SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompression = "gzip"
)

func (r SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed, SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatCompressionGzip:
		return true
	}
	return false
}

type SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding string

const (
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "number"
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "string"
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes  SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "bytes"
)

func (r SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber, SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString, SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat string

const (
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339    SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "rfc3339"
	SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "unix_millis"
)

func (r SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339, SinkNewResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewResponseFormatCloudflarePipelinesSinkParquetFormat struct {
	Type          SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatType        `json:"type" api:"required"`
	Compression   SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression `json:"compression"`
	RowGroupBytes int64                                                                `json:"row_group_bytes" api:"nullable"`
	JSON          sinkNewResponseFormatCloudflarePipelinesSinkParquetFormatJSON        `json:"-"`
}

// sinkNewResponseFormatCloudflarePipelinesSinkParquetFormatJSON contains the JSON
// metadata for the struct
// [SinkNewResponseFormatCloudflarePipelinesSinkParquetFormat]
type sinkNewResponseFormatCloudflarePipelinesSinkParquetFormatJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkNewResponseFormatCloudflarePipelinesSinkParquetFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkNewResponseFormatCloudflarePipelinesSinkParquetFormatJSON) RawJSON() string {
	return r.raw
}

func (r SinkNewResponseFormatCloudflarePipelinesSinkParquetFormat) implementsSinkNewResponseFormat() {
}

type SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatType string

const (
	SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatTypeParquet SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatType = "parquet"
)

func (r SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatType) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatTypeParquet:
		return true
	}
	return false
}

type SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression string

const (
	SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "uncompressed"
	SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy       SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "snappy"
	SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionGzip         SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "gzip"
	SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionZstd         SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "zstd"
	SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionLz4          SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "lz4"
)

func (r SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed, SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy, SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionGzip, SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionZstd, SinkNewResponseFormatCloudflarePipelinesSinkParquetFormatCompressionLz4:
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

// Specifies the compression applied to JSON sink output.
type SinkNewResponseFormatCompression string

const (
	SinkNewResponseFormatCompressionUncompressed SinkNewResponseFormatCompression = "uncompressed"
	SinkNewResponseFormatCompressionGzip         SinkNewResponseFormatCompression = "gzip"
	SinkNewResponseFormatCompressionSnappy       SinkNewResponseFormatCompression = "snappy"
	SinkNewResponseFormatCompressionZstd         SinkNewResponseFormatCompression = "zstd"
	SinkNewResponseFormatCompressionLz4          SinkNewResponseFormatCompression = "lz4"
)

func (r SinkNewResponseFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewResponseFormatCompressionUncompressed, SinkNewResponseFormatCompressionGzip, SinkNewResponseFormatCompressionSnappy, SinkNewResponseFormatCompressionZstd, SinkNewResponseFormatCompressionLz4:
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

// Defines the schema of the events in the data stream.
type SinkNewResponseSchema struct {
	Fields   []SourceField             `json:"fields"`
	Inferred bool                      `json:"inferred" api:"nullable"`
	JSON     sinkNewResponseSchemaJSON `json:"-"`
}

// sinkNewResponseSchemaJSON contains the JSON metadata for the struct
// [SinkNewResponseSchema]
type sinkNewResponseSchemaJSON struct {
	Fields      apijson.Field
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

type SinkListResponse struct {
	// Indicates a unique identifier for this sink.
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt time.Time `json:"modified_at" api:"required" format:"date-time"`
	// Defines the name of the Sink.
	Name string `json:"name" api:"required"`
	// Specifies the type of sink.
	Type SinkListResponseType `json:"type" api:"required"`
	// Defines the configuration of the R2 Sink.
	Config SinkListResponseConfig `json:"config"`
	// Defines the output data format of a sink.
	Format SinkListResponseFormat `json:"format"`
	// Defines the schema of the events in the data stream.
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
	AccountID string `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket" api:"required"`
	// This field can have the runtime type of
	// [SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNaming].
	FileNaming interface{} `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Table namespace
	Namespace string `json:"namespace"`
	// This field can have the runtime type of
	// [SinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioning].
	Partitioning interface{} `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// This field can have the runtime type of
	// [SinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy],
	// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy].
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
// [SinkListResponseConfigCloudflarePipelinesR2TablePublic],
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublic].
func (r SinkListResponseConfig) AsUnion() SinkListResponseConfigUnion {
	return r.union
}

// Defines the configuration of the R2 Sink.
//
// Union satisfied by [SinkListResponseConfigCloudflarePipelinesR2TablePublic] or
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublic].
type SinkListResponseConfigUnion interface {
	implementsSinkListResponseConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkListResponseConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkListResponseConfigCloudflarePipelinesR2TablePublic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublic{}),
		},
	)
}

// R2 Sink public configuration.
type SinkListResponseConfigCloudflarePipelinesR2TablePublic struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket" api:"required"`
	// Controls filename prefix/suffix and strategy.
	FileNaming SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNaming `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Data-layout partitioning for sinks.
	Partitioning SinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioning `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy `json:"rolling_policy"`
	JSON          sinkListResponseConfigCloudflarePipelinesR2TablePublicJSON          `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TablePublicJSON contains the JSON
// metadata for the struct [SinkListResponseConfigCloudflarePipelinesR2TablePublic]
type sinkListResponseConfigCloudflarePipelinesR2TablePublicJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TablePublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TablePublicJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseConfigCloudflarePipelinesR2TablePublic) implementsSinkListResponseConfig() {}

// Controls filename prefix/suffix and strategy.
type SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNaming struct {
	// The prefix to use in file name. i.e prefix-<uuid>.parquet
	Prefix string `json:"prefix"`
	// Filename generation strategy.
	Strategy SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy `json:"strategy"`
	// This will overwrite the default file suffix. i.e .parquet, use with caution
	Suffix string                                                               `json:"suffix"`
	JSON   sinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON contains
// the JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNaming]
type sinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON struct {
	Prefix      apijson.Field
	Strategy    apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNaming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON) RawJSON() string {
	return r.raw
}

// Filename generation strategy.
type SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy string

const (
	SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategySerial SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "serial"
	SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUID   SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "uuid"
	SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUIDV7 SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "uuid_v7"
	SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUlid   SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "ulid"
)

func (r SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy) IsKnown() bool {
	switch r {
	case SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategySerial, SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUID, SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUIDV7, SinkListResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUlid:
		return true
	}
	return false
}

// Data-layout partitioning for sinks.
type SinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioning struct {
	// The pattern of the date string
	TimePattern string                                                                 `json:"time_pattern"`
	JSON        sinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON contains
// the JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioning]
type sinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON struct {
	TimePattern apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON) RawJSON() string {
	return r.raw
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                                   `json:"interval_seconds"`
	JSON            sinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON contains
// the JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy]
type sinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON) RawJSON() string {
	return r.raw
}

// R2 Data Catalog Sink public configuration.
type SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublic struct {
	// Cloudflare Account ID
	AccountID string `json:"account_id" api:"required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket string `json:"bucket" api:"required"`
	// Table name
	TableName string `json:"table_name" api:"required"`
	// Table namespace
	Namespace string `json:"namespace"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy `json:"rolling_policy"`
	JSON          sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON          `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON contains
// the JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublic]
type sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	TableName     apijson.Field
	Namespace     apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublic) implementsSinkListResponseConfig() {
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                                              `json:"interval_seconds"`
	JSON            sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON `json:"-"`
}

// sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON
// contains the JSON metadata for the struct
// [SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy]
type sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON) RawJSON() string {
	return r.raw
}

// Defines the output data format of a sink.
type SinkListResponseFormat struct {
	Type SinkListResponseFormatType `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
	Compression     SinkListResponseFormatCompression     `json:"compression"`
	DecimalEncoding SinkListResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                 `json:"row_group_bytes" api:"nullable"`
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
// Possible runtime types of the union are
// [SinkListResponseFormatCloudflarePipelinesSinkJsonFormat],
// [SinkListResponseFormatCloudflarePipelinesSinkParquetFormat].
func (r SinkListResponseFormat) AsUnion() SinkListResponseFormatUnion {
	return r.union
}

// Defines the output data format of a sink.
//
// Union satisfied by [SinkListResponseFormatCloudflarePipelinesSinkJsonFormat] or
// [SinkListResponseFormatCloudflarePipelinesSinkParquetFormat].
type SinkListResponseFormatUnion interface {
	implementsSinkListResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkListResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseFormatCloudflarePipelinesSinkJsonFormat{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkListResponseFormatCloudflarePipelinesSinkParquetFormat{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkListResponseFormatCloudflarePipelinesSinkJsonFormat struct {
	Type SinkListResponseFormatCloudflarePipelinesSinkJsonFormatType `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
	Compression     SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompression     `json:"compression"`
	DecimalEncoding SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                                                   `json:"unstructured"`
	JSON            sinkListResponseFormatCloudflarePipelinesSinkJsonFormatJSON            `json:"-"`
}

// sinkListResponseFormatCloudflarePipelinesSinkJsonFormatJSON contains the JSON
// metadata for the struct
// [SinkListResponseFormatCloudflarePipelinesSinkJsonFormat]
type sinkListResponseFormatCloudflarePipelinesSinkJsonFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkListResponseFormatCloudflarePipelinesSinkJsonFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseFormatCloudflarePipelinesSinkJsonFormatJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseFormatCloudflarePipelinesSinkJsonFormat) implementsSinkListResponseFormat() {}

type SinkListResponseFormatCloudflarePipelinesSinkJsonFormatType string

const (
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTypeJson SinkListResponseFormatCloudflarePipelinesSinkJsonFormatType = "json"
)

func (r SinkListResponseFormatCloudflarePipelinesSinkJsonFormatType) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTypeJson:
		return true
	}
	return false
}

// Specifies the compression applied to JSON sink output.
type SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompression string

const (
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompression = "uncompressed"
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompressionGzip         SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompression = "gzip"
)

func (r SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompression) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed, SinkListResponseFormatCloudflarePipelinesSinkJsonFormatCompressionGzip:
		return true
	}
	return false
}

type SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding string

const (
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "number"
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "string"
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes  SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "bytes"
)

func (r SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber, SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString, SinkListResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat string

const (
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339    SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "rfc3339"
	SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "unix_millis"
)

func (r SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339, SinkListResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkListResponseFormatCloudflarePipelinesSinkParquetFormat struct {
	Type          SinkListResponseFormatCloudflarePipelinesSinkParquetFormatType        `json:"type" api:"required"`
	Compression   SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression `json:"compression"`
	RowGroupBytes int64                                                                 `json:"row_group_bytes" api:"nullable"`
	JSON          sinkListResponseFormatCloudflarePipelinesSinkParquetFormatJSON        `json:"-"`
}

// sinkListResponseFormatCloudflarePipelinesSinkParquetFormatJSON contains the JSON
// metadata for the struct
// [SinkListResponseFormatCloudflarePipelinesSinkParquetFormat]
type sinkListResponseFormatCloudflarePipelinesSinkParquetFormatJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkListResponseFormatCloudflarePipelinesSinkParquetFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkListResponseFormatCloudflarePipelinesSinkParquetFormatJSON) RawJSON() string {
	return r.raw
}

func (r SinkListResponseFormatCloudflarePipelinesSinkParquetFormat) implementsSinkListResponseFormat() {
}

type SinkListResponseFormatCloudflarePipelinesSinkParquetFormatType string

const (
	SinkListResponseFormatCloudflarePipelinesSinkParquetFormatTypeParquet SinkListResponseFormatCloudflarePipelinesSinkParquetFormatType = "parquet"
)

func (r SinkListResponseFormatCloudflarePipelinesSinkParquetFormatType) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCloudflarePipelinesSinkParquetFormatTypeParquet:
		return true
	}
	return false
}

type SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression string

const (
	SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "uncompressed"
	SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy       SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "snappy"
	SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionGzip         SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "gzip"
	SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionZstd         SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "zstd"
	SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionLz4          SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "lz4"
)

func (r SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompression) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed, SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy, SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionGzip, SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionZstd, SinkListResponseFormatCloudflarePipelinesSinkParquetFormatCompressionLz4:
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

// Specifies the compression applied to JSON sink output.
type SinkListResponseFormatCompression string

const (
	SinkListResponseFormatCompressionUncompressed SinkListResponseFormatCompression = "uncompressed"
	SinkListResponseFormatCompressionGzip         SinkListResponseFormatCompression = "gzip"
	SinkListResponseFormatCompressionSnappy       SinkListResponseFormatCompression = "snappy"
	SinkListResponseFormatCompressionZstd         SinkListResponseFormatCompression = "zstd"
	SinkListResponseFormatCompressionLz4          SinkListResponseFormatCompression = "lz4"
)

func (r SinkListResponseFormatCompression) IsKnown() bool {
	switch r {
	case SinkListResponseFormatCompressionUncompressed, SinkListResponseFormatCompressionGzip, SinkListResponseFormatCompressionSnappy, SinkListResponseFormatCompressionZstd, SinkListResponseFormatCompressionLz4:
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

// Defines the schema of the events in the data stream.
type SinkListResponseSchema struct {
	Fields   []SourceField              `json:"fields"`
	Inferred bool                       `json:"inferred" api:"nullable"`
	JSON     sinkListResponseSchemaJSON `json:"-"`
}

// sinkListResponseSchemaJSON contains the JSON metadata for the struct
// [SinkListResponseSchema]
type sinkListResponseSchemaJSON struct {
	Fields      apijson.Field
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

type SinkDeleteResponse = interface{}

type SinkGetResponse struct {
	// Indicates a unique identifier for this sink.
	ID         string    `json:"id" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt time.Time `json:"modified_at" api:"required" format:"date-time"`
	// Defines the name of the Sink.
	Name string `json:"name" api:"required"`
	// Specifies the type of sink.
	Type SinkGetResponseType `json:"type" api:"required"`
	// Defines the configuration of the R2 Sink.
	Config SinkGetResponseConfig `json:"config"`
	// Defines the output data format of a sink.
	Format SinkGetResponseFormat `json:"format"`
	// Defines the schema of the events in the data stream.
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
	AccountID string `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket" api:"required"`
	// This field can have the runtime type of
	// [SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNaming].
	FileNaming interface{} `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Table namespace
	Namespace string `json:"namespace"`
	// This field can have the runtime type of
	// [SinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioning].
	Partitioning interface{} `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// This field can have the runtime type of
	// [SinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy],
	// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy].
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
// [SinkGetResponseConfigCloudflarePipelinesR2TablePublic],
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublic].
func (r SinkGetResponseConfig) AsUnion() SinkGetResponseConfigUnion {
	return r.union
}

// Defines the configuration of the R2 Sink.
//
// Union satisfied by [SinkGetResponseConfigCloudflarePipelinesR2TablePublic] or
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublic].
type SinkGetResponseConfigUnion interface {
	implementsSinkGetResponseConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkGetResponseConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkGetResponseConfigCloudflarePipelinesR2TablePublic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublic{}),
		},
	)
}

// R2 Sink public configuration.
type SinkGetResponseConfigCloudflarePipelinesR2TablePublic struct {
	// Cloudflare Account ID for the bucket
	AccountID string `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket string `json:"bucket" api:"required"`
	// Controls filename prefix/suffix and strategy.
	FileNaming SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNaming `json:"file_naming"`
	// Jurisdiction this bucket is hosted in
	Jurisdiction string `json:"jurisdiction"`
	// Data-layout partitioning for sinks.
	Partitioning SinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioning `json:"partitioning"`
	// Subpath within the bucket to write to
	Path string `json:"path"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy `json:"rolling_policy"`
	JSON          sinkGetResponseConfigCloudflarePipelinesR2TablePublicJSON          `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TablePublicJSON contains the JSON
// metadata for the struct [SinkGetResponseConfigCloudflarePipelinesR2TablePublic]
type sinkGetResponseConfigCloudflarePipelinesR2TablePublicJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	FileNaming    apijson.Field
	Jurisdiction  apijson.Field
	Partitioning  apijson.Field
	Path          apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TablePublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TablePublicJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseConfigCloudflarePipelinesR2TablePublic) implementsSinkGetResponseConfig() {}

// Controls filename prefix/suffix and strategy.
type SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNaming struct {
	// The prefix to use in file name. i.e prefix-<uuid>.parquet
	Prefix string `json:"prefix"`
	// Filename generation strategy.
	Strategy SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy `json:"strategy"`
	// This will overwrite the default file suffix. i.e .parquet, use with caution
	Suffix string                                                              `json:"suffix"`
	JSON   sinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON contains the
// JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNaming]
type sinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON struct {
	Prefix      apijson.Field
	Strategy    apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNaming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingJSON) RawJSON() string {
	return r.raw
}

// Filename generation strategy.
type SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy string

const (
	SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategySerial SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "serial"
	SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUID   SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "uuid"
	SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUIDV7 SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "uuid_v7"
	SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUlid   SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy = "ulid"
)

func (r SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategy) IsKnown() bool {
	switch r {
	case SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategySerial, SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUID, SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUUIDV7, SinkGetResponseConfigCloudflarePipelinesR2TablePublicFileNamingStrategyUlid:
		return true
	}
	return false
}

// Data-layout partitioning for sinks.
type SinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioning struct {
	// The pattern of the date string
	TimePattern string                                                                `json:"time_pattern"`
	JSON        sinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON contains
// the JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioning]
type sinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON struct {
	TimePattern apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TablePublicPartitioningJSON) RawJSON() string {
	return r.raw
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                                  `json:"interval_seconds"`
	JSON            sinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON contains
// the JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy]
type sinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2TablePublicRollingPolicyJSON) RawJSON() string {
	return r.raw
}

// R2 Data Catalog Sink public configuration.
type SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublic struct {
	// Cloudflare Account ID
	AccountID string `json:"account_id" api:"required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket string `json:"bucket" api:"required"`
	// Table name
	TableName string `json:"table_name" api:"required"`
	// Table namespace
	Namespace string `json:"namespace"`
	// Rolling policy for file sinks (when & why to close a file and open a new one).
	RollingPolicy SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy `json:"rolling_policy"`
	JSON          sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON          `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON contains
// the JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublic]
type sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON struct {
	AccountID     apijson.Field
	Bucket        apijson.Field
	TableName     apijson.Field
	Namespace     apijson.Field
	RollingPolicy apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublic) implementsSinkGetResponseConfig() {
}

// Rolling policy for file sinks (when & why to close a file and open a new one).
type SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy struct {
	// Files will be rolled after reaching this number of bytes
	FileSizeBytes int64 `json:"file_size_bytes"`
	// Number of seconds of inactivity to wait before rolling over to a new file
	InactivitySeconds int64 `json:"inactivity_seconds"`
	// Number of seconds to wait before rolling over to a new file
	IntervalSeconds int64                                                                             `json:"interval_seconds"`
	JSON            sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON `json:"-"`
}

// sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON
// contains the JSON metadata for the struct
// [SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy]
type sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON struct {
	FileSizeBytes     apijson.Field
	InactivitySeconds apijson.Field
	IntervalSeconds   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseConfigCloudflarePipelinesR2DataCatalogTablePublicRollingPolicyJSON) RawJSON() string {
	return r.raw
}

// Defines the output data format of a sink.
type SinkGetResponseFormat struct {
	Type SinkGetResponseFormatType `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
	Compression     SinkGetResponseFormatCompression     `json:"compression"`
	DecimalEncoding SinkGetResponseFormatDecimalEncoding `json:"decimal_encoding"`
	RowGroupBytes   int64                                `json:"row_group_bytes" api:"nullable"`
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
// Possible runtime types of the union are
// [SinkGetResponseFormatCloudflarePipelinesSinkJsonFormat],
// [SinkGetResponseFormatCloudflarePipelinesSinkParquetFormat].
func (r SinkGetResponseFormat) AsUnion() SinkGetResponseFormatUnion {
	return r.union
}

// Defines the output data format of a sink.
//
// Union satisfied by [SinkGetResponseFormatCloudflarePipelinesSinkJsonFormat] or
// [SinkGetResponseFormatCloudflarePipelinesSinkParquetFormat].
type SinkGetResponseFormatUnion interface {
	implementsSinkGetResponseFormat()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SinkGetResponseFormatUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseFormatCloudflarePipelinesSinkJsonFormat{}),
			DiscriminatorValue: "json",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SinkGetResponseFormatCloudflarePipelinesSinkParquetFormat{}),
			DiscriminatorValue: "parquet",
		},
	)
}

type SinkGetResponseFormatCloudflarePipelinesSinkJsonFormat struct {
	Type SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatType `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
	Compression     SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompression     `json:"compression"`
	DecimalEncoding SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding `json:"decimal_encoding"`
	TimestampFormat SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat `json:"timestamp_format"`
	Unstructured    bool                                                                  `json:"unstructured"`
	JSON            sinkGetResponseFormatCloudflarePipelinesSinkJsonFormatJSON            `json:"-"`
}

// sinkGetResponseFormatCloudflarePipelinesSinkJsonFormatJSON contains the JSON
// metadata for the struct [SinkGetResponseFormatCloudflarePipelinesSinkJsonFormat]
type sinkGetResponseFormatCloudflarePipelinesSinkJsonFormatJSON struct {
	Type            apijson.Field
	Compression     apijson.Field
	DecimalEncoding apijson.Field
	TimestampFormat apijson.Field
	Unstructured    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SinkGetResponseFormatCloudflarePipelinesSinkJsonFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseFormatCloudflarePipelinesSinkJsonFormatJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseFormatCloudflarePipelinesSinkJsonFormat) implementsSinkGetResponseFormat() {}

type SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatType string

const (
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTypeJson SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatType = "json"
)

func (r SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatType) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTypeJson:
		return true
	}
	return false
}

// Specifies the compression applied to JSON sink output.
type SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompression string

const (
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompression = "uncompressed"
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompressionGzip         SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompression = "gzip"
)

func (r SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompression) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed, SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatCompressionGzip:
		return true
	}
	return false
}

type SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding string

const (
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "number"
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "string"
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes  SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "bytes"
)

func (r SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber, SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString, SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat string

const (
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339    SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "rfc3339"
	SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "unix_millis"
)

func (r SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339, SinkGetResponseFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkGetResponseFormatCloudflarePipelinesSinkParquetFormat struct {
	Type          SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatType        `json:"type" api:"required"`
	Compression   SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression `json:"compression"`
	RowGroupBytes int64                                                                `json:"row_group_bytes" api:"nullable"`
	JSON          sinkGetResponseFormatCloudflarePipelinesSinkParquetFormatJSON        `json:"-"`
}

// sinkGetResponseFormatCloudflarePipelinesSinkParquetFormatJSON contains the JSON
// metadata for the struct
// [SinkGetResponseFormatCloudflarePipelinesSinkParquetFormat]
type sinkGetResponseFormatCloudflarePipelinesSinkParquetFormatJSON struct {
	Type          apijson.Field
	Compression   apijson.Field
	RowGroupBytes apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SinkGetResponseFormatCloudflarePipelinesSinkParquetFormat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkGetResponseFormatCloudflarePipelinesSinkParquetFormatJSON) RawJSON() string {
	return r.raw
}

func (r SinkGetResponseFormatCloudflarePipelinesSinkParquetFormat) implementsSinkGetResponseFormat() {
}

type SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatType string

const (
	SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatTypeParquet SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatType = "parquet"
)

func (r SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatType) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatTypeParquet:
		return true
	}
	return false
}

type SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression string

const (
	SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "uncompressed"
	SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy       SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "snappy"
	SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionGzip         SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "gzip"
	SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionZstd         SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "zstd"
	SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionLz4          SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression = "lz4"
)

func (r SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompression) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed, SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy, SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionGzip, SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionZstd, SinkGetResponseFormatCloudflarePipelinesSinkParquetFormatCompressionLz4:
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

// Specifies the compression applied to JSON sink output.
type SinkGetResponseFormatCompression string

const (
	SinkGetResponseFormatCompressionUncompressed SinkGetResponseFormatCompression = "uncompressed"
	SinkGetResponseFormatCompressionGzip         SinkGetResponseFormatCompression = "gzip"
	SinkGetResponseFormatCompressionSnappy       SinkGetResponseFormatCompression = "snappy"
	SinkGetResponseFormatCompressionZstd         SinkGetResponseFormatCompression = "zstd"
	SinkGetResponseFormatCompressionLz4          SinkGetResponseFormatCompression = "lz4"
)

func (r SinkGetResponseFormatCompression) IsKnown() bool {
	switch r {
	case SinkGetResponseFormatCompressionUncompressed, SinkGetResponseFormatCompressionGzip, SinkGetResponseFormatCompressionSnappy, SinkGetResponseFormatCompressionZstd, SinkGetResponseFormatCompressionLz4:
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

// Defines the schema of the events in the data stream.
type SinkGetResponseSchema struct {
	Fields   []SourceField             `json:"fields"`
	Inferred bool                      `json:"inferred" api:"nullable"`
	JSON     sinkGetResponseSchemaJSON `json:"-"`
}

// sinkGetResponseSchemaJSON contains the JSON metadata for the struct
// [SinkGetResponseSchema]
type sinkGetResponseSchemaJSON struct {
	Fields      apijson.Field
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

type SinkNewParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Defines the name of the Sink.
	Name param.Field[string] `json:"name" api:"required"`
	// Specifies the type of sink.
	Type param.Field[SinkNewParamsType] `json:"type" api:"required"`
	// Defines the configuration of the R2 Sink.
	Config param.Field[SinkNewParamsConfigUnion] `json:"config"`
	// Defines the output data format of a sink.
	Format param.Field[SinkNewParamsFormatUnion] `json:"format"`
	// Defines the schema of the events in the data stream.
	Schema param.Field[SinkNewParamsSchema] `json:"schema"`
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
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket param.Field[string] `json:"bucket" api:"required"`
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
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// R2 Bucket to write to
	Bucket      param.Field[string]                                                   `json:"bucket" api:"required"`
	Credentials param.Field[SinkNewParamsConfigCloudflarePipelinesR2TableCredentials] `json:"credentials" api:"required"`
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
	AccessKeyID param.Field[string] `json:"access_key_id" api:"required" format:"var-str"`
	// Cloudflare Account ID for the bucket
	SecretAccessKey param.Field[string] `json:"secret_access_key" api:"required" format:"var-str"`
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
	Token param.Field[string] `json:"token" api:"required" format:"var-str"`
	// Cloudflare Account ID
	AccountID param.Field[string] `json:"account_id" api:"required" format:"uri"`
	// The R2 Bucket that hosts this catalog
	Bucket param.Field[string] `json:"bucket" api:"required"`
	// Table name
	TableName param.Field[string] `json:"table_name" api:"required"`
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

// Defines the output data format of a sink.
type SinkNewParamsFormat struct {
	Type param.Field[SinkNewParamsFormatType] `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
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

// Defines the output data format of a sink.
//
// Satisfied by [pipelines.SinkNewParamsFormatCloudflarePipelinesSinkJsonFormat],
// [pipelines.SinkNewParamsFormatCloudflarePipelinesSinkParquetFormat],
// [SinkNewParamsFormat].
type SinkNewParamsFormatUnion interface {
	implementsSinkNewParamsFormatUnion()
}

type SinkNewParamsFormatCloudflarePipelinesSinkJsonFormat struct {
	Type param.Field[SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatType] `json:"type" api:"required"`
	// Specifies the compression applied to JSON sink output.
	Compression     param.Field[SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompression]     `json:"compression"`
	DecimalEncoding param.Field[SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding] `json:"decimal_encoding"`
	TimestampFormat param.Field[SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormat] `json:"timestamp_format"`
	Unstructured    param.Field[bool]                                                                `json:"unstructured"`
}

func (r SinkNewParamsFormatCloudflarePipelinesSinkJsonFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsFormatCloudflarePipelinesSinkJsonFormat) implementsSinkNewParamsFormatUnion() {}

type SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatType string

const (
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTypeJson SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatType = "json"
)

func (r SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatType) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTypeJson:
		return true
	}
	return false
}

// Specifies the compression applied to JSON sink output.
type SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompression string

const (
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompression = "uncompressed"
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompressionGzip         SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompression = "gzip"
)

func (r SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompressionUncompressed, SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatCompressionGzip:
		return true
	}
	return false
}

type SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding string

const (
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "number"
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "string"
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes  SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding = "bytes"
)

func (r SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncoding) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingNumber, SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingString, SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatDecimalEncodingBytes:
		return true
	}
	return false
}

type SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormat string

const (
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339    SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "rfc3339"
	SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormat = "unix_millis"
)

func (r SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormat) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormatRfc3339, SinkNewParamsFormatCloudflarePipelinesSinkJsonFormatTimestampFormatUnixMillis:
		return true
	}
	return false
}

type SinkNewParamsFormatCloudflarePipelinesSinkParquetFormat struct {
	Type          param.Field[SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatType]        `json:"type" api:"required"`
	Compression   param.Field[SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression] `json:"compression"`
	RowGroupBytes param.Field[int64]                                                              `json:"row_group_bytes"`
}

func (r SinkNewParamsFormatCloudflarePipelinesSinkParquetFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SinkNewParamsFormatCloudflarePipelinesSinkParquetFormat) implementsSinkNewParamsFormatUnion() {
}

type SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatType string

const (
	SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatTypeParquet SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatType = "parquet"
)

func (r SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatType) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatTypeParquet:
		return true
	}
	return false
}

type SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression string

const (
	SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression = "uncompressed"
	SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy       SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression = "snappy"
	SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionGzip         SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression = "gzip"
	SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionZstd         SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression = "zstd"
	SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionLz4          SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression = "lz4"
)

func (r SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionUncompressed, SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionSnappy, SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionGzip, SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionZstd, SinkNewParamsFormatCloudflarePipelinesSinkParquetFormatCompressionLz4:
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

// Specifies the compression applied to JSON sink output.
type SinkNewParamsFormatCompression string

const (
	SinkNewParamsFormatCompressionUncompressed SinkNewParamsFormatCompression = "uncompressed"
	SinkNewParamsFormatCompressionGzip         SinkNewParamsFormatCompression = "gzip"
	SinkNewParamsFormatCompressionSnappy       SinkNewParamsFormatCompression = "snappy"
	SinkNewParamsFormatCompressionZstd         SinkNewParamsFormatCompression = "zstd"
	SinkNewParamsFormatCompressionLz4          SinkNewParamsFormatCompression = "lz4"
)

func (r SinkNewParamsFormatCompression) IsKnown() bool {
	switch r {
	case SinkNewParamsFormatCompressionUncompressed, SinkNewParamsFormatCompressionGzip, SinkNewParamsFormatCompressionSnappy, SinkNewParamsFormatCompressionZstd, SinkNewParamsFormatCompressionLz4:
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

// Defines the schema of the events in the data stream.
type SinkNewParamsSchema struct {
	Fields   param.Field[[]SourceFieldUnionParam] `json:"fields"`
	Inferred param.Field[bool]                    `json:"inferred"`
}

func (r SinkNewParamsSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SinkNewResponseEnvelope struct {
	Result SinkNewResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                        `json:"success" api:"required"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters sinks by name (case-insensitive substring).
	Name       param.Field[string]  `query:"name"`
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
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SinkDeleteResponseEnvelope struct {
	Result SinkDeleteResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                           `json:"success" api:"required"`
	JSON    sinkDeleteResponseEnvelopeJSON `json:"-"`
}

// sinkDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkDeleteResponseEnvelope]
type sinkDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkGetParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SinkGetResponseEnvelope struct {
	Result SinkGetResponse `json:"result" api:"required"`
	// Indicates whether the API call was successful.
	Success bool                        `json:"success" api:"required"`
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
