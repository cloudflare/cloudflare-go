// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pipelines

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/tidwall/gjson"
)

// PipelineService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineService] method instead.
type PipelineService struct {
	Options []option.RequestOption
}

// NewPipelineService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPipelineService(opts ...option.RequestOption) (r *PipelineService) {
	r = &PipelineService{}
	r.Options = opts
	return
}

// Create a new Pipeline.
func (r *PipelineService) New(ctx context.Context, params PipelineNewParams, opts ...option.RequestOption) (res *PipelineNewResponse, err error) {
	var env PipelineNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing Pipeline.
func (r *PipelineService) Update(ctx context.Context, pipelineName string, params PipelineUpdateParams, opts ...option.RequestOption) (res *PipelineUpdateResponse, err error) {
	var env PipelineUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pipelineName == "" {
		err = errors.New("missing required pipeline_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/%s", params.AccountID, pipelineName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, filter, and paginate Pipelines in an account.
func (r *PipelineService) List(ctx context.Context, params PipelineListParams, opts ...option.RequestOption) (res *PipelineListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a Pipeline.
func (r *PipelineService) Delete(ctx context.Context, pipelineName string, body PipelineDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pipelineName == "" {
		err = errors.New("missing required pipeline_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/%s", body.AccountID, pipelineName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get configuration details of a Pipeline.
func (r *PipelineService) Get(ctx context.Context, pipelineName string, query PipelineGetParams, opts ...option.RequestOption) (res *PipelineGetResponse, err error) {
	var env PipelineGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if pipelineName == "" {
		err = errors.New("missing required pipeline_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pipelines/%s", query.AccountID, pipelineName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Describes the configuration of a Pipeline.
type PipelineNewResponse struct {
	// Specifies the Pipeline identifier.
	ID          string                         `json:"id,required"`
	Destination PipelineNewResponseDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of Pipeline.
	Name   string                      `json:"name,required"`
	Source []PipelineNewResponseSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                 `json:"version,required"`
	JSON    pipelineNewResponseJSON `json:"-"`
}

// pipelineNewResponseJSON contains the JSON metadata for the struct
// [PipelineNewResponse]
type pipelineNewResponseJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseDestination struct {
	Batch       PipelineNewResponseDestinationBatch       `json:"batch,required"`
	Compression PipelineNewResponseDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineNewResponseDestinationFormat `json:"format,required"`
	Path   PipelineNewResponseDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineNewResponseDestinationType `json:"type,required"`
	JSON pipelineNewResponseDestinationJSON `json:"-"`
}

// pipelineNewResponseDestinationJSON contains the JSON metadata for the struct
// [PipelineNewResponseDestination]
type pipelineNewResponseDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                   `json:"max_rows,required"`
	JSON    pipelineNewResponseDestinationBatchJSON `json:"-"`
}

// pipelineNewResponseDestinationBatchJSON contains the JSON metadata for the
// struct [PipelineNewResponseDestinationBatch]
type pipelineNewResponseDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineNewResponseDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineNewResponseDestinationCompressionType `json:"type,required"`
	JSON pipelineNewResponseDestinationCompressionJSON `json:"-"`
}

// pipelineNewResponseDestinationCompressionJSON contains the JSON metadata for the
// struct [PipelineNewResponseDestinationCompression]
type pipelineNewResponseDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineNewResponseDestinationCompressionType string

const (
	PipelineNewResponseDestinationCompressionTypeNone    PipelineNewResponseDestinationCompressionType = "none"
	PipelineNewResponseDestinationCompressionTypeGzip    PipelineNewResponseDestinationCompressionType = "gzip"
	PipelineNewResponseDestinationCompressionTypeDeflate PipelineNewResponseDestinationCompressionType = "deflate"
)

func (r PipelineNewResponseDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineNewResponseDestinationCompressionTypeNone, PipelineNewResponseDestinationCompressionTypeGzip, PipelineNewResponseDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineNewResponseDestinationFormat string

const (
	PipelineNewResponseDestinationFormatJson PipelineNewResponseDestinationFormat = "json"
)

func (r PipelineNewResponseDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseDestinationFormatJson:
		return true
	}
	return false
}

type PipelineNewResponseDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                 `json:"prefix"`
	JSON   pipelineNewResponseDestinationPathJSON `json:"-"`
}

// pipelineNewResponseDestinationPathJSON contains the JSON metadata for the struct
// [PipelineNewResponseDestinationPath]
type pipelineNewResponseDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineNewResponseDestinationType string

const (
	PipelineNewResponseDestinationTypeR2 PipelineNewResponseDestinationType = "r2"
)

func (r PipelineNewResponseDestinationType) IsKnown() bool {
	switch r {
	case PipelineNewResponseDestinationTypeR2:
		return true
	}
	return false
}

type PipelineNewResponseSource struct {
	// Specifies the format of source data.
	Format PipelineNewResponseSourceFormat `json:"format,required"`
	Type   string                          `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                   `json:"cors"`
	JSON  pipelineNewResponseSourceJSON `json:"-"`
	union PipelineNewResponseSourceUnion
}

// pipelineNewResponseSourceJSON contains the JSON metadata for the struct
// [PipelineNewResponseSource]
type pipelineNewResponseSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineNewResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineNewResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineNewResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineNewResponseSourceUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [pipelines.PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSource].
func (r PipelineNewResponseSource) AsUnion() PipelineNewResponseSourceUnion {
	return r.union
}

// Union satisfied by
// [pipelines.PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource]
// or
// [pipelines.PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSource].
type PipelineNewResponseSourceUnion interface {
	implementsPipelineNewResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineNewResponseSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                    `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool                                                                    `json:"authentication"`
	CORS           PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON contains
// the JSON metadata for the struct
// [PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource]
type pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource) implementsPipelineNewResponseSource() {
}

// Specifies the format of source data.
type PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                    `json:"origins"`
	JSON    pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                       `json:"type,required"`
	JSON   pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSource]
type pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSource) implementsPipelineNewResponseSource() {
}

// Specifies the format of source data.
type PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineNewResponseSourceFormat string

const (
	PipelineNewResponseSourceFormatJson PipelineNewResponseSourceFormat = "json"
)

func (r PipelineNewResponseSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewResponseSourceFormatJson:
		return true
	}
	return false
}

// Describes the configuration of a Pipeline.
type PipelineUpdateResponse struct {
	// Specifies the Pipeline identifier.
	ID          string                            `json:"id,required"`
	Destination PipelineUpdateResponseDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of Pipeline.
	Name   string                         `json:"name,required"`
	Source []PipelineUpdateResponseSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                    `json:"version,required"`
	JSON    pipelineUpdateResponseJSON `json:"-"`
}

// pipelineUpdateResponseJSON contains the JSON metadata for the struct
// [PipelineUpdateResponse]
type pipelineUpdateResponseJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseDestination struct {
	Batch       PipelineUpdateResponseDestinationBatch       `json:"batch,required"`
	Compression PipelineUpdateResponseDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineUpdateResponseDestinationFormat `json:"format,required"`
	Path   PipelineUpdateResponseDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineUpdateResponseDestinationType `json:"type,required"`
	JSON pipelineUpdateResponseDestinationJSON `json:"-"`
}

// pipelineUpdateResponseDestinationJSON contains the JSON metadata for the struct
// [PipelineUpdateResponseDestination]
type pipelineUpdateResponseDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                      `json:"max_rows,required"`
	JSON    pipelineUpdateResponseDestinationBatchJSON `json:"-"`
}

// pipelineUpdateResponseDestinationBatchJSON contains the JSON metadata for the
// struct [PipelineUpdateResponseDestinationBatch]
type pipelineUpdateResponseDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineUpdateResponseDestinationCompressionType `json:"type,required"`
	JSON pipelineUpdateResponseDestinationCompressionJSON `json:"-"`
}

// pipelineUpdateResponseDestinationCompressionJSON contains the JSON metadata for
// the struct [PipelineUpdateResponseDestinationCompression]
type pipelineUpdateResponseDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineUpdateResponseDestinationCompressionType string

const (
	PipelineUpdateResponseDestinationCompressionTypeNone    PipelineUpdateResponseDestinationCompressionType = "none"
	PipelineUpdateResponseDestinationCompressionTypeGzip    PipelineUpdateResponseDestinationCompressionType = "gzip"
	PipelineUpdateResponseDestinationCompressionTypeDeflate PipelineUpdateResponseDestinationCompressionType = "deflate"
)

func (r PipelineUpdateResponseDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseDestinationCompressionTypeNone, PipelineUpdateResponseDestinationCompressionTypeGzip, PipelineUpdateResponseDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineUpdateResponseDestinationFormat string

const (
	PipelineUpdateResponseDestinationFormatJson PipelineUpdateResponseDestinationFormat = "json"
)

func (r PipelineUpdateResponseDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseDestinationFormatJson:
		return true
	}
	return false
}

type PipelineUpdateResponseDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                    `json:"prefix"`
	JSON   pipelineUpdateResponseDestinationPathJSON `json:"-"`
}

// pipelineUpdateResponseDestinationPathJSON contains the JSON metadata for the
// struct [PipelineUpdateResponseDestinationPath]
type pipelineUpdateResponseDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineUpdateResponseDestinationType string

const (
	PipelineUpdateResponseDestinationTypeR2 PipelineUpdateResponseDestinationType = "r2"
)

func (r PipelineUpdateResponseDestinationType) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseDestinationTypeR2:
		return true
	}
	return false
}

type PipelineUpdateResponseSource struct {
	// Specifies the format of source data.
	Format PipelineUpdateResponseSourceFormat `json:"format,required"`
	Type   string                             `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                      `json:"cors"`
	JSON  pipelineUpdateResponseSourceJSON `json:"-"`
	union PipelineUpdateResponseSourceUnion
}

// pipelineUpdateResponseSourceJSON contains the JSON metadata for the struct
// [PipelineUpdateResponseSource]
type pipelineUpdateResponseSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineUpdateResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineUpdateResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineUpdateResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineUpdateResponseSourceUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [pipelines.PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSource].
func (r PipelineUpdateResponseSource) AsUnion() PipelineUpdateResponseSourceUnion {
	return r.union
}

// Union satisfied by
// [pipelines.PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource]
// or
// [pipelines.PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSource].
type PipelineUpdateResponseSourceUnion interface {
	implementsPipelineUpdateResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineUpdateResponseSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                       `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool                                                                       `json:"authentication"`
	CORS           PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON
// contains the JSON metadata for the struct
// [PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource]
type pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource) implementsPipelineUpdateResponseSource() {
}

// Specifies the format of source data.
type PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                       `json:"origins"`
	JSON    pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                          `json:"type,required"`
	JSON   pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSource]
type pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSource) implementsPipelineUpdateResponseSource() {
}

// Specifies the format of source data.
type PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineUpdateResponseSourceFormat string

const (
	PipelineUpdateResponseSourceFormatJson PipelineUpdateResponseSourceFormat = "json"
)

func (r PipelineUpdateResponseSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateResponseSourceFormatJson:
		return true
	}
	return false
}

type PipelineListResponse struct {
	ResultInfo PipelineListResponseResultInfo `json:"result_info,required"`
	Results    []PipelineListResponseResult   `json:"results,required"`
	// Indicates whether the API call was successful.
	Success bool                     `json:"success,required"`
	JSON    pipelineListResponseJSON `json:"-"`
}

// pipelineListResponseJSON contains the JSON metadata for the struct
// [PipelineListResponse]
type pipelineListResponseJSON struct {
	ResultInfo  apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultInfo struct {
	// Indicates the number of items on current page.
	Count float64 `json:"count,required"`
	// Indicates the current page number.
	Page float64 `json:"page,required"`
	// Indicates the number of items per page.
	PerPage float64 `json:"per_page,required"`
	// Indicates the total number of items.
	TotalCount float64                            `json:"total_count,required"`
	JSON       pipelineListResponseResultInfoJSON `json:"-"`
}

// pipelineListResponseResultInfoJSON contains the JSON metadata for the struct
// [PipelineListResponseResultInfo]
type pipelineListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Describes the configuration of a Pipeline.
type PipelineListResponseResult struct {
	// Specifies the Pipeline identifier.
	ID          string                                 `json:"id,required"`
	Destination PipelineListResponseResultsDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of Pipeline.
	Name   string                              `json:"name,required"`
	Source []PipelineListResponseResultsSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                        `json:"version,required"`
	JSON    pipelineListResponseResultJSON `json:"-"`
}

// pipelineListResponseResultJSON contains the JSON metadata for the struct
// [PipelineListResponseResult]
type pipelineListResponseResultJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsDestination struct {
	Batch       PipelineListResponseResultsDestinationBatch       `json:"batch,required"`
	Compression PipelineListResponseResultsDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineListResponseResultsDestinationFormat `json:"format,required"`
	Path   PipelineListResponseResultsDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineListResponseResultsDestinationType `json:"type,required"`
	JSON pipelineListResponseResultsDestinationJSON `json:"-"`
}

// pipelineListResponseResultsDestinationJSON contains the JSON metadata for the
// struct [PipelineListResponseResultsDestination]
type pipelineListResponseResultsDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                           `json:"max_rows,required"`
	JSON    pipelineListResponseResultsDestinationBatchJSON `json:"-"`
}

// pipelineListResponseResultsDestinationBatchJSON contains the JSON metadata for
// the struct [PipelineListResponseResultsDestinationBatch]
type pipelineListResponseResultsDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineListResponseResultsDestinationCompressionType `json:"type,required"`
	JSON pipelineListResponseResultsDestinationCompressionJSON `json:"-"`
}

// pipelineListResponseResultsDestinationCompressionJSON contains the JSON metadata
// for the struct [PipelineListResponseResultsDestinationCompression]
type pipelineListResponseResultsDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineListResponseResultsDestinationCompressionType string

const (
	PipelineListResponseResultsDestinationCompressionTypeNone    PipelineListResponseResultsDestinationCompressionType = "none"
	PipelineListResponseResultsDestinationCompressionTypeGzip    PipelineListResponseResultsDestinationCompressionType = "gzip"
	PipelineListResponseResultsDestinationCompressionTypeDeflate PipelineListResponseResultsDestinationCompressionType = "deflate"
)

func (r PipelineListResponseResultsDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsDestinationCompressionTypeNone, PipelineListResponseResultsDestinationCompressionTypeGzip, PipelineListResponseResultsDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineListResponseResultsDestinationFormat string

const (
	PipelineListResponseResultsDestinationFormatJson PipelineListResponseResultsDestinationFormat = "json"
)

func (r PipelineListResponseResultsDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsDestinationFormatJson:
		return true
	}
	return false
}

type PipelineListResponseResultsDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                         `json:"prefix"`
	JSON   pipelineListResponseResultsDestinationPathJSON `json:"-"`
}

// pipelineListResponseResultsDestinationPathJSON contains the JSON metadata for
// the struct [PipelineListResponseResultsDestinationPath]
type pipelineListResponseResultsDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineListResponseResultsDestinationType string

const (
	PipelineListResponseResultsDestinationTypeR2 PipelineListResponseResultsDestinationType = "r2"
)

func (r PipelineListResponseResultsDestinationType) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsDestinationTypeR2:
		return true
	}
	return false
}

type PipelineListResponseResultsSource struct {
	// Specifies the format of source data.
	Format PipelineListResponseResultsSourceFormat `json:"format,required"`
	Type   string                                  `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                           `json:"cors"`
	JSON  pipelineListResponseResultsSourceJSON `json:"-"`
	union PipelineListResponseResultsSourceUnion
}

// pipelineListResponseResultsSourceJSON contains the JSON metadata for the struct
// [PipelineListResponseResultsSource]
type pipelineListResponseResultsSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineListResponseResultsSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineListResponseResultsSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineListResponseResultsSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineListResponseResultsSourceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [pipelines.PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSource].
func (r PipelineListResponseResultsSource) AsUnion() PipelineListResponseResultsSourceUnion {
	return r.union
}

// Union satisfied by
// [pipelines.PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSource]
// or
// [pipelines.PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSource].
type PipelineListResponseResultsSourceUnion interface {
	implementsPipelineListResponseResultsSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineListResponseResultsSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                            `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool                                                                            `json:"authentication"`
	CORS           PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON
// contains the JSON metadata for the struct
// [PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSource]
type pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSource) implementsPipelineListResponseResultsSource() {
}

// Specifies the format of source data.
type PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                            `json:"origins"`
	JSON    pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                               `json:"type,required"`
	JSON   pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSource]
type pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSource) implementsPipelineListResponseResultsSource() {
}

// Specifies the format of source data.
type PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineListResponseResultsSourceFormat string

const (
	PipelineListResponseResultsSourceFormatJson PipelineListResponseResultsSourceFormat = "json"
)

func (r PipelineListResponseResultsSourceFormat) IsKnown() bool {
	switch r {
	case PipelineListResponseResultsSourceFormatJson:
		return true
	}
	return false
}

// Describes the configuration of a Pipeline.
type PipelineGetResponse struct {
	// Specifies the Pipeline identifier.
	ID          string                         `json:"id,required"`
	Destination PipelineGetResponseDestination `json:"destination,required"`
	// Indicates the endpoint URL to send traffic.
	Endpoint string `json:"endpoint,required"`
	// Defines the name of Pipeline.
	Name   string                      `json:"name,required"`
	Source []PipelineGetResponseSource `json:"source,required"`
	// Indicates the version number of last saved configuration.
	Version float64                 `json:"version,required"`
	JSON    pipelineGetResponseJSON `json:"-"`
}

// pipelineGetResponseJSON contains the JSON metadata for the struct
// [PipelineGetResponse]
type pipelineGetResponseJSON struct {
	ID          apijson.Field
	Destination apijson.Field
	Endpoint    apijson.Field
	Name        apijson.Field
	Source      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseDestination struct {
	Batch       PipelineGetResponseDestinationBatch       `json:"batch,required"`
	Compression PipelineGetResponseDestinationCompression `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format PipelineGetResponseDestinationFormat `json:"format,required"`
	Path   PipelineGetResponseDestinationPath   `json:"path,required"`
	// Specifies the type of destination.
	Type PipelineGetResponseDestinationType `json:"type,required"`
	JSON pipelineGetResponseDestinationJSON `json:"-"`
}

// pipelineGetResponseDestinationJSON contains the JSON metadata for the struct
// [PipelineGetResponseDestination]
type pipelineGetResponseDestinationJSON struct {
	Batch       apijson.Field
	Compression apijson.Field
	Format      apijson.Field
	Path        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes int64 `json:"max_bytes,required"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS float64 `json:"max_duration_s,required"`
	// Specifies rough maximum number of rows per file.
	MaxRows int64                                   `json:"max_rows,required"`
	JSON    pipelineGetResponseDestinationBatchJSON `json:"-"`
}

// pipelineGetResponseDestinationBatchJSON contains the JSON metadata for the
// struct [PipelineGetResponseDestinationBatch]
type pipelineGetResponseDestinationBatchJSON struct {
	MaxBytes     apijson.Field
	MaxDurationS apijson.Field
	MaxRows      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PipelineGetResponseDestinationBatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationBatchJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type PipelineGetResponseDestinationCompressionType `json:"type,required"`
	JSON pipelineGetResponseDestinationCompressionJSON `json:"-"`
}

// pipelineGetResponseDestinationCompressionJSON contains the JSON metadata for the
// struct [PipelineGetResponseDestinationCompression]
type pipelineGetResponseDestinationCompressionJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseDestinationCompression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationCompressionJSON) RawJSON() string {
	return r.raw
}

// Specifies the desired compression algorithm and format.
type PipelineGetResponseDestinationCompressionType string

const (
	PipelineGetResponseDestinationCompressionTypeNone    PipelineGetResponseDestinationCompressionType = "none"
	PipelineGetResponseDestinationCompressionTypeGzip    PipelineGetResponseDestinationCompressionType = "gzip"
	PipelineGetResponseDestinationCompressionTypeDeflate PipelineGetResponseDestinationCompressionType = "deflate"
)

func (r PipelineGetResponseDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineGetResponseDestinationCompressionTypeNone, PipelineGetResponseDestinationCompressionTypeGzip, PipelineGetResponseDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineGetResponseDestinationFormat string

const (
	PipelineGetResponseDestinationFormatJson PipelineGetResponseDestinationFormat = "json"
)

func (r PipelineGetResponseDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseDestinationFormatJson:
		return true
	}
	return false
}

type PipelineGetResponseDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket string `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename string `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath string `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix string                                 `json:"prefix"`
	JSON   pipelineGetResponseDestinationPathJSON `json:"-"`
}

// pipelineGetResponseDestinationPathJSON contains the JSON metadata for the struct
// [PipelineGetResponseDestinationPath]
type pipelineGetResponseDestinationPathJSON struct {
	Bucket      apijson.Field
	Filename    apijson.Field
	Filepath    apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseDestinationPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseDestinationPathJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of destination.
type PipelineGetResponseDestinationType string

const (
	PipelineGetResponseDestinationTypeR2 PipelineGetResponseDestinationType = "r2"
)

func (r PipelineGetResponseDestinationType) IsKnown() bool {
	switch r {
	case PipelineGetResponseDestinationTypeR2:
		return true
	}
	return false
}

type PipelineGetResponseSource struct {
	// Specifies the format of source data.
	Format PipelineGetResponseSourceFormat `json:"format,required"`
	Type   string                          `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool `json:"authentication"`
	// This field can have the runtime type of
	// [PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS].
	CORS  interface{}                   `json:"cors"`
	JSON  pipelineGetResponseSourceJSON `json:"-"`
	union PipelineGetResponseSourceUnion
}

// pipelineGetResponseSourceJSON contains the JSON metadata for the struct
// [PipelineGetResponseSource]
type pipelineGetResponseSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r pipelineGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *PipelineGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = PipelineGetResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PipelineGetResponseSourceUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [pipelines.PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSource].
func (r PipelineGetResponseSource) AsUnion() PipelineGetResponseSourceUnion {
	return r.union
}

// Union satisfied by
// [pipelines.PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource]
// or
// [pipelines.PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSource].
type PipelineGetResponseSourceUnion interface {
	implementsPipelineGetResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PipelineGetResponseSourceUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSource{}),
		},
	)
}

type PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat `json:"format,required"`
	Type   string                                                                    `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication bool                                                                    `json:"authentication"`
	CORS           PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS `json:"cors"`
	JSON           pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON `json:"-"`
}

// pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON contains
// the JSON metadata for the struct
// [PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource]
type pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON struct {
	Format         apijson.Field
	Type           apijson.Field
	Authentication apijson.Field
	CORS           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSource) implementsPipelineGetResponseSource() {
}

// Specifies the format of source data.
type PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins []string                                                                    `json:"origins"`
	JSON    pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON `json:"-"`
}

// pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON
// contains the JSON metadata for the struct
// [PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS]
type pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON struct {
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORSJSON) RawJSON() string {
	return r.raw
}

type PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat `json:"format,required"`
	Type   string                                                                       `json:"type,required"`
	JSON   pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON   `json:"-"`
}

// pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON
// contains the JSON metadata for the struct
// [PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSource]
type pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON struct {
	Format      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceJSON) RawJSON() string {
	return r.raw
}

func (r PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSource) implementsPipelineGetResponseSource() {
}

// Specifies the format of source data.
type PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineGetResponseSourceFormat string

const (
	PipelineGetResponseSourceFormatJson PipelineGetResponseSourceFormat = "json"
)

func (r PipelineGetResponseSourceFormat) IsKnown() bool {
	switch r {
	case PipelineGetResponseSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewParams struct {
	// Specifies the public ID of the account.
	AccountID   param.Field[string]                       `path:"account_id,required"`
	Destination param.Field[PipelineNewParamsDestination] `json:"destination,required"`
	// Defines the name of Pipeline.
	Name   param.Field[string]                         `json:"name,required"`
	Source param.Field[[]PipelineNewParamsSourceUnion] `json:"source,required"`
}

func (r PipelineNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsDestination struct {
	Batch       param.Field[PipelineNewParamsDestinationBatch]       `json:"batch,required"`
	Compression param.Field[PipelineNewParamsDestinationCompression] `json:"compression,required"`
	Credentials param.Field[PipelineNewParamsDestinationCredentials] `json:"credentials,required"`
	// Specifies the format of data to deliver.
	Format param.Field[PipelineNewParamsDestinationFormat] `json:"format,required"`
	Path   param.Field[PipelineNewParamsDestinationPath]   `json:"path,required"`
	// Specifies the type of destination.
	Type param.Field[PipelineNewParamsDestinationType] `json:"type,required"`
}

func (r PipelineNewParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes param.Field[int64] `json:"max_bytes"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS param.Field[float64] `json:"max_duration_s"`
	// Specifies rough maximum number of rows per file.
	MaxRows param.Field[int64] `json:"max_rows"`
}

func (r PipelineNewParamsDestinationBatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type param.Field[PipelineNewParamsDestinationCompressionType] `json:"type"`
}

func (r PipelineNewParamsDestinationCompression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the desired compression algorithm and format.
type PipelineNewParamsDestinationCompressionType string

const (
	PipelineNewParamsDestinationCompressionTypeNone    PipelineNewParamsDestinationCompressionType = "none"
	PipelineNewParamsDestinationCompressionTypeGzip    PipelineNewParamsDestinationCompressionType = "gzip"
	PipelineNewParamsDestinationCompressionTypeDeflate PipelineNewParamsDestinationCompressionType = "deflate"
)

func (r PipelineNewParamsDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineNewParamsDestinationCompressionTypeNone, PipelineNewParamsDestinationCompressionTypeGzip, PipelineNewParamsDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

type PipelineNewParamsDestinationCredentials struct {
	// Specifies the R2 Bucket Access Key Id.
	AccessKeyID param.Field[string] `json:"access_key_id,required"`
	// Specifies the R2 Endpoint.
	Endpoint param.Field[string] `json:"endpoint,required"`
	// Specifies the R2 Bucket Secret Access Key.
	SecretAccessKey param.Field[string] `json:"secret_access_key,required"`
}

func (r PipelineNewParamsDestinationCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the format of data to deliver.
type PipelineNewParamsDestinationFormat string

const (
	PipelineNewParamsDestinationFormatJson PipelineNewParamsDestinationFormat = "json"
)

func (r PipelineNewParamsDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsDestinationFormatJson:
		return true
	}
	return false
}

type PipelineNewParamsDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket param.Field[string] `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename param.Field[string] `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath param.Field[string] `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix param.Field[string] `json:"prefix"`
}

func (r PipelineNewParamsDestinationPath) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of destination.
type PipelineNewParamsDestinationType string

const (
	PipelineNewParamsDestinationTypeR2 PipelineNewParamsDestinationType = "r2"
)

func (r PipelineNewParamsDestinationType) IsKnown() bool {
	switch r {
	case PipelineNewParamsDestinationTypeR2:
		return true
	}
	return false
}

type PipelineNewParamsSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineNewParamsSourceFormat] `json:"format,required"`
	Type   param.Field[string]                        `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication param.Field[bool]        `json:"authentication"`
	CORS           param.Field[interface{}] `json:"cors"`
}

func (r PipelineNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineNewParamsSource) implementsPipelineNewParamsSourceUnion() {}

// Satisfied by
// [pipelines.PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSource],
// [PipelineNewParamsSource].
type PipelineNewParamsSourceUnion interface {
	implementsPipelineNewParamsSourceUnion()
}

type PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                  `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication param.Field[bool]                                                                  `json:"authentication"`
	CORS           param.Field[PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS] `json:"cors"`
}

func (r PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource) implementsPipelineNewParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins param.Field[[]string] `json:"origins"`
}

func (r PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                     `json:"type,required"`
}

func (r PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSource) implementsPipelineNewParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineNewParamsSourceFormat string

const (
	PipelineNewParamsSourceFormatJson PipelineNewParamsSourceFormat = "json"
)

func (r PipelineNewParamsSourceFormat) IsKnown() bool {
	switch r {
	case PipelineNewParamsSourceFormatJson:
		return true
	}
	return false
}

type PipelineNewResponseEnvelope struct {
	// Describes the configuration of a Pipeline.
	Result PipelineNewResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                            `json:"success,required"`
	JSON    pipelineNewResponseEnvelopeJSON `json:"-"`
}

// pipelineNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PipelineNewResponseEnvelope]
type pipelineNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PipelineUpdateParams struct {
	// Specifies the public ID of the account.
	AccountID   param.Field[string]                          `path:"account_id,required"`
	Destination param.Field[PipelineUpdateParamsDestination] `json:"destination,required"`
	// Defines the name of Pipeline.
	Name   param.Field[string]                            `json:"name,required"`
	Source param.Field[[]PipelineUpdateParamsSourceUnion] `json:"source,required"`
}

func (r PipelineUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsDestination struct {
	Batch       param.Field[PipelineUpdateParamsDestinationBatch]       `json:"batch,required"`
	Compression param.Field[PipelineUpdateParamsDestinationCompression] `json:"compression,required"`
	// Specifies the format of data to deliver.
	Format param.Field[PipelineUpdateParamsDestinationFormat] `json:"format,required"`
	Path   param.Field[PipelineUpdateParamsDestinationPath]   `json:"path,required"`
	// Specifies the type of destination.
	Type        param.Field[PipelineUpdateParamsDestinationType]        `json:"type,required"`
	Credentials param.Field[PipelineUpdateParamsDestinationCredentials] `json:"credentials"`
}

func (r PipelineUpdateParamsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsDestinationBatch struct {
	// Specifies rough maximum size of files.
	MaxBytes param.Field[int64] `json:"max_bytes"`
	// Specifies duration to wait to aggregate batches files.
	MaxDurationS param.Field[float64] `json:"max_duration_s"`
	// Specifies rough maximum number of rows per file.
	MaxRows param.Field[int64] `json:"max_rows"`
}

func (r PipelineUpdateParamsDestinationBatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsDestinationCompression struct {
	// Specifies the desired compression algorithm and format.
	Type param.Field[PipelineUpdateParamsDestinationCompressionType] `json:"type"`
}

func (r PipelineUpdateParamsDestinationCompression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the desired compression algorithm and format.
type PipelineUpdateParamsDestinationCompressionType string

const (
	PipelineUpdateParamsDestinationCompressionTypeNone    PipelineUpdateParamsDestinationCompressionType = "none"
	PipelineUpdateParamsDestinationCompressionTypeGzip    PipelineUpdateParamsDestinationCompressionType = "gzip"
	PipelineUpdateParamsDestinationCompressionTypeDeflate PipelineUpdateParamsDestinationCompressionType = "deflate"
)

func (r PipelineUpdateParamsDestinationCompressionType) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsDestinationCompressionTypeNone, PipelineUpdateParamsDestinationCompressionTypeGzip, PipelineUpdateParamsDestinationCompressionTypeDeflate:
		return true
	}
	return false
}

// Specifies the format of data to deliver.
type PipelineUpdateParamsDestinationFormat string

const (
	PipelineUpdateParamsDestinationFormatJson PipelineUpdateParamsDestinationFormat = "json"
)

func (r PipelineUpdateParamsDestinationFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsDestinationFormatJson:
		return true
	}
	return false
}

type PipelineUpdateParamsDestinationPath struct {
	// Specifies the R2 Bucket to store files.
	Bucket param.Field[string] `json:"bucket,required"`
	// Specifies the name pattern to for individual data files.
	Filename param.Field[string] `json:"filename"`
	// Specifies the name pattern for directory.
	Filepath param.Field[string] `json:"filepath"`
	// Specifies the base directory within the bucket.
	Prefix param.Field[string] `json:"prefix"`
}

func (r PipelineUpdateParamsDestinationPath) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of destination.
type PipelineUpdateParamsDestinationType string

const (
	PipelineUpdateParamsDestinationTypeR2 PipelineUpdateParamsDestinationType = "r2"
)

func (r PipelineUpdateParamsDestinationType) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsDestinationTypeR2:
		return true
	}
	return false
}

type PipelineUpdateParamsDestinationCredentials struct {
	// Specifies the R2 Bucket Access Key Id.
	AccessKeyID param.Field[string] `json:"access_key_id,required"`
	// Specifies the R2 Endpoint.
	Endpoint param.Field[string] `json:"endpoint,required"`
	// Specifies the R2 Bucket Secret Access Key.
	SecretAccessKey param.Field[string] `json:"secret_access_key,required"`
}

func (r PipelineUpdateParamsDestinationCredentials) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineUpdateParamsSourceFormat] `json:"format,required"`
	Type   param.Field[string]                           `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication param.Field[bool]        `json:"authentication"`
	CORS           param.Field[interface{}] `json:"cors"`
}

func (r PipelineUpdateParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineUpdateParamsSource) implementsPipelineUpdateParamsSourceUnion() {}

// Satisfied by
// [pipelines.PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource],
// [pipelines.PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSource],
// [PipelineUpdateParamsSource].
type PipelineUpdateParamsSourceUnion interface {
	implementsPipelineUpdateParamsSourceUnion()
}

type PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                     `json:"type,required"`
	// Specifies authentication is required to send to this Pipeline.
	Authentication param.Field[bool]                                                                     `json:"authentication"`
	CORS           param.Field[PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS] `json:"cors"`
}

func (r PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSource) implementsPipelineUpdateParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat string

const (
	PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat = "json"
)

func (r PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceFormatJson:
		return true
	}
	return false
}

type PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS struct {
	// Specifies allowed origins to allow Cross Origin HTTP Requests.
	Origins param.Field[[]string] `json:"origins"`
}

func (r PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesHTTPSourceCORS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSource struct {
	// Specifies the format of source data.
	Format param.Field[PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat] `json:"format,required"`
	Type   param.Field[string]                                                                        `json:"type,required"`
}

func (r PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSource) implementsPipelineUpdateParamsSourceUnion() {
}

// Specifies the format of source data.
type PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat string

const (
	PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat = "json"
)

func (r PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsSourceWorkersPipelinesWorkersPipelinesBindingSourceFormatJson:
		return true
	}
	return false
}

// Specifies the format of source data.
type PipelineUpdateParamsSourceFormat string

const (
	PipelineUpdateParamsSourceFormatJson PipelineUpdateParamsSourceFormat = "json"
)

func (r PipelineUpdateParamsSourceFormat) IsKnown() bool {
	switch r {
	case PipelineUpdateParamsSourceFormatJson:
		return true
	}
	return false
}

type PipelineUpdateResponseEnvelope struct {
	// Describes the configuration of a Pipeline.
	Result PipelineUpdateResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                               `json:"success,required"`
	JSON    pipelineUpdateResponseEnvelopeJSON `json:"-"`
}

// pipelineUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PipelineUpdateResponseEnvelope]
type pipelineUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PipelineListParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies which page to retrieve.
	Page param.Field[string] `query:"page"`
	// Specifies the number of Pipelines per page.
	PerPage param.Field[string] `query:"per_page"`
	// Specifies the prefix of Pipeline name to search.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [PipelineListParams]'s query parameters as `url.Values`.
func (r PipelineListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PipelineDeleteParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PipelineGetParams struct {
	// Specifies the public ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PipelineGetResponseEnvelope struct {
	// Describes the configuration of a Pipeline.
	Result PipelineGetResponse `json:"result,required"`
	// Indicates whether the API call was successful.
	Success bool                            `json:"success,required"`
	JSON    pipelineGetResponseEnvelopeJSON `json:"-"`
}

// pipelineGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PipelineGetResponseEnvelope]
type pipelineGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PipelineGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pipelineGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
