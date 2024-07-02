// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SippyService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSippyService] method instead.
type SippyService struct {
	Options []option.RequestOption
}

// NewSippyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSippyService(opts ...option.RequestOption) (r *SippyService) {
	r = &SippyService{}
	r.Options = opts
	return
}

// Sets configuration for Sippy for an existing R2 bucket.
func (r *SippyService) Update(ctx context.Context, bucketName string, params SippyUpdateParams, opts ...option.RequestOption) (res *Sippy, err error) {
	var env SippyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables Sippy on this bucket
func (r *SippyService) Delete(ctx context.Context, bucketName string, body SippyDeleteParams, opts ...option.RequestOption) (res *SippyDeleteResponse, err error) {
	var env SippyDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets configuration for Sippy for an existing R2 bucket.
func (r *SippyService) Get(ctx context.Context, bucketName string, query SippyGetParams, opts ...option.RequestOption) (res *Sippy, err error) {
	var env SippyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/sippy", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Provider string

const (
	ProviderR2 Provider = "r2"
)

func (r Provider) IsKnown() bool {
	switch r {
	case ProviderR2:
		return true
	}
	return false
}

type Sippy struct {
	// Details about the configured destination bucket
	Destination SippyDestination `json:"destination"`
	// State of Sippy for this bucket
	Enabled bool `json:"enabled"`
	// Details about the configured source bucket
	Source SippySource `json:"source"`
	JSON   sippyJSON   `json:"-"`
}

// sippyJSON contains the JSON metadata for the struct [Sippy]
type sippyJSON struct {
	Destination apijson.Field
	Enabled     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Sippy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyJSON) RawJSON() string {
	return r.raw
}

// Details about the configured destination bucket
type SippyDestination struct {
	// ID of the Cloudflare API token used when writing objects to this bucket
	AccessKeyID string `json:"accessKeyId"`
	Account     string `json:"account"`
	// Name of the bucket on the provider
	Bucket   string               `json:"bucket"`
	Provider Provider             `json:"provider"`
	JSON     sippyDestinationJSON `json:"-"`
}

// sippyDestinationJSON contains the JSON metadata for the struct
// [SippyDestination]
type sippyDestinationJSON struct {
	AccessKeyID apijson.Field
	Account     apijson.Field
	Bucket      apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyDestination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyDestinationJSON) RawJSON() string {
	return r.raw
}

// Details about the configured source bucket
type SippySource struct {
	// Name of the bucket on the provider
	Bucket   string              `json:"bucket"`
	Provider SippySourceProvider `json:"provider"`
	// Region where the bucket resides (AWS only)
	Region string          `json:"region,nullable"`
	JSON   sippySourceJSON `json:"-"`
}

// sippySourceJSON contains the JSON metadata for the struct [SippySource]
type sippySourceJSON struct {
	Bucket      apijson.Field
	Provider    apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippySource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippySourceJSON) RawJSON() string {
	return r.raw
}

type SippySourceProvider string

const (
	SippySourceProviderAws SippySourceProvider = "aws"
	SippySourceProviderGcs SippySourceProvider = "gcs"
)

func (r SippySourceProvider) IsKnown() bool {
	switch r {
	case SippySourceProviderAws, SippySourceProviderGcs:
		return true
	}
	return false
}

type SippyDeleteResponse struct {
	Enabled SippyDeleteResponseEnabled `json:"enabled"`
	JSON    sippyDeleteResponseJSON    `json:"-"`
}

// sippyDeleteResponseJSON contains the JSON metadata for the struct
// [SippyDeleteResponse]
type sippyDeleteResponseJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SippyDeleteResponseEnabled bool

const (
	SippyDeleteResponseEnabledFalse SippyDeleteResponseEnabled = false
)

func (r SippyDeleteResponseEnabled) IsKnown() bool {
	switch r {
	case SippyDeleteResponseEnabledFalse:
		return true
	}
	return false
}

type SippyUpdateParams struct {
	// Account ID
	AccountID param.Field[string]        `path:"account_id,required"`
	Body      SippyUpdateParamsBodyUnion `json:"body,required"`
}

func (r SippyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SippyUpdateParamsBody struct {
	Destination param.Field[interface{}] `json:"destination,required"`
	Source      param.Field[interface{}] `json:"source,required"`
}

func (r SippyUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SippyUpdateParamsBody) implementsR2SippyUpdateParamsBodyUnion() {}

// Satisfied by [r2.SippyUpdateParamsBodyR2EnableSippyAws],
// [r2.SippyUpdateParamsBodyR2EnableSippyGcs], [SippyUpdateParamsBody].
type SippyUpdateParamsBodyUnion interface {
	implementsR2SippyUpdateParamsBodyUnion()
}

type SippyUpdateParamsBodyR2EnableSippyAws struct {
	// R2 bucket to copy objects to
	Destination param.Field[SippyUpdateParamsBodyR2EnableSippyAwsDestination] `json:"destination"`
	// AWS S3 bucket to copy objects from
	Source param.Field[SippyUpdateParamsBodyR2EnableSippyAwsSource] `json:"source"`
}

func (r SippyUpdateParamsBodyR2EnableSippyAws) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SippyUpdateParamsBodyR2EnableSippyAws) implementsR2SippyUpdateParamsBodyUnion() {}

// R2 bucket to copy objects to
type SippyUpdateParamsBodyR2EnableSippyAwsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]   `json:"accessKeyId"`
	Provider    param.Field[Provider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsBodyR2EnableSippyAwsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AWS S3 bucket to copy objects from
type SippyUpdateParamsBodyR2EnableSippyAwsSource struct {
	// Access Key ID of an IAM credential (ideally scoped to a single S3 bucket)
	AccessKeyID param.Field[string] `json:"accessKeyId"`
	// Name of the AWS S3 bucket
	Bucket   param.Field[string]                                              `json:"bucket"`
	Provider param.Field[SippyUpdateParamsBodyR2EnableSippyAwsSourceProvider] `json:"provider"`
	// Name of the AWS availability zone
	Region param.Field[string] `json:"region"`
	// Secret Access Key of an IAM credential (ideally scoped to a single S3 bucket)
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsBodyR2EnableSippyAwsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SippyUpdateParamsBodyR2EnableSippyAwsSourceProvider string

const (
	SippyUpdateParamsBodyR2EnableSippyAwsSourceProviderAws SippyUpdateParamsBodyR2EnableSippyAwsSourceProvider = "aws"
)

func (r SippyUpdateParamsBodyR2EnableSippyAwsSourceProvider) IsKnown() bool {
	switch r {
	case SippyUpdateParamsBodyR2EnableSippyAwsSourceProviderAws:
		return true
	}
	return false
}

type SippyUpdateParamsBodyR2EnableSippyGcs struct {
	// R2 bucket to copy objects to
	Destination param.Field[SippyUpdateParamsBodyR2EnableSippyGcsDestination] `json:"destination"`
	// GCS bucket to copy objects from
	Source param.Field[SippyUpdateParamsBodyR2EnableSippyGcsSource] `json:"source"`
}

func (r SippyUpdateParamsBodyR2EnableSippyGcs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SippyUpdateParamsBodyR2EnableSippyGcs) implementsR2SippyUpdateParamsBodyUnion() {}

// R2 bucket to copy objects to
type SippyUpdateParamsBodyR2EnableSippyGcsDestination struct {
	// ID of a Cloudflare API token. This is the value labelled "Access Key ID" when
	// creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	AccessKeyID param.Field[string]   `json:"accessKeyId"`
	Provider    param.Field[Provider] `json:"provider"`
	// Value of a Cloudflare API token. This is the value labelled "Secret Access Key"
	// when creating an API token from the
	// [R2 dashboard](https://dash.cloudflare.com/?to=/:account/r2/api-tokens).
	//
	// Sippy will use this token when writing objects to R2, so it is best to scope
	// this token to the bucket you're enabling Sippy for.
	SecretAccessKey param.Field[string] `json:"secretAccessKey"`
}

func (r SippyUpdateParamsBodyR2EnableSippyGcsDestination) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GCS bucket to copy objects from
type SippyUpdateParamsBodyR2EnableSippyGcsSource struct {
	// Name of the GCS bucket
	Bucket param.Field[string] `json:"bucket"`
	// Client email of an IAM credential (ideally scoped to a single GCS bucket)
	ClientEmail param.Field[string] `json:"clientEmail"`
	// Private Key of an IAM credential (ideally scoped to a single GCS bucket)
	PrivateKey param.Field[string]                                              `json:"privateKey"`
	Provider   param.Field[SippyUpdateParamsBodyR2EnableSippyGcsSourceProvider] `json:"provider"`
}

func (r SippyUpdateParamsBodyR2EnableSippyGcsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SippyUpdateParamsBodyR2EnableSippyGcsSourceProvider string

const (
	SippyUpdateParamsBodyR2EnableSippyGcsSourceProviderGcs SippyUpdateParamsBodyR2EnableSippyGcsSourceProvider = "gcs"
)

func (r SippyUpdateParamsBodyR2EnableSippyGcsSourceProvider) IsKnown() bool {
	switch r {
	case SippyUpdateParamsBodyR2EnableSippyGcsSourceProviderGcs:
		return true
	}
	return false
}

type SippyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   Sippy                 `json:"result,required"`
	// Whether the API call was successful
	Success SippyUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    sippyUpdateResponseEnvelopeJSON    `json:"-"`
}

// sippyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SippyUpdateResponseEnvelope]
type sippyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SippyUpdateResponseEnvelopeSuccess bool

const (
	SippyUpdateResponseEnvelopeSuccessTrue SippyUpdateResponseEnvelopeSuccess = true
)

func (r SippyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SippyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SippyDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type SippyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   SippyDeleteResponse   `json:"result,required"`
	// Whether the API call was successful
	Success SippyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    sippyDeleteResponseEnvelopeJSON    `json:"-"`
}

// sippyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SippyDeleteResponseEnvelope]
type sippyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SippyDeleteResponseEnvelopeSuccess bool

const (
	SippyDeleteResponseEnvelopeSuccessTrue SippyDeleteResponseEnvelopeSuccess = true
)

func (r SippyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SippyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SippyGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type SippyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   Sippy                 `json:"result,required"`
	// Whether the API call was successful
	Success SippyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    sippyGetResponseEnvelopeJSON    `json:"-"`
}

// sippyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SippyGetResponseEnvelope]
type sippyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SippyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sippyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SippyGetResponseEnvelopeSuccess bool

const (
	SippyGetResponseEnvelopeSuccessTrue SippyGetResponseEnvelopeSuccess = true
)

func (r SippyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SippyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
