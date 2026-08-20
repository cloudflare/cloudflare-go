// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// SinkholeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSinkholeService] method instead.
type SinkholeService struct {
	Options   []option.RequestOption
	Ingresses *SinkholeIngressService
}

// NewSinkholeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSinkholeService(opts ...option.RequestOption) (r *SinkholeService) {
	r = &SinkholeService{}
	r.Options = opts
	r.Ingresses = NewSinkholeIngressService(opts...)
	return
}

// Create a new sinkhole. Logs of large request bodies will be truncated, but the
// full request body can be recorded in R2. If you wish to record large request
// bodies in R2, include the R2 key ID, key secret, and bucket name in the request
// body.
func (r *SinkholeService) New(ctx context.Context, params SinkholeNewParams, opts ...option.RequestOption) (res *Sinkhole, err error) {
	var env SinkholeNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/intel/sinkholes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the name or R2 configuration of the specified sinkhole.
func (r *SinkholeService) Update(ctx context.Context, sinkholeID string, params SinkholeUpdateParams, opts ...option.RequestOption) (res *SinkholeUpdateResponse, err error) {
	var env SinkholeUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sinkholeID == "" {
		err = errors.New("missing required sinkhole_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/intel/sinkholes/%s", params.AccountID, sinkholeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists sinkholes owned by the account for redirecting malicious traffic.
func (r *SinkholeService) List(ctx context.Context, query SinkholeListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Sinkhole], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/intel/sinkholes", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists sinkholes owned by the account for redirecting malicious traffic.
func (r *SinkholeService) ListAutoPaging(ctx context.Context, query SinkholeListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Sinkhole] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete the specified sinkhole. The sinkhole must not have any active ingress
// rules defined. A 409 response code indicates that this condition is not met.
func (r *SinkholeService) Delete(ctx context.Context, sinkholeID string, body SinkholeDeleteParams, opts ...option.RequestOption) (res *SinkholeDeleteResponse, err error) {
	var env SinkholeDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sinkholeID == "" {
		err = errors.New("missing required sinkhole_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/intel/sinkholes/%s", body.AccountID, sinkholeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get the specified sinkhole by its unique identifier.
func (r *SinkholeService) Get(ctx context.Context, sinkholeID string, query SinkholeGetParams, opts ...option.RequestOption) (res *Sinkhole, err error) {
	var env SinkholeGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sinkholeID == "" {
		err = errors.New("missing required sinkhole_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/intel/sinkholes/%s", query.AccountID, sinkholeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Sinkhole struct {
	// The unique identifier for the sinkhole.
	ID string `json:"id"`
	// The account tag that owns this sinkhole.
	AccountTag string `json:"account_tag"`
	// The date and time when the sinkhole was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The date and time when the sinkhole was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the sinkhole.
	Name string `json:"name"`
	// The name of the R2 bucket to store results.
	R2Bucket string `json:"r2_bucket"`
	// The id of the R2 instance.
	R2ID string       `json:"r2_id"`
	JSON sinkholeJSON `json:"-"`
}

// sinkholeJSON contains the JSON metadata for the struct [Sinkhole]
type sinkholeJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	R2Bucket    apijson.Field
	R2ID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Sinkhole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeJSON) RawJSON() string {
	return r.raw
}

type SinkholeUpdateResponse = interface{}

type SinkholeDeleteResponse = interface{}

type SinkholeNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The name of the sinkhole.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the R2 bucket to store results. Required if you want to store large
	// request bodies in R2.
	R2Bucket param.Field[string] `json:"r2_bucket"`
	// The id of the R2 instance. Required if you want to store large request bodies in
	// R2.
	R2ID param.Field[string] `json:"r2_id"`
	// The secret key for the R2 API token. Required if you want to store large request
	// bodies in R2.
	R2Secret param.Field[string] `json:"r2_secret"`
}

func (r SinkholeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SinkholeNewResponseEnvelope struct {
	Errors   []SinkholeNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Sinkhole                           `json:"result"`
	JSON    sinkholeNewResponseEnvelopeJSON    `json:"-"`
}

// sinkholeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkholeNewResponseEnvelope]
type sinkholeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeNewResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           SinkholeNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SinkholeNewResponseEnvelopeErrors]
type sinkholeNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeNewResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    sinkholeNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SinkholeNewResponseEnvelopeErrorsSource]
type sinkholeNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeNewResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           SinkholeNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SinkholeNewResponseEnvelopeMessages]
type sinkholeNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeNewResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    sinkholeNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SinkholeNewResponseEnvelopeMessagesSource]
type sinkholeNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeNewResponseEnvelopeSuccess bool

const (
	SinkholeNewResponseEnvelopeSuccessTrue SinkholeNewResponseEnvelopeSuccess = true
)

func (r SinkholeNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SinkholeUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The name of the sinkhole.
	Name param.Field[string] `json:"name" api:"required"`
	// The name of the R2 bucket to store results. Required if you want to store large
	// request bodies in R2.
	R2Bucket param.Field[string] `json:"r2_bucket"`
	// The id of the R2 instance. Required if you want to store large request bodies in
	// R2.
	R2ID param.Field[string] `json:"r2_id"`
	// The secret key for the R2 API token. Required if you want to store large request
	// bodies in R2.
	R2Secret param.Field[string] `json:"r2_secret"`
}

func (r SinkholeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SinkholeUpdateResponseEnvelope struct {
	Errors   []SinkholeUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SinkholeUpdateResponse                `json:"result"`
	JSON    sinkholeUpdateResponseEnvelopeJSON    `json:"-"`
}

// sinkholeUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkholeUpdateResponseEnvelope]
type sinkholeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeUpdateResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SinkholeUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SinkholeUpdateResponseEnvelopeErrors]
type sinkholeUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    sinkholeUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SinkholeUpdateResponseEnvelopeErrorsSource]
type sinkholeUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeUpdateResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SinkholeUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SinkholeUpdateResponseEnvelopeMessages]
type sinkholeUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    sinkholeUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SinkholeUpdateResponseEnvelopeMessagesSource]
type sinkholeUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeUpdateResponseEnvelopeSuccess bool

const (
	SinkholeUpdateResponseEnvelopeSuccessTrue SinkholeUpdateResponseEnvelopeSuccess = true
)

func (r SinkholeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SinkholeListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SinkholeDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SinkholeDeleteResponseEnvelope struct {
	Errors   []SinkholeDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SinkholeDeleteResponse                `json:"result"`
	JSON    sinkholeDeleteResponseEnvelopeJSON    `json:"-"`
}

// sinkholeDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkholeDeleteResponseEnvelope]
type sinkholeDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeDeleteResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SinkholeDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SinkholeDeleteResponseEnvelopeErrors]
type sinkholeDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    sinkholeDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SinkholeDeleteResponseEnvelopeErrorsSource]
type sinkholeDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeDeleteResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SinkholeDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SinkholeDeleteResponseEnvelopeMessages]
type sinkholeDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    sinkholeDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SinkholeDeleteResponseEnvelopeMessagesSource]
type sinkholeDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeDeleteResponseEnvelopeSuccess bool

const (
	SinkholeDeleteResponseEnvelopeSuccessTrue SinkholeDeleteResponseEnvelopeSuccess = true
)

func (r SinkholeDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SinkholeGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SinkholeGetResponseEnvelope struct {
	Errors   []SinkholeGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  Sinkhole                           `json:"result"`
	JSON    sinkholeGetResponseEnvelopeJSON    `json:"-"`
}

// sinkholeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkholeGetResponseEnvelope]
type sinkholeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeGetResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           SinkholeGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SinkholeGetResponseEnvelopeErrors]
type sinkholeGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeGetResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    sinkholeGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SinkholeGetResponseEnvelopeErrorsSource]
type sinkholeGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeGetResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           SinkholeGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SinkholeGetResponseEnvelopeMessages]
type sinkholeGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeGetResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    sinkholeGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SinkholeGetResponseEnvelopeMessagesSource]
type sinkholeGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeGetResponseEnvelopeSuccess bool

const (
	SinkholeGetResponseEnvelopeSuccessTrue SinkholeGetResponseEnvelopeSuccess = true
)

func (r SinkholeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
