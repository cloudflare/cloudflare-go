// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

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
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// SuperSlurperJobService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperSlurperJobService] method instead.
type SuperSlurperJobService struct {
	Options []option.RequestOption
	Logs    *SuperSlurperJobLogService
}

// NewSuperSlurperJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuperSlurperJobService(opts ...option.RequestOption) (r *SuperSlurperJobService) {
	r = &SuperSlurperJobService{}
	r.Options = opts
	r.Logs = NewSuperSlurperJobLogService(opts...)
	return
}

// Create a job
func (r *SuperSlurperJobService) New(ctx context.Context, params SuperSlurperJobNewParams, opts ...option.RequestOption) (res *SuperSlurperJobNewResponse, err error) {
	var env SuperSlurperJobNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List jobs
func (r *SuperSlurperJobService) List(ctx context.Context, params SuperSlurperJobListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SuperSlurperJobListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs", params.AccountID)
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

// List jobs
func (r *SuperSlurperJobService) ListAutoPaging(ctx context.Context, params SuperSlurperJobListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SuperSlurperJobListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Abort all jobs
func (r *SuperSlurperJobService) AbortAll(ctx context.Context, body SuperSlurperJobAbortAllParams, opts ...option.RequestOption) (res *string, err error) {
	var env SuperSlurperJobAbortAllResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/slurper/jobs/abortAll", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SuperSlurperJobNewResponse struct {
	ID   string                         `json:"id"`
	JSON superSlurperJobNewResponseJSON `json:"-"`
}

// superSlurperJobNewResponseJSON contains the JSON metadata for the struct
// [SuperSlurperJobNewResponse]
type superSlurperJobNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobNewResponseJSON) RawJSON() string {
	return r.raw
}

type SuperSlurperJobListResponse struct {
	ID         string                            `json:"id"`
	CreatedAt  string                            `json:"createdAt"`
	FinishedAt string                            `json:"finishedAt,nullable"`
	Overwrite  bool                              `json:"overwrite"`
	Source     SuperSlurperJobListResponseSource `json:"source"`
	Status     SuperSlurperJobListResponseStatus `json:"status"`
	Target     SuperSlurperJobListResponseTarget `json:"target"`
	JSON       superSlurperJobListResponseJSON   `json:"-"`
}

// superSlurperJobListResponseJSON contains the JSON metadata for the struct
// [SuperSlurperJobListResponse]
type superSlurperJobListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	FinishedAt  apijson.Field
	Overwrite   apijson.Field
	Source      apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobListResponseJSON) RawJSON() string {
	return r.raw
}

type SuperSlurperJobListResponseSource struct {
	Bucket       string                                        `json:"bucket"`
	Endpoint     string                                        `json:"endpoint,nullable"`
	Jurisdiction SuperSlurperJobListResponseSourceJurisdiction `json:"jurisdiction"`
	PathPrefix   string                                        `json:"pathPrefix,nullable"`
	Vendor       SuperSlurperJobListResponseSourceVendor       `json:"vendor"`
	JSON         superSlurperJobListResponseSourceJSON         `json:"-"`
	union        SuperSlurperJobListResponseSourceUnion
}

// superSlurperJobListResponseSourceJSON contains the JSON metadata for the struct
// [SuperSlurperJobListResponseSource]
type superSlurperJobListResponseSourceJSON struct {
	Bucket       apijson.Field
	Endpoint     apijson.Field
	Jurisdiction apijson.Field
	PathPrefix   apijson.Field
	Vendor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r superSlurperJobListResponseSourceJSON) RawJSON() string {
	return r.raw
}

func (r *SuperSlurperJobListResponseSource) UnmarshalJSON(data []byte) (err error) {
	*r = SuperSlurperJobListResponseSource{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SuperSlurperJobListResponseSourceUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [r2.SuperSlurperJobListResponseSourceS3SourceResponseSchema],
// [r2.SuperSlurperJobListResponseSourceGcsSourceResponseSchema],
// [r2.SuperSlurperJobListResponseSourceR2SourceResponseSchema].
func (r SuperSlurperJobListResponseSource) AsUnion() SuperSlurperJobListResponseSourceUnion {
	return r.union
}

// Union satisfied by [r2.SuperSlurperJobListResponseSourceS3SourceResponseSchema],
// [r2.SuperSlurperJobListResponseSourceGcsSourceResponseSchema] or
// [r2.SuperSlurperJobListResponseSourceR2SourceResponseSchema].
type SuperSlurperJobListResponseSourceUnion interface {
	implementsSuperSlurperJobListResponseSource()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SuperSlurperJobListResponseSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperSlurperJobListResponseSourceS3SourceResponseSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperSlurperJobListResponseSourceGcsSourceResponseSchema{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SuperSlurperJobListResponseSourceR2SourceResponseSchema{}),
		},
	)
}

type SuperSlurperJobListResponseSourceS3SourceResponseSchema struct {
	Bucket     string                                                        `json:"bucket"`
	Endpoint   string                                                        `json:"endpoint,nullable"`
	PathPrefix string                                                        `json:"pathPrefix,nullable"`
	Vendor     SuperSlurperJobListResponseSourceS3SourceResponseSchemaVendor `json:"vendor"`
	JSON       superSlurperJobListResponseSourceS3SourceResponseSchemaJSON   `json:"-"`
}

// superSlurperJobListResponseSourceS3SourceResponseSchemaJSON contains the JSON
// metadata for the struct
// [SuperSlurperJobListResponseSourceS3SourceResponseSchema]
type superSlurperJobListResponseSourceS3SourceResponseSchemaJSON struct {
	Bucket      apijson.Field
	Endpoint    apijson.Field
	PathPrefix  apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperJobListResponseSourceS3SourceResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobListResponseSourceS3SourceResponseSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuperSlurperJobListResponseSourceS3SourceResponseSchema) implementsSuperSlurperJobListResponseSource() {
}

type SuperSlurperJobListResponseSourceS3SourceResponseSchemaVendor string

const (
	SuperSlurperJobListResponseSourceS3SourceResponseSchemaVendorS3 SuperSlurperJobListResponseSourceS3SourceResponseSchemaVendor = "s3"
)

func (r SuperSlurperJobListResponseSourceS3SourceResponseSchemaVendor) IsKnown() bool {
	switch r {
	case SuperSlurperJobListResponseSourceS3SourceResponseSchemaVendorS3:
		return true
	}
	return false
}

type SuperSlurperJobListResponseSourceGcsSourceResponseSchema struct {
	Bucket     string                                                         `json:"bucket"`
	PathPrefix string                                                         `json:"pathPrefix,nullable"`
	Vendor     SuperSlurperJobListResponseSourceGcsSourceResponseSchemaVendor `json:"vendor"`
	JSON       superSlurperJobListResponseSourceGcsSourceResponseSchemaJSON   `json:"-"`
}

// superSlurperJobListResponseSourceGcsSourceResponseSchemaJSON contains the JSON
// metadata for the struct
// [SuperSlurperJobListResponseSourceGcsSourceResponseSchema]
type superSlurperJobListResponseSourceGcsSourceResponseSchemaJSON struct {
	Bucket      apijson.Field
	PathPrefix  apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperJobListResponseSourceGcsSourceResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobListResponseSourceGcsSourceResponseSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuperSlurperJobListResponseSourceGcsSourceResponseSchema) implementsSuperSlurperJobListResponseSource() {
}

type SuperSlurperJobListResponseSourceGcsSourceResponseSchemaVendor string

const (
	SuperSlurperJobListResponseSourceGcsSourceResponseSchemaVendorGcs SuperSlurperJobListResponseSourceGcsSourceResponseSchemaVendor = "gcs"
)

func (r SuperSlurperJobListResponseSourceGcsSourceResponseSchemaVendor) IsKnown() bool {
	switch r {
	case SuperSlurperJobListResponseSourceGcsSourceResponseSchemaVendorGcs:
		return true
	}
	return false
}

type SuperSlurperJobListResponseSourceR2SourceResponseSchema struct {
	Bucket       string                                                              `json:"bucket"`
	Jurisdiction SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdiction `json:"jurisdiction"`
	PathPrefix   string                                                              `json:"pathPrefix,nullable"`
	Vendor       Provider                                                            `json:"vendor"`
	JSON         superSlurperJobListResponseSourceR2SourceResponseSchemaJSON         `json:"-"`
}

// superSlurperJobListResponseSourceR2SourceResponseSchemaJSON contains the JSON
// metadata for the struct
// [SuperSlurperJobListResponseSourceR2SourceResponseSchema]
type superSlurperJobListResponseSourceR2SourceResponseSchemaJSON struct {
	Bucket       apijson.Field
	Jurisdiction apijson.Field
	PathPrefix   apijson.Field
	Vendor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SuperSlurperJobListResponseSourceR2SourceResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobListResponseSourceR2SourceResponseSchemaJSON) RawJSON() string {
	return r.raw
}

func (r SuperSlurperJobListResponseSourceR2SourceResponseSchema) implementsSuperSlurperJobListResponseSource() {
}

type SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdiction string

const (
	SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdictionDefault SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdiction = "default"
	SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdictionEu      SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdiction = "eu"
	SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdictionFedramp SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdiction = "fedramp"
)

func (r SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdictionDefault, SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdictionEu, SuperSlurperJobListResponseSourceR2SourceResponseSchemaJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperJobListResponseSourceJurisdiction string

const (
	SuperSlurperJobListResponseSourceJurisdictionDefault SuperSlurperJobListResponseSourceJurisdiction = "default"
	SuperSlurperJobListResponseSourceJurisdictionEu      SuperSlurperJobListResponseSourceJurisdiction = "eu"
	SuperSlurperJobListResponseSourceJurisdictionFedramp SuperSlurperJobListResponseSourceJurisdiction = "fedramp"
)

func (r SuperSlurperJobListResponseSourceJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperJobListResponseSourceJurisdictionDefault, SuperSlurperJobListResponseSourceJurisdictionEu, SuperSlurperJobListResponseSourceJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperJobListResponseSourceVendor string

const (
	SuperSlurperJobListResponseSourceVendorS3  SuperSlurperJobListResponseSourceVendor = "s3"
	SuperSlurperJobListResponseSourceVendorGcs SuperSlurperJobListResponseSourceVendor = "gcs"
	SuperSlurperJobListResponseSourceVendorR2  SuperSlurperJobListResponseSourceVendor = "r2"
)

func (r SuperSlurperJobListResponseSourceVendor) IsKnown() bool {
	switch r {
	case SuperSlurperJobListResponseSourceVendorS3, SuperSlurperJobListResponseSourceVendorGcs, SuperSlurperJobListResponseSourceVendorR2:
		return true
	}
	return false
}

type SuperSlurperJobListResponseStatus string

const (
	SuperSlurperJobListResponseStatusRunning   SuperSlurperJobListResponseStatus = "running"
	SuperSlurperJobListResponseStatusPaused    SuperSlurperJobListResponseStatus = "paused"
	SuperSlurperJobListResponseStatusAborted   SuperSlurperJobListResponseStatus = "aborted"
	SuperSlurperJobListResponseStatusCompleted SuperSlurperJobListResponseStatus = "completed"
)

func (r SuperSlurperJobListResponseStatus) IsKnown() bool {
	switch r {
	case SuperSlurperJobListResponseStatusRunning, SuperSlurperJobListResponseStatusPaused, SuperSlurperJobListResponseStatusAborted, SuperSlurperJobListResponseStatusCompleted:
		return true
	}
	return false
}

type SuperSlurperJobListResponseTarget struct {
	Bucket       string                                        `json:"bucket"`
	Jurisdiction SuperSlurperJobListResponseTargetJurisdiction `json:"jurisdiction"`
	Vendor       Provider                                      `json:"vendor"`
	JSON         superSlurperJobListResponseTargetJSON         `json:"-"`
}

// superSlurperJobListResponseTargetJSON contains the JSON metadata for the struct
// [SuperSlurperJobListResponseTarget]
type superSlurperJobListResponseTargetJSON struct {
	Bucket       apijson.Field
	Jurisdiction apijson.Field
	Vendor       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SuperSlurperJobListResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobListResponseTargetJSON) RawJSON() string {
	return r.raw
}

type SuperSlurperJobListResponseTargetJurisdiction string

const (
	SuperSlurperJobListResponseTargetJurisdictionDefault SuperSlurperJobListResponseTargetJurisdiction = "default"
	SuperSlurperJobListResponseTargetJurisdictionEu      SuperSlurperJobListResponseTargetJurisdiction = "eu"
	SuperSlurperJobListResponseTargetJurisdictionFedramp SuperSlurperJobListResponseTargetJurisdiction = "fedramp"
)

func (r SuperSlurperJobListResponseTargetJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperJobListResponseTargetJurisdictionDefault, SuperSlurperJobListResponseTargetJurisdictionEu, SuperSlurperJobListResponseTargetJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperJobNewParams struct {
	AccountID param.Field[string]                              `path:"account_id,required"`
	Overwrite param.Field[bool]                                `json:"overwrite"`
	Source    param.Field[SuperSlurperJobNewParamsSourceUnion] `json:"source"`
	Target    param.Field[SuperSlurperJobNewParamsTarget]      `json:"target"`
}

func (r SuperSlurperJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperJobNewParamsSource struct {
	Bucket       param.Field[string]                                     `json:"bucket"`
	Endpoint     param.Field[string]                                     `json:"endpoint"`
	Jurisdiction param.Field[SuperSlurperJobNewParamsSourceJurisdiction] `json:"jurisdiction"`
	Secret       param.Field[interface{}]                                `json:"secret"`
	Vendor       param.Field[SuperSlurperJobNewParamsSourceVendor]       `json:"vendor"`
}

func (r SuperSlurperJobNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperJobNewParamsSource) implementsSuperSlurperJobNewParamsSourceUnion() {}

// Satisfied by [r2.SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchema],
// [r2.SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchema],
// [r2.SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchema],
// [SuperSlurperJobNewParamsSource].
type SuperSlurperJobNewParamsSourceUnion interface {
	implementsSuperSlurperJobNewParamsSourceUnion()
}

type SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchema struct {
	Bucket   param.Field[string]                                                      `json:"bucket"`
	Endpoint param.Field[string]                                                      `json:"endpoint"`
	Secret   param.Field[SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaSecret] `json:"secret"`
	Vendor   param.Field[SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaVendor] `json:"vendor"`
}

func (r SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchema) implementsSuperSlurperJobNewParamsSourceUnion() {
}

type SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaSecret struct {
	AccessKeyID     param.Field[string] `json:"accessKeyId"`
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaVendor string

const (
	SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaVendorS3 SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaVendor = "s3"
)

func (r SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaVendor) IsKnown() bool {
	switch r {
	case SuperSlurperJobNewParamsSourceR2SlurperS3SourceSchemaVendorS3:
		return true
	}
	return false
}

type SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchema struct {
	Bucket param.Field[string]                                                       `json:"bucket"`
	Secret param.Field[SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaSecret] `json:"secret"`
	Vendor param.Field[SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaVendor] `json:"vendor"`
}

func (r SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchema) implementsSuperSlurperJobNewParamsSourceUnion() {
}

type SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaSecret struct {
	ClientEmail param.Field[string] `json:"clientEmail"`
	PrivateKey  param.Field[string] `json:"privateKey"`
}

func (r SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaVendor string

const (
	SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaVendorGcs SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaVendor = "gcs"
)

func (r SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaVendor) IsKnown() bool {
	switch r {
	case SuperSlurperJobNewParamsSourceR2SlurperGcsSourceSchemaVendorGcs:
		return true
	}
	return false
}

type SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchema struct {
	Bucket       param.Field[string]                                                            `json:"bucket"`
	Jurisdiction param.Field[SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdiction] `json:"jurisdiction"`
	Secret       param.Field[SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaSecret]       `json:"secret"`
	Vendor       param.Field[Provider]                                                          `json:"vendor"`
}

func (r SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchema) implementsSuperSlurperJobNewParamsSourceUnion() {
}

type SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdiction string

const (
	SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdictionDefault SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdiction = "default"
	SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdictionEu      SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdiction = "eu"
	SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdictionFedramp SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdiction = "fedramp"
)

func (r SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdictionDefault, SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdictionEu, SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaSecret struct {
	AccessKeyID     param.Field[string] `json:"accessKeyId"`
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SuperSlurperJobNewParamsSourceR2SlurperR2SourceSchemaSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperJobNewParamsSourceJurisdiction string

const (
	SuperSlurperJobNewParamsSourceJurisdictionDefault SuperSlurperJobNewParamsSourceJurisdiction = "default"
	SuperSlurperJobNewParamsSourceJurisdictionEu      SuperSlurperJobNewParamsSourceJurisdiction = "eu"
	SuperSlurperJobNewParamsSourceJurisdictionFedramp SuperSlurperJobNewParamsSourceJurisdiction = "fedramp"
)

func (r SuperSlurperJobNewParamsSourceJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperJobNewParamsSourceJurisdictionDefault, SuperSlurperJobNewParamsSourceJurisdictionEu, SuperSlurperJobNewParamsSourceJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperJobNewParamsSourceVendor string

const (
	SuperSlurperJobNewParamsSourceVendorS3  SuperSlurperJobNewParamsSourceVendor = "s3"
	SuperSlurperJobNewParamsSourceVendorGcs SuperSlurperJobNewParamsSourceVendor = "gcs"
	SuperSlurperJobNewParamsSourceVendorR2  SuperSlurperJobNewParamsSourceVendor = "r2"
)

func (r SuperSlurperJobNewParamsSourceVendor) IsKnown() bool {
	switch r {
	case SuperSlurperJobNewParamsSourceVendorS3, SuperSlurperJobNewParamsSourceVendorGcs, SuperSlurperJobNewParamsSourceVendorR2:
		return true
	}
	return false
}

type SuperSlurperJobNewParamsTarget struct {
	Bucket       param.Field[string]                                     `json:"bucket"`
	Jurisdiction param.Field[SuperSlurperJobNewParamsTargetJurisdiction] `json:"jurisdiction"`
	Secret       param.Field[SuperSlurperJobNewParamsTargetSecret]       `json:"secret"`
	Vendor       param.Field[Provider]                                   `json:"vendor"`
}

func (r SuperSlurperJobNewParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperJobNewParamsTargetJurisdiction string

const (
	SuperSlurperJobNewParamsTargetJurisdictionDefault SuperSlurperJobNewParamsTargetJurisdiction = "default"
	SuperSlurperJobNewParamsTargetJurisdictionEu      SuperSlurperJobNewParamsTargetJurisdiction = "eu"
	SuperSlurperJobNewParamsTargetJurisdictionFedramp SuperSlurperJobNewParamsTargetJurisdiction = "fedramp"
)

func (r SuperSlurperJobNewParamsTargetJurisdiction) IsKnown() bool {
	switch r {
	case SuperSlurperJobNewParamsTargetJurisdictionDefault, SuperSlurperJobNewParamsTargetJurisdictionEu, SuperSlurperJobNewParamsTargetJurisdictionFedramp:
		return true
	}
	return false
}

type SuperSlurperJobNewParamsTargetSecret struct {
	AccessKeyID     param.Field[string] `json:"accessKeyId"`
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SuperSlurperJobNewParamsTargetSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SuperSlurperJobNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors"`
	Messages []string                   `json:"messages"`
	Result   SuperSlurperJobNewResponse `json:"result"`
	// Indicates if the API call was successful or not.
	Success SuperSlurperJobNewResponseEnvelopeSuccess `json:"success"`
	JSON    superSlurperJobNewResponseEnvelopeJSON    `json:"-"`
}

// superSlurperJobNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SuperSlurperJobNewResponseEnvelope]
type superSlurperJobNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperJobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SuperSlurperJobNewResponseEnvelopeSuccess bool

const (
	SuperSlurperJobNewResponseEnvelopeSuccessTrue SuperSlurperJobNewResponseEnvelopeSuccess = true
)

func (r SuperSlurperJobNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SuperSlurperJobNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SuperSlurperJobListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Limit     param.Field[int64]  `query:"limit"`
	Offset    param.Field[int64]  `query:"offset"`
}

// URLQuery serializes [SuperSlurperJobListParams]'s query parameters as
// `url.Values`.
func (r SuperSlurperJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SuperSlurperJobAbortAllParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type SuperSlurperJobAbortAllResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors"`
	Messages []string              `json:"messages"`
	Result   string                `json:"result"`
	// Indicates if the API call was successful or not.
	Success SuperSlurperJobAbortAllResponseEnvelopeSuccess `json:"success"`
	JSON    superSlurperJobAbortAllResponseEnvelopeJSON    `json:"-"`
}

// superSlurperJobAbortAllResponseEnvelopeJSON contains the JSON metadata for the
// struct [SuperSlurperJobAbortAllResponseEnvelope]
type superSlurperJobAbortAllResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SuperSlurperJobAbortAllResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r superSlurperJobAbortAllResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicates if the API call was successful or not.
type SuperSlurperJobAbortAllResponseEnvelopeSuccess bool

const (
	SuperSlurperJobAbortAllResponseEnvelopeSuccessTrue SuperSlurperJobAbortAllResponseEnvelopeSuccess = true
)

func (r SuperSlurperJobAbortAllResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SuperSlurperJobAbortAllResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
